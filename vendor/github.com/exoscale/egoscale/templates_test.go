package egoscale

import (
	"testing"
)

func TestTemplateResourceType(t *testing.T) {
	instance := &Template{}
	if instance.ResourceType() != "Template" {
		t.Errorf("ResourceType doesn't match")
	}
}

func TestCreateTemplate(t *testing.T) {
	req := &CreateTemplate{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Template)
}

func TestCopyTemplate(t *testing.T) {
	req := &CopyTemplate{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Template)
}

func TestUpdateTemplate(t *testing.T) {
	req := &UpdateTemplate{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Template)
}

func TestListTemplates(t *testing.T) {
	req := &ListTemplates{}
	_ = req.response().(*ListTemplatesResponse)
}

func TestDeleteTemplate(t *testing.T) {
	req := &DeleteTemplate{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestPrepareTemplate(t *testing.T) {
	req := &PrepareTemplate{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Template)
}

func TestRegisterTemplate(t *testing.T) {
	req := &RegisterTemplate{}
	_ = req.response().(*Template)
}

func TestListOSCategories(t *testing.T) {
	req := &ListOSCategories{}
	_ = req.response().(*ListOSCategoriesResponse)
}

func TestTemplate(t *testing.T) {
	instance := &Template{}
	if instance.ResourceType() != "Template" {
		t.Errorf("ResourceType doesn't match")
	}
}

func TestTemplateGet(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
		
		{ "listtemplatesresponse": {
			"count": 1,
			"template": [
			  {
				"account": "exostack",
				"checksum": "426b9c1efae42f84ea16e74a336f5202",
				"created": "2018-01-30T09:15:36+0100",
				"details": {
				  "username": "debian"
				},
				"displaytext": "Linux Debian 9 64-bit 10G Disk (2018-01-18-25e9ec)",
				"domain": "ROOT",
				"domainid": "4a8857b8-7235-4e31-a7ef-b8b44d180850",
				"format": "QCOW2",
				"hypervisor": "KVM",
				"id": "78c2cbe6-8e11-4722-b01f-bf06f4e28108",
				"isfeatured": true,
				"ispublic": true,
				"isready": true,
				"name": "Linux Debian 9 64-bit",
				"ostypeid": "113038d0-a8cd-4d20-92be-ea313f87c3ac",
				"ostypename": "Other PV (64-bit)",
				"passwordenabled": true,
				"size": 10737418240,
				"zoneid": "1128bd56-b4d9-4ac6-a7b9-c715b187ce11",
				"zonename": "ch-gva-2"
			  }
			]
		  }}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	zoneID := MustParseUUID("1128bd56-b4d9-4ac6-a7b9-c715b187ce11")
	id := MustParseUUID("78c2cbe6-8e11-4722-b01f-bf06f4e28108")
	template := &Template{IsFeatured: true,
		ZoneID: zoneID,
		ID:     id,
	}

	if err := cs.Get(template); err != nil {
		t.Error(err)
	}

	if template.Name == "" {
		t.Errorf("Expected a name, got empty name")
	}
}

func TestListTemplate(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
		
		{ "listtemplatesresponse": {
			"count": 1,
			"template": [
			  {
				"account": "exostack",
				"checksum": "3c80c71fcfe1e2e88c12ca7d39c294a0",
				"created": "2018-01-30T09:16:05+0100",
				"crossZones": false,
				"details": {
				  "username": "debian"
				},
				"displaytext": "Linux Debian 9 64-bit 200G Disk (2018-01-18-25e9ec)",
				"domain": "ROOT",
				"domainid": "4a8857b8-7235-4e31-a7ef-b8b44d180850",
				"format": "QCOW2",
				"hypervisor": "KVM",
				"id": "a8a4b773-32ce-4d1c-a19b-21da055ec5c6",
				"isdynamicallyscalable": false,
				"isextractable": false,
				"isfeatured": true,
				"ispublic": true,
				"isready": true,
				"name": "Linux Debian 9 64-bit",
				"ostypeid": "113038d0-a8cd-4d20-92be-ea313f87c3ac",
				"ostypename": "Other PV (64-bit)",
				"passwordenabled": true,
				"size": 214748364800,
				"sshkeyenabled": false,
				"tags": [],
				"templatetype": "USER",
				"zoneid": "4da1b188-dcd6-4ff5-b7fd-bde984055548",
				"zonename": "at-vie-1"
			  }
			]
		  }}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	zoneID := MustParseUUID("4da1b188-dcd6-4ff5-b7fd-bde984055548")
	id := MustParseUUID("a8a4b773-32ce-4d1c-a19b-21da055ec5c6")
	template := &Template{IsFeatured: true,
		ZoneID: zoneID,
		ID:     id,
	}

	temps, err := cs.List(template)
	if err != nil {
		t.Error(err)
	}

	if len(temps) != 1 {
		t.Fatalf("Expected one template, got %v", len(temps))
	}

	temp := temps[0].(*Template)

	if !temp.ZoneID.Equal(*zoneID) || !temp.ID.Equal(*id) {
		t.Errorf("Wrong result")
	}
}

func TestListTemplatesPaginate(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `{
	"listtemplatesresponse": {
		"count": 2,
		"template": [
			{
				"account": "exostack",
				"checksum": "3c80c71fcfe1e2e88c12ca7d39c294a0",
				"created": "2018-01-30T09:16:05+0100",
				"crossZones": false,
				"details": {
				  "username": "debian"
				},
				"displaytext": "Linux Debian 9 64-bit 200G Disk (2018-01-18-25e9ec)",
				"domain": "ROOT",
				"domainid": "4a8857b8-7235-4e31-a7ef-b8b44d180850",
				"format": "QCOW2",
				"hypervisor": "KVM",
				"id": "a8a4b773-32ce-4d1c-a19b-21da055ec5c6",
				"isdynamicallyscalable": false,
				"isextractable": false,
				"isfeatured": true,
				"ispublic": true,
				"isready": true,
				"name": "Linux Debian 9 64-bit",
				"ostypeid": "113038d0-a8cd-4d20-92be-ea313f87c3ac",
				"ostypename": "Other PV (64-bit)",
				"passwordenabled": true,
				"size": 214748364800,
				"sshkeyenabled": false,
				"tags": [],
				"templatetype": "USER",
				"zoneid": "4da1b188-dcd6-4ff5-b7fd-bde984055548",
				"zonename": "at-vie-1"
			  },
			  {
				"account": "exostack",
				"checksum": "3c80c71fcfe1e2e88c12ca7d39c294a0",
				"created": "2018-01-30T09:16:05+0100",
				"crossZones": false,
				"details": {
				  "username": "debian"
				},
				"displaytext": "Linux Debian 9 64-bit 200G Disk (2018-01-18-25e9ec)",
				"domain": "ROOT",
				"domainid": "4a8857b8-7235-4e31-a7ef-b8b44d180850",
				"format": "QCOW2",
				"hypervisor": "KVM",
				"id": "4a8857b9-7235-4e31-a7ef-b8b44d180850",
				"isdynamicallyscalable": false,
				"isextractable": false,
				"isfeatured": true,
				"ispublic": true,
				"isready": true,
				"name": "Linux Debian 9 64-bit",
				"ostypeid": "113038d0-a8cd-4d20-92be-ea313f87c3ac",
				"ostypename": "Other PV (64-bit)",
				"passwordenabled": true,
				"size": 214748364800,
				"sshkeyenabled": false,
				"tags": [],
				"templatetype": "USER",
				"zoneid": "4da1b188-dcd6-4ff5-b7fd-bde984055548",
				"zonename": "at-vie-1"
			}
		]
	}
}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	id := MustParseUUID("a8a4b773-32ce-4d1c-a19b-21da055ec5c6")
	template := &Template{IsFeatured: true}

	req, err := template.ListRequest()
	if err != nil {
		t.Error(err)
	}

	cs.Paginate(req, func(i interface{}, err error) bool {
		if err != nil {
			t.Fatal(err)
		}

		if !id.Equal(*i.(*Template).ID) {
			t.Errorf("Expected id '%s' but got %s", id, i.(*Template).ID)
		}
		return false
	})
}

func TestListTemplatesFailure(t *testing.T) {
	ts := newServer(response{200, jsonContentType, `
		
		{ "listtemplatesresponse": {
			"count": 1,
			"template": {}
		  }}`})
	defer ts.Close()

	cs := NewClient(ts.URL, "KEY", "SECRET")
	temps, err := cs.List(&Template{})
	if err == nil {
		t.Errorf("Expected an error got, %v", err)
	}

	if len(temps) != 0 {
		t.Fatalf("Expected 0 template, got %v", len(temps))
	}
}
