package handlers

import (
	"errors"

	"github.com/stevan-sdk/internal/clients/rest/httptransport"
)

type DefaultHeadersHandler struct {
	defaultHeaders map[string]string
	nextHandler    Handler
}

func NewDefaultHeadersHandler() *DefaultHeadersHandler {
	defaultHeaders := map[string]string{
		"User-Agent":   "liblab/0.0.8 go/1.18",
		"Content-type": "application/json",
	}

	return &DefaultHeadersHandler{
		defaultHeaders: defaultHeaders,
		nextHandler:    nil,
	}
}

func (h *DefaultHeadersHandler) Handle(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	nextRequest := request.Clone()

	for key, value := range h.defaultHeaders {
		nextRequest.SetHeader(key, value)
	}

	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse(err, nil)
	}

	return h.nextHandler.Handle(nextRequest)
}

func (h *DefaultHeadersHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}
