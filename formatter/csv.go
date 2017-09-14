package formatter

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/report"
)

type CSV struct {
	w        *csv.Writer
	messages []report.Message
}

func NewCSV(w io.Writer) CSV {
	return CSV{
		w:        csv.NewWriter(w),
		messages: make([]report.Message, 0),
	}
}

func (csv *CSV) Close() {
	for _, m := range csv.messages {
		csv.w.Write([]string{
			*m.Repo.Name,
			m.CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
			m.Message,
		})
	}
	csv.w.Flush()
}

func (csv *CSV) Format(m report.Message) {
	csv.messages = append(csv.messages, m)
}

func (csv *CSV) FormatError(err error) {
	fmt.Fprintln(os.Stderr, err)
}
