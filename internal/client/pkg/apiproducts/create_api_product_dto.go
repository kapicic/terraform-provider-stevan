package apiproducts

import (
	"github.com/stevan-sdk/internal/utils"
)

// The request schema to create an API product.
type CreateApiProductDto struct {
	// The name of the API product.
	Name *string `json:"name,omitempty" required:"true"`
	// The description of the API product.
	Description *string `json:"description,omitempty"`
	// description: A maximum of 5 user-defined labels are allowed on this resource.
	// Keys must not start with kong, stevan, insomnia, mesh, kic or _, which are reserved for Kong.
	// Keys are case-sensitive.
	//
	Labels map[string]string `json:"labels,omitempty"`
}

func (c *CreateApiProductDto) SetName(name string) {
	c.Name = &name
}

func (c *CreateApiProductDto) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateApiProductDto) SetDescription(description string) {
	c.Description = &description
}

func (c *CreateApiProductDto) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *CreateApiProductDto) SetLabels(labels map[string]string) {
	c.Labels = utils.CloneMap(labels)
}

func (c *CreateApiProductDto) GetLabels() map[string]string {
	if c == nil {
		return nil
	}
	return c.Labels
}
