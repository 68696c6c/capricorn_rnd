package golang

import (
	"fmt"
	"strings"
)

type Fields []Field

type Field struct {
	Name string
	Type IType // The Type is not composed like Var, Const, Function, and Struct because a struct Field has a type, but is not a type.  i.e. you can't do myVar := Field
	Tags Tags
}

func (f Field) Render() string {
	built := fmt.Sprintf(`%s %s %s`, f.Name, f.Type.GetReference(), string(f.Tags.Render()))
	result := strings.TrimSpace(built)
	return result
}

func (f Fields) Render() string {
	var builtValues []string
	for _, field := range f {
		valueString := field.Render()
		builtValues = append(builtValues, valueString)
	}
	if len(builtValues) == 0 {
		return ""
	}
	joinedValues := strings.Join(builtValues, "\n")
	result := strings.TrimSpace(joinedValues)
	return result
}
