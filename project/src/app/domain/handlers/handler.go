package handlers

import (
	"github.com/68696c6c/capricorn_rnd/golang"
)

type Handler struct {
	*golang.Function
	verb          string
	uri           string
	requestStruct *golang.Struct
}

func (h *Handler) GetVerb() string {
	return h.verb
}

func (h *Handler) GetUri() string {
	return h.uri
}

func (h *Handler) HasRequest() bool {
	return h.requestStruct != nil
}

func (h *Handler) GetRequestStruct() golang.IType {
	return h.requestStruct
}
