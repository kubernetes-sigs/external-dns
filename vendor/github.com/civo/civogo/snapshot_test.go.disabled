package civogo

import (
	"reflect"
	"testing"
)

func TestCreateSnapshot(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/snapshots/my-backup": `{
		  "id": "0ca69adc-ff39-4fc1-8f08-d91434e86fac",
		  "instance_id": "44aab548-61ca-11e5-860e-5cf9389be614",
		  "hostname": "server1.prod.example.com",
		  "template_id": "0b213794-d795-4483-8982-9f249c0262b9",
		  "openstack_snapshot_id": null,
		  "region": "lon1",
		  "name": "my-backup",
		  "safe": 1,
		  "size_gb": 0,
		  "state": "new",
		  "cron_timing": null,
		  "requested_at": null,
		  "completed_at": null
		}`,
	})
	defer server.Close()

	cfg := &SnapshotConfig{
		InstanceID: "44aab548-61ca-11e5-860e-5cf9389be614",
		Safe:       true,
		Cron:       "",
	}
	got, err := client.CreateSnapshot("my-backup", cfg)
	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}

	expected := &Snapshot{
		ID:            "0ca69adc-ff39-4fc1-8f08-d91434e86fac",
		InstanceID:    "44aab548-61ca-11e5-860e-5cf9389be614",
		Hostname:      "server1.prod.example.com",
		Template:      "0b213794-d795-4483-8982-9f249c0262b9",
		Region:        "lon1",
		Name:          "my-backup",
		Safe:          1,
		SizeGigabytes: 0,
		State:         "new",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}
}

func TestListSnapshots(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/snapshots": `[{
		  "id": "0ca69adc-ff39-4fc1-8f08-d91434e86fac",
		  "instance_id": "44aab548-61ca-11e5-860e-5cf9389be614",
		  "hostname": "server1.prod.example.com",
		  "template_id": "0b213794-d795-4483-8982-9f249c0262b9",
		  "openstack_snapshot_id": null,
		  "region": "lon1",
		  "name": "my-backup",
		  "safe": 1,
		  "size_gb": 0,
		  "state": "new",
		  "cron_timing": null,
		  "requested_at": null,
		  "completed_at": null
		}]`,
	})
	defer server.Close()
	got, err := client.ListSnapshots()

	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}
	expected := []Snapshot{{
		ID:            "0ca69adc-ff39-4fc1-8f08-d91434e86fac",
		InstanceID:    "44aab548-61ca-11e5-860e-5cf9389be614",
		Hostname:      "server1.prod.example.com",
		Template:      "0b213794-d795-4483-8982-9f249c0262b9",
		Region:        "lon1",
		Name:          "my-backup",
		Safe:          1,
		SizeGigabytes: 0,
		State:         "new",
	}}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}
}

func TestFindSnapshot(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/snapshots": `[
			{
				"id": "0ca69adc-ff39-4fc1-8f08-d91434e86fac",
				"instance_id": "44aab548-61ca-11e5-860e-5cf9389be614",
				"hostname": "server1.prod.example.com",
				"template_id": "0b213794-d795-4483-8982-9f249c0262b9",
				"openstack_snapshot_id": null,
				"region": "lon1",
				"name": "my-backup",
				"safe": 1,
				"size_gb": 0,
				"state": "new",
				"cron_timing": null,
				"requested_at": null,
				"completed_at": null
			},
			{
				"id": "aadec58e-26f4-43e7-8963-18739519ef76",
				"instance_id": "44aab548-61ca-11e5-860e-5cf9389be614",
				"hostname": "server1.prod.example.com",
				"template_id": "0b213794-d795-4483-8982-9f249c0262b9",
				"openstack_snapshot_id": null,
				"region": "lon1",
				"name": "other-backup",
				"safe": 1,
				"size_gb": 0,
				"state": "new",
				"cron_timing": null,
				"requested_at": null,
				"completed_at": null
			}
		]`,
	})
	defer server.Close()

	got, _ := client.FindSnapshot("ff39")
	if got.ID != "0ca69adc-ff39-4fc1-8f08-d91434e86fac" {
		t.Errorf("Expected %s, got %s", "0ca69adc-ff39-4fc1-8f08-d91434e86fac", got.ID)
	}

	got, _ = client.FindSnapshot("26f4")
	if got.ID != "aadec58e-26f4-43e7-8963-18739519ef76" {
		t.Errorf("Expected %s, got %s", "aadec58e-26f4-43e7-8963-18739519ef76", got.ID)
	}

	got, _ = client.FindSnapshot("my")
	if got.ID != "0ca69adc-ff39-4fc1-8f08-d91434e86fac" {
		t.Errorf("Expected %s, got %s", "0ca69adc-ff39-4fc1-8f08-d91434e86fac", got.ID)
	}

	got, _ = client.FindSnapshot("other")
	if got.ID != "aadec58e-26f4-43e7-8963-18739519ef76" {
		t.Errorf("Expected %s, got %s", "aadec58e-26f4-43e7-8963-18739519ef76", got.ID)
	}

	_, err := client.FindSnapshot("backup")
	if err.Error() != "MultipleMatchesError: unable to find backup because there were multiple matches" {
		t.Errorf("Expected %s, got %s", "unable to find backup because there were multiple matches", err.Error())
	}

	_, err = client.FindSnapshot("missing")
	if err.Error() != "ZeroMatchesError: unable to find missing, zero matches" {
		t.Errorf("Expected %s, got %s", "unable to find missing, zero matches", err.Error())
	}
}

func TestDeleteSnapshot(t *testing.T) {
	client, server, _ := NewClientForTesting(map[string]string{
		"/v2/snapshots/my-backup": `{"result": "success"}`,
	})
	defer server.Close()
	got, err := client.DeleteSnapshot("my-backup")
	if err != nil {
		t.Errorf("Request returned an error: %s", err)
		return
	}

	expected := &SimpleResponse{Result: "success"}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %+v, got %+v", expected, got)
	}
}
