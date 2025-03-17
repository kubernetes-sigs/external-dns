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

package utils

import (
	"fmt"
	"net/netip"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"

	"sigs.k8s.io/external-dns/endpoint"
)

// TTLFromAnnotations TODO: copied from source.go. Refactor to avoid duplication.
// TTLFromAnnotations extracts the TTL from the annotations of the given resource.
func TTLFromAnnotations(annotations map[string]string, resource string) endpoint.TTL {
	ttlNotConfigured := endpoint.TTL(0)
	ttlAnnotation, exists := annotations[ttlAnnotationKey]
	if !exists {
		return ttlNotConfigured
	}
	ttlValue, err := parseTTL(ttlAnnotation)
	if err != nil {
		log.Warnf("%s: \"%v\" is not a valid TTL value: %v", resource, ttlAnnotation, err)
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
// Note: for durations like "1.5s" the fraction is omitted (resulting in 1 second
// for the example).
func parseTTL(s string) (ttlSeconds int64, err error) {
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

func getAliasFromAnnotations(annotations map[string]string) bool {
	aliasAnnotation, exists := annotations[aliasAnnotationKey]
	return exists && aliasAnnotation == "true"
}

// suitableType returns the DNS resource record type suitable for the target.
// In this case type A/AAAA for IPs and type CNAME for everything else.
func suitableType(target string) string {
	netIP, err := netip.ParseAddr(target)
	if err == nil && netIP.Is4() {
		return endpoint.RecordTypeA
	} else if err == nil && netIP.Is6() {
		return endpoint.RecordTypeAAAA
	}
	return endpoint.RecordTypeCNAME
}

// ParseIngress parses an ingress string in the format "namespace/name" or "name".
// It returns the namespace and name extracted from the string, or an error if the format is invalid.
// If the namespace is not provided, it defaults to an empty string.
func ParseIngress(ingress string) (namespace, name string, err error) {
	parts := strings.Split(ingress, "/")
	if len(parts) == 2 {
		namespace, name = parts[0], parts[1]
	} else if len(parts) == 1 {
		name = parts[0]
	} else {
		err = fmt.Errorf("invalid ingress name (name or namespace/name) found %q", ingress)
	}

	return
}

// SelectorMatchesServiceSelector checks if all key-value pairs in the selector map
// are present and match the corresponding key-value pairs in the svcSelector map.
// It returns true if all pairs match, otherwise it returns false.
func SelectorMatchesServiceSelector(selector, svcSelector map[string]string) bool {
	for k, v := range selector {
		if lbl, ok := svcSelector[k]; !ok || lbl != v {
			return false
		}
	}
	return true
}

// TargetsFromTargetAnnotation gets endpoints from optional "target" annotation.
// Returns empty endpoints array if none are found.
func TargetsFromTargetAnnotation(annotations map[string]string) endpoint.Targets {
	var targets endpoint.Targets

	// Get the desired hostname of the ingress from the annotation.
	targetAnnotation, exists := annotations[targetAnnotationKey]
	if exists && targetAnnotation != "" {
		// splits the hostname annotation and removes the trailing periods
		targetsList := strings.Split(strings.Replace(targetAnnotation, " ", "", -1), ",")
		for _, targetHostname := range targetsList {
			targetHostname = strings.TrimSuffix(targetHostname, ".")
			targets = append(targets, targetHostname)
		}
	}
	return targets
}

// ParseAnnotationFilter parses an annotation filter string into a labels.Selector.
// Returns nil if the annotation filter is invalid.
func ParseAnnotationFilter(annotationFilter string) (labels.Selector, error) {
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
