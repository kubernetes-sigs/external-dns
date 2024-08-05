package civogo

import (
	"reflect"
	"testing"
)

func TestGetTemplateByCode(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/templates": `[{"id": "1", "code": "centos-7"},{"id": "2", "code": "ubuntu-18.04"}]`,
	})
	defer server.Close()

	got, err := client.GetTemplateByCode("ubuntu-18.04")
	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}
	if got.ID != "2" {
		t.Errorf("Expected %s, got %s", "12345", got.ID)
	}
}

func TestListTemplates(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/templates": `[{
		"id": "773aea72-d068-4cb7-8e08-28841acef0cb",
		"code": "ubuntu-18.04",
		"name": "Ubuntu 18.04",
		"account_id": null,
		"image_id": null,
		"volume_id": "61c2cfe2-f7f3-46e7-b5c9-7b03ae25ea86",
		"short_description": "Ubuntu 18.04",
		"description": "The freely available Ubuntu 18.04 OS, minimally installed with just OpenSSH server",
		"default_username": "ubuntu",
		"cloud_config": "#cloud-config contents"
	  }]`,
	})
	defer server.Close()
	got, err := client.ListTemplates()

	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}
	expected := []Template{
		{
			ID:               "773aea72-d068-4cb7-8e08-28841acef0cb",
			Code:             "ubuntu-18.04",
			Name:             "Ubuntu 18.04",
			VolumeID:         "61c2cfe2-f7f3-46e7-b5c9-7b03ae25ea86",
			ShortDescription: "Ubuntu 18.04",
			Description:      "The freely available Ubuntu 18.04 OS, minimally installed with just OpenSSH server",
			DefaultUsername:  "ubuntu",
			CloudConfig:      "#cloud-config contents",
		},
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}

}

func TestNewTemplate(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/templates": `{
		  "result": "success",
		  "id": "283d5ee6-fa9f-4e40-8e1a-bdc28812d593"
		}`,
	})
	defer server.Close()

	conf := &Template{
		Code:             "test",
		Name:             "Test",
		ImageID:          "811a8dfb-8202-49ad-b1ef-1e6320b20497",
		ShortDescription: "my custom",
		Description:      "my custom image from golang",
		DefaultUsername:  "root",
	}

	got, err := client.NewTemplate(conf)
	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}

	expected := &SimpleResponse{Result: "success", ID: "283d5ee6-fa9f-4e40-8e1a-bdc28812d593"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}

}

func TestUpdateTemplate(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/templates/283d5ee6-fa9f-4e40-8e1a-bdc28812d593": `{
		  "id": "283d5ee6-fa9f-4e40-8e1a-bdc28812d593",
		  "code": "my-linux-1.0",
		  "name": "My Linux 1.0",
		  "account_id": null,
		  "image_id": "811a8dfb-8202-49ad-b1ef-1e6320b20497",
		  "volume_id": null,
		  "short_description": "...",
		  "description": "...",
		  "default_username": "Ubuntu",
		  "cloud_config": "..."
		}`,
	})
	defer server.Close()

	conf := &Template{
		Code:             "my-linux-1.0",
		Name:             "My Linux 1.0",
		ImageID:          "811a8dfb-8202-49ad-b1ef-1e6320b20497",
		ShortDescription: "...",
		Description:      "...",
		DefaultUsername:  "Ubuntu",
	}

	got, err := client.UpdateTemplate("283d5ee6-fa9f-4e40-8e1a-bdc28812d593", conf)
	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}

	expected := &Template{
		ID:               "283d5ee6-fa9f-4e40-8e1a-bdc28812d593",
		Code:             "my-linux-1.0",
		Name:             "My Linux 1.0",
		ImageID:          "811a8dfb-8202-49ad-b1ef-1e6320b20497",
		ShortDescription: "...",
		Description:      "...",
		DefaultUsername:  "Ubuntu",
		CloudConfig:      "...",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}

}

func TestDeleteTemplate(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/templates/12346": `{"result": "success"}`,
	})
	defer server.Close()
	got, err := client.DeleteTemplate("12346")
	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}

	expected := &SimpleResponse{Result: "success"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}
}
