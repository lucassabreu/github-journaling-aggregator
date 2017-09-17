package filter

import (
	"regexp"

	"github.com/google/go-github/github"
)

// Filter will control if the github.Event should or not be shown
type Filter interface {
	// Filter should return if the event will be reported or not
	Filter(e *github.Event) bool
}

type defaultFilter struct {
}

func (df defaultFilter) Filter(e *github.Event) bool {
	return true
}

var DefaultFilter = new(defaultFilter)

type FilterGroup struct {
	filters []Filter
}

func NewFilterGroup() *FilterGroup {
	return &FilterGroup{
		filters: make([]Filter, 0),
	}
}

// Append a new filter to the group
func (fg *FilterGroup) Append(f Filter) {
	fg.filters = append(fg.filters, f)
}

func (fg *FilterGroup) Count() int {
	return len(fg.filters)
}

// Filter will call the other filters in the group, if one of then returns true, then the FilterGroup will return true, otherwise false
func (fg *FilterGroup) Filter(e *github.Event) bool {
	for _, f := range fg.filters {
		if f.Filter(e) {
			return true
		}
	}

	return false
}

type FilterFunc func(e *github.Event) bool

func (ff FilterFunc) Filter(e *github.Event) bool {
	return ff(e)
}

type RepositoryNameRegExpFilter struct {
	r *regexp.Regexp
}

func NewRepositoryNameRegExpFilter(r *regexp.Regexp) *RepositoryNameRegExpFilter {
	return &RepositoryNameRegExpFilter{r}
}
func (f *RepositoryNameRegExpFilter) Filter(e *github.Event) bool {
	return f.r.MatchString(*e.Repo.Name)
}
