package ddosx

import "fmt"

// DomainNotFoundError indicates a domain was not found
type DomainNotFoundError struct {
	Name string
}

func (e *DomainNotFoundError) Error() string {
	return fmt.Sprintf("Domain not found with name [%s]", e.Name)
}

// DomainAlreadyVerifiedError indicates a domain is already verified
type DomainAlreadyVerifiedError struct {
	Name string
}

func (e *DomainAlreadyVerifiedError) Error() string {
	return fmt.Sprintf("Domain already verified with name [%s]", e.Name)
}

// DomainPropertyNotFoundError indicates a domain property was not found
type DomainPropertyNotFoundError struct {
	ID string
}

func (e *DomainPropertyNotFoundError) Error() string {
	return fmt.Sprintf("Domain property not found with uuid [%s]", e.ID)
}

// RecordNotFoundError indicates a Record was not found
type RecordNotFoundError struct {
	ID string
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("Record not found with id [%s]", e.ID)
}

// DomainRecordNotFoundError indicates a Domain Record was not found
type DomainRecordNotFoundError struct {
	DomainName string
	ID         string
}

func (e *DomainRecordNotFoundError) Error() string {
	return fmt.Sprintf("Record not found with id [%s] for domain [%s]", e.ID, e.DomainName)
}

// DomainWAFNotFoundError indicates a WAF configuration was not found for domain
type DomainWAFNotFoundError struct {
	DomainName string
}

func (e *DomainWAFNotFoundError) Error() string {
	return fmt.Sprintf("WAF configuration not found for domain [%s]", e.DomainName)
}

// SSLNotFoundError indicates an SSL was not found
type SSLNotFoundError struct {
	ID string
}

func (e *SSLNotFoundError) Error() string {
	return fmt.Sprintf("SSL not found with id [%s]", e.ID)
}

// ACLGeoIPRuleNotFoundError indicates an ACL GeoIP rule was not found
type ACLGeoIPRuleNotFoundError struct {
	ID string
}

func (e *ACLGeoIPRuleNotFoundError) Error() string {
	return fmt.Sprintf("ACL GeoIP rule not found with id [%s]", e.ID)
}

// ACLIPRuleNotFoundError indicates an ACL IP rule was not found
type ACLIPRuleNotFoundError struct {
	ID string
}

func (e *ACLIPRuleNotFoundError) Error() string {
	return fmt.Sprintf("ACL IP rule not found with id [%s]", e.ID)
}

// WAFRuleSetNotFoundError indicates a WAF rule set was not found
type WAFRuleSetNotFoundError struct {
	ID string
}

func (e *WAFRuleSetNotFoundError) Error() string {
	return fmt.Sprintf("WAF rule set not found with id [%s]", e.ID)
}

// WAFRuleNotFoundError indicates a WAF rule was not found
type WAFRuleNotFoundError struct {
	ID string
}

func (e *WAFRuleNotFoundError) Error() string {
	return fmt.Sprintf("WAF rule not found with id [%s]", e.ID)
}

// WAFAdvancedRuleNotFoundError indicates a WAF advanced rule was not found
type WAFAdvancedRuleNotFoundError struct {
	ID string
}

func (e *WAFAdvancedRuleNotFoundError) Error() string {
	return fmt.Sprintf("WAF rule not found with id [%s]", e.ID)
}

// DomainCDNConfigurationNotFoundError indicates CDN configuration was not found for domain
type DomainCDNConfigurationNotFoundError struct {
	DomainName string
}

func (e *DomainCDNConfigurationNotFoundError) Error() string {
	return fmt.Sprintf("CDN configuration not found for domain [%s]", e.DomainName)
}

// CDNRuleNotFoundError indicates a CDN rule was not found
type CDNRuleNotFoundError struct {
	ID string
}

func (e *CDNRuleNotFoundError) Error() string {
	return fmt.Sprintf("CDN rule not found with id [%s]", e.ID)
}

// DomainHSTSConfigurationNotFoundError indicates HSTS configuration was not found for domain
type DomainHSTSConfigurationNotFoundError struct {
	DomainName string
}

func (e *DomainHSTSConfigurationNotFoundError) Error() string {
	return fmt.Sprintf("HSTS configuration not found for domain [%s]", e.DomainName)
}

// HSTSRuleNotFoundError indicates a HSTS rule was not found
type HSTSRuleNotFoundError struct {
	ID string
}

func (e *HSTSRuleNotFoundError) Error() string {
	return fmt.Sprintf("HSTS rule not found with id [%s]", e.ID)
}

// WAFLogNotFoundError indicates a WAF rule was not found
type WAFLogNotFoundError struct {
	ID string
}

func (e *WAFLogNotFoundError) Error() string {
	return fmt.Sprintf("WAF log not found with id [%s]", e.ID)
}

// WAFLogMatchNotFoundError indicates a WAF rule was not found
type WAFLogMatchNotFoundError struct {
	ID string
}

func (e *WAFLogMatchNotFoundError) Error() string {
	return fmt.Sprintf("WAF log match not found with id [%s]", e.ID)
}
