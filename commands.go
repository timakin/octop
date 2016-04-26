package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/timakin/octop/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "notification",
		Usage:  "octop notification - notification tracking",
		Action: command.CmdNotification,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "enterprise, e",
				Usage: "Access enterprise remote host",
			},
		},
	},
	{
		Name:   "issue",
		Usage:  "octop issue - issue tracking with selection of repo",
		Action: command.CmdIssue,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "enterprise, e",
				Usage: "Access enterprise remote host",
			},
		},
	},
	{
		Name:   "pr",
		Usage:  "octop pr - pull-reqs tracking with selection of repo",
		Action: command.CmdPr,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "enterprise, e",
				Usage: "Access enterprise remote host",
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
