package config

type Enum struct {
	Name        string   `yaml:"name,omitempty"`
	Description string   `yaml:"description,omitempty"`
	Type        string   `yaml:"type,omitempty"`
	Values      []string `yaml:"values,omitempty"`
}

type EnumOptions struct {
	PkgName    string
	FileName   string
	SpecPrefix string

	InputArgName string

	FromStringFuncNameSuffix string
	ScanFuncName             string
	ValueFuncName            string
}
