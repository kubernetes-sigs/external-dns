package linodego_test

import (
	"context"
	"strings"
	"testing"
)

func TestGetAccount(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	client, teardown := createTestClient(t, "fixtures/TestGetAccount")
	defer teardown()

	account, err := client.GetAccount(context.Background())
	if err != nil {
		t.Errorf("Error getting Account, expected struct, got error %v", err)
	}

	if !strings.Contains(account.Email, "@") {
		t.Error("Error accessing Account, expected Email to contain '@'")
	}
}
