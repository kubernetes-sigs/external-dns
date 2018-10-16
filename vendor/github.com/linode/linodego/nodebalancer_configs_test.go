package linodego_test

import (
	"context"
	"testing"

	"github.com/linode/linodego"
)

var (
	testNodeBalancerConfigCreateOpts = linodego.NodeBalancerConfigCreateOptions{
		Port:      80,
		Protocol:  linodego.ProtocolHTTP,
		Algorithm: linodego.AlgorithmRoundRobin,
	}
)

func TestCreateNodeBalancerConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	_, _, config, teardown, err := setupNodeBalancerConfig(t, "fixtures/TestCreateNodeBalancerConfig")
	defer teardown()

	if err != nil {
		t.Errorf("Error creating NodeBalancer Config, got error %v", err)
	}

	expected := testNodeBalancerConfigCreateOpts

	// cant compare Target, fixture IPs are sanitized
	if config.Port != expected.Port || config.Protocol != expected.Protocol {
		t.Errorf("NodeBalancerConfig did not match CreateOptions")
	}
}

func TestUpdateNodeBalancerConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, config, teardown, err := setupNodeBalancerConfig(t, "fixtures/TestUpdateNodeBalancerConfig")
	defer teardown()

	updateOpts := linodego.NodeBalancerConfigUpdateOptions{
		Port:      8080,
		Protocol:  linodego.ProtocolTCP,
		Algorithm: linodego.AlgorithmLeastConn,
	}
	configUpdated, err := client.UpdateNodeBalancerConfig(context.Background(), nodebalancer.ID, config.ID, updateOpts)

	if err != nil {
		t.Errorf("Error updating NodeBalancer Config, %s", err)
	}
	if configUpdated.Port != updateOpts.Port ||
		string(updateOpts.Algorithm) != string(configUpdated.Algorithm) ||
		string(updateOpts.Algorithm) != string(configUpdated.Algorithm) {
		t.Errorf("NodeBalancerConfig did not match UpdateOptions")
	}
}

func TestListNodeBalancerConfigs(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, _, teardown, err := setupNodeBalancerConfig(t, "fixtures/TestListNodeBalancerConfigs")
	defer teardown()

	listOpts := linodego.NewListOptions(0, "")
	configs, err := client.ListNodeBalancerConfigs(context.Background(), nodebalancer.ID, listOpts)
	if err != nil {
		t.Errorf("Error listing nodebalancers configs, expected array, got error %v", err)
	}
	if len(configs) != listOpts.Results {
		t.Errorf("Expected ListNodeBalancerConfigs to match API result count")
	}
}

func TestListNodeBalancerConfigsMultiplePages(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	// This fixture was hand-crafted to render an empty page 1 result, with a single result on page 2
	// "results:1,data:[],page:1,pages:2"  .. "results:1,data[{...}],page:2,pages:2"
	client, nodebalancer, _, teardown, err := setupNodeBalancerConfig(t, "fixtures/TestListNodeBalancerConfigsMultiplePages")
	defer teardown()

	listOpts := linodego.NewListOptions(0, "")
	configs, err := client.ListNodeBalancerConfigs(context.Background(), nodebalancer.ID, listOpts)
	if err != nil {
		t.Errorf("Error listing nodebalancers configs, expected array, got error %v", err)
	}
	if len(configs) != listOpts.Results {
		t.Errorf("Expected ListNodeBalancerConfigs count to match API results count")
	}
}

func TestGetNodeBalancerConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, nodebalancer, config, teardown, err := setupNodeBalancerConfig(t, "fixtures/TestGetNodeBalancerConfig")
	defer teardown()

	configGot, err := client.GetNodeBalancerConfig(context.Background(), nodebalancer.ID, config.ID)
	if configGot.Port != config.Port {
		t.Errorf("GetNodeBalancerConfig did not get the expected config")
	}
	if err != nil {
		t.Errorf("Error getting nodebalancer %d, got error %v", nodebalancer.ID, err)
	}
}

func setupNodeBalancerConfig(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.NodeBalancer, *linodego.NodeBalancerConfig, func(), error) {
	t.Helper()
	var fixtureTeardown func()
	client, nodebalancer, fixtureTeardown, err := setupNodeBalancer(t, fixturesYaml)
	if err != nil {
		t.Errorf("Error creating nodebalancer, got error %v", err)
	}

	createOpts := testNodeBalancerConfigCreateOpts
	config, err := client.CreateNodeBalancerConfig(context.Background(), nodebalancer.ID, createOpts)
	if err != nil {
		t.Errorf("Error creating NodeBalancer Config, got error %v", err)
	}

	teardown := func() {
		// delete the NodeBalancerConfig to excercise the code
		if err := client.DeleteNodeBalancerConfig(context.Background(), nodebalancer.ID, config.ID); err != nil {
			t.Errorf("Expected to delete a NodeBalancer Config, but got %v", err)
		}
		fixtureTeardown()
	}
	return client, nodebalancer, config, teardown, err
}
