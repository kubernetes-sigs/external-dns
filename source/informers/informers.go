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
	"os"
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	defaultRequestTimeout = 60
)

// CacheLifecycle covers the start/sync contract satisfied by the real
// controller-runtime cache in production and a test stub in tests.
type CacheLifecycle interface {
	Start(ctx context.Context) error
	WaitForCacheSync(ctx context.Context) bool
}

// StartAndWaitForCacheSync starts lc in a goroutine and waits for all its
// informers to sync, returning an error if the context expires first.
func StartAndWaitForCacheSync(ctx context.Context, lc CacheLifecycle) error {
	fmt.Println("Starting informers...")
	go func() {
		err := lc.Start(ctx)
		if err != nil {
			fmt.Println("fail informers...", err)
			_, _ = fmt.Fprintf(os.Stderr, "Error starting informers: %v\n", err)
		}
	}()
	if !lc.WaitForCacheSync(ctx) {
		if ctx.Err() != nil {
			return fmt.Errorf("cache failed to sync: %w", ctx.Err())
		}
		return fmt.Errorf("cache failed to sync")
	}
	return nil
}

type informerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

func WaitForCacheSync(ctx context.Context, factory informerFactory) error {
	return waitForCacheSync(ctx, factory.WaitForCacheSync)
}

func WaitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory) error {
	return waitForCacheSync(ctx, factory.WaitForCacheSync)
}

// waitForCacheSync waits for informer caches to sync with a default timeout.
// Returns an error if any cache fails to sync, wrapping the context error if a timeout occurred.
func waitForCacheSync[K comparable](ctx context.Context, waitFunc func(<-chan struct{}) map[K]bool) error {
	// The function receives a ctx but then creates a new timeout,
	// effectively overriding whatever deadline the caller may have set.
	// If the caller passed a context with a 30s timeout, this function ignores it and waits 60s anyway.
	timeout := defaultRequestTimeout * time.Second
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
