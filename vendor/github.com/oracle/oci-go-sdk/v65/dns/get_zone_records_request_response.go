// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetZoneRecordsRequest wrapper for the GetZoneRecords operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetZoneRecords.go.html to see an example of how to use GetZoneRecordsRequest.
type GetZoneRecordsRequest struct {

	// The name or OCID of the target zone.
	ZoneNameOrId *string `mandatory:"true" contributesTo:"path" name:"zoneNameOrId"`

	// The `If-None-Match` header field makes the request method conditional on
	// the absence of any current representation of the target resource, when
	// the field-value is `*`, or having a selected representation with an
	// entity-tag that does not match any of those listed in the field-value.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"If-None-Match"`

	// The `If-Modified-Since` header field makes a GET or HEAD request method
	// conditional on the selected representation's modification date being more
	// recent than the date provided in the field-value.  Transfer of the
	// selected representation's data is avoided if that data has not changed.
	IfModifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Modified-Since"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The version of the zone for which data is requested.
	ZoneVersion *string `mandatory:"false" contributesTo:"query" name:"zoneVersion"`

	// Search by domain.
	// Will match any record whose domain (case-insensitive) equals the provided value.
	Domain *string `mandatory:"false" contributesTo:"query" name:"domain"`

	// Search by domain.
	// Will match any record whose domain (case-insensitive) contains the provided value.
	DomainContains *string `mandatory:"false" contributesTo:"query" name:"domainContains"`

	// Search by record type.
	// Will match any record whose type (https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4) (case-insensitive) equals the provided value.
	Rtype *string `mandatory:"false" contributesTo:"query" name:"rtype"`

	// The field by which to sort records.
	SortBy GetZoneRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the resources.
	SortOrder GetZoneRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The OCID of the compartment the zone belongs to.
	// This parameter is deprecated and should be omitted.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope GetZoneRecordsScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the zone is associated with. Required when accessing a private zone by name.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetZoneRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetZoneRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetZoneRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetZoneRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetZoneRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetZoneRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetZoneRecordsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetZoneRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetZoneRecordsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetZoneRecordsScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetGetZoneRecordsScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetZoneRecordsResponse wrapper for the GetZoneRecords operation
type GetZoneRecordsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecordCollection instances
	RecordCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response GetZoneRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetZoneRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetZoneRecordsSortByEnum Enum with underlying type: string
type GetZoneRecordsSortByEnum string

// Set of constants representing the allowable values for GetZoneRecordsSortByEnum
const (
	GetZoneRecordsSortByDomain GetZoneRecordsSortByEnum = "domain"
	GetZoneRecordsSortByRtype  GetZoneRecordsSortByEnum = "rtype"
	GetZoneRecordsSortByTtl    GetZoneRecordsSortByEnum = "ttl"
)

var mappingGetZoneRecordsSortByEnum = map[string]GetZoneRecordsSortByEnum{
	"domain": GetZoneRecordsSortByDomain,
	"rtype":  GetZoneRecordsSortByRtype,
	"ttl":    GetZoneRecordsSortByTtl,
}

var mappingGetZoneRecordsSortByEnumLowerCase = map[string]GetZoneRecordsSortByEnum{
	"domain": GetZoneRecordsSortByDomain,
	"rtype":  GetZoneRecordsSortByRtype,
	"ttl":    GetZoneRecordsSortByTtl,
}

// GetGetZoneRecordsSortByEnumValues Enumerates the set of values for GetZoneRecordsSortByEnum
func GetGetZoneRecordsSortByEnumValues() []GetZoneRecordsSortByEnum {
	values := make([]GetZoneRecordsSortByEnum, 0)
	for _, v := range mappingGetZoneRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetZoneRecordsSortByEnumStringValues Enumerates the set of values in String for GetZoneRecordsSortByEnum
func GetGetZoneRecordsSortByEnumStringValues() []string {
	return []string{
		"domain",
		"rtype",
		"ttl",
	}
}

// GetMappingGetZoneRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetZoneRecordsSortByEnum(val string) (GetZoneRecordsSortByEnum, bool) {
	enum, ok := mappingGetZoneRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetZoneRecordsSortOrderEnum Enum with underlying type: string
type GetZoneRecordsSortOrderEnum string

// Set of constants representing the allowable values for GetZoneRecordsSortOrderEnum
const (
	GetZoneRecordsSortOrderAsc  GetZoneRecordsSortOrderEnum = "ASC"
	GetZoneRecordsSortOrderDesc GetZoneRecordsSortOrderEnum = "DESC"
)

var mappingGetZoneRecordsSortOrderEnum = map[string]GetZoneRecordsSortOrderEnum{
	"ASC":  GetZoneRecordsSortOrderAsc,
	"DESC": GetZoneRecordsSortOrderDesc,
}

var mappingGetZoneRecordsSortOrderEnumLowerCase = map[string]GetZoneRecordsSortOrderEnum{
	"asc":  GetZoneRecordsSortOrderAsc,
	"desc": GetZoneRecordsSortOrderDesc,
}

// GetGetZoneRecordsSortOrderEnumValues Enumerates the set of values for GetZoneRecordsSortOrderEnum
func GetGetZoneRecordsSortOrderEnumValues() []GetZoneRecordsSortOrderEnum {
	values := make([]GetZoneRecordsSortOrderEnum, 0)
	for _, v := range mappingGetZoneRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetZoneRecordsSortOrderEnumStringValues Enumerates the set of values in String for GetZoneRecordsSortOrderEnum
func GetGetZoneRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetZoneRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetZoneRecordsSortOrderEnum(val string) (GetZoneRecordsSortOrderEnum, bool) {
	enum, ok := mappingGetZoneRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetZoneRecordsScopeEnum Enum with underlying type: string
type GetZoneRecordsScopeEnum string

// Set of constants representing the allowable values for GetZoneRecordsScopeEnum
const (
	GetZoneRecordsScopeGlobal  GetZoneRecordsScopeEnum = "GLOBAL"
	GetZoneRecordsScopePrivate GetZoneRecordsScopeEnum = "PRIVATE"
)

var mappingGetZoneRecordsScopeEnum = map[string]GetZoneRecordsScopeEnum{
	"GLOBAL":  GetZoneRecordsScopeGlobal,
	"PRIVATE": GetZoneRecordsScopePrivate,
}

var mappingGetZoneRecordsScopeEnumLowerCase = map[string]GetZoneRecordsScopeEnum{
	"global":  GetZoneRecordsScopeGlobal,
	"private": GetZoneRecordsScopePrivate,
}

// GetGetZoneRecordsScopeEnumValues Enumerates the set of values for GetZoneRecordsScopeEnum
func GetGetZoneRecordsScopeEnumValues() []GetZoneRecordsScopeEnum {
	values := make([]GetZoneRecordsScopeEnum, 0)
	for _, v := range mappingGetZoneRecordsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetZoneRecordsScopeEnumStringValues Enumerates the set of values in String for GetZoneRecordsScopeEnum
func GetGetZoneRecordsScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingGetZoneRecordsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetZoneRecordsScopeEnum(val string) (GetZoneRecordsScopeEnum, bool) {
	enum, ok := mappingGetZoneRecordsScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
