package golang

import (
	"fmt"
	"strings"
)

// Represents an argument or return value.
type Value struct {
	*Type
	Name string
}

func ValueFromType(name string, t *Type) *Value {
	return &Value{
		Type: t,
		Name: name,
	}
}

func (v *Value) GetType() *Type {
	return v.Type
}

func (v *Value) CopyType() *Type {
	return copyType(v.Type)
}

func getJoinedValueString(values []*Value) string {
	var builtValues []string
	for _, v := range values {
		builtValues = append(builtValues, fmt.Sprintf("%s %s", v.Name, v.GetReference()))
	}
	joinedValues := strings.Join(builtValues, ", ")
	return strings.TrimSpace(joinedValues)
}
