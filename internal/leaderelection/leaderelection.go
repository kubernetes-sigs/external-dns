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
	DefaultLockName                    = "external-dns-leader-election"
	DefaultNamespace                   = "default"
	DefaultLeaderElectionLeaseDuration = 20 * time.Second
	DefaultLeaderElectionRenewDeadline = 15 * time.Second
	DefaultLeaderElectionRetryPeriod   = 5 * time.Second
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
		DefaultLeaderElectionRenewDeadline,
	)
	if err != nil {
		return nil, err
	}
	m.Interface = lock
	return m, nil
}

func (m *Manager) ConfigureElection(run func(ctx context.Context)) (*le.LeaderElector, error) {
	el, err := le.NewLeaderElector(le.LeaderElectionConfig{
		Lock:            m,
		LeaseDuration:   DefaultLeaderElectionLeaseDuration,
		RenewDeadline:   DefaultLeaderElectionRenewDeadline,
		RetryPeriod:     DefaultLeaderElectionRetryPeriod,
		ReleaseOnCancel: true, // release the lock when the context is cancelled
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

func getLeaderId() string {
	// Leader id, needs to be unique
	id, err := os.Hostname()
	if err != nil {
		return fmt.Sprintf("external-dns_%s", uuid.NewUUID())
	}
	id = fmt.Sprintf("%s_%s", id, uuid.NewUUID())
	return id
}
