package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// HealthCheck represents the health check configuration for an instance pool.
type HealthCheck struct {
	Port int32  `json:"port"`
	Path string `json:"path"`
}

// LoadBalancerBackend represents a backend instance being load-balanced
type LoadBalancerBackend struct {
	IP              string `json:"ip"`
	Protocol        string `json:"protocol,omitempty"`
	SourcePort      int32  `json:"source_port"`
	TargetPort      int32  `json:"target_port"`
	HealthCheckPort int32  `json:"health_check_port,omitempty"`
}

// InstancePool represents an instance pool configuration in a load balancer.
type InstancePool struct {
	Tags        []string    `json:"tags"`
	Names       []string    `json:"names"`
	Protocol    string      `json:"protocol,omitempty"`
	SourcePort  int32       `json:"source_port"`
	TargetPort  int32       `json:"target_port"`
	HealthCheck HealthCheck `json:"health_check"`
}

// LoadBalancerBackendConfig is the configuration for creating backends
type LoadBalancerBackendConfig struct {
	IP              string `json:"ip"`
	Protocol        string `json:"protocol,omitempty"`
	SourcePort      int32  `json:"source_port"`
	TargetPort      int32  `json:"target_port"`
	HealthCheckPort int32  `json:"health_check_port,omitempty"`
}

// LoadBalancerInstancePoolConfig represents an instance pool configuration in a load balancer.
type LoadBalancerInstancePoolConfig struct {
	Tags        []string    `json:"tags"`
	Names       []string    `json:"names"`
	Protocol    string      `json:"protocol,omitempty"`
	SourcePort  int32       `json:"source_port"`
	TargetPort  int32       `json:"target_port"`
	HealthCheck HealthCheck `json:"health_check"`
}

// LoadBalancer represents a load balancer configuration within Civo
type LoadBalancer struct {
	ID                           string                `json:"id"`
	Name                         string                `json:"name"`
	ServiceName                  string                `json:"service_name,omitempty"`
	NetworkID                    string                `json:"network_id,omitempty"`
	Algorithm                    string                `json:"algorithm"`
	Backends                     []LoadBalancerBackend `json:"backends,omitempty"`
	InstancePool                 []InstancePool        `json:"instance_pools,omitempty"`
	ExternalTrafficPolicy        string                `json:"external_traffic_policy,omitempty"`
	SessionAffinity              string                `json:"session_affinity,omitempty"`
	SessionAffinityConfigTimeout int32                 `json:"session_affinity_config_timeout,omitempty"`
	EnableProxyProtocol          string                `json:"enable_proxy_protocol,omitempty"`
	PublicIP                     string                `json:"public_ip"`
	PrivateIP                    string                `json:"private_ip"`
	FirewallID                   string                `json:"firewall_id"`
	ClusterID                    string                `json:"cluster_id,omitempty"`
	State                        string                `json:"state"`
	ReservedIPID                 string                `json:"reserved_ip_id,omitempty"`
	ReservedIPName               string                `json:"reserved_ip_name,omitempty"`
	ReservedIP                   string                `json:"reserved_ip,omitempty"`
	MaxConcurrentRequests        int                   `json:"max_concurrent_requests,omitempty"`
	Options                      *LoadBalancerOptions  `json:"options,omitempty"`
}

// LoadBalancerConfig represents a load balancer to be created
type LoadBalancerConfig struct {
	Region                       string                           `json:"region"`
	Name                         string                           `json:"name"`
	ServiceName                  string                           `json:"service_name,omitempty"`
	NetworkID                    string                           `json:"network_id,omitempty"`
	Algorithm                    string                           `json:"algorithm,omitempty"`
	Backends                     []LoadBalancerBackendConfig      `json:"backends,omitempty"`
	InstancePools                []LoadBalancerInstancePoolConfig `json:"instance_pools,omitempty"`
	ExternalTrafficPolicy        string                           `json:"external_traffic_policy,omitempty"`
	SessionAffinity              string                           `json:"session_affinity,omitempty"`
	SessionAffinityConfigTimeout int32                            `json:"session_affinity_config_timeout,omitempty"`
	EnableProxyProtocol          string                           `json:"enable_proxy_protocol,omitempty"`
	ClusterID                    string                           `json:"cluster_id,omitempty"`
	FirewallID                   string                           `json:"firewall_id,omitempty"`
	FirewallRules                string                           `json:"firewall_rule,omitempty"`
	MaxConcurrentRequests        *int                             `json:"max_concurrent_requests,omitempty"`
	LoadBalancerOptions          *LoadBalancerOptions             `json:"options,omitempty"`
}

// LoadBalancerOptions are additional loadbalancer options
type LoadBalancerOptions struct {
	ServerTimeout string `json:"server_timeout,omitempty"`
	ClientTimeout string `json:"client_timeout,omitempty"`
}

// LoadBalancerUpdateConfig represents a load balancer to be updated
type LoadBalancerUpdateConfig struct {
	Region                       string                           `json:"region"`
	Name                         string                           `json:"name,omitempty"`
	ServiceName                  string                           `json:"service_name,omitempty"`
	Algorithm                    string                           `json:"algorithm,omitempty"`
	Backends                     []LoadBalancerBackendConfig      `json:"backends,omitempty"`
	InstancePools                []LoadBalancerInstancePoolConfig `json:"instance_pools,omitempty"`
	ExternalTrafficPolicy        string                           `json:"external_traffic_policy,omitempty"`
	SessionAffinity              string                           `json:"session_affinity,omitempty"`
	SessionAffinityConfigTimeout int32                            `json:"session_affinity_config_timeout,omitempty"`
	EnableProxyProtocol          string                           `json:"enable_proxy_protocol,omitempty"`
	FirewallID                   string                           `json:"firewall_id,omitempty"`
	MaxConcurrentRequests        *int                             `json:"max_concurrent_requests,omitempty"`
	LoadBalancerOptions          *LoadBalancerOptions             `json:"options,omitempty"`
}

// ListLoadBalancers returns all load balancers owned by the calling API account
func (c *Client) ListLoadBalancers() ([]LoadBalancer, error) {
	resp, err := c.SendGetRequest("/v2/loadbalancers")
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := make([]LoadBalancer, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&loadbalancer); err != nil {
		return nil, decodeError(err)
	}

	return loadbalancer, nil
}

// GetLoadBalancer returns a load balancer
func (c *Client) GetLoadBalancer(id string) (*LoadBalancer, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/loadbalancers/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := &LoadBalancer{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&loadbalancer); err != nil {
		return nil, decodeError(err)
	}

	return loadbalancer, nil
}

// FindLoadBalancer finds a load balancer by either part of the ID or part of the name
func (c *Client) FindLoadBalancer(search string) (*LoadBalancer, error) {
	lbs, err := c.ListLoadBalancers()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := LoadBalancer{}

	for _, value := range lbs {
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

// CreateLoadBalancer creates a new load balancer
func (c *Client) CreateLoadBalancer(r *LoadBalancerConfig) (*LoadBalancer, error) {
	body, err := c.SendPostRequest("/v2/loadbalancers", r)
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := &LoadBalancer{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(loadbalancer); err != nil {
		return nil, err
	}

	return loadbalancer, nil
}

// UpdateLoadBalancer updates a load balancer
func (c *Client) UpdateLoadBalancer(id string, r *LoadBalancerUpdateConfig) (*LoadBalancer, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/loadbalancers/%s", id), r)
	if err != nil {
		return nil, decodeError(err)
	}

	loadbalancer := &LoadBalancer{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(loadbalancer); err != nil {
		return nil, err
	}

	return loadbalancer, nil
}

// DeleteLoadBalancer deletes a load balancer
func (c *Client) DeleteLoadBalancer(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/loadbalancers/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
