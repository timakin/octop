package command

import (
	"github.com/timakin/op/client"

	"fmt"

	"github.com/codegangsta/cli"
)

func CmdIssue(c *cli.Context) {
	instance := client.New()
	issues := instance.GetIssues()
	for _, issue := range issues {
		fmt.Print(issue.Issue.Labels)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Event)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Issue.Title)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Issue.Number)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Issue.User)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Issue.UpdatedAt)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Issue.CreatedAt)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Issue.State)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("==========================")
	}
}
