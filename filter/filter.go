package filter

import (
	"fmt"
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

type FilterGroup interface {
	Filter
	Append(...Filter)
	Count() int
}

type filterGroupBase struct {
	filters []Filter
}

func newFilterGroupBase() filterGroupBase {
	return filterGroupBase{
		filters: make([]Filter, 0),
	}
}

// Append a new filter to the group
func (fg *filterGroupBase) Append(filters ...Filter) {
	for _, f := range filters {
		fg.filters = append(fg.filters, f)
	}
}

func (fg *filterGroupBase) Count() int {
	return len(fg.filters)
}

func (fg *filterGroupBase) join(glue string) string {
	r := ""
	for _, f := range fg.filters {
		r = fmt.Sprintf("%s %s %s", r, glue, f)
	}
	return "(" + r[len(glue)+2:] + ")"
}

type OrGroup struct {
	filterGroupBase
}

func NewOrGroup() *OrGroup {
	return &OrGroup{
		filterGroupBase: newFilterGroupBase(),
	}
}

// Filter will call the other filters in the group, if one of then returns true, then the OrGroup will return true, otherwise false
func (fg *OrGroup) Filter(e *github.Event) bool {
	for _, f := range fg.filters {
		if f.Filter(e) {
			return true
		}
	}

	return false
}

func (fg *OrGroup) String() string {
	return fg.join("or")
}

type AndGroup struct {
	filterGroupBase
}

func NewAndGroup() *AndGroup {
	return &AndGroup{
		filterGroupBase: newFilterGroupBase(),
	}
}

// Filter will call the other filters in the group, if one of then returns false, then the AndGroup will return false, otherwise true
func (fg *AndGroup) Filter(e *github.Event) bool {
	for _, f := range fg.filters {
		if !f.Filter(e) {
			return false
		}
	}

	return true
}

func (fg *AndGroup) String() string {
	return fg.join("and")
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

func (f *RepositoryNameRegExpFilter) String() string {
	return fmt.Sprintf("Repo.Name like \"%s\"", f.r.String())
}

type Not struct {
	f Filter
}

func NewNot(f Filter) *Not {
	return &Not{f}
}

func (n *Not) Filter(e *github.Event) bool {
	return !n.f.Filter(e)
}

func (n *Not) String() string {
	return fmt.Sprintf("!%s", n.f)
}

type EqualsRepository struct {
	repoName string
}

func NewEqualsRepository(repoName string) *EqualsRepository {
	return &EqualsRepository{repoName}
}

func (er *EqualsRepository) Filter(e *github.Event) bool {
	return er.repoName == *e.Repo.Name
}

func (er *EqualsRepository) String() string {
	return fmt.Sprintf("Repo.Name == \"%s\"", er.repoName)
}
