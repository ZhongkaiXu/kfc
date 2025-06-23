package subsystems

import (
	"fmt"
	"os"
	"path"
)

type CpuSubsystem struct {
}

func (s *CpuSubsystem) Name() string {
	return "cpu"
}

func (s *CpuSubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(cgroupPath, true); err == nil {
		if res.CpuShare != "" {
			if err := os.WriteFile(path.Join(subsysCgroupPath, "cpu.weight"), []byte(res.CpuShare), 0644); err != nil {
				return fmt.Errorf("Set cgroup cpu share failed %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}
