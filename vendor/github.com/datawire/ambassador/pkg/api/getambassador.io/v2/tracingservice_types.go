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

type TraceSampling struct {
	Client  int `json:"client,omitempty"`
	Random  int `json:"random,omitempty"`
	Overall int `json:"overall,omitempty"`
}

type TraceConfig struct {
	AccessTokenFile   string `json:"access_token_file,omitempty"`
	CollectorCluster  string `json:"collector_cluster,omitempty"`
	CollectorEndpoint string `json:"collector_endpoint,omitempty"`
	TraceID128Bit     bool   `json:"trace_id_128bit,omitempty"`
	SharedSpanContext bool   `json:"shared_span_context,omitempty"`
	ServiceName       string `json:"service_name,omitempty"`
}

// TracingServiceSpec defines the desired state of TracingService
type TracingServiceSpec struct {
	AmbassadorID AmbassadorID `json:"ambassador_id,omitempty"`

	// +kubebuilder:validation:Enum={"lightstep","zipkin","datadog"}
	Driver     string         `json:"driver,omitempty"`
	Service    string         `json:"service,omitempty"`
	Sampling   *TraceSampling `json:"sampling,omitempty"`
	TagHeaders []string       `json:"tag_headers,omitempty"`
	Config     *TraceConfig   `json:"config,omitempty"`
}

// TracingService is the Schema for the tracingservices API
//
// +kubebuilder:object:root=true
type TracingService struct {
	metav1.TypeMeta   `json:""`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec TracingServiceSpec `json:"spec,omitempty"`
}

// TracingServiceList contains a list of TracingServices.
//
// +kubebuilder:object:root=true
type TracingServiceList struct {
	metav1.TypeMeta `json:""`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TracingService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TracingService{}, &TracingServiceList{})
}
