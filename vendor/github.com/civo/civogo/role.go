package civogo

import (
	"bytes"
	"encoding/json"
	"time"
)

// Role represents a set of permissions
type Role struct {
	ID          string    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Permissions string    `json:"permissions,omitempty"`
	BuiltIn     bool      `json:"built_in,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// ListRoles returns all roles (built-in and user defined)
func (c *Client) ListRoles() ([]Role, error) {
	resp, err := c.SendGetRequest("/v2/roles")
	if err != nil {
		return nil, decodeError(err)
	}

	roles := make([]Role, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&roles); err != nil {
		return nil, err
	}

	return roles, nil
}

// CreateRole creates a new role with a set of permissions for use within an organisation
func (c *Client) CreateRole(name, permissions string) (*Role, error) {
	data := map[string]string{"name": name, "permissions": permissions}
	resp, err := c.SendPostRequest("/v2/roles", data)
	if err != nil {
		return nil, decodeError(err)
	}

	role := &Role{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(role); err != nil {
		return nil, err
	}

	return role, nil
}

// DeleteRole removes a non-built-in role from an organisation
func (c *Client) DeleteRole(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest("/v2/roles/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
