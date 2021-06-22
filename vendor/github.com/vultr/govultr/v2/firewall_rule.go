package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// FireWallRuleService is the interface to interact with the firewall rule endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/firewall
type FireWallRuleService interface {
	Create(ctx context.Context, fwGroupID string, fwRuleReq *FirewallRuleReq) (*FirewallRule, error)
	Get(ctx context.Context, fwGroupID string, fwRuleID int) (*FirewallRule, error)
	Delete(ctx context.Context, fwGroupID string, fwRuleID int) error
	List(ctx context.Context, fwGroupID string, options *ListOptions) ([]FirewallRule, *Meta, error)
}

// FireWallRuleServiceHandler handles interaction with the firewall rule methods for the Vultr API
type FireWallRuleServiceHandler struct {
	client *Client
}

// FirewallRule represents a Vultr firewall rule
type FirewallRule struct {
	ID         int    `json:"id"`
	Action     string `json:"action"`
	Type       string `json:"type"`
	Protocol   string `json:"protocol"`
	Port       string `json:"port"`
	Subnet     string `json:"subnet"`
	SubnetSize int    `json:"subnet_size"`
	Source     string `json:"source"`
	Notes      string `json:"notes"`
}

// FirewallRuleReq struct used to create a FirewallRule.
type FirewallRuleReq struct {
	IPType     string `json:"ip_type"`
	Protocol   string `json:"protocol"`
	Subnet     string `json:"subnet"`
	SubnetSize int    `json:"subnet_size"`
	Port       string `json:"port,omitempty"`
	Source     string `json:"source,omitempty"`
	Notes      string `json:"notes,omitempty"`
}

type firewallRulesBase struct {
	FirewallRules []FirewallRule `json:"firewall_rules"`
	Meta          *Meta          `json:"meta"`
}

type firewallRuleBase struct {
	FirewallRule *FirewallRule `json:"firewall_rule"`
}

// Create will create a rule in a firewall group.
func (f *FireWallRuleServiceHandler) Create(ctx context.Context, fwGroupID string, fwRuleReq *FirewallRuleReq) (*FirewallRule, error) {
	uri := fmt.Sprintf("/v2/firewalls/%s/rules", fwGroupID)

	req, err := f.client.NewRequest(ctx, http.MethodPost, uri, fwRuleReq)
	if err != nil {
		return nil, err
	}

	firewallRule := new(firewallRuleBase)
	if err = f.client.DoWithContext(ctx, req, firewallRule); err != nil {
		return nil, err
	}

	return firewallRule.FirewallRule, nil
}

// Get will get a rule in a firewall group.
func (f *FireWallRuleServiceHandler) Get(ctx context.Context, fwGroupID string, fwRuleID int) (*FirewallRule, error) {
	uri := fmt.Sprintf("/v2/firewalls/%s/rules/%d", fwGroupID, fwRuleID)

	req, err := f.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	firewallRule := new(firewallRuleBase)
	if err = f.client.DoWithContext(ctx, req, firewallRule); err != nil {
		return nil, err
	}

	return firewallRule.FirewallRule, nil
}

// Delete will delete a firewall rule on your Vultr account
func (f *FireWallRuleServiceHandler) Delete(ctx context.Context, fwGroupID string, fwRuleID int) error {
	uri := fmt.Sprintf("/v2/firewalls/%s/rules/%d", fwGroupID, fwRuleID)

	req, err := f.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return f.client.DoWithContext(ctx, req, nil)
}

// List will return both ipv4 an ipv6 firewall rules that are defined within a firewall group
func (f *FireWallRuleServiceHandler) List(ctx context.Context, fwGroupID string, options *ListOptions) ([]FirewallRule, *Meta, error) {
	uri := fmt.Sprintf("/v2/firewalls/%s/rules", fwGroupID)

	req, err := f.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	firewallRule := new(firewallRulesBase)
	if err = f.client.DoWithContext(ctx, req, firewallRule); err != nil {
		return nil, nil, err
	}

	return firewallRule.FirewallRules, firewallRule.Meta, nil
}
