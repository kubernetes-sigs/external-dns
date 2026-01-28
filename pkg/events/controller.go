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
	"os"

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
	workers          = 1
	controllerName   = "external-dns"
	maxTriesPerEvent = 3
	maxQueuedEvents  = 100
)

type EventEmitter interface {
	Add(...Event)
}

type Controller struct {
	client          v1.EventsV1Interface
	queue           workqueue.TypedRateLimitingInterface[any]
	emitEvents      sets.Set[Reason]
	maxQueuedEvents int
	dryRun          bool
	hostname        string
}

func NewEventController(client v1.EventsV1Interface, cfg *Config) (*Controller, error) {
	queue := workqueue.NewTypedRateLimitingQueueWithConfig[any](
		workqueue.DefaultTypedControllerRateLimiter[any](),
		workqueue.TypedRateLimitingQueueConfig[any]{Name: controllerName},
	)
	hostname, _ := os.Hostname()
	return &Controller{
		client:          client,
		queue:           queue,
		emitEvents:      cfg.emitEvents,
		maxQueuedEvents: maxQueuedEvents,
		dryRun:          cfg.dryRun,
		hostname:        hostname,
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
	key, quit := ec.queue.Get()
	if quit {
		return false
	}
	defer ec.queue.Done(key)
	event, ok := key.(*eventsv1.Event)
	if !ok {
		log.Errorf("failed to convert key to Event: %q", key)
		return true
	}
	var dryRun []string
	if ec.dryRun {
		dryRun = []string{metav1.DryRunAll}
	}
	_, err := ec.client.Events(event.Namespace).Create(ctx, event, metav1.CreateOptions{
		DryRun: dryRun,
	})
	if err != nil && !apierrors.IsNotFound(err) {
		if ec.queue.NumRequeues(key) < maxTriesPerEvent {
			log.Errorf("not able to create event, retrying for key/%s. %v", key, err)
			return true
		}
		log.Errorf("dropping event %s/%s with key/%q after %d retries. %v", event.Namespace, event.Name, key, ec.queue.NumRequeues(key), err)
	}
	ec.queue.Forget(key)
	return true
}

func (ec *Controller) Add(events ...Event) {
	if ec.queue.Len() >= ec.maxQueuedEvents {
		log.Warnf("event queue is full, dropping %d events", len(events))
		return
	}
	for _, e := range events {
		event := e.event()
		if event == nil {
			continue
		}
		ec.emit(event)
	}
}

func (ec *Controller) emit(event *eventsv1.Event) {
	if !ec.emitEvents.Has(Reason(event.Reason)) {
		log.Debugf("skipping event %s/%s/%s with reason %s as not configured to emit", event.Kind, event.Namespace, event.Name, event.Reason)
		return
	}
	event.ReportingController = controllerName
	ec.queue.Add(event)
}
