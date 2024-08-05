package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// IP represents a serialized structure
type IP struct {
	ID         string     `json:"id"`
	Name       string     `json:"name,omitempty"`
	IP         string     `json:"ip,omitempty"`
	AssignedTo AssignedTo `json:"assigned_to,omitempty"`
}

// AssignedTo represents IP assigned to resource
type AssignedTo struct {
	ID string `json:"id"`
	// Type can be one of the following:
	// - instance
	// - loadbalancer
	Type string `json:"type"`
	Name string `json:"name"`
}

// CreateIPRequest is a struct for creating an IP
type CreateIPRequest struct {
	// Name is an optional parameter. If not provided, name will be the IP address
	Name string `json:"name,omitempty"`

	// Region is the region the IP will be created in
	Region string `json:"region"`
}

// PaginatedIPs is a paginated list of IPs
type PaginatedIPs struct {
	Page    int  `json:"page"`
	PerPage int  `json:"per_page"`
	Pages   int  `json:"pages"`
	Items   []IP `json:"items"`
}

// UpdateIPRequest is a struct for creating an IP
type UpdateIPRequest struct {
	Name string `json:"name" validate:"required"`
	// Region is the region the IP will be created in
	Region string `json:"region"`
}

// Actions for IP
type Actions struct {
	// Action is one of "assign", "unassign"
	Action       string `json:"action"`
	AssignToID   string `json:"assign_to_id"`
	AssignToType string `json:"assign_to_type"`
	// Region is the region the IP will be created in
	Region string `json:"region"`
}

// ListIPs returns all reserved IPs in that specific region
func (c *Client) ListIPs() (*PaginatedIPs, error) {
	resp, err := c.SendGetRequest("/v2/ips")
	if err != nil {
		return nil, decodeError(err)
	}

	ips := &PaginatedIPs{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&ips); err != nil {
		return nil, err
	}

	return ips, nil
}

// GetIP finds an reserved IP by the full ID
func (c *Client) GetIP(id string) (*IP, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/ips/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	var ip = IP{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&ip); err != nil {
		return nil, err
	}

	return &ip, nil
}

// FindIP finds an reserved IP by name or by IP
func (c *Client) FindIP(search string) (*IP, error) {
	ips, err := c.ListIPs()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := IP{}

	for _, value := range ips.Items {
		if value.IP == search || value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.IP, search) || strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// NewIP creates a new IP
func (c *Client) NewIP(v *CreateIPRequest) (*IP, error) {
	body, err := c.SendPostRequest("/v2/ips", v)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &IP{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateIP updates an IP
func (c *Client) UpdateIP(id string, v *UpdateIPRequest) (*IP, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/ips/%s", id), v)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &IP{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// AssignIP assigns a reserved IP to a Civo resource
func (c *Client) AssignIP(id, resourceID, resourceType, region string) (*SimpleResponse, error) {
	actions := &Actions{
		Action: "assign",
		Region: region,
	}

	if resourceID == "" || resourceType == "" {
		return nil, fmt.Errorf("resource ID and type are required")
	}

	actions.AssignToID = resourceID
	actions.AssignToType = resourceType

	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/ips/%s/actions", id), actions)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UnassignIP unassigns a reserved IP from a Civo resource
// UnassignIP is an idempotent operation. If you unassign on a unassigned IP, it will return a 200 OK.
func (c *Client) UnassignIP(id, region string) (*SimpleResponse, error) {
	actions := &Actions{
		Action: "unassign",
		Region: region,
	}

	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/ips/%s/actions", id), actions)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// DeleteIP deletes an IP
func (c *Client) DeleteIP(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/ips/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
