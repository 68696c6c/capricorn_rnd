package container

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain"
)

type containerMeta struct {
	domains         domain.Map
	structType      *golang.Struct
	singletonName   string
	dbFieldName     string
	loggerFieldName string
	errorsFieldName string
}

type Container struct {
	*golang.File
}

func NewContainer(pkg *golang.Package, domains domain.Map) *Container {
	containerStruct := golang.NewStruct("ServiceContainer", false, false)
	meta := containerMeta{
		structType:      containerStruct,
		singletonName:   "container",
		domains:         domains,
		dbFieldName:     "DB",
		loggerFieldName: "Logger",
		errorsFieldName: "Errors",
	}

	containerStruct.AddConstructor(makeConstructor(meta))

	containerStruct.AddField(golang.NewField(meta.dbFieldName, goat.MakeDbConnectionType(), true))
	containerStruct.AddField(golang.NewField(meta.loggerFieldName, goat.MakeLoggerType(), true))
	containerStruct.AddField(golang.NewField(meta.errorsFieldName, goat.MakeErrorsType(), true))

	for _, d := range domains {
		repoType := d.Repo.GetInterfaceType()
		containerStruct.AddField(golang.NewField(d.GetExternalRepoName(), repoType, true))

		if d.Service != nil {
			serviceType := d.Service.GetInterfaceType()
			containerStruct.AddField(golang.NewField(d.GetExternalServiceName(), serviceType, true))
		}
	}

	result := &Container{
		File: pkg.AddGoFile("app"),
	}
	result.AddVar(golang.NewVar(meta.singletonName, "", containerStruct.CopyType(), false))
	result.AddStruct(containerStruct)

	return result
}
