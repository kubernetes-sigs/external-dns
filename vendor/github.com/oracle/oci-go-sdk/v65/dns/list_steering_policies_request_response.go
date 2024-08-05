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

// ListSteeringPoliciesRequest wrapper for the ListSteeringPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListSteeringPolicies.go.html to see an example of how to use ListSteeringPoliciesRequest.
type ListSteeringPoliciesRequest struct {

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

	// The displayName of a resource.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The partial displayName of a resource. Will match any resource whose name
	// (case-insensitive) contains the provided value.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Search by health check monitor OCID.
	// Will match any resource whose health check monitor ID matches the provided value.
	HealthCheckMonitorId *string `mandatory:"false" contributesTo:"query" name:"healthCheckMonitorId"`

	// An RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) timestamp that states
	// all returned resources were created on or after the indicated time.
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// An RFC 3339 (https://www.ietf.org/rfc/rfc3339.txt) timestamp that states
	// all returned resources were created before the indicated time.
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// Search by steering template type.
	// Will match any resource whose template type matches the provided value.
	Template *string `mandatory:"false" contributesTo:"query" name:"template"`

	// The state of a resource.
	LifecycleState SteeringPolicySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field by which to sort steering policies. If unspecified, defaults to `timeCreated`.
	SortBy ListSteeringPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The order to sort the resources.
	SortOrder ListSteeringPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListSteeringPoliciesScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSteeringPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSteeringPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSteeringPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSteeringPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSteeringPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSteeringPolicySummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetSteeringPolicySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSteeringPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSteeringPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSteeringPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSteeringPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSteeringPoliciesScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetListSteeringPoliciesScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSteeringPoliciesResponse wrapper for the ListSteeringPolicies operation
type ListSteeringPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SteeringPolicySummary instances
	Items []SteeringPolicySummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of items that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSteeringPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSteeringPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSteeringPoliciesSortByEnum Enum with underlying type: string
type ListSteeringPoliciesSortByEnum string

// Set of constants representing the allowable values for ListSteeringPoliciesSortByEnum
const (
	ListSteeringPoliciesSortByDisplayname ListSteeringPoliciesSortByEnum = "displayName"
	ListSteeringPoliciesSortByTimecreated ListSteeringPoliciesSortByEnum = "timeCreated"
	ListSteeringPoliciesSortByTemplate    ListSteeringPoliciesSortByEnum = "template"
)

var mappingListSteeringPoliciesSortByEnum = map[string]ListSteeringPoliciesSortByEnum{
	"displayName": ListSteeringPoliciesSortByDisplayname,
	"timeCreated": ListSteeringPoliciesSortByTimecreated,
	"template":    ListSteeringPoliciesSortByTemplate,
}

var mappingListSteeringPoliciesSortByEnumLowerCase = map[string]ListSteeringPoliciesSortByEnum{
	"displayname": ListSteeringPoliciesSortByDisplayname,
	"timecreated": ListSteeringPoliciesSortByTimecreated,
	"template":    ListSteeringPoliciesSortByTemplate,
}

// GetListSteeringPoliciesSortByEnumValues Enumerates the set of values for ListSteeringPoliciesSortByEnum
func GetListSteeringPoliciesSortByEnumValues() []ListSteeringPoliciesSortByEnum {
	values := make([]ListSteeringPoliciesSortByEnum, 0)
	for _, v := range mappingListSteeringPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSteeringPoliciesSortByEnumStringValues Enumerates the set of values in String for ListSteeringPoliciesSortByEnum
func GetListSteeringPoliciesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
		"template",
	}
}

// GetMappingListSteeringPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSteeringPoliciesSortByEnum(val string) (ListSteeringPoliciesSortByEnum, bool) {
	enum, ok := mappingListSteeringPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSteeringPoliciesSortOrderEnum Enum with underlying type: string
type ListSteeringPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListSteeringPoliciesSortOrderEnum
const (
	ListSteeringPoliciesSortOrderAsc  ListSteeringPoliciesSortOrderEnum = "ASC"
	ListSteeringPoliciesSortOrderDesc ListSteeringPoliciesSortOrderEnum = "DESC"
)

var mappingListSteeringPoliciesSortOrderEnum = map[string]ListSteeringPoliciesSortOrderEnum{
	"ASC":  ListSteeringPoliciesSortOrderAsc,
	"DESC": ListSteeringPoliciesSortOrderDesc,
}

var mappingListSteeringPoliciesSortOrderEnumLowerCase = map[string]ListSteeringPoliciesSortOrderEnum{
	"asc":  ListSteeringPoliciesSortOrderAsc,
	"desc": ListSteeringPoliciesSortOrderDesc,
}

// GetListSteeringPoliciesSortOrderEnumValues Enumerates the set of values for ListSteeringPoliciesSortOrderEnum
func GetListSteeringPoliciesSortOrderEnumValues() []ListSteeringPoliciesSortOrderEnum {
	values := make([]ListSteeringPoliciesSortOrderEnum, 0)
	for _, v := range mappingListSteeringPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSteeringPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListSteeringPoliciesSortOrderEnum
func GetListSteeringPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSteeringPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSteeringPoliciesSortOrderEnum(val string) (ListSteeringPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListSteeringPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSteeringPoliciesScopeEnum Enum with underlying type: string
type ListSteeringPoliciesScopeEnum string

// Set of constants representing the allowable values for ListSteeringPoliciesScopeEnum
const (
	ListSteeringPoliciesScopeGlobal  ListSteeringPoliciesScopeEnum = "GLOBAL"
	ListSteeringPoliciesScopePrivate ListSteeringPoliciesScopeEnum = "PRIVATE"
)

var mappingListSteeringPoliciesScopeEnum = map[string]ListSteeringPoliciesScopeEnum{
	"GLOBAL":  ListSteeringPoliciesScopeGlobal,
	"PRIVATE": ListSteeringPoliciesScopePrivate,
}

var mappingListSteeringPoliciesScopeEnumLowerCase = map[string]ListSteeringPoliciesScopeEnum{
	"global":  ListSteeringPoliciesScopeGlobal,
	"private": ListSteeringPoliciesScopePrivate,
}

// GetListSteeringPoliciesScopeEnumValues Enumerates the set of values for ListSteeringPoliciesScopeEnum
func GetListSteeringPoliciesScopeEnumValues() []ListSteeringPoliciesScopeEnum {
	values := make([]ListSteeringPoliciesScopeEnum, 0)
	for _, v := range mappingListSteeringPoliciesScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSteeringPoliciesScopeEnumStringValues Enumerates the set of values in String for ListSteeringPoliciesScopeEnum
func GetListSteeringPoliciesScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingListSteeringPoliciesScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSteeringPoliciesScopeEnum(val string) (ListSteeringPoliciesScopeEnum, bool) {
	enum, ok := mappingListSteeringPoliciesScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
