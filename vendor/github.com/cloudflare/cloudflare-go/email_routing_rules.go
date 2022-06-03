package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrMissingRuleID = errors.New("required rule id missing")

type EmailRoutingRuleMatcher struct {
	Type  string `json:"type,omitempty"`
	Field string `json:"field,omitempty"`
	Value string `json:"value,omitempty"`
}

type EmailRoutingRuleAction struct {
	Type  string   `json:"type,omitempty"`
	Value []string `json:"value,omitempty"`
}

type EmailRoutingRule struct {
	Tag      string                    `json:"tag,omitempty"`
	Name     string                    `json:"name,omitempty"`
	Priority int                       `json:"priority,omitempty"`
	Enabled  *bool                     `json:"enabled,omitempty"`
	Matchers []EmailRoutingRuleMatcher `json:"matchers,omitempty"`
	Actions  []EmailRoutingRuleAction  `json:"actions,omitempty"`
}

type ListEmailRoutingRulesParameters struct {
	Enabled *bool `url:"enabled,omitempty"`
	ResultInfo
}

type ListEmailRoutingRuleResponse struct {
	Result     []EmailRoutingRule `json:"result"`
	ResultInfo `json:"result_info,omitempty"`
	Response
}

type CreateEmailRoutingRuleParameters struct {
	Matchers []EmailRoutingRuleMatcher `json:"matchers,omitempty"`
	Actions  []EmailRoutingRuleAction  `json:"actions,omitempty"`
	Name     string                    `json:"name,omitempty"`
	Enabled  *bool                     `json:"enabled,omitempty"`
	Priority int                       `json:"priority,omitempty"`
}

type CreateEmailRoutingRuleResponse struct {
	Result EmailRoutingRule `json:"result"`
	Response
}

type GetEmailRoutingRuleResponse struct {
	Result EmailRoutingRule `json:"result"`
	Response
}

type UpdateEmailRoutingRuleParameters struct {
	Matchers []EmailRoutingRuleMatcher `json:"matchers,omitempty"`
	Actions  []EmailRoutingRuleAction  `json:"actions,omitempty"`
	Name     string                    `json:"name,omitempty"`
	Enabled  *bool                     `json:"enabled,omitempty"`
	Priority int                       `json:"priority,omitempty"`
	RuleID   string
}

type EmailRoutingCatchAllRule struct {
	Tag      string                    `json:"tag,omitempty"`
	Name     string                    `json:"name,omitempty"`
	Enabled  *bool                     `json:"enabled,omitempty"`
	Matchers []EmailRoutingRuleMatcher `json:"matchers,omitempty"`
	Actions  []EmailRoutingRuleAction  `json:"actions,omitempty"`
}

type EmailRoutingCatchAllRuleResponse struct {
	Result EmailRoutingCatchAllRule `json:"result"`
	Response
}

// ListEmailRoutingRules Lists existing routing rules.
//
// API reference: https://api.cloudflare.com/#email-routing-routing-rules-list-routing-rules
func (api *API) ListEmailRoutingRules(ctx context.Context, rc *ResourceContainer, params ListEmailRoutingRulesParameters) ([]EmailRoutingRule, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []EmailRoutingRule{}, &ResultInfo{}, ErrMissingZoneID
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}
	if params.PerPage < 1 {
		params.PerPage = 50
	}
	if params.Page < 1 {
		params.Page = 1
	}

	var rules []EmailRoutingRule
	var rResponse ListEmailRoutingRuleResponse
	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/email/routing/rules", rc.Identifier), params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []EmailRoutingRule{}, &ResultInfo{}, err
		}

		err = json.Unmarshal(res, &rResponse)
		if err != nil {
			return []EmailRoutingRule{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}

		rules = append(rules, rResponse.Result...)
		params.ResultInfo = rResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return rules, &rResponse.ResultInfo, nil
}

// CreateEmailRoutingRule Rules consist of a set of criteria for matching emails (such as an email being sent to a specific custom email address) plus a set of actions to take on the email (like forwarding it to a specific destination address).
//
// API reference: https://api.cloudflare.com/#email-routing-routing-rules-create-routing-rule
func (api *API) CreateEmailRoutingRule(ctx context.Context, rc *ResourceContainer, params CreateEmailRoutingRuleParameters) (EmailRoutingRule, error) {
	if rc.Identifier == "" {
		return EmailRoutingRule{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/email/routing/rules", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return EmailRoutingRule{}, err
	}

	var r CreateEmailRoutingRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// GetEmailRoutingRule Get information for a specific routing rule already created.
//
// API reference: https://api.cloudflare.com/#email-routing-routing-rules-get-routing-rule
func (api *API) GetEmailRoutingRule(ctx context.Context, rc *ResourceContainer, ruleID string) (EmailRoutingRule, error) {
	if rc.Identifier == "" {
		return EmailRoutingRule{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/email/routing/rules/%s", rc.Identifier, ruleID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return EmailRoutingRule{}, err
	}

	var r GetEmailRoutingRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// UpdateEmailRoutingRule Update actions, matches, or enable/disable specific routing rules
//
// API reference: https://api.cloudflare.com/#email-routing-routing-rules-update-routing-rule
func (api *API) UpdateEmailRoutingRule(ctx context.Context, rc *ResourceContainer, params UpdateEmailRoutingRuleParameters) (EmailRoutingRule, error) {
	if rc.Identifier == "" {
		return EmailRoutingRule{}, ErrMissingZoneID
	}

	if params.RuleID == "" {
		return EmailRoutingRule{}, ErrMissingRuleID
	}

	uri := fmt.Sprintf("/zones/%s/email/routing/rules/%s", rc.Identifier, params.RuleID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return EmailRoutingRule{}, err
	}

	var r GetEmailRoutingRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DeleteEmailRoutingRule Delete a specific routing rule.
//
// API reference: https://api.cloudflare.com/#email-routing-routing-rules-delete-routing-rule
func (api *API) DeleteEmailRoutingRule(ctx context.Context, rc *ResourceContainer, ruleID string) (EmailRoutingRule, error) {
	if rc.Identifier == "" {
		return EmailRoutingRule{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/email/routing/rules/%s", rc.Identifier, ruleID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return EmailRoutingRule{}, err
	}

	var r GetEmailRoutingRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// GetEmailRoutingCatchAllRule Get information on the default catch-all routing rule.
//
// API reference: https://api.cloudflare.com/#email-routing-routing-rules-get-catch-all-rule
func (api *API) GetEmailRoutingCatchAllRule(ctx context.Context, rc *ResourceContainer) (EmailRoutingCatchAllRule, error) {
	if rc.Identifier == "" {
		return EmailRoutingCatchAllRule{}, ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/email/routing/rules/catch_all", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return EmailRoutingCatchAllRule{}, err
	}

	var r EmailRoutingCatchAllRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingCatchAllRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// UpdateEmailRoutingCatchAllRule Enable or disable catch-all routing rule, or change action to forward to specific destination address.
//
// API reference: https://api.cloudflare.com/#email-routing-routing-rules-update-catch-all-rule
func (api *API) UpdateEmailRoutingCatchAllRule(ctx context.Context, rc *ResourceContainer, params EmailRoutingCatchAllRule) (EmailRoutingCatchAllRule, error) {
	if rc.Identifier == "" {
		return EmailRoutingCatchAllRule{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/email/routing/rules/catch_all", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return EmailRoutingCatchAllRule{}, err
	}

	var r EmailRoutingCatchAllRuleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingCatchAllRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}
