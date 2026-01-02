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

// TODO: move all client generator logic here from store.go.
package source

import (
	"time"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// NewClientGenerator creates a SingletonClientGenerator from external-dns configuration.
// This factory function centralizes the logic for mapping external-dns Config to a
// client generator, ensuring consistent client creation across sources and controllers.
//
// The timeout behavior is special-cased: when cfg.UpdateEvents is true, the timeout
// is set to 0 (no timeout) to allow long-running watch operations.
func NewClientGenerator(cfg *externaldns.Config) *SingletonClientGenerator {
	return &SingletonClientGenerator{
		KubeConfig:     cfg.KubeConfig,
		APIServerURL:   cfg.APIServerURL,
		RequestTimeout: getRequestTimeout(cfg),
	}
}

func getRequestTimeout(cfg *externaldns.Config) time.Duration {
	if cfg.UpdateEvents {
		return 0
	}
	return cfg.RequestTimeout
}
