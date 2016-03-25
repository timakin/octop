package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/timakin/op/command"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "n",
		Usage:  "",
		Action: command.CmdNotification,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "init",
		Usage:  "",
		Action: command.CmdInit,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "i",
		Usage:  "",
		Action: command.CmdIssue,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "p",
		Usage:  "",
		Action: command.CmdPr,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
