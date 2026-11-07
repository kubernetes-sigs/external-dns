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

package annotations

import (
	"fmt"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TargetFromReference is a parsed target-from annotation value:
// <resource>[.<group>]/[<namespace>/]<name>#<template>.
type TargetFromReference struct {
	Resource schema.GroupResource
	// Namespace is empty when omitted; it then defaults to the annotated
	// object's namespace for namespaced resources.
	Namespace string
	Name      string
	Template  string
}

// ParseTargetFrom parses a target-from annotation value. The template is everything
// after the first '#', taken verbatim, so it needs no escaping.
func ParseTargetFrom(value string) (TargetFromReference, error) {
	ref, tmpl, ok := strings.Cut(value, "#")
	if !ok {
		return TargetFromReference{}, fmt.Errorf("missing '#<template>' in %q", value)
	}
	if strings.TrimSpace(tmpl) == "" {
		return TargetFromReference{}, fmt.Errorf("empty template in %q", value)
	}
	parts := strings.Split(ref, "/")
	var r TargetFromReference
	switch len(parts) {
	case 2:
		r.Name = parts[1]
	case 3:
		r.Namespace, r.Name = parts[1], parts[2]
		if r.Namespace == "" {
			return TargetFromReference{}, fmt.Errorf("empty namespace in %q", value)
		}
	default:
		return TargetFromReference{}, fmt.Errorf("expected <resource>[.<group>]/[<namespace>/]<name>#<template>, got %q", value)
	}
	r.Resource = schema.ParseGroupResource(parts[0])
	if r.Resource.Resource == "" || r.Name == "" {
		return TargetFromReference{}, fmt.Errorf("empty resource or name in %q", value)
	}
	r.Template = tmpl
	return r, nil
}
