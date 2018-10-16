package dynect

import (
	"log"
	"os"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

var (
	DynCustomerName string
	DynUsername     string
	DynPassword     string
	testZone        string
)

func getenv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

func init() {
	DynCustomerName = getenv("DYNECT_CUSTOMER_NAME", "go-dynect")
	DynUsername = getenv("DYNECT_USER_NAME", "dynect-user")
	DynPassword = getenv("DYNECT_PASSWORD", "p@55w0rd")
	testZone = getenv("DYNECT_TEST_ZONE", "go-dynect.test")
}

// test helper to begin recording or playback of vcr cassette
func withCassette(cassetteName string, f func(*recorder.Recorder)) {
	r, err := recorder.New(cassetteName)
	if err != nil {
		log.Fatal(err)
	}

	defer r.Stop()

	f(r)
}

// test helper to setup client with vcr cassette
func withClient(cassetteName string, f func(*Client)) {
	withCassette(cassetteName, func(r *recorder.Recorder) {
		c := NewClient(DynCustomerName)
		c.SetTransport(r)
		c.Verbose(true)

		f(c)
	})
}

// test helper to setup authenticated client with vcr cassette
func testWithClientSession(cassetteName string, t *testing.T, f func(*Client)) {
	withClient(cassetteName, func(c *Client) {
		if err := c.Login(DynUsername, DynPassword); err != nil {
			t.Fatal(err)
		}

		defer func() {
			if err := c.Logout(); err != nil {
				t.Error(err)
			}
		}()

		f(c)
	})
}

func TestLoginLogout(t *testing.T) {
	withClient("fixtures/login_logout", func(c *Client) {
		if err := c.Login(DynUsername, DynPassword); err != nil {
			t.Error(err)
		}

		if err := c.Logout(); err != nil {
			t.Error(err)
		}
	})
}

func TestZonesRequest(t *testing.T) {
	testWithClientSession("fixtures/zones_request", t, func(c *Client) {
		var resp ZonesResponse

		if err := c.Do("GET", "Zone", nil, &resp); err != nil {
			t.Fatal(err)
		}

		nresults := len(resp.Data)
		for i, zone := range resp.Data {
			parts := strings.Split(zone, "/")
			t.Logf("(%d/%d) %q", i+1, nresults, parts[len(parts)-2])
		}
	})
}

func TestFetchingAllZoneRecords(t *testing.T) {
	testWithClientSession("fixtures/fetching_all_zone_records", t, func(c *Client) {
		var resp AllRecordsResponse

		if err := c.Do("GET", "AllRecord/"+testZone, nil, &resp); err != nil {
			t.Error(err)
		}

		for _, zr := range resp.Data {
			parts := strings.Split(zr, "/")
			uri := strings.Join(parts[2:], "/")
			t.Log(uri)

			var record RecordResponse

			if err := c.Do("GET", uri, nil, &record); err != nil {
				t.Fatal(err)
			}

			t.Log("OK")
		}
	})
}
