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

// GetResolverEndpointRequest wrapper for the GetResolverEndpoint operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/GetResolverEndpoint.go.html to see an example of how to use GetResolverEndpointRequest.
type GetResolverEndpointRequest struct {

	// The OCID of the target resolver.
	ResolverId *string `mandatory:"true" contributesTo:"path" name:"resolverId"`

	// The name of the target resolver endpoint.
	ResolverEndpointName *string `mandatory:"true" contributesTo:"path" name:"resolverEndpointName"`

	// The `If-Modified-Since` header field makes a GET or HEAD request method
	// conditional on the selected representation's modification date being more
	// recent than the date provided in the field-value.  Transfer of the
	// selected representation's data is avoided if that data has not changed.
	IfModifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Modified-Since"`

	// The `If-None-Match` header field makes the request method conditional on
	// the absence of any current representation of the target resource, when
	// the field-value is `*`, or having a selected representation with an
	// entity-tag that does not match any of those listed in the field-value.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"If-None-Match"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope GetResolverEndpointScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetResolverEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetResolverEndpointRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetResolverEndpointRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetResolverEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetResolverEndpointRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetResolverEndpointScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetGetResolverEndpointScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetResolverEndpointResponse wrapper for the GetResolverEndpoint operation
type GetResolverEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ResolverEndpoint instance
	ResolverEndpoint `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response GetResolverEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetResolverEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetResolverEndpointScopeEnum Enum with underlying type: string
type GetResolverEndpointScopeEnum string

// Set of constants representing the allowable values for GetResolverEndpointScopeEnum
const (
	GetResolverEndpointScopeGlobal  GetResolverEndpointScopeEnum = "GLOBAL"
	GetResolverEndpointScopePrivate GetResolverEndpointScopeEnum = "PRIVATE"
)

var mappingGetResolverEndpointScopeEnum = map[string]GetResolverEndpointScopeEnum{
	"GLOBAL":  GetResolverEndpointScopeGlobal,
	"PRIVATE": GetResolverEndpointScopePrivate,
}

var mappingGetResolverEndpointScopeEnumLowerCase = map[string]GetResolverEndpointScopeEnum{
	"global":  GetResolverEndpointScopeGlobal,
	"private": GetResolverEndpointScopePrivate,
}

// GetGetResolverEndpointScopeEnumValues Enumerates the set of values for GetResolverEndpointScopeEnum
func GetGetResolverEndpointScopeEnumValues() []GetResolverEndpointScopeEnum {
	values := make([]GetResolverEndpointScopeEnum, 0)
	for _, v := range mappingGetResolverEndpointScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetResolverEndpointScopeEnumStringValues Enumerates the set of values in String for GetResolverEndpointScopeEnum
func GetGetResolverEndpointScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingGetResolverEndpointScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetResolverEndpointScopeEnum(val string) (GetResolverEndpointScopeEnum, bool) {
	enum, ok := mappingGetResolverEndpointScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
