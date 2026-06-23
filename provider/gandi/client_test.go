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

package gandi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/domain"
	"github.com/go-gandi/go-gandi/livedns"
	"github.com/go-gandi/go-gandi/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDomainClient(t *testing.T) {
	client := NewDomainClient(&domain.Domain{})
	assert.NotNil(t, client)
}

func TestNewLiveDNSClient(t *testing.T) {
	client := NewLiveDNSClient(&livedns.LiveDNS{})
	assert.NotNil(t, client)
}

func TestDomainClientListDomains(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v5/domain/domains", r.URL.Path)
		_ = json.NewEncoder(w).Encode([]domain.ListResponse{{FQDN: "example.com"}})
	}))
	t.Cleanup(server.Close)

	gandiClient := domain.New(config.Config{
		APIURL:              server.URL,
		PersonalAccessToken: "test-token",
	})

	adapter := NewDomainClient(gandiClient)
	domains, err := adapter.ListDomains()
	require.NoError(t, err)
	require.Len(t, domains, 1)
	assert.Equal(t, "example.com", domains[0].FQDN)
}

func TestLiveDNSClientGetDomainRecords(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/v5/livedns/domains/example.com/records", r.URL.Path)
		_ = json.NewEncoder(w).Encode([]livedns.DomainRecord{{
			RrsetName:   "www",
			RrsetType:   "A",
			RrsetTTL:    300,
			RrsetValues: []string{"203.0.113.1"},
		}})
	}))
	t.Cleanup(server.Close)

	liveClient := livedns.New(config.Config{
		APIURL:              server.URL,
		PersonalAccessToken: "test-token",
	})
	adapter := NewLiveDNSClient(liveClient)
	records, err := adapter.GetDomainRecords("example.com")
	require.NoError(t, err)
	require.Len(t, records, 1)
	assert.Equal(t, "www", records[0].RrsetName)
}

func TestLiveDNSClientCreateDomainRecordSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		_ = json.NewEncoder(w).Encode(types.StandardResponse{
			Code:    201,
			Message: "created",
			UUID:    "uuid-1",
			Status:  "success",
			Errors: []types.StandardError{{
				Location:    "name",
				Name:        "name",
				Description: "ok",
			}},
		})
	}))
	t.Cleanup(server.Close)

	liveClient := livedns.New(config.Config{
		APIURL:              server.URL,
		PersonalAccessToken: "test-token",
	})
	adapter := NewLiveDNSClient(liveClient)
	resp, err := adapter.CreateDomainRecord("example.com", "www", "A", 300, []string{"203.0.113.1"})
	require.NoError(t, err)
	assert.Equal(t, 201, resp.Code)
	assert.Equal(t, "created", resp.Message)
	require.Len(t, resp.Errors, 1)
	assert.Equal(t, "name", resp.Errors[0].Name)
}

func TestLiveDNSClientCreateDomainRecordError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(w, `{"message":"bad request"}`)
	}))
	t.Cleanup(server.Close)

	liveClient := livedns.New(config.Config{
		APIURL:              server.URL,
		PersonalAccessToken: "test-token",
	})
	adapter := NewLiveDNSClient(liveClient)
	_, err := adapter.CreateDomainRecord("example.com", "www", "A", 300, []string{"203.0.113.1"})
	require.Error(t, err)
}

func TestLiveDNSClientUpdateDomainRecordByNameAndType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPut, r.Method)
		assert.Equal(t, "/v5/livedns/domains/example.com/records/www/A", r.URL.Path)
		_ = json.NewEncoder(w).Encode(types.StandardResponse{Code: 200, Message: "updated"})
	}))
	t.Cleanup(server.Close)

	liveClient := livedns.New(config.Config{
		APIURL:              server.URL,
		PersonalAccessToken: "test-token",
	})
	adapter := NewLiveDNSClient(liveClient)
	resp, err := adapter.UpdateDomainRecordByNameAndType("example.com", "www", "A", 300, []string{"203.0.113.2"})
	require.NoError(t, err)
	assert.Equal(t, 200, resp.Code)
}

func TestLiveDNSClientDeleteDomainRecord(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodDelete, r.Method)
		assert.Equal(t, "/v5/livedns/domains/example.com/records/www/A", r.URL.Path)
		w.WriteHeader(http.StatusNoContent)
	}))
	t.Cleanup(server.Close)

	liveClient := livedns.New(config.Config{
		APIURL:              server.URL,
		PersonalAccessToken: "test-token",
	})
	adapter := NewLiveDNSClient(liveClient)
	err := adapter.DeleteDomainRecord("example.com", "www", "A")
	require.NoError(t, err)
}
