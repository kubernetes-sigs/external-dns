package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// FirewallRule is the struct of the firewall rule.
type FirewallRule struct {
	ID          string      `json:"id,omitempty"`
	Paused      bool        `json:"paused"`
	Description string      `json:"description"`
	Action      string      `json:"action"`
	Priority    interface{} `json:"priority"`
	Filter      Filter      `json:"filter"`
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	Products    []string    `json:"products,omitempty"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	Products    []string    `json:"products,omitempty"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======
	Products    []string    `json:"products,omitempty"`
>>>>>>> 6b7ce455e (update vendored files)
	CreatedOn   time.Time   `json:"created_on,omitempty"`
	ModifiedOn  time.Time   `json:"modified_on,omitempty"`
}

// FirewallRulesDetailResponse is the API response for the firewall
// rules.
type FirewallRulesDetailResponse struct {
	Result     []FirewallRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FirewallRuleResponse is the API response that is returned
// for requesting a single firewall rule on a zone.
type FirewallRuleResponse struct {
	Result     FirewallRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FirewallRules returns all firewall rules.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/get/#get-all-rules
func (api *API) FirewallRules(ctx context.Context, zoneID string, pageOpts PaginationOptions) ([]FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules", zoneID)
	v := url.Values{}

	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}

	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	if len(v) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []FirewallRule{}, err
	}

	var firewallDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallDetailResponse)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallDetailResponse.Result, nil
}

// FirewallRule returns a single firewall rule based on the ID.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/get/#get-by-rule-id
func (api *API) FirewallRule(ctx context.Context, zoneID, firewallRuleID string) (FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", zoneID, firewallRuleID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return FirewallRule{}, err
	}

	var firewallRuleResponse FirewallRuleResponse
	err = json.Unmarshal(res, &firewallRuleResponse)
	if err != nil {
		return FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRuleResponse.Result, nil
}

// CreateFirewallRules creates new firewall rules.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/post/
func (api *API) CreateFirewallRules(ctx context.Context, zoneID string, firewallRules []FirewallRule) ([]FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, firewallRules)
	if err != nil {
		return []FirewallRule{}, err
	}

	var firewallRulesDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallRulesDetailResponse)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRulesDetailResponse.Result, nil
}

// UpdateFirewallRule updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/put/#update-a-single-rule
func (api *API) UpdateFirewallRule(ctx context.Context, zoneID string, firewallRule FirewallRule) (FirewallRule, error) {
	if firewallRule.ID == "" {
		return FirewallRule{}, errors.Errorf("firewall rule ID cannot be empty")
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", zoneID, firewallRule.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, firewallRule)
	if err != nil {
		return FirewallRule{}, err
	}

	var firewallRuleResponse FirewallRuleResponse
	err = json.Unmarshal(res, &firewallRuleResponse)
	if err != nil {
		return FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRuleResponse.Result, nil
}

// UpdateFirewallRules updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/put/#update-multiple-rules
func (api *API) UpdateFirewallRules(ctx context.Context, zoneID string, firewallRules []FirewallRule) ([]FirewallRule, error) {
	for _, firewallRule := range firewallRules {
		if firewallRule.ID == "" {
			return []FirewallRule{}, errors.Errorf("firewall ID cannot be empty")
		}
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules", zoneID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, firewallRules)
	if err != nil {
		return []FirewallRule{}, err
	}

	var firewallRulesDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallRulesDetailResponse)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRulesDetailResponse.Result, nil
}

// DeleteFirewallRule deletes a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/delete/#delete-a-single-rule
func (api *API) DeleteFirewallRule(ctx context.Context, zoneID, firewallRuleID string) error {
	if firewallRuleID == "" {
		return errors.Errorf("firewall rule ID cannot be empty")
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", zoneID, firewallRuleID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFirewallRules deletes multiple firewall rules at once.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/delete/#delete-multiple-rules
func (api *API) DeleteFirewallRules(ctx context.Context, zoneID string, firewallRuleIDs []string) error {
	v := url.Values{}

	for _, ruleID := range firewallRuleIDs {
		v.Add("id", ruleID)
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules?%s", zoneID, v.Encode())

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// FirewallRule is the struct of the firewall rule.
type FirewallRule struct {
	ID          string      `json:"id,omitempty"`
	Paused      bool        `json:"paused"`
	Description string      `json:"description"`
	Action      string      `json:"action"`
	Priority    interface{} `json:"priority"`
	Filter      Filter      `json:"filter"`
	Products    []string    `json:"products,omitempty"`
	Ref         string      `json:"ref,omitempty"`
	CreatedOn   time.Time   `json:"created_on,omitempty"`
	ModifiedOn  time.Time   `json:"modified_on,omitempty"`
}

// FirewallRulesDetailResponse is the API response for the firewall
// rules.
type FirewallRulesDetailResponse struct {
	Result     []FirewallRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FirewallRuleResponse is the API response that is returned
// for requesting a single firewall rule on a zone.
type FirewallRuleResponse struct {
	Result     FirewallRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FirewallRuleCreateParams contains required and optional params
// for creating a firewall rule.
type FirewallRuleCreateParams struct {
	ID          string      `json:"id,omitempty"`
	Paused      bool        `json:"paused"`
	Description string      `json:"description"`
	Action      string      `json:"action"`
	Priority    interface{} `json:"priority"`
	Filter      Filter      `json:"filter"`
	Products    []string    `json:"products,omitempty"`
	Ref         string      `json:"ref,omitempty"`
}

// FirewallRuleUpdateParams contains required and optional params
// for updating a firewall rule.
type FirewallRuleUpdateParams struct {
	ID          string      `json:"id"`
	Paused      bool        `json:"paused"`
	Description string      `json:"description"`
	Action      string      `json:"action"`
	Priority    interface{} `json:"priority"`
	Filter      Filter      `json:"filter"`
	Products    []string    `json:"products,omitempty"`
	Ref         string      `json:"ref,omitempty"`
}

type FirewallRuleListParams struct {
	ResultInfo
}

// FirewallRules returns all firewall rules.
//
// Automatically paginates all results unless `params.PerPage` and `params.Page`
// is set.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/get/#get-all-rules
func (api *API) FirewallRules(ctx context.Context, rc *ResourceContainer, params FirewallRuleListParams) ([]FirewallRule, *ResultInfo, error) {
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

	var firewallRules []FirewallRule
	var fResponse FirewallRulesDetailResponse
	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/firewall/rules", rc.Identifier), params)

		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []FirewallRule{}, &ResultInfo{}, err
		}

		err = json.Unmarshal(res, &fResponse)
		if err != nil {
			return []FirewallRule{}, &ResultInfo{}, fmt.Errorf("failed to unmarshal filters JSON data: %w", err)
		}

		firewallRules = append(firewallRules, fResponse.Result...)
		params.ResultInfo = fResponse.ResultInfo.Next()

		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return firewallRules, &fResponse.ResultInfo, nil
}

// FirewallRule returns a single firewall rule based on the ID.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/get/#get-by-rule-id
func (api *API) FirewallRule(ctx context.Context, rc *ResourceContainer, firewallRuleID string) (FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", rc.Identifier, firewallRuleID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return FirewallRule{}, err
	}

	var firewallRuleResponse FirewallRuleResponse
	err = json.Unmarshal(res, &firewallRuleResponse)
	if err != nil {
		return FirewallRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return firewallRuleResponse.Result, nil
}

// CreateFirewallRules creates new firewall rules.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/post/
func (api *API) CreateFirewallRules(ctx context.Context, rc *ResourceContainer, params []FirewallRuleCreateParams) ([]FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return []FirewallRule{}, err
	}

	var firewallRulesDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallRulesDetailResponse)
	if err != nil {
		return []FirewallRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return firewallRulesDetailResponse.Result, nil
}

// UpdateFirewallRule updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/put/#update-a-single-rule
func (api *API) UpdateFirewallRule(ctx context.Context, rc *ResourceContainer, params FirewallRuleUpdateParams) (FirewallRule, error) {
	if params.ID == "" {
		return FirewallRule{}, fmt.Errorf("firewall rule ID cannot be empty")
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", rc.Identifier, params.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return FirewallRule{}, err
	}

	var firewallRuleResponse FirewallRuleResponse
	err = json.Unmarshal(res, &firewallRuleResponse)
	if err != nil {
		return FirewallRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return firewallRuleResponse.Result, nil
}

// UpdateFirewallRules updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/put/#update-multiple-rules
func (api *API) UpdateFirewallRules(ctx context.Context, rc *ResourceContainer, params []FirewallRuleUpdateParams) ([]FirewallRule, error) {
	for _, firewallRule := range params {
		if firewallRule.ID == "" {
			return []FirewallRule{}, fmt.Errorf("firewall ID cannot be empty")
		}
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return []FirewallRule{}, err
	}

	var firewallRulesDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallRulesDetailResponse)
	if err != nil {
		return []FirewallRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return firewallRulesDetailResponse.Result, nil
}

// DeleteFirewallRule deletes a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/delete/#delete-a-single-rule
func (api *API) DeleteFirewallRule(ctx context.Context, rc *ResourceContainer, firewallRuleID string) error {
	if firewallRuleID == "" {
		return fmt.Errorf("firewall rule ID cannot be empty")
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", rc.Identifier, firewallRuleID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFirewallRules deletes multiple firewall rules at once.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/delete/#delete-multiple-rules
func (api *API) DeleteFirewallRules(ctx context.Context, rc *ResourceContainer, firewallRuleIDs []string) error {
	v := url.Values{}

	for _, ruleID := range firewallRuleIDs {
		v.Add("id", ruleID)
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules?%s", rc.Identifier, v.Encode())

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return errors.Wrap(err, errMakeRequestError)
=======
		return err
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// FirewallRule is the struct of the firewall rule.
type FirewallRule struct {
	ID          string      `json:"id,omitempty"`
	Paused      bool        `json:"paused"`
	Description string      `json:"description"`
	Action      string      `json:"action"`
	Priority    interface{} `json:"priority"`
	Filter      Filter      `json:"filter"`
	CreatedOn   time.Time   `json:"created_on,omitempty"`
	ModifiedOn  time.Time   `json:"modified_on,omitempty"`
}

// FirewallRulesDetailResponse is the API response for the firewall
// rules.
type FirewallRulesDetailResponse struct {
	Result     []FirewallRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FirewallRuleResponse is the API response that is returned
// for requesting a single firewall rule on a zone.
type FirewallRuleResponse struct {
	Result     FirewallRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// FirewallRules returns all firewall rules.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/get/#get-all-rules
func (api *API) FirewallRules(zoneID string, pageOpts PaginationOptions) ([]FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules", zoneID)
	v := url.Values{}

	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}

	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	if len(v) > 0 {
		uri = uri + "?" + v.Encode()
	}

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errMakeRequestError)
	}

	var firewallDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallDetailResponse)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallDetailResponse.Result, nil
}

// FirewallRule returns a single firewall rule based on the ID.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/get/#get-by-rule-id
func (api *API) FirewallRule(zoneID, firewallRuleID string) (FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", zoneID, firewallRuleID)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return FirewallRule{}, errors.Wrap(err, errMakeRequestError)
	}

	var firewallRuleResponse FirewallRuleResponse
	err = json.Unmarshal(res, &firewallRuleResponse)
	if err != nil {
		return FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRuleResponse.Result, nil
}

// CreateFirewallRules creates new firewall rules.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/post/
func (api *API) CreateFirewallRules(zoneID string, firewallRules []FirewallRule) ([]FirewallRule, error) {
	uri := fmt.Sprintf("/zones/%s/firewall/rules", zoneID)

	res, err := api.makeRequest("POST", uri, firewallRules)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errMakeRequestError)
	}

	var firewallRulesDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallRulesDetailResponse)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRulesDetailResponse.Result, nil
}

// UpdateFirewallRule updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/put/#update-a-single-rule
func (api *API) UpdateFirewallRule(zoneID string, firewallRule FirewallRule) (FirewallRule, error) {
	if firewallRule.ID == "" {
		return FirewallRule{}, errors.Errorf("firewall rule ID cannot be empty")
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", zoneID, firewallRule.ID)

	res, err := api.makeRequest("PUT", uri, firewallRule)
	if err != nil {
		return FirewallRule{}, errors.Wrap(err, errMakeRequestError)
	}

	var firewallRuleResponse FirewallRuleResponse
	err = json.Unmarshal(res, &firewallRuleResponse)
	if err != nil {
		return FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRuleResponse.Result, nil
}

// UpdateFirewallRules updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/put/#update-multiple-rules
func (api *API) UpdateFirewallRules(zoneID string, firewallRules []FirewallRule) ([]FirewallRule, error) {
	for _, firewallRule := range firewallRules {
		if firewallRule.ID == "" {
			return []FirewallRule{}, errors.Errorf("firewall ID cannot be empty")
		}
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules", zoneID)

	res, err := api.makeRequest("PUT", uri, firewallRules)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errMakeRequestError)
	}

	var firewallRulesDetailResponse FirewallRulesDetailResponse
	err = json.Unmarshal(res, &firewallRulesDetailResponse)
	if err != nil {
		return []FirewallRule{}, errors.Wrap(err, errUnmarshalError)
	}

	return firewallRulesDetailResponse.Result, nil
}

// DeleteFirewallRule updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/delete/#delete-a-single-rule
func (api *API) DeleteFirewallRule(zoneID, firewallRuleID string) error {
	if firewallRuleID == "" {
		return errors.Errorf("firewall rule ID cannot be empty")
	}

	uri := fmt.Sprintf("/zones/%s/firewall/rules/%s", zoneID, firewallRuleID)

	_, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}

	return nil
}

// DeleteFirewallRules updates a single firewall rule.
//
// API reference: https://developers.cloudflare.com/firewall/api/cf-firewall-rules/delete/#delete-multiple-rules
func (api *API) DeleteFirewallRules(zoneID string, firewallRuleIDs []string) error {
	ids := strings.Join(firewallRuleIDs, ",")
	uri := fmt.Sprintf("/zones/%s/firewall/rules?id=%s", zoneID, ids)

	_, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	}

	return nil
}
