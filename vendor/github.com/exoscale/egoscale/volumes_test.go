package egoscale

import (
	"testing"
)

func TestVolume(t *testing.T) {
	instance := &Volume{}
	if instance.ResourceType() != "Volume" {
		t.Errorf("ResourceType doesn't match")
	}
}

func TestListVolumes(t *testing.T) {
	req := &ListVolumes{}
	_ = req.response().(*ListVolumesResponse)
}

func TestResizeVolume(t *testing.T) {
	req := &ResizeVolume{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Volume)
}

func TestListVolumeFailure(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listvolumesresponse": {
	"count": 1,
	"volume": {}
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	volume := &Volume{
		VirtualMachineID: MustParseUUID("9ccc3d5b-9dce-4302-a955-24b80b402f88"),
		Type:             "ROOT",
	}
	_, err := cs.List(volume)
	if err == nil {
		t.Errorf("Expected an error but got %v", err)
	}
}

func TestListVolumePaginate(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listvolumesresponse": {
	"count": 1,
	"volume": [
		{
			"account": "test",
			"created": "2018-03-23T00:41:14+0100",
			"destroyed": false,
			"deviceid": 0,
			"domain": "test",
			"domainid": "2083e04d-500f-48ef-8e3d-bae6805416cd",
			"id": "3613a751-5822-4d1d-b312-3036ef1acf86",
			"isextractable": true,
			"name": "ROOT-246634",
			"quiescevm": false,
			"serviceofferingdisplaytext": "Medium 4096mb 2cpu",
			"serviceofferingid": "5e5fb3c6-e076-429d-9b6c-b71f7b26760b",
			"serviceofferingname": "Medium",
			"size": 10737418240,
			"state": "Ready",
			"storagetype": "local",
			"tags": [],
			"templatedisplaytext": "Linux Ubuntu 16.04 LTS 64-bit 10G Disk (2018-03-02-5858e9)",
			"templateid": "4a0c4d65-8d88-40a5-b1be-549b211620b6",
			"templatename": "Linux Ubuntu 16.04 LTS 64-bit",
			"type": "ROOT",
			"virtualmachineid": "9ccc3d5b-9dce-4302-a955-24b80b402f88",
			"vmdisplayname": "test",
			"vmname": "test",
			"vmstate": "Running",
			"zoneid": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"zonename": "ch-gva-2"
		},
		{
			"account": "test",
			"created": "2018-03-23T00:41:14+0100",
			"destroyed": false,
			"deviceid": 0,
			"domain": "test",
			"domainid": "2083e04d-500f-48ef-8e3d-bae6805416cd",
			"id": "687a5461-6fc7-4376-9314-55a75bd5dee0",
			"isextractable": true,
			"name": "ROOT-246634",
			"quiescevm": false,
			"serviceofferingdisplaytext": "Medium 4096mb 2cpu",
			"serviceofferingid": "5e5fb3c6-e076-429d-9b6c-b71f7b26760b",
			"serviceofferingname": "Medium",
			"size": 10737418240,
			"state": "Ready",
			"storagetype": "local",
			"tags": [],
			"templatedisplaytext": "Linux Ubuntu 16.04 LTS 64-bit 10G Disk (2018-03-02-5858e9)",
			"templateid": "4a0c4d65-8d88-40a5-b1be-549b211620b6",
			"templatename": "Linux Ubuntu 16.04 LTS 64-bit",
			"type": "ROOT",
			"virtualmachineid": "9ccc3d5b-9dce-4302-a955-24b80b402f88",
			"vmdisplayname": "test",
			"vmname": "test",
			"vmstate": "Running",
			"zoneid": "1747ef5e-5451-41fd-9f1a-58913bae9702",
			"zonename": "ch-gva-2"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	volume := &Volume{
		VirtualMachineID: MustParseUUID("9ccc3d5b-9dce-4302-a955-24b80b402f88"),
		Type:             "ROOT",
	}

	req, err := volume.ListRequest()
	if err != nil {
		t.Error(err)
	}

	id := MustParseUUID("3613a751-5822-4d1d-b312-3036ef1acf86")
	cs.Paginate(req, func(i interface{}, err error) bool {
		if err != nil {
			t.Error(err)
			return true
		}

		if !i.(*Volume).ID.Equal(*id) {
			t.Errorf("Expected id %q but got %s", id, i.(*Volume).ID)
		}
		return false
	})
}

func TestListVolumeError(t *testing.T) {
	ts := newServer(response{431, jsonContentType, `
{"listvolumeresponse": {
	"cserrorcode": 9999,
	"errorcode": 431,
	"errortext": "Unable to execute API command listvolumes due to invalid value. Invalid parameter virtualmachineid value=9ccc3d5b-9dce-4302-a955-24b80b402f87 due to incorrect long value format, or entity does not exist or due to incorrect parameter annotation for the field in api cmd class.",
	"uuidList": []
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	volume := new(Volume)
	_, err := cs.List(volume)
	if err == nil {
		t.Error("An error was expected")
	}
}
