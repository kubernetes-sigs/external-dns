package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const lbPath = "/v2/load-balancers"

// LoadBalancerService is the interface to interact with the server endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/load-balancer
type LoadBalancerService interface {
	Create(ctx context.Context, createReq *LoadBalancerReq) (*LoadBalancer, error)
	Get(ctx context.Context, ID string) (*LoadBalancer, error)
	Update(ctx context.Context, ID string, updateReq *LoadBalancerReq) error
	Delete(ctx context.Context, ID string) error
	List(ctx context.Context, options *ListOptions) ([]LoadBalancer, *Meta, error)
	CreateForwardingRule(ctx context.Context, ID string, rule *ForwardingRule) (*ForwardingRule, error)
	GetForwardingRule(ctx context.Context, ID string, ruleID string) (*ForwardingRule, error)
	DeleteForwardingRule(ctx context.Context, ID string, RuleID string) error
	ListForwardingRules(ctx context.Context, ID string, options *ListOptions) ([]ForwardingRule, *Meta, error)
	ListFirewallRules(ctx context.Context, ID string, options *ListOptions) ([]LBFirewallRule, *Meta, error)
	GetFirewallRule(ctx context.Context, ID string, ruleID string) (*LBFirewallRule, error)
}

// LoadBalancerHandler handles interaction with the server methods for the Vultr API
type LoadBalancerHandler struct {
	client *Client
}

// LoadBalancer represent the structure of a load balancer
type LoadBalancer struct {
	ID              string           `json:"id,omitempty"`
	DateCreated     string           `json:"date_created,omitempty"`
	Region          string           `json:"region,omitempty"`
	Label           string           `json:"label,omitempty"`
	Status          string           `json:"status,omitempty"`
	IPV4            string           `json:"ipv4,omitempty"`
	IPV6            string           `json:"ipv6,omitempty"`
	Instances       []string         `json:"instances,omitempty"`
	HealthCheck     *HealthCheck     `json:"health_check,omitempty"`
	GenericInfo     *GenericInfo     `json:"generic_info,omitempty"`
	SSLInfo         *bool            `json:"has_ssl,omitempty"`
	ForwardingRules []ForwardingRule `json:"forwarding_rules,omitempty"`
	FirewallRules   []LBFirewallRule `json:"firewall_rules,omitempty"`
}

// LoadBalancerReq gives options for creating or updating a load balancer
type LoadBalancerReq struct {
	Region             string           `json:"region,omitempty"`
	Label              string           `json:"label,omitempty"`
	Instances          []string         `json:"instances"`
	HealthCheck        *HealthCheck     `json:"health_check,omitempty"`
	StickySessions     *StickySessions  `json:"sticky_session,omitempty"`
	ForwardingRules    []ForwardingRule `json:"forwarding_rules,omitempty"`
	SSL                *SSL             `json:"ssl,omitempty"`
	SSLRedirect        *bool            `json:"ssl_redirect,omitempty"`
	ProxyProtocol      *bool            `json:"proxy_protocol,omitempty"`
	BalancingAlgorithm string           `json:"balancing_algorithm,omitempty"`
	FirewallRules      []LBFirewallRule `json:"firewall_rules"`
	PrivateNetwork     *string          `json:"private_network,omitempty"`
}

// InstanceList represents instances that are attached to your load balancer
type InstanceList struct {
	InstanceList []string
}

// HealthCheck represents your health check configuration for your load balancer.
type HealthCheck struct {
	Protocol           string `json:"protocol,omitempty"`
	Port               int    `json:"port,omitempty"`
	Path               string `json:"path,omitempty"`
	CheckInterval      int    `json:"check_interval,omitempty"`
	ResponseTimeout    int    `json:"response_timeout,omitempty"`
	UnhealthyThreshold int    `json:"unhealthy_threshold,omitempty"`
	HealthyThreshold   int    `json:"healthy_threshold,omitempty"`
}

// GenericInfo represents generic configuration of your load balancer
type GenericInfo struct {
	BalancingAlgorithm string          `json:"balancing_algorithm,omitempty"`
	SSLRedirect        *bool           `json:"ssl_redirect,omitempty"`
	StickySessions     *StickySessions `json:"sticky_sessions,omitempty"`
	ProxyProtocol      *bool           `json:"proxy_protocol,omitempty"`
	PrivateNetwork     string          `json:"private_network,omitempty"`
}

// StickySessions represents cookie for your load balancer
type StickySessions struct {
	CookieName string `json:"cookie_name,omitempty"`
}

// ForwardingRules represent a list of forwarding rules
type ForwardingRules struct {
	ForwardRuleList []ForwardingRule `json:"forwarding_rules,omitempty"`
}

// ForwardingRule represent a single forwarding rule
type ForwardingRule struct {
	RuleID           string `json:"id,omitempty"`
	FrontendProtocol string `json:"frontend_protocol,omitempty"`
	FrontendPort     int    `json:"frontend_port,omitempty"`
	BackendProtocol  string `json:"backend_protocol,omitempty"`
	BackendPort      int    `json:"backend_port,omitempty"`
}

// LBFirewallRule represent a single firewall rule
type LBFirewallRule struct {
	RuleID string `json:"id,omitempty"`
	Port   int    `json:"port,omitempty"`
	IPType string `json:"ip_type,omitempty"`
	Source string `json:"source,omitempty"`
}

// SSL represents valid SSL config
type SSL struct {
	PrivateKey  string `json:"ssl_private_key,omitempty"`
	Certificate string `json:"ssl_certificate,omitempty"`
	Chain       string `json:"chain,omitempty"`
}

type lbsBase struct {
	LoadBalancers []LoadBalancer `json:"load_balancers"`
	Meta          *Meta          `json:"meta"`
}

type lbBase struct {
	LoadBalancer *LoadBalancer `json:"load_balancer"`
}

type lbRulesBase struct {
	ForwardingRules []ForwardingRule `json:"forwarding_rules"`
	Meta            *Meta            `json:"meta"`
}

type lbRuleBase struct {
	ForwardingRule *ForwardingRule `json:"forwarding_rule"`
}

type lbFWRulesBase struct {
	FirewallRules []LBFirewallRule `json:"firewall_rules"`
	Meta          *Meta            `json:"meta"`
}

type lbFWRuleBase struct {
	FirewallRule *LBFirewallRule `json:"firewall_rule"`
}

// Create a load balancer
func (l *LoadBalancerHandler) Create(ctx context.Context, createReq *LoadBalancerReq) (*LoadBalancer, error) {
	req, err := l.client.NewRequest(ctx, http.MethodPost, lbPath, createReq)
	if err != nil {
		return nil, err
	}

	var lb = new(lbBase)
	if err = l.client.DoWithContext(ctx, req, &lb); err != nil {
		return nil, err
	}

	return lb.LoadBalancer, nil
}

// Get a load balancer
func (l *LoadBalancerHandler) Get(ctx context.Context, ID string) (*LoadBalancer, error) {
	uri := fmt.Sprintf("%s/%s", lbPath, ID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var lb = new(lbBase)
	if err = l.client.DoWithContext(ctx, req, lb); err != nil {
		return nil, err
	}

	return lb.LoadBalancer, nil
}

// Update updates your your load balancer
func (l *LoadBalancerHandler) Update(ctx context.Context, ID string, updateReq *LoadBalancerReq) error {
	uri := fmt.Sprintf("%s/%s", lbPath, ID)
	req, err := l.client.NewRequest(ctx, http.MethodPatch, uri, updateReq)
	if err != nil {
		return err
	}

	return l.client.DoWithContext(ctx, req, nil)
}

// Delete a load balancer subscription.
func (l *LoadBalancerHandler) Delete(ctx context.Context, ID string) error {
	uri := fmt.Sprintf("%s/%s", lbPath, ID)
	req, err := l.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return l.client.DoWithContext(ctx, req, nil)
}

// List all load balancer subscriptions on the current account.
func (l *LoadBalancerHandler) List(ctx context.Context, options *ListOptions) ([]LoadBalancer, *Meta, error) {
	req, err := l.client.NewRequest(ctx, http.MethodGet, lbPath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	lbs := new(lbsBase)
	if err = l.client.DoWithContext(ctx, req, &lbs); err != nil {
		return nil, nil, err
	}

	return lbs.LoadBalancers, lbs.Meta, nil
}

// CreateForwardingRule will create a new forwarding rule for your load balancer subscription.
// Note the RuleID will be returned in the ForwardingRule struct
func (l *LoadBalancerHandler) CreateForwardingRule(ctx context.Context, ID string, rule *ForwardingRule) (*ForwardingRule, error) {
	uri := fmt.Sprintf("%s/%s/forwarding-rules", lbPath, ID)
	req, err := l.client.NewRequest(ctx, http.MethodPost, uri, rule)
	if err != nil {
		return nil, err
	}

	fwRule := new(lbRuleBase)
	if err = l.client.DoWithContext(ctx, req, fwRule); err != nil {
		return nil, err
	}

	return fwRule.ForwardingRule, nil
}

// GetForwardingRule will get a forwarding rule from your load balancer subscription.
func (l *LoadBalancerHandler) GetForwardingRule(ctx context.Context, ID string, ruleID string) (*ForwardingRule, error) {
	uri := fmt.Sprintf("%s/%s/forwarding-rules/%s", lbPath, ID, ruleID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	fwRule := new(lbRuleBase)
	if err = l.client.DoWithContext(ctx, req, fwRule); err != nil {
		return nil, err
	}

	return fwRule.ForwardingRule, nil
}

// ListForwardingRules lists all forwarding rules for a load balancer subscription
func (l *LoadBalancerHandler) ListForwardingRules(ctx context.Context, ID string, options *ListOptions) ([]ForwardingRule, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/forwarding-rules", lbPath, ID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	fwRules := new(lbRulesBase)
	if err = l.client.DoWithContext(ctx, req, &fwRules); err != nil {
		return nil, nil, err
	}

	return fwRules.ForwardingRules, fwRules.Meta, nil
}

// DeleteForwardingRule removes a forwarding rule from a load balancer subscription
func (l *LoadBalancerHandler) DeleteForwardingRule(ctx context.Context, ID string, RuleID string) error {
	uri := fmt.Sprintf("%s/%s/forwarding-rules/%s", lbPath, ID, RuleID)
	req, err := l.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return l.client.DoWithContext(ctx, req, nil)
}

// GetFirewallRule will get a firewall rule from your load balancer subscription.
func (l *LoadBalancerHandler) GetFirewallRule(ctx context.Context, ID string, ruleID string) (*LBFirewallRule, error) {
	uri := fmt.Sprintf("%s/%s/firewall-rules/%s", lbPath, ID, ruleID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	fwRule := new(lbFWRuleBase)
	if err = l.client.DoWithContext(ctx, req, fwRule); err != nil {
		return nil, err
	}

	return fwRule.FirewallRule, nil
}

// ListFirewallRules lists all firewall rules for a load balancer subscription
func (l *LoadBalancerHandler) ListFirewallRules(ctx context.Context, ID string, options *ListOptions) ([]LBFirewallRule, *Meta, error) {
	uri := fmt.Sprintf("%s/%s/firewall-rules", lbPath, ID)
	req, err := l.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	fwRules := new(lbFWRulesBase)
	if err = l.client.DoWithContext(ctx, req, &fwRules); err != nil {
		return nil, nil, err
	}

	return fwRules.FirewallRules, fwRules.Meta, nil
}
