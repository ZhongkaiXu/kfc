package main

// log "xxx" renamed package import "logrus"
import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const usage = `Kaii's First Container, is a simple container implementation`

func main() {
	app := &cli.App{
		Name:  "kfc",
		Usage: usage,
		// functions that will be executed before and after the app runs
		Commands: []*cli.Command{
			runCommand,
			initCommand,
		},
	}

	app.Before = func(ctx *cli.Context) error {
		log.SetFormatter(&log.TextFormatter{
			ForceColors:   true, // colors
			FullTimestamp: true, // time-stamp
		})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
