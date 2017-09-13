package formatter

import (
	"sort"

	"github.com/lucassabreu/github-journaling-aggregator/report"
)

type MessageSorter struct {
	messages []report.Message
}

func (es *MessageSorter) Len() int {
	return len(es.messages)
}

func (es *MessageSorter) Swap(i, j int) {
	es.messages[i], es.messages[j] = es.messages[j], es.messages[i]
}

func (es *MessageSorter) Less(i, j int) bool {
	return es.messages[i].CreatedAt.Before(*es.messages[j].CreatedAt)
}

func (es *MessageSorter) Sort(messages []report.Message) {
	es.messages = messages
	sort.Sort(es)
}
