package formatter

import (
	"io"
	"os"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/report"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/crypto/ssh/terminal"
)

type Table struct {
	sorter   MessageSorter
	tw       *tablewriter.Table
	messages []report.Message
	errors   []error
}

func NewGroupLineTable(w io.Writer) Table {
	t := NewTable(w)
	t.tw.SetHeader([]string{"Repository", "Created At", "What"})
	t.tw.SetAutoMergeCells(true)
	t.tw.SetRowLine(true)
	return t
}

func NewMDTable(w io.Writer) Table {
	t := NewTable(w)
	t.tw.SetHeader([]string{"Repository", "Created At", "What"})
	t.tw.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	t.tw.SetCenterSeparator("|")
	return t
}

func NewTable(w io.Writer) Table {
	tw := tablewriter.NewWriter(w)
	return Table{
		tw:       tw,
		messages: make([]report.Message, 0),
		errors:   make([]error, 0),
	}
}

func (t *Table) Close() {
	t.sorter.Sort(t.messages)

	var wM, wR int = 0, 0
	t.tw.SetHeader([]string{"Repository", "Created At", "What"})
	lines := make([][]string, len(t.messages))
	for i, m := range t.messages {
		if wM < len(m.Message) {
			wM = len(m.Message)
		}
		if wR < len(*m.Repo.Name) {
			wR = len(*m.Repo.Name)
		}
		lines[i] = []string{
			m.CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
			*m.Repo.Name,
			m.Message,
		}
	}

	if width, _, err := terminal.GetSize(int(os.Stdin.Fd())); err == nil {
		width = width - 30 - wR
		if width < wM {
			wM = width
		}
	}

	t.tw.SetColWidth(wM)
	t.tw.AppendBulk(lines)
	t.tw.Render()
}

func (t *Table) Format(m report.Message) {
	t.messages = append(t.messages, m)
}

func (t *Table) FormatError(err error) {
	t.errors = append(t.errors, err)
}
