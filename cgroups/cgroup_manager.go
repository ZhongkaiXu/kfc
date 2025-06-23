package cgroups

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/zhongkaixu/kfc/cgroups/subsystems"
)

type CgroupManager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	c.Resource = res
	for _, subSys := range subsystems.SubsystemIns {
		if err := subSys.Set(c.Path, res); err != nil {
			return err
		}
	}
	return nil
}

// `pid` is a integer, so we need to convert it to a string.
func (c *CgroupManager) Apply(pid int) error {
	if path, err := subsystems.GetCgroupPath(c.Path, true); err != nil {
		return err
	} else {
		os.WriteFile(path+"/cgroup.procs", []byte(strconv.Itoa(pid)), 0644)
	}
	return nil
}

func (c *CgroupManager) Destroy() error {
	if path, err := subsystems.GetCgroupPath(c.Path, false); err != nil {
		return err
	} else {
		if err := os.RemoveAll(path); err != nil {
			logrus.Warnf("remove cgroup fail %v", err)
		}
	}
	return nil
}
