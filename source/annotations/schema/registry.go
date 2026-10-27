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

package schema

import (
	"slices"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/types"
)

const defaultStrictMessage = "the resource will not receive a DNS record"

// defaultSpecs lists each supported annotation, the sources that consume it,
// and its validation rules. New(opts) builds a Schema around this list.
var defaultSpecs = []AnnotationSpec{
	{
		KeySuffix:        annotations.AccessSuffix,
		SupportedSources: []string{types.Service},
		Validators:       []Validator{ValidateOneOf("public", "private")},
		WarnMessage:      "the service still gets a DNS record either way, just with a different IP source",
		Documentation:    `Specifies whether the public or private interface address is used for headless/NodePort services. Accepted values: "public", "private".`,
	},
	{
		KeySuffix:        annotations.EndpointsTypeSuffix,
		SupportedSources: []string{types.Service},
		Validators:       []Validator{ValidateOneOf("NodeExternalIP", "HostIP")},
		WarnMessage:      "the service falls back to default headless endpoint publishing behavior (pod IPs unless --publish-host-ip is set)",
		StrictMessage:    "the service will not receive any DNS records",
		Documentation:    `Specifies the type of endpoints to use for headless services. Accepted values: "NodeExternalIP", "HostIP".`,
	},
	{
		KeySuffix: annotations.RecordTypeSuffix,
		SupportedSources: []string{
			types.AmbassadorHost, types.ContourHTTPProxy,
			types.GatewayHttpRoute, types.GatewayGrpcRoute, types.GatewayTlsRoute, types.GatewayTcpRoute, types.GatewayUdpRoute,
			types.GlooProxy,
			types.Ingress,
			types.IstioGateway, types.IstioVirtualService,
			types.KongTCPIngress,
			types.OpenShiftRoute,
			types.Service,
			types.SkipperRouteGroup,
			types.TraefikProxy,
		},
		Validators:    []Validator{ValidateOneOfFold("", endpoint.RecordTypePTR)},
		WarnMessage:   `the annotation does not control the DNS record type; only "ptr" has an effect (opting in to PTR record generation) and "" opts out`,
		StrictMessage: defaultStrictMessage,
		Documentation: `Opts the resource in or out of PTR record generation (--create-ptr). Accepted values: "ptr" to opt in, "" to opt out. See docs/advanced/ptr-records.md.`,
	},
	{
		KeySuffix: annotations.SetIdentifierSuffix,
		SupportedSources: []string{
			types.AmbassadorHost, types.ContourHTTPProxy,
			types.GatewayHttpRoute, types.GatewayGrpcRoute, types.GatewayTlsRoute, types.GatewayTcpRoute, types.GatewayUdpRoute,
			types.GlooProxy,
			types.Ingress,
			types.IstioGateway, types.IstioVirtualService,
			types.KongTCPIngress,
			types.OpenShiftRoute,
			types.Service,
			types.SkipperRouteGroup,
			types.TraefikProxy,
		},
		Validators:    []Validator{ValidateNonEmpty},
		WarnMessage:   "the set-identifier provider-specific property is ignored for this resource",
		StrictMessage: defaultStrictMessage,
		Documentation: "Distinguishes between multiple records with the same DNS name in routing policies (e.g. weighted, latency, or failover routing).",
	},
	{
		KeySuffix: annotations.TtlSuffix,
		SupportedSources: []string{
			types.AmbassadorHost, types.ContourHTTPProxy, types.F5TransportServer, types.F5VirtualServer,
			types.GatewayHttpRoute, types.GatewayGrpcRoute, types.GatewayTlsRoute, types.GatewayTcpRoute, types.GatewayUdpRoute,
			types.GlooProxy,
			types.Ingress,
			types.IstioGateway, types.IstioVirtualService,
			types.KongTCPIngress,
			types.Node,
			types.OpenShiftRoute,
			types.Pod,
			types.Service,
			types.SkipperRouteGroup,
			types.TraefikProxy,
			types.Unstructured,
		},
		Validators:    []Validator{ValidateTTL},
		WarnMessage:   "the TTL falls back to the provider's default",
		Documentation: `Specifies the TTL for the resource's DNS records, as an integer number of seconds or a Go duration string (e.g. "10m"). Must be between 1 and 2147483647 seconds.`,
	},
	{
		KeySuffix: annotations.HostnameSuffix,
		SupportedSources: []string{
			types.ContourHTTPProxy,
			types.GatewayHttpRoute, types.GatewayGrpcRoute, types.GatewayTlsRoute, types.GatewayTcpRoute, types.GatewayUdpRoute,
			types.Ingress,
			types.IstioGateway, types.IstioVirtualService,
			types.KongTCPIngress,
			types.OpenShiftRoute,
			types.Pod,
			types.Service,
			types.SkipperRouteGroup,
			types.TraefikProxy,
			types.Unstructured,
		},
		Validators:    []Validator{ValidateHostnames},
		WarnMessage:   "the hostnames are still published as-is and the provider may reject the invalid ones",
		StrictMessage: defaultStrictMessage,
		Documentation: "Comma-separated list of desired hostnames for the resource. Each must be a valid DNS name (RFC 1123).",
	},
	{
		KeySuffix:        annotations.InternalHostnameSuffix,
		SupportedSources: []string{types.Pod, types.Service},
		Validators:       []Validator{ValidateHostnames},
		WarnMessage:      "the hostnames are still published as-is and the provider may reject the invalid ones",
		StrictMessage:    defaultStrictMessage,
		Documentation:    "Comma-separated list of desired internal hostnames for the resource. Each must be a valid DNS name (RFC 1123).",
	},
	{
		KeySuffix:        annotations.NodePortFilterSuffix,
		SupportedSources: []string{types.Service},
		Validators:       []Validator{ValidateLabelSelector},
		WarnMessage:      "the service's NodePort records are not filtered by node labels",
		// No StrictMessage: an invalid selector falls back to every node, so even strict mode keeps the service.
		Documentation: "Kubernetes label selector limiting which Nodes a NodePort service's records point at, e.g. \"!karpenter.sh/nodepool\".",
	},
	{
		KeySuffix:        annotations.IngressHostnameSourceSuffix,
		SupportedSources: []string{types.Ingress},
		Validators:       []Validator{ValidateOneOfFold("defined-hosts-only", "annotation-only")},
		StrictMessage:    "the ingress will not receive any DNS records",
		Documentation:    `Controls which hostnames are used for the ingress: "defined-hosts-only" or "annotation-only" (case-insensitive).`,
	},
}

// Options configures a Schema at construction time. Fields correspond to runtime
// flags that change what the schema validates.
type Options struct {
	// Mode is the active validation mode. ModeOff disables all validation.
	Mode Mode
	// IgnoreHostnameAnnotation mirrors --ignore-hostname-annotation: when true,
	// IsValid skips the hostname (and, for the service source, internal-hostname)
	// annotation checks, because the source will not read those values at runtime.
	IgnoreHostnameAnnotation bool
}

// Schema validates annotations against a fixed list of AnnotationSpecs under a
// fixed set of Options. Constructed once and shared across sources; each source
// binds a per-call identity via WithSource.
type Schema struct {
	specs                    []AnnotationSpec
	mode                     Mode
	ignoreHostnameAnnotation bool
	source                   string
}

// WithSource returns a shallow copy of s bound to the given source name
// (e.g. types.Ingress). IsValid is a no-op until a source is bound.
func (s *Schema) WithSource(name string) *Schema {
	if s == nil {
		return nil
	}
	cp := *s
	cp.source = name
	return &cp
}

// New returns a Schema using the default spec list and the given options.
func New(opts Options) *Schema {
	return NewWithSpecs(defaultSpecs, opts)
}

// NewWithSpecs returns a Schema with an explicit spec list — intended for tests
// and future schema extensions that supply their own specs.
func NewWithSpecs(specs []AnnotationSpec, opts Options) *Schema {
	return &Schema{
		specs:                    specs,
		mode:                     opts.Mode,
		ignoreHostnameAnnotation: opts.IgnoreHostnameAnnotation,
	}
}

// Mode reports the Schema's active validation mode.
func (s *Schema) Mode() Mode {
	if s == nil {
		return ModeOff
	}
	return s.mode
}

// IsValid returns false iff an invalid or unsupported-for-source annotation is
// in strict mode. A nil, ModeOff, or source-unbound Schema always returns true.
// Warn-mode violations are logged only.
func (s *Schema) IsValid(entity metav1.Object) bool {
	if s == nil || s.mode == "" || s.mode == ModeOff || s.source == "" {
		return true
	}
	annots := entity.GetAnnotations()
	desc := describe(entity)
	valid := true
	for _, e := range s.specs {
		key := e.Key()
		value, ok := annots[key]
		if !ok {
			continue
		}
		if s.shouldSkipKey(key) {
			continue
		}
		if len(e.SupportedSources) > 0 && !slices.Contains(e.SupportedSources, s.source) {
			if e.IsStrict(s.mode) {
				log.Warnf("Excluding %s: annotation %s is not supported for source %q. %s", desc, key, s.source, e.StrictMessage)
				valid = false
			} else {
				log.Debugf("%s: annotation %s is not supported for source %q", desc, key, s.source)
			}
			continue
		}
		err := e.Validate(value)
		if err == nil {
			continue
		}
		if e.IsStrict(s.mode) {
			log.Warnf("Excluding %s: %q is not a valid %s value: %v. %s", desc, value, key, err, e.StrictMessage)
			valid = false
			continue
		}
		log.Warnf("%s: %q is not a valid %s value: %v. %s", desc, value, key, err, e.warnText())
	}
	return valid
}

// Service short-circuits internal-hostname too; other sources only gate hostname.
func (s *Schema) shouldSkipKey(key string) bool {
	if !s.ignoreHostnameAnnotation {
		return false
	}
	if key == annotations.HostnameKey {
		return true
	}
	return s.source == types.Service && key == annotations.InternalHostnameKey
}

func describe(entity metav1.Object) string {
	if ns := entity.GetNamespace(); ns != "" {
		return ns + "/" + entity.GetName()
	}
	return entity.GetName()
}
