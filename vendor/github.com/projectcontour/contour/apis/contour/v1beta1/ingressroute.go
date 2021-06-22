// Copyright © 2019 VMware
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

package v1beta1

import (
	projcontour "github.com/projectcontour/contour/apis/projectcontour/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IngressRouteSpec defines the spec of the CRD
type IngressRouteSpec struct {
	// Virtualhost appears at most once. If it is present, the object is considered
	// to be a "root".
	// +optional
	VirtualHost *VirtualHost `json:"virtualhost,omitempty"`
	// Routes are the ingress routes. If TCPProxy is present, Routes is ignored.
	Routes []Route `json:"routes,omitempty"`
	// TCPProxy holds TCP proxy information.
	// +optional
	TCPProxy *TCPProxy `json:"tcpproxy,omitempty"`
}

// VirtualHost appears at most once. If it is present, the object is considered
// to be a "root".
type VirtualHost struct {
	// The fully qualified domain name of the root of the ingress tree
	// all leaves of the DAG rooted at this object relate to the fqdn
	Fqdn string `json:"fqdn"`
	// If present describes tls properties. The SNI names that will be matched on
	// are described in fqdn, the tls.secretName secret must contain a
	// matching certificate
	// +optional
	TLS *TLS `json:"tls,omitempty"`
}

// TLS describes tls properties. The SNI names that will be matched on
// are described in fqdn, the tls.secretName secret must contain a
// matching certificate unless tls.passthrough is set to true.
type TLS struct {
	// required, the name of a secret in the current namespace
	SecretName string `json:"secretName,omitempty"`
	// Minimum TLS version this vhost should negotiate
	// +optional
	MinimumProtocolVersion string `json:"minimumProtocolVersion,omitempty"`
	// If Passthrough is set to true, the SecretName will be ignored
	// and the encrypted handshake will be passed through to the
	// backing cluster.
	// +optional
	Passthrough bool `json:"passthrough,omitempty"`
}

// Route contains the set of routes for a virtual host
type Route struct {
	// Match defines the prefix match
	Match string `json:"match"`
	// Services are the services to proxy traffic
	// +optional
	Services []Service `json:"services,omitempty"`
	// Delegate specifies that this route should be delegated to another IngressRoute
	// +optional
	Delegate *Delegate `json:"delegate,omitempty"`
	// Enables websocket support for the route
	// +optional
	EnableWebsockets bool `json:"enableWebsockets,omitempty"`
	// Allow this path to respond to insecure requests over HTTP which are normally
	// not permitted when a `virtualhost.tls` block is present.
	// +optional
	PermitInsecure bool `json:"permitInsecure,omitempty"`
	// Indicates that during forwarding, the matched prefix (or path) should be swapped with this value
	// +optional
	PrefixRewrite string `json:"prefixRewrite,omitempty"`
	// The timeout policy for this route
	// +optional
	TimeoutPolicy *TimeoutPolicy `json:"timeoutPolicy,omitempty"`
	// The retry policy for this route
	// +optional
	RetryPolicy *projcontour.RetryPolicy `json:"retryPolicy,omitempty"`
}

// TimeoutPolicy define the attributes associated with timeout
type TimeoutPolicy struct {
	// Timeout for receiving a response from the server after processing a request from client.
	// If not supplied the timeout duration is undefined.
	// +optional
	Request string `json:"request"`
}

// TCPProxy contains the set of services to proxy TCP connections.
type TCPProxy struct {
	// Services are the services to proxy traffic
	// +optional
	Services []Service `json:"services,omitempty"`
	// Delegate specifies that this tcpproxy should be delegated to another IngressRoute
	// +optional
	Delegate *Delegate `json:"delegate,omitempty"`
}

// Service defines an upstream to proxy traffic to
type Service struct {
	// Name is the name of Kubernetes service to proxy traffic.
	// Names defined here will be used to look up corresponding endpoints which contain the ips to route.
	Name string `json:"name"`
	// Port (defined as Integer) to proxy traffic to since a service can have multiple defined
	Port int `json:"port"`
	// Weight defines percentage of traffic to balance traffic
	// +optional
	// +kubebuilder:validation:Minimum=0
	Weight int64 `json:"weight,omitempty"`
	// HealthCheck defines optional healthchecks on the upstream service
	// +optional
	HealthCheck *HealthCheck `json:"healthCheck,omitempty"`
	// LB Algorithm to apply (see https://github.com/projectcontour/contour/blob/master/design/ingressroute-design.md#load-balancing)
	// +optional
	Strategy string `json:"strategy,omitempty"`
	// UpstreamValidation defines how to verify the backend service's certificate
	// +optional
	UpstreamValidation *projcontour.UpstreamValidation `json:"validation,omitempty"`
}

// HealthCheck defines health checks on the upstream service.
type HealthCheck struct {
	// HTTP endpoint used to perform health checks on upstream service
	Path string `json:"path"`
	// The value of the host header in the HTTP health check request.
	// If left empty (default value), the name "contour-envoy-healthcheck"
	// will be used.
	// +optional
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

// Delegate allows for delegating VHosts to other IngressRoutes
type Delegate struct {
	// Name of the IngressRoute
	Name string `json:"name"`
	// Namespace of the IngressRoute. Defaults to the current namespace if not supplied.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IngressRoute is an Ingress CRD specificiation
// +k8s:openapi-gen=true
// +kubebuilder:printcolumn:name="FQDN",type="string",JSONPath=".spec.virtualhost.fqdn",description="Fully qualified domain name"
// +kubebuilder:printcolumn:name="TLS Secret",type="string",JSONPath=".spec.virtualhost.tls.secretName",description="Secret with TLS credentials"
// +kubebuilder:printcolumn:name="First route",type="string",JSONPath=".spec.routes[0].match",description="First routes defined"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.currentStatus",description="The current status of the HTTPProxy"
// +kubebuilder:printcolumn:name="Status Description",type="string",JSONPath=".status.description",description="Description of the current status"
// +kubebuilder:resource:scope=Namespaced,path=ingressroutes,singular=ingressroute
type IngressRoute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec IngressRouteSpec `json:"spec"`
	// +optional
	projcontour.Status `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IngressRouteList is a list of IngressRoutes
type IngressRouteList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []IngressRoute `json:"items"`
}
