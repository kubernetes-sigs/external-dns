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

package registry

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

var (
	errInvalidHeritage = errors.New("heritage is unknown or not found")
)

// known keys
const (
	heritage = "external-dns"
)

// serializeLabel transforms endpoints labels into a external-dns format string
func serializeLabel(labels map[string]string, surroundQuotes bool) string {
	var tokens []string
	tokens = append(tokens, fmt.Sprintf("heritage=%s", heritage))
	var keys []string
	for key := range labels {
		keys = append(keys, key)
	}
	sort.Strings(keys) // sort for consistency

	for _, key := range keys {
		tokens = append(tokens, fmt.Sprintf("%s/%s=%s", heritage, key, labels[key]))
	}
	if surroundQuotes {
		return fmt.Sprintf("\"%s\"", strings.Join(tokens, ","))
	}
	return strings.Join(tokens, ",")
}

// deserializeLabel constructs endpoints labels from a provided format string
// if heritage set to another value is found then error is returned
// no heritage automatically assumes is not owned by external-dns
func deserializeLabel(labelText string) (map[string]string, error) {
	endpointLabels := map[string]string{}
	labelText = strings.Trim(labelText, "\"") // drop quotes
	tokens := strings.Split(labelText, ",")
	foundExternalDNSHeritage := false
	for _, token := range tokens {
		if len(strings.Split(token, "=")) != 2 {
			continue
		}
		key := strings.Split(token, "=")[0]
		val := strings.Split(token, "=")[1]
		if key == "heritage" && val != heritage {
			return nil, errInvalidHeritage
		}
		if key == "heritage" {
			foundExternalDNSHeritage = true
			continue
		}
		if strings.HasPrefix(key, heritage) {
			endpointLabels[strings.TrimPrefix(key, heritage+"/")] = val
		}
	}

	if !foundExternalDNSHeritage {
		return nil, errInvalidHeritage
	}

	return endpointLabels, nil
}
