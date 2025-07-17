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
	"regexp"
	"slices"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	runtime "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	recordPrefix        = "ExternalDns"
	ActionCreate Action = "resource created"
	ActionUpdate Action = "resource modified"
	ActionDelete Action = "resource deleted"
	RecordReady  Reason = "RecordReady"
	RecordError  Reason = "RecordError"
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
)

type Event struct {
	involvedObject ObjectReference
	message        string
	action         Action
	eType          string
	reason         Reason
	source         string
}

type ObjectReference struct {
	Kind       string
	ApiVersion string
	Namespace  string
	Name       string
	UID        types.UID
	Source     string
}

type ConfigOption func(*Config)
type Config struct {
	kubeConfig   string
	apiServerURL string
	emitEvents   sets.Set[string]
	dryRun       bool
}

func NewObjectReference(obj runtime.Object, source string) *ObjectReference {
	return &ObjectReference{
		Kind:       obj.GetObjectKind().GroupVersionKind().Kind,
		ApiVersion: obj.GetObjectKind().GroupVersionKind().GroupVersion().String(),
		Namespace:  obj.GetNamespace(),
		Name:       obj.GetName(),
		Source:     source,
	}
}

func NewEvent(obj *ObjectReference, msg string, a Action, r Reason) Event {
	if obj == nil {
		return Event{}
	}
	return Event{
		involvedObject: *obj,
		message:        msg,
		eType:          corev1.EventTypeNormal,
		action:         a,
		reason:         r,
		source:         obj.Source,
	}
}

func (i *Event) Reference() string {
	return fmt.Sprintf("%s/%s/%s", i.involvedObject.Kind, i.involvedObject.Namespace, i.involvedObject.Name)
}

// Sanitize input to comply with RFC 1123 subdomain naming requirements
func sanitize(input string) string {
	t := metav1.Time{Time: time.Now()}
	if input == "" {
		return fmt.Sprintf("a.%x", t.UnixNano())
	}
	sanitized := invalidChars.ReplaceAllString(strings.ToLower(input), "-")

	// the name should starts with an alphanumeric character
	if len(sanitized) > 0 && !startsWithAlphaNumeric.MatchString(sanitized) {
		sanitized = "a" + sanitized
	}

	// the name should end with an alphanumeric character
	if len(sanitized) > 0 && !endsWithAlphaNumeric.MatchString(sanitized) {
		sanitized = sanitized + "z"
	}

	sanitized = invalidChars.ReplaceAllString(sanitized, "-")

	return fmt.Sprintf("%v.%x", sanitized, t.UnixNano())
}

func (i *Event) transpose() *eventsv1.Event {
	if i.involvedObject.Name == "" {
		log.Debug("skipping event for resources as the name is not generated yet")
		return nil
	}
	message := i.message
	// https://github.com/kubernetes/api/blob/7da28ad7db85e33ab8be3b89e63cad4c07b37fb2/events/v1/types.go#L77
	if len(message) > 1024 {
		message = message[0:1021] + "..."
	}

	timestamp := metav1.MicroTime{Time: time.Now()}

	event := &eventsv1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sanitize(i.involvedObject.Name),
			Namespace: i.involvedObject.Namespace,
		},
		EventTime: timestamp,
		Series:    nil,
		// Related: &corev1.ObjectReference{
		// 	Kind:      i.involvedObject.Kind,
		// 	Namespace: i.involvedObject.Namespace,
		// 	Name:      i.involvedObject.Name,
		// 	UID:       i.involvedObject.UID,
		// },
		Regarding: corev1.ObjectReference{
			Kind:      i.involvedObject.Kind,
			Namespace: i.involvedObject.Namespace,
			Name:      i.involvedObject.Name,
			UID:       i.involvedObject.UID,
		},
		ReportingInstance:   "external-dns-controller",
		ReportingController: controllerName,
		Action:              string(i.action),
		Reason:              recordPrefix + string(i.reason),
		Note:                message,
		Type:                i.eType,
	}

	return event
}

func WithKubeConfig(kubeConfig string) ConfigOption {
	return func(c *Config) {
		c.kubeConfig = kubeConfig
	}
}

func WithAPIServerURL(apiServerURL string) ConfigOption {
	return func(c *Config) {
		c.apiServerURL = apiServerURL
	}
}

func WithDryRun(dryRun bool) ConfigOption {
	return func(c *Config) {
		c.dryRun = dryRun
	}
}

func WithEmitEvents(events []string) ConfigOption {
	return func(c *Config) {
		if len(events) > 0 {
			c.emitEvents = sets.New[string]()
			for _, event := range events {
				if slices.Contains([]string{string(RecordReady), string(RecordError)}, event) {
					c.emitEvents.Insert(recordPrefix + event)
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
