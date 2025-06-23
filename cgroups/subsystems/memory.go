package subsystems

import (
	"fmt"
	"os"
	"path"
)

type MemorySubsystem struct {
}

func (s *MemorySubsystem) Name() string {
	return "memory"
}

func (s *MemorySubsystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(cgroupPath, true); err == nil {
		if res.MemoryLimit != "" {
			if err := os.WriteFile(path.Join(subsysCgroupPath, "memory.max"), []byte(res.MemoryLimit), 0644); err != nil {
				return fmt.Errorf("set cgroup memory failed %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}
