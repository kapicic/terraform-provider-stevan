package apiproducts

import (
	restClient "github.com/stevan-sdk/internal/clients/rest"
	"github.com/stevan-sdk/internal/clients/rest/httptransport"
	"github.com/stevan-sdk/internal/unmarshal"
)

type ApiProductsService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewApiProductsService(baseUrl string, bearerToken string) *ApiProductsService {
	return &ApiProductsService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

// An API Product packages together associated APIs and their operations, backed by thorough documentation,
// with the objective of making API consumption straightforward for developers.
func (api *ApiProductsService) CreateApiProduct(createApiProductDto CreateApiProductDto) (*ApiProduct, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/api-products")

	request.Body = createApiProductDto

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[ApiProduct](resp)

}

// Returns an API product.
func (api *ApiProductsService) GetApiProduct(id string) (*ApiProduct, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/api-products/{id}")

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[ApiProduct](resp)

}

// Updates an API product.
func (api *ApiProductsService) UpdateApiProduct(id string, updateApiProductDto UpdateApiProductDto) (*ApiProduct, error) {
	request := httptransport.NewRequest("PATCH", api.baseUrl, "/api-products/{id}")

	request.Body = updateApiProductDto

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[ApiProduct](resp)

}

// Removes an individual API product.
func (api *ApiProductsService) DeleteApiProduct(id string) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/api-products/{id}")

	request.SetPathParam("id", id)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
