package command

import (
	"github.com/timakin/octop/client"

	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/timakin/octop/repl"
)

func CmdPr(c *cli.Context) {
	i := client.New()
	repoNotificationCounters := i.GetRepoNotificationCounters()

	selected, err := repl.RepoSelectInterface(repoNotificationCounters)
	if err != nil {
		log.Fatal(err)
	}

	pullreqs := i.GetPullRequests(selected[0].Owner, selected[0].Repo)

	for _, pullreq := range pullreqs {
		color.Cyan(*pullreq.Title)
		fmt.Print(*pullreq.User.Login + "\n")
		fmt.Print(*pullreq.HTMLURL + "\n")
		color.Cyan("----------------------------------------")
	}
}
