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
	structType    *golang.Struct
	interfaceType *golang.Interface
}

func NewService(pkg golang.IPackage, fileName string, meta Meta) *Service {
	if len(meta.Methods) == 0 {
		return nil
	}

	baseTypeName := utils.Pascal(fileName)
	serviceStruct := golang.NewStruct(baseTypeName+"Implementation", false, false)
	serviceInterface := golang.NewInterface(baseTypeName, false, false)

	repoFieldName := "repo"
	serviceStruct.AddField(golang.NewField(repoFieldName, meta.RepoType, false))

	serviceStruct.AddConstructor(makeConstructor(serviceStruct.Type, serviceInterface.Type, meta.RepoType, repoFieldName))

	for _, c := range meta.Methods {
		m := golang.NewFunction(utils.Pascal(c))
		serviceStruct.AddFunction(m)
		serviceInterface.AddFunction(m)
	}

	result := &Service{
		File:          pkg.AddGoFile(fileName),
		structType:    serviceStruct,
		interfaceType: serviceInterface,
	}
	result.AddStruct(serviceStruct)
	result.AddInterface(serviceInterface)

	return result
}

func (s *Service) GetInterfaceType() *golang.Interface {
	return s.interfaceType
}

func (s *Service) GetConstructor() *golang.Function {
	return s.structType.GetConstructor()
}
