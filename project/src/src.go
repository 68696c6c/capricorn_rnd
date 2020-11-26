package src

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
	"github.com/68696c6c/capricorn_rnd/project/src/cmd"
	"github.com/68696c6c/capricorn_rnd/project/src/db"
	"github.com/68696c6c/capricorn_rnd/project/src/http"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Meta struct {
	Module    string
	CmdMeta   *config.CmdMeta
	Commands  []cmd.Command
	Enums     []enum.Enum
	Resources []*model.Model
}

func Build(root *utils.Folder, meta Meta) {
	pkgSrc := golang.NewPackage("src", root.GetFullPath(), meta.Module)
	srcApp := app.NewApp(pkgSrc, meta.Enums, meta.Resources)

	cmd.Build(pkgSrc, meta.Commands, meta.CmdMeta)
	db.Build(pkgSrc, srcApp)
	http.Build(pkgSrc, srcApp)
	buildMainGo(pkgSrc)

	root.AddDirectory(pkgSrc)
}

type MainGo struct {
	*golang.File
}

func buildMainGo(pkgSrc *golang.Package) MainGo {
	return MainGo{
		File: pkgSrc.AddGoFile("main"),
	}
}
