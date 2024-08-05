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

// ListTsigKeysRequest wrapper for the ListTsigKeys operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListTsigKeys.go.html to see an example of how to use ListTsigKeysRequest.
type ListTsigKeysRequest struct {

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of a resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The name of a resource.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The state of a resource.
	LifecycleState TsigKeySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field by which to sort TSIG keys. If unspecified, defaults to `timeCreated`.
	SortBy ListTsigKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the resources.
	SortOrder ListTsigKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListTsigKeysScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTsigKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTsigKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTsigKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTsigKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTsigKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTsigKeySummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetTsigKeySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTsigKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTsigKeysSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTsigKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTsigKeysSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTsigKeysScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetListTsigKeysScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTsigKeysResponse wrapper for the ListTsigKeys operation
type ListTsigKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TsigKeySummary instances
	Items []TsigKeySummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTsigKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTsigKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTsigKeysSortByEnum Enum with underlying type: string
type ListTsigKeysSortByEnum string

// Set of constants representing the allowable values for ListTsigKeysSortByEnum
const (
	ListTsigKeysSortByName        ListTsigKeysSortByEnum = "name"
	ListTsigKeysSortByTimecreated ListTsigKeysSortByEnum = "timeCreated"
)

var mappingListTsigKeysSortByEnum = map[string]ListTsigKeysSortByEnum{
	"name":        ListTsigKeysSortByName,
	"timeCreated": ListTsigKeysSortByTimecreated,
}

var mappingListTsigKeysSortByEnumLowerCase = map[string]ListTsigKeysSortByEnum{
	"name":        ListTsigKeysSortByName,
	"timecreated": ListTsigKeysSortByTimecreated,
}

// GetListTsigKeysSortByEnumValues Enumerates the set of values for ListTsigKeysSortByEnum
func GetListTsigKeysSortByEnumValues() []ListTsigKeysSortByEnum {
	values := make([]ListTsigKeysSortByEnum, 0)
	for _, v := range mappingListTsigKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTsigKeysSortByEnumStringValues Enumerates the set of values in String for ListTsigKeysSortByEnum
func GetListTsigKeysSortByEnumStringValues() []string {
	return []string{
		"name",
		"timeCreated",
	}
}

// GetMappingListTsigKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTsigKeysSortByEnum(val string) (ListTsigKeysSortByEnum, bool) {
	enum, ok := mappingListTsigKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTsigKeysSortOrderEnum Enum with underlying type: string
type ListTsigKeysSortOrderEnum string

// Set of constants representing the allowable values for ListTsigKeysSortOrderEnum
const (
	ListTsigKeysSortOrderAsc  ListTsigKeysSortOrderEnum = "ASC"
	ListTsigKeysSortOrderDesc ListTsigKeysSortOrderEnum = "DESC"
)

var mappingListTsigKeysSortOrderEnum = map[string]ListTsigKeysSortOrderEnum{
	"ASC":  ListTsigKeysSortOrderAsc,
	"DESC": ListTsigKeysSortOrderDesc,
}

var mappingListTsigKeysSortOrderEnumLowerCase = map[string]ListTsigKeysSortOrderEnum{
	"asc":  ListTsigKeysSortOrderAsc,
	"desc": ListTsigKeysSortOrderDesc,
}

// GetListTsigKeysSortOrderEnumValues Enumerates the set of values for ListTsigKeysSortOrderEnum
func GetListTsigKeysSortOrderEnumValues() []ListTsigKeysSortOrderEnum {
	values := make([]ListTsigKeysSortOrderEnum, 0)
	for _, v := range mappingListTsigKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTsigKeysSortOrderEnumStringValues Enumerates the set of values in String for ListTsigKeysSortOrderEnum
func GetListTsigKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTsigKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTsigKeysSortOrderEnum(val string) (ListTsigKeysSortOrderEnum, bool) {
	enum, ok := mappingListTsigKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTsigKeysScopeEnum Enum with underlying type: string
type ListTsigKeysScopeEnum string

// Set of constants representing the allowable values for ListTsigKeysScopeEnum
const (
	ListTsigKeysScopeGlobal  ListTsigKeysScopeEnum = "GLOBAL"
	ListTsigKeysScopePrivate ListTsigKeysScopeEnum = "PRIVATE"
)

var mappingListTsigKeysScopeEnum = map[string]ListTsigKeysScopeEnum{
	"GLOBAL":  ListTsigKeysScopeGlobal,
	"PRIVATE": ListTsigKeysScopePrivate,
}

var mappingListTsigKeysScopeEnumLowerCase = map[string]ListTsigKeysScopeEnum{
	"global":  ListTsigKeysScopeGlobal,
	"private": ListTsigKeysScopePrivate,
}

// GetListTsigKeysScopeEnumValues Enumerates the set of values for ListTsigKeysScopeEnum
func GetListTsigKeysScopeEnumValues() []ListTsigKeysScopeEnum {
	values := make([]ListTsigKeysScopeEnum, 0)
	for _, v := range mappingListTsigKeysScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetListTsigKeysScopeEnumStringValues Enumerates the set of values in String for ListTsigKeysScopeEnum
func GetListTsigKeysScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingListTsigKeysScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTsigKeysScopeEnum(val string) (ListTsigKeysScopeEnum, bool) {
	enum, ok := mappingListTsigKeysScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
