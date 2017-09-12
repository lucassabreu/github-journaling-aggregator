package report

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
)

type Formatter interface {
	Format(*github.Event)
	FormatError(error)
	Close()
}

type Report struct {
	client     *github.Client
	username   string
	beginning  time.Time
	formatters []Formatter
}

func New(client *github.Client, username string, beginning time.Time) Report {
	return Report{
		client:     client,
		username:   username,
		beginning:  beginning,
		formatters: make([]Formatter, 0),
	}
}

func (r *Report) AttachFormatter(f Formatter) {
	r.formatters = append(r.formatters, f)
}

func (r *Report) format(e *github.Event) {
	for _, f := range r.formatters {
		f.Format(e)
	}
}

func (r *Report) formatError(err error) {
	for _, f := range r.formatters {
		f.FormatError(err)
	}
}

func (r *Report) Run() {
	r.getEvents()
	for _, f := range r.formatters {
		f.Close()
	}
}

func (r *Report) getEvents() {
	beginning := r.beginning.Add(-24 * time.Hour)
	beginning = time.Date(beginning.Year(), beginning.Month(), beginning.Day(), 23, 59, 59, int(time.Second)-1, time.Local)

	opt := &github.ListOptions{PerPage: 10}
	for {
		events, resp, err := r.client.Activity.ListEventsPerformedByUser(context.Background(), r.username, false, opt)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		for _, e := range events {
			if e.CreatedAt.Before(beginning) {
				return
			}
			r.format(e)
		}

		if resp.NextPage == 0 {
			return
		}
		opt.Page = resp.NextPage
	}

}
