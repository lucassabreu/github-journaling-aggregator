package filter

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/google/go-github/github"
)

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

type EqualsRepository struct {
	repoName string
}

func NewEqualsRepository(repoName string) *EqualsRepository {
	return &EqualsRepository{strings.ToLower(repoName)}
}

func (er *EqualsRepository) Filter(e *github.Event) bool {
	return er.repoName == strings.ToLower(*e.Repo.Name)
}

func (er *EqualsRepository) String() string {
	return fmt.Sprintf("Repo.Name == \"%s\"", er.repoName)
}
