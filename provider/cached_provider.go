/*
Copyright 2017 The Kubernetes Authors.

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

package provider

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/pkg/metrics"
	"sigs.k8s.io/external-dns/plan"
)

var (
	cachedRecordsCallsTotal = metrics.NewCounterVecWithOpts(
		prometheus.CounterOpts{
			Subsystem: "provider",
			Name:      "cache_records_calls",
			Help:      "Number of calls to the provider cache Records list.",
		},
		[]string{
			"from_cache",
		},
	)
	cachedApplyChangesCallsTotal = metrics.NewCounterWithOpts(
		prometheus.CounterOpts{
			Subsystem: "provider",
			Name:      "cache_apply_changes_calls",
			Help:      "Number of calls to the provider cache ApplyChanges.",
		},
	)
)

func init() {
	metrics.RegisterMetric.MustRegister(cachedRecordsCallsTotal)
	metrics.RegisterMetric.MustRegister(cachedApplyChangesCallsTotal)
}

type CachedProvider struct {
	Provider
	RefreshDelay   time.Duration
	PatchOnApply   bool
	lastRead       time.Time
	cache          []*endpoint.Endpoint
}

func NewCachedProvider(provider Provider, cfg *externaldns.Config) *CachedProvider {
	return &CachedProvider{
		Provider:     provider,
		RefreshDelay: cfg.ProviderCacheTime,
		PatchOnApply: cfg.ProviderCachePatchOnApply,
	}
}

func (c *CachedProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	if c.needRefresh() {
		log.Info("Records cache provider: refreshing records list cache")
		records, err := c.Provider.Records(ctx)
		if err != nil {
			c.cache = nil
			return nil, err
		}
		c.cache = records
		c.lastRead = time.Now()
		cachedRecordsCallsTotal.CounterVec.WithLabelValues("false").Inc()
	} else {
		log.Debug("Records cache provider: using records list from cache")
		cachedRecordsCallsTotal.CounterVec.WithLabelValues("true").Inc()
	}
	return c.cache, nil
}
func (c *CachedProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if !changes.HasChanges() {
		log.Info("Records cache provider: no changes to be applied")
		return nil
	}
	cachedApplyChangesCallsTotal.Counter.Inc()
	if err := c.Provider.ApplyChanges(ctx, changes); err != nil {
		c.Reset()
		return err
	}
	if c.PatchOnApply && c.cache != nil {
		c.cache = applyChangesToCache(c.cache, changes)
	} else {
		c.Reset()
	}
	return nil
}

func applyChangesToCache(cache []*endpoint.Endpoint, changes *plan.Changes) []*endpoint.Endpoint {
	remove := make(map[string]bool, len(changes.Delete)+len(changes.UpdateOld))
	for _, ep := range changes.Delete {
		remove[ep.DNSName+"/"+ep.RecordType] = true
	}
	for _, ep := range changes.UpdateOld {
		remove[ep.DNSName+"/"+ep.RecordType] = true
	}
	updated := make([]*endpoint.Endpoint, 0, len(cache))
	for _, ep := range cache {
		if !remove[ep.DNSName+"/"+ep.RecordType] {
			updated = append(updated, ep)
		}
	}
	updated = append(updated, changes.Create...)
	updated = append(updated, changes.UpdateNew...)
	return updated
}

func (c *CachedProvider) Reset() {
	c.cache = nil
	c.lastRead = time.Time{}
}

func (c *CachedProvider) needRefresh() bool {
	if c.cache == nil {
		log.Debug("Records cache provider is not initialized")
		return true
	}
	log.Debug("Records cache last Read: ", c.lastRead, "expiration: ", c.RefreshDelay, " provider expiration:", c.lastRead.Add(c.RefreshDelay), "expired: ", time.Now().After(c.lastRead.Add(c.RefreshDelay)))
	return time.Now().After(c.lastRead.Add(c.RefreshDelay))
}
