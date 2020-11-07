package ops

import (
	"github.com/68696c6c/capricorn_rnd/project/ops/local"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Ops struct {
	appEnv        local.AppEnv
	appEnvExample local.AppEnv
	makefile      local.Makefile
	dockerfile    local.Dockerfile
	dockerCompose local.DockerCompose
	gitIgnore     local.GitIgnore
	rootPath      string
}

func Build(root *utils.Folder, c local.Config) {
	rootPath := root.GetPath()
	root.AddRenderableFile(local.NewAppEnv(rootPath, c, false))
	root.AddRenderableFile(local.NewAppEnv(rootPath, c, true))
	root.AddRenderableFile(local.NewMakefile(rootPath, c))
	root.AddRenderableFile(local.NewDockerfile(rootPath, c))
	root.AddRenderableFile(local.NewDockerCompose(rootPath, c))
	root.AddRenderableFile(local.NewGitIgnore(rootPath, c))
}
