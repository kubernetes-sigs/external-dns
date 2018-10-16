package linodego_test

import (
	"context"
	"testing"

	"github.com/linode/linodego"
)

var (
	testNodeAddress                = "192.168.128.1"
	testNodePort                   = "8080"
	testNodeLabel                  = "test-label"
	testNodeWeight                 = 10
	testNodeBalancerNodeCreateOpts = linodego.NodeBalancerNodeCreateOptions{
		Address: testNodeAddress + ":" + testNodePort,
		Label:   testNodeLabel,
		Weight:  testNodeWeight,
		Mode:    linodego.ModeAccept,
	}
)

func TestCreateNodeBalancerNode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	_, _, _, node, teardown, err := setupNodeBalancerNode(t, "fixtures/TestCreateNodeBalancerNode")
	defer teardown()

	if err != nil {
		t.Errorf("Error creating NodeBalancer Node, got error %v", err)
	}

	expected := testNodeBalancerNodeCreateOpts

	// fixture sanitization breaks predictability for this test, verify the prefix
	if node.Address[:7] != expected.Address[:7] ||
		node.Label != expected.Label ||
		node.Weight != expected.Weight ||
		node.Mode != expected.Mode {
		t.Errorf("NodeBalancerNode did not match CreateOptions")
	}
}

func TestUpdateNodeBalancerNode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, config, node, teardown, err := setupNodeBalancerNode(t, "fixtures/TestUpdateNodeBalancerNode")
	defer teardown()

	updateOpts := linodego.NodeBalancerNodeUpdateOptions{
		Address: testNodeAddress + ":" + ("1" + testNodePort),
		Mode:    linodego.ModeDrain,
		Weight:  testNodeWeight + 90,
		Label:   testNodeLabel + "_r",
	}
	nodeUpdated, err := client.UpdateNodeBalancerNode(context.Background(), nodebalancer.ID, config.ID, node.ID, updateOpts)

	if err != nil {
		t.Errorf("Error updating NodeBalancer Node, %s", err)
	}

	// fixture sanitization breaks predictability for this test, verify the prefix
	if updateOpts.Address[:7] != nodeUpdated.Address[:7] ||
		string(updateOpts.Mode) != string(nodeUpdated.Mode) ||
		updateOpts.Label != nodeUpdated.Label ||
		updateOpts.Weight != nodeUpdated.Weight {
		t.Errorf("NodeBalancerNode did not match UpdateOptions")
	}
}

func TestListNodeBalancerNodes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, config, _, teardown, err := setupNodeBalancerNode(t, "fixtures/TestListNodeBalancerNodes")
	defer teardown()

	listOpts := linodego.NewListOptions(0, "")
	nodes, err := client.ListNodeBalancerNodes(context.Background(), nodebalancer.ID, config.ID, listOpts)
	if err != nil {
		t.Errorf("Error listing nodebalancers nodes, expected array, got error %v", err)
	}
	if len(nodes) != listOpts.Results {
		t.Errorf("Expected ListNodeBalancerNodes to match API result count")
	}
}

func TestListNodeBalancerNodesMultiplePages(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	// This fixture was hand-crafted to render an empty page 1 result, with a single result on page 2
	// "results:1,data:[],page:1,pages:2"  .. "results:1,data[{...}],page:2,pages:2"
	client, nodebalancer, config, _, teardown, err := setupNodeBalancerNode(t, "fixtures/TestListNodeBalancerNodesMultiplePages")
	defer teardown()

	listOpts := linodego.NewListOptions(0, "")
	nodes, err := client.ListNodeBalancerNodes(context.Background(), nodebalancer.ID, config.ID, listOpts)
	if err != nil {
		t.Errorf("Error listing nodebalancers configs, expected array, got error %v", err)
	}
	if len(nodes) != listOpts.Results {
		t.Errorf("Expected ListNodeBalancerNodes count to match API results count")
	}
}

func TestGetNodeBalancerNode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, config, node, teardown, err := setupNodeBalancerNode(t, "fixtures/TestGetNodeBalancerNode")
	defer teardown()

	nodeGot, err := client.GetNodeBalancerNode(context.Background(), nodebalancer.ID, config.ID, node.ID)
	if nodeGot.Address != node.Address {
		t.Errorf("GetNodeBalancerNode did not get the expected node")
	}
	if err != nil {
		t.Errorf("Error getting nodebalancer %d, got error %v", nodebalancer.ID, err)
	}
}

func setupNodeBalancerNode(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.NodeBalancer, *linodego.NodeBalancerConfig, *linodego.NodeBalancerNode, func(), error) {
	t.Helper()
	var fixtureTeardown func()
	client, nodebalancer, config, fixtureTeardown, err := setupNodeBalancerConfig(t, fixturesYaml)
	if err != nil {
		t.Errorf("Error creating nodebalancer config, got error %v", err)
	}

	createOpts := testNodeBalancerNodeCreateOpts
	node, err := client.CreateNodeBalancerNode(context.Background(), nodebalancer.ID, config.ID, createOpts)
	if err != nil {
		t.Errorf("Error creating NodeBalancer Config Node, got error %v", err)
	}

	teardown := func() {
		// delete the NodeBalancerNode to excercise the code
		if err := client.DeleteNodeBalancerNode(context.Background(), nodebalancer.ID, config.ID, node.ID); err != nil {
			t.Errorf("Expected to delete a NodeBalancer Config Node, but got %v", err)
		}
		fixtureTeardown()
	}
	return client, nodebalancer, config, node, teardown, err
}
