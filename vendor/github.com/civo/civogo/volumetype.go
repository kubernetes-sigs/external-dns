package civogo

import (
	"encoding/json"
)

// VolumeType represent the storage class related to a volume
// https://www.civo.com/api/volumes
type VolumeType struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Enabled     bool     `json:"enabled"`
	Labels      []string `json:"labels"`
}

// ListVolumeTypes returns a page of Instances owned by the calling API account
func (c *Client) ListVolumeTypes() ([]VolumeType, error) {
	resp, err := c.SendGetRequest("/v2/volumetypes")
	if err != nil {
		return nil, err
	}

	volumeTypes := make([]VolumeType, 0)
	if err := json.Unmarshal(resp, &volumeTypes); err != nil {
		return nil, err
	}

	return volumeTypes, nil
}
