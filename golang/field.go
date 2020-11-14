package golang

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/utils"
)

type Fields []*Field

type Field struct {
	Name          string
	Type          IType // The Type is not composed like Var, Const, Function, and Struct because a struct Field has a type, but is not a type.  i.e. you can't do myVar := Field
	Tags          Tags
	separatedName string
	isExported    bool
	isRequired    bool
}

// The only safe way to inflect names is to start from a separated string. Accepting a separated name and requiring the
// isExported arg makes it very explicit what kind of field this is and allows us to save the original separated name
// for future use when setting tags etc.
func NewField(separatedName string, t IType, isExported bool) *Field {
	name := utils.Camel(separatedName)
	if isExported {
		name = utils.Pascal(separatedName)
	}
	return &Field{
		Name:          name,
		Type:          t,
		separatedName: separatedName,
		isExported:    isExported,
		isRequired:    false,
	}
}

func (f *Field) SetRequired(isRequired bool) {
	if isRequired {
		f.isRequired = true
		f.AddTag("binding", []string{"required"})
	}
}

func (f *Field) SetJsonTag(omitEmpty bool) {
	values := []string{utils.Snake(f.separatedName)}
	if omitEmpty {
		values = append(values, "omitempty")
	}
	f.AddTag("json", values)
}

func (f *Field) AddTag(key string, values []string) {
	for _, t := range f.Tags {
		if key == t.Key {
			t.Values = append(t.Values, values...)
			return
		}
	}
	f.Tags = append(f.Tags, Tag{
		Key:    key,
		Values: values,
	})
}

func (f *Field) Render() string {
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
