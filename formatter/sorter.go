package formatter

import (
	"sort"

	"github.com/lucassabreu/github-journaling-aggregator/report"
)

type Sorter struct {
	messages []report.Message
	less     func(ms []report.Message, i, j int) bool
}

func (s *Sorter) Len() int {
	return len(s.messages)
}

func (s *Sorter) Less(i, j int) bool {
	return s.less(s.messages, i, j)
}

func (s *Sorter) Swap(i, j int) {
	s.messages[i], s.messages[j] = s.messages[j], s.messages[i]
}

func (s *Sorter) Sort(ms []report.Message, less func(ms []report.Message, i, j int) bool) []report.Message {
	s.messages = ms
	s.less = less
	sort.Sort(s)
	return s.messages
}

func (s *Sorter) SortByCreatedAt(ms []report.Message) []report.Message {
	return s.Sort(ms, func(ms []report.Message, i, j int) bool {
		return ms[i].CreatedAt.Before(*ms[j].CreatedAt)
	})
}
