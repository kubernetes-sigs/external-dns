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
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/kubernetes/scheme"
	runtime "sigs.k8s.io/controller-runtime/pkg/client"
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
		ref     ObjectReference
		message string
		source  string
		action  Action
		eType   EventType
		reason  Reason
	}

	ObjectReference struct {
		Kind       string
		ApiVersion string
		Namespace  string
		Name       string
		UID        types.UID
		Source     string
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
		RefObject() *ObjectReference
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
		Kind:       gvk.Kind,
		ApiVersion: gvk.GroupVersion().String(),
		Namespace:  obj.GetNamespace(),
		Name:       obj.GetName(),
		UID:        obj.GetUID(),
		Source:     source,
	}
}

func NewEvent(obj *ObjectReference, msg string, a Action, r Reason) Event {
	if obj == nil {
		return Event{}
	}
	return Event{
		ref:     *obj,
		message: msg,
		eType:   EventTypeNormal,
		action:  a,
		reason:  r,
		source:  obj.Source,
	}
}

// NewEventFromEndpoint creates an Event from an EndpointInfo with formatted message.
func NewEventFromEndpoint(ep EndpointInfo, a Action, r Reason) Event {
	if ep == nil || ep.RefObject() == nil {
		return Event{}
	}
	msg := fmt.Sprintf("(external-dns) record:%s,owner:%s,type:%s,ttl:%d,targets:%s",
		ep.GetDNSName(), ep.GetOwner(), ep.GetRecordType(), ep.GetRecordTTL(),
		strings.Join(ep.GetTargets(), ","))
	return NewEvent(ep.RefObject(), msg, a, r)
}

func (e *Event) description() string {
	return fmt.Sprintf("%s/%s/%s", e.ref.Kind, e.ref.Namespace, e.ref.Name)
}

func (e *Event) Action() Action {
	return e.action
}

func (e *Event) Reason() Reason {
	return e.reason
}

func (e *Event) EventType() EventType {
	return e.eType
}

func (e *Event) event() *eventsv1.Event {
	if e.ref.Name == "" {
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
	namespace := e.ref.Namespace
	if namespace == "" {
		namespace = "default"
	}

	event := &eventsv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sanitize(e.ref.Name),
			Namespace: namespace,
		},
		EventTime:           timestamp,
		ReportingInstance:   controllerName + "/source/" + e.ref.Source,
		ReportingController: controllerName,
		Action:              string(e.action),
		Reason:              string(e.reason),
		Note:                message,
		Type:                string(e.eType),
	}
	if e.ref.UID != "" {
		ref := e.ref.objectRef()
		event.Related = ref
		event.Regarding = *ref
	}
	return event
}

// Sanitize input to comply with RFC 1123 subdomain naming requirements
func sanitize(input string) string {
	t := metav1.Time{Time: time.Now()}
	if input == "" {
		return fmt.Sprintf("a.%x", t.UnixNano())
	}
	sanitized := invalidChars.ReplaceAllString(strings.ToLower(input), "-")

	// the name should start with an alphanumeric character
	if len(sanitized) > 0 && !startsWithAlphaNumeric.MatchString(sanitized) {
		sanitized = "a" + sanitized
	}

	// the name should end with an alphanumeric character
	if len(sanitized) > 0 && !endsWithAlphaNumeric.MatchString(sanitized) {
		sanitized += "z"
	}

	sanitized = invalidChars.ReplaceAllString(sanitized, "-")

	return fmt.Sprintf("%v.%x", sanitized, t.UnixNano())
}

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

func (r *ObjectReference) objectRef() *apiv1.ObjectReference {
	return &apiv1.ObjectReference{
		Kind:       r.Kind,
		Namespace:  r.Namespace,
		Name:       r.Name,
		UID:        r.UID,
		APIVersion: r.ApiVersion,
	}
}
