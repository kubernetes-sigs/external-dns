package ddosx

import "github.com/ukfast/sdk-go/pkg/connection"

// GetACLGeoIPRulesModeResponseBody represents an API response body containing ACLGeoIPRulesMode data
type GetACLGeoIPRulesModeResponseBody struct {
	connection.APIResponseBody

	Data struct {
		Mode ACLGeoIPRulesMode `json:"mode"`
	} `json:"data"`
}
