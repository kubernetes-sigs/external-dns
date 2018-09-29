package egoscale

import (
	"testing"
)

func TestListServiceOfferings(t *testing.T) {
	req := &ListServiceOfferings{}
	_ = req.response().(*ListServiceOfferingsResponse)
}
