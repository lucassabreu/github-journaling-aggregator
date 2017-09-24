package formatter

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/report"
	// markdown "github.com/shurcooL/github_flavored_markdown"
)

type HTML struct {
	w      io.Writer
	sorter Sorter

	since    time.Time
	messages []report.Message
	errors   []error
}

func NewHTML(w io.Writer, since time.Time) HTML {
	return HTML{
		w:        w,
		since:    since,
		messages: make([]report.Message, 0),
		errors:   make([]error, 0),
	}
}

func (h *HTML) Format(m report.Message) {
	h.messages = append(h.messages, m)
}

func (h *HTML) FormatError(err error) {
	h.errors = append(h.errors, err)
}

func (h *HTML) Close() {
	f, _ := Assets.Open("html.html")
	tpl, _ := ioutil.ReadAll(f)
	t, err := template.New("report").Parse(string(tpl))
	if err != nil {
		panic(err)
	}

	h.messages = h.sorter.SortByCreatedAt(h.messages)

	for i, m := range h.messages {
		// m.Message = string(markdown.Markdown([]byte(m.Message)))
		h.messages[i] = m
	}

	sinceString := h.since.Format("2006-01-02")
	nowString := time.Now().Format("2006-01-02")

	title := fmt.Sprintf(
		"Events between %s and %s",
		sinceString,
		nowString,
	)
	if sinceString == nowString {
		title = "Events in " + nowString
	}

	t.Execute(h.w, templateData{
		Title:    title,
		Messages: h.messages,
		Errors:   h.errors,
	})
}

type templateData struct {
	Title    string
	Messages []report.Message
	Errors   []error
}
