package linodego_test

import (
	"context"
	"testing"

	"github.com/linode/linodego"
)

func TestCreateVolume(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestCreateVolume")
	defer teardown()

	createOpts := linodego.VolumeCreateOptions{
		Label:  "linodego-test-volume",
		Region: "us-west",
	}
	volume, err := client.CreateVolume(context.Background(), createOpts)
	if err != nil {
		t.Errorf("Error listing volumes, expected struct, got error %v", err)
	}
	if volume.ID == 0 {
		t.Errorf("Expected a volumes id, but got 0")
	}

	if err := client.DeleteVolume(context.Background(), volume.ID); err != nil {
		t.Errorf("Expected to delete a volume, but got %v", err)
	}
}

func TestRenameVolume(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, volume, teardown, err := setupVolume(t, "fixtures/TestRenameVolume")
	defer teardown()

	volume, err = client.RenameVolume(context.Background(), volume.ID, "test-volume-renamed")
	if err != nil {
		t.Errorf("Error renaming volume, %s", err)
	}
}

func TestResizeVolume(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, volume, teardown, err := setupVolume(t, "fixtures/TestResizeVolume")
	defer teardown()

	if err != nil {
		t.Errorf("Error setting up volume test, %s", err)
	}

	if ok, err := client.ResizeVolume(context.Background(), volume.ID, volume.Size+1); err != nil {
		t.Errorf("Error resizing volume, %s", err)
	} else if !ok {
		t.Errorf("Error resizing volume")
	}
}

func TestListVolumes(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestListVolumes")
	defer teardown()

	volumes, err := client.ListVolumes(context.Background(), nil)
	if err != nil {
		t.Errorf("Error listing volumes, expected struct, got error %v", err)
	}
	if len(volumes) == 0 {
		t.Errorf("Expected a list of volumes, but got %v", volumes)
	}
}

func TestGetVolume(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, teardown := createTestClient(t, "fixtures/TestGetVolume")
	defer teardown()

	_, err := client.GetVolume(context.Background(), TestVolumeID)
	if err != nil {
		t.Errorf("Error getting volume %d, expected *LinodeVolume, got error %v", TestVolumeID, err)
	}
}

func TestWaitForVolumeLinodeID_nil(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}
	client, volume, teardown, err := setupVolume(t, "fixtures/TestWaitForVolumeLinodeID_nil")
	defer teardown()

	if err != nil {
		t.Errorf("Error setting up volume test, %s", err)
	}
	_, err = client.WaitForVolumeLinodeID(context.Background(), volume.ID, nil, 3)

	if err != nil {
		t.Errorf("Error getting volume %d, expected *LinodeVolume, got error %v", TestVolumeID, err)
	}
}

func TestWaitForVolumeLinodeID(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}

	client, instance, teardownInstance, err := setupInstance(t, "fixtures/TestWaitForVolumeLinodeID_linode")
	if err != nil {
		t.Errorf("Error setting up instance for volume test, %s", err)
	}

	defer teardownInstance()

	createConfigOpts := linodego.InstanceConfigCreateOptions{
		Label:   "test-instance-volume",
		Devices: linodego.InstanceConfigDeviceMap{},
	}
	config, err := client.CreateInstanceConfig(context.Background(), instance.ID, createConfigOpts)
	if err != nil {
		t.Errorf("Error setting up instance config for volume test, %s", err)
	}

	client, volume, teardownVolume, err := setupVolume(t, "fixtures/TestWaitForVolumeLinodeID")
	if err != nil {
		t.Errorf("Error setting up volume test, %s", err)
	}
	defer teardownVolume()

	attachOptions := linodego.VolumeAttachOptions{LinodeID: instance.ID, ConfigID: config.ID}
	if volumeAttached, err := client.AttachVolume(context.Background(), volume.ID, &attachOptions); err != nil {
		t.Errorf("Error attaching volume, %s", err)
	} else if volumeAttached.LinodeID == nil {
		t.Errorf("Could not attach test volume to test instance")
	}

	_, err = client.WaitForVolumeLinodeID(context.Background(), volume.ID, nil, 3)
	if err == nil {
		t.Errorf("Expected to timeout waiting for nil LinodeID on volume %d : %s", volume.ID, err)
	}

	_, err = client.WaitForVolumeLinodeID(context.Background(), volume.ID, &instance.ID, 3)
	if err != nil {
		t.Errorf("Error waiting for volume %d to attach to instance %d: %s", volume.ID, instance.ID, err)
	}
}

func setupVolume(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.Volume, func(), error) {
	t.Helper()
	var fixtureTeardown func()
	client, fixtureTeardown := createTestClient(t, fixturesYaml)
	createOpts := linodego.VolumeCreateOptions{
		Label:  "linodego-test-volume",
		Region: "us-west",
	}
	volume, err := client.CreateVolume(context.Background(), createOpts)
	if err != nil {
		t.Errorf("Error listing volumes, expected struct, got error %v", err)
	}

	teardown := func() {
		if err := client.DeleteVolume(context.Background(), volume.ID); err != nil {
			t.Errorf("Expected to delete a volume, but got %v", err)
		}
		fixtureTeardown()
	}
	return client, volume, teardown, err
}
