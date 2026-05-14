package linodego

import (
	"context"
	"encoding/json"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

// FirewallRuleSetType represents the type of rules a Rule Set contains.
// Valid values are "inbound" and "outbound".
type FirewallRuleSetType string

const (
	FirewallRuleSetTypeInbound  FirewallRuleSetType = "inbound"
	FirewallRuleSetTypeOutbound FirewallRuleSetType = "outbound"
)

// RuleSet represents the Rule Set resource.
// Note: created/updated/deleted are parsed via UnmarshalJSON into time.Time pointers.
type RuleSet struct {
	ID               int                 `json:"id"`
	Label            string              `json:"label"`
	Description      string              `json:"description,omitempty"`
	Type             FirewallRuleSetType `json:"type"`
	Rules            []FirewallRule      `json:"rules"`
	IsServiceDefined bool                `json:"is_service_defined"`
	Version          int                 `json:"version"`

	Created *time.Time `json:"-"`
	Updated *time.Time `json:"-"`
	Deleted *time.Time `json:"-"`
}

// UnmarshalJSON implements custom timestamp parsing for RuleSet.
func (r *RuleSet) UnmarshalJSON(b []byte) error {
	type Mask RuleSet

	aux := struct {
		*Mask

		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
		Deleted *parseabletime.ParseableTime `json:"deleted"`
	}{
		Mask: (*Mask)(r),
	}

	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}

	if aux.Created != nil {
		r.Created = (*time.Time)(aux.Created)
	}

	if aux.Updated != nil {
		r.Updated = (*time.Time)(aux.Updated)
	}

	if aux.Deleted != nil {
		r.Deleted = (*time.Time)(aux.Deleted)
	}

	return nil
}

// RuleSetCreateOptions fields accepted by CreateRuleSet.
type RuleSetCreateOptions struct {
	Label       string              `json:"label"`
	Description string              `json:"description,omitempty"`
	Type        FirewallRuleSetType `json:"type"`
	Rules       []FirewallRule      `json:"rules"`
}

// RuleSetUpdateOptions fields accepted by UpdateRuleSet.
// Omit a top-level field to leave it unchanged. If Rules is provided, it
// replaces the entire ordered rules array.
type RuleSetUpdateOptions struct {
	Label       *string         `json:"label,omitempty"`
	Description *string         `json:"description,omitempty"`
	Rules       *[]FirewallRule `json:"rules,omitempty"`
}

// ListFirewallRuleSets returns a paginated list of Rule Sets.
// Supports filtering (e.g., by label) via ListOptions.Filter.
func (c *Client) ListFirewallRuleSets(ctx context.Context, opts *ListOptions) ([]RuleSet, error) {
	return getPaginatedResults[RuleSet](ctx, c, "networking/firewalls/rulesets", opts)
}

// CreateFirewallRuleSet creates a new Rule Set.
func (c *Client) CreateFirewallRuleSet(ctx context.Context, opts RuleSetCreateOptions) (*RuleSet, error) {
	return doPOSTRequest[RuleSet](ctx, c, "networking/firewalls/rulesets", opts)
}

// GetFirewallRuleSet fetches a Rule Set by ID.
func (c *Client) GetFirewallRuleSet(ctx context.Context, rulesetID int) (*RuleSet, error) {
	e := formatAPIPath("networking/firewalls/rulesets/%d", rulesetID)
	return doGETRequest[RuleSet](ctx, c, e)
}

// UpdateFirewallRuleSet updates a Rule Set by ID.
func (c *Client) UpdateFirewallRuleSet(ctx context.Context, rulesetID int, opts RuleSetUpdateOptions) (*RuleSet, error) {
	e := formatAPIPath("networking/firewalls/rulesets/%d", rulesetID)
	return doPUTRequest[RuleSet](ctx, c, e, opts)
}

// DeleteFirewallRuleSet deletes a Rule Set by ID.
func (c *Client) DeleteFirewallRuleSet(ctx context.Context, rulesetID int) error {
	e := formatAPIPath("networking/firewalls/rulesets/%d", rulesetID)
	return doDELETERequest(ctx, c, e)
}
