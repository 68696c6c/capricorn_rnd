package golang

import (
	"strings"

	"github.com/68696c6c/capricorn_rnd/utils"
)

type Struct struct {
	Type
	*imports
	fields    Fields
	functions Functions
	receiver  Value
}

func NewStruct(typeName string, isPointer, isSlice bool) *Struct {
	return newStructFromType(NewType(typeName, isPointer, isSlice))
}

func NewStructMock(importPath, typeName string, isPointer, isSlice bool) *Struct {
	return newStructFromType(NewTypeMock(importPath, typeName, isPointer, isSlice))
}

func newStructFromType(t Type) *Struct {
	typeName := t.GetName()
	return &Struct{
		Type:    t,
		imports: newImports(),
		fields:  Fields{},
		receiver: Value{
			TypeRef: typeName,
			Name:    strings.ToLower(typeName[0:1]),
		},
	}
}

func (s *Struct) AddField(f *Field) {
	s.fields = append(s.fields, f)
}

func (s *Struct) AddConstructor(f *Function) {
	s.imports = mergeImports(*s.imports, f.GetImports())
	s.functions = append(s.functions, f)
}

func (s *Struct) AddFunction(f *Function) {
	f.SetReceiver(s.receiver)
	s.imports = mergeImports(*s.imports, f.GetImports())
	s.functions = append(s.functions, f)
}

func (s *Struct) GetStructFields() Fields {
	return s.fields
}

func (s *Struct) GetImports() imports {
	return *s.imports
}

func (s *Struct) GetReceiverName() string {
	return s.receiver.Name
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

func MakeHardModelStruct() *Struct {
	result := newStructFromType(makeBaseModelType())
	result.AddField(MakeModelField("id", MakeIdType(), true, false, true))
	result.AddField(MakeModelField("created_at", MakeTimeType(false), true, false, true))
	result.AddField(MakeModelField("updated_at", MakeTimeType(true), true, false, true))
	return result
}

func MakeSoftModelStruct() *Struct {
	result := MakeHardModelStruct()
	result.AddField(MakeModelField("deleted_at", MakeTimeType(true), true, false, true))
	return result
}
