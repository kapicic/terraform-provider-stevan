package runtimegroups

import (
	"github.com/stevan-sdk/internal/utils"
)

// The request schema for the create runtime group request.
type CreateRuntimeGroupRequest struct {
	// The name of the runtime group.
	Name *string `json:"name,omitempty" required:"true"`
	// The description of the runtime group in stevan.
	Description *string `json:"description,omitempty"`
	// The ClusterType value of the cluster associated with the Runtime Group.
	ClusterType *ClusterType `json:"cluster_type,omitempty"`
	// The auth type value of the cluster associated with the Runtime Group.
	AuthType *AuthType `json:"auth_type,omitempty"`
	// Labels to facilitate tagged search on runtime groups. Keys must be of length 1-63 characters, and cannot start with 'kong', 'stevan', 'mesh', 'kic', or '_'.
	Labels map[string]string `json:"labels,omitempty"`
}

func (c *CreateRuntimeGroupRequest) SetName(name string) {
	c.Name = &name
}

func (c *CreateRuntimeGroupRequest) GetName() *string {
	if c == nil {
		return nil
	}
	return c.Name
}

func (c *CreateRuntimeGroupRequest) SetDescription(description string) {
	c.Description = &description
}

func (c *CreateRuntimeGroupRequest) GetDescription() *string {
	if c == nil {
		return nil
	}
	return c.Description
}

func (c *CreateRuntimeGroupRequest) SetClusterType(clusterType ClusterType) {
	c.ClusterType = &clusterType
}

func (c *CreateRuntimeGroupRequest) GetClusterType() *ClusterType {
	if c == nil {
		return nil
	}
	return c.ClusterType
}

func (c *CreateRuntimeGroupRequest) SetAuthType(authType AuthType) {
	c.AuthType = &authType
}

func (c *CreateRuntimeGroupRequest) GetAuthType() *AuthType {
	if c == nil {
		return nil
	}
	return c.AuthType
}

func (c *CreateRuntimeGroupRequest) SetLabels(labels map[string]string) {
	c.Labels = utils.CloneMap(labels)
}

func (c *CreateRuntimeGroupRequest) GetLabels() map[string]string {
	if c == nil {
		return nil
	}
	return c.Labels
}

// The ClusterType value of the cluster associated with the Runtime Group.
type ClusterType string

const (
	CLUSTER_TYPE_CLUSTER_TYPE_HYBRID                  ClusterType = "CLUSTER_TYPE_HYBRID"
	CLUSTER_TYPE_CLUSTER_TYPE_K8_S_INGRESS_CONTROLLER ClusterType = "CLUSTER_TYPE_K8S_INGRESS_CONTROLLER"
	CLUSTER_TYPE_CLUSTER_TYPE_COMPOSITE               ClusterType = "CLUSTER_TYPE_COMPOSITE"
)

// The auth type value of the cluster associated with the Runtime Group.
type AuthType string

const (
	AUTH_TYPE_PINNED_CLIENT_CERTS AuthType = "pinned_client_certs"
	AUTH_TYPE_PKI_CLIENT_CERTS    AuthType = "pki_client_certs"
)
