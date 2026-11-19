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

// Package targetfrom resolves the target-from annotation: a record's targets
// rendered from another resource. See docs/proposal/007-target-from-resource.md.
package targetfrom

import (
	"context"
	"errors"
	"fmt"
	"net/netip"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/restmapper"
	"k8s.io/client-go/tools/cache"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
	"sigs.k8s.io/external-dns/source/informers"
	"sigs.k8s.io/external-dns/source/template"
)

const (
	// Each reset re-runs full discovery.
	mapperResetInterval = 30 * time.Second
	// Bounds a reconcile blocked on a reflector that cannot list.
	informerSyncTimeout = 30 * time.Second
)

// ClientsFunc returns the Resolver's clients. Called once, on first reference.
type ClientsFunc func() (dynamic.Interface, discovery.DiscoveryInterface, error)

type informerKey struct {
	gvr       schema.GroupVersionResource
	namespace string
}

// Resolver resolves target-from references through a dynamic client, with one
// lazily started informer per (GVR, namespace).
type Resolver struct {
	ctx        context.Context
	clientsFn  ClientsFunc
	clientOnce sync.Once
	clientErr  error
	client     dynamic.Interface
	mapper     meta.ResettableRESTMapper

	mu        sync.Mutex
	informers map[informerKey]kubeinformers.GenericInformer
	handlers  []func()
	lastReset time.Time
}

// NewResolver returns a Resolver whose informers live until ctx is done.
func NewResolver(ctx context.Context, clientsFn ClientsFunc) *Resolver {
	return &Resolver{
		ctx:       ctx,
		clientsFn: clientsFn,
		informers: map[informerKey]kubeinformers.GenericInformer{},
	}
}

// AddEventHandler registers handler on every informer, current and future, so
// a change to a referenced resource triggers a sync.
func (r *Resolver) AddEventHandler(handler func()) {
	if r == nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	r.handlers = append(r.handlers, handler)
	for _, inf := range r.informers {
		if inf == nil {
			continue
		}
		informers.MustAddEventHandler(inf.Informer(), eventHandler(handler))
	}
}

// TargetsFromAnnotations resolves the target-from annotation of obj. When ok,
// callers must use targets even if empty, never falling back to native discovery.
func (r *Resolver) TargetsFromAnnotations(ctx context.Context, kind string, obj metav1.Object) (endpoint.Targets, bool) {
	value, ok := obj.GetAnnotations()[annotations.TargetFromKey]
	if !ok {
		return nil, false
	}
	targets, err := r.Resolve(ctx, obj.GetNamespace(), value)
	if err != nil {
		log.Warnf("%s %s/%s: unresolved %s: %v", kind, obj.GetNamespace(), obj.GetName(), annotations.TargetFromKey, err)
		return nil, true
	}
	return targets, true
}

// Resolve resolves an annotation value for an object in ownerNamespace. Any
// error means the reference is unresolved.
func (r *Resolver) Resolve(ctx context.Context, ownerNamespace, value string) (endpoint.Targets, error) {
	if r == nil {
		return nil, errors.New("target-from is not supported by this source")
	}
	ref, err := annotations.ParseTargetFrom(value)
	if err != nil {
		return nil, err
	}
	if err := r.initClients(); err != nil {
		return nil, err
	}
	obj, err := r.get(ctx, ref, ownerNamespace)
	if err != nil {
		return nil, err
	}
	values, err := template.ExecUnstructured(ref.Template, obj.Object)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, errors.New("template rendered no targets, the field may not be populated yet")
	}
	for _, v := range values {
		if !isValidTarget(v) {
			// Not echoed: it can come from any field ExternalDNS can read.
			return nil, errors.New("rendered a value that is neither an IP address nor a hostname")
		}
	}
	return values, nil
}

func (r *Resolver) initClients() error {
	r.clientOnce.Do(func() {
		var dc discovery.DiscoveryInterface
		r.client, dc, r.clientErr = r.clientsFn()
		if r.clientErr == nil {
			r.mapper = restmapper.NewDeferredDiscoveryRESTMapper(memory.NewMemCacheClient(dc))
		}
	})
	return r.clientErr
}

func (r *Resolver) get(ctx context.Context, ref annotations.TargetFromReference, ownerNamespace string) (*unstructured.Unstructured, error) {
	gvr, namespaced, err := r.lookup(ref.Resource)
	if err != nil {
		return nil, err
	}
	namespace := ref.Namespace
	switch {
	case !namespaced && namespace != "":
		return nil, fmt.Errorf("%s is cluster-scoped, the namespace must be omitted", ref.Resource)
	case namespaced && namespace == "":
		namespace = ownerNamespace
	}

	inf := r.informerFor(ctx, informerKey{gvr: gvr, namespace: namespace})
	if inf == nil {
		return nil, fmt.Errorf("list and watch on %s in namespace %q are not granted, restart ExternalDNS once they are", gvr, namespace)
	}
	syncCtx, cancel := context.WithTimeout(ctx, informerSyncTimeout)
	defer cancel()
	if !cache.WaitForCacheSync(syncCtx.Done(), inf.Informer().HasSynced) {
		return nil, fmt.Errorf("informer for %s in namespace %q has not synced", gvr, namespace)
	}
	var obj runtime.Object
	if namespaced {
		obj, err = inf.Lister().ByNamespace(namespace).Get(ref.Name)
	} else {
		obj, err = inf.Lister().Get(ref.Name)
	}
	if err != nil {
		return nil, err
	}
	u, ok := obj.(*unstructured.Unstructured)
	if !ok {
		return nil, fmt.Errorf("unexpected object type %T", obj)
	}
	return u, nil
}

// lookup maps a group/resource to its preferred served version, like kubectl,
// so a CRD graduating to a new version does not break references.
func (r *Resolver) lookup(gr schema.GroupResource) (schema.GroupVersionResource, bool, error) {
	gvr, err := r.mapper.ResourceFor(gr.WithVersion(""))
	if meta.IsNoMatchError(err) && r.allowReset() {
		// The discovery cache never expires, and the mapper only retries on an
		// unpopulated one: reset it so a CRD installed after startup is found.
		r.mapper.Reset()
		gvr, err = r.mapper.ResourceFor(gr.WithVersion(""))
	}
	if err != nil {
		return schema.GroupVersionResource{}, false, err
	}
	// A template can rewrite any string into a tenant hostname: a Secret would leak through DNS.
	if gvr.Group == "" && gvr.Resource == "secrets" {
		return schema.GroupVersionResource{}, false, errors.New("secrets cannot be referenced")
	}
	gvk, err := r.mapper.KindFor(gvr)
	if err != nil {
		return schema.GroupVersionResource{}, false, err
	}
	mapping, err := r.mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return schema.GroupVersionResource{}, false, err
	}
	return gvr, mapping.Scope.Name() == meta.RESTScopeNameNamespace, nil
}

func (r *Resolver) allowReset() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	if time.Since(r.lastReset) < mapperResetInterval {
		return false
	}
	r.lastReset = time.Now()
	return true
}

// informerFor returns nil when list or watch is forbidden, as a reflector
// would log the denial forever.
func (r *Resolver) informerFor(ctx context.Context, key informerKey) kubeinformers.GenericInformer {
	r.mu.Lock()
	defer r.mu.Unlock()
	if inf, ok := r.informers[key]; ok {
		return inf
	}
	if err := r.canWatch(ctx, key); apierrors.IsForbidden(err) {
		r.informers[key] = nil
		return nil
	}
	inf := dynamicinformer.NewFilteredDynamicInformer(r.client, key.gvr, key.namespace, 0,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, nil)
	informers.MustSetTransform(inf.Informer(), informers.TransformerWithOptions[*unstructured.Unstructured](
		informers.TransformRemoveManagedFields(),
		informers.TransformRemoveLastAppliedConfig(),
	))
	for _, h := range r.handlers {
		informers.MustAddEventHandler(inf.Informer(), eventHandler(h))
	}
	go inf.Informer().RunWithContext(r.ctx)
	r.informers[key] = inf
	log.Debugf("Started target-from informer for %s in namespace %q", key.gvr, key.namespace)
	return inf
}

func (r *Resolver) canWatch(ctx context.Context, key informerKey) error {
	res := r.client.Resource(key.gvr).Namespace(key.namespace)
	list, err := res.List(ctx, metav1.ListOptions{Limit: 1})
	if err != nil {
		return err
	}
	// From the list's version, so existing objects are not replayed.
	w, err := res.Watch(ctx, metav1.ListOptions{ResourceVersion: list.GetResourceVersion()})
	if err != nil {
		return err
	}
	w.Stop()
	return nil
}

func isValidTarget(v string) bool {
	if _, err := netip.ParseAddr(v); err == nil {
		return true
	}
	return len(validation.IsFullyQualifiedDomainName(nil, strings.ToLower(v))) == 0
}

func eventHandler(fn func()) cache.ResourceEventHandler {
	return cache.ResourceEventHandlerFuncs{
		AddFunc:    func(any) { fn() },
		UpdateFunc: func(any, any) { fn() },
		DeleteFunc: func(any) { fn() },
	}
}
