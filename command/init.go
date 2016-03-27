package command

import (
	"fmt"
	"log"
	"reflect"

	"github.com/timakin/op/client"
	"github.com/timakin/op/repl"

	"github.com/codegangsta/cli"
	//"github.com/fatih/color"
)

type Tesss struct {
	name string
	yoyo string
}

func CmdInit(c *cli.Context) {
	i := client.New()
	repoNotificationCounters := i.GetRepoNotificationCounters()

	selected, err := repl.RepoSelectInterface(repoNotificationCounters)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(selected[0])
	fmt.Println(reflect.TypeOf(selected[0]))
	//issues := i.GetIssues(selected[0].Title,  selected[0].)

}
