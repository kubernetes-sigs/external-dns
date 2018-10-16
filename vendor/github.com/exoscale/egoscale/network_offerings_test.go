package egoscale

import (
	"testing"
)

func TestListNetworkOfferings(t *testing.T) {
	req := &ListNetworkOfferings{}
	_ = req.response().(*ListNetworkOfferingsResponse)
}

func TestUpdateNetworkOffering(t *testing.T) {
	req := &UpdateNetworkOffering{}
	_ = req.response().(*NetworkOffering)
}
