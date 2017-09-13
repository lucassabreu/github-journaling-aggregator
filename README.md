GitHub Journaling Aggregator
============================

Once GitHub is the main tool for many developers to control theirs source and issues. This tool is meant to help developers remember what they did by listing the recent events they generated in GitHub, such as pushs, creating issues, creating pull requests, commenting on pull requests, etc.

Install
-------

For now this cli is only available to install by `go get`, when the main GitHub common events were coded in this tool then will be releases. But for now to install (and update) it use:

```sh
go get -u github.com/lucassabreu/github-journaling-aggregator
```

Usage
-----

Here is the help from the cli:

```
Create a simple report using your activity feed at GitHub.

	Will receive a username, access token and a beginning date and generate a report based on the users activity feed on GitHub

Usage:
  github-journaling-aggregator <username> [flags]

Flags:
      --config string   config file (default is $HOME/.github-journaling-aggregator.yaml)
      --date string     set a beginning date (format 2017-12-31)
  -d, --days int        use today as beginning date
  -h, --help            help for github-journaling-aggregator
  -w, --last-week       use the last sunday as beginning date
  -t, --today           use today as beginning date (default)
      --token string    github access token (or user password), if not set $GITHUB_TOKEN will be used
  -y, --yesterday       use yesterday as beginning date
```
