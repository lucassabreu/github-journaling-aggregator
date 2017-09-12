package formatter

import (
	"fmt"
	"io"
	"sort"

	"github.com/google/go-github/github"
)

type eventSorter struct {
	events []*github.Event
}

func (es *eventSorter) Len() int {
	return len(es.events)
}

func (es *eventSorter) Swap(i, j int) {
	es.events[i], es.events[j] = es.events[j], es.events[i]
}

func (es *eventSorter) Less(i, j int) bool {
	return es.events[i].CreatedAt.Before(*es.events[j].CreatedAt)
}

func (es *eventSorter) Sort() {
	sort.Sort(es)
}

type Raw struct {
	w      io.Writer
	events []*github.Event
}

func NewRaw(w io.Writer) Raw {
	return Raw{
		w:      w,
		events: make([]*github.Event, 0),
	}
}

func (r *Raw) Close() {
	es := eventSorter{r.events}
	es.Sort()

	for _, e := range r.events {
		r.format(e)
	}
}

func (r *Raw) Format(e *github.Event) {
	r.events = append(r.events, e)
}

func (r *Raw) print(e *github.Event, what string) {
	fmt.Fprintf(r.w, "[%v]\t[%v]\t%s\n", *e.Repo.Name, e.CreatedAt, what)
}

func (r *Raw) format(e *github.Event) {
	pl, err := e.ParsePayload()
	if err != nil {
		r.FormatError(err)
		return
	}

	switch p := pl.(type) {
	case *github.CreateEvent:
		name := *e.Repo.Name
		if p.Ref != nil {
			name = *p.Ref
		}

		r.print(e, fmt.Sprintf("created %s \"%s\"", *p.RefType, name))

	case *github.IssueCommentEvent:
		r.print(e, fmt.Sprintf(
			"%s comment in issue %s#%s with content: \"%s\"",
			p.Action,
			e.Repo.Name,
			p.Issue.Number,
			p.Comment.Body))

	case *github.IssuesEvent:
		if p.Action == nil {
			break
		}
		action := *p.Action
		if action == "opened" || action == "closed" || action == "reopened" || action == "edited" {
			r.print(e, fmt.Sprintf("%s the issue %s#%d (%s)", action, *e.Repo.Name, *p.Issue.Number, *p.Issue.Title))
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
			if *p.PullRequest.Merged == false {
				action = "canceled"
			}
		}

		r.print(e, fmt.Sprintf("%s the pull request %s#%d (%s)", action, *e.Repo.Name, *p.PullRequest.Number, *p.PullRequest.Title))

	// case *github.MemberEvent:
	// case *github.MilestoneEvent:
	// case *github.PublicEvent:
	// case *github.PullRequestReviewCommentEvent:

	case *github.PushEvent:
		messages := make([]string, 0, len(p.Commits))
		for _, c := range p.Commits {
			if c.Message != nil {
				messages = append(messages, *c.Message)
			}
		}

		r.print(e, fmt.Sprintf("pushed %d commits to %s", len(p.Commits), *p.Ref))
	// ignore
	case *github.DeleteEvent:
	default:
		r.print(e, *e.Type)
	}
}

func (r *Raw) FormatError(err error) {
	fmt.Fprintln(r.w, err)
}
