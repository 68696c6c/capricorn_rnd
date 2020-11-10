package utils

import (
	"bytes"
	"text/template"

	"github.com/pkg/errors"
)

func ParseTemplate(name, temp string, data interface{}) (string, error) {
	var tpl bytes.Buffer
	t := template.Must(template.New(name).Parse(temp))
	err := t.Execute(&tpl, data)
	if err != nil {
		return "", errors.Wrapf(err, "failed parse template '%s'", name)
	}
	return tpl.String(), nil
}
