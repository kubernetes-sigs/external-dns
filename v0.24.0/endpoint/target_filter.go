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
	IsEnabled() bool
}

// TargetNetFilter holds a lists of valid target names
type TargetNetFilter struct {
	// filterNets define what targets to match
	filterNets []*net.IPNet
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
	return TargetNetFilter{filterNets: prepareTargetFilters(targetFilterNets), excludeNets: prepareTargetFilters(excludeNets)}
}

// Match checks whether a target can be found in the TargetNetFilter.
func (tf TargetNetFilter) Match(target string) bool {
	return matchTargetNetFilter(tf.filterNets, target, true) && !matchTargetNetFilter(tf.excludeNets, target, false)
}

// IsEnabled returns true if any filters or exclusions are set.
func (tf TargetNetFilter) IsEnabled() bool {
	return len(tf.filterNets) > 0 || len(tf.excludeNets) > 0
}

// matchTargetNetFilter determines if any `filters` match `target`.
// If no `filters` are provided, behavior depends on `emptyval`
// (empty `tf.filters` matches everything, while empty `tf.exclude` excludes nothing)
func matchTargetNetFilter(filters []*net.IPNet, target string, emptyval bool) bool {
	if len(filters) == 0 {
		return emptyval
	}

	ip := net.ParseIP(target)

	for _, filter := range filters {
		if filter.Contains(ip) {
			return true
		}
	}

	return false
}
