package main

import (
	"github.com/codegangsta/cli"

	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = "timakin"
	app.Email = "timaki.st@gmail.com"
	app.Usage = "octop [i/n/p]"

	app.Flags = GlobalFlags
	app.Commands = Commands
	app.CommandNotFound = CommandNotFound

	app.Run(os.Args)
}
