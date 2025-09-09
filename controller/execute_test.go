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
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"regexp"
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
