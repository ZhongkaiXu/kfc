package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupPath = "/sys/fs/cgroup"

func main() {
	// os.Args[0] is the path to the executable
	// If it is "/proc/self/exe", we are running in a container
	// If it is not, we are running in the host environment, it maybe main.go or a binary file
	fmt.Printf("os.Args[0] = %s\n", os.Args[0])
	if os.Args[0] == "/proc/self/exe" {
		// so we are running in a container, the child process
		fmt.Printf("Running in a container with PID %d\n", os.Getpid())
		fmt.Println()
		// run a stress command to consume memory
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 20m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{
		}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	// so we are running in the host environment, the parent process
	// create a child process with namespaces
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNS,
	}
	
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// start the child process here
	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)
	}else{
		// the pid of the child process is cmd.Process.Pid
		// cause we run the child process by the cmd
		fmt.Printf("child pid %d\n", cmd.Process.Pid)
		cgroupName := "test-mem-limit"
		cgPath := path.Join(cgroupPath, cgroupName)

		if err := os.MkdirAll(cgPath, 0755); err != nil {
			log.Fatal(err)
		}

		// 0700 means only the owner can read, write and execute
		if err := os.WriteFile(path.Join(cgPath, "memory.max"), []byte("50M"),0700); err != nil {
			log.Fatalf("set memory.max failed: %v",err)
		}

		pid := strconv.Itoa(cmd.Process.Pid)
		if err := os.WriteFile(path.Join(cgPath, "cgroup.procs"), []byte(pid), 0700); err != nil {
			log.Fatalf("add pid %s to cgroup %s failed: %v", pid, cgroupName, err)
		}
	}

	cmd.Process.Wait()
}