package project

import (
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/ops"
	"github.com/68696c6c/capricorn_rnd/project/src"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func Build(p *config.Project, o config.ProjectOptions) utils.Directory {
	projectDir := utils.NewFolder(o.BasePath, utils.Snake(p.Name))

	ops.Build(projectDir, p.Ops)

	src.Build(projectDir, p, o.Src)

	return projectDir
}
