package local

type Config struct {
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
