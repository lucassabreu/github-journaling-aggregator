package formatter

//go:generate go-bindata -ignore=assets/.gitignore -pkg $GOPACKAGE -o assets.go assets/

import (
	"html/template"
	"io"

	"github.com/lucassabreu/github-journaling-aggregator/report"
)

type HTML struct {
	w        io.Writer
	sorter   Sorter
	messages []report.Message
}

func NewHTML(w io.Writer) HTML {
	return HTML{
		w:        w,
		messages: make([]report.Message, 0),
	}
}

func (h *HTML) Format(m report.Message) {
	h.messages = append(h.messages, m)
}

func (h *HTML) Close() {
	data := MustAsset("assets/report.html")
	template.New("report").Parse(string(data))

	h.messages = h.sorter.SortByCreatedAt(h.messages)
	for _, m := range h.messages {
		h.print(m)
	}
}

func (h *HTML) print(m report.Message) {
}
