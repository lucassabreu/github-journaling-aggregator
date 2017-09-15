package filter

import "github.com/google/go-github/github"

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

func NewFilterGroup(filters []Filter) FilterGroup {
	return FilterGroup{
		filters: filters,
	}
}

// Append a new filter to the group
func (fg *FilterGroup) Append(f Filter) {
	fg.filters = append(fg.filters, f)
}

// Filter will call the other filters in the group, if one of then returns false, then the FilterGroup will return false, otherwise true
func (fg *FilterGroup) Filter(e *github.Event) bool {
	for _, f := range fg.filters {
		if !f.Filter(e) {
			return false
		}
	}

	return true
}

type FilterFunc func(e *github.Event) bool

func (ff FilterFunc) Filter(e *github.Event) bool {
	return ff(e)
}
