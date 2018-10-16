package linodego_test

import (
	"context"
	"testing"

	"github.com/linode/linodego"
)

var testSnapshotLabel = "snapshot-linodego-testing"

func TestListInstanceBackups(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	client, instance, backup, teardown, err := setupInstanceBackup(t, "fixtures/TestListInstanceBackups")
	defer teardown()

	backupGotten, err := client.GetInstanceSnapshot(context.Background(), instance.ID, backup.ID)
	if err != nil {
		t.Errorf("Error getting backup: %v", err)
	} else if backupGotten.Label != backup.Label {
		t.Errorf("Error getting backup, Labels dont match")
	}

	backups, err := client.GetInstanceBackups(context.Background(), instance.ID)
	if err != nil {
		t.Errorf("Error getting backups: %v", err)
	}

	if backups.Snapshot.InProgress == nil && backups.Snapshot.Current == nil {
		t.Errorf("Error getting snapshot: No Current or InProgress Snapshot")
	}

	if backups.Snapshot.InProgress != nil && backups.Snapshot.InProgress.Label != testSnapshotLabel {
		t.Errorf("Expected snapshot did not match inprogress snapshot: %v", backups.Snapshot.InProgress)
	} else if backups.Snapshot.Current != nil && backups.Snapshot.Current.Label != testSnapshotLabel {
		t.Errorf("Expected snapshot did not match current snapshot: %v", backups.Snapshot.Current)
	}

	_, err = client.WaitForSnapshotStatus(context.Background(), instance.ID, backup.ID, linodego.SnapshotSuccessful, 180)
	if err != nil {
		t.Errorf("Error waiting for snapshot: %v", err)
	}

	restoreOpts := linodego.RestoreInstanceOptions{
		LinodeID:  instance.ID,
		Overwrite: true,
	}

	ok, err := client.RestoreInstanceBackup(context.Background(), instance.ID, backup.ID, restoreOpts)
	if err != nil {
		t.Errorf("Error restoring backup: %v", err)
	} else if !ok {
		t.Errorf("Error restoring backup.")
	}

	ok, err = client.CancelInstanceBackups(context.Background(), instance.ID)
	if err != nil {
		t.Errorf("Error cancelling backups: %v", err)
	} else if !ok {
		t.Errorf("Error cancelling backups.")
	}
}

func setupInstanceBackup(t *testing.T, fixturesYaml string) (*linodego.Client, *linodego.Instance, *linodego.InstanceSnapshot, func(), error) {
	t.Helper()
	client, instance, fixtureTeardown, err := setupInstanceWithoutDisks(t, fixturesYaml)
	if err != nil {
		t.Errorf("Error creating instance, got error %v", err)
	}

	client.WaitForInstanceStatus(context.Background(), instance.ID, linodego.InstanceOffline, 180)
	createOpts := linodego.InstanceDiskCreateOptions{
		Size:       1,
		Label:      "snapshot-linodego-testing",
		Filesystem: "ext4",
	}
	disk, err := client.CreateInstanceDisk(context.Background(), instance.ID, createOpts)
	if err != nil {
		t.Errorf("Error creating Instance Disk: %v", err)
	}

	// wait for disk to finish provisioning
	event, err := client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionDiskCreate, disk.Created, 240)
	if err != nil {
		t.Errorf("Error waiting for instance snapshot: %v", err)
	}

	if event.Status == linodego.EventFailed {
		t.Errorf("Error creating instance disk: Disk Create Failed")
	}

	ok, err := client.EnableInstanceBackups(context.Background(), instance.ID)
	if err != nil {
		t.Errorf("Error enabling Instance Backups: %v", err)
	} else if !ok {
		t.Errorf("Error enabling Instance Backups.")
	}

	snapshot, err := client.CreateInstanceSnapshot(context.Background(), instance.ID, testSnapshotLabel)
	if err != nil {
		t.Errorf("Error creating instance snapshot: %v", err)
	}

	event, err = client.WaitForEventFinished(context.Background(), instance.ID, linodego.EntityLinode, linodego.ActionLinodeSnapshot, *instance.Created, 240)
	if err != nil {
		t.Errorf("Error waiting for instance snapshot: %v", err)
	}
	if event.Status == linodego.EventFailed {
		t.Errorf("Error taking instance snapshot: Snapshot Failed")
	}

	return client, instance, snapshot, fixtureTeardown, err
}
