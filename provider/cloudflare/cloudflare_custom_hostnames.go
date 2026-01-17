/*
Copyright 2026 The Kubernetes Authors.

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
	"errors"
	"fmt"
	"maps"
	"slices"

	"github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v5/option"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/provider"
)

// customHostname represents a Cloudflare custom hostname (v5 API compatible wrapper)
type customHostname struct {
	id                 string
	hostname           string
	customOriginServer string
	customOriginSNI    string
	ssl                *customHostnameSSL
}

// customHostnameSSL represents SSL configuration for custom hostname
type customHostnameSSL struct {
	sslType              string
	method               string
	bundleMethod         string
	certificateAuthority string
	settings             customHostnameSSLSettings
}

// customHostnameSSLSettings represents SSL settings for custom hostname
type customHostnameSSLSettings struct {
	minTLSVersion string
}

// for faster getCustomHostname() lookup
type customHostnameIndex struct {
	hostname string
}

type customHostnamesMap map[customHostnameIndex]customHostname

type CustomHostnamesConfig struct {
	Enabled              bool
	MinTLSVersion        string
	CertificateAuthority string
}

var recordTypeCustomHostnameSupported = map[string]bool{
	"A":     true,
	"CNAME": true,
}

func (z zoneService) CustomHostnames(ctx context.Context, zoneID string) autoPager[custom_hostnames.CustomHostnameListResponse] {
	params := custom_hostnames.CustomHostnameListParams{
		ZoneID: cloudflare.F(zoneID),
	}
	return z.service.CustomHostnames.ListAutoPaging(ctx, params)
}

func (z zoneService) DeleteCustomHostname(ctx context.Context, customHostnameID string, params custom_hostnames.CustomHostnameDeleteParams) error {
	_, err := z.service.CustomHostnames.Delete(ctx, customHostnameID, params)
	return err
}

func (z zoneService) CreateCustomHostname(ctx context.Context, zoneID string, ch customHostname) error {
	params := buildCustomHostnameNewParams(zoneID, ch)
	_, err := z.service.CustomHostnames.New(ctx, params,
		option.WithJSONSet("custom_origin_server", ch.customOriginServer))
	return err
}

// buildCustomHostnameNewParams builds the params for creating a custom hostname
func buildCustomHostnameNewParams(zoneID string, ch customHostname) custom_hostnames.CustomHostnameNewParams {
	params := custom_hostnames.CustomHostnameNewParams{
		ZoneID:   cloudflare.F(zoneID),
		Hostname: cloudflare.F(ch.hostname),
	}
	if ch.ssl != nil {
		sslParams := custom_hostnames.CustomHostnameNewParamsSSL{}
		if ch.ssl.method != "" {
			sslParams.Method = cloudflare.F(custom_hostnames.DCVMethod(ch.ssl.method))
		}
		if ch.ssl.sslType != "" {
			sslParams.Type = cloudflare.F(custom_hostnames.DomainValidationType(ch.ssl.sslType))
		}
		if ch.ssl.bundleMethod != "" {
			sslParams.BundleMethod = cloudflare.F(custom_hostnames.BundleMethod(ch.ssl.bundleMethod))
		}
		if ch.ssl.certificateAuthority != "" && ch.ssl.certificateAuthority != "none" {
			sslParams.CertificateAuthority = cloudflare.F(cloudflare.CertificateCA(ch.ssl.certificateAuthority))
		}
		if ch.ssl.settings.minTLSVersion != "" {
			sslParams.Settings = cloudflare.F(custom_hostnames.CustomHostnameNewParamsSSLSettings{
				MinTLSVersion: cloudflare.F(custom_hostnames.CustomHostnameNewParamsSSLSettingsMinTLSVersion(ch.ssl.settings.minTLSVersion)),
			})
		}
		params.SSL = cloudflare.F(sslParams)
	}
	return params
}

// submitCustomHostnameChanges implements Custom Hostname functionality for the Change, returns false if it fails
func (p *CloudFlareProvider) submitCustomHostnameChanges(ctx context.Context, zoneID string, change *cloudFlareChange, chs customHostnamesMap, logFields log.Fields) bool {
	// return early if disabled
	if !p.CustomHostnamesConfig.Enabled {
		return true
	}

	switch change.Action {
	case cloudFlareUpdate:
		return p.processCustomHostnameUpdate(ctx, zoneID, change, chs, logFields)
	case cloudFlareDelete:
		return p.processCustomHostnameDelete(ctx, zoneID, change, chs, logFields)
	case cloudFlareCreate:
		return p.processCustomHostnameCreate(ctx, zoneID, change, chs, logFields)
	}

	return true
}

func (p *CloudFlareProvider) processCustomHostnameUpdate(ctx context.Context, zoneID string, change *cloudFlareChange, chs customHostnamesMap, logFields log.Fields) bool {
	if !recordTypeCustomHostnameSupported[string(change.ResourceRecord.Type)] {
		return true
	}
	failedChange := false
	add, remove, _ := provider.Difference(change.CustomHostnamesPrev, slices.Collect(maps.Keys(change.CustomHostnames)))

	for _, changeCH := range remove {
		if prevCh, err := getCustomHostname(chs, changeCH); err == nil {
			prevChID := prevCh.id
			if prevChID != "" {
				log.WithFields(logFields).Infof("Removing previous custom hostname %q/%q", prevChID, changeCH)
				params := custom_hostnames.CustomHostnameDeleteParams{ZoneID: cloudflare.F(zoneID)}
				chErr := p.Client.DeleteCustomHostname(ctx, prevChID, params)
				if chErr != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to remove previous custom hostname %q/%q: %v", prevChID, changeCH, chErr)
				}
			}
		}
	}
	for _, changeCH := range add {
		log.WithFields(logFields).Infof("Adding custom hostname %q", changeCH)
		chErr := p.Client.CreateCustomHostname(ctx, zoneID, change.CustomHostnames[changeCH])
		if chErr != nil {
			failedChange = true
			log.WithFields(logFields).Errorf("failed to add custom hostname %q: %v", changeCH, chErr)
		}
	}
	return !failedChange
}

func (p *CloudFlareProvider) processCustomHostnameDelete(ctx context.Context, zoneID string, change *cloudFlareChange, chs customHostnamesMap, logFields log.Fields) bool {
	failedChange := false
	for _, changeCH := range change.CustomHostnames {
		if recordTypeCustomHostnameSupported[string(change.ResourceRecord.Type)] && changeCH.hostname != "" {
			log.WithFields(logFields).Infof("Deleting custom hostname %q", changeCH.hostname)
			if ch, err := getCustomHostname(chs, changeCH.hostname); err == nil {
				chID := ch.id
				params := custom_hostnames.CustomHostnameDeleteParams{ZoneID: cloudflare.F(zoneID)}
				chErr := p.Client.DeleteCustomHostname(ctx, chID, params)
				if chErr != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to delete custom hostname %q/%q: %v", chID, changeCH.hostname, chErr)
				}
			} else {
				log.WithFields(logFields).Warnf("failed to delete custom hostname %q: %v", changeCH.hostname, err)
			}
		}
	}
	return !failedChange
}

func (p *CloudFlareProvider) processCustomHostnameCreate(ctx context.Context, zoneID string, change *cloudFlareChange, chs customHostnamesMap, logFields log.Fields) bool {
	failedChange := false
	for _, changeCH := range change.CustomHostnames {
		if recordTypeCustomHostnameSupported[string(change.ResourceRecord.Type)] && changeCH.hostname != "" {
			log.WithFields(logFields).Infof("Creating custom hostname %q", changeCH.hostname)
			if ch, err := getCustomHostname(chs, changeCH.hostname); err == nil {
				if changeCH.customOriginServer == ch.customOriginServer {
					log.WithFields(logFields).Warnf("custom hostname %q already exists with the same origin %q, continue", changeCH.hostname, ch.customOriginServer)
				} else {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to create custom hostname, %q already exists with origin %q", changeCH.hostname, ch.customOriginServer)
				}
			} else {
				chErr := p.Client.CreateCustomHostname(ctx, zoneID, changeCH)
				if chErr != nil {
					failedChange = true
					log.WithFields(logFields).Errorf("failed to create custom hostname %q: %v", changeCH.hostname, chErr)
				}
			}
		}
	}
	return !failedChange
}

func getCustomHostname(chs customHostnamesMap, chName string) (customHostname, error) {
	if chName == "" {
		return customHostname{}, fmt.Errorf("failed to get custom hostname: %q is empty", chName)
	}
	if ch, ok := chs[customHostnameIndex{hostname: chName}]; ok {
		return ch, nil
	}
	return customHostname{}, fmt.Errorf("failed to get custom hostname: %q not found", chName)
}

func (p *CloudFlareProvider) newCustomHostname(hostname string, origin string) customHostname {
	return customHostname{
		hostname:           hostname,
		customOriginServer: origin,
		ssl:                getCustomHostnamesSSLOptions(p.CustomHostnamesConfig),
	}
}

func getCustomHostnamesSSLOptions(customHostnamesConfig CustomHostnamesConfig) *customHostnameSSL {
	ssl := &customHostnameSSL{
		sslType:      "dv",
		method:       "http",
		bundleMethod: "ubiquitous",
		settings: customHostnameSSLSettings{
			minTLSVersion: customHostnamesConfig.MinTLSVersion,
		},
	}
	// Set CertificateAuthority if provided
	// We're not able to set it at all (even with a blank) if you're not on an enterprise plan
	if customHostnamesConfig.CertificateAuthority != "none" {
		ssl.certificateAuthority = customHostnamesConfig.CertificateAuthority
	}
	return ssl
}

func newCustomHostnameIndex(ch customHostname) customHostnameIndex {
	return customHostnameIndex{hostname: ch.hostname}
}

// listCustomHostnamesWithPagination performs automatic pagination of results on requests to cloudflare.CustomHostnames
func (p *CloudFlareProvider) listCustomHostnamesWithPagination(ctx context.Context, zoneID string) (customHostnamesMap, error) {
	if !p.CustomHostnamesConfig.Enabled {
		return nil, nil
	}
	chs := make(customHostnamesMap)
	iter := p.Client.CustomHostnames(ctx, zoneID)
	customHostnames, err := listAllCustomHostnames(iter)
	if err != nil {
		convertedError := convertCloudflareError(err)
		if !errors.Is(convertedError, provider.SoftError) {
			log.Errorf("zone %q failed to fetch custom hostnames. Please check if \"Cloudflare for SaaS\" is enabled and API key permissions, %v", zoneID, err)
		}
		return nil, convertedError
	}
	for _, ch := range customHostnames {
		chs[newCustomHostnameIndex(ch)] = ch
	}
	return chs, nil
}

// listAllCustomHostnames extracts all custom hostnames from the iterator
func listAllCustomHostnames(iter autoPager[custom_hostnames.CustomHostnameListResponse]) ([]customHostname, error) {
	var customHostnames []customHostname
	for ch := range autoPagerIterator(iter) {
		customHostnames = append(customHostnames, customHostname{
			id:                 ch.ID,
			hostname:           ch.Hostname,
			customOriginServer: ch.CustomOriginServer,
			customOriginSNI:    ch.CustomOriginSNI,
		})
	}
	if iter.Err() != nil {
		return nil, iter.Err()
	}
	return customHostnames, nil
}
