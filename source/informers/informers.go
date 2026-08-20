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

package informers

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// DefaultCacheSyncTimeout is the maximum time in seconds to wait for informer
	// caches to complete initial sync. This is intentionally longer than the
	// per-request timeout: a cache sync may require multiple sequential API calls
	// (LIST + Watch handshake), so the total wait needs to exceed a single request duration.
	DefaultCacheSyncTimeout = 60
)

type informerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

func WaitForCacheSync(ctx context.Context, factory informerFactory, timeout time.Duration) error {
	return waitForCacheSync(ctx, factory.WaitForCacheSync, timeout)
}

func WaitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory, timeout time.Duration) error {
	return waitForCacheSync(ctx, factory.WaitForCacheSync, timeout)
}

// waitForCacheSync waits for informer caches to sync within the given timeout.
// If timeout is <= 0, the default is used.
// Returns an error if any cache fails to sync, wrapping the context error if a timeout occurred.
func waitForCacheSync[K comparable](ctx context.Context, waitFunc func(<-chan struct{}) map[K]bool, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = DefaultCacheSyncTimeout * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	for typ, done := range waitFunc(ctx.Done()) {
		if !done {
			if ctx.Err() != nil {
				return fmt.Errorf("failed to sync %v after %s: %w", typ, timeout, ctx.Err())
			}
			return fmt.Errorf("failed to sync %v", typ)
		}
	}
	return nil
}
