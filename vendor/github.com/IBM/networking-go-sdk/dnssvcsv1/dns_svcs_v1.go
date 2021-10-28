/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.38.0-07189efd-20210827-205025
 */

// Package dnssvcsv1 : Operations and models for the DnsSvcsV1 service
package dnssvcsv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/networking-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// DnsSvcsV1 : DNS Services API
//
// API Version: 1.0.0
type DnsSvcsV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.dns-svcs.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "dns_svcs"

// DnsSvcsV1Options : Service options
type DnsSvcsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewDnsSvcsV1UsingExternalConfig : constructs an instance of DnsSvcsV1 with passed in options and external configuration.
func NewDnsSvcsV1UsingExternalConfig(options *DnsSvcsV1Options) (dnsSvcs *DnsSvcsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	dnsSvcs, err = NewDnsSvcsV1(options)
	if err != nil {
		return
	}

	err = dnsSvcs.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = dnsSvcs.Service.SetServiceURL(options.URL)
	}
	return
}

// NewDnsSvcsV1 : constructs an instance of DnsSvcsV1 with passed in options.
func NewDnsSvcsV1(options *DnsSvcsV1Options) (service *DnsSvcsV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &DnsSvcsV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "dnsSvcs" suitable for processing requests.
func (dnsSvcs *DnsSvcsV1) Clone() *DnsSvcsV1 {
	if core.IsNil(dnsSvcs) {
		return nil
	}
	clone := *dnsSvcs
	clone.Service = dnsSvcs.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (dnsSvcs *DnsSvcsV1) SetServiceURL(url string) error {
	return dnsSvcs.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (dnsSvcs *DnsSvcsV1) GetServiceURL() string {
	return dnsSvcs.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (dnsSvcs *DnsSvcsV1) SetDefaultHeaders(headers http.Header) {
	dnsSvcs.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (dnsSvcs *DnsSvcsV1) SetEnableGzipCompression(enableGzip bool) {
	dnsSvcs.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (dnsSvcs *DnsSvcsV1) GetEnableGzipCompression() bool {
	return dnsSvcs.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (dnsSvcs *DnsSvcsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	dnsSvcs.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (dnsSvcs *DnsSvcsV1) DisableRetries() {
	dnsSvcs.Service.DisableRetries()
}

// ListDnszones : List DNS zones
// List the DNS zones for a given service instance.
func (dnsSvcs *DnsSvcsV1) ListDnszones(listDnszonesOptions *ListDnszonesOptions) (result *ListDnszones, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListDnszonesWithContext(context.Background(), listDnszonesOptions)
}

// ListDnszonesWithContext is an alternate form of the ListDnszones method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListDnszonesWithContext(ctx context.Context, listDnszonesOptions *ListDnszonesOptions) (result *ListDnszones, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDnszonesOptions, "listDnszonesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDnszonesOptions, "listDnszonesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listDnszonesOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDnszonesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListDnszones")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDnszonesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listDnszonesOptions.XCorrelationID))
	}

	if listDnszonesOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listDnszonesOptions.Offset))
	}
	if listDnszonesOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listDnszonesOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListDnszones)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDnszone : Create DNS zone
// Create a DNS zone for a given service instance.
func (dnsSvcs *DnsSvcsV1) CreateDnszone(createDnszoneOptions *CreateDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreateDnszoneWithContext(context.Background(), createDnszoneOptions)
}

// CreateDnszoneWithContext is an alternate form of the CreateDnszone method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreateDnszoneWithContext(ctx context.Context, createDnszoneOptions *CreateDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDnszoneOptions, "createDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDnszoneOptions, "createDnszoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createDnszoneOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createDnszoneOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createDnszoneOptions.Name != nil {
		body["name"] = createDnszoneOptions.Name
	}
	if createDnszoneOptions.Description != nil {
		body["description"] = createDnszoneOptions.Description
	}
	if createDnszoneOptions.Label != nil {
		body["label"] = createDnszoneOptions.Label
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDnszone)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDnszone : Delete DNS zone
// Delete a DNS zone.
func (dnsSvcs *DnsSvcsV1) DeleteDnszone(deleteDnszoneOptions *DeleteDnszoneOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeleteDnszoneWithContext(context.Background(), deleteDnszoneOptions)
}

// DeleteDnszoneWithContext is an alternate form of the DeleteDnszone method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeleteDnszoneWithContext(ctx context.Context, deleteDnszoneOptions *DeleteDnszoneOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDnszoneOptions, "deleteDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDnszoneOptions, "deleteDnszoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteDnszoneOptions.InstanceID,
		"dnszone_id":  *deleteDnszoneOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteDnszoneOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetDnszone : Get DNS zone
// Get details of a DNS zone.
func (dnsSvcs *DnsSvcsV1) GetDnszone(getDnszoneOptions *GetDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetDnszoneWithContext(context.Background(), getDnszoneOptions)
}

// GetDnszoneWithContext is an alternate form of the GetDnszone method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetDnszoneWithContext(ctx context.Context, getDnszoneOptions *GetDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDnszoneOptions, "getDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDnszoneOptions, "getDnszoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getDnszoneOptions.InstanceID,
		"dnszone_id":  *getDnszoneOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getDnszoneOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDnszone)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateDnszone : Update DNS zone
// Update the properties of a DNS zone.
func (dnsSvcs *DnsSvcsV1) UpdateDnszone(updateDnszoneOptions *UpdateDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdateDnszoneWithContext(context.Background(), updateDnszoneOptions)
}

// UpdateDnszoneWithContext is an alternate form of the UpdateDnszone method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdateDnszoneWithContext(ctx context.Context, updateDnszoneOptions *UpdateDnszoneOptions) (result *Dnszone, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDnszoneOptions, "updateDnszoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDnszoneOptions, "updateDnszoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateDnszoneOptions.InstanceID,
		"dnszone_id":  *updateDnszoneOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDnszoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateDnszone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateDnszoneOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateDnszoneOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateDnszoneOptions.Description != nil {
		body["description"] = updateDnszoneOptions.Description
	}
	if updateDnszoneOptions.Label != nil {
		body["label"] = updateDnszoneOptions.Label
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDnszone)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListResourceRecords : List resource records
// List the Resource Records for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) ListResourceRecords(listResourceRecordsOptions *ListResourceRecordsOptions) (result *ListResourceRecords, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListResourceRecordsWithContext(context.Background(), listResourceRecordsOptions)
}

// ListResourceRecordsWithContext is an alternate form of the ListResourceRecords method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListResourceRecordsWithContext(ctx context.Context, listResourceRecordsOptions *ListResourceRecordsOptions) (result *ListResourceRecords, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listResourceRecordsOptions, "listResourceRecordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listResourceRecordsOptions, "listResourceRecordsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listResourceRecordsOptions.InstanceID,
		"dnszone_id":  *listResourceRecordsOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/resource_records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listResourceRecordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListResourceRecords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listResourceRecordsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listResourceRecordsOptions.XCorrelationID))
	}

	if listResourceRecordsOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listResourceRecordsOptions.Offset))
	}
	if listResourceRecordsOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listResourceRecordsOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListResourceRecords)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateResourceRecord : Create resource record
// Create a resource record for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) CreateResourceRecord(createResourceRecordOptions *CreateResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreateResourceRecordWithContext(context.Background(), createResourceRecordOptions)
}

// CreateResourceRecordWithContext is an alternate form of the CreateResourceRecord method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreateResourceRecordWithContext(ctx context.Context, createResourceRecordOptions *CreateResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createResourceRecordOptions, "createResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createResourceRecordOptions, "createResourceRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createResourceRecordOptions.InstanceID,
		"dnszone_id":  *createResourceRecordOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/resource_records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createResourceRecordOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createResourceRecordOptions.Name != nil {
		body["name"] = createResourceRecordOptions.Name
	}
	if createResourceRecordOptions.Type != nil {
		body["type"] = createResourceRecordOptions.Type
	}
	if createResourceRecordOptions.Rdata != nil {
		body["rdata"] = createResourceRecordOptions.Rdata
	}
	if createResourceRecordOptions.TTL != nil {
		body["ttl"] = createResourceRecordOptions.TTL
	}
	if createResourceRecordOptions.Service != nil {
		body["service"] = createResourceRecordOptions.Service
	}
	if createResourceRecordOptions.Protocol != nil {
		body["protocol"] = createResourceRecordOptions.Protocol
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResourceRecord)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteResourceRecord : Delete resource record
// Delete a resource record.
func (dnsSvcs *DnsSvcsV1) DeleteResourceRecord(deleteResourceRecordOptions *DeleteResourceRecordOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeleteResourceRecordWithContext(context.Background(), deleteResourceRecordOptions)
}

// DeleteResourceRecordWithContext is an alternate form of the DeleteResourceRecord method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeleteResourceRecordWithContext(ctx context.Context, deleteResourceRecordOptions *DeleteResourceRecordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteResourceRecordOptions, "deleteResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteResourceRecordOptions, "deleteResourceRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteResourceRecordOptions.InstanceID,
		"dnszone_id":  *deleteResourceRecordOptions.DnszoneID,
		"record_id":   *deleteResourceRecordOptions.RecordID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/resource_records/{record_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteResourceRecordOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetResourceRecord : Get resource record
// Get details of a resource record.
func (dnsSvcs *DnsSvcsV1) GetResourceRecord(getResourceRecordOptions *GetResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetResourceRecordWithContext(context.Background(), getResourceRecordOptions)
}

// GetResourceRecordWithContext is an alternate form of the GetResourceRecord method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetResourceRecordWithContext(ctx context.Context, getResourceRecordOptions *GetResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getResourceRecordOptions, "getResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getResourceRecordOptions, "getResourceRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getResourceRecordOptions.InstanceID,
		"dnszone_id":  *getResourceRecordOptions.DnszoneID,
		"record_id":   *getResourceRecordOptions.RecordID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/resource_records/{record_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getResourceRecordOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResourceRecord)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateResourceRecord : Update resource record
// Update the properties of a resource record.
func (dnsSvcs *DnsSvcsV1) UpdateResourceRecord(updateResourceRecordOptions *UpdateResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdateResourceRecordWithContext(context.Background(), updateResourceRecordOptions)
}

// UpdateResourceRecordWithContext is an alternate form of the UpdateResourceRecord method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdateResourceRecordWithContext(ctx context.Context, updateResourceRecordOptions *UpdateResourceRecordOptions) (result *ResourceRecord, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateResourceRecordOptions, "updateResourceRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateResourceRecordOptions, "updateResourceRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateResourceRecordOptions.InstanceID,
		"dnszone_id":  *updateResourceRecordOptions.DnszoneID,
		"record_id":   *updateResourceRecordOptions.RecordID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/resource_records/{record_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateResourceRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateResourceRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateResourceRecordOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateResourceRecordOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateResourceRecordOptions.Name != nil {
		body["name"] = updateResourceRecordOptions.Name
	}
	if updateResourceRecordOptions.Rdata != nil {
		body["rdata"] = updateResourceRecordOptions.Rdata
	}
	if updateResourceRecordOptions.TTL != nil {
		body["ttl"] = updateResourceRecordOptions.TTL
	}
	if updateResourceRecordOptions.Service != nil {
		body["service"] = updateResourceRecordOptions.Service
	}
	if updateResourceRecordOptions.Protocol != nil {
		body["protocol"] = updateResourceRecordOptions.Protocol
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResourceRecord)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ExportResourceRecords : Export resource records to a zone file
// Export resource records to a zone file.
func (dnsSvcs *DnsSvcsV1) ExportResourceRecords(exportResourceRecordsOptions *ExportResourceRecordsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return dnsSvcs.ExportResourceRecordsWithContext(context.Background(), exportResourceRecordsOptions)
}

// ExportResourceRecordsWithContext is an alternate form of the ExportResourceRecords method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ExportResourceRecordsWithContext(ctx context.Context, exportResourceRecordsOptions *ExportResourceRecordsOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(exportResourceRecordsOptions, "exportResourceRecordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(exportResourceRecordsOptions, "exportResourceRecordsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *exportResourceRecordsOptions.InstanceID,
		"dnszone_id":  *exportResourceRecordsOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/export_resource_records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range exportResourceRecordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ExportResourceRecords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "text/plain; charset=utf-8")
	if exportResourceRecordsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*exportResourceRecordsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, &result)

	return
}

// ImportResourceRecords : Import resource records from a zone file
// Import resource records from a zone file.
func (dnsSvcs *DnsSvcsV1) ImportResourceRecords(importResourceRecordsOptions *ImportResourceRecordsOptions) (result *ImportResourceRecordsResp, response *core.DetailedResponse, err error) {
	return dnsSvcs.ImportResourceRecordsWithContext(context.Background(), importResourceRecordsOptions)
}

// ImportResourceRecordsWithContext is an alternate form of the ImportResourceRecords method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ImportResourceRecordsWithContext(ctx context.Context, importResourceRecordsOptions *ImportResourceRecordsOptions) (result *ImportResourceRecordsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(importResourceRecordsOptions, "importResourceRecordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(importResourceRecordsOptions, "importResourceRecordsOptions")
	if err != nil {
		return
	}
	if importResourceRecordsOptions.File == nil {
		err = fmt.Errorf("at least one of  or file must be supplied")
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *importResourceRecordsOptions.InstanceID,
		"dnszone_id":  *importResourceRecordsOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/import_resource_records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range importResourceRecordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ImportResourceRecords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if importResourceRecordsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*importResourceRecordsOptions.XCorrelationID))
	}

	if importResourceRecordsOptions.File != nil {
		builder.AddFormData("file", "filename",
			core.StringNilMapper(importResourceRecordsOptions.FileContentType), importResourceRecordsOptions.File)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImportResourceRecordsResp)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListPermittedNetworks : List permitted networks
// List the permitted networks for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) ListPermittedNetworks(listPermittedNetworksOptions *ListPermittedNetworksOptions) (result *ListPermittedNetworks, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListPermittedNetworksWithContext(context.Background(), listPermittedNetworksOptions)
}

// ListPermittedNetworksWithContext is an alternate form of the ListPermittedNetworks method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListPermittedNetworksWithContext(ctx context.Context, listPermittedNetworksOptions *ListPermittedNetworksOptions) (result *ListPermittedNetworks, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPermittedNetworksOptions, "listPermittedNetworksOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPermittedNetworksOptions, "listPermittedNetworksOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listPermittedNetworksOptions.InstanceID,
		"dnszone_id":  *listPermittedNetworksOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/permitted_networks`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPermittedNetworksOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListPermittedNetworks")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPermittedNetworksOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listPermittedNetworksOptions.XCorrelationID))
	}

	if listPermittedNetworksOptions.Offset != nil {
		builder.AddQuery("offset", fmt.Sprint(*listPermittedNetworksOptions.Offset))
	}
	if listPermittedNetworksOptions.Limit != nil {
		builder.AddQuery("limit", fmt.Sprint(*listPermittedNetworksOptions.Limit))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListPermittedNetworks)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePermittedNetwork : Create permitted network
// Create a permitted network for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) CreatePermittedNetwork(createPermittedNetworkOptions *CreatePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreatePermittedNetworkWithContext(context.Background(), createPermittedNetworkOptions)
}

// CreatePermittedNetworkWithContext is an alternate form of the CreatePermittedNetwork method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreatePermittedNetworkWithContext(ctx context.Context, createPermittedNetworkOptions *CreatePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPermittedNetworkOptions, "createPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPermittedNetworkOptions, "createPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createPermittedNetworkOptions.InstanceID,
		"dnszone_id":  *createPermittedNetworkOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/permitted_networks`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreatePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createPermittedNetworkOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createPermittedNetworkOptions.Type != nil {
		body["type"] = createPermittedNetworkOptions.Type
	}
	if createPermittedNetworkOptions.PermittedNetwork != nil {
		body["permitted_network"] = createPermittedNetworkOptions.PermittedNetwork
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeletePermittedNetwork : Remove permitted network
// Remove a permitted network.
func (dnsSvcs *DnsSvcsV1) DeletePermittedNetwork(deletePermittedNetworkOptions *DeletePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	return dnsSvcs.DeletePermittedNetworkWithContext(context.Background(), deletePermittedNetworkOptions)
}

// DeletePermittedNetworkWithContext is an alternate form of the DeletePermittedNetwork method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeletePermittedNetworkWithContext(ctx context.Context, deletePermittedNetworkOptions *DeletePermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePermittedNetworkOptions, "deletePermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePermittedNetworkOptions, "deletePermittedNetworkOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id":          *deletePermittedNetworkOptions.InstanceID,
		"dnszone_id":           *deletePermittedNetworkOptions.DnszoneID,
		"permitted_network_id": *deletePermittedNetworkOptions.PermittedNetworkID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/permitted_networks/{permitted_network_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeletePermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deletePermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deletePermittedNetworkOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetPermittedNetwork : Get permitted network
// Get details of a permitted network.
func (dnsSvcs *DnsSvcsV1) GetPermittedNetwork(getPermittedNetworkOptions *GetPermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetPermittedNetworkWithContext(context.Background(), getPermittedNetworkOptions)
}

// GetPermittedNetworkWithContext is an alternate form of the GetPermittedNetwork method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetPermittedNetworkWithContext(ctx context.Context, getPermittedNetworkOptions *GetPermittedNetworkOptions) (result *PermittedNetwork, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPermittedNetworkOptions, "getPermittedNetworkOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPermittedNetworkOptions, "getPermittedNetworkOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id":          *getPermittedNetworkOptions.InstanceID,
		"dnszone_id":           *getPermittedNetworkOptions.DnszoneID,
		"permitted_network_id": *getPermittedNetworkOptions.PermittedNetworkID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/permitted_networks/{permitted_network_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPermittedNetworkOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetPermittedNetwork")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPermittedNetworkOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getPermittedNetworkOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPermittedNetwork)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListLoadBalancers : List load balancers
// List the Global Load Balancers for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) ListLoadBalancers(listLoadBalancersOptions *ListLoadBalancersOptions) (result *ListLoadBalancers, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListLoadBalancersWithContext(context.Background(), listLoadBalancersOptions)
}

// ListLoadBalancersWithContext is an alternate form of the ListLoadBalancers method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListLoadBalancersWithContext(ctx context.Context, listLoadBalancersOptions *ListLoadBalancersOptions) (result *ListLoadBalancers, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listLoadBalancersOptions, "listLoadBalancersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listLoadBalancersOptions, "listLoadBalancersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listLoadBalancersOptions.InstanceID,
		"dnszone_id":  *listLoadBalancersOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/load_balancers`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listLoadBalancersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListLoadBalancers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listLoadBalancersOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listLoadBalancersOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListLoadBalancers)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateLoadBalancer : Create load balancer
// Create a load balancer for a given DNS zone.
func (dnsSvcs *DnsSvcsV1) CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreateLoadBalancerWithContext(context.Background(), createLoadBalancerOptions)
}

// CreateLoadBalancerWithContext is an alternate form of the CreateLoadBalancer method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreateLoadBalancerWithContext(ctx context.Context, createLoadBalancerOptions *CreateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createLoadBalancerOptions, "createLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createLoadBalancerOptions, "createLoadBalancerOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createLoadBalancerOptions.InstanceID,
		"dnszone_id":  *createLoadBalancerOptions.DnszoneID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/load_balancers`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createLoadBalancerOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createLoadBalancerOptions.Name != nil {
		body["name"] = createLoadBalancerOptions.Name
	}
	if createLoadBalancerOptions.Description != nil {
		body["description"] = createLoadBalancerOptions.Description
	}
	if createLoadBalancerOptions.Enabled != nil {
		body["enabled"] = createLoadBalancerOptions.Enabled
	}
	if createLoadBalancerOptions.TTL != nil {
		body["ttl"] = createLoadBalancerOptions.TTL
	}
	if createLoadBalancerOptions.FallbackPool != nil {
		body["fallback_pool"] = createLoadBalancerOptions.FallbackPool
	}
	if createLoadBalancerOptions.DefaultPools != nil {
		body["default_pools"] = createLoadBalancerOptions.DefaultPools
	}
	if createLoadBalancerOptions.AzPools != nil {
		body["az_pools"] = createLoadBalancerOptions.AzPools
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoadBalancer)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteLoadBalancer : Delete load balancer
// Delete a load balancer.
func (dnsSvcs *DnsSvcsV1) DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeleteLoadBalancerWithContext(context.Background(), deleteLoadBalancerOptions)
}

// DeleteLoadBalancerWithContext is an alternate form of the DeleteLoadBalancer method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeleteLoadBalancerWithContext(ctx context.Context, deleteLoadBalancerOptions *DeleteLoadBalancerOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteLoadBalancerOptions, "deleteLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteLoadBalancerOptions, "deleteLoadBalancerOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteLoadBalancerOptions.InstanceID,
		"dnszone_id":  *deleteLoadBalancerOptions.DnszoneID,
		"lb_id":       *deleteLoadBalancerOptions.LbID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/load_balancers/{lb_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteLoadBalancerOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetLoadBalancer : Get load balancer
// Get details of a load balancer.
func (dnsSvcs *DnsSvcsV1) GetLoadBalancer(getLoadBalancerOptions *GetLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetLoadBalancerWithContext(context.Background(), getLoadBalancerOptions)
}

// GetLoadBalancerWithContext is an alternate form of the GetLoadBalancer method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetLoadBalancerWithContext(ctx context.Context, getLoadBalancerOptions *GetLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getLoadBalancerOptions, "getLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getLoadBalancerOptions, "getLoadBalancerOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getLoadBalancerOptions.InstanceID,
		"dnszone_id":  *getLoadBalancerOptions.DnszoneID,
		"lb_id":       *getLoadBalancerOptions.LbID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/load_balancers/{lb_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getLoadBalancerOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoadBalancer)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateLoadBalancer : Update load balancer
// Update the properties of a load balancer.
func (dnsSvcs *DnsSvcsV1) UpdateLoadBalancer(updateLoadBalancerOptions *UpdateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdateLoadBalancerWithContext(context.Background(), updateLoadBalancerOptions)
}

// UpdateLoadBalancerWithContext is an alternate form of the UpdateLoadBalancer method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdateLoadBalancerWithContext(ctx context.Context, updateLoadBalancerOptions *UpdateLoadBalancerOptions) (result *LoadBalancer, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateLoadBalancerOptions, "updateLoadBalancerOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateLoadBalancerOptions, "updateLoadBalancerOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateLoadBalancerOptions.InstanceID,
		"dnszone_id":  *updateLoadBalancerOptions.DnszoneID,
		"lb_id":       *updateLoadBalancerOptions.LbID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/dnszones/{dnszone_id}/load_balancers/{lb_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateLoadBalancerOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateLoadBalancer")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateLoadBalancerOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateLoadBalancerOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateLoadBalancerOptions.Name != nil {
		body["name"] = updateLoadBalancerOptions.Name
	}
	if updateLoadBalancerOptions.Description != nil {
		body["description"] = updateLoadBalancerOptions.Description
	}
	if updateLoadBalancerOptions.Enabled != nil {
		body["enabled"] = updateLoadBalancerOptions.Enabled
	}
	if updateLoadBalancerOptions.TTL != nil {
		body["ttl"] = updateLoadBalancerOptions.TTL
	}
	if updateLoadBalancerOptions.FallbackPool != nil {
		body["fallback_pool"] = updateLoadBalancerOptions.FallbackPool
	}
	if updateLoadBalancerOptions.DefaultPools != nil {
		body["default_pools"] = updateLoadBalancerOptions.DefaultPools
	}
	if updateLoadBalancerOptions.AzPools != nil {
		body["az_pools"] = updateLoadBalancerOptions.AzPools
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLoadBalancer)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListPools : List load balancer pools
// List the load balancer pools.
func (dnsSvcs *DnsSvcsV1) ListPools(listPoolsOptions *ListPoolsOptions) (result *ListPools, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListPoolsWithContext(context.Background(), listPoolsOptions)
}

// ListPoolsWithContext is an alternate form of the ListPools method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListPoolsWithContext(ctx context.Context, listPoolsOptions *ListPoolsOptions) (result *ListPools, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listPoolsOptions, "listPoolsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listPoolsOptions, "listPoolsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listPoolsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/pools`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPoolsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListPools")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPoolsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listPoolsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListPools)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreatePool : Create load balancer pool
// Create a load balancer pool.
func (dnsSvcs *DnsSvcsV1) CreatePool(createPoolOptions *CreatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreatePoolWithContext(context.Background(), createPoolOptions)
}

// CreatePoolWithContext is an alternate form of the CreatePool method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreatePoolWithContext(ctx context.Context, createPoolOptions *CreatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createPoolOptions, "createPoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createPoolOptions, "createPoolOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createPoolOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/pools`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createPoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreatePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createPoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createPoolOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createPoolOptions.Name != nil {
		body["name"] = createPoolOptions.Name
	}
	if createPoolOptions.Description != nil {
		body["description"] = createPoolOptions.Description
	}
	if createPoolOptions.Enabled != nil {
		body["enabled"] = createPoolOptions.Enabled
	}
	if createPoolOptions.HealthyOriginsThreshold != nil {
		body["healthy_origins_threshold"] = createPoolOptions.HealthyOriginsThreshold
	}
	if createPoolOptions.Origins != nil {
		body["origins"] = createPoolOptions.Origins
	}
	if createPoolOptions.Monitor != nil {
		body["monitor"] = createPoolOptions.Monitor
	}
	if createPoolOptions.NotificationChannel != nil {
		body["notification_channel"] = createPoolOptions.NotificationChannel
	}
	if createPoolOptions.HealthcheckRegion != nil {
		body["healthcheck_region"] = createPoolOptions.HealthcheckRegion
	}
	if createPoolOptions.HealthcheckSubnets != nil {
		body["healthcheck_subnets"] = createPoolOptions.HealthcheckSubnets
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPool)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeletePool : Delete load balancer pool
// Delete a load balancer pool.
func (dnsSvcs *DnsSvcsV1) DeletePool(deletePoolOptions *DeletePoolOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeletePoolWithContext(context.Background(), deletePoolOptions)
}

// DeletePoolWithContext is an alternate form of the DeletePool method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeletePoolWithContext(ctx context.Context, deletePoolOptions *DeletePoolOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deletePoolOptions, "deletePoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deletePoolOptions, "deletePoolOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deletePoolOptions.InstanceID,
		"pool_id":     *deletePoolOptions.PoolID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/pools/{pool_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deletePoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeletePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deletePoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deletePoolOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetPool : Get load balancer pool
// Get details of a load balancer pool.
func (dnsSvcs *DnsSvcsV1) GetPool(getPoolOptions *GetPoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetPoolWithContext(context.Background(), getPoolOptions)
}

// GetPoolWithContext is an alternate form of the GetPool method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetPoolWithContext(ctx context.Context, getPoolOptions *GetPoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPoolOptions, "getPoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPoolOptions, "getPoolOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getPoolOptions.InstanceID,
		"pool_id":     *getPoolOptions.PoolID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/pools/{pool_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetPool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getPoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getPoolOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPool)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdatePool : Update load balancer pool
// Update the properties of a load balancer pool.
func (dnsSvcs *DnsSvcsV1) UpdatePool(updatePoolOptions *UpdatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdatePoolWithContext(context.Background(), updatePoolOptions)
}

// UpdatePoolWithContext is an alternate form of the UpdatePool method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdatePoolWithContext(ctx context.Context, updatePoolOptions *UpdatePoolOptions) (result *Pool, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updatePoolOptions, "updatePoolOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updatePoolOptions, "updatePoolOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updatePoolOptions.InstanceID,
		"pool_id":     *updatePoolOptions.PoolID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/pools/{pool_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updatePoolOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdatePool")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updatePoolOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updatePoolOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updatePoolOptions.Name != nil {
		body["name"] = updatePoolOptions.Name
	}
	if updatePoolOptions.Description != nil {
		body["description"] = updatePoolOptions.Description
	}
	if updatePoolOptions.Enabled != nil {
		body["enabled"] = updatePoolOptions.Enabled
	}
	if updatePoolOptions.HealthyOriginsThreshold != nil {
		body["healthy_origins_threshold"] = updatePoolOptions.HealthyOriginsThreshold
	}
	if updatePoolOptions.Origins != nil {
		body["origins"] = updatePoolOptions.Origins
	}
	if updatePoolOptions.Monitor != nil {
		body["monitor"] = updatePoolOptions.Monitor
	}
	if updatePoolOptions.NotificationChannel != nil {
		body["notification_channel"] = updatePoolOptions.NotificationChannel
	}
	if updatePoolOptions.HealthcheckRegion != nil {
		body["healthcheck_region"] = updatePoolOptions.HealthcheckRegion
	}
	if updatePoolOptions.HealthcheckSubnets != nil {
		body["healthcheck_subnets"] = updatePoolOptions.HealthcheckSubnets
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPool)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListMonitors : List load balancer monitors
// List the load balancer monitors.
func (dnsSvcs *DnsSvcsV1) ListMonitors(listMonitorsOptions *ListMonitorsOptions) (result *ListMonitors, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListMonitorsWithContext(context.Background(), listMonitorsOptions)
}

// ListMonitorsWithContext is an alternate form of the ListMonitors method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListMonitorsWithContext(ctx context.Context, listMonitorsOptions *ListMonitorsOptions) (result *ListMonitors, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listMonitorsOptions, "listMonitorsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listMonitorsOptions, "listMonitorsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listMonitorsOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/monitors`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listMonitorsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListMonitors")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listMonitorsOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listMonitorsOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListMonitors)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateMonitor : Create load balancer monitor
// Create a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) CreateMonitor(createMonitorOptions *CreateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreateMonitorWithContext(context.Background(), createMonitorOptions)
}

// CreateMonitorWithContext is an alternate form of the CreateMonitor method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreateMonitorWithContext(ctx context.Context, createMonitorOptions *CreateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createMonitorOptions, "createMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createMonitorOptions, "createMonitorOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createMonitorOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/monitors`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createMonitorOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createMonitorOptions.Name != nil {
		body["name"] = createMonitorOptions.Name
	}
	if createMonitorOptions.Description != nil {
		body["description"] = createMonitorOptions.Description
	}
	if createMonitorOptions.Type != nil {
		body["type"] = createMonitorOptions.Type
	}
	if createMonitorOptions.Port != nil {
		body["port"] = createMonitorOptions.Port
	}
	if createMonitorOptions.Interval != nil {
		body["interval"] = createMonitorOptions.Interval
	}
	if createMonitorOptions.Retries != nil {
		body["retries"] = createMonitorOptions.Retries
	}
	if createMonitorOptions.Timeout != nil {
		body["timeout"] = createMonitorOptions.Timeout
	}
	if createMonitorOptions.Method != nil {
		body["method"] = createMonitorOptions.Method
	}
	if createMonitorOptions.Path != nil {
		body["path"] = createMonitorOptions.Path
	}
	if createMonitorOptions.HeadersVar != nil {
		body["headers"] = createMonitorOptions.HeadersVar
	}
	if createMonitorOptions.AllowInsecure != nil {
		body["allow_insecure"] = createMonitorOptions.AllowInsecure
	}
	if createMonitorOptions.ExpectedCodes != nil {
		body["expected_codes"] = createMonitorOptions.ExpectedCodes
	}
	if createMonitorOptions.ExpectedBody != nil {
		body["expected_body"] = createMonitorOptions.ExpectedBody
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitor)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteMonitor : Delete load balancer monitor
// Delete a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) DeleteMonitor(deleteMonitorOptions *DeleteMonitorOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeleteMonitorWithContext(context.Background(), deleteMonitorOptions)
}

// DeleteMonitorWithContext is an alternate form of the DeleteMonitor method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeleteMonitorWithContext(ctx context.Context, deleteMonitorOptions *DeleteMonitorOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteMonitorOptions, "deleteMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteMonitorOptions, "deleteMonitorOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteMonitorOptions.InstanceID,
		"monitor_id":  *deleteMonitorOptions.MonitorID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/monitors/{monitor_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteMonitorOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetMonitor : Get load balancer monitor
// Get details of a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) GetMonitor(getMonitorOptions *GetMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetMonitorWithContext(context.Background(), getMonitorOptions)
}

// GetMonitorWithContext is an alternate form of the GetMonitor method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetMonitorWithContext(ctx context.Context, getMonitorOptions *GetMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getMonitorOptions, "getMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getMonitorOptions, "getMonitorOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getMonitorOptions.InstanceID,
		"monitor_id":  *getMonitorOptions.MonitorID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/monitors/{monitor_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getMonitorOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitor)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateMonitor : Update load balancer monitor
// Update the properties of a load balancer monitor.
func (dnsSvcs *DnsSvcsV1) UpdateMonitor(updateMonitorOptions *UpdateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdateMonitorWithContext(context.Background(), updateMonitorOptions)
}

// UpdateMonitorWithContext is an alternate form of the UpdateMonitor method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdateMonitorWithContext(ctx context.Context, updateMonitorOptions *UpdateMonitorOptions) (result *Monitor, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateMonitorOptions, "updateMonitorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateMonitorOptions, "updateMonitorOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateMonitorOptions.InstanceID,
		"monitor_id":  *updateMonitorOptions.MonitorID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/monitors/{monitor_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateMonitorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateMonitor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateMonitorOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateMonitorOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateMonitorOptions.Name != nil {
		body["name"] = updateMonitorOptions.Name
	}
	if updateMonitorOptions.Description != nil {
		body["description"] = updateMonitorOptions.Description
	}
	if updateMonitorOptions.Type != nil {
		body["type"] = updateMonitorOptions.Type
	}
	if updateMonitorOptions.Port != nil {
		body["port"] = updateMonitorOptions.Port
	}
	if updateMonitorOptions.Interval != nil {
		body["interval"] = updateMonitorOptions.Interval
	}
	if updateMonitorOptions.Retries != nil {
		body["retries"] = updateMonitorOptions.Retries
	}
	if updateMonitorOptions.Timeout != nil {
		body["timeout"] = updateMonitorOptions.Timeout
	}
	if updateMonitorOptions.Method != nil {
		body["method"] = updateMonitorOptions.Method
	}
	if updateMonitorOptions.Path != nil {
		body["path"] = updateMonitorOptions.Path
	}
	if updateMonitorOptions.HeadersVar != nil {
		body["headers"] = updateMonitorOptions.HeadersVar
	}
	if updateMonitorOptions.AllowInsecure != nil {
		body["allow_insecure"] = updateMonitorOptions.AllowInsecure
	}
	if updateMonitorOptions.ExpectedCodes != nil {
		body["expected_codes"] = updateMonitorOptions.ExpectedCodes
	}
	if updateMonitorOptions.ExpectedBody != nil {
		body["expected_body"] = updateMonitorOptions.ExpectedBody
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMonitor)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCustomResolvers : List custom resolvers
// List the custom resolvers.
func (dnsSvcs *DnsSvcsV1) ListCustomResolvers(listCustomResolversOptions *ListCustomResolversOptions) (result *CustomResolverList, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListCustomResolversWithContext(context.Background(), listCustomResolversOptions)
}

// ListCustomResolversWithContext is an alternate form of the ListCustomResolvers method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListCustomResolversWithContext(ctx context.Context, listCustomResolversOptions *ListCustomResolversOptions) (result *CustomResolverList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCustomResolversOptions, "listCustomResolversOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCustomResolversOptions, "listCustomResolversOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listCustomResolversOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCustomResolversOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListCustomResolvers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listCustomResolversOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listCustomResolversOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomResolverList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCustomResolver : Create a custom resolver
// Create a custom resolver.
func (dnsSvcs *DnsSvcsV1) CreateCustomResolver(createCustomResolverOptions *CreateCustomResolverOptions) (result *CustomResolver, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreateCustomResolverWithContext(context.Background(), createCustomResolverOptions)
}

// CreateCustomResolverWithContext is an alternate form of the CreateCustomResolver method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreateCustomResolverWithContext(ctx context.Context, createCustomResolverOptions *CreateCustomResolverOptions) (result *CustomResolver, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCustomResolverOptions, "createCustomResolverOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCustomResolverOptions, "createCustomResolverOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createCustomResolverOptions.InstanceID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCustomResolverOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateCustomResolver")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createCustomResolverOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createCustomResolverOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createCustomResolverOptions.Name != nil {
		body["name"] = createCustomResolverOptions.Name
	}
	if createCustomResolverOptions.Description != nil {
		body["description"] = createCustomResolverOptions.Description
	}
	if createCustomResolverOptions.Locations != nil {
		body["locations"] = createCustomResolverOptions.Locations
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomResolver)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomResolver : Delete a custom resolver
// Delete a custom resolver.
func (dnsSvcs *DnsSvcsV1) DeleteCustomResolver(deleteCustomResolverOptions *DeleteCustomResolverOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeleteCustomResolverWithContext(context.Background(), deleteCustomResolverOptions)
}

// DeleteCustomResolverWithContext is an alternate form of the DeleteCustomResolver method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeleteCustomResolverWithContext(ctx context.Context, deleteCustomResolverOptions *DeleteCustomResolverOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomResolverOptions, "deleteCustomResolverOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomResolverOptions, "deleteCustomResolverOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteCustomResolverOptions.InstanceID,
		"resolver_id": *deleteCustomResolverOptions.ResolverID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomResolverOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteCustomResolver")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteCustomResolverOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteCustomResolverOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetCustomResolver : Get a custom resolver
// Get details of a custom resolver.
func (dnsSvcs *DnsSvcsV1) GetCustomResolver(getCustomResolverOptions *GetCustomResolverOptions) (result *CustomResolver, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetCustomResolverWithContext(context.Background(), getCustomResolverOptions)
}

// GetCustomResolverWithContext is an alternate form of the GetCustomResolver method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetCustomResolverWithContext(ctx context.Context, getCustomResolverOptions *GetCustomResolverOptions) (result *CustomResolver, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCustomResolverOptions, "getCustomResolverOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCustomResolverOptions, "getCustomResolverOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getCustomResolverOptions.InstanceID,
		"resolver_id": *getCustomResolverOptions.ResolverID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCustomResolverOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetCustomResolver")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getCustomResolverOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getCustomResolverOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomResolver)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCustomResolver : Update the properties of a custom resolver
// Update the properties of a custom resolver.
func (dnsSvcs *DnsSvcsV1) UpdateCustomResolver(updateCustomResolverOptions *UpdateCustomResolverOptions) (result *CustomResolver, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdateCustomResolverWithContext(context.Background(), updateCustomResolverOptions)
}

// UpdateCustomResolverWithContext is an alternate form of the UpdateCustomResolver method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdateCustomResolverWithContext(ctx context.Context, updateCustomResolverOptions *UpdateCustomResolverOptions) (result *CustomResolver, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCustomResolverOptions, "updateCustomResolverOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCustomResolverOptions, "updateCustomResolverOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateCustomResolverOptions.InstanceID,
		"resolver_id": *updateCustomResolverOptions.ResolverID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCustomResolverOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateCustomResolver")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateCustomResolverOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateCustomResolverOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateCustomResolverOptions.Name != nil {
		body["name"] = updateCustomResolverOptions.Name
	}
	if updateCustomResolverOptions.Description != nil {
		body["description"] = updateCustomResolverOptions.Description
	}
	if updateCustomResolverOptions.Enabled != nil {
		body["enabled"] = updateCustomResolverOptions.Enabled
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomResolver)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddCustomResolverLocation : Add custom resolver location
// Add custom resolver location.
func (dnsSvcs *DnsSvcsV1) AddCustomResolverLocation(addCustomResolverLocationOptions *AddCustomResolverLocationOptions) (result *Location, response *core.DetailedResponse, err error) {
	return dnsSvcs.AddCustomResolverLocationWithContext(context.Background(), addCustomResolverLocationOptions)
}

// AddCustomResolverLocationWithContext is an alternate form of the AddCustomResolverLocation method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) AddCustomResolverLocationWithContext(ctx context.Context, addCustomResolverLocationOptions *AddCustomResolverLocationOptions) (result *Location, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addCustomResolverLocationOptions, "addCustomResolverLocationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addCustomResolverLocationOptions, "addCustomResolverLocationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *addCustomResolverLocationOptions.InstanceID,
		"resolver_id": *addCustomResolverLocationOptions.ResolverID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/locations`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addCustomResolverLocationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "AddCustomResolverLocation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if addCustomResolverLocationOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*addCustomResolverLocationOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if addCustomResolverLocationOptions.SubnetCrn != nil {
		body["subnet_crn"] = addCustomResolverLocationOptions.SubnetCrn
	}
	if addCustomResolverLocationOptions.Enabled != nil {
		body["enabled"] = addCustomResolverLocationOptions.Enabled
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLocation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCustomResolverLocation : Update custom resolver location
// Update custom resolver location.
func (dnsSvcs *DnsSvcsV1) UpdateCustomResolverLocation(updateCustomResolverLocationOptions *UpdateCustomResolverLocationOptions) (result *Location, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdateCustomResolverLocationWithContext(context.Background(), updateCustomResolverLocationOptions)
}

// UpdateCustomResolverLocationWithContext is an alternate form of the UpdateCustomResolverLocation method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdateCustomResolverLocationWithContext(ctx context.Context, updateCustomResolverLocationOptions *UpdateCustomResolverLocationOptions) (result *Location, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCustomResolverLocationOptions, "updateCustomResolverLocationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCustomResolverLocationOptions, "updateCustomResolverLocationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateCustomResolverLocationOptions.InstanceID,
		"resolver_id": *updateCustomResolverLocationOptions.ResolverID,
		"location_id": *updateCustomResolverLocationOptions.LocationID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/locations/{location_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCustomResolverLocationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateCustomResolverLocation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateCustomResolverLocationOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateCustomResolverLocationOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateCustomResolverLocationOptions.Enabled != nil {
		body["enabled"] = updateCustomResolverLocationOptions.Enabled
	}
	if updateCustomResolverLocationOptions.SubnetCrn != nil {
		body["subnet_crn"] = updateCustomResolverLocationOptions.SubnetCrn
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalLocation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomResolverLocation : Delete custom resolver location
// Delete custom resolver location.
func (dnsSvcs *DnsSvcsV1) DeleteCustomResolverLocation(deleteCustomResolverLocationOptions *DeleteCustomResolverLocationOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeleteCustomResolverLocationWithContext(context.Background(), deleteCustomResolverLocationOptions)
}

// DeleteCustomResolverLocationWithContext is an alternate form of the DeleteCustomResolverLocation method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeleteCustomResolverLocationWithContext(ctx context.Context, deleteCustomResolverLocationOptions *DeleteCustomResolverLocationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomResolverLocationOptions, "deleteCustomResolverLocationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomResolverLocationOptions, "deleteCustomResolverLocationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteCustomResolverLocationOptions.InstanceID,
		"resolver_id": *deleteCustomResolverLocationOptions.ResolverID,
		"location_id": *deleteCustomResolverLocationOptions.LocationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/locations/{location_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomResolverLocationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteCustomResolverLocation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteCustomResolverLocationOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteCustomResolverLocationOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// ListForwardingRules : List forwarding rules
// List the forwarding rules of the given custom resolver.
func (dnsSvcs *DnsSvcsV1) ListForwardingRules(listForwardingRulesOptions *ListForwardingRulesOptions) (result *ForwardingRuleList, response *core.DetailedResponse, err error) {
	return dnsSvcs.ListForwardingRulesWithContext(context.Background(), listForwardingRulesOptions)
}

// ListForwardingRulesWithContext is an alternate form of the ListForwardingRules method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) ListForwardingRulesWithContext(ctx context.Context, listForwardingRulesOptions *ListForwardingRulesOptions) (result *ForwardingRuleList, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listForwardingRulesOptions, "listForwardingRulesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listForwardingRulesOptions, "listForwardingRulesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *listForwardingRulesOptions.InstanceID,
		"resolver_id": *listForwardingRulesOptions.ResolverID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/forwarding_rules`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listForwardingRulesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "ListForwardingRules")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listForwardingRulesOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*listForwardingRulesOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalForwardingRuleList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateForwardingRule : Create a forwarding rule
// Create a forwarding rule for the given custom resolver.
func (dnsSvcs *DnsSvcsV1) CreateForwardingRule(createForwardingRuleOptions *CreateForwardingRuleOptions) (result *ForwardingRule, response *core.DetailedResponse, err error) {
	return dnsSvcs.CreateForwardingRuleWithContext(context.Background(), createForwardingRuleOptions)
}

// CreateForwardingRuleWithContext is an alternate form of the CreateForwardingRule method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) CreateForwardingRuleWithContext(ctx context.Context, createForwardingRuleOptions *CreateForwardingRuleOptions) (result *ForwardingRule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createForwardingRuleOptions, "createForwardingRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createForwardingRuleOptions, "createForwardingRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *createForwardingRuleOptions.InstanceID,
		"resolver_id": *createForwardingRuleOptions.ResolverID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/forwarding_rules`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createForwardingRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "CreateForwardingRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createForwardingRuleOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*createForwardingRuleOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if createForwardingRuleOptions.Description != nil {
		body["description"] = createForwardingRuleOptions.Description
	}
	if createForwardingRuleOptions.Type != nil {
		body["type"] = createForwardingRuleOptions.Type
	}
	if createForwardingRuleOptions.Match != nil {
		body["match"] = createForwardingRuleOptions.Match
	}
	if createForwardingRuleOptions.ForwardTo != nil {
		body["forward_to"] = createForwardingRuleOptions.ForwardTo
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalForwardingRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteForwardingRule : Delete a forwarding rule
// Delete a forwarding rule on the given custom resolver.
func (dnsSvcs *DnsSvcsV1) DeleteForwardingRule(deleteForwardingRuleOptions *DeleteForwardingRuleOptions) (response *core.DetailedResponse, err error) {
	return dnsSvcs.DeleteForwardingRuleWithContext(context.Background(), deleteForwardingRuleOptions)
}

// DeleteForwardingRuleWithContext is an alternate form of the DeleteForwardingRule method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) DeleteForwardingRuleWithContext(ctx context.Context, deleteForwardingRuleOptions *DeleteForwardingRuleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteForwardingRuleOptions, "deleteForwardingRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteForwardingRuleOptions, "deleteForwardingRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *deleteForwardingRuleOptions.InstanceID,
		"resolver_id": *deleteForwardingRuleOptions.ResolverID,
		"rule_id":     *deleteForwardingRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/forwarding_rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteForwardingRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "DeleteForwardingRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	if deleteForwardingRuleOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*deleteForwardingRuleOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = dnsSvcs.Service.Request(request, nil)

	return
}

// GetForwardingRule : Get a forwarding rule
// Get details of a forwarding rule on the given custom resolver.
func (dnsSvcs *DnsSvcsV1) GetForwardingRule(getForwardingRuleOptions *GetForwardingRuleOptions) (result *ForwardingRule, response *core.DetailedResponse, err error) {
	return dnsSvcs.GetForwardingRuleWithContext(context.Background(), getForwardingRuleOptions)
}

// GetForwardingRuleWithContext is an alternate form of the GetForwardingRule method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) GetForwardingRuleWithContext(ctx context.Context, getForwardingRuleOptions *GetForwardingRuleOptions) (result *ForwardingRule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getForwardingRuleOptions, "getForwardingRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getForwardingRuleOptions, "getForwardingRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *getForwardingRuleOptions.InstanceID,
		"resolver_id": *getForwardingRuleOptions.ResolverID,
		"rule_id":     *getForwardingRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/forwarding_rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getForwardingRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "GetForwardingRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getForwardingRuleOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*getForwardingRuleOptions.XCorrelationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalForwardingRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateForwardingRule : Update the properties of a forwarding rule
// Update the properties of a forwarding rule on the given custom resolver.
func (dnsSvcs *DnsSvcsV1) UpdateForwardingRule(updateForwardingRuleOptions *UpdateForwardingRuleOptions) (result *ForwardingRule, response *core.DetailedResponse, err error) {
	return dnsSvcs.UpdateForwardingRuleWithContext(context.Background(), updateForwardingRuleOptions)
}

// UpdateForwardingRuleWithContext is an alternate form of the UpdateForwardingRule method which supports a Context parameter
func (dnsSvcs *DnsSvcsV1) UpdateForwardingRuleWithContext(ctx context.Context, updateForwardingRuleOptions *UpdateForwardingRuleOptions) (result *ForwardingRule, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateForwardingRuleOptions, "updateForwardingRuleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateForwardingRuleOptions, "updateForwardingRuleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"instance_id": *updateForwardingRuleOptions.InstanceID,
		"resolver_id": *updateForwardingRuleOptions.ResolverID,
		"rule_id":     *updateForwardingRuleOptions.RuleID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsSvcs.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsSvcs.Service.Options.URL, `/instances/{instance_id}/custom_resolvers/{resolver_id}/forwarding_rules/{rule_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateForwardingRuleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_svcs", "V1", "UpdateForwardingRule")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if updateForwardingRuleOptions.XCorrelationID != nil {
		builder.AddHeader("X-Correlation-ID", fmt.Sprint(*updateForwardingRuleOptions.XCorrelationID))
	}

	body := make(map[string]interface{})
	if updateForwardingRuleOptions.Description != nil {
		body["description"] = updateForwardingRuleOptions.Description
	}
	if updateForwardingRuleOptions.Match != nil {
		body["match"] = updateForwardingRuleOptions.Match
	}
	if updateForwardingRuleOptions.ForwardTo != nil {
		body["forward_to"] = updateForwardingRuleOptions.ForwardTo
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsSvcs.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalForwardingRule)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddCustomResolverLocationOptions : The AddCustomResolverLocation options.
type AddCustomResolverLocationOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Custom resolver location, subnet CRN.
	SubnetCrn *string `json:"subnet_crn,omitempty"`

	// Enable/Disable custom resolver location.
	Enabled *bool `json:"enabled,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddCustomResolverLocationOptions : Instantiate AddCustomResolverLocationOptions
func (*DnsSvcsV1) NewAddCustomResolverLocationOptions(instanceID string, resolverID string) *AddCustomResolverLocationOptions {
	return &AddCustomResolverLocationOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *AddCustomResolverLocationOptions) SetInstanceID(instanceID string) *AddCustomResolverLocationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *AddCustomResolverLocationOptions) SetResolverID(resolverID string) *AddCustomResolverLocationOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetSubnetCrn : Allow user to set SubnetCrn
func (_options *AddCustomResolverLocationOptions) SetSubnetCrn(subnetCrn string) *AddCustomResolverLocationOptions {
	_options.SubnetCrn = core.StringPtr(subnetCrn)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *AddCustomResolverLocationOptions) SetEnabled(enabled bool) *AddCustomResolverLocationOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *AddCustomResolverLocationOptions) SetXCorrelationID(xCorrelationID string) *AddCustomResolverLocationOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddCustomResolverLocationOptions) SetHeaders(param map[string]string) *AddCustomResolverLocationOptions {
	options.Headers = param
	return options
}

// CreateCustomResolverOptions : The CreateCustomResolver options.
type CreateCustomResolverOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// Name of the custom resolver.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the custom resolver.
	Description *string `json:"description,omitempty"`

	// Locations on which the custom resolver will be running.
	Locations []LocationInput `json:"locations,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCustomResolverOptions : Instantiate CreateCustomResolverOptions
func (*DnsSvcsV1) NewCreateCustomResolverOptions(instanceID string) *CreateCustomResolverOptions {
	return &CreateCustomResolverOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateCustomResolverOptions) SetInstanceID(instanceID string) *CreateCustomResolverOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateCustomResolverOptions) SetName(name string) *CreateCustomResolverOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateCustomResolverOptions) SetDescription(description string) *CreateCustomResolverOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetLocations : Allow user to set Locations
func (_options *CreateCustomResolverOptions) SetLocations(locations []LocationInput) *CreateCustomResolverOptions {
	_options.Locations = locations
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateCustomResolverOptions) SetXCorrelationID(xCorrelationID string) *CreateCustomResolverOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomResolverOptions) SetHeaders(param map[string]string) *CreateCustomResolverOptions {
	options.Headers = param
	return options
}

// CreateDnszoneOptions : The CreateDnszone options.
type CreateDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// Name of DNS zone.
	Name *string `json:"name,omitempty"`

	// The text describing the purpose of a DNS zone.
	Description *string `json:"description,omitempty"`

	// The label of a DNS zone.
	Label *string `json:"label,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDnszoneOptions : Instantiate CreateDnszoneOptions
func (*DnsSvcsV1) NewCreateDnszoneOptions(instanceID string) *CreateDnszoneOptions {
	return &CreateDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateDnszoneOptions) SetInstanceID(instanceID string) *CreateDnszoneOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateDnszoneOptions) SetName(name string) *CreateDnszoneOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateDnszoneOptions) SetDescription(description string) *CreateDnszoneOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetLabel : Allow user to set Label
func (_options *CreateDnszoneOptions) SetLabel(label string) *CreateDnszoneOptions {
	_options.Label = core.StringPtr(label)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateDnszoneOptions) SetXCorrelationID(xCorrelationID string) *CreateDnszoneOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDnszoneOptions) SetHeaders(param map[string]string) *CreateDnszoneOptions {
	options.Headers = param
	return options
}

// CreateForwardingRuleOptions : The CreateForwardingRule options.
type CreateForwardingRuleOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Descriptive text of the forwarding rule.
	Description *string `json:"description,omitempty"`

	// Type of the forwarding rule.
	Type *string `json:"type,omitempty"`

	// The matching zone or hostname.
	Match *string `json:"match,omitempty"`

	// The upstream DNS servers will be forwarded to.
	ForwardTo []string `json:"forward_to,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateForwardingRuleOptions.Type property.
// Type of the forwarding rule.
const (
	CreateForwardingRuleOptions_Type_Zone = "zone"
)

// NewCreateForwardingRuleOptions : Instantiate CreateForwardingRuleOptions
func (*DnsSvcsV1) NewCreateForwardingRuleOptions(instanceID string, resolverID string) *CreateForwardingRuleOptions {
	return &CreateForwardingRuleOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateForwardingRuleOptions) SetInstanceID(instanceID string) *CreateForwardingRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *CreateForwardingRuleOptions) SetResolverID(resolverID string) *CreateForwardingRuleOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateForwardingRuleOptions) SetDescription(description string) *CreateForwardingRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateForwardingRuleOptions) SetType(typeVar string) *CreateForwardingRuleOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetMatch : Allow user to set Match
func (_options *CreateForwardingRuleOptions) SetMatch(match string) *CreateForwardingRuleOptions {
	_options.Match = core.StringPtr(match)
	return _options
}

// SetForwardTo : Allow user to set ForwardTo
func (_options *CreateForwardingRuleOptions) SetForwardTo(forwardTo []string) *CreateForwardingRuleOptions {
	_options.ForwardTo = forwardTo
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateForwardingRuleOptions) SetXCorrelationID(xCorrelationID string) *CreateForwardingRuleOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateForwardingRuleOptions) SetHeaders(param map[string]string) *CreateForwardingRuleOptions {
	options.Headers = param
	return options
}

// CreateLoadBalancerOptions : The CreateLoadBalancer options.
type CreateLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool IDs.
	AzPools []LoadBalancerAzPoolsItem `json:"az_pools,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateLoadBalancerOptions : Instantiate CreateLoadBalancerOptions
func (*DnsSvcsV1) NewCreateLoadBalancerOptions(instanceID string, dnszoneID string) *CreateLoadBalancerOptions {
	return &CreateLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateLoadBalancerOptions) SetInstanceID(instanceID string) *CreateLoadBalancerOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *CreateLoadBalancerOptions) SetDnszoneID(dnszoneID string) *CreateLoadBalancerOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateLoadBalancerOptions) SetName(name string) *CreateLoadBalancerOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateLoadBalancerOptions) SetDescription(description string) *CreateLoadBalancerOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *CreateLoadBalancerOptions) SetEnabled(enabled bool) *CreateLoadBalancerOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetTTL : Allow user to set TTL
func (_options *CreateLoadBalancerOptions) SetTTL(ttl int64) *CreateLoadBalancerOptions {
	_options.TTL = core.Int64Ptr(ttl)
	return _options
}

// SetFallbackPool : Allow user to set FallbackPool
func (_options *CreateLoadBalancerOptions) SetFallbackPool(fallbackPool string) *CreateLoadBalancerOptions {
	_options.FallbackPool = core.StringPtr(fallbackPool)
	return _options
}

// SetDefaultPools : Allow user to set DefaultPools
func (_options *CreateLoadBalancerOptions) SetDefaultPools(defaultPools []string) *CreateLoadBalancerOptions {
	_options.DefaultPools = defaultPools
	return _options
}

// SetAzPools : Allow user to set AzPools
func (_options *CreateLoadBalancerOptions) SetAzPools(azPools []LoadBalancerAzPoolsItem) *CreateLoadBalancerOptions {
	_options.AzPools = azPools
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *CreateLoadBalancerOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateLoadBalancerOptions) SetHeaders(param map[string]string) *CreateLoadBalancerOptions {
	options.Headers = param
	return options
}

// CreateMonitorOptions : The CreateMonitor options.
type CreateMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The name of the load balancer monitor.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	HeadersVar []HealthcheckHeader `json:"headers,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTPS monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateMonitorOptions.Type property.
// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
const (
	CreateMonitorOptions_Type_Http  = "HTTP"
	CreateMonitorOptions_Type_Https = "HTTPS"
	CreateMonitorOptions_Type_Tcp   = "TCP"
)

// Constants associated with the CreateMonitorOptions.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	CreateMonitorOptions_Method_Get  = "GET"
	CreateMonitorOptions_Method_Head = "HEAD"
)

// NewCreateMonitorOptions : Instantiate CreateMonitorOptions
func (*DnsSvcsV1) NewCreateMonitorOptions(instanceID string) *CreateMonitorOptions {
	return &CreateMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateMonitorOptions) SetInstanceID(instanceID string) *CreateMonitorOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateMonitorOptions) SetName(name string) *CreateMonitorOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateMonitorOptions) SetDescription(description string) *CreateMonitorOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateMonitorOptions) SetType(typeVar string) *CreateMonitorOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetPort : Allow user to set Port
func (_options *CreateMonitorOptions) SetPort(port int64) *CreateMonitorOptions {
	_options.Port = core.Int64Ptr(port)
	return _options
}

// SetInterval : Allow user to set Interval
func (_options *CreateMonitorOptions) SetInterval(interval int64) *CreateMonitorOptions {
	_options.Interval = core.Int64Ptr(interval)
	return _options
}

// SetRetries : Allow user to set Retries
func (_options *CreateMonitorOptions) SetRetries(retries int64) *CreateMonitorOptions {
	_options.Retries = core.Int64Ptr(retries)
	return _options
}

// SetTimeout : Allow user to set Timeout
func (_options *CreateMonitorOptions) SetTimeout(timeout int64) *CreateMonitorOptions {
	_options.Timeout = core.Int64Ptr(timeout)
	return _options
}

// SetMethod : Allow user to set Method
func (_options *CreateMonitorOptions) SetMethod(method string) *CreateMonitorOptions {
	_options.Method = core.StringPtr(method)
	return _options
}

// SetPath : Allow user to set Path
func (_options *CreateMonitorOptions) SetPath(path string) *CreateMonitorOptions {
	_options.Path = core.StringPtr(path)
	return _options
}

// SetHeadersVar : Allow user to set HeadersVar
func (_options *CreateMonitorOptions) SetHeadersVar(headersVar []HealthcheckHeader) *CreateMonitorOptions {
	_options.HeadersVar = headersVar
	return _options
}

// SetAllowInsecure : Allow user to set AllowInsecure
func (_options *CreateMonitorOptions) SetAllowInsecure(allowInsecure bool) *CreateMonitorOptions {
	_options.AllowInsecure = core.BoolPtr(allowInsecure)
	return _options
}

// SetExpectedCodes : Allow user to set ExpectedCodes
func (_options *CreateMonitorOptions) SetExpectedCodes(expectedCodes string) *CreateMonitorOptions {
	_options.ExpectedCodes = core.StringPtr(expectedCodes)
	return _options
}

// SetExpectedBody : Allow user to set ExpectedBody
func (_options *CreateMonitorOptions) SetExpectedBody(expectedBody string) *CreateMonitorOptions {
	_options.ExpectedBody = core.StringPtr(expectedBody)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateMonitorOptions) SetXCorrelationID(xCorrelationID string) *CreateMonitorOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateMonitorOptions) SetHeaders(param map[string]string) *CreateMonitorOptions {
	options.Headers = param
	return options
}

// CreatePermittedNetworkOptions : The CreatePermittedNetwork options.
type CreatePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// Permitted network data for VPC.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePermittedNetworkOptions.Type property.
// The type of a permitted network.
const (
	CreatePermittedNetworkOptions_Type_Vpc = "vpc"
)

// NewCreatePermittedNetworkOptions : Instantiate CreatePermittedNetworkOptions
func (*DnsSvcsV1) NewCreatePermittedNetworkOptions(instanceID string, dnszoneID string) *CreatePermittedNetworkOptions {
	return &CreatePermittedNetworkOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreatePermittedNetworkOptions) SetInstanceID(instanceID string) *CreatePermittedNetworkOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *CreatePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *CreatePermittedNetworkOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreatePermittedNetworkOptions) SetType(typeVar string) *CreatePermittedNetworkOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetPermittedNetwork : Allow user to set PermittedNetwork
func (_options *CreatePermittedNetworkOptions) SetPermittedNetwork(permittedNetwork *PermittedNetworkVpc) *CreatePermittedNetworkOptions {
	_options.PermittedNetwork = permittedNetwork
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreatePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *CreatePermittedNetworkOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePermittedNetworkOptions) SetHeaders(param map[string]string) *CreatePermittedNetworkOptions {
	options.Headers = param
	return options
}

// CreatePoolOptions : The CreatePool options.
type CreatePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	HealthyOriginsThreshold *int64 `json:"healthy_origins_threshold,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []OriginInput `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Health check region of VSIs.
	HealthcheckRegion *string `json:"healthcheck_region,omitempty"`

	// Health check subnet CRN.
	HealthcheckSubnets []string `json:"healthcheck_subnets,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreatePoolOptions.HealthcheckRegion property.
// Health check region of VSIs.
const (
	CreatePoolOptions_HealthcheckRegion_AuSyd   = "au-syd"
	CreatePoolOptions_HealthcheckRegion_EuDu    = "eu-du"
	CreatePoolOptions_HealthcheckRegion_EuGb    = "eu-gb"
	CreatePoolOptions_HealthcheckRegion_JpTok   = "jp-tok"
	CreatePoolOptions_HealthcheckRegion_UsEast  = "us-east"
	CreatePoolOptions_HealthcheckRegion_UsSouth = "us-south"
)

// NewCreatePoolOptions : Instantiate CreatePoolOptions
func (*DnsSvcsV1) NewCreatePoolOptions(instanceID string) *CreatePoolOptions {
	return &CreatePoolOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreatePoolOptions) SetInstanceID(instanceID string) *CreatePoolOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreatePoolOptions) SetName(name string) *CreatePoolOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreatePoolOptions) SetDescription(description string) *CreatePoolOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *CreatePoolOptions) SetEnabled(enabled bool) *CreatePoolOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetHealthyOriginsThreshold : Allow user to set HealthyOriginsThreshold
func (_options *CreatePoolOptions) SetHealthyOriginsThreshold(healthyOriginsThreshold int64) *CreatePoolOptions {
	_options.HealthyOriginsThreshold = core.Int64Ptr(healthyOriginsThreshold)
	return _options
}

// SetOrigins : Allow user to set Origins
func (_options *CreatePoolOptions) SetOrigins(origins []OriginInput) *CreatePoolOptions {
	_options.Origins = origins
	return _options
}

// SetMonitor : Allow user to set Monitor
func (_options *CreatePoolOptions) SetMonitor(monitor string) *CreatePoolOptions {
	_options.Monitor = core.StringPtr(monitor)
	return _options
}

// SetNotificationChannel : Allow user to set NotificationChannel
func (_options *CreatePoolOptions) SetNotificationChannel(notificationChannel string) *CreatePoolOptions {
	_options.NotificationChannel = core.StringPtr(notificationChannel)
	return _options
}

// SetHealthcheckRegion : Allow user to set HealthcheckRegion
func (_options *CreatePoolOptions) SetHealthcheckRegion(healthcheckRegion string) *CreatePoolOptions {
	_options.HealthcheckRegion = core.StringPtr(healthcheckRegion)
	return _options
}

// SetHealthcheckSubnets : Allow user to set HealthcheckSubnets
func (_options *CreatePoolOptions) SetHealthcheckSubnets(healthcheckSubnets []string) *CreatePoolOptions {
	_options.HealthcheckSubnets = healthcheckSubnets
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreatePoolOptions) SetXCorrelationID(xCorrelationID string) *CreatePoolOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreatePoolOptions) SetHeaders(param map[string]string) *CreatePoolOptions {
	options.Headers = param
	return options
}

// CreateResourceRecordOptions : The CreateResourceRecord options.
type CreateResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Name of the resource record.
	Name *string `json:"name,omitempty"`

	// Type of the resource record.
	Type *string `json:"type,omitempty"`

	// Content of the resource record.
	Rdata ResourceRecordInputRdataIntf `json:"rdata,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Only used for SRV record.
	Service *string `json:"service,omitempty"`

	// Only used for SRV record.
	Protocol *string `json:"protocol,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateResourceRecordOptions.Type property.
// Type of the resource record.
const (
	CreateResourceRecordOptions_Type_A     = "A"
	CreateResourceRecordOptions_Type_Aaaa  = "AAAA"
	CreateResourceRecordOptions_Type_Cname = "CNAME"
	CreateResourceRecordOptions_Type_Mx    = "MX"
	CreateResourceRecordOptions_Type_Ptr   = "PTR"
	CreateResourceRecordOptions_Type_Srv   = "SRV"
	CreateResourceRecordOptions_Type_Txt   = "TXT"
)

// NewCreateResourceRecordOptions : Instantiate CreateResourceRecordOptions
func (*DnsSvcsV1) NewCreateResourceRecordOptions(instanceID string, dnszoneID string) *CreateResourceRecordOptions {
	return &CreateResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *CreateResourceRecordOptions) SetInstanceID(instanceID string) *CreateResourceRecordOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *CreateResourceRecordOptions) SetDnszoneID(dnszoneID string) *CreateResourceRecordOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateResourceRecordOptions) SetName(name string) *CreateResourceRecordOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetType : Allow user to set Type
func (_options *CreateResourceRecordOptions) SetType(typeVar string) *CreateResourceRecordOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetRdata : Allow user to set Rdata
func (_options *CreateResourceRecordOptions) SetRdata(rdata ResourceRecordInputRdataIntf) *CreateResourceRecordOptions {
	_options.Rdata = rdata
	return _options
}

// SetTTL : Allow user to set TTL
func (_options *CreateResourceRecordOptions) SetTTL(ttl int64) *CreateResourceRecordOptions {
	_options.TTL = core.Int64Ptr(ttl)
	return _options
}

// SetService : Allow user to set Service
func (_options *CreateResourceRecordOptions) SetService(service string) *CreateResourceRecordOptions {
	_options.Service = core.StringPtr(service)
	return _options
}

// SetProtocol : Allow user to set Protocol
func (_options *CreateResourceRecordOptions) SetProtocol(protocol string) *CreateResourceRecordOptions {
	_options.Protocol = core.StringPtr(protocol)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *CreateResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *CreateResourceRecordOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateResourceRecordOptions) SetHeaders(param map[string]string) *CreateResourceRecordOptions {
	options.Headers = param
	return options
}

// DeleteCustomResolverLocationOptions : The DeleteCustomResolverLocation options.
type DeleteCustomResolverLocationOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Custom resolver location ID.
	LocationID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomResolverLocationOptions : Instantiate DeleteCustomResolverLocationOptions
func (*DnsSvcsV1) NewDeleteCustomResolverLocationOptions(instanceID string, resolverID string, locationID string) *DeleteCustomResolverLocationOptions {
	return &DeleteCustomResolverLocationOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
		LocationID: core.StringPtr(locationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteCustomResolverLocationOptions) SetInstanceID(instanceID string) *DeleteCustomResolverLocationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *DeleteCustomResolverLocationOptions) SetResolverID(resolverID string) *DeleteCustomResolverLocationOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetLocationID : Allow user to set LocationID
func (_options *DeleteCustomResolverLocationOptions) SetLocationID(locationID string) *DeleteCustomResolverLocationOptions {
	_options.LocationID = core.StringPtr(locationID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteCustomResolverLocationOptions) SetXCorrelationID(xCorrelationID string) *DeleteCustomResolverLocationOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomResolverLocationOptions) SetHeaders(param map[string]string) *DeleteCustomResolverLocationOptions {
	options.Headers = param
	return options
}

// DeleteCustomResolverOptions : The DeleteCustomResolver options.
type DeleteCustomResolverOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomResolverOptions : Instantiate DeleteCustomResolverOptions
func (*DnsSvcsV1) NewDeleteCustomResolverOptions(instanceID string, resolverID string) *DeleteCustomResolverOptions {
	return &DeleteCustomResolverOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteCustomResolverOptions) SetInstanceID(instanceID string) *DeleteCustomResolverOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *DeleteCustomResolverOptions) SetResolverID(resolverID string) *DeleteCustomResolverOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteCustomResolverOptions) SetXCorrelationID(xCorrelationID string) *DeleteCustomResolverOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomResolverOptions) SetHeaders(param map[string]string) *DeleteCustomResolverOptions {
	options.Headers = param
	return options
}

// DeleteDnszoneOptions : The DeleteDnszone options.
type DeleteDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDnszoneOptions : Instantiate DeleteDnszoneOptions
func (*DnsSvcsV1) NewDeleteDnszoneOptions(instanceID string, dnszoneID string) *DeleteDnszoneOptions {
	return &DeleteDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteDnszoneOptions) SetInstanceID(instanceID string) *DeleteDnszoneOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *DeleteDnszoneOptions) SetDnszoneID(dnszoneID string) *DeleteDnszoneOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteDnszoneOptions) SetXCorrelationID(xCorrelationID string) *DeleteDnszoneOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDnszoneOptions) SetHeaders(param map[string]string) *DeleteDnszoneOptions {
	options.Headers = param
	return options
}

// DeleteForwardingRuleOptions : The DeleteForwardingRule options.
type DeleteForwardingRuleOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a forwarding rule.
	RuleID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteForwardingRuleOptions : Instantiate DeleteForwardingRuleOptions
func (*DnsSvcsV1) NewDeleteForwardingRuleOptions(instanceID string, resolverID string, ruleID string) *DeleteForwardingRuleOptions {
	return &DeleteForwardingRuleOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
		RuleID:     core.StringPtr(ruleID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteForwardingRuleOptions) SetInstanceID(instanceID string) *DeleteForwardingRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *DeleteForwardingRuleOptions) SetResolverID(resolverID string) *DeleteForwardingRuleOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *DeleteForwardingRuleOptions) SetRuleID(ruleID string) *DeleteForwardingRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteForwardingRuleOptions) SetXCorrelationID(xCorrelationID string) *DeleteForwardingRuleOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteForwardingRuleOptions) SetHeaders(param map[string]string) *DeleteForwardingRuleOptions {
	options.Headers = param
	return options
}

// DeleteLoadBalancerOptions : The DeleteLoadBalancer options.
type DeleteLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer.
	LbID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteLoadBalancerOptions : Instantiate DeleteLoadBalancerOptions
func (*DnsSvcsV1) NewDeleteLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *DeleteLoadBalancerOptions {
	return &DeleteLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		LbID:       core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteLoadBalancerOptions) SetInstanceID(instanceID string) *DeleteLoadBalancerOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *DeleteLoadBalancerOptions) SetDnszoneID(dnszoneID string) *DeleteLoadBalancerOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetLbID : Allow user to set LbID
func (_options *DeleteLoadBalancerOptions) SetLbID(lbID string) *DeleteLoadBalancerOptions {
	_options.LbID = core.StringPtr(lbID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *DeleteLoadBalancerOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteLoadBalancerOptions) SetHeaders(param map[string]string) *DeleteLoadBalancerOptions {
	options.Headers = param
	return options
}

// DeleteMonitorOptions : The DeleteMonitor options.
type DeleteMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteMonitorOptions : Instantiate DeleteMonitorOptions
func (*DnsSvcsV1) NewDeleteMonitorOptions(instanceID string, monitorID string) *DeleteMonitorOptions {
	return &DeleteMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID:  core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteMonitorOptions) SetInstanceID(instanceID string) *DeleteMonitorOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetMonitorID : Allow user to set MonitorID
func (_options *DeleteMonitorOptions) SetMonitorID(monitorID string) *DeleteMonitorOptions {
	_options.MonitorID = core.StringPtr(monitorID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteMonitorOptions) SetXCorrelationID(xCorrelationID string) *DeleteMonitorOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteMonitorOptions) SetHeaders(param map[string]string) *DeleteMonitorOptions {
	options.Headers = param
	return options
}

// DeletePermittedNetworkOptions : The DeletePermittedNetwork options.
type DeletePermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePermittedNetworkOptions : Instantiate DeletePermittedNetworkOptions
func (*DnsSvcsV1) NewDeletePermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *DeletePermittedNetworkOptions {
	return &DeletePermittedNetworkOptions{
		InstanceID:         core.StringPtr(instanceID),
		DnszoneID:          core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeletePermittedNetworkOptions) SetInstanceID(instanceID string) *DeletePermittedNetworkOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *DeletePermittedNetworkOptions) SetDnszoneID(dnszoneID string) *DeletePermittedNetworkOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (_options *DeletePermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *DeletePermittedNetworkOptions {
	_options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeletePermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *DeletePermittedNetworkOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePermittedNetworkOptions) SetHeaders(param map[string]string) *DeletePermittedNetworkOptions {
	options.Headers = param
	return options
}

// DeletePoolOptions : The DeletePool options.
type DeletePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeletePoolOptions : Instantiate DeletePoolOptions
func (*DnsSvcsV1) NewDeletePoolOptions(instanceID string, poolID string) *DeletePoolOptions {
	return &DeletePoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID:     core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeletePoolOptions) SetInstanceID(instanceID string) *DeletePoolOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetPoolID : Allow user to set PoolID
func (_options *DeletePoolOptions) SetPoolID(poolID string) *DeletePoolOptions {
	_options.PoolID = core.StringPtr(poolID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeletePoolOptions) SetXCorrelationID(xCorrelationID string) *DeletePoolOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeletePoolOptions) SetHeaders(param map[string]string) *DeletePoolOptions {
	options.Headers = param
	return options
}

// DeleteResourceRecordOptions : The DeleteResourceRecord options.
type DeleteResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a resource record.
	RecordID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteResourceRecordOptions : Instantiate DeleteResourceRecordOptions
func (*DnsSvcsV1) NewDeleteResourceRecordOptions(instanceID string, dnszoneID string, recordID string) *DeleteResourceRecordOptions {
	return &DeleteResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		RecordID:   core.StringPtr(recordID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *DeleteResourceRecordOptions) SetInstanceID(instanceID string) *DeleteResourceRecordOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *DeleteResourceRecordOptions) SetDnszoneID(dnszoneID string) *DeleteResourceRecordOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetRecordID : Allow user to set RecordID
func (_options *DeleteResourceRecordOptions) SetRecordID(recordID string) *DeleteResourceRecordOptions {
	_options.RecordID = core.StringPtr(recordID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *DeleteResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *DeleteResourceRecordOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteResourceRecordOptions) SetHeaders(param map[string]string) *DeleteResourceRecordOptions {
	options.Headers = param
	return options
}

// ExportResourceRecordsOptions : The ExportResourceRecords options.
type ExportResourceRecordsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewExportResourceRecordsOptions : Instantiate ExportResourceRecordsOptions
func (*DnsSvcsV1) NewExportResourceRecordsOptions(instanceID string, dnszoneID string) *ExportResourceRecordsOptions {
	return &ExportResourceRecordsOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ExportResourceRecordsOptions) SetInstanceID(instanceID string) *ExportResourceRecordsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *ExportResourceRecordsOptions) SetDnszoneID(dnszoneID string) *ExportResourceRecordsOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ExportResourceRecordsOptions) SetXCorrelationID(xCorrelationID string) *ExportResourceRecordsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ExportResourceRecordsOptions) SetHeaders(param map[string]string) *ExportResourceRecordsOptions {
	options.Headers = param
	return options
}

// GetCustomResolverOptions : The GetCustomResolver options.
type GetCustomResolverOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCustomResolverOptions : Instantiate GetCustomResolverOptions
func (*DnsSvcsV1) NewGetCustomResolverOptions(instanceID string, resolverID string) *GetCustomResolverOptions {
	return &GetCustomResolverOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetCustomResolverOptions) SetInstanceID(instanceID string) *GetCustomResolverOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *GetCustomResolverOptions) SetResolverID(resolverID string) *GetCustomResolverOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetCustomResolverOptions) SetXCorrelationID(xCorrelationID string) *GetCustomResolverOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomResolverOptions) SetHeaders(param map[string]string) *GetCustomResolverOptions {
	options.Headers = param
	return options
}

// GetDnszoneOptions : The GetDnszone options.
type GetDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDnszoneOptions : Instantiate GetDnszoneOptions
func (*DnsSvcsV1) NewGetDnszoneOptions(instanceID string, dnszoneID string) *GetDnszoneOptions {
	return &GetDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetDnszoneOptions) SetInstanceID(instanceID string) *GetDnszoneOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *GetDnszoneOptions) SetDnszoneID(dnszoneID string) *GetDnszoneOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetDnszoneOptions) SetXCorrelationID(xCorrelationID string) *GetDnszoneOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDnszoneOptions) SetHeaders(param map[string]string) *GetDnszoneOptions {
	options.Headers = param
	return options
}

// GetForwardingRuleOptions : The GetForwardingRule options.
type GetForwardingRuleOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a forwarding rule.
	RuleID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetForwardingRuleOptions : Instantiate GetForwardingRuleOptions
func (*DnsSvcsV1) NewGetForwardingRuleOptions(instanceID string, resolverID string, ruleID string) *GetForwardingRuleOptions {
	return &GetForwardingRuleOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
		RuleID:     core.StringPtr(ruleID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetForwardingRuleOptions) SetInstanceID(instanceID string) *GetForwardingRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *GetForwardingRuleOptions) SetResolverID(resolverID string) *GetForwardingRuleOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *GetForwardingRuleOptions) SetRuleID(ruleID string) *GetForwardingRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetForwardingRuleOptions) SetXCorrelationID(xCorrelationID string) *GetForwardingRuleOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetForwardingRuleOptions) SetHeaders(param map[string]string) *GetForwardingRuleOptions {
	options.Headers = param
	return options
}

// GetLoadBalancerOptions : The GetLoadBalancer options.
type GetLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer.
	LbID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLoadBalancerOptions : Instantiate GetLoadBalancerOptions
func (*DnsSvcsV1) NewGetLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *GetLoadBalancerOptions {
	return &GetLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		LbID:       core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetLoadBalancerOptions) SetInstanceID(instanceID string) *GetLoadBalancerOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *GetLoadBalancerOptions) SetDnszoneID(dnszoneID string) *GetLoadBalancerOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetLbID : Allow user to set LbID
func (_options *GetLoadBalancerOptions) SetLbID(lbID string) *GetLoadBalancerOptions {
	_options.LbID = core.StringPtr(lbID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *GetLoadBalancerOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetLoadBalancerOptions) SetHeaders(param map[string]string) *GetLoadBalancerOptions {
	options.Headers = param
	return options
}

// GetMonitorOptions : The GetMonitor options.
type GetMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMonitorOptions : Instantiate GetMonitorOptions
func (*DnsSvcsV1) NewGetMonitorOptions(instanceID string, monitorID string) *GetMonitorOptions {
	return &GetMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID:  core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetMonitorOptions) SetInstanceID(instanceID string) *GetMonitorOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetMonitorID : Allow user to set MonitorID
func (_options *GetMonitorOptions) SetMonitorID(monitorID string) *GetMonitorOptions {
	_options.MonitorID = core.StringPtr(monitorID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetMonitorOptions) SetXCorrelationID(xCorrelationID string) *GetMonitorOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetMonitorOptions) SetHeaders(param map[string]string) *GetMonitorOptions {
	options.Headers = param
	return options
}

// GetPermittedNetworkOptions : The GetPermittedNetwork options.
type GetPermittedNetworkOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a permitted network.
	PermittedNetworkID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPermittedNetworkOptions : Instantiate GetPermittedNetworkOptions
func (*DnsSvcsV1) NewGetPermittedNetworkOptions(instanceID string, dnszoneID string, permittedNetworkID string) *GetPermittedNetworkOptions {
	return &GetPermittedNetworkOptions{
		InstanceID:         core.StringPtr(instanceID),
		DnszoneID:          core.StringPtr(dnszoneID),
		PermittedNetworkID: core.StringPtr(permittedNetworkID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetPermittedNetworkOptions) SetInstanceID(instanceID string) *GetPermittedNetworkOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *GetPermittedNetworkOptions) SetDnszoneID(dnszoneID string) *GetPermittedNetworkOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetPermittedNetworkID : Allow user to set PermittedNetworkID
func (_options *GetPermittedNetworkOptions) SetPermittedNetworkID(permittedNetworkID string) *GetPermittedNetworkOptions {
	_options.PermittedNetworkID = core.StringPtr(permittedNetworkID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetPermittedNetworkOptions) SetXCorrelationID(xCorrelationID string) *GetPermittedNetworkOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPermittedNetworkOptions) SetHeaders(param map[string]string) *GetPermittedNetworkOptions {
	options.Headers = param
	return options
}

// GetPoolOptions : The GetPool options.
type GetPoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetPoolOptions : Instantiate GetPoolOptions
func (*DnsSvcsV1) NewGetPoolOptions(instanceID string, poolID string) *GetPoolOptions {
	return &GetPoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID:     core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetPoolOptions) SetInstanceID(instanceID string) *GetPoolOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetPoolID : Allow user to set PoolID
func (_options *GetPoolOptions) SetPoolID(poolID string) *GetPoolOptions {
	_options.PoolID = core.StringPtr(poolID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetPoolOptions) SetXCorrelationID(xCorrelationID string) *GetPoolOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPoolOptions) SetHeaders(param map[string]string) *GetPoolOptions {
	options.Headers = param
	return options
}

// GetResourceRecordOptions : The GetResourceRecord options.
type GetResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a resource record.
	RecordID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetResourceRecordOptions : Instantiate GetResourceRecordOptions
func (*DnsSvcsV1) NewGetResourceRecordOptions(instanceID string, dnszoneID string, recordID string) *GetResourceRecordOptions {
	return &GetResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		RecordID:   core.StringPtr(recordID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *GetResourceRecordOptions) SetInstanceID(instanceID string) *GetResourceRecordOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *GetResourceRecordOptions) SetDnszoneID(dnszoneID string) *GetResourceRecordOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetRecordID : Allow user to set RecordID
func (_options *GetResourceRecordOptions) SetRecordID(recordID string) *GetResourceRecordOptions {
	_options.RecordID = core.StringPtr(recordID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *GetResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *GetResourceRecordOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetResourceRecordOptions) SetHeaders(param map[string]string) *GetResourceRecordOptions {
	options.Headers = param
	return options
}

// ImportResourceRecordsOptions : The ImportResourceRecords options.
type ImportResourceRecordsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// file to upload.
	File io.ReadCloser `json:"-"`

	// The content type of file.
	FileContentType *string `json:"-"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewImportResourceRecordsOptions : Instantiate ImportResourceRecordsOptions
func (*DnsSvcsV1) NewImportResourceRecordsOptions(instanceID string, dnszoneID string) *ImportResourceRecordsOptions {
	return &ImportResourceRecordsOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ImportResourceRecordsOptions) SetInstanceID(instanceID string) *ImportResourceRecordsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *ImportResourceRecordsOptions) SetDnszoneID(dnszoneID string) *ImportResourceRecordsOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetFile : Allow user to set File
func (_options *ImportResourceRecordsOptions) SetFile(file io.ReadCloser) *ImportResourceRecordsOptions {
	_options.File = file
	return _options
}

// SetFileContentType : Allow user to set FileContentType
func (_options *ImportResourceRecordsOptions) SetFileContentType(fileContentType string) *ImportResourceRecordsOptions {
	_options.FileContentType = core.StringPtr(fileContentType)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ImportResourceRecordsOptions) SetXCorrelationID(xCorrelationID string) *ImportResourceRecordsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ImportResourceRecordsOptions) SetHeaders(param map[string]string) *ImportResourceRecordsOptions {
	options.Headers = param
	return options
}

// ListCustomResolversOptions : The ListCustomResolvers options.
type ListCustomResolversOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCustomResolversOptions : Instantiate ListCustomResolversOptions
func (*DnsSvcsV1) NewListCustomResolversOptions(instanceID string) *ListCustomResolversOptions {
	return &ListCustomResolversOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListCustomResolversOptions) SetInstanceID(instanceID string) *ListCustomResolversOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListCustomResolversOptions) SetXCorrelationID(xCorrelationID string) *ListCustomResolversOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCustomResolversOptions) SetHeaders(param map[string]string) *ListCustomResolversOptions {
	options.Headers = param
	return options
}

// ListDnszonesOptions : The ListDnszones options.
type ListDnszonesOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Specify how many resource records to skip over, the default value is 0.
	Offset *int64 `json:"-"`

	// Specify how many resource records are returned, the default value is 200.
	Limit *int64 `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDnszonesOptions : Instantiate ListDnszonesOptions
func (*DnsSvcsV1) NewListDnszonesOptions(instanceID string) *ListDnszonesOptions {
	return &ListDnszonesOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListDnszonesOptions) SetInstanceID(instanceID string) *ListDnszonesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListDnszonesOptions) SetXCorrelationID(xCorrelationID string) *ListDnszonesOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListDnszonesOptions) SetOffset(offset int64) *ListDnszonesOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListDnszonesOptions) SetLimit(limit int64) *ListDnszonesOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDnszonesOptions) SetHeaders(param map[string]string) *ListDnszonesOptions {
	options.Headers = param
	return options
}

// ListForwardingRulesOptions : The ListForwardingRules options.
type ListForwardingRulesOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListForwardingRulesOptions : Instantiate ListForwardingRulesOptions
func (*DnsSvcsV1) NewListForwardingRulesOptions(instanceID string, resolverID string) *ListForwardingRulesOptions {
	return &ListForwardingRulesOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListForwardingRulesOptions) SetInstanceID(instanceID string) *ListForwardingRulesOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *ListForwardingRulesOptions) SetResolverID(resolverID string) *ListForwardingRulesOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListForwardingRulesOptions) SetXCorrelationID(xCorrelationID string) *ListForwardingRulesOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListForwardingRulesOptions) SetHeaders(param map[string]string) *ListForwardingRulesOptions {
	options.Headers = param
	return options
}

// ListLoadBalancersOptions : The ListLoadBalancers options.
type ListLoadBalancersOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListLoadBalancersOptions : Instantiate ListLoadBalancersOptions
func (*DnsSvcsV1) NewListLoadBalancersOptions(instanceID string, dnszoneID string) *ListLoadBalancersOptions {
	return &ListLoadBalancersOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListLoadBalancersOptions) SetInstanceID(instanceID string) *ListLoadBalancersOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *ListLoadBalancersOptions) SetDnszoneID(dnszoneID string) *ListLoadBalancersOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListLoadBalancersOptions) SetXCorrelationID(xCorrelationID string) *ListLoadBalancersOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListLoadBalancersOptions) SetHeaders(param map[string]string) *ListLoadBalancersOptions {
	options.Headers = param
	return options
}

// ListMonitorsOptions : The ListMonitors options.
type ListMonitorsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListMonitorsOptions : Instantiate ListMonitorsOptions
func (*DnsSvcsV1) NewListMonitorsOptions(instanceID string) *ListMonitorsOptions {
	return &ListMonitorsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListMonitorsOptions) SetInstanceID(instanceID string) *ListMonitorsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListMonitorsOptions) SetXCorrelationID(xCorrelationID string) *ListMonitorsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListMonitorsOptions) SetHeaders(param map[string]string) *ListMonitorsOptions {
	options.Headers = param
	return options
}

// ListPermittedNetworksOptions : The ListPermittedNetworks options.
type ListPermittedNetworksOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Specify how many resource records to skip over, the default value is 0.
	Offset *int64 `json:"-"`

	// Specify how many resource records are returned, the default value is 200.
	Limit *int64 `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPermittedNetworksOptions : Instantiate ListPermittedNetworksOptions
func (*DnsSvcsV1) NewListPermittedNetworksOptions(instanceID string, dnszoneID string) *ListPermittedNetworksOptions {
	return &ListPermittedNetworksOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListPermittedNetworksOptions) SetInstanceID(instanceID string) *ListPermittedNetworksOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *ListPermittedNetworksOptions) SetDnszoneID(dnszoneID string) *ListPermittedNetworksOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListPermittedNetworksOptions) SetXCorrelationID(xCorrelationID string) *ListPermittedNetworksOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListPermittedNetworksOptions) SetOffset(offset int64) *ListPermittedNetworksOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListPermittedNetworksOptions) SetLimit(limit int64) *ListPermittedNetworksOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPermittedNetworksOptions) SetHeaders(param map[string]string) *ListPermittedNetworksOptions {
	options.Headers = param
	return options
}

// ListPoolsOptions : The ListPools options.
type ListPoolsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPoolsOptions : Instantiate ListPoolsOptions
func (*DnsSvcsV1) NewListPoolsOptions(instanceID string) *ListPoolsOptions {
	return &ListPoolsOptions{
		InstanceID: core.StringPtr(instanceID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListPoolsOptions) SetInstanceID(instanceID string) *ListPoolsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListPoolsOptions) SetXCorrelationID(xCorrelationID string) *ListPoolsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPoolsOptions) SetHeaders(param map[string]string) *ListPoolsOptions {
	options.Headers = param
	return options
}

// ListResourceRecordsOptions : The ListResourceRecords options.
type ListResourceRecordsOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Specify how many resource records to skip over, the default value is 0.
	Offset *int64 `json:"-"`

	// Specify how many resource records are returned, the default value is 200.
	Limit *int64 `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListResourceRecordsOptions : Instantiate ListResourceRecordsOptions
func (*DnsSvcsV1) NewListResourceRecordsOptions(instanceID string, dnszoneID string) *ListResourceRecordsOptions {
	return &ListResourceRecordsOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *ListResourceRecordsOptions) SetInstanceID(instanceID string) *ListResourceRecordsOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *ListResourceRecordsOptions) SetDnszoneID(dnszoneID string) *ListResourceRecordsOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *ListResourceRecordsOptions) SetXCorrelationID(xCorrelationID string) *ListResourceRecordsOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetOffset : Allow user to set Offset
func (_options *ListResourceRecordsOptions) SetOffset(offset int64) *ListResourceRecordsOptions {
	_options.Offset = core.Int64Ptr(offset)
	return _options
}

// SetLimit : Allow user to set Limit
func (_options *ListResourceRecordsOptions) SetLimit(limit int64) *ListResourceRecordsOptions {
	_options.Limit = core.Int64Ptr(limit)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListResourceRecordsOptions) SetHeaders(param map[string]string) *ListResourceRecordsOptions {
	options.Headers = param
	return options
}

// LoadBalancerAzPoolsItem : LoadBalancerAzPoolsItem struct
type LoadBalancerAzPoolsItem struct {
	// Availability zone.
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// List of load balancer pools.
	Pools []string `json:"pools,omitempty"`
}

// UnmarshalLoadBalancerAzPoolsItem unmarshals an instance of LoadBalancerAzPoolsItem from the specified map of raw messages.
func UnmarshalLoadBalancerAzPoolsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoadBalancerAzPoolsItem)
	err = core.UnmarshalPrimitive(m, "availability_zone", &obj.AvailabilityZone)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pools", &obj.Pools)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PoolHealthcheckVsisItem : PoolHealthcheckVsisItem struct
type PoolHealthcheckVsisItem struct {
	// Health check VSI subnet CRN.
	Subnet *string `json:"subnet,omitempty"`

	// healthcheck VSI ip address.
	Ipv4Address *string `json:"ipv4_address,omitempty"`

	// ipv4 cidr block.
	Ipv4CidrBlock *string `json:"ipv4_cidr_block,omitempty"`

	// vpc crn.
	Vpc *string `json:"vpc,omitempty"`
}

// UnmarshalPoolHealthcheckVsisItem unmarshals an instance of PoolHealthcheckVsisItem from the specified map of raw messages.
func UnmarshalPoolHealthcheckVsisItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PoolHealthcheckVsisItem)
	err = core.UnmarshalPrimitive(m, "subnet", &obj.Subnet)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ipv4_address", &obj.Ipv4Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ipv4_cidr_block", &obj.Ipv4CidrBlock)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "vpc", &obj.Vpc)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsImportErrorModelError : RecordsImportErrorModelError struct
type RecordsImportErrorModelError struct {
	// Internal service error when DNS resource created fails by internal error.
	Code *string `json:"code" validate:"required"`

	// An internal error occurred. Try again later.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalRecordsImportErrorModelError unmarshals an instance of RecordsImportErrorModelError from the specified map of raw messages.
func UnmarshalRecordsImportErrorModelError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordsImportErrorModelError)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdata : Content of the resource record.
// Models which "extend" this model:
// - ResourceRecordInputRdataRdataARecord
// - ResourceRecordInputRdataRdataAaaaRecord
// - ResourceRecordInputRdataRdataCnameRecord
// - ResourceRecordInputRdataRdataMxRecord
// - ResourceRecordInputRdataRdataSrvRecord
// - ResourceRecordInputRdataRdataTxtRecord
// - ResourceRecordInputRdataRdataPtrRecord
type ResourceRecordInputRdata struct {
	// IPv4 address.
	Ip *string `json:"ip,omitempty"`

	// Canonical name.
	Cname *string `json:"cname,omitempty"`

	// Hostname of Exchange server.
	Exchange *string `json:"exchange,omitempty"`

	// Preference of the MX record.
	Preference *int64 `json:"preference,omitempty"`

	// Port number of the target server.
	Port *int64 `json:"port,omitempty"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority,omitempty"`

	// Hostname of the target server.
	Target *string `json:"target,omitempty"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight,omitempty"`

	// Human readable text.
	Text *string `json:"text,omitempty"`

	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

func (*ResourceRecordInputRdata) isaResourceRecordInputRdata() bool {
	return true
}

type ResourceRecordInputRdataIntf interface {
	isaResourceRecordInputRdata() bool
}

// UnmarshalResourceRecordInputRdata unmarshals an instance of ResourceRecordInputRdata from the specified map of raw messages.
func UnmarshalResourceRecordInputRdata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdata)
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cname", &obj.Cname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "exchange", &obj.Exchange)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preference", &obj.Preference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "priority", &obj.Priority)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "weight", &obj.Weight)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ptrdname", &obj.Ptrdname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdata : Content of the resource record.
// Models which "extend" this model:
// - ResourceRecordUpdateInputRdataRdataARecord
// - ResourceRecordUpdateInputRdataRdataAaaaRecord
// - ResourceRecordUpdateInputRdataRdataCnameRecord
// - ResourceRecordUpdateInputRdataRdataMxRecord
// - ResourceRecordUpdateInputRdataRdataSrvRecord
// - ResourceRecordUpdateInputRdataRdataTxtRecord
// - ResourceRecordUpdateInputRdataRdataPtrRecord
type ResourceRecordUpdateInputRdata struct {
	// IPv4 address.
	Ip *string `json:"ip,omitempty"`

	// Canonical name.
	Cname *string `json:"cname,omitempty"`

	// Hostname of Exchange server.
	Exchange *string `json:"exchange,omitempty"`

	// Preference of the MX record.
	Preference *int64 `json:"preference,omitempty"`

	// Port number of the target server.
	Port *int64 `json:"port,omitempty"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority,omitempty"`

	// Hostname of the target server.
	Target *string `json:"target,omitempty"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight,omitempty"`

	// Human readable text.
	Text *string `json:"text,omitempty"`

	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname,omitempty"`
}

func (*ResourceRecordUpdateInputRdata) isaResourceRecordUpdateInputRdata() bool {
	return true
}

type ResourceRecordUpdateInputRdataIntf interface {
	isaResourceRecordUpdateInputRdata() bool
}

// UnmarshalResourceRecordUpdateInputRdata unmarshals an instance of ResourceRecordUpdateInputRdata from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdata)
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cname", &obj.Cname)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "exchange", &obj.Exchange)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preference", &obj.Preference)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "priority", &obj.Priority)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "weight", &obj.Weight)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ptrdname", &obj.Ptrdname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCustomResolverLocationOptions : The UpdateCustomResolverLocation options.
type UpdateCustomResolverLocationOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Custom resolver location ID.
	LocationID *string `json:"-" validate:"required,ne="`

	// Enable/Disable custom resolver location.
	Enabled *bool `json:"enabled,omitempty"`

	// Subnet CRN.
	SubnetCrn *string `json:"subnet_crn,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCustomResolverLocationOptions : Instantiate UpdateCustomResolverLocationOptions
func (*DnsSvcsV1) NewUpdateCustomResolverLocationOptions(instanceID string, resolverID string, locationID string) *UpdateCustomResolverLocationOptions {
	return &UpdateCustomResolverLocationOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
		LocationID: core.StringPtr(locationID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateCustomResolverLocationOptions) SetInstanceID(instanceID string) *UpdateCustomResolverLocationOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *UpdateCustomResolverLocationOptions) SetResolverID(resolverID string) *UpdateCustomResolverLocationOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetLocationID : Allow user to set LocationID
func (_options *UpdateCustomResolverLocationOptions) SetLocationID(locationID string) *UpdateCustomResolverLocationOptions {
	_options.LocationID = core.StringPtr(locationID)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *UpdateCustomResolverLocationOptions) SetEnabled(enabled bool) *UpdateCustomResolverLocationOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetSubnetCrn : Allow user to set SubnetCrn
func (_options *UpdateCustomResolverLocationOptions) SetSubnetCrn(subnetCrn string) *UpdateCustomResolverLocationOptions {
	_options.SubnetCrn = core.StringPtr(subnetCrn)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateCustomResolverLocationOptions) SetXCorrelationID(xCorrelationID string) *UpdateCustomResolverLocationOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCustomResolverLocationOptions) SetHeaders(param map[string]string) *UpdateCustomResolverLocationOptions {
	options.Headers = param
	return options
}

// UpdateCustomResolverOptions : The UpdateCustomResolver options.
type UpdateCustomResolverOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// Name of the custom resolver.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the custom resolver.
	Description *string `json:"description,omitempty"`

	// Whether the custom resolver is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCustomResolverOptions : Instantiate UpdateCustomResolverOptions
func (*DnsSvcsV1) NewUpdateCustomResolverOptions(instanceID string, resolverID string) *UpdateCustomResolverOptions {
	return &UpdateCustomResolverOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateCustomResolverOptions) SetInstanceID(instanceID string) *UpdateCustomResolverOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *UpdateCustomResolverOptions) SetResolverID(resolverID string) *UpdateCustomResolverOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateCustomResolverOptions) SetName(name string) *UpdateCustomResolverOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateCustomResolverOptions) SetDescription(description string) *UpdateCustomResolverOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *UpdateCustomResolverOptions) SetEnabled(enabled bool) *UpdateCustomResolverOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateCustomResolverOptions) SetXCorrelationID(xCorrelationID string) *UpdateCustomResolverOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCustomResolverOptions) SetHeaders(param map[string]string) *UpdateCustomResolverOptions {
	options.Headers = param
	return options
}

// UpdateDnszoneOptions : The UpdateDnszone options.
type UpdateDnszoneOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The text describing the purpose of a DNS zone.
	Description *string `json:"description,omitempty"`

	// The label of a DNS zone.
	Label *string `json:"label,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDnszoneOptions : Instantiate UpdateDnszoneOptions
func (*DnsSvcsV1) NewUpdateDnszoneOptions(instanceID string, dnszoneID string) *UpdateDnszoneOptions {
	return &UpdateDnszoneOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateDnszoneOptions) SetInstanceID(instanceID string) *UpdateDnszoneOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *UpdateDnszoneOptions) SetDnszoneID(dnszoneID string) *UpdateDnszoneOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateDnszoneOptions) SetDescription(description string) *UpdateDnszoneOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetLabel : Allow user to set Label
func (_options *UpdateDnszoneOptions) SetLabel(label string) *UpdateDnszoneOptions {
	_options.Label = core.StringPtr(label)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateDnszoneOptions) SetXCorrelationID(xCorrelationID string) *UpdateDnszoneOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDnszoneOptions) SetHeaders(param map[string]string) *UpdateDnszoneOptions {
	options.Headers = param
	return options
}

// UpdateForwardingRuleOptions : The UpdateForwardingRule options.
type UpdateForwardingRuleOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a custom resolver.
	ResolverID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a forwarding rule.
	RuleID *string `json:"-" validate:"required,ne="`

	// Descriptive text of the forwarding rule.
	Description *string `json:"description,omitempty"`

	// The matching zone or hostname.
	Match *string `json:"match,omitempty"`

	// The upstream DNS servers will be forwarded to.
	ForwardTo []string `json:"forward_to,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateForwardingRuleOptions : Instantiate UpdateForwardingRuleOptions
func (*DnsSvcsV1) NewUpdateForwardingRuleOptions(instanceID string, resolverID string, ruleID string) *UpdateForwardingRuleOptions {
	return &UpdateForwardingRuleOptions{
		InstanceID: core.StringPtr(instanceID),
		ResolverID: core.StringPtr(resolverID),
		RuleID:     core.StringPtr(ruleID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateForwardingRuleOptions) SetInstanceID(instanceID string) *UpdateForwardingRuleOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetResolverID : Allow user to set ResolverID
func (_options *UpdateForwardingRuleOptions) SetResolverID(resolverID string) *UpdateForwardingRuleOptions {
	_options.ResolverID = core.StringPtr(resolverID)
	return _options
}

// SetRuleID : Allow user to set RuleID
func (_options *UpdateForwardingRuleOptions) SetRuleID(ruleID string) *UpdateForwardingRuleOptions {
	_options.RuleID = core.StringPtr(ruleID)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateForwardingRuleOptions) SetDescription(description string) *UpdateForwardingRuleOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetMatch : Allow user to set Match
func (_options *UpdateForwardingRuleOptions) SetMatch(match string) *UpdateForwardingRuleOptions {
	_options.Match = core.StringPtr(match)
	return _options
}

// SetForwardTo : Allow user to set ForwardTo
func (_options *UpdateForwardingRuleOptions) SetForwardTo(forwardTo []string) *UpdateForwardingRuleOptions {
	_options.ForwardTo = forwardTo
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateForwardingRuleOptions) SetXCorrelationID(xCorrelationID string) *UpdateForwardingRuleOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateForwardingRuleOptions) SetHeaders(param map[string]string) *UpdateForwardingRuleOptions {
	options.Headers = param
	return options
}

// UpdateLoadBalancerOptions : The UpdateLoadBalancer options.
type UpdateLoadBalancerOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer.
	LbID *string `json:"-" validate:"required,ne="`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool IDs.
	AzPools []LoadBalancerAzPoolsItem `json:"az_pools,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateLoadBalancerOptions : Instantiate UpdateLoadBalancerOptions
func (*DnsSvcsV1) NewUpdateLoadBalancerOptions(instanceID string, dnszoneID string, lbID string) *UpdateLoadBalancerOptions {
	return &UpdateLoadBalancerOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		LbID:       core.StringPtr(lbID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateLoadBalancerOptions) SetInstanceID(instanceID string) *UpdateLoadBalancerOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *UpdateLoadBalancerOptions) SetDnszoneID(dnszoneID string) *UpdateLoadBalancerOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetLbID : Allow user to set LbID
func (_options *UpdateLoadBalancerOptions) SetLbID(lbID string) *UpdateLoadBalancerOptions {
	_options.LbID = core.StringPtr(lbID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateLoadBalancerOptions) SetName(name string) *UpdateLoadBalancerOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateLoadBalancerOptions) SetDescription(description string) *UpdateLoadBalancerOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *UpdateLoadBalancerOptions) SetEnabled(enabled bool) *UpdateLoadBalancerOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetTTL : Allow user to set TTL
func (_options *UpdateLoadBalancerOptions) SetTTL(ttl int64) *UpdateLoadBalancerOptions {
	_options.TTL = core.Int64Ptr(ttl)
	return _options
}

// SetFallbackPool : Allow user to set FallbackPool
func (_options *UpdateLoadBalancerOptions) SetFallbackPool(fallbackPool string) *UpdateLoadBalancerOptions {
	_options.FallbackPool = core.StringPtr(fallbackPool)
	return _options
}

// SetDefaultPools : Allow user to set DefaultPools
func (_options *UpdateLoadBalancerOptions) SetDefaultPools(defaultPools []string) *UpdateLoadBalancerOptions {
	_options.DefaultPools = defaultPools
	return _options
}

// SetAzPools : Allow user to set AzPools
func (_options *UpdateLoadBalancerOptions) SetAzPools(azPools []LoadBalancerAzPoolsItem) *UpdateLoadBalancerOptions {
	_options.AzPools = azPools
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateLoadBalancerOptions) SetXCorrelationID(xCorrelationID string) *UpdateLoadBalancerOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateLoadBalancerOptions) SetHeaders(param map[string]string) *UpdateLoadBalancerOptions {
	options.Headers = param
	return options
}

// UpdateMonitorOptions : The UpdateMonitor options.
type UpdateMonitorOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer monitor.
	MonitorID *string `json:"-" validate:"required,ne="`

	// The name of the load balancer monitor.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	HeadersVar []HealthcheckHeader `json:"headers,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTP and HTTPS
	// monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateMonitorOptions.Type property.
// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
const (
	UpdateMonitorOptions_Type_Http  = "HTTP"
	UpdateMonitorOptions_Type_Https = "HTTPS"
	UpdateMonitorOptions_Type_Tcp   = "TCP"
)

// Constants associated with the UpdateMonitorOptions.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	UpdateMonitorOptions_Method_Get  = "GET"
	UpdateMonitorOptions_Method_Head = "HEAD"
)

// NewUpdateMonitorOptions : Instantiate UpdateMonitorOptions
func (*DnsSvcsV1) NewUpdateMonitorOptions(instanceID string, monitorID string) *UpdateMonitorOptions {
	return &UpdateMonitorOptions{
		InstanceID: core.StringPtr(instanceID),
		MonitorID:  core.StringPtr(monitorID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateMonitorOptions) SetInstanceID(instanceID string) *UpdateMonitorOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetMonitorID : Allow user to set MonitorID
func (_options *UpdateMonitorOptions) SetMonitorID(monitorID string) *UpdateMonitorOptions {
	_options.MonitorID = core.StringPtr(monitorID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateMonitorOptions) SetName(name string) *UpdateMonitorOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateMonitorOptions) SetDescription(description string) *UpdateMonitorOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetType : Allow user to set Type
func (_options *UpdateMonitorOptions) SetType(typeVar string) *UpdateMonitorOptions {
	_options.Type = core.StringPtr(typeVar)
	return _options
}

// SetPort : Allow user to set Port
func (_options *UpdateMonitorOptions) SetPort(port int64) *UpdateMonitorOptions {
	_options.Port = core.Int64Ptr(port)
	return _options
}

// SetInterval : Allow user to set Interval
func (_options *UpdateMonitorOptions) SetInterval(interval int64) *UpdateMonitorOptions {
	_options.Interval = core.Int64Ptr(interval)
	return _options
}

// SetRetries : Allow user to set Retries
func (_options *UpdateMonitorOptions) SetRetries(retries int64) *UpdateMonitorOptions {
	_options.Retries = core.Int64Ptr(retries)
	return _options
}

// SetTimeout : Allow user to set Timeout
func (_options *UpdateMonitorOptions) SetTimeout(timeout int64) *UpdateMonitorOptions {
	_options.Timeout = core.Int64Ptr(timeout)
	return _options
}

// SetMethod : Allow user to set Method
func (_options *UpdateMonitorOptions) SetMethod(method string) *UpdateMonitorOptions {
	_options.Method = core.StringPtr(method)
	return _options
}

// SetPath : Allow user to set Path
func (_options *UpdateMonitorOptions) SetPath(path string) *UpdateMonitorOptions {
	_options.Path = core.StringPtr(path)
	return _options
}

// SetHeadersVar : Allow user to set HeadersVar
func (_options *UpdateMonitorOptions) SetHeadersVar(headersVar []HealthcheckHeader) *UpdateMonitorOptions {
	_options.HeadersVar = headersVar
	return _options
}

// SetAllowInsecure : Allow user to set AllowInsecure
func (_options *UpdateMonitorOptions) SetAllowInsecure(allowInsecure bool) *UpdateMonitorOptions {
	_options.AllowInsecure = core.BoolPtr(allowInsecure)
	return _options
}

// SetExpectedCodes : Allow user to set ExpectedCodes
func (_options *UpdateMonitorOptions) SetExpectedCodes(expectedCodes string) *UpdateMonitorOptions {
	_options.ExpectedCodes = core.StringPtr(expectedCodes)
	return _options
}

// SetExpectedBody : Allow user to set ExpectedBody
func (_options *UpdateMonitorOptions) SetExpectedBody(expectedBody string) *UpdateMonitorOptions {
	_options.ExpectedBody = core.StringPtr(expectedBody)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateMonitorOptions) SetXCorrelationID(xCorrelationID string) *UpdateMonitorOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMonitorOptions) SetHeaders(param map[string]string) *UpdateMonitorOptions {
	options.Headers = param
	return options
}

// UpdatePoolOptions : The UpdatePool options.
type UpdatePoolOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a load balancer pool.
	PoolID *string `json:"-" validate:"required,ne="`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	HealthyOriginsThreshold *int64 `json:"healthy_origins_threshold,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []OriginInput `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Health check region of VSIs.
	HealthcheckRegion *string `json:"healthcheck_region,omitempty"`

	// Health check subnet CRNs.
	HealthcheckSubnets []string `json:"healthcheck_subnets,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdatePoolOptions.HealthcheckRegion property.
// Health check region of VSIs.
const (
	UpdatePoolOptions_HealthcheckRegion_AuSyd   = "au-syd"
	UpdatePoolOptions_HealthcheckRegion_EuDu    = "eu-du"
	UpdatePoolOptions_HealthcheckRegion_EuGb    = "eu-gb"
	UpdatePoolOptions_HealthcheckRegion_JpTok   = "jp-tok"
	UpdatePoolOptions_HealthcheckRegion_UsEast  = "us-east"
	UpdatePoolOptions_HealthcheckRegion_UsSouth = "us-south"
)

// NewUpdatePoolOptions : Instantiate UpdatePoolOptions
func (*DnsSvcsV1) NewUpdatePoolOptions(instanceID string, poolID string) *UpdatePoolOptions {
	return &UpdatePoolOptions{
		InstanceID: core.StringPtr(instanceID),
		PoolID:     core.StringPtr(poolID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdatePoolOptions) SetInstanceID(instanceID string) *UpdatePoolOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetPoolID : Allow user to set PoolID
func (_options *UpdatePoolOptions) SetPoolID(poolID string) *UpdatePoolOptions {
	_options.PoolID = core.StringPtr(poolID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdatePoolOptions) SetName(name string) *UpdatePoolOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdatePoolOptions) SetDescription(description string) *UpdatePoolOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetEnabled : Allow user to set Enabled
func (_options *UpdatePoolOptions) SetEnabled(enabled bool) *UpdatePoolOptions {
	_options.Enabled = core.BoolPtr(enabled)
	return _options
}

// SetHealthyOriginsThreshold : Allow user to set HealthyOriginsThreshold
func (_options *UpdatePoolOptions) SetHealthyOriginsThreshold(healthyOriginsThreshold int64) *UpdatePoolOptions {
	_options.HealthyOriginsThreshold = core.Int64Ptr(healthyOriginsThreshold)
	return _options
}

// SetOrigins : Allow user to set Origins
func (_options *UpdatePoolOptions) SetOrigins(origins []OriginInput) *UpdatePoolOptions {
	_options.Origins = origins
	return _options
}

// SetMonitor : Allow user to set Monitor
func (_options *UpdatePoolOptions) SetMonitor(monitor string) *UpdatePoolOptions {
	_options.Monitor = core.StringPtr(monitor)
	return _options
}

// SetNotificationChannel : Allow user to set NotificationChannel
func (_options *UpdatePoolOptions) SetNotificationChannel(notificationChannel string) *UpdatePoolOptions {
	_options.NotificationChannel = core.StringPtr(notificationChannel)
	return _options
}

// SetHealthcheckRegion : Allow user to set HealthcheckRegion
func (_options *UpdatePoolOptions) SetHealthcheckRegion(healthcheckRegion string) *UpdatePoolOptions {
	_options.HealthcheckRegion = core.StringPtr(healthcheckRegion)
	return _options
}

// SetHealthcheckSubnets : Allow user to set HealthcheckSubnets
func (_options *UpdatePoolOptions) SetHealthcheckSubnets(healthcheckSubnets []string) *UpdatePoolOptions {
	_options.HealthcheckSubnets = healthcheckSubnets
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdatePoolOptions) SetXCorrelationID(xCorrelationID string) *UpdatePoolOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdatePoolOptions) SetHeaders(param map[string]string) *UpdatePoolOptions {
	options.Headers = param
	return options
}

// UpdateResourceRecordOptions : The UpdateResourceRecord options.
type UpdateResourceRecordOptions struct {
	// The unique identifier of a service instance.
	InstanceID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a DNS zone.
	DnszoneID *string `json:"-" validate:"required,ne="`

	// The unique identifier of a resource record.
	RecordID *string `json:"-" validate:"required,ne="`

	// Name of the resource record.
	Name *string `json:"name,omitempty"`

	// Content of the resource record.
	Rdata ResourceRecordUpdateInputRdataIntf `json:"rdata,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Only used for SRV record.
	Service *string `json:"service,omitempty"`

	// Only used for SRV record.
	Protocol *string `json:"protocol,omitempty"`

	// Uniquely identifying a request.
	XCorrelationID *string `json:"-"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateResourceRecordOptions : Instantiate UpdateResourceRecordOptions
func (*DnsSvcsV1) NewUpdateResourceRecordOptions(instanceID string, dnszoneID string, recordID string) *UpdateResourceRecordOptions {
	return &UpdateResourceRecordOptions{
		InstanceID: core.StringPtr(instanceID),
		DnszoneID:  core.StringPtr(dnszoneID),
		RecordID:   core.StringPtr(recordID),
	}
}

// SetInstanceID : Allow user to set InstanceID
func (_options *UpdateResourceRecordOptions) SetInstanceID(instanceID string) *UpdateResourceRecordOptions {
	_options.InstanceID = core.StringPtr(instanceID)
	return _options
}

// SetDnszoneID : Allow user to set DnszoneID
func (_options *UpdateResourceRecordOptions) SetDnszoneID(dnszoneID string) *UpdateResourceRecordOptions {
	_options.DnszoneID = core.StringPtr(dnszoneID)
	return _options
}

// SetRecordID : Allow user to set RecordID
func (_options *UpdateResourceRecordOptions) SetRecordID(recordID string) *UpdateResourceRecordOptions {
	_options.RecordID = core.StringPtr(recordID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateResourceRecordOptions) SetName(name string) *UpdateResourceRecordOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetRdata : Allow user to set Rdata
func (_options *UpdateResourceRecordOptions) SetRdata(rdata ResourceRecordUpdateInputRdataIntf) *UpdateResourceRecordOptions {
	_options.Rdata = rdata
	return _options
}

// SetTTL : Allow user to set TTL
func (_options *UpdateResourceRecordOptions) SetTTL(ttl int64) *UpdateResourceRecordOptions {
	_options.TTL = core.Int64Ptr(ttl)
	return _options
}

// SetService : Allow user to set Service
func (_options *UpdateResourceRecordOptions) SetService(service string) *UpdateResourceRecordOptions {
	_options.Service = core.StringPtr(service)
	return _options
}

// SetProtocol : Allow user to set Protocol
func (_options *UpdateResourceRecordOptions) SetProtocol(protocol string) *UpdateResourceRecordOptions {
	_options.Protocol = core.StringPtr(protocol)
	return _options
}

// SetXCorrelationID : Allow user to set XCorrelationID
func (_options *UpdateResourceRecordOptions) SetXCorrelationID(xCorrelationID string) *UpdateResourceRecordOptions {
	_options.XCorrelationID = core.StringPtr(xCorrelationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateResourceRecordOptions) SetHeaders(param map[string]string) *UpdateResourceRecordOptions {
	options.Headers = param
	return options
}

// CustomResolver : custom resolver details.
type CustomResolver struct {
	// Identifier of the custom resolver.
	ID *string `json:"id,omitempty"`

	// Name of the custom resolver.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the custom resolver.
	Description *string `json:"description,omitempty"`

	// Whether the custom resolver is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Healthy state of the custom resolver.
	Health *string `json:"health,omitempty"`

	// Locations on which the custom resolver will be running.
	Locations []Location `json:"locations,omitempty"`

	// the time when a custom resolver is created, RFC3339 format.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// the recent time when a custom resolver is modified, RFC3339 format.
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`
}

// Constants associated with the CustomResolver.Health property.
// Healthy state of the custom resolver.
const (
	CustomResolver_Health_Critical = "CRITICAL"
	CustomResolver_Health_Degraded = "DEGRADED"
	CustomResolver_Health_Healthy  = "HEALTHY"
)

// UnmarshalCustomResolver unmarshals an instance of CustomResolver from the specified map of raw messages.
func UnmarshalCustomResolver(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomResolver)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health", &obj.Health)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "locations", &obj.Locations, UnmarshalLocation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomResolverList : List custom resolvers response.
type CustomResolverList struct {
	// An array of custom resolvers.
	CustomResolvers []CustomResolver `json:"custom_resolvers,omitempty"`
}

// UnmarshalCustomResolverList unmarshals an instance of CustomResolverList from the specified map of raw messages.
func UnmarshalCustomResolverList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomResolverList)
	err = core.UnmarshalModel(m, "custom_resolvers", &obj.CustomResolvers, UnmarshalCustomResolver)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Dnszone : DNS zone details.
type Dnszone struct {
	// Unique identifier of a DNS zone.
	ID *string `json:"id,omitempty"`

	// the time when a DNS zone is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a DNS zone is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Unique identifier of a service instance.
	InstanceID *string `json:"instance_id,omitempty"`

	// Name of DNS zone.
	Name *string `json:"name,omitempty"`

	// The text describing the purpose of a DNS zone.
	Description *string `json:"description,omitempty"`

	// State of DNS zone.
	State *string `json:"state,omitempty"`

	// The label of a DNS zone.
	Label *string `json:"label,omitempty"`
}

// Constants associated with the Dnszone.State property.
// State of DNS zone.
const (
	Dnszone_State_Active            = "active"
	Dnszone_State_Deleted           = "deleted"
	Dnszone_State_Disabled          = "disabled"
	Dnszone_State_PendingDelete     = "pending_delete"
	Dnszone_State_PendingNetworkAdd = "pending_network_add"
)

// UnmarshalDnszone unmarshals an instance of Dnszone from the specified map of raw messages.
func UnmarshalDnszone(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Dnszone)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_id", &obj.InstanceID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "label", &obj.Label)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FirstHref : href.
type FirstHref struct {
	// href.
	Href *string `json:"href,omitempty"`
}

// UnmarshalFirstHref unmarshals an instance of FirstHref from the specified map of raw messages.
func UnmarshalFirstHref(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FirstHref)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ForwardingRule : forwarding rule details.
type ForwardingRule struct {
	// Identifier of the forwarding rule.
	ID *string `json:"id,omitempty"`

	// Descriptive text of the forwarding rule.
	Description *string `json:"description,omitempty"`

	// Type of the forwarding rule.
	Type *string `json:"type,omitempty"`

	// The matching zone or hostname.
	Match *string `json:"match,omitempty"`

	// The upstream DNS servers will be forwarded to.
	ForwardTo []string `json:"forward_to,omitempty"`

	// the time when a forwarding rule is created, RFC3339 format.
	CreatedOn *strfmt.DateTime `json:"created_on,omitempty"`

	// the recent time when a forwarding rule is modified, RFC3339 format.
	ModifiedOn *strfmt.DateTime `json:"modified_on,omitempty"`
}

// Constants associated with the ForwardingRule.Type property.
// Type of the forwarding rule.
const (
	ForwardingRule_Type_Default = "default"
	ForwardingRule_Type_Zone    = "zone"
)

// UnmarshalForwardingRule unmarshals an instance of ForwardingRule from the specified map of raw messages.
func UnmarshalForwardingRule(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ForwardingRule)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "match", &obj.Match)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "forward_to", &obj.ForwardTo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ForwardingRuleList : An array of forwarding rules.
type ForwardingRuleList struct {
	ForwardingRules []ForwardingRule `json:"forwarding_rules,omitempty"`
}

// UnmarshalForwardingRuleList unmarshals an instance of ForwardingRuleList from the specified map of raw messages.
func UnmarshalForwardingRuleList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ForwardingRuleList)
	err = core.UnmarshalModel(m, "forwarding_rules", &obj.ForwardingRules, UnmarshalForwardingRule)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// HealthcheckHeader : The HTTP header of health check request.
type HealthcheckHeader struct {
	// The name of HTTP request header.
	Name *string `json:"name" validate:"required"`

	// The value of HTTP request header.
	Value []string `json:"value" validate:"required"`
}

// NewHealthcheckHeader : Instantiate HealthcheckHeader (Generic Model Constructor)
func (*DnsSvcsV1) NewHealthcheckHeader(name string, value []string) (_model *HealthcheckHeader, err error) {
	_model = &HealthcheckHeader{
		Name:  core.StringPtr(name),
		Value: value,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalHealthcheckHeader unmarshals an instance of HealthcheckHeader from the specified map of raw messages.
func UnmarshalHealthcheckHeader(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HealthcheckHeader)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportResourceRecordsResp : Import DNS records response.
type ImportResourceRecordsResp struct {
	// Number of records parsed from the zone file.
	TotalRecordsParsed *int64 `json:"total_records_parsed" validate:"required"`

	// Number of records imported successfully.
	RecordsAdded *int64 `json:"records_added" validate:"required"`

	// Number of records failed import.
	RecordsFailed *int64 `json:"records_failed" validate:"required"`

	// Number of records classified by type.
	RecordsAddedByType *RecordStatsByType `json:"records_added_by_type" validate:"required"`

	// Number of records classified by type.
	RecordsFailedByType *RecordStatsByType `json:"records_failed_by_type" validate:"required"`

	// Error messages.
	Messages []RecordsImportMessageModel `json:"messages,omitempty"`

	// Number of records parsed from the zone file.
	Errors []RecordsImportErrorModel `json:"errors,omitempty"`
}

// UnmarshalImportResourceRecordsResp unmarshals an instance of ImportResourceRecordsResp from the specified map of raw messages.
func UnmarshalImportResourceRecordsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportResourceRecordsResp)
	err = core.UnmarshalPrimitive(m, "total_records_parsed", &obj.TotalRecordsParsed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "records_added", &obj.RecordsAdded)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "records_failed", &obj.RecordsFailed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "records_added_by_type", &obj.RecordsAddedByType, UnmarshalRecordStatsByType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "records_failed_by_type", &obj.RecordsFailedByType, UnmarshalRecordStatsByType)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "messages", &obj.Messages, UnmarshalRecordsImportMessageModel)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalRecordsImportErrorModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDnszones : List DNS zones response.
type ListDnszones struct {
	// An array of DNS zones.
	Dnszones []Dnszone `json:"dnszones" validate:"required"`

	// Specify how many DNS zones to skip over, the default value is 0.
	Offset *int64 `json:"offset" validate:"required"`

	// Specify how many DNS zones are returned, the default value is 10.
	Limit *int64 `json:"limit" validate:"required"`

	// Total number of DNS zones.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next,omitempty"`
}

// UnmarshalListDnszones unmarshals an instance of ListDnszones from the specified map of raw messages.
func UnmarshalListDnszones(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListDnszones)
	err = core.UnmarshalModel(m, "dnszones", &obj.Dnszones, UnmarshalDnszone)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListDnszones) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// ListLoadBalancers : List Global Load Balancers response.
type ListLoadBalancers struct {
	// An array of Global Load Balancers.
	LoadBalancers []LoadBalancer `json:"load_balancers" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of Global Load Balancers per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of Global Load Balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of Global Load Balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}

// UnmarshalListLoadBalancers unmarshals an instance of ListLoadBalancers from the specified map of raw messages.
func UnmarshalListLoadBalancers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListLoadBalancers)
	err = core.UnmarshalModel(m, "load_balancers", &obj.LoadBalancers, UnmarshalLoadBalancer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListMonitors : List load balancer monitors response.
type ListMonitors struct {
	// An array of load balancer monitors.
	Monitors []Monitor `json:"monitors" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of load balancer monitors per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of load balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of load balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}

// UnmarshalListMonitors unmarshals an instance of ListMonitors from the specified map of raw messages.
func UnmarshalListMonitors(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListMonitors)
	err = core.UnmarshalModel(m, "monitors", &obj.Monitors, UnmarshalMonitor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListPermittedNetworks : List permitted networks response.
type ListPermittedNetworks struct {
	// An array of permitted networks.
	PermittedNetworks []PermittedNetwork `json:"permitted_networks" validate:"required"`

	// Specify how many permitted networks to skip over, the default value is 0.
	Offset *int64 `json:"offset" validate:"required"`

	// Specify how many permitted networks are returned, the default value is 10.
	Limit *int64 `json:"limit" validate:"required"`

	// Total number of permitted networks.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next,omitempty"`
}

// UnmarshalListPermittedNetworks unmarshals an instance of ListPermittedNetworks from the specified map of raw messages.
func UnmarshalListPermittedNetworks(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPermittedNetworks)
	err = core.UnmarshalModel(m, "permitted_networks", &obj.PermittedNetworks, UnmarshalPermittedNetwork)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListPermittedNetworks) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// ListPools : List load balancer pools response.
type ListPools struct {
	// An array of load balancer pools.
	Pools []Pool `json:"pools" validate:"required"`

	// Page number.
	Offset *int64 `json:"offset" validate:"required"`

	// Number of load balancer pools per page.
	Limit *int64 `json:"limit" validate:"required"`

	// Number of load balancers.
	Count *int64 `json:"count" validate:"required"`

	// Total number of load balancers.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next" validate:"required"`
}

// UnmarshalListPools unmarshals an instance of ListPools from the specified map of raw messages.
func UnmarshalListPools(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPools)
	err = core.UnmarshalModel(m, "pools", &obj.Pools, UnmarshalPool)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListResourceRecords : List Resource Records response.
type ListResourceRecords struct {
	// An array of resource records.
	ResourceRecords []ResourceRecord `json:"resource_records" validate:"required"`

	// Specify how many resource records to skip over, the default value is 0.
	Offset *int64 `json:"offset" validate:"required"`

	// Specify how many resource records are returned, the default value is 20.
	Limit *int64 `json:"limit" validate:"required"`

	// Total number of resource records.
	TotalCount *int64 `json:"total_count" validate:"required"`

	// href.
	First *FirstHref `json:"first" validate:"required"`

	// href.
	Next *NextHref `json:"next,omitempty"`
}

// UnmarshalListResourceRecords unmarshals an instance of ListResourceRecords from the specified map of raw messages.
func UnmarshalListResourceRecords(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListResourceRecords)
	err = core.UnmarshalModel(m, "resource_records", &obj.ResourceRecords, UnmarshalResourceRecord)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "offset", &obj.Offset)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "limit", &obj.Limit)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total_count", &obj.TotalCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "first", &obj.First, UnmarshalFirstHref)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "next", &obj.Next, UnmarshalNextHref)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Retrieve the value to be passed to a request to access the next page of results
func (resp *ListResourceRecords) GetNextOffset() (*int64, error) {
	if core.IsNil(resp.Next) {
		return nil, nil
	}
	offset, err := core.GetQueryParam(resp.Next.Href, "offset")
	if err != nil || offset == nil {
		return nil, err
	}
	var offsetValue int64
	offsetValue, err = strconv.ParseInt(*offset, 10, 64)
	if err != nil {
		return nil, err
	}
	return core.Int64Ptr(offsetValue), nil
}

// LoadBalancer : Load balancer details.
type LoadBalancer struct {
	// Identifier of the load balancer.
	ID *string `json:"id,omitempty"`

	// Name of the load balancer.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Healthy state of the load balancer.
	Health *string `json:"health,omitempty"`

	// The pool ID to use when all other pools are detected as unhealthy.
	FallbackPool *string `json:"fallback_pool,omitempty"`

	// A list of pool IDs ordered by their failover priority. Pools defined here are used by default, or when region_pools
	// are not configured for a given region.
	DefaultPools []string `json:"default_pools,omitempty"`

	// Map availability zones to pool IDs.
	AzPools []LoadBalancerAzPoolsItem `json:"az_pools,omitempty"`

	// The time when a load balancer is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// The recent time when a load balancer is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`
}

// Constants associated with the LoadBalancer.Health property.
// Healthy state of the load balancer.
const (
	LoadBalancer_Health_Critical = "CRITICAL"
	LoadBalancer_Health_Degraded = "DEGRADED"
	LoadBalancer_Health_Healthy  = "HEALTHY"
)

// UnmarshalLoadBalancer unmarshals an instance of LoadBalancer from the specified map of raw messages.
func UnmarshalLoadBalancer(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LoadBalancer)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ttl", &obj.TTL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health", &obj.Health)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "fallback_pool", &obj.FallbackPool)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "default_pools", &obj.DefaultPools)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "az_pools", &obj.AzPools, UnmarshalLoadBalancerAzPoolsItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Location : Custom resolver location.
type Location struct {
	// Location ID.
	ID *string `json:"id,omitempty"`

	// Subnet CRN.
	SubnetCrn *string `json:"subnet_crn,omitempty"`

	// Whether the location is enabled for the custom resolver.
	Enabled *bool `json:"enabled,omitempty"`

	// Whether the DNS server in this location is healthy or not.
	Healthy *bool `json:"healthy,omitempty"`

	// The ip address of this dns server.
	DnsServerIp *string `json:"dns_server_ip,omitempty"`
}

// UnmarshalLocation unmarshals an instance of Location from the specified map of raw messages.
func UnmarshalLocation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Location)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "subnet_crn", &obj.SubnetCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthy", &obj.Healthy)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dns_server_ip", &obj.DnsServerIp)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// LocationInput : Request to add custom resolver location.
type LocationInput struct {
	// Custom resolver location, subnet CRN.
	SubnetCrn *string `json:"subnet_crn" validate:"required"`

	// Enable/Disable custom resolver location.
	Enabled *bool `json:"enabled,omitempty"`
}

// NewLocationInput : Instantiate LocationInput (Generic Model Constructor)
func (*DnsSvcsV1) NewLocationInput(subnetCrn string) (_model *LocationInput, err error) {
	_model = &LocationInput{
		SubnetCrn: core.StringPtr(subnetCrn),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalLocationInput unmarshals an instance of LocationInput from the specified map of raw messages.
func UnmarshalLocationInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(LocationInput)
	err = core.UnmarshalPrimitive(m, "subnet_crn", &obj.SubnetCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Monitor : Load balancer monitor details.
type Monitor struct {
	// Identifier of the load balancer monitor.
	ID *string `json:"id,omitempty"`

	// The name of the load balancer monitor.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer monitor.
	Description *string `json:"description,omitempty"`

	// The protocol to use for the health check. Currently supported protocols are 'HTTP','HTTPS' and 'TCP'.
	Type *string `json:"type,omitempty"`

	// Port number to connect to for the health check. Required for TCP checks. HTTP and HTTPS checks should only define
	// the port when using a non-standard port (HTTP: default 80, HTTPS: default 443).
	Port *int64 `json:"port,omitempty"`

	// The interval between each health check. Shorter intervals may improve failover time, but will increase load on the
	// origins as we check from multiple locations.
	Interval *int64 `json:"interval,omitempty"`

	// The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted
	// immediately.
	Retries *int64 `json:"retries,omitempty"`

	// The timeout (in seconds) before marking the health check as failed.
	Timeout *int64 `json:"timeout,omitempty"`

	// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
	Method *string `json:"method,omitempty"`

	// The endpoint path to health check against. This parameter is only valid for HTTP and HTTPS monitors.
	Path *string `json:"path,omitempty"`

	// The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The
	// User-Agent header cannot be overridden. This parameter is only valid for HTTP and HTTPS monitors.
	HeadersVar []HealthcheckHeader `json:"headers,omitempty"`

	// Do not validate the certificate when monitor use HTTPS. This parameter is currently only valid for HTTPS monitors.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// The expected HTTP response code or code range of the health check. This parameter is only valid for HTTP and HTTPS
	// monitors.
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be
	// marked as unhealthy. This parameter is only valid for HTTP and HTTPS monitors.
	ExpectedBody *string `json:"expected_body,omitempty"`

	// the time when a load balancer monitor is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer monitor is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`
}

// Constants associated with the Monitor.Method property.
// The method to use for the health check applicable to HTTP/HTTPS based checks, the default value is 'GET'.
const (
	Monitor_Method_Get  = "GET"
	Monitor_Method_Head = "HEAD"
)

// UnmarshalMonitor unmarshals an instance of Monitor from the specified map of raw messages.
func UnmarshalMonitor(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Monitor)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "interval", &obj.Interval)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "retries", &obj.Retries)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "timeout", &obj.Timeout)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "method", &obj.Method)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "headers", &obj.HeadersVar, UnmarshalHealthcheckHeader)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allow_insecure", &obj.AllowInsecure)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_codes", &obj.ExpectedCodes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expected_body", &obj.ExpectedBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NextHref : href.
type NextHref struct {
	// href.
	Href *string `json:"href,omitempty"`
}

// UnmarshalNextHref unmarshals an instance of NextHref from the specified map of raw messages.
func UnmarshalNextHref(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NextHref)
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Origin : Origin server.
type Origin struct {
	// The name of the origin server.
	Name *string `json:"name,omitempty"`

	// Description of the origin server.
	Description *string `json:"description,omitempty"`

	// The address of the origin server. It can be a hostname or an IP address.
	Address *string `json:"address,omitempty"`

	// Whether the origin server is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The health state of the origin server.
	Health *bool `json:"health,omitempty"`

	// The failure reason of the origin server if it is unhealthy.
	HealthFailureReason *string `json:"health_failure_reason,omitempty"`
}

// UnmarshalOrigin unmarshals an instance of Origin from the specified map of raw messages.
func UnmarshalOrigin(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Origin)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health", &obj.Health)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health_failure_reason", &obj.HealthFailureReason)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// OriginInput : The request data of origin server.
type OriginInput struct {
	// The name of the origin server.
	Name *string `json:"name,omitempty"`

	// Description of the origin server.
	Description *string `json:"description,omitempty"`

	// The address of the origin server. It can be a hostname or an IP address.
	Address *string `json:"address,omitempty"`

	// Whether the origin server is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// UnmarshalOriginInput unmarshals an instance of OriginInput from the specified map of raw messages.
func UnmarshalOriginInput(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(OriginInput)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "address", &obj.Address)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PermittedNetwork : Permitted network details.
type PermittedNetwork struct {
	// Unique identifier of a permitted network.
	ID *string `json:"id,omitempty"`

	// The time when a permitted network is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// The recent time when a permitted network is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Permitted network data for VPC.
	PermittedNetwork *PermittedNetworkVpc `json:"permitted_network,omitempty"`

	// The type of a permitted network.
	Type *string `json:"type,omitempty"`

	// The state of a permitted network.
	State *string `json:"state,omitempty"`
}

// Constants associated with the PermittedNetwork.Type property.
// The type of a permitted network.
const (
	PermittedNetwork_Type_Vpc = "vpc"
)

// Constants associated with the PermittedNetwork.State property.
// The state of a permitted network.
const (
	PermittedNetwork_State_Active            = "ACTIVE"
	PermittedNetwork_State_RemovalInProgress = "REMOVAL_IN_PROGRESS"
)

// UnmarshalPermittedNetwork unmarshals an instance of PermittedNetwork from the specified map of raw messages.
func UnmarshalPermittedNetwork(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PermittedNetwork)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "permitted_network", &obj.PermittedNetwork, UnmarshalPermittedNetworkVpc)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "state", &obj.State)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PermittedNetworkVpc : Permitted network data for VPC.
type PermittedNetworkVpc struct {
	// CRN string uniquely identifies a VPC.
	VpcCrn *string `json:"vpc_crn" validate:"required"`
}

// NewPermittedNetworkVpc : Instantiate PermittedNetworkVpc (Generic Model Constructor)
func (*DnsSvcsV1) NewPermittedNetworkVpc(vpcCrn string) (_model *PermittedNetworkVpc, err error) {
	_model = &PermittedNetworkVpc{
		VpcCrn: core.StringPtr(vpcCrn),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPermittedNetworkVpc unmarshals an instance of PermittedNetworkVpc from the specified map of raw messages.
func UnmarshalPermittedNetworkVpc(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PermittedNetworkVpc)
	err = core.UnmarshalPrimitive(m, "vpc_crn", &obj.VpcCrn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Pool : Load balancer pool details.
type Pool struct {
	// Identifier of the load balancer pool.
	ID *string `json:"id,omitempty"`

	// Name of the load balancer pool.
	Name *string `json:"name,omitempty"`

	// Descriptive text of the load balancer pool.
	Description *string `json:"description,omitempty"`

	// Whether the load balancer pool is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins
	// falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
	HealthyOriginsThreshold *int64 `json:"healthy_origins_threshold,omitempty"`

	// The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy
	// origins, provided the pool itself is healthy.
	Origins []Origin `json:"origins,omitempty"`

	// The ID of the load balancer monitor to be associated to this pool.
	Monitor *string `json:"monitor,omitempty"`

	// The notification channel.
	NotificationChannel *string `json:"notification_channel,omitempty"`

	// Healthy state of the load balancer pool.
	Health *string `json:"health,omitempty"`

	// Health check region of VSIs.
	HealthcheckRegion *string `json:"healthcheck_region,omitempty"`

	// Health check subnet CRNs.
	HealthcheckSubnets []string `json:"healthcheck_subnets,omitempty"`

	// Health check VSI information.
	HealthcheckVsis []PoolHealthcheckVsisItem `json:"healthcheck_vsis,omitempty"`

	// the time when a load balancer pool is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a load balancer pool is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`
}

// Constants associated with the Pool.Health property.
// Healthy state of the load balancer pool.
const (
	Pool_Health_Critical = "CRITICAL"
	Pool_Health_Degraded = "DEGRADED"
	Pool_Health_Healthy  = "HEALTHY"
)

// Constants associated with the Pool.HealthcheckRegion property.
// Health check region of VSIs.
const (
	Pool_HealthcheckRegion_AuSyd   = "au-syd"
	Pool_HealthcheckRegion_EuDu    = "eu-du"
	Pool_HealthcheckRegion_EuGb    = "eu-gb"
	Pool_HealthcheckRegion_JpTok   = "jp-tok"
	Pool_HealthcheckRegion_UsEast  = "us-east"
	Pool_HealthcheckRegion_UsSouth = "us-south"
)

// UnmarshalPool unmarshals an instance of Pool from the specified map of raw messages.
func UnmarshalPool(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Pool)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthy_origins_threshold", &obj.HealthyOriginsThreshold)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "origins", &obj.Origins, UnmarshalOrigin)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "monitor", &obj.Monitor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "notification_channel", &obj.NotificationChannel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "health", &obj.Health)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthcheck_region", &obj.HealthcheckRegion)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "healthcheck_subnets", &obj.HealthcheckSubnets)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "healthcheck_vsis", &obj.HealthcheckVsis, UnmarshalPoolHealthcheckVsisItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordStatsByType : Number of records classified by type.
type RecordStatsByType struct {
	// Number of records, type A.
	A *int64 `json:"A" validate:"required"`

	// Number of records, type AAAA.
	AAAA *int64 `json:"AAAA" validate:"required"`

	// Number of records, type CNAME.
	CNAME *int64 `json:"CNAME" validate:"required"`

	// Number of records, type SRV.
	SRV *int64 `json:"SRV" validate:"required"`

	// Number of records, type TXT.
	TXT *int64 `json:"TXT" validate:"required"`

	// Number of records, type MX.
	MX *int64 `json:"MX" validate:"required"`

	// Number of records, type PTR.
	PTR *int64 `json:"PTR" validate:"required"`
}

// UnmarshalRecordStatsByType unmarshals an instance of RecordStatsByType from the specified map of raw messages.
func UnmarshalRecordStatsByType(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordStatsByType)
	err = core.UnmarshalPrimitive(m, "A", &obj.A)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "AAAA", &obj.AAAA)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "CNAME", &obj.CNAME)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "SRV", &obj.SRV)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "TXT", &obj.TXT)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "MX", &obj.MX)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "PTR", &obj.PTR)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsImportErrorModel : RecordsImportErrorModel struct
type RecordsImportErrorModel struct {
	// resource record content in zone file.
	ResourceRecord *string `json:"resource_record" validate:"required"`

	Error *RecordsImportErrorModelError `json:"error" validate:"required"`
}

// UnmarshalRecordsImportErrorModel unmarshals an instance of RecordsImportErrorModel from the specified map of raw messages.
func UnmarshalRecordsImportErrorModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordsImportErrorModel)
	err = core.UnmarshalPrimitive(m, "resource_record", &obj.ResourceRecord)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "error", &obj.Error, UnmarshalRecordsImportErrorModelError)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RecordsImportMessageModel : RecordsImportMessageModel struct
type RecordsImportMessageModel struct {
	// Code to classify import DNS records error.
	Code *string `json:"code" validate:"required"`

	// Message to describe import DNS records error.
	Message *string `json:"message" validate:"required"`
}

// UnmarshalRecordsImportMessageModel unmarshals an instance of RecordsImportMessageModel from the specified map of raw messages.
func UnmarshalRecordsImportMessageModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RecordsImportMessageModel)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecord : Resource record details.
type ResourceRecord struct {
	// Identifier of the resource record.
	ID *string `json:"id,omitempty"`

	// the time when a resource record is created.
	CreatedOn *string `json:"created_on,omitempty"`

	// the recent time when a resource record is modified.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// Name of the resource record.
	Name *string `json:"name,omitempty"`

	// Type of the resource record.
	Type *string `json:"type,omitempty"`

	// Time to live in second.
	TTL *int64 `json:"ttl,omitempty"`

	// Content of the resource record.
	Rdata interface{} `json:"rdata,omitempty"`

	// Only used for SRV record.
	Service *string `json:"service,omitempty"`

	// Only used for SRV record.
	Protocol *string `json:"protocol,omitempty"`
}

// Constants associated with the ResourceRecord.Type property.
// Type of the resource record.
const (
	ResourceRecord_Type_A     = "A"
	ResourceRecord_Type_Aaaa  = "AAAA"
	ResourceRecord_Type_Cname = "CNAME"
	ResourceRecord_Type_Mx    = "MX"
	ResourceRecord_Type_Ptr   = "PTR"
	ResourceRecord_Type_Srv   = "SRV"
	ResourceRecord_Type_Txt   = "TXT"
)

// UnmarshalResourceRecord unmarshals an instance of ResourceRecord from the specified map of raw messages.
func UnmarshalResourceRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecord)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_on", &obj.CreatedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "modified_on", &obj.ModifiedOn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ttl", &obj.TTL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "rdata", &obj.Rdata)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "service", &obj.Service)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "protocol", &obj.Protocol)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdataRdataARecord : The content of type-A resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataARecord struct {
	// IPv4 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordInputRdataRdataARecord : Instantiate ResourceRecordInputRdataRdataARecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataARecord(ip string) (_model *ResourceRecordInputRdataRdataARecord, err error) {
	_model = &ResourceRecordInputRdataRdataARecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataARecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataARecord unmarshals an instance of ResourceRecordInputRdataRdataARecord from the specified map of raw messages.
func UnmarshalResourceRecordInputRdataRdataARecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdataRdataARecord)
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdataRdataAaaaRecord : The content of type-AAAA resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataAaaaRecord struct {
	// IPv6 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordInputRdataRdataAaaaRecord : Instantiate ResourceRecordInputRdataRdataAaaaRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataAaaaRecord(ip string) (_model *ResourceRecordInputRdataRdataAaaaRecord, err error) {
	_model = &ResourceRecordInputRdataRdataAaaaRecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataAaaaRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataAaaaRecord unmarshals an instance of ResourceRecordInputRdataRdataAaaaRecord from the specified map of raw messages.
func UnmarshalResourceRecordInputRdataRdataAaaaRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdataRdataAaaaRecord)
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdataRdataCnameRecord : The content of type-CNAME resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataCnameRecord struct {
	// Canonical name.
	Cname *string `json:"cname" validate:"required"`
}

// NewResourceRecordInputRdataRdataCnameRecord : Instantiate ResourceRecordInputRdataRdataCnameRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataCnameRecord(cname string) (_model *ResourceRecordInputRdataRdataCnameRecord, err error) {
	_model = &ResourceRecordInputRdataRdataCnameRecord{
		Cname: core.StringPtr(cname),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataCnameRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataCnameRecord unmarshals an instance of ResourceRecordInputRdataRdataCnameRecord from the specified map of raw messages.
func UnmarshalResourceRecordInputRdataRdataCnameRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdataRdataCnameRecord)
	err = core.UnmarshalPrimitive(m, "cname", &obj.Cname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdataRdataMxRecord : The content of type-MX resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataMxRecord struct {
	// Hostname of Exchange server.
	Exchange *string `json:"exchange" validate:"required"`

	// Preference of the MX record.
	Preference *int64 `json:"preference" validate:"required"`
}

// NewResourceRecordInputRdataRdataMxRecord : Instantiate ResourceRecordInputRdataRdataMxRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataMxRecord(exchange string, preference int64) (_model *ResourceRecordInputRdataRdataMxRecord, err error) {
	_model = &ResourceRecordInputRdataRdataMxRecord{
		Exchange:   core.StringPtr(exchange),
		Preference: core.Int64Ptr(preference),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataMxRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataMxRecord unmarshals an instance of ResourceRecordInputRdataRdataMxRecord from the specified map of raw messages.
func UnmarshalResourceRecordInputRdataRdataMxRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdataRdataMxRecord)
	err = core.UnmarshalPrimitive(m, "exchange", &obj.Exchange)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preference", &obj.Preference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdataRdataPtrRecord : The content of type-PTR resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataPtrRecord struct {
	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname" validate:"required"`
}

// NewResourceRecordInputRdataRdataPtrRecord : Instantiate ResourceRecordInputRdataRdataPtrRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataPtrRecord(ptrdname string) (_model *ResourceRecordInputRdataRdataPtrRecord, err error) {
	_model = &ResourceRecordInputRdataRdataPtrRecord{
		Ptrdname: core.StringPtr(ptrdname),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataPtrRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataPtrRecord unmarshals an instance of ResourceRecordInputRdataRdataPtrRecord from the specified map of raw messages.
func UnmarshalResourceRecordInputRdataRdataPtrRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdataRdataPtrRecord)
	err = core.UnmarshalPrimitive(m, "ptrdname", &obj.Ptrdname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdataRdataSrvRecord : The content of type-SRV resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataSrvRecord struct {
	// Port number of the target server.
	Port *int64 `json:"port" validate:"required"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority" validate:"required"`

	// Hostname of the target server.
	Target *string `json:"target" validate:"required"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight" validate:"required"`
}

// NewResourceRecordInputRdataRdataSrvRecord : Instantiate ResourceRecordInputRdataRdataSrvRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataSrvRecord(port int64, priority int64, target string, weight int64) (_model *ResourceRecordInputRdataRdataSrvRecord, err error) {
	_model = &ResourceRecordInputRdataRdataSrvRecord{
		Port:     core.Int64Ptr(port),
		Priority: core.Int64Ptr(priority),
		Target:   core.StringPtr(target),
		Weight:   core.Int64Ptr(weight),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataSrvRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataSrvRecord unmarshals an instance of ResourceRecordInputRdataRdataSrvRecord from the specified map of raw messages.
func UnmarshalResourceRecordInputRdataRdataSrvRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdataRdataSrvRecord)
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "priority", &obj.Priority)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "weight", &obj.Weight)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordInputRdataRdataTxtRecord : The content of type-TXT resource record.
// This model "extends" ResourceRecordInputRdata
type ResourceRecordInputRdataRdataTxtRecord struct {
	// Human readable text.
	Text *string `json:"text" validate:"required"`
}

// NewResourceRecordInputRdataRdataTxtRecord : Instantiate ResourceRecordInputRdataRdataTxtRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordInputRdataRdataTxtRecord(text string) (_model *ResourceRecordInputRdataRdataTxtRecord, err error) {
	_model = &ResourceRecordInputRdataRdataTxtRecord{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordInputRdataRdataTxtRecord) isaResourceRecordInputRdata() bool {
	return true
}

// UnmarshalResourceRecordInputRdataRdataTxtRecord unmarshals an instance of ResourceRecordInputRdataRdataTxtRecord from the specified map of raw messages.
func UnmarshalResourceRecordInputRdataRdataTxtRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordInputRdataRdataTxtRecord)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdataRdataARecord : The content of type-A resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataARecord struct {
	// IPv4 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataARecord : Instantiate ResourceRecordUpdateInputRdataRdataARecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataARecord(ip string) (_model *ResourceRecordUpdateInputRdataRdataARecord, err error) {
	_model = &ResourceRecordUpdateInputRdataRdataARecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataARecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataARecord unmarshals an instance of ResourceRecordUpdateInputRdataRdataARecord from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdataRdataARecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataARecord)
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdataRdataAaaaRecord : The content of type-AAAA resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataAaaaRecord struct {
	// IPv6 address.
	Ip *string `json:"ip" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataAaaaRecord : Instantiate ResourceRecordUpdateInputRdataRdataAaaaRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataAaaaRecord(ip string) (_model *ResourceRecordUpdateInputRdataRdataAaaaRecord, err error) {
	_model = &ResourceRecordUpdateInputRdataRdataAaaaRecord{
		Ip: core.StringPtr(ip),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataAaaaRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataAaaaRecord unmarshals an instance of ResourceRecordUpdateInputRdataRdataAaaaRecord from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdataRdataAaaaRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataAaaaRecord)
	err = core.UnmarshalPrimitive(m, "ip", &obj.Ip)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdataRdataCnameRecord : The content of type-CNAME resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataCnameRecord struct {
	// Canonical name.
	Cname *string `json:"cname" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataCnameRecord : Instantiate ResourceRecordUpdateInputRdataRdataCnameRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataCnameRecord(cname string) (_model *ResourceRecordUpdateInputRdataRdataCnameRecord, err error) {
	_model = &ResourceRecordUpdateInputRdataRdataCnameRecord{
		Cname: core.StringPtr(cname),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataCnameRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataCnameRecord unmarshals an instance of ResourceRecordUpdateInputRdataRdataCnameRecord from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdataRdataCnameRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataCnameRecord)
	err = core.UnmarshalPrimitive(m, "cname", &obj.Cname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdataRdataMxRecord : The content of type-MX resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataMxRecord struct {
	// Hostname of Exchange server.
	Exchange *string `json:"exchange" validate:"required"`

	// Preference of the MX record.
	Preference *int64 `json:"preference" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataMxRecord : Instantiate ResourceRecordUpdateInputRdataRdataMxRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataMxRecord(exchange string, preference int64) (_model *ResourceRecordUpdateInputRdataRdataMxRecord, err error) {
	_model = &ResourceRecordUpdateInputRdataRdataMxRecord{
		Exchange:   core.StringPtr(exchange),
		Preference: core.Int64Ptr(preference),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataMxRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataMxRecord unmarshals an instance of ResourceRecordUpdateInputRdataRdataMxRecord from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdataRdataMxRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataMxRecord)
	err = core.UnmarshalPrimitive(m, "exchange", &obj.Exchange)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preference", &obj.Preference)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdataRdataPtrRecord : The content of type-PTR resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataPtrRecord struct {
	// Hostname of the relevant A or AAAA record.
	Ptrdname *string `json:"ptrdname" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataPtrRecord : Instantiate ResourceRecordUpdateInputRdataRdataPtrRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataPtrRecord(ptrdname string) (_model *ResourceRecordUpdateInputRdataRdataPtrRecord, err error) {
	_model = &ResourceRecordUpdateInputRdataRdataPtrRecord{
		Ptrdname: core.StringPtr(ptrdname),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataPtrRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataPtrRecord unmarshals an instance of ResourceRecordUpdateInputRdataRdataPtrRecord from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdataRdataPtrRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataPtrRecord)
	err = core.UnmarshalPrimitive(m, "ptrdname", &obj.Ptrdname)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdataRdataSrvRecord : The content of type-SRV resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataSrvRecord struct {
	// Port number of the target server.
	Port *int64 `json:"port" validate:"required"`

	// Priority of the SRV record.
	Priority *int64 `json:"priority" validate:"required"`

	// Hostname of the target server.
	Target *string `json:"target" validate:"required"`

	// Weight of distributing queries among multiple target servers.
	Weight *int64 `json:"weight" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataSrvRecord : Instantiate ResourceRecordUpdateInputRdataRdataSrvRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataSrvRecord(port int64, priority int64, target string, weight int64) (_model *ResourceRecordUpdateInputRdataRdataSrvRecord, err error) {
	_model = &ResourceRecordUpdateInputRdataRdataSrvRecord{
		Port:     core.Int64Ptr(port),
		Priority: core.Int64Ptr(priority),
		Target:   core.StringPtr(target),
		Weight:   core.Int64Ptr(weight),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataSrvRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataSrvRecord unmarshals an instance of ResourceRecordUpdateInputRdataRdataSrvRecord from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdataRdataSrvRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataSrvRecord)
	err = core.UnmarshalPrimitive(m, "port", &obj.Port)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "priority", &obj.Priority)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "target", &obj.Target)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "weight", &obj.Weight)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResourceRecordUpdateInputRdataRdataTxtRecord : The content of type-TXT resource record.
// This model "extends" ResourceRecordUpdateInputRdata
type ResourceRecordUpdateInputRdataRdataTxtRecord struct {
	// Human readable text.
	Text *string `json:"text" validate:"required"`
}

// NewResourceRecordUpdateInputRdataRdataTxtRecord : Instantiate ResourceRecordUpdateInputRdataRdataTxtRecord (Generic Model Constructor)
func (*DnsSvcsV1) NewResourceRecordUpdateInputRdataRdataTxtRecord(text string) (_model *ResourceRecordUpdateInputRdataRdataTxtRecord, err error) {
	_model = &ResourceRecordUpdateInputRdataRdataTxtRecord{
		Text: core.StringPtr(text),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

func (*ResourceRecordUpdateInputRdataRdataTxtRecord) isaResourceRecordUpdateInputRdata() bool {
	return true
}

// UnmarshalResourceRecordUpdateInputRdataRdataTxtRecord unmarshals an instance of ResourceRecordUpdateInputRdataRdataTxtRecord from the specified map of raw messages.
func UnmarshalResourceRecordUpdateInputRdataRdataTxtRecord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResourceRecordUpdateInputRdataRdataTxtRecord)
	err = core.UnmarshalPrimitive(m, "text", &obj.Text)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
