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

package annotations

import (
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"
)

func hasAliasFromAnnotations(annotations map[string]string) bool {
	aliasAnnotation, ok := annotations[AliasKey]
	return ok && aliasAnnotation == "true"
}

// TTLFromAnnotations extracts the TTL from the annotations of the given resource.
func TTLFromAnnotations(annotations map[string]string, resource string) endpoint.TTL {
	ttlNotConfigured := endpoint.TTL(0)
	ttlAnnotation, ok := annotations[TtlKey]
	if !ok {
		return ttlNotConfigured
	}
	ttlValue, err := parseTTL(ttlAnnotation)
	if err != nil {
		log.Warnf("%s: %q is not a valid TTL value: %v", resource, ttlAnnotation, err)
		return ttlNotConfigured
	}
	if ttlValue < ttlMinimum || ttlValue > ttlMaximum {
		log.Warnf("TTL value %q must be between [%d, %d]", ttlValue, ttlMinimum, ttlMaximum)
		return ttlNotConfigured
	}
	return endpoint.TTL(ttlValue)
}

// parseTTL parses TTL from string, returning duration in seconds.
// parseTTL supports both integers like "600" and durations based
// on Go Duration like "10m", hence "600" and "10m" represent the same value.
//
// Note: for durations like "1.5s" the fraction is omitted (resulting in 1 second for the example).
func parseTTL(s string) (int64, error) {
	ttlDuration, errDuration := time.ParseDuration(s)
	if errDuration != nil {
		ttlInt, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, errDuration
		}
		return ttlInt, nil
	}

	return int64(ttlDuration.Seconds()), nil
}

// ParseFilter parses an annotation filter string into a labels.Selector.
// Returns nil if the annotation filter is invalid.
func ParseFilter(annotationFilter string) (labels.Selector, error) {
	labelSelector, err := metav1.ParseToLabelSelector(annotationFilter)
	if err != nil {
		return nil, err
	}
	selector, err := metav1.LabelSelectorAsSelector(labelSelector)
	if err != nil {
		return nil, err
	}
	return selector, nil
}

// TargetsFromTargetAnnotation gets endpoints from optional "target" annotation.
// Returns empty endpoints array if none are found.
func TargetsFromTargetAnnotation(annotations map[string]string) endpoint.Targets {
	var targets endpoint.Targets
	// Get the desired hostname of the ingress from the annotation.
	targetAnnotation, ok := annotations[TargetKey]
	if ok && targetAnnotation != "" {
		// splits the hostname annotation and removes the trailing periods
		targetsList := SplitHostnameAnnotation(targetAnnotation)
		for _, targetHostname := range targetsList {
			targetHostname = strings.TrimSuffix(targetHostname, ".")
			targets = append(targets, targetHostname)
		}
	}
	return targets
}

// HostnamesFromAnnotations extracts the hostnames from the given annotations map.
// It returns a slice of hostnames if the HostnameKey annotation is present, otherwise it returns nil.
func HostnamesFromAnnotations(input map[string]string) []string {
	return extractHostnamesFromAnnotations(input, HostnameKey)
}

// InternalHostnamesFromAnnotations extracts the internal hostnames from the given annotations map.
// It returns a slice of internal hostnames if the InternalHostnameKey annotation is present, otherwise it returns nil.
func InternalHostnamesFromAnnotations(input map[string]string) []string {
	return extractHostnamesFromAnnotations(input, InternalHostnameKey)
}

// SplitHostnameAnnotation splits a comma-separated hostname annotation string into a slice of hostnames.
// It trims any leading or trailing whitespace and removes any spaces within the anno
func SplitHostnameAnnotation(input string) []string {
	return strings.Split(strings.TrimSpace(strings.ReplaceAll(input, " ", "")), ",")
}

func extractHostnamesFromAnnotations(input map[string]string, key string) []string {
	annotation, ok := input[key]
	if !ok {
		return nil
	}
	return SplitHostnameAnnotation(annotation)
}
