package golang

import (
	"fmt"
	"strings"
)

type Tags []Tag

type Tag struct {
	Key    string
	Values []string
}

func (t Tag) getValues() string {
	var builtValues []string
	for _, v := range t.Values {
		builtValues = append(builtValues, v)
	}
	joinedValues := strings.Join(builtValues, ",")
	return strings.TrimSpace(joinedValues)
}

func (t Tag) Render() []byte {
	result := fmt.Sprintf(`%s:"%s"`, t.Key, t.getValues())
	return []byte(result)
}

func (t Tags) Render() []byte {
	var builtValues []string
	for _, tag := range t {
		valueString := string(tag.Render())
		builtValues = append(builtValues, valueString)
	}
	if len(builtValues) == 0 {
		return []byte("")
	}
	joinedValues := strings.TrimSpace(strings.Join(builtValues, " "))
	result := fmt.Sprintf("`%s`", joinedValues)
	return []byte(result)
}
