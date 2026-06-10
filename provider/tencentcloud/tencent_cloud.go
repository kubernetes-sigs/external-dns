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
	"slices"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	dnspod "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dnspod/v20210323"
	privatedns "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/privatedns/v20201028"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	defaultDnsTTL = 600
	defaultPvtTTL = 60
	pageSize      = 50
	apexRecord    = "@"
	dnsDomain     = "dnspod.tencentcloudapi.com"     // Public  DNS for Internet
	zoneDomain    = "privatedns.tencentcloudapi.com" // Private DNS for VPC
)

type PublicDNSAPI interface {
	CreateRecord(req *dnspod.CreateRecordRequest) (res *dnspod.CreateRecordResponse, err error)
	CreateRecordBatch(req *dnspod.CreateRecordBatchRequest) (res *dnspod.CreateRecordBatchResponse, err error)
	DeleteRecord(req *dnspod.DeleteRecordRequest) (res *dnspod.DeleteRecordResponse, err error)
	DeleteRecordBatch(req *dnspod.DeleteRecordBatchRequest) (res *dnspod.DeleteRecordBatchResponse, err error)
	DescribeDomainList(req *dnspod.DescribeDomainListRequest) (res *dnspod.DescribeDomainListResponse, err error)
	DescribeRecordList(req *dnspod.DescribeRecordListRequest) (res *dnspod.DescribeRecordListResponse, err error)
}

type PrivateDNSAPI interface {
	CreatePrivateZoneRecord(req *privatedns.CreatePrivateZoneRecordRequest) (res *privatedns.CreatePrivateZoneRecordResponse, err error)
	CreatePrivateZoneRecordList(req *privatedns.CreatePrivateZoneRecordListRequest) (res *privatedns.CreatePrivateZoneRecordListResponse, err error)
	DeletePrivateZoneRecord(req *privatedns.DeletePrivateZoneRecordRequest) (res *privatedns.DeletePrivateZoneRecordResponse, err error)
	DescribePrivateZoneList(req *privatedns.DescribePrivateZoneListRequest) (res *privatedns.DescribePrivateZoneListResponse, err error)
	DescribePrivateZoneRecordList(req *privatedns.DescribePrivateZoneRecordListRequest) (res *privatedns.DescribePrivateZoneRecordListResponse, err error)
}

type TencentCloudProvider struct {
	provider.BaseProvider
	domainFilter *endpoint.DomainFilter
	zoneIDFilter *provider.ZoneIDFilter // Private Zone only
	vpcID        string                 // Private Zone only
	dryRun       bool
	privateZone  bool
	dnsApi       PublicDNSAPI
	pvtApi       PrivateDNSAPI
}

type tencentCloudConfig struct {
	RegionId  string `json:"regionId"    yaml:"regionId"`
	SecretId  string `json:"secretId"    yaml:"secretId"`
	SecretKey string `json:"secretKey"   yaml:"secretKey"`
	VPCId     string `json:"vpcId"       yaml:"vpcId"`
}

type dnsRecordGroup struct {
	DomainId   uint64
	Domain     string
	SubDomain  string
	RecordType string
	RecordList []*dnspod.RecordListItem
}

type zoneRecordGroup struct {
	ZoneId     string
	Domain     string
	SubDomain  string
	RecordType string
	RecordList []*privatedns.PrivateZoneRecord
}

type extendEndpoint struct {
	*endpoint.Endpoint
	ZoneId string
	Domain string
}

// New creates an Tencent Cloud provider from the given configuration.
func New(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	return newProvider(cfg.TencentCloudConfigFile, domainFilter, provider.NewZoneIDFilter(cfg.ZoneIDFilter), cfg.TencentCloudZoneType, cfg.DryRun)
}

// newProvider creates a new Tencent Cloud provider.
//
// Returns the provider or an error if a provider could not be created.
func newProvider(configFile string, domainFilter *endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, zoneType string, dryRun bool) (*TencentCloudProvider, error) {
	cfg := tencentCloudConfig{}
	if configFile != "" {
		contents, err := os.ReadFile(configFile)
		if err != nil {
			return nil, fmt.Errorf("Failed to read TencentCloud config file '%s': %w", configFile, err)
		}
		err = json.Unmarshal(contents, &cfg)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse TencentCloud config file '%s': %w", configFile, err)
		}
	}

	cred := common.NewCredential(cfg.SecretId, cfg.SecretKey)

	// Public DNS service
	dnsProfile := profile.NewClientProfile()
	dnsProfile.HttpProfile.Endpoint = dnsDomain
	dnsApi, err := dnspod.NewClient(cred, cfg.RegionId, dnsProfile)

	if err != nil {
		return nil, fmt.Errorf("Failed to create TencentCloud DNSPod client: %w", err)
	}

	// Private DNS service
	zoneProfile := profile.NewClientProfile()
	zoneProfile.HttpProfile.Endpoint = zoneDomain
	zoneApi, err := privatedns.NewClient(cred, cfg.RegionId, zoneProfile)

	if err != nil {
		return nil, fmt.Errorf("Failed to create TencentCloud PrivateDNS client: %w", err)
	}

	provider := &TencentCloudProvider{
		domainFilter: domainFilter,
		zoneIDFilter: &zoneIDFilter,
		dryRun:       dryRun,
		dnsApi:       dnsApi,
		pvtApi:       zoneApi,
		vpcID:        cfg.VPCId,
		privateZone:  zoneType == "private",
	}

	return provider, nil
}

func (p *TencentCloudProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	if p.privateZone {
		return p.zoneRecords()
	}
	return p.dnsRecords()
}

func (p *TencentCloudProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if !changes.HasChanges() {
		return nil
	}

	if p.privateZone {
		return p.applyChangesForZone(changes)
	}

	return p.applyChangesForDNS(changes)
}

func (p *TencentCloudProvider) dnsRecords() ([]*endpoint.Endpoint, error) {
	log.Infof("Retrieving Tencent Cloud domain DNS records")
	domains, err := p.getDomainList()
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0)

	for _, domain := range domains {
		domainName := *domain.Name
		records, err := p.getDomainRecordList(domainName)
		if err != nil {
			return nil, err
		}

		endpointMap := make(map[string]*endpoint.Endpoint)
		for _, record := range records {
			if !provider.SupportedRecordType(*record.Type) {
				continue
			}

			name := getDNSName(*record.Name, domainName)
			key := toRecordKey(*record.Type, name, nil)
			target := wrapWithQuotes(*record.Type, *record.Value)

			if _, exist := endpointMap[key]; !exist {
				recordType := *record.Type
				ttl := endpoint.TTL(*record.TTL)
				// recordID := strconv.FormatUint(*record.RecordId, 10)
				endpointMap[key] = endpoint.NewEndpointWithTTL(name, recordType, ttl, target)
			} else {
				endpointMap[key].Targets = append(endpointMap[key].Targets, target)
			}
		}

		for _, ep := range endpointMap {
			endpoints = append(endpoints, ep)
		}
	}
	log.Infof("Found %d TencentCloud domain DNS record(s).", len(endpoints))

	return endpoints, nil
}

func (p *TencentCloudProvider) getDomainList() ([]*dnspod.DomainListItem, error) {
	request := dnspod.NewDescribeDomainListRequest()
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(pageSize)

	domainList := make([]*dnspod.DomainListItem, 0)
	totalCount := int64(pageSize)
	for *request.Offset < totalCount {
		response, err := p.dnsApi.DescribeDomainList(request)
		if err != nil {
			return nil, err
		}

		for _, domain := range response.Response.DomainList {
			if p.domainFilter == nil || !p.domainFilter.IsConfigured() || p.domainFilter.Match(*domain.Name) {
				domainList = append(domainList, domain)
			}
		}

		if response.Response.DomainCountInfo == nil || response.Response.DomainCountInfo.AllTotal == nil {
			break
		}
		totalCount = int64(*response.Response.DomainCountInfo.AllTotal)
		request.Offset = common.Int64Ptr(*request.Offset + int64(len(response.Response.DomainList)))
	}
	return domainList, nil
}

func (p *TencentCloudProvider) getDomainRecordMap(domain string, domainId uint64) (map[string]*dnsRecordGroup, error) {
	records, err := p.getDomainRecordList(domain)
	if err != nil {
		return nil, err
	}

	recordGroupMap := make(map[string]*dnsRecordGroup)
	for _, record := range records {
		key := toRecordKey(*record.Type, domain, record.Name)

		if _, exists := recordGroupMap[key]; !exists {
			recordGroupMap[key] = &dnsRecordGroup{
				DomainId:   domainId,
				Domain:     domain,
				SubDomain:  *record.Name,
				RecordType: *record.Type,
				RecordList: make([]*dnspod.RecordListItem, 0),
			}
		}
		recordGroupMap[key].RecordList = append(recordGroupMap[key].RecordList, record)
	}
	return recordGroupMap, nil
}

func (p *TencentCloudProvider) getDomainRecordList(domain string) ([]*dnspod.RecordListItem, error) {
	request := dnspod.NewDescribeRecordListRequest()
	request.Domain = &domain
	request.Offset = common.Uint64Ptr(0)
	request.Limit = common.Uint64Ptr(pageSize)

	recordList := make([]*dnspod.RecordListItem, 0)
	totalCount := uint64(pageSize)
	for *request.Offset < totalCount {
		response, err := p.dnsApi.DescribeRecordList(request)
		if err != nil {
			return nil, err
		}

		for _, record := range response.Response.RecordList {
			if *record.Name == apexRecord && *record.Type == endpoint.RecordTypeNS {
				continue
			}
			recordList = append(recordList, record)
		}

		if response.Response.RecordCountInfo == nil || response.Response.RecordCountInfo.TotalCount == nil {
			break
		}
		totalCount = *response.Response.RecordCountInfo.TotalCount
		request.Offset = common.Uint64Ptr(*request.Offset + uint64(len(response.Response.RecordList)))
	}
	return recordList, nil
}

func (p *TencentCloudProvider) applyChangesForDNS(changes *plan.Changes) error {
	log.Infof("Apply changes to TencentCloud domain DNS: %++v", *changes)
	domains, err := p.getDomainList()
	if err != nil {
		return err
	}

	zoneIDMapper := provider.ZoneIDName{}
	recordGroupMap := make(map[string]*dnsRecordGroup)
	for _, domain := range domains {
		zoneIDMapper.Add(strconv.FormatUint(*domain.DomainId, 10), *domain.Name)
		if recordGroups, err := p.getDomainRecordMap(*domain.Name, *domain.DomainId); err == nil {
			for key, group := range recordGroups {
				recordGroupMap[key] = group
			}
		}
	}

	deleteRecords := make([]*uint64, 0)
	for _, change := range [][]*endpoint.Endpoint{changes.Delete, changes.UpdateOld} {
		for _, ep := range change {
			key := toEndpointKey(ep)
			if group, exists := recordGroupMap[key]; exists {
				for _, record := range group.RecordList {
					target := wrapWithQuotes(*record.Type, *record.Value)
					if slices.Contains(ep.Targets, target) {
						deleteRecords = append(deleteRecords, record.RecordId)
					}
				}
			}
		}
	}

	createEndpoints := make(map[string][]*extendEndpoint)
	for _, change := range [][]*endpoint.Endpoint{changes.Create, changes.UpdateNew} {
		for _, ep := range change {
			zoneId, zoneName := zoneIDMapper.FindZone(ep.DNSName)
			if zoneId != "" {
				createEndpoints[zoneId] = append(
					createEndpoints[zoneId],
					&extendEndpoint{
						Endpoint: ep,
						ZoneId:   zoneId,
						Domain:   zoneName,
					},
				)
			}
		}
	}

	var errors []error
	if len(deleteRecords) > 0 {
		if err := p.deleteDNSRecords(deleteRecords); err != nil {
			errors = append(errors, err)
		}
	}

	if len(createEndpoints) > 0 {
		for zoneId, endpoints := range createEndpoints {
			domainId, _ := strconv.ParseUint(zoneId, 10, 64)
			for _, ep := range endpoints {
				if err := p.createRecord(domainId, ep.Domain, ep.Endpoint); err != nil {
					errors = append(errors, err...)
				}
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("Failed to apply changes to TencentCloud domain DNS: %v", errors)
	}

	return nil
}

func (p *TencentCloudProvider) createRecord(domainId uint64, domain string, ep *endpoint.Endpoint) []error {
	var errs []error
	if domainId == 0 || domain == "" || ep == nil {
		return append(errs, fmt.Errorf("Invalid input for creating DNS record"))
	}

	if p.dryRun {
		log.Infof("Dry run: create TencentCloud domain DNS record %s with endpoint %v", domain, ep)
		return nil
	}

	subDomain := getSubDomain(domain, ep)
	ttl := uint64(defaultDnsTTL)
	if ep.RecordTTL.IsConfigured() {
		ttl = uint64(ep.RecordTTL)
	}

	request := dnspod.NewCreateRecordRequest()
	request.Domain = &domain
	request.DomainId = &domainId
	request.SubDomain = &subDomain
	request.RecordType = &ep.RecordType
	request.RecordLine = common.StringPtr("默认")
	request.TTL = &ttl

	for _, target := range ep.Targets {
		target := unwrapQuotes(ep.RecordType, target)
		req := *request
		req.Value = &target

		if _, err := p.dnsApi.CreateRecord(&req); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (p *TencentCloudProvider) deleteDNSRecords(recordIds []*uint64) error {
	request := dnspod.NewDeleteRecordBatchRequest()
	request.RecordIdList = recordIds

	if p.dryRun {
		log.Infof("Dry run: delete TencentCloud domain DNS records %v", recordIds)
		return nil
	}

	if _, err := p.dnsApi.DeleteRecordBatch(request); err != nil {
		return err
	}

	return nil
}

func (p *TencentCloudProvider) zoneRecords() ([]*endpoint.Endpoint, error) {
	log.Infof("Retrieving TencentCloud zone DNS records")
	privateZones, err := p.getZones()
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0)

	for _, zone := range privateZones {
		records, err := p.getZoneRecords(*zone.ZoneId)
		if err != nil {
			return nil, err
		}

		endpointMap := make(map[string]*endpoint.Endpoint)
		for _, record := range records {
			if !provider.SupportedRecordType(*record.RecordType) {
				continue
			}

			name := getDNSName(*record.SubDomain, *zone.Domain)
			key := toRecordKey(*record.RecordType, name, nil)
			target := wrapWithQuotes(*record.RecordType, *record.RecordValue)

			if _, exist := endpointMap[key]; !exist {
				ttl := endpoint.TTL(*record.TTL)
				recordType := *record.RecordType
				// recordID := *record.RecordId
				endpointMap[key] = endpoint.NewEndpointWithTTL(name, recordType, ttl, target)
			} else {
				endpointMap[key].Targets = append(endpointMap[key].Targets, target)
			}
		}
		for _, ep := range endpointMap {
			endpoints = append(endpoints, ep)
		}
	}
	log.Infof("Found %d TencentCloud zone DNS record(s).", len(endpoints))

	return endpoints, nil
}

func (p *TencentCloudProvider) getZones() ([]*privatedns.PrivateZone, error) {
	filters := []*privatedns.Filter{
		{
			Name: common.StringPtr("Vpc"),
			Values: []*string{
				common.StringPtr(p.vpcID),
			},
		},
	}

	if p.zoneIDFilter != nil && p.zoneIDFilter.IsConfigured() {
		zoneIDs := make([]*string, len(p.zoneIDFilter.ZoneIDs))
		for index, zoneID := range p.zoneIDFilter.ZoneIDs {
			zoneIDs[index] = common.StringPtr(zoneID)
		}
		filters = append(filters, &privatedns.Filter{
			Name:   common.StringPtr("ZoneId"),
			Values: zoneIDs,
		})
	}

	request := privatedns.NewDescribePrivateZoneListRequest()
	request.Filters = filters
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(pageSize)

	privateZones := make([]*privatedns.PrivateZone, 0)
	totalCount := int64(pageSize)
	for *request.Offset < totalCount {
		response, err := p.pvtApi.DescribePrivateZoneList(request)
		if err != nil {
			return nil, err
		}

		for _, privateZone := range response.Response.PrivateZoneSet {
			if p.domainFilter != nil && !p.domainFilter.Match(*privateZone.Domain) {
				continue
			}
			privateZones = append(privateZones, privateZone)
		}

		if response.Response.TotalCount == nil {
			break
		}
		totalCount = *response.Response.TotalCount
		request.Offset = common.Int64Ptr(*request.Offset + int64(len(response.Response.PrivateZoneSet)))
	}

	return privateZones, nil
}

func (p *TencentCloudProvider) getZoneRecordMap(domain string, zoneId string) (map[string]*zoneRecordGroup, error) {
	records, err := p.getZoneRecords(zoneId)
	if err != nil {
		return nil, err
	}

	recordGroupMap := make(map[string]*zoneRecordGroup)
	for _, record := range records {
		key := toRecordKey(*record.RecordType, domain, record.SubDomain)

		if _, exists := recordGroupMap[key]; !exists {
			recordGroupMap[key] = &zoneRecordGroup{
				ZoneId:     zoneId,
				Domain:     domain,
				SubDomain:  *record.SubDomain,
				RecordType: *record.RecordType,
				RecordList: make([]*privatedns.PrivateZoneRecord, 0),
			}
		}
		recordGroupMap[key].RecordList = append(recordGroupMap[key].RecordList, record)
	}
	return recordGroupMap, nil
}

func (p *TencentCloudProvider) getZoneRecords(zoneID string) ([]*privatedns.PrivateZoneRecord, error) {
	request := privatedns.NewDescribePrivateZoneRecordListRequest()
	request.ZoneId = common.StringPtr(zoneID)
	request.Offset = common.Int64Ptr(0)
	request.Limit = common.Int64Ptr(pageSize)

	records := make([]*privatedns.PrivateZoneRecord, 0)
	totalCount := int64(pageSize)
	for *request.Offset < totalCount {
		response, err := p.pvtApi.DescribePrivateZoneRecordList(request)
		if err != nil {
			return nil, err
		}

		for _, record := range response.Response.RecordSet {
			if *record.SubDomain == apexRecord && *record.RecordType == endpoint.RecordTypeNS {
				continue
			}
			records = append(records, record)
		}

		if response.Response.TotalCount == nil {
			break
		}
		totalCount = *response.Response.TotalCount
		request.Offset = common.Int64Ptr(*request.Offset + int64(len(response.Response.RecordSet)))
	}
	return records, nil
}

func (p *TencentCloudProvider) applyChangesForZone(changes *plan.Changes) error {
	log.Infof("Apply changes to TencentCloud zone DNS: %++v", *changes)
	zones, err := p.getZones()
	if err != nil {
		return err
	}

	zoneIDMapper := provider.ZoneIDName{}
	recordGroupMap := make(map[string]*zoneRecordGroup)
	for _, zone := range zones {
		zoneIDMapper.Add(*zone.ZoneId, *zone.Domain)
		recordGroups, err := p.getZoneRecordMap(*zone.Domain, *zone.ZoneId)
		if err != nil {
			continue
		}

		// Tencent Cloud PrivateDNS requires each private zone to keep at least one record.
		containsBase := false

		for key, group := range recordGroups {
			recordGroupMap[key] = group
			if !containsBase && containsBaseRecord(group.RecordList) {
				containsBase = true
			}
		}

		if !containsBase {
			txtRecord := endpoint.NewEndpoint(*zone.Domain, endpoint.RecordTypeTXT, "tencent_provider_record")
			if err := p.createZoneRecord(*zone.ZoneId, *zone.Domain, txtRecord); err != nil {
				return err
			}
		}
	}

	deleteRecordIdsMap := make(map[string][]*string, 0)
	for _, change := range [][]*endpoint.Endpoint{changes.Delete, changes.UpdateOld} {
		for _, ep := range change {
			key := toEndpointKey(ep)
			if group, exists := recordGroupMap[key]; exists {
				zoneId := group.ZoneId
				for _, record := range group.RecordList {
					target := wrapWithQuotes(*record.RecordType, *record.RecordValue)
					if slices.Contains(ep.Targets, target) {
						if _, exist := deleteRecordIdsMap[zoneId]; !exist {
							deleteRecordIdsMap[zoneId] = []*string{record.RecordId}
						} else {
							deleteRecordIdsMap[zoneId] = append(deleteRecordIdsMap[zoneId], record.RecordId)
						}
					}
				}
			}
		}
	}

	createEndpoints := make(map[string][]*extendEndpoint)
	for _, change := range [][]*endpoint.Endpoint{changes.Create, changes.UpdateNew} {
		for _, ep := range change {
			zoneId, zoneName := zoneIDMapper.FindZone(ep.DNSName)
			if zoneId != "" {
				createEndpoints[zoneId] = append(
					createEndpoints[zoneId],
					&extendEndpoint{
						Endpoint: ep,
						ZoneId:   zoneId,
						Domain:   zoneName,
					},
				)
			}
		}
	}

	var errors []error
	for zoneId, recordIds := range deleteRecordIdsMap {
		if err := p.deleteZoneRecords(&zoneId, recordIds); err != nil {
			errors = append(errors, err)
		}
	}

	if len(createEndpoints) > 0 {
		for zoneId, endpoints := range createEndpoints {
			for _, ep := range endpoints {
				if err := p.createZoneRecord(zoneId, ep.Domain, ep.Endpoint); err != nil {
					errors = append(errors, err)
				}
			}
		}
	}

	if len(errors) > 0 {
		return fmt.Errorf("Failed to apply changes to TencentCloud zone DNS: %v", errors)
	}

	return nil
}

func (p *TencentCloudProvider) createZoneRecord(zoneId string, domain string, ep *endpoint.Endpoint) error {
	if zoneId == "" || domain == "" || ep == nil {
		return fmt.Errorf("Invalid input for creating zone DNS record")
	}

	if p.dryRun {
		log.Infof("Dry run: create TencentCloud zone DNS record %s with endpoint %v", zoneId, ep)
		return nil
	}

	subDomain := getSubDomain(domain, ep)
	ttl := int64(defaultPvtTTL)
	if ep.RecordTTL.IsConfigured() {
		ttl = int64(ep.RecordTTL)
	}

	request := privatedns.NewCreatePrivateZoneRecordRequest()
	request.ZoneId = &zoneId
	request.SubDomain = &subDomain
	request.RecordType = &ep.RecordType
	request.TTL = &ttl

	var errs []error
	for _, target := range ep.Targets {
		target := unwrapQuotes(ep.RecordType, target)
		req := *request
		req.RecordValue = &target

		if _, err := p.pvtApi.CreatePrivateZoneRecord(&req); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("Failed to create %d records: %v", len(errs), errs)
	}

	return nil
}

func (p *TencentCloudProvider) deleteZoneRecords(zoneId *string, recordIds []*string) error {
	request := privatedns.NewDeletePrivateZoneRecordRequest()
	request.ZoneId = zoneId
	request.RecordIdSet = recordIds

	if p.dryRun {
		log.Infof("Dry run: delete TencentCloud zone DNS records %v in zone %s", recordIds, *zoneId)
		return nil
	}

	if _, err := p.pvtApi.DeletePrivateZoneRecord(request); err != nil {
		return err
	}

	return nil
}

func containsBaseRecord(records []*privatedns.PrivateZoneRecord) bool {
	for _, r := range records {
		if *r.SubDomain == apexRecord && *r.RecordType == endpoint.RecordTypeTXT && *r.RecordValue == "tencent_provider_record" {
			return true
		}
	}
	return false
}

func toRecordKey(recordType, domain string, subDomain *string) string {
	if subDomain != nil {
		domain = getDNSName(*subDomain, domain)
	}
	return fmt.Sprintf("%s:%s", recordType, domain)
}

func toEndpointKey(ep *endpoint.Endpoint) string {
	return fmt.Sprintf("%s:%s", ep.RecordType, ep.DNSName)
}

func getSubDomain(domain string, ep *endpoint.Endpoint) string {
	name := strings.TrimSuffix(ep.DNSName, ".")
	domain = strings.TrimSuffix(domain, ".")
	name = strings.TrimSuffix(name, domain)
	name = strings.TrimSuffix(name, ".")

	if name == "" {
		return apexRecord
	}
	return name
}

func getDNSName(subDomain, domain string) string {
	if subDomain == apexRecord {
		return domain
	}
	return subDomain + "." + domain
}

func wrapWithQuotes(recordType, target string) string {
	if recordType == endpoint.RecordTypeTXT && strings.HasPrefix(target, `heritage=`) {
		return fmt.Sprintf("%q", target)
	}
	return target
}

func unwrapQuotes(recordType, target string) string {
	if recordType == endpoint.RecordTypeTXT && strings.HasPrefix(target, `"heritage=`) {
		return strings.Trim(target, `"`)
	}
	return target
}
