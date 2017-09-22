package filter

import (
	"fmt"
	"regexp"

	"github.com/google/go-github/github"
)

type EventNameRegExpFilter struct {
	r *regexp.Regexp
}

func NewEventNameRegExpFilter(r *regexp.Regexp) *EventNameRegExpFilter {
	return &EventNameRegExpFilter{r}
}

func (f *EventNameRegExpFilter) Filter(e *github.Event) bool {
	return f.r.MatchString(*e.Repo.Name)
}

func (f *EventNameRegExpFilter) String() string {
	return fmt.Sprintf("EventName like \"%s\"", f.r.String())
}

type EqualsEventName struct {
	eventName string
}

func NewEqualsEventName(eventName string) *EqualsEventName {
	return &EqualsEventName{eventName}
}

func (er *EqualsEventName) Filter(e *github.Event) bool {
	return er.eventName == *e.Repo.Name
}

func (er *EqualsEventName) String() string {
	return fmt.Sprintf("EventName == \"%s\"", er.eventName)
}
