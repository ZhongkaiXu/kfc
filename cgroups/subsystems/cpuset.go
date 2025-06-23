package subsystems

import (
	"fmt"
	"os"
	"path"
)

type CpusetSubsystem struct {
}

func (s *CpusetSubsystem) Name() string {
	return "cpuset"
}

func (s *CpusetSubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(cgroupPath, true); err == nil {
		if res.CpuSet != "" {
			if err := os.WriteFile(path.Join(subsysCgroupPath, "cpuset.cpus"), []byte(res.CpuSet), 0644); err != nil {
				return fmt.Errorf("Set cgroup cpuset.cpus failed %v", err)
			} else {
				if err := os.WriteFile(path.Join(subsysCgroupPath, "cpuset.mems"), []byte(res.CpuSet), 0644); err != nil {
					return fmt.Errorf("Set cgroup cpuset.mems failed %v", err)
				}
			}
		}
		return nil
	} else {
		return err
	}
}
