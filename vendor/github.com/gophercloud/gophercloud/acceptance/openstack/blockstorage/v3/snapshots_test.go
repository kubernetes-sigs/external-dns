// +build acceptance blockstorage

package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/snapshots"
	"github.com/gophercloud/gophercloud/pagination"
)

func TestSnapshotsList(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	volume1, err := CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}

	defer DeleteVolume(t, client, volume1)

	snapshot1, err := CreateSnapshot(t, client, volume1)
	if err != nil {
		t.Fatalf("Unable to create snapshot: %v", err)
	}

	defer DeleteSnapshot(t, client, snapshot1)

	volume2, err := CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}

	defer DeleteVolume(t, client, volume2)

	snapshot2, err := CreateSnapshot(t, client, volume2)
	if err != nil {
		t.Fatalf("Unable to create snapshot: %v", err)
	}

	defer DeleteSnapshot(t, client, snapshot2)

	pages := 0
	err = snapshots.List(client, snapshots.ListOpts{Limit: 1}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := snapshots.ExtractSnapshots(page)
		if err != nil {
			t.Fatalf("Unable to extract snapshots: %v", err)
		}

		if len(actual) != 1 {
			t.Fatalf("Expected 1 snapshot, got %d", len(actual))
		}

		tools.PrintResource(t, actual[0])

		return true, nil
	})

	if pages != 2 {
		t.Fatalf("Expected 2 pages, saw %d", pages)
	}
}

func TestSnapshotsCreateDelete(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	volume, err := CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}
	defer DeleteVolume(t, client, volume)

	snapshot, err := CreateSnapshot(t, client, volume)
	if err != nil {
		t.Fatalf("Unable to create snapshot: %v", err)
	}
	defer DeleteSnapshot(t, client, snapshot)

	newSnapshot, err := snapshots.Get(client, snapshot.ID).Extract()
	if err != nil {
		t.Errorf("Unable to retrieve snapshot: %v", err)
	}

	tools.PrintResource(t, newSnapshot)
}
