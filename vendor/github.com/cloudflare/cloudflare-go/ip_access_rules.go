package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type ListIPAccessRulesOrderOption string
type ListIPAccessRulesMatchOption string
type IPAccessRulesModeOption string

const (
	IPAccessRulesConfigurationTarget  ListIPAccessRulesOrderOption = "configuration.target"
	IPAccessRulesConfigurationValue   ListIPAccessRulesOrderOption = "configuration.value"
	IPAccessRulesMatchOptionAll       ListIPAccessRulesMatchOption = "all"
	IPAccessRulesMatchOptionAny       ListIPAccessRulesMatchOption = "any"
	IPAccessRulesModeBlock            IPAccessRulesModeOption      = "block"
	IPAccessRulesModeChallenge        IPAccessRulesModeOption      = "challenge"
	IPAccessRulesModeJsChallenge      IPAccessRulesModeOption      = "js_challenge"
	IPAccessRulesModeManagedChallenge IPAccessRulesModeOption      = "managed_challenge"
	IPAccessRulesModeWhitelist        IPAccessRulesModeOption      = "whitelist"
)

type ListIPAccessRulesFilters struct {
	Configuration IPAccessRuleConfiguration    `json:"configuration,omitempty"`
	Match         ListIPAccessRulesMatchOption `json:"match,omitempty"`
	Mode          IPAccessRulesModeOption      `json:"mode,omitempty"`
	Notes         string                       `json:"notes,omitempty"`
}

type ListIPAccessRulesParams struct {
	Direction string                       `url:"direction,omitempty"`
	Filters   ListIPAccessRulesFilters     `url:"filters,omitempty"`
	Order     ListIPAccessRulesOrderOption `url:"order,omitempty"`
	PaginationOptions
}

type IPAccessRuleConfiguration struct {
	Target string `json:"target,omitempty"`
	Value  string `json:"value,omitempty"`
}

type IPAccessRule struct {
	AllowedModes  []IPAccessRulesModeOption `json:"allowed_modes"`
	Configuration IPAccessRuleConfiguration `json:"configuration"`
	CreatedOn     string                    `json:"created_on"`
	ID            string                    `json:"id"`
	Mode          IPAccessRulesModeOption   `json:"mode"`
	ModifiedOn    string                    `json:"modified_on"`
	Notes         string                    `json:"notes"`
}

type ListIPAccessRulesResponse struct {
	Result     []IPAccessRule `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// ListIPAccessRules fetches IP Access rules of a zone/user/account. You can
// filter the results using several optional parameters.
//
// API references:
//   - https://developers.cloudflare.com/api/operations/ip-access-rules-for-a-user-list-ip-access-rules
//   - https://developers.cloudflare.com/api/operations/ip-access-rules-for-a-zone-list-ip-access-rules
//   - https://developers.cloudflare.com/api/operations/ip-access-rules-for-an-account-list-ip-access-rules
func (api *API) ListIPAccessRules(ctx context.Context, rc *ResourceContainer, params ListIPAccessRulesParams) ([]IPAccessRule, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []IPAccessRule{}, &ResultInfo{}, ErrMissingResourceIdentifier
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/firewall/access_rules/rules", rc.Level, rc.Identifier), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []IPAccessRule{}, &ResultInfo{}, err
	}

	result := ListIPAccessRulesResponse{}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return []IPAccessRule{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, &result.ResultInfo, nil
}
