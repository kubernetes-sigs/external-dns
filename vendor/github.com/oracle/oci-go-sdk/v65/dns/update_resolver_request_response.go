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

// UpdateResolverRequest wrapper for the UpdateResolver operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateResolver.go.html to see an example of how to use UpdateResolverRequest.
type UpdateResolverRequest struct {

	// The OCID of the target resolver.
	ResolverId *string `mandatory:"true" contributesTo:"path" name:"resolverId"`

	// New data for the resolver.
	UpdateResolverDetails `contributesTo:"body"`

	// The `If-Match` header field makes the request method conditional on the
	// existence of at least one current representation of the target resource,
	// when the field-value is `*`, or having a current representation of the
	// target resource that has an entity-tag matching a member of the list of
	// entity-tags provided in the field-value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"If-Match"`

	// The `If-Unmodified-Since` header field makes the request method
	// conditional on the selected representation's last modification date being
	// earlier than or equal to the date provided in the field-value.  This
	// field accomplishes the same purpose as If-Match for cases where the user
	// agent does not have an entity-tag for the representation.
	IfUnmodifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Unmodified-Since"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope UpdateResolverScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateResolverRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateResolverRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateResolverRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateResolverRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateResolverRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateResolverScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetUpdateResolverScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateResolverResponse wrapper for the UpdateResolver operation
type UpdateResolverResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Resolver instance
	Resolver `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response UpdateResolverResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateResolverResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UpdateResolverScopeEnum Enum with underlying type: string
type UpdateResolverScopeEnum string

// Set of constants representing the allowable values for UpdateResolverScopeEnum
const (
	UpdateResolverScopeGlobal  UpdateResolverScopeEnum = "GLOBAL"
	UpdateResolverScopePrivate UpdateResolverScopeEnum = "PRIVATE"
)

var mappingUpdateResolverScopeEnum = map[string]UpdateResolverScopeEnum{
	"GLOBAL":  UpdateResolverScopeGlobal,
	"PRIVATE": UpdateResolverScopePrivate,
}

var mappingUpdateResolverScopeEnumLowerCase = map[string]UpdateResolverScopeEnum{
	"global":  UpdateResolverScopeGlobal,
	"private": UpdateResolverScopePrivate,
}

// GetUpdateResolverScopeEnumValues Enumerates the set of values for UpdateResolverScopeEnum
func GetUpdateResolverScopeEnumValues() []UpdateResolverScopeEnum {
	values := make([]UpdateResolverScopeEnum, 0)
	for _, v := range mappingUpdateResolverScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateResolverScopeEnumStringValues Enumerates the set of values in String for UpdateResolverScopeEnum
func GetUpdateResolverScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingUpdateResolverScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateResolverScopeEnum(val string) (UpdateResolverScopeEnum, bool) {
	enum, ok := mappingUpdateResolverScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
