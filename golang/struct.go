package golang

import (
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Struct struct {
	*Type
	fields      Fields
	constructor *Function
}

// Use this for generating structs.
func NewStruct(typeName string, isPointer, isSlice bool) *Struct {
	return StructFromType(NewType(typeName, isPointer, isSlice))
}

// Use this for mocking builtin and vendor structs.
func MockStruct(importPath, typeName string, isPointer, isSlice bool) *Struct {
	return StructFromType(MockType(importPath, typeName, isPointer, isSlice))
}

func StructFromType(t *Type) *Struct {
	return &Struct{
		Type:   t,
		fields: Fields{},
	}
}

func (s *Struct) GetType() *Type {
	return s.Type
}

func (s *Struct) CopyType() *Type {
	return copyType(s.Type)
}

func (s *Struct) AddField(f *Field) {
	removePackageRefIfSamePackage(s.GetPackage(), f.Type)
	s.fields = append(s.fields, f)
}

func (s *Struct) AddConstructor(f *Function) {
	s.constructor = f
	s.imports = mergeImports(*s.imports, f.getImports())
	s.functions = append(s.functions, f)
}

func (s *Struct) GetConstructor() *Function {
	return s.constructor
}

func (s *Struct) GetStructFields() Fields {
	return s.fields
}

func (s *Struct) getImports() imports {
	return *s.imports
}

func (s *Struct) Render() string {
	var template = `
type {{ .Name }} struct {
	{{- range $key, $value := .Fields }}
	{{ $value.Render }}
	{{- end }}
}

{{ .Functions.Render }}
`
	result, err := utils.ParseTemplate("template_struct", template, struct {
		Name      string
		Fields    utils.Renderable
		Functions utils.Renderable
	}{
		Name:      s.Name,
		Fields:    s.fields,
		Functions: s.functions,
	})
	if err != nil {
		panic(err)
	}
	return result
}
