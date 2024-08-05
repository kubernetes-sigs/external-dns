package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

// KubernetesClusterPoolUpdateConfig is used to create a new cluster pool
type KubernetesClusterPoolUpdateConfig struct {
	ID               string            `json:"id,omitempty"`
	Count            int               `json:"count,omitempty"`
	Size             string            `json:"size,omitempty"`
	Labels           map[string]string `json:"labels,omitempty"`
	Taints           []corev1.Taint    `json:"taints"`
	PublicIPNodePool bool              `json:"public_ip_node_pool,omitempty"`
	Region           string            `json:"region,omitempty"`
}

// ListKubernetesClusterPools returns all the pools for a kubernetes cluster
func (c *Client) ListKubernetesClusterPools(cid string) ([]KubernetesPool, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/pools", cid))
	if err != nil {
		return nil, decodeError(err)
	}

	pools := make([]KubernetesPool, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&pools); err != nil {
		return nil, decodeError(err)
	}

	return pools, nil
}

// CreateKubernetesClusterPool update a single kubernetes cluster by its full ID
func (c *Client) CreateKubernetesClusterPool(id string, i *KubernetesClusterPoolUpdateConfig) (*SimpleResponse, error) {
	i.Region = c.Region
	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/pools", id), i)
	if err != nil {
		return nil, decodeError(err)
	}
	return c.DecodeSimpleResponse(resp)
}

// GetKubernetesClusterPool returns a pool for a kubernetes cluster
func (c *Client) GetKubernetesClusterPool(cid, pid string) (*KubernetesPool, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/pools/%s", cid, pid))
	if err != nil {
		return nil, decodeError(err)
	}

	pool := &KubernetesPool{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&pool); err != nil {
		return nil, decodeError(err)
	}

	return pool, nil
}

// FindKubernetesClusterPool finds a pool by either part of the ID
func (c *Client) FindKubernetesClusterPool(cid, search string) (*KubernetesPool, error) {
	pools, err := c.ListKubernetesClusterPools(cid)
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := KubernetesPool{}

	for _, value := range pools {
		if value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.ID, search) {
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

// DeleteKubernetesClusterPoolInstance deletes a instance from pool
func (c *Client) DeleteKubernetesClusterPoolInstance(cid, pid, id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/pools/%s/instances/%s", cid, pid, id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UpdateKubernetesClusterPool updates a pool for a kubernetes cluster
func (c *Client) UpdateKubernetesClusterPool(cid, pid string, config *KubernetesClusterPoolUpdateConfig) (*KubernetesPool, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/pools/%s", cid, pid), config)
	if err != nil {
		return nil, decodeError(err)
	}

	pool := &KubernetesPool{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&pool); err != nil {
		return nil, decodeError(err)
	}

	return pool, nil
}

// DeleteKubernetesClusterPool delete a pool inside the cluster
func (c *Client) DeleteKubernetesClusterPool(id, poolID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/kubernetes/clusters/%s/pools/%s", id, poolID))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
