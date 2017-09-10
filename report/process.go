package report

import (
	"fmt"

	"github.com/google/go-github/github"
)

func (r *report) sendReportEvent(ge *github.Event, d string) {
	e := reportEvent{
		ge:       *ge,
		repo:     *ge.Repo.FullName,
		username: fmt.Sprintf("%s (%s)", ge.Actor.Login, ge.Actor.Name),
		when:     *ge.CreatedAt,
		detail:   d,
	}

	r.reportEventsChan <- &e
}

func (r *report) processOtherEvents(ge *github.Event) {
	r.warningsChan <- fmt.Errorf("Event %s will be ignored", ge.Type)
}

func (r *report) processCreateEvent(ge *github.Event, p github.CreateEvent) {
	name := p.Ref
	if p.RefType == nil || *p.RefType != "repository" {
		name = ge.Repo.FullName
	}
	r.sendReportEvent(ge, fmt.Sprintf("Created the %s: \"%s\"", p.RefType, name))
}

func (r *report) processIssueCommentEvent(ge *github.Event, p github.IssueCommentEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred IssueCommentEventIssueCommentEvent"))
}

func (r *report) processIssuesEvent(ge *github.Event, p github.IssuesEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred IssuesEvent"))
}

func (r *report) processMemberEvent(ge *github.Event, p github.MemberEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred MemberEvent"))
}

func (r *report) processMilestoneEvent(ge *github.Event, p github.MilestoneEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred MilestoneEvent"))
}

func (r *report) processPublicEvent(ge *github.Event, p github.PublicEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred PublicEvent"))
}

func (r *report) processPullRequestEvent(ge *github.Event, p github.PullRequestEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred PullRequestEvent"))
}

func (r *report) processPullRequestReviewCommentEvent(ge *github.Event, p github.PullRequestReviewCommentEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred PullRequestReviewCommentEvent"))
}

func (r *report) processPushEvent(ge *github.Event, p github.PushEvent) {
	r.sendReportEvent(ge, fmt.Sprintf("Event occurred PushEvent"))
}
