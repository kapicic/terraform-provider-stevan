package routes

import (
	restClient "github.com/stevan-sdk/internal/clients/rest"
	"github.com/stevan-sdk/internal/clients/rest/httptransport"
	"github.com/stevan-sdk/internal/unmarshal"
)

type RoutesService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewRoutesService(baseUrl string, bearerToken string) *RoutesService {
	return &RoutesService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

// Create a new route
func (api *RoutesService) CreateRoute(runtimeGroupId string, route Route) (*Route, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes")

	request.Body = route

	request.SetPathParam("runtimeGroupId", runtimeGroupId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[Route](resp)

}

// Get a route using ID or name.
func (api *RoutesService) GetRoute(runtimeGroupId string, routeId string) (*Route, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes/{route_id}")

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("route_id", routeId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[Route](resp)

}

// Create or Update route using ID or name.
func (api *RoutesService) UpsertRoute(runtimeGroupId string, routeId string, route Route) (*Route, error) {
	request := httptransport.NewRequest("PUT", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes/{route_id}")

	request.Body = route

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("route_id", routeId)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[Route](resp)

}

// Delete a route.
func (api *RoutesService) DeleteRoute(runtimeGroupId string, routeId string) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/runtime-groups/{runtimeGroupId}/core-entities/routes/{route_id}")

	request.SetPathParam("runtimeGroupId", runtimeGroupId)
	request.SetPathParam("route_id", routeId)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
