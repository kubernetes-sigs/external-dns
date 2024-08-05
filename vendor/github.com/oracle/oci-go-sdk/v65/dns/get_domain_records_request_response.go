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

// GetDomainRecordsRequest wrapper for the GetDomainRecords operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetDomainRecords.go.html to see an example of how to use GetDomainRecordsRequest.
type GetDomainRecordsRequest struct {

	// The name or OCID of the target zone.
	ZoneNameOrId *string `mandatory:"true" contributesTo:"path" name:"zoneNameOrId"`

	// The target fully-qualified domain name (FQDN) within the target zone.
	Domain *string `mandatory:"true" contributesTo:"path" name:"domain"`

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

	// Search by record type.
	// Will match any record whose type (https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4) (case-insensitive) equals the provided value.
	Rtype *string `mandatory:"false" contributesTo:"query" name:"rtype"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope GetDomainRecordsScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the zone is associated with. Required when accessing a private zone by name.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// The field by which to sort records.
	SortBy GetDomainRecordsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the resources.
	SortOrder GetDomainRecordsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The OCID of the compartment the zone belongs to.
	// This parameter is deprecated and should be omitted.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDomainRecordsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDomainRecordsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDomainRecordsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDomainRecordsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDomainRecordsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetDomainRecordsScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetGetDomainRecordsScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDomainRecordsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetDomainRecordsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetDomainRecordsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetDomainRecordsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDomainRecordsResponse wrapper for the GetDomainRecords operation
type GetDomainRecordsResponse struct {

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

func (response GetDomainRecordsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDomainRecordsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDomainRecordsScopeEnum Enum with underlying type: string
type GetDomainRecordsScopeEnum string

// Set of constants representing the allowable values for GetDomainRecordsScopeEnum
const (
	GetDomainRecordsScopeGlobal  GetDomainRecordsScopeEnum = "GLOBAL"
	GetDomainRecordsScopePrivate GetDomainRecordsScopeEnum = "PRIVATE"
)

var mappingGetDomainRecordsScopeEnum = map[string]GetDomainRecordsScopeEnum{
	"GLOBAL":  GetDomainRecordsScopeGlobal,
	"PRIVATE": GetDomainRecordsScopePrivate,
}

var mappingGetDomainRecordsScopeEnumLowerCase = map[string]GetDomainRecordsScopeEnum{
	"global":  GetDomainRecordsScopeGlobal,
	"private": GetDomainRecordsScopePrivate,
}

// GetGetDomainRecordsScopeEnumValues Enumerates the set of values for GetDomainRecordsScopeEnum
func GetGetDomainRecordsScopeEnumValues() []GetDomainRecordsScopeEnum {
	values := make([]GetDomainRecordsScopeEnum, 0)
	for _, v := range mappingGetDomainRecordsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDomainRecordsScopeEnumStringValues Enumerates the set of values in String for GetDomainRecordsScopeEnum
func GetGetDomainRecordsScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingGetDomainRecordsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDomainRecordsScopeEnum(val string) (GetDomainRecordsScopeEnum, bool) {
	enum, ok := mappingGetDomainRecordsScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDomainRecordsSortByEnum Enum with underlying type: string
type GetDomainRecordsSortByEnum string

// Set of constants representing the allowable values for GetDomainRecordsSortByEnum
const (
	GetDomainRecordsSortByRtype GetDomainRecordsSortByEnum = "rtype"
	GetDomainRecordsSortByTtl   GetDomainRecordsSortByEnum = "ttl"
)

var mappingGetDomainRecordsSortByEnum = map[string]GetDomainRecordsSortByEnum{
	"rtype": GetDomainRecordsSortByRtype,
	"ttl":   GetDomainRecordsSortByTtl,
}

var mappingGetDomainRecordsSortByEnumLowerCase = map[string]GetDomainRecordsSortByEnum{
	"rtype": GetDomainRecordsSortByRtype,
	"ttl":   GetDomainRecordsSortByTtl,
}

// GetGetDomainRecordsSortByEnumValues Enumerates the set of values for GetDomainRecordsSortByEnum
func GetGetDomainRecordsSortByEnumValues() []GetDomainRecordsSortByEnum {
	values := make([]GetDomainRecordsSortByEnum, 0)
	for _, v := range mappingGetDomainRecordsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDomainRecordsSortByEnumStringValues Enumerates the set of values in String for GetDomainRecordsSortByEnum
func GetGetDomainRecordsSortByEnumStringValues() []string {
	return []string{
		"rtype",
		"ttl",
	}
}

// GetMappingGetDomainRecordsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDomainRecordsSortByEnum(val string) (GetDomainRecordsSortByEnum, bool) {
	enum, ok := mappingGetDomainRecordsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetDomainRecordsSortOrderEnum Enum with underlying type: string
type GetDomainRecordsSortOrderEnum string

// Set of constants representing the allowable values for GetDomainRecordsSortOrderEnum
const (
	GetDomainRecordsSortOrderAsc  GetDomainRecordsSortOrderEnum = "ASC"
	GetDomainRecordsSortOrderDesc GetDomainRecordsSortOrderEnum = "DESC"
)

var mappingGetDomainRecordsSortOrderEnum = map[string]GetDomainRecordsSortOrderEnum{
	"ASC":  GetDomainRecordsSortOrderAsc,
	"DESC": GetDomainRecordsSortOrderDesc,
}

var mappingGetDomainRecordsSortOrderEnumLowerCase = map[string]GetDomainRecordsSortOrderEnum{
	"asc":  GetDomainRecordsSortOrderAsc,
	"desc": GetDomainRecordsSortOrderDesc,
}

// GetGetDomainRecordsSortOrderEnumValues Enumerates the set of values for GetDomainRecordsSortOrderEnum
func GetGetDomainRecordsSortOrderEnumValues() []GetDomainRecordsSortOrderEnum {
	values := make([]GetDomainRecordsSortOrderEnum, 0)
	for _, v := range mappingGetDomainRecordsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDomainRecordsSortOrderEnumStringValues Enumerates the set of values in String for GetDomainRecordsSortOrderEnum
func GetGetDomainRecordsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetDomainRecordsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDomainRecordsSortOrderEnum(val string) (GetDomainRecordsSortOrderEnum, bool) {
	enum, ok := mappingGetDomainRecordsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
