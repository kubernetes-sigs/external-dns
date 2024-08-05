package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// LoadBalancerPool represents a load balancer pool's properties.
type LoadBalancerPool struct {
	ID                string                      `json:"id,omitempty"`
	CreatedOn         *time.Time                  `json:"created_on,omitempty"`
	ModifiedOn        *time.Time                  `json:"modified_on,omitempty"`
	Description       string                      `json:"description"`
	Name              string                      `json:"name"`
	Enabled           bool                        `json:"enabled"`
	MinimumOrigins    int                         `json:"minimum_origins,omitempty"`
	Monitor           string                      `json:"monitor,omitempty"`
	Origins           []LoadBalancerOrigin        `json:"origins"`
	NotificationEmail string                      `json:"notification_email,omitempty"`
	Latitude          *float32                    `json:"latitude,omitempty"`
	Longitude         *float32                    `json:"longitude,omitempty"`
	LoadShedding      *LoadBalancerLoadShedding   `json:"load_shedding,omitempty"`
	OriginSteering    *LoadBalancerOriginSteering `json:"origin_steering,omitempty"`

	// CheckRegions defines the geographic region(s) from where to run health-checks from - e.g. "WNAM", "WEU", "SAF", "SAM".
	// Providing a null/empty value means "all regions", which may not be available to all plan types.
	CheckRegions []string `json:"check_regions"`
}

// LoadBalancerOrigin represents a Load Balancer origin's properties.
type LoadBalancerOrigin struct {
	Name    string              `json:"name"`
	Address string              `json:"address"`
	Enabled bool                `json:"enabled"`
	Weight  float64             `json:"weight"`
	Header  map[string][]string `json:"header"`
}

// LoadBalancerOriginSteering controls origin selection for new sessions and traffic without session affinity.
type LoadBalancerOriginSteering struct {
	// Policy defaults to "random" (weighted) when empty or unspecified.
	Policy string `json:"policy,omitempty"`
}

// LoadBalancerMonitor represents a load balancer monitor's properties.
type LoadBalancerMonitor struct {
	ID              string              `json:"id,omitempty"`
	CreatedOn       *time.Time          `json:"created_on,omitempty"`
	ModifiedOn      *time.Time          `json:"modified_on,omitempty"`
	Type            string              `json:"type"`
	Description     string              `json:"description"`
	Method          string              `json:"method"`
	Path            string              `json:"path"`
	Header          map[string][]string `json:"header"`
	Timeout         int                 `json:"timeout"`
	Retries         int                 `json:"retries"`
	Interval        int                 `json:"interval"`
	Port            uint16              `json:"port,omitempty"`
	ExpectedBody    string              `json:"expected_body"`
	ExpectedCodes   string              `json:"expected_codes"`
	FollowRedirects bool                `json:"follow_redirects"`
	AllowInsecure   bool                `json:"allow_insecure"`
	ProbeZone       string              `json:"probe_zone"`
}

// LoadBalancer represents a load balancer's properties.
type LoadBalancer struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	ID                        string                     `json:"id,omitempty"`
	CreatedOn                 *time.Time                 `json:"created_on,omitempty"`
	ModifiedOn                *time.Time                 `json:"modified_on,omitempty"`
	Description               string                     `json:"description"`
	Name                      string                     `json:"name"`
	TTL                       int                        `json:"ttl,omitempty"`
	FallbackPool              string                     `json:"fallback_pool"`
	DefaultPools              []string                   `json:"default_pools"`
	RegionPools               map[string][]string        `json:"region_pools"`
	PopPools                  map[string][]string        `json:"pop_pools"`
	Proxied                   bool                       `json:"proxied"`
	Enabled                   *bool                      `json:"enabled,omitempty"`
	Persistence               string                     `json:"session_affinity,omitempty"`
	PersistenceTTL            int                        `json:"session_affinity_ttl,omitempty"`
	SessionAffinityAttributes *SessionAffinityAttributes `json:"session_affinity_attributes,omitempty"`

	// SteeringPolicy controls pool selection logic.
	// "off" select pools in DefaultPools order
	// "geo" select pools based on RegionPools/PopPools
	// "dynamic_latency" select pools based on RTT (requires health checks)
	// "random" selects pools in a random order
	// "" maps to "geo" if RegionPools or PopPools have entries otherwise "off"
	SteeringPolicy string `json:"steering_policy,omitempty"`
}

// SessionAffinityAttributes represents the fields used to set attributes in a load balancer session affinity cookie.
type SessionAffinityAttributes struct {
	SameSite string `json:"samesite,omitempty"`
	Secure   string `json:"secure,omitempty"`
}

// LoadBalancerOriginHealth represents the health of the origin.
type LoadBalancerOriginHealth struct {
	Healthy       bool     `json:"healthy,omitempty"`
	RTT           Duration `json:"rtt,omitempty"`
	FailureReason string   `json:"failure_reason,omitempty"`
	ResponseCode  int      `json:"response_code,omitempty"`
}

// LoadBalancerPoolPopHealth represents the health of the pool for given PoP.
type LoadBalancerPoolPopHealth struct {
	Healthy bool                                  `json:"healthy,omitempty"`
	Origins []map[string]LoadBalancerOriginHealth `json:"origins,omitempty"`
}

// LoadBalancerPoolHealth represents the healthchecks from different PoPs for a pool.
type LoadBalancerPoolHealth struct {
	ID        string                               `json:"pool_id,omitempty"`
	PopHealth map[string]LoadBalancerPoolPopHealth `json:"pop_health,omitempty"`
}

// loadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type loadBalancerPoolResponse struct {
	Response
	Result LoadBalancerPool `json:"result"`
}

// loadBalancerPoolListResponse represents the response from the List Pools endpoint.
type loadBalancerPoolListResponse struct {
	Response
	Result     []LoadBalancerPool `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// loadBalancerMonitorResponse represents the response from the load balancer monitor endpoints.
type loadBalancerMonitorResponse struct {
	Response
	Result LoadBalancerMonitor `json:"result"`
}

// loadBalancerMonitorListResponse represents the response from the List Monitors endpoint.
type loadBalancerMonitorListResponse struct {
	Response
	Result     []LoadBalancerMonitor `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// loadBalancerResponse represents the response from the load balancer endpoints.
type loadBalancerResponse struct {
	Response
	Result LoadBalancer `json:"result"`
}

// loadBalancerListResponse represents the response from the List Load Balancers endpoint.
type loadBalancerListResponse struct {
	Response
	Result     []LoadBalancer `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}

// loadBalancerPoolHealthResponse represents the response from the Pool Health Details endpoint.
type loadBalancerPoolHealthResponse struct {
	Response
	Result LoadBalancerPoolHealth `json:"result"`
}

// CreateLoadBalancerPool creates a new load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-create-pool
func (api *API) CreateLoadBalancerPool(pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools"
	res, err := api.makeRequest("POST", uri, pool)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-list-pools
func (api *API) ListLoadBalancerPools() ([]LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerPoolDetails returns the details for a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-details
func (api *API) LoadBalancerPoolDetails(poolID string) (LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools/" + poolID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerPool disables and deletes a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-delete-pool
func (api *API) DeleteLoadBalancerPool(poolID string) error {
	uri := api.userBaseURL("/user") + "/load_balancers/pools/" + poolID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// ModifyLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-update-pool
func (api *API) ModifyLoadBalancerPool(pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools/" + pool.ID
	res, err := api.makeRequest("PUT", uri, pool)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancerMonitor creates a new load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-create-monitor
func (api *API) CreateLoadBalancerMonitor(monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors"
	res, err := api.makeRequest("POST", uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerMonitors lists load balancer monitors connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-list-monitors
func (api *API) ListLoadBalancerMonitors() ([]LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerMonitorDetails returns the details for a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-monitor-details
func (api *API) LoadBalancerMonitorDetails(monitorID string) (LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors/" + monitorID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerMonitor disables and deletes a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-delete-monitor
func (api *API) DeleteLoadBalancerMonitor(monitorID string) error {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors/" + monitorID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// ModifyLoadBalancerMonitor modifies a configured load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-update-monitor
func (api *API) ModifyLoadBalancerMonitor(monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors/" + monitor.ID
	res, err := api.makeRequest("PUT", uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancer creates a new load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-create-load-balancer
func (api *API) CreateLoadBalancer(zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers"
	res, err := api.makeRequest("POST", uri, lb)
	if err != nil {
		return LoadBalancer{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancers lists load balancers configured on a zone.
//
// API reference: https://api.cloudflare.com/#load-balancers-list-load-balancers
func (api *API) ListLoadBalancers(zoneID string) ([]LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerDetails returns the details for a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-load-balancer-details
func (api *API) LoadBalancerDetails(zoneID, lbID string) (LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers/" + lbID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LoadBalancer{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancer disables and deletes a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-delete-load-balancer
func (api *API) DeleteLoadBalancer(zoneID, lbID string) error {
	uri := "/zones/" + zoneID + "/load_balancers/" + lbID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// ModifyLoadBalancer modifies a configured load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-update-load-balancer
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	ID             string              `json:"id,omitempty"`
	CreatedOn      *time.Time          `json:"created_on,omitempty"`
	ModifiedOn     *time.Time          `json:"modified_on,omitempty"`
	Description    string              `json:"description"`
	Name           string              `json:"name"`
	TTL            int                 `json:"ttl,omitempty"`
	FallbackPool   string              `json:"fallback_pool"`
	DefaultPools   []string            `json:"default_pools"`
	RegionPools    map[string][]string `json:"region_pools"`
	PopPools       map[string][]string `json:"pop_pools"`
	Proxied        bool                `json:"proxied"`
	Enabled        *bool               `json:"enabled,omitempty"`
	Persistence    string              `json:"session_affinity,omitempty"`
	PersistenceTTL int                 `json:"session_affinity_ttl,omitempty"`
||||||| parent of 5ce8c7613 (update vendored files)
	ID             string              `json:"id,omitempty"`
	CreatedOn      *time.Time          `json:"created_on,omitempty"`
	ModifiedOn     *time.Time          `json:"modified_on,omitempty"`
	Description    string              `json:"description"`
	Name           string              `json:"name"`
	TTL            int                 `json:"ttl,omitempty"`
	FallbackPool   string              `json:"fallback_pool"`
	DefaultPools   []string            `json:"default_pools"`
	RegionPools    map[string][]string `json:"region_pools"`
	PopPools       map[string][]string `json:"pop_pools"`
	Proxied        bool                `json:"proxied"`
	Enabled        *bool               `json:"enabled,omitempty"`
	Persistence    string              `json:"session_affinity,omitempty"`
	PersistenceTTL int                 `json:"session_affinity_ttl,omitempty"`
=======
	ID                        string                     `json:"id,omitempty"`
	CreatedOn                 *time.Time                 `json:"created_on,omitempty"`
	ModifiedOn                *time.Time                 `json:"modified_on,omitempty"`
	Description               string                     `json:"description"`
	Name                      string                     `json:"name"`
	TTL                       int                        `json:"ttl,omitempty"`
	FallbackPool              string                     `json:"fallback_pool"`
	DefaultPools              []string                   `json:"default_pools"`
	RegionPools               map[string][]string        `json:"region_pools"`
	PopPools                  map[string][]string        `json:"pop_pools"`
	Proxied                   bool                       `json:"proxied"`
	Enabled                   *bool                      `json:"enabled,omitempty"`
	Persistence               string                     `json:"session_affinity,omitempty"`
	PersistenceTTL            int                        `json:"session_affinity_ttl,omitempty"`
	SessionAffinityAttributes *SessionAffinityAttributes `json:"session_affinity_attributes,omitempty"`
>>>>>>> 5ce8c7613 (update vendored files)

	// SteeringPolicy controls pool selection logic.
	// "off" select pools in DefaultPools order
	// "geo" select pools based on RegionPools/PopPools
	// "dynamic_latency" select pools based on RTT (requires health checks)
	// "random" selects pools in a random order
	// "" maps to "geo" if RegionPools or PopPools have entries otherwise "off"
	SteeringPolicy string `json:"steering_policy,omitempty"`
}

// SessionAffinityAttributes represents the fields used to set attributes in a load balancer session affinity cookie.
type SessionAffinityAttributes struct {
	SameSite string `json:"samesite,omitempty"`
	Secure   string `json:"secure,omitempty"`
}

// LoadBalancerOriginHealth represents the health of the origin.
type LoadBalancerOriginHealth struct {
	Healthy       bool     `json:"healthy,omitempty"`
	RTT           Duration `json:"rtt,omitempty"`
	FailureReason string   `json:"failure_reason,omitempty"`
	ResponseCode  int      `json:"response_code,omitempty"`
}

// LoadBalancerPoolPopHealth represents the health of the pool for given PoP.
type LoadBalancerPoolPopHealth struct {
	Healthy bool                                  `json:"healthy,omitempty"`
	Origins []map[string]LoadBalancerOriginHealth `json:"origins,omitempty"`
}

// LoadBalancerPoolHealth represents the healthchecks from different PoPs for a pool.
type LoadBalancerPoolHealth struct {
	ID        string                               `json:"pool_id,omitempty"`
	PopHealth map[string]LoadBalancerPoolPopHealth `json:"pop_health,omitempty"`
}

// loadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type loadBalancerPoolResponse struct {
	Response
	Result LoadBalancerPool `json:"result"`
}

// loadBalancerPoolListResponse represents the response from the List Pools endpoint.
type loadBalancerPoolListResponse struct {
	Response
	Result     []LoadBalancerPool `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// loadBalancerMonitorResponse represents the response from the load balancer monitor endpoints.
type loadBalancerMonitorResponse struct {
	Response
	Result LoadBalancerMonitor `json:"result"`
}

// loadBalancerMonitorListResponse represents the response from the List Monitors endpoint.
type loadBalancerMonitorListResponse struct {
	Response
	Result     []LoadBalancerMonitor `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// loadBalancerResponse represents the response from the load balancer endpoints.
type loadBalancerResponse struct {
	Response
	Result LoadBalancer `json:"result"`
}

// loadBalancerListResponse represents the response from the List Load Balancers endpoint.
type loadBalancerListResponse struct {
	Response
	Result     []LoadBalancer `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}

// loadBalancerPoolHealthResponse represents the response from the Pool Health Details endpoint.
type loadBalancerPoolHealthResponse struct {
	Response
	Result LoadBalancerPoolHealth `json:"result"`
}

// CreateLoadBalancerPool creates a new load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-create-pool
func (api *API) CreateLoadBalancerPool(pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools"
	res, err := api.makeRequest("POST", uri, pool)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-list-pools
func (api *API) ListLoadBalancerPools() ([]LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerPoolDetails returns the details for a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-details
func (api *API) LoadBalancerPoolDetails(poolID string) (LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools/" + poolID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerPool disables and deletes a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-delete-pool
func (api *API) DeleteLoadBalancerPool(poolID string) error {
	uri := api.userBaseURL("/user") + "/load_balancers/pools/" + poolID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// ModifyLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-update-pool
func (api *API) ModifyLoadBalancerPool(pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/pools/" + pool.ID
	res, err := api.makeRequest("PUT", uri, pool)
	if err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancerMonitor creates a new load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-create-monitor
func (api *API) CreateLoadBalancerMonitor(monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors"
	res, err := api.makeRequest("POST", uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerMonitors lists load balancer monitors connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-list-monitors
func (api *API) ListLoadBalancerMonitors() ([]LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerMonitorDetails returns the details for a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-monitor-details
func (api *API) LoadBalancerMonitorDetails(monitorID string) (LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors/" + monitorID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerMonitor disables and deletes a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-delete-monitor
func (api *API) DeleteLoadBalancerMonitor(monitorID string) error {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors/" + monitorID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// ModifyLoadBalancerMonitor modifies a configured load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-update-monitor
func (api *API) ModifyLoadBalancerMonitor(monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := api.userBaseURL("/user") + "/load_balancers/monitors/" + monitor.ID
	res, err := api.makeRequest("PUT", uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancer creates a new load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-create-load-balancer
func (api *API) CreateLoadBalancer(zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers"
	res, err := api.makeRequest("POST", uri, lb)
	if err != nil {
		return LoadBalancer{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancers lists load balancers configured on a zone.
//
// API reference: https://api.cloudflare.com/#load-balancers-list-load-balancers
func (api *API) ListLoadBalancers(zoneID string) ([]LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerDetails returns the details for a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-load-balancer-details
func (api *API) LoadBalancerDetails(zoneID, lbID string) (LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers/" + lbID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return LoadBalancer{}, errors.Wrap(err, errMakeRequestError)
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancer disables and deletes a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-delete-load-balancer
func (api *API) DeleteLoadBalancer(zoneID, lbID string) error {
	uri := "/zones/" + zoneID + "/load_balancers/" + lbID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
}

// ModifyLoadBalancer modifies a configured load balancer.
//
<<<<<<< HEAD
// API reference: https://api.cloudflare.com/#load-balancers-modify-a-load-balancer
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// API reference: https://api.cloudflare.com/#load-balancers-modify-a-load-balancer
=======
// API reference: https://api.cloudflare.com/#load-balancers-update-load-balancer
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	ID             string              `json:"id,omitempty"`
	CreatedOn      *time.Time          `json:"created_on,omitempty"`
	ModifiedOn     *time.Time          `json:"modified_on,omitempty"`
	Description    string              `json:"description"`
	Name           string              `json:"name"`
	TTL            int                 `json:"ttl,omitempty"`
	FallbackPool   string              `json:"fallback_pool"`
	DefaultPools   []string            `json:"default_pools"`
	RegionPools    map[string][]string `json:"region_pools"`
	PopPools       map[string][]string `json:"pop_pools"`
	Proxied        bool                `json:"proxied"`
	Enabled        *bool               `json:"enabled,omitempty"`
	Persistence    string              `json:"session_affinity,omitempty"`
	PersistenceTTL int                 `json:"session_affinity_ttl,omitempty"`
||||||| parent of 6b7ce455e (update vendored files)
	ID             string              `json:"id,omitempty"`
	CreatedOn      *time.Time          `json:"created_on,omitempty"`
	ModifiedOn     *time.Time          `json:"modified_on,omitempty"`
	Description    string              `json:"description"`
	Name           string              `json:"name"`
	TTL            int                 `json:"ttl,omitempty"`
	FallbackPool   string              `json:"fallback_pool"`
	DefaultPools   []string            `json:"default_pools"`
	RegionPools    map[string][]string `json:"region_pools"`
	PopPools       map[string][]string `json:"pop_pools"`
	Proxied        bool                `json:"proxied"`
	Enabled        *bool               `json:"enabled,omitempty"`
	Persistence    string              `json:"session_affinity,omitempty"`
	PersistenceTTL int                 `json:"session_affinity_ttl,omitempty"`
=======
	ID                        string                     `json:"id,omitempty"`
	CreatedOn                 *time.Time                 `json:"created_on,omitempty"`
	ModifiedOn                *time.Time                 `json:"modified_on,omitempty"`
	Description               string                     `json:"description"`
	Name                      string                     `json:"name"`
	TTL                       int                        `json:"ttl,omitempty"`
	FallbackPool              string                     `json:"fallback_pool"`
	DefaultPools              []string                   `json:"default_pools"`
	RegionPools               map[string][]string        `json:"region_pools"`
	PopPools                  map[string][]string        `json:"pop_pools"`
	Proxied                   bool                       `json:"proxied"`
	Enabled                   *bool                      `json:"enabled,omitempty"`
	Persistence               string                     `json:"session_affinity,omitempty"`
	PersistenceTTL            int                        `json:"session_affinity_ttl,omitempty"`
	SessionAffinityAttributes *SessionAffinityAttributes `json:"session_affinity_attributes,omitempty"`
	Rules                     []*LoadBalancerRule        `json:"rules,omitempty"`
>>>>>>> 6b7ce455e (update vendored files)

	// SteeringPolicy controls pool selection logic.
	// "off" select pools in DefaultPools order
	// "geo" select pools based on RegionPools/PopPools
	// "dynamic_latency" select pools based on RTT (requires health checks)
	// "random" selects pools in a random order
	// "proximity" select pools based on 'distance' from request
	// "" maps to "geo" if RegionPools or PopPools have entries otherwise "off"
	SteeringPolicy string `json:"steering_policy,omitempty"`
}

// LoadBalancerLoadShedding contains the settings for controlling load shedding
type LoadBalancerLoadShedding struct {
	DefaultPercent float32 `json:"default_percent,omitempty"`
	DefaultPolicy  string  `json:"default_policy,omitempty"`
	SessionPercent float32 `json:"session_percent,omitempty"`
	SessionPolicy  string  `json:"session_policy,omitempty"`
}

// LoadBalancerRule represents a single rule entry for a Load Balancer. Each rules
// is run one after the other in priority order. Disabled rules are skipped.
type LoadBalancerRule struct {
	// Name is required but is only used for human readability
	Name string `json:"name"`
	// Priority controls the order of rule execution the lowest value will be invoked first
	Priority int  `json:"priority"`
	Disabled bool `json:"disabled"`

	Condition string                    `json:"condition"`
	Overrides LoadBalancerRuleOverrides `json:"overrides"`

	// Terminates flag this rule as 'terminating'. No further rules will
	// be executed after this one.
	Terminates bool `json:"terminates,omitempty"`

	// FixedResponse if set and the condition is true we will not run
	// routing logic but rather directly respond with the provided fields.
	// FixedResponse implies terminates.
	FixedResponse *LoadBalancerFixedResponseData `json:"fixed_response,omitempty"`
}

// LoadBalancerFixedResponseData contains all the data needed to generate
// a fixed response from a Load Balancer. This behavior can be enabled via Rules.
type LoadBalancerFixedResponseData struct {
	// MessageBody data to write into the http body
	MessageBody string `json:"message_body,omitempty"`
	// StatusCode the http status code to response with
	StatusCode int `json:"status_code,omitempty"`
	// ContentType value of the http 'content-type' header
	ContentType string `json:"content_type,omitempty"`
	// Location value of the http 'location' header
	Location string `json:"location,omitempty"`
}

// LoadBalancerRuleOverrides are the set of field overridable by the rules system.
type LoadBalancerRuleOverrides struct {
	// session affinity
	Persistence    string `json:"session_affinity,omitempty"`
	PersistenceTTL *uint  `json:"session_affinity_ttl,omitempty"`

	SessionAffinityAttrs *LoadBalancerRuleOverridesSessionAffinityAttrs `json:"session_affinity_attributes,omitempty"`

	TTL uint `json:"ttl,omitempty"`

	SteeringPolicy string `json:"steering_policy,omitempty"`
	FallbackPool   string `json:"fallback_pool,omitempty"`

	DefaultPools []string            `json:"default_pools,omitempty"`
	PoPPools     map[string][]string `json:"pop_pools,omitempty"`
	RegionPools  map[string][]string `json:"region_pools,omitempty"`
}

// LoadBalancerRuleOverridesSessionAffinityAttrs mimics SessionAffinityAttributes without the
// DrainDuration field as that field can not be overwritten via rules.
type LoadBalancerRuleOverridesSessionAffinityAttrs struct {
	SameSite string `json:"samesite,omitempty"`
	Secure   string `json:"secure,omitempty"`
}

// SessionAffinityAttributes represents the fields used to set attributes in a load balancer session affinity cookie.
type SessionAffinityAttributes struct {
	SameSite      string `json:"samesite,omitempty"`
	Secure        string `json:"secure,omitempty"`
	DrainDuration int    `json:"drain_duration,omitempty"`
}

// LoadBalancerOriginHealth represents the health of the origin.
type LoadBalancerOriginHealth struct {
	Healthy       bool     `json:"healthy,omitempty"`
	RTT           Duration `json:"rtt,omitempty"`
	FailureReason string   `json:"failure_reason,omitempty"`
	ResponseCode  int      `json:"response_code,omitempty"`
}

// LoadBalancerPoolPopHealth represents the health of the pool for given PoP.
type LoadBalancerPoolPopHealth struct {
	Healthy bool                                  `json:"healthy,omitempty"`
	Origins []map[string]LoadBalancerOriginHealth `json:"origins,omitempty"`
}

// LoadBalancerPoolHealth represents the healthchecks from different PoPs for a pool.
type LoadBalancerPoolHealth struct {
	ID        string                               `json:"pool_id,omitempty"`
	PopHealth map[string]LoadBalancerPoolPopHealth `json:"pop_health,omitempty"`
}

// loadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type loadBalancerPoolResponse struct {
	Response
	Result LoadBalancerPool `json:"result"`
}

// loadBalancerPoolListResponse represents the response from the List Pools endpoint.
type loadBalancerPoolListResponse struct {
	Response
	Result     []LoadBalancerPool `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// loadBalancerMonitorResponse represents the response from the load balancer monitor endpoints.
type loadBalancerMonitorResponse struct {
	Response
	Result LoadBalancerMonitor `json:"result"`
}

// loadBalancerMonitorListResponse represents the response from the List Monitors endpoint.
type loadBalancerMonitorListResponse struct {
	Response
	Result     []LoadBalancerMonitor `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// loadBalancerResponse represents the response from the load balancer endpoints.
type loadBalancerResponse struct {
	Response
	Result LoadBalancer `json:"result"`
}

// loadBalancerListResponse represents the response from the List Load Balancers endpoint.
type loadBalancerListResponse struct {
	Response
	Result     []LoadBalancer `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}

// loadBalancerPoolHealthResponse represents the response from the Pool Health Details endpoint.
type loadBalancerPoolHealthResponse struct {
	Response
	Result LoadBalancerPoolHealth `json:"result"`
}

// CreateLoadBalancerPool creates a new load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-create-pool
func (api *API) CreateLoadBalancerPool(ctx context.Context, pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, pool)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-list-pools
func (api *API) ListLoadBalancerPools(ctx context.Context) ([]LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerPoolListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerPoolDetails returns the details for a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-details
func (api *API) LoadBalancerPoolDetails(ctx context.Context, poolID string) (LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s", api.userBaseURL("/user"), poolID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerPool disables and deletes a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-delete-pool
func (api *API) DeleteLoadBalancerPool(ctx context.Context, poolID string) error {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s", api.userBaseURL("/user"), poolID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// ModifyLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-update-pool
func (api *API) ModifyLoadBalancerPool(ctx context.Context, pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s", api.userBaseURL("/user"), pool.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, pool)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancerMonitor creates a new load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-create-monitor
func (api *API) CreateLoadBalancerMonitor(ctx context.Context, monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancerMonitors lists load balancer monitors connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-list-monitors
func (api *API) ListLoadBalancerMonitors(ctx context.Context) ([]LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerMonitorListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerMonitorDetails returns the details for a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-monitor-details
func (api *API) LoadBalancerMonitorDetails(ctx context.Context, monitorID string) (LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors/%s", api.userBaseURL("/user"), monitorID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancerMonitor disables and deletes a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-delete-monitor
func (api *API) DeleteLoadBalancerMonitor(ctx context.Context, monitorID string) error {
	uri := fmt.Sprintf("%s/load_balancers/monitors/%s", api.userBaseURL("/user"), monitorID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// ModifyLoadBalancerMonitor modifies a configured load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-update-monitor
func (api *API) ModifyLoadBalancerMonitor(ctx context.Context, monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors/%s", api.userBaseURL("/user"), monitor.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// CreateLoadBalancer creates a new load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-create-load-balancer
func (api *API) CreateLoadBalancer(ctx context.Context, zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, lb)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListLoadBalancers lists load balancers configured on a zone.
//
// API reference: https://api.cloudflare.com/#load-balancers-list-load-balancers
func (api *API) ListLoadBalancers(ctx context.Context, zoneID string) ([]LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// LoadBalancerDetails returns the details for a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-load-balancer-details
func (api *API) LoadBalancerDetails(ctx context.Context, zoneID, lbID string) (LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", zoneID, lbID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteLoadBalancer disables and deletes a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-delete-load-balancer
func (api *API) DeleteLoadBalancer(ctx context.Context, zoneID, lbID string) error {
	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", zoneID, lbID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// ModifyLoadBalancer modifies a configured load balancer.
//
<<<<<<< HEAD
// API reference: https://api.cloudflare.com/#load-balancers-modify-a-load-balancer
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
func (api *API) ModifyLoadBalancer(zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers/" + lb.ID
	res, err := api.makeRequest("PUT", uri, lb)
||||||| parent of 6b7ce455e (update vendored files)
// API reference: https://api.cloudflare.com/#load-balancers-modify-a-load-balancer
func (api *API) ModifyLoadBalancer(zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := "/zones/" + zoneID + "/load_balancers/" + lb.ID
	res, err := api.makeRequest("PUT", uri, lb)
=======
// API reference: https://api.cloudflare.com/#load-balancers-update-load-balancer
func (api *API) ModifyLoadBalancer(ctx context.Context, zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", zoneID, lb.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, lb)
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// PoolHealthDetails fetches the latest healtcheck details for a single pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-health-details
func (api *API) PoolHealthDetails(ctx context.Context, poolID string) (LoadBalancerPoolHealth, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s/health", api.userBaseURL("/user"), poolID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerPoolHealth{}, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// LoadBalancerPool represents a load balancer pool's properties.
type LoadBalancerPool struct {
	ID                string                      `json:"id,omitempty"`
	CreatedOn         *time.Time                  `json:"created_on,omitempty"`
	ModifiedOn        *time.Time                  `json:"modified_on,omitempty"`
	Description       string                      `json:"description"`
	Name              string                      `json:"name"`
	Enabled           bool                        `json:"enabled"`
	MinimumOrigins    int                         `json:"minimum_origins,omitempty"`
	Monitor           string                      `json:"monitor,omitempty"`
	Origins           []LoadBalancerOrigin        `json:"origins"`
	NotificationEmail string                      `json:"notification_email,omitempty"`
	Latitude          *float32                    `json:"latitude,omitempty"`
	Longitude         *float32                    `json:"longitude,omitempty"`
	LoadShedding      *LoadBalancerLoadShedding   `json:"load_shedding,omitempty"`
	OriginSteering    *LoadBalancerOriginSteering `json:"origin_steering,omitempty"`

	// CheckRegions defines the geographic region(s) from where to run health-checks from - e.g. "WNAM", "WEU", "SAF", "SAM".
	// Providing a null/empty value means "all regions", which may not be available to all plan types.
	CheckRegions []string `json:"check_regions"`
}

// LoadBalancerOrigin represents a Load Balancer origin's properties.
type LoadBalancerOrigin struct {
	Name    string              `json:"name"`
	Address string              `json:"address"`
	Enabled bool                `json:"enabled"`
	Weight  float64             `json:"weight"`
	Header  map[string][]string `json:"header"`
}

// LoadBalancerOriginSteering controls origin selection for new sessions and traffic without session affinity.
type LoadBalancerOriginSteering struct {
	// Policy defaults to "random" (weighted) when empty or unspecified.
	Policy string `json:"policy,omitempty"`
}

// LoadBalancerMonitor represents a load balancer monitor's properties.
type LoadBalancerMonitor struct {
	ID              string              `json:"id,omitempty"`
	CreatedOn       *time.Time          `json:"created_on,omitempty"`
	ModifiedOn      *time.Time          `json:"modified_on,omitempty"`
	Type            string              `json:"type"`
	Description     string              `json:"description"`
	Method          string              `json:"method"`
	Path            string              `json:"path"`
	Header          map[string][]string `json:"header"`
	Timeout         int                 `json:"timeout"`
	Retries         int                 `json:"retries"`
	Interval        int                 `json:"interval"`
	Port            uint16              `json:"port,omitempty"`
	ExpectedBody    string              `json:"expected_body"`
	ExpectedCodes   string              `json:"expected_codes"`
	FollowRedirects bool                `json:"follow_redirects"`
	AllowInsecure   bool                `json:"allow_insecure"`
	ProbeZone       string              `json:"probe_zone"`
}

// LoadBalancer represents a load balancer's properties.
type LoadBalancer struct {
	ID                        string                     `json:"id,omitempty"`
	CreatedOn                 *time.Time                 `json:"created_on,omitempty"`
	ModifiedOn                *time.Time                 `json:"modified_on,omitempty"`
	Description               string                     `json:"description"`
	Name                      string                     `json:"name"`
	TTL                       int                        `json:"ttl,omitempty"`
	FallbackPool              string                     `json:"fallback_pool"`
	DefaultPools              []string                   `json:"default_pools"`
	RegionPools               map[string][]string        `json:"region_pools"`
	PopPools                  map[string][]string        `json:"pop_pools"`
	CountryPools              map[string][]string        `json:"country_pools"`
	Proxied                   bool                       `json:"proxied"`
	Enabled                   *bool                      `json:"enabled,omitempty"`
	Persistence               string                     `json:"session_affinity,omitempty"`
	PersistenceTTL            int                        `json:"session_affinity_ttl,omitempty"`
	SessionAffinityAttributes *SessionAffinityAttributes `json:"session_affinity_attributes,omitempty"`
	Rules                     []*LoadBalancerRule        `json:"rules,omitempty"`
	RandomSteering            *RandomSteering            `json:"random_steering,omitempty"`

	// SteeringPolicy controls pool selection logic.
	// "off" select pools in DefaultPools order
	// "geo" select pools based on RegionPools/PopPools/CountryPools
	// "dynamic_latency" select pools based on RTT (requires health checks)
	// "random" selects pools in a random order
	// "proximity" select pools based on 'distance' from request
	// "" maps to "geo" if RegionPools or PopPools or CountryPools have entries otherwise "off"
	SteeringPolicy string `json:"steering_policy,omitempty"`
}

// LoadBalancerLoadShedding contains the settings for controlling load shedding.
type LoadBalancerLoadShedding struct {
	DefaultPercent float32 `json:"default_percent,omitempty"`
	DefaultPolicy  string  `json:"default_policy,omitempty"`
	SessionPercent float32 `json:"session_percent,omitempty"`
	SessionPolicy  string  `json:"session_policy,omitempty"`
}

// LoadBalancerRule represents a single rule entry for a Load Balancer. Each rules
// is run one after the other in priority order. Disabled rules are skipped.
type LoadBalancerRule struct {
	Overrides LoadBalancerRuleOverrides `json:"overrides"`

	// Name is required but is only used for human readability
	Name string `json:"name"`

	Condition string `json:"condition"`

	// Priority controls the order of rule execution the lowest value will be invoked first
	Priority int `json:"priority"`

	// FixedResponse if set and the condition is true we will not run
	// routing logic but rather directly respond with the provided fields.
	// FixedResponse implies terminates.
	FixedResponse *LoadBalancerFixedResponseData `json:"fixed_response,omitempty"`

	Disabled bool `json:"disabled"`

	// Terminates flag this rule as 'terminating'. No further rules will
	// be executed after this one.
	Terminates bool `json:"terminates,omitempty"`
}

// LoadBalancerFixedResponseData contains all the data needed to generate
// a fixed response from a Load Balancer. This behavior can be enabled via Rules.
type LoadBalancerFixedResponseData struct {
	// MessageBody data to write into the http body
	MessageBody string `json:"message_body,omitempty"`
	// StatusCode the http status code to response with
	StatusCode int `json:"status_code,omitempty"`
	// ContentType value of the http 'content-type' header
	ContentType string `json:"content_type,omitempty"`
	// Location value of the http 'location' header
	Location string `json:"location,omitempty"`
}

// LoadBalancerRuleOverrides are the set of field overridable by the rules system.
type LoadBalancerRuleOverrides struct {
	// session affinity
	Persistence    string `json:"session_affinity,omitempty"`
	PersistenceTTL *uint  `json:"session_affinity_ttl,omitempty"`

	SessionAffinityAttrs *LoadBalancerRuleOverridesSessionAffinityAttrs `json:"session_affinity_attributes,omitempty"`

	TTL uint `json:"ttl,omitempty"`

	SteeringPolicy string `json:"steering_policy,omitempty"`
	FallbackPool   string `json:"fallback_pool,omitempty"`

	DefaultPools []string            `json:"default_pools,omitempty"`
	PoPPools     map[string][]string `json:"pop_pools,omitempty"`
	RegionPools  map[string][]string `json:"region_pools,omitempty"`
	CountryPools map[string][]string `json:"country_pools,omitempty"`

	RandomSteering *RandomSteering `json:"random_steering,omitempty"`
}

// RandomSteering represents fields used to set pool weights on a load balancer
// with "random" steering policy.
type RandomSteering struct {
	DefaultWeight float64            `json:"default_weight,omitempty"`
	PoolWeights   map[string]float64 `json:"pool_weights,omitempty"`
}

// LoadBalancerRuleOverridesSessionAffinityAttrs mimics SessionAffinityAttributes without the
// DrainDuration field as that field can not be overwritten via rules.
type LoadBalancerRuleOverridesSessionAffinityAttrs struct {
	SameSite             string `json:"samesite,omitempty"`
	Secure               string `json:"secure,omitempty"`
	ZeroDowntimeFailover string `json:"zero_downtime_failover,omitempty"`
}

// SessionAffinityAttributes represents the fields used to set attributes in a load balancer session affinity cookie.
type SessionAffinityAttributes struct {
	SameSite             string `json:"samesite,omitempty"`
	Secure               string `json:"secure,omitempty"`
	DrainDuration        int    `json:"drain_duration,omitempty"`
	ZeroDowntimeFailover string `json:"zero_downtime_failover,omitempty"`
}

// LoadBalancerOriginHealth represents the health of the origin.
type LoadBalancerOriginHealth struct {
	Healthy       bool     `json:"healthy,omitempty"`
	RTT           Duration `json:"rtt,omitempty"`
	FailureReason string   `json:"failure_reason,omitempty"`
	ResponseCode  int      `json:"response_code,omitempty"`
}

// LoadBalancerPoolPopHealth represents the health of the pool for given PoP.
type LoadBalancerPoolPopHealth struct {
	Healthy bool                                  `json:"healthy,omitempty"`
	Origins []map[string]LoadBalancerOriginHealth `json:"origins,omitempty"`
}

// LoadBalancerPoolHealth represents the healthchecks from different PoPs for a pool.
type LoadBalancerPoolHealth struct {
	ID        string                               `json:"pool_id,omitempty"`
	PopHealth map[string]LoadBalancerPoolPopHealth `json:"pop_health,omitempty"`
}

// loadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type loadBalancerPoolResponse struct {
	Response
	Result LoadBalancerPool `json:"result"`
}

// loadBalancerPoolListResponse represents the response from the List Pools endpoint.
type loadBalancerPoolListResponse struct {
	Response
	Result     []LoadBalancerPool `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// loadBalancerMonitorResponse represents the response from the load balancer monitor endpoints.
type loadBalancerMonitorResponse struct {
	Response
	Result LoadBalancerMonitor `json:"result"`
}

// loadBalancerMonitorListResponse represents the response from the List Monitors endpoint.
type loadBalancerMonitorListResponse struct {
	Response
	Result     []LoadBalancerMonitor `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// loadBalancerResponse represents the response from the load balancer endpoints.
type loadBalancerResponse struct {
	Response
	Result LoadBalancer `json:"result"`
}

// loadBalancerListResponse represents the response from the List Load Balancers endpoint.
type loadBalancerListResponse struct {
	Response
	Result     []LoadBalancer `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}

// loadBalancerPoolHealthResponse represents the response from the Pool Health Details endpoint.
type loadBalancerPoolHealthResponse struct {
	Response
	Result LoadBalancerPoolHealth `json:"result"`
}

// CreateLoadBalancerPool creates a new load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-create-pool
func (api *API) CreateLoadBalancerPool(ctx context.Context, pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, pool)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-list-pools
func (api *API) ListLoadBalancerPools(ctx context.Context) ([]LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerPoolListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// LoadBalancerPoolDetails returns the details for a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-details
func (api *API) LoadBalancerPoolDetails(ctx context.Context, poolID string) (LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s", api.userBaseURL("/user"), poolID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteLoadBalancerPool disables and deletes a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-delete-pool
func (api *API) DeleteLoadBalancerPool(ctx context.Context, poolID string) error {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s", api.userBaseURL("/user"), poolID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// ModifyLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-update-pool
func (api *API) ModifyLoadBalancerPool(ctx context.Context, pool LoadBalancerPool) (LoadBalancerPool, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s", api.userBaseURL("/user"), pool.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, pool)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateLoadBalancerMonitor creates a new load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-create-monitor
func (api *API) CreateLoadBalancerMonitor(ctx context.Context, monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListLoadBalancerMonitors lists load balancer monitors connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-list-monitors
func (api *API) ListLoadBalancerMonitors(ctx context.Context) ([]LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerMonitorListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// LoadBalancerMonitorDetails returns the details for a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-monitor-details
func (api *API) LoadBalancerMonitorDetails(ctx context.Context, monitorID string) (LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors/%s", api.userBaseURL("/user"), monitorID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteLoadBalancerMonitor disables and deletes a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-delete-monitor
func (api *API) DeleteLoadBalancerMonitor(ctx context.Context, monitorID string) error {
	uri := fmt.Sprintf("%s/load_balancers/monitors/%s", api.userBaseURL("/user"), monitorID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// ModifyLoadBalancerMonitor modifies a configured load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-update-monitor
func (api *API) ModifyLoadBalancerMonitor(ctx context.Context, monitor LoadBalancerMonitor) (LoadBalancerMonitor, error) {
	uri := fmt.Sprintf("%s/load_balancers/monitors/%s", api.userBaseURL("/user"), monitor.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, monitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateLoadBalancer creates a new load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-create-load-balancer
func (api *API) CreateLoadBalancer(ctx context.Context, zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, lb)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListLoadBalancers lists load balancers configured on a zone.
//
// API reference: https://api.cloudflare.com/#load-balancers-list-load-balancers
func (api *API) ListLoadBalancers(ctx context.Context, zoneID string) ([]LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// LoadBalancerDetails returns the details for a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-load-balancer-details
func (api *API) LoadBalancerDetails(ctx context.Context, zoneID, lbID string) (LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", zoneID, lbID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteLoadBalancer disables and deletes a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-delete-load-balancer
func (api *API) DeleteLoadBalancer(ctx context.Context, zoneID, lbID string) error {
	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", zoneID, lbID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// ModifyLoadBalancer modifies a configured load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-update-load-balancer
func (api *API) ModifyLoadBalancer(ctx context.Context, zoneID string, lb LoadBalancer) (LoadBalancer, error) {
	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", zoneID, lb.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, lb)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// PoolHealthDetails fetches the latest healtcheck details for a single pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-health-details
func (api *API) PoolHealthDetails(ctx context.Context, poolID string) (LoadBalancerPoolHealth, error) {
	uri := fmt.Sprintf("%s/load_balancers/pools/%s/health", api.userBaseURL("/user"), poolID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return LoadBalancerPoolHealth{}, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return LoadBalancerPoolHealth{}, errors.Wrap(err, errMakeRequestError)
=======
		return LoadBalancerPoolHealth{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
	}
	var r loadBalancerPoolHealthResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPoolHealth{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"encoding/json"
=======
	"context"
	"errors"
	"fmt"
	"net/http"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"time"

	"github.com/goccy/go-json"
)

// LoadBalancerPool represents a load balancer pool's properties.
type LoadBalancerPool struct {
	ID                string                      `json:"id,omitempty"`
	CreatedOn         *time.Time                  `json:"created_on,omitempty"`
	ModifiedOn        *time.Time                  `json:"modified_on,omitempty"`
	Description       string                      `json:"description"`
	Name              string                      `json:"name"`
	Enabled           bool                        `json:"enabled"`
	MinimumOrigins    *int                        `json:"minimum_origins,omitempty"`
	Monitor           string                      `json:"monitor,omitempty"`
	Origins           []LoadBalancerOrigin        `json:"origins"`
	NotificationEmail string                      `json:"notification_email,omitempty"`
	Latitude          *float32                    `json:"latitude,omitempty"`
	Longitude         *float32                    `json:"longitude,omitempty"`
	LoadShedding      *LoadBalancerLoadShedding   `json:"load_shedding,omitempty"`
	OriginSteering    *LoadBalancerOriginSteering `json:"origin_steering,omitempty"`
	Healthy           *bool                       `json:"healthy,omitempty"`

	// CheckRegions defines the geographic region(s) from where to run health-checks from - e.g. "WNAM", "WEU", "SAF", "SAM".
	// Providing a null/empty value means "all regions", which may not be available to all plan types.
	CheckRegions []string `json:"check_regions"`
}

// LoadBalancerOrigin represents a Load Balancer origin's properties.
type LoadBalancerOrigin struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Enabled bool   `json:"enabled"`
	// Weight of this origin relative to other origins in the pool.
	// Based on the configured weight the total traffic is distributed
	// among origins within the pool.
	//
	// When LoadBalancerOriginSteering.Policy="least_outstanding_requests", this
	// weight is used to scale the origin's outstanding requests.
	// When LoadBalancerOriginSteering.Policy="least_connections", this
	// weight is used to scale the origin's open connections.
	Weight float64             `json:"weight"`
	Header map[string][]string `json:"header"`
	// The virtual network subnet ID the origin belongs in.
	// Virtual network must also belong to the account.
	VirtualNetworkID string `json:"virtual_network_id,omitempty"`
}

// LoadBalancerOriginSteering controls origin selection for new sessions and traffic without session affinity.
type LoadBalancerOriginSteering struct {
	// Policy determines the type of origin steering policy to use.
	// It defaults to "random" (weighted) when empty or unspecified.
	//
	// "random": Select an origin randomly.
	//
	// "hash": Select an origin by computing a hash over the CF-Connecting-IP address.
	//
	// "least_outstanding_requests": Select an origin by taking into consideration origin weights,
	// as well as each origin's number of outstanding requests. Origins with more pending requests
	// are weighted proportionately less relative to others.
	//
	// "least_connections": Select an origin by taking into consideration origin weights,
	// as well as each origin's number of open connections. Origins with more open connections
	// are weighted proportionately less relative to others. Supported for HTTP/1 and HTTP/2 connections.
	Policy string `json:"policy,omitempty"`
}

// LoadBalancerMonitor represents a load balancer monitor's properties.
type LoadBalancerMonitor struct {
	ID              string              `json:"id,omitempty"`
	CreatedOn       *time.Time          `json:"created_on,omitempty"`
	ModifiedOn      *time.Time          `json:"modified_on,omitempty"`
	Type            string              `json:"type"`
	Description     string              `json:"description"`
	Method          string              `json:"method"`
	Path            string              `json:"path"`
	Header          map[string][]string `json:"header"`
	Timeout         int                 `json:"timeout"`
	Retries         int                 `json:"retries"`
	Interval        int                 `json:"interval"`
	ConsecutiveUp   int                 `json:"consecutive_up"`
	ConsecutiveDown int                 `json:"consecutive_down"`
	Port            uint16              `json:"port,omitempty"`
	ExpectedBody    string              `json:"expected_body"`
	ExpectedCodes   string              `json:"expected_codes"`
	FollowRedirects bool                `json:"follow_redirects"`
	AllowInsecure   bool                `json:"allow_insecure"`
	ProbeZone       string              `json:"probe_zone"`
}

// LoadBalancer represents a load balancer's properties.
type LoadBalancer struct {
	ID                        string                     `json:"id,omitempty"`
	CreatedOn                 *time.Time                 `json:"created_on,omitempty"`
	ModifiedOn                *time.Time                 `json:"modified_on,omitempty"`
	Description               string                     `json:"description"`
	Name                      string                     `json:"name"`
	TTL                       int                        `json:"ttl,omitempty"`
	FallbackPool              string                     `json:"fallback_pool"`
	DefaultPools              []string                   `json:"default_pools"`
	RegionPools               map[string][]string        `json:"region_pools"`
	PopPools                  map[string][]string        `json:"pop_pools"`
	CountryPools              map[string][]string        `json:"country_pools"`
	Proxied                   bool                       `json:"proxied"`
	Enabled                   *bool                      `json:"enabled,omitempty"`
	Persistence               string                     `json:"session_affinity,omitempty"`
	PersistenceTTL            int                        `json:"session_affinity_ttl,omitempty"`
	SessionAffinityAttributes *SessionAffinityAttributes `json:"session_affinity_attributes,omitempty"`
	Rules                     []*LoadBalancerRule        `json:"rules,omitempty"`
	RandomSteering            *RandomSteering            `json:"random_steering,omitempty"`
	AdaptiveRouting           *AdaptiveRouting           `json:"adaptive_routing,omitempty"`
	LocationStrategy          *LocationStrategy          `json:"location_strategy,omitempty"`

	// SteeringPolicy controls pool selection logic.
	//
	// "off": Select pools in DefaultPools order.
	//
	// "geo": Select pools based on RegionPools/PopPools/CountryPools.
	// For non-proxied requests, the country for CountryPools is determined by LocationStrategy.
	//
	// "dynamic_latency": Select pools based on RTT (requires health checks).
	//
	// "random": Selects pools in a random order.
	//
	// "proximity": Use the pools' latitude and longitude to select the closest pool using
	// the Cloudflare PoP location for proxied requests or the location determined by
	// LocationStrategy for non-proxied requests.
	//
	// "least_outstanding_requests": Select a pool by taking into consideration
	// RandomSteering weights, as well as each pool's number of outstanding requests.
	// Pools with more pending requests are weighted proportionately less relative to others.
	//
	// "least_connections": Select a pool by taking into consideration
	// RandomSteering weights, as well as each pool's number of open connections.
	// Pools with more open connections are weighted proportionately less relative to others.
	// Supported for HTTP/1 and HTTP/2 connections.
	//
	// "": Maps to "geo" if RegionPools or PopPools or CountryPools have entries otherwise "off".
	SteeringPolicy string `json:"steering_policy,omitempty"`
}

// LoadBalancerLoadShedding contains the settings for controlling load shedding.
type LoadBalancerLoadShedding struct {
	DefaultPercent float32 `json:"default_percent,omitempty"`
	DefaultPolicy  string  `json:"default_policy,omitempty"`
	SessionPercent float32 `json:"session_percent,omitempty"`
	SessionPolicy  string  `json:"session_policy,omitempty"`
}

// LoadBalancerRule represents a single rule entry for a Load Balancer. Each rules
// is run one after the other in priority order. Disabled rules are skipped.
type LoadBalancerRule struct {
	Overrides LoadBalancerRuleOverrides `json:"overrides"`

	// Name is required but is only used for human readability
	Name string `json:"name"`

	Condition string `json:"condition"`

	// Priority controls the order of rule execution the lowest value will be invoked first
	Priority int `json:"priority"`

	// FixedResponse if set and the condition is true we will not run
	// routing logic but rather directly respond with the provided fields.
	// FixedResponse implies terminates.
	FixedResponse *LoadBalancerFixedResponseData `json:"fixed_response,omitempty"`

	Disabled bool `json:"disabled"`

	// Terminates flag this rule as 'terminating'. No further rules will
	// be executed after this one.
	Terminates bool `json:"terminates,omitempty"`
}

// LoadBalancerFixedResponseData contains all the data needed to generate
// a fixed response from a Load Balancer. This behavior can be enabled via Rules.
type LoadBalancerFixedResponseData struct {
	// MessageBody data to write into the http body
	MessageBody string `json:"message_body,omitempty"`
	// StatusCode the http status code to response with
	StatusCode int `json:"status_code,omitempty"`
	// ContentType value of the http 'content-type' header
	ContentType string `json:"content_type,omitempty"`
	// Location value of the http 'location' header
	Location string `json:"location,omitempty"`
}

// LoadBalancerRuleOverrides are the set of field overridable by the rules system.
type LoadBalancerRuleOverrides struct {
	// session affinity
	Persistence    string `json:"session_affinity,omitempty"`
	PersistenceTTL *uint  `json:"session_affinity_ttl,omitempty"`

	SessionAffinityAttrs *LoadBalancerRuleOverridesSessionAffinityAttrs `json:"session_affinity_attributes,omitempty"`

	TTL uint `json:"ttl,omitempty"`

	SteeringPolicy string `json:"steering_policy,omitempty"`
	FallbackPool   string `json:"fallback_pool,omitempty"`

	DefaultPools []string            `json:"default_pools,omitempty"`
	PoPPools     map[string][]string `json:"pop_pools,omitempty"`
	RegionPools  map[string][]string `json:"region_pools,omitempty"`
	CountryPools map[string][]string `json:"country_pools,omitempty"`

	RandomSteering   *RandomSteering   `json:"random_steering,omitempty"`
	AdaptiveRouting  *AdaptiveRouting  `json:"adaptive_routing,omitempty"`
	LocationStrategy *LocationStrategy `json:"location_strategy,omitempty"`
}

// RandomSteering configures pool weights.
//
// SteeringPolicy="random": A random pool is selected with probability
// proportional to pool weights.
//
// SteeringPolicy="least_outstanding_requests": Use pool weights to
// scale each pool's outstanding requests.
//
// SteeringPolicy="least_connections": Use pool weights to
// scale each pool's open connections.
type RandomSteering struct {
	DefaultWeight float64            `json:"default_weight,omitempty"`
	PoolWeights   map[string]float64 `json:"pool_weights,omitempty"`
}

// AdaptiveRouting controls features that modify the routing of requests
// to pools and origins in response to dynamic conditions, such as during
// the interval between active health monitoring requests.
// For example, zero-downtime failover occurs immediately when an origin
// becomes unavailable due to HTTP 521, 522, or 523 response codes.
// If there is another healthy origin in the same pool, the request is
// retried once against this alternate origin.
type AdaptiveRouting struct {
	// FailoverAcrossPools extends zero-downtime failover of requests to healthy origins
	// from alternate pools, when no healthy alternate exists in the same pool, according
	// to the failover order defined by traffic and origin steering.
	// When set false (the default) zero-downtime failover will only occur between origins
	// within the same pool. See SessionAffinityAttributes for control over when sessions
	// are broken or reassigned.
	FailoverAcrossPools *bool `json:"failover_across_pools,omitempty"`
}

// LocationStrategy controls location-based steering for non-proxied requests.
// See SteeringPolicy to learn how steering is affected.
type LocationStrategy struct {
	// PreferECS determines whether the EDNS Client Subnet (ECS) GeoIP should
	// be preferred as the authoritative location.
	//
	// "always": Always prefer ECS.
	//
	// "never": Never prefer ECS.
	//
	// "proximity": (default) Prefer ECS only when SteeringPolicy="proximity".
	//
	// "geo": Prefer ECS only when SteeringPolicy="geo".
	PreferECS string `json:"prefer_ecs,omitempty"`
	// Mode determines the authoritative location when ECS is not preferred,
	// does not exist in the request, or its GeoIP lookup is unsuccessful.
	//
	// "pop": (default) Use the Cloudflare PoP location.
	//
	// "resolver_ip": Use the DNS resolver GeoIP location.
	// If the GeoIP lookup is unsuccessful, use the Cloudflare PoP location.
	Mode string `json:"mode,omitempty"`
}

// LoadBalancerRuleOverridesSessionAffinityAttrs mimics SessionAffinityAttributes without the
// DrainDuration field as that field can not be overwritten via rules.
type LoadBalancerRuleOverridesSessionAffinityAttrs struct {
	SameSite             string   `json:"samesite,omitempty"`
	Secure               string   `json:"secure,omitempty"`
	ZeroDowntimeFailover string   `json:"zero_downtime_failover,omitempty"`
	Headers              []string `json:"headers,omitempty"`
	RequireAllHeaders    *bool    `json:"require_all_headers,omitempty"`
}

// SessionAffinityAttributes represents additional configuration options for session affinity.
type SessionAffinityAttributes struct {
	SameSite             string   `json:"samesite,omitempty"`
	Secure               string   `json:"secure,omitempty"`
	DrainDuration        int      `json:"drain_duration,omitempty"`
	ZeroDowntimeFailover string   `json:"zero_downtime_failover,omitempty"`
	Headers              []string `json:"headers,omitempty"`
	RequireAllHeaders    bool     `json:"require_all_headers,omitempty"`
}

// LoadBalancerOriginHealth represents the health of the origin.
type LoadBalancerOriginHealth struct {
	Healthy       bool     `json:"healthy,omitempty"`
	RTT           Duration `json:"rtt,omitempty"`
	FailureReason string   `json:"failure_reason,omitempty"`
	ResponseCode  int      `json:"response_code,omitempty"`
}

// LoadBalancerPoolPopHealth represents the health of the pool for given PoP.
type LoadBalancerPoolPopHealth struct {
	Healthy bool                                  `json:"healthy,omitempty"`
	Origins []map[string]LoadBalancerOriginHealth `json:"origins,omitempty"`
}

// LoadBalancerPoolHealth represents the healthchecks from different PoPs for a pool.
type LoadBalancerPoolHealth struct {
	ID        string                               `json:"pool_id,omitempty"`
	PopHealth map[string]LoadBalancerPoolPopHealth `json:"pop_health,omitempty"`
}

// loadBalancerPoolResponse represents the response from the load balancer pool endpoints.
type loadBalancerPoolResponse struct {
	Response
	Result LoadBalancerPool `json:"result"`
}

// loadBalancerPoolListResponse represents the response from the List Pools endpoint.
type loadBalancerPoolListResponse struct {
	Response
	Result     []LoadBalancerPool `json:"result"`
	ResultInfo ResultInfo         `json:"result_info"`
}

// loadBalancerMonitorResponse represents the response from the load balancer monitor endpoints.
type loadBalancerMonitorResponse struct {
	Response
	Result LoadBalancerMonitor `json:"result"`
}

// loadBalancerMonitorListResponse represents the response from the List Monitors endpoint.
type loadBalancerMonitorListResponse struct {
	Response
	Result     []LoadBalancerMonitor `json:"result"`
	ResultInfo ResultInfo            `json:"result_info"`
}

// loadBalancerResponse represents the response from the load balancer endpoints.
type loadBalancerResponse struct {
	Response
	Result LoadBalancer `json:"result"`
}

// loadBalancerListResponse represents the response from the List Load Balancers endpoint.
type loadBalancerListResponse struct {
	Response
	Result     []LoadBalancer `json:"result"`
	ResultInfo ResultInfo     `json:"result_info"`
}

// loadBalancerPoolHealthResponse represents the response from the Pool Health Details endpoint.
type loadBalancerPoolHealthResponse struct {
	Response
	Result LoadBalancerPoolHealth `json:"result"`
}

type CreateLoadBalancerPoolParams struct {
	LoadBalancerPool LoadBalancerPool
}

type ListLoadBalancerPoolParams struct {
	PaginationOptions
}

type UpdateLoadBalancerPoolParams struct {
	LoadBalancer LoadBalancerPool
}

type CreateLoadBalancerMonitorParams struct {
	LoadBalancerMonitor LoadBalancerMonitor
}

type ListLoadBalancerMonitorParams struct {
	PaginationOptions
}

type UpdateLoadBalancerMonitorParams struct {
	LoadBalancerMonitor LoadBalancerMonitor
}

type CreateLoadBalancerParams struct {
	LoadBalancer LoadBalancer
}

type ListLoadBalancerParams struct {
	PaginationOptions
}

type UpdateLoadBalancerParams struct {
	LoadBalancer LoadBalancer
}

var (
	ErrMissingPoolID         = errors.New("missing required pool ID")
	ErrMissingMonitorID      = errors.New("missing required monitor ID")
	ErrMissingLoadBalancerID = errors.New("missing required load balancer ID")
)

// CreateLoadBalancerPool creates a new load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-create-pool
func (api *API) CreateLoadBalancerPool(ctx context.Context, rc *ResourceContainer, params CreateLoadBalancerPoolParams) (LoadBalancerPool, error) {
	if rc.Level == ZoneRouteLevel {
		return LoadBalancerPool{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = "/user/load_balancers/pools"
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools", rc.Identifier)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.LoadBalancerPool)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListLoadBalancerPools lists load balancer pools connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-list-pools
func (api *API) ListLoadBalancerPools(ctx context.Context, rc *ResourceContainer, params ListLoadBalancerPoolParams) ([]LoadBalancerPool, error) {
	if rc.Level == ZoneRouteLevel {
		return []LoadBalancerPool{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = "/user/load_balancers/pools"
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools", rc.Identifier)
	}

	uri = buildURI(uri, params.PaginationOptions)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerPoolListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetLoadBalancerPool returns the details for a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-details
func (api *API) GetLoadBalancerPool(ctx context.Context, rc *ResourceContainer, poolID string) (LoadBalancerPool, error) {
	if rc.Level == ZoneRouteLevel {
		return LoadBalancerPool{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if poolID == "" {
		return LoadBalancerPool{}, ErrMissingPoolID
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = fmt.Sprintf("/user/load_balancers/pools/%s", poolID)
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", rc.Identifier, poolID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteLoadBalancerPool disables and deletes a load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-delete-pool
func (api *API) DeleteLoadBalancerPool(ctx context.Context, rc *ResourceContainer, poolID string) error {
	if rc.Level == ZoneRouteLevel {
		return fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if poolID == "" {
		return ErrMissingPoolID
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = fmt.Sprintf("/user/load_balancers/pools/%s", poolID)
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", rc.Identifier, poolID)
	}

	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}

	return nil
}

// UpdateLoadBalancerPool modifies a configured load balancer pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-update-pool
func (api *API) UpdateLoadBalancerPool(ctx context.Context, rc *ResourceContainer, params UpdateLoadBalancerPoolParams) (LoadBalancerPool, error) {
	if rc.Level == ZoneRouteLevel {
		return LoadBalancerPool{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if params.LoadBalancer.ID == "" {
		return LoadBalancerPool{}, ErrMissingPoolID
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = fmt.Sprintf("/user/load_balancers/pools/%s", params.LoadBalancer.ID)
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s", rc.Identifier, params.LoadBalancer.ID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.LoadBalancer)
	if err != nil {
		return LoadBalancerPool{}, err
	}
	var r loadBalancerPoolResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerPool{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateLoadBalancerMonitor creates a new load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-create-monitor
func (api *API) CreateLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, params CreateLoadBalancerMonitorParams) (LoadBalancerMonitor, error) {
	if rc.Level == ZoneRouteLevel {
		return LoadBalancerMonitor{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = "/user/load_balancers/monitors"
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors", rc.Identifier)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.LoadBalancerMonitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListLoadBalancerMonitors lists load balancer monitors connected to an account.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-list-monitors
func (api *API) ListLoadBalancerMonitors(ctx context.Context, rc *ResourceContainer, params ListLoadBalancerMonitorParams) ([]LoadBalancerMonitor, error) {
	if rc.Level == ZoneRouteLevel {
		return []LoadBalancerMonitor{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = "/user/load_balancers/monitors"
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors", rc.Identifier)
	}

	uri = buildURI(uri, params.PaginationOptions)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerMonitorListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetLoadBalancerMonitor returns the details for a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-monitor-details
func (api *API) GetLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, monitorID string) (LoadBalancerMonitor, error) {
	if rc.Level == ZoneRouteLevel {
		return LoadBalancerMonitor{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if monitorID == "" {
		return LoadBalancerMonitor{}, ErrMissingMonitorID
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = fmt.Sprintf("/user/load_balancers/monitors/%s", monitorID)
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", rc.Identifier, monitorID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteLoadBalancerMonitor disables and deletes a load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-delete-monitor
func (api *API) DeleteLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, monitorID string) error {
	if rc.Level == ZoneRouteLevel {
		return fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if monitorID == "" {
		return ErrMissingMonitorID
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = fmt.Sprintf("/user/load_balancers/monitors/%s", monitorID)
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", rc.Identifier, monitorID)
	}

	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// UpdateLoadBalancerMonitor modifies a configured load balancer monitor.
//
// API reference: https://api.cloudflare.com/#load-balancer-monitors-update-monitor
func (api *API) UpdateLoadBalancerMonitor(ctx context.Context, rc *ResourceContainer, params UpdateLoadBalancerMonitorParams) (LoadBalancerMonitor, error) {
	if rc.Level == ZoneRouteLevel {
		return LoadBalancerMonitor{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if params.LoadBalancerMonitor.ID == "" {
		return LoadBalancerMonitor{}, ErrMissingMonitorID
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = fmt.Sprintf("/user/load_balancers/monitors/%s", params.LoadBalancerMonitor.ID)
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/monitors/%s", rc.Identifier, params.LoadBalancerMonitor.ID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.LoadBalancerMonitor)
	if err != nil {
		return LoadBalancerMonitor{}, err
	}
	var r loadBalancerMonitorResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancerMonitor{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// CreateLoadBalancer creates a new load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-create-load-balancer
func (api *API) CreateLoadBalancer(ctx context.Context, rc *ResourceContainer, params CreateLoadBalancerParams) (LoadBalancer, error) {
	if rc.Level != ZoneRouteLevel {
		return LoadBalancer{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf("/zones/%s/load_balancers", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.LoadBalancer)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListLoadBalancers lists load balancers configured on a zone.
//
// API reference: https://api.cloudflare.com/#load-balancers-list-load-balancers
func (api *API) ListLoadBalancers(ctx context.Context, rc *ResourceContainer, params ListLoadBalancerParams) ([]LoadBalancer, error) {
	if rc.Level != ZoneRouteLevel {
		return []LoadBalancer{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/load_balancers", rc.Identifier), params.PaginationOptions)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r loadBalancerListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetLoadBalancer returns the details for a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-load-balancer-details
func (api *API) GetLoadBalancer(ctx context.Context, rc *ResourceContainer, loadbalancerID string) (LoadBalancer, error) {
	if rc.Level != ZoneRouteLevel {
		return LoadBalancer{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if loadbalancerID == "" {
		return LoadBalancer{}, ErrMissingLoadBalancerID
	}

	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", rc.Identifier, loadbalancerID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteLoadBalancer disables and deletes a load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-delete-load-balancer
func (api *API) DeleteLoadBalancer(ctx context.Context, rc *ResourceContainer, loadbalancerID string) error {
	if rc.Level != ZoneRouteLevel {
		return fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if loadbalancerID == "" {
		return ErrMissingLoadBalancerID
	}

	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", rc.Identifier, loadbalancerID)

	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
	}
	return nil
}

// UpdateLoadBalancer modifies a configured load balancer.
//
// API reference: https://api.cloudflare.com/#load-balancers-update-load-balancer
func (api *API) UpdateLoadBalancer(ctx context.Context, rc *ResourceContainer, params UpdateLoadBalancerParams) (LoadBalancer, error) {
	if rc.Level != ZoneRouteLevel {
		return LoadBalancer{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if params.LoadBalancer.ID == "" {
		return LoadBalancer{}, ErrMissingLoadBalancerID
	}

	uri := fmt.Sprintf("/zones/%s/load_balancers/%s", rc.Identifier, params.LoadBalancer.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.LoadBalancer)
	if err != nil {
		return LoadBalancer{}, err
	}
	var r loadBalancerResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return LoadBalancer{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetLoadBalancerPoolHealth fetches the latest healtcheck details for a single
// pool.
//
// API reference: https://api.cloudflare.com/#load-balancer-pools-pool-health-details
func (api *API) GetLoadBalancerPoolHealth(ctx context.Context, rc *ResourceContainer, poolID string) (LoadBalancerPoolHealth, error) {
	if rc.Level == ZoneRouteLevel {
		return LoadBalancerPoolHealth{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if poolID == "" {
		return LoadBalancerPoolHealth{}, ErrMissingPoolID
	}

	var uri string
	if rc.Level == UserRouteLevel {
		uri = fmt.Sprintf("/user/load_balancers/pools/%s/health", poolID)
	} else {
		uri = fmt.Sprintf("/accounts/%s/load_balancers/pools/%s/health", rc.Identifier, poolID)
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return LoadBalancerPoolHealth{}, err
	}
	var r loadBalancerPoolHealthResponse
	if err := json.Unmarshal(res, &r); err != nil {
<<<<<<< HEAD
		return LoadBalancerPoolHealth{}, errors.Wrap(err, errUnmarshalError)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		return LoadBalancerPoolHealth{}, errors.Wrap(err, errUnmarshalError)
=======
		return LoadBalancerPoolHealth{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
	return r.Result, nil
}
