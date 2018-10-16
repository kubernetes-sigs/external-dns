package egoscale

import (
	"encoding/json"
	"testing"
)

func TestCIDRMustParse(t *testing.T) {
	defer func() {
		recover()
	}()
	MustParseCIDR("foo")
	t.Error("invalid cidr should panic")
}

func TestCIDRMarshalJSON(t *testing.T) {
	nic := &Nic{
		IP6CIDR: MustParseCIDR("::/0"),
	}
	j, err := json.Marshal(nic)
	if err != nil {
		t.Fatal(err)
	}
	s := string(j)
	expected := `{"ip6cidr":"::/0"}`
	if expected != s {
		t.Errorf("bad json serialization, got %q, expected %s", s, expected)
	}
}

func TestCIDRUnmarshalJSONFailure(t *testing.T) {
	ss := []string{
		`{"ip6cidr": 123}`,
		`{"ip6cidr": "123"}`,
		`{"ip6cidr": "192.168.0.1/33"}`,
	}
	nic := &Nic{}
	for _, s := range ss {
		if err := json.Unmarshal([]byte(s), nic); err == nil {
			t.Errorf("an error was expected, %#v", nic)
		}
	}
}
