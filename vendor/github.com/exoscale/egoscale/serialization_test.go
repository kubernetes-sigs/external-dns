package egoscale

import (
	"net"
	"net/url"
	"strings"
	"testing"
)

func TestPrepareValues(t *testing.T) {
	type tag struct {
		Name      string `json:"name"`
		IsVisible bool   `json:"isvisible,omitempty"`
	}

	tr := true
	f := false

	profile := struct {
		IgnoreMe    string
		Zone        string            `json:"myzone,omitempty"`
		Name        string            `json:"name"`
		NoName      string            `json:"omitempty"`
		ID          int               `json:"id"`
		UserID      uint              `json:"user_id"`
		IsGreat     bool              `json:"is_great"`
		IsActive    *bool             `json:"is_active"`
		IsAlive     *bool             `json:"is_alive"`
		Num         float64           `json:"num"`
		Bytes       []byte            `json:"bytes"`
		IDs         []string          `json:"ids,omitempty"`
		TagPointers []*tag            `json:"tagpointers,omitempty"`
		Tags        []tag             `json:"tags,omitempty"`
		Map         map[string]string `json:"map"`
		IP          net.IP            `json:"ip,omitempty"`
		UUID        *UUID             `json:"uuid,omitempty"`
		UUIDs       []UUID            `json:"uuids,omitempty"`
		CIDR        *CIDR             `json:"cidr,omitempty"`
		CIDRList    []CIDR            `json:"cidrlist,omitempty"`
		MAC         MACAddress        `json:"mac,omitempty"`
	}{
		IgnoreMe: "bar",
		Name:     "world",
		NoName:   "foo",
		ID:       1,
		UserID:   uint(2),
		IsActive: &tr,
		IsAlive:  &f,
		Num:      3.14,
		Bytes:    []byte("exo"),
		IDs:      []string{"1", "2", "three"},
		TagPointers: []*tag{
			{Name: "foo"},
			{Name: "bar", IsVisible: false},
		},
		Tags: []tag{
			{Name: "foo"},
			{Name: "bar", IsVisible: false},
		},
		Map: map[string]string{
			"foo": "bar",
		},
		IP:   net.IPv4(192, 168, 0, 11),
		UUID: MustParseUUID("5361a11b-615c-42bf-9bdb-e2c3790ada14"),
		UUIDs: []UUID{
			*MustParseUUID("5361a11b-615c-42bf-9bdb-e2c3790ada14"),
			*MustParseUUID("5361a11b-615c-42bf-9bdb-e2c3790ada11"),
		},
		CIDR: MustParseCIDR("192.168.0.0/32"),
		CIDRList: []CIDR{
			*MustParseCIDR("192.168.0.0/32"),
			*MustParseCIDR("::/0"),
		},
		MAC: MAC48(0x01, 0x23, 0x45, 0x67, 0x89, 0xab),
	}

	params := url.Values{}
	err := prepareValues("", params, profile)
	if err != nil {
		t.Fatal(err)
	}

	if _, ok := params["myzone"]; ok {
		t.Errorf("myzone params shouldn't be set, got %v", params.Get("myzone"))
	}

	if params.Get("is_active") != "true" {
		t.Errorf("IsActive params wasn't properly set, got %v", params.Get("IsActive"))
	}

	if params.Get("is_alive") != "false" {
		t.Errorf("IsAlive params wasn't properly set, got %v", params.Get("IsAlive"))
	}

	if params.Get("NoName") != "foo" {
		t.Errorf("NoName params wasn't properly set, got %v", params.Get("NoName"))
	}

	if params.Get("name") != "world" {
		t.Errorf("name params wasn't properly set, got %v", params.Get("name"))
	}

	if params.Get("bytes") != "ZXhv" {
		t.Errorf("bytes params wasn't properly encoded in base 64, got %v", params.Get("bytes"))
	}

	if params.Get("ids") != "1,2,three" {
		t.Errorf("array of strings, wasn't property encoded, got %v", params.Get("ids"))
	}

	if _, ok := params["ignoreme"]; ok {
		t.Errorf("IgnoreMe key was set")
	}

	v := params.Get("tags[0].name")
	if v != "foo" {
		t.Errorf("expected tags to be serialized as foo, got %#v", v)
	}

	v = params.Get("tagpointers[0].name")
	if v != "foo" {
		t.Errorf("expected tag pointers to be serialized as foo, got %#v", v)
	}

	v = params.Get("map[0].foo")
	if v != "bar" {
		t.Errorf("expected map to be serialized as .foo => \"bar\", got .foo => %#v", v)
	}

	v = params.Get("is_great")
	if v != "false" {
		t.Errorf("expected bool to be serialized as \"false\", got %#v", v)
	}

	v = params.Get("ip")
	if v != "192.168.0.11" {
		t.Errorf(`expected ip to be serialized as "192.168.0.11", got %q`, v)
	}

	v = params.Get("uuid")
	if v != "5361a11b-615c-42bf-9bdb-e2c3790ada14" {
		t.Errorf(`expected uuid to be serialized as "5361a11b-615c-42bf-9bdb-e2c3790ada14", got %q`, v)
	}

	v = params.Get("uuids")
	if !strings.Contains(v, "5361a11b-615c-42bf-9bdb-e2c3790ada14,") {
		t.Errorf(`expected uuids to contains "5361a11b-615c-42bf-9bdb-e2c3790ada14", got %q`, v)
	}

	v = params.Get("cidr")
	if v != "192.168.0.0/32" {
		t.Errorf(`expected cidr to be serialized as "192.168.0.0/32", got %q`, v)
	}

	v = params.Get("cidrlist")
	if v != "192.168.0.0/32,::/0" {
		t.Errorf(`expected cidrlist to be serialized as "192.168.0.0/32,::/0", got %q`, v)
	}

	v = params.Get("mac")
	if v != "01:23:45:67:89:ab" {
		t.Errorf(`expected mac to be serialized as "01:23:45:67:89:ab", got %q`, v)
	}
}

func TestPrepareValuesStringRequired(t *testing.T) {
	profile := struct {
		RequiredField string `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesBoolRequired(t *testing.T) {
	profile := struct {
		RequiredField bool `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err != nil {
		t.Fatal(nil)
	}
	if params.Get("requiredfield") != "false" {
		t.Errorf("bool params wasn't set to false (default value)")
	}
}

func TestPrepareValuesBoolPtrRequired(t *testing.T) {
	profile := struct {
		RequiredField *bool `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesIntRequired(t *testing.T) {
	profile := struct {
		RequiredField int64 `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesUintRequired(t *testing.T) {
	profile := struct {
		RequiredField uint64 `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesBytesRequired(t *testing.T) {
	profile := struct {
		RequiredField []byte `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesSliceString(t *testing.T) {
	profile := struct {
		RequiredField []string `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesIP(t *testing.T) {
	profile := struct {
		RequiredField net.IP `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesIPZero(t *testing.T) {
	profile := struct {
		RequiredField net.IP `json:"requiredfield"`
	}{
		RequiredField: net.IPv4zero,
	}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesMap(t *testing.T) {
	profile := struct {
		RequiredField map[string]string `json:"requiredfield"`
	}{}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err == nil {
		t.Errorf("It should have failed")
	}
}

func TestPrepareValuesBoolPtr(t *testing.T) {
	tru := new(bool)
	f := new(bool)
	*tru = true
	*f = false

	profile := struct {
		IsOne   bool  `json:"is_one,omitempty"`
		IsTwo   bool  `json:"is_two,omitempty"`
		IsThree *bool `json:"is_three,omitempty"`
		IsFour  *bool `json:"is_four,omitempty"`
		IsFive  *bool `json:"is_five,omitempty"`
	}{
		IsOne:   true,
		IsTwo:   false,
		IsThree: tru,
		IsFour:  f,
	}

	params := url.Values{}
	err := prepareValues("", params, &profile)
	if err != nil {
		t.Error(err)
	}
	if params["is_one"][0] != "true" {
		t.Errorf("Expected is_one to be true")
	}
	if isTwo, ok := params["is_two"]; ok {
		t.Errorf("Expected is_two to be missing, got %v", isTwo)
	}
	if params["is_three"][0] != "true" {
		t.Errorf("Expected is_three to be true")
	}
	if params["is_four"][0] != "false" {
		t.Errorf("Expected is_four to be false")
	}
	if isFive, ok := params["is_five"]; ok {
		t.Errorf("Expected is_five to be missing, got %v", isFive)
	}
}
