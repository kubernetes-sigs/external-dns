package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"fmt"
)

// NOTE: Everything in this file is deprecated. See `dns_firewall.go` instead.

// VirtualDNS represents a Virtual DNS configuration.
//
// Deprecated: Use DNSFirewallCluster instead.
type VirtualDNS struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	OriginIPs            []string `json:"origin_ips"`
	VirtualDNSIPs        []string `json:"virtual_dns_ips"`
	MinimumCacheTTL      uint     `json:"minimum_cache_ttl"`
	MaximumCacheTTL      uint     `json:"maximum_cache_ttl"`
	DeprecateAnyRequests bool     `json:"deprecate_any_requests"`
	ModifiedOn           string   `json:"modified_on"`
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// VirtualDNSAnalyticsMetrics represents a group of aggregated Virtual DNS metrics.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// VirtualDNSAnalyticsMetrics respresents a group of aggregated Virtual DNS metrics.
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// VirtualDNSAnalyticsMetrics respresents a group of aggregated Virtual DNS metrics.
=======
// VirtualDNSAnalyticsMetrics represents a group of aggregated Virtual DNS metrics.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// VirtualDNSAnalyticsMetrics respresents a group of aggregated Virtual DNS metrics.
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// VirtualDNSAnalyticsMetrics respresents a group of aggregated Virtual DNS metrics.
=======
// VirtualDNSAnalyticsMetrics represents a group of aggregated Virtual DNS metrics.
>>>>>>> 6b7ce455e (update vendored files)
type VirtualDNSAnalyticsMetrics struct {
	QueryCount         *int64   `json:"queryCount"`
	UncachedCount      *int64   `json:"uncachedCount"`
	StaleCount         *int64   `json:"staleCount"`
	ResponseTimeAvg    *float64 `json:"responseTimeAvg"`
	ResponseTimeMedian *float64 `json:"responseTimeMedian"`
	ResponseTime90th   *float64 `json:"responseTime90th"`
	ResponseTime99th   *float64 `json:"responseTime99th"`
}

// VirtualDNSAnalytics represents a set of aggregated Virtual DNS metrics.
// TODO: Add the queried data and not only the aggregated values.
type VirtualDNSAnalytics struct {
	Totals VirtualDNSAnalyticsMetrics `json:"totals"`
	Min    VirtualDNSAnalyticsMetrics `json:"min"`
	Max    VirtualDNSAnalyticsMetrics `json:"max"`
}

// VirtualDNSUserAnalyticsOptions represents range and dimension selection on analytics endpoint
type VirtualDNSUserAnalyticsOptions struct {
	Metrics []string
	Since   *time.Time
	Until   *time.Time
}

// VirtualDNSResponse represents a Virtual DNS response.
type VirtualDNSResponse struct {
	Response
	Result *VirtualDNS `json:"result"`
}

// VirtualDNSListResponse represents an array of Virtual DNS responses.
type VirtualDNSListResponse struct {
	Response
	Result []*VirtualDNS `json:"result"`
}

// VirtualDNSAnalyticsResponse represents a Virtual DNS analytics response.
type VirtualDNSAnalyticsResponse struct {
	Response
	Result VirtualDNSAnalytics `json:"result"`
}

// CreateVirtualDNS creates a new Virtual DNS cluster.
//
// API reference: https://api.cloudflare.com/#virtual-dns-users--create-a-virtual-dns-cluster
func (api *API) CreateVirtualDNS(ctx context.Context, v *VirtualDNS) (*VirtualDNS, error) {
	uri := fmt.Sprintf("%s/virtual_dns", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, v)
	if err != nil {
		return nil, err
	}

	response := &VirtualDNSResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response.Result, nil
}

// VirtualDNS fetches a single virtual DNS cluster.
//
// API reference: https://api.cloudflare.com/#virtual-dns-users--get-a-virtual-dns-cluster
func (api *API) VirtualDNS(ctx context.Context, virtualDNSID string) (*VirtualDNS, error) {
	uri := fmt.Sprintf("%s/virtual_dns/%s", api.userBaseURL("/user"), virtualDNSID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response := &VirtualDNSResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response.Result, nil
}

// ListVirtualDNS lists the virtual DNS clusters associated with an account.
//
// API reference: https://api.cloudflare.com/#virtual-dns-users--get-virtual-dns-clusters
func (api *API) ListVirtualDNS(ctx context.Context) ([]*VirtualDNS, error) {
	uri := fmt.Sprintf("%s/virtual_dns", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response := &VirtualDNSListResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}

	return response.Result, nil
}

// UpdateVirtualDNS updates a Virtual DNS cluster.
//
// API reference: https://api.cloudflare.com/#virtual-dns-users--modify-a-virtual-dns-cluster
func (api *API) UpdateVirtualDNS(ctx context.Context, virtualDNSID string, vv VirtualDNS) error {
	uri := fmt.Sprintf("%s/virtual_dns/%s", api.userBaseURL("/user"), virtualDNSID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, vv)
	if err != nil {
		return err
	}

	response := &VirtualDNSResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}

	return nil
}

// DeleteVirtualDNS deletes a Virtual DNS cluster. Note that this cannot be
// undone, and will stop all traffic to that cluster.
//
// API reference: https://api.cloudflare.com/#virtual-dns-users--delete-a-virtual-dns-cluster
func (api *API) DeleteVirtualDNS(ctx context.Context, virtualDNSID string) error {
	uri := fmt.Sprintf("%s/virtual_dns/%s", api.userBaseURL("/user"), virtualDNSID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	response := &VirtualDNSResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}

	return nil
}

// encode encodes non-nil fields into URL encoded form.
func (o VirtualDNSUserAnalyticsOptions) encode() string {
	v := url.Values{}
	if o.Since != nil {
		v.Set("since", (*o.Since).UTC().Format(time.RFC3339))
	}
	if o.Until != nil {
		v.Set("until", (*o.Until).UTC().Format(time.RFC3339))
	}
	if o.Metrics != nil {
		v.Set("metrics", strings.Join(o.Metrics, ","))
	}
	return v.Encode()
}

// VirtualDNSUserAnalytics retrieves analytics report for a specified dimension and time range
func (api *API) VirtualDNSUserAnalytics(ctx context.Context, virtualDNSID string, o VirtualDNSUserAnalyticsOptions) (VirtualDNSAnalytics, error) {
	uri := fmt.Sprintf("%s/virtual_dns/%s/dns_analytics/report?%s", api.userBaseURL("/user"), virtualDNSID, o.encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return VirtualDNSAnalytics{}, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// VirtualDNS represents a Virtual DNS configuration.
type VirtualDNS struct {
	ID                   string   `json:"id"`
	Name                 string   `json:"name"`
	OriginIPs            []string `json:"origin_ips"`
	VirtualDNSIPs        []string `json:"virtual_dns_ips"`
	MinimumCacheTTL      uint     `json:"minimum_cache_ttl"`
	MaximumCacheTTL      uint     `json:"maximum_cache_ttl"`
	DeprecateAnyRequests bool     `json:"deprecate_any_requests"`
	ModifiedOn           string   `json:"modified_on"`
}

// VirtualDNSAnalyticsMetrics represents a group of aggregated Virtual DNS metrics.
//
// Deprecated: Use DNSFirewallAnalyticsMetrics instead.
type VirtualDNSAnalyticsMetrics DNSFirewallAnalyticsMetrics

// VirtualDNSAnalytics represents a set of aggregated Virtual DNS metrics.
//
// Deprecated: Use DNSFirewallAnalytics instead.
type VirtualDNSAnalytics struct {
	Totals VirtualDNSAnalyticsMetrics `json:"totals"`
	Min    VirtualDNSAnalyticsMetrics `json:"min"`
	Max    VirtualDNSAnalyticsMetrics `json:"max"`
}

// VirtualDNSUserAnalyticsOptions represents range and dimension selection on analytics endpoint
//
// Deprecated: Use DNSFirewallUserAnalyticsOptions instead.
type VirtualDNSUserAnalyticsOptions DNSFirewallUserAnalyticsOptions

// VirtualDNSResponse represents a Virtual DNS response.
//
// Deprecated: This internal type will be removed in the future.
type VirtualDNSResponse struct {
	Response
	Result *VirtualDNS `json:"result"`
}

// VirtualDNSListResponse represents an array of Virtual DNS responses.
//
// Deprecated: This internal type will be removed in the future.
type VirtualDNSListResponse struct {
	Response
	Result []*VirtualDNS `json:"result"`
}

// VirtualDNSAnalyticsResponse represents a Virtual DNS analytics response.
//
// Deprecated: This internal type will be removed in the future.
type VirtualDNSAnalyticsResponse struct {
	Response
	Result VirtualDNSAnalytics `json:"result"`
}

// CreateVirtualDNS creates a new Virtual DNS cluster.
//
// Deprecated: Use CreateDNSFirewallCluster instead.
func (api *API) CreateVirtualDNS(ctx context.Context, v *VirtualDNS) (*VirtualDNS, error) {
	if v == nil {
		return nil, fmt.Errorf("cluster must not be nil")
	}

	res, err := api.CreateDNSFirewallCluster(ctx, v.vdnsUpgrade())
	return res.vdnsDowngrade(), err
}

// VirtualDNS fetches a single virtual DNS cluster.
//
// Deprecated: Use DNSFirewallCluster instead.
func (api *API) VirtualDNS(ctx context.Context, virtualDNSID string) (*VirtualDNS, error) {
	res, err := api.DNSFirewallCluster(ctx, virtualDNSID)
	return res.vdnsDowngrade(), err
}

// ListVirtualDNS lists the virtual DNS clusters associated with an account.
//
// Deprecated: Use ListDNSFirewallClusters instead.
func (api *API) ListVirtualDNS(ctx context.Context) ([]*VirtualDNS, error) {
	res, err := api.ListDNSFirewallClusters(ctx)
	if res == nil {
		return nil, err
	}

	clusters := make([]*VirtualDNS, 0, len(res))
	for _, cluster := range res {
		clusters = append(clusters, cluster.vdnsDowngrade())
	}

	return clusters, err
}

// UpdateVirtualDNS updates a Virtual DNS cluster.
//
// Deprecated: Use UpdateDNSFirewallCluster instead.
func (api *API) UpdateVirtualDNS(ctx context.Context, virtualDNSID string, vv VirtualDNS) error {
	return api.UpdateDNSFirewallCluster(ctx, virtualDNSID, vv.vdnsUpgrade())
}

// DeleteVirtualDNS deletes a Virtual DNS cluster. Note that this cannot be
// undone, and will stop all traffic to that cluster.
//
// Deprecated: Use DeleteDNSFirewallCluster instead.
func (api *API) DeleteVirtualDNS(ctx context.Context, virtualDNSID string) error {
	return api.DeleteDNSFirewallCluster(ctx, virtualDNSID)
}

// VirtualDNSUserAnalytics retrieves analytics report for a specified dimension and time range
//
// Deprecated: Use DNSFirewallUserAnalytics instead.
func (api *API) VirtualDNSUserAnalytics(ctx context.Context, virtualDNSID string, o VirtualDNSUserAnalyticsOptions) (VirtualDNSAnalytics, error) {
<<<<<<< HEAD
	uri := fmt.Sprintf("%s/virtual_dns/%s/dns_analytics/report?%s", api.userBaseURL("/user"), virtualDNSID, o.encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return VirtualDNSAnalytics{}, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return VirtualDNSAnalytics{}, errors.Wrap(err, errMakeRequestError)
=======
		return VirtualDNSAnalytics{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	uri := fmt.Sprintf("%s/virtual_dns/%s/dns_analytics/report?%s", api.userBaseURL("/user"), virtualDNSID, o.encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return VirtualDNSAnalytics{}, err
=======
	res, err := api.DNSFirewallUserAnalytics(ctx, virtualDNSID, DNSFirewallUserAnalyticsOptions(o))
	return res.vdnsDowngrade(), err
}

// --- Compatibility helper functions ---

func (v VirtualDNS) vdnsUpgrade() DNSFirewallCluster {
	return DNSFirewallCluster{
		ID:                   v.ID,
		Name:                 v.Name,
		OriginIPs:            v.OriginIPs,
		DNSFirewallIPs:       v.VirtualDNSIPs,
		MinimumCacheTTL:      v.MinimumCacheTTL,
		MaximumCacheTTL:      v.MaximumCacheTTL,
		DeprecateAnyRequests: v.DeprecateAnyRequests,
		ModifiedOn:           v.ModifiedOn,
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
	}
}

func (v *DNSFirewallCluster) vdnsDowngrade() *VirtualDNS {
	if v == nil {
		return nil
	}

	return &VirtualDNS{
		ID:                   v.ID,
		Name:                 v.Name,
		OriginIPs:            v.OriginIPs,
		VirtualDNSIPs:        v.DNSFirewallIPs,
		MinimumCacheTTL:      v.MinimumCacheTTL,
		MaximumCacheTTL:      v.MaximumCacheTTL,
		DeprecateAnyRequests: v.DeprecateAnyRequests,
		ModifiedOn:           v.ModifiedOn,
	}
}

func (v DNSFirewallAnalytics) vdnsDowngrade() VirtualDNSAnalytics {
	return VirtualDNSAnalytics{
		Totals: VirtualDNSAnalyticsMetrics(v.Totals),
		Min:    VirtualDNSAnalyticsMetrics(v.Min),
		Max:    VirtualDNSAnalyticsMetrics(v.Max),
	}
}
