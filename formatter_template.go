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

var DefaultTemplateFormatter = &TemplateFormatter{Template: template.Must(template.New("default").Parse(
	`=[{{.Magenta}}{{.Bold}}{{.Host}}{{.Reset}}]=[{{.LevelColor}}{{.Level}}{{.Reset}}]=[{{.Yellow}}{{.TimeRFC3339}}{{.Reset}}]=> {{.Blue}}{{.Bold}}message{{.Reset}}={{.Message}}{{range $key, $value := .Extras}}  {{$.Blue}}{{$.Bold}}{{$key}}{{$.Reset}}={{$value}}{{end}}`))}
