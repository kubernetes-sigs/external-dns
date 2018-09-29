package egoscale

import (
	"context"
	"testing"
	"time"
)

func TestListZonesAPIName(t *testing.T) {
	req := &ListZones{}
	_ = req.response().(*ListZonesResponse)
}

func TestZoneGet(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 1,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"localstorageenabled": true,
			"name": "ch-gva-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "f9a2983b-42e5-3b12-ae74-0b1f54cd6204"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	zone := &Zone{Name: "ch-gva-2"}
	if err := cs.Get(zone); err != nil {
		t.Error(err)
	}

	if zone.Name != "ch-gva-2" {
		t.Errorf("Expected CH-GVA-2, got %q", zone.Name)
	}
}

func TestListZones(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"localstorageenabled": true,
			"name": "ch-gva-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "f9a2983b-42e5-3b12-ae74-0b1f54cd6204"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "381d0a95-ed4a-4ad9-b41c-b97073c1a433",
			"localstorageenabled": true,
			"name": "ch-dk-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "23a24359-121a-38af-a938-e225c97c397b"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "b0fcd72f-47ad-4779-a64f-fe4de007ec72",
			"localstorageenabled": true,
			"name": "at-vie-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "a2a8345d-7daa-3316-8d90-5b8e49706764"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "de88c980-78f6-467c-a431-71bcc88e437f",
			"localstorageenabled": true,
			"name": "de-fra-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "c4bdb9f2-c28d-36a3-bbc5-f91fc69527e6"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	zone := new(Zone)
	zones, err := cs.List(zone)
	if err != nil {
		t.Error(err)
	}

	if len(zones) != 4 {
		t.Errorf("Four zones were expected, got %d", len(zones))
	}

	if zones[2].(*Zone).Name != "at-vie-1" {
		t.Errorf("Expected VIE1 to be third, got %#v", zones[2])
	}
}

func TestListZonesTypeError(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": []}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	_, err := cs.List(&Zone{})
	if err == nil {
		t.Errorf("An error was expected")
	}
}

func TestListZonesPaginateError(t *testing.T) {
	ts := newServer(response{431, jsonContentType, `
{
	"listzonesresponse": {
		"cserrorcode": 9999,
		"errorcode": 431,
		"errortext": "Unable to execute API command listzones due to invalid value. Invalid parameter id value=1747ef5e-5451-41fd-9f1a-58913bae9701 due to incorrect long value format, or entity does not exist or due to incorrect parameter annotation for the field in api cmd class.",
		"uuidList": []
	}
}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	zone := &Zone{
		ID: MustParseUUID("1747ef5e-5451-41fd-9f1a-58913bae9701"),
	}

	req, _ := zone.ListRequest()

	cs.Paginate(req, func(i interface{}, e error) bool {
		if e != nil {
			return false
		}
		t.Errorf("No zones were expected")
		return true
	})
}

func TestListZonesPaginate(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"localstorageenabled": true,
			"name": "ch-gva-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "f9a2983b-42e5-3b12-ae74-0b1f54cd6204"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "381d0a95-ed4a-4ad9-b41c-b97073c1a433",
			"localstorageenabled": true,
			"name": "ch-dk-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "23a24359-121a-38af-a938-e225c97c397b"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "b0fcd72f-47ad-4779-a64f-fe4de007ec72",
			"localstorageenabled": true,
			"name": "at-vie-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "a2a8345d-7daa-3316-8d90-5b8e49706764"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "de88c980-78f6-467c-a431-71bcc88e437f",
			"localstorageenabled": true,
			"name": "de-fra-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "c4bdb9f2-c28d-36a3-bbc5-f91fc69527e6"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	zone := new(Zone)
	req, _ := zone.ListRequest()

	counter := 0
	cs.Paginate(req, func(i interface{}, e error) bool {
		if e != nil {
			t.Error(e)
			return false
		}
		z := i.(*Zone)
		if z.Name == "" {
			t.Errorf("Zone Name not set")
		}
		counter++
		return true
	})

	if counter != 4 {
		t.Errorf("Four zones were expected, got %d", counter)
	}
}

func TestListZonesPaginateBreak(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"localstorageenabled": true,
			"name": "ch-gva-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "f9a2983b-42e5-3b12-ae74-0b1f54cd6204"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "381d0a95-ed4a-4ad9-b41c-b97073c1a433",
			"localstorageenabled": true,
			"name": "ch-dk-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "23a24359-121a-38af-a938-e225c97c397b"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	zone := new(Zone)
	req, _ := zone.ListRequest()

	cs.Paginate(req, func(i interface{}, e error) bool {
		if e != nil {
			t.Error(e)
			return false
		}
		z := i.(*Zone)
		if z.Name == "" {
			t.Errorf("Zone Name not set")
		}
		return false
	})
}

func TestListZonesAsyncError(t *testing.T) {
	ts := newServer(response{431, jsonContentType, `
{
	"listzonesresponse": {
		"cserrorcode": 9999,
		"errorcode": 431,
		"errortext": "Unable to execute API command listzones due to invalid value. Invalid parameter id value=1747ef5e-5451-41fd-9f1a-58913bae9701 due to incorrect long value format, or entity does not exist or due to incorrect parameter annotation for the field in api cmd class.",
		"uuidList": []
	}
}
`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	zone := &Zone{
		ID: MustParseUUID("1747ef5e-5451-41fd-9f1a-58913bae9701"),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	outChan, errChan := cs.AsyncListWithContext(ctx, zone)

	var err error
	for {
		select {
		case _, ok := <-outChan:
			if ok {
				t.Errorf("No zones were expected")
			} else {
				outChan = nil
			}
		case e, ok := <-errChan:
			if ok {
				err = e
			}
			errChan = nil
		}

		if outChan == nil && errChan == nil {
			break
		}
	}

	if err == nil {
		t.Errorf("An error was expected!")
	}
}

func TestListZonesAsync(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"localstorageenabled": true,
			"name": "ch-gva-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "f9a2983b-42e5-3b12-ae74-0b1f54cd6204"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "381d0a95-ed4a-4ad9-b41c-b97073c1a433",
			"localstorageenabled": true,
			"name": "ch-dk-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "23a24359-121a-38af-a938-e225c97c397b"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "b0fcd72f-47ad-4779-a64f-fe4de007ec72",
			"localstorageenabled": true,
			"name": "at-vie-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "a2a8345d-7daa-3316-8d90-5b8e49706764"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "de88c980-78f6-467c-a431-71bcc88e437f",
			"localstorageenabled": true,
			"name": "de-fra-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "c4bdb9f2-c28d-36a3-bbc5-f91fc69527e6"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	zone := new(Zone)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	outChan, errChan := cs.AsyncListWithContext(ctx, zone)

	counter := 0
	for {
		select {
		case z, ok := <-outChan:
			if ok {
				zone := z.(*Zone)
				if zone.Name == "" {
					t.Errorf("Zone Name is empty")
				}
				counter++
			} else {
				outChan = nil
			}
		case e, ok := <-errChan:
			if ok {
				t.Error(e)
			}
			errChan = nil
		}

		if outChan == nil && errChan == nil {
			break
		}
	}

	if counter != 4 {
		t.Errorf("Four zones were expected, got %d", counter)
	}
}

func TestListZonesTwoPages(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"localstorageenabled": true,
			"name": "ch-gva-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "f9a2983b-42e5-3b12-ae74-0b1f54cd6204"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "381d0a95-ed4a-4ad9-b41c-b97073c1a433",
			"localstorageenabled": true,
			"name": "ch-dk-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "23a24359-121a-38af-a938-e225c97c397b"
		}
	]
}}`}, response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "b0fcd72f-47ad-4779-a64f-fe4de007ec72",
			"localstorageenabled": true,
			"name": "at-vie-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "a2a8345d-7daa-3316-8d90-5b8e49706764"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "de88c980-78f6-467c-a431-71bcc88e437f",
			"localstorageenabled": true,
			"name": "de-fra-1",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "c4bdb9f2-c28d-36a3-bbc5-f91fc69527e6"
		}
	]
}}`}, response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zones": null
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	cs.PageSize = 2

	zone := new(Zone)
	zones, err := cs.List(zone)
	if err != nil {
		t.Error(err)
	}

	if len(zones) != 4 {
		t.Errorf("Four zones were expected, got %d", len(zones))
	}
}

func TestListZonesError(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listzonesresponse": {
	"count": 4,
	"zone": [
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"localstorageenabled": true,
			"name": "ch-gva-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "f9a2983b-42e5-3b12-ae74-0b1f54cd6204"
		},
		{
			"allocationstate": "Enabled",
			"dhcpprovider": "VirtualRouter",
			"id": "381d0a95-ed4a-4ad9-b41c-b97073c1a433",
			"localstorageenabled": true,
			"name": "ch-dk-2",
			"networktype": "Basic",
			"securitygroupsenabled": true,
			"tags": [],
			"zonetoken": "23a24359-121a-38af-a938-e225c97c397b"
		}
	]
}}`}, response{400, jsonContentType, `
{"listzonesresponse": {
	"cserrorcode": 9999,
	"errorcode": 431,
	"errortext": "Unable to execute API command listzones",
	"uuidList": []
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	cs.PageSize = 2

	zone := new(Zone)
	_, err := cs.List(zone)
	if err == nil {
		t.Error("An error was expected")
	}
}

func TestListZonesTimeout(t *testing.T) {
	ts := newSleepyServer(time.Second, 200, jsonContentType, `
{"listzonesresponse": {
	"count": 4
}}`)
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	cs.HTTPClient.Timeout = time.Millisecond

	zone := new(Zone)
	_, err := cs.List(zone)
	if err == nil {
		t.Errorf("An error was expected")
	}
}
