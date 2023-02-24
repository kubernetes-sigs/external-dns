# ZoneZonePost

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | **string** | access control list | [optional] [default to 0.0.0.0/0,::/0]
**ContactEmail** | **string** | contact e-mail for the zone | [optional] [default to *]
**DefaultTTL** | **int32** | default time to live | [optional] [default to null]
**Description** | **string** | description of the zone | [optional] [default to *]
**DnsName** | **string** | zone name | [default to null]
**ExpireTime** | **int32** | expire time | [optional] [default to null]
**IsReverseZone** | **bool** | if the zone is a reverse zone or not | [optional] [default to false]
**Name** | **string** | user given name | [default to null]
**Primaries** | **[]string** | primary name server for secondary zone | [optional] [default to null]
**PrimaryNameServer** | **string** | primary name server. FQDN | [optional] [default to null]
**RefreshTime** | **int32** | refresh time | [optional] [default to null]
**RetryTime** | **int32** | retry time | [optional] [default to null]
**Type_** | **string** | zone type | [optional] [default to TYPE_.PRIMARY]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

