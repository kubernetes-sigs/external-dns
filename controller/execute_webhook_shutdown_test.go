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

package controller

import (
	"context"
	"os"
	"os/exec"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestExecuteWebhookShutdownHelper(t *testing.T) {
	if os.Getenv("GO_WANT_WEBHOOK_SHUTDOWN_HELPER") != "1" {
		t.SkipNow()
	}
	log.SetOutput(os.Stdout)
	os.Args = []string{"external-dns", "--source=fake", "--provider=inmemory", "--policy=sync", "--webhook-server", "--metrics-address=127.0.0.1:0"}
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	execute(ctx)
}

func TestExecuteWebhookCanceledContextReturns(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, os.Args[0], "-test.run=^TestExecuteWebhookShutdownHelper$")
	cmd.Env = append(os.Environ(), "GO_WANT_WEBHOOK_SHUTDOWN_HELPER=1")
	output, err := cmd.CombinedOutput()
	require.NoError(t, ctx.Err(), "webhook mode ignored context cancellation")
	require.NoError(t, err, "%s", output)
}
