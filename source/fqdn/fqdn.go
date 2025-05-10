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
	"net/netip"
	"strings"
	"text/template"
)

func ParseTemplate(fqdnTemplate string) (tmpl *template.Template, err error) {
	if fqdnTemplate == "" {
		return nil, nil
	}
	funcs := template.FuncMap{
		"trimPrefix": strings.TrimPrefix,
		"replace":    replace,
		"isIPv6":     isIPv6String,
		"isIPv4":     isIPv4String,
	}
	return template.New("endpoint").Funcs(funcs).Parse(fqdnTemplate)
}

// replace all instances of oldValue with newValue in target string.
// adheres to syntax from https://masterminds.github.io/sprig/strings.html.
func replace(oldValue, newValue, target string) string {
	//lint:ignore QF1004 Using strings.Replace due to argument order requirements
	return strings.Replace(target, oldValue, newValue, -1)
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
