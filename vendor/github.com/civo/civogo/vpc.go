package civogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// VPC API provides methods for the VPC umbrella which contains:
// - Networks (/v2/vpc/networks)
// - Firewalls (/v2/vpc/firewalls)
// - Load Balancers (/v2/vpc/loadbalancers)
// - Reserved IPs (/v2/vpc/ips)
//
// These are aliases for the existing endpoints. Both paths work identically.

// =============================================================================
// VPC Networks - /v2/vpc/networks
// =============================================================================

// GetDefaultVPCNetwork finds the default private network for an account using VPC API path
func (c *Client) GetDefaultVPCNetwork() (*Network, error) {
	resp, err := c.SendGetRequest("/v2/vpc/networks")
	if err != nil {
		return nil, decodeError(err)
	}

	networks := make([]Network, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&networks); err != nil {
		return nil, fmt.Errorf("could not decode networks: %w", err)
	}
	for _, network := range networks {
		if network.Default {
			return &network, nil
		}
	}

	return nil, errors.New("no default network found")
}

// GetVPCNetwork gets a network with ID using VPC API path
func (c *Client) GetVPCNetwork(id string) (*Network, error) {
	resp, err := c.SendGetRequest("/v2/vpc/networks/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	network := Network{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&network)
	return &network, err
}

// NewVPCNetwork creates a new private network using VPC API path
func (c *Client) NewVPCNetwork(label string) (*NetworkResult, error) {
	nc := NetworkConfig{Label: label, Region: c.Region}
	body, err := c.SendPostRequest("/v2/vpc/networks", nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListVPCNetworks lists all private networks using VPC API path
func (c *Client) ListVPCNetworks() ([]Network, error) {
	resp, err := c.SendGetRequest("/v2/vpc/networks")
	if err != nil {
		return nil, decodeError(err)
	}

	networks := make([]Network, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&networks); err != nil {
		return nil, err
	}

	return networks, nil
}

// FindVPCNetwork finds a network by either part of the ID or part of the name using VPC API path
func (c *Client) FindVPCNetwork(search string) (*Network, error) {
	networks, err := c.ListVPCNetworks()
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

// RenameVPCNetwork renames an existing private network using VPC API path
func (c *Client) RenameVPCNetwork(label, id string) (*NetworkResult, error) {
	nc := NetworkConfig{Label: label, Region: c.Region}
	body, err := c.SendPutRequest("/v2/vpc/networks/"+id, nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteVPCNetwork deletes a private network using VPC API path
func (c *Client) DeleteVPCNetwork(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/vpc/networks/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// CreateVPCNetwork creates a new network using VPC API path
func (c *Client) CreateVPCNetwork(nc NetworkConfig) (*NetworkResult, error) {
	body, err := c.SendPostRequest("/v2/vpc/networks", nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateVPCNetwork updates an existing network using VPC API path
func (c *Client) UpdateVPCNetwork(id string, nc NetworkConfig) (*NetworkResult, error) {
	body, err := c.SendPutRequest("/v2/vpc/networks/"+id, nc)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &NetworkResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// =============================================================================
// VPC Subnets - /v2/vpc/networks/{network_id}/subnets
// =============================================================================

// GetVPCSubnet gets a subnet with ID using VPC API path
func (c *Client) GetVPCSubnet(networkID, subnetID string) (*Subnet, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/vpc/networks/%s/subnets/%s", networkID, subnetID))
	if err != nil {
		return nil, decodeError(err)
	}

	subnet := Subnet{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&subnet)
	return &subnet, err
}

// ListVPCSubnets lists all subnets for a private network using VPC API path
func (c *Client) ListVPCSubnets(networkID string) ([]Subnet, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/vpc/networks/%s/subnets", networkID))
	if err != nil {
		return nil, decodeError(err)
	}

	subnets := make([]Subnet, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&subnets); err != nil {
		return nil, err
	}

	return subnets, nil
}

// CreateVPCSubnet creates a new subnet for a private network using VPC API path
func (c *Client) CreateVPCSubnet(networkID string, subnet SubnetConfig) (*Subnet, error) {
	body, err := c.SendPostRequest(fmt.Sprintf("/v2/vpc/networks/%s/subnets", networkID), subnet)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &Subnet{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// FindVPCSubnet finds a subnet by either part of the ID or part of the name using VPC API path
func (c *Client) FindVPCSubnet(search, networkID string) (*Subnet, error) {
	subnets, err := c.ListVPCSubnets(networkID)
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

// AttachVPCSubnetToInstance attaches a subnet to an instance using VPC API path
func (c *Client) AttachVPCSubnetToInstance(networkID, subnetID string, route *CreateRoute) (*Route, error) {
	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/vpc/networks/%s/subnets/%s/routes", networkID, subnetID), route)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &Route{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// DetachVPCSubnetFromInstance detaches a subnet from an instance using VPC API path
func (c *Client) DetachVPCSubnetFromInstance(networkID, subnetID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/vpc/networks/%s/subnets/%s/routes", networkID, subnetID))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// DeleteVPCSubnet deletes a subnet using VPC API path
func (c *Client) DeleteVPCSubnet(networkID, subnetID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/vpc/networks/%s/subnets/%s", networkID, subnetID))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// =============================================================================
// VPC Firewalls - /v2/vpc/firewalls
// =============================================================================

// ListVPCFirewalls returns all firewall owned by the calling API account using VPC API path
func (c *Client) ListVPCFirewalls() ([]Firewall, error) {
	resp, err := c.SendGetRequest("/v2/vpc/firewalls")
	if err != nil {
		return nil, decodeError(err)
	}

	firewall := make([]Firewall, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&firewall); err != nil {
		return nil, err
	}

	return firewall, nil
}

// FindVPCFirewall finds a firewall by either part of the ID or part of the name using VPC API path
func (c *Client) FindVPCFirewall(search string) (*Firewall, error) {
	firewalls, err := c.ListVPCFirewalls()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Firewall{}

	for _, value := range firewalls {
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

// NewVPCFirewall creates a new firewall record using VPC API path
func (c *Client) NewVPCFirewall(firewall *FirewallConfig) (*FirewallResult, error) {
	body, err := c.SendPostRequest("/v2/vpc/firewalls", firewall)
	if err != nil {
		return nil, decodeError(err)
	}

	result := &FirewallResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// RenameVPCFirewall renames a firewall using VPC API path
func (c *Client) RenameVPCFirewall(id string, f *FirewallConfig) (*SimpleResponse, error) {
	f.Region = c.Region
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/vpc/firewalls/%s", id), f)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// DeleteVPCFirewall deletes a firewall using VPC API path
func (c *Client) DeleteVPCFirewall(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest("/v2/vpc/firewalls/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// =============================================================================
// VPC Firewall Rules - /v2/vpc/firewalls/{id}/rules
// =============================================================================

// NewVPCFirewallRule creates a new rule within a firewall using VPC API path
func (c *Client) NewVPCFirewallRule(r *FirewallRuleConfig) (*FirewallRule, error) {
	if len(r.FirewallID) == 0 {
		err := fmt.Errorf("the firewall ID is empty")
		return nil, IDisEmptyError.wrap(err)
	}

	r.Region = c.Region

	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/vpc/firewalls/%s/rules", r.FirewallID), r)
	if err != nil {
		return nil, decodeError(err)
	}

	rule := &FirewallRule{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(rule); err != nil {
		return nil, err
	}

	return rule, nil
}

// ListVPCFirewallRules gets all rules for a firewall using VPC API path
func (c *Client) ListVPCFirewallRules(id string) ([]FirewallRule, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/vpc/firewalls/%s/rules", id))
	if err != nil {
		return nil, decodeError(err)
	}

	firewallRule := make([]FirewallRule, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&firewallRule); err != nil {
		return nil, err
	}

	return firewallRule, nil
}

// FindVPCFirewallRule finds a firewall rule by ID or part of the same using VPC API path
func (c *Client) FindVPCFirewallRule(firewallID string, search string) (*FirewallRule, error) {
	firewallsRules, err := c.ListVPCFirewallRules(firewallID)
	if err != nil {
		return nil, decodeError(err)
	}

	found := -1

	for i, firewallRule := range firewallsRules {
		if strings.Contains(firewallRule.ID, search) {
			if found != -1 {
				err := fmt.Errorf("unable to find %s because there were multiple matches", search)
				return nil, MultipleMatchesError.wrap(err)
			}
			found = i
		}
	}

	if found == -1 {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}

	return &firewallsRules[found], nil
}

// DeleteVPCFirewallRule deletes a firewall rule using VPC API path
func (c *Client) DeleteVPCFirewallRule(id string, ruleID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/vpc/firewalls/%s/rules/%s", id, ruleID))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// =============================================================================
// VPC Load Balancers - /v2/vpc/loadbalancers
// =============================================================================

// ListVPCLoadBalancers returns all load balancers owned by the calling API account using VPC API path
func (c *Client) ListVPCLoadBalancers() ([]LoadBalancer, error) {
	resp, err := c.SendGetRequest("/v2/vpc/loadbalancers")
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := make([]LoadBalancer, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&loadbalancer); err != nil {
		return nil, decodeError(err)
	}

	return loadbalancer, nil
}

// GetVPCLoadBalancer returns a load balancer using VPC API path
func (c *Client) GetVPCLoadBalancer(id string) (*LoadBalancer, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/vpc/loadbalancers/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := &LoadBalancer{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&loadbalancer); err != nil {
		return nil, decodeError(err)
	}

	return loadbalancer, nil
}

// FindVPCLoadBalancer finds a load balancer by either part of the ID or part of the name using VPC API path
func (c *Client) FindVPCLoadBalancer(search string) (*LoadBalancer, error) {
	lbs, err := c.ListVPCLoadBalancers()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := LoadBalancer{}

	for _, value := range lbs {
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

// CreateVPCLoadBalancer creates a new load balancer using VPC API path
func (c *Client) CreateVPCLoadBalancer(r *LoadBalancerConfig) (*LoadBalancer, error) {
	body, err := c.SendPostRequest("/v2/vpc/loadbalancers", r)
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := &LoadBalancer{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(loadbalancer); err != nil {
		return nil, err
	}

	return loadbalancer, nil
}

// UpdateVPCLoadBalancer updates a load balancer using VPC API path
func (c *Client) UpdateVPCLoadBalancer(id string, r *LoadBalancerUpdateConfig) (*LoadBalancer, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/vpc/loadbalancers/%s", id), r)
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := &LoadBalancer{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(loadbalancer); err != nil {
		return nil, err
	}

	return loadbalancer, nil
}

// DeleteVPCLoadBalancer deletes a load balancer using VPC API path
func (c *Client) DeleteVPCLoadBalancer(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/vpc/loadbalancers/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// =============================================================================
// VPC Reserved IPs - /v2/vpc/ips
// =============================================================================

// ListVPCIPs returns all reserved IPs in that specific region using VPC API path
func (c *Client) ListVPCIPs() (*PaginatedIPs, error) {
	resp, err := c.SendGetRequest("/v2/vpc/ips")
	if err != nil {
		return nil, decodeError(err)
	}

	ips := &PaginatedIPs{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&ips); err != nil {
		return nil, err
	}

	return ips, nil
}

// GetVPCIP finds a reserved IP by the full ID using VPC API path
func (c *Client) GetVPCIP(id string) (*IP, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/vpc/ips/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	var ip = IP{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&ip); err != nil {
		return nil, err
	}

	return &ip, nil
}

// FindVPCIP finds a reserved IP by name or by IP using VPC API path
func (c *Client) FindVPCIP(search string) (*IP, error) {
	ips, err := c.ListVPCIPs()
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

// NewVPCIP creates a new reserved IP using VPC API path
func (c *Client) NewVPCIP(v *CreateIPRequest) (*IP, error) {
	body, err := c.SendPostRequest("/v2/vpc/ips", v)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &IP{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateVPCIP updates a reserved IP using VPC API path
func (c *Client) UpdateVPCIP(id string, v *UpdateIPRequest) (*IP, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/vpc/ips/%s", id), v)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &IP{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// AssignVPCIP assigns a reserved IP to a Civo resource using VPC API path
func (c *Client) AssignVPCIP(id, resourceID, resourceType, region string) (*SimpleResponse, error) {
	actions := &Actions{
		Action: "assign",
		Region: region,
	}

	if resourceID == "" || resourceType == "" {
		return nil, fmt.Errorf("resource ID and type are required")
	}

	actions.AssignToID = resourceID
	actions.AssignToType = resourceType

	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/vpc/ips/%s/actions", id), actions)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UnassignVPCIP unassigns a reserved IP from a Civo resource using VPC API path
func (c *Client) UnassignVPCIP(id, region string) (*SimpleResponse, error) {
	actions := &Actions{
		Action: "unassign",
		Region: region,
	}

	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/vpc/ips/%s/actions", id), actions)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// DeleteVPCIP deletes a reserved IP using VPC API path
func (c *Client) DeleteVPCIP(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/vpc/ips/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
