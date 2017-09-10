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

	r := report{
		client:    githubclient.NewGithubClient(username, token, nil),
		username:  username,
		beginning: beginning,

		githubEventsChan: make(chan *github.Event),
	}

	return r.run()
}

type reportEventType int

const (
	_                = iota
	PULL_REQUEST int = iota
	COMMENT_PULL_REQUEST
	COMMIT
)

type reportEvent struct {
	eventType reportEventType
	repo      string
	username  string
	when      time.Time
	detail    string
	ge        github.Event
}

func (re *reportEvent) String() string {
	return fmt.Sprintf(
		"Event %s ocurred at %s into repository %s with this context: %s",
		re.ge.Type,
		re.when,
		re.repo,
		re.detail)
}

type report struct {
	client    *github.Client
	username  string
	beginning time.Time

	err              error
	done             chan struct{}
	githubEventsChan chan *github.Event
	reportEventsChan chan *reportEvent
	warningsChan     chan error
	warnings         []error
}

func (r *report) run() error {
	defer close(r.warningsChan)
	r.done = make(chan struct{})

	r.warningsChan = make(chan error)
	r.warnings = make([]error, 0)

	r.githubEventsChan = make(chan *github.Event)
	r.reportEventsChan = make(chan *reportEvent)

	go r.warningCollect()
	go r.fetchEvents()
	go r.processEvents()

	for re := range r.reportEventsChan {
		fmt.Println(re)
	}

	<-r.done
	return r.err
}

func (r *report) warningCollect() {
	for w := range r.warningsChan {
		r.warnings = append(r.warnings, w)
	}
}

func (r *report) processEvents() {
	defer func() {
		close(r.reportEventsChan)
		close(r.done)
	}()

	for {
		select {
		case ge := <-r.githubEventsChan:
			if ge == nil {
				continue
			}

			if ge.Type == nil {
				fmt.Println(ge)
				return
			}

			p, err := ge.ParsePayload()
			if err != nil {
				r.warningsChan <- err
				continue
			}

			switch p := p.(type) {
			case github.CreateEvent:
				go r.processCreateEvent(ge, p)
			case github.IssueCommentEvent:
				go r.processIssueCommentEvent(ge, p)
			case github.IssuesEvent:
				go r.processIssuesEvent(ge, p)
			case github.MemberEvent:
				go r.processMemberEvent(ge, p)
			case github.MilestoneEvent:
				go r.processMilestoneEvent(ge, p)
			case github.PublicEvent:
				go r.processPublicEvent(ge, p)
			case github.PullRequestEvent:
				go r.processPullRequestEvent(ge, p)
			case github.PullRequestReviewCommentEvent:
				go r.processPullRequestReviewCommentEvent(ge, p)
			case github.PushEvent:
				go r.processPushEvent(ge, p)
			default:
				go r.processOtherEvents(ge)
			}
		case <-r.done:
			return
		}
	}

	close(r.done)
}

func (r *report) fetchEvents() {
	beginning := r.beginning.Add(-24 * time.Hour)
	beginning = time.Date(beginning.Year(), beginning.Month(), beginning.Day(), 23, 59, 59, int(time.Second)-1, time.Local)

	defer close(r.githubEventsChan)

	opt := &github.ListOptions{
		PerPage: 10,
	}
	for {
		events, resp, err := r.client.Activity.ListEventsPerformedByUser(context.Background(), r.username, false, opt)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		for _, e := range events {
			if e.CreatedAt.Before(beginning) {
				return
			}

			select {
			case r.githubEventsChan <- e:
			case <-r.done:
				return
			}
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}
