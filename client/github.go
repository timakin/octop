package client

import (
	"log"
	"sort"
	"strings"

	"github.com/google/go-github/github"
	"github.com/patrickmn/go-cache"
)

func (i Instance) GetNotifications(owner string, repo string) FilteredNotifications {
	cacheKey := "notifications_" + owner + "_" + repo
	if cv, found := i.cache.Get(cacheKey); found {
		cachedNotifications := cv.(FilteredNotifications)
		return cachedNotifications
	}
	opt := &github.NotificationListOptions{All: true}
	notifications, _, err := i.ghCli.Activity.ListRepositoryNotifications(owner, repo, opt)
	if err != nil {
		log.Fatal(err)
	}
	filtered := i.filterNotifications(notifications)

	i.cache.Set(cacheKey, filtered, cache.DefaultExpiration)
	return filtered
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
			URL:   *issue.HTMLURL,
		})
	}
	return res
}

func (i Instance) filterNotifications(notifications []github.Notification) FilteredNotifications {
	var filtered FilteredNotifications
	for _, notification := range notifications {
		notification := notification
		filtered = append(filtered, &FilteredNotification{
			Title:      *notification.Subject.Title,
			Repository: notification.Repository,
			URL:        i.toHTMLURL(notification),
		})
	}
	return filtered
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
	notifications := i.GetNotifications(*owner, *repoName)
	unreadRepositoryNotifications := NotificationFilter(notifications, func(n *FilteredNotification) bool {
		return *n.Repository.Owner.Login == *owner && *n.Repository.Name == *repoName
	})
	return len(unreadRepositoryNotifications)
}

func (i Instance) toHTMLURL(n github.Notification) string {
	s := strings.Replace(*n.Subject.URL, "api.", "", 1)
	s = strings.Replace(s, "repos/", "", 1)
	return s
}
