package egoscale

import (
	"net/url"
	"testing"
)

func TestCreateAffinityGroup(t *testing.T) {
	req := &CreateAffinityGroup{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*AffinityGroup)
}

func TestDeleteAffinityGroup(t *testing.T) {
	req := &DeleteAffinityGroup{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestListAffinityGroups(t *testing.T) {
	req := &ListAffinityGroups{}
	_ = req.response().(*ListAffinityGroupsResponse)
}

func TestListAffinityGroupTypes(t *testing.T) {
	req := &ListAffinityGroupTypes{}
	_ = req.response().(*ListAffinityGroupTypesResponse)
}

func TestUpdateVMAffinityGroup(t *testing.T) {
	req := &UpdateVMAffinityGroup{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*VirtualMachine)
}

func TestUpdateVMOnBeforeSend(t *testing.T) {
	req := &UpdateVMAffinityGroup{}
	params := url.Values{}

	if err := req.onBeforeSend(params); err != nil {
		t.Error(err)
	}

	if _, ok := params["affinitygroupids"]; !ok {
		t.Errorf("affinitygroupids should have been set")
	}
}
