/*
Copyright 2018 The Kubernetes Authors.

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

package awssd

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	sd "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	sdtypes "github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	sdDefaultRecordTTL = 300

	sdNamespaceTypePublic  = "public"
	sdNamespaceTypePrivate = "private"

	sdInstanceAttrIPV4  = "AWS_INSTANCE_IPV4"
	sdInstanceAttrIPV6  = "AWS_INSTANCE_IPV6"
	sdInstanceAttrCname = "AWS_INSTANCE_CNAME"
	sdInstanceAttrAlias = "AWS_ALIAS_DNS_NAME"
)

var (
	// matches ELB with hostname format load-balancer.us-east-1.elb.amazonaws.com
	sdElbHostnameRegex = regexp.MustCompile(`.+\.[^.]+\.elb\.amazonaws\.com$`)

	// matches NLB with hostname format load-balancer.elb.us-east-1.amazonaws.com
	sdNlbHostnameRegex = regexp.MustCompile(`.+\.elb\.[^.]+\.amazonaws\.com$`)
)

// AWSSDClient is the subset of the AWS Cloud Map API that we actually use. Add methods as required.
// Signatures must match exactly. Taken from https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/servicediscovery
type AWSSDClient interface {
	CreateService(ctx context.Context, params *sd.CreateServiceInput, optFns ...func(*sd.Options)) (*sd.CreateServiceOutput, error)
	DeregisterInstance(ctx context.Context, params *sd.DeregisterInstanceInput, optFns ...func(*sd.Options)) (*sd.DeregisterInstanceOutput, error)
	DiscoverInstances(ctx context.Context, params *sd.DiscoverInstancesInput, optFns ...func(*sd.Options)) (*sd.DiscoverInstancesOutput, error)
	ListNamespaces(ctx context.Context, params *sd.ListNamespacesInput, optFns ...func(*sd.Options)) (*sd.ListNamespacesOutput, error)
	ListServices(ctx context.Context, params *sd.ListServicesInput, optFns ...func(*sd.Options)) (*sd.ListServicesOutput, error)
	RegisterInstance(ctx context.Context, params *sd.RegisterInstanceInput, optFns ...func(*sd.Options)) (*sd.RegisterInstanceOutput, error)
	UpdateService(ctx context.Context, params *sd.UpdateServiceInput, optFns ...func(*sd.Options)) (*sd.UpdateServiceOutput, error)
	DeleteService(ctx context.Context, params *sd.DeleteServiceInput, optFns ...func(*sd.Options)) (*sd.DeleteServiceOutput, error)
}

// AWSSDProvider is an implementation of Provider for AWS Cloud Map.
type AWSSDProvider struct {
	provider.BaseProvider
	client AWSSDClient
	dryRun bool
	// only consider namespaces ending in this suffix
	namespaceFilter endpoint.DomainFilter
	// filter namespace by type (private or public)
	namespaceTypeFilter sdtypes.NamespaceFilter
	// enables service without instances cleanup
	cleanEmptyService bool
	// filter services for removal
	ownerID string
	// tags to be added to the service
	tags []sdtypes.Tag
}

// NewAWSSDProvider initializes a new AWS Cloud Map based Provider.
func NewAWSSDProvider(domainFilter endpoint.DomainFilter, namespaceType string, dryRun, cleanEmptyService bool, ownerID string, tags map[string]string, client AWSSDClient) (*AWSSDProvider, error) {
	p := &AWSSDProvider{
		client:              client,
		dryRun:              dryRun,
		namespaceFilter:     domainFilter,
		namespaceTypeFilter: newSdNamespaceFilter(namespaceType),
		cleanEmptyService:   cleanEmptyService,
		ownerID:             ownerID,
		tags:                awsTags(tags),
	}

	return p, nil
}

// newSdNamespaceFilter initialized AWS SD Namespace Filter based on given string config
func newSdNamespaceFilter(namespaceTypeConfig string) sdtypes.NamespaceFilter {
	switch namespaceTypeConfig {
	case sdNamespaceTypePublic:
		return sdtypes.NamespaceFilter{
			Name:   sdtypes.NamespaceFilterNameType,
			Values: []string{string(sdtypes.NamespaceTypeDnsPublic)},
		}
	case sdNamespaceTypePrivate:
		return sdtypes.NamespaceFilter{
			Name:   sdtypes.NamespaceFilterNameType,
			Values: []string{string(sdtypes.NamespaceTypeDnsPrivate)},
		}
	default:
		return sdtypes.NamespaceFilter{}
	}
}

// awsTags converts user supplied tags to AWS format
func awsTags(tags map[string]string) []sdtypes.Tag {
	awsTags := make([]sdtypes.Tag, 0, len(tags))
	for k, v := range tags {
		awsTags = append(awsTags, sdtypes.Tag{Key: aws.String(k), Value: aws.String(v)})
	}
	return awsTags
}

// Records returns list of all endpoints.
func (p *AWSSDProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, err error) {
	namespaces, err := p.ListNamespaces(ctx)
	if err != nil {
		return nil, err
	}

	for _, ns := range namespaces {
		services, err := p.ListServicesByNamespaceID(ctx, ns.Id)
		if err != nil {
			return nil, err
		}

		for _, srv := range services {
			resp, err := p.client.DiscoverInstances(ctx, &sd.DiscoverInstancesInput{
				NamespaceName: ns.Name,
				ServiceName:   srv.Name,
			})
			if err != nil {
				return nil, err
			}

			if len(resp.Instances) == 0 {
				if err := p.DeleteService(ctx, srv); err != nil {
					log.Errorf("Failed to delete service %q, error: %s", *srv.Name, err)
				}
				continue
			}

			endpoints = append(endpoints, p.instancesToEndpoint(ns, srv, resp.Instances))
		}
	}

	return endpoints, nil
}

func (p *AWSSDProvider) instancesToEndpoint(ns *sdtypes.NamespaceSummary, srv *sdtypes.Service, instances []sdtypes.HttpInstanceSummary) *endpoint.Endpoint {
	// DNS name of the record is a concatenation of service and namespace
	recordName := *srv.Name + "." + *ns.Name

	labels := endpoint.NewLabels()
	labels[endpoint.AWSSDDescriptionLabel] = *srv.Description

	newEndpoint := &endpoint.Endpoint{
		DNSName:   recordName,
		RecordTTL: endpoint.TTL(*srv.DnsConfig.DnsRecords[0].TTL),
		Targets:   make(endpoint.Targets, 0, len(instances)),
		Labels:    labels,
	}

	for _, inst := range instances {
		// CNAME
		if inst.Attributes[sdInstanceAttrCname] != "" && srv.DnsConfig.DnsRecords[0].Type == sdtypes.RecordTypeCname {
			newEndpoint.RecordType = endpoint.RecordTypeCNAME
			newEndpoint.Targets = append(newEndpoint.Targets, inst.Attributes[sdInstanceAttrCname])

			// ALIAS
		} else if inst.Attributes[sdInstanceAttrAlias] != "" {
			newEndpoint.RecordType = endpoint.RecordTypeCNAME
			newEndpoint.Targets = append(newEndpoint.Targets, inst.Attributes[sdInstanceAttrAlias])

			// IPv4-based target
		} else if inst.Attributes[sdInstanceAttrIPV4] != "" {
			newEndpoint.RecordType = endpoint.RecordTypeA
			newEndpoint.Targets = append(newEndpoint.Targets, inst.Attributes[sdInstanceAttrIPV4])

			// IPv6-based target
		} else if inst.Attributes[sdInstanceAttrIPV6] != "" {
			newEndpoint.RecordType = endpoint.RecordTypeAAAA
			newEndpoint.Targets = append(newEndpoint.Targets, inst.Attributes[sdInstanceAttrIPV6])
		} else {
			log.Warnf("Invalid instance \"%v\" found in service \"%v\"", inst, srv.Name)
		}
	}

	return newEndpoint
}

// ApplyChanges applies Kubernetes changes in endpoints to AWS API
func (p *AWSSDProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// return early if there is nothing to change
	if len(changes.Create) == 0 && len(changes.Delete) == 0 && len(changes.UpdateNew) == 0 {
		log.Info("All records are already up to date")
		return nil
	}

	// convert updates to delete and create operation if applicable (updates not supported)
	creates, deletes := p.updatesToCreates(changes)
	changes.Delete = append(changes.Delete, deletes...)
	changes.Create = append(changes.Create, creates...)

	namespaces, err := p.ListNamespaces(ctx)
	if err != nil {
		return err
	}

	// Deletes must be executed first to support update case.
	// When just list of targets is updated `[1.2.3.4] -> [1.2.3.4, 1.2.3.5]` it is translated to:
	// ```
	// deletes = [1.2.3.4]
	// creates = [1.2.3.4, 1.2.3.5]
	// ```
	// then when deletes are executed after creates it will miss the `1.2.3.4` instance.
	err = p.submitDeletes(ctx, namespaces, changes.Delete)
	if err != nil {
		return err
	}

	err = p.submitCreates(ctx, namespaces, changes.Create)
	if err != nil {
		return err
	}

	return nil
}

func (p *AWSSDProvider) updatesToCreates(changes *plan.Changes) (creates []*endpoint.Endpoint, deletes []*endpoint.Endpoint) {
	updateNewMap := map[string]*endpoint.Endpoint{}
	for _, e := range changes.UpdateNew {
		updateNewMap[e.DNSName] = e
	}

	for _, old := range changes.UpdateOld {
		current := updateNewMap[old.DNSName]

		if !old.Targets.Same(current.Targets) {
			// when targets differ the old instances need to be de-registered first
			deletes = append(deletes, old)
		}

		// always register (or re-register) instance with the current data
		creates = append(creates, current)
	}

	return creates, deletes
}

func (p *AWSSDProvider) submitCreates(ctx context.Context, namespaces []*sdtypes.NamespaceSummary, changes []*endpoint.Endpoint) error {
	changesByNamespaceID := p.changesByNamespaceID(namespaces, changes)

	for nsID, changeList := range changesByNamespaceID {
		services, err := p.ListServicesByNamespaceID(ctx, aws.String(nsID))
		if err != nil {
			return err
		}

		for _, ch := range changeList {
			_, srvName := p.parseHostname(ch.DNSName)

			srv := services[srvName]
			if srv == nil {
				// when service is missing create a new one
				srv, err = p.CreateService(ctx, &nsID, &srvName, ch)
				if err != nil {
					return err
				}
				// update local list of services
				services[*srv.Name] = srv
			} else if ch.RecordTTL.IsConfigured() && *srv.DnsConfig.DnsRecords[0].TTL != int64(ch.RecordTTL) {
				// update service when TTL differ
				err = p.UpdateService(ctx, srv, ch)
				if err != nil {
					return err
				}
			}

			err = p.RegisterInstance(ctx, srv, ch)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *AWSSDProvider) submitDeletes(ctx context.Context, namespaces []*sdtypes.NamespaceSummary, changes []*endpoint.Endpoint) error {
	changesByNamespaceID := p.changesByNamespaceID(namespaces, changes)

	for nsID, changeList := range changesByNamespaceID {
		services, err := p.ListServicesByNamespaceID(ctx, aws.String(nsID))
		if err != nil {
			return err
		}

		for _, ch := range changeList {
			hostname := ch.DNSName
			_, srvName := p.parseHostname(hostname)

			srv := services[srvName]
			if srv == nil {
				return fmt.Errorf("service \"%s\" is missing when trying to delete \"%v\"", srvName, hostname)
			}

			err := p.DeregisterInstance(ctx, srv, ch)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ListNamespaces returns all namespaces matching defined namespace filter
func (p *AWSSDProvider) ListNamespaces(ctx context.Context) ([]*sdtypes.NamespaceSummary, error) {
	namespaces := make([]*sdtypes.NamespaceSummary, 0)

	paginator := sd.NewListNamespacesPaginator(p.client, &sd.ListNamespacesInput{
		Filters: []sdtypes.NamespaceFilter{p.namespaceTypeFilter},
	})
	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, ns := range resp.Namespaces {
			if !p.namespaceFilter.Match(*ns.Name) {
				continue
			}
			namespaces = append(namespaces, &ns)
		}
	}

	return namespaces, nil
}

// ListServicesByNamespaceID returns list of services in given namespace.
func (p *AWSSDProvider) ListServicesByNamespaceID(ctx context.Context, namespaceID *string) (map[string]*sdtypes.Service, error) {
	services := make([]sdtypes.ServiceSummary, 0)

	paginator := sd.NewListServicesPaginator(p.client, &sd.ListServicesInput{
		Filters: []sdtypes.ServiceFilter{{
			Name:   sdtypes.ServiceFilterNameNamespaceId,
			Values: []string{*namespaceID},
		}},
		MaxResults: aws.Int32(100),
	})
	for paginator.HasMorePages() {
		resp, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		services = append(services, resp.Services...)
	}

	servicesMap := make(map[string]*sdtypes.Service)
	for _, serviceSummary := range services {
		service := &sdtypes.Service{
			Arn:                     serviceSummary.Arn,
			CreateDate:              serviceSummary.CreateDate,
			Description:             serviceSummary.Description,
			DnsConfig:               serviceSummary.DnsConfig,
			HealthCheckConfig:       serviceSummary.HealthCheckConfig,
			HealthCheckCustomConfig: serviceSummary.HealthCheckCustomConfig,
			Id:                      serviceSummary.Id,
			InstanceCount:           serviceSummary.InstanceCount,
			Name:                    serviceSummary.Name,
			NamespaceId:             namespaceID,
			Type:                    serviceSummary.Type,
		}

		servicesMap[*service.Name] = service
	}
	return servicesMap, nil
}

// CreateService creates a new service in AWS API. Returns the created service.
func (p *AWSSDProvider) CreateService(ctx context.Context, namespaceID *string, srvName *string, ep *endpoint.Endpoint) (*sdtypes.Service, error) {
	log.Infof("Creating a new service \"%s\" in \"%s\" namespace", *srvName, *namespaceID)

	srvType := p.serviceTypeFromEndpoint(ep)
	routingPolicy := p.routingPolicyFromEndpoint(ep)

	ttl := int64(sdDefaultRecordTTL)
	if ep.RecordTTL.IsConfigured() {
		ttl = int64(ep.RecordTTL)
	}

	if !p.dryRun {
		out, err := p.client.CreateService(ctx, &sd.CreateServiceInput{
			Name:        srvName,
			Description: aws.String(ep.Labels[endpoint.AWSSDDescriptionLabel]),
			DnsConfig: &sdtypes.DnsConfig{
				RoutingPolicy: routingPolicy,
				DnsRecords: []sdtypes.DnsRecord{{
					Type: srvType,
					TTL:  aws.Int64(ttl),
				}},
			},
			NamespaceId: namespaceID,
			Tags:        p.tags,
		})
		if err != nil {
			return nil, err
		}

		return out.Service, nil
	}

	// return mock service summary in case of dry run
	return &sdtypes.Service{Id: aws.String("dry-run-service"), Name: aws.String("dry-run-service")}, nil
}

// UpdateService updates the specified service with information from provided endpoint.
func (p *AWSSDProvider) UpdateService(ctx context.Context, service *sdtypes.Service, ep *endpoint.Endpoint) error {
	log.Infof("Updating service \"%s\"", *service.Name)

	srvType := p.serviceTypeFromEndpoint(ep)

	ttl := int64(sdDefaultRecordTTL)
	if ep.RecordTTL.IsConfigured() {
		ttl = int64(ep.RecordTTL)
	}

	if !p.dryRun {
		_, err := p.client.UpdateService(ctx, &sd.UpdateServiceInput{
			Id: service.Id,
			Service: &sdtypes.ServiceChange{
				Description: aws.String(ep.Labels[endpoint.AWSSDDescriptionLabel]),
				DnsConfig: &sdtypes.DnsConfigChange{
					DnsRecords: []sdtypes.DnsRecord{{
						Type: srvType,
						TTL:  aws.Int64(ttl),
					}},
				},
			},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteService deletes empty Service from AWS API if its owner id match
func (p *AWSSDProvider) DeleteService(ctx context.Context, service *sdtypes.Service) error {
	log.Debugf("Check if service \"%s\" owner id match and it can be deleted", *service.Name)
	if !p.dryRun && p.cleanEmptyService {
		// convert ownerID string to service description format
		label := endpoint.NewLabels()
		label[endpoint.OwnerLabelKey] = p.ownerID
		label[endpoint.AWSSDDescriptionLabel] = label.SerializePlain(false)

		if strings.HasPrefix(*service.Description, label[endpoint.AWSSDDescriptionLabel]) {
			log.Infof("Deleting service \"%s\"", *service.Name)
			_, err := p.client.DeleteService(ctx, &sd.DeleteServiceInput{
				Id: aws.String(*service.Id),
			})
			return err
		}
		log.Debugf("Skipping service removal %s because owner id does not match, found: \"%s\", required: \"%s\"", *service.Name, *service.Description, label[endpoint.AWSSDDescriptionLabel])
	}
	return nil
}

// RegisterInstance creates a new instance in given service.
func (p *AWSSDProvider) RegisterInstance(ctx context.Context, service *sdtypes.Service, ep *endpoint.Endpoint) error {
	for _, target := range ep.Targets {
		log.Infof("Registering a new instance \"%s\" for service \"%s\" (%s)", target, *service.Name, *service.Id)

		attr := make(map[string]string)

		switch ep.RecordType {
		case endpoint.RecordTypeCNAME:
			if p.isAWSLoadBalancer(target) {
				attr[sdInstanceAttrAlias] = target
			} else {
				attr[sdInstanceAttrCname] = target
			}
		case endpoint.RecordTypeA:
			attr[sdInstanceAttrIPV4] = target
		case endpoint.RecordTypeAAAA:
			attr[sdInstanceAttrIPV6] = target
		default:
			return fmt.Errorf("invalid endpoint type (%v)", ep)
		}

		if !p.dryRun {
			_, err := p.client.RegisterInstance(ctx, &sd.RegisterInstanceInput{
				ServiceId:  service.Id,
				Attributes: attr,
				InstanceId: aws.String(p.targetToInstanceID(target)),
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// DeregisterInstance removes an instance from given service.
func (p *AWSSDProvider) DeregisterInstance(ctx context.Context, service *sdtypes.Service, ep *endpoint.Endpoint) error {
	for _, target := range ep.Targets {
		log.Infof("De-registering an instance \"%s\" for service \"%s\" (%s)", target, *service.Name, *service.Id)

		if !p.dryRun {
			_, err := p.client.DeregisterInstance(ctx, &sd.DeregisterInstanceInput{
				InstanceId: aws.String(p.targetToInstanceID(target)),
				ServiceId:  service.Id,
			})
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Instance ID length is limited by AWS API to 64 characters. For longer strings SHA-256 hash will be used instead of
// the verbatim target to limit the length.
func (p *AWSSDProvider) targetToInstanceID(target string) string {
	if len(target) > 64 {
		hash := sha256.Sum256([]byte(strings.ToLower(target)))
		return hex.EncodeToString(hash[:])
	}

	return strings.ToLower(target)
}

func (p *AWSSDProvider) changesByNamespaceID(namespaces []*sdtypes.NamespaceSummary, changes []*endpoint.Endpoint) map[string][]*endpoint.Endpoint {
	changesByNsID := make(map[string][]*endpoint.Endpoint)

	for _, ns := range namespaces {
		changesByNsID[*ns.Id] = []*endpoint.Endpoint{}
	}

	for _, c := range changes {
		// trim the trailing dot from hostname if any
		hostname := strings.TrimSuffix(c.DNSName, ".")
		nsName, _ := p.parseHostname(hostname)

		matchingNamespaces := matchingNamespaces(nsName, namespaces)
		if len(matchingNamespaces) == 0 {
			log.Warnf("Skipping record %s because no namespace matching record DNS Name was detected ", c.String())
			continue
		}
		for _, ns := range matchingNamespaces {
			changesByNsID[*ns.Id] = append(changesByNsID[*ns.Id], c)
		}
	}

	// separating a change could lead to empty sub changes, remove them here.
	for zone, change := range changesByNsID {
		if len(change) == 0 {
			delete(changesByNsID, zone)
		}
	}

	return changesByNsID
}

// returns list of all namespaces matching given hostname
func matchingNamespaces(hostname string, namespaces []*sdtypes.NamespaceSummary) []*sdtypes.NamespaceSummary {
	matchingNamespaces := make([]*sdtypes.NamespaceSummary, 0)

	for _, ns := range namespaces {
		if *ns.Name == hostname {
			matchingNamespaces = append(matchingNamespaces, ns)
		}
	}

	return matchingNamespaces
}

// parse hostname to namespace (domain) and service
func (p *AWSSDProvider) parseHostname(hostname string) (namespace string, service string) {
	parts := strings.Split(hostname, ".")
	service = parts[0]
	namespace = strings.Join(parts[1:], ".")
	return
}

// determine service routing policy based on endpoint type
func (p *AWSSDProvider) routingPolicyFromEndpoint(ep *endpoint.Endpoint) sdtypes.RoutingPolicy {
	if ep.RecordType == endpoint.RecordTypeA || ep.RecordType == endpoint.RecordTypeAAAA {
		return sdtypes.RoutingPolicyMultivalue
	}

	return sdtypes.RoutingPolicyWeighted
}

// determine service type (A, AAAA, CNAME) from given endpoint
func (p *AWSSDProvider) serviceTypeFromEndpoint(ep *endpoint.Endpoint) sdtypes.RecordType {
	switch ep.RecordType {
	case endpoint.RecordTypeCNAME:
		// FIXME service type is derived from the first target only. Theoretically this may be problem.
		// But I don't see a scenario where one endpoint contains targets of different types.
		if p.isAWSLoadBalancer(ep.Targets[0]) {
			// ALIAS target uses DNS record of type A
			return sdtypes.RecordTypeA
		}
		return sdtypes.RecordTypeCname
	case endpoint.RecordTypeAAAA:
		return sdtypes.RecordTypeAaaa
	default:
		return sdtypes.RecordTypeA
	}
}

// determine if a given hostname belongs to an AWS load balancer
func (p *AWSSDProvider) isAWSLoadBalancer(hostname string) bool {
	matchElb := sdElbHostnameRegex.MatchString(hostname)
	matchNlb := sdNlbHostnameRegex.MatchString(hostname)

	return matchElb || matchNlb
}
