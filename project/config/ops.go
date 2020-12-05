package config

type OpsMeta struct {
	ImportEnums      string
	ImportMigrations string
	ServiceNameApp   string
	ServiceNameDb    string
	AppBinaryName    string
}

type OpsOptions struct {
	ServiceNameApp NameTemplate
	ServiceNameDb  NameTemplate
	AppBinaryName  NameTemplate
}

// @TODO: need to merge the YAML values with a set of defaults so that users don't need to include the ops section in the spec.
type Ops struct {
	Workdir      string   `yaml:"workdir,omitempty"`
	AppHTTPAlias string   `yaml:"app_http_alias,omitempty"`
	MainDatabase Database `yaml:"database,omitempty"`
}

type Database struct {
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	Name     string `yaml:"name,omitempty"`
	Debug    string `yaml:"debug,omitempty"`
}
