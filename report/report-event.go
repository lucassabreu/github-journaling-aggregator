package report

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

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
