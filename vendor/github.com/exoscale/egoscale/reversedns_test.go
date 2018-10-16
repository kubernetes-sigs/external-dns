package egoscale

import (
	"testing"
)

func TestDeleteReverseDNSFromVirtualMachine(t *testing.T) {
	req := &DeleteReverseDNSFromVirtualMachine{}
	_ = req.response().(*booleanResponse)
}

func TestDeleteReverseDNSFromPublicIPAddress(t *testing.T) {
	req := &DeleteReverseDNSFromPublicIPAddress{}
	_ = req.response().(*booleanResponse)
}

func TestQueryReverseDNSForVirtualMachine(t *testing.T) {
	req := &QueryReverseDNSForVirtualMachine{}
	_ = req.response().(*VirtualMachine)
}

func TestQueryReverseDNSForPublicIPAddress(t *testing.T) {
	req := &QueryReverseDNSForPublicIPAddress{}
	_ = req.response().(*IPAddress)
}

func TestUpdateReverseDNSForVirtualMachine(t *testing.T) {
	req := &UpdateReverseDNSForVirtualMachine{}
	_ = req.response().(*VirtualMachine)
}

func TestUpdateReverseDNSForPublicIPAddress(t *testing.T) {
	req := &UpdateReverseDNSForPublicIPAddress{}
	_ = req.response().(*IPAddress)
}
