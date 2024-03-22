package runtimegroups

import (
	"github.com/stevan-sdk/internal/utils"
)

type RuntimeGroup struct {
	// The runtime group ID.
	Id *string `json:"id,omitempty"`
	// The name of the runtime group.
	Name *string `json:"name,omitempty"`
	// The description of the runtime group in stevan.
	Description *string `json:"description,omitempty"`
	// Labels to facilitate tagged search on runtime groups. Keys must be of length 1-63 characters, and cannot start with 'kong', 'stevan', 'mesh', 'kic', or '_'.
	Labels map[string]string `json:"labels,omitempty"`
	// CP configuration object for related access endpoints.
	Config *Config `json:"config,omitempty"`
	// An ISO-8604 timestamp representation of runtime group creation date.
	CreatedAt *string `json:"created_at,omitempty"`
	// An ISO-8604 timestamp representation of runtime group update date.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (r *RuntimeGroup) SetId(id string) {
	r.Id = &id
}

func (r *RuntimeGroup) GetId() *string {
	if r == nil {
		return nil
	}
	return r.Id
}

func (r *RuntimeGroup) SetName(name string) {
	r.Name = &name
}

func (r *RuntimeGroup) GetName() *string {
	if r == nil {
		return nil
	}
	return r.Name
}

func (r *RuntimeGroup) SetDescription(description string) {
	r.Description = &description
}

func (r *RuntimeGroup) GetDescription() *string {
	if r == nil {
		return nil
	}
	return r.Description
}

func (r *RuntimeGroup) SetLabels(labels map[string]string) {
	r.Labels = utils.CloneMap(labels)
}

func (r *RuntimeGroup) GetLabels() map[string]string {
	if r == nil {
		return nil
	}
	return r.Labels
}

func (r *RuntimeGroup) SetConfig(config Config) {
	r.Config = &config
}

func (r *RuntimeGroup) GetConfig() *Config {
	if r == nil {
		return nil
	}
	return r.Config
}

func (r *RuntimeGroup) SetCreatedAt(createdAt string) {
	r.CreatedAt = &createdAt
}

func (r *RuntimeGroup) GetCreatedAt() *string {
	if r == nil {
		return nil
	}
	return r.CreatedAt
}

func (r *RuntimeGroup) SetUpdatedAt(updatedAt string) {
	r.UpdatedAt = &updatedAt
}

func (r *RuntimeGroup) GetUpdatedAt() *string {
	if r == nil {
		return nil
	}
	return r.UpdatedAt
}

// CP configuration object for related access endpoints.
type Config struct {
	// Control Plane Endpoint.
	ControlPlaneEndpoint *string `json:"control_plane_endpoint,omitempty"`
	// Telemetry Endpoint.
	TelemetryEndpoint *string `json:"telemetry_endpoint,omitempty"`
}

func (c *Config) SetControlPlaneEndpoint(controlPlaneEndpoint string) {
	c.ControlPlaneEndpoint = &controlPlaneEndpoint
}

func (c *Config) GetControlPlaneEndpoint() *string {
	if c == nil {
		return nil
	}
	return c.ControlPlaneEndpoint
}

func (c *Config) SetTelemetryEndpoint(telemetryEndpoint string) {
	c.TelemetryEndpoint = &telemetryEndpoint
}

func (c *Config) GetTelemetryEndpoint() *string {
	if c == nil {
		return nil
	}
	return c.TelemetryEndpoint
}
