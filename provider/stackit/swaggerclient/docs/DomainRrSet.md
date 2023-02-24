# DomainRrSet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | if the record set is active or not | [optional] [default to null]
**Comment** | **string** | comment | [optional] [default to null]
**Error_** | **string** | Error shows error in case create/update/delete failed | [optional] [default to null]
**Id** | **string** | rr set id | [default to null]
**Name** | **string** | name of the record which should be a valid domain according to rfc1035 Section 2.3.4 | [default to null]
**Records** | [**[]DomainRecord**](domain.Record.md) | records | [default to null]
**State** | **string** | record set state | [default to null]
**Ttl** | **int32** | time to live | [default to null]
**Type_** | **string** | record set type | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

