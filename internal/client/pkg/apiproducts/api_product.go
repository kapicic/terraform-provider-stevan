package apiproducts

import (
	"github.com/stevan-sdk/internal/utils"
)

type ApiProduct struct {
	// The API product ID.
	Id *string `json:"id,omitempty" required:"true"`
	// The name of the API product
	Name *string `json:"name,omitempty" required:"true"`
	// The description of the API product
	Description *string `json:"description,omitempty" required:"true"`
	// The list of portal identifiers which this API product is published to
	PortalIds []string `json:"portal_ids,omitempty" required:"true"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt *string `json:"created_at,omitempty" required:"true"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt *string `json:"updated_at,omitempty" required:"true"`
	// description: A maximum of 5 user-defined labels are allowed on this resource.
	// Keys must not start with kong, stevan, insomnia, mesh, kic or _, which are reserved for Kong.
	// Keys are case-sensitive.
	//
	Labels map[string]string `json:"labels,omitempty" required:"true"`
}

func (a *ApiProduct) SetId(id string) {
	a.Id = &id
}

func (a *ApiProduct) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *ApiProduct) SetName(name string) {
	a.Name = &name
}

func (a *ApiProduct) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *ApiProduct) SetDescription(description string) {
	a.Description = &description
}

func (a *ApiProduct) GetDescription() *string {
	if a == nil {
		return nil
	}
	return a.Description
}

func (a *ApiProduct) SetPortalIds(portalIds []string) {
	a.PortalIds = portalIds
}

func (a *ApiProduct) GetPortalIds() []string {
	if a == nil {
		return nil
	}
	return a.PortalIds
}

func (a *ApiProduct) SetCreatedAt(createdAt string) {
	a.CreatedAt = &createdAt
}

func (a *ApiProduct) GetCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *ApiProduct) SetUpdatedAt(updatedAt string) {
	a.UpdatedAt = &updatedAt
}

func (a *ApiProduct) GetUpdatedAt() *string {
	if a == nil {
		return nil
	}
	return a.UpdatedAt
}

func (a *ApiProduct) SetLabels(labels map[string]string) {
	a.Labels = utils.CloneMap(labels)
}

func (a *ApiProduct) GetLabels() map[string]string {
	if a == nil {
		return nil
	}
	return a.Labels
}
