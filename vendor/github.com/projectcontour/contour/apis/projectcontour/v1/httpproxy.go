// Copyright Project Contour Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HTTPProxySpec defines the spec of the CRD.
type HTTPProxySpec struct {
	// Virtualhost appears at most once. If it is present, the object is considered
	// to be a "root" HTTPProxy.
	// +optional
	VirtualHost *VirtualHost `json:"virtualhost,omitempty"`
	// Routes are the ingress routes. If TCPProxy is present, Routes is ignored.
	//  +optional
	Routes []Route `json:"routes,omitempty"`
	// TCPProxy holds TCP proxy information.
	// +optional
	TCPProxy *TCPProxy `json:"tcpproxy,omitempty"`
	// Includes allow for specific routing configuration to be included from another HTTPProxy,
	// possibly in another namespace.
	// +optional
	Includes []Include `json:"includes,omitempty"`
	// IngressClassName optionally specifies the ingress class to use for this
	// HTTPProxy. This replaces the deprecated `kubernetes.io/ingress.class`
	// annotation. For backwards compatibility, when that annotation is set, it
	// is given precedence over this field.
	// +optional
	IngressClassName string `json:"ingressClassName,omitempty"`
}

// Include describes a set of policies that can be applied to an HTTPProxy in a namespace.
type Include struct {
	// Name of the HTTPProxy
	Name string `json:"name"`
	// Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.
	// +optional
	Namespace string `json:"namespace,omitempty"`
	// Conditions are a set of rules that are applied to included HTTPProxies.
	// In effect, they are added onto the Conditions of included HTTPProxy Route
	// structs.
	// When applied, they are merged using AND, with one exception:
	// There can be only one Prefix MatchCondition per Conditions slice.
	// More than one Prefix, or contradictory Conditions, will make the
	// include invalid.
	// +optional
	Conditions []MatchCondition `json:"conditions,omitempty"`
}

// MatchCondition are a general holder for matching rules for HTTPProxies.
// One of Prefix or Header must be provided.
type MatchCondition struct {
	// Prefix defines a prefix match for a request.
	// +optional
	Prefix string `json:"prefix,omitempty"`

	// Header specifies the header condition to match.
	// +optional
	Header *HeaderMatchCondition `json:"header,omitempty"`
}

// HeaderMatchCondition specifies how to conditionally match against HTTP
// headers. The Name field is required, but only one of the remaining
// fields should be be provided.
type HeaderMatchCondition struct {
	// Name is the name of the header to match against. Name is required.
	// Header names are case insensitive.
	Name string `json:"name"`

	// Present specifies that condition is true when the named header
	// is present, regardless of its value. Note that setting Present
	// to false does not make the condition true if the named header
	// is absent.
	// +optional
	Present bool `json:"present,omitempty"`

	// NotPresent specifies that condition is true when the named header
	// is not present. Note that setting NotPresent to false does not
	// make the condition true if the named header is present.
	// +optional
	NotPresent bool `json:"notpresent,omitempty"`

	// Contains specifies a substring that must be present in
	// the header value.
	// +optional
	Contains string `json:"contains,omitempty"`

	// NotContains specifies a substring that must not be present
	// in the header value.
	// +optional
	NotContains string `json:"notcontains,omitempty"`

	// Exact specifies a string that the header value must be equal to.
	// +optional
	Exact string `json:"exact,omitempty"`

	// NoExact specifies a string that the header value must not be
	// equal to. The condition is true if the header has any other value.
	// +optional
	NotExact string `json:"notexact,omitempty"`
}

// ExtensionServiceReference names an ExtensionService resource.
type ExtensionServiceReference struct {
	// API version of the referent.
	// If this field is not specified, the default "projectcontour.io/v1alpha1" will be used
	//
	// +optional
	// +kubebuilder:validation:MinLength=1
	APIVersion string `json:"apiVersion,omitempty" protobuf:"bytes,5,opt,name=apiVersion"`

	// Namespace of the referent.
	// If this field is not specifies, the namespace of the resource that targets the referent will be used.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/
	//
	// +optional
	// +kubebuilder:validation:MinLength=1
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,2,opt,name=namespace"`

	// Name of the referent.
	//
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	//
	// +required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
}

// AuthorizationServer configures an external server to authenticate
// client requests. The external server must implement the v3 Envoy
// external authorization GRPC protocol (https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/auth/v3/external_auth.proto).
type AuthorizationServer struct {
	// ExtensionServiceRef specifies the extension resource that will authorize client requests.
	//
	// +required
	ExtensionServiceRef ExtensionServiceReference `json:"extensionRef"`

	// AuthPolicy sets a default authorization policy for client requests.
	// This policy will be used unless overridden by individual routes.
	//
	// +optional
	AuthPolicy *AuthorizationPolicy `json:"authPolicy,omitempty"`

	// ResponseTimeout configures maximum time to wait for a check response from the authorization server.
	// Timeout durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// The string "infinity" is also a valid input and specifies no timeout.
	//
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`
	ResponseTimeout string `json:"responseTimeout,omitempty"`

	// If FailOpen is true, the client request is forwarded to the upstream service
	// even if the authorization server fails to respond. This field should not be
	// set in most cases. It is intended for use only while migrating applications
	// from internal authorization to Contour external authorization.
	//
	// +optional
	FailOpen bool `json:"failOpen,omitempty"`
}

// AuthorizationPolicy modifies how client requests are authenticated.
type AuthorizationPolicy struct {
	// When true, this field disables client request authentication
	// for the scope of the policy.
	//
	// +optional
	Disabled bool `json:"disabled,omitempty"`

	// Context is a set of key/value pairs that are sent to the
	// authentication server in the check request. If a context
	// is provided at an enclosing scope, the entries are merged
	// such that the inner scope overrides matching keys from the
	// outer scope.
	//
	// +optional
	Context map[string]string `json:"context,omitempty"`
}

// VirtualHost appears at most once. If it is present, the object is considered
// to be a "root".
type VirtualHost struct {
	// The fully qualified domain name of the root of the ingress tree
	// all leaves of the DAG rooted at this object relate to the fqdn.
	Fqdn string `json:"fqdn"`

	// If present the fields describes TLS properties of the virtual
	// host. The SNI names that will be matched on are described in fqdn,
	// the tls.secretName secret must contain a certificate that itself
	// contains a name that matches the FQDN.
	//
	// +optional
	TLS *TLS `json:"tls,omitempty"`

	// This field configures an extension service to perform
	// authorization for this virtual host. Authorization can
	// only be configured on virtual hosts that have TLS enabled.
	// If the TLS configuration requires client certificate
	// validation, the client certificate is always included in the
	// authentication check request.
	//
	// +optional
	Authorization *AuthorizationServer `json:"authorization,omitempty"`
	// Specifies the cross-origin policy to apply to the VirtualHost.
	// +optional
	CORSPolicy *CORSPolicy `json:"corsPolicy,omitempty"`
	// The policy for rate limiting on the virtual host.
	// +optional
	RateLimitPolicy *RateLimitPolicy `json:"rateLimitPolicy,omitempty"`
}

// TLS describes tls properties. The SNI names that will be matched on
// are described in the HTTPProxy's Spec.VirtualHost.Fqdn field.
type TLS struct {
	// SecretName is the name of a TLS secret in the current namespace.
	// Either SecretName or Passthrough must be specified, but not both.
	// If specified, the named secret must contain a matching certificate
	// for the virtual host's FQDN.
	SecretName string `json:"secretName,omitempty"`
	// MinimumProtocolVersion is the minimum TLS version this vhost should
	// negotiate. Valid options are `1.2` (default) and `1.3`. Any other value
	// defaults to TLS 1.2.
	// +optional
	MinimumProtocolVersion string `json:"minimumProtocolVersion,omitempty"`
	// Passthrough defines whether the encrypted TLS handshake will be
	// passed through to the backing cluster. Either Passthrough or
	// SecretName must be specified, but not both.
	// +optional
	Passthrough bool `json:"passthrough,omitempty"`
	// ClientValidation defines how to verify the client certificate
	// when an external client establishes a TLS connection to Envoy.
	//
	// This setting:
	//
	// 1. Enables TLS client certificate validation.
	// 2. Specifies how the client certificate will be validated (i.e.
	//    validation required or skipped).
	//
	// Note: Setting client certificate validation to be skipped should
	// be only used in conjunction with an external authorization server that
	// performs client validation as Contour will ensure client certificates
	// are passed along.
	//
	// +optional
	ClientValidation *DownstreamValidation `json:"clientValidation,omitempty"`

	// EnableFallbackCertificate defines if the vhost should allow a default certificate to
	// be applied which handles all requests which don't match the SNI defined in this vhost.
	EnableFallbackCertificate bool `json:"enableFallbackCertificate,omitempty"`
}

// CORSHeaderValue specifies the value of the string headers returned by a cross-domain request.
// +kubebuilder:validation:Pattern="^[a-zA-Z0-9!#$%&'*+.^_`|~-]+$"
type CORSHeaderValue string

// CORSPolicy allows setting the CORS policy
type CORSPolicy struct {
	// Specifies whether the resource allows credentials.
	//  +optional
	AllowCredentials bool `json:"allowCredentials,omitempty"`
	// AllowOrigin specifies the origins that will be allowed to do CORS requests. "*" means
	// allow any origin.
	// +kubebuilder:validation:Required
	AllowOrigin []string `json:"allowOrigin"`
	// AllowMethods specifies the content for the *access-control-allow-methods* header.
	// +kubebuilder:validation:Required
	AllowMethods []CORSHeaderValue `json:"allowMethods"`
	// AllowHeaders specifies the content for the *access-control-allow-headers* header.
	//  +optional
	AllowHeaders []CORSHeaderValue `json:"allowHeaders,omitempty"`
	// ExposeHeaders Specifies the content for the *access-control-expose-headers* header.
	//  +optional
	ExposeHeaders []CORSHeaderValue `json:"exposeHeaders,omitempty"`
	// MaxAge indicates for how long the results of a preflight request can be cached.
	// MaxAge durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// Only positive values are allowed while 0 disables the cache requiring a preflight OPTIONS
	// check for all cross-origin requests.
	//  +optional
	MaxAge string `json:"maxAge,omitempty"`
}

// Route contains the set of routes for a virtual host.
type Route struct {
	// Conditions are a set of rules that are applied to a Route.
	// When applied, they are merged using AND, with one exception:
	// There can be only one Prefix MatchCondition per Conditions slice.
	// More than one Prefix, or contradictory Conditions, will make the
	// route invalid.
	// +optional
	Conditions []MatchCondition `json:"conditions,omitempty"`
	// Services are the services to proxy traffic.
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:Required
	Services []Service `json:"services"`
	// Enables websocket support for the route.
	// +optional
	EnableWebsockets bool `json:"enableWebsockets,omitempty"`
	// Allow this path to respond to insecure requests over HTTP which are normally
	// not permitted when a `virtualhost.tls` block is present.
	// +optional
	PermitInsecure bool `json:"permitInsecure,omitempty"`
	// AuthPolicy updates the authorization policy that was set
	// on the root HTTPProxy object for client requests that
	// match this route.
	// +optional
	AuthPolicy *AuthorizationPolicy `json:"authPolicy,omitempty"`
	// The timeout policy for this route.
	// +optional
	TimeoutPolicy *TimeoutPolicy `json:"timeoutPolicy,omitempty"`
	// The retry policy for this route.
	// +optional
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`
	// The health check policy for this route.
	// +optional
	HealthCheckPolicy *HTTPHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
	// The load balancing policy for this route.
	// +optional
	LoadBalancerPolicy *LoadBalancerPolicy `json:"loadBalancerPolicy,omitempty"`
	// The policy for rewriting the path of the request URL
	// after the request has been routed to a Service.
	//
	// +optional
	PathRewritePolicy *PathRewritePolicy `json:"pathRewritePolicy,omitempty"`
	// The policy for managing request headers during proxying.
	// +optional
	RequestHeadersPolicy *HeadersPolicy `json:"requestHeadersPolicy,omitempty"`
	// The policy for managing response headers during proxying.
	// Rewriting the 'Host' header is not supported.
	// +optional
	ResponseHeadersPolicy *HeadersPolicy `json:"responseHeadersPolicy,omitempty"`
	// The policy for rate limiting on the route.
	// +optional
	RateLimitPolicy *RateLimitPolicy `json:"rateLimitPolicy,omitempty"`
}

// RateLimitPolicy defines rate limiting parameters.
type RateLimitPolicy struct {
	// Local defines local rate limiting parameters, i.e. parameters
	// for rate limiting that occurs within each Envoy pod as requests
	// are handled.
	// +optional
	Local *LocalRateLimitPolicy `json:"local,omitempty"`

	// Global defines global rate limiting parameters, i.e. parameters
	// defining descriptors that are sent to an external rate limit
	// service (RLS) for a rate limit decision on each request.
	// +optional
	Global *GlobalRateLimitPolicy `json:"global,omitempty"`
}

// LocalRateLimitPolicy defines local rate limiting parameters.
type LocalRateLimitPolicy struct {
	// Requests defines how many requests per unit of time should
	// be allowed before rate limiting occurs.
	// +required
	// +kubebuilder:validation:Minimum=1
	Requests uint32 `json:"requests"`

	// Unit defines the period of time within which requests
	// over the limit will be rate limited. Valid values are
	// "second", "minute" and "hour".
	// +kubebuilder:validation:Enum=second;minute;hour
	// +required
	Unit string `json:"unit"`

	// Burst defines the number of requests above the requests per
	// unit that should be allowed within a short period of time.
	// +optional
	Burst uint32 `json:"burst,omitempty"`

	// ResponseStatusCode is the HTTP status code to use for responses
	// to rate-limited requests. Codes must be in the 400-599 range
	// (inclusive). If not specified, the Envoy default of 429 (Too
	// Many Requests) is used.
	// +optional
	// +kubebuilder:validation:Minimum=400
	// +kubebuilder:validation:Maximum=599
	ResponseStatusCode uint32 `json:"responseStatusCode,omitempty"`

	// ResponseHeadersToAdd is an optional list of response headers to
	// set when a request is rate-limited.
	// +optional
	ResponseHeadersToAdd []HeaderValue `json:"responseHeadersToAdd,omitempty"`
}

// GlobalRateLimitPolicy defines global rate limiting parameters.
type GlobalRateLimitPolicy struct {
	// Descriptors defines the list of descriptors that will
	// be generated and sent to the rate limit service. Each
	// descriptor contains 1+ key-value pair entries.
	// +required
	// +kubebuilder:validation:MinItems=1
	Descriptors []RateLimitDescriptor `json:"descriptors,omitempty"`
}

// RateLimitDescriptor defines a list of key-value pair generators.
type RateLimitDescriptor struct {
	// Entries is the list of key-value pair generators.
	// +required
	// +kubebuilder:validation:MinItems=1
	Entries []RateLimitDescriptorEntry `json:"entries,omitempty"`
}

// RateLimitDescriptorEntry is a key-value pair generator. Exactly
// one field on this struct must be non-nil.
type RateLimitDescriptorEntry struct {
	// GenericKey defines a descriptor entry with a static key and value.
	// +optional
	GenericKey *GenericKeyDescriptor `json:"genericKey,omitempty"`

	// RequestHeader defines a descriptor entry that's populated only if
	// a given header is present on the request. The descriptor key is static,
	// and the descriptor value is equal to the value of the header.
	// +optional
	RequestHeader *RequestHeaderDescriptor `json:"requestHeader,omitempty"`

	// RequestHeaderValueMatch defines a descriptor entry that's populated
	// if the request's headers match a set of 1+ match criteria. The
	// descriptor key is "header_match", and the descriptor value is static.
	// +optional
	RequestHeaderValueMatch *RequestHeaderValueMatchDescriptor `json:"requestHeaderValueMatch,omitempty"`

	// RemoteAddress defines a descriptor entry with a key of "remote_address"
	// and a value equal to the client's IP address (from x-forwarded-for).
	// +optional
	RemoteAddress *RemoteAddressDescriptor `json:"remoteAddress,omitempty"`
}

// GenericKeyDescriptor defines a descriptor entry with a static key and
// value.
type GenericKeyDescriptor struct {
	// Key defines the key of the descriptor entry. If not set, the
	// key is set to "generic_key".
	// +optional
	Key string `json:"key,omitempty"`

	// Value defines the value of the descriptor entry.
	// +required
	// +kubebuilder:validation:MinLength=1
	Value string `json:"value,omitempty"`
}

// RequestHeaderDescriptor defines a descriptor entry that's populated only
// if a given header is present on the request. The value of the descriptor
// entry is equal to the value of the header (if present).
type RequestHeaderDescriptor struct {
	// HeaderName defines the name of the header to look for on the request.
	// +required
	// +kubebuilder:validation:MinLength=1
	HeaderName string `json:"headerName,omitempty"`

	// DescriptorKey defines the key to use on the descriptor entry.
	// +required
	// +kubebuilder:validation:MinLength=1
	DescriptorKey string `json:"descriptorKey,omitempty"`
}

// RequestHeaderValueMatchDescriptor defines a descriptor entry that's populated
// if the request's headers match a set of 1+ match criteria. The descriptor key
// is "header_match", and the descriptor value is statically defined.
type RequestHeaderValueMatchDescriptor struct {
	// Headers is a list of 1+ match criteria to apply against the request
	// to determine whether to populate the descriptor entry or not.
	// +kubebuilder:validation:MinItems=1
	Headers []HeaderMatchCondition `json:"headers,omitempty"`

	// ExpectMatch defines whether the request must positively match the match
	// criteria in order to generate a descriptor entry (i.e. true), or not
	// match the match criteria in order to generate a descriptor entry (i.e. false).
	// The default is true.
	// +kubebuilder:default=true
	ExpectMatch bool `json:"expectMatch,omitempty"`

	// Value defines the value of the descriptor entry.
	// +required
	// +kubebuilder:validation:MinLength=1
	Value string `json:"value,omitempty"`
}

// RemoteAddressDescriptor defines a descriptor entry with a key of
// "remote_address" and a value equal to the client's IP address
// (from x-forwarded-for).
type RemoteAddressDescriptor struct{}

// TCPProxy contains the set of services to proxy TCP connections.
type TCPProxy struct {
	// The load balancing policy for the backend services. Note that the
	// `Cookie` and `RequestHash` load balancing strategies cannot be used
	// here.
	// +optional
	LoadBalancerPolicy *LoadBalancerPolicy `json:"loadBalancerPolicy,omitempty"`
	// Services are the services to proxy traffic
	// +optional
	Services []Service `json:"services"`
	// Include specifies that this tcpproxy should be delegated to another HTTPProxy.
	// +optional
	Include *TCPProxyInclude `json:"include,omitempty"`
	// IncludesDeprecated allow for specific routing configuration to be appended to another HTTPProxy in another namespace.
	//
	// Exists due to a mistake when developing HTTPProxy and the field was marked plural
	// when it should have been singular. This field should stay to not break backwards compatibility to v1 users.
	// +optional
	IncludesDeprecated *TCPProxyInclude `json:"includes,omitempty"`
	// The health check policy for this tcp proxy
	// +optional
	HealthCheckPolicy *TCPHealthCheckPolicy `json:"healthCheckPolicy,omitempty"`
}

// TCPProxyInclude describes a target HTTPProxy document which contains the TCPProxy details.
type TCPProxyInclude struct {
	// Name of the child HTTPProxy
	Name string `json:"name"`
	// Namespace of the HTTPProxy to include. Defaults to the current namespace if not supplied.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// Service defines an Kubernetes Service to proxy traffic.
type Service struct {
	// Name is the name of Kubernetes service to proxy traffic.
	// Names defined here will be used to look up corresponding endpoints which contain the ips to route.
	Name string `json:"name"`
	// Port (defined as Integer) to proxy traffic to since a service can have multiple defined.
	//
	// +required
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65536
	// +kubebuilder:validation:ExclusiveMinimum=false
	// +kubebuilder:validation:ExclusiveMaximum=true
	Port int `json:"port"`
	// Protocol may be used to specify (or override) the protocol used to reach this Service.
	// Values may be tls, h2, h2c. If omitted, protocol-selection falls back on Service annotations.
	// +kubebuilder:validation:Enum=h2;h2c;tls
	// +optional
	Protocol *string `json:"protocol,omitempty"`
	// Weight defines percentage of traffic to balance traffic
	// +optional
	// +kubebuilder:validation:Minimum=0
	Weight int64 `json:"weight,omitempty"`
	// UpstreamValidation defines how to verify the backend service's certificate
	// +optional
	UpstreamValidation *UpstreamValidation `json:"validation,omitempty"`
	// If Mirror is true the Service will receive a read only mirror of the traffic for this route.
	Mirror bool `json:"mirror,omitempty"`
	// The policy for managing request headers during proxying.
	// Rewriting the 'Host' header is not supported.
	// +optional
	RequestHeadersPolicy *HeadersPolicy `json:"requestHeadersPolicy,omitempty"`
	// The policy for managing response headers during proxying.
	// Rewriting the 'Host' header is not supported.
	// +optional
	ResponseHeadersPolicy *HeadersPolicy `json:"responseHeadersPolicy,omitempty"`
}

// HTTPHealthCheckPolicy defines health checks on the upstream service.
type HTTPHealthCheckPolicy struct {
	// HTTP endpoint used to perform health checks on upstream service
	Path string `json:"path"`
	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the name "contour-envoy-healthcheck"
	// will be used.
	Host string `json:"host,omitempty"`
	// The interval (seconds) between health checks
	// +optional
	IntervalSeconds int64 `json:"intervalSeconds"`
	// The time to wait (seconds) for a health check response
	// +optional
	TimeoutSeconds int64 `json:"timeoutSeconds"`
	// The number of unhealthy health checks required before a host is marked unhealthy
	// +optional
	// +kubebuilder:validation:Minimum=0
	UnhealthyThresholdCount int64 `json:"unhealthyThresholdCount"`
	// The number of healthy health checks required before a host is marked healthy
	// +optional
	// +kubebuilder:validation:Minimum=0
	HealthyThresholdCount int64 `json:"healthyThresholdCount"`
}

// TCPHealthCheckPolicy defines health checks on the upstream service.
type TCPHealthCheckPolicy struct {
	// The interval (seconds) between health checks
	// +optional
	IntervalSeconds int64 `json:"intervalSeconds"`
	// The time to wait (seconds) for a health check response
	// +optional
	TimeoutSeconds int64 `json:"timeoutSeconds"`
	// The number of unhealthy health checks required before a host is marked unhealthy
	// +optional
	UnhealthyThresholdCount uint32 `json:"unhealthyThresholdCount"`
	// The number of healthy health checks required before a host is marked healthy
	// +optional
	HealthyThresholdCount uint32 `json:"healthyThresholdCount"`
}

// TimeoutPolicy configures timeouts that are used for handling network requests.
//
// TimeoutPolicy durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
// The string "infinity" is also a valid input and specifies no timeout.
// A value of "0s" will be treated as if the field were not set, i.e. by using Envoy's default behavior.
//
// Example input values: "300ms", "5s", "1m".
type TimeoutPolicy struct {
	// Timeout for receiving a response from the server after processing a request from client.
	// If not supplied, Envoy's default value of 15s applies.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`
	Response string `json:"response,omitempty"`

	// Timeout after which, if there are no active requests for this route, the connection between
	// Envoy and the backend or Envoy and the external client will be closed.
	// If not specified, there is no per-route idle timeout, though a connection manager-wide
	// stream_idle_timeout default of 5m still applies.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`
	Idle string `json:"idle,omitempty"`
}

// RetryOn is a string type alias with validation to ensure that the value is valid.
// +kubebuilder:validation:Enum="5xx";gateway-error;reset;connect-failure;retriable-4xx;refused-stream;retriable-status-codes;retriable-headers;cancelled;deadline-exceeded;internal;resource-exhausted;unavailable
type RetryOn string

// RetryPolicy defines the attributes associated with retrying policy.
type RetryPolicy struct {
	// NumRetries is maximum allowed number of retries.
	// If not supplied, the number of retries is one.
	// +optional
	// +kubebuilder:validation:Minimum=0
	NumRetries int64 `json:"count"`
	// PerTryTimeout specifies the timeout per retry attempt.
	// Ignored if NumRetries is not supplied.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`
	PerTryTimeout string `json:"perTryTimeout,omitempty"`
	// RetryOn specifies the conditions on which to retry a request.
	//
	// Supported [HTTP conditions](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-on):
	//
	// - `5xx`
	// - `gateway-error`
	// - `reset`
	// - `connect-failure`
	// - `retriable-4xx`
	// - `refused-stream`
	// - `retriable-status-codes`
	// - `retriable-headers`
	//
	// Supported [gRPC conditions](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/router_filter#x-envoy-retry-grpc-on):
	//
	// - `cancelled`
	// - `deadline-exceeded`
	// - `internal`
	// - `resource-exhausted`
	// - `unavailable`
	// +optional
	RetryOn []RetryOn `json:"retryOn,omitempty"`
	// RetriableStatusCodes specifies the HTTP status codes that should be retried.
	//
	// This field is only respected when you include `retriable-status-codes` in the `RetryOn` field.
	// +optional
	RetriableStatusCodes []uint32 `json:"retriableStatusCodes,omitempty"`
}

// ReplacePrefix describes a path prefix replacement.
type ReplacePrefix struct {
	// Prefix specifies the URL path prefix to be replaced.
	//
	// If Prefix is specified, it must exactly match the MatchCondition
	// prefix that is rendered by the chain of including HTTPProxies
	// and only that path prefix will be replaced by Replacement.
	// This allows HTTPProxies that are included through multiple
	// roots to only replace specific path prefixes, leaving others
	// unmodified.
	//
	// If Prefix is not specified, all routing prefixes rendered
	// by the include chain will be replaced.
	//
	// +optional
	// +kubebuilder:validation:MinLength=1
	Prefix string `json:"prefix,omitempty"`

	// Replacement is the string that the routing path prefix
	// will be replaced with. This must not be empty.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Replacement string `json:"replacement"`
}

// PathRewritePolicy specifies how a request URL path should be
// rewritten. This rewriting takes place after a request is routed
// and has no subsequent effects on the proxy's routing decision.
// No HTTP headers or body content is rewritten.
//
// Exactly one field in this struct may be specified.
type PathRewritePolicy struct {
	// ReplacePrefix describes how the path prefix should be replaced.
	// +optional
	ReplacePrefix []ReplacePrefix `json:"replacePrefix,omitempty"`
}

// HeaderHashOptions contains options to configure a HTTP request header hash
// policy, used in request attribute hash based load balancing.
type HeaderHashOptions struct {
	// HeaderName is the name of the HTTP request header that will be used to
	// calculate the hash key. If the header specified is not present on a
	// request, no hash will be produced.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	HeaderName string `json:"headerName,omitempty"`
}

// RequestHashPolicy contains configuration for an individual hash policy
// on a request attribute.
type RequestHashPolicy struct {
	// Terminal is a flag that allows for short-circuiting computing of a hash
	// for a given request. If set to true, and the request attribute specified
	// in the attribute hash options is present, no further hash policies will
	// be used to calculate a hash for the request.
	Terminal bool `json:"terminal,omitempty"`

	// HeaderHashOptions should be set when request header hash based load
	// balancing is desired. It must be the only hash option field set,
	// otherwise this request hash policy object will be ignored.
	// +kubebuilder:validation:Required
	HeaderHashOptions *HeaderHashOptions `json:"headerHashOptions,omitempty"`
}

// LoadBalancerPolicy defines the load balancing policy.
type LoadBalancerPolicy struct {
	// Strategy specifies the policy used to balance requests
	// across the pool of backend pods. Valid policy names are
	// `Random`, `RoundRobin`, `WeightedLeastRequest`, `Cookie`,
	// and `RequestHash`. If an unknown strategy name is specified
	// or no policy is supplied, the default `RoundRobin` policy
	// is used.
	Strategy string `json:"strategy,omitempty"`

	// RequestHashPolicies contains a list of hash policies to apply when the
	// `RequestHash` load balancing strategy is chosen. If an element of the
	// supplied list of hash policies is invalid, it will be ignored. If the
	// list of hash policies is empty after validation, the load balancing
	// strategy will fall back the the default `RoundRobin`.
	RequestHashPolicies []RequestHashPolicy `json:"requestHashPolicies,omitempty"`
}

// HeadersPolicy defines how headers are managed during forwarding.
// The `Host` header is treated specially and if set in a HTTP response
// will be used as the SNI server name when forwarding over TLS. It is an
// error to attempt to set the `Host` header in a HTTP response.
type HeadersPolicy struct {
	// Set specifies a list of HTTP header values that will be set in the HTTP header.
	// If the header does not exist it will be added, otherwise it will be overwritten with the new value.
	// +optional
	Set []HeaderValue `json:"set,omitempty"`
	// Remove specifies a list of HTTP header names to remove.
	// +optional
	Remove []string `json:"remove,omitempty"`
}

// HeaderValue represents a header name/value pair
type HeaderValue struct {
	// Name represents a key of a header
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`
	// Value represents the value of a header specified by a key
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Value string `json:"value"`
}

// UpstreamValidation defines how to verify the backend service's certificate
type UpstreamValidation struct {
	// Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend
	CACertificate string `json:"caSecret"`
	// Key which is expected to be present in the 'subjectAltName' of the presented certificate
	SubjectName string `json:"subjectName"`
}

// DownstreamValidation defines how to verify the client certificate.
type DownstreamValidation struct {
	// Name of a Kubernetes secret that contains a CA certificate bundle.
	// The client certificate must validate against the certificates in the bundle.
	// If specified and SkipClientCertValidation is true, client certificates will
	// be required on requests.
	// +optional
	// +kubebuilder:validation:MinLength=1
	CACertificate string `json:"caSecret,omitempty"`

	// SkipClientCertValidation disables downstream client certificate
	// validation. Defaults to false. This field is intended to be used in
	// conjunction with external authorization in order to enable the external
	// authorization server to validate client certificates. When this field
	// is set to true, client certificates are requested but not verified by
	// Envoy. If CACertificate is specified, client certificates are required on
	// requests, but not verified. If external authorization is in use, they are
	// presented to the external authorization server.
	// +optional
	SkipClientCertValidation bool `json:"skipClientCertValidation"`
}

// HTTPProxyStatus reports the current state of the HTTPProxy.
type HTTPProxyStatus struct {
	// +optional
	CurrentStatus string `json:"currentStatus,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// LoadBalancer contains the current status of the load balancer.
	LoadBalancer corev1.LoadBalancerStatus `json:"loadBalancer,omitempty"`
	// +optional
	// Conditions contains information about the current status of the HTTPProxy,
	// in an upstream-friendly container.
	//
	// Contour will update a single condition, `Valid`, that is in normal-true polarity.
	// That is, when `currentStatus` is `valid`, the `Valid` condition will be `status: true`,
	// and vice versa.
	//
	// Contour will leave untouched any other Conditions set in this block,
	// in case some other controller wants to add a Condition.
	//
	// If you are another controller owner and wish to add a condition, you *should*
	// namespace your condition with a label, like `controller.domain.com/ConditionName`.
	// +patchMergeKey=type
	// +patchStrategy=merge
	// +listType=map
	// +listMapKey=type
	Conditions []DetailedCondition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPProxy is an Ingress CRD specification.
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="FQDN",type="string",JSONPath=".spec.virtualhost.fqdn",description="Fully qualified domain name"
// +kubebuilder:printcolumn:name="TLS Secret",type="string",JSONPath=".spec.virtualhost.tls.secretName",description="Secret with TLS credentials"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.currentStatus",description="The current status of the HTTPProxy"
// +kubebuilder:printcolumn:name="Status Description",type="string",JSONPath=".status.description",description="Description of the current status"
// +kubebuilder:resource:scope=Namespaced,path=httpproxies,shortName=proxy;proxies,singular=httpproxy
// +kubebuilder:subresource:status
type HTTPProxy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec HTTPProxySpec `json:"spec"`
	// Status is a container for computed information about the HTTPProxy.
	// +optional
	Status HTTPProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPProxyList is a list of HTTPProxies.
type HTTPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []HTTPProxy `json:"items"`
}
