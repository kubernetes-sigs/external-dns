package domain

import (
	"fmt"
	"github.com/transip/gotransip/v6"
	"github.com/transip/gotransip/v6/repository"
	"github.com/transip/gotransip/v6/rest"
	"net/url"
)

// Repository can be used to get a list of your domains,
// order new ones and changing specific domain properties
type Repository repository.RestRepository

// GetAll returns all domains listed in your account
func (r *Repository) GetAll() ([]Domain, error) {
	var response domainsResponse
	err := r.Client.Get(rest.Request{Endpoint: "/domains"}, &response)

	return response.Domains, err
}

// GetAllByTags returns a list of all Domains that match the tags provided
func (r *Repository) GetAllByTags(tags []string) ([]Domain, error) {
	var response domainsResponse
	restRequest := rest.Request{Endpoint: "/domains", Parameters: url.Values{"tags": tags}}
	err := r.Client.Get(restRequest, &response)

	return response.Domains, err
}

// GetSelection returns a limited list of all domains in your account,
// specify how many and which page/chunk of domains you want to retrieve
func (r *Repository) GetSelection(page int, itemsPerPage int) ([]Domain, error) {
	var response domainsResponse
	params := url.Values{
		"pageSize": []string{fmt.Sprintf("%d", itemsPerPage)},
		"page":     []string{fmt.Sprintf("%d", page)},
	}

	restRequest := rest.Request{Endpoint: "/domains", Parameters: params}
	err := r.Client.Get(restRequest, &response)

	return response.Domains, err
}

// GetByDomainName returns a Domain struct for a specific domain name.
//
// Requires a domainName, for example: 'example.com'
func (r *Repository) GetByDomainName(domainName string) (Domain, error) {
	var response domainWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Domain, err
}

// Register allows you to registers a new domain.
// You can set the contacts, nameservers and DNS entries immediately, but it’s not mandatory for registration.
func (r *Repository) Register(domainRegister Register) error {
	restRequest := rest.Request{Endpoint: "/domains", Body: &domainRegister}

	return r.Client.Post(restRequest)
}

// Transfer allows you to transfer a domain to TransIP using its transfer key
// (or ‘EPP code’) by specifying it in the authCode parameter
func (r *Repository) Transfer(domainTransfer Transfer) error {
	restRequest := rest.Request{Endpoint: "/domains", Body: &domainTransfer}

	return r.Client.Post(restRequest)
}

// Update an existing domain.
// To apply or release a lock, change the IsTransferLocked property.
// To change tags, update the tags property.
func (r *Repository) Update(domain Domain) error {
	requestBody := domainWrapper{Domain: domain}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s", domain.Name), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// Cancel cancels the specified domain.
// Depending on the time you want to cancel the domain,
// specify gotransip.CancellationTimeEnd or gotransip.CancellationTimeImmediately for the endTime attribute.
func (r *Repository) Cancel(domainName string, endTime gotransip.CancellationTime) error {
	var requestBody gotransip.CancellationRequest
	requestBody.EndTime = endTime
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s", domainName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetBranding returns a Branding struct for the given domain.
// Branding can be altered using the method below
func (r *Repository) GetBranding(domainName string) (Branding, error) {
	var response domainBrandingWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/branding", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Branding, err
}

// UpdateBranding allows you to change the branding information on a domain
func (r *Repository) UpdateBranding(domainName string, branding Branding) error {
	requestBody := domainBrandingWrapper{Branding: branding}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/branding", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetContacts returns a list of contacts for the given domain name
func (r *Repository) GetContacts(domainName string) ([]WhoisContact, error) {
	var response contactsWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/contacts", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Contacts, err
}

// UpdateContacts allows you to replace the whois contacts currently on a domain
func (r *Repository) UpdateContacts(domainName string, contacts []WhoisContact) error {
	requestBody := contactsWrapper{Contacts: contacts}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/contacts", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetDNSEntries returns a list of all DNS entries for a domain by domainName
func (r *Repository) GetDNSEntries(domainName string) ([]DNSEntry, error) {
	var response dnsEntriesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.DNSEntries, err
}

// AddDNSEntry allows you to add a single dns entry to a domain
func (r *Repository) AddDNSEntry(domainName string, dnsEntry DNSEntry) error {
	requestBody := dnsEntryWrapper{DNSEntry: dnsEntry}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Post(restRequest)
}

// UpdateDNSEntry updates the content of a single DNS entry,
// the dns entry is identified by the 'Name', 'Expire' and 'Type' properties of the DNSEntry struct
func (r *Repository) UpdateDNSEntry(domainName string, dnsEntry DNSEntry) error {
	requestBody := dnsEntryWrapper{DNSEntry: dnsEntry}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// ReplaceDNSEntries will wipe the entire zone replacing it with the given dns entries
func (r *Repository) ReplaceDNSEntries(domainName string, dnsEntries []DNSEntry) error {
	requestBody := dnsEntriesWrapper{DNSEntries: dnsEntries}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// RemoveDNSEntry allows you to remove a single DNS entry from a domain
func (r *Repository) RemoveDNSEntry(domainName string, dnsEntry DNSEntry) error {
	requestBody := dnsEntryWrapper{DNSEntry: dnsEntry}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/dns", domainName), Body: &requestBody}

	return r.Client.Delete(restRequest)
}

// GetDNSSecEntries returns a list of all DNS Sec entries for a domain by domainName
func (r *Repository) GetDNSSecEntries(domainName string) ([]DNSSecEntry, error) {
	var response dnsSecEntriesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/dnssec", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.DNSSecEntries, err
}

// ReplaceDNSSecEntries allows you to replace all DNSSEC entries with the ones that are provided
func (r *Repository) ReplaceDNSSecEntries(domainName string, dnsSecEntries []DNSSecEntry) error {
	requestBody := dnsSecEntriesWrapper{DNSSecEntries: dnsSecEntries}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/dnssec", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetNameservers will list all nameservers currently set for a domain.
func (r *Repository) GetNameservers(domainName string) ([]Nameserver, error) {
	var response nameserversWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/nameservers", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Nameservers, err
}

// UpdateNameservers allows you to change the nameservers for a domain
func (r *Repository) UpdateNameservers(domainName string, nameservers []Nameserver) error {
	requestBody := nameserversWrapper{Nameservers: nameservers}
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/nameservers", domainName), Body: &requestBody}

	return r.Client.Put(restRequest)
}

// GetDomainAction allows you to get the current domain action running for the given domain.
// Domain actions are kept track of by TransIP. Domain actions include, for example, changing nameservers.
func (r *Repository) GetDomainAction(domainName string) (Action, error) {
	var response actionWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/actions", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Action, err
}

// RetryDomainAction allows you to retry a failed domain action.
// Domain actions can fail due to wrong information, this method allows you to retry an action.
func (r *Repository) RetryDomainAction(domainName string, authCode string, dnsEntries []DNSEntry, nameservers []Nameserver, contacts []WhoisContact) error {
	var requestBody retryActionWrapper
	requestBody.AuthCode = authCode
	requestBody.DNSEntries = dnsEntries
	requestBody.Nameservers = nameservers
	requestBody.Contacts = contacts
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/actions", domainName), Body: &requestBody}

	return r.Client.Patch(restRequest)
}

// CancelDomainAction allows you to cancel a domain action while it is still pending or being processed
func (r *Repository) CancelDomainAction(domainName string) error {
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/actions", domainName)}

	return r.Client.Delete(restRequest)
}

// GetSSLCertificates allows you to get a list of SSL certificates for a specific domain
func (r *Repository) GetSSLCertificates(domainName string) ([]SslCertificate, error) {
	var response certificatesWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/ssl", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Certificates, err
}

// GetSSLCertificateByID allows you to get a single SSL certificate by id.
func (r *Repository) GetSSLCertificateByID(domainName string, certificateID int64) (SslCertificate, error) {
	var response certificateWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/ssl/%d", domainName, certificateID)}
	err := r.Client.Get(restRequest, &response)

	return response.Certificate, err
}

// GetWHOIS will return the WHOIS information for a domain name as a string
func (r *Repository) GetWHOIS(domainName string) (string, error) {
	var response whoisWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domains/%s/whois", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Whois, err
}

// OrderWhitelabel allows you to order a whitelabel account.
// Note that you do not need to order a whitelabel account for every registered domain name.
func (r *Repository) OrderWhitelabel() error {
	restRequest := rest.Request{Endpoint: "/whitelabel"}

	return r.Client.Post(restRequest)
}

// GetAvailability method allows you to check the availability for a domain name
func (r *Repository) GetAvailability(domainName string) (Availability, error) {
	var response availabilityWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/domain-availability/%s", domainName)}
	err := r.Client.Get(restRequest, &response)

	return response.Availability, err
}

// GetAvailabilityForMultipleDomains method allows you to check the availability for a list of domain names
func (r *Repository) GetAvailabilityForMultipleDomains(domainNames []string) ([]Availability, error) {
	var response availabilityListWrapper
	var requestBody multipleAvailabilityRequest
	requestBody.DomainNames = domainNames

	restRequest := rest.Request{Endpoint: "/domain-availability", Body: requestBody}
	err := r.Client.Get(restRequest, &response)

	return response.AvailabilityList, err
}

// GetTLDs will return a list of all available TLDs currently offered by TransIP
func (r *Repository) GetTLDs() ([]Tld, error) {
	var response tldsWrapper
	restRequest := rest.Request{Endpoint: "/tlds"}
	err := r.Client.Get(restRequest, &response)

	return response.Tlds, err
}

// GetTLDByTLD returns information about a specific TLD.
// General details such as price, renewal price and minimum registration length are outlined.
func (r *Repository) GetTLDByTLD(tld string) (Tld, error) {
	var response tldWrapper
	restRequest := rest.Request{Endpoint: fmt.Sprintf("/tlds/%s", tld)}
	err := r.Client.Get(restRequest, &response)

	return response.Tld, err
}
