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

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testToken          = "test"
	testRecordContent  = "acme"
	testRecordContent2 = "foo"
	txtRecordType      = "TXT"
	testTTL            = 10
)

func setupTest(t *testing.T) (*http.ServeMux, *Client) {
	t.Helper()

	mux := http.NewServeMux()

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	client := NewClient(testToken)
	client.baseURL, _ = url.Parse(server.URL)

	return mux, client
}

func TestClient_Zone(t *testing.T) {
	mux, client := setupTest(t)

	expected := Zone{
		Name: "example.com",
		Records: []ZoneRecord{
			{
				Name:         "test.example.com",
				Type:         txtRecordType,
				TTL:          10,
				ShortAnswers: []string{"test1"},
			},
		},
	}

	mux.Handle("/v2/zones/example.com", validationHandler{
		method: http.MethodGet,
		next:   handleJSONResponse(expected),
	})

	zone, err := client.Zone(context.Background(), "example.com")
	require.NoError(t, err)

	assert.Equal(t, expected, zone)
}

func TestClient_Zones(t *testing.T) {
	mux, client := setupTest(t)

	expected := []Zone{{Name: "example.com"}}

	mux.Handle("/v2/zones", validationHandler{
		method: http.MethodGet,
		next:   handleJSONResponse(ListZones{Zones: expected}),
	})

	zones, err := client.Zones(context.Background())
	require.NoError(t, err)

	assert.Equal(t, expected, zones)
}

func TestClient_ZonesWithRecords(t *testing.T) {
	mux, client := setupTest(t)

	expected := []Zone{
		{
			Name: "example.com",
			Records: []ZoneRecord{
				{
					Name:         "test.example.com",
					Type:         txtRecordType,
					TTL:          10,
					ShortAnswers: []string{"test1"},
				},
			},
		},
	}

	mux.Handle("/v2/zones", validationHandler{
		method: http.MethodGet,
		next:   handleJSONResponse(ListZones{Zones: []Zone{{Name: expected[0].Name}}}),
	})
	mux.Handle("/v2/zones/example.com", validationHandler{
		method: http.MethodGet,
		next:   handleJSONResponse(expected[0]),
	})

	zones, err := client.ZonesWithRecords(context.Background())
	require.NoError(t, err)

	assert.Equal(t, expected, zones)
}

func TestClient_Zone_error(t *testing.T) {
	mux, client := setupTest(t)

	mux.Handle("/v2/zones/example.com", validationHandler{
		method: http.MethodGet,
		next:   handleAPIError(),
	})

	_, err := client.Zone(context.Background(), "example.com")
	require.Error(t, err)
}

func TestClient_RRSet(t *testing.T) {
	mux, client := setupTest(t)

	expected := RRSet{
		TTL: testTTL,
		Records: []ResourceRecords{
			{Content: []string{testRecordContent}},
		},
	}

	mux.Handle("/v2/zones/example.com/foo.example.com/"+txtRecordType, validationHandler{
		method: http.MethodGet,
		next:   handleJSONResponse(expected),
	})

	rrSet, err := client.RRSet(context.Background(), "example.com", "foo.example.com", txtRecordType)
	require.NoError(t, err)

	assert.Equal(t, expected, rrSet)
}

func TestClient_RRSet_error(t *testing.T) {
	mux, client := setupTest(t)

	mux.Handle("/v2/zones/example.com/foo.example.com/"+txtRecordType, validationHandler{
		method: http.MethodGet,
		next:   handleAPIError(),
	})

	_, err := client.RRSet(context.Background(), "example.com", "foo.example.com", txtRecordType)
	require.Error(t, err)
}

func TestClient_DeleteRRSetRecord_Remove(t *testing.T) {
	mux, client := setupTest(t)
	rrSet := RRSet{
		TTL: 10,
		Records: []ResourceRecords{
			{
				Content: []string{"1"},
			},
			{
				Content: []string{"2"},
			},
			{
				Content: []string{"3"},
			},
			{
				Content: []string{"4"},
			},
		},
	}
	mux.HandleFunc("/v2/zones/test.example.com/foo.test.example.com/"+txtRecordType,
		func(writer http.ResponseWriter, request *http.Request) {
			switch request.Method {
			case http.MethodGet:
				handleJSONResponse(rrSet)(writer, request)
			case http.MethodDelete:
			default:
				http.Error(writer, "wrong method", http.StatusNotFound)
			}
		})

	err := client.DeleteRRSetRecord(context.Background(),
		"test.example.com", "foo.test.example.com", txtRecordType, "1", "2", "3", "4")
	require.NoError(t, err)
}

func TestClient_DeleteRRSetRecord_Update(t *testing.T) {
	mux, client := setupTest(t)
	rrSet := RRSet{
		TTL: 10,
		Records: []ResourceRecords{
			{
				Content: []string{"1"},
			},
			{
				Content: []string{"2"},
			},
			{
				Content: []string{"3"},
			},
			{
				Content: []string{"4"},
			},
		},
	}
	mux.HandleFunc("/v2/zones/test.example.com/foo.test.example.com/"+txtRecordType,
		func(writer http.ResponseWriter, request *http.Request) {
			switch request.Method {
			case http.MethodGet:
				handleJSONResponse(rrSet).ServeHTTP(writer, request)
			case http.MethodPut:
				handleRRSet([]ResourceRecords{
					{
						Content: []string{"1"},
					},
					{
						Content: []string{"4"},
					},
				}).ServeHTTP(writer, request)
			default:
				http.Error(writer, "wrong method", http.StatusNotFound)
			}
		})

	err := client.DeleteRRSetRecord(context.Background(),
		"test.example.com", "foo.test.example.com.", txtRecordType, "2", "3")
	require.NoError(t, err)
}

func TestClient_DeleteRRSet(t *testing.T) {
	mux, client := setupTest(t)

	mux.Handle("/v2/zones/test.example.com/my.test.example.com/"+txtRecordType,
		validationHandler{method: http.MethodDelete})

	err := client.DeleteRRSet(context.Background(),
		"test.example.com", "my.test.example.com", txtRecordType)
	require.NoError(t, err)
}

func TestClient_DeleteRRSet_error(t *testing.T) {
	mux, client := setupTest(t)

	mux.Handle("/v2/zones/test.example.com/my.test.example.com/"+txtRecordType, validationHandler{
		method: http.MethodDelete,
		next:   handleAPIError(),
	})

	err := client.DeleteRRSet(context.Background(),
		"test.example.com", "my.test.example.com", txtRecordType)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "delete record request: 500: oops")
}

func TestClient_AddRRSet(t *testing.T) {
	testCases := []struct {
		desc          string
		zone          string
		recordName    string
		value         string
		handledDomain string
		handlers      map[string]http.Handler
		wantErr       bool
	}{
		{
			desc:       "success add",
			zone:       "test.example.com",
			recordName: "my.test.example.com",
			value:      testRecordContent,
			handlers: map[string]http.Handler{
				// createRRSet
				"/v2/zones/test.example.com/my.test.example.com/" + txtRecordType: validationHandler{
					method: http.MethodPost,
					next:   handleRRSet([]ResourceRecords{{Content: []string{testRecordContent}}}),
				},
			},
		},
		{
			desc:       "success update",
			zone:       "test.example.com",
			recordName: "my.test.example.com",
			value:      testRecordContent,
			handlers: map[string]http.Handler{
				"/v2/zones/test.example.com/my.test.example.com/" + txtRecordType: http.HandlerFunc(
					func(rw http.ResponseWriter, req *http.Request) {
						switch req.Method {
						case http.MethodGet: // GetRRSet
							data := RRSet{
								TTL:     testTTL,
								Records: []ResourceRecords{{Content: []string{testRecordContent2}}},
							}
							handleJSONResponse(data).ServeHTTP(rw, req)
						case http.MethodPut: // updateRRSet
							expected := []ResourceRecords{
								{Content: []string{testRecordContent}},
								{Content: []string{testRecordContent2}},
							}
							handleRRSet(expected).ServeHTTP(rw, req)
						default:
							http.Error(rw, "wrong method", http.StatusMethodNotAllowed)
						}
					}),
			},
		},
		{
			desc:       "not in the zone",
			zone:       "test.example.com",
			recordName: "notfound.example.com",
			value:      testRecordContent,
			wantErr:    true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			mux, cl := setupTest(t)

			for pattern, handler := range test.handlers {
				mux.Handle(pattern, handler)
			}

			err := cl.AddZoneRRSet(context.Background(),
				test.zone, test.recordName, txtRecordType, []string{test.value}, testTTL)
			if test.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
		})
	}
}

type validationHandler struct {
	method string
	next   http.Handler
}

func (v validationHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if req.Header.Get("Authorization") != fmt.Sprintf("%s %s", tokenHeader, testToken) {
		rw.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(rw).Encode(APIError{Message: "token up for parsing was not passed through the context"})
		return
	}

	if req.Method != v.method {
		http.Error(rw, "wrong method", http.StatusMethodNotAllowed)
		return
	}

	if v.next != nil {
		v.next.ServeHTTP(rw, req)
	}
}

func handleAPIError() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(rw).Encode(APIError{Message: "oops"})
	}
}

func handleJSONResponse(data interface{}) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		err := json.NewEncoder(rw).Encode(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleRRSet(expected []ResourceRecords) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		body := RRSet{}

		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		if body.TTL != testTTL {
			http.Error(rw, "wrong ttl", http.StatusInternalServerError)
			return
		}
		if !reflect.DeepEqual(body.Records, expected) {
			http.Error(rw, "wrong resource records", http.StatusInternalServerError)
		}
	}
}
