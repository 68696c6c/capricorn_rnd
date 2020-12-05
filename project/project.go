package project

import (
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/ops"
	"github.com/68696c6c/capricorn_rnd/project/src"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func Build(p *config.Project, o config.ProjectOptions) utils.Directory {
	projectDir := utils.NewFolder(o.BasePath, utils.Snake(p.Name))

	projectSrc := src.Build(projectDir, p, o.Src)

	ops.Build(projectDir, p.Ops, config.OpsMeta{
		ImportEnums:      projectSrc.GetApp().GetEnums().GetImport(),
		ImportMigrations: projectSrc.GetDb().GetImportMigrations(false),
		AppBinaryName:    o.Ops.ServiceNameApp.Parse(p.Name),
		ServiceNameApp:   o.Ops.ServiceNameApp.Parse(p.Name),
		ServiceNameDb:    o.Ops.ServiceNameDb.Parse(p.Name),
	})

	return projectDir
}
