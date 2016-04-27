package command

import (
	"github.com/timakin/octop/client"

	"fmt"
	"log"
	"net/url"

	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/timakin/octop/repl"
)

func CmdPr(c *cli.Context) {
	baseUrl := ""
	if c.String("enterprise") != "" {
		gheHost := c.String("enterprise")
		baseUrl = "https://" + gheHost
	}

	i := client.New()
	if baseUrl != "" {
		remoteHost, err := url.Parse(baseUrl)
		if err != nil {
			e := errors.Wrap(err, "Specifield remote host cannot be parsed.")
			fmt.Print(e.Error())
		}
		i.SetRemoteHost(remoteHost)
	}
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
