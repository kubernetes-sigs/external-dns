package domain

import (
<<<<<<< HEAD
	"time"

	"github.com/go-gandi/go-gandi/internal/client"
)

// Domain is the API client to the Gandi v5 Domain API
type Domain struct {
	client client.Gandi
}

// New returns an instance of the Domain API client
func New(apikey string, sharingid string, debug bool, dryRun bool) *Domain {
	client := client.New(apikey, sharingid, debug, dryRun)
	client.SetEndpoint("domain/")
	return &Domain{client: *client}
}

// NewFromClient returns an instance of the Domain API client
func NewFromClient(g client.Gandi) *Domain {
	g.SetEndpoint("domain/")
	return &Domain{client: g}
}

// Contact represents a contact associated with a domain
type Contact struct {
	Country         string                 `json:"country"`
	Email           string                 `json:"email"`
	FamilyName      string                 `json:"family"`
	GivenName       string                 `json:"given"`
	StreetAddr      string                 `json:"streetaddr"`
	ContactType     int                    `json:"type"`
	BrandNumber     string                 `json:"brand_number,omitempty"`
	City            string                 `json:"city,omitempty"`
	DataObfuscated  *bool                  `json:"data_obfuscated,omitempty"`
	Fax             string                 `json:"fax,omitempty"`
	Language        string                 `json:"lang,omitempty"`
	MailObfuscated  *bool                  `json:"mail_obfuscated,omitempty"`
	Mobile          string                 `json:"mobile,omitempty"`
	OrgName         string                 `json:"orgname,omitempty"`
	Phone           string                 `json:"phone,omitempty"`
	Siren           string                 `json:"siren,omitempty"`
	State           string                 `json:"state,omitempty"`
	Validation      string                 `json:"validation,omitempty"`
	Zip             string                 `json:"zip,omitempty"`
	ExtraParameters map[string]interface{} `json:"extra_parameters,omitempty"`
}

// ResponseDates represents all the dates associated with a domain
type ResponseDates struct {
	RegistryCreatedAt   *time.Time `json:"registry_created_at"`
	UpdatedAt           *time.Time `json:"updated_at"`
	AuthInfoExpiresAt   *time.Time `json:"authinfo_expires_at,omitempty"`
	CreatedAt           *time.Time `json:"created_at,omitempty"`
	DeletesAt           *time.Time `json:"deletes_at,omitempty"`
	HoldBeginsAt        *time.Time `json:"hold_begins_at,omitempty"`
	HoldEndsAt          *time.Time `json:"hold_ends_at,omitempty"`
	PendingDeleteEndsAt *time.Time `json:"pending_delete_ends_at,omitempty"`
	RegistryEndsAt      *time.Time `json:"registry_ends_at,omitempty"`
	RenewBeginsAt       *time.Time `json:"renew_begins_at,omitempty"`
	RenewEndsAt         *time.Time `json:"renew_ends_at,omitempty"`
}

// NameServerConfig represents the name server configuration for a domain
type NameServerConfig struct {
	Current string   `json:"current"`
	Hosts   []string `json:"hosts,omitempty"`
}

// ListResponse is the response object returned by listing domains
type ListResponse struct {
	AutoRenew   *bool             `json:"autorenew"`
	Dates       *ResponseDates    `json:"dates"`
	DomainOwner string            `json:"domain_owner"`
	FQDN        string            `json:"fqdn"`
	FQDNUnicode string            `json:"fqdn_unicode"`
	Href        string            `json:"href"`
	ID          string            `json:"id"`
	NameServer  *NameServerConfig `json:"nameserver"`
	OrgaOwner   string            `json:"orga_owner"`
	Owner       string            `json:"owner"`
	Status      []string          `json:"status"`
	TLD         string            `json:"tld"`
	SharingID   string            `json:"sharing_id,omitempty"`
	Tags        []string          `json:"tags,omitempty"`
}

// AutoRenew is the auto renewal information for the domain
type AutoRenew struct {
	Href     string       `json:"href,omitempty"`
	Dates    *[]time.Time `json:"dates,omitempty"`
	Duration int          `json:"duration,omitempty"`
	Enabled  *bool        `json:"enabled,omitempty"`
	OrgID    string       `json:"org_id,omitempty"`
}

// SharingSpace defines the Organisation that owns the domain
type SharingSpace struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Details describes a single domain
type Details struct {
	AutoRenew    *AutoRenew     `json:"autorenew"`
	CanTLDLock   *bool          `json:"can_tld_lock"`
	Contacts     *Contacts      `json:"contacts"`
	Dates        *ResponseDates `json:"dates"`
	FQDN         string         `json:"fqdn"`
	FQDNUnicode  string         `json:"fqdn_unicode"`
	Href         string         `json:"href"`
	Nameservers  []string       `json:"nameservers,omitempty"`
	Services     []string       `json:"services"`
	SharingSpace *SharingSpace  `json:"sharing_space"`
	Status       []string       `json:"status"`
	TLD          string         `json:"tld"`
	AuthInfo     string         `json:"authinfo,omitempty"`
	ID           string         `json:"id,omitempty"`
	SharingID    string         `json:"sharing_id,omitempty"`
	Tags         []string       `json:"tags,omitempty"`
	TrusteeRoles []string       `json:"trustee_roles,omitempty"`
}

// CreateRequest is used to request a new domain
type CreateRequest struct {
	FQDN     string   `json:"fqdn"`
	Owner    *Contact `json:"owner"`
	Admin    *Contact `json:"admin,omitempty"`
	Billing  *Contact `json:"bill,omitempty"`
	Claims   string   `json:"claims,omitempty"`
	Currency string   `json:"currency,omitempty"`
	// Duration in years between 1 and 10
	Duration       int    `json:"duration,omitempty"`
	EnforcePremium bool   `json:"enforce_premium,omitempty"`
	Lang           string `json:"lang,omitempty"`
	// NameserverIPs sets the Glue Records for the domain
	NameserverIPs map[string]string `json:"nameserver_ips,omitempty"`
	Nameservers   []string          `json:"nameservers,omitempty"`
	Price         int               `json:"price,omitempty"`
	ReselleeID    string            `json:"resellee_id,omitempty"`
	// SMD is a Signed Mark Data file; if used, `TLDPeriod` must be "sunrise"
	SMD       string   `json:"smd,omitempty"`
	Tech      *Contact `json:"tech,omitempty"`
	TLDPeriod string   `json:"tld_period,omitempty"`
}

// Contacts is the set of contacts associated with a Domain
type Contacts struct {
	Admin   *Contact `json:"admin,omitempty"`
	Billing *Contact `json:"bill,omitempty"`
	Owner   *Contact `json:"owner,omitempty"`
	Tech    *Contact `json:"tech,omitempty"`
}

// Nameservers represents a list of nameservers
type Nameservers struct {
	Nameservers []string `json:"nameservers,omitempty"`
}

// ListDomains requests the set of Domains
// It returns a slice of domains and any error encountered
func (g *Domain) ListDomains() (domains []ListResponse, err error) {
	_, err = g.client.Get("domains", nil, &domains)
	return
}

// GetDomain requests a single Domain
// It returns a Details object and any error encountered
func (g *Domain) GetDomain(domain string) (domainResponse Details, err error) {
	_, err = g.client.Get("domains/"+domain, nil, &domainResponse)
	return
}

// CreateDomain creates a single Domain
func (g *Domain) CreateDomain(req CreateRequest) (err error) {
	_, err = g.client.Post("domains", req, nil)
	return
}

// GetNameServers returns the configured nameservers for a domain
func (g *Domain) GetNameServers(domain string) (nameservers []string, err error) {
	_, err = g.client.Get("domains/"+domain+"/nameservers", nil, &nameservers)
	return
}

// UpdateNameServers sets the list of the nameservers for a domain
func (g *Domain) UpdateNameServers(domain string, ns []string) (err error) {
	_, err = g.client.Put("domains/"+domain+"/nameservers", Nameservers{ns}, nil)
	return
}

// GetContacts returns the contact objects for a domain
func (g *Domain) GetContacts(domain string) (contacts Contacts, err error) {
	_, err = g.client.Get("domains/"+domain+"/contacts", nil, &contacts)
	return
}

// SetContacts sets the contact objects for a domain
func (g *Domain) SetContacts(domain string, contacts Contacts) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/contacts", contacts, nil)
	return
}

// SetAutoRenew enables or disables auto renew on the given Domain
func (g *Domain) SetAutoRenew(domain string, autorenew bool) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/autorenew", AutoRenew{Enabled: &autorenew}, nil)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"encoding/json"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Domain API client
func New(config config.Config) *Domain {
	client := client.New(config.APIKey, config.PersonalAccessToken, config.APIURL, config.SharingID, config.Debug, config.DryRun, config.Timeout)
	client.SetEndpoint("domain/")
	return &Domain{client: *client}
}

// NewFromClient returns an instance of the Domain API client
func NewFromClient(g client.Gandi) *Domain {
	g.SetEndpoint("domain/")
	return &Domain{client: g}
}

// ListDomains requests the set of Domains
// It returns a slice of domains and any error encountered
func (g *Domain) ListDomains() (domains []ListResponse, err error) {
	_, elements, err := g.client.GetCollection("domains", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var key ListResponse
		err := json.Unmarshal(element, &key)
		if err != nil {
			return nil, err
		}
		domains = append(domains, key)
	}
	return domains, nil
}

// GetDomain requests a single Domain
// It returns a Details object and any error encountered
func (g *Domain) GetDomain(domain string) (domainResponse Details, err error) {
	_, err = g.client.Get("domains/"+domain, nil, &domainResponse)
	return
}

// CreateDomain creates a single Domain
func (g *Domain) CreateDomain(req CreateRequest) (err error) {
	_, err = g.client.Post("domains", req, nil)
	return
}

// GetNameServers returns the configured nameservers for a domain
func (g *Domain) GetNameServers(domain string) (nameservers []string, err error) {
	_, err = g.client.Get("domains/"+domain+"/nameservers", nil, &nameservers)
	return
}

// UpdateNameServers sets the list of the nameservers for a domain
func (g *Domain) UpdateNameServers(domain string, ns []string) (err error) {
	_, err = g.client.Put("domains/"+domain+"/nameservers", Nameservers{ns}, nil)
	return
}

// GetContacts returns the contact objects for a domain
func (g *Domain) GetContacts(domain string) (contacts Contacts, err error) {
	_, err = g.client.Get("domains/"+domain+"/contacts", nil, &contacts)
	return
}

// SetContacts sets the contact objects for a domain
func (g *Domain) SetContacts(domain string, contacts Contacts) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/contacts", contacts, nil)
	return
}

// SetAutoRenew enables or disables auto renew on the given Domain
func (g *Domain) SetAutoRenew(domain string, autorenew bool) (err error) {
	_, err = g.client.Patch("domains/"+domain+"/autorenew", AutoRenew{Enabled: &autorenew}, nil)
	return
}

func (g *Domain) ListDNSSECKeys(domain string) (keys []DNSSECKey, err error) {
	_, elements, err := g.client.GetCollection("domains/"+domain+"/dnskeys", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var key DNSSECKey
		err := json.Unmarshal(element, &key)
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}
	return keys, nil
}

func (g *Domain) CreateDNSSECKey(domain string, key DNSSECKeyCreateRequest) (err error) {
	_, err = g.client.Post("domains/"+domain+"/dnskeys", key, nil)
	return
}

func (g *Domain) DeleteDNSSECKey(domain string, keyid string) (err error) {
	_, err = g.client.Delete("domains/"+domain+"/dnskeys/"+keyid, nil, nil)
	return
}

func (g *Domain) CreateGlueRecord(domain string, gluerecord GlueRecordCreateRequest) (err error) {
	_, err = g.client.Post("domains/"+domain+"/hosts", gluerecord, nil)
	return
}

func (g *Domain) ListGlueRecords(domain string) (gluerecords []GlueRecord, err error) {
	_, elements, err := g.client.GetCollection("domains/"+domain+"/hosts", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var gluerecord GlueRecord
		err := json.Unmarshal(element, &gluerecord)
		if err != nil {
			return nil, err
		}
		gluerecords = append(gluerecords, gluerecord)
	}
	return gluerecords, nil
}

func (g *Domain) GetGlueRecord(domain string, name string) (gluerecord GlueRecord, err error) {
	_, err = g.client.Get("domains/"+domain+"/hosts/"+name, nil, &gluerecord)
	return
}

func (g *Domain) UpdateGlueRecord(domain string, name string, ips []string) (err error) {
	_, err = g.client.Put("domains/"+domain+"/hosts/"+name, GlueRecordUpdateRequest{ips}, nil)
	return
}

func (g *Domain) DeleteGlueRecord(domain string, name string) (err error) {
	_, err = g.client.Delete("domains/"+domain+"/hosts/"+name, nil, nil)
	return
}

func (g *Domain) CreateWebRedirection(domain string, webredir WebRedirectionCreateRequest) (err error) {
	_, err = g.client.Post("domains/"+domain+"/webredirs", webredir, nil)
	return
}

func (g *Domain) ListWebRedirections(domain string) (webredirs []WebRedirection, err error) {
	_, elements, err := g.client.GetCollection("domains/"+domain+"/webredirs", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var webredir WebRedirection
		err := json.Unmarshal(element, &webredir)
		if err != nil {
			return nil, err
		}
		webredirs = append(webredirs, webredir)
	}
	return webredirs, nil
}

func (g *Domain) GetWebRedirection(domain string, host string) (webredir WebRedirection, err error) {
	_, err = g.client.Get("domains/"+domain+"/webredirs/"+host, nil, &webredir)
	return
}

func (g *Domain) DeleteWebRedirection(domain string, host string) (err error) {
	_, err = g.client.Delete("domains/"+domain+"/webredirs/"+host, nil, nil)
	return
}

func (g *Domain) EnableLiveDNS(domain string) (err error) {
	_, err = g.client.Post("domains/"+domain+"/livedns", nil, nil)
	return
}

func (g *Domain) GetLiveDNS(domain string) (livedns LiveDNS, err error) {
	_, err = g.client.Get("domains/"+domain+"/livedns", nil, &livedns)
	return
}

func (g *Domain) GetTags(domain string) (tags []string, err error) {
	_, err = g.client.Get("domains/"+domain+"/tags", nil, &tags)
	return
}

func (g *Domain) SetTags(domain string, tags []string) (err error) {
	_, err = g.client.Put("domains/"+domain+"/tags", Tags{tags}, nil)
	return
}

func (g *Domain) DeleteTags(domain string) (err error) {
	_, err = g.client.Delete("domains/"+domain+"/tags", nil, nil)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return
}
