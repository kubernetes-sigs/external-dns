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

package leaderelection

import (
	"context"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/uuid"
	restclient "k8s.io/client-go/rest"
	le "k8s.io/client-go/tools/leaderelection"
	rlock "k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
)

const inClusterNamespacePath = "/var/run/secrets/kubernetes.io/serviceaccount/namespace"

var (
	DefaultLockName      = "external-dns-leader-election"
	DefaultNamespace     = "default"
	DefaultLeaseDuration = 20 * time.Second // specifies how long the lease is valid
	DefaultRenewDeadline = 15 * time.Second // specifies the amount of time that the current node has to renew the lease before it expires
	DefaultRetryPeriod   = 5 * time.Second  // specifies the amount of time that the current holder of a lease has last updated the lease
)

type Manager struct {
	*restclient.Config
	rlock.Interface
	Namespace string
	LockName  string
	Identity  string
}

func NewLeaderElectionManager() (*Manager, error) {
	m := &Manager{}

	m.Namespace = getInClusterNamespace()
	m.LockName = DefaultLockName
	m.Identity = getLeaderId()

	cfg, err := ctrl.GetConfig()
	if err != nil {
		return nil, err
	}
	m.Config = cfg
	// Create a new lock. This will be used to create a Lease resource in the cluster.
	lock, err := rlock.NewFromKubeconfig(
		// Default resource lock to "leases".
		rlock.LeasesResourceLock,
		m.Namespace,
		m.LockName,
		rlock.ResourceLockConfig{
			Identity: m.Identity,
		},
		cfg,
		DefaultRenewDeadline,
	)
	if err != nil {
		return nil, err
	}
	m.Interface = lock
	return m, nil
}

func (m *Manager) ConfigureElection(run func(ctx context.Context)) (*le.LeaderElector, error) {
	el, err := le.NewLeaderElector(le.LeaderElectionConfig{
		Lock:            m.Interface,
		LeaseDuration:   DefaultLeaseDuration,
		RenewDeadline:   DefaultRenewDeadline,
		RetryPeriod:     DefaultRetryPeriod,
		ReleaseOnCancel: true, // release the lock when the context is canceled
		Name:            m.LockName,
		// election checker
		WatchDog: le.NewLeaderHealthzAdaptor(time.Second * 10),
		Callbacks: le.LeaderCallbacks{
			OnStartedLeading: run,
			OnStoppedLeading: func() {
				log.Info("leader election lost")
			},
			OnNewLeader: func(identity string) {
				log.Infof("the leader identity is '%s'\n", identity)
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return el, nil
}

// getInClusterNamespace retrieves the namespace in which the pod is running by reading the namespace file.
// If the namespace file does not exist or cannot be read, it falls back to the default namespace.
func getInClusterNamespace() string {
	// Check whether the namespace file exists.
	// If not, we are not running in cluster so can't guess the namespace.
	if _, err := os.Stat(inClusterNamespacePath); err != nil {
		log.Debugf("not running in-cluster, fallback to %s namespace\n", DefaultNamespace)
		return DefaultNamespace
	}

	// Load the namespace file and return its content
	namespace, err := os.ReadFile(inClusterNamespacePath)
	if err != nil {
		log.Debugf("not running in-cluster, fallback to %s namespace\n", DefaultNamespace)
		return DefaultLockName
	}
	return string(namespace)
}

// getLeaderId generates a unique leader ID by combining the hostname and a UUID.
// If the hostname cannot be retrieved, it falls back to using "external-dns" with a UUID.
func getLeaderId() string {
	// Leader id, needs to be unique
	id, err := os.Hostname()
	if err != nil {
		return fmt.Sprintf("external-dns_%s", uuid.NewUUID())
	}
	id = fmt.Sprintf("%s_%s", id, uuid.NewUUID())
	return id
}
