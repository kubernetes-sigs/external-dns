package cloudflare

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubmitRulesetChanges_Disabled(t *testing.T) {
	p := &CloudFlareProvider{
		RulesetsConfig: RulesetsConfig{
			Enabled: false,
		},
	}
	ctx := context.Background()
	// Should return true (success) when disabled
	success := p.submitRulesetChanges(ctx, "zone-1", []*cloudFlareChange{})
	assert.True(t, success)
}

func TestSubmitRulesetChanges_Enabled(t *testing.T) {
	p := &CloudFlareProvider{
		RulesetsConfig: RulesetsConfig{
			Enabled: true,
		},
	}
	ctx := context.Background()
	// Placeholder behavior: currently just logs and returns true
	success := p.submitRulesetChanges(ctx, "zone-1", []*cloudFlareChange{})
	assert.True(t, success)
}
