package egoscale

import (
	"testing"
)

func TestResourceMetadata(t *testing.T) {
	var _ Command = (*ListResourceDetails)(nil)
}

func TestListResourceDetailss(t *testing.T) {
	req := &ListResourceDetails{}
	if req.name() != "listResourceDetails" {
		t.Errorf("API call doesn't match")
	}
	if req.description() != "List resource detail(s)" {
		t.Errorf("API call doesn't match")
	}
	_ = req.response().(*ListResourceDetailsResponse)
}
