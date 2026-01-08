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
	"encoding/json"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/rulesets"
	log "github.com/sirupsen/logrus"
)

// submitRulesetChanges handles the creation/update/deletion of Rulesets
// For this MVP, we will only support creating/updating rulesets based on the annotation value.
// The value is expected to be a JSON string or a reference.
func (p *CloudFlareProvider) submitRulesetChanges(ctx context.Context, zoneID string, changes []*cloudFlareChange) bool {
	var failed bool

	for _, change := range changes {
		if change.Ruleset == "" {
			continue
		}

		// We assume the annotation value is a JSON representing the Ruleset.
		// We expect an object that can be unmarshaled into a map
		// PLUS an "id" field to identify which ruleset to update.
		// AND optionally an "action" field (e.g. "delete").
		var rulesetDef map[string]any

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

		// Check for deletion action
		if action, ok := rulesetDef["action"].(string); ok && action == "delete" {
			log.Infof("Deleting Ruleset %s for zone %s", id, zoneID)
			err = p.Client.DeleteRuleset(ctx, zoneID, id)
			if err != nil {
				log.Errorf("Failed to delete ruleset %s: %v", id, err)
				failed = true
			}
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
