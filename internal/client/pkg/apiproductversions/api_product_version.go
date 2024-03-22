package apiproductversions

type ApiProductVersion struct {
	// The API product version identifier.
	Id *string `json:"id,omitempty"`
	// The version of the API product
	Name           *string                `json:"name,omitempty" required:"true"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty" required:"true"`
	// The publish status of the API product version
	PublishStatus *PublishStatus_1 `json:"publish_status,omitempty" required:"true"`
	// Indicates if this API product version is deprecated
	Deprecated *bool `json:"deprecated,omitempty"`
	// An ISO-8601 timestamp representation of entity creation date.
	CreatedAt *string `json:"created_at,omitempty"`
	// An ISO-8601 timestamp representation of entity update date.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (a *ApiProductVersion) SetId(id string) {
	a.Id = &id
}

func (a *ApiProductVersion) GetId() *string {
	if a == nil {
		return nil
	}
	return a.Id
}

func (a *ApiProductVersion) SetName(name string) {
	a.Name = &name
}

func (a *ApiProductVersion) GetName() *string {
	if a == nil {
		return nil
	}
	return a.Name
}

func (a *ApiProductVersion) SetGatewayService(gatewayService GatewayServicePayload) {
	a.GatewayService = &gatewayService
}

func (a *ApiProductVersion) GetGatewayService() *GatewayServicePayload {
	if a == nil {
		return nil
	}
	return a.GatewayService
}

func (a *ApiProductVersion) SetPublishStatus(publishStatus PublishStatus_1) {
	a.PublishStatus = &publishStatus
}

func (a *ApiProductVersion) GetPublishStatus() *PublishStatus_1 {
	if a == nil {
		return nil
	}
	return a.PublishStatus
}

func (a *ApiProductVersion) SetDeprecated(deprecated bool) {
	a.Deprecated = &deprecated
}

func (a *ApiProductVersion) GetDeprecated() *bool {
	if a == nil {
		return nil
	}
	return a.Deprecated
}

func (a *ApiProductVersion) SetCreatedAt(createdAt string) {
	a.CreatedAt = &createdAt
}

func (a *ApiProductVersion) GetCreatedAt() *string {
	if a == nil {
		return nil
	}
	return a.CreatedAt
}

func (a *ApiProductVersion) SetUpdatedAt(updatedAt string) {
	a.UpdatedAt = &updatedAt
}

func (a *ApiProductVersion) GetUpdatedAt() *string {
	if a == nil {
		return nil
	}
	return a.UpdatedAt
}

// The publish status of the API product version
type PublishStatus_1 string

const (
	PUBLISH_STATUS_1_UNPUBLISHED PublishStatus_1 = "unpublished"
	PUBLISH_STATUS_1_PUBLISHED   PublishStatus_1 = "published"
)
