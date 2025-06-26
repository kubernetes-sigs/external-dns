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
	"unicode"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func ParseTemplate(input string) (tmpl *template.Template, err error) {
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
	var hostnames []string
	for _, name := range strings.Split(buf.String(), ",") {
		name = strings.TrimFunc(name, unicode.IsSpace)
		name = strings.TrimSuffix(name, ".")
		hostnames = append(hostnames, name)
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
