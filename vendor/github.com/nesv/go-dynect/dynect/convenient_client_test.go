package dynect

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

// test helper to setup convenient client with vcr cassette
func withConvenientClient(cassetteName string, f func(*ConvenientClient)) {
	withCassette(cassetteName, func(r *recorder.Recorder) {
		c := NewConvenientClient(DynCustomerName)
		c.SetTransport(r)
		c.Verbose(true)

		f(c)
	})
}

// test helper to setup authenticated convenient client with vcr cassette
func testWithConvenientClientSession(cassetteName string, t *testing.T, f func(*ConvenientClient)) {
	withConvenientClient(cassetteName, func(c *ConvenientClient) {
		if err := c.Login(DynUsername, DynPassword); err != nil {
			t.Error(err)
		}

		defer func() {
			if err := c.Logout(); err != nil {
				t.Error(err)
			}
		}()

		f(c)
	})
}

func TestConvenientLoginLogout(t *testing.T) {
	withConvenientClient("fixtures/login_logout", func(c *ConvenientClient) {
		if err := c.Login(DynUsername, DynPassword); err != nil {
			t.Error(err)
		}

		if err := c.Logout(); err != nil {
			t.Error(err)
		}
	})
}

func TestConvenientGetA(t *testing.T) {
	testWithConvenientClientSession("fixtures/convenient_get_a", t, func(c *ConvenientClient) {
		actual := Record{
			Zone: testZone,
			Type: "A",
			FQDN: "foobar." + testZone,
		}

		if err := c.GetRecordID(&actual); err != nil {
			t.Fatal(err)
		}

		if err := c.GetRecord(&actual); err != nil {
			t.Fatal(err)
		}

		if actual.Value != "10.9.8.7" {
			t.Fatalf("Incorrect value %q for %q (expected %q)", actual.Value, actual.FQDN, "foobar.go-dynect.test.")
		}

		t.Log("OK")
	})
}

func TestConvenientGetANotFound(t *testing.T) {
	testWithConvenientClientSession("fixtures/convenient_get_a_not_found", t, func(c *ConvenientClient) {
		actual := Record{
			Zone: testZone,
			Type: "A",
			FQDN: "unknown." + testZone,
		}

		if err := c.GetRecordID(&actual); err == nil {
			t.Fatalf("Expected error getting %q", actual.FQDN)
		} else if !strings.HasPrefix(err.Error(), "Failed to find Dyn record id:") {
			t.Fatalf("Expected error %q for %q (actual error %q)", "Failed to find Dyn record id:", actual.FQDN, err.Error())
		}

		t.Log("OK")
	})
}

func TestConvenientGetCNAME(t *testing.T) {
	testWithConvenientClientSession("fixtures/convenient_get_cname", t, func(c *ConvenientClient) {
		actual := Record{
			Zone: testZone,
			Type: "CNAME",
			FQDN: "foo." + testZone,
		}

		if err := c.GetRecordID(&actual); err != nil {
			t.Fatal(err)
		}

		if err := c.GetRecord(&actual); err != nil {
			t.Fatal(err)
		}

		if actual.Value != "foobar.go-dynect.test." {
			t.Fatalf("Incorrect value %q (expected %q)", actual.Value, "foobar.go-dynect.test.")
		}

		t.Log("OK")
	})
}

func TestConvenientCreateMX(t *testing.T) {
	testWithConvenientClientSession("fixtures/convenient_create_mx", t, func(c *ConvenientClient) {
		record := Record{
			Zone:  testZone,
			Type:  "MX",
			Value: "123 mx.example.com.",
			TTL:   "12345",
		}

		if err := c.CreateRecord(&record); err != nil {
			t.Fatal(err)
		}

		if err := c.PublishZone(testZone); err != nil {
			t.Fatal(err)
		}

		if err := c.GetRecordID(&record); err != nil {
			t.Fatal(err)
		}

		if err := c.GetRecord(&record); err != nil {
			t.Fatal(err)
		}

		if record.FQDN != testZone {
			t.Fatalf("Expected FQDN %q (actual %q)", testZone, record.FQDN)
		}

		id, err := strconv.Atoi(record.ID)
		if err != nil || id <= 0 {
			t.Fatalf("Expected ID to be positive integer (actual %q)", record.ID)
		}

		ttl, err := strconv.Atoi(record.TTL)
		if err != nil || ttl != 12345 {
			t.Fatalf("Expected ID to be 12345 (actual %q)", record.TTL)
		}

		t.Log("OK")
	})
}

func TestConvenientCreateZone(t *testing.T) {
	testWithConvenientClientSession("fixtures/convenient_create_zone", t, func(c *ConvenientClient) {
		subZone := fmt.Sprintf("subzone.%s", testZone)

		if err := c.CreateZone(subZone, "admin@example.com", "day", "1800"); err != nil {
			t.Fatal(err)
		}

		if err := c.PublishZone(subZone); err != nil {
			t.Fatal(err)
		}

		z := &Zone{Zone: subZone}

		if err := c.GetZone(z); err != nil {
			t.Fatal(err)
		}

		if z.Zone != subZone {
			t.Fatalf("Expected Zone of %q (actual %q)", subZone, z.Zone)
		}

		if z.Type != "Primary" {
			t.Fatalf("Expected Zone Type of %q (actual %q)", "Primary", z.Type)
		}

		if z.SerialStyle != "day" {
			t.Fatalf("Expected SerialStyle of %q (actual %q)", "day", z.SerialStyle)
		}

		if z.Serial == "" {
			t.Fatalf("Expected non-empty Serial (actual %q)", z.Serial)
		}

		t.Log("OK")
	})
}

func TestConvenientDeleteZone(t *testing.T) {
	testWithConvenientClientSession("fixtures/convenient_delete_zone", t, func(c *ConvenientClient) {
		subZone := fmt.Sprintf("zone-%s", testZone)

		if err := c.DeleteZone(subZone); err != nil {
			t.Fatal(err)
		}

		z := &Zone{Zone: subZone}

		if err := c.GetZone(z); err == nil {
			t.Fatalf("Zone %q not deleted", subZone)
		}

		t.Log("OK")
	})
}

func TestConvenientDeleteSubZone(t *testing.T) {
	testWithConvenientClientSession("fixtures/convenient_delete_sub_zone", t, func(c *ConvenientClient) {
		subZone := fmt.Sprintf("subzone.%s", testZone)

		if err := c.DeleteZone(subZone); err != nil {
			t.Fatal(err)
		}

		if err := c.DeleteZoneNode(subZone); err != nil {
			t.Fatal(err)
		}

		if err := c.PublishZone(testZone); err != nil {
			t.Fatal(err)
		}

		z := &Zone{Zone: subZone}

		if err := c.GetZone(z); err == nil {
			t.Fatalf("Zone %q not deleted", subZone)
		}

		t.Log("OK")
	})
}
