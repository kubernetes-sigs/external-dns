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
	"context"

	log "github.com/sirupsen/logrus"
	eventsv1 "k8s.io/api/events/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	v1 "k8s.io/client-go/kubernetes/typed/events/v1"
	"k8s.io/client-go/util/workqueue"
)

const (
	workers            = 1
	controllerName     = "external-dns"
	maxRetriesPerEvent = 3
	maxQueuedEvents    = 100
)

type EventEmitter interface {
	Add(...Event)
}

type Controller struct {
	client          v1.EventsV1Interface
	queue           workqueue.TypedRateLimitingInterface[*eventsv1.Event]
	emitEvents      sets.Set[Reason]
	maxQueuedEvents int
	createOpts      metav1.CreateOptions
}

func NewEventController(client v1.EventsV1Interface, cfg *Config) (*Controller, error) {
	queue := workqueue.NewTypedRateLimitingQueueWithConfig(
		workqueue.DefaultTypedControllerRateLimiter[*eventsv1.Event](),
		workqueue.TypedRateLimitingQueueConfig[*eventsv1.Event]{Name: controllerName},
	)
	createOpts := metav1.CreateOptions{}
	if cfg.dryRun {
		createOpts.DryRun = []string{metav1.DryRunAll}
	}
	return &Controller{
		client:          client,
		queue:           queue,
		emitEvents:      cfg.emitEvents,
		maxQueuedEvents: maxQueuedEvents,
		createOpts:      createOpts,
	}, nil
}

func (ec *Controller) Run(ctx context.Context) {
	if len(ec.emitEvents) == 0 {
		return
	}
	go ec.run(ctx)
}

func (ec *Controller) run(ctx context.Context) {
	log.Info("event Controller started")
	defer log.Info("event Controller terminated")
	defer utilruntime.HandleCrash()
	var waitGroup wait.Group
	for range workers {
		waitGroup.StartWithContext(ctx, func(ctx context.Context) {
			for ec.processNextWorkItem(ctx) {
			}
		})
	}
	<-ctx.Done()
	ec.queue.ShutDownWithDrain()
	waitGroup.Wait()
}

func (ec *Controller) processNextWorkItem(ctx context.Context) bool {
	event, quit := ec.queue.Get()
	if quit {
		return false
	}
	defer ec.queue.Done(event)
	_, err := ec.client.Events(event.Namespace).Create(ctx, event, ec.createOpts)
	switch {
	case apierrors.IsNotFound(err):
		log.Warnf("dropping event %s/%s: namespace not found. %v", event.Namespace, event.Name, err)
	case err != nil && ec.queue.NumRequeues(event) < maxRetriesPerEvent:
		log.Errorf("not able to create event %s/%s, retrying. %v", event.Namespace, event.Name, err)
		ec.queue.AddRateLimited(event)
		return true
	case err != nil:
		log.Errorf("dropping event %s/%s after %d retries. %v", event.Namespace, event.Name, ec.queue.NumRequeues(event), err)
	}
	ec.queue.Forget(event)
	return true
}

func (ec *Controller) Add(events ...Event) {
	dropped := 0
	for _, e := range events {
		if ec.queue.Len() >= ec.maxQueuedEvents {
			dropped++
			continue
		}
		event := e.event()
		if event == nil {
			continue
		}
		ec.emit(event)
	}
	if dropped > 0 {
		log.Warnf("event queue is full, dropped %d events", dropped)
	}
}

func (ec *Controller) emit(event *eventsv1.Event) {
	if !ec.emitEvents.Has(Reason(event.Reason)) {
		log.Debugf("skipping event %s/%s/%s with reason %s as not configured to emit", event.Kind, event.Namespace, event.Name, event.Reason)
		return
	}
	ec.queue.Add(event)
}
