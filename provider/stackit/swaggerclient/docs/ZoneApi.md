# {{classname}}

All URIs are relative to *https://api.dns.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ProjectsProjectIdZonesGet**](ZoneApi.md#V1ProjectsProjectIdZonesGet) | **Get** /v1/projects/{projectId}/zones | All get selected zones
[**V1ProjectsProjectIdZonesPost**](ZoneApi.md#V1ProjectsProjectIdZonesPost) | **Post** /v1/projects/{projectId}/zones | Post create a new zone
[**V1ProjectsProjectIdZonesZoneIdDelete**](ZoneApi.md#V1ProjectsProjectIdZonesZoneIdDelete) | **Delete** /v1/projects/{projectId}/zones/{zoneId} | Delete delete a zone
[**V1ProjectsProjectIdZonesZoneIdGet**](ZoneApi.md#V1ProjectsProjectIdZonesZoneIdGet) | **Get** /v1/projects/{projectId}/zones/{zoneId} | Get a single zone
[**V1ProjectsProjectIdZonesZoneIdPut**](ZoneApi.md#V1ProjectsProjectIdZonesZoneIdPut) | **Put** /v1/projects/{projectId}/zones/{zoneId} | Put update an existing zone
[**V1ProjectsProjectIdZonesZoneIdRestoresPost**](ZoneApi.md#V1ProjectsProjectIdZonesZoneIdRestoresPost) | **Post** /v1/projects/{projectId}/zones/{zoneId}/restores | Restore  an inactive zone but will not restore the record sets

# **V1ProjectsProjectIdZonesGet**
> ZoneResponseZoneAll V1ProjectsProjectIdZonesGet(ctx, projectId, optional)
All get selected zones

All zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
 **optional** | ***ZoneApiV1ProjectsProjectIdZonesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ZoneApiV1ProjectsProjectIdZonesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| page | [default to 1]
 **pageSize** | **optional.Int32**| page size | [default to 100]
 **dnsNameEq** | **optional.String**| filter dns name equal | 
 **dnsNameLike** | **optional.String**| filter dns name like | 
 **typeEq** | **optional.String**| filter type | 
 **primaryNameServerEq** | **optional.String**| filter primary name server equal | 
 **primaryNameServerLike** | **optional.String**| filter primary name server like | 
 **isReverseZoneEq** | **optional.Bool**| filter reverse zone equal | 
 **activeEq** | **optional.Bool**| filter active equal | 
 **createdGt** | **optional.String**| filter created greater with utc timestamp | 
 **createdLt** | **optional.String**| filter created lesser with utc timestamp | 
 **createdGte** | **optional.String**| filter created greater equal with utc timestamp | 
 **createdLte** | **optional.String**| filter created lesser equal with utc timestamp | 
 **orderByDnsName** | **optional.String**| order by dns name | 
 **orderByCreated** | **optional.String**| order by created | 

### Return type

[**ZoneResponseZoneAll**](zone.ResponseZoneAll.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesPost**
> ZoneResponseZone V1ProjectsProjectIdZonesPost(ctx, body, projectId)
Post create a new zone

Post zone create a new zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ZoneZonePost**](ZoneZonePost.md)| zone to create | 
  **projectId** | **string**| project id | 

### Return type

[**ZoneResponseZone**](zone.ResponseZone.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdDelete**
> SerializerMessage V1ProjectsProjectIdZonesZoneIdDelete(ctx, projectId, zoneId)
Delete delete a zone

Delete delete a zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 

### Return type

[**SerializerMessage**](serializer.Message.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdGet**
> ZoneResponseZone V1ProjectsProjectIdZonesZoneIdGet(ctx, projectId, zoneId)
Get a single zone

Get zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 

### Return type

[**ZoneResponseZone**](zone.ResponseZone.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdPut**
> ZoneResponseZone V1ProjectsProjectIdZonesZoneIdPut(ctx, body, projectId, zoneId)
Put update an existing zone

Put update an existing zone

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DomainUpdateZone**](DomainUpdateZone.md)| zone to update | 
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 

### Return type

[**ZoneResponseZone**](zone.ResponseZone.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdRestoresPost**
> SerializerMessage V1ProjectsProjectIdZonesZoneIdRestoresPost(ctx, projectId, zoneId)
Restore  an inactive zone but will not restore the record sets

Restore  an inactive zone but will not restore the record sets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 

### Return type

[**SerializerMessage**](serializer.Message.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

