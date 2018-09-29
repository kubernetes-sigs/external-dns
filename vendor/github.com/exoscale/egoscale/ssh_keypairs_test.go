package egoscale

import (
	"testing"
)

func TestResetSSHKeyForVirtualMachine(t *testing.T) {
	req := &ResetSSHKeyForVirtualMachine{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestRegisterSSHKeyPair(t *testing.T) {
	req := &RegisterSSHKeyPair{}
	_ = req.response().(*SSHKeyPair)
}

func TestCreateSSHKeyPair(t *testing.T) {
	req := &CreateSSHKeyPair{}
	_ = req.response().(*SSHKeyPair)
}

func TestDeleteSSHKeyPair(t *testing.T) {
	req := &DeleteSSHKeyPair{}
	_ = req.response().(*booleanResponse)
}

func TestListSSHKeyPairsResponse(t *testing.T) {
	req := &ListSSHKeyPairs{}
	_ = req.response().(*ListSSHKeyPairsResponse)
}

func TestGetSSHKeyPair(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listsshkeypairsresponse": {
	"count": 1,
	"sshkeypair": [
		{
			"fingerprint": "07:97:32:04:80:23:b9:a2:a2:46:fe:ab:a6:4b:20:76",
			"name": "yoan@herp"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	ssh := &SSHKeyPair{
		Name: "yoan@herp",
	}
	if err := cs.Get(ssh); err != nil {
		t.Error(err)
	}

	if ssh.Fingerprint != "07:97:32:04:80:23:b9:a2:a2:46:fe:ab:a6:4b:20:76" {
		t.Errorf("Fingerprint doesn't match, got %v", ssh.Fingerprint)
	}
}

func TestListSSHKeyPairs(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listsshkeypairsresponse": {
	"count": 2,
	"sshkeypair": [
		{
			"fingerprint": "07:97:32:04:80:23:b9:a2:a2:46:fe:ab:a6:4b:20:76",
			"name": "yoan@herp"
		},
		{
			"fingerprint": "9e:97:54:95:82:22:eb:f8:9b:4f:28:6f:c7:f5:58:83",
			"name": "yoan@derp"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	ssh := &SSHKeyPair{}

	sshs, err := cs.List(ssh)
	if err != nil {
		t.Error(err)
	}

	if len(sshs) != 2 {
		t.Errorf("Expected two ssh keys, got %v", len(sshs))
	}
}

func TestListSSHKeysFailure(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listsshkeypairsresponse": {
	"count": 2,
	"sshkeypair": {}
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	ssh := &SSHKeyPair{}

	sshs, err := cs.List(ssh)
	if err == nil {
		t.Errorf("Expected an error, got %v", err)
	}

	if len(sshs) != 0 {
		t.Errorf("Expected two ssh keys, got %v", len(sshs))
	}
}

func TestListSSHKeyPaginate(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listsshkeypairsresponse": {
	"count": 2,
	"sshkeypair": [
		{
			"fingerprint": "07:97:32:04:80:23:b9:a2:a2:46:fe:ab:a6:4b:20:76",
			"name": "yoan@herp"
		},
		{
			"fingerprint": "9e:97:54:95:82:22:eb:f8:9b:4f:28:6f:c7:f5:58:83",
			"name": "hello"
		}
	]
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	ssh := &SSHKeyPair{}

	req, err := ssh.ListRequest()
	if err != nil {
		t.Error(err)
	}

	cs.Paginate(req, func(i interface{}, err error) bool {

		if i.(*SSHKeyPair).Name != "yoan@herp" {
			t.Errorf("Expected yoan@herp name, got %s", i.(*SSHKeyPair).Name)
		}
		return false
	})
}

func TestGetSSHKeyPairNotFound(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
{"listsshkeypairsresponse": {
	"count": 0,
	"sshkeypair": []
}}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	ssh := &SSHKeyPair{
		Name: "foo",
	}
	if err := cs.Get(ssh); err == nil {
		t.Errorf("An error was expected")
	}
}
