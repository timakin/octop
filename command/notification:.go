package command

import (
	"github.com/timakin/op/client"

	"fmt"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

func CmdNotification(c *cli.Context) {
	notifications := client.GetNotifications()
	for _, notification := range notifications {
		color.Cyan("> > > > > > > > > > > > > >")
		fmt.Print(*notification.Subject.Title)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*notification.Subject.URL)
		fmt.Print("\n")
		fmt.Print("\n")
	}
}
