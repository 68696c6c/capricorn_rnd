package app

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/src/app/container"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"
	"github.com/68696c6c/capricorn_rnd/project/src/app/enum"
)

type App struct {
	*golang.Package
	domains   domain.Map
	container *container.Container
	config    container.Config
}

func NewApp(pkg golang.IPackage, p *config.Project, o config.AppOptions) *App {
	pkgApp := pkg.AddPackage(o.PkgName)
	appEnums := enum.NewEnums(pkgApp, o.Enums, p.Enums)
	appDomains := domain.NewDomains(pkgApp, o.Domain, p.Resources, &appEnums)
	return &App{
		Package:   pkgApp,
		domains:   appDomains,
		container: container.NewContainer(pkgApp, o.ServiceContainer, appDomains),
		config:    container.NewConfig(pkgApp, o.ServiceContainerConfig),
	}
}

func (a *App) GetDomains() domain.Map {
	return a.domains
}

func (a *App) GetContainerType() golang.IType {
	return a.container.GetContainerType()
}

func (a *App) GetContainerConstructor() golang.IType {
	return a.container.GetConstructor()
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
