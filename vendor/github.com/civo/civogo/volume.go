package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Volume is a block of attachable storage for our IAAS products
// https://www.civo.com/api/volumes
type Volume struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	InstanceID    string    `json:"instance_id"`
	ClusterID     string    `json:"cluster_id"`
	NetworkID     string    `json:"network_id"`
	MountPoint    string    `json:"mountpoint"`
	Status        string    `json:"status"`
	SizeGigabytes int       `json:"size_gb"`
	Bootable      bool      `json:"bootable"`
	CreatedAt     time.Time `json:"created_at"`
}

// VolumeResult is the response from one of our simple API calls
type VolumeResult struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Result string `json:"result"`
}

// VolumeConfig are the settings required to create a new Volume
type VolumeConfig struct {
	Name          string `json:"name"`
	Namespace     string `json:"namespace"`
	ClusterID     string `json:"cluster_id"`
	NetworkID     string `json:"network_id"`
	Region        string `json:"region"`
	SizeGigabytes int    `json:"size_gb"`
	Bootable      bool   `json:"bootable"`
}

// ListVolumes returns all volumes owned by the calling API account
// https://www.civo.com/api/volumes#list-volumes
func (c *Client) ListVolumes() ([]Volume, error) {
	resp, err := c.SendGetRequest("/v2/volumes")
	if err != nil {
		return nil, decodeError(err)
	}

	var volumes = make([]Volume, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&volumes); err != nil {
		return nil, err
	}

	return volumes, nil
}

// ListVolumesForCluster returns all volumes for a cluster
func (c *Client) ListVolumesForCluster(clusterID string) ([]Volume, error) {
	cluster, err := c.FindKubernetesCluster(clusterID)
	if err != nil {
		return nil, err
	}

	volumes, err := c.ListVolumes()
	if err != nil {
		return nil, decodeError(err)
	}

	var vols []Volume
	for _, vol := range volumes {
		if vol.ClusterID != "" {
			if cluster.ID == vol.ClusterID {
				vols = append(vols, vol)
			}
		}
	}
	return vols, nil
}

// ListDanglingVolumes returns all dangling volumes (Volumes which have a cluster ID set but that cluster doesn't exist anymore)
func (c *Client) ListDanglingVolumes() ([]Volume, error) {
	clusters, err := c.ListKubernetesClusters()
	if err != nil {
		return nil, decodeError(err)
	}

	var clusterIDs []string
	for _, cluster := range clusters.Items {
		clusterIDs = append(clusterIDs, cluster.ID)
	}

	volumes, err := c.ListVolumes()
	if err != nil {
		return nil, decodeError(err)
	}

	var danglingVolumes = make([]Volume, 0)
	for _, volume := range volumes {
		if volume.ClusterID != "" {
			if !findString(clusterIDs, volume.ClusterID) {
				danglingVolumes = append(danglingVolumes, volume)
			}
		}
	}
	return danglingVolumes, nil
}

func findString(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// GetVolume finds a volume by the full ID
func (c *Client) GetVolume(id string) (*Volume, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/volumes/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	var volume = Volume{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&volume); err != nil {
		return nil, err
	}

	return &volume, nil
}

// FindVolume finds a volume by either part of the ID or part of the name
func (c *Client) FindVolume(search string) (*Volume, error) {
	volumes, err := c.ListVolumes()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Volume{}

	for _, value := range volumes {
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

// NewVolume creates a new volume
// https://www.civo.com/api/volumes#create-a-new-volume
func (c *Client) NewVolume(v *VolumeConfig) (*VolumeResult, error) {
	body, err := c.SendPostRequest("/v2/volumes", v)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &VolumeResult{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// ResizeVolume resizes a volume
// https://www.civo.com/api/volumes#resizing-a-volume
func (c *Client) ResizeVolume(id string, size int) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/volumes/%s/resize", id), map[string]interface{}{
		"size_gb": size,
		"region":  c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// AttachVolume attaches a volume to an instance
// https://www.civo.com/api/volumes#attach-a-volume-to-an-instance
func (c *Client) AttachVolume(id string, instance string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/volumes/%s/attach", id), map[string]string{
		"instance_id": instance,
		"region":      c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// DetachVolume attach volume from any instances
// https://www.civo.com/api/volumes#attach-a-volume-to-an-instance
func (c *Client) DetachVolume(id string) (*SimpleResponse, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/volumes/%s/detach", id), map[string]string{
		"region": c.Region,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	response, err := c.DecodeSimpleResponse(resp)
	return response, err
}

// DeleteVolume deletes a volumes
// https://www.civo.com/api/volumes#deleting-a-volume
func (c *Client) DeleteVolume(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/volumes/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
