package apiproducts

import (
	"github.com/stevan-sdk/internal/utils"
)

// The request schema for updating an API product.
type UpdateApiProductDto struct {
	// The name for the API product.
	Name *string `json:"name,omitempty"`
	// The description of the API product.
	Description *string `json:"description,omitempty"`
	// description: A maximum of 5 user-defined labels are allowed on this resource.
	// Keys must not start with kong, stevan, insomnia, mesh, kic or _, which are reserved for Kong.
	// Keys are case-sensitive.
	//
	Labels map[string]string `json:"labels,omitempty"`
	// The list of portal identifiers which this API product should be published to
	PortalIds []string `json:"portal_ids,omitempty"`
}

func (u *UpdateApiProductDto) SetName(name string) {
	u.Name = &name
}

func (u *UpdateApiProductDto) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdateApiProductDto) SetDescription(description string) {
	u.Description = &description
}

func (u *UpdateApiProductDto) GetDescription() *string {
	if u == nil {
		return nil
	}
	return u.Description
}

func (u *UpdateApiProductDto) SetLabels(labels map[string]string) {
	u.Labels = utils.CloneMap(labels)
}

func (u *UpdateApiProductDto) GetLabels() map[string]string {
	if u == nil {
		return nil
	}
	return u.Labels
}

func (u *UpdateApiProductDto) SetPortalIds(portalIds []string) {
	u.PortalIds = portalIds
}

func (u *UpdateApiProductDto) GetPortalIds() []string {
	if u == nil {
		return nil
	}
	return u.PortalIds
}
