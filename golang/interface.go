package golang

import "github.com/68696c6c/capricorn_rnd/utils"

type Interface struct {
	Type
	*imports
	functions Functions
}

func NewInterfaceFromType(t Type) *Interface {
	return &Interface{
		Type:    t,
		imports: newImports(),
	}
}

func (s *Interface) AddFunction(f *Function) {
	s.imports = mergeImports(*s.imports, f.GetImports())
	s.functions = append(s.functions, f)
}

func (s *Interface) GetImports() imports {
	return *s.imports
}

func (s *Interface) Render() string {
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
		Name:      s.Name,
		Functions: s.functions,
	})
	if err != nil {
		panic(err)
	}
	return result
}
