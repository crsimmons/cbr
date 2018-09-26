package main

import (
	"fmt"
	"os"

	"github.com/crsimmons/cbr/commands"
	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:     "cbr",
		Usage:    "A tool for backing up and restoring Concourse",
		Commands: commands.Commands,
		Action: func(c *cli.Context) error {
			fmt.Println("Hello World")
			return nil
		},
	}

	app.Run(os.Args)
}
