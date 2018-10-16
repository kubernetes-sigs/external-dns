package egoscale

import (
	"net/url"
	"testing"
)

func TestNetwork(t *testing.T) {
	instance := &Network{}
	if instance.ResourceType() != "Network" {
		t.Errorf("ResourceType doesn't match")
	}
}

func TestListNetworks(t *testing.T) {
	req := &ListNetworks{}
	_ = req.response().(*ListNetworksResponse)
}

func TestCreateNetwork(t *testing.T) {
	req := &CreateNetwork{}
	_ = req.response().(*Network)
}

func TestRestartNetwork(t *testing.T) {
	req := &RestartNetwork{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Network)
}

func TestUpdateNetwork(t *testing.T) {
	req := &UpdateNetwork{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Network)
}

func TestDeleteNetwork(t *testing.T) {
	req := &DeleteNetwork{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestCreateNetworkOnBeforeSend(t *testing.T) {
	req := &CreateNetwork{}
	params := url.Values{}

	if err := req.onBeforeSend(params); err != nil {
		t.Error(err)
	}

	if _, ok := params["name"]; !ok {
		t.Errorf("name should have been set")
	}
	if _, ok := params["displaytext"]; !ok {
		t.Errorf("displaytext should have been set")
	}
}

func TestListNetworkEmpty(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listnetworksresponse": {
	"count": 0,
	"network": []
  }}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	network := new(Network)
	networks, err := cs.List(network)
	if err != nil {
		t.Fatal(err)
	}

	if len(networks) != 0 {
		t.Errorf("zero networks were expected, got %d", len(networks))
	}
}

func TestListNetworkFailure(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listnetworksresponse": {
	"count": 3456,
	"network": {}
  }}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	network := new(Network)
	networks, err := cs.List(network)
	if err == nil {
		t.Errorf("error was expected, got %v", err)
	}

	if len(networks) != 0 {
		t.Errorf("zero networks were expected, got %d", len(networks))
	}
}

func TestFindNetwork(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listnetworksresponse": {
	"count": 1,
	"network": [
	  {
		"account": "exoscale-1",
		"acltype": "Account",
		"broadcastdomaintype": "Vxlan",
		"canusefordeploy": true,
		"displaytext": "fddd",
		"domain": "exoscale-1",
		"domainid": "5b2f621e-3eb6-4a14-a315-d4d7d62f28ff",
		"id": "772cee7a-631f-4c0e-ad2b-27776f260d71",
		"ispersistent": true,
		"issystem": false,
		"name": "klmfsdvds",
		"networkofferingavailability": "Optional",
		"networkofferingconservemode": true,
		"networkofferingdisplaytext": "Private Network",
		"networkofferingid": "eb35f4e6-0ecc-412e-9925-e469bf03d8fd",
		"networkofferingname": "PrivNet",
		"physicalnetworkid": "07f747f5-b445-487f-b2d7-81a5a512989e",
		"related": "772cee7a-631f-4c0e-ad2b-27776f260d71",
		"restartrequired": false,
		"service": [
		  {
			"name": "PrivateNetwork"
		  }
		],
		"specifyipranges": false,
		"state": "Implemented",
		"strechedl2subnet": false,
		"tags": [],
		"traffictype": "Guest",
		"type": "Isolated",
		"zoneid": "1128bd56-b4d9-4ac6-a7b9-c715b187ce11",
		"zonename": "ch-gva-2"
	  }
	]
  }}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	networks, err := cs.List(&Network{Name: "klmfsdvds", CanUseForDeploy: true, Type: "Isolated"})
	if err != nil {
		t.Fatal(err)
	}

	if len(networks) != 1 {
		t.Fatalf("One network was expected, got %d", len(networks))
	}

	net, ok := networks[0].(*Network)
	if !ok {
		t.Errorf("unable to type inference *Network, got %v", net)
	}

	if networks[0].(*Network).Name != "klmfsdvds" {
		t.Errorf("klmfsdvds network name was expected, got %s", networks[0].(*Network).Name)
	}

}
