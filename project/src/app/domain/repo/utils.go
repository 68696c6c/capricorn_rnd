package repo

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"github.com/68696c6c/capricorn_rnd/project/src/app/domain/model"
	"github.com/68696c6c/capricorn_rnd/utils"
)

type methodMeta struct {
	modelType           model.Type
	modelPlural         string
	modelArgName        string
	receiverName        string
	dbFieldRef          string
	dbFieldName         string
	dbType              *golang.Type
	queryArgName        string
	queryType           *golang.Type
	baseQueryFuncName   string
	filterQueryFuncName string
	pageQueryFuncName   string
	repoStructType      *golang.Type
	repoInterfaceType   *golang.Type
}

func makeMethodMeta(modelType model.Type, recName string, repoStruct, repoInterface *golang.Type) methodMeta {
	dbFieldName := "db"
	return methodMeta{
		modelType:           modelType,
		modelPlural:         strings.ToLower(utils.Plural(modelType.Name)),
		modelArgName:        "model",
		receiverName:        recName,
		dbFieldRef:          fmt.Sprintf("%s.%s", recName, dbFieldName),
		dbFieldName:         dbFieldName,
		dbType:              goat.MakeDbConnectionType(),
		queryArgName:        "query",
		queryType:           goat.MakeQueryType(),
		baseQueryFuncName:   "getBaseQuery",
		filterQueryFuncName: "getFilteredQuery",
		pageQueryFuncName:   "applyPaginationToQuery",
		repoStructType:      repoStruct,
		repoInterfaceType:   repoInterface,
	}
}
