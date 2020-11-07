package local

import "github.com/68696c6c/gonad/utils"

const appEnvTemplate = `
DB_HOST={{ .MainDatabase.Host }}
DB_PORT={{ .MainDatabase.Port }}
DB_USERNAME={{ .MainDatabase.Username }}
DB_PASSWORD={{ .MainDatabase.Password }}
DB_DATABASE={{ .MainDatabase.Name }}
DB_DEBUG={{ .MainDatabase.Debug }}
`

const appExampleEnvTemplate = `
DB_HOST={{ .MainDatabase.Host }}
DB_PORT={{ .MainDatabase.Port }}
DB_USERNAME={{ .MainDatabase.Username }}
DB_PASSWORD=
DB_DATABASE={{ .MainDatabase.Name }}
DB_DEBUG={{ .MainDatabase.Debug }}
`

type AppEnv struct {
	*utils.File
	data      Config
	isExample bool
}

func NewAppEnv(basePath string, c Config, isExample bool) AppEnv {
	name := ".app"
	if isExample {
		name += ".example"
	}
	file := utils.NewFile(basePath, name, "env")
	return AppEnv{
		File: file,
		data: c,
	}
}

func (a AppEnv) Render() []byte {
	t := appEnvTemplate
	if a.isExample {
		t = appExampleEnvTemplate
	}
	result, err := utils.ParseTemplate(a.FullPath, t, a.data)
	if err != nil {
		panic(err)
	}
	return result
}
