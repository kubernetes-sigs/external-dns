# {{classname}}

All URIs are relative to *https://api.dns.stackit.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ProjectsProjectIdZonesZoneIdRrsetsGet**](RecordSetApi.md#V1ProjectsProjectIdZonesZoneIdRrsetsGet) | **Get** /v1/projects/{projectId}/zones/{zoneId}/rrsets | All get selected RRSets
[**V1ProjectsProjectIdZonesZoneIdRrsetsPost**](RecordSetApi.md#V1ProjectsProjectIdZonesZoneIdRrsetsPost) | **Post** /v1/projects/{projectId}/zones/{zoneId}/rrsets | Post record set
[**V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdDelete**](RecordSetApi.md#V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdDelete) | **Delete** /v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId} | Delete a record set
[**V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdGet**](RecordSetApi.md#V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdGet) | **Get** /v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId} | Get a single rrset
[**V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdPatch**](RecordSetApi.md#V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdPatch) | **Patch** /v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId} | Patch updates a record set
[**V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRecordsPatch**](RecordSetApi.md#V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRecordsPatch) | **Patch** /v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}/records | PatchRecords updates a record in a rrset
[**V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRestoresPost**](RecordSetApi.md#V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRestoresPost) | **Post** /v1/projects/{projectId}/zones/{zoneId}/rrsets/{rrSetId}/restores | Restore record set

# **V1ProjectsProjectIdZonesZoneIdRrsetsGet**
> RrsetResponseRrSetAll V1ProjectsProjectIdZonesZoneIdRrsetsGet(ctx, projectId, zoneId, optional)
All get selected RRSets

All RRSet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 
 **optional** | ***RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecordSetApiV1ProjectsProjectIdZonesZoneIdRrsetsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| page | [default to 1]
 **pageSize** | **optional.Int32**| page size | [default to 100]
 **nameEq** | **optional.String**| filter name equal | 
 **nameLike** | **optional.String**| filter name like | 
 **typeEq** | **optional.String**| filter type | 
 **activeEq** | **optional.Bool**| filter active equal | 
 **createdGt** | **optional.String**| filter created greater with utc timestamp | 
 **createdLt** | **optional.String**| filter created lesser with utc timestamp | 
 **createdGte** | **optional.String**| filter created greater equal with utc timestamp | 
 **createdLte** | **optional.String**| filter created lesser equal with utc timestamp | 
 **orderByName** | **optional.String**| order by name | 
 **orderByCreated** | **optional.String**| order by created | 

### Return type

[**RrsetResponseRrSetAll**](rrset.ResponseRRSetAll.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdRrsetsPost**
> RrsetResponseRrSet V1ProjectsProjectIdZonesZoneIdRrsetsPost(ctx, body, projectId, zoneId)
Post record set

Post record set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrsetRrSetPost**](RrsetRrSetPost.md)| record set to create | 
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 

### Return type

[**RrsetResponseRrSet**](rrset.ResponseRRSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdDelete**
> SerializerMessage V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdDelete(ctx, projectId, zoneId, rrSetId)
Delete a record set

Delete a record set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 
  **rrSetId** | **string**| record set id | 

### Return type

[**SerializerMessage**](serializer.Message.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdGet**
> RrsetResponseRrSet V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdGet(ctx, projectId, zoneId, rrSetId)
Get a single rrset

Get rrset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 
  **rrSetId** | **string**| record set id | 

### Return type

[**RrsetResponseRrSet**](rrset.ResponseRRSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdPatch**
> RrsetResponseRrSet V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdPatch(ctx, body, projectId, zoneId, rrSetId)
Patch updates a record set

Patch record set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrsetRrSetPatch**](RrsetRrSetPatch.md)| record set to patch | 
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 
  **rrSetId** | **string**| record set id | 

### Return type

[**RrsetResponseRrSet**](rrset.ResponseRRSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRecordsPatch**
> RrsetResponseRrSet V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRecordsPatch(ctx, body, projectId, zoneId, rrSetId)
PatchRecords updates a record in a rrset

PatchRecords rrset updates a record in a rrset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RrsetRecordPatch**](RrsetRecordPatch.md)| rrset to update | 
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 
  **rrSetId** | **string**| record set id | 

### Return type

[**RrsetResponseRrSet**](rrset.ResponseRRSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRestoresPost**
> SerializerMessage V1ProjectsProjectIdZonesZoneIdRrsetsRrSetIdRestoresPost(ctx, projectId, zoneId, rrSetId)
Restore record set

Restore record set

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **string**| project id | 
  **zoneId** | **string**| zone id | 
  **rrSetId** | **string**| record set id | 

### Return type

[**SerializerMessage**](serializer.Message.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

