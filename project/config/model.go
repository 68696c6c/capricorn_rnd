package config

type Model struct {
	Name      string   `yaml:"name,omitempty"`
	Delete    string   `yaml:"delete,omitempty"`
	BelongsTo []string `yaml:"belongs_to,omitempty"`
	HasMany   []string `yaml:"has_many,omitempty"`
	Fields    []Field  `yaml:"fields,omitempty"`
	Actions   []Action `yaml:"actions,omitempty"`
	Custom    []string `yaml:"custom,omitempty"`
}

type Field struct {
	Name     string `yaml:"name,omitempty"`
	Type     string `yaml:"type,omitempty"`
	Enum     string `yaml:"enum,omitempty"`
	Required bool   `yaml:"required,omitempty"`
	Unique   bool   `yaml:"unique,omitempty"`
	Indexed  bool   `yaml:"indexed,omitempty"`
}

type ModelOptions struct {
	FileNameTemplate NameTemplate
	TypeNameTemplate NameTemplate
}
