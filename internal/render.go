package internal

import (
	"bytes"
	"strings"
	"text/template"
)

func Render(text string, data interface{}) (string, error) {
	tmpl, err := template.New("temp").Parse(text)
	if err != nil {
		return "", nil
	}
	var result bytes.Buffer
	err = tmpl.Execute(&result, data)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(result.String()), nil
}
