package client

import (
	"log"
	"sort"

	"github.com/google/go-github/github"
	"github.com/patrickmn/go-cache"
)

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

func (i Instance) GetIssues(owner string, repo string) ResponseContents {
	opt := &github.IssueListByRepoOptions{State: "open"}
	issues, _, err := i.ghCli.Issues.ListByRepo(owner, repo, opt)
	if err != nil {
		log.Fatal(err)
	}

	return i.convertIssuesToResContents(issues)
}

func (i Instance) convertIssuesToResContents(issues []github.Issue) ResponseContents {
	var res ResponseContents
	for _, issue := range issues {
		issue := issue
		res = append(res, &ResponseContent{
			Title: *issue.Title,
			Owner: *issue.User.Login,
			Body:  *issue.Body,
		})
	}
	return res
}

func (i Instance) GetPullRequests(owner string, repo string) []github.PullRequest {
	opt := &github.PullRequestListOptions{}
	pullreqs, _, err := i.ghCli.PullRequests.List(owner, repo, opt)
	if err != nil {
		log.Fatal(err)
	}

	pullreqs = PullReqFilter(pullreqs, func(p github.PullRequest) bool {
		isOpen := *p.State == "open"
		return isOpen
	})
	return pullreqs
}

func (i Instance) GetRepoNotificationCounters() RepoNotificationCounters {
	repos := i.GetListFollowingRepository()
	repoNotificationCounters := make(RepoNotificationCounters, len(repos))
	for index, repo := range repos {
		repo := repo
		unreadCount := i.countUnreadRepositoryNotification(repo.Owner.Login, repo.Name)
		repoNotificationCounter := &RepoNotificationCounter{
			Repository:              &repo,
			UnreadNotificationCount: unreadCount,
		}
		repoNotificationCounters[index] = repoNotificationCounter
	}
	sort.Sort(repoNotificationCounters)
	return repoNotificationCounters
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
