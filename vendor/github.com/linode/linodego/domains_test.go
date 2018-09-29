package linodego_test

import (
	"context"
	"testing"

	"github.com/linode/linodego"
)

var (
	testDomainCreateOpts = linodego.DomainCreateOptions{
		Domain:   randLabel() + "-linodego-testing.com",
		Type:     linodego.DomainTypeMaster,
		SOAEmail: "example@example.com",
	}
)

func TestCreateDomain(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	_, domain, teardown, err := setupDomain(t, "fixtures/TestCreateDomain")
	defer teardown()

	if err != nil {
		t.Errorf("Error creating domain: %v", err)
	}

	// when comparing fixtures to random value Domain will differ
	if domain.SOAEmail != testDomainCreateOpts.SOAEmail {
		t.Errorf("Domain returned does not match domain create request")
	}
}

func TestUpdateDomain(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, domain, teardown, err := setupDomain(t, "fixtures/TestUpdateDomain")
	defer teardown()

	updateOpts := linodego.DomainUpdateOptions{
		Domain: "linodego-renamed-domain.com",
	}
	domain, err = client.UpdateDomain(context.Background(), domain.ID, updateOpts)
	if err != nil {
		t.Errorf("Error renaming domain, %s", err)
	}
}

func TestListDomains(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, _, teardown, err := setupDomain(t, "fixtures/TestListDomains")
	defer teardown()

	domains, err := client.ListDomains(context.Background(), nil)
	if err != nil {
		t.Errorf("Error listing domains, expected struct, got error %v", err)
	}
	if len(domains) == 0 {
		t.Errorf("Expected a list of domains, but got %v", domains)
	}
}

func TestGetDomain(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, domain, teardown, err := setupDomain(t, "fixtures/TestGetDomain")
	defer teardown()

	_, err = client.GetDomain(context.Background(), domain.ID)
	if err != nil {
		t.Errorf("Error getting domain %d, expected *Domain, got error %v", domain.ID, err)
	}
}

func setupDomain(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.Domain, func(), error) {
	t.Helper()
	var fixtureTeardown func()
	client, fixtureTeardown := createTestClient(t, fixturesYaml)
	createOpts := testDomainCreateOpts
	domain, err := client.CreateDomain(context.Background(), createOpts)
	if err != nil {
		t.Errorf("Error listing domains, expected struct, got error %v", err)
	}

	teardown := func() {
		if err := client.DeleteDomain(context.Background(), domain.ID); err != nil {
			t.Errorf("Expected to delete a domain, but got %v", err)
		}
		fixtureTeardown()
	}
	return client, domain, teardown, err
}
