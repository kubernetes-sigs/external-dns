package cloudflare

import (
	"context"
	"encoding/json"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/rulesets"
	log "github.com/sirupsen/logrus"
)

// RulesetsConfig contains configuration for Ruleset management
type RulesetsConfig struct {
	Enabled bool
}

// submitRulesetChanges handles the creation/update/deletion of Rulesets
// For this MVP, we will only support creating/updating rulesets based on the annotation value.
// The value is expected to be a JSON string or a reference.
func (p *CloudFlareProvider) submitRulesetChanges(ctx context.Context, zoneID string, changes []*cloudFlareChange) bool {
	if !p.RulesetsConfig.Enabled {
		return true
	}

	var failed bool

	for _, change := range changes {
		if change.Ruleset == "" {
			continue
		}

		// We assume the annotation value is a JSON representing the Ruleset.
		// We expect an object that can be unmarshaled into a map
		// PLUS an "id" field to identify which ruleset to update.
		var rulesetDef map[string]interface{}

		err := json.Unmarshal([]byte(change.Ruleset), &rulesetDef)
		if err != nil {
			log.Errorf("Failed to unmarshal ruleset JSON for zone %s: %v", zoneID, err)
			failed = true
			continue
		}

		idVal, ok := rulesetDef["id"]
		if !ok {
			log.Errorf("Ruleset JSON must contain an 'id' field for zone %s", zoneID)
			failed = true
			continue
		}

		id, ok := idVal.(string)
		if !ok || id == "" {
			log.Errorf("Ruleset JSON 'id' field must be a non-empty string for zone %s", zoneID)
			failed = true
			continue
		}

		log.Infof("Updating Ruleset %s for zone %s", id, zoneID)

		// Use WithRequestBody to pass the raw JSON (as map) directly, avoiding complex struct mapping
		// for the union types in Rules.
		_, err = p.Client.UpdateRuleset(ctx, id, rulesets.RulesetUpdateParams{
			ZoneID: cloudflare.F(zoneID),
		}, option.WithRequestBody("application/json", rulesetDef))

		if err != nil {
			log.Errorf("Failed to update ruleset %s: %v", id, err)
			failed = true
		}
	}

	return !failed
}
