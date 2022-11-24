/*
Copyright 2017 The Kubernetes Authors.

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

package hetzner

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"strings"
	"testing"

	hdns "github.com/jobstoit/hetzner-dns-go/dns"
	"github.com/jobstoit/hetzner-dns-go/dns/schema"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

type testEnv struct {
	Server   *httptest.Server
	Mux      *http.ServeMux
	Context  context.Context
	Provider *HetznerProvider
	State    *testState
}

type testState struct {
	Zones   []schema.Zone
	Records []schema.Record
}

func (env *testEnv) Teardown() {
	env.Server.Close()
	env.Server = nil
	env.Mux = nil
	env.Provider = nil
	env.Context.Done()
	env.Context = nil
}

func findRecord(records []schema.Record, test func(schema.Record) bool) (index int, ret schema.Record) {
	index = -1
	for i, r := range records {
		if test(r) {
			ret = r
			index = i
			break
		}
	}
	return
}

func filterRecords(records []schema.Record, test func(schema.Record) bool) (ret []schema.Record) {
	for _, r := range records {
		if test(r) {
			ret = append(ret, r)
		}
	}
	return
}

func filterZones(zones []schema.Zone, test func(schema.Zone) bool) (ret []schema.Zone) {
	for _, z := range zones {
		if test(z) {
			ret = append(ret, z)
		}
	}
	return
}

func orderEndpoints(eps []*endpoint.Endpoint) []*endpoint.Endpoint {
	sort.Slice(eps, func(i, j int) bool {
		a := eps[i]
		b := eps[j]

		dnsComp := strings.Compare(a.DNSName, b.DNSName)

		if dnsComp == -1 {
			return true
		} else if dnsComp == 1 {
			return false
		}

		typeComp := strings.Compare(a.RecordType, b.RecordType)

		if typeComp == -1 {
			return true
		} else if typeComp == 1 {
			return false
		}

		return false
	})
	return eps
}

func newTestEnv(t *testing.T) (testEnv, error) {
	log.SetLevel(log.TraceLevel)
	ctx := context.Background()
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	client := hdns.NewClient(
		hdns.WithEndpoint(server.URL),
		hdns.WithToken("32CharactersTokenxxxxxxxXxxxxxxx"),
		hdns.WithApplication("testing", hdns.Version),
	)
	provider, err := NewHetznerProvider(ctx, endpoint.NewDomainFilter([]string{}), false, 50)
	if err != nil {
		return testEnv{}, fmt.Errorf("could not create testEnv: %s", err)
	}
	provider.Client = *client

	state := testState{
		Zones: []schema.Zone{
			{
				ID:   "1",
				Name: "hetzner.com",
			},
			{
				ID:   "2",
				Name: "hetzner.cloud",
			}},
		Records: []schema.Record{
			{
				ID:     "1",
				Type:   "A",
				Name:   "@",
				ZoneID: "1",
				Value:  "6.7.8.9",
			},
			{
				ID:     "2",
				Type:   "A",
				Name:   "upd",
				ZoneID: "2",
				Value:  "1.2.3.4",
			},
			{
				ID:     "3",
				Type:   "TXT",
				Name:   "del",
				ZoneID: "2",
				Value:  "\"heritage=external-dns,external-dns/owner=default\"",
			},
		},
	}

	mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		res := schema.ZoneListResponse{}
		res.Zones = state.Zones

		zoneNameFilter := r.URL.Query().Get("name")
		if zoneNameFilter != "" {
			res.Zones = filterZones(state.Zones, func(z schema.Zone) bool { return z.Name == zoneNameFilter })
		}

		json.NewEncoder(w).Encode(res) // nolint: errcheck
	})

	mux.HandleFunc("/records", func(w http.ResponseWriter, r *http.Request) {
		defaultResp := schema.RecordListResponse{
			Records: state.Records,
		}
		switch r.Method {
		case "GET":
			resp := schema.RecordListResponse{
				Records: state.Records,
			}
			zoneIDFilter := r.URL.Query().Get("zone_id")
			if zoneIDFilter != "" {
				resp.Records = filterRecords(state.Records, func(r schema.Record) bool { return r.ZoneID == zoneIDFilter })
			}
			json.NewEncoder(w).Encode(resp) // nolint: errcheck
			return
		case "POST":
			resp := schema.RecordResponse{}
			req := schema.RecordCreateRequest{}
			json.NewDecoder(r.Body).Decode(&req)
			newRecord := schema.Record{
				Name: req.Name,
				Ttl: func() int {
					if req.Ttl != nil {
						return *req.Ttl
					}
					return 0
				}(),
				Type:   req.Type,
				Value:  req.Value,
				ZoneID: req.ZoneID,
			}
			state.Records = append(state.Records, newRecord)
			resp.Record = newRecord
			json.NewEncoder(w).Encode(resp) // nolint: errcheck
			return
		}
		json.NewEncoder(w).Encode(defaultResp) // nolint: errcheck
		return
	})

	mux.HandleFunc("/records/", func(w http.ResponseWriter, r *http.Request) {
		defaultResp := schema.RecordResponse{
			Record: schema.Record{},
		}
		recordID := strings.TrimPrefix(r.URL.Path, "/records/")

		resp := schema.RecordResponse{
			Record: schema.Record{},
		}

		recordIndex, targetedRecord := findRecord(state.Records, func(r schema.Record) bool { return r.ID == recordID })
		if recordIndex == -1 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		switch r.Method {
		case "PUT":
			req := schema.RecordCreateRequest{}
			json.NewDecoder(r.Body).Decode(&req)

			adaptedRecord := schema.Record{
				Name: req.Name,
				Ttl: func() int {
					if req.Ttl != nil {
						return *req.Ttl
					}
					return 0
				}(),
				Type:   req.Type,
				Value:  req.Value,
				ZoneID: req.ZoneID,
			}
			state.Records[recordIndex] = adaptedRecord
			resp.Record = adaptedRecord

			json.NewEncoder(w).Encode(resp) // nolint: errcheck
			return
		case "DELETE":
			if len(state.Records) > recordIndex+1 {
				state.Records = append(state.Records[:recordIndex], state.Records[recordIndex+1:]...)
			} else if len(state.Records) > recordIndex {
				state.Records = state.Records[:recordIndex]
			} else {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			resp.Record = targetedRecord

			json.NewEncoder(w).Encode(resp) // nolint: errcheck
			return
		}
		json.NewEncoder(w).Encode(defaultResp) // nolint: errcheck
		return
	})

	return testEnv{
		Server:   server,
		Mux:      mux,
		Context:  ctx,
		Provider: provider,
		State:    &state,
	}, nil
}

type cstm_assert struct {
	t *testing.T
}

func newAssert(t *testing.T) *cstm_assert {
	return &cstm_assert{
		t: t,
	}
}

func (a cstm_assert) NotNil(x any) bool {
	if x == nil {
		a.t.Error("expected value but got nil")
	}

	return x != nil
}

func (a cstm_assert) NoError(err error) bool {
	if err != nil {
		a.t.Errorf("unexpected error: %v", err)
	}

	return err == nil
}

func (a cstm_assert) Error(err error) bool {
	if err == nil {
		a.t.Error("missing expected error")
	}

	return err != nil
}

func (a cstm_assert) EqStr(expected, actual string) bool {
	if expected != actual {
		a.t.Errorf("expected '%s' but got '%s'", expected, actual)
	}

	return expected == actual
}

func (a cstm_assert) EqInt(expected, actual int) bool {
	if expected != actual {
		a.t.Errorf("expected %d but got %d", expected, actual)
	}

	return expected == actual
}

func TestClientNoToken(t *testing.T) {
	os.Unsetenv("HDNS_TOKEN")

	env, err := newTestEnv(t)

	if nil == err {
		defer env.Teardown()
		t.Error("Failed to trigger expected error")
	} else if err.Error() != "could not create testEnv: no token found" {
		t.Fatalf("Missing authorization token triggered unexpected error message: %s", err)
	}
}

func TestZoneList(t *testing.T) {
	os.Setenv("HDNS_TOKEN", "32CharactersTokenxxxxxxxXxxxxxxx")
	env, err := newTestEnv(t)
	if err != nil {
		t.Error(err)
	}
	defer env.Teardown()
	as := newAssert(t)

	zones, err := env.Provider.Zones(env.Context)
	as.NoError(err)
	as.EqInt(2, len(zones))

	env.Provider.domainFilter = endpoint.NewDomainFilter([]string{"hetzner.cloud"})

	zones, err = env.Provider.Zones(env.Context)
	as.NoError(err)
	as.EqInt(1, len(zones))
}

func TestRecordList(t *testing.T) {
	os.Setenv("HDNS_TOKEN", "32CharactersTokenxxxxxxxXxxxxxxx")
	env, err := newTestEnv(t)
	if err != nil {
		t.Error(err)
	}
	defer env.Teardown()

	as := newAssert(t)

	eps, err := env.Provider.Records(env.Context)
	as.NoError(err)
	as.EqInt(3, len(eps))

	env.Provider.domainFilter = endpoint.NewDomainFilter([]string{"hetzner.com"})

	eps, err = env.Provider.Records(env.Context)
	as.NoError(err)
	as.EqInt(1, len(eps))

	env.Provider.domainFilter = endpoint.NewDomainFilter([]string{"example.com"})

	eps, err = env.Provider.Records(env.Context)
	as.NoError(err)
	as.EqInt(0, len(eps))
}

func TestApplyChanges(t *testing.T) {
	os.Setenv("HDNS_TOKEN", "32CharactersTokenxxxxxxxXxxxxxxx")
	env, err := newTestEnv(t)
	if err != nil {
		t.Error(err)
	}
	defer env.Teardown()

	as := newAssert(t)

	changes := plan.Changes{
		Create: []*endpoint.Endpoint{
			{
				DNSName:    "new.hetzner.com",
				RecordType: "A",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("4.3.2.1"),
				Labels:     endpoint.Labels{},
			},
			{
				DNSName:    "newnottl.hetzner.com",
				RecordType: "A",
				Targets:    endpoint.NewTargets("4.3.2.1"),
				Labels:     endpoint.Labels{},
			},
		},
		UpdateNew: []*endpoint.Endpoint{
			{
				DNSName:    "upd.hetzner.cloud",
				RecordType: "A",
				RecordTTL:  500,
				Targets:    endpoint.NewTargets("1.2.3.4", "5.6.7.8"),
				Labels:     endpoint.Labels{},
			},
		},
		Delete: []*endpoint.Endpoint{
			{
				DNSName:    "del.hetzner.cloud",
				RecordType: "TXT",
				RecordTTL:  300,
				Targets:    endpoint.NewTargets("\"heritage=external-dns,external-dns/owner=default\""),
				Labels:     endpoint.Labels{},
			},
		},
	}

	exceptedOutcome := []*endpoint.Endpoint{}
	exceptedOutcome = append(exceptedOutcome,
		&endpoint.Endpoint{
			Targets:    endpoint.NewTargets("6.7.8.9"),
			RecordType: "A",
			DNSName:    "hetzner.com",
			Labels:     endpoint.Labels{},
			RecordTTL:  endpoint.TTL(0),
		},
	)
	exceptedOutcome = append(exceptedOutcome, changes.Create...)
	exceptedOutcome = append(exceptedOutcome, changes.UpdateNew...)

	err = env.Provider.ApplyChanges(env.Context, &changes)
	as.NoError(err)

	endpoints, err := env.Provider.Records(env.Context)

	exceptedOutcome = orderEndpoints(exceptedOutcome)
	endpoints = orderEndpoints(endpoints)

	as.NoError(err)
	assert.Equal(t, exceptedOutcome, endpoints)
}
