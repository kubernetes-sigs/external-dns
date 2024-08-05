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

// ListResolversRequest wrapper for the ListResolvers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListResolvers.go.html to see an example of how to use ListResolversRequest.
type ListResolversRequest struct {

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The displayName of a resource.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of a resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The order to sort the resources.
	SortOrder ListResolversSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort resolvers.
	SortBy ListResolversSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The state of a resource.
	LifecycleState ResolverSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListResolversScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResolversRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResolversRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResolversRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResolversRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResolversRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResolversSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResolversSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResolversSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResolversSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResolverSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetResolverSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResolversScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetListResolversScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResolversResponse wrapper for the ListResolvers operation
type ListResolversResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ResolverSummary instances
	Items []ResolverSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResolversResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResolversResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResolversSortOrderEnum Enum with underlying type: string
type ListResolversSortOrderEnum string

// Set of constants representing the allowable values for ListResolversSortOrderEnum
const (
	ListResolversSortOrderAsc  ListResolversSortOrderEnum = "ASC"
	ListResolversSortOrderDesc ListResolversSortOrderEnum = "DESC"
)

var mappingListResolversSortOrderEnum = map[string]ListResolversSortOrderEnum{
	"ASC":  ListResolversSortOrderAsc,
	"DESC": ListResolversSortOrderDesc,
}

var mappingListResolversSortOrderEnumLowerCase = map[string]ListResolversSortOrderEnum{
	"asc":  ListResolversSortOrderAsc,
	"desc": ListResolversSortOrderDesc,
}

// GetListResolversSortOrderEnumValues Enumerates the set of values for ListResolversSortOrderEnum
func GetListResolversSortOrderEnumValues() []ListResolversSortOrderEnum {
	values := make([]ListResolversSortOrderEnum, 0)
	for _, v := range mappingListResolversSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResolversSortOrderEnumStringValues Enumerates the set of values in String for ListResolversSortOrderEnum
func GetListResolversSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResolversSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResolversSortOrderEnum(val string) (ListResolversSortOrderEnum, bool) {
	enum, ok := mappingListResolversSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResolversSortByEnum Enum with underlying type: string
type ListResolversSortByEnum string

// Set of constants representing the allowable values for ListResolversSortByEnum
const (
	ListResolversSortByDisplayname ListResolversSortByEnum = "displayName"
	ListResolversSortByTimecreated ListResolversSortByEnum = "timeCreated"
)

var mappingListResolversSortByEnum = map[string]ListResolversSortByEnum{
	"displayName": ListResolversSortByDisplayname,
	"timeCreated": ListResolversSortByTimecreated,
}

var mappingListResolversSortByEnumLowerCase = map[string]ListResolversSortByEnum{
	"displayname": ListResolversSortByDisplayname,
	"timecreated": ListResolversSortByTimecreated,
}

// GetListResolversSortByEnumValues Enumerates the set of values for ListResolversSortByEnum
func GetListResolversSortByEnumValues() []ListResolversSortByEnum {
	values := make([]ListResolversSortByEnum, 0)
	for _, v := range mappingListResolversSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResolversSortByEnumStringValues Enumerates the set of values in String for ListResolversSortByEnum
func GetListResolversSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListResolversSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResolversSortByEnum(val string) (ListResolversSortByEnum, bool) {
	enum, ok := mappingListResolversSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResolversScopeEnum Enum with underlying type: string
type ListResolversScopeEnum string

// Set of constants representing the allowable values for ListResolversScopeEnum
const (
	ListResolversScopeGlobal  ListResolversScopeEnum = "GLOBAL"
	ListResolversScopePrivate ListResolversScopeEnum = "PRIVATE"
)

var mappingListResolversScopeEnum = map[string]ListResolversScopeEnum{
	"GLOBAL":  ListResolversScopeGlobal,
	"PRIVATE": ListResolversScopePrivate,
}

var mappingListResolversScopeEnumLowerCase = map[string]ListResolversScopeEnum{
	"global":  ListResolversScopeGlobal,
	"private": ListResolversScopePrivate,
}

// GetListResolversScopeEnumValues Enumerates the set of values for ListResolversScopeEnum
func GetListResolversScopeEnumValues() []ListResolversScopeEnum {
	values := make([]ListResolversScopeEnum, 0)
	for _, v := range mappingListResolversScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetListResolversScopeEnumStringValues Enumerates the set of values in String for ListResolversScopeEnum
func GetListResolversScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingListResolversScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResolversScopeEnum(val string) (ListResolversScopeEnum, bool) {
	enum, ok := mappingListResolversScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
