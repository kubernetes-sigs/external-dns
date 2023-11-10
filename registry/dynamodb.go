/*
Copyright 2023 The Kubernetes Authors.

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
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/sets"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// DynamoDBAPI is the subset of the AWS Route53 API that we actually use.  Add methods as required. Signatures must match exactly.
type DynamoDBAPI interface {
	DescribeTableWithContext(ctx aws.Context, input *dynamodb.DescribeTableInput, opts ...request.Option) (*dynamodb.DescribeTableOutput, error)
	ScanPagesWithContext(ctx aws.Context, input *dynamodb.ScanInput, fn func(*dynamodb.ScanOutput, bool) bool, opts ...request.Option) error
	BatchExecuteStatementWithContext(aws.Context, *dynamodb.BatchExecuteStatementInput, ...request.Option) (*dynamodb.BatchExecuteStatementOutput, error)
}

// DynamoDBRegistry implements registry interface with ownership implemented via an AWS DynamoDB table.
type DynamoDBRegistry struct {
	provider provider.Provider
	ownerID  string // refers to the owner id of the current instance

	dynamodbAPI DynamoDBAPI
	table       string

	// For migration from TXT registry
	mapper              nameMapper
	wildcardReplacement string
	managedRecordTypes  []string
	excludeRecordTypes  []string
	txtEncryptAESKey    []byte

	// cache the dynamodb records owned by us.
	labels         map[endpoint.EndpointKey]endpoint.Labels
	orphanedLabels sets.Set[endpoint.EndpointKey]

	// cache the records in memory and update on an interval instead.
	recordsCache            []*endpoint.Endpoint
	recordsCacheRefreshTime time.Time
	cacheInterval           time.Duration
}

const dynamodbAttributeMigrate = "dynamodb/needs-migration"

// DynamoDB allows a maximum batch size of 25 items.
var dynamodbMaxBatchSize uint8 = 25

// NewDynamoDBRegistry returns a new DynamoDBRegistry object.
func NewDynamoDBRegistry(provider provider.Provider, ownerID string, dynamodbAPI DynamoDBAPI, table string, txtPrefix, txtSuffix, txtWildcardReplacement string, managedRecordTypes, excludeRecordTypes []string, txtEncryptAESKey []byte, cacheInterval time.Duration) (*DynamoDBRegistry, error) {
	if ownerID == "" {
		return nil, errors.New("owner id cannot be empty")
	}
	if table == "" {
		return nil, errors.New("table cannot be empty")
	}

	if len(txtEncryptAESKey) == 0 {
		txtEncryptAESKey = nil
	} else if len(txtEncryptAESKey) != 32 {
		return nil, errors.New("the AES Encryption key must have a length of 32 bytes")
	}
	if len(txtPrefix) > 0 && len(txtSuffix) > 0 {
		return nil, errors.New("txt-prefix and txt-suffix are mutually exclusive")
	}

	mapper := newaffixNameMapper(txtPrefix, txtSuffix, txtWildcardReplacement)

	return &DynamoDBRegistry{
		provider:            provider,
		ownerID:             ownerID,
		dynamodbAPI:         dynamodbAPI,
		table:               table,
		mapper:              mapper,
		wildcardReplacement: txtWildcardReplacement,
		managedRecordTypes:  managedRecordTypes,
		excludeRecordTypes:  excludeRecordTypes,
		txtEncryptAESKey:    txtEncryptAESKey,
		cacheInterval:       cacheInterval,
	}, nil
}

func (im *DynamoDBRegistry) GetDomainFilter() endpoint.DomainFilter {
	return im.provider.GetDomainFilter()
}

func (im *DynamoDBRegistry) OwnerID() string {
	return im.ownerID
}

// Records returns the current records from the registry.
func (im *DynamoDBRegistry) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	// If we have the zones cached AND we have refreshed the cache since the
	// last given interval, then just use the cached results.
	if im.recordsCache != nil && time.Since(im.recordsCacheRefreshTime) < im.cacheInterval {
		log.Debug("Using cached records.")
		return im.recordsCache, nil
	}

	if im.labels == nil {
		if err := im.readLabels(ctx); err != nil {
			return nil, err
		}
	}

	records, err := im.provider.Records(ctx)
	if err != nil {
		return nil, err
	}

	orphanedLabels := sets.KeySet(im.labels)
	endpoints := make([]*endpoint.Endpoint, 0, len(records))
	labelMap := map[endpoint.EndpointKey]endpoint.Labels{}
	txtRecordsMap := map[endpoint.EndpointKey]*endpoint.Endpoint{}
	for _, record := range records {
		key := record.Key()
		if labels := im.labels[key]; labels != nil {
			record.Labels = labels
			orphanedLabels.Delete(key)
		} else {
			record.Labels = endpoint.NewLabels()

			if record.RecordType == endpoint.RecordTypeTXT {
				// We simply assume that TXT records for the TXT registry will always have only one target.
				if labels, err := endpoint.NewLabelsFromString(record.Targets[0], im.txtEncryptAESKey); err == nil {
					endpointName, recordType := im.mapper.toEndpointName(record.DNSName)
					key := endpoint.EndpointKey{
						DNSName:       endpointName,
						SetIdentifier: record.SetIdentifier,
					}
					if recordType == endpoint.RecordTypeAAAA {
						key.RecordType = recordType
					}
					labelMap[key] = labels
					txtRecordsMap[key] = record
					continue
				}
			}
		}

		endpoints = append(endpoints, record)
	}

	im.orphanedLabels = orphanedLabels

	// Migrate label data from TXT registry.
	if len(labelMap) > 0 {
		for _, ep := range endpoints {
			if _, ok := im.labels[ep.Key()]; ok {
				continue
			}

			dnsNameSplit := strings.Split(ep.DNSName, ".")
			// If specified, replace a leading asterisk in the generated txt record name with some other string
			if im.wildcardReplacement != "" && dnsNameSplit[0] == "*" {
				dnsNameSplit[0] = im.wildcardReplacement
			}
			dnsName := strings.Join(dnsNameSplit, ".")
			key := endpoint.EndpointKey{
				DNSName:       dnsName,
				SetIdentifier: ep.SetIdentifier,
			}
			if ep.RecordType == endpoint.RecordTypeAAAA {
				key.RecordType = ep.RecordType
			}
			if labels, ok := labelMap[key]; ok {
				for k, v := range labels {
					ep.Labels[k] = v
				}
				ep.SetProviderSpecificProperty(dynamodbAttributeMigrate, "true")
				delete(txtRecordsMap, key)
			}
		}
	}

	// Remove any unused TXT ownership records owned by us
	if len(txtRecordsMap) > 0 && !plan.IsManagedRecord(endpoint.RecordTypeTXT, im.managedRecordTypes, im.excludeRecordTypes) {
		log.Infof("Old TXT ownership records will not be deleted because \"TXT\" is not in the set of managed record types.")
	}
	for _, record := range txtRecordsMap {
		record.Labels[endpoint.OwnerLabelKey] = im.ownerID
		endpoints = append(endpoints, record)
	}

	// Update the cache.
	if im.cacheInterval > 0 {
		im.recordsCache = endpoints
		im.recordsCacheRefreshTime = time.Now()
	}

	return endpoints, nil
}

// ApplyChanges updates the DNS provider and DynamoDB table with the changes.
func (im *DynamoDBRegistry) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	filteredChanges := &plan.Changes{
		Create:    changes.Create,
		UpdateNew: endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.UpdateNew),
		UpdateOld: endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.UpdateOld),
		Delete:    endpoint.FilterEndpointsByOwnerID(im.ownerID, changes.Delete),
	}

	statements := make([]*dynamodb.BatchStatementRequest, 0, len(filteredChanges.Create)+len(filteredChanges.UpdateNew))
	for _, r := range filteredChanges.Create {
		if r.Labels == nil {
			r.Labels = make(map[string]string)
		}
		r.Labels[endpoint.OwnerLabelKey] = im.ownerID

		key := r.Key()
		oldLabels := im.labels[key]
		if oldLabels == nil {
			statements = im.appendInsert(statements, key, r.Labels)
		} else {
			im.orphanedLabels.Delete(key)
			statements = im.appendUpdate(statements, key, oldLabels, r.Labels)
		}

		im.labels[key] = r.Labels
		if im.cacheInterval > 0 {
			im.addToCache(r)
		}
	}

	for _, r := range filteredChanges.Delete {
		delete(im.labels, r.Key())
		if im.cacheInterval > 0 {
			im.removeFromCache(r)
		}
	}

	oldLabels := make(map[endpoint.EndpointKey]endpoint.Labels, len(filteredChanges.UpdateOld))
	needMigration := map[endpoint.EndpointKey]bool{}
	for _, r := range filteredChanges.UpdateOld {
		oldLabels[r.Key()] = r.Labels

		if _, ok := r.GetProviderSpecificProperty(dynamodbAttributeMigrate); ok {
			needMigration[r.Key()] = true
		}

		// remove old version of record from cache
		if im.cacheInterval > 0 {
			im.removeFromCache(r)
		}
	}

	for _, r := range filteredChanges.UpdateNew {
		key := r.Key()
		if needMigration[key] {
			statements = im.appendInsert(statements, key, r.Labels)
			// Invalidate the records cache so the next sync deletes the TXT ownership record
			im.recordsCache = nil
		} else {
			statements = im.appendUpdate(statements, key, oldLabels[key], r.Labels)
		}

		// add new version of record to caches
		im.labels[key] = r.Labels
		if im.cacheInterval > 0 {
			im.addToCache(r)
		}
	}

	err := im.executeStatements(ctx, statements, func(request *dynamodb.BatchStatementRequest, response *dynamodb.BatchStatementResponse) error {
		var context string
		if strings.HasPrefix(*request.Statement, "INSERT") {
			if aws.StringValue(response.Error.Code) == "DuplicateItem" {
				// We lost a race with a different owner or another owner has an orphaned ownership record.
				key := fromDynamoKey(request.Parameters[0])
				for i, endpoint := range filteredChanges.Create {
					if endpoint.Key() == key {
						log.Infof("Skipping endpoint %v because owner does not match", endpoint)
						filteredChanges.Create = append(filteredChanges.Create[:i], filteredChanges.Create[i+1:]...)
						// The dynamodb insertion failed; remove from our cache.
						im.removeFromCache(endpoint)
						delete(im.labels, key)
						return nil
					}
				}
			}
			context = fmt.Sprintf("inserting dynamodb record %q", aws.StringValue(request.Parameters[0].S))
		} else {
			context = fmt.Sprintf("updating dynamodb record %q", aws.StringValue(request.Parameters[1].S))
		}
		return fmt.Errorf("%s: %s: %s", context, aws.StringValue(response.Error.Code), aws.StringValue(response.Error.Message))
	})
	if err != nil {
		im.recordsCache = nil
		im.labels = nil
		return err
	}

	// When caching is enabled, disable the provider from using the cache.
	if im.cacheInterval > 0 {
		ctx = context.WithValue(ctx, provider.RecordsContextKey, nil)
	}
	err = im.provider.ApplyChanges(ctx, filteredChanges)
	if err != nil {
		im.recordsCache = nil
		im.labels = nil
		return err
	}

	statements = make([]*dynamodb.BatchStatementRequest, 0, len(filteredChanges.Delete)+len(im.orphanedLabels))
	for _, r := range filteredChanges.Delete {
		statements = im.appendDelete(statements, r.Key())
	}
	for r := range im.orphanedLabels {
		statements = im.appendDelete(statements, r)
		delete(im.labels, r)
	}
	im.orphanedLabels = nil
	return im.executeStatements(ctx, statements, func(request *dynamodb.BatchStatementRequest, response *dynamodb.BatchStatementResponse) error {
		im.labels = nil
		return fmt.Errorf("deleting dynamodb record %q: %s: %s", aws.StringValue(request.Parameters[0].S), aws.StringValue(response.Error.Code), aws.StringValue(response.Error.Message))
	})
}

// AdjustEndpoints modifies the endpoints as needed by the specific provider.
func (im *DynamoDBRegistry) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	return im.provider.AdjustEndpoints(endpoints)
}

func (im *DynamoDBRegistry) readLabels(ctx context.Context) error {
	table, err := im.dynamodbAPI.DescribeTableWithContext(ctx, &dynamodb.DescribeTableInput{
		TableName: aws.String(im.table),
	})
	if err != nil {
		return fmt.Errorf("describing table %q: %w", im.table, err)
	}

	foundKey := false
	for _, def := range table.Table.AttributeDefinitions {
		if aws.StringValue(def.AttributeName) == "k" {
			if aws.StringValue(def.AttributeType) != "S" {
				return fmt.Errorf("table %q attribute \"k\" must have type \"S\"", im.table)
			}
			foundKey = true
		}
	}
	if !foundKey {
		return fmt.Errorf("table %q must have attribute \"k\" of type \"S\"", im.table)
	}

	if aws.StringValue(table.Table.KeySchema[0].AttributeName) != "k" {
		return fmt.Errorf("table %q must have hash key \"k\"", im.table)
	}
	if len(table.Table.KeySchema) > 1 {
		return fmt.Errorf("table %q must not have a range key", im.table)
	}

	labels := map[endpoint.EndpointKey]endpoint.Labels{}
	err = im.dynamodbAPI.ScanPagesWithContext(ctx, &dynamodb.ScanInput{
		TableName:        aws.String(im.table),
		FilterExpression: aws.String("o = :ownerval"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":ownerval": {S: aws.String(im.ownerID)},
		},
		ProjectionExpression: aws.String("k,l"),
		ConsistentRead:       aws.Bool(true),
	}, func(output *dynamodb.ScanOutput, last bool) bool {
		for _, item := range output.Items {
			labels[fromDynamoKey(item["k"])] = fromDynamoLabels(item["l"], im.ownerID)
		}
		return true
	})
	if err != nil {
		return fmt.Errorf("querying dynamodb: %w", err)
	}

	im.labels = labels
	return nil
}

func fromDynamoKey(key *dynamodb.AttributeValue) endpoint.EndpointKey {
	split := strings.SplitN(aws.StringValue(key.S), "#", 3)
	return endpoint.EndpointKey{
		DNSName:       split[0],
		RecordType:    split[1],
		SetIdentifier: split[2],
	}
}

func toDynamoKey(key endpoint.EndpointKey) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{
		S: aws.String(fmt.Sprintf("%s#%s#%s", key.DNSName, key.RecordType, key.SetIdentifier)),
	}
}

func fromDynamoLabels(label *dynamodb.AttributeValue, owner string) endpoint.Labels {
	labels := endpoint.NewLabels()
	for k, v := range label.M {
		labels[k] = aws.StringValue(v.S)
	}
	labels[endpoint.OwnerLabelKey] = owner
	return labels
}

func toDynamoLabels(labels endpoint.Labels) *dynamodb.AttributeValue {
	labelMap := make(map[string]*dynamodb.AttributeValue, len(labels))
	for k, v := range labels {
		if k == endpoint.OwnerLabelKey {
			continue
		}
		labelMap[k] = &dynamodb.AttributeValue{S: aws.String(v)}
	}
	return &dynamodb.AttributeValue{M: labelMap}
}

func (im *DynamoDBRegistry) appendInsert(statements []*dynamodb.BatchStatementRequest, key endpoint.EndpointKey, new endpoint.Labels) []*dynamodb.BatchStatementRequest {
	return append(statements, &dynamodb.BatchStatementRequest{
		Statement: aws.String(fmt.Sprintf("INSERT INTO %q VALUE {'k':?, 'o':?, 'l':?}", im.table)),
		Parameters: []*dynamodb.AttributeValue{
			toDynamoKey(key),
			{S: aws.String(im.ownerID)},
			toDynamoLabels(new),
		},
		ConsistentRead: aws.Bool(true),
	})
}

func (im *DynamoDBRegistry) appendUpdate(statements []*dynamodb.BatchStatementRequest, key endpoint.EndpointKey, old endpoint.Labels, new endpoint.Labels) []*dynamodb.BatchStatementRequest {
	if len(old) == len(new) {
		equal := true
		for k, v := range old {
			if newV, exists := new[k]; !exists || v != newV {
				equal = false
				break
			}
		}
		if equal {
			return statements
		}
	}

	return append(statements, &dynamodb.BatchStatementRequest{
		Statement: aws.String(fmt.Sprintf("UPDATE %q SET \"l\"=? WHERE \"k\"=?", im.table)),
		Parameters: []*dynamodb.AttributeValue{
			toDynamoLabels(new),
			toDynamoKey(key),
		},
	})
}

func (im *DynamoDBRegistry) appendDelete(statements []*dynamodb.BatchStatementRequest, key endpoint.EndpointKey) []*dynamodb.BatchStatementRequest {
	return append(statements, &dynamodb.BatchStatementRequest{
		Statement: aws.String(fmt.Sprintf("DELETE FROM %q WHERE \"k\"=? AND \"o\"=?", im.table)),
		Parameters: []*dynamodb.AttributeValue{
			toDynamoKey(key),
			{S: aws.String(im.ownerID)},
		},
	})
}

func (im *DynamoDBRegistry) executeStatements(ctx context.Context, statements []*dynamodb.BatchStatementRequest, handleErr func(request *dynamodb.BatchStatementRequest, response *dynamodb.BatchStatementResponse) error) error {
	for len(statements) > 0 {
		var chunk []*dynamodb.BatchStatementRequest
		if len(statements) > int(dynamodbMaxBatchSize) {
			chunk = statements[:dynamodbMaxBatchSize]
			statements = statements[dynamodbMaxBatchSize:]
		} else {
			chunk = statements
			statements = nil
		}

		output, err := im.dynamodbAPI.BatchExecuteStatementWithContext(ctx, &dynamodb.BatchExecuteStatementInput{
			Statements: chunk,
		})
		if err != nil {
			return err
		}

		for i, response := range output.Responses {
			request := chunk[i]
			if response.Error == nil {
				op, _, _ := strings.Cut(*request.Statement, " ")
				var key string
				if op == "UPDATE" {
					key = *request.Parameters[1].S
				} else {
					key = *request.Parameters[0].S
				}
				log.Infof("%s dynamodb record %q", op, key)
			} else {
				if err := handleErr(request, response); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (im *DynamoDBRegistry) addToCache(ep *endpoint.Endpoint) {
	if im.recordsCache != nil {
		im.recordsCache = append(im.recordsCache, ep)
	}
}

func (im *DynamoDBRegistry) removeFromCache(ep *endpoint.Endpoint) {
	if im.recordsCache == nil || ep == nil {
		return
	}

	for i, e := range im.recordsCache {
		if e.DNSName == ep.DNSName && e.RecordType == ep.RecordType && e.SetIdentifier == ep.SetIdentifier && e.Targets.Same(ep.Targets) {
			// We found a match; delete the endpoint from the cache.
			im.recordsCache = append(im.recordsCache[:i], im.recordsCache[i+1:]...)
			return
		}
	}
}
