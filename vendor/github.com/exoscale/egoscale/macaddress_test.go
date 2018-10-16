package egoscale

import (
	"encoding/json"
	"testing"
)

func TestMACAddressMustParse(t *testing.T) {
	defer func() {
		recover()
	}()
	MustParseMAC("foo")
	t.Error("invalid mac should panic")
}

func TestMACAddressMarshalJSON(t *testing.T) {
	nic := &Nic{
		MACAddress: MAC48(0x01, 0x23, 0x45, 0x67, 0x89, 0xab),
	}
	j, err := json.Marshal(nic)
	if err != nil {
		t.Fatal(err)
	}
	s := string(j)
	expected := `{"macaddress":"01:23:45:67:89:ab"}`
	if expected != s {
		t.Errorf("bad json serialization, got %q, expected %s", s, expected)
	}
}

func TestMACAddressUnmarshalJSON(t *testing.T) {
	s := `{"macaddress": "01:23:45:67:89:ab"}`
	nic := &Nic{}
	if err := json.Unmarshal([]byte(s), nic); err != nil {
		t.Errorf("no errors were expected, %s", err)
	} else {
		if nic.MACAddress == nil {
			t.Errorf("a macaddress was expected, got %+v", nic)
		}
	}
}

func TestMACAddressUnmarshalJSONFailure(t *testing.T) {
	ss := []string{
		`{"macaddress": 123}`,
		`{"macaddress": "123"}`,
		`{"macaddress": "01:23:45:67:89:a"}`,
		`{"macaddress": "01:23:45:67:89:ab\""}`,
	}
	nic := &Nic{}
	for _, s := range ss {
		if err := json.Unmarshal([]byte(s), nic); err == nil {
			t.Errorf("an error was expected, %#v", nic)
		}
	}
}
