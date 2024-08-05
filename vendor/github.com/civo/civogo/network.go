package civogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Network represents a private network for instances to connect to
type Network struct {
	ID                    string   `json:"id"`
	Name                  string   `json:"name,omitempty"`
	Default               bool     `json:"default"`
	CIDR                  string   `json:"cidr,omitempty"`
	CIDRV6                string   `json:"cidr_v6,omitempty"`
	Label                 string   `json:"label,omitempty"`
	Status                string   `json:"status,omitempty"`
	IPv4Enabled           bool     `json:"ipv4_enabled,omitempty"`
	IPv6Enabled           bool     `json:"ipv6_enabled,omitempty"`
	NameserversV4         []string `json:"nameservers_v4,omitempty"`
	NameserversV6         []string `json:"nameservers_v6,omitempty"`
	VlanID                int      `json:"vlan_id" validate:"required" schema:"vlan_id"`
	HardwareAddr          string   `json:"hardware_addr,omitempty" schema:"hardware_addr"`
	GatewayIPv4           string   `json:"gateway_ipv4" validate:"required" schema:"gateway_ipv4"`
	AllocationPoolV4Start string   `json:"allocation_pool_v4_start" validate:"required" schema:"allocation_pool_v4_start"`
	AllocationPoolV4End   string   `json:"allocation_pool_v4_end" validate:"required" schema:"allocation_pool_v4_end"`
}

// Subnet represents a subnet within a private network
type Subnet struct {
	ID         string `json:"id"`
	Name       string `json:"name,omitempty"`
	NetworkID  string `json:"network_id"`
	SubnetSize string `json:"subnet_size,omitempty"`
	Status     string `json:"status,omitempty"`
}

// SubnetConfig contains incoming request parameters for the subnet object
type SubnetConfig struct {
	Name string `json:"name" validate:"required" schema:"name"`
}

// Route represents a route within a subnet
type Route struct {
	ID           string `json:"id"`
	SubnetID     string `json:"subnet_id"`
	NetworkID    string `json:"network_id"`
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

// CreateRoute contains incoming request parameters for creating a route object
type CreateRoute struct {
	ResourceID   string `json:"resource_id"`
	ResourceType string `json:"resource_type"`
}

// VLANConnectConfig represents the connection of a network to a VLAN
type VLANConnectConfig struct {
	// VLanID is the ID of the VLAN to connect to
	VlanID int `json:"vlan_id" validate:"required" schema:"vlan_id"`

	// HardwareAddr is the base interface(default: eth0) at which we want to setup VLAN.
	HardwareAddr string `json:"hardware_addr,omitempty" schema:"hardware_addr"`

	// CIDRv4 is the CIDR of the VLAN to connect to
	CIDRv4 string `json:"cidr_v4" validate:"required" schema:"cidr_v4"`

	// GatewayIP is the gateway IP address
	GatewayIPv4 string `json:"gateway_ipv4" validate:"required" schema:"gateway_ipv4"`

	// AllocationPoolV4Start address of the allocation pool
	AllocationPoolV4Start string `json:"allocation_pool_v4_start" validate:"required" schema:"allocation_pool_v4_start"`

	// AllocationPoolV4End address of the allocation pool
	AllocationPoolV4End string `json:"allocation_pool_v4_end" validate:"required" schema:"allocation_pool_v4_end"`
}

// NetworkConfig contains incoming request parameters for the network object
type NetworkConfig struct {
	Label         string             `json:"label" validate:"required" schema:"label"`
	Default       string             `json:"default" schema:"default"`
	IPv4Enabled   *bool              `json:"ipv4_enabled"`
	NameserversV4 []string           `json:"nameservers_v4"`
	CIDRv4        string             `json:"cidr_v4"`
	IPv6Enabled   *bool              `json:"ipv6_enabled"`
	NameserversV6 []string           `json:"nameservers_v6"`
	Region        string             `json:"region"`
	VLanConfig    *VLANConnectConfig `json:"vlan_connect,omitempty"`
}

// NetworkResult represents the result from a network create/update call
type NetworkResult struct {
	ID     string `json:"id"`
	Label  string `json:"label"`
	Result string `json:"result"`
}

// GetDefaultNetwork finds the default private network for an account
func (c *Client) GetDefaultNetwork() (*Network, error) {
	resp, err := c.SendGetRequest("/v2/networks")
	if err != nil {
		return nil, decodeError(err)
	}

	networks := make([]Network, 0)
	json.NewDecoder(bytes.NewReader(resp)).Decode(&networks)
	for _, network := range networks {
		if network.Default {
			return &network, nil
		}
	}

	return nil, errors.New("no default network found")
}

// GetNetwork gets a network with ID
func (c *Client) GetNetwork(id string) (*Network, error) {
	resp, err := c.SendGetRequest("/v2/networks/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	network := Network{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&network)
	return &network, err
}

// NewNetwork creates a new private network
func (c *Client) NewNetwork(label string) (*NetworkResult, error) {
	nc := NetworkConfig{Label: label, Region: c.Region}
	body, err := c.SendPostRequest("/v2/networks", nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListNetworks list all private networks
func (c *Client) ListNetworks() ([]Network, error) {
	resp, err := c.SendGetRequest("/v2/networks")
	if err != nil {
		return nil, decodeError(err)
	}

	networks := make([]Network, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&networks); err != nil {
		return nil, err
	}

	return networks, nil
}

// FindNetwork finds a network by either part of the ID or part of the name
func (c *Client) FindNetwork(search string) (*Network, error) {
	networks, err := c.ListNetworks()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Network{}

	for _, value := range networks {
		if value.Name == search || value.ID == search || value.Label == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Name, search) || strings.Contains(value.ID, search) || strings.Contains(value.Label, search) {
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

// RenameNetwork renames an existing private network
func (c *Client) RenameNetwork(label, id string) (*NetworkResult, error) {
	nc := NetworkConfig{Label: label, Region: c.Region}
	body, err := c.SendPutRequest("/v2/networks/"+id, nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteNetwork deletes a private network
func (c *Client) DeleteNetwork(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/networks/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// GetSubnet gets a subnet with ID
func (c *Client) GetSubnet(networkID, subnetID string) (*Subnet, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/networks/%s/subnets/", networkID) + subnetID)
	if err != nil {
		return nil, decodeError(err)
	}

	subnet := Subnet{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&subnet)
	return &subnet, err
}

// ListSubnets list all subnets for a private network
func (c *Client) ListSubnets(networkID string) ([]Subnet, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/networks/%s/subnets", networkID))
	if err != nil {
		return nil, decodeError(err)
	}

	subnets := make([]Subnet, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&subnets); err != nil {
		return nil, err
	}

	return subnets, nil
}

// CreateSubnet creates a new subnet for a private network
func (c *Client) CreateSubnet(networkID string, subnet SubnetConfig) (*Subnet, error) {
	body, err := c.SendPostRequest(fmt.Sprintf("/v2/networks/%s/subnets", networkID), subnet)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &Subnet{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// FindSubnet finds a subnet by either part of the ID or part of the name
func (c *Client) FindSubnet(search, networkID string) (*Subnet, error) {
	subnets, err := c.ListSubnets(networkID)
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Subnet{}

	for _, value := range subnets {
		if value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
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

// AttachSubnetToInstance attaches a subnet to an instance
func (c *Client) AttachSubnetToInstance(networkID, subnetID string, route *CreateRoute) (*Route, error) {
	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/networks/%s/subnets/%s/routes", networkID, subnetID), route)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &Route{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// DetachSubnetFromInstance detaches a subnet from an instance
func (c *Client) DetachSubnetFromInstance(networkID, subnetID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/networks/%s/subnets/%s/routes", networkID, subnetID))
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err

}

// DeleteSubnet deletes a subnet
func (c *Client) DeleteSubnet(networkID, subnetID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/networks/%s/subnets/%s", networkID, subnetID))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// CreateNetwork creates a new network
func (c *Client) CreateNetwork(nc NetworkConfig) (*NetworkResult, error) {
	body, err := c.SendPostRequest("/v2/networks", nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateNetwork updates an existing network
func (c *Client) UpdateNetwork(id string, nc NetworkConfig) (*NetworkResult, error) {
	body, err := c.SendPutRequest("/v2/networks/"+id, nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
