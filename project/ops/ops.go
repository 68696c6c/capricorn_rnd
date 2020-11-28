package ops

import (
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/ops/local"
	"github.com/68696c6c/capricorn_rnd/utils"
)

func Build(root utils.Directory, c config.Ops) {
	rootPath := root.GetFullPath()
	root.AddRenderableFile(local.NewAppEnv(rootPath, c, false))
	root.AddRenderableFile(local.NewAppEnv(rootPath, c, true))
	root.AddRenderableFile(local.NewMakefile(rootPath, c))
	root.AddRenderableFile(local.NewDockerfile(rootPath, c))
	root.AddRenderableFile(local.NewDockerCompose(rootPath, c))
	root.AddRenderableFile(local.NewGitIgnore(rootPath, c))
}
