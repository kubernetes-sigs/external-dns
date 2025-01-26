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
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes/fake"
	rs "k8s.io/client-go/rest"
	le "k8s.io/client-go/tools/leaderelection"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/external-dns/internal/testutils"
)

func TestNewLeaderElectionManager_Success(t *testing.T) {
	manager, err := NewLeaderElectionManager()
	require.NoError(t, err)
	require.NotNil(t, manager)
	require.Equal(t, DefaultNamespace, manager.Namespace)
	require.Equal(t, DefaultLockName, manager.LockName)
	require.NotEmpty(t, manager.Identity)
	require.NotNil(t, manager.Config)
	require.NotNil(t, manager.Interface)
}

func TestConfigureElection_Success(t *testing.T) {
	manager, err := NewLeaderElectionManager()
	require.NoError(t, err)
	require.NotNil(t, manager)

	elector, err := manager.ConfigureElection(mockRun)
	require.NoError(t, err)
	require.NotNil(t, elector)
	require.False(t, elector.IsLeader())
}

func TestNewLeaderElectionManager_ConfigError(t *testing.T) {
	// Temporarily replace the GetConfig function to simulate an error
	originalGetConfig := ctrl.GetConfig
	ctrl.GetConfig = func() (*rs.Config, error) {
		return nil, fmt.Errorf("config error")
	}
	defer func() { ctrl.GetConfig = originalGetConfig }()

	manager, err := NewLeaderElectionManager()
	require.Error(t, err)
	require.Nil(t, manager)
}

func NewK8sConfigForTesting() (*fake.Clientset, rs.Config) {
	mockClient := fake.NewSimpleClientset()
	return mockClient, rs.Config{
		Host: "http://localhost",
	}
}

func TestConfigureElection_Run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	buf := testutils.LogsToBuffer(t)
	_, cfg := NewK8sConfigForTesting()

	manager := &Manager{
		Config:    &cfg,
		Interface: &fakeLock{identity: "test-identity"},
		Namespace: DefaultNamespace,
		LockName:  DefaultLockName,
		Identity:  "test-identity",
	}

	var elector *le.LeaderElector

	elector, err := manager.ConfigureElection(func(ctx context.Context) {
		time.Sleep(1 * time.Microsecond)
		require.True(t, elector.IsLeader())
		cancel()
	})
	require.NoError(t, err)
	elector.Run(ctx)
	require.False(t, elector.IsLeader())

	require.Contains(t, buf.String(), "attempting to acquire leader lease")
	require.Contains(t, buf.String(), "update leader election record:test-identity")
	require.Contains(t, buf.String(), "leader election lost")
}
