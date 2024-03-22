package handlers

import (
	"github.com/stevan-sdk/internal/clients/rest/httptransport"
	"github.com/stevan-sdk/internal/validation"
)

type RequestValidationHandler struct {
	nextHandler Handler
}

func NewRequestValidationHandler() *RequestValidationHandler {
	return &RequestValidationHandler{
		nextHandler: nil,
	}
}

func (h *RequestValidationHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	err := validation.ValidateData(request.Body)
	if err != nil {
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	err = validation.ValidateData(request.Options)
	if err != nil {
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	return h.nextHandler.Handle(request)
}

func (h *RequestValidationHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}
