package command

import (
	"github.com/timakin/octop/client"

	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/timakin/octop/repl"
)

func CmdNotification(c *cli.Context) {
	i := client.New()
	repoNotificationCounters := i.GetRepoNotificationCounters()

	selected, err := repl.RepoSelectInterface(repoNotificationCounters)
	if err != nil {
		log.Fatal(err)
	}

	notifications := i.GetNotifications(selected[0].Owner, selected[0].Repo)
	for _, notification := range notifications {
		color.Cyan(notification.Title)
		fmt.Print(notification.URL + "\n")
		color.Cyan("-------------------------------------------------")
	}
}
