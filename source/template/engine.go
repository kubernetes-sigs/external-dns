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
	"reflect"
	"strings"
	"text/template"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes/scheme"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/sets"
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
func NewEngine(fqdnTemplates, targetTemplates, fqdnTargetTemplates []string, combineFQDN bool) (Engine, error) {
	fqdnTmpl, err := validateAndParse(fqdnTemplates, "--fqdn-template")
	if err != nil {
		return Engine{}, err
	}
	targetTmpl, err := validateAndParse(targetTemplates, "--target-template")
	if err != nil {
		return Engine{}, err
	}
	fqdnTargetTmpl, err := validateAndParse(fqdnTargetTemplates, "--fqdn-target-template")
	if err != nil {
		return Engine{}, err
	}
	return Engine{fqdn: fqdnTmpl, target: targetTmpl, fqdnTarget: fqdnTargetTmpl, combine: combineFQDN}, nil
}

func validateAndParse(templates []string, flag string) (*template.Template, error) {
	if err := validateTemplates(templates, flag); err != nil {
		return nil, err
	}
	if log.IsLevelEnabled(log.DebugLevel) {
		log.Debugf("%s: %s", flag, strings.Join(templates, ","))
	}
	t, err := parseTemplate(strings.Join(templates, ","))
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", flag, err)
	}
	return t, nil
}

// IsConfigured reports whether the FQDN template is set and ready to use.
func (e Engine) IsConfigured() bool {
	return e.fqdn != nil
}

// HasDNSNameTemplate reports whether any name-generating template (fqdn or fqdn-target) is
// configured. The target template alone cannot generate DNS names, so it is excluded.
// Use this to decide whether template output should replace resource-derived DNS names.
func (e Engine) HasDNSNameTemplate() bool {
	return e.fqdn != nil || e.fqdnTarget != nil
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

func (e Engine) execTarget(obj kubeObject) ([]string, error) {
	return execTemplate(e.target, obj)
}

func (e Engine) execFQDNTarget(obj kubeObject) ([]string, error) {
	return execTemplate(e.fqdnTarget, obj)
}

// ApplyFQDNTargetTemplate combines existing endpoints with those derived from the fqdn-target
// template (host:target pairs). Used standalone by sources whose second template pass derives
// targets from the resource itself (pod IPs, service ClusterIP, node addresses).
func (e Engine) ApplyFQDNTargetTemplate(existing []*endpoint.Endpoint, obj kubeObject) ([]*endpoint.Endpoint, error) {
	return e.CombineWithEndpoints(existing, func() ([]*endpoint.Endpoint, error) {
		return e.endpointsFromFQDNTargetTemplate(obj)
	})
}

// ApplyTemplates runs both template passes in sequence — fqdn-target first, then
// fqdn+target — returning the combined result. Use this for sources where the template
// itself provides the DNS target (gloo, traefik, f5, unstructured). Sources that derive
// targets from the resource (pod, service, node) call ApplyFQDNTargetTemplate and then
// their own resource-specific CombineWithEndpoints.
func (e Engine) ApplyTemplates(existing []*endpoint.Endpoint, obj kubeObject) ([]*endpoint.Endpoint, error) {
	eps, err := e.ApplyFQDNTargetTemplate(existing, obj)
	if err != nil {
		return nil, err
	}
	return e.CombineWithEndpoints(eps, func() ([]*endpoint.Endpoint, error) {
		return e.endpointsFromTemplate(obj)
	})
}

func (e Engine) endpointsFromFQDNTargetTemplate(obj kubeObject) ([]*endpoint.Endpoint, error) {
	pairs, err := e.execFQDNTarget(obj)
	if err != nil || len(pairs) == 0 {
		return nil, err
	}
	kind := strings.ToLower(obj.GetObjectKind().GroupVersionKind().Kind)
	eps := make([]*endpoint.Endpoint, 0, len(pairs))
	for _, pair := range pairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) != 2 {
			log.Debugf("Skipping invalid host:target pair %q from %s %s/%s: missing ':' separator",
				pair, kind, obj.GetNamespace(), obj.GetName())
			continue
		}
		host := strings.TrimSpace(parts[0])
		target := strings.TrimSpace(parts[1])
		if host == "" || target == "" {
			log.Debugf("Skipping incomplete host:target pair %q from %s %s/%s: field may not yet be populated",
				pair, kind, obj.GetNamespace(), obj.GetName())
			continue
		}
		eps = append(eps, endpoint.NewEndpoint(host, endpoint.SuitableType(target), target))
	}
	return endpoint.MergeEndpoints(eps), nil
}

func (e Engine) endpointsFromTemplate(obj kubeObject) ([]*endpoint.Endpoint, error) {
	hostnames, err := e.ExecFQDN(obj)
	if err != nil || len(hostnames) == 0 {
		return nil, err
	}
	targets, err := e.execTarget(obj)
	if err != nil {
		return nil, err
	}
	return endpoint.EndpointsForHostsAndTargets(hostnames, targets), nil
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

// validateTemplates validates each template string individually for syntax errors,
// then checks cumulatively for cross-value {{ define }} block conflicts.
// Duplicate and blank strings are silently skipped.
// validateTemplates parses templates cumulatively so that both syntax errors and
// cross-value {{ define }} block conflicts are caught. Go only reports redefinition
// when both definitions appear in the same Parse call.
func validateTemplates(templates []string, flagName string) error {
	var joined []string
	for i, tmpl := range templates {
		joined = append(joined, tmpl)
		t, err := baseTemplate.Clone()
		if err != nil {
			return err
		}
		if _, err = t.Parse(strings.Join(joined, ",")); err != nil {
			return fmt.Errorf("%s[%d] %q: %w", flagName, i, tmpl, err)
		}
	}
	return nil
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
	hostnames := make(sets.Set[string], len(hosts))
	for _, name := range hosts {
		name = strings.TrimSpace(name)
		name = strings.TrimSuffix(name, ".")
		if name != "" {
			hostnames.Insert(name)
		}
	}
	return sets.Sorted(hostnames), nil
}
