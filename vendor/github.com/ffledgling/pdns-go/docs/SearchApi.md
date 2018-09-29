# \SearchApi

All URIs are relative to *http://localhost:8081/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchData**](SearchApi.md#SearchData) | **Get** /servers/{server_id}/search-data | Search the data inside PowerDNS
[**SearchLog**](SearchApi.md#SearchLog) | **Get** /servers/{server_id}/search-log | Query the log, filtered by search_term.


# **SearchData**
> SearchResults SearchData(ctx, serverId, q, max)
Search the data inside PowerDNS

Search the data inside PowerDNS for search_term and return at most max_results. This includes zones, records and comments. The * character can be used in search_term as a wildcard character and the ? character can be used as a wildcard for a single character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **q** | **string**| The string to search for | 
  **max** | **int32**| Maximum number of entries to return | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLog**
> []string SearchLog(ctx, serverId, q)
Query the log, filtered by search_term.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **serverId** | **string**| The id of the server to retrieve | 
  **q** | **string**| The string to search for | 

### Return type

**[]string**

### Authorization

[APIKeyHeader](../README.md#APIKeyHeader)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

