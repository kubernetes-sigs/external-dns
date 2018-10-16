// +build acceptance blockstorage

package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumetypes"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestVolumeTypesList(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	listOpts := volumetypes.ListOpts{
		Sort:  "name:asc",
		Limit: 1,
	}

	allPages, err := volumetypes.List(client, listOpts).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve volumetypes: %v", err)
	}

	allVolumeTypes, err := volumetypes.ExtractVolumeTypes(allPages)
	if err != nil {
		t.Fatalf("Unable to extract volumetypes: %v", err)
	}

	for _, vt := range allVolumeTypes {
		tools.PrintResource(t, vt)
	}

	if len(allVolumeTypes) > 0 {
		vt, err := volumetypes.Get(client, allVolumeTypes[0].ID).Extract()
		if err != nil {
			t.Fatalf("Error retrieving volume type: %v", err)
		}

		tools.PrintResource(t, vt)
	}
}

func TestVolumeTypesCRUD(t *testing.T) {
	client, err := clients.NewBlockStorageV3Client()
	if err != nil {
		t.Fatalf("Unable to create a blockstorage client: %v", err)
	}

	createOpts := volumetypes.CreateOpts{
		Name:        "create_from_gophercloud",
		ExtraSpecs:  map[string]string{"volume_backend_name": "fake_backend_name"},
		Description: "create_from_gophercloud",
	}

	vt, err := volumetypes.Create(client, createOpts).Extract()
	if err != nil {
		t.Fatalf("Unable to create volumetype: %v", err)
	}

	th.AssertEquals(t, true, vt.IsPublic)

	tools.PrintResource(t, vt)

	defer volumetypes.Delete(client, vt.ID)

	var isPublic = false

	newVT, err := volumetypes.Update(client, vt.ID, volumetypes.UpdateOpts{
		Name:     "updated_volume_type",
		IsPublic: &isPublic,
	}).Extract()

	th.AssertEquals(t, "updated_volume_type", newVT.Name)
	th.AssertEquals(t, false, newVT.IsPublic)

	tools.PrintResource(t, newVT)
}
