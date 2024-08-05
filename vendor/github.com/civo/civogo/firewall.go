package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// Firewall represents list of rule in Civo's infrastructure
type Firewall struct {
	ID                string         `json:"id"`
	Name              string         `json:"name,omitempty"`
	RulesCount        int            `json:"rules_count,omitempty"`
	InstanceCount     int            `json:"instance_count"`
	ClusterCount      int            `json:"cluster_count"`
	LoadBalancerCount int            `json:"loadbalancer_count"`
	NetworkID         string         `json:"network_id,omitempty"`
	Rules             []FirewallRule `json:"rules,omitempty"`
}

// FirewallResult is the response from the Civo Firewall APIs
type FirewallResult struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Result string `json:"result"`
}

// FirewallRule represents a single rule for a given firewall, regarding
// which ports to open and which protocol, to which CIDR
type FirewallRule struct {
	ID         string   `json:"id,omitempty"`
	FirewallID string   `json:"firewall_id,omitempty"`
	Protocol   string   `json:"protocol"`
	StartPort  string   `json:"start_port"`
	EndPort    string   `json:"end_port"`
	Cidr       []string `json:"cidr"`
	Direction  string   `json:"direction"`
	Action     string   `json:"action"`
	Label      string   `json:"label,omitempty"`
	Ports      string   `json:"ports,omitempty"`
}

// FirewallRuleConfig is how you specify the details when creating a new rule
type FirewallRuleConfig struct {
	FirewallID string   `json:"firewall_id"`
	Region     string   `json:"region"`
	Protocol   string   `json:"protocol"`
	StartPort  string   `json:"start_port"`
	EndPort    string   `json:"end_port"`
	Cidr       []string `json:"cidr"`
	Direction  string   `json:"direction"`
	Action     string   `json:"action"`
	Label      string   `json:"label,omitempty"`
	// Ports will be chosen over StartPort,EndPort if both are provided
	Ports string `json:"ports,omitempty"`
}

// FirewallConfig is how you specify the details when creating a new firewall
type FirewallConfig struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	NetworkID string `json:"network_id"`
	// CreateRules if not send the value will be nil, that mean the default rules will be created
	CreateRules *bool          `json:"create_rules,omitempty"`
	Rules       []FirewallRule `json:"rules,omitempty"`
}

// ListFirewalls returns all firewall owned by the calling API account
func (c *Client) ListFirewalls() ([]Firewall, error) {
	resp, err := c.SendGetRequest("/v2/firewalls")
	if err != nil {
		return nil, decodeError(err)
	}

	firewall := make([]Firewall, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&firewall); err != nil {
		return nil, err
	}

	return firewall, nil
}

// FindFirewall finds a firewall by either part of the ID or part of the name
func (c *Client) FindFirewall(search string) (*Firewall, error) {
	firewalls, err := c.ListFirewalls()
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

// NewFirewall creates a new firewall record
func (c *Client) NewFirewall(firewall *FirewallConfig) (*FirewallResult, error) {
	body, err := c.SendPostRequest("/v2/firewalls", firewall)
	if err != nil {
		return nil, decodeError(err)
	}

	result := &FirewallResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// RenameFirewall rename firewall
func (c *Client) RenameFirewall(id string, f *FirewallConfig) (*SimpleResponse, error) {
	f.Region = c.Region
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/firewalls/%s", id), f)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// DeleteFirewall deletes an firewall
func (c *Client) DeleteFirewall(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest("/v2/firewalls/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// NewFirewallRule creates a new rule within a firewall
func (c *Client) NewFirewallRule(r *FirewallRuleConfig) (*FirewallRule, error) {
	if len(r.FirewallID) == 0 {
		err := fmt.Errorf("the firewall ID is empty")
		return nil, IDisEmptyError.wrap(err)
	}

	r.Region = c.Region

	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/firewalls/%s/rules", r.FirewallID), r)
	if err != nil {
		return nil, decodeError(err)
	}

	rule := &FirewallRule{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(rule); err != nil {
		return nil, err
	}

	return rule, nil
}

// ListFirewallRules get all rules for a firewall
func (c *Client) ListFirewallRules(id string) ([]FirewallRule, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/firewalls/%s/rules", id))
	if err != nil {
		return nil, decodeError(err)
	}

	firewallRule := make([]FirewallRule, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&firewallRule); err != nil {
		return nil, err
	}

	return firewallRule, nil
}

// FindFirewallRule finds a firewall Rule by ID or part of the same
func (c *Client) FindFirewallRule(firewallID string, search string) (*FirewallRule, error) {
	firewallsRules, err := c.ListFirewallRules(firewallID)
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

// DeleteFirewallRule deletes an firewall
func (c *Client) DeleteFirewallRule(id string, ruleID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/firewalls/%s/rules/%s", id, ruleID))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
