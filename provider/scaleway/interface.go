package scaleway

import (
	domain "github.com/scaleway/scaleway-sdk-go/api/domain/v2alpha2"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// DomainAPI is an interface matching the domain.API struct
type DomainAPI interface {
	ListDNSZones(req *domain.ListDNSZonesRequest, opts ...scw.RequestOption) (*domain.ListDNSZonesResponse, error)
	ListDNSZoneRecords(req *domain.ListDNSZoneRecordsRequest, opts ...scw.RequestOption) (*domain.ListDNSZoneRecordsResponse, error)
	UpdateDNSZoneRecords(req *domain.UpdateDNSZoneRecordsRequest, opts ...scw.RequestOption) (*domain.UpdateDNSZoneRecordsResponse, error)
}
