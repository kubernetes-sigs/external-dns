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
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/core/v1"
	eventsv1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes/fake"
	eventsclient "k8s.io/client-go/kubernetes/typed/events/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"

	clienttesting "k8s.io/client-go/testing"
)

func TestNewEventController_Success(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {}))
	defer svr.Close()

	mockKubeCfgDir := filepath.Join(t.TempDir(), ".kube")
	mockKubeCfgPath := filepath.Join(mockKubeCfgDir, "config")
	err := os.MkdirAll(mockKubeCfgDir, 0755)
	require.NoError(t, err)

	kubeCfgTemplate := `
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: %s
  name: test-cluster
contexts:
- context:
    cluster: test-cluster
    user: test-user
  name: test-context
current-context: test-context
users:
- name: test-user
  user:
    token: fake-token
`
	err = os.WriteFile(mockKubeCfgPath, fmt.Appendf(nil, kubeCfgTemplate, svr.URL), os.FileMode(0755))
	require.NoError(t, err)

	restConfig, err := clientcmd.BuildConfigFromFlags(svr.URL, mockKubeCfgPath)
	require.NoError(t, err)
	client, err := eventsclient.NewForConfig(restConfig)
	require.NoError(t, err)

	cfg := NewConfig(
		WithEmitEvents([]string{string(RecordReady)}),
	)
	ctrl, err := NewEventController(client, cfg)
	require.NoError(t, err)
	require.NotNil(t, ctrl)
	require.False(t, ctrl.dryRun)
}

func TestController_Run_NoEmitEvents(t *testing.T) {
	kClient := fake.NewClientset()
	ctrl := &Controller{
		client:     kClient.EventsV1(),
		emitEvents: sets.New[Reason](),
	}

	require.NotPanics(t, func() {
		ctrl.Run(t.Context())
	})
}

func TestController_Run_EmitEvents(t *testing.T) {
	log.SetLevel(log.ErrorLevel)
	ctx := t.Context()

	eventCreated := make(chan struct{})
	kubeClient := fake.NewClientset()
	kubeClient.PrependReactor("create", "events", func(_ clienttesting.Action) (bool, runtime.Object, error) {
		eventCreated <- struct{}{}
		return true, nil, nil
	})

	eventsClient := kubeClient.EventsV1()
	ctrl := &Controller{
		client:     eventsClient,
		emitEvents: sets.New[Reason](RecordReady),
		queue: workqueue.NewTypedRateLimitingQueueWithConfig[any](
			workqueue.DefaultTypedControllerRateLimiter[any](),
			workqueue.TypedRateLimitingQueueConfig[any]{Name: controllerName},
		),
		hostname:        controllerName,
		maxQueuedEvents: 1,
	}

	ctrl.Run(ctx)

	event := NewEvent(NewObjectReference(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-object",
			Namespace: v1.NamespaceDefault,
			UID:       "9de3fc19-8aeb-4e76-865d-ada955403103",
		},
	}, "fake-source"), "record created", ActionCreate, RecordReady)

	ctrl.Add(event)

	select {
	case <-eventCreated:
	case <-time.After(wait.ForeverTestTimeout):
		t.Fatal("event not created")
	}
}

func TestController_Queue_EmitEvents(t *testing.T) {
	log.SetLevel(log.ErrorLevel)

	eventsClient := fake.NewClientset().EventsV1()
	ctrl := &Controller{
		client:     eventsClient,
		emitEvents: sets.New[Reason](RecordReady),
		queue: workqueue.NewTypedRateLimitingQueueWithConfig[any](
			workqueue.DefaultTypedControllerRateLimiter[any](),
			workqueue.TypedRateLimitingQueueConfig[any]{Name: controllerName},
		),
		hostname:        controllerName,
		maxQueuedEvents: 1,
	}

	event := NewEvent(NewObjectReference(&v1.Pod{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Pod",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "fake-object",
			Namespace: v1.NamespaceDefault,
			UID:       "9de3fc19-8aeb-4e76-865d-ada955403103",
		},
	}, "fake-source"), "record created", ActionCreate, RecordReady)

	ctrl.Add(event)

	queueItem, shutdown := ctrl.queue.Get()
	require.False(t, shutdown)
	value, ok := queueItem.(*eventsv1.Event)
	assert.True(t, ok)
	assert.NotNil(t, value)

	assert.Contains(t, value.Name, "fake-object.")
	assert.Contains(t, value.Reason, RecordReady)
}
