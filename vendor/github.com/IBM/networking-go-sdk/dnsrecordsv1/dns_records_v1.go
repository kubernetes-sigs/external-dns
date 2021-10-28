/**
 * (C) Copyright IBM Corp. 2020.
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
 * IBM OpenAPI SDK Code Generator Version: 3.20.0-debb9f29-20201203-202043
 */

// Package dnsrecordsv1 : Operations and models for the DnsRecordsV1 service
package dnsrecordsv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/networking-go-sdk/common"
)

// DnsRecordsV1 : DNS records
//
// Version: 1.0.1
type DnsRecordsV1 struct {
	Service *core.BaseService

	// Full crn of the service instance.
	Crn *string

	// Zone identifier (zone id).
	ZoneIdentifier *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "dns_records"

// DnsRecordsV1Options : Service options
type DnsRecordsV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full crn of the service instance.
	Crn *string `validate:"required"`

	// Zone identifier (zone id).
	ZoneIdentifier *string `validate:"required"`
}

// NewDnsRecordsV1UsingExternalConfig : constructs an instance of DnsRecordsV1 with passed in options and external configuration.
func NewDnsRecordsV1UsingExternalConfig(options *DnsRecordsV1Options) (dnsRecords *DnsRecordsV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	dnsRecords, err = NewDnsRecordsV1(options)
	if err != nil {
		return
	}

	err = dnsRecords.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = dnsRecords.Service.SetServiceURL(options.URL)
	}
	return
}

// NewDnsRecordsV1 : constructs an instance of DnsRecordsV1 with passed in options.
func NewDnsRecordsV1(options *DnsRecordsV1Options) (service *DnsRecordsV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
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

	service = &DnsRecordsV1{
		Service:        baseService,
		Crn:            options.Crn,
		ZoneIdentifier: options.ZoneIdentifier,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "dnsRecords" suitable for processing requests.
func (dnsRecords *DnsRecordsV1) Clone() *DnsRecordsV1 {
	if core.IsNil(dnsRecords) {
		return nil
	}
	clone := *dnsRecords
	clone.Service = dnsRecords.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (dnsRecords *DnsRecordsV1) SetServiceURL(url string) error {
	return dnsRecords.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (dnsRecords *DnsRecordsV1) GetServiceURL() string {
	return dnsRecords.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (dnsRecords *DnsRecordsV1) SetDefaultHeaders(headers http.Header) {
	dnsRecords.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (dnsRecords *DnsRecordsV1) SetEnableGzipCompression(enableGzip bool) {
	dnsRecords.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (dnsRecords *DnsRecordsV1) GetEnableGzipCompression() bool {
	return dnsRecords.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (dnsRecords *DnsRecordsV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	dnsRecords.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (dnsRecords *DnsRecordsV1) DisableRetries() {
	dnsRecords.Service.DisableRetries()
}

// ListAllDnsRecords : List all DNS records
// List all DNS records for a given zone of a service instance.
func (dnsRecords *DnsRecordsV1) ListAllDnsRecords(listAllDnsRecordsOptions *ListAllDnsRecordsOptions) (result *ListDnsrecordsResp, response *core.DetailedResponse, err error) {
	return dnsRecords.ListAllDnsRecordsWithContext(context.Background(), listAllDnsRecordsOptions)
}

// ListAllDnsRecordsWithContext is an alternate form of the ListAllDnsRecords method which supports a Context parameter
func (dnsRecords *DnsRecordsV1) ListAllDnsRecordsWithContext(ctx context.Context, listAllDnsRecordsOptions *ListAllDnsRecordsOptions) (result *ListDnsrecordsResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listAllDnsRecordsOptions, "listAllDnsRecordsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *dnsRecords.Crn,
		"zone_identifier": *dnsRecords.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsRecords.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsRecords.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/dns_records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listAllDnsRecordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_records", "V1", "ListAllDnsRecords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listAllDnsRecordsOptions.Type != nil {
		builder.AddQuery("type", fmt.Sprint(*listAllDnsRecordsOptions.Type))
	}
	if listAllDnsRecordsOptions.Name != nil {
		builder.AddQuery("name", fmt.Sprint(*listAllDnsRecordsOptions.Name))
	}
	if listAllDnsRecordsOptions.Content != nil {
		builder.AddQuery("content", fmt.Sprint(*listAllDnsRecordsOptions.Content))
	}
	if listAllDnsRecordsOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listAllDnsRecordsOptions.Page))
	}
	if listAllDnsRecordsOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listAllDnsRecordsOptions.PerPage))
	}
	if listAllDnsRecordsOptions.Order != nil {
		builder.AddQuery("order", fmt.Sprint(*listAllDnsRecordsOptions.Order))
	}
	if listAllDnsRecordsOptions.Direction != nil {
		builder.AddQuery("direction", fmt.Sprint(*listAllDnsRecordsOptions.Direction))
	}
	if listAllDnsRecordsOptions.Match != nil {
		builder.AddQuery("match", fmt.Sprint(*listAllDnsRecordsOptions.Match))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsRecords.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListDnsrecordsResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateDnsRecord : Create DNS record
// Add a new DNS record for a given zone for a given service instance.
func (dnsRecords *DnsRecordsV1) CreateDnsRecord(createDnsRecordOptions *CreateDnsRecordOptions) (result *DnsrecordResp, response *core.DetailedResponse, err error) {
	return dnsRecords.CreateDnsRecordWithContext(context.Background(), createDnsRecordOptions)
}

// CreateDnsRecordWithContext is an alternate form of the CreateDnsRecord method which supports a Context parameter
func (dnsRecords *DnsRecordsV1) CreateDnsRecordWithContext(ctx context.Context, createDnsRecordOptions *CreateDnsRecordOptions) (result *DnsrecordResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createDnsRecordOptions, "createDnsRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *dnsRecords.Crn,
		"zone_identifier": *dnsRecords.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsRecords.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsRecords.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/dns_records`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDnsRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_records", "V1", "CreateDnsRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createDnsRecordOptions.Name != nil {
		body["name"] = createDnsRecordOptions.Name
	}
	if createDnsRecordOptions.Type != nil {
		body["type"] = createDnsRecordOptions.Type
	}
	if createDnsRecordOptions.TTL != nil {
		body["ttl"] = createDnsRecordOptions.TTL
	}
	if createDnsRecordOptions.Content != nil {
		body["content"] = createDnsRecordOptions.Content
	}
	if createDnsRecordOptions.Priority != nil {
		body["priority"] = createDnsRecordOptions.Priority
	}
	if createDnsRecordOptions.Data != nil {
		body["data"] = createDnsRecordOptions.Data
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
	response, err = dnsRecords.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDnsrecordResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteDnsRecord : Delete DNS record
// Delete a DNS record given its id.
func (dnsRecords *DnsRecordsV1) DeleteDnsRecord(deleteDnsRecordOptions *DeleteDnsRecordOptions) (result *DeleteDnsrecordResp, response *core.DetailedResponse, err error) {
	return dnsRecords.DeleteDnsRecordWithContext(context.Background(), deleteDnsRecordOptions)
}

// DeleteDnsRecordWithContext is an alternate form of the DeleteDnsRecord method which supports a Context parameter
func (dnsRecords *DnsRecordsV1) DeleteDnsRecordWithContext(ctx context.Context, deleteDnsRecordOptions *DeleteDnsRecordOptions) (result *DeleteDnsrecordResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDnsRecordOptions, "deleteDnsRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDnsRecordOptions, "deleteDnsRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":                  *dnsRecords.Crn,
		"zone_identifier":      *dnsRecords.ZoneIdentifier,
		"dnsrecord_identifier": *deleteDnsRecordOptions.DnsrecordIdentifier,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsRecords.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsRecords.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/dns_records/{dnsrecord_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDnsRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_records", "V1", "DeleteDnsRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsRecords.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteDnsrecordResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetDnsRecord : Get DNS record
// Get the details of a DNS record for a given zone under a given service instance.
func (dnsRecords *DnsRecordsV1) GetDnsRecord(getDnsRecordOptions *GetDnsRecordOptions) (result *DnsrecordResp, response *core.DetailedResponse, err error) {
	return dnsRecords.GetDnsRecordWithContext(context.Background(), getDnsRecordOptions)
}

// GetDnsRecordWithContext is an alternate form of the GetDnsRecord method which supports a Context parameter
func (dnsRecords *DnsRecordsV1) GetDnsRecordWithContext(ctx context.Context, getDnsRecordOptions *GetDnsRecordOptions) (result *DnsrecordResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDnsRecordOptions, "getDnsRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDnsRecordOptions, "getDnsRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":                  *dnsRecords.Crn,
		"zone_identifier":      *dnsRecords.ZoneIdentifier,
		"dnsrecord_identifier": *getDnsRecordOptions.DnsrecordIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsRecords.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsRecords.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/dns_records/{dnsrecord_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDnsRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_records", "V1", "GetDnsRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = dnsRecords.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDnsrecordResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateDnsRecord : Update DNS record
// Update an existing DNS record for a given zone under a given service instance.
func (dnsRecords *DnsRecordsV1) UpdateDnsRecord(updateDnsRecordOptions *UpdateDnsRecordOptions) (result *DnsrecordResp, response *core.DetailedResponse, err error) {
	return dnsRecords.UpdateDnsRecordWithContext(context.Background(), updateDnsRecordOptions)
}

// UpdateDnsRecordWithContext is an alternate form of the UpdateDnsRecord method which supports a Context parameter
func (dnsRecords *DnsRecordsV1) UpdateDnsRecordWithContext(ctx context.Context, updateDnsRecordOptions *UpdateDnsRecordOptions) (result *DnsrecordResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDnsRecordOptions, "updateDnsRecordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDnsRecordOptions, "updateDnsRecordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":                  *dnsRecords.Crn,
		"zone_identifier":      *dnsRecords.ZoneIdentifier,
		"dnsrecord_identifier": *updateDnsRecordOptions.DnsrecordIdentifier,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = dnsRecords.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(dnsRecords.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/dns_records/{dnsrecord_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDnsRecordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("dns_records", "V1", "UpdateDnsRecord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateDnsRecordOptions.Name != nil {
		body["name"] = updateDnsRecordOptions.Name
	}
	if updateDnsRecordOptions.Type != nil {
		body["type"] = updateDnsRecordOptions.Type
	}
	if updateDnsRecordOptions.TTL != nil {
		body["ttl"] = updateDnsRecordOptions.TTL
	}
	if updateDnsRecordOptions.Content != nil {
		body["content"] = updateDnsRecordOptions.Content
	}
	if updateDnsRecordOptions.Priority != nil {
		body["priority"] = updateDnsRecordOptions.Priority
	}
	if updateDnsRecordOptions.Proxied != nil {
		body["proxied"] = updateDnsRecordOptions.Proxied
	}
	if updateDnsRecordOptions.Data != nil {
		body["data"] = updateDnsRecordOptions.Data
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
	response, err = dnsRecords.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDnsrecordResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateDnsRecordOptions : The CreateDnsRecord options.
type CreateDnsRecordOptions struct {
	// Required for all record types except SRV.
	Name *string `json:"name,omitempty"`

	// dns record type.
	Type *string `json:"type,omitempty"`

	// dns record ttl value.
	TTL *int64 `json:"ttl,omitempty"`

	// dns record content.
	Content *string `json:"content,omitempty"`

	// For MX records only.
	Priority *int64 `json:"priority,omitempty"`

	// For LOC, SRV and CAA records only.
	Data interface{} `json:"data,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateDnsRecordOptions.Type property.
// dns record type.
const (
	CreateDnsRecordOptions_Type_A     = "A"
	CreateDnsRecordOptions_Type_Aaaa  = "AAAA"
	CreateDnsRecordOptions_Type_Caa   = "CAA"
	CreateDnsRecordOptions_Type_Cname = "CNAME"
	CreateDnsRecordOptions_Type_Loc   = "LOC"
	CreateDnsRecordOptions_Type_Mx    = "MX"
	CreateDnsRecordOptions_Type_Ns    = "NS"
	CreateDnsRecordOptions_Type_Spf   = "SPF"
	CreateDnsRecordOptions_Type_Srv   = "SRV"
	CreateDnsRecordOptions_Type_Txt   = "TXT"
)

// NewCreateDnsRecordOptions : Instantiate CreateDnsRecordOptions
func (*DnsRecordsV1) NewCreateDnsRecordOptions() *CreateDnsRecordOptions {
	return &CreateDnsRecordOptions{}
}

// SetName : Allow user to set Name
func (options *CreateDnsRecordOptions) SetName(name string) *CreateDnsRecordOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetType : Allow user to set Type
func (options *CreateDnsRecordOptions) SetType(typeVar string) *CreateDnsRecordOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetTTL : Allow user to set TTL
func (options *CreateDnsRecordOptions) SetTTL(ttl int64) *CreateDnsRecordOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetContent : Allow user to set Content
func (options *CreateDnsRecordOptions) SetContent(content string) *CreateDnsRecordOptions {
	options.Content = core.StringPtr(content)
	return options
}

// SetPriority : Allow user to set Priority
func (options *CreateDnsRecordOptions) SetPriority(priority int64) *CreateDnsRecordOptions {
	options.Priority = core.Int64Ptr(priority)
	return options
}

// SetData : Allow user to set Data
func (options *CreateDnsRecordOptions) SetData(data interface{}) *CreateDnsRecordOptions {
	options.Data = data
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDnsRecordOptions) SetHeaders(param map[string]string) *CreateDnsRecordOptions {
	options.Headers = param
	return options
}

// DeleteDnsRecordOptions : The DeleteDnsRecord options.
type DeleteDnsRecordOptions struct {
	// Identifier of DNS record.
	DnsrecordIdentifier *string `json:"dnsrecord_identifier" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDnsRecordOptions : Instantiate DeleteDnsRecordOptions
func (*DnsRecordsV1) NewDeleteDnsRecordOptions(dnsrecordIdentifier string) *DeleteDnsRecordOptions {
	return &DeleteDnsRecordOptions{
		DnsrecordIdentifier: core.StringPtr(dnsrecordIdentifier),
	}
}

// SetDnsrecordIdentifier : Allow user to set DnsrecordIdentifier
func (options *DeleteDnsRecordOptions) SetDnsrecordIdentifier(dnsrecordIdentifier string) *DeleteDnsRecordOptions {
	options.DnsrecordIdentifier = core.StringPtr(dnsrecordIdentifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDnsRecordOptions) SetHeaders(param map[string]string) *DeleteDnsRecordOptions {
	options.Headers = param
	return options
}

// DeleteDnsrecordRespResult : result.
type DeleteDnsrecordRespResult struct {
	// dns record id.
	ID *string `json:"id" validate:"required"`
}

// UnmarshalDeleteDnsrecordRespResult unmarshals an instance of DeleteDnsrecordRespResult from the specified map of raw messages.
func UnmarshalDeleteDnsrecordRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteDnsrecordRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDnsRecordOptions : The GetDnsRecord options.
type GetDnsRecordOptions struct {
	// Identifier of DNS record.
	DnsrecordIdentifier *string `json:"dnsrecord_identifier" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDnsRecordOptions : Instantiate GetDnsRecordOptions
func (*DnsRecordsV1) NewGetDnsRecordOptions(dnsrecordIdentifier string) *GetDnsRecordOptions {
	return &GetDnsRecordOptions{
		DnsrecordIdentifier: core.StringPtr(dnsrecordIdentifier),
	}
}

// SetDnsrecordIdentifier : Allow user to set DnsrecordIdentifier
func (options *GetDnsRecordOptions) SetDnsrecordIdentifier(dnsrecordIdentifier string) *GetDnsRecordOptions {
	options.DnsrecordIdentifier = core.StringPtr(dnsrecordIdentifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetDnsRecordOptions) SetHeaders(param map[string]string) *GetDnsRecordOptions {
	options.Headers = param
	return options
}

// ListAllDnsRecordsOptions : The ListAllDnsRecords options.
type ListAllDnsRecordsOptions struct {
	// Type of DNS records to display.
	Type *string `json:"type,omitempty"`

	// Value of name field to filter by.
	Name *string `json:"name,omitempty"`

	// Value of content field to filter by.
	Content *string `json:"content,omitempty"`

	// Page number of paginated results.
	Page *int64 `json:"page,omitempty"`

	// Maximum number of DNS records per page.
	PerPage *int64 `json:"per_page,omitempty"`

	// Field by which to order list of DNS records.
	Order *string `json:"order,omitempty"`

	// Direction in which to order results [ascending/descending order].
	Direction *string `json:"direction,omitempty"`

	// Whether to match all (all) or atleast one search parameter (any).
	Match *string `json:"match,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListAllDnsRecordsOptions.Order property.
// Field by which to order list of DNS records.
const (
	ListAllDnsRecordsOptions_Order_Content = "content"
	ListAllDnsRecordsOptions_Order_Name    = "name"
	ListAllDnsRecordsOptions_Order_Proxied = "proxied"
	ListAllDnsRecordsOptions_Order_TTL     = "ttl"
	ListAllDnsRecordsOptions_Order_Type    = "type"
)

// Constants associated with the ListAllDnsRecordsOptions.Direction property.
// Direction in which to order results [ascending/descending order].
const (
	ListAllDnsRecordsOptions_Direction_Asc  = "asc"
	ListAllDnsRecordsOptions_Direction_Desc = "desc"
)

// Constants associated with the ListAllDnsRecordsOptions.Match property.
// Whether to match all (all) or atleast one search parameter (any).
const (
	ListAllDnsRecordsOptions_Match_All = "all"
	ListAllDnsRecordsOptions_Match_Any = "any"
)

// NewListAllDnsRecordsOptions : Instantiate ListAllDnsRecordsOptions
func (*DnsRecordsV1) NewListAllDnsRecordsOptions() *ListAllDnsRecordsOptions {
	return &ListAllDnsRecordsOptions{}
}

// SetType : Allow user to set Type
func (options *ListAllDnsRecordsOptions) SetType(typeVar string) *ListAllDnsRecordsOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetName : Allow user to set Name
func (options *ListAllDnsRecordsOptions) SetName(name string) *ListAllDnsRecordsOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetContent : Allow user to set Content
func (options *ListAllDnsRecordsOptions) SetContent(content string) *ListAllDnsRecordsOptions {
	options.Content = core.StringPtr(content)
	return options
}

// SetPage : Allow user to set Page
func (options *ListAllDnsRecordsOptions) SetPage(page int64) *ListAllDnsRecordsOptions {
	options.Page = core.Int64Ptr(page)
	return options
}

// SetPerPage : Allow user to set PerPage
func (options *ListAllDnsRecordsOptions) SetPerPage(perPage int64) *ListAllDnsRecordsOptions {
	options.PerPage = core.Int64Ptr(perPage)
	return options
}

// SetOrder : Allow user to set Order
func (options *ListAllDnsRecordsOptions) SetOrder(order string) *ListAllDnsRecordsOptions {
	options.Order = core.StringPtr(order)
	return options
}

// SetDirection : Allow user to set Direction
func (options *ListAllDnsRecordsOptions) SetDirection(direction string) *ListAllDnsRecordsOptions {
	options.Direction = core.StringPtr(direction)
	return options
}

// SetMatch : Allow user to set Match
func (options *ListAllDnsRecordsOptions) SetMatch(match string) *ListAllDnsRecordsOptions {
	options.Match = core.StringPtr(match)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListAllDnsRecordsOptions) SetHeaders(param map[string]string) *ListAllDnsRecordsOptions {
	options.Headers = param
	return options
}

// UpdateDnsRecordOptions : The UpdateDnsRecord options.
type UpdateDnsRecordOptions struct {
	// Identifier of DNS record.
	DnsrecordIdentifier *string `json:"dnsrecord_identifier" validate:"required,ne="`

	// Required for all record types except SRV.
	Name *string `json:"name,omitempty"`

	// dns record type.
	Type *string `json:"type,omitempty"`

	// dns record ttl value.
	TTL *int64 `json:"ttl,omitempty"`

	// content of dns record.
	Content *string `json:"content,omitempty"`

	// For MX records only.
	Priority *int64 `json:"priority,omitempty"`

	// proxied.
	Proxied *bool `json:"proxied,omitempty"`

	// For LOC, SRV and CAA records only.
	Data interface{} `json:"data,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateDnsRecordOptions.Type property.
// dns record type.
const (
	UpdateDnsRecordOptions_Type_A     = "A"
	UpdateDnsRecordOptions_Type_Aaaa  = "AAAA"
	UpdateDnsRecordOptions_Type_Caa   = "CAA"
	UpdateDnsRecordOptions_Type_Cname = "CNAME"
	UpdateDnsRecordOptions_Type_Loc   = "LOC"
	UpdateDnsRecordOptions_Type_Mx    = "MX"
	UpdateDnsRecordOptions_Type_Ns    = "NS"
	UpdateDnsRecordOptions_Type_Spf   = "SPF"
	UpdateDnsRecordOptions_Type_Srv   = "SRV"
	UpdateDnsRecordOptions_Type_Txt   = "TXT"
)

// NewUpdateDnsRecordOptions : Instantiate UpdateDnsRecordOptions
func (*DnsRecordsV1) NewUpdateDnsRecordOptions(dnsrecordIdentifier string) *UpdateDnsRecordOptions {
	return &UpdateDnsRecordOptions{
		DnsrecordIdentifier: core.StringPtr(dnsrecordIdentifier),
	}
}

// SetDnsrecordIdentifier : Allow user to set DnsrecordIdentifier
func (options *UpdateDnsRecordOptions) SetDnsrecordIdentifier(dnsrecordIdentifier string) *UpdateDnsRecordOptions {
	options.DnsrecordIdentifier = core.StringPtr(dnsrecordIdentifier)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateDnsRecordOptions) SetName(name string) *UpdateDnsRecordOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetType : Allow user to set Type
func (options *UpdateDnsRecordOptions) SetType(typeVar string) *UpdateDnsRecordOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetTTL : Allow user to set TTL
func (options *UpdateDnsRecordOptions) SetTTL(ttl int64) *UpdateDnsRecordOptions {
	options.TTL = core.Int64Ptr(ttl)
	return options
}

// SetContent : Allow user to set Content
func (options *UpdateDnsRecordOptions) SetContent(content string) *UpdateDnsRecordOptions {
	options.Content = core.StringPtr(content)
	return options
}

// SetPriority : Allow user to set Priority
func (options *UpdateDnsRecordOptions) SetPriority(priority int64) *UpdateDnsRecordOptions {
	options.Priority = core.Int64Ptr(priority)
	return options
}

// SetProxied : Allow user to set Proxied
func (options *UpdateDnsRecordOptions) SetProxied(proxied bool) *UpdateDnsRecordOptions {
	options.Proxied = core.BoolPtr(proxied)
	return options
}

// SetData : Allow user to set Data
func (options *UpdateDnsRecordOptions) SetData(data interface{}) *UpdateDnsRecordOptions {
	options.Data = data
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDnsRecordOptions) SetHeaders(param map[string]string) *UpdateDnsRecordOptions {
	options.Headers = param
	return options
}

// DeleteDnsrecordResp : dns record delete response.
type DeleteDnsrecordResp struct {
	// success response.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result.
	Result *DeleteDnsrecordRespResult `json:"result" validate:"required"`
}

// UnmarshalDeleteDnsrecordResp unmarshals an instance of DeleteDnsrecordResp from the specified map of raw messages.
func UnmarshalDeleteDnsrecordResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteDnsrecordResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDeleteDnsrecordRespResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DnsrecordDetails : dns record details.
type DnsrecordDetails struct {
	// dns record identifier.
	ID *string `json:"id,omitempty"`

	// created on.
	CreatedOn *string `json:"created_on,omitempty"`

	// modified date.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// dns record name.
	Name *string `json:"name,omitempty"`

	// dns record type.
	Type *string `json:"type,omitempty"`

	// dns record content.
	Content *string `json:"content,omitempty"`

	// zone identifier.
	ZoneID *string `json:"zone_id,omitempty"`

	// zone name.
	ZoneName *string `json:"zone_name,omitempty"`

	// proxiable.
	Proxiable *bool `json:"proxiable,omitempty"`

	// proxied.
	Proxied *bool `json:"proxied,omitempty"`

	// dns record ttl value.
	TTL *int64 `json:"ttl,omitempty"`

	// Relevant only to MX type records.
	Priority *int64 `json:"priority,omitempty"`

	// Data details for the DNS record. Only for LOC, SRV, CAA records.
	Data interface{} `json:"data,omitempty"`
}

// Constants associated with the DnsrecordDetails.Type property.
// dns record type.
const (
	DnsrecordDetails_Type_A     = "A"
	DnsrecordDetails_Type_Aaaa  = "AAAA"
	DnsrecordDetails_Type_Caa   = "CAA"
	DnsrecordDetails_Type_Cname = "CNAME"
	DnsrecordDetails_Type_Loc   = "LOC"
	DnsrecordDetails_Type_Mx    = "MX"
	DnsrecordDetails_Type_Ns    = "NS"
	DnsrecordDetails_Type_Spf   = "SPF"
	DnsrecordDetails_Type_Srv   = "SRV"
	DnsrecordDetails_Type_Txt   = "TXT"
)

// UnmarshalDnsrecordDetails unmarshals an instance of DnsrecordDetails from the specified map of raw messages.
func UnmarshalDnsrecordDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DnsrecordDetails)
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
	err = core.UnmarshalPrimitive(m, "content", &obj.Content)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "zone_id", &obj.ZoneID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "zone_name", &obj.ZoneName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "proxiable", &obj.Proxiable)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "proxied", &obj.Proxied)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ttl", &obj.TTL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "priority", &obj.Priority)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data", &obj.Data)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DnsrecordResp : dns record response.
type DnsrecordResp struct {
	// success response.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// dns record details.
	Result *DnsrecordDetails `json:"result" validate:"required"`
}

// UnmarshalDnsrecordResp unmarshals an instance of DnsrecordResp from the specified map of raw messages.
func UnmarshalDnsrecordResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DnsrecordResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDnsrecordDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDnsrecordsResp : dns records list response.
type ListDnsrecordsResp struct {
	// success response.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// dns record list.
	Result []DnsrecordDetails `json:"result" validate:"required"`

	// result information.
	ResultInfo *ResultInfo `json:"result_info" validate:"required"`
}

// UnmarshalListDnsrecordsResp unmarshals an instance of ListDnsrecordsResp from the specified map of raw messages.
func UnmarshalListDnsrecordsResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListDnsrecordsResp)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "errors", &obj.Errors)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "messages", &obj.Messages)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDnsrecordDetails)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "result_info", &obj.ResultInfo, UnmarshalResultInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ResultInfo : result information.
type ResultInfo struct {
	// page.
	Page *int64 `json:"page" validate:"required"`

	// per page.
	PerPage *int64 `json:"per_page" validate:"required"`

	// count.
	Count *int64 `json:"count" validate:"required"`

	// total count.
	TotalCount *int64 `json:"total_count" validate:"required"`
}

// UnmarshalResultInfo unmarshals an instance of ResultInfo from the specified map of raw messages.
func UnmarshalResultInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResultInfo)
	err = core.UnmarshalPrimitive(m, "page", &obj.Page)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "per_page", &obj.PerPage)
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
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
