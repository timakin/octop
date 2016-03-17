package client

import (
	"fmt"
	"log"
	"sort"

	"github.com/google/go-github/github"
)

type NotificationOptions struct {
	repositoryName string
	mentioned      bool
}

type RepoNotificationCounter struct {
	Repository              *github.Repository
	UnreadNotificationCount int
}

type RepoNotificationCounters []RepoNotificationCounter

func (r RepoNotificationCounters) Len() int {
	return len(r)
}

func (r RepoNotificationCounters) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RepoNotificationCounters) Less(i, j int) bool {
	return r[i].UnreadNotificationCount < r[j].UnreadNotificationCount
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
	var sortRepoCandidate RepoNotificationCounters
	repos := GetListFollowingRepository()
	for _, repo := range repos {
		unreadCount := countUnreadRepositoryNotification(repo.Owner.Login, repo.Name)
		repoNotificationCounter := &RepoNotificationCounter{
			Repository:              &repo,
			UnreadNotificationCount: unreadCount,
		}
		sortRepoCandidate := append(sortRepoCandidate, *repoNotificationCounter)
	}
	sort.Sort(sortRepoCandidate)
	for _, v := range sortRepoCandidate {
		fmt.Print("======================")
		fmt.Print(v.UnreadNotificationCount)
		fmt.Print(*v.Repository.Owner.Name)
		fmt.Print(*v.Repository.Name)
		fmt.Print("======================")
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
