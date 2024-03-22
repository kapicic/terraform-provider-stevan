package rest

import (
	"github.com/stevan-sdk/internal/clients/rest/handlers"
	"github.com/stevan-sdk/internal/clients/rest/hooks"
	"github.com/stevan-sdk/internal/clients/rest/httptransport"
)

type RestClient struct {
	handlers *handlers.HandlerChain
}

func NewRestClient(baseUrl string, bearerToken string) *RestClient {
	defaultHeadersHandler := handlers.NewDefaultHeadersHandler()
	retryHandler := handlers.NewRetryHandler()
	bearerTokenHandler := handlers.NewBearerTokenHandler(bearerToken)
	hookHandler := handlers.NewHookHandler(hooks.NewCustomHook())
	requestValidationHandler := handlers.NewRequestValidationHandler()
	terminatingHandler := handlers.NewTerminatingHandler()

	handlers := handlers.BuildHandlerChain().
		AddHandler(defaultHeadersHandler).
		AddHandler(retryHandler).
		AddHandler(bearerTokenHandler).
		AddHandler(hookHandler).
		AddHandler(requestValidationHandler).
		AddHandler(terminatingHandler)

	return &RestClient{
		handlers: handlers,
	}
}

func (client *RestClient) Call(request httptransport.Request) (*httptransport.Response, *httptransport.ErrorResponse) {
	return client.handlers.CallApi(request)
}
