# Zone

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Opaque zone id (string), assigned by the server, should not be interpreted by the application. Guaranteed to be safe for embedding in URLs. | [optional] [default to null]
**Name** | **string** | Name of the zone (e.g. “example.com.”) MUST have a trailing dot | [optional] [default to null]
**Type_** | **string** | Set to “Zone” | [optional] [default to null]
**Url** | **string** | API endpoint for this zone | [optional] [default to null]
**Kind** | **string** | Zone kind, one of “Native”, “Master”, “Slave” | [optional] [default to null]
**Rrsets** | [**[]RrSet**](RRSet.md) | RRSets in this zone | [optional] [default to null]
**Serial** | **int32** | The SOA serial number | [optional] [default to null]
**NotifiedSerial** | **int32** | The SOA serial notifications have been sent out for | [optional] [default to null]
**Masters** | **[]string** |  List of IP addresses configured as a master for this zone (“Slave” type zones only) | [optional] [default to null]
**Dnssec** | **bool** | Whether or not this zone is DNSSEC signed (inferred from presigned being true XOR presence of at least one cryptokey with active being true) | [optional] [default to null]
**Nsec3param** | **string** | The NSEC3PARAM record | [optional] [default to null]
**Nsec3narrow** | **bool** | Whether or not the zone uses NSEC3 narrow | [optional] [default to null]
**Presigned** | **bool** | Whether or not the zone is pre-signed | [optional] [default to null]
**SoaEdit** | **string** | The SOA-EDIT metadata item | [optional] [default to null]
**SoaEditApi** | **string** | The SOA-EDIT-API metadata item | [optional] [default to null]
**ApiRectify** | **bool** |  Whether or not the zone will be rectified on data changes via the API | [optional] [default to null]
**Zone** | **string** | MAY contain a BIND-style zone file when creating a zone | [optional] [default to null]
**Account** | **string** | MAY be set. Its value is defined by local policy | [optional] [default to null]
**Nameservers** | **[]string** | MAY be sent in client bodies during creation, and MUST NOT be sent by the server. Simple list of strings of nameserver names, including the trailing dot. Not required for slave zones. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


