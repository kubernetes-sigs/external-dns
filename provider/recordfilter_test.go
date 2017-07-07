package provider

import "testing"

func TestRecordTypeFilter(t *testing.T) {
	var records = []struct {
		rtype  string
		expect bool
	}{
		{
			"A",
			true,
		},
		{
			"CNAME",
			true,
		},
		{
			"TXT",
			true,
		},
		{
			"MX",
			false,
		},
	}
	for _, r := range records {
		got := recordTypeFilter(r.rtype)
		if r.expect != got {
			t.Errorf("wrong record type %s: expect %v, but got %v", r.rtype, r.expect, got)
		}

	}
}
