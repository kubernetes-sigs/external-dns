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

package aws

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	route53types "github.com/aws/aws-sdk-go-v2/service/route53/types"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	defaultAWSProfile = "default"
	recordTTL         = 300
	// From the experiments, it seems that the default MaxItems applied is 100,
	// and that, on the server side, there is a hard limit of 300 elements per page.
	// After a discussion with AWS representatives, clients should accept
	// when fewer items are returned, and still paginate accordingly.
	// As we are using the standard AWS client, this should already be compliant.
	// Hence, if AWS ever decides to raise this limit, we will automatically reduce the pressure on rate limits
	route53PageSize int32 = 300
	// providerSpecificAlias specifies whether a CNAME endpoint maps to an AWS ALIAS record.
	providerSpecificAlias            = "alias"
	providerSpecificTargetHostedZone = "aws/target-hosted-zone"
	// providerSpecificEvaluateTargetHealth specifies whether an AWS ALIAS record
	// has the EvaluateTargetHealth field set to true. Present iff the endpoint
	// has a `providerSpecificAlias` value of `true`.
	providerSpecificEvaluateTargetHealth       = "aws/evaluate-target-health"
	providerSpecificWeight                     = "aws/weight"
	providerSpecificRegion                     = "aws/region"
	providerSpecificFailover                   = "aws/failover"
	providerSpecificGeolocationContinentCode   = "aws/geolocation-continent-code"
	providerSpecificGeolocationCountryCode     = "aws/geolocation-country-code"
	providerSpecificGeolocationSubdivisionCode = "aws/geolocation-subdivision-code"
	providerSpecificMultiValueAnswer           = "aws/multi-value-answer"
	providerSpecificHealthCheckID              = "aws/health-check-id"
	sameZoneAlias                              = "same-zone"
	// Currently supported up to 10 health checks or hosted zones.
	// https://docs.aws.amazon.com/Route53/latest/APIReference/API_ListTagsForResources.html#API_ListTagsForResources_RequestSyntax
	batchSize = 10
)

// see elb: https://docs.aws.amazon.com/general/latest/gr/elb.html
var canonicalHostedZones = map[string]string{
	// Application Load Balancers and Classic Load Balancers
	"us-east-2.elb.amazonaws.com":         "Z3AADJGX6KTTL2",
	"us-east-1.elb.amazonaws.com":         "Z35SXDOTRQ7X7K",
	"us-west-1.elb.amazonaws.com":         "Z368ELLRRE2KJ0",
	"us-west-2.elb.amazonaws.com":         "Z1H1FL5HABSF5",
	"ca-central-1.elb.amazonaws.com":      "ZQSVJUPU6J1EY",
	"ca-west-1.elb.amazonaws.com":         "Z06473681N0SF6OS049SD",
	"ap-east-1.elb.amazonaws.com":         "Z3DQVH9N71FHZ0",
	"ap-south-1.elb.amazonaws.com":        "ZP97RAFLXTNZK",
	"ap-south-2.elb.amazonaws.com":        "Z0173938T07WNTVAEPZN",
	"ap-northeast-2.elb.amazonaws.com":    "ZWKZPGTI48KDX",
	"ap-northeast-3.elb.amazonaws.com":    "Z5LXEXXYW11ES",
	"ap-southeast-1.elb.amazonaws.com":    "Z1LMS91P8CMLE5",
	"ap-southeast-2.elb.amazonaws.com":    "Z1GM3OXH4ZPM65",
	"ap-southeast-3.elb.amazonaws.com":    "Z08888821HLRG5A9ZRTER",
	"ap-southeast-4.elb.amazonaws.com":    "Z09517862IB2WZLPXG76F",
	"ap-southeast-5.elb.amazonaws.com":    "Z06010284QMVVW7WO5J",
	"ap-southeast-7.elb.amazonaws.com":    "Z0390008CMBRTHFGWBCB",
	"ap-northeast-1.elb.amazonaws.com":    "Z14GRHDCWA56QT",
	"eu-central-1.elb.amazonaws.com":      "Z215JYRZR1TBD5",
	"eu-central-2.elb.amazonaws.com":      "Z06391101F2ZOEP8P5EB3",
	"eu-west-1.elb.amazonaws.com":         "Z32O12XQLNTSW2",
	"eu-west-2.elb.amazonaws.com":         "ZHURV8PSTC4K8",
	"eu-west-3.elb.amazonaws.com":         "Z3Q77PNBQS71R4",
	"eu-north-1.elb.amazonaws.com":        "Z23TAZ6LKFMNIO",
	"eu-south-1.elb.amazonaws.com":        "Z3ULH7SSC9OV64",
	"eu-south-2.elb.amazonaws.com":        "Z0956581394HF5D5LXGAP",
	"sa-east-1.elb.amazonaws.com":         "Z2P70J7HTTTPLU",
	"cn-north-1.elb.amazonaws.com.cn":     "Z1GDH35T77C1KE",
	"cn-northwest-1.elb.amazonaws.com.cn": "ZM7IZAIOVVDZF",
	"us-gov-west-1.elb.amazonaws.com":     "Z33AYJ8TM3BH4J",
	"us-gov-east-1.elb.amazonaws.com":     "Z166TLBEWOO7G0",
	"mx-central-1.elb.amazonaws.com":      "Z023552324OKD1BB28BH5",
	"me-central-1.elb.amazonaws.com":      "Z08230872XQRWHG2XF6I",
	"me-south-1.elb.amazonaws.com":        "ZS929ML54UICD",
	"af-south-1.elb.amazonaws.com":        "Z268VQBMOI5EKX",
	"il-central-1.elb.amazonaws.com":      "Z09170902867EHPV2DABU",
	// Network Load Balancers https://docs.aws.amazon.com/general/latest/gr/elb.html#elb_region
	"elb.us-east-2.amazonaws.com":         "ZLMOA37VPKANP",
	"elb.us-east-1.amazonaws.com":         "Z26RNL4JYFTOTI",
	"elb.us-west-1.amazonaws.com":         "Z24FKFUX50B4VW",
	"elb.us-west-2.amazonaws.com":         "Z18D5FSROUN65G",
	"elb.ca-central-1.amazonaws.com":      "Z2EPGBW3API2WT",
	"elb.ca-west-1.amazonaws.com":         "Z02754302KBB00W2LKWZ9",
	"elb.ap-east-1.amazonaws.com":         "Z12Y7K3UBGUAD1",
	"elb.ap-south-1.amazonaws.com":        "ZVDDRBQ08TROA",
	"elb.ap-south-2.amazonaws.com":        "Z0711778386UTO08407HT",
	"elb.ap-northeast-3.amazonaws.com":    "Z1GWIQ4HH19I5X",
	"elb.ap-northeast-2.amazonaws.com":    "ZIBE1TIR4HY56",
	"elb.ap-southeast-1.amazonaws.com":    "ZKVM4W9LS7TM",
	"elb.ap-southeast-2.amazonaws.com":    "ZCT6FZBF4DROD",
	"elb.ap-southeast-3.amazonaws.com":    "Z01971771FYVNCOVWJU1G",
	"elb.ap-southeast-4.amazonaws.com":    "Z01156963G8MIIL7X90IV",
	"elb.ap-southeast-5.amazonaws.com":    "Z026317210H9ACVTRO6FB",
	"elb.ap-southeast-7.amazonaws.com":    "Z054363131YWATEMWRG5L",
	"elb.ap-northeast-1.amazonaws.com":    "Z31USIVHYNEOWT",
	"elb.eu-central-1.amazonaws.com":      "Z3F0SRJ5LGBH90",
	"elb.eu-central-2.amazonaws.com":      "Z02239872DOALSIDCX66S",
	"elb.eu-west-1.amazonaws.com":         "Z2IFOLAFXWLO4F",
	"elb.eu-west-2.amazonaws.com":         "ZD4D7Y8KGAS4G",
	"elb.eu-west-3.amazonaws.com":         "Z1CMS0P5QUZ6D5",
	"elb.eu-north-1.amazonaws.com":        "Z1UDT6IFJ4EJM",
	"elb.eu-south-1.amazonaws.com":        "Z23146JA1KNAFP",
	"elb.eu-south-2.amazonaws.com":        "Z1011216NVTVYADP1SSV",
	"elb.sa-east-1.amazonaws.com":         "ZTK26PT1VY4CU",
	"elb.cn-north-1.amazonaws.com.cn":     "Z3QFB96KMJ7ED6",
	"elb.cn-northwest-1.amazonaws.com.cn": "ZQEIKTCZ8352D",
	"elb.us-gov-west-1.amazonaws.com":     "ZMG1MZ2THAWF1",
	"elb.us-gov-east-1.amazonaws.com":     "Z1ZSMQQ6Q24QQ8",
	"elb.mx-central-1.amazonaws.com":      "Z02031231H3ID6HYJ9A7U",
	"elb.me-central-1.amazonaws.com":      "Z00282643NTTLPANJJG2P",
	"elb.me-south-1.amazonaws.com":        "Z3QSRYVP46NYYV",
	"elb.af-south-1.amazonaws.com":        "Z203XCE67M25HM",
	"elb.il-central-1.amazonaws.com":      "Z0313266YDI6ZRHTGQY4",
	// Global Accelerator
	"awsglobalaccelerator.com": "Z2BJ6XQ5FK7U4H",
	// Cloudfront and AWS API Gateway edge-optimized endpoints
	"cloudfront.net": "Z2FDTNDATAQYW2",
	// VPC Endpoint (PrivateLink) https://github.com/kubernetes-sigs/external-dns/issues/3429#issuecomment-1440415806
	"eu-west-2.vpce.amazonaws.com":      "Z7K1066E3PUKB",
	"us-east-2.vpce.amazonaws.com":      "ZC8PG0KIFKBRI",
	"af-south-1.vpce.amazonaws.com":     "Z09302161J80N9A7UTP7U",
	"ap-east-1.vpce.amazonaws.com":      "Z2LIHJ7PKBEMWN",
	"ap-northeast-1.vpce.amazonaws.com": "Z2E726K9Y6RL4W",
	"ap-northeast-2.vpce.amazonaws.com": "Z27UANNT0PRK1T",
	"ap-northeast-3.vpce.amazonaws.com": "Z376B5OMM2JZL2",
	"ap-south-1.vpce.amazonaws.com":     "Z2KVTB3ZLFM7JR",
	"ap-south-2.vpce.amazonaws.com":     "Z0952991RWSF5AHIQDIY",
	"ap-southeast-1.vpce.amazonaws.com": "Z18LLCSTV4NVNL",
	"ap-southeast-2.vpce.amazonaws.com": "ZDK2GCRPAFKGO",
	"ap-southeast-3.vpce.amazonaws.com": "Z03881013RZ9BYYZO8N5W",
	"ap-southeast-4.vpce.amazonaws.com": "Z07508191CO1RNBX3X3AU",
	"ca-central-1.vpce.amazonaws.com":   "ZRCXCF510Y6P9",
	"eu-central-1.vpce.amazonaws.com":   "Z273ZU8SZ5RJPC",
	"eu-central-2.vpce.amazonaws.com":   "Z045369019J4FUQ4S272E",
	"eu-north-1.vpce.amazonaws.com":     "Z3OWWK6JFDEDGC",
	"eu-south-1.vpce.amazonaws.com":     "Z2A5FDNRLY7KZG",
	"eu-south-2.vpce.amazonaws.com":     "Z014396544HENR57XQCJ",
	"eu-west-1.vpce.amazonaws.com":      "Z38GZ743OKFT7T",
	"eu-west-3.vpce.amazonaws.com":      "Z1DWHTMFP0WECP",
	"me-central-1.vpce.amazonaws.com":   "Z07122992YCEUCB9A9570",
	"me-south-1.vpce.amazonaws.com":     "Z3B95P3VBGEQGY",
	"sa-east-1.vpce.amazonaws.com":      "Z2LXUWEVLCVZIB",
	"us-east-1.vpce.amazonaws.com":      "Z7HUB22UULQXV",
	"us-gov-east-1.vpce.amazonaws.com":  "Z2MU5TEIGO9WXB",
	"us-gov-west-1.vpce.amazonaws.com":  "Z12529ZODG2B6H",
	"us-west-1.vpce.amazonaws.com":      "Z12I86A8N7VCZO",
	"us-west-2.vpce.amazonaws.com":      "Z1YSA3EXCYUU9Z",
	// AWS API Gateway (Regional endpoints)
	// See: https://docs.aws.amazon.com/general/latest/gr/apigateway.html
	"execute-api.us-east-2.amazonaws.com":      "ZOJJZC49E0EPZ",
	"execute-api.us-east-1.amazonaws.com":      "Z1UJRXOUMOOFQ8",
	"execute-api.us-west-1.amazonaws.com":      "Z2MUQ32089INYE",
	"execute-api.us-west-2.amazonaws.com":      "Z2OJLYMUO9EFXC",
	"execute-api.af-south-1.amazonaws.com":     "Z2DHW2332DAMTN",
	"execute-api.ap-east-1.amazonaws.com":      "Z3FD1VL90ND7K5",
	"execute-api.ap-south-1.amazonaws.com":     "Z3VO1THU9YC4UR",
	"execute-api.ap-northeast-2.amazonaws.com": "Z20JF4UZKIW1U8",
	"execute-api.ap-southeast-1.amazonaws.com": "ZL327KTPIQFUL",
	"execute-api.ap-southeast-2.amazonaws.com": "Z2RPCDW04V8134",
	"execute-api.ap-northeast-1.amazonaws.com": "Z1YSHQZHG15GKL",
	"execute-api.ca-central-1.amazonaws.com":   "Z19DQILCV0OWEC",
	"execute-api.eu-central-1.amazonaws.com":   "Z1U9ULNL0V5AJ3",
	"execute-api.eu-west-1.amazonaws.com":      "ZLY8HYME6SFDD",
	"execute-api.eu-west-2.amazonaws.com":      "ZJ5UAJN8Y3Z2Q",
	"execute-api.eu-south-1.amazonaws.com":     "Z3BT4WSQ9TDYZV",
	"execute-api.eu-west-3.amazonaws.com":      "Z3KY65QIEKYHQQ",
	"execute-api.eu-south-2.amazonaws.com":     "Z02499852UI5HEQ5JVWX3",
	"execute-api.eu-north-1.amazonaws.com":     "Z3UWIKFBOOGXPP",
	"execute-api.me-south-1.amazonaws.com":     "Z20ZBPC0SS8806",
	"execute-api.me-central-1.amazonaws.com":   "Z08780021BKYYY8U0YHTV",
	"execute-api.sa-east-1.amazonaws.com":      "ZCMLWB8V5SYIT",
	"execute-api.us-gov-east-1.amazonaws.com":  "Z3SE9ATJYCRCZJ",
	"execute-api.us-gov-west-1.amazonaws.com":  "Z1K6XKP9SAGWDV",
}

// Route53API is the subset of the AWS Route53 API that we actually use.  Add methods as required. Signatures must match exactly.
// https://github.com/aws/aws-sdk-go-v2/tree/main/service/route53
type Route53API interface {
	ListResourceRecordSets(ctx context.Context, input *route53.ListResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ListResourceRecordSetsOutput, error)
	ChangeResourceRecordSets(ctx context.Context, input *route53.ChangeResourceRecordSetsInput, optFns ...func(options *route53.Options)) (*route53.ChangeResourceRecordSetsOutput, error)
	CreateHostedZone(ctx context.Context, input *route53.CreateHostedZoneInput, optFns ...func(*route53.Options)) (*route53.CreateHostedZoneOutput, error)
	ListHostedZones(ctx context.Context, input *route53.ListHostedZonesInput, optFns ...func(options *route53.Options)) (*route53.ListHostedZonesOutput, error)
	ListTagsForResources(ctx context.Context, input *route53.ListTagsForResourcesInput, optFns ...func(options *route53.Options)) (*route53.ListTagsForResourcesOutput, error)
}

// Route53Change wrapper to handle ownership relation throughout the provider implementation
type Route53Change struct {
	route53types.Change
	OwnedRecord string
	sizeBytes   int
	sizeValues  int
}

type Route53Changes []*Route53Change

type profiledZone struct {
	profile string
	zone    *route53types.HostedZone
}

func (cs Route53Changes) Route53Changes() []route53types.Change {
	ret := []route53types.Change{}
	for _, c := range cs {
		ret = append(ret, c.Change)
	}
	return ret
}

type zoneTags map[string]map[string]string

// filterZonesByTags filters the provided zones map by matching the tags against the provider's zoneTagFilter.
// It removes any zones from the map that do not match the filter criteria.
func (z zoneTags) filterZonesByTags(p *AWSProvider, zones map[string]*profiledZone) {
	for zone, tags := range z {
		if !p.zoneTagFilter.Match(tags) {
			delete(zones, zone)
		}
	}
}

// append adds tags to the ZoneTags for a given zoneID.
func (z zoneTags) append(id string, tags []route53types.Tag) {
	zoneId := fmt.Sprintf("/hostedzone/%s", id)
	if _, exists := z[zoneId]; !exists {
		z[zoneId] = make(map[string]string)
	}
	for _, tag := range tags {
		z[zoneId][*tag.Key] = *tag.Value
	}
}

type zonesListCache struct {
	age      time.Time
	duration time.Duration
	zones    map[string]*profiledZone
}

// AWSProvider is an implementation of Provider for AWS Route53.
type AWSProvider struct {
	provider.BaseProvider
	clients               map[string]Route53API
	dryRun                bool
	batchChangeSize       int
	batchChangeSizeBytes  int
	batchChangeSizeValues int
	batchChangeInterval   time.Duration
	evaluateTargetHealth  bool
	// only consider hosted zones managing domains ending in this suffix
	domainFilter endpoint.DomainFilter
	// filter hosted zones by id
	zoneIDFilter provider.ZoneIDFilter
	// filter hosted zones by type (e.g. private or public)
	zoneTypeFilter provider.ZoneTypeFilter
	// filter hosted zones by tags
	zoneTagFilter provider.ZoneTagFilter
	// extend filter for subdomains in the zone (e.g. first.us-east-1.example.com)
	zoneMatchParent bool
	preferCNAME     bool
	zonesCache      *zonesListCache
	// queue for collecting changes to submit them in the next iteration, but after all other changes
	failedChangesQueue map[string]Route53Changes
}

// AWSConfig contains configuration to create a new AWS provider.
type AWSConfig struct {
	DomainFilter          endpoint.DomainFilter
	ZoneIDFilter          provider.ZoneIDFilter
	ZoneTypeFilter        provider.ZoneTypeFilter
	ZoneTagFilter         provider.ZoneTagFilter
	ZoneMatchParent       bool
	BatchChangeSize       int
	BatchChangeSizeBytes  int
	BatchChangeSizeValues int
	BatchChangeInterval   time.Duration
	EvaluateTargetHealth  bool
	PreferCNAME           bool
	DryRun                bool
	ZoneCacheDuration     time.Duration
}

// NewAWSProvider initializes a new AWS Route53 based Provider.
func NewAWSProvider(awsConfig AWSConfig, clients map[string]Route53API) (*AWSProvider, error) {
	provider := &AWSProvider{
		clients:               clients,
		domainFilter:          awsConfig.DomainFilter,
		zoneIDFilter:          awsConfig.ZoneIDFilter,
		zoneTypeFilter:        awsConfig.ZoneTypeFilter,
		zoneTagFilter:         awsConfig.ZoneTagFilter,
		zoneMatchParent:       awsConfig.ZoneMatchParent,
		batchChangeSize:       awsConfig.BatchChangeSize,
		batchChangeSizeBytes:  awsConfig.BatchChangeSizeBytes,
		batchChangeSizeValues: awsConfig.BatchChangeSizeValues,
		batchChangeInterval:   awsConfig.BatchChangeInterval,
		evaluateTargetHealth:  awsConfig.EvaluateTargetHealth,
		preferCNAME:           awsConfig.PreferCNAME,
		dryRun:                awsConfig.DryRun,
		zonesCache:            &zonesListCache{duration: awsConfig.ZoneCacheDuration},
		failedChangesQueue:    make(map[string]Route53Changes),
	}

	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *AWSProvider) Zones(ctx context.Context) (map[string]*route53types.HostedZone, error) {
	zones, err := p.zones(ctx)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*route53types.HostedZone, len(zones))
	for id, zone := range zones {
		result[id] = zone.zone
	}
	return result, nil
}

// zones returns the list of zones per AWS profile
func (p *AWSProvider) zones(ctx context.Context) (map[string]*profiledZone, error) {
	if p.zonesCache.zones != nil && time.Since(p.zonesCache.age) < p.zonesCache.duration {
		log.Debug("Using cached zones list")
		return p.zonesCache.zones, nil
	}
	log.Debug("Refreshing zones list cache")

	zones := make(map[string]*profiledZone)

	for profile, client := range p.clients {
		paginator := route53.NewListHostedZonesPaginator(client, &route53.ListHostedZonesInput{})

		for paginator.HasMorePages() {
			resp, err := paginator.NextPage(ctx)
			if err != nil {
				var te *route53types.ThrottlingException
				if errors.As(err, &te) {
					log.Infof("Skipping AWS profile %q due to provider side throttling: %v", profile, te.ErrorMessage())
					continue
				}
				// nothing to do here. Falling through to general error handling
				return nil, provider.NewSoftError(fmt.Errorf("failed to list hosted zones: %w", err))
			}
			var zonesToTagFilter []string
			for _, zone := range resp.HostedZones {
				if !p.zoneIDFilter.Match(*zone.Id) {
					continue
				}

				if !p.zoneTypeFilter.Match(zone) {
					continue
				}

				if !p.domainFilter.Match(*zone.Name) {
					if !p.zoneMatchParent {
						continue
					}
					if !p.domainFilter.MatchParent(*zone.Name) {
						continue
					}
				}

				if !p.zoneTagFilter.IsEmpty() {
					zonesToTagFilter = append(zonesToTagFilter, cleanZoneID(*zone.Id))
				}

				zones[*zone.Id] = &profiledZone{
					profile: profile,
					zone:    &zone,
				}
			}

			if len(zonesToTagFilter) > 0 {
				if zTags, err := p.tagsForZone(ctx, zonesToTagFilter, profile); err != nil {
					return nil, provider.NewSoftErrorf("failed to list tags for zones %w", err)
				} else {
					zTags.filterZonesByTags(p, zones)
				}
			}
		}
	}

	if log.IsLevelEnabled(log.DebugLevel) {
		for _, zone := range zones {
			log.Debugf("Considering zone: %s (domain: %s)", *zone.zone.Id, *zone.zone.Name)
		}
	}

	if p.zonesCache.duration > time.Duration(0) {
		p.zonesCache.zones = zones
		p.zonesCache.age = time.Now()
	}

	return zones, nil
}

// wildcardUnescape converts \\052.abc back to *.abc
// Route53 stores wildcards escaped: http://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html?shortFooter=true#domain-name-format-asterisk
func wildcardUnescape(s string) string {
	return strings.Replace(s, "\\052", "*", 1)
}

// See https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DomainNameFormat.html
// convertOctalToAscii decodes inputs that contain octal escape sequences into their original ASCII characters.
// The function returns converted string where any octal escape sequences have been replaced with their corresponding ASCII characters.
func convertOctalToAscii(input string) string {
	if !containsOctalSequence(input) {
		return input
	}
	result, err := strconv.Unquote("\"" + input + "\"")
	if err != nil {
		return input
	}
	return result
}

// validateDomainName checks if the domain name contains valid octal escape sequences.
func containsOctalSequence(domain string) bool {
	// Pattern to match valid octal escape sequences
	octalEscapePattern := `\\[0-3][0-7]{2}`
	octalEscapeRegex := regexp.MustCompile(octalEscapePattern)
	return octalEscapeRegex.MatchString(domain)
}

// Records returns the list of records in a given hosted zone.
func (p *AWSProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.zones(ctx)
	if err != nil {
		return nil, provider.NewSoftErrorf("records retrieval failed: %w", err)
	}

	return p.records(ctx, zones)
}

func (p *AWSProvider) records(ctx context.Context, zones map[string]*profiledZone) ([]*endpoint.Endpoint, error) {
	endpoints := make([]*endpoint.Endpoint, 0)

	for _, z := range zones {
		client := p.clients[z.profile]

		paginator := route53.NewListResourceRecordSetsPaginator(client, &route53.ListResourceRecordSetsInput{
			HostedZoneId: z.zone.Id,
			MaxItems:     aws.Int32(route53PageSize),
		})

		for paginator.HasMorePages() {
			resp, err := paginator.NextPage(ctx)
			if err != nil {
				return nil, provider.NewSoftErrorf("failed to list resource records sets for zone %s using aws profile %q: %w", *z.zone.Id, z.profile, err)
			}

			for _, r := range resp.ResourceRecordSets {
				newEndpoints := make([]*endpoint.Endpoint, 0)

				if !p.SupportedRecordType(r.Type) {
					continue
				}

				name := convertOctalToAscii(wildcardUnescape(*r.Name))

				var ttl endpoint.TTL
				if r.TTL != nil {
					ttl = endpoint.TTL(*r.TTL)
				}

				if len(r.ResourceRecords) > 0 {
					targets := make([]string, len(r.ResourceRecords))
					for idx, rr := range r.ResourceRecords {
						targets[idx] = *rr.Value
					}

					ep := endpoint.NewEndpointWithTTL(name, string(r.Type), ttl, targets...)
					if r.Type == endpoint.RecordTypeCNAME {
						ep = ep.WithProviderSpecific(providerSpecificAlias, "false")
					}
					newEndpoints = append(newEndpoints, ep)
				}

				if r.AliasTarget != nil {
					// Alias records don't have TTLs so provide the default to match the TXT generation
					if ttl == 0 {
						ttl = recordTTL
					}
					ep := endpoint.
						NewEndpointWithTTL(name, string(r.Type), ttl, *r.AliasTarget.DNSName).
						WithProviderSpecific(providerSpecificEvaluateTargetHealth, fmt.Sprintf("%t", r.AliasTarget.EvaluateTargetHealth)).
						WithProviderSpecific(providerSpecificAlias, "true")
					newEndpoints = append(newEndpoints, ep)
				}

				for _, ep := range newEndpoints {
					if r.SetIdentifier != nil {
						ep.SetIdentifier = *r.SetIdentifier
						switch {
						case r.Weight != nil:
							ep.WithProviderSpecific(providerSpecificWeight, fmt.Sprintf("%d", *r.Weight))
						case r.Region != "":
							ep.WithProviderSpecific(providerSpecificRegion, string(r.Region))
						case r.Failover != "":
							ep.WithProviderSpecific(providerSpecificFailover, string(r.Failover))
						case r.MultiValueAnswer != nil && *r.MultiValueAnswer:
							ep.WithProviderSpecific(providerSpecificMultiValueAnswer, "")
						case r.GeoLocation != nil:
							if r.GeoLocation.ContinentCode != nil {
								ep.WithProviderSpecific(providerSpecificGeolocationContinentCode, *r.GeoLocation.ContinentCode)
							} else {
								if r.GeoLocation.CountryCode != nil {
									ep.WithProviderSpecific(providerSpecificGeolocationCountryCode, *r.GeoLocation.CountryCode)
								}
								if r.GeoLocation.SubdivisionCode != nil {
									ep.WithProviderSpecific(providerSpecificGeolocationSubdivisionCode, *r.GeoLocation.SubdivisionCode)
								}
							}
						default:
							// one of the above needs to be set, otherwise SetIdentifier doesn't make sense
						}
					}

					if r.HealthCheckId != nil {
						ep.WithProviderSpecific(providerSpecificHealthCheckID, *r.HealthCheckId)
					}

					endpoints = append(endpoints, ep)
				}
			}
		}
	}

	return endpoints, nil
}

// Identify if old and new endpoints require DELETE/CREATE instead of UPDATE.
func (p *AWSProvider) requiresDeleteCreate(old *endpoint.Endpoint, new *endpoint.Endpoint) bool {
	// a change of record type
	if old.RecordType != new.RecordType {
		return true
	}

	// an ALIAS record change to/from an A
	if old.RecordType == endpoint.RecordTypeA {
		oldAlias, _ := old.GetProviderSpecificProperty(providerSpecificAlias)
		newAlias, _ := new.GetProviderSpecificProperty(providerSpecificAlias)
		if oldAlias != newAlias {
			return true
		}
	}

	// a set identifier change
	if old.SetIdentifier != new.SetIdentifier {
		return true
	}

	// a change of routing policy
	// default to true for geolocation properties if any geolocation property exists in old/new but not the other
	for _, propType := range [7]string{providerSpecificWeight, providerSpecificRegion, providerSpecificFailover,
		providerSpecificFailover, providerSpecificGeolocationContinentCode, providerSpecificGeolocationCountryCode,
		providerSpecificGeolocationSubdivisionCode} {
		_, oldPolicy := old.GetProviderSpecificProperty(propType)
		_, newPolicy := new.GetProviderSpecificProperty(propType)
		if oldPolicy != newPolicy {
			return true
		}
	}

	return false
}

func (p *AWSProvider) createUpdateChanges(newEndpoints, oldEndpoints []*endpoint.Endpoint) Route53Changes {
	var deletes []*endpoint.Endpoint
	var creates []*endpoint.Endpoint
	var updates []*endpoint.Endpoint

	for i, new := range newEndpoints {
		old := oldEndpoints[i]
		if p.requiresDeleteCreate(old, new) {
			deletes = append(deletes, old)
			creates = append(creates, new)
		} else {
			// Safe to perform an UPSERT.
			updates = append(updates, new)
		}
	}

	combined := make(Route53Changes, 0, len(deletes)+len(creates)+len(updates))
	combined = append(combined, p.newChanges(route53types.ChangeActionCreate, creates)...)
	combined = append(combined, p.newChanges(route53types.ChangeActionUpsert, updates)...)
	combined = append(combined, p.newChanges(route53types.ChangeActionDelete, deletes)...)
	return combined
}

// GetDomainFilter generates a filter to exclude any domain that is not controlled by the provider
func (p *AWSProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	zones, err := p.Zones(context.Background())
	if err != nil {
		log.Errorf("failed to list zones: %v", err)
		return endpoint.DomainFilter{}
	}
	zoneNames := []string(nil)
	for _, z := range zones {
		zoneNames = append(zoneNames, *z.Name, "."+*z.Name)
	}
	log.Infof("Applying provider record filter for domains: %v", zoneNames)
	return endpoint.NewDomainFilter(zoneNames)
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *AWSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, err := p.zones(ctx)
	if err != nil {
		return provider.NewSoftErrorf("failed to list zones, not applying changes: %w", err)
	}

	updateChanges := p.createUpdateChanges(changes.UpdateNew, changes.UpdateOld)

	combinedChanges := make(Route53Changes, 0, len(changes.Delete)+len(changes.Create)+len(updateChanges))
	combinedChanges = append(combinedChanges, p.newChanges(route53types.ChangeActionCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, p.newChanges(route53types.ChangeActionDelete, changes.Delete)...)
	combinedChanges = append(combinedChanges, updateChanges...)

	return p.submitChanges(ctx, combinedChanges, zones)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *AWSProvider) submitChanges(ctx context.Context, changes Route53Changes, zones map[string]*profiledZone) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		log.Info("All records are already up to date")
		return nil
	}

	// separate into per-zone change sets to be passed to the API.
	changesByZone := changesByZone(zones, changes)
	if len(changesByZone) == 0 {
		log.Info("All records are already up to date, there are no changes for the matching hosted zones")
	}

	var failedZones []string
	debugLevel := log.DebugLevel
	for z, cs := range changesByZone {
		log := log.WithFields(log.Fields{
			"zoneName": *zones[z].zone.Name,
			"zoneID":   z,
			"profile":  zones[z].profile,
		})

		var failedUpdate bool

		// group changes into new changes and into changes that failed in a previous iteration and are retried
		retriedChanges, newChanges := findChangesInQueue(cs, p.failedChangesQueue[z])
		p.failedChangesQueue[z] = nil

		batchCs := append(batchChangeSet(newChanges, p.batchChangeSize, p.batchChangeSizeBytes, p.batchChangeSizeValues),
			batchChangeSet(retriedChanges, p.batchChangeSize, p.batchChangeSizeBytes, p.batchChangeSizeValues)...)
		for i, b := range batchCs {
			if len(b) == 0 {
				continue
			}

			for _, c := range b {
				log.Infof("Desired change: %s %s %s", c.Action, *c.ResourceRecordSet.Name, c.ResourceRecordSet.Type)
			}

			if !p.dryRun {
				params := &route53.ChangeResourceRecordSetsInput{
					HostedZoneId: aws.String(z),
					ChangeBatch: &route53types.ChangeBatch{
						Changes: b.Route53Changes(),
					},
				}

				successfulChanges := 0

				client := p.clients[zones[z].profile]
				if _, err := client.ChangeResourceRecordSets(ctx, params); err != nil {
					log.Errorf("Failure in zone %s when submitting change batch: %v", *zones[z].zone.Name, err)

					changesByOwnership := groupChangesByNameAndOwnershipRelation(b)

					if len(changesByOwnership) > 1 {
						log.Debug("Trying to submit change sets one-by-one instead")
						for _, changes := range changesByOwnership {
							if log.Logger.IsLevelEnabled(debugLevel) {
								for _, c := range changes {
									log.Debugf("Desired change: %s %s %s", c.Action, *c.ResourceRecordSet.Name, c.ResourceRecordSet.Type)
								}
							}
							params.ChangeBatch = &route53types.ChangeBatch{
								Changes: changes.Route53Changes(),
							}
							if _, err := client.ChangeResourceRecordSets(ctx, params); err != nil {
								failedUpdate = true
								log.Errorf("Failed submitting change (error: %v), it will be retried in a separate change batch in the next iteration", err)
								p.failedChangesQueue[z] = append(p.failedChangesQueue[z], changes...)
							} else {
								successfulChanges = successfulChanges + len(changes)
							}
						}
					} else {
						failedUpdate = true
					}
				} else {
					successfulChanges = len(b)
				}

				if successfulChanges > 0 {
					// z is the R53 Hosted Zone ID already as aws.StringValue
					log.Infof("%d record(s) were successfully updated", successfulChanges)
				}

				if i != len(batchCs)-1 {
					time.Sleep(p.batchChangeInterval)
				}
			}
		}

		if failedUpdate {
			failedZones = append(failedZones, z)
		}
	}

	if len(failedZones) > 0 {
		return provider.NewSoftErrorf("failed to submit all changes for the following zones: %v", failedZones)
	}

	return nil
}

// newChanges returns a collection of Changes based on the given records and action.
func (p *AWSProvider) newChanges(action route53types.ChangeAction, endpoints []*endpoint.Endpoint) Route53Changes {
	changes := make(Route53Changes, 0, len(endpoints))

	for _, endpoint := range endpoints {
		change := p.newChange(action, endpoint)
		changes = append(changes, change)
	}

	return changes
}

// AdjustEndpoints modifies the provided endpoints (coming from various sources) to match
// the endpoints that the provider returns in `Records` so that the change plan will not have
// unneeded (potentially failing) changes.
// Example: CNAME endpoints pointing to ELBs will have a `alias` provider-specific property
// added to match the endpoints generated from existing alias records in Route53.
func (p *AWSProvider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	// Holds CNAME targets that we will treat as Alias records. Such records are
	// hard coded to 'A' type aliases but we also need their 'AAAA' counterparts.
	var aliasCnameAaaaEndpoints []*endpoint.Endpoint

	for _, ep := range endpoints {
		alias := false

		if aliasString, ok := ep.GetProviderSpecificProperty(providerSpecificAlias); ok {
			alias = aliasString == "true"
			if alias {
				if !slices.Contains([]string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME}, ep.RecordType) {
					ep.DeleteProviderSpecificProperty(providerSpecificAlias)
				}
			} else {
				if ep.RecordType == endpoint.RecordTypeCNAME {
					if aliasString != "false" {
						ep.SetProviderSpecificProperty(providerSpecificAlias, "false")
					}
				} else {
					ep.DeleteProviderSpecificProperty(providerSpecificAlias)
				}
			}
		} else if ep.RecordType == endpoint.RecordTypeCNAME {
			alias = useAlias(ep, p.preferCNAME)
			log.Debugf("Modifying endpoint: %v, setting %s=%v", ep, providerSpecificAlias, alias)
			ep.SetProviderSpecificProperty(providerSpecificAlias, strconv.FormatBool(alias))
		}

		if alias {
			if ep.RecordTTL.IsConfigured() {
				log.Debugf("Modifying endpoint: %v, setting ttl=%v", ep, recordTTL)
				ep.RecordTTL = recordTTL
			}
			if prop, ok := ep.GetProviderSpecificProperty(providerSpecificEvaluateTargetHealth); ok {
				if prop != "true" && prop != "false" {
					ep.SetProviderSpecificProperty(providerSpecificEvaluateTargetHealth, "false")
				}
			} else {
				ep.SetProviderSpecificProperty(providerSpecificEvaluateTargetHealth, strconv.FormatBool(p.evaluateTargetHealth))
			}

			if ep.RecordType == endpoint.RecordTypeCNAME {
				// This needs to match two records from Route53, one alias for 'A' (IPv4)
				// and one alias for 'AAAA' (IPv6).
				aliasCnameAaaaEndpoints = append(aliasCnameAaaaEndpoints, &endpoint.Endpoint{
					DNSName:          ep.DNSName,
					Targets:          ep.Targets,
					RecordType:       endpoint.RecordTypeAAAA,
					RecordTTL:        ep.RecordTTL,
					Labels:           ep.Labels,
					ProviderSpecific: ep.ProviderSpecific,
					SetIdentifier:    ep.SetIdentifier,
				})
				ep.RecordType = endpoint.RecordTypeA
			}
		} else {
			ep.DeleteProviderSpecificProperty(providerSpecificEvaluateTargetHealth)
		}
	}

	endpoints = append(endpoints, aliasCnameAaaaEndpoints...)
	return endpoints, nil
}

// newChange returns a route53 Change
// returned Change is based on the given record by the given action, e.g.
// action=ChangeActionCreate returns a change for creation of the record and
// action=ChangeActionDelete returns a change for deletion of the record.
func (p *AWSProvider) newChange(action route53types.ChangeAction, ep *endpoint.Endpoint) *Route53Change {
	change := &Route53Change{
		Change: route53types.Change{
			Action: action,
			ResourceRecordSet: &route53types.ResourceRecordSet{
				Name: aws.String(ep.DNSName),
			},
		},
	}
	change.ResourceRecordSet.Type = route53types.RRType(ep.RecordType)
	if targetHostedZone := isAWSAlias(ep); targetHostedZone != "" {
		evalTargetHealth := p.evaluateTargetHealth
		if prop, ok := ep.GetProviderSpecificProperty(providerSpecificEvaluateTargetHealth); ok {
			evalTargetHealth = prop == "true"
		}
		change.ResourceRecordSet.AliasTarget = &route53types.AliasTarget{
			DNSName:              aws.String(ep.Targets[0]),
			HostedZoneId:         aws.String(cleanZoneID(targetHostedZone)),
			EvaluateTargetHealth: evalTargetHealth,
		}
		change.sizeBytes += len([]byte(ep.Targets[0]))
		change.sizeValues += 1
	} else {
		if !ep.RecordTTL.IsConfigured() {
			change.ResourceRecordSet.TTL = aws.Int64(recordTTL)
		} else {
			change.ResourceRecordSet.TTL = aws.Int64(int64(ep.RecordTTL))
		}
		change.ResourceRecordSet.ResourceRecords = make([]route53types.ResourceRecord, len(ep.Targets))
		for idx, val := range ep.Targets {
			change.ResourceRecordSet.ResourceRecords[idx] = route53types.ResourceRecord{
				Value: aws.String(val),
			}
			change.sizeBytes += len([]byte(val))
			change.sizeValues += 1
		}
	}

	if action == route53types.ChangeActionUpsert {
		// If the value of the Action element is UPSERT, each ResourceRecord element and each character in a Value
		// element is counted twice
		change.sizeBytes *= 2
		change.sizeValues *= 2
	}

	setIdentifier := ep.SetIdentifier
	if setIdentifier != "" {
		change.ResourceRecordSet.SetIdentifier = aws.String(setIdentifier)
		if prop, ok := ep.GetProviderSpecificProperty(providerSpecificWeight); ok {
			weight, err := strconv.ParseInt(prop, 10, 64)
			if err != nil {
				log.Errorf("Failed parsing value of %s: %s: %v; using weight of 0", providerSpecificWeight, prop, err)
				weight = 0
			}
			change.ResourceRecordSet.Weight = aws.Int64(weight)
		}
		if prop, ok := ep.GetProviderSpecificProperty(providerSpecificRegion); ok {
			change.ResourceRecordSet.Region = route53types.ResourceRecordSetRegion(prop)
		}
		if prop, ok := ep.GetProviderSpecificProperty(providerSpecificFailover); ok {
			change.ResourceRecordSet.Failover = route53types.ResourceRecordSetFailover(prop)
		}
		if _, ok := ep.GetProviderSpecificProperty(providerSpecificMultiValueAnswer); ok {
			change.ResourceRecordSet.MultiValueAnswer = aws.Bool(true)
		}

		geolocation := &route53types.GeoLocation{}
		useGeolocation := false
		if prop, ok := ep.GetProviderSpecificProperty(providerSpecificGeolocationContinentCode); ok {
			geolocation.ContinentCode = aws.String(prop)
			useGeolocation = true
		} else {
			if prop, ok := ep.GetProviderSpecificProperty(providerSpecificGeolocationCountryCode); ok {
				geolocation.CountryCode = aws.String(prop)
				useGeolocation = true
			}
			if prop, ok := ep.GetProviderSpecificProperty(providerSpecificGeolocationSubdivisionCode); ok {
				geolocation.SubdivisionCode = aws.String(prop)
				useGeolocation = true
			}
		}
		if useGeolocation {
			change.ResourceRecordSet.GeoLocation = geolocation
		}
	}

	if prop, ok := ep.GetProviderSpecificProperty(providerSpecificHealthCheckID); ok {
		change.ResourceRecordSet.HealthCheckId = aws.String(prop)
	}

	if ownedRecord, ok := ep.Labels[endpoint.OwnedRecordLabelKey]; ok {
		change.OwnedRecord = ownedRecord
	}

	return change
}

// searches for `changes` that are contained in `queue` and returns the `changes` separated by whether they were found in the queue (`foundChanges`) or not (`notFoundChanges`)
func findChangesInQueue(changes Route53Changes, queue Route53Changes) (foundChanges, notFoundChanges Route53Changes) {
	if queue == nil {
		return Route53Changes{}, changes
	}

	for _, c := range changes {
		found := false
		for _, qc := range queue {
			if c == qc {
				foundChanges = append(foundChanges, c)
				found = true
				break
			}
		}
		if !found {
			notFoundChanges = append(notFoundChanges, c)
		}
	}

	return
}

// group the given changes by name and ownership relation to ensure these are always submitted in the same transaction to Route53;
// grouping by name is done to always submit changes with the same name but different set identifier together,
// grouping by ownership relation is done to always submit changes of records and e.g. their corresponding TXT registry records together
func groupChangesByNameAndOwnershipRelation(cs Route53Changes) map[string]Route53Changes {
	changesByOwnership := make(map[string]Route53Changes)
	for _, v := range cs {
		key := v.OwnedRecord
		if key == "" {
			key = *v.ResourceRecordSet.Name
		}
		changesByOwnership[key] = append(changesByOwnership[key], v)
	}
	return changesByOwnership
}

func (p *AWSProvider) tagsForZone(ctx context.Context, zoneIDs []string, profile string) (zoneTags, error) {
	client := p.clients[profile]

	result := zoneTags{}

	for i := 0; i < len(zoneIDs); i += batchSize {
		batch := zoneIDs[i:min(i+batchSize, len(zoneIDs))]
		if len(batch) == 0 {
			break
		}
		response, err := client.ListTagsForResources(ctx, &route53.ListTagsForResourcesInput{
			ResourceType: route53types.TagResourceTypeHostedzone,
			ResourceIds:  batch,
		})
		if err != nil {
			return nil, provider.NewSoftErrorf("failed to list tags for zones. %v", err)
		}

		for _, res := range response.ResourceTagSets {
			result.append(*res.ResourceId, res.Tags)
		}
	}
	return result, nil
}

// count bytes for all changes values
func countChangeBytes(cs Route53Changes) int {
	count := 0
	for _, c := range cs {
		count += c.sizeBytes
	}
	return count
}

// count total value count for all changes
func countChangeValues(cs Route53Changes) int {
	count := 0
	for _, c := range cs {
		count += c.sizeValues
	}
	return count
}

func batchChangeSet(cs Route53Changes, batchSize int, batchSizeBytes int, batchSizeValues int) []Route53Changes {
	if len(cs) <= batchSize && countChangeBytes(cs) <= batchSizeBytes && countChangeValues(cs) <= batchSizeValues {
		res := sortChangesByActionNameType(cs)
		return []Route53Changes{res}
	}

	batchChanges := make([]Route53Changes, 0)

	changesByOwnership := groupChangesByNameAndOwnershipRelation(cs)

	names := make([]string, 0)
	for v := range changesByOwnership {
		names = append(names, v)
	}
	sort.Strings(names)

	currentBatch := Route53Changes{}
	for k, name := range names {
		v := changesByOwnership[name]
		vBytes := countChangeBytes(v)
		vValues := countChangeValues(v)
		if len(v) > batchSize {
			log.Warnf("Total changes for %v exceeds max batch size of %d, total changes: %d; changes will not be performed", k, batchSize, len(v))
			continue
		}
		if vBytes > batchSizeBytes {
			log.Warnf("Total changes for %v exceeds max batch size bytes of %d, total changes bytes: %d; changes will not be performed", k, batchSizeBytes, vBytes)
			continue
		}
		if vValues > batchSizeValues {
			log.Warnf("Total changes for %v exceeds max batch size values of %d, total changes values: %d; changes will not be performed", k, batchSizeValues, vValues)
			continue
		}

		bytes := countChangeBytes(currentBatch) + vBytes
		values := countChangeValues(currentBatch) + vValues

		if len(currentBatch)+len(v) > batchSize || bytes > batchSizeBytes || values > batchSizeValues {
			// currentBatch would be too large if we add this changeset;
			// add currentBatch to batchChanges and start a new currentBatch
			batchChanges = append(batchChanges, sortChangesByActionNameType(currentBatch))
			currentBatch = append(Route53Changes{}, v...)
		} else {
			currentBatch = append(currentBatch, v...)
		}
	}
	if len(currentBatch) > 0 {
		// add final currentBatch
		batchChanges = append(batchChanges, sortChangesByActionNameType(currentBatch))
	}

	return batchChanges
}

func sortChangesByActionNameType(cs Route53Changes) Route53Changes {
	sort.SliceStable(cs, func(i, j int) bool {
		if cs[i].Action > cs[j].Action {
			return true
		}
		if cs[i].Action < cs[j].Action {
			return false
		}
		if *cs[i].ResourceRecordSet.Name < *cs[j].ResourceRecordSet.Name {
			return true
		}
		if *cs[i].ResourceRecordSet.Name > *cs[j].ResourceRecordSet.Name {
			return false
		}
		return cs[i].ResourceRecordSet.Type < cs[j].ResourceRecordSet.Type
	})

	return cs
}

// changesByZone separates a multi-zone change into a single change per zone.
func changesByZone(zones map[string]*profiledZone, changeSet Route53Changes) map[string]Route53Changes {
	changes := make(map[string]Route53Changes)

	for _, z := range zones {
		changes[*z.zone.Id] = Route53Changes{}
	}

	for _, c := range changeSet {
		hostname := provider.EnsureTrailingDot(*c.ResourceRecordSet.Name)

		zones := suitableZones(hostname, zones)
		if len(zones) == 0 {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", *c.ResourceRecordSet.Name)
			continue
		}
		for _, z := range zones {
			if c.ResourceRecordSet.AliasTarget != nil && *c.ResourceRecordSet.AliasTarget.HostedZoneId == sameZoneAlias {
				// alias record is to be created; target needs to be in the same zone as endpoint
				// if it's not, this will fail
				rrset := *c.ResourceRecordSet
				aliasTarget := *rrset.AliasTarget
				aliasTarget.HostedZoneId = aws.String(cleanZoneID(*z.zone.Id))
				rrset.AliasTarget = &aliasTarget
				c = &Route53Change{
					Change: route53types.Change{
						Action:            c.Action,
						ResourceRecordSet: &rrset,
					},
				}
			}
			changes[*z.zone.Id] = append(changes[*z.zone.Id], c)
			log.Debugf("Adding %s to zone %s [Id: %s]", hostname, *z.zone.Name, *z.zone.Id)
		}
	}

	// separating a change could lead to empty sub changes, remove them here.
	for zone, change := range changes {
		if len(change) == 0 {
			delete(changes, zone)
		}
	}

	return changes
}

// suitableZones returns all suitable private zones and the most suitable public zone
//
//	for a given hostname and a set of zones.
func suitableZones(hostname string, zones map[string]*profiledZone) []*profiledZone {
	var matchingZones []*profiledZone
	var publicZone *profiledZone

	for _, z := range zones {
		if *z.zone.Name == hostname || strings.HasSuffix(hostname, "."+*z.zone.Name) {
			if z.zone.Config == nil || !z.zone.Config.PrivateZone {
				// Only select the best matching public zone
				if publicZone == nil || len(*z.zone.Name) > len(*publicZone.zone.Name) {
					publicZone = z
				}
			} else {
				// Include all private zones
				matchingZones = append(matchingZones, z)
			}
		}
	}

	if publicZone != nil {
		matchingZones = append(matchingZones, publicZone)
	}

	return matchingZones
}

// useAlias determines if AWS ALIAS should be used.
func useAlias(ep *endpoint.Endpoint, preferCNAME bool) bool {
	if preferCNAME {
		return false
	}

	if ep.RecordType == endpoint.RecordTypeCNAME && len(ep.Targets) > 0 {
		return canonicalHostedZone(ep.Targets[0]) != ""
	}

	return false
}

// isAWSAlias determines if a given endpoint is supposed to create an AWS Alias record
// and (if so) returns the target hosted zone ID
func isAWSAlias(ep *endpoint.Endpoint) string {
	isAlias, exists := ep.GetProviderSpecificProperty(providerSpecificAlias)
	if exists && isAlias == "true" && slices.Contains([]string{endpoint.RecordTypeA, endpoint.RecordTypeAAAA}, ep.RecordType) && len(ep.Targets) > 0 {
		// alias records can only point to canonical hosted zones (e.g. to ELBs) or other records in the same zone

		if hostedZoneID, ok := ep.GetProviderSpecificProperty(providerSpecificTargetHostedZone); ok {
			// existing Endpoint where we got the target hosted zone from the Route53 data
			return hostedZoneID
		}

		// check if the target is in a canonical hosted zone
		if canonicalHostedZone := canonicalHostedZone(ep.Targets[0]); canonicalHostedZone != "" {
			return canonicalHostedZone
		}

		// if not, target needs to be in the same zone
		return sameZoneAlias
	}
	return ""
}

// canonicalHostedZone returns the matching canonical zone for a given hostname.
func canonicalHostedZone(hostname string) string {
	// strings.HasSuffix is optimized for this specific task and avoids the overhead associated with compiling and executing a regular expression.
	if strings.HasSuffix(hostname, "aws.com") || strings.HasSuffix(hostname, "aws.com.cn") || strings.HasSuffix(hostname, "tor.com") || strings.HasSuffix(hostname, "ont.com") || strings.HasSuffix(hostname, "ont.net") {
		parts := strings.Split(hostname, ".")
		// iterate from the second-last part (zone) towards the beginning
		for i := len(parts) - 2; i >= 0; i-- {
			suffix := strings.Join(parts[i:], ".")
			if zone, exists := canonicalHostedZones[suffix]; exists {
				return zone
			}
		}
	}

	if strings.HasSuffix(hostname, ".amazonaws.com") {
		// hostname is an AWS hostname, but could not find canonical hosted zone.
		// This could mean that a new region has been added but is not supported yet.
		log.Warnf("Could not find canonical hosted zone for domain %s. This may be because your region is not supported yet.", hostname)
	}

	return ""
}

// cleanZoneID removes the "/hostedzone/" prefix
func cleanZoneID(id string) string {
	return strings.TrimPrefix(id, "/hostedzone/")
}

func (p *AWSProvider) SupportedRecordType(recordType route53types.RRType) bool {
	switch recordType {
	case route53types.RRTypeMx:
		return true
	default:
		return provider.SupportedRecordType(string(recordType))
	}
}
