package model

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Model struct {
	*golang.File `yaml:"-"`
	*fields      `yaml:"-"`
	Name         string   `yaml:"name,omitempty"`
	Delete       string   `yaml:"delete,omitempty"`
	BelongsTo    []string `yaml:"belongs_to,omitempty"`
	HasMany      []string `yaml:"has_many,omitempty"`
	Fields       []Field  `yaml:"fields,omitempty"`
	Actions      []string `yaml:"actions,omitempty"`
	Custom       []string `yaml:"custom,omitempty"`
	built        bool     `yaml:"-"`
	fileName     string   `yaml:"-"`
}

func (m *Model) Build(pkg *golang.Package, enums *enum.Enums, fileName string) {
	if m.built {
		return
	}

	m.File = pkg.AddGoFile(fileName)
	m.fileName = fileName
	m.fields = newFields()

	// Build the base model fields.
	m.AddBaseFields()

	// Build the foreign ID fields for the Belongs-To relations.
	for _, relation := range m.BelongsTo {
		m.AddBelongsToIdField(relation)
	}

	// Build the user-defined fields.
	for _, f := range m.Fields {
		m.AddUserDefinedField(enums, f)
	}

	// Build the Belongs-To fields that GORM will hydrate the relation in to.
	for _, relation := range m.BelongsTo {
		m.AddBelongsToTargetField(relation)
	}

	// Build the Has-Many fields.
	for _, relation := range m.HasMany {
		m.AddHasManyField(relation)
	}

	// Build the struct using the accumulated fields.
	m.AddModelStruct()

	m.built = true
}

// Determine the base model type, compose it into our type, and register the base model fields.
func (m *Model) AddBaseFields() {
	modelType := getModelSoftDelete()
	if m.Delete == "hard" {
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

func (m *Model) AddUserDefinedField(enums *enum.Enums, f Field) {
	fieldType := golang.NewTypeFromReference(f.Type)
	eType, isEnum := enums.GetEnumType(f.Type)
	if isEnum {
		fieldType = eType
		m.AddImportsApp(eType.GetImport())
	}
	m.AddAllField(makeField(f.Name, fieldType, true))
}

func (m *Model) AddBelongsToIdField(relation string) {
	name := utils.Pascal(utils.Singular(relation) + "_id")
	m.AddAllField(makeField(name, getIdType(), true))
}

func (m *Model) AddBelongsToTargetField(relation string) {
	name := utils.Pascal(utils.Plural(relation))
	relModel := getAssumedModelType(m.PKG.GetBaseImport(), relation, true)
	m.AddImportsApp(relModel.GetImport())
	m.AddModelField(makeField(name, relModel, true))
}

func (m *Model) AddHasManyField(relation string) {
	name := utils.Pascal(relation)
	relModel := getAssumedModelType(m.PKG.GetBaseImport(), relation, true)
	m.AddImportsApp(relModel.GetImport())
	relSlice := golang.MakeSliceType(false, relModel)
	m.AddModelField(makeField(name, relSlice, true))
}

func (m *Model) AddModelStruct() {
	result := golang.NewStructFromType(golang.Type{
		Import:    m.PKG.GetImport(),
		Package:   m.PKG.GetName(),
		Name:      utils.Pascal(m.fileName),
		IsPointer: false,
		IsSlice:   false,
	})
	for _, f := range m.modelFields {
		result.AddField(f)
	}
	m.AddStruct(result)
}
