package egoscale

import (
	"testing"
)

func TestAddIPToNic(t *testing.T) {
	req := &AddIPToNic{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*NicSecondaryIP)
}

func TestRemoveIPFromNic(t *testing.T) {
	req := &RemoveIPFromNic{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestListNicsAPIName(t *testing.T) {
	req := &ListNics{}
	_ = req.response().(*ListNicsResponse)
}

func TestActivateIP6(t *testing.T) {
	req := &ActivateIP6{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Nic)
}

func TestListNicInvalid(t *testing.T) {
	ts := newServer()
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	nic := new(Nic)

	_, err := cs.List(nic)
	if err == nil {
		t.Error("An error was expected")
	}
}

func TestListNicError(t *testing.T) {
	ts := newServer(response{431, jsonContentType, `
{"listnicresponse": {
	"cserrorcode": 9999,
	"errorcode": 431,
	"errortext": "Unable to execute API command listnics due to missing parameter virtualmachineid",
	"uuidList": []
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")

	nic := &Nic{
		VirtualMachineID: MustParseUUID("f9c61d37-d2f5-4bed-b8d2-73edc0a0f61e"),
	}

	_, err := cs.List(nic)
	if err == nil {
		t.Error("An error was expected")
	}
}
