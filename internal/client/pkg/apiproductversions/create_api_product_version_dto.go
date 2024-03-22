package apiproductversions

// The request schema to create a version of an API product.
type CreateApiProductVersionDto struct {
	// The version name of the API product version.
	Name *string `json:"name,omitempty" required:"true"`
	// The publish status of the API product version.
	PublishStatus *PublishStatus `json:"publish_status,omitempty"`
	// Indicates if the version of the API product is deprecated.
	Deprecated     *bool                  `json:"deprecated,omitempty"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
}

func (c *CreateApiProductVersionDto) SetName(name string) {
	c.Name = &name
}

func (c *CreateApiProductVersionDto) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateApiProductVersionDto) SetPublishStatus(publishStatus PublishStatus) {
	c.PublishStatus = &publishStatus
}

func (c *CreateApiProductVersionDto) GetPublishStatus() *PublishStatus {
	if c == nil {
		return nil
	}
	return c.PublishStatus
}

func (c *CreateApiProductVersionDto) SetDeprecated(deprecated bool) {
	c.Deprecated = &deprecated
}

func (c *CreateApiProductVersionDto) GetDeprecated() *bool {
	if c == nil {
		return nil
	}
	return c.Deprecated
}

func (c *CreateApiProductVersionDto) SetGatewayService(gatewayService GatewayServicePayload) {
	c.GatewayService = &gatewayService
}

func (c *CreateApiProductVersionDto) GetGatewayService() *GatewayServicePayload {
	if c == nil {
		return nil
	}
	return c.GatewayService
}

// The publish status of the API product version.
type PublishStatus string

const (
	PUBLISH_STATUS_UNPUBLISHED PublishStatus = "unpublished"
	PUBLISH_STATUS_PUBLISHED   PublishStatus = "published"
)
