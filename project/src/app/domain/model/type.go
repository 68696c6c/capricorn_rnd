package model

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Type struct {
	*golang.Struct
	*fields
	hardDelete bool
}

func newModel(fileName string, hardDelete bool) *Type {
	typeName := utils.Pascal(fileName)
	return &Type{
		Struct:     golang.NewStruct(typeName, false, false),
		fields:     newFields(),
		hardDelete: hardDelete,
	}
}

func (m *Type) buildFields() {
	for _, f := range m.modelFields {
		m.AddField(f)
	}
}

// Determine the base model type, compose it into our type, and register the base model fields.
func (m *Type) addBaseFields() {
	modelType := goat.MakeSoftModelStruct()
	if m.hardDelete {
		modelType = goat.MakeHardModelStruct()
	}

	// This "field" is the composition of the base model struct type into this struct type.
	m.AddModelField(golang.NewField("", modelType, false))
	m.AddImportsVendor(goat.ImportGoat)

	// The fields that this struct receives from the base model are not declared on this struct, but they still need
	// database fields made for them.
	for _, f := range modelType.GetStructFields() {
		m.AddDbField(f)
	}
}

func (m *Type) addUserDefinedField(enums *enum.Enums, f Field) {
	fType := f.Type
	if fType == "email" {
		fType = "string"
	}
	fieldType := golang.MockTypeFromReference(fType)
	eType, isEnum := enums.GetEnumType(fType)
	if isEnum {
		fieldType = eType
		m.AddImportsApp(eType.GetImport())
	}
	field := goat.MakeModelField(f.Name, fieldType, true, f.Required, false)
	m.AddAllField(field)
}

func (m *Type) addBelongsToIdField(relation string) {
	name := utils.Pascal(utils.Singular(relation) + "_id")
	field := goat.MakeModelField(name, goat.MakeIdType(), true, true, false)
	m.AddAllField(field)
}

// Can't use relModel.GetName() to name the field because in a DDD app the name will always be "Model".
func (m *Type) addBelongsToTargetField(relName string, relModel golang.IType) {
	fieldName := utils.Pascal(utils.Singular(relName))
	field := goat.MakeModelField(fieldName, relModel, true, false, true)
	m.AddModelField(field)
	m.AddImportsApp(relModel.GetImport())
}

// Can't use relModel.GetName() to name the field because in a DDD app the name will always be "Model".
func (m *Type) addHasManyField(relName string, relModel golang.IType) {
	fieldName := utils.Pascal(utils.Plural(relName))
	relSlice := golang.MakeSliceType(false, relModel)
	field := goat.MakeModelField(fieldName, relSlice, true, false, true)
	m.AddModelField(field)
	m.AddImportsApp(relModel.GetImport())
}
