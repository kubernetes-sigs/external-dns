package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Snapshot is a backup of an instance
type Snapshot struct {
	ID            string    `json:"id"`
	InstanceID    string    `json:"instance_id"`
	Hostname      string    `json:"hostname"`
	Template      string    `json:"template_id"`
	Region        string    `json:"region"`
	Name          string    `json:"name"`
	Safe          int       `json:"safe"`
	SizeGigabytes int       `json:"size_gb"`
	State         string    `json:"state"`
	Cron          string    `json:"cron_timing,omitempty"`
	RequestedAt   time.Time `json:"requested_at,omitempty"`
	CompletedAt   time.Time `json:"completed_at,omitempty"`
}

// SnapshotConfig represents the options required for creating a new snapshot
type SnapshotConfig struct {
	InstanceID string `json:"instance_id"`
	Safe       bool   `json:"safe"`
	Cron       string `json:"cron_timing"`
	Region     string `json:"region"`
}

// CreateSnapshot create a new or update an old snapshot
func (c *Client) CreateSnapshot(name string, r *SnapshotConfig) (*Snapshot, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/snapshots/%s", name), r)
	if err != nil {
		return nil, decodeError(err)
	}

	var n = &Snapshot{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(n); err != nil {
		return nil, err
	}

	return n, nil
}

// ListSnapshots returns a list of all snapshots within the current account
func (c *Client) ListSnapshots() ([]Snapshot, error) {
	resp, err := c.SendGetRequest("/v2/snapshots")
	if err != nil {
		return nil, decodeError(err)
	}

	snapshots := make([]Snapshot, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&snapshots); err != nil {
		return nil, err
	}

	return snapshots, nil
}

// FindSnapshot finds a snapshot by either part of the ID or part of the name
func (c *Client) FindSnapshot(search string) (*Snapshot, error) {
	snapshots, err := c.ListSnapshots()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Snapshot{}

	for _, value := range snapshots {
		if value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// DeleteSnapshot deletes a snapshot
func (c *Client) DeleteSnapshot(name string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/snapshots/%s", name))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
