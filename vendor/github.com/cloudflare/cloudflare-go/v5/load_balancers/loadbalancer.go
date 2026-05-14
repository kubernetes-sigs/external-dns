// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// LoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLoadBalancerService] method instead.
type LoadBalancerService struct {
	Options  []option.RequestOption
	Monitors *MonitorService
	Pools    *PoolService
	Previews *PreviewService
	Regions  *RegionService
	Searches *SearchService
}

// NewLoadBalancerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLoadBalancerService(opts ...option.RequestOption) (r *LoadBalancerService) {
	r = &LoadBalancerService{}
	r.Options = opts
	r.Monitors = NewMonitorService(opts...)
	r.Pools = NewPoolService(opts...)
	r.Previews = NewPreviewService(opts...)
	r.Regions = NewRegionService(opts...)
	r.Searches = NewSearchService(opts...)
	return
}

// Create a new load balancer.
func (r *LoadBalancerService) New(ctx context.Context, params LoadBalancerNewParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	var env LoadBalancerNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a configured load balancer.
func (r *LoadBalancerService) Update(ctx context.Context, loadBalancerID string, params LoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	var env LoadBalancerUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", params.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List configured load balancers.
func (r *LoadBalancerService) List(ctx context.Context, query LoadBalancerListParams, opts ...option.RequestOption) (res *pagination.SinglePage[LoadBalancer], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List configured load balancers.
func (r *LoadBalancerService) ListAutoPaging(ctx context.Context, query LoadBalancerListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LoadBalancer] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a configured load balancer.
func (r *LoadBalancerService) Delete(ctx context.Context, loadBalancerID string, body LoadBalancerDeleteParams, opts ...option.RequestOption) (res *LoadBalancerDeleteResponse, err error) {
	var env LoadBalancerDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", body.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Apply changes to an existing load balancer, overwriting the supplied properties.
func (r *LoadBalancerService) Edit(ctx context.Context, loadBalancerID string, params LoadBalancerEditParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	var env LoadBalancerEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", params.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single configured load balancer.
func (r *LoadBalancerService) Get(ctx context.Context, loadBalancerID string, query LoadBalancerGetParams, opts ...option.RequestOption) (res *LoadBalancer, err error) {
	var env LoadBalancerGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if loadBalancerID == "" {
		err = errors.New("missing required load_balancer_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/load_balancers/%s", query.ZoneID, loadBalancerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type AdaptiveRouting struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools bool                `json:"failover_across_pools"`
	JSON                adaptiveRoutingJSON `json:"-"`
}

// adaptiveRoutingJSON contains the JSON metadata for the struct [AdaptiveRouting]
type adaptiveRoutingJSON struct {
	FailoverAcrossPools apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AdaptiveRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adaptiveRoutingJSON) RawJSON() string {
	return r.raw
}

// Controls features that modify the routing of requests to pools and origins in
// response to dynamic conditions, such as during the interval between active
// health monitoring requests. For example, zero-downtime failover occurs
// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
// response codes. If there is another healthy origin in the same pool, the request
// is retried once against this alternate origin.
type AdaptiveRoutingParam struct {
	// Extends zero-downtime failover of requests to healthy origins from alternate
	// pools, when no healthy alternate exists in the same pool, according to the
	// failover order defined by traffic and origin steering. When set false (the
	// default) zero-downtime failover will only occur between origins within the same
	// pool. See `session_affinity_attributes` for control over when sessions are
	// broken or reassigned.
	FailoverAcrossPools param.Field[bool] `json:"failover_across_pools"`
}

func (r AdaptiveRoutingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// WNAM: Western North America, ENAM: Eastern North America, WEU: Western Europe,
// EEU: Eastern Europe, NSAM: Northern South America, SSAM: Southern South America,
// OC: Oceania, ME: Middle East, NAF: North Africa, SAF: South Africa, SAS:
// Southern Asia, SEAS: South East Asia, NEAS: North East Asia, ALL_REGIONS: all
// regions (ENTERPRISE customers only).
type CheckRegion string

const (
	CheckRegionWnam       CheckRegion = "WNAM"
	CheckRegionEnam       CheckRegion = "ENAM"
	CheckRegionWeu        CheckRegion = "WEU"
	CheckRegionEeu        CheckRegion = "EEU"
	CheckRegionNsam       CheckRegion = "NSAM"
	CheckRegionSsam       CheckRegion = "SSAM"
	CheckRegionOc         CheckRegion = "OC"
	CheckRegionMe         CheckRegion = "ME"
	CheckRegionNaf        CheckRegion = "NAF"
	CheckRegionSaf        CheckRegion = "SAF"
	CheckRegionSas        CheckRegion = "SAS"
	CheckRegionSeas       CheckRegion = "SEAS"
	CheckRegionNeas       CheckRegion = "NEAS"
	CheckRegionAllRegions CheckRegion = "ALL_REGIONS"
)

func (r CheckRegion) IsKnown() bool {
	switch r {
	case CheckRegionWnam, CheckRegionEnam, CheckRegionWeu, CheckRegionEeu, CheckRegionNsam, CheckRegionSsam, CheckRegionOc, CheckRegionMe, CheckRegionNaf, CheckRegionSaf, CheckRegionSas, CheckRegionSeas, CheckRegionNeas, CheckRegionAllRegions:
		return true
	}
	return false
}

type DefaultPools = string

type DefaultPoolsParam = string

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type FilterOptions struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable bool `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy bool              `json:"healthy,nullable"`
	JSON    filterOptionsJSON `json:"-"`
}

// filterOptionsJSON contains the JSON metadata for the struct [FilterOptions]
type filterOptionsJSON struct {
	Disable     apijson.Field
	Healthy     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterOptionsJSON) RawJSON() string {
	return r.raw
}

// Filter options for a particular resource type (pool or origin). Use null to
// reset.
type FilterOptionsParam struct {
	// If set true, disable notifications for this type of resource (pool or origin).
	Disable param.Field[bool] `json:"disable"`
	// If present, send notifications only for this health status (e.g. false for only
	// DOWN events). Use null to reset (all events).
	Healthy param.Field[bool] `json:"healthy"`
}

func (r FilterOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type Header struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host []Host     `json:"Host"`
	JSON headerJSON `json:"-"`
}

// headerJSON contains the JSON metadata for the struct [Header]
type headerJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Header) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r headerJSON) RawJSON() string {
	return r.raw
}

// The request header is used to pass additional information with an HTTP request.
// Currently supported header is 'Host'.
type HeaderParam struct {
	// The 'Host' header allows to override the hostname set in the HTTP request.
	// Current support is 1 'Host' header override per origin.
	Host param.Field[[]HostParam] `json:"Host"`
}

func (r HeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Host = string

type HostParam = string

type LoadBalancer struct {
	ID string `json:"id"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting AdaptiveRouting `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools map[string][]string `json:"country_pools"`
	CreatedOn    string              `json:"created_on"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools []DefaultPools `json:"default_pools"`
	// Object description.
	Description string `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled bool `json:"enabled"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool string `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy LocationStrategy `json:"location_strategy"`
	ModifiedOn       string           `json:"modified_on"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name string `json:"name"`
	// List of networks where Load Balancer or Pool is enabled.
	Networks []string `json:"networks"`
	// Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	POPPools map[string][]string `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied bool `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering RandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools map[string][]string `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules []Rules `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are: - `"cookie"`: On the first
	// request to a proxied load balancer, a cookie is generated, encoding information
	// of which origin the request will be forwarded to. Subsequent requests, by the
	// same client to the same load balancer, will be sent to the origin server the
	// cookie encodes, for the duration of the cookie and as long as the origin server
	// remains healthy. If the cookie has expired or the origin server is unhealthy,
	// then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the
	// same as `"cookie"` except the initial origin selection is stable and based on
	// the client's ip address. - `"header"`: On the first request to a proxied load
	// balancer, a session key based on the configured HTTP headers (see
	// `session_affinity_attributes.headers`) is generated, encoding the request
	// headers used for storing in the load balancer session state which origin the
	// request will be forwarded to. Subsequent requests to the load balancer with the
	// same headers will be sent to the same origin server, for the duration of the
	// session and as long as the origin server remains healthy. If the session has
	// been idle for the duration of `session_affinity_ttl` seconds or the origin
	// server is unhealthy, then a new origin server is calculated and used. See
	// `headers` in `session_affinity_attributes` for additional required
	// configuration.
	SessionAffinity SessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes SessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are: -
	// `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless
	// explicitly set. The accepted range of values is between [1800, 604800]. -
	// `"header"`: The current default of 1800 seconds will be used unless explicitly
	// set. The accepted range of values is between [30, 3600]. Note: With session
	// affinity by header, sessions only expire after they haven't been used for the
	// number of seconds specified.
	SessionAffinityTTL float64 `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy SteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL      float64          `json:"ttl"`
	ZoneName string           `json:"zone_name"`
	JSON     loadBalancerJSON `json:"-"`
}

// loadBalancerJSON contains the JSON metadata for the struct [LoadBalancer]
type loadBalancerJSON struct {
	ID                        apijson.Field
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	CreatedOn                 apijson.Field
	DefaultPools              apijson.Field
	Description               apijson.Field
	Enabled                   apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	ModifiedOn                apijson.Field
	Name                      apijson.Field
	Networks                  apijson.Field
	POPPools                  apijson.Field
	Proxied                   apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	Rules                     apijson.Field
	SessionAffinity           apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTTL        apijson.Field
	SteeringPolicy            apijson.Field
	TTL                       apijson.Field
	ZoneName                  apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *LoadBalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerJSON) RawJSON() string {
	return r.raw
}

// Configures load shedding policies and percentages for the pool.
type LoadShedding struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent float64 `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy LoadSheddingDefaultPolicy `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent float64 `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy LoadSheddingSessionPolicy `json:"session_policy"`
	JSON          loadSheddingJSON          `json:"-"`
}

// loadSheddingJSON contains the JSON metadata for the struct [LoadShedding]
type loadSheddingJSON struct {
	DefaultPercent apijson.Field
	DefaultPolicy  apijson.Field
	SessionPercent apijson.Field
	SessionPolicy  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadShedding) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadSheddingJSON) RawJSON() string {
	return r.raw
}

// The default policy to use when load shedding. A random policy randomly sheds a
// given percent of requests. A hash policy computes a hash over the
// CF-Connecting-IP address and sheds all requests originating from a percent of
// IPs.
type LoadSheddingDefaultPolicy string

const (
	LoadSheddingDefaultPolicyRandom LoadSheddingDefaultPolicy = "random"
	LoadSheddingDefaultPolicyHash   LoadSheddingDefaultPolicy = "hash"
)

func (r LoadSheddingDefaultPolicy) IsKnown() bool {
	switch r {
	case LoadSheddingDefaultPolicyRandom, LoadSheddingDefaultPolicyHash:
		return true
	}
	return false
}

// Only the hash policy is supported for existing sessions (to avoid exponential
// decay).
type LoadSheddingSessionPolicy string

const (
	LoadSheddingSessionPolicyHash LoadSheddingSessionPolicy = "hash"
)

func (r LoadSheddingSessionPolicy) IsKnown() bool {
	switch r {
	case LoadSheddingSessionPolicyHash:
		return true
	}
	return false
}

// Configures load shedding policies and percentages for the pool.
type LoadSheddingParam struct {
	// The percent of traffic to shed from the pool, according to the default policy.
	// Applies to new sessions and traffic without session affinity.
	DefaultPercent param.Field[float64] `json:"default_percent"`
	// The default policy to use when load shedding. A random policy randomly sheds a
	// given percent of requests. A hash policy computes a hash over the
	// CF-Connecting-IP address and sheds all requests originating from a percent of
	// IPs.
	DefaultPolicy param.Field[LoadSheddingDefaultPolicy] `json:"default_policy"`
	// The percent of existing sessions to shed from the pool, according to the session
	// policy.
	SessionPercent param.Field[float64] `json:"session_percent"`
	// Only the hash policy is supported for existing sessions (to avoid exponential
	// decay).
	SessionPolicy param.Field[LoadSheddingSessionPolicy] `json:"session_policy"`
}

func (r LoadSheddingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LocationStrategy struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode LocationStrategyMode `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferECS LocationStrategyPreferECS `json:"prefer_ecs"`
	JSON      locationStrategyJSON      `json:"-"`
}

// locationStrategyJSON contains the JSON metadata for the struct
// [LocationStrategy]
type locationStrategyJSON struct {
	Mode        apijson.Field
	PreferECS   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationStrategy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r locationStrategyJSON) RawJSON() string {
	return r.raw
}

// Determines the authoritative location when ECS is not preferred, does not exist
// in the request, or its GeoIP lookup is unsuccessful.
//
//   - `"pop"`: Use the Cloudflare PoP location.
//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
//     unsuccessful, use the Cloudflare PoP location.
type LocationStrategyMode string

const (
	LocationStrategyModePOP        LocationStrategyMode = "pop"
	LocationStrategyModeResolverIP LocationStrategyMode = "resolver_ip"
)

func (r LocationStrategyMode) IsKnown() bool {
	switch r {
	case LocationStrategyModePOP, LocationStrategyModeResolverIP:
		return true
	}
	return false
}

// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
// authoritative location.
//
// - `"always"`: Always prefer ECS.
// - `"never"`: Never prefer ECS.
// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
type LocationStrategyPreferECS string

const (
	LocationStrategyPreferECSAlways    LocationStrategyPreferECS = "always"
	LocationStrategyPreferECSNever     LocationStrategyPreferECS = "never"
	LocationStrategyPreferECSProximity LocationStrategyPreferECS = "proximity"
	LocationStrategyPreferECSGeo       LocationStrategyPreferECS = "geo"
)

func (r LocationStrategyPreferECS) IsKnown() bool {
	switch r {
	case LocationStrategyPreferECSAlways, LocationStrategyPreferECSNever, LocationStrategyPreferECSProximity, LocationStrategyPreferECSGeo:
		return true
	}
	return false
}

// Controls location-based steering for non-proxied requests. See `steering_policy`
// to learn how steering is affected.
type LocationStrategyParam struct {
	// Determines the authoritative location when ECS is not preferred, does not exist
	// in the request, or its GeoIP lookup is unsuccessful.
	//
	//   - `"pop"`: Use the Cloudflare PoP location.
	//   - `"resolver_ip"`: Use the DNS resolver GeoIP location. If the GeoIP lookup is
	//     unsuccessful, use the Cloudflare PoP location.
	Mode param.Field[LocationStrategyMode] `json:"mode"`
	// Whether the EDNS Client Subnet (ECS) GeoIP should be preferred as the
	// authoritative location.
	//
	// - `"always"`: Always prefer ECS.
	// - `"never"`: Never prefer ECS.
	// - `"proximity"`: Prefer ECS only when `steering_policy="proximity"`.
	// - `"geo"`: Prefer ECS only when `steering_policy="geo"`.
	PreferECS param.Field[LocationStrategyPreferECS] `json:"prefer_ecs"`
}

func (r LocationStrategyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type NotificationFilter struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin FilterOptions `json:"origin,nullable"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool FilterOptions          `json:"pool,nullable"`
	JSON notificationFilterJSON `json:"-"`
}

// notificationFilterJSON contains the JSON metadata for the struct
// [NotificationFilter]
type notificationFilterJSON struct {
	Origin      apijson.Field
	Pool        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NotificationFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r notificationFilterJSON) RawJSON() string {
	return r.raw
}

// Filter pool and origin health notifications by resource type or health status.
// Use null to reset.
type NotificationFilterParam struct {
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Origin param.Field[FilterOptionsParam] `json:"origin"`
	// Filter options for a particular resource type (pool or origin). Use null to
	// reset.
	Pool param.Field[FilterOptionsParam] `json:"pool"`
}

func (r NotificationFilterParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Origin struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address string `json:"address"`
	// This field shows up only if the origin is disabled. This field is set with the
	// time the origin was disabled.
	DisabledAt time.Time `json:"disabled_at" format:"date-time"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled bool `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header Header `json:"header"`
	// A human-identifiable name for the origin.
	Name string `json:"name"`
	// The port for upstream connections. A value of 0 means the default port for the
	// protocol will be used.
	Port int64 `json:"port"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID string `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight float64    `json:"weight"`
	JSON   originJSON `json:"-"`
}

// originJSON contains the JSON metadata for the struct [Origin]
type originJSON struct {
	Address          apijson.Field
	DisabledAt       apijson.Field
	Enabled          apijson.Field
	Header           apijson.Field
	Name             apijson.Field
	Port             apijson.Field
	VirtualNetworkID apijson.Field
	Weight           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Origin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originJSON) RawJSON() string {
	return r.raw
}

type OriginParam struct {
	// The IP address (IPv4 or IPv6) of the origin, or its publicly addressable
	// hostname. Hostnames entered here should resolve directly to the origin, and not
	// be a hostname proxied by Cloudflare. To set an internal/reserved address,
	// virtual_network_id must also be set.
	Address param.Field[string] `json:"address"`
	// Whether to enable (the default) this origin within the pool. Disabled origins
	// will not receive traffic and are excluded from health checks. The origin will
	// only be disabled for the current pool.
	Enabled param.Field[bool] `json:"enabled"`
	// The request header is used to pass additional information with an HTTP request.
	// Currently supported header is 'Host'.
	Header param.Field[HeaderParam] `json:"header"`
	// A human-identifiable name for the origin.
	Name param.Field[string] `json:"name"`
	// The port for upstream connections. A value of 0 means the default port for the
	// protocol will be used.
	Port param.Field[int64] `json:"port"`
	// The virtual network subnet ID the origin belongs in. Virtual network must also
	// belong to the account.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id"`
	// The weight of this origin relative to other origins in the pool. Based on the
	// configured weight the total traffic is distributed among origins within the
	// pool.
	//
	//   - `origin_steering.policy="least_outstanding_requests"`: Use weight to scale the
	//     origin's outstanding requests.
	//   - `origin_steering.policy="least_connections"`: Use weight to scale the origin's
	//     open connections.
	Weight param.Field[float64] `json:"weight"`
}

func (r OriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type OriginSteering struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy OriginSteeringPolicy `json:"policy"`
	JSON   originSteeringJSON   `json:"-"`
}

// originSteeringJSON contains the JSON metadata for the struct [OriginSteering]
type originSteeringJSON struct {
	Policy      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originSteeringJSON) RawJSON() string {
	return r.raw
}

// The type of origin steering policy to use.
//
//   - `"random"`: Select an origin randomly.
//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
//     address.
//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
//     origin weights, as well as each origin's number of outstanding requests.
//     Origins with more pending requests are weighted proportionately less relative
//     to others.
//   - `"least_connections"`: Select an origin by taking into consideration origin
//     weights, as well as each origin's number of open connections. Origins with
//     more open connections are weighted proportionately less relative to others.
//     Supported for HTTP/1 and HTTP/2 connections.
type OriginSteeringPolicy string

const (
	OriginSteeringPolicyRandom                   OriginSteeringPolicy = "random"
	OriginSteeringPolicyHash                     OriginSteeringPolicy = "hash"
	OriginSteeringPolicyLeastOutstandingRequests OriginSteeringPolicy = "least_outstanding_requests"
	OriginSteeringPolicyLeastConnections         OriginSteeringPolicy = "least_connections"
)

func (r OriginSteeringPolicy) IsKnown() bool {
	switch r {
	case OriginSteeringPolicyRandom, OriginSteeringPolicyHash, OriginSteeringPolicyLeastOutstandingRequests, OriginSteeringPolicyLeastConnections:
		return true
	}
	return false
}

// Configures origin steering for the pool. Controls how origins are selected for
// new sessions and traffic without session affinity.
type OriginSteeringParam struct {
	// The type of origin steering policy to use.
	//
	//   - `"random"`: Select an origin randomly.
	//   - `"hash"`: Select an origin by computing a hash over the CF-Connecting-IP
	//     address.
	//   - `"least_outstanding_requests"`: Select an origin by taking into consideration
	//     origin weights, as well as each origin's number of outstanding requests.
	//     Origins with more pending requests are weighted proportionately less relative
	//     to others.
	//   - `"least_connections"`: Select an origin by taking into consideration origin
	//     weights, as well as each origin's number of open connections. Origins with
	//     more open connections are weighted proportionately less relative to others.
	//     Supported for HTTP/1 and HTTP/2 connections.
	Policy param.Field[OriginSteeringPolicy] `json:"policy"`
}

func (r OriginSteeringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type RandomSteering struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight float64 `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights map[string]float64 `json:"pool_weights"`
	JSON        randomSteeringJSON `json:"-"`
}

// randomSteeringJSON contains the JSON metadata for the struct [RandomSteering]
type randomSteeringJSON struct {
	DefaultWeight apijson.Field
	PoolWeights   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RandomSteering) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r randomSteeringJSON) RawJSON() string {
	return r.raw
}

// Configures pool weights.
//
//   - `steering_policy="random"`: A random pool is selected with probability
//     proportional to pool weights.
//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
//     pool's outstanding requests.
//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
//     open connections.
type RandomSteeringParam struct {
	// The default weight for pools in the load balancer that are not specified in the
	// pool_weights map.
	DefaultWeight param.Field[float64] `json:"default_weight"`
	// A mapping of pool IDs to custom weights. The weight is relative to other pools
	// in the load balancer.
	PoolWeights param.Field[map[string]float64] `json:"pool_weights"`
}

func (r RandomSteeringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type Rules struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition string `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled bool `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse RulesFixedResponse `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name string `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides RulesOverrides `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority int64 `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates bool      `json:"terminates"`
	JSON       rulesJSON `json:"-"`
}

// rulesJSON contains the JSON metadata for the struct [Rules]
type rulesJSON struct {
	Condition     apijson.Field
	Disabled      apijson.Field
	FixedResponse apijson.Field
	Name          apijson.Field
	Overrides     apijson.Field
	Priority      apijson.Field
	Terminates    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Rules) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesJSON) RawJSON() string {
	return r.raw
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type RulesFixedResponse struct {
	// The http 'Content-Type' header to include in the response.
	ContentType string `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location string `json:"location"`
	// Text to include as the http body.
	MessageBody string `json:"message_body"`
	// The http status code to respond with.
	StatusCode int64                  `json:"status_code"`
	JSON       rulesFixedResponseJSON `json:"-"`
}

// rulesFixedResponseJSON contains the JSON metadata for the struct
// [RulesFixedResponse]
type rulesFixedResponseJSON struct {
	ContentType apijson.Field
	Location    apijson.Field
	MessageBody apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesFixedResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesFixedResponseJSON) RawJSON() string {
	return r.raw
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type RulesOverrides struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting AdaptiveRouting `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools map[string][]string `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools []DefaultPools `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool string `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy LocationStrategy `json:"location_strategy"`
	// Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	POPPools map[string][]string `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering RandomSteering `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools map[string][]string `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are: - `"cookie"`: On the first
	// request to a proxied load balancer, a cookie is generated, encoding information
	// of which origin the request will be forwarded to. Subsequent requests, by the
	// same client to the same load balancer, will be sent to the origin server the
	// cookie encodes, for the duration of the cookie and as long as the origin server
	// remains healthy. If the cookie has expired or the origin server is unhealthy,
	// then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the
	// same as `"cookie"` except the initial origin selection is stable and based on
	// the client's ip address. - `"header"`: On the first request to a proxied load
	// balancer, a session key based on the configured HTTP headers (see
	// `session_affinity_attributes.headers`) is generated, encoding the request
	// headers used for storing in the load balancer session state which origin the
	// request will be forwarded to. Subsequent requests to the load balancer with the
	// same headers will be sent to the same origin server, for the duration of the
	// session and as long as the origin server remains healthy. If the session has
	// been idle for the duration of `session_affinity_ttl` seconds or the origin
	// server is unhealthy, then a new origin server is calculated and used. See
	// `headers` in `session_affinity_attributes` for additional required
	// configuration.
	SessionAffinity SessionAffinity `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes SessionAffinityAttributes `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are: -
	// `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless
	// explicitly set. The accepted range of values is between [1800, 604800]. -
	// `"header"`: The current default of 1800 seconds will be used unless explicitly
	// set. The accepted range of values is between [30, 3600]. Note: With session
	// affinity by header, sessions only expire after they haven't been used for the
	// number of seconds specified.
	SessionAffinityTTL float64 `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy SteeringPolicy `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL  float64            `json:"ttl"`
	JSON rulesOverridesJSON `json:"-"`
}

// rulesOverridesJSON contains the JSON metadata for the struct [RulesOverrides]
type rulesOverridesJSON struct {
	AdaptiveRouting           apijson.Field
	CountryPools              apijson.Field
	DefaultPools              apijson.Field
	FallbackPool              apijson.Field
	LocationStrategy          apijson.Field
	POPPools                  apijson.Field
	RandomSteering            apijson.Field
	RegionPools               apijson.Field
	SessionAffinity           apijson.Field
	SessionAffinityAttributes apijson.Field
	SessionAffinityTTL        apijson.Field
	SteeringPolicy            apijson.Field
	TTL                       apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RulesOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesOverridesJSON) RawJSON() string {
	return r.raw
}

// A rule object containing conditions and overrides for this load balancer to
// evaluate.
type RulesParam struct {
	// The condition expressions to evaluate. If the condition evaluates to true, the
	// overrides or fixed_response in this rule will be applied. An empty condition is
	// always true. For more details on condition expressions, please see
	// https://developers.cloudflare.com/load-balancing/understand-basics/load-balancing-rules/expressions.
	Condition param.Field[string] `json:"condition"`
	// Disable this specific rule. It will no longer be evaluated by this load
	// balancer.
	Disabled param.Field[bool] `json:"disabled"`
	// A collection of fields used to directly respond to the eyeball instead of
	// routing to a pool. If a fixed_response is supplied the rule will be marked as
	// terminates.
	FixedResponse param.Field[RulesFixedResponseParam] `json:"fixed_response"`
	// Name of this rule. Only used for human readability.
	Name param.Field[string] `json:"name"`
	// A collection of overrides to apply to the load balancer when this rule's
	// condition is true. All fields are optional.
	Overrides param.Field[RulesOverridesParam] `json:"overrides"`
	// The order in which rules should be executed in relation to each other. Lower
	// values are executed first. Values do not need to be sequential. If no value is
	// provided for any rule the array order of the rules field will be used to assign
	// a priority.
	Priority param.Field[int64] `json:"priority"`
	// If this rule's condition is true, this causes rule evaluation to stop after
	// processing this rule.
	Terminates param.Field[bool] `json:"terminates"`
}

func (r RulesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of fields used to directly respond to the eyeball instead of
// routing to a pool. If a fixed_response is supplied the rule will be marked as
// terminates.
type RulesFixedResponseParam struct {
	// The http 'Content-Type' header to include in the response.
	ContentType param.Field[string] `json:"content_type"`
	// The http 'Location' header to include in the response.
	Location param.Field[string] `json:"location"`
	// Text to include as the http body.
	MessageBody param.Field[string] `json:"message_body"`
	// The http status code to respond with.
	StatusCode param.Field[int64] `json:"status_code"`
}

func (r RulesFixedResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A collection of overrides to apply to the load balancer when this rule's
// condition is true. All fields are optional.
type RulesOverridesParam struct {
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]DefaultPoolsParam] `json:"default_pools"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	POPPools param.Field[map[string][]string] `json:"pop_pools"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are: - `"cookie"`: On the first
	// request to a proxied load balancer, a cookie is generated, encoding information
	// of which origin the request will be forwarded to. Subsequent requests, by the
	// same client to the same load balancer, will be sent to the origin server the
	// cookie encodes, for the duration of the cookie and as long as the origin server
	// remains healthy. If the cookie has expired or the origin server is unhealthy,
	// then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the
	// same as `"cookie"` except the initial origin selection is stable and based on
	// the client's ip address. - `"header"`: On the first request to a proxied load
	// balancer, a session key based on the configured HTTP headers (see
	// `session_affinity_attributes.headers`) is generated, encoding the request
	// headers used for storing in the load balancer session state which origin the
	// request will be forwarded to. Subsequent requests to the load balancer with the
	// same headers will be sent to the same origin server, for the duration of the
	// session and as long as the origin server remains healthy. If the session has
	// been idle for the duration of `session_affinity_ttl` seconds or the origin
	// server is unhealthy, then a new origin server is calculated and used. See
	// `headers` in `session_affinity_attributes` for additional required
	// configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are: -
	// `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless
	// explicitly set. The accepted range of values is between [1800, 604800]. -
	// `"header"`: The current default of 1800 seconds will be used unless explicitly
	// set. The accepted range of values is between [30, 3600]. Note: With session
	// affinity by header, sessions only expire after they haven't been used for the
	// number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r RulesOverridesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the type of session affinity the load balancer should use unless
// specified as `"none"`. The supported types are: - `"cookie"`: On the first
// request to a proxied load balancer, a cookie is generated, encoding information
// of which origin the request will be forwarded to. Subsequent requests, by the
// same client to the same load balancer, will be sent to the origin server the
// cookie encodes, for the duration of the cookie and as long as the origin server
// remains healthy. If the cookie has expired or the origin server is unhealthy,
// then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the
// same as `"cookie"` except the initial origin selection is stable and based on
// the client's ip address. - `"header"`: On the first request to a proxied load
// balancer, a session key based on the configured HTTP headers (see
// `session_affinity_attributes.headers`) is generated, encoding the request
// headers used for storing in the load balancer session state which origin the
// request will be forwarded to. Subsequent requests to the load balancer with the
// same headers will be sent to the same origin server, for the duration of the
// session and as long as the origin server remains healthy. If the session has
// been idle for the duration of `session_affinity_ttl` seconds or the origin
// server is unhealthy, then a new origin server is calculated and used. See
// `headers` in `session_affinity_attributes` for additional required
// configuration.
type SessionAffinity string

const (
	SessionAffinityNone     SessionAffinity = "none"
	SessionAffinityCookie   SessionAffinity = "cookie"
	SessionAffinityIPCookie SessionAffinity = "ip_cookie"
	SessionAffinityHeader   SessionAffinity = "header"
)

func (r SessionAffinity) IsKnown() bool {
	switch r {
	case SessionAffinityNone, SessionAffinityCookie, SessionAffinityIPCookie, SessionAffinityHeader:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type SessionAffinityAttributes struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration float64 `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers []string `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are: - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	// specified by the `headers` session affinity attribute, otherwise sessions aren't
	// created. - `"false"`: Load balancing requests must contain _at least one_ of the
	// HTTP headers specified by the `headers` session affinity attribute, otherwise
	// sessions aren't created.
	RequireAllHeaders bool `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite SessionAffinityAttributesSamesite `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure SessionAffinityAttributesSecure `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are: - `"none"`: No failover
	// takes place for sessions pinned to the origin (default). - `"temporary"`:
	// Traffic will be sent to another other healthy origin until the originally pinned
	// origin is available; note that this can potentially result in heavy origin
	// flapping. - `"sticky"`: The session affinity cookie is updated and subsequent
	// requests are sent to the new origin. Note: Zero-downtime failover with sticky
	// sessions is currently not supported for session affinity by header.
	ZeroDowntimeFailover SessionAffinityAttributesZeroDowntimeFailover `json:"zero_downtime_failover"`
	JSON                 sessionAffinityAttributesJSON                 `json:"-"`
}

// sessionAffinityAttributesJSON contains the JSON metadata for the struct
// [SessionAffinityAttributes]
type sessionAffinityAttributesJSON struct {
	DrainDuration        apijson.Field
	Headers              apijson.Field
	RequireAllHeaders    apijson.Field
	Samesite             apijson.Field
	Secure               apijson.Field
	ZeroDowntimeFailover apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *SessionAffinityAttributes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sessionAffinityAttributesJSON) RawJSON() string {
	return r.raw
}

// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
// when using value "None", the secure attribute can not be set to "Never".
type SessionAffinityAttributesSamesite string

const (
	SessionAffinityAttributesSamesiteAuto   SessionAffinityAttributesSamesite = "Auto"
	SessionAffinityAttributesSamesiteLax    SessionAffinityAttributesSamesite = "Lax"
	SessionAffinityAttributesSamesiteNone   SessionAffinityAttributesSamesite = "None"
	SessionAffinityAttributesSamesiteStrict SessionAffinityAttributesSamesite = "Strict"
)

func (r SessionAffinityAttributesSamesite) IsKnown() bool {
	switch r {
	case SessionAffinityAttributesSamesiteAuto, SessionAffinityAttributesSamesiteLax, SessionAffinityAttributesSamesiteNone, SessionAffinityAttributesSamesiteStrict:
		return true
	}
	return false
}

// Configures the Secure attribute on session affinity cookie. Value "Always"
// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
// indicates the Secure attribute will not be set, and "Auto" will set the Secure
// attribute depending if Always Use HTTPS is enabled.
type SessionAffinityAttributesSecure string

const (
	SessionAffinityAttributesSecureAuto   SessionAffinityAttributesSecure = "Auto"
	SessionAffinityAttributesSecureAlways SessionAffinityAttributesSecure = "Always"
	SessionAffinityAttributesSecureNever  SessionAffinityAttributesSecure = "Never"
)

func (r SessionAffinityAttributesSecure) IsKnown() bool {
	switch r {
	case SessionAffinityAttributesSecureAuto, SessionAffinityAttributesSecureAlways, SessionAffinityAttributesSecureNever:
		return true
	}
	return false
}

// Configures the zero-downtime failover between origins within a pool when session
// affinity is enabled. This feature is currently incompatible with Argo, Tiered
// Cache, and Bandwidth Alliance. The supported values are: - `"none"`: No failover
// takes place for sessions pinned to the origin (default). - `"temporary"`:
// Traffic will be sent to another other healthy origin until the originally pinned
// origin is available; note that this can potentially result in heavy origin
// flapping. - `"sticky"`: The session affinity cookie is updated and subsequent
// requests are sent to the new origin. Note: Zero-downtime failover with sticky
// sessions is currently not supported for session affinity by header.
type SessionAffinityAttributesZeroDowntimeFailover string

const (
	SessionAffinityAttributesZeroDowntimeFailoverNone      SessionAffinityAttributesZeroDowntimeFailover = "none"
	SessionAffinityAttributesZeroDowntimeFailoverTemporary SessionAffinityAttributesZeroDowntimeFailover = "temporary"
	SessionAffinityAttributesZeroDowntimeFailoverSticky    SessionAffinityAttributesZeroDowntimeFailover = "sticky"
)

func (r SessionAffinityAttributesZeroDowntimeFailover) IsKnown() bool {
	switch r {
	case SessionAffinityAttributesZeroDowntimeFailoverNone, SessionAffinityAttributesZeroDowntimeFailoverTemporary, SessionAffinityAttributesZeroDowntimeFailoverSticky:
		return true
	}
	return false
}

// Configures attributes for session affinity.
type SessionAffinityAttributesParam struct {
	// Configures the drain duration in seconds. This field is only used when session
	// affinity is enabled on the load balancer.
	DrainDuration param.Field[float64] `json:"drain_duration"`
	// Configures the names of HTTP headers to base session affinity on when header
	// `session_affinity` is enabled. At least one HTTP header name must be provided.
	// To specify the exact cookies to be used, include an item in the following
	// format: `"cookie:<cookie-name-1>,<cookie-name-2>"` (example) where everything
	// after the colon is a comma-separated list of cookie names. Providing only
	// `"cookie"` will result in all cookies being used. The default max number of HTTP
	// header names that can be provided depends on your plan: 5 for Enterprise, 1 for
	// all other plans.
	Headers param.Field[[]string] `json:"headers"`
	// When header `session_affinity` is enabled, this option can be used to specify
	// how HTTP headers on load balancing requests will be used. The supported values
	// are: - `"true"`: Load balancing requests must contain _all_ of the HTTP headers
	// specified by the `headers` session affinity attribute, otherwise sessions aren't
	// created. - `"false"`: Load balancing requests must contain _at least one_ of the
	// HTTP headers specified by the `headers` session affinity attribute, otherwise
	// sessions aren't created.
	RequireAllHeaders param.Field[bool] `json:"require_all_headers"`
	// Configures the SameSite attribute on session affinity cookie. Value "Auto" will
	// be translated to "Lax" or "None" depending if Always Use HTTPS is enabled. Note:
	// when using value "None", the secure attribute can not be set to "Never".
	Samesite param.Field[SessionAffinityAttributesSamesite] `json:"samesite"`
	// Configures the Secure attribute on session affinity cookie. Value "Always"
	// indicates the Secure attribute will be set in the Set-Cookie header, "Never"
	// indicates the Secure attribute will not be set, and "Auto" will set the Secure
	// attribute depending if Always Use HTTPS is enabled.
	Secure param.Field[SessionAffinityAttributesSecure] `json:"secure"`
	// Configures the zero-downtime failover between origins within a pool when session
	// affinity is enabled. This feature is currently incompatible with Argo, Tiered
	// Cache, and Bandwidth Alliance. The supported values are: - `"none"`: No failover
	// takes place for sessions pinned to the origin (default). - `"temporary"`:
	// Traffic will be sent to another other healthy origin until the originally pinned
	// origin is available; note that this can potentially result in heavy origin
	// flapping. - `"sticky"`: The session affinity cookie is updated and subsequent
	// requests are sent to the new origin. Note: Zero-downtime failover with sticky
	// sessions is currently not supported for session affinity by header.
	ZeroDowntimeFailover param.Field[SessionAffinityAttributesZeroDowntimeFailover] `json:"zero_downtime_failover"`
}

func (r SessionAffinityAttributesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Steering Policy for this load balancer.
//
//   - `"off"`: Use `default_pools`.
//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
//     requests, the country for `country_pools` is determined by
//     `location_strategy`.
//   - `"random"`: Select a pool randomly.
//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
//     default_pools (requires pool health checks).
//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
//     pool using the Cloudflare PoP location for proxied requests or the location
//     determined by `location_strategy` for non-proxied requests.
//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of outstanding
//     requests. Pools with more pending requests are weighted proportionately less
//     relative to others.
//   - `"least_connections"`: Select a pool by taking into consideration
//     `random_steering` weights, as well as each pool's number of open connections.
//     Pools with more open connections are weighted proportionately less relative to
//     others. Supported for HTTP/1 and HTTP/2 connections.
//   - `""`: Will map to `"geo"` if you use
//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
type SteeringPolicy string

const (
	SteeringPolicyOff                      SteeringPolicy = "off"
	SteeringPolicyGeo                      SteeringPolicy = "geo"
	SteeringPolicyRandom                   SteeringPolicy = "random"
	SteeringPolicyDynamicLatency           SteeringPolicy = "dynamic_latency"
	SteeringPolicyProximity                SteeringPolicy = "proximity"
	SteeringPolicyLeastOutstandingRequests SteeringPolicy = "least_outstanding_requests"
	SteeringPolicyLeastConnections         SteeringPolicy = "least_connections"
	SteeringPolicyEmpty                    SteeringPolicy = ""
)

func (r SteeringPolicy) IsKnown() bool {
	switch r {
	case SteeringPolicyOff, SteeringPolicyGeo, SteeringPolicyRandom, SteeringPolicyDynamicLatency, SteeringPolicyProximity, SteeringPolicyLeastOutstandingRequests, SteeringPolicyLeastConnections, SteeringPolicyEmpty:
		return true
	}
	return false
}

type LoadBalancerDeleteResponse struct {
	ID   string                         `json:"id"`
	JSON loadBalancerDeleteResponseJSON `json:"-"`
}

// loadBalancerDeleteResponseJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponse]
type loadBalancerDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerNewParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]DefaultPoolsParam] `json:"default_pools,required"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool,required"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// List of networks where Load Balancer or Pool is enabled.
	Networks param.Field[[]string] `json:"networks"`
	// Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	POPPools param.Field[map[string][]string] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]RulesParam] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are: - `"cookie"`: On the first
	// request to a proxied load balancer, a cookie is generated, encoding information
	// of which origin the request will be forwarded to. Subsequent requests, by the
	// same client to the same load balancer, will be sent to the origin server the
	// cookie encodes, for the duration of the cookie and as long as the origin server
	// remains healthy. If the cookie has expired or the origin server is unhealthy,
	// then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the
	// same as `"cookie"` except the initial origin selection is stable and based on
	// the client's ip address. - `"header"`: On the first request to a proxied load
	// balancer, a session key based on the configured HTTP headers (see
	// `session_affinity_attributes.headers`) is generated, encoding the request
	// headers used for storing in the load balancer session state which origin the
	// request will be forwarded to. Subsequent requests to the load balancer with the
	// same headers will be sent to the same origin server, for the duration of the
	// session and as long as the origin server remains healthy. If the session has
	// been idle for the duration of `session_affinity_ttl` seconds or the origin
	// server is unhealthy, then a new origin server is calculated and used. See
	// `headers` in `session_affinity_attributes` for additional required
	// configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are: -
	// `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless
	// explicitly set. The accepted range of values is between [1800, 604800]. -
	// `"header"`: The current default of 1800 seconds will be used unless explicitly
	// set. The accepted range of values is between [30, 3600]. Note: With session
	// affinity by header, sessions only expire after they haven't been used for the
	// number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancer          `json:"result,required"`
	// Whether the API call was successful.
	Success LoadBalancerNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerNewResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerNewResponseEnvelope]
type loadBalancerNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LoadBalancerNewResponseEnvelopeSuccess bool

const (
	LoadBalancerNewResponseEnvelopeSuccessTrue LoadBalancerNewResponseEnvelopeSuccess = true
)

func (r LoadBalancerNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerUpdateParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]DefaultPoolsParam] `json:"default_pools,required"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool,required"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled param.Field[bool] `json:"enabled"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// List of networks where Load Balancer or Pool is enabled.
	Networks param.Field[[]string] `json:"networks"`
	// Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	POPPools param.Field[map[string][]string] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]RulesParam] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are: - `"cookie"`: On the first
	// request to a proxied load balancer, a cookie is generated, encoding information
	// of which origin the request will be forwarded to. Subsequent requests, by the
	// same client to the same load balancer, will be sent to the origin server the
	// cookie encodes, for the duration of the cookie and as long as the origin server
	// remains healthy. If the cookie has expired or the origin server is unhealthy,
	// then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the
	// same as `"cookie"` except the initial origin selection is stable and based on
	// the client's ip address. - `"header"`: On the first request to a proxied load
	// balancer, a session key based on the configured HTTP headers (see
	// `session_affinity_attributes.headers`) is generated, encoding the request
	// headers used for storing in the load balancer session state which origin the
	// request will be forwarded to. Subsequent requests to the load balancer with the
	// same headers will be sent to the same origin server, for the duration of the
	// session and as long as the origin server remains healthy. If the session has
	// been idle for the duration of `session_affinity_ttl` seconds or the origin
	// server is unhealthy, then a new origin server is calculated and used. See
	// `headers` in `session_affinity_attributes` for additional required
	// configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are: -
	// `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless
	// explicitly set. The accepted range of values is between [1800, 604800]. -
	// `"header"`: The current default of 1800 seconds will be used unless explicitly
	// set. The accepted range of values is between [30, 3600]. Note: With session
	// affinity by header, sessions only expire after they haven't been used for the
	// number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancer          `json:"result,required"`
	// Whether the API call was successful.
	Success LoadBalancerUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerUpdateResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerUpdateResponseEnvelope]
type loadBalancerUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LoadBalancerUpdateResponseEnvelopeSuccess bool

const (
	LoadBalancerUpdateResponseEnvelopeSuccessTrue LoadBalancerUpdateResponseEnvelopeSuccess = true
)

func (r LoadBalancerUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerListParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LoadBalancerDeleteParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LoadBalancerDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   LoadBalancerDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success LoadBalancerDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerDeleteResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerDeleteResponseEnvelope]
type loadBalancerDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LoadBalancerDeleteResponseEnvelopeSuccess bool

const (
	LoadBalancerDeleteResponseEnvelopeSuccessTrue LoadBalancerDeleteResponseEnvelopeSuccess = true
)

func (r LoadBalancerDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerEditParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Controls features that modify the routing of requests to pools and origins in
	// response to dynamic conditions, such as during the interval between active
	// health monitoring requests. For example, zero-downtime failover occurs
	// immediately when an origin becomes unavailable due to HTTP 521, 522, or 523
	// response codes. If there is another healthy origin in the same pool, the request
	// is retried once against this alternate origin.
	AdaptiveRouting param.Field[AdaptiveRoutingParam] `json:"adaptive_routing"`
	// A mapping of country codes to a list of pool IDs (ordered by their failover
	// priority) for the given country. Any country not explicitly defined will fall
	// back to using the corresponding region_pool mapping if it exists else to
	// default_pools.
	CountryPools param.Field[map[string][]string] `json:"country_pools"`
	// A list of pool IDs ordered by their failover priority. Pools defined here are
	// used by default, or when region_pools are not configured for a given region.
	DefaultPools param.Field[[]DefaultPoolsParam] `json:"default_pools"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// Whether to enable (the default) this load balancer.
	Enabled param.Field[bool] `json:"enabled"`
	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool param.Field[string] `json:"fallback_pool"`
	// Controls location-based steering for non-proxied requests. See `steering_policy`
	// to learn how steering is affected.
	LocationStrategy param.Field[LocationStrategyParam] `json:"location_strategy"`
	// The DNS hostname to associate with your Load Balancer. If this hostname already
	// exists as a DNS record in Cloudflare's DNS, the Load Balancer will take
	// precedence and the DNS record will not be used.
	Name param.Field[string] `json:"name"`
	// Enterprise only: A mapping of Cloudflare PoP identifiers to a list of pool IDs
	// (ordered by their failover priority) for the PoP (datacenter). Any PoPs not
	// explicitly defined will fall back to using the corresponding country_pool, then
	// region_pool mapping if it exists else to default_pools.
	POPPools param.Field[map[string][]string] `json:"pop_pools"`
	// Whether the hostname should be gray clouded (false) or orange clouded (true).
	Proxied param.Field[bool] `json:"proxied"`
	// Configures pool weights.
	//
	//   - `steering_policy="random"`: A random pool is selected with probability
	//     proportional to pool weights.
	//   - `steering_policy="least_outstanding_requests"`: Use pool weights to scale each
	//     pool's outstanding requests.
	//   - `steering_policy="least_connections"`: Use pool weights to scale each pool's
	//     open connections.
	RandomSteering param.Field[RandomSteeringParam] `json:"random_steering"`
	// A mapping of region codes to a list of pool IDs (ordered by their failover
	// priority) for the given region. Any regions not explicitly defined will fall
	// back to using default_pools.
	RegionPools param.Field[map[string][]string] `json:"region_pools"`
	// BETA Field Not General Access: A list of rules for this load balancer to
	// execute.
	Rules param.Field[[]RulesParam] `json:"rules"`
	// Specifies the type of session affinity the load balancer should use unless
	// specified as `"none"`. The supported types are: - `"cookie"`: On the first
	// request to a proxied load balancer, a cookie is generated, encoding information
	// of which origin the request will be forwarded to. Subsequent requests, by the
	// same client to the same load balancer, will be sent to the origin server the
	// cookie encodes, for the duration of the cookie and as long as the origin server
	// remains healthy. If the cookie has expired or the origin server is unhealthy,
	// then a new origin server is calculated and used. - `"ip_cookie"`: Behaves the
	// same as `"cookie"` except the initial origin selection is stable and based on
	// the client's ip address. - `"header"`: On the first request to a proxied load
	// balancer, a session key based on the configured HTTP headers (see
	// `session_affinity_attributes.headers`) is generated, encoding the request
	// headers used for storing in the load balancer session state which origin the
	// request will be forwarded to. Subsequent requests to the load balancer with the
	// same headers will be sent to the same origin server, for the duration of the
	// session and as long as the origin server remains healthy. If the session has
	// been idle for the duration of `session_affinity_ttl` seconds or the origin
	// server is unhealthy, then a new origin server is calculated and used. See
	// `headers` in `session_affinity_attributes` for additional required
	// configuration.
	SessionAffinity param.Field[SessionAffinity] `json:"session_affinity"`
	// Configures attributes for session affinity.
	SessionAffinityAttributes param.Field[SessionAffinityAttributesParam] `json:"session_affinity_attributes"`
	// Time, in seconds, until a client's session expires after being created. Once the
	// expiry time has been reached, subsequent requests may get sent to a different
	// origin server. The accepted ranges per `session_affinity` policy are: -
	// `"cookie"` / `"ip_cookie"`: The current default of 23 hours will be used unless
	// explicitly set. The accepted range of values is between [1800, 604800]. -
	// `"header"`: The current default of 1800 seconds will be used unless explicitly
	// set. The accepted range of values is between [30, 3600]. Note: With session
	// affinity by header, sessions only expire after they haven't been used for the
	// number of seconds specified.
	SessionAffinityTTL param.Field[float64] `json:"session_affinity_ttl"`
	// Steering Policy for this load balancer.
	//
	//   - `"off"`: Use `default_pools`.
	//   - `"geo"`: Use `region_pools`/`country_pools`/`pop_pools`. For non-proxied
	//     requests, the country for `country_pools` is determined by
	//     `location_strategy`.
	//   - `"random"`: Select a pool randomly.
	//   - `"dynamic_latency"`: Use round trip time to select the closest pool in
	//     default_pools (requires pool health checks).
	//   - `"proximity"`: Use the pools' latitude and longitude to select the closest
	//     pool using the Cloudflare PoP location for proxied requests or the location
	//     determined by `location_strategy` for non-proxied requests.
	//   - `"least_outstanding_requests"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of outstanding
	//     requests. Pools with more pending requests are weighted proportionately less
	//     relative to others.
	//   - `"least_connections"`: Select a pool by taking into consideration
	//     `random_steering` weights, as well as each pool's number of open connections.
	//     Pools with more open connections are weighted proportionately less relative to
	//     others. Supported for HTTP/1 and HTTP/2 connections.
	//   - `""`: Will map to `"geo"` if you use
	//     `region_pools`/`country_pools`/`pop_pools` otherwise `"off"`.
	SteeringPolicy param.Field[SteeringPolicy] `json:"steering_policy"`
	// Time to live (TTL) of the DNS entry for the IP address returned by this load
	// balancer. This only applies to gray-clouded (unproxied) load balancers.
	TTL param.Field[float64] `json:"ttl"`
}

func (r LoadBalancerEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancer          `json:"result,required"`
	// Whether the API call was successful.
	Success LoadBalancerEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerEditResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerEditResponseEnvelope]
type loadBalancerEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LoadBalancerEditResponseEnvelopeSuccess bool

const (
	LoadBalancerEditResponseEnvelopeSuccessTrue LoadBalancerEditResponseEnvelopeSuccess = true
)

func (r LoadBalancerEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type LoadBalancerGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type LoadBalancerGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   LoadBalancer          `json:"result,required"`
	// Whether the API call was successful.
	Success LoadBalancerGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponseEnvelope]
type loadBalancerGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LoadBalancerGetResponseEnvelopeSuccess bool

const (
	LoadBalancerGetResponseEnvelopeSuccessTrue LoadBalancerGetResponseEnvelopeSuccess = true
)

func (r LoadBalancerGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancerGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
