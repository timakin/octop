package command

import (
	"github.com/timakin/octop/client"
	"github.com/timakin/octop/constants"

	"fmt"
	"net/url"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/timakin/octop/repl"
)

func CmdIssue(c *cli.Context) {
	baseUrl := ""
	if c.String("enterprise") != "" {
		gheHost := c.String("enterprise")
		baseUrl = "https://" + gheHost
	}

	i := client.New()
	if baseUrl != "" {
		remoteHost, err := url.Parse(baseUrl)
		if err != nil {
			e := errors.Wrap(err, constants.ErrRemoteHostIsInvalid)
			fmt.Print(e.Error())
		}
		i.SetRemoteHost(remoteHost)
	}
	repoNotificationCounters := i.GetRepoNotificationCounters()

	selected, err := repl.RepoSelectInterface(repoNotificationCounters)
	if err != nil {
		e := errors.Wrap(err, constants.ErrRepoSelectFailed)
		fmt.Print(e.Error())
	}

	issues := i.GetIssues(selected[0].Owner, selected[0].Repo)

	for _, issue := range issues {
		color.Cyan(issue.Title)
		fmt.Print(issue.Owner + "\n")
		fmt.Print(issue.URL + "\n")
		color.Cyan("---------------------------------------")
	}
}
