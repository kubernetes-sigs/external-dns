package egoscale

import (
	"net"
	"testing"
)

func TestIPAddress(t *testing.T) {
	instance := &IPAddress{}
	if instance.ResourceType() != "PublicIpAddress" {
		t.Errorf("ResourceType doesn't match")
	}
}

func TestAssociateIPAddress(t *testing.T) {
	req := &AssociateIPAddress{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*IPAddress)
}

func TestDisassociateIPAddress(t *testing.T) {
	req := &DisassociateIPAddress{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestListPublicIPAddresses(t *testing.T) {
	req := &ListPublicIPAddresses{}
	_ = req.response().(*ListPublicIPAddressesResponse)
}

func TestUpdateIPAddress(t *testing.T) {
	req := &UpdateIPAddress{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*IPAddress)
}

func TestGetIPAddress(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listpublicipaddressesresponse": {
	"count": 1,
	"publicipaddress": [
		{
			"account": "yoan.blanc@exoscale.ch",
			"allocated": "2017-12-18T08:16:56+0100",
			"domain": "yoan.blanc@exoscale.ch",
			"domainid": "17b3bc0c-8aed-4920-be3a-afdd6daf4314",
			"forvirtualnetwork": false,
			"id": "adc31802-9413-47ea-a252-266f7eadaf00",
			"ipaddress": "109.100.242.45",
			"iselastic": true,
			"isportable": false,
			"issourcenat": false,
			"isstaticnat": false,
			"issystem": false,
			"networkid": "00304a04-e7ea-4e77-a786-18bc64347bf7",
			"physicalnetworkid": "01f747f5-b445-487f-b2d7-81a5a512989e",
			"state": "Allocated",
			"tags": [],
			"zoneid": "1128bd56-b4e9-4ac6-a7b9-c715b187ce11",
			"zonename": "ch-gva-2"
		}
	]
}}`})

	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	eip := &IPAddress{
		IPAddress: net.ParseIP("109.100.242.45"),
		IsElastic: true,
	}
	if err := cs.Get(eip); err != nil {
		t.Error(err)
	}

	if eip.Account != "yoan.blanc@exoscale.ch" {
		t.Errorf("Account doesn't match, got %v", eip.Account)
	}
}

func TestListIPAddressFailure(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
	{"listpublicipaddressesresponse": {
		"count": 1,
		"publicipaddress": {}
	}`})

	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	eip := &IPAddress{}

	_, err := cs.List(eip)
	if err == nil {
		t.Errorf("Expected an error, got %v", err)
	}

}
