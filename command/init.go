package command

import (
	"github.com/timakin/op/client"

	"github.com/codegangsta/cli"
	//"github.com/fatih/color"
)

func CmdInit(c *cli.Context) {
	client.SelectRepository()
}
