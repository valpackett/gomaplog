package gomaplog

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

type TemplateFormatter struct {
	Template *template.Template
}

func (formatter *TemplateFormatter) Format(event LogEvent) ([]byte, error) {
	for k, v := range event.Extras {
		event.Extras[k] = fmt.Sprintf("%v", v)
	}
	var buf bytes.Buffer
	err := formatter.Template.Execute(&buf, &event)
	if err != nil {
		return []byte{}, err
	} else {
		return buf.Bytes(), nil
	}
}

func Collapse(str string) string {
	s := str
	s = strings.Join(strings.Split(s, "\n"), " ")
	s = strings.Join(strings.Split(s, "  "), " ")
	return s
}

var DefaultTemplateFormatter = &TemplateFormatter{Template: template.Must(template.New("default").Funcs(template.FuncMap{"Collapse": Collapse}).Parse(
	`=[{{.Magenta}}{{.Bold}}{{Collapse .Host}}{{.Reset}}]=[{{.LevelColor}}{{.Level}}{{.Reset}}]=[{{.Yellow}}{{.TimeRFC3339}}{{.Reset}}]=> {{.Blue}}{{.Bold}}message{{.Reset}}={{Collapse .Message}}{{range $key, $value := .Extras}}  {{$.Blue}}{{$.Bold}}{{Collapse $key}}{{$.Reset}}={{Collapse $value}}{{end}}`))}
