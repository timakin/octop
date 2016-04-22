package command

import (
	"github.com/timakin/octop/client"

	"fmt"

	"github.com/codegangsta/cli"
)

func CmdIssue(c *cli.Context) {
	instance := client.New()
	issues := instance.GetIssues("rails", "rails")
	for _, issue := range issues {
		fmt.Print(issue)
		fmt.Print("\n")
		fmt.Print("==========================")
	}
}
