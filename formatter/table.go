package formatter

import (
	"io"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/report"
	"github.com/olekukonko/tablewriter"
)

type Table struct {
	sorter   MessageSorter
	tw       *tablewriter.Table
	messages []report.Message
	errors   []error
}

func NewGroupLineTable(w io.Writer) Table {
	tw := tablewriter.NewWriter(w)
	tw.SetHeader([]string{"Repository", "Created At", "What"})
	tw.SetAutoMergeCells(true)
	tw.SetRowLine(true)

	return Table{
		tw:       tw,
		messages: make([]report.Message, 0),
		errors:   make([]error, 0),
	}
}

func NewMDTable(w io.Writer) Table {
	tw := tablewriter.NewWriter(w)
	tw.SetHeader([]string{"Repository", "Created At", "What"})
	tw.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	tw.SetCenterSeparator("|")

	return Table{
		tw:       tw,
		messages: make([]report.Message, 0),
		errors:   make([]error, 0),
	}
}

func NewTable(w io.Writer) Table {
	return Table{
		tw:       tablewriter.NewWriter(w),
		messages: make([]report.Message, 0),
		errors:   make([]error, 0),
	}
}

func (t *Table) Close() {
	t.sorter.Sort(t.messages)

	t.tw.SetHeader([]string{"Repository", "Created At", "What"})
	for _, m := range t.messages {
		t.tw.Append([]string{
			m.CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
			*m.Repo.Name,
			m.Message,
		})
	}
	t.tw.Render()
}

func (t *Table) Format(m report.Message) {
	t.messages = append(t.messages, m)
}

func (t *Table) FormatError(err error) {
	t.errors = append(t.errors, err)
}
