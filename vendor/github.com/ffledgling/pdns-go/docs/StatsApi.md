# \StatsApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStats**](StatsApi.md#GetStats) | **Get** /servers/{server_id}/statistics | Query statistics.


# **GetStats**
> []StatisticItem GetStats(ctx, serverId)
Query statistics.

Query PowerDNS internal statistics. Returns a list of StatisticItem elements.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 

### Return type

[**[]StatisticItem**](StatisticItem.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

