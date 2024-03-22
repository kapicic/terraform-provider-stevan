package apiproductversions

import (
	restClient "github.com/stevan-sdk/internal/clients/rest"
	"github.com/stevan-sdk/internal/clients/rest/httptransport"
	"github.com/stevan-sdk/internal/unmarshal"
)

type ApiProductVersionsService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewApiProductVersionsService(baseUrl string, bearerToken string) *ApiProductVersionsService {
	return &ApiProductVersionsService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

// Creates a version of an API product.
func (api *ApiProductVersionsService) CreateApiProductVersion(apiProductId string, createApiProductVersionDto CreateApiProductVersionDto) (*ApiProductVersion, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/api-products/{apiProductId}/product-versions")

	request.Body = createApiProductVersionDto

	request.SetPathParam("apiProductId", apiProductId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[ApiProductVersion](resp)

}

// Returns a version of an API product.
func (api *ApiProductVersionsService) GetApiProductVersion(apiProductId string, id string) (*ApiProductVersion, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/api-products/{apiProductId}/product-versions/{id}")

	request.SetPathParam("apiProductId", apiProductId)
	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[ApiProductVersion](resp)

}

// Updates an API product version.
func (api *ApiProductVersionsService) UpdateApiProductVersion(apiProductId string, id string, updateApiProductVersionDto UpdateApiProductVersionDto) (*ApiProductVersion, error) {
	request := httptransport.NewRequest("PATCH", api.baseUrl, "/api-products/{apiProductId}/product-versions/{id}")

	request.Body = updateApiProductVersionDto

	request.SetPathParam("apiProductId", apiProductId)
	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[ApiProductVersion](resp)

}

// Removes an API product version.
func (api *ApiProductVersionsService) DeleteApiProductVersion(apiProductId string, id string) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/api-products/{apiProductId}/product-versions/{id}")

	request.SetPathParam("apiProductId", apiProductId)
	request.SetPathParam("id", id)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
