package scaleway

import (
	domain "github.com/scaleway/scaleway-sdk-go/api/domain/v2alpha2"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// DomainAPI is an interface matching the domain.API struct
type DomainAPI interface {
	ListTasks(req *domain.ListTasksRequest, opts ...scw.RequestOption) (*domain.ListTasksResponse, error)
	BuyDomain(req *domain.BuyDomainRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	RenewDomain(req *domain.RenewDomainRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	TransferDomain(req *domain.TransferDomainRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	TradeDomain(req *domain.TradeDomainRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	RegisterExternalDomain(req *domain.RegisterExternalDomainRequest, opts ...scw.RequestOption) (*domain.RegisterExternalDomainResponse, error)
	DeleteExternalDomain(req *domain.DeleteExternalDomainRequest, opts ...scw.RequestOption) (*domain.DeleteExternalDomainResponse, error)
	ListContacts(req *domain.ListContactsRequest, opts ...scw.RequestOption) (*domain.ListContactsResponse, error)
	GetContact(req *domain.GetContactRequest, opts ...scw.RequestOption) (*domain.Contact, error)
	UpdateContact(req *domain.UpdateContactRequest, opts ...scw.RequestOption) (*domain.Contact, error)
	ListDomains(req *domain.ListDomainsRequest, opts ...scw.RequestOption) (*domain.ListDomainsResponse, error)
	GetDomain(req *domain.GetDomainRequest, opts ...scw.RequestOption) (*domain.GetDomainResponse, error)
	UpdateDomain(req *domain.UpdateDomainRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	LockDomainTransfer(req *domain.LockDomainTransferRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	UnlockDomainTransfer(req *domain.UnlockDomainTransferRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	EnableDomainAutoRenew(req *domain.EnableDomainAutoRenewRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	DisableDomainAutoRenew(req *domain.DisableDomainAutoRenewRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	GetDomainAuthCode(req *domain.GetDomainAuthCodeRequest, opts ...scw.RequestOption) (*domain.GetDomainAuthCodeResponse, error)
	EnableDomainDNSSEC(req *domain.EnableDomainDNSSECRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	DisableDomainDNSSEC(req *domain.DisableDomainDNSSECRequest, opts ...scw.RequestOption) (*domain.Domain, error)
	CreateDNSZone(req *domain.CreateDNSZoneRequest, opts ...scw.RequestOption) (*domain.DNSZone, error)
	UpdateDNSZone(req *domain.UpdateDNSZoneRequest, opts ...scw.RequestOption) (*domain.DNSZone, error)
	CopyDNSZone(req *domain.CopyDNSZoneRequest, opts ...scw.RequestOption) (*domain.DNSZone, error)
	DeleteDNSZone(req *domain.DeleteDNSZoneRequest, opts ...scw.RequestOption) (*domain.DeleteDNSZoneResponse, error)
	ListDNSZoneNameservers(req *domain.ListDNSZoneNameserversRequest, opts ...scw.RequestOption) (*domain.ListDNSZoneNameserversResponse, error)
	UpdateDNSZoneNameservers(req *domain.UpdateDNSZoneNameserversRequest, opts ...scw.RequestOption) (*domain.UpdateDNSZoneNameserversResponse, error)
	ClearDNSZoneRecords(req *domain.ClearDNSZoneRecordsRequest, opts ...scw.RequestOption) (*domain.ClearDNSZoneRecordsResponse, error)
	ExportRawDNSZone(req *domain.ExportRawDNSZoneRequest, opts ...scw.RequestOption) (*scw.File, error)
	ImportRawDNSZone(req *domain.ImportRawDNSZoneRequest, opts ...scw.RequestOption) (*domain.ImportRawDNSZoneResponse, error)
	ImportProviderDNSZone(req *domain.ImportProviderDNSZoneRequest, opts ...scw.RequestOption) (*domain.ImportProviderDNSZoneResponse, error)
	RefreshDNSZone(req *domain.RefreshDNSZoneRequest, opts ...scw.RequestOption) (*domain.RefreshDNSZoneResponse, error)
	ListDNSZoneVersions(req *domain.ListDNSZoneVersionsRequest, opts ...scw.RequestOption) (*domain.ListDNSZoneVersionsResponse, error)
	ListDNSZoneVersionRecords(req *domain.ListDNSZoneVersionRecordsRequest, opts ...scw.RequestOption) (*domain.ListDNSZoneVersionRecordsResponse, error)
	GetDNSZoneVersionDiff(req *domain.GetDNSZoneVersionDiffRequest, opts ...scw.RequestOption) (*domain.GetDNSZoneVersionDiffResponse, error)
	RestoreDNSZoneVersion(req *domain.RestoreDNSZoneVersionRequest, opts ...scw.RequestOption) (*domain.RestoreDNSZoneVersionResponse, error)
	CreateSSLCertificate(req *domain.CreateSSLCertificateRequest, opts ...scw.RequestOption) (*domain.ZoneSSL, error)
	ListSSLCertificates(req *domain.ListSSLCertificatesRequest, opts ...scw.RequestOption) (*domain.ListSSLCertificatesResponse, error)
	DeleteSSLCertificate(req *domain.DeleteSSLCertificateRequest, opts ...scw.RequestOption) (*domain.DeleteSSLCertificateResponse, error)
	GetDNSZoneTsigKey(req *domain.GetDNSZoneTsigKeyRequest, opts ...scw.RequestOption) (*domain.GetDNSZoneTsigKeyResponse, error)
	DeleteDNSZoneTsigKey(req *domain.DeleteDNSZoneTsigKeyRequest, opts ...scw.RequestOption) error
	ListDNSZones(req *domain.ListDNSZonesRequest, opts ...scw.RequestOption) (*domain.ListDNSZonesResponse, error)
	ListDNSZoneRecords(req *domain.ListDNSZoneRecordsRequest, opts ...scw.RequestOption) (*domain.ListDNSZoneRecordsResponse, error)
	UpdateDNSZoneRecords(req *domain.UpdateDNSZoneRecordsRequest, opts ...scw.RequestOption) (*domain.UpdateDNSZoneRecordsResponse, error)
}
