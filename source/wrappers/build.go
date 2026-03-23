/*
Copyright 2026 The Kubernetes Authors.

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

package wrappers

import (
	"context"

	"sigs.k8s.io/external-dns/source"
)

// Build creates all named sources using cfg's ClientGenerator and wraps them
// with the standard pipeline (dedup, optional NAT64, optional target filter,
// post-processor). Inject a custom ClientGenerator via source.WithClientGenerator.
func Build(ctx context.Context, cfg *source.Config) (source.Source, error) {
	sources, err := source.ByNames(ctx, cfg, cfg.ClientGenerator())
	if err != nil {
		return nil, err
	}
	opts := NewConfig(
		WithDefaultTargets(cfg.DefaultTargets),
		WithForceDefaultTargets(cfg.ForceDefaultTargets),
		WithNAT64Networks(cfg.NAT64Networks),
		WithTargetNetFilter(cfg.TargetNetFilter),
		WithExcludeTargetNets(cfg.ExcludeTargetNets),
		WithMinTTL(cfg.MinTTL),
		WithProvider(cfg.Provider),
		WithPreferAlias(cfg.PreferAlias),
	)
	return wrapSources(sources, opts)
}
