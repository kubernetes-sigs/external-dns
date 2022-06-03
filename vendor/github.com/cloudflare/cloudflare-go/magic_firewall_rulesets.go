package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"errors"
)

const (
	// MagicFirewallRulesetKindRoot specifies a root Ruleset.
	MagicFirewallRulesetKindRoot = "root"

	// MagicFirewallRulesetPhaseMagicTransit specifies the Magic Transit Ruleset phase.
	MagicFirewallRulesetPhaseMagicTransit = "magic_transit"

	// MagicFirewallRulesetRuleActionSkip specifies a skip (allow) action.
	MagicFirewallRulesetRuleActionSkip MagicFirewallRulesetRuleAction = "skip"

	// MagicFirewallRulesetRuleActionBlock specifies a block action.
	MagicFirewallRulesetRuleActionBlock MagicFirewallRulesetRuleAction = "block"
)

// MagicFirewallRulesetRuleAction specifies the action for a Firewall rule.
type MagicFirewallRulesetRuleAction string

// MagicFirewallRuleset contains information about a Firewall Ruleset.
type MagicFirewallRuleset struct {
	ID          string                     `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Kind        string                     `json:"kind"`
	Version     string                     `json:"version,omitempty"`
	LastUpdated *time.Time                 `json:"last_updated,omitempty"`
	Phase       string                     `json:"phase"`
	Rules       []MagicFirewallRulesetRule `json:"rules"`
}

// MagicFirewallRulesetRuleActionParameters specifies the action parameters for a Firewall rule.
type MagicFirewallRulesetRuleActionParameters struct {
	Ruleset string `json:"ruleset,omitempty"`
}

// MagicFirewallRulesetRule contains information about a single Magic Firewall rule.
type MagicFirewallRulesetRule struct {
	ID               string                                    `json:"id,omitempty"`
	Version          string                                    `json:"version,omitempty"`
	Action           MagicFirewallRulesetRuleAction            `json:"action"`
	ActionParameters *MagicFirewallRulesetRuleActionParameters `json:"action_parameters,omitempty"`
	Expression       string                                    `json:"expression"`
	Description      string                                    `json:"description"`
	LastUpdated      *time.Time                                `json:"last_updated,omitempty"`
	Ref              string                                    `json:"ref,omitempty"`
	Enabled          bool                                      `json:"enabled"`
}

// CreateMagicFirewallRulesetRequest contains data for a new Firewall ruleset.
type CreateMagicFirewallRulesetRequest struct {
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Kind        string                     `json:"kind"`
	Phase       string                     `json:"phase"`
	Rules       []MagicFirewallRulesetRule `json:"rules"`
}

// UpdateMagicFirewallRulesetRequest contains data for a Magic Firewall ruleset update.
type UpdateMagicFirewallRulesetRequest struct {
	Description string                     `json:"description"`
	Rules       []MagicFirewallRulesetRule `json:"rules"`
}

// ListMagicFirewallRulesetResponse contains a list of Magic Firewall rulesets.
type ListMagicFirewallRulesetResponse struct {
	Response
	Result []MagicFirewallRuleset `json:"result"`
}

// GetMagicFirewallRulesetResponse contains a single Magic Firewall Ruleset.
type GetMagicFirewallRulesetResponse struct {
	Response
	Result MagicFirewallRuleset `json:"result"`
}

// CreateMagicFirewallRulesetResponse contains response data when creating a new Magic Firewall ruleset.
type CreateMagicFirewallRulesetResponse struct {
	Response
	Result MagicFirewallRuleset `json:"result"`
}

// UpdateMagicFirewallRulesetResponse contains response data when updating an existing Magic Firewall ruleset.
type UpdateMagicFirewallRulesetResponse struct {
	Response
	Result MagicFirewallRuleset `json:"result"`
}

// ListMagicFirewallRulesets lists all Rulesets for a given account
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
//
// Deprecated: Use `ListZoneRuleset` or `ListAccountRuleset` instead.
func (api *API) ListMagicFirewallRulesets(ctx context.Context, accountID string) ([]MagicFirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []MagicFirewallRuleset{}, err
	}

	result := ListMagicFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []MagicFirewallRuleset{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// GetMagicFirewallRuleset returns a specific Magic Firewall Ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-get-a-ruleset
//
// Deprecated: Use `GetZoneRuleset` or `GetAccountRuleset` instead.
func (api *API) GetMagicFirewallRuleset(ctx context.Context, accountID, ID string) (MagicFirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", accountID, ID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return MagicFirewallRuleset{}, err
	}

	result := GetMagicFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicFirewallRuleset{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// CreateMagicFirewallRuleset creates a Magic Firewall ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-list-rulesets
//
// Deprecated: Use `CreateZoneRuleset` or `CreateAccountRuleset` instead.
func (api *API) CreateMagicFirewallRuleset(ctx context.Context, accountID, name, description string, rules []MagicFirewallRulesetRule) (MagicFirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets", accountID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri,
		CreateMagicFirewallRulesetRequest{
			Name:        name,
			Description: description,
			Kind:        MagicFirewallRulesetKindRoot,
			Phase:       MagicFirewallRulesetPhaseMagicTransit,
			Rules:       rules})
	if err != nil {
		return MagicFirewallRuleset{}, err
	}

	result := CreateMagicFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicFirewallRuleset{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

// DeleteMagicFirewallRuleset deletes a Magic Firewall ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-delete-ruleset
//
// Deprecated: Use `DeleteZoneRuleset` or `DeleteAccountRuleset` instead.
func (api *API) DeleteMagicFirewallRuleset(ctx context.Context, accountID, ID string) error {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", accountID, ID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return err
	}

	// Firewall API is not implementing the standard response blob but returns an empty response (204) in case
	// of a success. So we are checking for the response body size here
	if len(res) > 0 {
		return fmt.Errorf(errMakeRequestError+": %w", errors.New(string(res)))
	}

	return nil
}

// UpdateMagicFirewallRuleset updates a Magic Firewall ruleset
//
// API reference: https://api.cloudflare.com/#rulesets-update-ruleset
//
// Deprecated: Use `UpdateZoneRuleset` or `UpdateAccountRuleset` instead.
func (api *API) UpdateMagicFirewallRuleset(ctx context.Context, accountID, ID string, description string, rules []MagicFirewallRulesetRule) (MagicFirewallRuleset, error) {
	uri := fmt.Sprintf("/accounts/%s/rulesets/%s", accountID, ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri,
		UpdateMagicFirewallRulesetRequest{Description: description, Rules: rules})
	if err != nil {
		return MagicFirewallRuleset{}, err
	}

	result := UpdateMagicFirewallRulesetResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return MagicFirewallRuleset{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}
