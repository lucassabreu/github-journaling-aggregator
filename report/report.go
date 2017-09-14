package report

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/go-github/github"
)

type Message struct {
	github.Event

	Payload interface{}
	Message string
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
	username   string
	beginning  time.Time
	formatters []Formatter
}

// New Report
func New(client *github.Client, username string, beginning time.Time) Report {
	return Report{
		client:     client,
		username:   username,
		beginning:  beginning,
		formatters: make([]Formatter, 0),
	}
}

// AttachFormatter to receive the messages
func (r *Report) AttachFormatter(f Formatter) {
	r.formatters = append(r.formatters, f)
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
			r.formatError(err)
			return
		}

		for _, e := range events {
			if e.CreatedAt.Before(beginning) {
				return
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

		r.format(Message{*e, p, fmt.Sprintf("created %s \"%s\"", *p.RefType, name)})

	case *github.IssueCommentEvent:
		r.format(Message{*e, p, fmt.Sprintf(
			"%v comment in issue %v#%d with content: \"%v\"",
			*p.Action,
			*e.Repo.Name,
			*p.Issue.Number,
			*p.Comment.Body)})

	case *github.IssuesEvent:
		if p.Action == nil {
			break
		}
		action := *p.Action
		if action == "opened" || action == "closed" || action == "reopened" || action == "edited" {
			r.format(Message{*e, p, fmt.Sprintf("%s the issue %s#%d (%s)", action, *e.Repo.Name, *p.Issue.Number, *p.Issue.Title)})
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
		r.format(Message{*e, p, fmt.Sprintf("%s the pull request %s#%d (%s)", action, *e.Repo.Name, *p.PullRequest.Number, *p.PullRequest.Title)})

		// case *github.MemberEvent:
		// case *github.MilestoneEvent:
		// case *github.PublicEvent:
	case *github.PullRequestReviewCommentEvent:
		if p.Action == nil {
			break
		}
		r.format(Message{*e, p, fmt.Sprintf(
			"%s a comment in the pull request %s#%d with \"%s\"",
			*p.Action,
			*e.Repo.Name,
			*p.PullRequest.Number,
			*p.Comment.Body,
		)})

	case *github.PushEvent:
		for _, c := range p.Commits {
			if !*c.Distinct {
				continue
			}
			r.format(Message{
				*e,
				p,
				fmt.Sprintf("pushed commit %s with message: %v",
					*c.SHA,
					strings.Split(*c.Message, "\n")[0]),
			})
		}

	// ignore
	case *github.DeleteEvent:
	default:
		r.format(Message{*e, p, *e.Type})
	}
	return nil
}
