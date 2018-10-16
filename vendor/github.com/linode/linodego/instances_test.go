package linodego_test

import (
	"context"
	"testing"

	"github.com/linode/linodego"
)

func TestListInstances(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestListInstances")
	defer teardown()

	linodes, err := client.ListInstances(context.Background(), nil)
	if err != nil {
		t.Errorf("Error listing instances, expected struct, got error %v", err)
	}
	if len(linodes) == 0 {
		t.Errorf("Expected a list of instances, but got %v", linodes)
	}
}

func TestGetInstance(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestGetInstance")
	defer teardown()

	instance, err := client.GetInstance(context.Background(), TestInstanceID)
	if err != nil {
		t.Errorf("Error getting instance TestInstanceID, expected *LinodeInstance, got error %v", err)
	}
	if instance.Specs.Disk <= 0 {
		t.Errorf("Error in instance TestInstanceID spec for disk size, %v", instance.Specs)
	}
}

func TestListInstanceDisks(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestListInstanceDisks")
	defer teardown()

	disks, err := client.ListInstanceDisks(context.Background(), TestInstanceID, nil)
	if err != nil {
		t.Errorf("Error listing instance disks, expected struct, got error %v", err)
	}
	if len(disks) == 0 {
		t.Errorf("Expected a list of instance disks, but got %v", disks)
	}
}

func TestListInstanceConfigs(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestListInstanceConfigs")
	defer teardown()

	configs, err := client.ListInstanceConfigs(context.Background(), TestInstanceID, nil)
	if err != nil {
		t.Errorf("Error listing instance configs, expected struct, got error %v", err)
	}
	if len(configs) == 0 {
		t.Errorf("Expected a list of instance configs, but got %v", configs)
	}
}

func TestListInstanceVolumes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestListInstanceVolumes")
	defer teardown()

	volumes, err := client.ListInstanceVolumes(context.Background(), TestInstanceID, nil)
	if err != nil {
		t.Errorf("Error listing instance volumes, expected struct, got error %v", err)
	}
	if len(volumes) == 0 {
		t.Errorf("Expected an list of instance volumes, but got %v", volumes)
	}
}

func setupInstance(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.Instance, func(), error) {
	t.Helper()
	client, fixtureTeardown := createTestClient(t, fixturesYaml)
	createOpts := linodego.InstanceCreateOptions{
		Label:    "linodego-test-instance",
		RootPass: "R34lBAdP455",
		Region:   "us-west",
		Type:     "g6-nanode-1",
	}
	instance, err := client.CreateInstance(context.Background(), createOpts)
	if err != nil {
		t.Errorf("Error creating test Instance: %s", err)
	}

	teardown := func() {
		if err := client.DeleteInstance(context.Background(), instance.ID); err != nil {
			t.Errorf("Error deleting test Instance: %s", err)
		}
		fixtureTeardown()
	}
	return client, instance, teardown, err
}

func setupInstanceWithoutDisks(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.Instance, func(), error) {
	t.Helper()
	client, fixtureTeardown := createTestClient(t, fixturesYaml)
	falseBool := false
	createOpts := linodego.InstanceCreateOptions{
		Label:  "linodego-test-instance",
		Region: "us-west",
		Type:   "g6-nanode-1",
		Booted: &falseBool,
	}
	instance, err := client.CreateInstance(context.Background(), createOpts)
	if err != nil {
		t.Errorf("Error creating test Instance: %s", err)
		return nil, nil, fixtureTeardown, err
	}
	configOpts := linodego.InstanceConfigCreateOptions{
		Label: "linodego-test-config",
	}
	_, err = client.CreateInstanceConfig(context.Background(), instance.ID, configOpts)
	if err != nil {
		t.Errorf("Error creating config: %s", err)
		return nil, nil, fixtureTeardown, err
	}

	teardown := func() {
		if err := client.DeleteInstance(context.Background(), instance.ID); err != nil {
			t.Errorf("Error deleting test Instance: %s", err)
		}
		fixtureTeardown()
	}
	return client, instance, teardown, err
}
