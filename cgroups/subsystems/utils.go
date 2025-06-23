package subsystems

import (
	"fmt"
	"os"
	"path"
)

func FindCgroupMountPoint() string {
	return "/sys/fs/cgroup"
}

// Find the cgroup mount point
// Based on Cgroup v2 so we don't need to specify the subsystem
/*
	cgroupPath: name of the cgroup, e.g., "kfc-cgroup"
	autoCreate: whether to create the cgroup if it does not exist
*/
func GetCgroupPath(cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountPoint()
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err != nil {
				return "", fmt.Errorf("error creating cgroup path %s: %v", cgroupPath, err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("cgroup path error %v", err)
	}
}
