package formatter

import (
	"fmt"
	"html/template"
	"io"
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
	t, err := template.New("report").Parse(tpl)
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

const tpl = `
<html>
<head>
  <title>{{ .Title }}</title>
  <link rel="stylesheet" href="https://fonts.googleapis.com/icon?family=Material+Icons">
  <link rel="stylesheet" href="https://code.getmdl.io/1.3.0/material.indigo-pink.min.css">
  <style>
    table td.message {
      white-space: normal;
    }
  </style>
</head>
<body>
<div class="mdl-layout mdl-layout--fixed-header mdl-js-layout mdl-color--grey-100">
  <header class="mdl-layout__header mdl-layout__header--scroll mdl-color--grey-100 mdl-color-text--grey-800">
    <div class="mdl-layout__header-row">
      <span class="mdl-layout-title">{{ .Title }}</span>
    </div>
  </header>
  <main class="mdl-layout__content">
    <ul class="mdl-list mdl-cell mdl-cell--12-col">
      {{ range .Errors }}
        <li class="mdl-list__item">{{ .Error }}</li>
      {{ end }}
    </ul>

    <table class="mdl-data-table mdl-js-data-table mdl-shadow--2dp mdl-cell mdl-cell--12-col">
      <thead>
        <th class="mdl-data-table__cell--non-numeric">Event Name</th>
        <th class="mdl-data-table__cell--non-numeric">Repo Name</th>
        <th class="mdl-data-table__cell--non-numeric">Message</th>
      </thead>
      <tbody>
        {{ range .Messages }}
          <tr>
            <td class="mdl-data-table__cell--non-numeric">{{ .EventName }}</td>
            <td class="mdl-data-table__cell--non-numeric">{{ .Repo.Name }}</td>
            <td class="mdl-data-table__cell--non-numeric message">{{ .Message }}</td>
          </tr>
        {{ end }}
      </tbody>
    </table>
  </main>
</div>
<script defer src="https://code.getmdl.io/1.3.0/material.min.js"></script>
</body>
</html>
`
