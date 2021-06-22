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
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MappingSpec defines the desired state of Mapping
type MappingSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	Prefix                string                  `json:"prefix,omitempty"`
	PrefixRegex           bool                    `json:"prefix_regex,omitempty"`
	PrefixExact           bool                    `json:"prefix_exact,omitempty"`
	Service               string                  `json:"service,omitempty"`
	AddRequestHeaders     map[string]AddedHeader  `json:"add_request_headers,omitempty"`
	AddResponseHeaders    map[string]AddedHeader  `json:"add_response_headers,omitempty"`
	AddLinkerdHeaders     bool                    `json:"add_linkerd_headers,omitempty"`
	AutoHostRewrite       bool                    `json:"auto_host_rewrite,omitempty"`
	CaseSensitive         bool                    `json:"case_sensitive,omitempty"`
	EnableIPv4            bool                    `json:"enable_ipv4,omitempty"`
	EnableIPv6            bool                    `json:"enable_ipv6,omitempty"`
	CircuitBreakers       []*CircuitBreaker       `json:"circuit_breakers,omitempty"`
	KeepAlive             *KeepAlive              `json:"keepalive,omitempty"`
	CORS                  *CORS                   `json:"cors,omitempty"`
	RetryPolicy           *RetryPolicy            `json:"retry_policy,omitempty"`
	GRPC                  bool                    `json:"grpc,omitempty"`
	HostRedirect          bool                    `json:"host_redirect,omitempty"`
	HostRewrite           string                  `json:"host_rewrite,omitempty"`
	Method                string                  `json:"method,omitempty"`
	MethodRegex           bool                    `json:"method_regex,omitempty"`
	OutlierDetection      string                  `json:"outlier_detection,omitempty"`
	PathRedirect          string                  `json:"path_redirect,omitempty"`
	Priority              string                  `json:"priority,omitempty"`
	Precedence            int                     `json:"precedence,omitempty"`
	ClusterTag            string                  `json:"cluster_tag,omitempty"`
	RemoveRequestHeaders  StringOrStringList      `json:"remove_request_headers,omitempty"`
	RemoveResponseHeaders StringOrStringList      `json:"remove_response_headers,omitempty"`
	Resolver              string                  `json:"resolver,omitempty"`
	Rewrite               string                  `json:"rewrite,omitempty"`
	RegexRewrite          map[string]BoolOrString `json:"regex_rewrite,omitempty"`
	Shadow                bool                    `json:"shadow,omitempty"`
	ConnectTimeoutMs      int                     `json:"connect_timeout_ms,omitempty"`
	ClusterIdleTimeoutMs  int                     `json:"cluster_idle_timeout_ms,omitempty"`
	TimeoutMs             int                     `json:"timeout_ms,omitempty"`
	IdleTimeoutMs         int                     `json:"idle_timeout_ms,omitempty"`
	TLS                   BoolOrString            `json:"tls,omitempty"`
	UseWebsocket          bool                    `json:"use_websocket,omitempty"`
	Weight                int                     `json:"weight,omitempty"`
	BypassAuth            bool                    `json:"bypass_auth,omitempty"`
	Modules               []UntypedDict           `json:"modules,omitempty"`
	Host                  string                  `json:"host,omitempty"`
	HostRegex             bool                    `json:"host_regex,omitempty"`
	Headers               map[string]BoolOrString `json:"headers,omitempty"`
	RegexHeaders          map[string]BoolOrString `json:"regex_headers,omitempty"`
	Labels                MappingLabels           `json:"labels,omitempty"`
	EnvoyOverride         UntypedDict             `json:"envoy_override,omitempty"`
	LoadBalancer          *LoadBalancer           `json:"load_balancer,omitempty"`
	QueryParameters       map[string]BoolOrString `json:"query_parameters,omitempty"`
	RegexQueryParameters  map[string]BoolOrString `json:"regex_query_parameters,omitempty"`
}

// Python: MappingLabels = Dict[str, Union[str,'MappingLabels']]
type MappingLabels map[string]StringOrMappingLabels

// StringOrMapping labels is the `Union[str,'MappingLabels']` part of
// the MappingLabels type.
//
// See the remarks about schema on custom types in `./common.go`.
//
// +kubebuilder:validation:Type=""
type StringOrMappingLabels struct {
	String *string
	Labels MappingLabels
}

// MarshalJSON is important both so that we generate the proper
// output, and to trigger controller-gen to not try to generate
// jsonschema for our sub-fields:
// https://github.com/kubernetes-sigs/controller-tools/pull/427
func (o StringOrMappingLabels) MarshalJSON() ([]byte, error) {
	switch {
	case o.String == nil && o.Labels == nil:
		return json.Marshal(nil)
	case o.String == nil && o.Labels != nil:
		return json.Marshal(o.Labels)
	case o.String != nil && o.Labels == nil:
		return json.Marshal(o.String)
	case o.String != nil && o.Labels != nil:
		panic("invalid StringOrMappingLabels")
	}
	panic("not reached")
}

func (o *StringOrMappingLabels) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*o = StringOrMappingLabels{}
		return nil
	}

	var err error

	var labels MappingLabels
	if err = json.Unmarshal(data, &labels); err == nil {
		*o = StringOrMappingLabels{Labels: labels}
		return nil
	}

	var str string
	if err = json.Unmarshal(data, &str); err == nil {
		*o = StringOrMappingLabels{String: &str}
		return nil
	}

	return err
}

// +kubebuilder:validation:Type="d6e-union:string,boolean,object"
type AddedHeader struct {
	String *string
	Bool   *bool
	Object *UntypedDict
}

// MarshalJSON is important both so that we generate the proper
// output, and to trigger controller-gen to not try to generate
// jsonschema for our sub-fields:
// https://github.com/kubernetes-sigs/controller-tools/pull/427
func (o AddedHeader) MarshalJSON() ([]byte, error) {
	switch {
	case o.String != nil:
		return json.Marshal(*o.String)
	case o.Bool != nil:
		return json.Marshal(*o.Bool)
	case o.Object != nil:
		return json.Marshal(*o.Object)
	default:
		return json.Marshal(nil)
	}
}

func (o *AddedHeader) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		*o = AddedHeader{}
		return nil
	}

	var err error

	var str string
	if err = json.Unmarshal(data, &str); err == nil {
		*o = AddedHeader{String: &str}
		return nil
	}

	var b bool
	if err = json.Unmarshal(data, &b); err == nil {
		*o = AddedHeader{Bool: &b}
		return nil
	}

	var obj UntypedDict
	if err = json.Unmarshal(data, &obj); err == nil {
		*o = AddedHeader{Object: &obj}
		return nil
	}

	return err
}

type KeepAlive struct {
	Probes   int `json:"probes,omitempty"`
	IdleTime int `json:"idle_time,omitempty"`
	Interval int `json:"interval,omitempty"`
}

type CORS struct {
	Origins        StringOrStringList `json:"origins,omitempty"`
	Methods        StringOrStringList `json:"methods,omitempty"`
	Headers        StringOrStringList `json:"headers,omitempty"`
	Credentials    bool               `json:"credentials,omitempty"`
	ExposedHeaders StringOrStringList `json:"exposed_headers,omitempty"`
	MaxAge         string             `json:"max_age,omitempty"`
}

type RetryPolicy struct {
	// +kubebuilder:validation:Enum={"5xx","gateway-error","connect-failure","retriable-4xx","refused-stream","retriable-status-codes"}
	RetryOn       string `json:"retry_on,omitempty"`
	NumRetries    int    `json:"num_retries,omitempty"`
	PerTryTimeout string `json:"per_try_timeout,omitempty"`
}

type LoadBalancer struct {
	// +kubebuilder:validation:Enum={"round_robin","ring_hash","maglev","least_request"}
	// +kubebuilder:validation:Required
	Policy   string              `json:"policy,omitempty"`
	Cookie   *LoadBalancerCookie `json:"cookie,omitempty"`
	Header   string              `json:"header,omitempty"`
	SourceIp bool                `json:"source_ip,omitempty"`
}

type LoadBalancerCookie struct {
	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
	Ttl  string `json:"ttl,omitempty"`
}

// MappingStatus defines the observed state of Mapping
type MappingStatus struct {
	// +kubebuilder:validation:Enum={"","Inactive","Running"}
	State string `json:"state,omitempty"`

	Reason string `json:"reason,omitempty"`
}

// Mapping is the Schema for the mappings API
//
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Prefix",type=string,JSONPath=`.spec.prefix`
// +kubebuilder:printcolumn:name="Service",type=string,JSONPath=`.spec.service`
// +kubebuilder:printcolumn:name="State",type=string,JSONPath=`.status.state`
// +kubebuilder:printcolumn:name="Reason",type=string,JSONPath=`.status.reason`
type Mapping struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MappingSpec   `json:"spec,omitempty"`
	Status MappingStatus `json:"status,omitempty"`
}

// MappingList contains a list of Mappings.
//
// +kubebuilder:object:root=true
type MappingList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mapping `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Mapping{}, &MappingList{})
}
