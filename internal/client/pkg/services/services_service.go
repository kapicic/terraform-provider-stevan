package services

import (
	restClient "github.com/stevan-sdk/internal/clients/rest"
	"github.com/stevan-sdk/internal/clients/rest/httptransport"
	"github.com/stevan-sdk/internal/unmarshal"
)

type ServicesService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewServicesService(baseUrl string, bearerToken string) *ServicesService {
	return &ServicesService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

// Create a new service
func (api *ServicesService) CreateService(runtimeGroupId string, service Service) (*CreateService_201Response, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services")

	request.Body = service

	request.SetPathParam("runtimeGroupId", runtimeGroupId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[CreateService_201Response](resp)

}

// Get a service using ID or name.
func (api *ServicesService) GetService(runtimeGroupId string, serviceId string) (*GetService_200Response, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services/{service_id}")

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("service_id", serviceId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[GetService_200Response](resp)

}

// Create or Update service using ID or name.
func (api *ServicesService) UpsertService(runtimeGroupId string, serviceId string, service Service) (*UpsertService_200Response, error) {
	request := httptransport.NewRequest("PUT", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services/{service_id}")

	request.Body = service

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("service_id", serviceId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[UpsertService_200Response](resp)

}

// Delete a service
func (api *ServicesService) DeleteService(runtimeGroupId string, serviceId string) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/services/{service_id}")

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("service_id", serviceId)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
