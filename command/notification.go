package command

import (
	"github.com/timakin/octop/client"

	"fmt"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

func CmdNotification(c *cli.Context) {
	i := client.New()
	notifications := i.GetNotifications()
	for _, notification := range notifications {
		fmt.Print(*notification.Subject.Title)
		fmt.Print("\n")
		fmt.Print(*notification.Subject.URL)
		fmt.Print("\n")
		color.Cyan("-------------------------------------------------\n")
	}
}
