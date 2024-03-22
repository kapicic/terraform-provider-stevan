package runtimegroups

import (
	"github.com/stevan-sdk/internal/utils"
)

// The request schema for the update runtime group request.
type UpdateRuntimeGroupRequest struct {
	// The name of the runtime group.
	Name *string `json:"name,omitempty"`
	// The description of the runtime group in stevan.
	Description *string `json:"description,omitempty"`
	// The auth type value of the cluster associated with the Runtime Group.
	AuthType *AuthType_1 `json:"auth_type,omitempty"`
	// Labels to facilitate tagged search on runtime groups. Keys must be of length 1-63 characters, and cannot start with 'kong', 'stevan', 'mesh', 'kic', or '_'.
	Labels map[string]string `json:"labels,omitempty"`
}

func (u *UpdateRuntimeGroupRequest) SetName(name string) {
	u.Name = &name
}

func (u *UpdateRuntimeGroupRequest) GetName() *string {
	if u == nil {
		return nil
	}
	return u.Name
}

func (u *UpdateRuntimeGroupRequest) SetDescription(description string) {
	u.Description = &description
}

func (u *UpdateRuntimeGroupRequest) GetDescription() *string {
	if u == nil {
		return nil
	}
	return u.Description
}

func (u *UpdateRuntimeGroupRequest) SetAuthType(authType AuthType_1) {
	u.AuthType = &authType
}

func (u *UpdateRuntimeGroupRequest) GetAuthType() *AuthType_1 {
	if u == nil {
		return nil
	}
	return u.AuthType
}

func (u *UpdateRuntimeGroupRequest) SetLabels(labels map[string]string) {
	u.Labels = utils.CloneMap(labels)
}

func (u *UpdateRuntimeGroupRequest) GetLabels() map[string]string {
	if u == nil {
		return nil
	}
	return u.Labels
}

// The auth type value of the cluster associated with the Runtime Group.
type AuthType_1 string

const (
	AUTH_TYPE_1_PINNED_CLIENT_CERTS AuthType_1 = "pinned_client_certs"
	AUTH_TYPE_1_PKI_CLIENT_CERTS    AuthType_1 = "pki_client_certs"
)
