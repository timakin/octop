package client

import (
	"fmt"
	"log"

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

func EventFilter(vs []github.Event, f func(github.Event) bool) []github.Event {
	vsf := make([]github.Event, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func NotificationFilter(vs []github.Notification, f func(github.Notification) bool) []github.Notification {
	vsf := make([]github.Notification, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func SortByNotificationCount(vs []github.Notification, f func(github.Notification) int) []github.Notification {
	vsf := make([]github.Notification, 0)
	for _, v := range vs {
		count := countUnreadRepositoryNotification(v.Repository.Owner.Name, v.Repository.Name)
		// countとrepositoryのmapを作る
	}
	// countをキーにしてソートする
}

func GetPullRequests() []github.Event {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	opt := &github.ListOptions{PerPage: 50}
	events, _, err := ghCli.Activity.ListRepositoryEvents("rails", "rails", opt)
	if err != nil {
		panic(err)
	}
	pullreqs := EventFilter(events, func(e github.Event) bool {
		return *e.Type == "PullRequestEvent"
	})
	return pullreqs
}

func SelectRepository() {
	repos := GetListFollowingRepository()
	fmt.Print(repos)
	for _, repo := range repos {
		nofiticationCount := countUnreadRepositoryNotification(repo.Owner.Login, repo.Name)
		fmt.Print(""*repo.Owner.Login + "/" + *repo.Name)
		fmt.Print("\n")
	}
}

func GetListFollowingRepository() []github.Repository {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)
	opt := &github.ListOptions{PerPage: 100}
	userId := getAuthenticatedUserId()
	Repositories, _, err := ghCli.Activity.ListWatched(*userId, opt)
	if err != nil {
		log.Fatal(err)
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

func countUnreadRepositoryNotification(owner *string, repoName *string) int {
	// TODO: API叩き過ぎるのでcache c.f. https://github.com/patrickmn/go-cache/blob/master/README.md
	notifications := GetNotifications()
	unreadRepositoryNotifications := NotificationFilter(notifications, func(n github.Notification) bool {
		return *n.Repository.Owner.Name == *owner && *n.Repository.Name == *repoName
	})
	return len(unreadRepositoryNotifications)
}
