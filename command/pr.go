package command

import (
	"github.com/timakin/octop/client"

	"encoding/json"
	"fmt"

	"github.com/codegangsta/cli"
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
	instance := client.New()
	pullreqs := instance.GetPullRequests("rails", "rails")
	for _, pullreq := range pullreqs {
		fmt.Print(*pullreq.Title)
		fmt.Print("\n")
		fmt.Print(*pullreq.State)
		fmt.Print("\n")
		fmt.Print(*pullreq.Body)
		fmt.Print("\n")
	}
}
