package dynect

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// ConvenientClient A client with extra helper methods for common actions
type ConvenientClient struct {
	Client
}

// NewConvenientClient Creates a new ConvenientClient
func NewConvenientClient(customerName string) *ConvenientClient {
	return &ConvenientClient{
		Client{
			CustomerName: customerName,
			Transport:    &http.Transport{Proxy: http.ProxyFromEnvironment},
		}}
}

// CreateZone method to create a zone
func (c *ConvenientClient) CreateZone(zone, rname, serialStyle, ttl string) error {
	url := fmt.Sprintf("Zone/%s/", zone)
	data := &CreateZoneBlock{
		RName:       rname,
		SerialStyle: serialStyle,
		TTL:         ttl,
	}

	if err := c.Do("POST", url, data, nil); err != nil {
		return fmt.Errorf("Failed to create zone: %s", err)
	}

	return nil
}

// GetZone method to read a zone
func (c *ConvenientClient) GetZone(z *Zone) error {
	url := fmt.Sprintf("Zone/%s", z.Zone)
	data := &ZoneResponse{}

	if err := c.Do("GET", url, nil, data); err != nil {
		return fmt.Errorf("Failed to get zone: %s", err)
	}

	z.Serial = strconv.Itoa(data.Data.Serial)
	z.SerialStyle = data.Data.SerialStyle
	z.Zone = data.Data.Zone
	z.Type = data.Data.ZoneType

	return nil
}

// PublishZone Publish a specific zone and the changes for the current session
func (c *ConvenientClient) PublishZone(zone string) error {
	url := fmt.Sprintf("Zone/%s", zone)
	data := &PublishZoneBlock{
		Publish: true,
	}

	if err := c.Do("PUT", url, data, nil); err != nil {
		return fmt.Errorf("Failed to publish zone: %s", err)
	}

	return nil
}

// DeleteZoneNode method to delete everything in a zone
func (c *ConvenientClient) DeleteZoneNode(zone string) error {
	parentZone := strings.Join(strings.Split(zone, ".")[1:], ".")
	url := fmt.Sprintf("Node/%s/%s", parentZone, zone)

	if err := c.Do("DELETE", url, nil, nil); err != nil {
		return fmt.Errorf("Failed to delete zone node: %s", err)
	}

	return nil
}

// DeleteZone method to delete a zone
func (c *ConvenientClient) DeleteZone(zone string) error {
	url := fmt.Sprintf("Zone/%s/", zone)

	if err := c.Do("DELETE", url, nil, nil); err != nil {
		return fmt.Errorf("Failed to delete zone: %s", err)
	}

	return nil
}

// GetRecordID finds the dns record ID by fetching all records for a FQDN
func (c *ConvenientClient) GetRecordID(record *Record) error {
	finalID := ""
	url := fmt.Sprintf("AllRecord/%s/%s", record.Zone, record.FQDN)
	var records AllRecordsResponse
	err := c.Do("GET", url, nil, &records)
	if err != nil {
		return fmt.Errorf("Failed to find Dyn record id: %s", err)
	}
	for _, recordURL := range records.Data {
		id := strings.TrimPrefix(recordURL, fmt.Sprintf("/REST/%sRecord/%s/%s/", record.Type, record.Zone, record.FQDN))
		if !strings.Contains(id, "/") && id != "" {
			finalID = id
			log.Printf("[INFO] Found Dyn record ID: %s", id)
		}
	}
	if finalID == "" {
		return fmt.Errorf("Failed to find Dyn record id!")
	}

	record.ID = finalID
	return nil
}

// CreateRecord Method to create a DNS record
func (c *ConvenientClient) CreateRecord(record *Record) error {
	if record.FQDN == "" && record.Name == "" {
		record.FQDN = record.Zone
	} else if record.FQDN == "" {
		record.FQDN = fmt.Sprintf("%s.%s", record.Name, record.Zone)
	}
	rdata, err := buildRData(record)
	if err != nil {
		return fmt.Errorf("Failed to create Dyn RData: %s", err)
	}
	url := fmt.Sprintf("%sRecord/%s/%s", record.Type, record.Zone, record.FQDN)
	data := &RecordRequest{
		RData: rdata,
		TTL:   record.TTL,
	}
	return c.Do("POST", url, data, nil)
}

// UpdateRecord Method to update a DNS record
func (c *ConvenientClient) UpdateRecord(record *Record) error {
	if record.FQDN == "" {
		record.FQDN = fmt.Sprintf("%s.%s", record.Name, record.Zone)
	}
	rdata, err := buildRData(record)
	if err != nil {
		return fmt.Errorf("Failed to create Dyn RData: %s", err)
	}
	url := fmt.Sprintf("%sRecord/%s/%s/%s", record.Type, record.Zone, record.FQDN, record.ID)
	data := &RecordRequest{
		RData: rdata,
		TTL:   record.TTL,
	}
	return c.Do("PUT", url, data, nil)
}

// DeleteRecord Method to delete a DNS record
func (c *ConvenientClient) DeleteRecord(record *Record) error {
	if record.FQDN == "" {
		record.FQDN = fmt.Sprintf("%s.%s", record.Name, record.Zone)
	}
	// safety check that we have an ID, otherwise we could accidentally delete everything
	if record.ID == "" {
		return fmt.Errorf("No ID found! We can't continue!")
	}
	url := fmt.Sprintf("%sRecord/%s/%s/%s", record.Type, record.Zone, record.FQDN, record.ID)
	return c.Do("DELETE", url, nil, nil)
}

// GetRecord Method to get record details
func (c *ConvenientClient) GetRecord(record *Record) error {
	url := fmt.Sprintf("%sRecord/%s/%s/%s", record.Type, record.Zone, record.FQDN, record.ID)
	var rec RecordResponse
	err := c.Do("GET", url, nil, &rec)
	if err != nil {
		return err
	}

	record.Zone = rec.Data.Zone
	record.FQDN = rec.Data.FQDN
	record.Name = strings.TrimSuffix(rec.Data.FQDN, "."+rec.Data.Zone)
	record.Type = rec.Data.RecordType
	record.TTL = strconv.Itoa(rec.Data.TTL)

	switch rec.Data.RecordType {
	case "A", "AAAA":
		record.Value = rec.Data.RData.Address
	case "ALIAS":
		record.Value = rec.Data.RData.Alias
	case "CNAME":
		record.Value = rec.Data.RData.CName
	case "MX":
		record.Value = fmt.Sprintf("%d %s", rec.Data.RData.Preference, rec.Data.RData.Exchange)
	case "NS":
		record.Value = rec.Data.RData.NSDName
	case "SOA":
		record.Value = rec.Data.RData.RName
	case "TXT", "SPF":
		record.Value = rec.Data.RData.TxtData
	default:
		fmt.Println("unknown response", rec)
		return fmt.Errorf("Invalid Dyn record type: %s", rec.Data.RecordType)
	}

	return nil
}

func buildRData(r *Record) (DataBlock, error) {
	var rdata DataBlock

	switch r.Type {
	case "A", "AAAA":
		rdata = DataBlock{
			Address: r.Value,
		}
	case "ALIAS":
		rdata = DataBlock{
			Alias: r.Value,
		}
	case "CNAME":
		rdata = DataBlock{
			CName: r.Value,
		}
	case "MX":
		rdata = DataBlock{}
		fmt.Sscanf(r.Value, "%d %s", &rdata.Preference, &rdata.Exchange)
	case "NS":
		rdata = DataBlock{
			NSDName: r.Value,
		}
	case "SOA":
		rdata = DataBlock{
			RName: r.Value,
		}
	case "TXT", "SPF":
		rdata = DataBlock{
			TxtData: r.Value,
		}
	default:
		return rdata, fmt.Errorf("Invalid Dyn record type: %s", r.Type)
	}

	return rdata, nil
}
