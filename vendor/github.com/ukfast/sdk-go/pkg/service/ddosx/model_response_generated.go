package ddosx

import "github.com/ukfast/sdk-go/pkg/connection"

// GetDomainSliceResponseBody represents an API response body containing []Domain data
type GetDomainSliceResponseBody struct {
	connection.APIResponseBody
	Data []Domain `json:"data"`
}

// GetDomainResponseBody represents an API response body containing Domain data
type GetDomainResponseBody struct {
	connection.APIResponseBody
	Data Domain `json:"data"`
}

// GetDomainPropertySliceResponseBody represents an API response body containing []DomainProperty data
type GetDomainPropertySliceResponseBody struct {
	connection.APIResponseBody
	Data []DomainProperty `json:"data"`
}

// GetDomainPropertyResponseBody represents an API response body containing DomainProperty data
type GetDomainPropertyResponseBody struct {
	connection.APIResponseBody
	Data DomainProperty `json:"data"`
}

// GetRecordSliceResponseBody represents an API response body containing []Record data
type GetRecordSliceResponseBody struct {
	connection.APIResponseBody
	Data []Record `json:"data"`
}

// GetRecordResponseBody represents an API response body containing Record data
type GetRecordResponseBody struct {
	connection.APIResponseBody
	Data Record `json:"data"`
}

// GetWAFSliceResponseBody represents an API response body containing []WAF data
type GetWAFSliceResponseBody struct {
	connection.APIResponseBody
	Data []WAF `json:"data"`
}

// GetWAFResponseBody represents an API response body containing WAF data
type GetWAFResponseBody struct {
	connection.APIResponseBody
	Data WAF `json:"data"`
}

// GetWAFRuleSetSliceResponseBody represents an API response body containing []WAFRuleSet data
type GetWAFRuleSetSliceResponseBody struct {
	connection.APIResponseBody
	Data []WAFRuleSet `json:"data"`
}

// GetWAFRuleSetResponseBody represents an API response body containing WAFRuleSet data
type GetWAFRuleSetResponseBody struct {
	connection.APIResponseBody
	Data WAFRuleSet `json:"data"`
}

// GetWAFRuleSliceResponseBody represents an API response body containing []WAFRule data
type GetWAFRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []WAFRule `json:"data"`
}

// GetWAFRuleResponseBody represents an API response body containing WAFRule data
type GetWAFRuleResponseBody struct {
	connection.APIResponseBody
	Data WAFRule `json:"data"`
}

// GetWAFAdvancedRuleSliceResponseBody represents an API response body containing []WAFAdvancedRule data
type GetWAFAdvancedRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []WAFAdvancedRule `json:"data"`
}

// GetWAFAdvancedRuleResponseBody represents an API response body containing WAFAdvancedRule data
type GetWAFAdvancedRuleResponseBody struct {
	connection.APIResponseBody
	Data WAFAdvancedRule `json:"data"`
}

// GetSSLSliceResponseBody represents an API response body containing []SSL data
type GetSSLSliceResponseBody struct {
	connection.APIResponseBody
	Data []SSL `json:"data"`
}

// GetSSLResponseBody represents an API response body containing SSL data
type GetSSLResponseBody struct {
	connection.APIResponseBody
	Data SSL `json:"data"`
}

// GetSSLContentSliceResponseBody represents an API response body containing []SSLContent data
type GetSSLContentSliceResponseBody struct {
	connection.APIResponseBody
	Data []SSLContent `json:"data"`
}

// GetSSLContentResponseBody represents an API response body containing SSLContent data
type GetSSLContentResponseBody struct {
	connection.APIResponseBody
	Data SSLContent `json:"data"`
}

// GetSSLPrivateKeySliceResponseBody represents an API response body containing []SSLPrivateKey data
type GetSSLPrivateKeySliceResponseBody struct {
	connection.APIResponseBody
	Data []SSLPrivateKey `json:"data"`
}

// GetSSLPrivateKeyResponseBody represents an API response body containing SSLPrivateKey data
type GetSSLPrivateKeyResponseBody struct {
	connection.APIResponseBody
	Data SSLPrivateKey `json:"data"`
}

// GetACLGeoIPRuleSliceResponseBody represents an API response body containing []ACLGeoIPRule data
type GetACLGeoIPRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []ACLGeoIPRule `json:"data"`
}

// GetACLGeoIPRuleResponseBody represents an API response body containing ACLGeoIPRule data
type GetACLGeoIPRuleResponseBody struct {
	connection.APIResponseBody
	Data ACLGeoIPRule `json:"data"`
}

// GetACLIPRuleSliceResponseBody represents an API response body containing []ACLIPRule data
type GetACLIPRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []ACLIPRule `json:"data"`
}

// GetACLIPRuleResponseBody represents an API response body containing ACLIPRule data
type GetACLIPRuleResponseBody struct {
	connection.APIResponseBody
	Data ACLIPRule `json:"data"`
}

// GetCDNRuleSliceResponseBody represents an API response body containing []CDNRule data
type GetCDNRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []CDNRule `json:"data"`
}

// GetCDNRuleResponseBody represents an API response body containing CDNRule data
type GetCDNRuleResponseBody struct {
	connection.APIResponseBody
	Data CDNRule `json:"data"`
}

// GetHSTSConfigurationSliceResponseBody represents an API response body containing []HSTSConfiguration data
type GetHSTSConfigurationSliceResponseBody struct {
	connection.APIResponseBody
	Data []HSTSConfiguration `json:"data"`
}

// GetHSTSConfigurationResponseBody represents an API response body containing HSTSConfiguration data
type GetHSTSConfigurationResponseBody struct {
	connection.APIResponseBody
	Data HSTSConfiguration `json:"data"`
}

// GetHSTSRuleSliceResponseBody represents an API response body containing []HSTSRule data
type GetHSTSRuleSliceResponseBody struct {
	connection.APIResponseBody
	Data []HSTSRule `json:"data"`
}

// GetHSTSRuleResponseBody represents an API response body containing HSTSRule data
type GetHSTSRuleResponseBody struct {
	connection.APIResponseBody
	Data HSTSRule `json:"data"`
}

// GetWAFLogSliceResponseBody represents an API response body containing []WAFLog data
type GetWAFLogSliceResponseBody struct {
	connection.APIResponseBody
	Data []WAFLog `json:"data"`
}

// GetWAFLogResponseBody represents an API response body containing WAFLog data
type GetWAFLogResponseBody struct {
	connection.APIResponseBody
	Data WAFLog `json:"data"`
}

// GetWAFLogMatchSliceResponseBody represents an API response body containing []WAFLogMatch data
type GetWAFLogMatchSliceResponseBody struct {
	connection.APIResponseBody
	Data []WAFLogMatch `json:"data"`
}

// GetWAFLogMatchResponseBody represents an API response body containing WAFLogMatch data
type GetWAFLogMatchResponseBody struct {
	connection.APIResponseBody
	Data WAFLogMatch `json:"data"`
}
