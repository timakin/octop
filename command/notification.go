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
		color.Cyan("【" + *notification.Repository.Name + "】 : " + *notification.Subject.Title)
		fmt.Print(*notification.Subject.URL + "\n")
		color.Cyan("-------------------------------------------------")
	}
}
