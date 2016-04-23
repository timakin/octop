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
		fmt.Print(*pullreq.Title)
		fmt.Print("\n")
		fmt.Print(*pullreq.State)
		fmt.Print("\n")
		fmt.Print(*pullreq.Body)
		fmt.Print("\n")
		color.Cyan("----------------------------------------\n")
	}
}
