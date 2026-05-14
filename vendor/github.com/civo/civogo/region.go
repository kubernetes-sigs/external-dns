package civogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Region represents a geographical/DC region for Civo resources
type Region struct {
	Code          string  `json:"code"`
	Name          string  `json:"name"`
	Type          string  `json:"type"`
	OutOfCapacity bool    `json:"out_of_capacity"`
	Country       string  `json:"country"`
	CountryName   string  `json:"country_name"`
	Features      Feature `json:"features"`
	Default       bool    `json:"default"`
}

// Feature represent a all feature inside a region
type Feature struct {
	Iaas              bool `json:"iaas"`
	Kubernetes        bool `json:"kubernetes"`
	ObjectStore       bool `json:"object_store"`
	LoadBalancer      bool `json:"loadbalancer"`
	GPU               bool `json:"gpu"`
	DBaaS             bool `json:"dbaas"`
	Volume            bool `json:"volume"`
	PaaS              bool `json:"paas"`
	PublicIPNodePools bool `json:"public_ip_node_pools"`
}

// CreateRegionRequest is the request to create a new region
type CreateRegionRequest struct {
	Code           string   `json:"code"`
	CountryISOCode string   `json:"country_iso_code" `
	Private        bool     `json:"private,omitempty"`
	AccountIDs     []string `json:"account_ids,omitempty"`
	// Kubeconfig should be a base64 encoded kubeconfig content
	Kubeconfig string `json:"kubeconfig"`
	// ComputeSoftDeletionHours can only be configured for private regions.
	ComputeSoftDeletionHours *int            `json:"compute_soft_deletion_hours" `
	Features                 map[string]bool `json:"features" `
}

// DisconnectRegionRequest is the request to disconnect a region
type DisconnectRegionRequest struct {
	Code string `json:"code"`
}

// ConnectRegionRequest is the request to connect a region
type ConnectRegionRequest struct {
	Code string `json:"code"`
}

// ListRegions returns all load balancers owned by the calling API account
func (c *Client) ListRegions() ([]Region, error) {
	resp, err := c.SendGetRequest("/v2/regions")
	if err != nil {
		return nil, decodeError(err)
	}

	regions := make([]Region, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&regions); err != nil {
		return nil, err
	}

	return regions, nil
}

// FindRegion is a function to find a region
func (c *Client) FindRegion(search string) (*Region, error) {
	allregion, err := c.ListRegions()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Region{}
	search = strings.ToUpper(search)

	for _, value := range allregion {
		name := strings.ToUpper(value.Name)
		code := strings.ToUpper(value.Code)

		if name == search || code == search {
			exactMatch = true
			result = value
		} else if strings.Contains(name, search) || strings.Contains(code, search) {
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

// GetDefaultRegion finds the default region for an account
func (c *Client) GetDefaultRegion() (*Region, error) {
	allregion, err := c.ListRegions()
	if err != nil {
		return nil, decodeError(err)
	}

	for _, region := range allregion {
		if region.Default {
			return &region, nil
		}
	}

	return nil, errors.New("no default region found")
}

// CreateRegion is a function to create a region
func (c *Client) CreateRegion(r *CreateRegionRequest) (*Region, error) {
	resp, err := c.SendPostRequest("/v2/regions", r)
	if err != nil {
		return nil, decodeError(err)
	}

	region := Region{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&region); err != nil {
		return nil, err
	}

	return &region, nil
}

// ConnectRegion connects a region to CivoAPI
func (c *Client) ConnectRegion(r *ConnectRegionRequest) error {
	_, err := c.SendPostRequest("/v2/regions/connect", r)
	if err != nil {
		return decodeError(err)
	}
	return nil
}

// DisconnectRegion disconnects a region to CivoAPI
func (c *Client) DisconnectRegion(r *DisconnectRegionRequest) error {
	_, err := c.SendPostRequest("/v2/regions/disconnect", r)
	if err != nil {
		return decodeError(err)
	}
	return nil
}
