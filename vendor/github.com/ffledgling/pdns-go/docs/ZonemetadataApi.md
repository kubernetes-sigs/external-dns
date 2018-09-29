# \ZonemetadataApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetadata**](ZonemetadataApi.md#CreateMetadata) | **Post** /servers/{server_id}/zones/{zone_id}/metadata | Creates a set of metadata entries
[**DeleteMetadata**](ZonemetadataApi.md#DeleteMetadata) | **Delete** /servers/{server_id}/zones/{zone_id}/metadata/{metadata_kind} | Delete all items of a single kind of domain metadata.
[**GetMetadata**](ZonemetadataApi.md#GetMetadata) | **Get** /servers/{server_id}/zones/{zone_id}/metadata/{metadata_kind} | Get the content of a single kind of domain metadata as a list of MetaData objects.
[**ListMetadata**](ZonemetadataApi.md#ListMetadata) | **Get** /servers/{server_id}/zones/{zone_id}/metadata | Get all the MetaData associated with the zone.
[**ModifyMetadata**](ZonemetadataApi.md#ModifyMetadata) | **Put** /servers/{server_id}/zones/{zone_id}/metadata/{metadata_kind} | Modify the content of a single kind of domain metadata.


# **CreateMetadata**
> CreateMetadata(ctx, serverId, zoneId, metadata)
Creates a set of metadata entries

Creates a set of metadata entries of given kind for the zone. Existing metadata entries for the zone with the same kind are not overwritten.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**|  | 
  **metadata** | [**[]Metadata**](Metadata.md)| List of metadata to add/create | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteMetadata**
> DeleteMetadata(ctx, serverId, zoneId, metadataKind)
Delete all items of a single kind of domain metadata.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 
  **metadataKind** | **string**| ??? | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMetadata**
> Metadata GetMetadata(ctx, serverId, zoneId, metadataKind)
Get the content of a single kind of domain metadata as a list of MetaData objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 
  **metadataKind** | **string**| ??? | 

### Return type

[**Metadata**](Metadata.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListMetadata**
> []Metadata ListMetadata(ctx, serverId, zoneId)
Get all the MetaData associated with the zone.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

[**[]Metadata**](Metadata.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyMetadata**
> ModifyMetadata(ctx, serverId, zoneId, metadataKind, metadata)
Modify the content of a single kind of domain metadata.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**|  | 
  **metadataKind** | **string**| The kind of metadata | 
  **metadata** | [**Metadata**](Metadata.md)| metadata to add/create | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

