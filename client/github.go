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

func GetListFollowingRepository() []github.Repository {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	opt := &github.ListOptions{PerPage: 50}
	userId := getAuthenticatedUserId()
	Repositories, _, err := ghCli.Activity.ListWatched(*userId, opt)
	if err != nil {
		panic(err)
	}
	return Repositories
}

// TODO: 引数でghCli引き回すのはアホなので、github cli 共通化 with interface
func getAuthenticatedUserId() *string {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	User, _, err := ghCli.Users.Get("")
	if err != nil {
		panic(err)
	}
	return User.Login
}
