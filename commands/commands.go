package commands

import (
	"fmt"

	cli "gopkg.in/urfave/cli.v2"
)

var Commands = []*cli.Command{
	{
		Name:    "backup",
		Aliases: []string{"b"},
		Usage:   "backup",
		Action: func(c *cli.Context) error {
			fmt.Println("backing up")
			return nil
		},
	},
}
