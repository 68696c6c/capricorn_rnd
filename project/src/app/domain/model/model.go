package model

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Model struct {
	*golang.Struct
	*fields
	hardDelete bool
}

func Build(pkg golang.IPackage, o config.ModelOptions, resource config.Model, enums *enum.Enums) *Model {
	fileName := o.FileNameTemplate.Parse(resource.Name)
	file := pkg.AddGoFile(fileName)
	pkgImport := pkg.GetImport()

	name := o.TypeNameTemplate.Parse(resource.Name)
	model := newModel(name, resource.Delete == "hard")
	file.AddStruct(model.Struct)

	// Build the base model fields.
	model.addBaseFields()

	// Build the foreign ID fields for the Belongs-To relations.
	for _, relation := range resource.BelongsTo {
		model.addBelongsToIdField(relation)
	}

	// Build the user-defined fields.
	for _, f := range resource.Fields {
		model.addUserDefinedField(enums, f)
	}

	// Build the Belongs-To fields that GORM will hydrate the relation in to.
	for _, relation := range resource.BelongsTo {
		relModel := getAssumedModelType(pkgImport, relation, true)
		model.addBelongsToTargetField(relation, relModel)
	}

	// Build the Has-Many fields.
	for _, relation := range resource.HasMany {
		relModel := getAssumedModelType(pkgImport, relation, true)
		model.addHasManyField(relation, relModel)
	}

	// Build the struct using the accumulated fields.
	model.buildFields()

	return model
}

func newModel(fileName string, hardDelete bool) *Model {
	typeName := utils.Pascal(fileName)
	return &Model{
		Struct:     golang.NewStruct(typeName, false, false),
		fields:     newFields(),
		hardDelete: hardDelete,
	}
}

func (m *Model) buildFields() {
	for _, f := range m.modelFields {
		m.AddField(f)
	}
}

// Determine the base model type, compose it into our type, and register the base model fields.
func (m *Model) addBaseFields() {
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

func (m *Model) addUserDefinedField(enums *enum.Enums, f config.Field) {
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

func (m *Model) addBelongsToIdField(relation string) {
	name := utils.Pascal(utils.Singular(relation) + "_id")
	field := goat.MakeModelField(name, goat.MakeTypeId(), true, true, false)
	m.AddAllField(field)
}

func (m *Model) addBelongsToTargetField(relName string, relModel golang.IType) {
	fieldName := utils.Pascal(utils.Singular(relName))
	field := goat.MakeModelField(fieldName, relModel, true, false, true)
	m.AddModelField(field)
}

func (m *Model) addHasManyField(relName string, relModel golang.IType) {
	fieldName := utils.Pascal(utils.Plural(relName))
	relSlice := golang.MakeSliceType(false, relModel)
	field := goat.MakeModelField(fieldName, relSlice, true, false, true)
	m.AddModelField(field)
}

func getAssumedModelType(imp, inputName string, isPointer bool) golang.IType {
	return golang.MockStruct(imp, utils.Pascal(utils.Singular(inputName)), isPointer, false)
}
