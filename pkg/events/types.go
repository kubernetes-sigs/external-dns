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

package events

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	runtime "sigs.k8s.io/controller-runtime/pkg/client"

	"sigs.k8s.io/external-dns/internal/sets"
)

const (
	ActionCreate  Action = "Created"
	ActionUpdate  Action = "Updated"
	ActionDelete  Action = "Deleted"
	ActionFailed  Action = "FailedSync"
	RecordReady   Reason = "RecordReady"
	RecordDeleted Reason = "RecordDeleted"
	RecordError   Reason = "RecordError"

	EventTypeNormal  EventType = EventType(apiv1.EventTypeNormal)
	EventTypeWarning EventType = EventType(apiv1.EventTypeWarning)
)

var (
	invalidChars           = regexp.MustCompile(`[^a-z0-9.\-]`)
	startsWithAlphaNumeric = regexp.MustCompile(`^[a-z0-9]`)
	endsWithAlphaNumeric   = regexp.MustCompile(`[a-z0-9]$`)
)

type (
	// Action values for actions
	Action string
	// Reason types of Event Reasons
	Reason string
	// EventType values for event types
	EventType    string
	ConfigOption func(*Config)

	Event struct {
		refs    []ObjectReference
		message string
		action  Action
		eType   EventType
		reason  Reason
	}

	// ObjectReference holds metadata about a Kubernetes object for event correlation.
	ObjectReference struct {
		kind       string
		apiVersion string
		namespace  string
		name       string
		uid        types.UID
		source     string
	}

	Config struct {
		emitEvents sets.Set[Reason]
		dryRun     bool
	}

	// EndpointInfo defines the interface for endpoint data needed to create events.
	EndpointInfo interface {
		GetDNSName() string
		GetRecordType() string
		GetRecordTTL() int64
		GetTargets() []string
		GetOwner() string
		RefObjects() []*ObjectReference
	}
)

func NewObjectReference(obj runtime.Object, source string) *ObjectReference {
	// Kubernetes API doesn't populate TypeMeta (Kind/APIVersion) when retrieving
	// objects via informers. Look up the Kind from the scheme without mutating the object.
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Kind == "" {
		gvks, _, err := scheme.Scheme.ObjectKinds(obj)
		if err == nil && len(gvks) > 0 {
			gvk = gvks[0]
		} else {
			// Fallback to reflection for types not in scheme
			gvk = schema.GroupVersionKind{Kind: reflect.TypeOf(obj).Elem().Name()}
		}
	}
	return &ObjectReference{
		kind:       gvk.Kind,
		apiVersion: gvk.GroupVersion().String(),
		namespace:  obj.GetNamespace(),
		name:       obj.GetName(),
		uid:        obj.GetUID(),
		source:     source,
	}
}

func NewEvent(obj *ObjectReference, msg string, a Action, r Reason) Event {
	if obj == nil {
		return Event{}
	}
	return Event{
		refs:    []ObjectReference{*obj},
		message: msg,
		eType:   EventTypeNormal,
		action:  a,
		reason:  r,
	}
}

// NewEventFromEndpoint creates an Event from an EndpointInfo with formatted message.
// All ref objects on the endpoint are stored in the event; one Kubernetes event is
// emitted per ref when the event is processed by the Controller.
func NewEventFromEndpoint(ep EndpointInfo, a Action, r Reason) Event {
	if ep == nil {
		return Event{}
	}
	var refs []ObjectReference
	for _, ref := range ep.RefObjects() {
		if ref != nil {
			refs = append(refs, *ref)
		}
	}
	if len(refs) == 0 {
		return Event{}
	}
	msg := fmt.Sprintf("(external-dns) record:%s,owner:%s,type:%s,ttl:%d,targets:%s",
		ep.GetDNSName(), ep.GetOwner(), ep.GetRecordType(), ep.GetRecordTTL(),
		strings.Join(ep.GetTargets(), ","))
	return Event{
		refs:    refs,
		message: msg,
		eType:   EventTypeNormal,
		action:  a,
		reason:  r,
	}
}

// Action returns the action associated with the event (e.g. Created, Updated, Deleted).
func (e *Event) Action() Action {
	return e.action
}

// Reason returns the reason associated with the event (e.g. RecordReady, RecordError).
func (e *Event) Reason() Reason {
	return e.reason
}

// EventType returns the Kubernetes event type (Normal or Warning).
func (e *Event) EventType() EventType {
	return e.eType
}

// events returns one Kubernetes event per ref stored in the Event.
func (e *Event) events() []*eventsv1.Event {
	result := make([]*eventsv1.Event, 0, len(e.refs))
	for _, ref := range e.refs {
		if ev := e.eventForRef(ref); ev != nil {
			result = append(result, ev)
		}
	}
	return result
}

func (e *Event) eventForRef(ref ObjectReference) *eventsv1.Event {
	if ref.name == "" {
		log.Debug("skipping event for resources as the name is not generated yet")
		return nil
	}
	message := e.message
	// https://github.com/kubernetes/api/blob/7da28ad7db85e33ab8be3b89e63cad4c07b37fb2/events/v1/types.go#L77
	if len(message) > 1024 {
		message = message[0:1021] + "..."
	}

	timestamp := metav1.MicroTime{Time: time.Now()}

	// Events are namespaced resources. For cluster-scoped objects like Nodes,
	// the namespace is empty, so we default to "default" namespace.
	namespace := ref.namespace
	if namespace == "" {
		namespace = "default"
	}

	event := &eventsv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sanitize(ref.name, timestamp.Time),
			Namespace: namespace,
		},
		EventTime:           timestamp,
		ReportingInstance:   controllerName + "/source/" + ref.source,
		ReportingController: controllerName,
		Action:              string(e.action),
		Reason:              string(e.reason),
		Note:                message,
		Type:                string(e.eType),
	}
	objRef := ref.objectRef()
	event.Regarding = *objRef
	if ref.uid != "" {
		event.Related = objRef
	}
	return event
}

// sanitize converts input to a valid RFC 1123 subdomain name, appending a hex
// timestamp suffix for uniqueness. t should be the same timestamp used for EventTime.
func sanitize(input string, t time.Time) string {
	suffix := fmt.Sprintf(".%x", t.UnixNano())
	sanitized := invalidChars.ReplaceAllString(strings.ToLower(input), "-")

	// the name should start with an alphanumeric character
	if !startsWithAlphaNumeric.MatchString(sanitized) {
		sanitized = "a" + sanitized
	}

	// truncate to leave room for the suffix, keeping total ≤ 253 (RFC 1123 subdomain limit)
	if maxBase := 253 - len(suffix); len(sanitized) > maxBase {
		sanitized = sanitized[:maxBase]
	}

	// the name should end with an alphanumeric character
	if !endsWithAlphaNumeric.MatchString(sanitized) {
		sanitized += "z"
	}

	return sanitized + suffix
}

// WithDryRun returns a ConfigOption that sets dry-run mode; events are skipped when enabled.
func WithDryRun(dryRun bool) ConfigOption {
	return func(c *Config) {
		c.dryRun = dryRun
	}
}

func WithEmitEvents(events []string) ConfigOption {
	return func(c *Config) {
		if len(events) > 0 {
			c.emitEvents = sets.New[Reason]()
			for _, event := range events {
				if slices.Contains([]string{string(RecordReady), string(RecordError)}, event) {
					c.emitEvents.Insert(Reason(event))
				}
			}
		}
	}
}

func NewConfig(opts ...ConfigOption) *Config {
	c := &Config{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *Config) IsEnabled() bool {
	return len(c.emitEvents) > 0
}

// description returns a human-readable summary of the reference.
// Namespaced resources use "kind/namespace/name"; cluster-scoped resources omit the namespace: "kind/name".
func (r *ObjectReference) description() string {
	if r.namespace == "" {
		return fmt.Sprintf("%s/%s", r.kind, r.name)
	}
	return fmt.Sprintf("%s/%s/%s", r.kind, r.namespace, r.name)
}

// Key returns a stable string that uniquely identifies this object reference
// in the form "source/namespace/name".
func (r *ObjectReference) Key() string {
	return r.source + "/" + r.namespace + "/" + r.name
}

// Kind returns the Kubernetes kind of the referenced object (e.g. "Service", "Ingress").
func (r *ObjectReference) Kind() string {
	return r.kind
}

// Namespace returns the namespace of the referenced Kubernetes object.
func (r *ObjectReference) Namespace() string {
	return r.namespace
}

// Name returns the name of the referenced Kubernetes object.
func (r *ObjectReference) Name() string {
	return r.name
}

// Source returns the source identifier of the ObjectReference (e.g. "ingress", "service").
func (r *ObjectReference) Source() string {
	return r.source
}

// UID returns the UID of the referenced Kubernetes object.
func (r *ObjectReference) UID() types.UID {
	return r.uid
}

func (r *ObjectReference) objectRef() *apiv1.ObjectReference {
	return &apiv1.ObjectReference{
		Kind:       r.kind,
		Namespace:  r.namespace,
		Name:       r.name,
		UID:        r.uid,
		APIVersion: r.apiVersion,
	}
}

// NewObjectReferenceFromParts constructs an ObjectReference directly from its components.
// Use this when you don't have a runtime.Object (e.g., in tests or when constructing
// references from non-Kubernetes data sources).
func NewObjectReferenceFromParts(kind, apiVersion, namespace, name string, uid types.UID, source string) *ObjectReference {
	return &ObjectReference{
		kind:       kind,
		apiVersion: apiVersion,
		namespace:  namespace,
		name:       name,
		uid:        uid,
		source:     source,
	}
}
