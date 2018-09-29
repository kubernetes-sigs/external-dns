package egoscale

import (
	"encoding/json"
	"testing"
)

func TestUUIDMustParse(t *testing.T) {
	defer func() {
		recover()
	}()
	MustParseUUID("foo")
	t.Error("invalid uuid should pazone")
}

func TestUUIDMarshalJSON(t *testing.T) {
	zone := &Zone{
		ID: MustParseUUID("5361a11b-615c-42bf-9bdb-e2c3790ada14"),
	}
	j, err := json.Marshal(zone)
	if err != nil {
		t.Fatal(err)
	}
	s := string(j)
	expected := `{"id":"5361a11b-615c-42bf-9bdb-e2c3790ada14"}`
	if expected != s {
		t.Errorf("bad json serialization, got %q, expected %s", s, expected)
	}
}

func TestUUIDUnmarshalJSONFailure(t *testing.T) {
	ss := []string{
		`{"id": 1}`,
		`{"id": "1"}`,
		`{"id": "5361a11b-615c-42bf-9bdb-e2c3790ada1"}`,
	}
	zone := &Zone{}
	for _, s := range ss {
		if err := json.Unmarshal([]byte(s), zone); err == nil {
			t.Errorf("an error was expected, %#v", zone)
		}
	}
}
