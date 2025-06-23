package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/zhongkaixu/kfc/cgroups"
	"github.com/zhongkaixu/kfc/cgroups/subsystems"
	"github.com/zhongkaixu/kfc/container"
)

func Run(tty bool, command string, res *subsystems.ResourceConfig) {
	parent := container.NewParentProcess(tty, command)
	if parent == nil {
		log.Errorf("Fail to create parent process")
		return
	}
	if err := parent.Start(); err != nil {
		log.Error(err)
	}
	cgroupManager := cgroups.NewCgroupManager(fmt.Sprintf("kfc-cgroup-%d", parent.Process.Pid))
	defer cgroupManager.Destroy()
	if err := cgroupManager.Set(res); err != nil {
		log.Errorf("Fail to set cgroup: %v", err)
		return
	}
	cgroupManager.Apply(parent.Process.Pid)
	parent.Wait()
	cgroupManager.Destroy()
	os.Exit(0)
}
