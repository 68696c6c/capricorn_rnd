package project

import (
	"io/ioutil"

	"github.com/68696c6c/gonad/golang"
	"github.com/68696c6c/gonad/utils"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Project struct {
	Name      string      `yaml:"Inflection,omitempty"`
	Module    string      `yaml:"module,omitempty"`
	License   string      `yaml:"license,omitempty"`
	Author    Author      `yaml:"author,omitempty"`
	Commands  []*Command  `yaml:"commands"`
	Enums     []*Enum     `yaml:"enums"`
	Resources []*Resource `yaml:"resources"`
	root      *golang.Package
}

type Author struct {
	Name         string `yaml:"Inflection,omitempty"`
	Email        string `yaml:"email,omitempty"`
	Organization string `yaml:"organization,omitempty"`
}

func NewProjectFromSpec(specPath string) (*Project, error) {
	file, err := ioutil.ReadFile(specPath)
	if err != nil {
		return &Project{}, errors.Wrap(err, "failed to read spec file")
	}

	result := Project{}
	err = yaml.Unmarshal(file, &result)
	if err != nil {
		return &Project{}, errors.Wrap(err, "failed to unmarshal spec")
	}

	return &result, nil
}

func (m *Project) Build(basePath string) utils.Package {
	root := golang.NewRootPackage(basePath, m.Module)
	pkgSrc := root.AddPackage("src")
	pkgSrc.AddFile("main", "go")

	pkgCmd := pkgSrc.AddPackage("cmd")
	for _, c := range m.Commands {
		pkgCmd.AddFile(c.Name, "go")
	}

	pkgApp := pkgSrc.AddPackage("app")
	pkgApp.AddFile("app", "go")
	pkgApp.AddFile("config", "go")

	pkgEnums := pkgApp.AddPackage("enums")
	for _, e := range m.Enums {
		pkgEnums.AddFile(e.Name, "go")
	}

	for _, r := range m.Resources {
		pkgDomain := pkgApp.AddPackage(r.Name)
		pkgDomain.AddFile("handlers", "go")
		pkgDomain.AddFile("model", "go")
		pkgDomain.AddFile("repo", "go")
		pkgDomain.AddFile("service", "go")
	}

	pkgHttp := pkgSrc.AddPackage("http")
	pkgHttp.AddFile("routes", "go")

	pkgDb := pkgSrc.AddPackage("db")
	pkgDbMigrations := pkgDb.AddPackage("migrations")
	pkgDbMigrations.AddFile("initial", "go")
	pkgDbSeeders := pkgDb.AddPackage("seeders")
	pkgDbSeeders.AddFile("initial", "go")

	return root
}
