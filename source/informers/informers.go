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
	"reflect"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	defaultInformerSyncTimeout = 60
)

type informerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[reflect.Type]bool
}

// WaitForCacheSync waits for all informers in the factory to sync their caches.
// If sync fails or times out, it logs a warning and continues rather than returning an error.
// This "soft error" approach prevents crash loops that can overwhelm the Kubernetes API server
// with repeated LIST/WATCH calls. Common causes of sync failures include:
//   - Missing RBAC permissions (check ServiceAccount permissions)
//   - API server latency or network issues
//   - Resource contention from other controllers
//
// The function returns nil to allow the application to continue operating with potentially
// stale cache data, which is preferable to crashing repeatedly.
func WaitForCacheSync(ctx context.Context, factory informerFactory, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = defaultInformerSyncTimeout * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	syncResults := factory.WaitForCacheSync(ctx.Done())
	allSynced := true
	for typ, done := range syncResults {
		if !done {
			allSynced = false
			select {
			case <-ctx.Done():
				log.Warnf("Cache sync for %v timed out after %s: %v. "+
					"This may indicate RBAC issues, API server latency, or network problems. "+
					"Continuing with potentially stale data.", typ, timeout, ctx.Err())
			default:
				log.Warnf("Cache sync for %v failed within timeout %s. "+
					"This may indicate RBAC issues or other problems. "+
					"Continuing with potentially stale data.", typ, timeout)
			}
		}
	}

	if !allSynced {
		log.Warn("Not all informer caches synced successfully. " +
			"ExternalDNS will continue but may have incomplete data. " +
			"Please check RBAC permissions and API server connectivity.")
	}

	return nil
}

type dynamicInformerFactory interface {
	WaitForCacheSync(stopCh <-chan struct{}) map[schema.GroupVersionResource]bool
}

// WaitForDynamicCacheSync waits for all dynamic informers in the factory to sync their caches.
// Similar to WaitForCacheSync, it uses soft error handling to prevent crash loops.
// See WaitForCacheSync documentation for details on error handling behavior.
func WaitForDynamicCacheSync(ctx context.Context, factory dynamicInformerFactory, timeout time.Duration) error {
	if timeout <= 0 {
		timeout = defaultInformerSyncTimeout * time.Second
	}
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	syncResults := factory.WaitForCacheSync(ctx.Done())
	allSynced := true
	for typ, done := range syncResults {
		if !done {
			allSynced = false
			select {
			case <-ctx.Done():
				log.Warnf("Cache sync for %v timed out after %s: %v. "+
					"This may indicate RBAC issues, API server latency, or network problems. "+
					"Continuing with potentially stale data.", typ, timeout, ctx.Err())
			default:
				log.Warnf("Cache sync for %v failed within timeout %s. "+
					"This may indicate RBAC issues or other problems. "+
					"Continuing with potentially stale data.", typ, timeout)
			}
		}
	}

	if !allSynced {
		log.Warn("Not all informer caches synced successfully. " +
			"ExternalDNS will continue but may have incomplete data. " +
			"Please check RBAC permissions and API server connectivity.")
	}

	return nil
}
