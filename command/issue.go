package command

import (
	"github.com/timakin/octop/client"

	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/timakin/octop/repl"
)

func CmdIssue(c *cli.Context) {
	i := client.New()
	repoNotificationCounters := i.GetRepoNotificationCounters()

	selected, err := repl.RepoSelectInterface(repoNotificationCounters)
	if err != nil {
		log.Fatal(err)
	}

	issues := i.GetIssues(selected[0].Owner, selected[0].Repo)

	for _, issue := range issues {
		color.Cyan(issue.Title)
		fmt.Print(issue.URL + "\n")
		color.Cyan("---------------------------------------")
	}
}
