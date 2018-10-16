package linodego_test

import (
	"context"
	"testing"

	"github.com/linode/linodego"
)

func TestListStackscripts(t *testing.T) {
	client, teardown := createTestClient(t, "fixtures/TestListStackscripts")
	defer teardown()

	filterOpt := linodego.NewListOptions(1, "")
	stackscripts, err := client.ListStackscripts(context.Background(), filterOpt)
	if err != nil {
		t.Errorf("Error listing stackscripts, expected struct - error %v", err)
	}
	if len(stackscripts) == 0 {
		t.Errorf("Expected a list of public stackscripts - %v", stackscripts)
	}
}
