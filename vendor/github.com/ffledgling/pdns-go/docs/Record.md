# Record

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** | The content of this record | [default to null]
**Disabled** | **bool** | Whether or not this record is disabled | [default to null]
**SetPtr** | **bool** | If set to true, the server will find the matching reverse zone and create a PTR there. Existing PTR records are replaced. If no matching reverse Zone, an error is thrown. Only valid in client bodies, only valid for A and AAAA types. Not returned by the server. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


