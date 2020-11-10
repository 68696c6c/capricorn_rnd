package golang

import "github.com/68696c6c/capricorn_rnd/utils"

type Struct struct {
	Type
	fields Fields
}

func (s *Struct) AddField(f Field) {
	s.fields = append(s.fields, f)
}

func (s *Struct) GetStructFields() []Field {
	return s.fields
}

func (s Struct) Render() []byte {
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

// func hmm() Struct {
// 	userSiteModel := Struct{
// 		Type: Type{
// 			Import:    "github.com/68696c6c/src/app/user_sites",
// 			Package:   "user_sites",
// 			Name:      "Model",
// 			IsPointer: true,
// 			IsSlice:   false,
// 		},
// 	}
// 	userModel := Struct{
// 		Type: Type{
// 			Import:    "github.com/68696c6c/src/app/users",
// 			Package:   "users",
// 			Name:      "Model",
// 			IsPointer: false,
// 			IsSlice:   false,
// 		},
// 		fields: []*Field{
// 			{
// 				Type: GetModelSoftDelete(),
// 			},
// 			makeField("email", getStringType(), true),
// 			makeField("first_name", getStringType(), true),
// 			makeField("last_name", getStringType(), true),
// 			{
// 				Name: "UserSites",
// 				Type: MakeSliceType(false, userSiteModel),
// 				Tags: nil,
// 			},
// 		},
// 	}
// 	return userModel
// }
