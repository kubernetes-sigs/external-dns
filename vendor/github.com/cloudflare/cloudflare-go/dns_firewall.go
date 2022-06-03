package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// DNSFirewallCluster represents a DNS Firewall configuration.
type DNSFirewallCluster struct {
	ID                   string   `json:"id,omitempty"`
	Name                 string   `json:"name"`
	OriginIPs            []string `json:"origin_ips"`
	DNSFirewallIPs       []string `json:"dns_firewall_ips,omitempty"`
	MinimumCacheTTL      uint     `json:"minimum_cache_ttl,omitempty"`
	MaximumCacheTTL      uint     `json:"maximum_cache_ttl,omitempty"`
	DeprecateAnyRequests bool     `json:"deprecate_any_requests"`
	ModifiedOn           string   `json:"modified_on,omitempty"`
}

// DNSFirewallAnalyticsMetrics represents a group of aggregated DNS Firewall metrics.
type DNSFirewallAnalyticsMetrics struct {
	QueryCount         *int64   `json:"queryCount"`
	UncachedCount      *int64   `json:"uncachedCount"`
	StaleCount         *int64   `json:"staleCount"`
	ResponseTimeAvg    *float64 `json:"responseTimeAvg"`
	ResponseTimeMedian *float64 `json:"responseTimeMedian"`
	ResponseTime90th   *float64 `json:"responseTime90th"`
	ResponseTime99th   *float64 `json:"responseTime99th"`
}

// DNSFirewallAnalytics represents a set of aggregated DNS Firewall metrics.
// TODO: Add the queried data and not only the aggregated values.
type DNSFirewallAnalytics struct {
	Totals DNSFirewallAnalyticsMetrics `json:"totals"`
	Min    DNSFirewallAnalyticsMetrics `json:"min"`
	Max    DNSFirewallAnalyticsMetrics `json:"max"`
}

// DNSFirewallUserAnalyticsOptions represents range and dimension selection on analytics endpoint.
type DNSFirewallUserAnalyticsOptions struct {
	Metrics []string
	Since   *time.Time
	Until   *time.Time
}

// dnsFirewallResponse represents a DNS Firewall response.
type dnsFirewallResponse struct {
	Response
	Result *DNSFirewallCluster `json:"result"`
}

// dnsFirewallListResponse represents an array of DNS Firewall responses.
type dnsFirewallListResponse struct {
	Response
	Result []*DNSFirewallCluster `json:"result"`
}

// dnsFirewallAnalyticsResponse represents a DNS Firewall analytics response.
type dnsFirewallAnalyticsResponse struct {
	Response
	Result DNSFirewallAnalytics `json:"result"`
}

// CreateDNSFirewallCluster creates a new DNS Firewall cluster.
//
// API reference: https://api.cloudflare.com/#dns-firewall-create-dns-firewall-cluster
func (api *API) CreateDNSFirewallCluster(ctx context.Context, v DNSFirewallCluster) (*DNSFirewallCluster, error) {
	uri := fmt.Sprintf("%s/dns_firewall", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, v)
	if err != nil {
		return nil, err
	}

	response := &dnsFirewallResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// DNSFirewallCluster fetches a single DNS Firewall cluster.
//
// API reference: https://api.cloudflare.com/#dns-firewall-dns-firewall-cluster-details
func (api *API) DNSFirewallCluster(ctx context.Context, clusterID string) (*DNSFirewallCluster, error) {
	uri := fmt.Sprintf("%s/dns_firewall/%s", api.userBaseURL("/user"), clusterID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response := &dnsFirewallResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// ListDNSFirewallClusters lists the DNS Firewall clusters associated with an account.
//
// API reference: https://api.cloudflare.com/#dns-firewall-list-dns-firewall-clusters
func (api *API) ListDNSFirewallClusters(ctx context.Context) ([]*DNSFirewallCluster, error) {
	uri := fmt.Sprintf("%s/dns_firewall", api.userBaseURL("/user"))
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	response := &dnsFirewallListResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// UpdateDNSFirewallCluster updates a DNS Firewall cluster.
//
// API reference: https://api.cloudflare.com/#dns-firewall-update-dns-firewall-cluster
func (api *API) UpdateDNSFirewallCluster(ctx context.Context, clusterID string, vv DNSFirewallCluster) error {
	uri := fmt.Sprintf("%s/dns_firewall/%s", api.userBaseURL("/user"), clusterID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, vv)
	if err != nil {
		return err
	}

	response := &dnsFirewallResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// DeleteDNSFirewallCluster deletes a DNS Firewall cluster. Note that this cannot be
// undone, and will stop all traffic to that cluster.
//
// API reference: https://api.cloudflare.com/#dns-firewall-delete-dns-firewall-cluster
func (api *API) DeleteDNSFirewallCluster(ctx context.Context, clusterID string) error {
	uri := fmt.Sprintf("%s/dns_firewall/%s", api.userBaseURL("/user"), clusterID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	response := &dnsFirewallResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// encode encodes non-nil fields into URL encoded form.
func (o DNSFirewallUserAnalyticsOptions) encode() string {
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

// DNSFirewallUserAnalytics retrieves analytics report for a specified dimension and time range.
func (api *API) DNSFirewallUserAnalytics(ctx context.Context, clusterID string, o DNSFirewallUserAnalyticsOptions) (DNSFirewallAnalytics, error) {
	uri := fmt.Sprintf("%s/dns_firewall/%s/dns_analytics/report?%s", api.userBaseURL("/user"), clusterID, o.encode())
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DNSFirewallAnalytics{}, err
	}

	response := dnsFirewallAnalyticsResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return DNSFirewallAnalytics{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}
