package command

import (
	"fmt"
	"log"

	"github.com/timakin/op/client"
	"github.com/timakin/op/repl"

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
	selectedRes, err := repl.ResSelectInterface(issues)
	fmt.Println(selectedRes)
}
