package apiproductversions

type GatewayServicePayload struct {
	// The identifier of a gateway service associated with the version of the API product.
	Id *string `json:"id,omitempty" required:"true"`
	// The identifier of the control plane that the gateway service resides in
	ControlPlaneId *string `json:"control_plane_id,omitempty" required:"true"`
}

func (g *GatewayServicePayload) SetId(id string) {
	g.Id = &id
}

func (g *GatewayServicePayload) GetId() *string {
	if g == nil {
		return nil
	}
	return g.Id
}

func (g *GatewayServicePayload) SetControlPlaneId(controlPlaneId string) {
	g.ControlPlaneId = &controlPlaneId
}

func (g *GatewayServicePayload) GetControlPlaneId() *string {
	if g == nil {
		return nil
	}
	return g.ControlPlaneId
}
