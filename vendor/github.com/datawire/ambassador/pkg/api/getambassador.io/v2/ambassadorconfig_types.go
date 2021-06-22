// Copyright 2020 Datawire.  All rights reserved
//
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

///////////////////////////////////////////////////////////////////////////
// Important: Run "make update-yaml" to regenerate code after modifying
// this file.
///////////////////////////////////////////////////////////////////////////

package v2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ModuleSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	Config UntypedDict `json:"config,omitempty"`
}

// A Module defines system-wide configuration.  The type of module is
// controlled by the .metadata.name; valid names are "ambassador" or
// "tls".
//
// https://www.getambassador.io/docs/latest/topics/running/ambassador/#the-ambassador-module
// https://www.getambassador.io/docs/latest/topics/running/tls/#tls-module-deprecated
//
// +kubebuilder:object:root=true
type Module struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ModuleSpec `json:"spec,omitempty"`
}

// ModuleList contains a list of Modules.
//
// +kubebuilder:object:root=true
type ModuleList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Module `json:"items"`
}

type Features struct {
	// The diagnostic service (at /ambassador/v0/diag/) defaults on, but
	// you can disable the api route. It will remain accessible on
	// diag_port.
	Diagnostics bool `json:"diagnostics,omitempty"`

	// Should we automatically add Linkerd `l5d-dst-override` headers?
	LinkerdHeaders bool `json:"linkerd_headers,omitempty"`

	// Should we enable the gRPC-http11 bridge?
	GrpcHttp11Bridge bool `json:"grpc_http11_bridge,omitempty"`

	// Should we enable the grpc-Web protocol?
	GrpcWeb bool `json:"grpc_web,omitempty"`

	// Should we enable http/1.0 protocol?
	Http10 bool `json:"http10,omitempty"`

	// Should we do IPv4 DNS lookups when contacting services? Defaults to true,
	// but can be overridden in a [`Mapping`](/reference/mappings).
	Ipv4Dns bool `json:"ipv4_dns,omitempty"`

	// Should we do IPv6 DNS lookups when contacting services? Defaults to false,
	// but can be overridden in a [`Mapping`](/reference/mappings).
	Ipv6Dns bool `json:"ipv6_dns,omitempty"`

	// liveness_probe defaults on, but you can disable the api route.
	// It will remain accessible on diag_port.
	LivenessProbe bool `json:"liveness_probe,omitempty"`

	// readiness_probe defaults on, but you can disable the api route.
	// It will remain accessible on diag_port.
	ReadinessProbe bool `json:"readiness_probe,omitempty"`

	// xff_num_trusted_hops controls the how Envoy sets the trusted
	// client IP address of a request. If you have a proxy in front
	// of Ambassador, Envoy will set the trusted client IP to the
	// address of that proxy. To preserve the orginal client IP address,
	// setting x_num_trusted_hops: 1 will tell Envoy to use the client IP
	// address in X-Forwarded-For. Please see the envoy documentation for
	// more information: https://www.envoyproxy.io/docs/envoy/latest/configuration/http_conn_man/headers#x-forwarded-for
	XffNumTrustedHops int `json:"xff_num_trusted_hops,omitempty"`

	// proxy_proto controls whether Envoy will honor the PROXY
	// protocol on incoming requests.
	ProxyProto bool `json:"proxy_proto,omitempty"`

	// remote_address controls whether Envoy will trust the remote
	// address of incoming connections or rely exclusively on the
	// X-Forwarded_For header.
	RemoteAddress bool `json:"remote_address,omitempty"`

	// Ambassador lets through only the HTTP requests with
	// `X-FORWARDED-PROTO: https` header set, and redirects all the other
	// requests to HTTPS if this field is set to true. Note that `use_remote_address`
	// must be set to false for this feature to work as expected.
	XForwardedProtoRedirect bool `json:"x_forwarded_proto_redirect,omitempty"`
}

// AmbassadorConfigSpec defines the desired state of AmbassadorConfig
type AmbassadorConfigSpec struct {
	// Common to all Ambassador objects (and optional).
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	// admin_port is the port where Ambassador's Envoy will listen for
	// low-level admin requests. You should almost never need to change
	// this.
	AdminPort int `json:"admin_port,omitempty"`

	// diag_port is the port where Ambassador will listen for requests
	// to the diagnostic service.
	DiagPort int `json:"diag_port,omitempty"`

	// By default Envoy sets server_name response header to 'envoy'
	// Override it with this variable
	ServerName string `json:"server_name,omitempty"`

	// If present, service_port will be the port Ambassador listens
	// on for microservice access. If not present, Ambassador will
	// use 8443 if TLS is configured, 8080 otherwise.
	ServicePort int `json:"service_port,omitempty"`

	Features *Features `json:"features,omitempty"`

	// run a custom lua script on every request. see below for more details.
	LuaScripts string `json:"lua_scripts,omitempty"`

	// +kubebuilder:validation:Enum={"text", "json"}
	EnvoyLogType string `json:"envoy_log_type,omitempty"`

	// envoy_log_path defines the path of log envoy will use. By default this is standard output
	EnvoyLogPath string `json:"envoy_log_path,omitempty"`

	LoadBalancer *LoadBalancer `json:"load_balancer,omitempty"`

	CircuitBreakers *CircuitBreaker `json:"circuit_breakers,omitempty"`

	RetryPolicy *RetryPolicy `json:"retry_policy,omitempty"`

	Cors *CORS `json:"cors,omitempty"`

	// Set the default upstream-connection idle timeout. If not set (the default), upstream
	// connections will never be closed due to idling.
	ClusterIdleTimeoutMS int `json:"cluster_idle_timeout_ms,omitempty"`

	// +kubebuilder:validation:Enum={"safe", "unsafe"}
	RegexType string `json:"regex_type,omitempty"`

	// This field controls the RE2 “program size” which is a rough estimate of how complex a compiled regex is to
	// evaluate.  A regex that has a program size greater than the configured value will fail to compile.
	RegexMaxSize int `json:"regex_max_size,omitempty"`
}

// AmbassadorConfigStatus defines the observed state of AmbassadorConfig
type AmbassadorConfigStatus struct {
}

/*
// AmbassadorConfig is the Schema for the ambassadorconfigs API
//
// +kubebuilder:object:root=true
type AmbassadorConfig struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AmbassadorConfigSpec   `json:"spec,omitempty"`
	Status AmbassadorConfigStatus `json:"status,omitempty"`
}

// AmbassadorConfigList contains a list of AmbassadorConfigs.
//
// +kubebuilder:object:root=true
type AmbassadorConfigList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AmbassadorConfig `json:"items"`
}
*/

func init() {
	SchemeBuilder.Register(&Module{}, &ModuleList{})
	//SchemeBuilder.Register(&AmbassadorConfig{}, &AmbassadorConfigList{})
}
