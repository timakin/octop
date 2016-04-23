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
		Name:   "n",
		Usage:  "octop n - notification tracking",
		Action: command.CmdNotification,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "i",
		Usage:  "octop i - issue tracking with selection of repo",
		Action: command.CmdIssue,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "p",
		Usage:  "octop p - pull-reqs tracking with selection of repo",
		Action: command.CmdPr,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
