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
	enums     enum.Enums
	domains   domain.Map
	container *container.Container
	config    container.Config
}

func NewApp(pkg *golang.Package, enums []enum.Enum, resources []*model.Model) *App {
	pkgApp := pkg.AddPackage("app")
	appEnums := enum.NewEnums(pkgApp, enums)
	appDomains := domain.NewDomains(pkgApp, resources, &appEnums)
	return &App{
		Package:   pkgApp,
		enums:     appEnums,
		domains:   appDomains,
		container: container.NewContainer(pkgApp, appDomains),
		config:    container.NewConfig(pkgApp),
	}
}

func (a *App) GetDomains() domain.Map {
	return a.domains
}

func (a *App) GetContainerType() golang.IType {
	return a.container.GetContainerType()
}

func (a *App) GetErrorHandlerFieldName() string {
	return a.container.ErrorHandlerField().Name
}

func (a *App) GetDomainRepoFieldName(domainKey string) (string, error) {
	repoField, err := a.container.GetDomainRepoField(domainKey)
	if err != nil {
		return "", err
	}
	return repoField.Name, nil
}
