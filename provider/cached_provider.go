package provider

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

var (
	cachedRecordsCallsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "provider",
			Name:      "cache_records_calls",
			Help:      "Number of calls to the provider cache Records list.",
		},
		[]string{
			"from_cache",
		},
	)
	cachedApplyChangesCallsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: "external_dns",
			Subsystem: "provider",
			Name:      "cache_apply_changes_calls",
			Help:      "Number of calls to the provider cache ApplyChanges.",
		},
	)
)

type CachedProvider struct {
	Provider
	RefreshDelay time.Duration
	err          error
	lastRead     time.Time
	cache        []*endpoint.Endpoint
}

func (c *CachedProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	if c.needRefresh() {
		log.Info("Records cache provider: refreshing records list cache")
		c.cache, c.err = c.Provider.Records(ctx)
		if c.err != nil {
			log.Errorf("Records cache provider: list records failed: %v", c.err)
		}
		c.lastRead = time.Now()
		cachedRecordsCallsTotal.WithLabelValues("false").Inc()
	} else {
		log.Info("Records cache provider: using records list from cache")
		cachedRecordsCallsTotal.WithLabelValues("true").Inc()
	}
	return c.cache, c.err
}
func (c *CachedProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if !changes.HasChanges() {
		return nil
	}
	c.Reset()
	cachedApplyChangesCallsTotal.Inc()
	return c.Provider.ApplyChanges(ctx, changes)
}

func (c *CachedProvider) Reset() {
	c.err = nil
	c.cache = nil
	c.lastRead = time.Time{}
}

func (c *CachedProvider) needRefresh() bool {
	if c.cache == nil || c.err != nil {
		log.Debug("Records cache provider is not initialized")
		return true
	}
	log.Debug("Records cache last Read: ", c.lastRead, "expiration: ", c.RefreshDelay, " provider expiration:", c.lastRead.Add(c.RefreshDelay), "expired: ", time.Now().After(c.lastRead.Add(c.RefreshDelay)))
	return time.Now().After(c.lastRead.Add(c.RefreshDelay))
}

func init() {
	prometheus.MustRegister(cachedRecordsCallsTotal)
}
