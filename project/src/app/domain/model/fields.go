package model

import "github.com/68696c6c/capricorn_rnd/golang"

type fields struct {
	allFields      golang.Fields `desc:"all fields, including base model fields, relations, and user defined fields"`
	modelFields    golang.Fields `desc:"fields that are defined on the model struct. includes relations but not base model fields"`
	databaseFields golang.Fields `desc:"fields that exist in the database.  includes base model fields but not relations"`
}

func newFields() *fields {
	return &fields{
		allFields:      golang.Fields{},
		modelFields:    golang.Fields{},
		databaseFields: golang.Fields{},
	}
}

func (f *fields) AddAllField(field *golang.Field) {
	f.allFields = append(f.allFields, field)
	f.modelFields = append(f.modelFields, field)
	f.databaseFields = append(f.databaseFields, field)
}

func (f *fields) AddModelField(field *golang.Field) {
	f.allFields = append(f.allFields, field)
	f.modelFields = append(f.modelFields, field)
}

func (f *fields) AddDbField(field *golang.Field) {
	f.allFields = append(f.allFields, field)
	f.databaseFields = append(f.databaseFields, field)
}
