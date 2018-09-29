# \ZonecryptokeyApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCryptokey**](ZonecryptokeyApi.md#CreateCryptokey) | **Post** /servers/{server_id}/zones/{zone_id}/cryptokeys | Creates a Cryptokey
[**DeleteCryptokey**](ZonecryptokeyApi.md#DeleteCryptokey) | **Delete** /servers/{server_id}/zones/{zone_id}/cryptokeys/{cryptokey_id} | This method deletes a key specified by cryptokey_id.
[**GetCryptokey**](ZonecryptokeyApi.md#GetCryptokey) | **Get** /servers/{server_id}/zones/{zone_id}/cryptokeys/{cryptokey_id} | Returns all data about the CryptoKey, including the privatekey.
[**ListCryptokeys**](ZonecryptokeyApi.md#ListCryptokeys) | **Get** /servers/{server_id}/zones/{zone_id}/cryptokeys | Get all CryptoKeys for a zone, except the privatekey
[**ModifyCryptokey**](ZonecryptokeyApi.md#ModifyCryptokey) | **Put** /servers/{server_id}/zones/{zone_id}/cryptokeys/{cryptokey_id} | This method (de)activates a key from zone_name specified by cryptokey_id


# **CreateCryptokey**
> Cryptokey CreateCryptokey(ctx, serverId, zoneId, cryptokey)
Creates a Cryptokey

This method adds a new key to a zone. The key can either be generated or imported by supplying the content parameter. if content, bits and algo are null, a key will be generated based on the default-ksk-algorithm and default-ksk-size settings for a KSK and the default-zsk-algorithm and default-zsk-size options for a ZSK.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**|  | 
  **cryptokey** | [**Cryptokey**](Cryptokey.md)| Add a Cryptokey | 

### Return type

[**Cryptokey**](Cryptokey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCryptokey**
> DeleteCryptokey(ctx, serverId, zoneId, cryptokeyId)
This method deletes a key specified by cryptokey_id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 
  **cryptokeyId** | **string**| The id value of the Cryptokey | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCryptokey**
> Cryptokey GetCryptokey(ctx, serverId, zoneId, cryptokeyId)
Returns all data about the CryptoKey, including the privatekey.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 
  **cryptokeyId** | **string**| The id value of the CryptoKey | 

### Return type

[**Cryptokey**](Cryptokey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCryptokeys**
> []Cryptokey ListCryptokeys(ctx, serverId, zoneId)
Get all CryptoKeys for a zone, except the privatekey

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**| The id of the zone to retrieve | 

### Return type

[**[]Cryptokey**](Cryptokey.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ModifyCryptokey**
> ModifyCryptokey(ctx, serverId, zoneId, cryptokeyId, cryptokey)
This method (de)activates a key from zone_name specified by cryptokey_id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **zoneId** | **string**|  | 
  **cryptokeyId** | **string**| Cryptokey to manipulate | 
  **cryptokey** | [**Cryptokey**](Cryptokey.md)| the Cryptokey | 

### Return type

 (empty response body)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

