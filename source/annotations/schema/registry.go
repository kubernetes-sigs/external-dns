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

// Registry is external-dns's contract for which annotations are supported by which
// sources, along with their validation rules and warn/strict messages.
var Registry = []AnnotationSpec{
	{
		Key:              annotations.AccessKey,
		SupportedSources: []string{types.Service},
		Config: Config{
			Validators:    []Validator{ValidateOneOf("public", "private")},
			WarnMessage:   "the service still gets a DNS record either way, just with a different IP source",
			Documentation: `Specifies whether the public or private interface address is used for headless/NodePort services. Accepted values: "public", "private".`,
		},
	},
	{
		Key:              annotations.EndpointsTypeKey,
		SupportedSources: []string{types.Service},
		Config: Config{
			Validators:    []Validator{ValidateOneOf("NodeExternalIP", "HostIP")},
			WarnMessage:   "the service falls back to default headless endpoint publishing behavior (pod IPs unless --publish-host-ip is set)",
			StrictMessage: "the service will not receive any DNS records",
			Documentation: `Specifies the type of endpoints to use for headless services. Accepted values: "NodeExternalIP", "HostIP".`,
		},
	},
	{
		Key: annotations.RecordTypeKey,
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
		Config: Config{
			Validators: []Validator{ValidateOneOf(
				endpoint.RecordTypeA,
				endpoint.RecordTypeAAAA,
				endpoint.RecordTypeCNAME,
				endpoint.RecordTypeTXT,
				endpoint.RecordTypeSRV,
				endpoint.RecordTypeNS,
				endpoint.RecordTypePTR,
				endpoint.RecordTypeMX,
				endpoint.RecordTypeNAPTR,
				endpoint.RecordTypeDNAME,
			)},
			WarnMessage:   "the record-type provider-specific property is ignored for this resource",
			StrictMessage: "the resource will not receive a DNS record",
			Documentation: "Overrides the DNS record type for the resource's endpoints. Accepted values: A, AAAA, CNAME, TXT, SRV, NS, PTR, MX, NAPTR, DNAME.",
		},
	},
	{
		Key: annotations.SetIdentifierKey,
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
		Config: Config{
			Validators:    []Validator{ValidateNonEmpty},
			WarnMessage:   "the set-identifier provider-specific property is ignored for this resource",
			StrictMessage: "the resource will not receive a DNS record",
			Documentation: "Distinguishes between multiple records with the same DNS name in routing policies (e.g. weighted, latency, or failover routing).",
		},
	},
	{
		Key: annotations.TtlKey,
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
		Config: Config{
			Validators:    []Validator{ValidateTTL},
			WarnMessage:   "the TTL falls back to the provider's default",
			Documentation: `Specifies the TTL for the resource's DNS records, as an integer number of seconds or a Go duration string (e.g. "10m"). Must be between 1 and 2147483647 seconds.`,
		},
	},
	{
		Key: annotations.HostnameKey,
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
		Config: Config{
			Validators:    []Validator{ValidateHostnames},
			WarnMessage:   "the malformed hostname list is skipped entirely",
			StrictMessage: "the resource will not receive a DNS record",
			Documentation: "Comma-separated list of desired hostnames for the resource. Each must be a valid DNS name (RFC 1123).",
		},
	},
	{
		Key:              annotations.InternalHostnameKey,
		SupportedSources: []string{types.Pod, types.Service},
		Config: Config{
			Validators:    []Validator{ValidateHostnames},
			WarnMessage:   "the malformed internal hostname list is skipped entirely",
			StrictMessage: "the resource will not receive a DNS record",
			Documentation: "Comma-separated list of desired internal hostnames for the resource. Each must be a valid DNS name (RFC 1123).",
		},
	},
	{
		Key:              annotations.IngressHostnameSourceKey,
		SupportedSources: []string{types.Ingress},
		Config: Config{
			Validators:    []Validator{ValidateOneOfFold("defined-hosts-only", "annotation-only")},
			StrictMessage: "the ingress will not receive any DNS records",
			Documentation: `Controls which hostnames are used for the ingress: "defined-hosts-only" or "annotation-only" (case-insensitive).`,
		},
	},
}

// IsValid checks entity's annotations against Registry for source and mode. In warn
// mode it only logs; in strict mode an invalid or unsupported annotation excludes the
// object from the informer's index.
func IsValid[T metav1.Object](entity T, source string, mode Mode) bool {
	annots := entity.GetAnnotations()
	desc := describe(entity)
	valid := true
	for _, e := range Registry {
		value, ok := annots[e.Key]
		if !ok {
			continue
		}
		if len(e.SupportedSources) > 0 && !slices.Contains(e.SupportedSources, source) {
			if e.IsStrict(mode) {
				log.Warnf("Excluding %s: annotation %s is not supported for source %q. %s", desc, e.Key, source, e.StrictMessage)
				valid = false
			} else {
				log.Debugf("%s: annotation %s is not supported for source %q", desc, e.Key, source)
			}
			continue
		}
		err := e.Validate(value)
		if err == nil {
			continue
		}
		if e.IsStrict(mode) {
			log.Warnf("Excluding %s: %q is not a valid %s value: %v. %s", desc, value, e.Key, err, e.StrictMessage)
			valid = false
			continue
		}
		log.Warnf("%s: %q is not a valid %s value: %v. %s", desc, value, e.Key, err, e.warnText())
	}
	return valid
}

// describe returns a "namespace/name" identifier for entity, or just "name" for
// cluster-scoped objects (empty namespace), for use in log messages.
func describe(entity metav1.Object) string {
	if ns := entity.GetNamespace(); ns != "" {
		return ns + "/" + entity.GetName()
	}
	return entity.GetName()
}
