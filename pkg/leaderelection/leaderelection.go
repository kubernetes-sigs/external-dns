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
	"time"

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rs "k8s.io/client-go/rest"
	le "k8s.io/client-go/tools/leaderelection"
	rl "k8s.io/client-go/tools/leaderelection/resourcelock"
	ctrl "sigs.k8s.io/controller-runtime"
)

var (
	DefaultLockName      = "external-dns-leader-election"
	DefaultNamespace     = metav1.NamespaceDefault
	DefaultLeaseDuration = 20 * time.Second // specifies how long the lease is valid
	DefaultRenewDeadline = 15 * time.Second // specifies the amount of time that the current node has to renew the lease before it expires
	DefaultRetryPeriod   = 5 * time.Second  // specifies the amount of time that the current holder of a lease has last updated the lease
)

type Manager struct {
	*rs.Config
	rl.Interface
	Namespace string
	LockName  string
	Identity  string
}

func NewLeaderElectionManager() (*Manager, error) {
	m := &Manager{}

	m.Namespace = getNamespace()
	m.LockName = DefaultLockName
	m.Identity = getLeaderId()
	m.Config = ctrl.GetConfigOrDie()
	// explicitly set the resource lock to leases to avoid the deprecated endpoints resource lock
	lock, _ := rl.NewFromKubeconfig(
		rl.LeasesResourceLock,
		m.Namespace,
		m.LockName,
		rl.ResourceLockConfig{
			Identity: m.Identity,
		},
		m.Config,
		DefaultRenewDeadline,
	)
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
				// this is where you would put your code to cleanup, release resources and stop being the leader
				log.Info("leader election lost")
			},
			OnNewLeader: func(identity string) {
				if identity == m.Identity {
					log.Infof("the leader identity is '%s'\n", identity)
				}
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return el, nil
}
