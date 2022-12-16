/*
Copyright 2020 The Kubernetes Authors.

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

package adguardhome

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	adguard "github.com/markussiebert/go-adguardhome-client/client"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// The regex hat should identify and parse all custom rules in adguard home that are used to handle dns records
var regex = regexp.MustCompile(`(?P<dnsName>[a-zA-Z\*.-]+)\$dnstype=(?P<recordType>[A-Z]+),dnsrewrite=NOERROR;[A-Z]+;(?P<target>.*)`)

// AdguardHomeProvider is an implementation of Provider for Vultr DNS.
type AdguardHomeProvider struct {
	provider.BaseProvider
	client adguard.ClientWithResponses

	domainFilter endpoint.DomainFilter
	DryRun       bool
}

// Prepare basicAuth with credentials
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}

// NewAdguardHomeProvider initializes a new Vultr BNS based provider
func NewAdguardHomeProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*AdguardHomeProvider, error) {
	adguardHomeURL, adguardHomeUrlOk := os.LookupEnv("ADGUARD_HOME_URL")
	if !adguardHomeUrlOk {
		return nil, fmt.Errorf("no url was found in environment variable ADGUARD_HOME_URL")
	}

	adguardHomeUser, adguardHomeUserOk := os.LookupEnv("ADGUARD_HOME_USER")
	if !adguardHomeUserOk {
		return nil, fmt.Errorf("no user was found in environment variable ADGUARD_HOME_USER")
	}

	adguardHomePass, adguardHomePassOk := os.LookupEnv("ADGUARD_HOME_PASS")
	if !adguardHomePassOk {
		return nil, fmt.Errorf("no password was found in environment variable ADGUARD_HOME_PASS")
	}

	client, clientOk := adguard.NewClientWithResponses(adguardHomeURL, adguard.ClientOption(func(client *adguard.AdguardHomeClient) error {
		client.RequestEditors = append(client.RequestEditors,
			adguard.RequestEditorFn(
				func(ctx context.Context, req *http.Request) error {
					req.Header.Add("Authorization", basicAuth(adguardHomeUser, adguardHomePass))
					return nil
				},
			),
			adguard.RequestEditorFn(
				func(ctx context.Context, req *http.Request) error {
					req.Header.Add("User-Agent", "ExternalDNS")
					return nil
				},
			),
		)
		return nil
	}))

	if clientOk != nil {
		return nil, fmt.Errorf("failed to create the adguard home api client")
	}

	p := &AdguardHomeProvider{
		client:       *client,
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}

	return p, nil
}

// ApplyChanges implements Provider, syncing desired state with the AdguardHome server Local DNS.
func (p *AdguardHomeProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Get current filter values
	resp, err := p.client.FilteringStatusWithResponse(ctx)
	if err != nil {
		log.Errorf("Error %s", err)
		return err
	}

	debugJ, _ := json.MarshalIndent(changes, "", "    ")
	log.Debugf("changes, %s", string(debugJ))
	userRules := *resp.JSON200.UserRules

	for _, createValue := range append(changes.Create, changes.UpdateNew...) {
		newRule := genAdguardCustomRuleRecord(createValue)
		log.Debugf("add custom rule %s", newRule)
		userRules = append(userRules, newRule)
	}

	for _, deleteValue := range append(changes.UpdateOld, changes.Delete...) {
		userRules = removeRule(userRules, genAdguardCustomRuleRecord(deleteValue), false)
	}

	apply, err := p.client.FilteringSetRulesWithResponse(ctx, adguard.SetRulesRequest{
		Rules: &userRules,
	})

	if apply.StatusCode() != 200 {
		log.Errorf("Failed to sync records %s", apply.Status())
	}
	return err
}

func removeRule(rules []string, ruleToRemove string, finishAfterFirst bool) []string {
	userRules := rules
	for userRuleIndex, userRule := range rules {
		if userRule == ruleToRemove {
			log.Debugf("deleted rule %s", ruleToRemove)
			userRules = append(userRules[:userRuleIndex], userRules[userRuleIndex+1:]...)
			if finishAfterFirst {
				return userRules
			}
		}
	}
	return userRules
}

// Records implements Provider, populating a slice of endpoints from
// AdguardHome local DNS.
func (p *AdguardHomeProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	resp, err := p.client.FilteringStatusWithResponse(ctx)
	if err != nil {
		log.Errorf("Error %s", err)
	}

	log.WithFields(log.Fields{
		"func":  "Records",
		"rules": resp.JSON200.UserRules,
	}).Debugf("retrieved rules")
	var ret []*endpoint.Endpoint
	for _, customRule := range *resp.JSON200.UserRules {
		parsed := parseAdguardCustomRuleRecord(customRule)
		if parsed != nil {
			ret = append(ret, parsed)
		}
	}

	return ret, nil
}

func genAdguardCustomRuleRecord(e *endpoint.Endpoint) string {
	return fmt.Sprintf("%s$dnstype=%s,dnsrewrite=NOERROR;%s;%s", e.DNSName, e.RecordType, e.RecordType, e.Targets.String())
}

func parseAdguardCustomRuleRecord(s string) *endpoint.Endpoint {

	matched := regex.FindStringSubmatch(s)

	if matched == nil {
		return nil
	}

	log.Debugf("successfully parsed rule %s", s)

	endpoint := &endpoint.Endpoint{
		DNSName:    matched[1],
		RecordType: matched[2],
		Targets:    endpoint.Targets{matched[3]},
	}
	debugJ, _ := json.MarshalIndent(endpoint, "", "    ")
	log.Debugf("returning endpoint %s", string(debugJ))
	return endpoint
}
