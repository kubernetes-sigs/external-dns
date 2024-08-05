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

// UpdateSteeringPolicyRequest wrapper for the UpdateSteeringPolicy operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/UpdateSteeringPolicy.go.html to see an example of how to use UpdateSteeringPolicyRequest.
type UpdateSteeringPolicyRequest struct {

	// The OCID of the target steering policy.
	SteeringPolicyId *string `mandatory:"true" contributesTo:"path" name:"steeringPolicyId"`

	// New data for the steering policy.
	UpdateSteeringPolicyDetails `contributesTo:"body"`

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
	Scope UpdateSteeringPolicyScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateSteeringPolicyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateSteeringPolicyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateSteeringPolicyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateSteeringPolicyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateSteeringPolicyRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateSteeringPolicyScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetUpdateSteeringPolicyScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateSteeringPolicyResponse wrapper for the UpdateSteeringPolicy operation
type UpdateSteeringPolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SteeringPolicy instance
	SteeringPolicy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response UpdateSteeringPolicyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateSteeringPolicyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UpdateSteeringPolicyScopeEnum Enum with underlying type: string
type UpdateSteeringPolicyScopeEnum string

// Set of constants representing the allowable values for UpdateSteeringPolicyScopeEnum
const (
	UpdateSteeringPolicyScopeGlobal  UpdateSteeringPolicyScopeEnum = "GLOBAL"
	UpdateSteeringPolicyScopePrivate UpdateSteeringPolicyScopeEnum = "PRIVATE"
)

var mappingUpdateSteeringPolicyScopeEnum = map[string]UpdateSteeringPolicyScopeEnum{
	"GLOBAL":  UpdateSteeringPolicyScopeGlobal,
	"PRIVATE": UpdateSteeringPolicyScopePrivate,
}

var mappingUpdateSteeringPolicyScopeEnumLowerCase = map[string]UpdateSteeringPolicyScopeEnum{
	"global":  UpdateSteeringPolicyScopeGlobal,
	"private": UpdateSteeringPolicyScopePrivate,
}

// GetUpdateSteeringPolicyScopeEnumValues Enumerates the set of values for UpdateSteeringPolicyScopeEnum
func GetUpdateSteeringPolicyScopeEnumValues() []UpdateSteeringPolicyScopeEnum {
	values := make([]UpdateSteeringPolicyScopeEnum, 0)
	for _, v := range mappingUpdateSteeringPolicyScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSteeringPolicyScopeEnumStringValues Enumerates the set of values in String for UpdateSteeringPolicyScopeEnum
func GetUpdateSteeringPolicyScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingUpdateSteeringPolicyScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSteeringPolicyScopeEnum(val string) (UpdateSteeringPolicyScopeEnum, bool) {
	enum, ok := mappingUpdateSteeringPolicyScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
