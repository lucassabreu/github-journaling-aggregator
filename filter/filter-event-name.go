package filter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
)

type TypeRegExpFilter struct {
	r *regexp.Regexp
}

func NewTypeRegExpFilter(r *regexp.Regexp) *TypeRegExpFilter {
	return &TypeRegExpFilter{r}
}

func (f *TypeRegExpFilter) Filter(e *github.Event) bool {
	return f.r.MatchString(strings.ToLower(*e.Type))
}

func (f *TypeRegExpFilter) String() string {
	return fmt.Sprintf("Type like \"%s\"", f.r.String())
}

type EqualsType struct {
	typeName string
}

func NewEqualsType(typeName string) *EqualsType {
	return &EqualsType{strings.ToLower(typeName)}
}

func (er *EqualsType) Filter(e *github.Event) bool {
	return er.typeName == strings.ToLower(*e.Type)
}

func (er *EqualsType) String() string {
	return fmt.Sprintf("Type == \"%s\"", er.typeName)
}
