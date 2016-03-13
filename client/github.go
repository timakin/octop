package client

import (
	"fmt"

	"github.com/google/go-github/github"
)

type NotificationOptions struct {
	repositoryName string
	mentioned      bool
}

func GetNotifications() {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	repos, _, err := ghCli.Repositories.List("", nil)
	fmt.Print(repos)
	fmt.Print(err)
}

func GetIssues() []github.IssueEvent {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	opt := &github.ListOptions{PerPage: 10}
	issueEvents, _, err := ghCli.Issues.ListRepositoryEvents("rails", "rails", opt)
	if err != nil {
		panic(err)
	}
	return issueEvents
}

func GetPullRequests() {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	pullreqs, _, err := ghCli.Repositories.List("", nil)
	fmt.Print(pullreqs)
	fmt.Print(err)
}
