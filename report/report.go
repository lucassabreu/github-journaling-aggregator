package report

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"github.com/lucassabreu/github-journaling-aggregator/githubclient"
)

func RunReportGen(username string, token string, beginning time.Time) error {
	client := githubclient.NewGithubClient(username, token, nil)
	done := make(chan struct{})

	events := getAllEventsSince(done, client, username, beginning)
	for e := range events {
		fmt.Println(e.CreatedAt, beginning)
	}

	return nil
}

func getAllEventsSince(done <-chan struct{}, client *github.Client, username string, beginning time.Time) <-chan *github.Event {
	out := make(chan *github.Event)

	beginning = beginning.Add(-24 * time.Hour)
	beginning = time.Date(beginning.Year(), beginning.Month(), beginning.Day(), 23, 59, 59, int(time.Second)-1, time.Local)

	go func() {
		defer close(out)

		opt := &github.ListOptions{
			PerPage: 10,
		}
		for {
			events, resp, err := client.Activity.ListEventsPerformedByUser(context.Background(), username, false, opt)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				break
			}

			for _, e := range events {
				if e.CreatedAt.Before(beginning) {
					return
				}

				select {
				case out <- e:
				case <-done:
					return
				}
			}

			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}
	}()

	return out
}
