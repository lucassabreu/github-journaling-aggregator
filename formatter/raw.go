package formatter

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/report"
)

type Raw struct {
	w        io.Writer
	sorter   Sorter
	messages []report.Message
}

func NewRaw(w io.Writer) Raw {
	return Raw{
		w:        w,
		messages: make([]report.Message, 0),
	}
}

func (r *Raw) Close() {
	r.messages = r.sorter.SortByCreatedAt(r.messages)
	for _, m := range r.messages {
		r.print(m)
	}
}

func (r *Raw) Format(m report.Message) {
	r.messages = append(r.messages, m)
}

func (r *Raw) print(m report.Message) {
	fmt.Fprintf(r.w, "%v\t%s\t%s\n", *m.Repo.Name, m.CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"), m.Message)
}

func (r *Raw) FormatError(err error) {
	fmt.Fprintln(os.Stderr, err)
}
