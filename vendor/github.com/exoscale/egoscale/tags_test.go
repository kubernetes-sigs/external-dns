package egoscale

import (
	"testing"
)

func TestCreateTags(t *testing.T) {
	req := &CreateTags{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestDeleteTags(t *testing.T) {
	req := &DeleteTags{}
	_ = req.response().(*AsyncJobResult)
	_ = req.asyncResponse().(*booleanResponse)
}

func TestListTags(t *testing.T) {
	req := &ListTags{}
	_ = req.response().(*ListTagsResponse)
}
