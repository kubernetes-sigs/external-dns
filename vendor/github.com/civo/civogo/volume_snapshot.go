package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// VolumeSnapshot is the point-in-time copy of a Volume
type VolumeSnapshot struct {
	Name                string `json:"name"`
	SnapshotID          string `json:"snapshot_id"`
	SnapshotDescription string `json:"snapshot_description"`
	VolumeID            string `json:"volume_id"`
	InstanceID          string `json:"instance_id,omitempty"`
	SourceVolumeName    string `json:"source_volume_name"`
	RestoreSize         int    `json:"restore_size"`
	State               string `json:"state"`
	CreationTime        string `json:"creation_time,omitempty"`
}

// VolumeSnapshotConfig is the configuration for creating a new VolumeSnapshot
type VolumeSnapshotConfig struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Region      string `json:"region"`
}

// ListVolumeSnapshots returns all snapshots owned by the calling API account
func (c *Client) ListVolumeSnapshots() ([]VolumeSnapshot, error) {
	resp, err := c.SendGetRequest("/v2/snapshots?resource_type=volume")
	if err != nil {
		return nil, decodeError(err)
	}

	var volumeSnapshots = make([]VolumeSnapshot, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&volumeSnapshots); err != nil {
		return nil, err
	}

	return volumeSnapshots, nil
}

// GetVolumeSnapshot retrieves a volume snapshot based on the provided snapshot ID.
func (c *Client) GetVolumeSnapshot(id string) (*VolumeSnapshot, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/snapshots/%s?resource_type=volume", id))
	if err != nil {
		return nil, decodeError(err)
	}
	var volumeSnapshot = VolumeSnapshot{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&volumeSnapshot); err != nil {
		return nil, err
	}
	return &volumeSnapshot, nil
}

// DeleteVolumeSnapshot deletes a volume snapshot
func (c *Client) DeleteVolumeSnapshot(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/snapshots/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
