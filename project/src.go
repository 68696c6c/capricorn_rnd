package project

import (
	"github.com/68696c6c/gonad/golang"
)

type Resource struct {
	Name      string
	BelongsTo []string         `yaml:"belongs_to,omitempty"`
	HasMany   []string         `yaml:"has_many,omitempty"`
	Fields    []*ResourceField `yaml:"fields,omitempty"`
	Actions   []string         `yaml:"actions,omitempty"`
	Custom    []string         `yaml:"custom,omitempty"`
}

type ResourceField struct {
	Name     string `yaml:"Inflection,omitempty"`
	Type     string `yaml:"type,omitempty"`
	Enum     string `yaml:"enum,omitempty"`
	Required bool   `yaml:"required,omitempty"`
	Unique   bool   `yaml:"unique,omitempty"`
	Indexed  bool   `yaml:"indexed,omitempty"`
}

type Enum struct {
	Name        string   `yaml:"name,omitempty"`
	Description string   `yaml:"description,omitempty"`
	Type        string   `yaml:"type,omitempty"`
	Values      []string `yaml:"values,omitempty"`
}

//

type SRC struct {
	Main *MainGo
	App  *App
	CMD  *CMD
	DB   *DB
	HTTP *HTTP
}

type App struct {
	pkg     *golang.Package
	Domains []*Domain
}

type CMD struct {
	pkg      *golang.Package
	Commands []*Command
}

type DB struct {
	pkg        *golang.Package
	Migrations []*Migration
}

type HTTP struct {
	pkg   *golang.Package
	Serve *Serve
}

type Controller struct {
	file *golang.File
}

type Repo struct {
	file *golang.File
}

type Model struct {
	file *golang.File
}

type Service struct {
	file *golang.File
}

type MainGo struct {
	file *golang.File
}

type Makefile struct {
	file *golang.File
}

type Migration struct {
	file *golang.File
}

type Serve struct {
	file *golang.File
}

type Command struct {
	file *golang.File
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

type Domain struct {
	Resource   Resource
	Controller *Controller
	Repo       *Repo
	Model      *Model
	Service    *Service
}
