/*
Copyright 2022 The Kubernetes Authors.

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

package tencentcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/tencentcloud/cloudapi"
)

const (
	TencentCloudEmptyPrefix = "@"
	DefaultAPIRate          = 9
)

func NewTencentCloudProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, configFile string, zoneType string, dryRun bool) (*TencentCloudProvider, error) {
	cfg := tencentCloudConfig{}
	if configFile != "" {
		contents, err := os.ReadFile(configFile)
		if err != nil {
			return nil, fmt.Errorf("failed to read Tencent Cloud config file '%s': %w", configFile, err)
		}
		err = json.Unmarshal(contents, &cfg)
		if err != nil {
			return nil, fmt.Errorf("failed to parse Tencent Cloud config file '%s': %w", configFile, err)
		}
	}

	var apiService cloudapi.TencentAPIService = cloudapi.NewTencentAPIService(cfg.RegionId, DefaultAPIRate, cfg.SecretId, cfg.SecretKey, cfg.InternetEndpoint)
	if dryRun {
		apiService = cloudapi.NewReadOnlyAPIService(cfg.RegionId, DefaultAPIRate, cfg.SecretId, cfg.SecretKey, cfg.InternetEndpoint)
	}

	tencentCloudProvider := &TencentCloudProvider{
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		apiService:   apiService,
		vpcID:        cfg.VPCId,
		privateZone:  zoneType == "private",
	}

	return tencentCloudProvider, nil
}

type TencentCloudProvider struct {
	provider.BaseProvider
	logger       *log.Logger
	apiService   cloudapi.TencentAPIService
	domainFilter endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter // Private Zone only
	vpcID        string                // Private Zone only
	privateZone  bool
}

type tencentCloudConfig struct {
	RegionId         string `json:"regionId" yaml:"regionId"`
	SecretId         string `json:"secretId" yaml:"secretId"`
	SecretKey        string `json:"secretKey" yaml:"secretKey"`
	VPCId            string `json:"vpcId" yaml:"vpcId"`
	InternetEndpoint bool   `json:"internetEndpoint" yaml:"internetEndpoint"`
}

func (p *TencentCloudProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	if p.privateZone {
		return p.privateZoneRecords()
	}
	return p.dnsRecords()
}

func (p *TencentCloudProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if !changes.HasChanges() {
		return nil
	}

	log.Infof("apply changes. %s", cloudapi.JsonWrapper(changes))

	if p.privateZone {
		return p.applyChangesForPrivateZone(changes)
	}
	return p.applyChangesForDNS(changes)
}

func getSubDomain(domain string, endpoint *endpoint.Endpoint) string {
	name := endpoint.DNSName
	name = name[:len(name)-len(domain)]
	name = strings.TrimSuffix(name, ".")

	if name == "" {
		return TencentCloudEmptyPrefix
	}
	return name
}

func getDnsDomain(subDomain string, domain string) string {
	if subDomain == TencentCloudEmptyPrefix {
		return domain
	}
	return subDomain + "." + domain
}
