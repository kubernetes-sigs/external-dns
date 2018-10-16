package egoscale

import (
	"strings"
	"testing"
)

func TestSnapshot(t *testing.T) {
	instance := &Snapshot{}
	if instance.ResourceType() != "Snapshot" {
		t.Errorf("ResourceType doesn't match")
	}
}

func TestCreateSnapshot(t *testing.T) {
	req := &CreateSnapshot{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*Snapshot)
}

func TestListSnapshots(t *testing.T) {
	req := &ListSnapshots{}
	_ = req.response().(*ListSnapshotsResponse)
}

func TestDeleteSnapshot(t *testing.T) {
	req := &DeleteSnapshot{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestRevertSnapshot(t *testing.T) {
	req := &RevertSnapshot{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestSnapshotStateString(t *testing.T) {
	if BackingUp != SnapshotState(3) {
		t.Error("bad enum value", (int)(BackingUp), 3)
	}

	if BackingUp.String() != "BackingUp" {
		t.Error("mismatch", BackingUp, "BackingUp")
	}
	s := SnapshotState(45)

	if !strings.Contains(s.String(), "45") {
		t.Error("bad state", s.String())
	}
}
