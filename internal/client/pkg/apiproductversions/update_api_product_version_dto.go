package apiproductversions

// The request schema for updating a version of an API product.
type UpdateApiProductVersionDto struct {
	// The version name of the API product version.
	Name *string `json:"name,omitempty"`
	// The publish status of the API product version.
	PublishStatus *PublishStatus_2 `json:"publish_status,omitempty"`
	// Indicates if the version of the API product is deprecated.
	Deprecated *bool `json:"deprecated,omitempty"`
	// When set to `true`, and all the following conditions are true:
	// - version of the API product deprecation has changed from `false` -> `true`
	// - version of the API product is published
	//
	// then consumers of the now deprecated verion of the API product will be notified.
	//
	Notify         *bool                  `json:"notify,omitempty"`
	GatewayService *GatewayServicePayload `json:"gateway_service,omitempty"`
}

func (u *UpdateApiProductVersionDto) SetName(name string) {
	u.Name = &name
}

func (u *UpdateApiProductVersionDto) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdateApiProductVersionDto) SetPublishStatus(publishStatus PublishStatus_2) {
	u.PublishStatus = &publishStatus
}

func (u *UpdateApiProductVersionDto) GetPublishStatus() *PublishStatus_2 {
	if u == nil {
		return nil
	}
	return u.PublishStatus
}

func (u *UpdateApiProductVersionDto) SetDeprecated(deprecated bool) {
	u.Deprecated = &deprecated
}

func (u *UpdateApiProductVersionDto) GetDeprecated() *bool {
	if u == nil {
		return nil
	}
	return u.Deprecated
}

func (u *UpdateApiProductVersionDto) SetNotify(notify bool) {
	u.Notify = &notify
}

func (u *UpdateApiProductVersionDto) GetNotify() *bool {
	if u == nil {
		return nil
	}
	return u.Notify
}

func (u *UpdateApiProductVersionDto) SetGatewayService(gatewayService GatewayServicePayload) {
	u.GatewayService = &gatewayService
}

func (u *UpdateApiProductVersionDto) GetGatewayService() *GatewayServicePayload {
	if u == nil {
		return nil
	}
	return u.GatewayService
}

// The publish status of the API product version.
type PublishStatus_2 string

const (
	PUBLISH_STATUS_2_UNPUBLISHED PublishStatus_2 = "unpublished"
	PUBLISH_STATUS_2_PUBLISHED   PublishStatus_2 = "published"
)
