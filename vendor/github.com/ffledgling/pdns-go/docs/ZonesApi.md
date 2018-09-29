# \ZonesApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AxfrExportZone**](ZonesApi.md#AxfrExportZone) | **Get** /servers/{server_id}/zones/{zone_id}/export | Returns the zone in AXFR format.
[**AxfrRetrieveZone**](ZonesApi.md#AxfrRetrieveZone) | **Put** /servers/{server_id}/zones/{zone_id}/axfr-retrieve | Send a DNS NOTIFY to all slaves.
[**CheckZone**](ZonesApi.md#CheckZone) | **Get** /servers/{server_id}/zones/{zone_id}/check | Verify zone contents/configuration.
[**CreateZone**](ZonesApi.md#CreateZone) | **Post** /servers/{server_id}/zones | Creates a new domain, returns the Zone on creation.
[**DeleteZone**](ZonesApi.md#DeleteZone) | **Delete** /servers/{server_id}/zones/{zone_id} | Deletes this zone, all attached metadata and rrsets.
[**ListZone**](ZonesApi.md#ListZone) | **Get** /servers/{server_id}/zones/{zone_id} | zone managed by a server
[**ListZones**](ZonesApi.md#ListZones) | **Get** /servers/{server_id}/zones | List all Zones in a server
[**NotifyZone**](ZonesApi.md#NotifyZone) | **Put** /servers/{server_id}/zones/{zone_id}/notify | Send a DNS NOTIFY to all slaves.
[**PatchZone**](ZonesApi.md#PatchZone) | **Patch** /servers/{server_id}/zones/{zone_id} | Creates/modifies/deletes RRsets present in the payload and their comments. Returns 204 No Content on success.
[**PutZone**](ZonesApi.md#PutZone) | **Put** /servers/{server_id}/zones/{zone_id} | Modifies basic zone data (metadata).
[**RectifyZone**](ZonesApi.md#RectifyZone) | **Put** /servers/{server_id}/zones/{zone_id}/rectify | Rectify the zone data.


# **AxfrExportZone**
> string AxfrExportZone(ctx, serverId, zoneId)
Returns the zone in AXFR format.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AxfrRetrieveZone**
> AxfrRetrieveZone(ctx, serverId, zoneId)
Send a DNS NOTIFY to all slaves.

Fails when zone kind is not Master or Slave, or master and slave are disabled in the configuration. Only works for Slave if renotify is on. Clients MUST NOT send a body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckZone**
> string CheckZone(ctx, serverId, zoneId)
Verify zone contents/configuration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateZone**
> Zone CreateZone(ctx, serverId, optional)
Creates a new domain, returns the Zone on creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **string**| The id of the server to retrieve | 
 **rrsets** | **bool**| “true” (default) or “false”, whether to include the “rrsets” in the response Zone object. | [default to true]

### Return type

[**Zone**](Zone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteZone**
> DeleteZone(ctx, serverId, zoneId)
Deletes this zone, all attached metadata and rrsets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListZone**
> Zone ListZone(ctx, serverId, zoneId)
zone managed by a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

[**Zone**](Zone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListZones**
> []Zone ListZones(ctx, serverId)
List all Zones in a server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 

### Return type

[**[]Zone**](Zone.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NotifyZone**
> NotifyZone(ctx, serverId, zoneId)
Send a DNS NOTIFY to all slaves.

Fails when zone kind is not Master or Slave, or master and slave are disabled in the configuration. Only works for Slave if renotify is on. Clients MUST NOT send a body.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchZone**
> PatchZone(ctx, serverId, zoneId, zoneStruct)
Creates/modifies/deletes RRsets present in the payload and their comments. Returns 204 No Content on success.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**|  | 
  **zoneStruct** | [**Zone**](Zone.md)| The zone struct to patch with | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutZone**
> PutZone(ctx, serverId, zoneId, zoneStruct)
Modifies basic zone data (metadata).

Allowed fields in client body: all except id, url and name. Returns 204 No Content on success.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**|  | 
  **zoneStruct** | [**Zone**](Zone.md)| The zone struct to patch with | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RectifyZone**
> string RectifyZone(ctx, serverId, zoneId)
Rectify the zone data.

This does not take into account the API-RECTIFY metadata. Fails on slave zones and zones that do not have DNSSEC.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

**string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

