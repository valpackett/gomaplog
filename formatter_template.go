package gomaplog

import (
	"bytes"
	"text/template"
)

type TemplateFormatter struct {
	Template *template.Template
}

func (formatter *TemplateFormatter) Format(event LogEvent) ([]byte, error) {
	var buf bytes.Buffer
	err := formatter.Template.Execute(&buf, &event)
	if err != nil {
		return []byte{}, err
	} else {
		return buf.Bytes(), nil
	}
}

func DefaultTemplateFormatter() *TemplateFormatter {
	tpl := template.Must(template.New("default").Parse(
		`=[{{.Magenta}}{{.Bold}}{{.Host}}{{.Reset}}]=[{{.LevelColor}}{{.Level}}{{.Reset}}]=[{{.Yellow}}{{.TimeISO}}{{.Reset}}]=> {{.Blue}}{{.Bold}}message{{.Reset}}={{.Message}}{{range $key, $value := .Extras}}  {{$.Blue}}{{$.Bold}}{{$key}}{{$.Reset}}={{$value}}{{end}}`))
	return &TemplateFormatter{Template: tpl}
}
