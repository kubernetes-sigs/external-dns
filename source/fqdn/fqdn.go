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

package fqdn

import (
	"bytes"
	"fmt"
	"net/netip"
	"strings"
	"text/template"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"sigs.k8s.io/external-dns/endpoint"
)

func ParseTemplate(input string) (*template.Template, error) {
	if input == "" {
		return nil, nil
	}
	funcs := template.FuncMap{
		"contains":   strings.Contains,
		"trimPrefix": strings.TrimPrefix,
		"trimSuffix": strings.TrimSuffix,
		"trim":       strings.TrimSpace,
		"toLower":    strings.ToLower,
		"replace":    replace,
		"isIPv6":     isIPv6String,
		"isIPv4":     isIPv4String,
	}
	return template.New("endpoint").Funcs(funcs).Parse(input)
}

type kubeObject interface {
	runtime.Object
	metav1.Object
}

func ExecTemplate(tmpl *template.Template, obj kubeObject) ([]string, error) {
	if obj == nil {
		return nil, fmt.Errorf("object is nil")
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, obj); err != nil {
		kind := obj.GetObjectKind().GroupVersionKind().Kind
		return nil, fmt.Errorf("failed to apply template on %s %s/%s: %w", kind, obj.GetNamespace(), obj.GetName(), err)
	}
	hosts := strings.Split(buf.String(), ",")
	hostnames := make([]string, 0, len(hosts))
	for _, name := range hosts {
		name = strings.TrimSpace(name)
		name = strings.TrimSuffix(name, ".")
		if name != "" {
			hostnames = append(hostnames, name)
		}
	}
	return hostnames, nil
}

// replace all instances of oldValue with newValue in target string.
// adheres to syntax from https://masterminds.github.io/sprig/strings.html.
func replace(oldValue, newValue, target string) string {
	return strings.ReplaceAll(target, oldValue, newValue)
}

// isIPv6String reports whether the target string is an IPv6 address,
// including IPv4-mapped IPv6 addresses.
func isIPv6String(target string) bool {
	netIP, err := netip.ParseAddr(target)
	if err != nil {
		return false
	}
	return netIP.Is6()
}

// isIPv4String reports whether the target string is an IPv4 address.
func isIPv4String(target string) bool {
	netIP, err := netip.ParseAddr(target)
	if err != nil {
		return false
	}
	return netIP.Is4()
}

// CombineWithTemplatedEndpoints merges annotation-based endpoints with template-based endpoints
// according to the FQDN template configuration.
//
// Logic:
//   - If fqdnTemplate is nil, returns original endpoints unchanged
//   - If combineFQDNAnnotation is true, appends templated endpoints to existing
//   - If combineFQDNAnnotation is false and endpoints is empty, uses templated endpoints
//   - If combineFQDNAnnotation is false and endpoints exist, returns original unchanged
func CombineWithTemplatedEndpoints(
	endpoints []*endpoint.Endpoint,
	fqdnTemplate *template.Template,
	combineFQDNAnnotation bool,
	templateFunc func() ([]*endpoint.Endpoint, error),
) ([]*endpoint.Endpoint, error) {
	if fqdnTemplate == nil {
		return endpoints, nil
	}

	if !combineFQDNAnnotation && len(endpoints) > 0 {
		return endpoints, nil
	}

	templatedEndpoints, err := templateFunc()
	if err != nil {
		return nil, err
	}

	if combineFQDNAnnotation {
		return append(endpoints, templatedEndpoints...), nil
	}
	return templatedEndpoints, nil
}
