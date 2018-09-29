package egoscale

import (
	"strings"
	"testing"
)

func TestRecordString(t *testing.T) {
	if AAAA != Record(1) {
		t.Error("bad enum value", (int)(AAAA), 1)
	}

	if AAAA.String() != "AAAA" {
		t.Error("mismatch", AAAA, "AAAA")
	}
	s := Record(45)

	if !strings.Contains(s.String(), "45") {
		t.Error("bad record", s.String())
	}
}
