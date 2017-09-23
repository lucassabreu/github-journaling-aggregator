package report

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/lucassabreu/github-journaling-aggregator/filter"
)

// Message represents some event from GitHub, but with a little preprocessing
type Message struct {
	github.Event

	EventName string
	Payload   interface{}
	Message   string
}

// Formatter is a interface for output Formatters to implement
type Formatter interface {
	Format(Message)
	FormatError(error)
	Close()
}

// Report will read data from GitHub and forward then to the Formatters
type Report struct {
	client     *github.Client
	user       *github.User
	beginning  time.Time
	formatters []Formatter
	filter     filter.Filter
}

// New Report
func New(client *github.Client, beginning time.Time) Report {
	return Report{
		client:     client,
		beginning:  beginning,
		formatters: make([]Formatter, 0),
		filter:     filter.DefaultFilter,
	}
}

// AttachFormatter to receive the messages
func (r *Report) AttachFormatter(f Formatter) {
	r.formatters = append(r.formatters, f)
}

func (r *Report) SetFilter(filter filter.Filter) {
	r.filter = filter
}

func (r *Report) Run() {
	err := r.setUser()
	if err != nil {
		r.formatError(err)
		return
	}

	r.getEvents()
	for _, f := range r.formatters {
		f.Close()
	}
}

func (r *Report) setUser() (err error) {
	r.user, _, err = r.client.Users.Get(context.Background(), "")
	return
}

func (r *Report) getEvents() {
	beginning := r.beginning.Add(-24 * time.Hour)
	beginning = time.Date(beginning.Year(), beginning.Month(), beginning.Day(), 23, 59, 59, int(time.Second)-1, time.Local)

	opt := &github.ListOptions{PerPage: 10}
	for {
		events, resp, err := r.client.Activity.ListEventsPerformedByUser(context.Background(), *r.user.Login, false, opt)

		if err != nil {
			r.formatError(err)
			return
		}

		for _, e := range events {
			if e.CreatedAt.Before(beginning) {
				return
			}

			if !r.filter.Filter(e) {
				continue
			}

			err := r.forward(e)
			if err != nil {
				r.formatError(err)
			}
		}

		if resp.NextPage == 0 {
			return
		}
		opt.Page = resp.NextPage
	}
}

func (r *Report) formatError(err error) {
	for _, f := range r.formatters {
		f.FormatError(err)
	}
}

func (r *Report) format(m Message) {
	for _, f := range r.formatters {
		f.Format(m)
	}
}

func (r *Report) forward(e *github.Event) error {
	pl, err := e.ParsePayload()
	if err != nil {
		return err
	}

	switch p := pl.(type) {
	case *github.CreateEvent:
		name := *e.Repo.Name
		if p.Ref != nil {
			name = *p.Ref
		}

		t := fmt.Sprintf("created %s", *p.RefType)

		r.format(Message{
			*e,
			t,
			p,
			fmt.Sprintf("%s \"%s\"", t, name),
		})

	case *github.IssueCommentEvent:
		t := fmt.Sprintf("%s comment", *p.Action)
		r.format(Message{
			*e,
			t,
			p,
			fmt.Sprintf(
				"%s in issue %v#%d with content: \"%v\"",
				t,
				*e.Repo.Name,
				*p.Issue.Number,
				*p.Comment.Body,
			),
		})

	case *github.IssuesEvent:
		if p.Action == nil {
			break
		}
		action := *p.Action
		if action == "opened" || action == "closed" || action == "reopened" || action == "edited" {
			r.format(Message{
				*e,
				action + " issue",
				p,
				fmt.Sprintf("%s the issue %s#%d (%s)", action, *e.Repo.Name, *p.Issue.Number, *p.Issue.Title),
			})
		}

	case *github.PullRequestEvent:
		if p.Action == nil {
			break
		}
		action := *p.Action
		if !(action == "opened" || action == "reopened" || action == "edited" || action == "closed") {
			break
		}

		if action == "closed" {
			action = "merged"
			if !*p.PullRequest.Merged {
				action = "canceled"
			}
		}
		r.format(Message{
			*e,
			action + " pull request",
			p,
			fmt.Sprintf("%s the pull request %s#%d (%s)", action, *e.Repo.Name, *p.PullRequest.Number, *p.PullRequest.Title),
		})

	case *github.PullRequestReviewCommentEvent:
		if p.Action == nil {
			break
		}
		r.format(Message{
			*e,
			*p.Action + " pull request comment",
			p,
			fmt.Sprintf(
				"%s a comment in the pull request %s#%d with \"%s\"",
				*p.Action,
				*e.Repo.Name,
				*p.PullRequest.Number,
				*p.Comment.Body,
			),
		})

	case *github.PushEvent:
		for _, c := range p.Commits {
			if !*c.Distinct {
				continue
			}
			r.format(Message{
				*e,
				"pushed commit",
				p,
				fmt.Sprintf("pushed commit %s with message: %v",
					*c.SHA,
					strings.Split(*c.Message, "\n")[0]),
			})
		}

	// ignore
	case *github.DeleteEvent:
	default:
		r.format(Message{*e, "unknown", p, *e.Type})
	}
	return nil
}
