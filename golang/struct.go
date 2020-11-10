package golang

import "github.com/68696c6c/capricorn_rnd/utils"

type Struct struct {
	Type
	*imports
	fields    Fields
	functions Functions
	receiver  Value
}

func NewStructFromType(t Type) *Struct {
	typeName := t.GetName()
	return &Struct{
		Type:    t,
		imports: newImports(),
		fields:  Fields{},
		receiver: Value{
			TypeRef: t.GetReference(),
			Name:    typeName[0:1],
		},
	}
}

func (s *Struct) AddField(f Field) {
	s.fields = append(s.fields, f)
}

func (s *Struct) AddFunction(f *Function) {
	f.SetReceiver(s.receiver)
	s.imports = mergeImports(*s.imports, f.GetImports())
	s.functions = append(s.functions, f)
}

func (s *Struct) GetStructFields() []Field {
	return s.fields
}

func (s *Struct) GetImports() imports {
	return *s.imports
}

func (s *Struct) SetReceiverName(name string) {
	s.receiver.Name = name
}

func (s *Struct) SetReceiverTypeRef(typeRef string) {
	s.receiver.TypeRef = typeRef
}

func (s *Struct) GetReceiver() Value {
	return s.receiver
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
