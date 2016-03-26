package client

import (
	"time"

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

type NotificationOptions struct {
	repositoryName string
	mentioned      bool
}

type RepoNotificationCounter struct {
	Repository              *github.Repository
	UnreadNotificationCount int
}

type ResponseContent struct {
	Title string
	Body  string
}

type RepoNotificationCounters []*RepoNotificationCounter
type ResponseContents []*ResponseContent
