package src

import (
	"github.com/68696c6c/gonad/golang"
	"github.com/68696c6c/gonad/project/src/app"
	"github.com/68696c6c/gonad/project/src/app/domain/model"
	"github.com/68696c6c/gonad/project/src/app/enum"
	"github.com/68696c6c/gonad/project/src/cmd"
	"github.com/68696c6c/gonad/project/src/db"
	"github.com/68696c6c/gonad/project/src/http"
	"github.com/68696c6c/gonad/utils"
)

type Meta struct {
	BasePath  string
	Module    string
	Commands  []cmd.Command
	Enums     []enum.Enum
	Resources []model.Model
}

type SRC struct {
	Main MainGo
	App  app.App
	CMD  cmd.CMD
	DB   db.DB
	HTTP http.HTTP
	root *golang.Package
}

func NewSRC(meta Meta) SRC {
	root := golang.NewRootPackage(meta.BasePath, meta.Module)
	pkgSrc := root.AddPackage("src")

	srcApp := app.NewApp(pkgSrc, meta.Enums, meta.Resources)

	return SRC{
		Main: NewMainGo(pkgSrc),
		App:  srcApp,
		CMD:  cmd.NewCMD(pkgSrc, meta.Commands),
		DB:   db.NewDB(pkgSrc, srcApp),
		HTTP: http.NewHTTP(pkgSrc, srcApp),
		root: root,
	}
}

func (s SRC) GetDirectory() utils.Directory {
	return s.root
}

type MainGo struct {
	file *golang.File
}

func NewMainGo(pkg *golang.Package) MainGo {
	return MainGo{
		file: pkg.AddGoFile("main"),
	}
}
