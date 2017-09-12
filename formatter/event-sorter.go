package formatter

import (
	"sort"

	"github.com/google/go-github/github"
)

type EventSorter struct {
	events []*github.Event
}

func (es *EventSorter) Len() int {
	return len(es.events)
}

func (es *EventSorter) Swap(i, j int) {
	es.events[i], es.events[j] = es.events[j], es.events[i]
}

func (es *EventSorter) Less(i, j int) bool {
	return es.events[i].CreatedAt.Before(*es.events[j].CreatedAt)
}

func (es *EventSorter) Sort(events []*github.Event) {
	es.events = events
	sort.Sort(es)
}
