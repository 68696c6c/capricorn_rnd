package service

import (
	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type Meta struct {
	Methods  []string
	RepoType *golang.Interface
}

type Service struct {
	*golang.File
	serviceStruct    *golang.Struct
	serviceInterface *golang.Interface
}

func NewService(pkg *golang.Package, fileName string, meta Meta) *Service {
	repoType := determineRepoType(pkg.GetName(), meta.RepoType)

	baseTypeName := utils.Pascal(fileName)
	serviceStruct := golang.NewStruct(baseTypeName+"Implementation", false, false)
	serviceInterface := golang.NewInterface(baseTypeName, false, false)

	repoFieldName := "repo"
	serviceStruct.AddField(golang.NewField(repoFieldName, repoType, false))

	serviceStruct.AddConstructor(makeConstructor(serviceStruct, serviceInterface, repoType, repoFieldName))

	for _, c := range meta.Methods {
		m := golang.NewFunction(utils.Pascal(c))
		serviceStruct.AddFunction(m)
		serviceInterface.AddFunction(m)
	}

	result := &Service{
		File: pkg.AddGoFile(fileName),
	}
	result.AddStruct(serviceStruct)
	result.AddInterface(serviceInterface)

	return result
}

func determineRepoType(servicePkgName string, inputRepoType *golang.Interface) *golang.Interface {
	result := inputRepoType

	// If we are generating a DDD app, the model and repo will be in the same package and references to the model in the
	// repo file should not include the package name.
	if result.Package == servicePkgName {
		result.Package = ""
	}

	return result
}
