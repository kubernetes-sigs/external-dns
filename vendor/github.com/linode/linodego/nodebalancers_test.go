package linodego_test

import (
	"context"
	"strings"
	"testing"

	"github.com/linode/linodego"
)

var (
	clientConnThrottle         = 20
	label                      = randString(12, lowerBytes, digits) + "-linodego-testing"
	testNodeBalancerCreateOpts = linodego.NodeBalancerCreateOptions{
		Label:              &label,
		Region:             "us-west",
		ClientConnThrottle: &clientConnThrottle,
	}
)

func TestCreateNodeBalancer(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	_, nodebalancer, teardown, err := setupNodeBalancer(t, "fixtures/TestCreateNodeBalancer")
	defer teardown()

	if err != nil {
		t.Errorf("Error creating nodebalancer: %v", err)
	}

	// when comparing fixtures to random value Label will differ, compare the known suffix
	if !strings.Contains(*nodebalancer.Label, "-linodego-testing") {
		t.Errorf("nodebalancer returned does not match nodebalancer create request")
	}
}

func TestUpdateNodeBalancer(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, teardown, err := setupNodeBalancer(t, "fixtures/TestUpdateNodeBalancer")
	defer teardown()

	renamedLabel := *nodebalancer.Label + "_r"
	updateOpts := linodego.NodeBalancerUpdateOptions{
		Label: &renamedLabel,
	}
	nodebalancer, err = client.UpdateNodeBalancer(context.Background(), nodebalancer.ID, updateOpts)

	if err != nil {
		t.Errorf("Error renaming nodebalancer, %s", err)
	}

	if !strings.Contains(*nodebalancer.Label, "-linodego-testing_r") {
		t.Errorf("nodebalancer returned does not match nodebalancer create request")
	}
}

func TestListNodeBalancers(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, _, teardown, err := setupNodeBalancer(t, "fixtures/TestListNodeBalancers")
	defer teardown()

	nodebalancers, err := client.ListNodeBalancers(context.Background(), nil)
	if err != nil {
		t.Errorf("Error listing nodebalancers, expected struct, got error %v", err)
	}
	if len(nodebalancers) == 0 {
		t.Errorf("Expected a list of nodebalancers, but got %v", nodebalancers)
	}
}

func TestGetNodeBalancer(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, teardown, err := setupNodeBalancer(t, "fixtures/TestGetNodeBalancer")
	defer teardown()

	_, err = client.GetNodeBalancer(context.Background(), nodebalancer.ID)
	if err != nil {
		t.Errorf("Error getting nodebalancer %d, expected *NodeBalancer, got error %v", nodebalancer.ID, err)
	}
}

func setupNodeBalancer(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.NodeBalancer, func(), error) {
	t.Helper()
	var fixtureTeardown func()
	client, fixtureTeardown := createTestClient(t, fixturesYaml)
	createOpts := testNodeBalancerCreateOpts
	nodebalancer, err := client.CreateNodeBalancer(context.Background(), createOpts)
	if err != nil {
		t.Errorf("Error listing nodebalancers, expected struct, got error %v", err)
	}

	teardown := func() {
		if err := client.DeleteNodeBalancer(context.Background(), nodebalancer.ID); err != nil {
			t.Errorf("Expected to delete a nodebalancer, but got %v", err)
		}
		fixtureTeardown()
	}
	return client, nodebalancer, teardown, err
}
