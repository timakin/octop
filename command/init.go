package command

import (
	"fmt"
	"log"

	"github.com/timakin/octop/client"
	"github.com/timakin/octop/repl"

	"github.com/codegangsta/cli"
	//"github.com/fatih/color"
)

func CmdInit(c *cli.Context) {
	i := client.New()
	repoNotificationCounters := i.GetRepoNotificationCounters()

	selected, err := repl.RepoSelectInterface(repoNotificationCounters)
	if err != nil {
		log.Fatal(err)
	}

	issues := i.GetIssues(selected[0].Owner, selected[0].Repo)
	fmt.Println(issues)
}
