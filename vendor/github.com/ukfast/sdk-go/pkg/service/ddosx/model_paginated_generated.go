package ddosx

import "github.com/ukfast/sdk-go/pkg/connection"

// PaginatedDomain represents a paginated collection of Domain
type PaginatedDomain struct {
	*connection.PaginatedBase
	Items []Domain
}

// NewPaginatedDomain returns a pointer to an initialized PaginatedDomain struct
func NewPaginatedDomain(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Domain) *PaginatedDomain {
	return &PaginatedDomain{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedDomainProperty represents a paginated collection of DomainProperty
type PaginatedDomainProperty struct {
	*connection.PaginatedBase
	Items []DomainProperty
}

// NewPaginatedDomainProperty returns a pointer to an initialized PaginatedDomainProperty struct
func NewPaginatedDomainProperty(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []DomainProperty) *PaginatedDomainProperty {
	return &PaginatedDomainProperty{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedRecord represents a paginated collection of Record
type PaginatedRecord struct {
	*connection.PaginatedBase
	Items []Record
}

// NewPaginatedRecord returns a pointer to an initialized PaginatedRecord struct
func NewPaginatedRecord(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []Record) *PaginatedRecord {
	return &PaginatedRecord{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedWAFRuleSet represents a paginated collection of WAFRuleSet
type PaginatedWAFRuleSet struct {
	*connection.PaginatedBase
	Items []WAFRuleSet
}

// NewPaginatedWAFRuleSet returns a pointer to an initialized PaginatedWAFRuleSet struct
func NewPaginatedWAFRuleSet(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []WAFRuleSet) *PaginatedWAFRuleSet {
	return &PaginatedWAFRuleSet{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedWAFRule represents a paginated collection of WAFRule
type PaginatedWAFRule struct {
	*connection.PaginatedBase
	Items []WAFRule
}

// NewPaginatedWAFRule returns a pointer to an initialized PaginatedWAFRule struct
func NewPaginatedWAFRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []WAFRule) *PaginatedWAFRule {
	return &PaginatedWAFRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedWAFAdvancedRule represents a paginated collection of WAFAdvancedRule
type PaginatedWAFAdvancedRule struct {
	*connection.PaginatedBase
	Items []WAFAdvancedRule
}

// NewPaginatedWAFAdvancedRule returns a pointer to an initialized PaginatedWAFAdvancedRule struct
func NewPaginatedWAFAdvancedRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []WAFAdvancedRule) *PaginatedWAFAdvancedRule {
	return &PaginatedWAFAdvancedRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedSSL represents a paginated collection of SSL
type PaginatedSSL struct {
	*connection.PaginatedBase
	Items []SSL
}

// NewPaginatedSSL returns a pointer to an initialized PaginatedSSL struct
func NewPaginatedSSL(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []SSL) *PaginatedSSL {
	return &PaginatedSSL{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedACLGeoIPRule represents a paginated collection of ACLGeoIPRule
type PaginatedACLGeoIPRule struct {
	*connection.PaginatedBase
	Items []ACLGeoIPRule
}

// NewPaginatedACLGeoIPRule returns a pointer to an initialized PaginatedACLGeoIPRule struct
func NewPaginatedACLGeoIPRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ACLGeoIPRule) *PaginatedACLGeoIPRule {
	return &PaginatedACLGeoIPRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedACLIPRule represents a paginated collection of ACLIPRule
type PaginatedACLIPRule struct {
	*connection.PaginatedBase
	Items []ACLIPRule
}

// NewPaginatedACLIPRule returns a pointer to an initialized PaginatedACLIPRule struct
func NewPaginatedACLIPRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []ACLIPRule) *PaginatedACLIPRule {
	return &PaginatedACLIPRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedCDNRule represents a paginated collection of CDNRule
type PaginatedCDNRule struct {
	*connection.PaginatedBase
	Items []CDNRule
}

// NewPaginatedCDNRule returns a pointer to an initialized PaginatedCDNRule struct
func NewPaginatedCDNRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []CDNRule) *PaginatedCDNRule {
	return &PaginatedCDNRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedHSTSRule represents a paginated collection of HSTSRule
type PaginatedHSTSRule struct {
	*connection.PaginatedBase
	Items []HSTSRule
}

// NewPaginatedHSTSRule returns a pointer to an initialized PaginatedHSTSRule struct
func NewPaginatedHSTSRule(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []HSTSRule) *PaginatedHSTSRule {
	return &PaginatedHSTSRule{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedWAFLog represents a paginated collection of WAFLog
type PaginatedWAFLog struct {
	*connection.PaginatedBase
	Items []WAFLog
}

// NewPaginatedWAFLog returns a pointer to an initialized PaginatedWAFLog struct
func NewPaginatedWAFLog(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []WAFLog) *PaginatedWAFLog {
	return &PaginatedWAFLog{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}

// PaginatedWAFLogMatch represents a paginated collection of WAFLogMatch
type PaginatedWAFLogMatch struct {
	*connection.PaginatedBase
	Items []WAFLogMatch
}

// NewPaginatedWAFLogMatch returns a pointer to an initialized PaginatedWAFLogMatch struct
func NewPaginatedWAFLogMatch(getFunc connection.PaginatedGetFunc, parameters connection.APIRequestParameters, pagination connection.APIResponseMetadataPagination, items []WAFLogMatch) *PaginatedWAFLogMatch {
	return &PaginatedWAFLogMatch{
		Items:         items,
		PaginatedBase: connection.NewPaginatedBase(parameters, pagination, getFunc),
	}
}
