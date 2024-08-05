package v2

import (
	"context"
	"time"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// DNSDomain represents a DNS domain.
type DNSDomain struct {
	CreatedAt   *time.Time
	ID          *string `req-for:"delete"`
	UnicodeName *string `req-for:"create"`
}

// DNSDomainRecord represents a DNS record.
type DNSDomainRecord struct {
	Content   *string `req-for:"create"`
	CreatedAt *time.Time
	ID        *string `req-for:"delete,update"`
	Name      *string `req-for:"create"`
	Priority  *int64
	TTL       *int64
	Type      *string `req-for:"create"`
	UpdatedAt *time.Time
}

func dnsDomainFromAPI(d *oapi.DnsDomain) *DNSDomain {
	return &DNSDomain{
		CreatedAt:   d.CreatedAt,
		ID:          d.Id,
		UnicodeName: d.UnicodeName,
	}
}

func dnsDomainRecordFromAPI(d *oapi.DnsDomainRecord) *DNSDomainRecord {
	var t *string
	if d.Type != nil {
		x := string(*d.Type)
		t = &x
	}
	return &DNSDomainRecord{
		Content:   d.Content,
		CreatedAt: d.CreatedAt,
		ID:        d.Id,
		Name:      d.Name,
		Priority:  d.Priority,
		TTL:       d.Ttl,
		Type:      t,
		UpdatedAt: d.UpdatedAt,
	}
}

// ListDNSDomains returns the list of DNS domains.
func (c *Client) ListDNSDomains(ctx context.Context, zone string) ([]DNSDomain, error) {
	var list []DNSDomain

	resp, err := c.ListDnsDomainsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.DnsDomains != nil {
		for _, domain := range *resp.JSON200.DnsDomains {
			list = append(list, *dnsDomainFromAPI(&domain))
		}
	}

	return list, nil
}

// GetDNSDomain returns DNS domain details.
func (c *Client) GetDNSDomain(ctx context.Context, zone, id string) (*DNSDomain, error) {
	resp, err := c.GetDnsDomainWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return dnsDomainFromAPI(resp.JSON200), nil
}

// DeleteDNSDomain deletes a DNS domain.
func (c *Client) DeleteDNSDomain(ctx context.Context, zone string, domain *DNSDomain) error {
	if err := validateOperationParams(domain, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteDnsDomainWithResponse(apiv2.WithZone(ctx, zone), *domain.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// CreateDNSDomain adds a new DNS domain.
func (c *Client) CreateDNSDomain(
	ctx context.Context,
	zone string,
	domain *DNSDomain,
) (*DNSDomain, error) {
	if err := validateOperationParams(domain, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateDnsDomainWithResponse(apiv2.WithZone(ctx, zone), oapi.CreateDnsDomainJSONRequestBody{
		UnicodeName: domain.UnicodeName,
	})
	if err != nil {
		return nil, err
	}

	r, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetDNSDomain(ctx, zone, *r.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// GetDNSDomainZoneFile returns zone file of a DNS domain.
func (c *Client) GetDNSDomainZoneFile(ctx context.Context, zone, id string) ([]byte, error) {
	resp, err := c.GetDnsDomainZoneFileWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

// ListDNSDomainRecords returns the list of records for DNS domain.
func (c *Client) ListDNSDomainRecords(ctx context.Context, zone, id string) ([]DNSDomainRecord, error) {
	var list []DNSDomainRecord

	resp, err := c.ListDnsDomainRecordsWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200.DnsDomainRecords != nil {
		for _, record := range *resp.JSON200.DnsDomainRecords {
			list = append(list, *dnsDomainRecordFromAPI(&record))
		}
	}

	return list, nil
}

// GetDNSDomainRecord returns a single DNS domain record.
func (c *Client) GetDNSDomainRecord(ctx context.Context, zone, domainID, recordID string) (*DNSDomainRecord, error) {
	resp, err := c.GetDnsDomainRecordWithResponse(apiv2.WithZone(ctx, zone), domainID, recordID)
	if err != nil {
		return nil, err
	}

	return dnsDomainRecordFromAPI(resp.JSON200), nil
}

// CreateDNSDomainRecord adds a new DNS record for domain.
func (c *Client) CreateDNSDomainRecord(
	ctx context.Context,
	zone string,
	domainID string,
	record *DNSDomainRecord,
) (*DNSDomainRecord, error) {
	if err := validateOperationParams(record, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateDnsDomainRecordWithResponse(apiv2.WithZone(ctx, zone), domainID, oapi.CreateDnsDomainRecordJSONRequestBody{
		Content:  *record.Content,
		Name:     *record.Name,
		Priority: record.Priority,
		Ttl:      record.TTL,
		Type:     oapi.CreateDnsDomainRecordJSONBodyType(*record.Type),
	})
	if err != nil {
		return nil, err
	}

	r, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetDNSDomainRecord(ctx, zone, domainID, *r.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// DeleteDNSDomainRecord deletes a DNS domain record.
func (c *Client) DeleteDNSDomainRecord(ctx context.Context, zone, domainID string, record *DNSDomainRecord) error {
	if err := validateOperationParams(record, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteDnsDomainRecordWithResponse(apiv2.WithZone(ctx, zone), domainID, *record.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// UpdateDNSDomainRecord updates existing DNS domain record.
func (c *Client) UpdateDNSDomainRecord(ctx context.Context, zone, domainID string, record *DNSDomainRecord) error {
	if err := validateOperationParams(record, "update"); err != nil {
		return err
	}

	resp, err := c.UpdateDnsDomainRecordWithResponse(apiv2.WithZone(ctx, zone), domainID, *record.ID, oapi.UpdateDnsDomainRecordJSONRequestBody{
		Content:  record.Content,
		Name:     record.Name,
		Priority: record.Priority,
		Ttl:      record.TTL,
	})
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
