package repo

import (
	"fmt"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/config"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type methodMeta struct {
	modelType         golang.IType
	modelPlural       string
	receiverName      string
	dbFieldRef        string
	dbType            *golang.Type
	queryType         *golang.Type
	repoStructType    *golang.Struct
	repoInterfaceType *golang.Interface
	constructor       *golang.Function
}

func makeMethodMeta(o config.RepoOptions, domainMeta *config.DomainMeta, repoStruct *golang.Struct, repoInterface *golang.Interface) *methodMeta {
	recName := repoStruct.GetReceiverName()
	return &methodMeta{
		modelType:         domainMeta.GetModelType(),
		modelPlural:       utils.Camel(domainMeta.NamePlural),
		receiverName:      recName,
		dbFieldRef:        fmt.Sprintf("%s.%s", recName, o.DbFieldName),
		dbType:            goat.MakeTypeDbConnection(),
		queryType:         goat.MakeTypeQuery(),
		repoStructType:    repoStruct,
		repoInterfaceType: repoInterface,
	}
}

func (m *methodMeta) AddConstructor(constructor *golang.Function) {
	m.constructor = constructor
}
