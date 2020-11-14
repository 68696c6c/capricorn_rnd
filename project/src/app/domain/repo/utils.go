package repo

import (
	"fmt"
	"github.com/68696c6c/capricorn_rnd/project/goat"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
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
	dbType              golang.Type
	queryArgName        string
	queryType           golang.Type
	baseQueryFuncName   string
	filterQueryFuncName string
	pageQueryFuncName   string
	repoStructType      golang.Type
	repoInterfaceType   golang.Type
}

func makeMethodMeta(modelType model.Type, recName string, repoStruct, repoInterface golang.Type) methodMeta {
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

// Determine the correct reference to the model.
func determineModelType(repoPkgName string, rawModelType model.Type) model.Type {
	result := rawModelType

	// Repos always operate on pointers to models.
	result.IsPointer = true

	// If we are generating a DDD app, the model and repo will be in the same package and references to the model in the
	// repo file should not include the package name.
	if result.Package == repoPkgName {
		result.Package = ""
	}

	return result
}
