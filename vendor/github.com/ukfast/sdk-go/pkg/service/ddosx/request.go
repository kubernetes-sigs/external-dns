package ddosx

import (
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/ptr"
)

// CreateRecordRequest represents a request to create a DDoSX record
type CreateRecordRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name            string     `json:"name" validate:"required"`
	SafeDNSRecordID int        `json:"safedns_record_id,omitempty"`
	SSLID           string     `json:"ssl_id,omitempty"`
	Type            RecordType `json:"type,omitempty"`
	Content         string     `json:"content,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateRecordRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateDomainRequest represents a request to create a DDoSX domain
type CreateDomainRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name string `json:"name" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateDomainRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// DeleteDomainRequest represents a DDoSX domain removal request
type DeleteDomainRequest struct {
	Summary     string `json:"summary"`
	Description string `json:"description"`
}

// PatchDomainPropertyRequest represents a DDoSX Domain Property patch request
type PatchDomainPropertyRequest struct {
	Value interface{} `json:"value,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchDomainPropertyRequest) Validate() *connection.ValidationError {
	return nil
}

// PatchRecordRequest represents a DDoSX Record patch request
type PatchRecordRequest struct {
	SafeDNSRecordID int        `json:"safedns_record_id,omitempty"`
	SSLID           string     `json:"ssl_id,omitempty"`
	Name            string     `json:"name,omitempty"`
	Type            RecordType `json:"type,omitempty"`
	Content         string     `json:"content,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchRecordRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateACLGeoIPRuleRequest represents a DDoSX GeoIP ACL rule
type CreateACLGeoIPRuleRequest struct {
	connection.APIRequestBodyDefaultValidator

	Code string `json:"code" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateACLGeoIPRuleRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchACLGeoIPRuleRequest represents a DDoSX GeoIP ACL rule patch request
type PatchACLGeoIPRuleRequest struct {
	connection.APIRequestBodyDefaultValidator

	Code string `json:"code,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchACLGeoIPRuleRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateACLIPRuleRequest represents a DDoSX IP ACL rule
type CreateACLIPRuleRequest struct {
	connection.APIRequestBodyDefaultValidator

	IP   connection.IPAddress `json:"ip" validate:"required"`
	URI  string               `json:"uri"`
	Mode ACLIPMode            `json:"mode" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateACLIPRuleRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchACLIPRuleRequest represents a DDoSX IP ACL rule patch request
type PatchACLIPRuleRequest struct {
	IP   connection.IPAddress `json:"ip,omitempty"`
	URI  *string              `json:"uri,omitempty"`
	Mode ACLIPMode            `json:"mode,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchACLIPRuleRequest) Validate() *connection.ValidationError {
	return nil
}

// PatchACLGeoIPRulesModeRequest represents a DDoSX IP ACL rule mode patch request
type PatchACLGeoIPRulesModeRequest struct {
	Mode ACLGeoIPRulesMode `json:"mode,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchACLGeoIPRulesModeRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateWAFRequest represents a DDoSX WAF create request
type CreateWAFRequest struct {
	connection.APIRequestBodyDefaultValidator

	Mode          WAFMode          `json:"mode" validate:"required"`
	ParanoiaLevel WAFParanoiaLevel `json:"paranoia_level" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateWAFRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchWAFRequest represents a DDoSX WAF patch request
type PatchWAFRequest struct {
	Mode          WAFMode          `json:"mode,omitempty"`
	ParanoiaLevel WAFParanoiaLevel `json:"paranoia_level,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchWAFRequest) Validate() *connection.ValidationError {
	return nil
}

// PatchWAFRuleSetRequest represents a DDoSX WAF rule set patch request
type PatchWAFRuleSetRequest struct {
	Active *bool `json:"active,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchWAFRuleSetRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateWAFRuleRequest represents a DDoSX WAF rule create request
type CreateWAFRuleRequest struct {
	connection.APIRequestBodyDefaultValidator

	URI string               `json:"uri" validate:"required"`
	IP  connection.IPAddress `json:"ip" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateWAFRuleRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchWAFRuleRequest represents a DDoSX WAF rule patch request
type PatchWAFRuleRequest struct {
	URI string               `json:"uri,omitempty"`
	IP  connection.IPAddress `json:"ip,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchWAFRuleRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateWAFAdvancedRuleRequest represents a DDoSX WAF advanced rule create request
type CreateWAFAdvancedRuleRequest struct {
	connection.APIRequestBodyDefaultValidator

	Section  WAFAdvancedRuleSection  `json:"section" validate:"required"`
	Modifier WAFAdvancedRuleModifier `json:"modifier" validate:"required"`
	Phrase   string                  `json:"phrase" validate:"required"`
	IP       connection.IPAddress    `json:"ip" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateWAFAdvancedRuleRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchWAFAdvancedRuleRequest represents a DDoSX WAF advanced rule patch request
type PatchWAFAdvancedRuleRequest struct {
	Section  WAFAdvancedRuleSection  `json:"section,omitempty"`
	Modifier WAFAdvancedRuleModifier `json:"modifier,omitempty"`
	Phrase   string                  `json:"phrase,omitempty"`
	IP       connection.IPAddress    `json:"ip,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchWAFAdvancedRuleRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateSSLRequest represents a DDoSX SSL create request
type CreateSSLRequest struct {
	connection.APIRequestBodyDefaultValidator

	FriendlyName string `json:"friendly_name" validate:"required"`
	UKFastSSLID  int    `json:"ukfast_ssl_id,omitempty"`
	Key          string `json:"key,omitempty"`
	Certificate  string `json:"certificate,omitempty"`
	CABundle     string `json:"ca_bundle,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateSSLRequest) Validate() *connection.ValidationError {
	if c.UKFastSSLID < 1 {
		if c.Key == "" {
			return connection.NewValidationError("Key must be provided when UKFastSSLID isn't provided")
		}
		if c.Certificate == "" {
			return connection.NewValidationError("Certificate must be provided when UKFastSSLID isn't provided")
		}
	}

	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchSSLRequest represents a DDoSX SSL create request
type PatchSSLRequest struct {
	FriendlyName string `json:"friendly_name,omitempty"`
	UKFastSSLID  int    `json:"ukfast_ssl_id,omitempty"`
	Key          string `json:"key,omitempty"`
	Certificate  string `json:"certificate,omitempty"`
	CABundle     string `json:"ca_bundle,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchSSLRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateCDNRuleRequest represents a DDoSX CDN rule create request
type CreateCDNRuleRequest struct {
	connection.APIRequestBodyDefaultValidator

	URI                  string                       `json:"uri" validate:"required"`
	CacheControl         CDNRuleCacheControl          `json:"cache_control" validate:"required"`
	CacheControlDuration *CDNRuleCacheControlDuration `json:"cache_control_duration,omitempty"`
	MimeTypes            []string                     `json:"mime_types" validate:"required"`
	Type                 CDNRuleType                  `json:"type" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateCDNRuleRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchCDNRuleRequest represents a DDoSX CDN rule patch request
type PatchCDNRuleRequest struct {
	URI                  string                       `json:"uri,omitempty"`
	CacheControl         CDNRuleCacheControl          `json:"cache_control,omitempty"`
	CacheControlDuration *CDNRuleCacheControlDuration `json:"cache_control_duration,omitempty"`
	MimeTypes            []string                     `json:"mime_types,omitempty"`
	Type                 CDNRuleType                  `json:"type,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchCDNRuleRequest) Validate() *connection.ValidationError {
	return nil
}

// PurgeCDNRequest represents a DDoSX CDN purge request
type PurgeCDNRequest struct {
	connection.APIRequestBodyDefaultValidator

	RecordName string `json:"record_name" validate:"required"`
	URI        string `json:"uri" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PurgeCDNRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateHSTSRuleRequest represents a DDoSX HSTS rule create request
type CreateHSTSRuleRequest struct {
	connection.APIRequestBodyDefaultValidator

	MaxAge            int          `json:"max_age"`
	Preload           bool         `json:"preload"`
	IncludeSubdomains bool         `json:"include_subdomains"`
	Type              HSTSRuleType `json:"type" validate:"required"`
	RecordName        *string      `json:"record_name,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateHSTSRuleRequest) Validate() *connection.ValidationError {
	if c.Type == HSTSRuleTypeRecord && ptr.ToStringOrDefault(c.RecordName) == "" {
		return connection.NewValidationError("RecordName must be specified with Type 'HSTSRuleTypeRecord'")
	}

	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchHSTSRuleRequest represents a DDoSX HSTS rule patch request
type PatchHSTSRuleRequest struct {
	MaxAge            *int  `json:"max_age,omitempty"`
	Preload           *bool `json:"preload,omitempty"`
	IncludeSubdomains *bool `json:"include_subdomains,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchHSTSRuleRequest) Validate() *connection.ValidationError {
	return nil
}
