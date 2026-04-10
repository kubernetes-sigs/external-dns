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
	DefaultCacheSyncTimeout = 60
)

type informerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

// WaitForCacheSync waits for all informers in the factory to sync their caches.
// Returns an error if any informer fails to sync within the given timeout.
func WaitForCacheSync(ctx context.Context, factory informerFactory, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = DefaultCacheSyncTimeout * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	syncResults := factory.WaitForCacheSync(ctx.Done())
	for typ, done := range syncResults {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("cache sync for %v timed out after %s: %w", typ, timeout, ctx.Err())
			default:
				return fmt.Errorf("cache sync for %v failed", typ)
			}
		}
	}

	return nil
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

// WaitForDynamicCacheSync waits for all dynamic informers in the factory to sync their caches.
// Returns an error if any informer fails to sync within the given timeout.
func WaitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = DefaultCacheSyncTimeout * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	syncResults := factory.WaitForCacheSync(ctx.Done())
	for typ, done := range syncResults {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("cache sync for %v timed out after %s: %w", typ, timeout, ctx.Err())
			default:
				return fmt.Errorf("cache sync for %v failed", typ)
			}
		}
	}

	return nil
}
