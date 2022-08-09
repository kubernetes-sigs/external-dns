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

package endpoint

import (
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

// TargetFilterInterface defines the interface to select matching targets for a specific provider or runtime
type TargetFilterInterface interface {
	Match(target string) bool
	IsConfigured() bool
}

// TargetNetFilter holds a lists of valid target names
type TargetNetFilter struct {
	// FilterNets define what targets to match
	FilterNets []*net.IPNet
	// excludeNets define what targets not to match
	excludeNets []*net.IPNet
}

// prepareTargetFilters provides consistent trimming for filters/exclude params
func prepareTargetFilters(filters []string) []*net.IPNet {
	fs := make([]*net.IPNet, 0)

	for _, filter := range filters {
		filter = strings.TrimSpace(filter)

		_, filterNet, err := net.ParseCIDR(filter)
		if err != nil {
			log.Errorf("Invalid target net filter: %s", filter)

			continue
		}

		fs = append(fs, filterNet)
	}
	return fs
}

// NewTargetNetFilterWithExclusions returns a new TargetNetFilter, given a list of matches and exclusions
func NewTargetNetFilterWithExclusions(targetFilterNets []string, excludeNets []string) TargetNetFilter {
	return TargetNetFilter{FilterNets: prepareTargetFilters(targetFilterNets), excludeNets: prepareTargetFilters(excludeNets)}
}

// NewTargetNetFilter returns a new TargetNetFilter given a comma separated list of targets
func NewTargetNetFilter(targetFilterNets []string) TargetNetFilter {
	return TargetNetFilter{FilterNets: prepareTargetFilters(targetFilterNets)}
}

// Match checks whether a target can be found in the TargetNetFilter.
func (tf TargetNetFilter) Match(target string) bool {
	return matchTargetNetFilter(tf.FilterNets, target, true) && !matchTargetNetFilter(tf.excludeNets, target, false)
}

// matchTargetNetFilter determines if any `filters` match `target`.
// If no `filters` are provided, behavior depends on `emptyval`
// (empty `tf.filters` matches everything, while empty `tf.exclude` excludes nothing)
func matchTargetNetFilter(filters []*net.IPNet, target string, emptyval bool) bool {
	if len(filters) == 0 {
		return emptyval
	}

	for _, filter := range filters {
		ip := net.ParseIP(target)

		if filter.Contains(ip) {
			return true
		}
	}

	return false
}

// IsConfigured returns true if TargetFilter is configured, false otherwise
func (tf TargetNetFilter) IsConfigured() bool {
	if len(tf.FilterNets) == 1 {
		return tf.FilterNets[0].Network() != ""
	}
	return len(tf.FilterNets) > 0
}
