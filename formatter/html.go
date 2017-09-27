package formatter

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"text/template"
	"time"

	"github.com/lucassabreu/github-journaling-aggregator/report"
	"github.com/russross/blackfriday"
)

type templateData struct {
	Title    string
	Messages []report.Message
	Errors   []error
}

var issueWithRepoRegExp = regexp.MustCompile("[^ ]*/[^ ]#[0-9]*")
var funcs = template.FuncMap{
	"markdown": func(s string) string {
		return string(blackfriday.MarkdownCommon([]byte(s)))
	},
	"derefstr": func(s *string) string {
		return *s
	},
	"title": strings.Title,
	"date_to_str": func(t *time.Time) string {
		return t.In(time.Local).Format("2006-01-02 15:04:05")
	},
}

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
	tpl, _ := getAssetContent("/html.html")
	t := template.New("report")
	t.Funcs(funcs)
	t, err := t.Parse(string(tpl))
	if err != nil {
		panic(err)
	}

	h.messages = h.sorter.SortByCreatedAt(h.messages)

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

	err = t.Execute(h.w, templateData{
		Title:    title,
		Messages: h.messages,
		Errors:   h.errors,
	})

	if err != nil {
		panic(err)
	}
}
