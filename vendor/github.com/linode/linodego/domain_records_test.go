package linodego_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/linode/linodego"
)

var (
	testDomainRecordCreateOpts = linodego.DomainRecordCreateOptions{
		Target: "127.0.0.1",
		Type:   linodego.RecordTypeA,
		Name:   "a",
	}
)

func TestCreateDomainRecord(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	_, _, record, teardown, err := setupDomainRecord(t, "fixtures/TestCreateDomainRecord")
	defer teardown()

	if err != nil {
		t.Errorf("Error creating domain record, got error %v", err)
	}

	expected := testDomainRecordCreateOpts

	// cant compare Target, fixture IPs are sanitized
	if record.Name != expected.Name || record.Type != expected.Type {
		t.Errorf("DomainRecord did not match CreateOptions")
	}
}

func TestUpdateDomainRecord(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, domain, record, teardown, err := setupDomainRecord(t, "fixtures/TestUpdateDomainRecord")
	defer teardown()

	updateOpts := linodego.DomainRecordUpdateOptions{
		Name: "renamed",
	}
	recordUpdated, err := client.UpdateDomainRecord(context.Background(), domain.ID, record.ID, updateOpts)

	if err != nil {
		t.Errorf("Error updating domain record, %s", err)
	}
	if recordUpdated.Name != "renamed" || record.Type != recordUpdated.Type || recordUpdated.Target != record.Target {
		t.Errorf("DomainRecord did not match UpdateOptions")
	}
}

func TestListDomainRecords(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, domain, record, teardown, err := setupDomainRecord(t, "fixtures/TestListDomainRecords")
	defer teardown()

	filter, err := json.Marshal(map[string]interface{}{
		"name": record.Name,
	})

	listOpts := linodego.NewListOptions(0, string(filter))
	records, err := client.ListDomainRecords(context.Background(), domain.ID, listOpts)
	if err != nil {
		t.Errorf("Error listing domains records, expected array, got error %v", err)
	}
	if len(records) != 1 {
		t.Errorf("Expected ListDomainRecords to match one result")
	}
}

func TestListDomainRecordsMultiplePages(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, domain, record, teardown, err := setupDomainRecord(t, "fixtures/TestListDomainRecordsMultiplePages")
	defer teardown()

	filter, err := json.Marshal(map[string]interface{}{
		"name": record.Name,
	})

	listOpts := linodego.NewListOptions(0, string(filter))
	records, err := client.ListDomainRecords(context.Background(), domain.ID, listOpts)
	if err != nil {
		t.Errorf("Error listing domains records, expected array, got error %v", err)
	}
	if len(records) != 2 {
		t.Errorf("Expected ListDomainRecords to match one result")
	}
}

func TestGetDomainRecord(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, domain, record, teardown, err := setupDomainRecord(t, "fixtures/TestGetDomainRecord")
	defer teardown()

	recordGot, err := client.GetDomainRecord(context.Background(), domain.ID, record.ID)
	if recordGot.Name != record.Name {
		t.Errorf("GetDomainRecord did not get the expected record")
	}
	if err != nil {
		t.Errorf("Error getting domain %d, got error %v", domain.ID, err)
	}
}

func setupDomainRecord(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.Domain, *linodego.DomainRecord, func(), error) {
	t.Helper()
	var fixtureTeardown func()
	client, domain, fixtureTeardown, err := setupDomain(t, fixturesYaml)
	if err != nil {
		t.Errorf("Error creating domain, got error %v", err)
	}

	createOpts := testDomainRecordCreateOpts
	record, err := client.CreateDomainRecord(context.Background(), domain.ID, createOpts)
	if err != nil {
		t.Errorf("Error creating domain record, got error %v", err)
	}

	teardown := func() {
		// delete the DomainRecord to excercise the code
		if err := client.DeleteDomainRecord(context.Background(), domain.ID, record.ID); err != nil {
			t.Errorf("Expected to delete a domain record, but got %v", err)
		}
		fixtureTeardown()
	}
	return client, domain, record, teardown, err
}
