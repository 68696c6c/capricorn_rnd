package handlers

import (
	"fmt"
	"strings"

	"github.com/68696c6c/capricorn_rnd/golang"
)

const (
	verbGet    = "GET"
	verbPost   = "POST"
	verbPut    = "PUT"
	verbDelete = "DELETE"

	paramNameId = "id"
)

type Handler struct {
	verb          string
	uri           string
	handlerFunc   *golang.Function
	requestStruct *golang.Struct
}

func (h *Handler) renderHandlerChain(errorsRef, repoRef string) string {
	var result []string
	if h.requestStruct != nil {
		result = append(result, fmt.Sprintf("goat.BindMiddleware(%s{})", h.requestStruct.GetReference()))
	}
	handlerCall := fmt.Sprintf("%s(%s, %s)", h.handlerFunc.GetReference(), errorsRef, repoRef)
	result = append(result, handlerCall)
	return strings.Join(result, ", ")
}

type handlerGroupMeta struct {
	ContextArg           *golang.Value
	ErrorsArg            *golang.Value
	RepoArg              *golang.Value
	SingleName           string
	PluralName           string
	ModelType            *golang.Struct
	RequestCreateType    *golang.Struct
	RequestUpdateType    *golang.Struct
	ResourceResponseType *golang.Struct
	ListResponseType     *golang.Struct
	RepoPageFuncName     string
	RepoFilterFuncName   string
	ParamNameId          string
}
