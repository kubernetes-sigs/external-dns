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

package controller

import (
	"bytes"
	"context"
	"errors"
	"os"
	"os/exec"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	provider "sigs.k8s.io/external-dns/provider/factory"
	"sigs.k8s.io/external-dns/source"
	"sigs.k8s.io/external-dns/source/wrappers"
)

// Logger
func TestConfigureLogger(t *testing.T) {
	tests := []struct {
		name       string
		cfg        *externaldns.Config
		wantLevel  log.Level
		wantJSON   bool
		wantErr    bool
		wantErrMsg string
	}{
		{
			name: "Default log format and level",
			cfg: &externaldns.Config{
				LogLevel:  "info",
				LogFormat: "text",
			},
			wantLevel: log.InfoLevel,
		},
		{
			name: "JSON log format",
			cfg: &externaldns.Config{
				LogLevel:  "debug",
				LogFormat: "json",
			},
			wantLevel: log.DebugLevel,
			wantJSON:  true,
		},
		{
			name: "Invalid log level",
			cfg: &externaldns.Config{
				LogLevel:  "invalid",
				LogFormat: "text",
			},
			wantLevel:  log.InfoLevel,
			wantErr:    true,
			wantErrMsg: "not a valid logrus Level: \"invalid\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				err := configureLogger(tt.cfg)
				require.Error(t, err)
			} else {
				// Save and restore logger state to avoid leaking between tests
				logger := log.StandardLogger()
				prevFormatter := logger.Formatter
				prevLevel := log.GetLevel()
				t.Cleanup(func() {
					log.SetLevel(prevLevel)
					logger.SetFormatter(prevFormatter)
				})

				err := configureLogger(tt.cfg)
				require.NoError(t, err)

				assert.Equal(t, tt.wantLevel, log.GetLevel())

				if tt.wantJSON {
					assert.IsType(t, &log.JSONFormatter{}, log.StandardLogger().Formatter)
				} else {
					assert.IsType(t, &log.TextFormatter{}, log.StandardLogger().Formatter)
				}
			}
		})
	}
}

// Helper used by runExecuteSubprocess.
func TestHelperProcess(_ *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	// Parse args after the "--" sentinel.
	idx := -1
	for i, a := range os.Args {
		if a == "--" {
			idx = i
			break
		}
	}
	var args []string
	if idx >= 0 {
		args = os.Args[idx+1:]
	}
	os.Args = append([]string{"external-dns"}, args...)
	Execute()
}

// runExecuteSubprocess runs Execute in a separate process and returns exit code.
func runExecuteSubprocess(t *testing.T, args []string) (int, error) {
	t.Helper()
	// make sure the subprocess does not run forever
	ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
	defer cancel()

	// TODO: investigate why -test.run=TestHelperProcess
	cmdArgs := append([]string{"-test.run=TestHelperProcess", "--"}, args...)
	cmd := exec.CommandContext(ctx, os.Args[0], cmdArgs...)
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return -1, ctx.Err()
	}
	if err == nil {
		return 0, nil
	}
	ee := &exec.ExitError{}
	if errors.As(err, &ee) {
		return ee.ExitCode(), nil
	}
	return -1, err
}

func TestExecuteOnceDryRunExitsZero(t *testing.T) {
	// Use :0 for an ephemeral metrics port.
	code, err := runExecuteSubprocess(t, []string{
		"--source", "fake",
		"--provider", "inmemory",
		"--once",
		"--dry-run",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.Equal(t, 0, code)
}

func TestExecuteUnknownProviderExitsNonZero(t *testing.T) {
	code, err := runExecuteSubprocess(t, []string{
		"--source", "fake",
		"--provider", "unknown",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

func TestExecuteValidationErrorNoSources(t *testing.T) {
	code, err := runExecuteSubprocess(t, []string{
		"--provider", "inmemory",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

func TestExecuteFlagParsingErrorInvalidLogFormat(t *testing.T) {
	code, err := runExecuteSubprocess(t, []string{
		"--log-format", "invalid",
		// Provide minimal required flags to keep errors focused on parsing
		"--source", "fake",
		"--provider", "inmemory",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

// Config validation failure triggers log.Fatalf.
func TestExecuteConfigValidationErrorExitsNonZero(t *testing.T) {
	code, err := runExecuteSubprocess(t, []string{
		"--source", "fake",
		// Choose a provider with validation that fails without required flags
		"--provider", "azure",
		// No --azure-config-file provided
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

// buildSource failure triggers log.Fatal.
func TestExecuteBuildSourceErrorExitsNonZero(t *testing.T) {
	// Use a valid source name (ingress) and an invalid kubeconfig path to
	// force client creation failure inside buildSource.
	code, err := runExecuteSubprocess(t, []string{
		"--source", "ingress",
		"--kubeconfig", "this/path/does/not/exist",
		"--provider", "inmemory",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

// RunOnce error exits non-zero.
func TestExecuteRunOnceErrorExitsNonZero(t *testing.T) {
	// Connector source dials a TCP server; use a closed port to fail.
	code, err := runExecuteSubprocess(t, []string{
		"--source", "connector",
		"--connector-source-server", "127.0.0.1:1",
		"--provider", "inmemory",
		"--once",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

// Run loop error exits non-zero.
func TestExecuteRunLoopErrorExitsNonZero(t *testing.T) {
	code, err := runExecuteSubprocess(t, []string{
		"--source", "connector",
		"--connector-source-server", "127.0.0.1:1",
		"--provider", "inmemory",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

// buildController registry-creation failure triggers log.Fatal.
func TestExecuteBuildControllerErrorExitsNonZero(t *testing.T) {
	code, err := runExecuteSubprocess(t, []string{
		"--source", "fake",
		"--provider", "inmemory",
		"--registry", "dynamodb",
		// Force NewDynamoDBRegistry to fail validation by using empty owner id
		"--txt-owner-id", "",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

// Controller run loop stops on context cancel.
func TestControllerRunCancelContextStopsLoop(t *testing.T) {
	// Minimal controller using fake source and inmemory provider.
	cfg := &externaldns.Config{
		Sources:    []string{"fake"},
		Provider:   "inmemory",
		LogLevel:   "error",
		LogFormat:  "text",
		Policy:     "sync",
		Registry:   "txt",
		TXTOwnerID: "test-owner",
	}
	sCfg, err := source.NewSourceConfig(cfg)
	require.NoError(t, err)
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()
	src, err := wrappers.Build(ctx, sCfg)
	require.NoError(t, err)
	domainFilter := endpoint.NewDomainFilterWithOptions(
		endpoint.WithDomainFilter(cfg.DomainFilter),
		endpoint.WithDomainExclude(cfg.DomainExclude),
		endpoint.WithRegexDomainFilter(cfg.RegexDomainFilter),
		endpoint.WithRegexDomainExclude(cfg.RegexDomainExclude),
	)
	p, err := provider.Select(ctx, cfg, domainFilter)
	require.NoError(t, err)
	ctrl, err := buildController(ctx, cfg, sCfg, src, p, domainFilter)
	require.NoError(t, err)

	done := make(chan struct{})
	go func() {
		ctrl.Run(ctx)
		close(done)
	}()
	cancel()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("controller did not stop after context cancellation")
	}
}
