package linodego_test

import (
	"context"

	. "github.com/linode/linodego"

	"testing"
)

func TestGetImage_missing(t *testing.T) {
	client, teardown := createTestClient(t, "fixtures/TestGetImage_missing")
	defer teardown()

	i, err := client.GetImage(context.Background(), "does-not-exist")
	if err == nil {
		t.Errorf("should have received an error requesting a missing image, got %v", i)
	}
	e, ok := err.(*Error)
	if !ok {
		t.Errorf("should have received an Error requesting a missing image, got %v", e)
	}

	if e.Code != 404 {
		t.Errorf("should have received a 404 Code requesting a missing image, got %v", e.Code)
	}
}

func TestGetImage_found(t *testing.T) {
	client, teardown := createTestClient(t, "fixtures/TestGetImage_found")
	defer teardown()

	i, err := client.GetImage(context.Background(), "linode/ubuntu16.04lts")
	if err != nil {
		t.Errorf("Error getting image, expected struct, got %v and error %v", i, err)
	}
	if i.ID != "linode/ubuntu16.04lts" {
		t.Errorf("Expected a specific image, but got a different one %v", i)
	}
}
func TestListImages(t *testing.T) {
	client, teardown := createTestClient(t, "fixtures/TestListImages")
	defer teardown()

	i, err := client.ListImages(context.Background(), nil)
	if err != nil {
		t.Errorf("Error listing images, expected struct, got error %v", err)
	}
	if len(i) == 0 {
		t.Errorf("Expected a list of images, but got none %v", i)
	}
}
