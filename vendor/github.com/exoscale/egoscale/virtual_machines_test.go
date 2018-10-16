package egoscale

import (
	"net"
	"net/url"
	"testing"
)

func TestVirtualMachine(t *testing.T) {
	instance := &VirtualMachine{}
	if instance.ResourceType() != "UserVM" {
		t.Errorf("ResourceType doesn't match")
	}
}

func TestDeployVirtualMachine(t *testing.T) {
	req := &DeployVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestDestroyVirtualMachine(t *testing.T) {
	req := &DestroyVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestRebootVirtualMachine(t *testing.T) {
	req := &RebootVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestStartVirtualMachine(t *testing.T) {
	req := &StartVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestStopVirtualMachine(t *testing.T) {
	req := &StopVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestResetPasswordForVirtualMachine(t *testing.T) {
	req := &ResetPasswordForVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestUpdateVirtualMachine(t *testing.T) {
	req := &UpdateVirtualMachine{}
	_ = req.response().(*VirtualMachine)
}

func TestListVirtualMachines(t *testing.T) {
	req := &ListVirtualMachines{}
	_ = req.response().(*ListVirtualMachinesResponse)
}

func TestGetVMPassword(t *testing.T) {
	req := &GetVMPassword{}
	_ = req.response().(*Password)
}

func TestRestoreVirtualMachine(t *testing.T) {
	req := &RestoreVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestChangeServiceForVirtualMachine(t *testing.T) {
	req := &ChangeServiceForVirtualMachine{}
	_ = req.response().(*VirtualMachine)
}

func TestScaleVirtualMachine(t *testing.T) {
	req := &ScaleVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestRecoverVirtualMachine(t *testing.T) {
	req := &RecoverVirtualMachine{}
	_ = req.response().(*VirtualMachine)
}

func TestExpungeVirtualMachine(t *testing.T) {
	req := &ExpungeVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestGetVirtualMachineUserData(t *testing.T) {
	req := &GetVirtualMachineUserData{}
	_ = req.response().(*VirtualMachineUserData)
}

func TestMigrateVirtualMachine(t *testing.T) {
	req := &MigrateVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestAddNicToVirtualMachine(t *testing.T) {
	req := &AddNicToVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestRemoveNicFromVirtualMachine(t *testing.T) {
	req := &RemoveNicFromVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestUpdateDefaultNicForVirtualMachine(t *testing.T) {
	req := &UpdateDefaultNicForVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestDeployOnBeforeSend(t *testing.T) {
	req := &DeployVirtualMachine{
		SecurityGroupNames: []string{"default"},
	}
	params := url.Values{}

	if err := req.onBeforeSend(params); err != nil {
		t.Error(err)
	}
}

func TestDeployOnBeforeSendNoSG(t *testing.T) {
	req := &DeployVirtualMachine{}
	params := url.Values{}

	// CS will pick the default oiine
	if err := req.onBeforeSend(params); err != nil {
		t.Error(err)
	}
}

func TestDeployOnBeforeSendBothSG(t *testing.T) {
	req := &DeployVirtualMachine{
		SecurityGroupIDs:   []UUID{*MustParseUUID("f2b4e439-2b23-441c-ba66-0e25cdfe1b2b")},
		SecurityGroupNames: []string{"foo"},
	}
	params := url.Values{}

	if err := req.onBeforeSend(params); err == nil {
		t.Errorf("DeployVM should only accept SG ids or names")
	}
}

func TestDeployOnBeforeSendBothAG(t *testing.T) {
	req := &DeployVirtualMachine{
		AffinityGroupIDs:   []UUID{*MustParseUUID("f2b4e439-2b23-441c-ba66-0e25cdfe1b2b")},
		AffinityGroupNames: []string{"foo"},
	}
	params := url.Values{}

	if err := req.onBeforeSend(params); err == nil {
		t.Errorf("DeployVM should only accept SG ids or names")
	}
}

func TestGetVirtualMachine(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listvirtualmachinesresponse": {
	"count": 1,
	"virtualmachine": [
		{
			"account": "yoan.blanc@exoscale.ch",
			"affinitygroup": [],
			"cpunumber": 1,
			"cpuspeed": 2198,
			"cpuused": "0%",
			"created": "2018-01-19T14:37:08+0100",
			"diskioread": 0,
			"diskiowrite": 13734,
			"diskkbsread": 0,
			"diskkbswrite": 94342,
			"displayname": "test",
			"displayvm": true,
			"domain": "ROOT",
			"domainid": "1874276d-4cac-448b-aa5e-de00fd4157e8",
			"haenable": false,
			"hostid": "70c12af4-b1cb-4133-97dd-3579bb88a8ce",
			"hostname": "virt-hv-pp005.dk2.p.exoscale.net",
			"hypervisor": "KVM",
			"id": "69069d5e-1591-4214-937e-4c8cba63fcfb",
			"instancename": "i-2-188150-VM",
			"isdynamicallyscalable": false,
			"keypair": "test-yoan",
			"memory": 1024,
			"name": "test",
			"networkkbsread": 5542,
			"networkkbswrite": 8813,
			"nic": [
				{
					"broadcasturi": "vlan://untagged",
					"gateway": "159.100.248.1",
					"id": "75d1367c-319a-4658-b31d-1a26496061ff",
					"ipaddress": "159.100.251.247",
					"isdefault": true,
					"macaddress": "06:6d:cc:00:00:3c",
					"netmask": "255.255.252.0",
					"networkid": "d48bfccc-c11f-438f-8177-9cf6a40dc4f8",
					"networkname": "defaultGuestNetwork",
					"traffictype": "Guest",
					"type": "Shared"
				}
			],
			"oscategoryid": "ca158095-a6d2-4b0c-95e3-9a2e5123cbfc",
			"passwordenabled": true,
			"securitygroup": [
				{
					"account": "default",
					"description": "",
					"id": "41471067-d3c2-41ef-ae45-f57e00078843",
					"name": "default security group",
					"tags": []
				}
			],
			"serviceofferingid": "84925525-7825-418b-845b-1aed179bbc40",
			"serviceofferingname": "Tiny",
			"state": "Running",
			"tags": [],
			"templatedisplaytext": "Linux CentOS 7.4 64-bit 10G Disk (2018-01-08-d617dd)",
			"templateid": "934b4d48-d82e-42f5-8f14-b34de3af9854",
			"templatename": "Linux CentOS 7.4 64-bit",
			"zoneid": "381d0a95-ed3a-4ad9-b41c-b97073c1a433",
			"zonename": "ch-dk-2"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	vm := &VirtualMachine{
		ID: MustParseUUID("69069d5e-1591-4214-937e-4c8cba63fcfb"),
	}
	if err := cs.Get(vm); err != nil {
		t.Error(err)
	}

	if vm.Account != "yoan.blanc@exoscale.ch" {
		t.Errorf("Account doesn't match, got %v", vm.Account)
	}
}

func TestGetVirtualMachinePassword(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"getvmpasswordresponse": {
	"password": {
		"encryptedpassword": "test"
	}
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	req := &GetVMPassword{
		ID: MustParseUUID("69069d5e-1591-4214-937e-4c8cba63fcfb"),
	}
	resp, err := cs.Request(req)
	if err != nil {
		t.Error(err)
	}
	if resp.(*Password).EncryptedPassword != "test" {
		t.Errorf("Encrypted password missing")
	}
}

func TestListMachines(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listvirtualmachinesresponse": {
	"count": 3,
	"virtualmachine": [
		{
			"id": "84752707-a1d6-4e93-8207-bafeda83fe15"
		},
		{
			"id": "f93238e1-cc6e-484b-9650-fe8921631b7b"
		},
		{
			"id": "487eda20-eea1-43f7-9456-e870a359b173"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	req := &VirtualMachine{}
	vms, err := cs.List(req)
	if err != nil {
		t.Error(err)
	}

	if len(vms) != 3 {
		t.Errorf("Expected three vms, got %d", len(vms))
	}
}

func TestListMachinesFailure(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listvirtualmachinesresponse": {
	"count": 3,
	"virtualmachine": {}
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	req := &VirtualMachine{}
	vms, err := cs.List(req)
	if err == nil {
		t.Errorf("Expected an error got %v", err)
	}

	if len(vms) != 0 {
		t.Errorf("Expected 0 vms, got %d", len(vms))
	}
}

func TestListMachinesPaginate(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listvirtualmachinesresponse": {
	"count": 3,
	"virtualmachine": [
		{
			"id": "84752707-a1d6-4e93-8207-bafeda83fe15"
		},
		{
			"id": "f93238e1-cc6e-484b-9650-fe8921631b7b"
		},
		{
			"id": "487eda20-eea1-43f7-9456-e870a359b173"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	vm := &VirtualMachine{}
	req, err := vm.ListRequest()
	if err != nil {
		t.Error(err)
	}

	id := MustParseUUID("84752707-a1d6-4e93-8207-bafeda83fe15")
	cs.Paginate(req, func(i interface{}, err error) bool {
		if !i.(*VirtualMachine).ID.Equal(*id) {
			t.Errorf("Expected id %q, got %q", id, i.(*VirtualMachine).ID)
		}
		return false
	})

}

func TestNicHelpers(t *testing.T) {
	vm := &VirtualMachine{
		ID: MustParseUUID("25ce0763-f34d-435a-8b84-08466908355a"),
		Nic: []Nic{
			{
				ID:           MustParseUUID("e3b9c165-f3c3-4672-be54-08bfa6bac6fe"),
				IsDefault:    true,
				MACAddress:   MustParseMAC("06:aa:14:00:00:18"),
				IPAddress:    net.ParseIP("192.168.0.10"),
				Gateway:      net.ParseIP("192.168.0.1"),
				Netmask:      net.ParseIP("255.255.255.0"),
				NetworkID:    MustParseUUID("d48bfccc-c11f-438f-8177-9cf6a40dc4d8"),
				NetworkName:  "defaultGuestNetwork",
				BroadcastURI: "vlan://untagged",
				TrafficType:  "Guest",
				Type:         "Shared",
			}, {
				BroadcastURI: "vxlan://001",
				ID:           MustParseUUID("10b8ffc8-62b3-4b87-82d0-fb7f31bc99b6"),
				IsDefault:    false,
				MACAddress:   MustParseMAC("0a:7b:5e:00:25:fa"),
				NetworkID:    MustParseUUID("5f1033fe-2abd-4dda-80b6-c946e21a78ec"),
				NetworkName:  "privNetForBasicZone1",
				TrafficType:  "Guest",
				Type:         "Isolated",
			}, {
				BroadcastURI: "vxlan://002",
				ID:           MustParseUUID("10b8ffc8-62b3-4b87-82d0-fb7f31bc99b7"),
				IsDefault:    false,
				MACAddress:   MustParseMAC("0a:7b:5e:00:25:ff"),
				NetworkID:    MustParseUUID("5f1033fe-2abd-4dda-80b6-c946e21a72ec"),
				NetworkName:  "privNetForBasicZone2",
				TrafficType:  "Guest",
				Type:         "Isolated",
			},
		},
	}

	nic := vm.DefaultNic()
	if nic.IPAddress.String() != "192.168.0.10" {
		t.Errorf("Default NIC doesn't match")
	}

	ip := vm.IP()
	if ip.String() != "192.168.0.10" {
		t.Errorf("IP Address doesn't match")
	}

	nic1 := vm.NicByID(*MustParseUUID("e3b9c165-f3c3-4672-be54-08bfa6bac6fe"))
	if nic1.ID != nil && !nic.ID.Equal(*nic1.ID) {
		t.Errorf("NicByID does not match %#v %#v", nic, nic1)
	}

	if len(vm.NicsByType("Isolated")) != 2 {
		t.Errorf("Isolated nics count does not match")
	}

	if len(vm.NicsByType("Shared")) != 1 {
		t.Errorf("Shared nics count does not match")
	}

	if len(vm.NicsByType("Dummy")) != 0 {
		t.Errorf("Dummy nics count does not match")
	}

	if vm.NicByNetworkID(*MustParseUUID("5f1033fe-2abd-4dda-80b6-c946e21a78ec")) == nil {
		t.Errorf("NetworkID nic wasn't found")
	}

	if vm.NicByNetworkID(*MustParseUUID("5f1033fe-2abd-4dda-80b6-c946e21a78ed")) != nil {
		t.Errorf("NetworkID nic was found??")
	}
}

func TestNicNoDefault(t *testing.T) {
	vm := &VirtualMachine{
		Nic: []Nic{},
	}

	// code coverage...
	nic := vm.DefaultNic()
	if nic != nil {
		t.Errorf("Default NIC wasn't nil?")
	}
}
