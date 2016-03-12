package client

import (
	"fmt"

	"github.com/google/go-github/github"
)

type NotificationOptions struct {
	repositoryName string
	mentioned      bool
}

func GetNotifications(token string) {
	httpClient := Authenticate(token)
	ghCli := github.NewClient(httpClient)
	repos, _, err := ghCli.Repositories.List("", nil)
	fmt.Print(repos)
	fmt.Print(err)
}
