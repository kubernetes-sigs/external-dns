package linodego_test

import (
	"context"
	"testing"
)

func TestListRegions(t *testing.T) {
	client, teardown := createTestClient(t, "fixtures/TestListRegions")
	defer teardown()

	regions, err := client.ListRegions(context.Background(), nil)
	if err != nil {
		t.Errorf("Error listing regions, expected struct - error %v", err)
	}
	if len(regions) == 0 {
		t.Errorf("Expected a list of regions - %v", regions)
	}
}
