package cloudflare

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubmitRulesetChanges_Disabled(t *testing.T) {
	mockClient := NewMockCloudFlareClient()
	p := &CloudFlareProvider{
		Client: mockClient,
		RulesetsConfig: RulesetsConfig{
			Enabled: false,
		},
	}
	ctx := context.Background()
	// Should return true (success) when disabled
	success := p.submitRulesetChanges(ctx, "zone-1", []*cloudFlareChange{
		{
			Ruleset: `{"id": "test-id", "kind": "zone"}`,
		},
	})
	assert.True(t, success)
	assert.Len(t, mockClient.Actions, 0)
}

func TestSubmitRulesetChanges_Enabled(t *testing.T) {
	mockClient := NewMockCloudFlareClient()
	p := &CloudFlareProvider{
		Client: mockClient,
		RulesetsConfig: RulesetsConfig{
			Enabled: true,
		},
	}
	ctx := context.Background()

	changes := []*cloudFlareChange{
		{
			Ruleset: `{"id": "test-ruleset-id", "kind": "zone", "phase": "http_request_firewall_custom", "rules": []}`,
		},
	}

	success := p.submitRulesetChanges(ctx, "zone-1", changes)
	assert.True(t, success)

	// Verify action
	if assert.Len(t, mockClient.Actions, 1) {
		assert.Equal(t, "UpdateRuleset", mockClient.Actions[0].Name)
		assert.Equal(t, "test-ruleset-id", mockClient.Actions[0].RecordId)
	}
}

func TestSubmitRulesetChanges_InvalidJSON(t *testing.T) {
	mockClient := NewMockCloudFlareClient()
	p := &CloudFlareProvider{
		Client: mockClient,
		RulesetsConfig: RulesetsConfig{
			Enabled: true,
		},
	}
	ctx := context.Background()

	changes := []*cloudFlareChange{
		{
			Ruleset: `{"id": "test-ids", "kind": `, // Invalid JSON
		},
	}

	success := p.submitRulesetChanges(ctx, "zone-1", changes)
	assert.False(t, success) // Should return false -> failed
	assert.Len(t, mockClient.Actions, 0)
}

func TestSubmitRulesetChanges_MissingID(t *testing.T) {
	mockClient := NewMockCloudFlareClient()
	p := &CloudFlareProvider{
		Client: mockClient,
		RulesetsConfig: RulesetsConfig{
			Enabled: true,
		},
	}
	ctx := context.Background()

	changes := []*cloudFlareChange{
		{
			Ruleset: `{"kind": "zone"}`, // Missing ID
		},
	}

	success := p.submitRulesetChanges(ctx, "zone-1", changes)
	assert.False(t, success) // Should return false -> failed
	assert.Len(t, mockClient.Actions, 0)
}
