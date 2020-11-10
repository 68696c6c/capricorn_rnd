package model

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
)

type Model struct {
	*golang.File `yaml:"-"`
	Name         string   `yaml:"name,omitempty"`
	Delete       string   `yaml:"delete,omitempty"`
	BelongsTo    []string `yaml:"belongs_to,omitempty"`
	HasMany      []string `yaml:"has_many,omitempty"`
	Fields       []Field  `yaml:"fields,omitempty"`
	Actions      []Action `yaml:"actions,omitempty"`
	Custom       []string `yaml:"custom,omitempty"`
	modelType    *Type    `yaml:"-"`
}

func (m *Model) Build(pkg *golang.Package, enums *enum.Enums, fileName string) Type {
	if m.modelType != nil {
		return *m.modelType
	}

	m.File = pkg.AddGoFile(fileName)
	model := newModel(m.PKG.GetBaseImport(), pkg.GetName(), fileName, m.Delete == "hard")

	// Build the base model fields.
	model.addBaseFields()

	// Build the foreign ID fields for the Belongs-To relations.
	for _, relation := range m.BelongsTo {
		model.addBelongsToIdField(relation)
	}

	// Build the user-defined fields.
	for _, f := range m.Fields {
		model.addUserDefinedField(enums, f)
	}

	// Build the Belongs-To fields that GORM will hydrate the relation in to.
	for _, relation := range m.BelongsTo {
		model.addBelongsToTargetField(relation)
	}

	// Build the Has-Many fields.
	for _, relation := range m.HasMany {
		model.addHasManyField(relation)
	}

	// Build the struct using the accumulated fields.
	model.buildFields()

	m.AddStruct(model.Struct)
	m.modelType = model

	return *m.modelType
}
