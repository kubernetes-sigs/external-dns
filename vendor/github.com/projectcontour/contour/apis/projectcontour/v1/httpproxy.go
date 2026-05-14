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
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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

// Namespace refers to a Kubernetes namespace. It must be a RFC 1123 label.
//
// This validation is based off of the corresponding Kubernetes validation:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/util/validation/validation.go#L187
//
// This is used for Namespace name validation here:
// https://github.com/kubernetes/apimachinery/blob/02cfb53916346d085a6c6c7c66f882e3c6b0eca6/pkg/api/validation/generic.go#L63
//
// Valid values include:
//
// * "example"
//
// Invalid values include:
//
// * "example.com" - "." is an invalid character
//
// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?$`
// +kubebuilder:validation:MinLength=1
// +kubebuilder:validation:MaxLength=63
type Namespace string

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
	// include invalid. Exact and Regex match conditions are not allowed
	// on includes.
	// +optional
	Conditions []MatchCondition `json:"conditions,omitempty"`
}

// MatchCondition are a general holder for matching rules for HTTPProxies.
// One of Prefix, Exact, Regex, Header or QueryParameter must be provided.
type MatchCondition struct {
	// Prefix defines a prefix match for a request.
	// +optional
	Prefix string `json:"prefix,omitempty"`

	// Exact defines a exact match for a request.
	// This field is not allowed in include match conditions.
	// +optional
	Exact string `json:"exact,omitempty"`

	// Regex defines a regex match for a request.
	// This field is not allowed in include match conditions.
	// +optional
	Regex string `json:"regex,omitempty"`

	// Header specifies the header condition to match.
	// +optional
	Header *HeaderMatchCondition `json:"header,omitempty"`

	// QueryParameter specifies the query parameter condition to match.
	// +optional
	QueryParameter *QueryParameterMatchCondition `json:"queryParameter,omitempty"`
}

// HeaderMatchCondition specifies how to conditionally match against HTTP
// headers. The Name field is required, only one of Present, NotPresent,
// Contains, NotContains, Exact, NotExact and Regex can be set.
// For negative matching rules only (e.g. NotContains or NotExact) you can set
// TreatMissingAsEmpty.
// IgnoreCase has no effect for Regex.
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

	// IgnoreCase specifies that string matching should be case insensitive.
	// Note that this has no effect on the Regex parameter.
	// +optional
	IgnoreCase bool `json:"ignoreCase,omitempty"`

	// Exact specifies a string that the header value must be equal to.
	// +optional
	Exact string `json:"exact,omitempty"`

	// NoExact specifies a string that the header value must not be
	// equal to. The condition is true if the header has any other value.
	// +optional
	NotExact string `json:"notexact,omitempty"`

	// Regex specifies a regular expression pattern that must match the header
	// value.
	// +optional
	Regex string `json:"regex,omitempty"`

	// TreatMissingAsEmpty specifies if the header match rule specified header
	// does not exist, this header value will be treated as empty. Defaults to false.
	// Unlike the underlying Envoy implementation this is **only** supported for
	// negative matches (e.g. NotContains, NotExact).
	// +optional
	TreatMissingAsEmpty bool `json:"treatMissingAsEmpty,omitempty"`
}

// QueryParameterMatchCondition specifies how to conditionally match against HTTP
// query parameters. The Name field is required, only one of Exact, Prefix,
// Suffix, Regex, Contains and Present can be set. IgnoreCase has no effect
// for Regex.
type QueryParameterMatchCondition struct {
	// Name is the name of the query parameter to match against. Name is required.
	// Query parameter names are case insensitive.
	Name string `json:"name"`

	// Exact specifies a string that the query parameter value must be equal to.
	// +optional
	Exact string `json:"exact,omitempty"`

	// Prefix defines a prefix match for the query parameter value.
	// +optional
	Prefix string `json:"prefix,omitempty"`

	// Suffix defines a suffix match for a query parameter value.
	// +optional
	Suffix string `json:"suffix,omitempty"`

	// Regex specifies a regular expression pattern that must match the query
	// parameter value.
	// +optional
	Regex string `json:"regex,omitempty"`

	// Contains specifies a substring that must be present in
	// the query parameter value.
	// +optional
	Contains string `json:"contains,omitempty"`

	// IgnoreCase specifies that string matching should be case insensitive.
	// Note that this has no effect on the Regex parameter.
	// +optional
	IgnoreCase bool `json:"ignoreCase,omitempty"`

	// Present specifies that condition is true when the named query parameter
	// is present, regardless of its value. Note that setting Present
	// to false does not make the condition true if the named query parameter
	// is absent.
	// +optional
	Present bool `json:"present,omitempty"`
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
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name,omitempty" protobuf:"bytes,3,opt,name=name"`
}

// AuthorizationServer configures an external server to authenticate
// client requests. The external server must implement the v3 Envoy
// external authorization GRPC protocol (https://www.envoyproxy.io/docs/envoy/latest/api-v3/service/auth/v3/external_auth.proto).
type AuthorizationServer struct {
	// ExtensionServiceRef specifies the extension resource that will authorize client requests.
	//
	// +optional
	ExtensionServiceRef ExtensionServiceReference `json:"extensionRef,omitempty"`

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

	// WithRequestBody specifies configuration for sending the client request's body to authorization server.
	// +optional
	WithRequestBody *AuthorizationServerBufferSettings `json:"withRequestBody,omitempty"`
}

// AuthorizationServerBufferSettings enables ExtAuthz filter to buffer client request data and send it as part of authorization request
type AuthorizationServerBufferSettings struct {
	// MaxRequestBytes sets the maximum size of message body ExtAuthz filter will hold in-memory.
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:default=1024
	MaxRequestBytes uint32 `json:"maxRequestBytes,omitempty"`

	// If AllowPartialMessage is true, then Envoy will buffer the body until MaxRequestBytes are reached.
	// +optional
	AllowPartialMessage bool `json:"allowPartialMessage,omitempty"`

	// If PackAsBytes is true, the body sent to Authorization Server is in raw bytes.
	// +optional
	PackAsBytes bool `json:"packAsBytes,omitempty"`
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
	//
	// +kubebuilder:validation:Pattern="^(\\*\\.)?[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$"
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
	// Providers to use for verifying JSON Web Tokens (JWTs) on the virtual host.
	// +optional
	JWTProviders []JWTProvider `json:"jwtProviders,omitempty"`

	// IPAllowFilterPolicy is a list of ipv4/6 filter rules for which matching
	// requests should be allowed. All other requests will be denied.
	// Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined.
	// The rules defined here may be overridden in a Route.
	IPAllowFilterPolicy []IPFilterPolicy `json:"ipAllowPolicy,omitempty"`

	// IPDenyFilterPolicy is a list of ipv4/6 filter rules for which matching
	// requests should be denied. All other requests will be allowed.
	// Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined.
	// The rules defined here may be overridden in a Route.
	IPDenyFilterPolicy []IPFilterPolicy `json:"ipDenyPolicy,omitempty"`
}

// JWTProvider defines how to verify JWTs on requests.
type JWTProvider struct {
	// Unique name for the provider.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	Name string `json:"name"`

	// Whether the provider should apply to all
	// routes in the HTTPProxy/its includes by
	// default. At most one provider can be marked
	// as the default. If no provider is marked
	// as the default, individual routes must explicitly
	// identify the provider they require.
	// +optional
	Default bool `json:"default,omitempty"`

	// Issuer that JWTs are required to have in the "iss" field.
	// If not provided, JWT issuers are not checked.
	// +optional
	Issuer string `json:"issuer,omitempty"`

	// Audiences that JWTs are allowed to have in the "aud" field.
	// If not provided, JWT audiences are not checked.
	// +optional
	Audiences []string `json:"audiences,omitempty"`

	// Remote JWKS to use for verifying JWT signatures.
	// +kubebuilder:validation:Required
	RemoteJWKS RemoteJWKS `json:"remoteJWKS"`

	// Whether the JWT should be forwarded to the backend
	// service after successful verification. By default,
	// the JWT is not forwarded.
	// +optional
	ForwardJWT bool `json:"forwardJWT,omitempty"`
}

// RemoteJWKS defines how to fetch a JWKS from an HTTP endpoint.
type RemoteJWKS struct {
	// The URI for the JWKS.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	URI string `json:"uri"`

	// UpstreamValidation defines how to verify the JWKS's TLS certificate.
	// +optional
	UpstreamValidation *UpstreamValidation `json:"validation,omitempty"`

	// How long to wait for a response from the URI.
	// If not specified, a default of 1s applies.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+)$`
	Timeout string `json:"timeout,omitempty"`

	// How long to cache the JWKS locally. If not specified,
	// Envoy's default of 5m applies.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+)$`
	CacheDuration string `json:"cacheDuration,omitempty"`

	// The DNS IP address resolution policy for the JWKS URI.
	// When configured as "v4", the DNS resolver will only perform a lookup
	// for addresses in the IPv4 family. If "v6" is configured, the DNS resolver
	// will only perform a lookup for addresses in the IPv6 family.
	// If "all" is configured, the DNS resolver
	// will perform a lookup for addresses in both the IPv4 and IPv6 family.
	// If "auto" is configured, the DNS resolver will first perform a lookup
	// for addresses in the IPv6 family and fallback to a lookup for addresses
	// in the IPv4 family. If not specified, the Contour-wide setting defined
	// in the config file or ContourConfiguration applies (defaults to "auto").
	//
	// See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html#envoy-v3-api-enum-config-cluster-v3-cluster-dnslookupfamily
	// for more information.
	// +optional
	// +kubebuilder:validation:Enum=auto;v4;v6
	DNSLookupFamily string `json:"dnsLookupFamily,omitempty"`
}

// TLS describes tls properties. The SNI names that will be matched on
// are described in the HTTPProxy's Spec.VirtualHost.Fqdn field.
type TLS struct {
	// SecretName is the name of a TLS secret.
	// Either SecretName or Passthrough must be specified, but not both.
	// If specified, the named secret must contain a matching certificate
	// for the virtual host's FQDN.
	// The name can be optionally prefixed with namespace "namespace/name".
	// When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.
	SecretName string `json:"secretName,omitempty"`
	// MinimumProtocolVersion is the minimum TLS version this vhost should
	// negotiate. Valid options are `1.2` (default) and `1.3`. Any other value
	// defaults to TLS 1.2.
	// +optional
	MinimumProtocolVersion string `json:"minimumProtocolVersion,omitempty"`
	// MaximumProtocolVersion is the maximum TLS version this vhost should
	// negotiate. Valid options are `1.2` and `1.3` (default). Any other value
	// defaults to TLS 1.3.
	// +optional
	MaximumProtocolVersion string `json:"maximumProtocolVersion,omitempty"`
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
	// +optional
	AllowCredentials bool `json:"allowCredentials,omitempty"`
	// AllowOrigin specifies the origins that will be allowed to do CORS requests.
	// Allowed values include "*" which signifies any origin is allowed, an exact
	// origin of the form "scheme://host[:port]" (where port is optional), or a valid
	// regex pattern.
	// Note that regex patterns are validated and a simple "glob" pattern (e.g. *.foo.com)
	// will be rejected or produce unexpected matches when applied as a regex.
	//
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	AllowOrigin []string `json:"allowOrigin"`
	// AllowMethods specifies the content for the *access-control-allow-methods* header.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	AllowMethods []CORSHeaderValue `json:"allowMethods"`
	// AllowHeaders specifies the content for the *access-control-allow-headers* header.
	// +optional
	// +kubebuilder:validation:MinItems=1
	AllowHeaders []CORSHeaderValue `json:"allowHeaders,omitempty"`
	// ExposeHeaders Specifies the content for the *access-control-expose-headers* header.
	// +optional
	// +kubebuilder:validation:MinItems=1
	ExposeHeaders []CORSHeaderValue `json:"exposeHeaders,omitempty"`
	// MaxAge indicates for how long the results of a preflight request can be cached.
	// MaxAge durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// Only positive values are allowed while 0 disables the cache requiring a preflight OPTIONS
	// check for all cross-origin requests.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|0)$`
	MaxAge string `json:"maxAge,omitempty"`
	// AllowPrivateNetwork specifies whether to allow private network requests.
	// See https://developer.chrome.com/blog/private-network-access-preflight.
	AllowPrivateNetwork bool `json:"allowPrivateNetwork,omitempty"`
}

// Route contains the set of routes for a virtual host.
type Route struct {
	// Conditions are a set of rules that are applied to a Route.
	// When applied, they are merged using AND, with one exception:
	// There can be only one Prefix, Exact or Regex MatchCondition
	// per Conditions slice. More than one of these condition types,
	// or contradictory Conditions, will make the route invalid.
	// +optional
	Conditions []MatchCondition `json:"conditions,omitempty"`
	// Services are the services to proxy traffic.
	// +optional
	Services []Service `json:"services,omitempty"`
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
	//
	// You may dynamically rewrite the Host header to be forwarded
	// upstream to the content of a request header using
	// the below format "%REQ(X-Header-Name)%". If the value of the header
	// is empty, it is ignored.
	//
	// *NOTE: Pay attention to the potential security implications of using this option.
	// Provided header must come from trusted source.
	//
	// **NOTE: The header rewrite is only done while forwarding and has no bearing
	// on the routing decision.
	//
	// +optional
	RequestHeadersPolicy *HeadersPolicy `json:"requestHeadersPolicy,omitempty"`
	// The policy for managing response headers during proxying.
	// Rewriting the 'Host' header is not supported.
	// +optional
	ResponseHeadersPolicy *HeadersPolicy `json:"responseHeadersPolicy,omitempty"`
	// The policies for rewriting Set-Cookie header attributes. Note that
	// rewritten cookie names must be unique in this list. Order rewrite
	// policies are specified in does not matter.
	// +optional
	CookieRewritePolicies []CookieRewritePolicy `json:"cookieRewritePolicies,omitempty"`
	// The policy for rate limiting on the route.
	// +optional
	RateLimitPolicy *RateLimitPolicy `json:"rateLimitPolicy,omitempty"`

	// RequestRedirectPolicy defines an HTTP redirection.
	// +optional
	RequestRedirectPolicy *HTTPRequestRedirectPolicy `json:"requestRedirectPolicy,omitempty"`

	// DirectResponsePolicy returns an arbitrary HTTP response directly.
	// +optional
	DirectResponsePolicy *HTTPDirectResponsePolicy `json:"directResponsePolicy,omitempty"`

	// The policy to define when to handle redirects responses internally.
	// +optional
	InternalRedirectPolicy *HTTPInternalRedirectPolicy `json:"internalRedirectPolicy,omitempty"`

	// The policy for verifying JWTs for requests to this route.
	// +optional
	JWTVerificationPolicy *JWTVerificationPolicy `json:"jwtVerificationPolicy,omitempty"`

	// IPAllowFilterPolicy is a list of ipv4/6 filter rules for which matching
	// requests should be allowed. All other requests will be denied.
	// Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined.
	// The rules defined here override any rules set on the root HTTPProxy.
	IPAllowFilterPolicy []IPFilterPolicy `json:"ipAllowPolicy,omitempty"`

	// IPDenyFilterPolicy is a list of ipv4/6 filter rules for which matching
	// requests should be denied. All other requests will be allowed.
	// Only one of IPAllowFilterPolicy and IPDenyFilterPolicy can be defined.
	// The rules defined here override any rules set on the root HTTPProxy.
	IPDenyFilterPolicy []IPFilterPolicy `json:"ipDenyPolicy,omitempty"`
}

type JWTVerificationPolicy struct {
	// Require names a specific JWT provider (defined in the virtual host)
	// to require for the route. If specified, this field overrides the
	// default provider if one exists. If this field is not specified,
	// the default provider will be required if one exists. At most one of
	// this field or the "disabled" field can be specified.
	// +optional
	Require string `json:"require,omitempty"`

	// Disabled defines whether to disable all JWT verification for this
	// route. This can be used to opt specific routes out of the default
	// JWT provider for the HTTPProxy. At most one of this field or the
	// "require" field can be specified.
	// +optional
	Disabled bool `json:"disabled,omitempty"`
}

// IPFilterSource indicates which IP should be considered for filtering
// +kubebuilder:validation:Enum=Peer;Remote
type IPFilterSource string

const (
	IPFilterSourcePeer   IPFilterSource = "Peer"
	IPFilterSourceRemote IPFilterSource = "Remote"
)

type IPFilterPolicy struct {
	// Source indicates how to determine the ip address to filter on, and can be
	// one of two values:
	//  - `Remote` filters on the ip address of the client, accounting for PROXY and
	//    X-Forwarded-For as needed.
	//  - `Peer` filters on the ip of the network request, ignoring PROXY and
	//    X-Forwarded-For.
	Source IPFilterSource `json:"source"`

	// CIDR is a CIDR block of ipv4 or ipv6 addresses to filter on. This can also be
	// a bare IP address (without a mask) to filter on exactly one address.
	CIDR string `json:"cidr"`
}

type HTTPDirectResponsePolicy struct {
	// StatusCode is the HTTP response status to be returned.
	// +required
	// +kubebuilder:validation:Minimum=200
	// +kubebuilder:validation:Maximum=599
	StatusCode int `json:"statusCode"`

	// Body is the content of the response body.
	// If this setting is omitted, no body is included in the generated response.
	//
	// Note: Body is not recommended to set too long
	// otherwise it can have significant resource usage impacts.
	//
	// +optional
	Body string `json:"body,omitempty"`
}

// HTTPRequestRedirectPolicy defines configuration for redirecting a request.
type HTTPRequestRedirectPolicy struct {
	// Scheme is the scheme to be used in the value of the `Location`
	// header in the response.
	// When empty, the scheme of the request is used.
	// +optional
	// +kubebuilder:validation:Enum=http;https
	Scheme *string `json:"scheme,omitempty"`

	// Hostname is the precise hostname to be used in the value of the `Location`
	// header in the response.
	// When empty, the hostname of the request is used.
	// No wildcards are allowed.
	// +optional
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=253
	// +kubebuilder:validation:Pattern=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
	Hostname *string `json:"hostname,omitempty"`

	// Port is the port to be used in the value of the `Location`
	// header in the response.
	// When empty, port (if specified) of the request is used.
	// +optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port *int32 `json:"port,omitempty"`

	// StatusCode is the HTTP status code to be used in response.
	// +optional
	// +kubebuilder:default=302
	StatusCode *RedirectResponseCode `json:"statusCode,omitempty"`

	// Path allows for redirection to a different path from the
	// original on the request. The path must start with a
	// leading slash.
	//
	// Note: Only one of Path or Prefix can be defined.
	//
	// +optional
	// +kubebuilder:validation:Pattern=`^\/.*$`
	Path *string `json:"path,omitempty"`

	// Prefix defines the value to swap the matched prefix or path with.
	// The prefix must start with a leading slash.
	//
	// Note: Only one of Path or Prefix can be defined.
	//
	// +optional
	// +kubebuilder:validation:Pattern=`^\/.*$`
	Prefix *string `json:"prefix,omitempty"`
}

// RedirectResponseCode is a uint32 type alias with validation to ensure that the value is valid.
// +kubebuilder:validation:Enum=301;302;303;307;308
type RedirectResponseCode uint32

type HTTPInternalRedirectPolicy struct {
	// MaxInternalRedirects An internal redirect is not handled, unless the number of previous internal
	// redirects that a downstream request has encountered is lower than this value.
	// +optional
	MaxInternalRedirects uint32 `json:"maxInternalRedirects,omitempty"`

	// RedirectResponseCodes If unspecified, only 302 will be treated as internal redirect.
	// Only 301, 302, 303, 307 and 308 are valid values.
	// +optional
	RedirectResponseCodes []RedirectResponseCode `json:"redirectResponseCodes,omitempty"`

	// AllowCrossSchemeRedirect Allow internal redirect to follow a target URI with a different scheme
	// than the value of x-forwarded-proto.
	// SafeOnly allows same scheme redirect and safe cross scheme redirect, which means if the downstream
	// scheme is HTTPS, both HTTPS and HTTP redirect targets are allowed, but if the downstream scheme
	// is HTTP, only HTTP redirect targets are allowed.
	// +kubebuilder:validation:Enum=Always;Never;SafeOnly
	// +kubebuilder:default=Never
	// +optional
	AllowCrossSchemeRedirect string `json:"allowCrossSchemeRedirect,omitempty"`

	// If DenyRepeatedRouteRedirect is true, rejects redirect targets that are pointing to a route that has
	// been followed by a previous redirect from the current route.
	// +optional
	DenyRepeatedRouteRedirect bool `json:"denyRepeatedRouteRedirect,omitempty"`
}

type CookieRewritePolicy struct {
	// Name is the name of the cookie for which attributes will be rewritten.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=4096
	// +kubebuilder:validation:Pattern=`^[^()<>@,;:\\"\/[\]?={} \t\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`
	Name string `json:"name"`

	// PathRewrite enables rewriting the Set-Cookie Path element.
	// If not set, Path will not be rewritten.
	// +optional
	PathRewrite *CookiePathRewrite `json:"pathRewrite,omitempty"`

	// DomainRewrite enables rewriting the Set-Cookie Domain element.
	// If not set, Domain will not be rewritten.
	// +optional
	DomainRewrite *CookieDomainRewrite `json:"domainRewrite,omitempty"`

	// Secure enables rewriting the Set-Cookie Secure element.
	// If not set, Secure attribute will not be rewritten.
	// +optional
	Secure *bool `json:"secure,omitempty"`

	// SameSite enables rewriting the Set-Cookie SameSite element.
	// If not set, SameSite attribute will not be rewritten.
	// +optional
	// +kubebuilder:validation:Enum=Strict;Lax;None
	SameSite *string `json:"sameSite,omitempty"`
}

type CookiePathRewrite struct {
	// Value is the value to rewrite the Path attribute to.
	// For now this is required.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=4096
	// +kubebuilder:validation:Pattern=`^[^;\x7f\x00\x01\x02\x03\x04\x05\x06\x07\x08\x09\x0a\x0b\x0c\x0d\x0e\x0f\x10\x11\x12\x13\x14\x15\x16\x17\x18\x19\x1a\x1b\x1c\x1d\x1e\x1f]+$`
	Value string `json:"value"`
}

type CookieDomainRewrite struct {
	// Value is the value to rewrite the Domain attribute to.
	// For now this is required.
	// +required
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=4096
	// +kubebuilder:validation:Pattern="^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$"
	Value string `json:"value"`
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
	// Disabled configures the HTTPProxy to not use
	// the default global rate limit policy defined by the Contour configuration.
	// +optional
	Disabled bool `json:"disabled,omitempty"`

	// Descriptors defines the list of descriptors that will
	// be generated and sent to the rate limit service. Each
	// descriptor contains 1+ key-value pair entries.
	// +optional
	// +kubebuilder:validation:MinItems=1
	Descriptors []RateLimitDescriptor `json:"descriptors,omitempty" yaml:"descriptors,omitempty"`
}

// RateLimitDescriptor defines a list of key-value pair generators.
type RateLimitDescriptor struct {
	// Entries is the list of key-value pair generators.
	// +required
	// +kubebuilder:validation:MinItems=1
	Entries []RateLimitDescriptorEntry `json:"entries,omitempty" yaml:"entries,omitempty"`
}

// RateLimitDescriptorEntry is a key-value pair generator. Exactly
// one field on this struct must be non-nil.
type RateLimitDescriptorEntry struct {
	// GenericKey defines a descriptor entry with a static key and value.
	// +optional
	GenericKey *GenericKeyDescriptor `json:"genericKey,omitempty" yaml:"genericKey,omitempty"`

	// RequestHeader defines a descriptor entry that's populated only if
	// a given header is present on the request. The descriptor key is static,
	// and the descriptor value is equal to the value of the header.
	// +optional
	RequestHeader *RequestHeaderDescriptor `json:"requestHeader,omitempty" yaml:"requestHeader,omitempty"`

	// RequestHeaderValueMatch defines a descriptor entry that's populated
	// if the request's headers match a set of 1+ match criteria. The
	// descriptor key is "header_match", and the descriptor value is static.
	// +optional
	RequestHeaderValueMatch *RequestHeaderValueMatchDescriptor `json:"requestHeaderValueMatch,omitempty" yaml:"requestHeaderValueMatch,omitempty"`

	// RemoteAddress defines a descriptor entry with a key of "remote_address"
	// and a value equal to the client's IP address (from x-forwarded-for).
	// +optional
	RemoteAddress *RemoteAddressDescriptor `json:"remoteAddress,omitempty" yaml:"remoteAddress,omitempty"`
}

// GenericKeyDescriptor defines a descriptor entry with a static key and
// value.
type GenericKeyDescriptor struct {
	// Key defines the key of the descriptor entry. If not set, the
	// key is set to "generic_key".
	// +optional
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value defines the value of the descriptor entry.
	// +required
	// +kubebuilder:validation:MinLength=1
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}

// RequestHeaderDescriptor defines a descriptor entry that's populated only
// if a given header is present on the request. The value of the descriptor
// entry is equal to the value of the header (if present).
type RequestHeaderDescriptor struct {
	// HeaderName defines the name of the header to look for on the request.
	// +required
	// +kubebuilder:validation:MinLength=1
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`

	// DescriptorKey defines the key to use on the descriptor entry.
	// +required
	// +kubebuilder:validation:MinLength=1
	DescriptorKey string `json:"descriptorKey,omitempty" yaml:"descriptorKey,omitempty"`
}

// RequestHeaderValueMatchDescriptor defines a descriptor entry that's populated
// if the request's headers match a set of 1+ match criteria. The descriptor key
// is "header_match", and the descriptor value is statically defined.
type RequestHeaderValueMatchDescriptor struct {
	// Headers is a list of 1+ match criteria to apply against the request
	// to determine whether to populate the descriptor entry or not.
	// +kubebuilder:validation:MinItems=1
	Headers []HeaderMatchCondition `json:"headers,omitempty" yaml:"headers,omitempty"`

	// ExpectMatch defines whether the request must positively match the match
	// criteria in order to generate a descriptor entry (i.e. true), or not
	// match the match criteria in order to generate a descriptor entry (i.e. false).
	// The default is true.
	// +kubebuilder:default=true
	ExpectMatch bool `json:"expectMatch,omitempty" yaml:"expectMatch,omitempty"`

	// Value defines the value of the descriptor entry.
	// +required
	// +kubebuilder:validation:MinLength=1
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
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
	// HealthPort is the port for this service healthcheck.
	// If not specified, Port is used for service healthchecks.
	//
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	// +optional
	HealthPort int `json:"healthPort,omitempty"`
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
	// If Mirror is true, then fractional mirroring can be enabled by optionally setting the Weight
	// field. Legal values for Weight are 1-100. Omitting the Weight field will result in 100% mirroring.
	// NOTE: Setting Weight explicitly to 0 will unexpectedly result in 100% traffic mirroring. This
	// occurs since we cannot distinguish omitted fields from those explicitly set to their default
	// values
	Mirror bool `json:"mirror,omitempty"`
	// The policy for managing request headers during proxying.
	// +optional
	RequestHeadersPolicy *HeadersPolicy `json:"requestHeadersPolicy,omitempty"`
	// The policy for managing response headers during proxying.
	// Rewriting the 'Host' header is not supported.
	// +optional
	ResponseHeadersPolicy *HeadersPolicy `json:"responseHeadersPolicy,omitempty"`
	// The policies for rewriting Set-Cookie header attributes.
	// +optional
	CookieRewritePolicies []CookieRewritePolicy `json:"cookieRewritePolicies,omitempty"`
	// Slow start will gradually increase amount of traffic to a newly added endpoint.
	// +optional
	SlowStartPolicy *SlowStartPolicy `json:"slowStartPolicy,omitempty"`
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
	// The ranges of HTTP response statuses considered healthy. Follow half-open
	// semantics, i.e. for each range the start is inclusive and the end is exclusive.
	// Must be within the range [100,600). If not specified, only a 200 response status
	// is considered healthy.
	// +optional
	ExpectedStatuses []HTTPStatusRange `json:"expectedStatuses,omitempty"`
}

type HTTPStatusRange struct {
	// The start (inclusive) of a range of HTTP status codes.
	// +kubebuilder:validation:Minimum=100
	// +kubebuilder:validation:Maximum=599
	Start int64 `json:"start"`
	// The end (exclusive) of a range of HTTP status codes.
	// +kubebuilder:validation:Minimum=101
	// +kubebuilder:validation:Maximum=600
	End int64 `json:"end"`
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

	// Timeout for how long the proxy should wait while there is no activity during single request/response (for HTTP/1.1) or stream (for HTTP/2).
	// Timeout will not trigger while HTTP/1.1 connection is idle between two consecutive requests.
	// If not specified, there is no per-route idle timeout, though a connection manager-wide
	// stream_idle_timeout default of 5m still applies.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`
	Idle string `json:"idle,omitempty"`

	// Timeout for how long connection from the proxy to the upstream service is kept when there are no active requests.
	// If not supplied, Envoy's default value of 1h applies.
	// +optional
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`
	IdleConnection string `json:"idleConnection,omitempty"`
}

// RetryOn is a string type alias with validation to ensure that the value is valid.
// +kubebuilder:validation:Enum="5xx";gateway-error;reset;reset-before-request;connect-failure;envoy-ratelimited;retriable-4xx;refused-stream;retriable-status-codes;retriable-headers;http3-post-connect-failure;cancelled;deadline-exceeded;internal;resource-exhausted;unavailable
type RetryOn string

// RetryPolicy defines the attributes associated with retrying policy.
type RetryPolicy struct {
	// NumRetries is maximum allowed number of retries.
	// If set to -1, then retries are disabled.
	// If set to 0 or not supplied, the value is set
	// to the Envoy default of 1.
	// +optional
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=-1
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
	// - `reset-before-request`
	// - `connect-failure`
	// - `envoy-ratelimited`
	// - `retriable-4xx`
	// - `refused-stream`
	// - `retriable-status-codes`
	// - `retriable-headers`
	// - `http3-post-connect-failure`
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

// QueryParameterHashOptions contains options to configure a query parameter based hash
// policy, used in request attribute hash based load balancing.
type QueryParameterHashOptions struct {
	// ParameterName is the name of the HTTP request query parameter that will be used to
	// calculate the hash key. If the query parameter specified is not present on a
	// request, no hash will be produced.
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	ParameterName string `json:"parameterName,omitempty"`
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
	// +optional
	HeaderHashOptions *HeaderHashOptions `json:"headerHashOptions,omitempty"`

	// QueryParameterHashOptions should be set when request query parameter hash based load
	// balancing is desired. It must be the only hash option field set,
	// otherwise this request hash policy object will be ignored.
	// +optional
	QueryParameterHashOptions *QueryParameterHashOptions `json:"queryParameterHashOptions,omitempty"`

	// HashSourceIP should be set to true when request source IP hash based
	// load balancing is desired. It must be the only hash option field set,
	// otherwise this request hash policy object will be ignored.
	// +optional
	HashSourceIP bool `json:"hashSourceIP,omitempty"`
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
	// strategy will fall back to the default `RoundRobin`.
	RequestHashPolicies []RequestHashPolicy `json:"requestHashPolicies,omitempty"`
}

// HeadersPolicy defines how headers are managed during forwarding.
// The `Host` header is treated specially and if set in a HTTP request
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
// +kubebuilder:validation:XValidation:message="subjectNames[0] must equal subjectName if set",rule="has(self.subjectNames) ? self.subjectNames[0] == self.subjectName : true"
type UpstreamValidation struct {
	// Name or namespaced name of the Kubernetes secret used to validate the certificate presented by the backend.
	// The secret must contain key named ca.crt.
	// The name can be optionally prefixed with namespace "namespace/name".
	// When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.
	// Max length should be the actual max possible length of a namespaced name (63 + 253 + 1 = 317)
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=317
	CACertificate string `json:"caSecret"`
	// Key which is expected to be present in the 'subjectAltName' of the presented certificate.
	// Deprecated: migrate to using the plural field subjectNames.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:MaxLength=250
	SubjectName string `json:"subjectName"`
	// List of keys, of which at least one is expected to be present in the 'subjectAltName of the
	// presented certificate.
	// +optional
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:MaxItems=8
	SubjectNames []string `json:"subjectNames"`
}

// DownstreamValidation defines how to verify the client certificate.
type DownstreamValidation struct {
	// Name of a Kubernetes secret that contains a CA certificate bundle.
	// The secret must contain key named ca.crt.
	// The client certificate must validate against the certificates in the bundle.
	// If specified and SkipClientCertValidation is true, client certificates will
	// be required on requests.
	// The name can be optionally prefixed with namespace "namespace/name".
	// When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.
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

	// ForwardClientCertificate adds the selected data from the passed client TLS certificate
	// to the x-forwarded-client-cert header.
	// +optional
	ForwardClientCertificate *ClientCertificateDetails `json:"forwardClientCertificate,omitempty"`

	// Name of a Kubernetes opaque secret that contains a concatenated list of PEM encoded CRLs.
	// The secret must contain key named crl.pem.
	// This field will be used to verify that a client certificate has not been revoked.
	// CRLs must be available from all CAs, unless crlOnlyVerifyLeafCert is true.
	// Large CRL lists are not supported since individual secrets are limited to 1MiB in size.
	// The name can be optionally prefixed with namespace "namespace/name".
	// When cross-namespace reference is used, TLSCertificateDelegation resource must exist in the namespace to grant access to the secret.
	// +optional
	// +kubebuilder:validation:MinLength=1
	CertificateRevocationList string `json:"crlSecret,omitempty"`

	// If this option is set to true, only the certificate at the end of the
	// certificate chain will be subject to validation by CRL.
	// +optional
	OnlyVerifyLeafCertCrl bool `json:"crlOnlyVerifyLeafCert"`

	// OptionalClientCertificate when set to true will request a client certificate
	// but allow the connection to continue if the client does not provide one.
	// If a client certificate is sent, it will be verified according to the
	// other properties, which includes disabling validation if
	// SkipClientCertValidation is set. Defaults to false.
	// +optional
	OptionalClientCertificate bool `json:"optionalClientCertificate"`
}

// ClientCertificateDetails defines which parts of the client certificate will be forwarded.
type ClientCertificateDetails struct {
	// Subject of the client cert.
	// +optional
	Subject bool `json:"subject"`
	// Client cert in URL encoded PEM format.
	// +optional
	Cert bool `json:"cert"`
	// Client cert chain (including the leaf cert) in URL encoded PEM format.
	// +optional
	Chain bool `json:"chain"`
	// DNS type Subject Alternative Names of the client cert.
	// +optional
	DNS bool `json:"dns"`
	// URI type Subject Alternative Name of the client cert.
	// +optional
	URI bool `json:"uri"`
}

// HTTPProxyStatus reports the current state of the HTTPProxy.
type HTTPProxyStatus struct {
	// +optional
	CurrentStatus string `json:"currentStatus,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// LoadBalancer contains the current status of the load balancer.
	LoadBalancer core_v1.LoadBalancerStatus `json:"loadBalancer,omitempty"`
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
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`

	Spec HTTPProxySpec `json:"spec"`
	// Status is a container for computed information about the HTTPProxy.
	// +optional
	// +kubebuilder:default={currentStatus: "NotReconciled", description:"Waiting for controller"}
	Status HTTPProxyStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPProxyList is a list of HTTPProxies.
type HTTPProxyList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []HTTPProxy `json:"items"`
}

// SlowStartPolicy will gradually increase amount of traffic to a newly added endpoint.
// It can be used only with RoundRobin and WeightedLeastRequest load balancing strategies.
type SlowStartPolicy struct {
	// The duration of slow start window.
	// Duration is expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// +required
	// +kubebuilder:validation:Pattern=`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+)$`
	Window string `json:"window"`

	// The speed of traffic increase over the slow start window.
	// Defaults to 1.0, so that endpoint would get linearly increasing amount of traffic.
	// When increasing the value for this parameter, the speed of traffic ramp-up increases non-linearly.
	// The value of aggression parameter should be greater than 0.0.
	//
	// More info: https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/slow_start
	//
	// +optional
	// +kubebuilder:default=`1.0`
	// +kubebuilder:validation:Pattern=`^([0-9]+([.][0-9]+)?|[.][0-9]+)$`
	Aggression string `json:"aggression"`

	// The minimum or starting percentage of traffic to send to new endpoints.
	// A non-zero value helps avoid a too small initial weight, which may cause endpoints in slow start mode to receive no traffic in the beginning of the slow start window.
	// If not specified, the default is 10%.
	// +optional
	// +kubebuilder:default=10
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=100
	MinimumWeightPercent uint32 `json:"minWeightPercent"`
}

// +kubebuilder:validation:Enum=grpcroutes;tlsroutes;extensionservices;backendtlspolicies
type Feature string
