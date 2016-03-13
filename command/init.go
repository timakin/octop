package command

import (
	"github.com/timakin/op/client"

	"fmt"

	"github.com/codegangsta/cli"
	//"github.com/fatih/color"
)

func CmdInit(c *cli.Context) {
	repos := client.GetListFollowingRepository()
	for _, repo := range repos {
		fmt.Print(*repo.Organization.Name + *repo.Name)
		fmt.Print("\n")
		fmt.Print("\n")
	}

}
