package egoscale

import (
	"testing"
)

func TestListInstanceGroups(t *testing.T) {
	req := &ListInstanceGroups{}
	_ = req.response().(*ListInstanceGroupsResponse)
}

func TestCreateInstanceGroup(t *testing.T) {
	req := &CreateInstanceGroup{}
	_ = req.response().(*InstanceGroup)
}

func TestUpdateInstanceGroup(t *testing.T) {
	req := &UpdateInstanceGroup{}
	_ = req.response().(*InstanceGroup)
}

func TestDeleteInstanceGroup(t *testing.T) {
	req := &DeleteInstanceGroup{}
	_ = req.response().(*booleanResponse)
}
