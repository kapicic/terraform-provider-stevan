package client

import (
	"github.com/stevan-sdk/pkg/apiproducts"
	"github.com/stevan-sdk/pkg/apiproductversions"
	"github.com/stevan-sdk/pkg/routes"
	"github.com/stevan-sdk/pkg/runtimegroups"
	"github.com/stevan-sdk/pkg/services"
)

type Client struct {
	Routes             *routes.RoutesService
	Services           *services.ServicesService
	ApiProducts        *apiproducts.ApiProductsService
	ApiProductVersions *apiproductversions.ApiProductVersionsService
	RuntimeGroups      *runtimegroups.RuntimeGroupsService
}

func NewClient(baseUrl string, bearerToken string) *Client {

	return &Client{
		Routes:             routes.NewRoutesService(baseUrl, bearerToken),
		Services:           services.NewServicesService(baseUrl, bearerToken),
		ApiProducts:        apiproducts.NewApiProductsService(baseUrl, bearerToken),
		ApiProductVersions: apiproductversions.NewApiProductVersionsService(baseUrl, bearerToken),
		RuntimeGroups:      runtimegroups.NewRuntimeGroupsService(baseUrl, bearerToken),
	}
}
