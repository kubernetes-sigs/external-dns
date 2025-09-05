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
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"os/signal"
	"reflect"
	"regexp"
	"runtime"
	"syscall"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/provider"
	fakeprovider "sigs.k8s.io/external-dns/provider/fakes"
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
			wantErrMsg: "failed to parse log level",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				// Capture and suppress fatal exit; restore logger after test
				logger := log.StandardLogger()
				prevOut := logger.Out
				prevExit := logger.ExitFunc
				b := new(bytes.Buffer)
				var captureLogFatal bool
				logger.ExitFunc = func(int) { captureLogFatal = true }
				logger.SetOutput(b)
				t.Cleanup(func() {
					logger.SetOutput(prevOut)
					logger.ExitFunc = prevExit
				})

				configureLogger(tt.cfg)

				assert.True(t, captureLogFatal)
				assert.Contains(t, b.String(), tt.wantErrMsg)
			} else {
				// Save and restore logger state to avoid leaking between tests
				logger := log.StandardLogger()
				prevFormatter := logger.Formatter
				prevLevel := log.GetLevel()
				t.Cleanup(func() {
					log.SetLevel(prevLevel)
					logger.SetFormatter(prevFormatter)
				})

				configureLogger(tt.cfg)
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

func TestSelectRegistry(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *externaldns.Config
		provider provider.Provider
		wantErr  bool
		wantType string
	}{
		{
			name: "DynamoDB registry",
			cfg: &externaldns.Config{
				Registry:               "dynamodb",
				AWSDynamoDBRegion:      "us-west-2",
				AWSDynamoDBTable:       "test-table",
				TXTOwnerID:             "owner-id",
				TXTWildcardReplacement: "wildcard",
				ManagedDNSRecordTypes:  []string{"A", "CNAME"},
				ExcludeDNSRecordTypes:  []string{"TXT"},
				TXTCacheInterval:       60,
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "DynamoDBRegistry",
		},
		{
			name: "Noop registry",
			cfg: &externaldns.Config{
				Registry: "noop",
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "NoopRegistry",
		},
		{
			name: "TXT registry",
			cfg: &externaldns.Config{
				Registry:               "txt",
				TXTPrefix:              "prefix",
				TXTOwnerID:             "owner-id",
				TXTCacheInterval:       60,
				TXTWildcardReplacement: "wildcard",
				ManagedDNSRecordTypes:  []string{"A", "CNAME"},
				ExcludeDNSRecordTypes:  []string{"TXT"},
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "TXTRegistry",
		},
		{
			name: "AWS-SD registry",
			cfg: &externaldns.Config{
				Registry:   "aws-sd",
				TXTOwnerID: "owner-id",
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  false,
			wantType: "AWSSDRegistry",
		},
		{
			name: "Unknown registry",
			cfg: &externaldns.Config{
				Registry: "unknown",
			},
			provider: &fakeprovider.MockProvider{},
			wantErr:  true,
			wantType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				// Capture fatal without exiting; avoid brittle output assertions
				logger := log.StandardLogger()
				prevOut := logger.Out
				prevExit := logger.ExitFunc
				var fatalCalled bool
				logger.ExitFunc = func(int) { fatalCalled = true }
				// Capture log output
				b := new(bytes.Buffer)
				logger.SetOutput(b)
				t.Cleanup(func() {
					logger.SetOutput(prevOut)
					logger.ExitFunc = prevExit
				})

				_, err := selectRegistry(tt.cfg, tt.provider)
				assert.NoError(t, err)
				assert.True(t, fatalCalled)
			} else {
				reg, err := selectRegistry(tt.cfg, tt.provider)
				assert.NoError(t, err)
				assert.Contains(t, reflect.TypeOf(reg).String(), tt.wantType)
			}
		})
	}
}

// Provider
func TestBuildProvider(t *testing.T) {
	tests := []struct {
		name          string
		cfg           *externaldns.Config
		expectedType  string
		expectedError string
	}{
		{
			name: "aws provider",
			cfg: &externaldns.Config{
				Provider: "aws",
			},
			expectedType: "*aws.AWSProvider",
		},
		{
			name: "rfc2136 provider",
			cfg: &externaldns.Config{
				Provider:             "rfc2136",
				RFC2136TSIGSecretAlg: "hmac-sha256",
			},
			expectedType: "*rfc2136.rfc2136Provider",
		},
		{
			name: "gandi provider",
			cfg: &externaldns.Config{
				Provider: "gandi",
			},
			expectedError: "no environment variable GANDI_KEY or GANDI_PAT provided",
		},
		{
			name: "inmemory provider",
			cfg: &externaldns.Config{
				Provider: "inmemory",
			},
			expectedType: "*inmemory.InMemoryProvider",
		},
		{
			name: "inmemory cached provider",
			cfg: &externaldns.Config{
				Provider:          "inmemory",
				ProviderCacheTime: 10 * time.Millisecond,
			},
			expectedType: "*provider.CachedProvider",
		},
		{
			name: "coredns provider",
			cfg: &externaldns.Config{
				Provider: "coredns",
			},
			expectedType: "coredns.coreDNSProvider",
		},
		{
			name: "pihole provider",
			cfg: &externaldns.Config{
				Provider:         "pihole",
				PiholeApiVersion: "6",
				PiholeServer:     "http://localhost:8080",
			},
			expectedType: "*pihole.PiholeProvider",
		},
		{
			name: "dnsimple provider",
			cfg: &externaldns.Config{
				Provider: "dnsimple",
			},
			expectedError: "no dnsimple oauth token provided",
		},
		{
			name: "unknown provider",
			cfg: &externaldns.Config{
				Provider: "unknown",
			},
			expectedError: "unknown dns provider: unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domainFilter := endpoint.NewDomainFilter([]string{"example.com"})

			p, err := buildProvider(t.Context(), tt.cfg, domainFilter)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectedError)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, p)
				assert.Contains(t, reflect.TypeOf(p).String(), tt.expectedType)
			}
		})
	}
}

func TestCreateDomainFilter(t *testing.T) {
	tests := []struct {
		name                 string
		cfg                  *externaldns.Config
		expectedDomainFilter *endpoint.DomainFilter
		isConfigured         bool
	}{
		{
			name: "RegexDomainFilter",
			cfg: &externaldns.Config{
				RegexDomainFilter:    regexp.MustCompile(`example\.com`),
				RegexDomainExclusion: regexp.MustCompile(`excluded\.example\.com`),
			},
			expectedDomainFilter: endpoint.NewRegexDomainFilter(regexp.MustCompile(`example\.com`), regexp.MustCompile(`excluded\.example\.com`)),
			isConfigured:         true,
		},
		{
			name: "RegexDomainWithoutExclusionFilter",
			cfg: &externaldns.Config{
				RegexDomainFilter: regexp.MustCompile(`example\.com`),
			},
			expectedDomainFilter: endpoint.NewRegexDomainFilter(regexp.MustCompile(`example\.com`), nil),
			isConfigured:         true,
		},
		{
			name: "DomainFilterWithExclusions",
			cfg: &externaldns.Config{
				DomainFilter:   []string{"example.com"},
				ExcludeDomains: []string{"excluded.example.com"},
			},
			expectedDomainFilter: endpoint.NewDomainFilterWithExclusions([]string{"example.com"}, []string{"excluded.example.com"}),
			isConfigured:         true,
		},
		{
			name: "DomainFilterWithExclusionsOnly",
			cfg: &externaldns.Config{
				ExcludeDomains: []string{"excluded.example.com"},
			},
			expectedDomainFilter: endpoint.NewDomainFilterWithExclusions([]string{}, []string{"excluded.example.com"}),
			isConfigured:         true,
		},
		{
			name: "EmptyDomainFilter",
			cfg: &externaldns.Config{
				DomainFilter:   []string{},
				ExcludeDomains: []string{},
			},
			expectedDomainFilter: endpoint.NewDomainFilterWithExclusions([]string{}, []string{}),
			isConfigured:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := createDomainFilter(tt.cfg)
			assert.Equal(t, tt.isConfigured, filter.IsConfigured())
			assert.Equal(t, tt.expectedDomainFilter, filter)
		})
	}
}

func getRandomPort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func TestServeMetrics(t *testing.T) {
	// Use a fresh DefaultServeMux for this test (do not restore to avoid data race with server goroutine)
	http.DefaultServeMux = http.NewServeMux()

	port, err := getRandomPort()
	require.NoError(t, err)
	address := fmt.Sprintf("localhost:%d", port)

	go serveMetrics(fmt.Sprintf(":%d", port))

	// Wait for the TCP socket to be ready
	require.Eventually(t, func() bool {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			return false
		}
		_ = conn.Close()
		return true
	}, 2*time.Second, 10*time.Millisecond, "server not ready with port open in time")

	resp, err := http.Get(fmt.Sprintf("http://%s/healthz", address))
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	_ = resp.Body.Close()

	resp, err = http.Get(fmt.Sprintf("http://%s/metrics", address))
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	_ = resp.Body.Close()
}

func TestHandleSigterm(t *testing.T) {
	cancelCalled := make(chan bool, 1)
	cancel := func() { cancelCalled <- true }

	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer log.SetOutput(os.Stderr)

	go handleSigterm(cancel)

	// Simulate sending a SIGTERM signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	defer signal.Stop(sigChan)
	err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	assert.NoError(t, err)

	// Wait for cancel to be called
	select {
	case <-cancelCalled:
		assert.Contains(t, logOutput.String(), "Received SIGTERM. Terminating...")
	case sig := <-sigChan:
		assert.Equal(t, syscall.SIGTERM, sig)
	case <-time.After(1 * time.Second):
		t.Fatal("cancel function was not called")
	}
}

func TestBuildSource(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}))
	defer svr.Close()

	tests := []struct {
		name          string
		cfg           *externaldns.Config
		expectedError bool
	}{
		{
			name: "Valid configuration with sources",
			cfg: &externaldns.Config{
				APIServerURL:   svr.URL,
				Sources:        []string{"fake"},
				RequestTimeout: 6 * time.Millisecond,
			},
			expectedError: false,
		},
		{
			name: "Empty sources configuration",
			cfg: &externaldns.Config{
				APIServerURL:   svr.URL,
				Sources:        []string{},
				RequestTimeout: 6 * time.Millisecond,
			},
			expectedError: false,
		},
		{
			name: "Update events enabled",
			cfg: &externaldns.Config{
				KubeConfig:   "path-to-kubeconfig-not-exists",
				APIServerURL: svr.URL,
				Sources:      []string{"ingress"},
				UpdateEvents: true,
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, err := buildSource(t.Context(), tt.cfg)

			if tt.expectedError {
				assert.Error(t, err)
				assert.Nil(t, src)
			} else {
				require.NoError(t, err)
				assert.NotNil(t, src)
			}
		})
	}
}

func TestBuildSourceWithWrappers(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}))
	defer svr.Close()

	tests := []struct {
		name    string
		cfg     *externaldns.Config
		asserts func(*testing.T, *externaldns.Config)
	}{
		{
			name: "configuration with target filter wrapper",
			cfg: &externaldns.Config{
				APIServerURL:    svr.URL,
				Sources:         []string{"fake"},
				TargetNetFilter: []string{"10.0.0.0/8"},
			},
			asserts: func(t *testing.T, cfg *externaldns.Config) {
				assert.True(t, cfg.IsSourceWrapperInstrumented("target-filter"))
			},
		},
		{
			name: "configuration without target filter wrapper",
			cfg: &externaldns.Config{
				APIServerURL: svr.URL,
				Sources:      []string{"fake"},
			},
			asserts: func(t *testing.T, cfg *externaldns.Config) {
				assert.True(t, cfg.IsSourceWrapperInstrumented("dedup"))
				assert.True(t, cfg.IsSourceWrapperInstrumented("nat64"))
				assert.False(t, cfg.IsSourceWrapperInstrumented("target-filter"))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := buildSource(t.Context(), tt.cfg)
			require.NoError(t, err)
			tt.asserts(t, tt.cfg)
		})
	}
}

// Helper used by runExecuteSubprocess.
func TestHelperProcess(t *testing.T) {
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

// runExecuteSubprocess runs Execute in a separate process and returns exit code and output.
func runExecuteSubprocess(t *testing.T, args []string) (int, string, error) {
	t.Helper()
	cmdArgs := append([]string{"-test.run=TestHelperProcess", "--"}, args...)
	cmd := exec.Command(os.Args[0], cmdArgs...)
	cmd.Env = append(os.Environ(), "GO_WANT_HELPER_PROCESS=1")
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf
	err := cmd.Run()
	output := buf.String()
	if err == nil {
		return 0, output, nil
	}
	ee := &exec.ExitError{}
	if errors.As(err, &ee) {
		return ee.ExitCode(), output, nil
	}
	return -1, output, err
}

func TestExecuteOnceDryRunExitsZero(t *testing.T) {
	// Use :0 for an ephemeral metrics port.
	code, _, err := runExecuteSubprocess(t, []string{
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
	code, _, err := runExecuteSubprocess(t, []string{
		"--source", "fake",
		"--provider", "unknown",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

func TestExecuteValidationErrorNoSources(t *testing.T) {
	code, _, err := runExecuteSubprocess(t, []string{
		"--provider", "inmemory",
		"--metrics-address", ":0",
	})
	require.NoError(t, err)
	assert.NotEqual(t, 0, code)
}

func TestExecuteFlagParsingErrorInvalidLogFormat(t *testing.T) {
	code, _, err := runExecuteSubprocess(t, []string{
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
	code, _, err := runExecuteSubprocess(t, []string{
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
	code, _, err := runExecuteSubprocess(t, []string{
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
	code, _, err := runExecuteSubprocess(t, []string{
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
	code, _, err := runExecuteSubprocess(t, []string{
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
	code, _, err := runExecuteSubprocess(t, []string{
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

// ValidateConfig triggers log.Fatalf (in-process).
func TestExecuteConfigValidationFatalInProcess(t *testing.T) {
	// Prepare args to trigger validation error before any goroutines start
	prevArgs := os.Args
	os.Args = []string{
		"external-dns",
		"--source", "fake",
		"--provider", "inmemory",
		"--ignore-hostname-annotation", // triggers validation: FQDN template required when ignoring annotations
		"--metrics-address", ":0",
	}
	t.Cleanup(func() { os.Args = prevArgs })

	// Capture logs and replace Fatalf with Goexit to stop only the Execute goroutine
	logger := log.StandardLogger()
	prevExit := logger.ExitFunc
	prevOut := logger.Out
	buf := new(bytes.Buffer)
	logger.SetOutput(buf)
	logger.ExitFunc = func(int) { runtime.Goexit() }
	t.Cleanup(func() { logger.ExitFunc = prevExit; logger.SetOutput(prevOut) })

	done := make(chan struct{})
	go func() {
		defer close(done)
		Execute()
	}()

	select {
	case <-done:
		// ok
	case <-time.After(2 * time.Second):
		t.Fatal("Execute did not exit after validation fatal")
	}

	// Do not assert on logger text to avoid flakiness with global logger
}

// Run path with --events; shut down via SIGTERM.
func TestExecuteDefaultRunWithEventsStopsOnSigterm(t *testing.T) {
	// Use a fresh DefaultServeMux for this test (do not restore to avoid data race with server goroutine)
	http.DefaultServeMux = http.NewServeMux()

	// Prepare args to run Execute without --once and with --events
	prevArgs := os.Args
	os.Args = []string{
		"external-dns",
		"--source", "fake",
		"--provider", "inmemory",
		"--events",
		"--dry-run",
		"--metrics-address", ":0",
	}
	t.Cleanup(func() { os.Args = prevArgs })

	// Prevent log.Fatal from terminating the test process
	logger := log.StandardLogger()
	prevExit := logger.ExitFunc
	logger.ExitFunc = func(int) { runtime.Goexit() }
	t.Cleanup(func() { logger.ExitFunc = prevExit })

	done := make(chan struct{})
	go func() {
		defer close(done)
		Execute()
	}()

	// Give goroutines time to start
	time.Sleep(50 * time.Millisecond)

	// Send SIGTERM to trigger handleSigterm(cancel)
	_ = syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	select {
	case <-done:
		// ok
	case <-time.After(2 * time.Second):
		t.Fatal("Execute did not stop after SIGTERM")
	}
}

// Webhook server path; pre-bind 127.0.0.1:8888 to force a bind failure.
func TestExecuteWebhookServerFailsPortInUseInProcess(t *testing.T) {
	// Use a fresh DefaultServeMux for this test (do not restore to avoid data race with server goroutine)
	http.DefaultServeMux = http.NewServeMux()

	// Pre-bind the webhook server port so it is unavailable
	l, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		// If we cannot bind, assume something else is bound already, which is fine for this test
	} else {
		t.Cleanup(func() { _ = l.Close() })
	}

	prevArgs := os.Args
	os.Args = []string{
		"external-dns",
		"--source", "fake",
		"--provider", "inmemory",
		"--webhook-server",
		"--metrics-address", ":0",
	}
	t.Cleanup(func() { os.Args = prevArgs })

	logger := log.StandardLogger()
	prevExit := logger.ExitFunc
	logger.ExitFunc = func(int) { runtime.Goexit() }
	t.Cleanup(func() { logger.ExitFunc = prevExit })

	done := make(chan struct{})
	go func() {
		defer close(done)
		Execute()
	}()

	select {
	case <-done:
		// ok
	case <-time.After(2 * time.Second):
		t.Fatal("Execute did not exit after webhook server fatal")
	}
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	src, err := buildSource(ctx, cfg)
	require.NoError(t, err)
	domainFilter := createDomainFilter(cfg)
	p, err := buildProvider(ctx, cfg, domainFilter)
	require.NoError(t, err)
	ctrl, err := buildController(ctx, cfg, src, p, domainFilter)
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
