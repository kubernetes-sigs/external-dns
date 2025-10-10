/*
Copyright 2014 The Kubernetes Authors.

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

package fake

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
=======
	"context"

>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	core "k8s.io/client-go/testing"
)

// Deprecated: use CreateWithEventNamespaceWithContext instead.
func (c *fakeEvents) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	return c.CreateWithEventNamespaceWithContext(context.Background(), event)
}

func (c *fakeEvents) CreateWithEventNamespaceWithContext(_ context.Context, event *v1.Event) (*v1.Event, error) {
	var action core.CreateActionImpl
	if c.Namespace() != "" {
		action = core.NewCreateAction(c.Resource(), c.Namespace(), event)
	} else {
		action = core.NewCreateAction(c.Resource(), event.GetNamespace(), event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
//
// Deprecated: use UpdateWithEventNamespaceWithContext instead.
func (c *fakeEvents) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	return c.UpdateWithEventNamespaceWithContext(context.Background(), event)
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
func (c *fakeEvents) UpdateWithEventNamespaceWithContext(_ context.Context, event *v1.Event) (*v1.Event, error) {
	var action core.UpdateActionImpl
	if c.Namespace() != "" {
		action = core.NewUpdateAction(c.Resource(), c.Namespace(), event)
	} else {
		action = core.NewUpdateAction(c.Resource(), event.GetNamespace(), event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// PatchWithEventNamespace patches an existing event. Returns the copy of the event the server returns, or an error.
// TODO: Should take a PatchType as an argument probably.
//
// Deprecated: use PatchWithEventNamespaceWithContext instead.
func (c *fakeEvents) PatchWithEventNamespace(event *v1.Event, data []byte) (*v1.Event, error) {
	return c.PatchWithEventNamespaceWithContext(context.Background(), event, data)
}

// PatchWithEventNamespaceWithContext patches an existing event. Returns the copy of the event the server returns, or an error.
// TODO: Should take a PatchType as an argument probably.
func (c *fakeEvents) PatchWithEventNamespaceWithContext(_ context.Context, event *v1.Event, data []byte) (*v1.Event, error) {
	// TODO: Should be configurable to support additional patch strategies.
	pt := types.StrategicMergePatchType
	var action core.PatchActionImpl
	if c.Namespace() != "" {
		action = core.NewPatchAction(c.Resource(), c.Namespace(), event.Name, pt, data)
	} else {
		action = core.NewPatchAction(c.Resource(), event.GetNamespace(), event.Name, pt, data)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Search returns a list of events matching the specified object.
//
// Deprecated: use SearchWithContext instead.
func (c *fakeEvents) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	return c.SearchWithContext(context.Background(), scheme, objOrRef)
}

// SearchWithContext returns a list of events matching the specified object.
func (c *fakeEvents) SearchWithContext(_ context.Context, scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	var action core.ListActionImpl
	if c.Namespace() != "" {
		action = core.NewListAction(c.Resource(), c.Kind(), c.Namespace(), metav1.ListOptions{})
	} else {
<<<<<<< HEAD
		action = core.NewListAction(eventsResource, eventsKind, v1.NamespaceDefault, metav1.ListOptions{})
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"k8s.io/api/core/v1"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"k8s.io/api/core/v1"
=======
	v1 "k8s.io/api/core/v1"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	types "k8s.io/apimachinery/pkg/types"
	core "k8s.io/client-go/testing"
)

func (c *FakeEvents) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	var action core.CreateActionImpl
	if c.ns != "" {
		action = core.NewCreateAction(eventsResource, c.ns, event)
	} else {
		action = core.NewCreateAction(eventsResource, event.GetNamespace(), event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
func (c *FakeEvents) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	var action core.UpdateActionImpl
	if c.ns != "" {
		action = core.NewUpdateAction(eventsResource, c.ns, event)
	} else {
		action = core.NewUpdateAction(eventsResource, event.GetNamespace(), event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// PatchWithEventNamespace patches an existing event. Returns the copy of the event the server returns, or an error.
// TODO: Should take a PatchType as an argument probably.
func (c *FakeEvents) PatchWithEventNamespace(event *v1.Event, data []byte) (*v1.Event, error) {
	// TODO: Should be configurable to support additional patch strategies.
	pt := types.StrategicMergePatchType
	var action core.PatchActionImpl
	if c.ns != "" {
		action = core.NewPatchAction(eventsResource, c.ns, event.Name, pt, data)
	} else {
		action = core.NewPatchAction(eventsResource, event.GetNamespace(), event.Name, pt, data)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Search returns a list of events matching the specified object.
func (c *FakeEvents) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	var action core.ListActionImpl
	if c.ns != "" {
		action = core.NewListAction(eventsResource, eventsKind, c.ns, metav1.ListOptions{})
<<<<<<< HEAD
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	} else {
		action = core.NewListAction(eventsResource, eventsKind, v1.NamespaceDefault, metav1.ListOptions{})
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
		action = core.NewListAction(eventsResource, eventsKind, v1.NamespaceDefault, metav1.ListOptions{})
=======
		action = core.NewListAction(c.Resource(), c.Kind(), v1.NamespaceDefault, metav1.ListOptions{})
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	}
	obj, err := c.Fake.Invokes(action, &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.EventList), err
}

func (c *fakeEvents) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	action := core.GenericActionImpl{}
	action.Verb = "get-field-selector"
	action.Resource = c.Resource()

	c.Fake.Invokes(action, nil)
	return fields.Everything()
}
