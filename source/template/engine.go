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

package template

import (
	"bytes"
	"fmt"
	"maps"
	"reflect"
	"slices"
	"strings"
	"text/template"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"

	"sigs.k8s.io/external-dns/endpoint"
)

// Engine holds the parsed Go templates used to derive DNS names and targets
// from Kubernetes objects. It is shared across source implementations.
// The zero value is valid and represents a no-op engine.
type Engine struct {
	// fqdn is the template that generates fully-qualified domain names from a Kubernetes object.
	// Parsed from --fqdn-template.
	fqdn *template.Template
	// target is the optional template that overrides the DNS target values.
	// Parsed from --target-template.
	target *template.Template
	// fqdnTarget is the optional template that generates "hostname:target" pairs in a single pass,
	// superseding fqdn and target when set.
	// Parsed from --fqdn-target-template.
	fqdnTarget *template.Template
	// combine controls whether template-derived endpoints are merged with annotation-derived endpoints.
	// Set by --combine-fqdn-annotation.
	combine bool
}

// NewEngine parses the provided Go template strings into a Engine.
// An empty string leaves the corresponding template unset; IsConfigured reflects
// whether the FQDN template was provided. Returns an error on the first parse failure.
func NewEngine(fqdnStr, targetStr, fqdnTargetStr string, combineFQDN bool) (Engine, error) {
	fqdnTmpl, err := parseTemplate(fqdnStr)
	if err != nil {
		return Engine{}, fmt.Errorf("parse --fqdn-template: %w", err)
	}
	targetTmpl, err := parseTemplate(targetStr)
	if err != nil {
		return Engine{}, fmt.Errorf("parse --target-template: %w", err)
	}
	fqdnTargetTmpl, err := parseTemplate(fqdnTargetStr)
	if err != nil {
		return Engine{}, fmt.Errorf("parse --fqdn-target-template: %w", err)
	}
	return Engine{
		fqdn:       fqdnTmpl,
		target:     targetTmpl,
		fqdnTarget: fqdnTargetTmpl,
		combine:    combineFQDN,
	}, nil
}

// IsConfigured reports whether the FQDN template is set and ready to use.
func (e Engine) IsConfigured() bool {
	return e.fqdn != nil
}

// Combining reports whether the engine is configured to combine template-based
// endpoints with annotation-based endpoints.
func (e Engine) Combining() bool {
	return e.combine
}

// ExecFQDN executes the FQDN template against a Kubernetes object and returns hostnames.
func (e Engine) ExecFQDN(obj kubeObject) ([]string, error) {
	return execTemplate(e.fqdn, obj)
}

// ExecTarget executes the Target template against a Kubernetes object and returns targets.
func (e Engine) ExecTarget(obj kubeObject) ([]string, error) {
	return execTemplate(e.target, obj)
}

// ExecFQDNTarget executes the FQDNTarget template against a Kubernetes object and returns hostname:target pairs.
func (e Engine) ExecFQDNTarget(obj kubeObject) ([]string, error) {
	return execTemplate(e.fqdnTarget, obj)
}

// CombineWithEndpoints merges annotation-based endpoints with template-based endpoints.
func (e Engine) CombineWithEndpoints(
	endpoints []*endpoint.Endpoint,
	templateFunc func() ([]*endpoint.Endpoint, error),
) ([]*endpoint.Endpoint, error) {
	if e.fqdn == nil && e.target == nil && e.fqdnTarget == nil {
		return endpoints, nil
	}

	if !e.combine && len(endpoints) > 0 {
		return endpoints, nil
	}

	templatedEndpoints, err := templateFunc()
	if err != nil {
		return nil, fmt.Errorf("failed to get endpoints from template: %w", err)
	}

	if e.combine {
		return append(endpoints, templatedEndpoints...), nil
	}
	return templatedEndpoints, nil
}

func parseTemplate(input string) (*template.Template, error) {
	if strings.TrimSpace(input) == "" {
		return nil, nil //nolint:nilnil // nil template signals "not configured"; callers check IsConfigured()
	}
	// Clone is cheaper than re-registering all functions on a new template each call.
	t, err := baseTemplate.Clone()
	if err != nil {
		return nil, err
	}
	if _, err = t.Parse(input); err != nil {
		return nil, fmt.Errorf("%q: %w", input, err)
	}
	return t, nil
}

type kubeObject interface {
	runtime.Object
	metav1.Object
}

func execTemplate(tmpl *template.Template, obj kubeObject) ([]string, error) {
	if tmpl == nil {
		return []string{}, nil
	}
	if obj == nil {
		return nil, fmt.Errorf("object is nil")
	}
	// Kubernetes API doesn't populate TypeMeta (Kind/APIVersion) when retrieving
	// objects via informers, because the client already knows what type it requested.
	// Set it so templates can use .Kind and .APIVersion.
	// TODO: all sources to transform Informer().SetTransform()
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Kind == "" {
		gvks, _, err := scheme.Scheme.ObjectKinds(obj)
		if err == nil && len(gvks) > 0 {
			gvk = gvks[0]
		} else {
			// Fallback to reflection for types not in scheme
			gvk = schema.GroupVersionKind{Kind: reflect.TypeOf(obj).Elem().Name()}
		}
		obj.GetObjectKind().SetGroupVersionKind(gvk)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, obj); err != nil {
		kind := obj.GetObjectKind().GroupVersionKind().Kind
		return nil, fmt.Errorf("failed to apply template on %s %s/%s: %w", kind, obj.GetNamespace(), obj.GetName(), err)
	}
	hosts := strings.Split(buf.String(), ",")
	hostnames := make(map[string]struct{}, len(hosts))
	for _, name := range hosts {
		name = strings.TrimSpace(name)
		name = strings.TrimSuffix(name, ".")
		if name != "" {
			hostnames[name] = struct{}{}
		}
	}
	return slices.Sorted(maps.Keys(hostnames)), nil
}
