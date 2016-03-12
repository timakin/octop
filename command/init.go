package command

import (
	"github.com/timakin/op/client"

	"bufio"
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	//"github.com/fatih/color"
)

func CmdInit(c *cli.Context) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter token: ")
	token, _ := reader.ReadString('\n')
	client.GetNotifications(token)
}
