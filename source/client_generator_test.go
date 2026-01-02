/*
Copyright 2025 The Kubernetes Authors.

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

package source

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

func TestNewClientGenerator(t *testing.T) {
	cfg := &externaldns.Config{
		KubeConfig:     "/path/to/kubeconfig",
		APIServerURL:   "https://api.example.com",
		RequestTimeout: 30 * time.Second,
		UpdateEvents:   false,
	}

	gen := NewClientGenerator(cfg)

	assert.Equal(t, "/path/to/kubeconfig", gen.KubeConfig)
	assert.Equal(t, "https://api.example.com", gen.APIServerURL)
	assert.Equal(t, 30*time.Second, gen.RequestTimeout)
}

func TestNewClientGenerator_UpdateEvents(t *testing.T) {
	cfg := &externaldns.Config{
		KubeConfig:     "/path/to/kubeconfig",
		APIServerURL:   "https://api.example.com",
		RequestTimeout: 30 * time.Second,
		UpdateEvents:   true, // Special case
	}

	gen := NewClientGenerator(cfg)

	assert.Equal(t, time.Duration(0), gen.RequestTimeout, "UpdateEvents should set timeout to 0")
}
