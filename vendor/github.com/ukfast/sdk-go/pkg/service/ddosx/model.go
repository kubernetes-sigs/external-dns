//go:generate go run ../../gen/model_response/main.go -package ddosx -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package ddosx -source model.go -destination model_paginated_generated.go

package ddosx

import (
	"time"

	"github.com/ukfast/go-durationstring"
	"github.com/ukfast/sdk-go/pkg/connection"
)

type DomainStatus string

func (s DomainStatus) String() string {
	return string(s)
}

const (
	DomainStatusConfigured    DomainStatus = "Configured"
	DomainStatusNotConfigured DomainStatus = "Not Configured"
	DomainStatusPending       DomainStatus = "Pending"
	DomainStatusFailed        DomainStatus = "Failed"
	DomainStatusCancelling    DomainStatus = "Cancelling"
	DomainStatusCancelled     DomainStatus = "Cancelled"
)

type DomainPropertyName string

func (e DomainPropertyName) String() string {
	return string(e)
}

const (
	DomainPropertyNameClientMaxBodySize DomainPropertyName = "client_max_body_size"
	DomainPropertyNameProxyTimeout      DomainPropertyName = "proxy_timeout"
	DomainPropertyNameIPv6Enabled       DomainPropertyName = "ipv6_enabled"
	DomainPropertyNameSecureOrigin      DomainPropertyName = "secure_origin"
)

var DomainPropertyNameEnum connection.EnumSlice = []connection.Enum{
	DomainPropertyNameClientMaxBodySize,
	DomainPropertyNameProxyTimeout,
	DomainPropertyNameIPv6Enabled,
	DomainPropertyNameSecureOrigin,
}

// ParseDomainPropertyName attempts to parse a DomainPropertyName from string
func ParseDomainPropertyName(s string) (DomainPropertyName, error) {
	e, err := connection.ParseEnum(s, DomainPropertyNameEnum)
	if err != nil {
		return "", err
	}

	return e.(DomainPropertyName), err
}

type RecordType string

func (s RecordType) String() string {
	return string(s)
}

const (
	RecordTypeA    RecordType = "A"
	RecordTypeAAAA RecordType = "AAAA"
)

type WAFMode string

func (s WAFMode) String() string {
	return string(s)
}

const (
	WAFModeOn            WAFMode = "On"
	WAFModeOff           WAFMode = "Off"
	WAFModeDetectionOnly WAFMode = "DetectionOnly"
)

var WAFModeEnum connection.EnumSlice = []connection.Enum{
	WAFModeOn,
	WAFModeOff,
	WAFModeDetectionOnly,
}

// ParseWAFMode attempts to parse a WAFMode from string
func ParseWAFMode(s string) (WAFMode, error) {
	e, err := connection.ParseEnum(s, WAFModeEnum)
	if err != nil {
		return "", err
	}

	return e.(WAFMode), err
}

type WAFParanoiaLevel string

func (s WAFParanoiaLevel) String() string {
	return string(s)
}

const (
	WAFParanoiaLevelLow     WAFParanoiaLevel = "Low"
	WAFParanoiaLevelMedium  WAFParanoiaLevel = "Medium"
	WAFParanoiaLevelHigh    WAFParanoiaLevel = "High"
	WAFParanoiaLevelHighest WAFParanoiaLevel = "Highest"
)

var WAFParanoiaLevelEnum connection.EnumSlice = []connection.Enum{
	WAFParanoiaLevelLow,
	WAFParanoiaLevelMedium,
	WAFParanoiaLevelHigh,
	WAFParanoiaLevelHighest,
}

// ParseWAFParanoiaLevel attempts to parse a WAFParanoiaLevel from string
func ParseWAFParanoiaLevel(s string) (WAFParanoiaLevel, error) {
	e, err := connection.ParseEnum(s, WAFParanoiaLevelEnum)
	if err != nil {
		return "", err
	}

	return e.(WAFParanoiaLevel), err
}

type WAFRuleSetName string

func (s WAFRuleSetName) String() string {
	return string(s)
}

const (
	WAFRuleSetNameIPRepution                             WAFRuleSetName = "IP Reputation"
	WAFRuleSetNameMethodEnforcement                      WAFRuleSetName = "Method Enforcement"
	WAFRuleSetNameScannerDetection                       WAFRuleSetName = "Scanner Detection"
	WAFRuleSetNameProtocolEnforcement                    WAFRuleSetName = "Protocol Enforcement"
	WAFRuleSetNameProtocolAttack                         WAFRuleSetName = "Protocol Attack"
	WAFRuleSetNameApplicationAttackLocalFileInclusion    WAFRuleSetName = "Application Attack (Local File Inclusion)"
	WAFRuleSetNameApplicationAttackRemoteFileInclusion   WAFRuleSetName = "Application Attack (Remote File Inclusion)"
	WAFRuleSetNameApplicationAttackRemoteCodeExecution   WAFRuleSetName = "Application Attack (Remote Code Execution)"
	WAFRuleSetNameApplicationAttackPHP                   WAFRuleSetName = "Application Attack PHP"
	WAFRuleSetNameApplicationAttackXSSCrossSiteScripting WAFRuleSetName = "Application Attack XSS (Cross Site Scripting)"
	WAFRuleSetNameApplicationAttackSQLISQLInjection      WAFRuleSetName = "Application Attack SQLI (SQL Injection)"
	WAFRuleSetNameApplicationAttackSessionFixation       WAFRuleSetName = "Application Attack Session Fixation"
	WAFRuleSetNameDataDeakages                           WAFRuleSetName = "Data Leakages"
	WAFRuleSetNameDataLeakageSQL                         WAFRuleSetName = "Data Leakage SQL"
	WAFRuleSetNameDataLeakageJava                        WAFRuleSetName = "Data Leakage Java"
	WAFRuleSetNameDataLeakagePHP                         WAFRuleSetName = "Data Leakage PHP"
	WAFRuleSetNameDataLeakageIIS                         WAFRuleSetName = "Data Leakage IIS"
)

type WAFAdvancedRuleSection string

func (s WAFAdvancedRuleSection) String() string {
	return string(s)
}

const (
	WAFAdvancedRuleSectionArgs           WAFAdvancedRuleSection = "ARGS"
	WAFAdvancedRuleSectionMatchedVars    WAFAdvancedRuleSection = "MATCHED_VARS"
	WAFAdvancedRuleSectionRemoteHost     WAFAdvancedRuleSection = "REMOTE_HOST"
	WAFAdvancedRuleSectionRequestBody    WAFAdvancedRuleSection = "REQUEST_BODY"
	WAFAdvancedRuleSectionRequestCookies WAFAdvancedRuleSection = "REQUEST_COOKIES"
	WAFAdvancedRuleSectionRequestHeaders WAFAdvancedRuleSection = "REQUEST_HEADERS"
	WAFAdvancedRuleSectionRequestURI     WAFAdvancedRuleSection = "REQUEST_URI"
)

var WAFAdvancedRuleSectionEnum connection.EnumSlice = []connection.Enum{
	WAFAdvancedRuleSectionArgs,
	WAFAdvancedRuleSectionMatchedVars,
	WAFAdvancedRuleSectionRemoteHost,
	WAFAdvancedRuleSectionRequestBody,
	WAFAdvancedRuleSectionRequestCookies,
	WAFAdvancedRuleSectionRequestHeaders,
	WAFAdvancedRuleSectionRequestURI,
}

// ParseWAFAdvancedRuleSection attempts to parse a WAFAdvancedRuleSection from string
func ParseWAFAdvancedRuleSection(s string) (WAFAdvancedRuleSection, error) {
	e, err := connection.ParseEnum(s, WAFAdvancedRuleSectionEnum)
	if err != nil {
		return "", err
	}

	return e.(WAFAdvancedRuleSection), err
}

type WAFAdvancedRuleModifier string

func (s WAFAdvancedRuleModifier) String() string {
	return string(s)
}

const (
	WAFAdvancedRuleModifierBeginsWith   WAFAdvancedRuleModifier = "beginsWith"
	WAFAdvancedRuleModifierEndsWith     WAFAdvancedRuleModifier = "endsWith"
	WAFAdvancedRuleModifierContains     WAFAdvancedRuleModifier = "contains"
	WAFAdvancedRuleModifierContainsWord WAFAdvancedRuleModifier = "containsWord"
)

var WAFAdvancedRuleModifierEnum connection.EnumSlice = []connection.Enum{
	WAFAdvancedRuleModifierBeginsWith,
	WAFAdvancedRuleModifierEndsWith,
	WAFAdvancedRuleModifierContains,
	WAFAdvancedRuleModifierContainsWord,
}

// ParseWAFAdvancedRuleModifier attempts to parse a WAFAdvancedRuleModifier from string
func ParseWAFAdvancedRuleModifier(s string) (WAFAdvancedRuleModifier, error) {
	e, err := connection.ParseEnum(s, WAFAdvancedRuleModifierEnum)
	if err != nil {
		return "", err
	}

	return e.(WAFAdvancedRuleModifier), err
}

type ACLIPMode string

func (s ACLIPMode) String() string {
	return string(s)
}

const (
	ACLIPModeAllow ACLIPMode = "Allow"
	ACLIPModeDeny  ACLIPMode = "Deny"
)

var ACLIPModeEnum connection.EnumSlice = []connection.Enum{
	ACLIPModeAllow,
	ACLIPModeDeny,
}

// ParseACLIPMode attempts to parse a ACLIPMode from string
func ParseACLIPMode(s string) (ACLIPMode, error) {
	e, err := connection.ParseEnum(s, ACLIPModeEnum)
	if err != nil {
		return "", err
	}

	return e.(ACLIPMode), err
}

type ACLGeoIPRulesMode string

func (s ACLGeoIPRulesMode) String() string {
	return string(s)
}

const (
	ACLGeoIPRulesModeWhitelist ACLGeoIPRulesMode = "Whitelist"
	ACLGeoIPRulesModeBlacklist ACLGeoIPRulesMode = "Blacklist"
)

var ACLGeoIPRulesModeEnum connection.EnumSlice = []connection.Enum{
	ACLGeoIPRulesModeWhitelist,
	ACLGeoIPRulesModeBlacklist,
}

// ParseACLGeoIPRulesMode attempts to parse a ACLGeoIPRulesMode from string
func ParseACLGeoIPRulesMode(s string) (ACLGeoIPRulesMode, error) {
	e, err := connection.ParseEnum(s, ACLGeoIPRulesModeEnum)
	if err != nil {
		return "", err
	}

	return e.(ACLGeoIPRulesMode), err
}

type CDNRuleCacheControl string

func (e CDNRuleCacheControl) String() string {
	return string(e)
}

const (
	CDNRuleCacheControlCustom CDNRuleCacheControl = "Custom"
	CDNRuleCacheControlOrigin CDNRuleCacheControl = "Origin"
)

var CDNRuleCacheControlEnum connection.EnumSlice = []connection.Enum{
	CDNRuleCacheControlCustom,
	CDNRuleCacheControlOrigin,
}

// ParseCDNRuleCacheControl attempts to parse a CDNRuleCacheControl from string
func ParseCDNRuleCacheControl(s string) (CDNRuleCacheControl, error) {
	e, err := connection.ParseEnum(s, CDNRuleCacheControlEnum)
	if err != nil {
		return "", err
	}

	return e.(CDNRuleCacheControl), err
}

type CDNRuleType string

func (e CDNRuleType) String() string {
	return string(e)
}

const (
	CDNRuleTypeGlobal CDNRuleType = "global"
	CDNRuleTypePerURI CDNRuleType = "per-uri"
)

var CDNRuleTypeEnum connection.EnumSlice = []connection.Enum{CDNRuleTypeGlobal, CDNRuleTypePerURI}

// ParseCDNRuleType attempts to parse a CDNRuleType from string
func ParseCDNRuleType(s string) (CDNRuleType, error) {
	e, err := connection.ParseEnum(s, CDNRuleTypeEnum)
	if err != nil {
		return "", err
	}

	return e.(CDNRuleType), err
}

type HSTSRuleType string

func (e HSTSRuleType) String() string {
	return string(e)
}

const (
	HSTSRuleTypeDomain HSTSRuleType = "domain"
	HSTSRuleTypeRecord HSTSRuleType = "record"
)

var HSTSRuleTypeEnum connection.EnumSlice = []connection.Enum{HSTSRuleTypeDomain, HSTSRuleTypeRecord}

// ParseHSTSRuleType attempts to parse a HSTSRuleType from string
func ParseHSTSRuleType(s string) (HSTSRuleType, error) {
	e, err := connection.ParseEnum(s, HSTSRuleTypeEnum)
	if err != nil {
		return "", err
	}

	return e.(HSTSRuleType), err
}

// Domain represents a DDoSX domain
// +genie:model_response
// +genie:model_paginated
type Domain struct {
	SafeDNSZoneID *int               `json:"safedns_zone_id"`
	Name          string             `json:"name"`
	Status        DomainStatus       `json:"status"`
	DNSActive     bool               `json:"dns_active"`
	CDNActive     bool               `json:"cdn_active"`
	WAFActive     bool               `json:"waf_active"`
	ExternalDNS   *DomainExternalDNS `json:"external_dns"`
}

// DomainExternalDNS represents a DDoSX domain external DNS configuration
type DomainExternalDNS struct {
	Verified           bool   `json:"verified"`
	VerificationString string `json:"verification_string"`
	Target             string `json:"target"`
}

// DomainProperty represents a DDoSX domain property
// +genie:model_response
// +genie:model_paginated
type DomainProperty struct {
	ID    string             `json:"id"`
	Name  DomainPropertyName `json:"name"`
	Value interface{}        `json:"value"`
}

// Record represents a DDoSX record
// +genie:model_response
// +genie:model_paginated
type Record struct {
	ID              string     `json:"id"`
	DomainName      string     `json:"domain_name"`
	SafeDNSRecordID *int       `json:"safedns_record_id"`
	SSLID           *string    `json:"ssl_id"`
	Name            string     `json:"name"`
	Type            RecordType `json:"type"`
	Content         string     `json:"content"`
}

// WAF represents a DDoSX WAF configuration
// +genie:model_response
type WAF struct {
	Mode          WAFMode          `json:"mode"`
	ParanoiaLevel WAFParanoiaLevel `json:"paranoia_level"`
}

// WAFRuleSet represents a DDoSX WAF rule set
// +genie:model_response
// +genie:model_paginated
type WAFRuleSet struct {
	ID     string         `json:"id"`
	Name   WAFRuleSetName `json:"name"`
	Active bool           `json:"active"`
}

// WAFRule represents a DDoSX WAF rule
// +genie:model_response
// +genie:model_paginated
type WAFRule struct {
	ID  string               `json:"id"`
	URI string               `json:"uri"`
	IP  connection.IPAddress `json:"ip"`
}

// WAFAdvancedRule represents a DDoSX WAF advanced rule
// +genie:model_response
// +genie:model_paginated
type WAFAdvancedRule struct {
	ID       string                  `json:"id"`
	Section  WAFAdvancedRuleSection  `json:"section"`
	Modifier WAFAdvancedRuleModifier `json:"modifier"`
	Phrase   string                  `json:"phrase"`
	IP       connection.IPAddress    `json:"ip"`
}

// SSL represents a DDoSX SSL
// +genie:model_response
// +genie:model_paginated
type SSL struct {
	ID           string   `json:"id"`
	UKFastSSLID  *int     `json:"ukfast_ssl_id"`
	Domains      []string `json:"domains"`
	FriendlyName string   `json:"friendly_name"`
}

// SSLContent represents a DDoSX SSL content
// +genie:model_response
type SSLContent struct {
	Certificate string `json:"certificate"`
	CABundle    string `json:"ca_bundle"`
}

// SSLPrivateKey represents a DDoSX SSL private key
// +genie:model_response
type SSLPrivateKey struct {
	Key string `json:"key"`
}

// ACLGeoIPRule represents a DDoSX ACL GeoIP rule
// +genie:model_response
// +genie:model_paginated
type ACLGeoIPRule struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// ACLIPRule represents a DDoSX ACL IP rule
// +genie:model_response
// +genie:model_paginated
type ACLIPRule struct {
	ID   string               `json:"id"`
	IP   connection.IPAddress `json:"ip"`
	URI  string               `json:"uri"`
	Mode ACLIPMode            `json:"mode"`
}

// CDNRule represents a DDoSX CDN rule
// +genie:model_response
// +genie:model_paginated
type CDNRule struct {
	ID           string              `json:"id"`
	URI          string              `json:"uri"`
	CacheControl CDNRuleCacheControl `json:"cache_control"`
	// CacheControlDuration specifies the cache control duration. May be nil if duration not applicable
	CacheControlDuration CDNRuleCacheControlDuration `json:"cache_control_duration"`
	MimeTypes            []string                    `json:"mime_types"`
	Type                 CDNRuleType                 `json:"type"`
}

// CDNRuleCacheControlDuration represents a DDoSX CDN rule duration
type CDNRuleCacheControlDuration struct {
	Years   int `json:"years"`
	Months  int `json:"months"`
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

// Duration returns the cache control duration as time.Duration
func (d *CDNRuleCacheControlDuration) Duration() time.Duration {
	day := time.Hour * 24
	td := time.Duration(d.Years) * day * 365
	td = td + time.Duration(d.Months)*day*30
	td = td + time.Duration(d.Days)*day
	td = td + time.Duration(d.Hours)*time.Hour
	return td + time.Duration(d.Minutes)*time.Minute
}

func (d *CDNRuleCacheControlDuration) String() string {
	return durationstring.NewDuration(d.Years, d.Months, d.Days, d.Hours, d.Minutes, 0, 0, 0, 0).String()
}

// ParseCDNRuleCacheControlDuration parses string s and returns a pointer to an
// initialised CDNRuleCacheControlDuration
func ParseCDNRuleCacheControlDuration(s string) (*CDNRuleCacheControlDuration, error) {
	d, err := durationstring.Parse(s)
	if err != nil {
		return nil, err
	}

	return &CDNRuleCacheControlDuration{
		Years:   d.Years,
		Months:  d.Months,
		Days:    d.Days,
		Hours:   d.Hours,
		Minutes: d.Minutes,
	}, nil
}

// HSTSConfiguration represents HSTS configuration for a DDoSX domain
// +genie:model_response
type HSTSConfiguration struct {
	Enabled bool `json:"enabled"`
}

// HSTSRule represents HSTS rule for a DDoSX domain
// +genie:model_response
// +genie:model_paginated
type HSTSRule struct {
	ID                string       `json:"id"`
	MaxAge            int          `json:"max_age"`
	Preload           bool         `json:"preload"`
	IncludeSubdomains bool         `json:"include_subdomains"`
	Type              HSTSRuleType `json:"type"`
	RecordName        *string      `json:"record_name"`
}

// WAFLog represents a WAF log entry
// +genie:model_response
// +genie:model_paginated
type WAFLog struct {
	ID        string               `json:"id"`
	Host      string               `json:"host"`
	ClientIP  connection.IPAddress `json:"client_ip"`
	Request   string               `json:"request"`
	CreatedAt connection.DateTime  `json:"created_at"`
}

// WAFLogMatch represents a WAF log match
// +genie:model_response
// +genie:model_paginated
type WAFLogMatch struct {
	ID          string               `json:"id"`
	LogID       string               `json:"log_id"`
	Host        string               `json:"host"`
	ClientIP    connection.IPAddress `json:"client_ip"`
	RequestURI  string               `json:"request_uri"`
	CreatedAt   connection.DateTime  `json:"created_at"`
	CountryCode string               `json:"country_code"`
	Method      string               `json:"method"`
	Content     string               `json:"content"`
	Message     string               `json:"message"`
	MatchData   string               `json:"match_data"`
	URIPart     string               `json:"uri_part"`
	Value       string               `json:"value"`
}
