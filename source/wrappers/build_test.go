/*
Copyright 2026 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package wrappers

import (
	"fmt"
	"testing"
	"time"

	openshift "github.com/openshift/client-go/route/clientset/versioned"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	istioclient "istio.io/client-go/pkg/clientset/versioned"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	gateway "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/types"
)

// stubClientGenerator satisfies source.ClientGenerator for sources that do not
// require any Kubernetes client (e.g. the "fake" source).
type stubClientGenerator struct{}

func (stubClientGenerator) KubeClient() (kubernetes.Interface, error) {
	return nil, fmt.Errorf("KubeClient: not available in stub")
}
func (stubClientGenerator) GatewayClient() (gateway.Interface, error) {
	return nil, fmt.Errorf("GatewayClient: not available in stub")
}
func (stubClientGenerator) IstioClient() (istioclient.Interface, error) {
	return nil, fmt.Errorf("IstioClient: not available in stub")
}
func (stubClientGenerator) DynamicKubernetesClient() (dynamic.Interface, error) {
	return nil, fmt.Errorf("DynamicKubernetesClient: not available in stub")
}
func (stubClientGenerator) OpenShiftClient() (openshift.Interface, error) {
	return nil, fmt.Errorf("OpenShiftClient: not available in stub")
}
func (stubClientGenerator) RESTConfig() (*rest.Config, error) {
	return nil, fmt.Errorf("RESTConfig: not available in stub")
}

func TestBuildWrappedSource(t *testing.T) {
	tests := []struct {
		name    string
		cfg     *source.Config
		wantErr bool
	}{
		{
			name: "fake source with no extra wrappers",
			cfg: source.NewSourceConfig(&externaldns.Config{
				Sources: []string{types.Fake},
			}),
		},
		{
			name: "fake source with target filter wrapper",
			cfg: source.NewSourceConfig(&externaldns.Config{
				Sources:         []string{types.Fake},
				TargetNetFilter: []string{"10.0.0.0/8"},
			}),
		},
		{
			name: "fake source with NAT64 networks",
			cfg: source.NewSourceConfig(&externaldns.Config{
				Sources:       []string{types.Fake},
				NAT64Networks: []string{"2001:db8::/96"},
			}),
		},
		{
			name: "fake source with minTTL, provider, and preferAlias",
			cfg: source.NewSourceConfig(&externaldns.Config{
				Sources:     []string{types.Fake},
				MinTTL:      300 * time.Second,
				Provider:    "aws",
				PreferAlias: true,
			}),
		},
		{
			name: "fake source with exclude target nets",
			cfg: source.NewSourceConfig(&externaldns.Config{
				Sources:           []string{types.Fake},
				TargetNetFilter:   []string{"10.0.0.0/8"},
				ExcludeTargetNets: []string{"10.1.0.0/16"},
			}),
		},
		{
			name: "unknown source returns error",
			cfg: source.NewSourceConfig(&externaldns.Config{
				Sources: []string{"does-not-exist"},
			}),
			wantErr: true,
		},
		{
			name: "invalid NAT64 network returns error",
			cfg: source.NewSourceConfig(&externaldns.Config{
				Sources:       []string{types.Fake},
				NAT64Networks: []string{"not-a-cidr"},
			}),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, err := BuildWrappedSource(t.Context(), tt.cfg, stubClientGenerator{})
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.NotNil(t, src)
		})
	}
}
