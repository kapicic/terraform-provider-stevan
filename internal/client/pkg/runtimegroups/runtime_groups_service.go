package runtimegroups

import (
	restClient "github.com/stevan-sdk/internal/clients/rest"
	"github.com/stevan-sdk/internal/clients/rest/httptransport"
	"github.com/stevan-sdk/internal/unmarshal"
)

type RuntimeGroupsService struct {
	client  *restClient.RestClient
	baseUrl string
}

func NewRuntimeGroupsService(baseUrl string, bearerToken string) *RuntimeGroupsService {
	return &RuntimeGroupsService{
		client:  restClient.NewRestClient(baseUrl, bearerToken),
		baseUrl: baseUrl,
	}
}

// Create a runtime group in the stevan Organization.
func (api *RuntimeGroupsService) CreateRuntimeGroup(createRuntimeGroupRequest CreateRuntimeGroupRequest) (*RuntimeGroup, error) {
	request := httptransport.NewRequest("POST", api.baseUrl, "/runtime-groups")

	request.Body = createRuntimeGroupRequest

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[RuntimeGroup](resp)

}

// Returns information about a team from a given team ID.
func (api *RuntimeGroupsService) GetRuntimeGroup(id string) (*RuntimeGroup, error) {
	request := httptransport.NewRequest("GET", api.baseUrl, "/runtime-groups/{id}")

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[RuntimeGroup](resp)

}

// Update an individual runtime group.
func (api *RuntimeGroupsService) UpdateRuntimeGroup(id string, updateRuntimeGroupRequest UpdateRuntimeGroupRequest) (*RuntimeGroup, error) {
	request := httptransport.NewRequest("PATCH", api.baseUrl, "/runtime-groups/{id}")

	request.Body = updateRuntimeGroupRequest

	request.SetPathParam("id", id)

	resp, err := api.client.Call(request)
	if err != nil {
		return nil, err.GetError()
	}

	return unmarshal.ToObject[RuntimeGroup](resp)

}

// Delete an individual runtime group.
func (api *RuntimeGroupsService) DeleteRuntimeGroup(id string) error {
	request := httptransport.NewRequest("DELETE", api.baseUrl, "/runtime-groups/{id}")

	request.SetPathParam("id", id)

	_, err := api.client.Call(request)
	if err != nil {
		return err.GetError()
	}

	return nil

}
