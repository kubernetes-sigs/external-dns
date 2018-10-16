package egoscale

import (
	"testing"
)

func TestListEvents(t *testing.T) {
	req := &ListEvents{}
	_ = req.response().(*ListEventsResponse)
}

func TestListEventTypes(t *testing.T) {
	req := &ListEventTypes{}
	_ = req.response().(*ListEventTypesResponse)
}
