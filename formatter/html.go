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
	errors   []error
}

func NewHTML(w io.Writer) HTML {
	return HTML{
		w:        w,
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

	t.Execute(h.w, templateData{
		Title:    "Example",
		Messages: h.messages,
		Errors:   h.errors,
	})
}

type templateData struct {
	Title    string
	Messages []report.Message
	Errors   []error
}

const tpl = `<html>
	<head>
		<title>{{ .Title }}</title>
	</head>
	<body>
		<table>
			<thead>
				<th>Event Name</th>
				<th>Message</th>
			</thead>
			<tbody>
				{{ range .Messages }}<tr>
					<td>{{ .EventName }}</td>
					<td>{{ .Message }}</td>
				</ti>{{ end }}
			</tbody>
		</table>
	</body>
</html>`
