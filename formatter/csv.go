package formatter

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/report"
)

type CSV struct {
	w        *csv.Writer
	sorter   Sorter
	messages []report.Message
}

func NewCSV(w io.Writer) CSV {
	return CSV{
		w:        csv.NewWriter(w),
		messages: make([]report.Message, 0),
	}
}

func (csv *CSV) Close() {
	csv.messages = csv.sorter.SortByCreatedAt(csv.messages)
	for _, m := range csv.messages {
		err := csv.w.Write([]string{
			*m.Repo.Name,
			m.CreatedAt.In(time.Local).Format("2006-01-02 15:04:05"),
			m.Message,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
	csv.w.Flush()
}

func (csv *CSV) Format(m report.Message) {
	csv.messages = append(csv.messages, m)
}

func (csv *CSV) FormatError(err error) {
	fmt.Fprintln(os.Stderr, err)
}
