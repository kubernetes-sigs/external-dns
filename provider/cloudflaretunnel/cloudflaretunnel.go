package cloudflaretunnel

import (
	"context"
	"fmt"
	"os"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

type CloudFlareTunnelProvider struct {
	provider.BaseProvider
	Client            *cloudflare.API
	domainFilter      endpoint.DomainFilter
	DryRun            bool
	accountId         string
	tunnelId          string
	zoneNameIDMapper  provider.ZoneIDName
	zoneIDFilter      provider.ZoneIDFilter
	DNSRecordsPerPage int
}

var defaultOriginRequestConfig = cloudflare.OriginRequestConfig{
	NoTLSVerify: boolPtr(true),
	Http2Origin: boolPtr(true),
}

func NewCloudFlareTunnelProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool, dnsRecordsPerPage int) (*CloudFlareTunnelProvider, error) {
	var (
		client *cloudflare.API
		err    error
	)

	token, ok := os.LookupEnv("CF_API_TOKEN")
	if ok {
		if strings.HasPrefix(token, "file:") {
			tokenBytes, err := os.ReadFile(strings.TrimPrefix(token, "file:"))
			if err != nil {
				return nil, fmt.Errorf("failed to read CF_API_TOKEN from file: %w", err)
			}
			token = string(tokenBytes)
		}
		client, err = cloudflare.NewWithAPIToken(token)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize cloudflare provider: %v", err)
		}
	} else {
		client, err = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
		if err != nil {
			return nil, fmt.Errorf("failed to initialize cloudflare provider: %v", err)
		}
	}

	accountId, ok := os.LookupEnv("CF_ACCOUNT_ID")
	if !ok {
		return nil, fmt.Errorf("failed to get cloudflare account id: please set env, CF_ACCOUNT_ID")
	}

	tunnelId, ok := os.LookupEnv("CF_TUNNEL_ID")
	if !ok {
		return nil, fmt.Errorf("failed to get cloudflare tunnel id: please set env, CF_TUNNEL_ID")
	}

	provider := &CloudFlareTunnelProvider{
		Client:           client,
		accountId:        accountId,
		DryRun:           dryRun,
		tunnelId:         tunnelId,
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		zoneNameIDMapper: provider.ZoneIDName{},
	}
	return provider, nil
}

func (p *CloudFlareTunnelProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	configResult, err := p.Client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(p.accountId), p.tunnelId)
	if err != nil {
		return nil, fmt.Errorf("failed to get tunnel configs: %v", err)
	}

	endpoints := []*endpoint.Endpoint{}
	for _, config := range configResult.Config.Ingress {
		endpoint := endpoint.NewEndpoint(config.Hostname, "A", config.Service)
		endpoints = append(endpoints, endpoint)
	}
	return endpoints, nil
}

func (p *CloudFlareTunnelProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	err := p.updateZoneIdMapper(ctx)
	if err != nil {
		return fmt.Errorf("failed to update zoneidmapper: %v", err)
	}

	oldConfigResult, err := p.Client.GetTunnelConfiguration(ctx, cloudflare.AccountIdentifier(p.accountId), p.tunnelId)
	if err != nil {
		return fmt.Errorf("failed to get tunnel configs: %v", err)
	}

	ingresses := make(map[string]cloudflare.UnvalidatedIngressRule)
	var catchAll cloudflare.UnvalidatedIngressRule
	for _, v := range oldConfigResult.Config.Ingress {
		if v.Hostname == "" {
			catchAll = v
			break
		}
		ingresses[v.Hostname] = v
	}

	param := cloudflare.TunnelConfigurationParams{TunnelID: p.tunnelId, Config: oldConfigResult.Config}
	param.Config.Ingress = []cloudflare.UnvalidatedIngressRule{}

	for _, endpoint := range changes.Create {
		if endpoint.RecordType != "A" {
			continue
		}
		ingresses[endpoint.DNSName] = cloudflare.UnvalidatedIngressRule{
			Hostname:      endpoint.DNSName,
			Service:       convertHttps(endpoint.Targets[0]),
			OriginRequest: &defaultOriginRequestConfig,
		}

		zoneID, _ := p.zoneNameIDMapper.FindZone(endpoint.DNSName)
		if zoneID == "" {
			fmt.Println("zoneid is empty. skipping...")
			continue
		}

		params := cloudflare.CreateDNSRecordParams{
			Name:    endpoint.DNSName,
			TTL:     1, // auto
			Proxied: boolPtr(true),
			Type:    "CNAME",
			Content: fmt.Sprintf("%v.cfargotunnel.com", p.tunnelId),
		}
		_, err := p.Client.CreateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			fmt.Printf("failed to create dns record: %v\n", err)
			continue
		}
	}

	for i, desired := range changes.UpdateNew {
		if desired.RecordType != "A" {
			continue
		}

		current := changes.UpdateOld[i]
		if !(ingresses[desired.DNSName].Hostname == current.DNSName && ingresses[desired.DNSName].Service == current.Targets[0]) {
			log.Println("failed to update dns")
			continue
		}

		ingresses[desired.DNSName] = cloudflare.UnvalidatedIngressRule{
			Hostname:      desired.DNSName,
			Service:       convertHttps(desired.Targets[0]),
			OriginRequest: &defaultOriginRequestConfig,
		}

		zoneID, _ := p.zoneNameIDMapper.FindZone(desired.DNSName)
		if zoneID == "" {
			fmt.Println("zoneid is empty. skipping...")
			continue
		}

		records, err := p.listDNSRecordsWithAutoPagination(ctx, zoneID)
		if err != nil {
			return err
		}
		recordID := p.getRecordID(records, cloudflare.DNSRecord{
			Name:    desired.DNSName,
			Type:    "CNAME",
			Content: desired.Targets[0],
		})

		params := cloudflare.UpdateDNSRecordParams{
			ID: recordID,
		}
		_, err = p.Client.UpdateDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			fmt.Printf("failed to update dns record: %v\n", err)
			continue
		}
	}

	for _, endpoint := range changes.Delete {
		if endpoint.RecordType != "A" {
			continue
		}
		delete(ingresses, endpoint.DNSName)

		zoneID, _ := p.zoneNameIDMapper.FindZone(endpoint.DNSName)
		if zoneID == "" {
			fmt.Println("zoneid is empty. skipping...")
			continue
		}

		records, err := p.listDNSRecordsWithAutoPagination(ctx, zoneID)
		if err != nil {
			return err
		}
		recordID := p.getRecordID(records, cloudflare.DNSRecord{
			Name:    endpoint.DNSName,
			Type:    "CNAME",
			Content: endpoint.Targets[0],
		})
		err = p.Client.DeleteDNSRecord(ctx, cloudflare.ZoneIdentifier(zoneID), recordID)
		if err != nil {
			fmt.Printf("failed to delete dns record: %v\n", err)
			continue
		}
	}

	for _, v := range ingresses {
		//fmt.Println(v.Hostname, ":", v.Service)
		param.Config.Ingress = append(param.Config.Ingress, v)
	}
	param.Config.Ingress = append(param.Config.Ingress, catchAll)

	if p.DryRun {
		return nil
	}

	_, err = p.Client.UpdateTunnelConfiguration(ctx, cloudflare.AccountIdentifier(p.accountId), param)
	if err != nil {
		return fmt.Errorf("failed to update tunnel configs: %v", err)
	}
	fmt.Println("successfully update tunnel config")
	return nil
}

// Zones returns the list of hosted zones.
func (p *CloudFlareTunnelProvider) Zones(ctx context.Context) ([]cloudflare.Zone, error) {
	result := []cloudflare.Zone{}

	// if there is a zoneIDfilter configured
	// && if the filter isn't just a blank string (used in tests)
	if len(p.zoneIDFilter.ZoneIDs) > 0 && p.zoneIDFilter.ZoneIDs[0] != "" {
		log.Debugln("zoneIDFilter configured. only looking up zone IDs defined")
		for _, zoneID := range p.zoneIDFilter.ZoneIDs {
			log.Debugf("looking up zone %s", zoneID)
			detailResponse, err := p.Client.ZoneDetails(ctx, zoneID)
			if err != nil {
				log.Errorf("zone %s lookup failed, %v", zoneID, err)
				return result, err
			}
			log.WithFields(log.Fields{
				"zoneName": detailResponse.Name,
				"zoneID":   detailResponse.ID,
			}).Debugln("adding zone for consideration")
			result = append(result, detailResponse)
		}
		return result, nil
	}

	log.Debugln("no zoneIDFilter configured, looking at all zones")

	zonesResponse, err := p.Client.ListZonesContext(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range zonesResponse.Result {
		if !p.domainFilter.Match(zone.Name) {
			log.Debugf("zone %s not in domain filter", zone.Name)
			continue
		}
		result = append(result, zone)
	}
	return result, nil
}

func (p CloudFlareTunnelProvider) updateZoneIdMapper(ctx context.Context) error {
	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}

	for _, z := range zones {
		p.zoneNameIDMapper.Add(z.ID, z.Name)
	}
	return nil
}

func (p *CloudFlareTunnelProvider) getRecordID(records []cloudflare.DNSRecord, record cloudflare.DNSRecord) string {
	for _, zoneRecord := range records {
		if zoneRecord.Name == record.Name && zoneRecord.Type == record.Type && zoneRecord.Content == record.Content {
			return zoneRecord.ID
		}
	}
	return ""
}

// listDNSRecords performs automatic pagination of results on requests to cloudflare.ListDNSRecords with custom per_page values
func (p *CloudFlareTunnelProvider) listDNSRecordsWithAutoPagination(ctx context.Context, zoneID string) ([]cloudflare.DNSRecord, error) {
	var records []cloudflare.DNSRecord
	resultInfo := cloudflare.ResultInfo{PerPage: p.DNSRecordsPerPage, Page: 1}
	params := cloudflare.ListDNSRecordsParams{ResultInfo: resultInfo}
	for {
		pageRecords, resultInfo, err := p.Client.ListDNSRecords(ctx, cloudflare.ZoneIdentifier(zoneID), params)
		if err != nil {
			return nil, err
		}

		records = append(records, pageRecords...)
		params.ResultInfo = resultInfo.Next()
		if params.ResultInfo.Done() {
			break
		}
	}
	return records, nil
}

func convertHttps(target string) string {
	return fmt.Sprintf("https://%v:443", target)
}

// boolPtr is used as a helper function to return a pointer to a boolean
// Needed because some parameters require a pointer.
func boolPtr(b bool) *bool {
	return &b
}
