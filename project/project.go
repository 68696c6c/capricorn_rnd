package project

import (
	"io/ioutil"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/ops"
	"github.com/68696c6c/capricorn_rnd/project/ops/local"
	"github.com/68696c6c/capricorn_rnd/project/src"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/project/src/cmd"
	"github.com/68696c6c/capricorn_rnd/utils"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type Project struct {
	Name      string        `yaml:"name,omitempty"`
	Module    string        `yaml:"module,omitempty"`
	License   string        `yaml:"license,omitempty"`
	Author    Author        `yaml:"author,omitempty"`
	Ops       local.Config  `yaml:"ops"`
	Commands  []cmd.Command `yaml:"commands"`
	Enums     []enum.Enum   `yaml:"enums"`
	Resources []model.Model `yaml:"resources"`
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

func (p *Project) Build(basePath string) utils.Directory {
	projectDir := utils.NewFolder(basePath, utils.Snake(p.Name))

	ops.Build(projectDir, p.Ops)

	src.Build(projectDir, src.Meta{
		// BasePath:  projectDir.GetPath(),
		Module:    p.Module,
		Commands:  p.Commands,
		Enums:     p.Enums,
		Resources: p.Resources,
	})

	// src.NewSRC(projectDir, src.Meta{
	// 	// BasePath:  projectDir.GetPath(),
	// 	Module:    p.Module,
	// 	Commands:  p.Commands,
	// 	Enums:     p.Enums,
	// 	Resources: p.Resources,
	// })
	// projectDir.AddDirectory(srcPkg.GetDirectory())

	return projectDir
}
