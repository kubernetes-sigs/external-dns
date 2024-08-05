package simplehosting

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Simple Hosting API client
func New(config config.Config) *SimpleHosting {
	client := client.New(config.APIKey, config.PersonalAccessToken, config.APIURL, config.SharingID, config.Debug, config.DryRun, config.Timeout)
	client.SetEndpoint("simplehosting/")
	return &SimpleHosting{client: *client}
}

// ListInstances requests the list of SimpleHosting instances
func (g *SimpleHosting) ListInstances() (instances []Instance, err error) {
	_, elements, err := g.client.GetCollection("instances", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var instance Instance
		err := json.Unmarshal(element, &instance)
		if err != nil {
			return nil, err
		}
		instances = append(instances, instance)
	}
	return instances, nil
}

// GetInstance requests a single Instance
func (g *SimpleHosting) GetInstance(instanceId string) (simplehostingResponse Instance, err error) {
	_, err = g.client.Get("instances/"+instanceId, nil, &simplehostingResponse)
	return
}

// CreateInstance creates a SimpleHosting instance
func (g *SimpleHosting) CreateInstance(req CreateInstanceRequest) (instanceId string, err error) {
	header, err := g.client.Post("instances", req, nil)
	if err != nil {
		return "", err
	}
	// We have to extract the instance ID from the
	// Content-Location response header.
	location := header.Get("Content-Location")
	endpoint := g.client.GetEndpoint() + "instances/"
	if strings.HasPrefix(location, endpoint) {
		return strings.TrimPrefix(location, endpoint), nil
	} else {
		return "", fmt.Errorf("Could not extract the instance ID from '%s'", location)
	}
}

// CreateInstance deletes a SimpleHosting instance
func (g *SimpleHosting) DeleteInstance(instanceId string) (response ErrorResponse, err error) {
	_, err = g.client.Delete("instances/"+instanceId, nil, &response)
	return
}

// // GetVhost requests a single Vhost
func (g *SimpleHosting) GetVhost(instanceId string, fqdn string) (response Vhost, err error) {
	_, err = g.client.Get("instances/"+instanceId+"/vhosts/"+fqdn, nil, &response)
	return
}

// ListVhosts lists vhosts of a Simple Hosting instance
func (g *SimpleHosting) ListVhosts(instanceId string) (vhosts []Vhost, err error) {
	_, elements, err := g.client.GetCollection("instances/"+instanceId+"/vhosts", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var vhost Vhost
		err := json.Unmarshal(element, &vhost)
		if err != nil {
			return nil, err
		}
		vhosts = append(vhosts, vhost)
	}
	return vhosts, nil
}

// ListVhosts creates a vhost for a Simple Hosting instance
func (g *SimpleHosting) CreateVhost(instanceId string, req CreateVhostRequest) (response Vhost, err error) {
	_, err = g.client.Post("instances/"+instanceId+"/vhosts", req, &response)
	if err != nil {
		return Vhost{}, err
	}
	return
}

// UpdateVhost updates a vhost for a Simple Hosting instance
func (g *SimpleHosting) UpdateVhost(instanceId string, fqdn string, req PatchVhostRequest) (response PatchVhostResponse, err error) {
	_, err = g.client.Patch("instances/"+instanceId+"/vhosts/"+fqdn, req, &response)
	if err != nil {
		return PatchVhostResponse{}, err
	}
	return
}

// ListVhosts deletes vhosts of a Simple Hosting instance
func (g *SimpleHosting) DeleteVhost(instanceId string, fqdn string) (response ErrorResponse, err error) {
	_, err = g.client.Delete("instances/"+instanceId+"/vhosts/"+fqdn, nil, &response)
	return
}
