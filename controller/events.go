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

package controller

import (
	"strings"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/events"
	"sigs.k8s.io/external-dns/plan"
)

// Delete events are reported by the registry: the source object is usually gone.
const registrySource = "registry"

// resourceKinds maps resource label prefixes to Kinds, for `involvedObject.kind=` selectors.
var resourceKinds = map[string]string{
	"service":         "Service",
	"ingress":         "Ingress",
	"pod":             "Pod",
	"node":            "Node",
	"crd":             "DNSEndpoint",
	"httproute":       "HTTPRoute",
	"grpcroute":       "GRPCRoute",
	"tcproute":        "TCPRoute",
	"udproute":        "UDPRoute",
	"tlsroute":        "TLSRoute",
	"gateway":         "Gateway",
	"virtualservice":  "VirtualService",
	"route":           "Route",
	"ingressroute":    "IngressRoute",
	"ingressroutetcp": "IngressRouteTCP",
	"ingressrouteudp": "IngressRouteUDP",
}

// emitChangeEvent emits a Kubernetes event for each DNS record change.
// Deletes use RecordDeleted on success and RecordError on failure.
func emitChangeEvent(e events.EventEmitter, ch *plan.Changes, reason events.Reason) {
	if e == nil {
		return
	}
	for _, ep := range ch.Create {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionCreate, reason))
	}
	for _, ep := range ch.UpdateNew {
		e.Add(events.NewEventFromEndpoint(ep, events.ActionUpdate, reason))
	}
	deleteReason := events.RecordDeleted
	if reason == events.RecordError {
		deleteReason = events.RecordError
	}
	for _, ep := range ch.Delete {
		ref := refFromResourceLabel(ep.Labels[endpoint.ResourceLabelKey])
		if ref == nil {
			continue
		}
		e.Add(events.NewEventFromEndpoint(deletedEndpoint{Endpoint: ep, ref: ref}, events.ActionDelete, deleteReason))
	}
}

// deletedEndpoint attaches a ref rebuilt from the resource label, as registry records carry none.
type deletedEndpoint struct {
	*endpoint.Endpoint
	ref *events.ObjectReference
}

func (d deletedEndpoint) RefObjects() []*events.ObjectReference {
	return []*events.ObjectReference{d.ref}
}

// refFromResourceLabel parses "kind/namespace/name" or "kind/name". No UID: the object may be gone.
func refFromResourceLabel(resource string) *events.ObjectReference {
	parts := strings.Split(resource, "/")
	var kind, namespace, name string
	switch len(parts) {
	case 2:
		kind, name = parts[0], parts[1]
	case 3:
		kind, namespace, name = parts[0], parts[1], parts[2]
	default:
		return nil
	}
	if kind == "" || name == "" {
		return nil
	}
	if k, ok := resourceKinds[strings.ToLower(kind)]; ok {
		kind = k
	}
	return events.NewObjectReferenceFromParts(kind, "", namespace, name, "", registrySource)
}
