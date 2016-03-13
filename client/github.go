package client

import (
	"fmt"

	"github.com/google/go-github/github"
)

type NotificationOptions struct {
	repositoryName string
	mentioned      bool
}

func GetNotifications() []github.Notification {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	opt := &github.NotificationListOptions{All: true}
	notifications, _, err := ghCli.Activity.ListNotifications(opt)
	if err != nil {
		panic(err)
	}
	return notifications
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
