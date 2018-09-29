// +build acceptance blockstorage

package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/pagination"
)

func TestVolumesList(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	volume1, err := CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}
	defer DeleteVolume(t, client, volume1)

	volume2, err := CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}
	defer DeleteVolume(t, client, volume2)

	pages := 0
	err = volumes.List(client, volumes.ListOpts{Limit: 1}).EachPage(func(page pagination.Page) (bool, error) {
		pages++

		actual, err := volumes.ExtractVolumes(page)
		if err != nil {
			t.Fatalf("Unable to extract volumes: %v", err)
		}

		if len(actual) != 1 {
			t.Fatalf("Expected 1 volume, got %d", len(actual))
		}

		tools.PrintResource(t, actual[0])

		return true, nil
	})

	if pages != 2 {
		t.Fatalf("Expected 2 pages, saw %d", pages)
	}
}

func TestVolumesCreateDelete(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create blockstorage client: %v", err)
	}

	volume, err := CreateVolume(t, client)
	if err != nil {
		t.Fatalf("Unable to create volume: %v", err)
	}
	defer DeleteVolume(t, client, volume)

	newVolume, err := volumes.Get(client, volume.ID).Extract()
	if err != nil {
		t.Errorf("Unable to retrieve volume: %v", err)
	}

	tools.PrintResource(t, newVolume)
}
