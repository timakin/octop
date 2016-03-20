package client

import (
	"fmt"
	"log"
	"sort"

	"github.com/google/go-github/github"
	"github.com/patrickmn/go-cache"
)

type NotificationOptions struct {
	repositoryName string
	mentioned      bool
}

type RepoNotificationCounter struct {
	Repository              *github.Repository
	UnreadNotificationCount int
}

type RepoNotificationCounters []*RepoNotificationCounter

func (i Instance) GetNotifications() []github.Notification {
	if cv, found := i.cache.Get("notifications"); found {
		cachedNotifications := cv.([]github.Notification)
		return cachedNotifications
	}
	opt := &github.NotificationListOptions{All: true}
	notifications, _, err := i.ghCli.Activity.ListNotifications(opt)
	if err != nil {
		log.Fatal(err)
	}
	i.cache.Set("notifications", notifications, cache.DefaultExpiration)
	return notifications
}

func (i Instance) GetIssues() []github.IssueEvent {
	opt := &github.ListOptions{PerPage: 10}
	issueEvents, _, err := i.ghCli.Issues.ListRepositoryEvents("rails", "rails", opt)
	if err != nil {
		log.Fatal(err)
	}
	return issueEvents
}

func (i Instance) GetPullRequests() []github.Event {
	opt := &github.ListOptions{PerPage: 100}
	events, _, err := i.ghCli.Activity.ListRepositoryEvents("rails", "rails", opt)
	if err != nil {
		log.Fatal(err)
	}
	pullreqs := EventFilter(events, func(e github.Event) bool {
		return *e.Type == "PullRequestEvent"
	})
	return pullreqs
}

func (i Instance) SelectRepository() {
	repos := i.GetListFollowingRepository()
	sortRepoCandidate := make(RepoNotificationCounters, len(repos))
	for index, repo := range repos {
		repo := repo
		unreadCount := i.countUnreadRepositoryNotification(repo.Owner.Login, repo.Name)
		repoNotificationCounter := &RepoNotificationCounter{
			Repository:              &repo,
			UnreadNotificationCount: unreadCount,
		}
		sortRepoCandidate[index] = repoNotificationCounter
	}
	sort.Sort(sortRepoCandidate)
	for _, v := range sortRepoCandidate {
		fmt.Print("======================\n")
		fmt.Print(v.UnreadNotificationCount)
		fmt.Print("\n")
		fmt.Print(*v.Repository.Owner.Login)
		fmt.Print("\n")
		fmt.Print(*v.Repository.Name)
		fmt.Print("\n")
	}
}

func (i Instance) GetListFollowingRepository() []github.Repository {
	opt := &github.ListOptions{PerPage: 100}
	userId := i.getAuthenticatedUserId()
	Repositories, _, err := i.ghCli.Activity.ListWatched(*userId, opt)
	if err != nil {
		log.Fatal(err)
	}
	return Repositories
}

func (i Instance) getAuthenticatedUserId() *string {
	User, _, err := i.ghCli.Users.Get("")
	if err != nil {
		log.Fatal(err)
	}
	return User.Login
}

func (i Instance) countUnreadRepositoryNotification(owner *string, repoName *string) int {
	notifications := i.GetNotifications()
	unreadRepositoryNotifications := NotificationFilter(notifications, func(n github.Notification) bool {
		return *n.Repository.Owner.Login == *owner && *n.Repository.Name == *repoName
	})
	return len(unreadRepositoryNotifications)
}
