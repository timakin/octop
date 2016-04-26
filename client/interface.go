package client

import (
	"time"

	"net/url"

	"github.com/google/go-github/github"
	"github.com/patrickmn/go-cache"
)

type Instance struct {
	cache *cache.Cache
	ghCli *github.Client
}

func New() *Instance {
	httpClient := newAuthenticatedClient()
	ghCli := github.NewClient(httpClient)

	I := &Instance{
		cache: cache.New(5*time.Minute, 30*time.Second),
		ghCli: ghCli,
	}

	return I
}

func (i *Instance) SetRemoteHost(baseUrl *url.URL) {
	i.ghCli.BaseURL = baseUrl
}

type NotificationOptions struct {
	repositoryName string
	mentioned      bool
}

type FilteredNotification struct {
	Title      string
	Repository *github.Repository
	URL        string
}

type RepoNotificationCounter struct {
	Repository              *github.Repository
	UnreadNotificationCount int
}

type ResponseContent struct {
	Title string
	Body  string
	Owner string
	URL   string
}

type RepoNotificationCounters []*RepoNotificationCounter
type ResponseContents []*ResponseContent
type FilteredNotifications []*FilteredNotification
