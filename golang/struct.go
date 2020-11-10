package golang

import "github.com/68696c6c/capricorn_rnd/utils"

type Struct struct {
	Type
	*imports
	fields Fields
}

func NewStructFromType(t Type) *Struct {
	return &Struct{
		Type:    t,
		imports: newImports(),
		fields:  Fields{},
	}
}

func (s *Struct) AddField(f Field) {
	s.fields = append(s.fields, f)
}

func (s *Struct) GetStructFields() []Field {
	return s.fields
}

func (s *Struct) GetImports() imports {
	return *s.imports
}

func (s Struct) Render() string {
	var template = `type {{ .Name }} struct {
	{{- range $key, $value := .Fields }}
	{{ $value.Render }}
	{{- end }}
}`
	result, err := utils.ParseTemplate("template_struct", template, struct {
		Name   string
		Fields utils.Renderable
	}{
		Name:   s.Name,
		Fields: s.fields,
	})
	if err != nil {
		panic(err)
	}
	return result
}
