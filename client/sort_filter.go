package client

import (
	"github.com/google/go-github/github"
)

func (r RepoNotificationCounters) Len() int {
	return len(r)
}

func (r RepoNotificationCounters) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r RepoNotificationCounters) Less(i, j int) bool {
	return r[i].UnreadNotificationCount > r[j].UnreadNotificationCount
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

func IssueEventFilter(vs []github.IssueEvent, f func(github.IssueEvent) bool) []github.IssueEvent {
	vsf := make([]github.IssueEvent, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func PullReqFilter(vs []github.PullRequestEvent, f func(github.PullRequestEvent) bool) []github.PullRequestEvent {
	vsf := make([]github.PullRequestEvent, 0)
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
