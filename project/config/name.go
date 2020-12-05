package config

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/utils"
)

const (
	nameTemplateResourceSingular = ":singular:"
	nameTemplateResourcePlural   = ":plural:"
	nameTemplateAppName          = ":app:"
)

type NameTemplate string

func NameTemplateF(format string, a ...interface{}) NameTemplate {
	return NameTemplate(fmt.Sprintf(format, a...))
}

func (n NameTemplate) Parse(resourceName string) string {
	template := string(n)
	if strings.Contains(template, nameTemplateResourceSingular) {
		return strings.Replace(template, nameTemplateResourceSingular, utils.Singular(resourceName), -1)
	}
	if strings.Contains(template, nameTemplateResourcePlural) {
		return strings.Replace(template, nameTemplateResourcePlural, utils.Plural(resourceName), -1)
	}
	if strings.Contains(template, nameTemplateAppName) {
		return strings.Replace(template, nameTemplateAppName, utils.Snake(resourceName), -1)
	}
	return template
}
