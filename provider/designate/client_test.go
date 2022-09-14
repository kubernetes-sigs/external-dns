package designate

import (
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	_ "github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
)

func TestClientZones(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()

	tmpfile, err := setupClientEnv(ts)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()

	client, err := newDesignateClient()
	if err != nil {
		t.Fatal(err)
	}

	expect := map[string]string{
		"a86dba58-0043-4cc6-a1bb-69d5e86f3ca3": "example.org.",
		"34c4561c-9205-4386-9df5-167436f5a222": "foo.example.com.",
	}
	res := make(map[string]string)
	client.ForEachZone(func(zone *zones.Zone) error {
		res[zone.ID] = zone.Name
		return nil
	})
	if diff := cmp.Diff(expect, res); diff != "" {
		t.Fatalf("unexpected zones:\n%s", diff)
	}
}

func TestClientRecordSets(t *testing.T) {
	ts := setupTestServer()
	defer ts.Close()

	tmpfile, err := setupClientEnv(ts)
	if err != nil {
		t.Fatal(err)
	}
	defer tmpfile.Close()

	client, err := newDesignateClient()
	if err != nil {
		t.Fatal(err)
	}

	expect := map[string][]string{
		"example.org./A":     {"10.1.0.2"},
		"foo.example.org./A": {"10.1.0.3", "10.1.0.4"},
	}
	res := make(map[string][]string)
	client.ForEachRecordSet("2150b1bf-dee2-4221-9d85-11f7886fb15f", func(rs *recordsets.RecordSet) error {
		key := fmt.Sprintf("%s/%s", rs.Name, rs.Type)
		res[key] = rs.Records
		return nil
	})

	if diff := cmp.Diff(expect, res); diff != "" {
		t.Fatalf("unexpected recordsets:\n%s", diff)
	}
}

func setupTestServer() *httptest.Server {
	dnsMux := http.NewServeMux()

	dnsMux.Handle("/v2/zones", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, zoneListOutput)
	}))
	dnsMux.Handle("/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, zoneOutput)
	}))

	// dnsTs := httptest.NewTLSServer(nil)
	// dnsTs.Config = &http.Server{Handler: dnsMux}
	// defer dnsTs.Close()

	dnsMux.Handle("/v3/auth/tokens", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(`{
		  "token": {
		    "catalog": [
		      {
		        "id": "9615c2dfac3b4b19935226d4c9d4afce",
		        "name": "designate",
		        "type": "dns",
		        "endpoints": [
		          {
		            "id": "3d3cc3a273b54d0490ac43d6572e4c48",
		            "region": "RegionOne",
		            "region_id": "RegionOne",
		            "interface": "public",
		            "url": "https://` + r.Host + `"
		          }
		        ]
		      }
		    ]
		  }
		}`))
	}))

	return httptest.NewTLSServer(dnsMux)
}

func setupClientEnv(ts *httptest.Server) (*os.File, error) {
	block := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: ts.Certificate().Raw,
	}
	tmpfile, err := ioutil.TempFile("", "os-test.crt")
	if err != nil {
		return nil, err
	}
	if err := pem.Encode(tmpfile, block); err != nil {
		return nil, err
	}
	if err := tmpfile.Close(); err != nil {
		return nil, err
	}

	os.Setenv("OS_AUTH_URL", ts.URL+"/v3")
	os.Setenv("OS_USERNAME", "username")
	os.Setenv("OS_PASSWORD", "password")
	os.Setenv("OS_USER_DOMAIN_NAME", "Default")
	os.Setenv("OPENSTACK_CA_FILE", tmpfile.Name())
	return tmpfile, nil
}

const (
	zoneListOutput = `
{
    "links": {
      "self": "http://example.com:9001/v2/zones"
    },
    "metadata": {
      "total_count": 2
    },
    "zones": [
        {
            "id": "a86dba58-0043-4cc6-a1bb-69d5e86f3ca3",
            "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "name": "example.org.",
            "email": "joe@example.org",
            "ttl": 7200,
            "serial": 1404757531,
            "status": "ACTIVE",
            "action": "CREATE",
            "description": "This is an example zone.",
            "masters": [],
            "type": "PRIMARY",
            "transferred_at": null,
            "version": 1,
            "created_at": "2014-07-07T18:25:31.275934",
            "updated_at": null,
            "links": {
              "self": "https://127.0.0.1:9001/v2/zones/a86dba58-0043-4cc6-a1bb-69d5e86f3ca3"
            }
        },
        {
            "id": "34c4561c-9205-4386-9df5-167436f5a222",
            "pool_id": "572ba08c-d929-4c70-8e42-03824bb24ca2",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "name": "foo.example.com.",
            "email": "joe@foo.example.com",
            "ttl": 7200,
            "serial": 1488053571,
            "status": "ACTIVE",
            "action": "CREATE",
            "description": "This is another example zone.",
            "masters": ["example.com."],
            "type": "PRIMARY",
            "transferred_at": null,
            "version": 1,
            "created_at": "2014-07-07T18:25:31.275934",
            "updated_at": "2015-02-25T20:23:01.234567",
            "links": {
              "self": "https://127.0.0.1:9001/v2/zones/34c4561c-9205-4386-9df5-167436f5a222"
            }
        }
    ]
}
`

	zoneOutput = `
{
    "recordsets": [
        {
            "description": "This is an example record set.",
            "links": {
                "self": "https://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/f7b10e9b-0cae-4a91-b162-562bc6096648"
            },
            "updated_at": null,
            "records": [
                "10.1.0.2"
            ],
            "ttl": 3600,
            "id": "f7b10e9b-0cae-4a91-b162-562bc6096648",
            "name": "example.org.",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "zone_id": "2150b1bf-dee2-4221-9d85-11f7886fb15f",
            "zone_name": "example.com.",
            "created_at": "2014-10-24T19:59:44.000000",
            "version": 1,
            "type": "A",
            "status": "PENDING",
            "action": "CREATE"
        },
        {
            "description": "This is another example record set.",
            "links": {
                "self": "https://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets/7423aeaf-b354-4bd7-8aba-2e831567b478"
            },
            "updated_at": "2017-03-04T14:29:07.000000",
            "records": [
                "10.1.0.3",
                "10.1.0.4"
            ],
            "ttl": 3600,
            "id": "7423aeaf-b354-4bd7-8aba-2e831567b478",
            "name": "foo.example.org.",
            "project_id": "4335d1f0-f793-11e2-b778-0800200c9a66",
            "zone_id": "2150b1bf-dee2-4221-9d85-11f7886fb15f",
            "zone_name": "example.com.",
            "created_at": "2014-10-24T19:59:44.000000",
            "version": 1,
            "type": "A",
            "status": "PENDING",
            "action": "CREATE"
        }
    ],
    "links": {
        "self": "http://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets"
    },
    "metadata": {
        "total_count": 2
    }
}
`
)
