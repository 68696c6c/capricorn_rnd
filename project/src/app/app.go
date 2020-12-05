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
	enums     enum.Enums
	domains   *domain.Domains
	container *container.Container
	config    container.Config
}

func NewApp(rootPath string, p *config.Project, o config.AppOptions) *App {
	pkgApp := golang.NewPackage(o.PkgName, rootPath, p.Module)
	appEnums := enum.NewEnums(pkgApp, o.Enums, p.Enums)
	appDomains := domain.NewDomains(pkgApp, o.Domain, p.Resources, &appEnums)
	return &App{
		Package:   pkgApp,
		enums:     appEnums,
		domains:   appDomains,
		container: container.NewContainer(pkgApp, o.ServiceContainer, appDomains),
		config:    container.NewConfig(pkgApp, o.ServiceContainerConfig),
	}
}

func (a *App) GetEnums() enum.Enums {
	return a.enums
}

func (a *App) GetDomains() *domain.Domains {
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
