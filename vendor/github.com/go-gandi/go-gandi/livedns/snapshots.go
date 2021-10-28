package livedns

import (
	"time"

	"github.com/go-gandi/go-gandi/internal/client"
)

// Snapshot represents a point in time record of a domain
type Snapshot struct {
	Automatic    *bool          `json:"automatic,omitempty"`
	CreatedAt    time.Time      `json:"created_at,omitempty"`
	ID           string         `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	SnapshotHREF string         `json:"snapshot_href,omitempty"`
	ZoneData     []DomainRecord `json:"zone_data,omitempty"`
}

// ListSnapshots lists all snapshots for a domain
func (g *LiveDNS) ListSnapshots(fqdn string) (snapshots []Snapshot, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/snapshots", nil, &snapshots)
	return
}

// CreateSnapshot creates a snapshot for a domain
func (g *LiveDNS) CreateSnapshot(fqdn string) (response client.StandardResponse, err error) {
	_, err = g.client.Post("domains/"+fqdn+"/snapshots", nil, &response)
	return
}

// GetSnapshot returns a snapshot for a domain
func (g *LiveDNS) GetSnapshot(fqdn, snapUUID string) (snapshot Snapshot, err error) {
	_, err = g.client.Get("domains/"+fqdn+"/snapshots/"+snapUUID, nil, &snapshot)
	return
}

// DeleteSnapshot deletes a snapshot for a domain
func (g *LiveDNS) DeleteSnapshot(fqdn, snapUUID string) (err error) {
	_, err = g.client.Delete("domains/"+fqdn+"/snapshots/"+snapUUID, nil, nil)
	return
}
