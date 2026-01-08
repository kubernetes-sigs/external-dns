/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
