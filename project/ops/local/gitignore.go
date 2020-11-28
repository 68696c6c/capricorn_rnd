package local

import (
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/utils"
)

const gitignoreTemplate = `.DS_Store
.idea
src/vendor
.app.env
`

type GitIgnore struct {
	*utils.File
	data config.Ops
}

func NewGitIgnore(basePath string, c config.Ops) GitIgnore {
	file := utils.NewFile(basePath, ".gitignore", "")
	return GitIgnore{
		File: file,
		data: c,
	}
}

func (g GitIgnore) Render() string {
	result, err := utils.ParseTemplate(g.FullPath, gitignoreTemplate, g.data)
	if err != nil {
		panic(err)
	}
	return result
}
