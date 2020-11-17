package app

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/src/app/container"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
)

type App struct {
	*golang.Package
	Enums     enum.Enums
	Domains   domain.Map
	Container *container.Container
	Config    container.Config
}

func NewApp(pkg *golang.Package, enums []enum.Enum, resources []*model.Model) App {
	pkgApp := pkg.AddPackage("app")
	appEnums := enum.NewEnums(pkgApp, enums)
	appDomains := domain.NewDomains(pkgApp, resources, &appEnums)
	return App{
		Package:   pkgApp,
		Enums:     appEnums,
		Domains:   appDomains,
		Container: container.NewContainer(pkgApp, appDomains),
		Config:    container.NewConfig(pkgApp),
	}
}
