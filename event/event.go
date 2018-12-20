package event

import (
	"fmt"
	"sync"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	client kubernetes.Interface
	kubeOnce sync.Once
)

// Set the client if it has not been set already
func InitializeClient(kc kubernetes.Interface) {
	kubeOnce.Do(func() {
		client = kc
	})
}

// Sends an event using kubeClient
func Emit(msg string) error {
	if client == nil {
		return nil
	}
	var err error
	now := metav1.NewTime(time.Now())
	event := &v1.Event{
		ObjectMeta: metav1.ObjectMeta{
			Name: fmt.Sprintf("%v.%x", "external-dns-event", time.Now().Nanosecond()),
			Namespace: "default",
		},
		Reason: "Started",
		Message: msg,
		FirstTimestamp: now,
		LastTimestamp:  now,
		Source: v1.EventSource{
			Component: "external-dns-event",
		},
		Count: 1,
		Type: v1.EventTypeNormal,
	}

	_, err = client.CoreV1().Events("default").Create(event)
	if err != nil {
		return err
	}
	return nil
}
