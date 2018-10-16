# \ConfigApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfig**](ConfigApi.md#GetConfig) | **Get** /servers/{server_id}/config | Returns all ConfigSettings for a single server
[**GetConfigSetting**](ConfigApi.md#GetConfigSetting) | **Get** /servers/{server_id}/config/{config_setting_name} | Returns a specific ConfigSetting for a single server


# **GetConfig**
> []ConfigSetting GetConfig(ctx, serverId)
Returns all ConfigSettings for a single server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 

### Return type

[**[]ConfigSetting**](ConfigSetting.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigSetting**
> ConfigSetting GetConfigSetting(ctx, serverId, configSettingName)
Returns a specific ConfigSetting for a single server

NOT IMPLEMENTED

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **configSettingName** | **string**| The name of the setting to retrieve | 

### Return type

[**ConfigSetting**](ConfigSetting.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

