OctoPatrol (op command)
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
$ go get -d github.com/timakin/op
```

After installed binary packed octopatrol command, authenticate your account for repository tracking.

```
op init
> Token: ***************************************
> Authentication Verified!
```

_Get your access token following this post._
https://help.github.com/articles/creating-an-access-token-for-command-line-use/

## Notification tracking
```
op
> list all repositories that you are watching.
> And decided the repo, display events.

[ Repo ] open_issues_count: xxx
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
Type       | name |   created_at   |

[Issue]    | .... | 2016-01-03-xxx |
[Issue]    | .... | 2016-01-02-xxx |
[Issue]    | .... | 2016-01-01-xxx |
[Pull-req] | .... | 2015-12-31-xxx |
[Pull-req] | .... | 2015-12-30-xxx |
```

```
op issue [ -r repo_name ] [-p participating]

> list all repos you fallow, and select repository seeing names.
> and show 10 latest issues.

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
Type       | name |   created_at   |

[Issue]    | .... | 2016-01-03-xxx |
[Issue]    | .... | 2016-01-02-xxx |
[Issue]    | .... | 2016-01-01-xxx |
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
```

```
op pr [ -r repo_name ]

>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
Type       | name |   created_at   |

[Pull-req] | .... | 2015-12-31-xxx |
[Pull-req] | .... | 2015-12-30-xxx |
[Pull-req] | .... | 2015-12-29-xxx |
[Pull-req] | .... | 2015-12-28-xxx |
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
```

```
op explore

> hot topics in all github news streams.
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

1. Fork ([https://github.com/timakin/op/fork](https://github.com/timakin/op/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[timakin](https://github.com/timakin)
