package livedns

import (
	"encoding/json"

	"github.com/go-gandi/go-gandi/types"
)

// ListSnapshots lists all snapshots for a domain
func (g *LiveDNS) ListSnapshots(fqdn string) (snapshots []Snapshot, err error) {
	_, elements, err := g.client.GetCollection("domains/"+fqdn+"/snapshots", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var snapshot Snapshot
		err := json.Unmarshal(element, &snapshot)
		if err != nil {
			return nil, err
		}
		snapshots = append(snapshots, snapshot)
	}
	return snapshots, nil
}

// CreateSnapshot creates a snapshot for a domain
func (g *LiveDNS) CreateSnapshot(fqdn string) (response types.StandardResponse, err error) {
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
