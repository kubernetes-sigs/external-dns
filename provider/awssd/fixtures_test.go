/*
Copyright 2025 The Kubernetes Authors.

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
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sigs.k8s.io/external-dns/endpoint"

	sd "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	sdtypes "github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
)

var (
	// Compile time checks for interface conformance
	_                    AWSSDClient = &AWSSDClientStub{}
	ErrNamespaceNotFound             = errors.New("namespace not found")
)

type AWSSDClientStub struct {
	// map[namespace_id]namespace
	namespaces map[string]*types.Namespace

	// map[namespace_id] => map[service_id]instance
	services map[string]map[string]*types.Service

	// map[service_id] => map[inst_id]instance
	instances map[string]map[string]*types.Instance

	// []inst_id
	deregistered []string
}

func (s *AWSSDClientStub) CreateService(_ context.Context, input *servicediscovery.CreateServiceInput, _ ...func(*servicediscovery.Options)) (*servicediscovery.CreateServiceOutput, error) {
	srv := &types.Service{
		Id:               input.Name,
		DnsConfig:        input.DnsConfig,
		Name:             input.Name,
		Description:      input.Description,
		CreateDate:       aws.Time(time.Now()),
		CreatorRequestId: input.CreatorRequestId,
	}

	nsServices, ok := s.services[*input.NamespaceId]
	if !ok {
		nsServices = make(map[string]*types.Service)
		s.services[*input.NamespaceId] = nsServices
	}
	nsServices[*srv.Id] = srv

	return &servicediscovery.CreateServiceOutput{
		Service: srv,
	}, nil
}

func (s *AWSSDClientStub) DeregisterInstance(_ context.Context, input *servicediscovery.DeregisterInstanceInput, _ ...func(options *servicediscovery.Options)) (*servicediscovery.DeregisterInstanceOutput, error) {
	serviceInstances := s.instances[*input.ServiceId]
	delete(serviceInstances, *input.InstanceId)
	s.deregistered = append(s.deregistered, *input.InstanceId)

	return &servicediscovery.DeregisterInstanceOutput{}, nil
}

func (s *AWSSDClientStub) GetService(_ context.Context, input *servicediscovery.GetServiceInput, _ ...func(options *servicediscovery.Options)) (*servicediscovery.GetServiceOutput, error) {
	for _, entry := range s.services {
		srv, ok := entry[*input.Id]
		if ok {
			return &servicediscovery.GetServiceOutput{
				Service: srv,
			}, nil
		}
	}

	return nil, errors.New("service not found")
}

func (s *AWSSDClientStub) DiscoverInstances(_ context.Context, input *sd.DiscoverInstancesInput, _ ...func(options *sd.Options)) (*sd.DiscoverInstancesOutput, error) {
	instances := make([]sdtypes.HttpInstanceSummary, 0)

	var foundNs bool
	for _, ns := range s.namespaces {
		if *ns.Name == *input.NamespaceName {
			foundNs = true

			for _, srv := range s.services[*ns.Id] {
				if *srv.Name == *input.ServiceName {
					for _, inst := range s.instances[*srv.Id] {
						instances = append(instances, *instanceToHTTPInstanceSummary(inst))
					}
				}
			}
		}
	}

	if !foundNs {
		return nil, ErrNamespaceNotFound
	}

	return &sd.DiscoverInstancesOutput{
		Instances: instances,
	}, nil
}

func (s *AWSSDClientStub) ListNamespaces(_ context.Context, input *sd.ListNamespacesInput, _ ...func(options *sd.Options)) (*sd.ListNamespacesOutput, error) {
	namespaces := make([]sdtypes.NamespaceSummary, 0)

	for _, ns := range s.namespaces {
		if len(input.Filters) > 0 && input.Filters[0].Name == sdtypes.NamespaceFilterNameType {
			if ns.Type != sdtypes.NamespaceType(input.Filters[0].Values[0]) {
				// skip namespaces not matching filter
				continue
			}
		}
		namespaces = append(namespaces, *namespaceToNamespaceSummary(ns))
	}

	return &sd.ListNamespacesOutput{
		Namespaces: namespaces,
	}, nil
}

func (s *AWSSDClientStub) ListServices(_ context.Context, input *sd.ListServicesInput, _ ...func(options *sd.Options)) (*sd.ListServicesOutput, error) {
	services := make([]sdtypes.ServiceSummary, 0)

	// get namespace filter
	if len(input.Filters) == 0 || input.Filters[0].Name != sdtypes.ServiceFilterNameNamespaceId {
		return nil, errors.New("missing namespace filter")
	}
	nsID := input.Filters[0].Values[0]

	for _, srv := range s.services[nsID] {
		services = append(services, *serviceToServiceSummary(srv))
	}

	return &sd.ListServicesOutput{
		Services: services,
	}, nil
}

func (s *AWSSDClientStub) RegisterInstance(ctx context.Context, input *sd.RegisterInstanceInput, _ ...func(options *sd.Options)) (*sd.RegisterInstanceOutput, error) {
	srvInstances, ok := s.instances[*input.ServiceId]
	if !ok {
		srvInstances = make(map[string]*sdtypes.Instance)
		s.instances[*input.ServiceId] = srvInstances
	}

	srvInstances[*input.InstanceId] = &sdtypes.Instance{
		Id:               input.InstanceId,
		Attributes:       input.Attributes,
		CreatorRequestId: input.CreatorRequestId,
	}

	return &sd.RegisterInstanceOutput{}, nil
}

func (s *AWSSDClientStub) UpdateService(ctx context.Context, input *sd.UpdateServiceInput, _ ...func(options *sd.Options)) (*sd.UpdateServiceOutput, error) {
	out, err := s.GetService(ctx, &sd.GetServiceInput{Id: input.Id})
	if err != nil {
		return nil, err
	}

	origSrv := out.Service
	updateSrv := input.Service

	origSrv.Description = updateSrv.Description
	origSrv.DnsConfig.DnsRecords = updateSrv.DnsConfig.DnsRecords

	return &sd.UpdateServiceOutput{}, nil
}

func (s *AWSSDClientStub) DeleteService(ctx context.Context, input *sd.DeleteServiceInput, _ ...func(options *sd.Options)) (*sd.DeleteServiceOutput, error) {
	out, err := s.GetService(ctx, &sd.GetServiceInput{Id: input.Id})
	if err != nil {
		return nil, err
	}

	service := out.Service
	namespace := s.services[*service.NamespaceId]
	delete(namespace, *input.Id)

	return &sd.DeleteServiceOutput{}, nil
}

func newTestAWSSDProvider(api AWSSDClient, domainFilter *endpoint.DomainFilter, namespaceTypeFilter, ownerID string) *AWSSDProvider {
	return &AWSSDProvider{
		client:              api,
		dryRun:              false,
		namespaceFilter:     domainFilter,
		namespaceTypeFilter: newSdNamespaceFilter(namespaceTypeFilter),
		cleanEmptyService:   true,
		ownerID:             ownerID,
	}
}

func instanceToHTTPInstanceSummary(instance *sdtypes.Instance) *sdtypes.HttpInstanceSummary {
	if instance == nil {
		return nil
	}

	return &sdtypes.HttpInstanceSummary{
		InstanceId: instance.Id,
		Attributes: instance.Attributes,
	}
}

func namespaceToNamespaceSummary(namespace *sdtypes.Namespace) *sdtypes.NamespaceSummary {
	if namespace == nil {
		return nil
	}

	return &sdtypes.NamespaceSummary{
		Id:   namespace.Id,
		Type: namespace.Type,
		Name: namespace.Name,
		Arn:  namespace.Arn,
	}
}

func serviceToServiceSummary(service *sdtypes.Service) *sdtypes.ServiceSummary {
	if service == nil {
		return nil
	}

	return &sdtypes.ServiceSummary{
		Arn:                     service.Arn,
		CreateDate:              service.CreateDate,
		Description:             service.Description,
		DnsConfig:               service.DnsConfig,
		HealthCheckConfig:       service.HealthCheckConfig,
		HealthCheckCustomConfig: service.HealthCheckCustomConfig,
		Id:                      service.Id,
		InstanceCount:           service.InstanceCount,
		Name:                    service.Name,
		Type:                    service.Type,
	}
}

func testHelperAWSSDServicesMapsEqual(t *testing.T, expected map[string]*sdtypes.Service, services map[string]*sdtypes.Service) {
	require.Len(t, services, len(expected))

	for _, srv := range services {
		testHelperAWSSDServicesEqual(t, expected[*srv.Name], srv)
	}
}

func testHelperAWSSDServicesEqual(t *testing.T, expected *sdtypes.Service, srv *sdtypes.Service) {
	assert.Equal(t, *expected.Description, *srv.Description)
	assert.Equal(t, *expected.Name, *srv.Name)
	assert.True(t, reflect.DeepEqual(*expected.DnsConfig, *srv.DnsConfig))
}
