package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
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
	},
	// Action is the function that will be executed
	Action: func(ctx *cli.Context) error {
		// if user foget to provide a command
		if ctx.Args().Len() == 0 {
			return fmt.Errorf("missing container command")
		}

		// the command to run in the container
		cmd := ctx.Args().Get(0)
		// if user need a tty
		tty := ctx.Bool("tty")
		Run(tty, cmd)
		return nil
	},
}

// init command is the first command that will be executed when the container starts
var initCommand = &cli.Command{
	Name:  "init",
	Usage: "Initialize the container environment",
	// get command line arguments, init the container environment
	Action: func(ctx *cli.Context) error {
		log.Infof("init come on")
		cmd := ctx.Args().Get(0)
		log.Infof("init command: %s", cmd)
		err := container.RunContainerInitProcess(cmd, nil)
		return err
	},
}
