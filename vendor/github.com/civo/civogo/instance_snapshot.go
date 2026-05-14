package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// InstanceSnapshot represents a snapshot of an instance
type InstanceSnapshot struct {
	ID              string                 `json:"id"`
	Name            string                 `json:"name"`
	Description     string                 `json:"description,omitempty"`
	IncludedVolumes []string               `json:"included_volumes,omitempty"`
	Status          InstanceSnapshotStatus `json:"status,omitempty"`
	CreatedAt       time.Time              `json:"created_at,omitempty"`
}

// InstanceSnapshotStatus represents the status of an instance snapshot
type InstanceSnapshotStatus struct {
	State   string                         `json:"state"`
	Volumes []InstanceSnapshotVolumeStatus `json:"volumes,omitempty"`
}

// InstanceSnapshotVolumeStatus represents the status of a volume in an instance snapshot
type InstanceSnapshotVolumeStatus struct {
	ID    string `json:"id"`
	State string `json:"state"`
}

// CreateInstanceSnapshotParams represents the parameters for creating a new instance snapshot
type CreateInstanceSnapshotParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// UpdateInstanceSnapshotParams represents the parameters for updating an instance snapshot
type UpdateInstanceSnapshotParams struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// RestoreInstanceSnapshotParams represents the parameters for restoring an instance snapshot
type RestoreInstanceSnapshotParams struct {
	Description       string `json:"description,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	PrivateIPv4       string `json:"private_ipv4,omitempty"`
	OverwriteExisting bool   `json:"overwrite_existing,omitempty"`
}

// CreateInstanceSnapshot creates a new snapshot of an instance
func (c *Client) CreateInstanceSnapshot(instanceID string, params *CreateInstanceSnapshotParams) (*InstanceSnapshot, error) {
	url := fmt.Sprintf("/v2/instances/%s/snapshots", instanceID)
	resp, err := c.SendPostRequest(url, params)
	if err != nil {
		return nil, decodeError(err)
	}

	snapshot := &InstanceSnapshot{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(snapshot); err != nil {
		return nil, err
	}

	return snapshot, nil
}

// GetInstanceSnapshot gets a snapshot of an instance by ID or name
func (c *Client) GetInstanceSnapshot(instanceID, snapshotID string) (*InstanceSnapshot, error) {
	url := fmt.Sprintf("/v2/instances/%s/snapshots/%s", instanceID, snapshotID)
	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	snapshot := &InstanceSnapshot{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(snapshot); err != nil {
		return nil, err
	}

	return snapshot, nil
}

// ListInstanceSnapshots lists all snapshots for an instance
func (c *Client) ListInstanceSnapshots(instanceID string) ([]InstanceSnapshot, error) {
	url := fmt.Sprintf("/v2/instances/%s/snapshots", instanceID)
	resp, err := c.SendGetRequest(url)
	if err != nil {
		return nil, decodeError(err)
	}

	snapshots := make([]InstanceSnapshot, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&snapshots); err != nil {
		return nil, err
	}

	return snapshots, nil
}

// UpdateInstanceSnapshot updates a snapshot of an instance
func (c *Client) UpdateInstanceSnapshot(instanceID, snapshotID string, params *UpdateInstanceSnapshotParams) (*InstanceSnapshot, error) {
	url := fmt.Sprintf("/v2/instances/%s/snapshots/%s", instanceID, snapshotID)
	resp, err := c.SendPutRequest(url, params)
	if err != nil {
		return nil, decodeError(err)
	}

	snapshot := &InstanceSnapshot{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(snapshot); err != nil {
		return nil, err
	}

	return snapshot, nil
}

// DeleteInstanceSnapshot deletes a snapshot of an instance
func (c *Client) DeleteInstanceSnapshot(instanceID, snapshotID string) error {
	url := fmt.Sprintf("/v2/instances/%s/snapshots/%s", instanceID, snapshotID)
	_, err := c.SendDeleteRequest(url)
	if err != nil {
		return decodeError(err)
	}

	return nil
}

// RestoreInstanceSnapshot restores a snapshot of an instance
func (c *Client) RestoreInstanceSnapshot(instanceID, snapshotID string, params *RestoreInstanceSnapshotParams) (*InstanceRestoreInfo, error) {
	url := fmt.Sprintf("/v2/instances/%s/snapshots/%s/restore", instanceID, snapshotID)
	body, err := c.SendPostRequest(url, params)
	if err != nil {
		return nil, decodeError(err)
	}
	var instanceRestoreInfo InstanceRestoreInfo
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&instanceRestoreInfo); err != nil {
		return nil, decodeError(err)
	}

	return &instanceRestoreInfo, nil
}
