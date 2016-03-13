package command

import (
	"github.com/timakin/op/client"

	"fmt"

	"github.com/codegangsta/cli"
)

func CmdIssue(c *cli.Context) {
	issues := client.GetIssues()
	for _, issue := range issues {
		fmt.Print(issue.Issue.Labels)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Issue.Title)
		fmt.Print("\n")
		fmt.Print("\n")
	}
}
