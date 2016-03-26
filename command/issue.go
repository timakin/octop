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
		fmt.Print(*issue.Title)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.Body)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print(*issue.UpdatedAt)
		fmt.Print("\n")
		fmt.Print("\n")
		fmt.Print("==========================")
	}
}
