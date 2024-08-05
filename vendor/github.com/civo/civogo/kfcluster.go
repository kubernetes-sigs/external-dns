package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// KfCluster represents a cluster with Kubeflow installed.
type KfCluster struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	NetworkID     string    `json:"network_id"`
	FirewallID    string    `json:"firewall_id,omitempty"`
	Size          string    `json:"size,omitempty"`
	KubeflowReady string    `json:"kubeflow_ready,omitempty"`
	DashboardURL  string    `json:"dashboard_url,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

// CreateKfClusterReq is the request for creating a KfCluster.
type CreateKfClusterReq struct {
	Name       string `json:"name" validate:"required"`
	NetworkID  string `json:"network_id" validate:"required"`
	FirewallID string `json:"firewall_id,omitempty"`
	Size       string `json:"size,omitempty"`
	Region     string `json:"region,omitempty"`
}

// UpdateKfClusterReq is the request for updating a KfCluster.
type UpdateKfClusterReq struct {
	Name   string `json:"name,omitempty"`
	Region string `json:"region,omitempty"`
	// Size string `json:"size"`
}

// PaginatedKfClusters returns a paginated list of KfCluster object
type PaginatedKfClusters struct {
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
	Pages   int         `json:"pages"`
	Items   []KfCluster `json:"items"`
}

// ListKfClusters returns all applications in that specific region
func (c *Client) ListKfClusters() (*PaginatedKfClusters, error) {
	resp, err := c.SendGetRequest("/v2/kfclusters")
	if err != nil {
		return nil, decodeError(err)
	}

	kfc := &PaginatedKfClusters{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&kfc); err != nil {
		return nil, decodeError(err)
	}

	return kfc, nil
}

// GetKfCluster returns a kubeflow cluster by ID
func (c *Client) GetKfCluster(id string) (*KfCluster, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/kfclusters/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	kfc := &KfCluster{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&kfc); err != nil {
		return nil, decodeError(err)
	}

	return kfc, nil
}

// FindKfCluster finds a kubeflow cluster by either part of the ID or part of the name
func (c *Client) FindKfCluster(search string) (*KfCluster, error) {
	kfClusters, err := c.ListKfClusters()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := KfCluster{}

	for _, value := range kfClusters.Items {
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

// CreateKfCluster creates a new kubeflow cluster
func (c *Client) CreateKfCluster(req CreateKfClusterReq) (*KfCluster, error) {
	req.Region = c.Region
	body, err := c.SendPostRequest("/v2/kfclusters", req)
	if err != nil {
		return nil, decodeError(err)
	}

	var kfc KfCluster
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&kfc); err != nil {
		return nil, err
	}

	return &kfc, nil
}

// UpdateKfCluster updates a kubeflow cluster
func (c *Client) UpdateKfCluster(id string, kfc *UpdateKfClusterReq) (*KfCluster, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/kfclusters/%s", id), kfc)
	if err != nil {
		return nil, decodeError(err)
	}

	updatedKfCluster := &KfCluster{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(updatedKfCluster); err != nil {
		return nil, err
	}

	return updatedKfCluster, nil
}

// DeleteKfCluster deletes an application
func (c *Client) DeleteKfCluster(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/kfclusters/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
