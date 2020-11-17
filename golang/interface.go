package golang

import "github.com/68696c6c/capricorn_rnd/utils"

type Interface struct {
	*Type
}

func NewInterface(typeName string, isPointer, isSlice bool) *Interface {
	return &Interface{
		Type: NewType(typeName, isPointer, isSlice),
	}
}

func (i *Interface) GetType() *Type {
	return i.Type
}

func (i *Interface) AddFunction(f *Function) {
	i.imports = mergeImports(*i.imports, f.getImports())
	i.functions = append(i.functions, f)
}

func (i *Interface) Render() string {
	var template = `
type {{ .Name }} interface {
	{{- range $key, $value := .Functions }}
	{{ $value.GetSignature }}
	{{- end }}
}`
	result, err := utils.ParseTemplate("template_interface", template, struct {
		Name      string
		Functions utils.Renderable
	}{
		Name:      i.Name,
		Functions: i.functions,
	})
	if err != nil {
		panic(err)
	}
	return result
}
