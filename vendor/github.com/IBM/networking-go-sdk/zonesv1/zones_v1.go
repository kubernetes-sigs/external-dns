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

// Package zonesv1 : Operations and models for the ZonesV1 service
package zonesv1

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

// ZonesV1 : CIS Zones
//
// Version: 1.0.1
type ZonesV1 struct {
	Service *core.BaseService

	// Full url-encoded CRN of the service instance.
	Crn *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.cis.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "zones"

// ZonesV1Options : Service options
type ZonesV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// Full url-encoded CRN of the service instance.
	Crn *string `validate:"required"`
}

// NewZonesV1UsingExternalConfig : constructs an instance of ZonesV1 with passed in options and external configuration.
func NewZonesV1UsingExternalConfig(options *ZonesV1Options) (zones *ZonesV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	zones, err = NewZonesV1(options)
	if err != nil {
		return
	}

	err = zones.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = zones.Service.SetServiceURL(options.URL)
	}
	return
}

// NewZonesV1 : constructs an instance of ZonesV1 with passed in options.
func NewZonesV1(options *ZonesV1Options) (service *ZonesV1, err error) {
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

	service = &ZonesV1{
		Service: baseService,
		Crn:     options.Crn,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "zones" suitable for processing requests.
func (zones *ZonesV1) Clone() *ZonesV1 {
	if core.IsNil(zones) {
		return nil
	}
	clone := *zones
	clone.Service = zones.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (zones *ZonesV1) SetServiceURL(url string) error {
	return zones.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (zones *ZonesV1) GetServiceURL() string {
	return zones.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (zones *ZonesV1) SetDefaultHeaders(headers http.Header) {
	zones.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (zones *ZonesV1) SetEnableGzipCompression(enableGzip bool) {
	zones.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (zones *ZonesV1) GetEnableGzipCompression() bool {
	return zones.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (zones *ZonesV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	zones.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (zones *ZonesV1) DisableRetries() {
	zones.Service.DisableRetries()
}

// ListZones : List all zones
// List all zones for a service instance.
func (zones *ZonesV1) ListZones(listZonesOptions *ListZonesOptions) (result *ListZonesResp, response *core.DetailedResponse, err error) {
	return zones.ListZonesWithContext(context.Background(), listZonesOptions)
}

// ListZonesWithContext is an alternate form of the ListZones method which supports a Context parameter
func (zones *ZonesV1) ListZonesWithContext(ctx context.Context, listZonesOptions *ListZonesOptions) (result *ListZonesResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listZonesOptions, "listZonesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *zones.Crn,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zones.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zones.Service.Options.URL, `/v1/{crn}/zones`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listZonesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones", "V1", "ListZones")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listZonesOptions.Page != nil {
		builder.AddQuery("page", fmt.Sprint(*listZonesOptions.Page))
	}
	if listZonesOptions.PerPage != nil {
		builder.AddQuery("per_page", fmt.Sprint(*listZonesOptions.PerPage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListZonesResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateZone : Create zone
// Add a new zone for a given service instance.
func (zones *ZonesV1) CreateZone(createZoneOptions *CreateZoneOptions) (result *ZoneResp, response *core.DetailedResponse, err error) {
	return zones.CreateZoneWithContext(context.Background(), createZoneOptions)
}

// CreateZoneWithContext is an alternate form of the CreateZone method which supports a Context parameter
func (zones *ZonesV1) CreateZoneWithContext(ctx context.Context, createZoneOptions *CreateZoneOptions) (result *ZoneResp, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(createZoneOptions, "createZoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn": *zones.Crn,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zones.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zones.Service.Options.URL, `/v1/{crn}/zones`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones", "V1", "CreateZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createZoneOptions.Name != nil {
		body["name"] = createZoneOptions.Name
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
	response, err = zones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZoneResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// DeleteZone : Delete zone
// Delete a zone given its id.
func (zones *ZonesV1) DeleteZone(deleteZoneOptions *DeleteZoneOptions) (result *DeleteZoneResp, response *core.DetailedResponse, err error) {
	return zones.DeleteZoneWithContext(context.Background(), deleteZoneOptions)
}

// DeleteZoneWithContext is an alternate form of the DeleteZone method which supports a Context parameter
func (zones *ZonesV1) DeleteZoneWithContext(ctx context.Context, deleteZoneOptions *DeleteZoneOptions) (result *DeleteZoneResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteZoneOptions, "deleteZoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteZoneOptions, "deleteZoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zones.Crn,
		"zone_identifier": *deleteZoneOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zones.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zones.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones", "V1", "DeleteZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDeleteZoneResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// GetZone : Get zone
// Get the details of a zone for a given service instance and given zone id.
func (zones *ZonesV1) GetZone(getZoneOptions *GetZoneOptions) (result *ZoneResp, response *core.DetailedResponse, err error) {
	return zones.GetZoneWithContext(context.Background(), getZoneOptions)
}

// GetZoneWithContext is an alternate form of the GetZone method which supports a Context parameter
func (zones *ZonesV1) GetZoneWithContext(ctx context.Context, getZoneOptions *GetZoneOptions) (result *ZoneResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getZoneOptions, "getZoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getZoneOptions, "getZoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zones.Crn,
		"zone_identifier": *getZoneOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zones.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zones.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones", "V1", "GetZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZoneResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// UpdateZone : Update zone
// Update the paused field of the zone.
func (zones *ZonesV1) UpdateZone(updateZoneOptions *UpdateZoneOptions) (result *ZoneResp, response *core.DetailedResponse, err error) {
	return zones.UpdateZoneWithContext(context.Background(), updateZoneOptions)
}

// UpdateZoneWithContext is an alternate form of the UpdateZone method which supports a Context parameter
func (zones *ZonesV1) UpdateZoneWithContext(ctx context.Context, updateZoneOptions *UpdateZoneOptions) (result *ZoneResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateZoneOptions, "updateZoneOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateZoneOptions, "updateZoneOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zones.Crn,
		"zone_identifier": *updateZoneOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zones.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zones.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateZoneOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones", "V1", "UpdateZone")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateZoneOptions.Paused != nil {
		body["paused"] = updateZoneOptions.Paused
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
	response, err = zones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZoneResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// ZoneActivationCheck : Check zone
// Perform activation check on zone for status.
func (zones *ZonesV1) ZoneActivationCheck(zoneActivationCheckOptions *ZoneActivationCheckOptions) (result *ZoneActivationcheckResp, response *core.DetailedResponse, err error) {
	return zones.ZoneActivationCheckWithContext(context.Background(), zoneActivationCheckOptions)
}

// ZoneActivationCheckWithContext is an alternate form of the ZoneActivationCheck method which supports a Context parameter
func (zones *ZonesV1) ZoneActivationCheckWithContext(ctx context.Context, zoneActivationCheckOptions *ZoneActivationCheckOptions) (result *ZoneActivationcheckResp, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(zoneActivationCheckOptions, "zoneActivationCheckOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(zoneActivationCheckOptions, "zoneActivationCheckOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"crn":             *zones.Crn,
		"zone_identifier": *zoneActivationCheckOptions.ZoneIdentifier,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = zones.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(zones.Service.Options.URL, `/v1/{crn}/zones/{zone_identifier}/activation_check`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range zoneActivationCheckOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("zones", "V1", "ZoneActivationCheck")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = zones.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalZoneActivationcheckResp)
	if err != nil {
		return
	}
	response.Result = result

	return
}

// CreateZoneOptions : The CreateZone options.
type CreateZoneOptions struct {
	// name.
	Name *string `json:"name,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateZoneOptions : Instantiate CreateZoneOptions
func (*ZonesV1) NewCreateZoneOptions() *CreateZoneOptions {
	return &CreateZoneOptions{}
}

// SetName : Allow user to set Name
func (options *CreateZoneOptions) SetName(name string) *CreateZoneOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateZoneOptions) SetHeaders(param map[string]string) *CreateZoneOptions {
	options.Headers = param
	return options
}

// DeleteZoneOptions : The DeleteZone options.
type DeleteZoneOptions struct {
	// Identifier of zone.
	ZoneIdentifier *string `json:"zone_identifier" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteZoneOptions : Instantiate DeleteZoneOptions
func (*ZonesV1) NewDeleteZoneOptions(zoneIdentifier string) *DeleteZoneOptions {
	return &DeleteZoneOptions{
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *DeleteZoneOptions) SetZoneIdentifier(zoneIdentifier string) *DeleteZoneOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteZoneOptions) SetHeaders(param map[string]string) *DeleteZoneOptions {
	options.Headers = param
	return options
}

// DeleteZoneRespResult : result.
type DeleteZoneRespResult struct {
	// id.
	ID *string `json:"id" validate:"required"`
}

// UnmarshalDeleteZoneRespResult unmarshals an instance of DeleteZoneRespResult from the specified map of raw messages.
func UnmarshalDeleteZoneRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteZoneRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetZoneOptions : The GetZone options.
type GetZoneOptions struct {
	// Zone identifier.
	ZoneIdentifier *string `json:"zone_identifier" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetZoneOptions : Instantiate GetZoneOptions
func (*ZonesV1) NewGetZoneOptions(zoneIdentifier string) *GetZoneOptions {
	return &GetZoneOptions{
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *GetZoneOptions) SetZoneIdentifier(zoneIdentifier string) *GetZoneOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetZoneOptions) SetHeaders(param map[string]string) *GetZoneOptions {
	options.Headers = param
	return options
}

// ListZonesOptions : The ListZones options.
type ListZonesOptions struct {
	// Page number of paginated results.
	Page *int64 `json:"page,omitempty"`

	// Maximum number of zones per page.
	PerPage *int64 `json:"per_page,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListZonesOptions : Instantiate ListZonesOptions
func (*ZonesV1) NewListZonesOptions() *ListZonesOptions {
	return &ListZonesOptions{}
}

// SetPage : Allow user to set Page
func (options *ListZonesOptions) SetPage(page int64) *ListZonesOptions {
	options.Page = core.Int64Ptr(page)
	return options
}

// SetPerPage : Allow user to set PerPage
func (options *ListZonesOptions) SetPerPage(perPage int64) *ListZonesOptions {
	options.PerPage = core.Int64Ptr(perPage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListZonesOptions) SetHeaders(param map[string]string) *ListZonesOptions {
	options.Headers = param
	return options
}

// UpdateZoneOptions : The UpdateZone options.
type UpdateZoneOptions struct {
	// Zone identifier.
	ZoneIdentifier *string `json:"zone_identifier" validate:"required,ne="`

	// paused.
	Paused *bool `json:"paused,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateZoneOptions : Instantiate UpdateZoneOptions
func (*ZonesV1) NewUpdateZoneOptions(zoneIdentifier string) *UpdateZoneOptions {
	return &UpdateZoneOptions{
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *UpdateZoneOptions) SetZoneIdentifier(zoneIdentifier string) *UpdateZoneOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetPaused : Allow user to set Paused
func (options *UpdateZoneOptions) SetPaused(paused bool) *UpdateZoneOptions {
	options.Paused = core.BoolPtr(paused)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateZoneOptions) SetHeaders(param map[string]string) *UpdateZoneOptions {
	options.Headers = param
	return options
}

// ZoneActivationCheckOptions : The ZoneActivationCheck options.
type ZoneActivationCheckOptions struct {
	// Identifier of zone.
	ZoneIdentifier *string `json:"zone_identifier" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewZoneActivationCheckOptions : Instantiate ZoneActivationCheckOptions
func (*ZonesV1) NewZoneActivationCheckOptions(zoneIdentifier string) *ZoneActivationCheckOptions {
	return &ZoneActivationCheckOptions{
		ZoneIdentifier: core.StringPtr(zoneIdentifier),
	}
}

// SetZoneIdentifier : Allow user to set ZoneIdentifier
func (options *ZoneActivationCheckOptions) SetZoneIdentifier(zoneIdentifier string) *ZoneActivationCheckOptions {
	options.ZoneIdentifier = core.StringPtr(zoneIdentifier)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ZoneActivationCheckOptions) SetHeaders(param map[string]string) *ZoneActivationCheckOptions {
	options.Headers = param
	return options
}

// ZoneActivationcheckRespResult : result.
type ZoneActivationcheckRespResult struct {
	// id.
	ID *string `json:"id" validate:"required"`
}

// UnmarshalZoneActivationcheckRespResult unmarshals an instance of ZoneActivationcheckRespResult from the specified map of raw messages.
func UnmarshalZoneActivationcheckRespResult(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ZoneActivationcheckRespResult)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteZoneResp : delete zone response.
type DeleteZoneResp struct {
	// success.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result.
	Result *DeleteZoneRespResult `json:"result" validate:"required"`
}

// UnmarshalDeleteZoneResp unmarshals an instance of DeleteZoneResp from the specified map of raw messages.
func UnmarshalDeleteZoneResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DeleteZoneResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalDeleteZoneRespResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListZonesResp : list zones response.
type ListZonesResp struct {
	// success.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// zone list.
	Result []ZoneDetails `json:"result" validate:"required"`

	// result information.
	ResultInfo *ResultInfo `json:"result_info" validate:"required"`
}

// UnmarshalListZonesResp unmarshals an instance of ListZonesResp from the specified map of raw messages.
func UnmarshalListZonesResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListZonesResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalZoneDetails)
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

// ZoneActivationcheckResp : zone activation check response.
type ZoneActivationcheckResp struct {
	// success.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// result.
	Result *ZoneActivationcheckRespResult `json:"result" validate:"required"`
}

// UnmarshalZoneActivationcheckResp unmarshals an instance of ZoneActivationcheckResp from the specified map of raw messages.
func UnmarshalZoneActivationcheckResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ZoneActivationcheckResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalZoneActivationcheckRespResult)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ZoneDetails : zone details.
type ZoneDetails struct {
	// id.
	ID *string `json:"id,omitempty"`

	// created date.
	CreatedOn *string `json:"created_on,omitempty"`

	// modified date.
	ModifiedOn *string `json:"modified_on,omitempty"`

	// name.
	Name *string `json:"name,omitempty"`

	// original registrar.
	OriginalRegistrar *string `json:"original_registrar,omitempty"`

	// orginal dns host.
	OriginalDnshost *string `json:"original_dnshost,omitempty"`

	// status.
	Status *string `json:"status,omitempty"`

	// paused.
	Paused *bool `json:"paused,omitempty"`

	// orginal name servers.
	OriginalNameServers []string `json:"original_name_servers,omitempty"`

	// name servers.
	NameServers []string `json:"name_servers,omitempty"`
}

// UnmarshalZoneDetails unmarshals an instance of ZoneDetails from the specified map of raw messages.
func UnmarshalZoneDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ZoneDetails)
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
	err = core.UnmarshalPrimitive(m, "original_registrar", &obj.OriginalRegistrar)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "original_dnshost", &obj.OriginalDnshost)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "paused", &obj.Paused)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "original_name_servers", &obj.OriginalNameServers)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name_servers", &obj.NameServers)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ZoneResp : zone response.
type ZoneResp struct {
	// success.
	Success *bool `json:"success" validate:"required"`

	// errors.
	Errors [][]string `json:"errors" validate:"required"`

	// messages.
	Messages [][]string `json:"messages" validate:"required"`

	// zone details.
	Result *ZoneDetails `json:"result" validate:"required"`
}

// UnmarshalZoneResp unmarshals an instance of ZoneResp from the specified map of raw messages.
func UnmarshalZoneResp(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ZoneResp)
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
	err = core.UnmarshalModel(m, "result", &obj.Result, UnmarshalZoneDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
