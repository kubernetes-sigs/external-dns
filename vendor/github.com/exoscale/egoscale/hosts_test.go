package egoscale

import (
	"testing"
)

func TestListHosts(t *testing.T) {
	req := &ListHosts{}
	_ = req.response().(*ListHostsResponse)
}

func TestUpdateHost(t *testing.T) {
	req := &UpdateHost{}
	_ = req.response().(*Host)
}
