OctoPatrol (octop command)
========

OctoPatrol is the command line tool for watching notifications from repositories you watch.

## Motivation
Checking OSS repositories with Github Notification APIs is supported by several softwares, or Github send e-mails.
Or an official Github page's navigation bar shows the unread notifications.

However, after following over 5 active repositories, notifications are so noisy to check.
They overwhelms and disturbs other contacts, so you must controll expressive info.

In my case, any tool doesn't fit for this solution.
So created efficient and convenient environment for catching up fresh discussions.
And wished to finish off all tasks on terminal display as I could.

## Use cases
- See notifications from repositories you follow
- Check received mensions from your private team
- Watch repository issues and pull-reqs separately

## Installation
```bash
$ go install github.com/timakin/octop
```

After installed binary packed octopatrol command, authenticate your account for repository tracking.

_Get your access token following this post._
https://help.github.com/articles/creating-an-access-token-for-command-line-use/

## Notification tracking
```
octop n - notification tracking
octop i - issue tracking, with an interactive selection of repo
octop p - pull-reqs tracking, with an interactive selection of repo
```

## Reference

- notifications
  - [user-notifications](https://developer.github.com/v3/activity/notifications/#list-your-notifications)
  - [repo-notifications](https://developer.github.com/v3/activity/notifications/#list-your-notifications-in-a-repository)
- received_events
 - [user-events](https://developer.github.com/v3/activity/events/#list-events-that-a-user-has-received)
 - [orgs-events](https://developer.github.com/v3/activity/events/#list-events-for-an-organization)
- event_types
 - [types](https://developer.github.com/v3/activity/events/types/) 
- notifications
  - [user](https://developer.github.com/v3/activity/notifications/#list-your-notifications)
  - [repo](https://developer.github.com/v3/activity/notifications/#list-your-notifications-in-a-repository)
- issues
  - [repo-issues](https://developer.github.com/v3/activity/events/#list-issue-events-for-a-repository)
  - [orgs-issues](https://developer.github.com/v3/activity/events/#list-public-events-for-an-organization)
  - [issue-events](https://developer.github.com/v3/issues/events/)
- Use Cases
 - [Qiita](http://qiita.com/awakia/items/bd4cdfab2b552e2151ad)

## Contribution

1. Fork ([https://github.com/timakin/octop/fork](https://github.com/timakin/octop/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[timakin](https://github.com/timakin)
