package command

import (
	"github.com/timakin/op/client"

	"encoding/json"
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
)

func MapToStruct(m map[string]interface{}, val interface{}) error {
	tmp, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tmp, val)
	if err != nil {
		return err
	}
	return nil
}

func CmdPr(c *cli.Context) {
	pullreqs := client.GetPullRequests()
	for _, pullreq := range pullreqs {
		var pullreqPayload github.PullRequestEvent

		//		switch v := pullreq.Payload().(type) {
		//		case github.PullRequestEvent:
		//			fmt.Println(*v.PullRequest.Title)
		//		default:
		//			fmt.Println(v)
		//			fmt.Println("invalid payload!")
		//		}
		//fmt.Print(pullreq.Payload().(github.PullRequestEvent).PullRequest.Title)
		err := json.Unmarshal(*pullreq.RawPayload, &pullreqPayload)
		if err != nil {
			panic(err)
		}
		fmt.Print(*pullreqPayload.PullRequest.Title)

		fmt.Print("\n")
		fmt.Print("\n")
	}
}
