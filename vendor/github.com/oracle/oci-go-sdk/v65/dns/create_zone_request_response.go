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

// CreateZoneRequest wrapper for the CreateZone operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateZone.go.html to see an example of how to use CreateZoneRequest.
type CreateZoneRequest struct {

	// Details for creating a new zone.
	CreateZoneDetails CreateZoneBaseDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the compartment the zone belongs to.
	// This parameter is deprecated and should be omitted.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope CreateZoneScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the resource is associated with.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateZoneRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateZoneRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateZoneRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateZoneRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateZoneRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateZoneScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetCreateZoneScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateZoneResponse wrapper for the CreateZone operation
type CreateZoneResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Zone instance
	Zone `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`

	// The full URI of the resource related to the request.
	Location *string `presentIn:"header" name:"location"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response CreateZoneResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateZoneResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateZoneScopeEnum Enum with underlying type: string
type CreateZoneScopeEnum string

// Set of constants representing the allowable values for CreateZoneScopeEnum
const (
	CreateZoneScopeGlobal  CreateZoneScopeEnum = "GLOBAL"
	CreateZoneScopePrivate CreateZoneScopeEnum = "PRIVATE"
)

var mappingCreateZoneScopeEnum = map[string]CreateZoneScopeEnum{
	"GLOBAL":  CreateZoneScopeGlobal,
	"PRIVATE": CreateZoneScopePrivate,
}

var mappingCreateZoneScopeEnumLowerCase = map[string]CreateZoneScopeEnum{
	"global":  CreateZoneScopeGlobal,
	"private": CreateZoneScopePrivate,
}

// GetCreateZoneScopeEnumValues Enumerates the set of values for CreateZoneScopeEnum
func GetCreateZoneScopeEnumValues() []CreateZoneScopeEnum {
	values := make([]CreateZoneScopeEnum, 0)
	for _, v := range mappingCreateZoneScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateZoneScopeEnumStringValues Enumerates the set of values in String for CreateZoneScopeEnum
func GetCreateZoneScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingCreateZoneScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateZoneScopeEnum(val string) (CreateZoneScopeEnum, bool) {
	enum, ok := mappingCreateZoneScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
