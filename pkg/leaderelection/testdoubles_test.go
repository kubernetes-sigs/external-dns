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

	log "github.com/sirupsen/logrus"
	rl "k8s.io/client-go/tools/leaderelection/resourcelock"
)

// Mock function to simulate the run function in ConfigureElection
func mockRun(ctx context.Context) {}

type fakeLock struct {
	identity string
}

// Get is a dummy to allow us to have a fakeLock for testing.
func (fl *fakeLock) Get(ctx context.Context) (ler *rl.LeaderElectionRecord, rawRecord []byte, err error) {
	log.Info("get leader election record:", fl.identity)
	return &rl.LeaderElectionRecord{
		HolderIdentity: fl.identity,
	}, nil, nil
}

// Create is a dummy to allow us to have a fakeLock for testing.
func (fl *fakeLock) Create(ctx context.Context, ler rl.LeaderElectionRecord) error {
	log.Info("create leader election record:", fl.identity)
	return nil
}

// Update is a dummy to allow us to have a fakeLock for testing.
func (fl *fakeLock) Update(ctx context.Context, ler rl.LeaderElectionRecord) error {
	log.Info("update leader election record:", fl.identity)
	return nil
}

// RecordEvent is a dummy to allow us to have a fakeLock for testing.
func (fl *fakeLock) RecordEvent(string) {}

// Identity is a dummy to allow us to have a fakeLock for testing.
func (fl *fakeLock) Identity() string {
	return fl.identity
}

// Describe is a dummy to allow us to have a fakeLock for testing.
func (fl *fakeLock) Describe() string {
	return fl.identity
}

// logsToBuffer redirects log output to a buffer for testing purposes
// func logsToBuffer(t *testing.T) *bytes.Buffer {
// 	t.Helper()
// 	buf := new(bytes.Buffer)
// 	log.SetOutput(buf)
// 	klog.SetOutput(buf)
// 	flags := &flag.FlagSet{}
// 	klog.InitFlags(flags)
// 	_ = flags.Set("logtostderr", "false")
// 	_ = flags.Set("alsologtostderr", "false")
// 	_ = flags.Set("stderrthreshold", "4")
// 	return buf
// }
