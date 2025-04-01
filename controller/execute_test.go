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
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"regexp"
	"syscall"
	"testing"
	"time"

	"context"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

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
			provider: &MockProvider{},
			wantErr:  false,
			wantType: "DynamoDBRegistry",
		},
		{
			name: "Noop registry",
			cfg: &externaldns.Config{
				Registry: "noop",
			},
			provider: &MockProvider{},
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
				TXTNewFormatOnly:       true,
			},
			provider: &MockProvider{},
			wantErr:  false,
			wantType: "TXTRegistry",
		},
		{
			name: "AWS-SD registry",
			cfg: &externaldns.Config{
				Registry:   "aws-sd",
				TXTOwnerID: "owner-id",
			},
			provider: &MockProvider{},
			wantErr:  false,
			wantType: "AWSSDRegistry",
		},
		{
			name: "Unknown registry",
			cfg: &externaldns.Config{
				Registry: "unknown",
			},
			provider: &MockProvider{},
			wantErr:  true,
			wantType: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				defer func() { log.StandardLogger().ExitFunc = nil }()
				b := new(bytes.Buffer)
				log.StandardLogger().ExitFunc = func(int) {}
				log.StandardLogger().SetOutput(b)

				_, err := selectRegistry(tt.cfg, tt.provider)
				assert.NoError(t, err)
				assert.Contains(t, b.String(), "unknown registry: unknown")
			} else {
				reg, err := selectRegistry(tt.cfg, tt.provider)
				assert.NoError(t, err)
				assert.Contains(t, reflect.TypeOf(reg).String(), tt.wantType)
			}
		})
	}
}

func TestCreateDomainFilter(t *testing.T) {
	tests := []struct {
		name                 string
		cfg                  *externaldns.Config
		expectedDomainFilter endpoint.DomainFilter
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

func TestHandleSigterm(t *testing.T) {
	cancelCalled := make(chan bool, 1)
	cancel := func() {
		cancelCalled <- true
	}

	var logOutput bytes.Buffer
	log.SetOutput(&logOutput)
	defer log.SetOutput(os.Stderr)

	go handleSigterm(cancel)

	// Simulate sending a SIGTERM signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)
	err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	assert.NoError(t, err)

	// Wait for the cancel function to be called
	select {
	case <-cancelCalled:
		assert.Contains(t, logOutput.String(), "Received SIGTERM. Terminating...")
	case sig := <-sigChan:
		assert.Equal(t, syscall.SIGTERM, sig)
	case <-time.After(1 * time.Second):
		t.Fatal("cancel function was not called")
	}
}

func TestServeMetrics(t *testing.T) {
	l, _ := net.Listen("tcp", ":0")
	_ = l.Close()
	_, port, _ := net.SplitHostPort(l.Addr().String())

	go serveMetrics(fmt.Sprintf(":%s", port))
	resp, err := http.Get(fmt.Sprintf("http://localhost:%s", port) + "/healthz")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = http.Get(fmt.Sprintf("http://localhost:%s", port) + "/metrics")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

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
				defer func() { log.StandardLogger().ExitFunc = nil }()

				b := new(bytes.Buffer)
				var captureLogFatal bool
				log.StandardLogger().ExitFunc = func(int) { captureLogFatal = true }
				log.StandardLogger().SetOutput(b)

				configureLogger(tt.cfg)

				assert.True(t, captureLogFatal)
				assert.Contains(t, b.String(), tt.wantErrMsg)
			} else {
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

// mocks
type MockProvider struct{}

func (m *MockProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (p *MockProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {

	return nil
}

func (m *MockProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return nil, nil
}

func (m *MockProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	return nil
}
