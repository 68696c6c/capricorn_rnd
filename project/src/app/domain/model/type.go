package model

import (
	"path"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Type struct {
	*golang.Struct
	*fields
	baseImport string
	hardDelete bool
}

func newModel(baseImport, pkgName, fileName string, hardDelete bool) *Type {
	return &Type{
		Struct: golang.NewStructFromType(golang.Type{
			Import:    path.Join(baseImport, pkgName),
			Package:   pkgName,
			Name:      utils.Pascal(fileName),
			IsPointer: false,
			IsSlice:   false,
		}),
		fields:     newFields(),
		baseImport: baseImport,
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
	modelType := getModelSoftDelete()
	if m.hardDelete {
		modelType = getModelHardDelete()
	}

	// This "field" is the composition of the base model struct type into this struct type.
	m.AddModelField(golang.Field{
		Type: modelType,
	})
	m.AddImportsVendor(golang.ImportGoat)

	// The fields that this struct receives from the base model are not declared on this struct, but they still need
	// database fields made for them.
	for _, f := range modelType.GetStructFields() {
		m.AddDbField(f)
	}
}

func (m *Type) addUserDefinedField(enums *enum.Enums, f Field) {
	fieldType := golang.NewTypeFromReference(f.Type)
	eType, isEnum := enums.GetEnumType(f.Type)
	if isEnum {
		fieldType = eType
		m.AddImportsApp(eType.GetImport())
	}
	m.AddAllField(makeField(f.Name, fieldType, true))
}

func (m *Type) addBelongsToIdField(relation string) {
	name := utils.Pascal(utils.Singular(relation) + "_id")
	m.AddAllField(makeField(name, golang.MakeIdType(), true))
}

func (m *Type) addBelongsToTargetField(relation string) {
	name := utils.Pascal(utils.Plural(relation))
	relModel := getAssumedModelType(m.baseImport, relation, true)
	m.AddImportsApp(relModel.GetImport())
	m.AddModelField(makeField(name, relModel, true))
}

func (m *Type) addHasManyField(relation string) {
	name := utils.Pascal(relation)
	relModel := getAssumedModelType(m.baseImport, relation, true)
	m.AddImportsApp(relModel.GetImport())
	relSlice := golang.MakeSliceType(false, relModel)
	m.AddModelField(makeField(name, relSlice, true))
}
