package cerrlist

import "github.com/lucassabreu/github-journaling-aggregator/syncker"

type ConcurrencyErrorList struct {
	s      syncker.Syncker
	errors []error
}

func New(errors []error) ConcurrencyErrorList {
	return ConcurrencyErrorList{
		s:      syncker.New(),
		errors: errors,
	}
}

func (cel *ConcurrencyErrorList) Append(err error) {
	cel.s.Sync(func() {
		cel.errors = append(cel.errors, err)
	})
}

func (cel ConcurrencyErrorList) GetErrors() []error {
	return cel.errors
}
