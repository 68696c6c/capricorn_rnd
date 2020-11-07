package app

import (
	"github.com/68696c6c/gonad/golang"
	"github.com/68696c6c/gonad/project/src/app/container"
	"github.com/68696c6c/gonad/project/src/app/domain"
	"github.com/68696c6c/gonad/project/src/app/domain/model"
	"github.com/68696c6c/gonad/project/src/app/enum"
)

type App struct {
	pkg       *golang.Package
	Enums     enum.Enums
	Domains   domain.Map
	Container container.Container
	Config    container.Config
}

func NewApp(pkg *golang.Package, enums []enum.Enum, resources []model.Model) App {
	pkgApp := pkg.AddPackage("app")
	appEnums := enum.NewEnums(pkgApp, enums)
	appDomains := domain.NewDomains(pkgApp, resources)
	return App{
		pkg:       pkgApp,
		Enums:     appEnums,
		Domains:   appDomains,
		Container: container.NewContainer(pkgApp),
		Config:    container.NewConfig(pkgApp),
	}
}