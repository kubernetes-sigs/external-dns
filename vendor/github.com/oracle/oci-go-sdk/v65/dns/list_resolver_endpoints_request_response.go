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

// ListResolverEndpointsRequest wrapper for the ListResolverEndpoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListResolverEndpoints.go.html to see an example of how to use ListResolverEndpointsRequest.
type ListResolverEndpointsRequest struct {

	// The OCID of the target resolver.
	ResolverId *string `mandatory:"true" contributesTo:"path" name:"resolverId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The name of a resource.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The order to sort the resources.
	SortOrder ListResolverEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort resolver endpoints.
	SortBy ListResolverEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The state of a resource.
	LifecycleState ResolverEndpointSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListResolverEndpointsScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResolverEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResolverEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResolverEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResolverEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResolverEndpointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResolverEndpointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResolverEndpointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResolverEndpointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResolverEndpointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResolverEndpointSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetResolverEndpointSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResolverEndpointsScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetListResolverEndpointsScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResolverEndpointsResponse wrapper for the ListResolverEndpoints operation
type ListResolverEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ResolverEndpointSummary instances
	Items []ResolverEndpointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResolverEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResolverEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResolverEndpointsSortOrderEnum Enum with underlying type: string
type ListResolverEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListResolverEndpointsSortOrderEnum
const (
	ListResolverEndpointsSortOrderAsc  ListResolverEndpointsSortOrderEnum = "ASC"
	ListResolverEndpointsSortOrderDesc ListResolverEndpointsSortOrderEnum = "DESC"
)

var mappingListResolverEndpointsSortOrderEnum = map[string]ListResolverEndpointsSortOrderEnum{
	"ASC":  ListResolverEndpointsSortOrderAsc,
	"DESC": ListResolverEndpointsSortOrderDesc,
}

var mappingListResolverEndpointsSortOrderEnumLowerCase = map[string]ListResolverEndpointsSortOrderEnum{
	"asc":  ListResolverEndpointsSortOrderAsc,
	"desc": ListResolverEndpointsSortOrderDesc,
}

// GetListResolverEndpointsSortOrderEnumValues Enumerates the set of values for ListResolverEndpointsSortOrderEnum
func GetListResolverEndpointsSortOrderEnumValues() []ListResolverEndpointsSortOrderEnum {
	values := make([]ListResolverEndpointsSortOrderEnum, 0)
	for _, v := range mappingListResolverEndpointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResolverEndpointsSortOrderEnumStringValues Enumerates the set of values in String for ListResolverEndpointsSortOrderEnum
func GetListResolverEndpointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResolverEndpointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResolverEndpointsSortOrderEnum(val string) (ListResolverEndpointsSortOrderEnum, bool) {
	enum, ok := mappingListResolverEndpointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResolverEndpointsSortByEnum Enum with underlying type: string
type ListResolverEndpointsSortByEnum string

// Set of constants representing the allowable values for ListResolverEndpointsSortByEnum
const (
	ListResolverEndpointsSortByName        ListResolverEndpointsSortByEnum = "name"
	ListResolverEndpointsSortByTimecreated ListResolverEndpointsSortByEnum = "timeCreated"
)

var mappingListResolverEndpointsSortByEnum = map[string]ListResolverEndpointsSortByEnum{
	"name":        ListResolverEndpointsSortByName,
	"timeCreated": ListResolverEndpointsSortByTimecreated,
}

var mappingListResolverEndpointsSortByEnumLowerCase = map[string]ListResolverEndpointsSortByEnum{
	"name":        ListResolverEndpointsSortByName,
	"timecreated": ListResolverEndpointsSortByTimecreated,
}

// GetListResolverEndpointsSortByEnumValues Enumerates the set of values for ListResolverEndpointsSortByEnum
func GetListResolverEndpointsSortByEnumValues() []ListResolverEndpointsSortByEnum {
	values := make([]ListResolverEndpointsSortByEnum, 0)
	for _, v := range mappingListResolverEndpointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResolverEndpointsSortByEnumStringValues Enumerates the set of values in String for ListResolverEndpointsSortByEnum
func GetListResolverEndpointsSortByEnumStringValues() []string {
	return []string{
		"name",
		"timeCreated",
	}
}

// GetMappingListResolverEndpointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResolverEndpointsSortByEnum(val string) (ListResolverEndpointsSortByEnum, bool) {
	enum, ok := mappingListResolverEndpointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResolverEndpointsScopeEnum Enum with underlying type: string
type ListResolverEndpointsScopeEnum string

// Set of constants representing the allowable values for ListResolverEndpointsScopeEnum
const (
	ListResolverEndpointsScopeGlobal  ListResolverEndpointsScopeEnum = "GLOBAL"
	ListResolverEndpointsScopePrivate ListResolverEndpointsScopeEnum = "PRIVATE"
)

var mappingListResolverEndpointsScopeEnum = map[string]ListResolverEndpointsScopeEnum{
	"GLOBAL":  ListResolverEndpointsScopeGlobal,
	"PRIVATE": ListResolverEndpointsScopePrivate,
}

var mappingListResolverEndpointsScopeEnumLowerCase = map[string]ListResolverEndpointsScopeEnum{
	"global":  ListResolverEndpointsScopeGlobal,
	"private": ListResolverEndpointsScopePrivate,
}

// GetListResolverEndpointsScopeEnumValues Enumerates the set of values for ListResolverEndpointsScopeEnum
func GetListResolverEndpointsScopeEnumValues() []ListResolverEndpointsScopeEnum {
	values := make([]ListResolverEndpointsScopeEnum, 0)
	for _, v := range mappingListResolverEndpointsScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetListResolverEndpointsScopeEnumStringValues Enumerates the set of values in String for ListResolverEndpointsScopeEnum
func GetListResolverEndpointsScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingListResolverEndpointsScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResolverEndpointsScopeEnum(val string) (ListResolverEndpointsScopeEnum, bool) {
	enum, ok := mappingListResolverEndpointsScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
