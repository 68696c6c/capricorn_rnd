package local

import "github.com/68696c6c/gonad/utils"

const gitignoreTemplate = `
.DS_Store
.idea
vendor
.app.env
`

type GitIgnore struct {
	*utils.File
	data Config
}

func NewGitIgnore(basePath string, c Config) GitIgnore {
	file := utils.NewFile(basePath, ".gitignore", "")
	return GitIgnore{
		File: file,
		data: c,
	}
}

func (g GitIgnore) Render() []byte {
	result, err := utils.ParseTemplate(g.FullPath, gitignoreTemplate, g.data)
	if err != nil {
		panic(err)
	}
	return result
}
