package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/zhongkaixu/kfc/cgroups/subsystems"
	"github.com/zhongkaixu/kfc/container"
)

// run container
var runCommand = &cli.Command{
	Name:  "run",
	Usage: "Run a command in a new container with namespace and cgroup limits",
	// Flags means command line options, use '-tty' to enable TTY
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "tty",
			Usage: "enable TTY",
		},
		&cli.StringFlag{
			Name:  "m",
			Usage: "memory limit",
		},
		&cli.StringFlag{
			Name:  "cpushare",
			Usage: "cpushare limit",
		},
		&cli.StringFlag{
			Name:  "cpuset",
			Usage: "cpuset limit",
		},
	},
	// Action is the function that will be executed
	Action: func(ctx *cli.Context) error {
		log.Infof("Welcome to KFC")
		// if user foget to provide a command
		if ctx.Args().Len() == 0 {
			return fmt.Errorf("missing container command")
		}
		// cmdArray := ctx.Args().Slice()
		resConf := &subsystems.ResourceConfig{
			MemoryLimit: ctx.String("m"),
			CpuShare:    ctx.String("cpushare"),
			CpuSet:      ctx.String("cpuset"),
		}
		// if user need a tty
		tty := ctx.Bool("tty")
		Run(tty, ctx.Args().First(), resConf)
		return nil
	},
}

// init command is the first command that will be executed when the container starts
var initCommand = &cli.Command{
	Name:  "init",
	Usage: "Initialize the container environment",
	// get command line arguments, init the container environment
	Action: func(ctx *cli.Context) error {
		log.Infof("Starting container init process")
		cmd := ctx.Args().First()
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
