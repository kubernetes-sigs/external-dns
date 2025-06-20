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

func WaitForCacheSync(ctx context.Context, factory informerFactory) error {
	timeout := defaultRequestTimeout * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	for typ, done := range factory.WaitForCacheSync(ctx.Done()) {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("failed to sync %v: %w with timeout %s", typ, ctx.Err(), timeout)
			default:
				return fmt.Errorf("failed to sync %v with timeout %s", typ, timeout)
			}
		}
	}
	return nil
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

func WaitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory) error {
	timeout := defaultRequestTimeout * time.Second
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()
	for typ, done := range factory.WaitForCacheSync(ctx.Done()) {
		if !done {
			select {
			case <-ctx.Done():
				return fmt.Errorf("failed to sync %v: %w with timeout %s", typ, ctx.Err(), timeout)
			default:
				return fmt.Errorf("failed to sync %v with timeout %s", typ, timeout)
			}
		}
	}
	return nil
}
