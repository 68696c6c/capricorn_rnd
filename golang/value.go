package golang

import (
	"fmt"
	"strings"
)

// Represents an argument or return value.
type Value struct {
	TypeRef string
	Name    string
}

func getJoinedValueString(values []Value) string {
	var builtValues []string
	for _, v := range values {
		builtValues = append(builtValues, fmt.Sprintf("%s %s", v.Name, v.TypeRef))
	}
	joinedValues := strings.Join(builtValues, ", ")
	return strings.TrimSpace(joinedValues)
}
