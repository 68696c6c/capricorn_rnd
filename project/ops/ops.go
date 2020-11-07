package ops

import (
	"github.com/68696c6c/gonad/project/ops/local"
)

type Ops struct {
	AppEnv        local.AppEnv
	AppEnvExample local.AppEnv
	Makefile      local.Makefile
	Dockerfile    local.Dockerfile
	DockerCompose local.DockerCompose
	GitIgnore     local.GitIgnore
}

func NewOps(basePath string, c local.Config) Ops {
	return Ops{
		AppEnv:        local.NewAppEnv(basePath, c, false),
		AppEnvExample: local.NewAppEnv(basePath, c, true),
		Makefile:      local.NewMakefile(basePath, c),
		Dockerfile:    local.NewDockerfile(basePath, c),
		DockerCompose: local.NewDockerCompose(basePath, c),
		GitIgnore:     local.NewGitIgnore(basePath, c),
	}
}
