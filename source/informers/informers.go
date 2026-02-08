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
	defaultRequestTimeout = 60
)

type informerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

// controllerInformerFactory is the subset of controller-runtime cache.Cache needed for start and sync.
type controllerInformerFactory interface {
	Start(ctx context.Context) error
	WaitForCacheSync(ctx context.Context) bool
}

func WaitForCacheSync(ctx context.Context, factory informerFactory) error {
	return waitForCacheSync(ctx, factory.WaitForCacheSync)
}

func WaitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory) error {
	return waitForCacheSync(ctx, factory.WaitForCacheSync)
}

// StartAndWaitForCacheSync starts a controller-runtime cache in a goroutine and waits
// for it to sync with a default timeout. This is the controller-runtime equivalent of
// WaitForCacheSync / WaitForDynamicCacheSync.
func StartAndWaitForCacheSync(ctx context.Context, c controllerInformerFactory) error {
	startErrCh := make(chan error, 1)
	go func() {
		startErrCh <- c.Start(ctx)
	}()

	timeout := defaultRequestTimeout * time.Second
	syncCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if !c.WaitForCacheSync(syncCtx) {
		// Check if Start itself failed (e.g. already started)
		select {
		case err := <-startErrCh:
			if err != nil {
				return fmt.Errorf("failed to start controller-runtime cache: %w", err)
			}
		default:
		}
		// Include timeout and context error (if any) in the error message
		return fmt.Errorf("failed to sync controller-runtime cache after %s: %w", timeout, syncCtx.Err())
	}
	return nil
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
