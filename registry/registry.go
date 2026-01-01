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

package registry

import (
	"context"
	"fmt"

	aws_dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/aws"
	"sigs.k8s.io/external-dns/registry/awssd"
	"sigs.k8s.io/external-dns/registry/dynamodb"
	"sigs.k8s.io/external-dns/registry/noop"
	"sigs.k8s.io/external-dns/registry/txt"
)

const (
	DYNAMODB = "dynamodb"
	NOOP     = "noop"
	TXT      = "txt"
	AWSSD    = "aws-sd"
)

// Registry is an interface which should enables ownership concept in external-dns
// Records() returns ALL records registered with DNS provider
// each entry includes owner information
// ApplyChanges(changes *plan.Changes) propagates the changes to the DNS Provider API and correspondingly updates ownership depending on type of registry being used
type Registry interface {
	Records(ctx context.Context) ([]*endpoint.Endpoint, error)
	ApplyChanges(ctx context.Context, changes *plan.Changes) error
	AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error)
	GetDomainFilter() endpoint.DomainFilterInterface
	OwnerID() string
}

// SelectRegistry selects the appropriate registry implementation based on the configuration in cfg.
// It initializes and returns a registry along with any error encountered during setup.
// Supported registry types include: dynamodb, noop, txt, and aws-sd.
func SelectRegistry(cfg *externaldns.Config, p provider.Provider) (Registry, error) {
	var r Registry
	var err error
	switch cfg.Registry {
	case DYNAMODB:
		r, err = dynamodb.NewDynamoDBRegistry(
			p, cfg.TXTOwnerID, aws_dynamodb.NewFromConfig(aws.CreateDefaultV2Config(cfg), dynamodb.WithRegion(cfg.AWSDynamoDBRegion)),
			cfg.AWSDynamoDBTable, cfg.TXTPrefix, cfg.TXTSuffix, cfg.TXTWildcardReplacement, cfg.ManagedDNSRecordTypes,
			cfg.ExcludeDNSRecordTypes, []byte(cfg.TXTEncryptAESKey), cfg.TXTCacheInterval)
	case NOOP:
		r, err = noop.NewNoopRegistry(p)
	case TXT:
		r, err = txt.NewTXTRegistry(
			p, cfg.TXTPrefix, cfg.TXTSuffix, cfg.TXTOwnerID,
			cfg.TXTCacheInterval, cfg.TXTWildcardReplacement,
			cfg.ManagedDNSRecordTypes, cfg.ExcludeDNSRecordTypes,
			cfg.TXTEncryptEnabled, []byte(cfg.TXTEncryptAESKey), cfg.TXTOwnerOld)
	case AWSSD:
		r, err = awssd.NewAWSSDRegistry(p, cfg.TXTOwnerID)
	default:
		err = fmt.Errorf("unknown registry: %s", cfg.Registry)
	}
	return r, err
}
