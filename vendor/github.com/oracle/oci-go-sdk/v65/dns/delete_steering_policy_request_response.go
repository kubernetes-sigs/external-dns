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

// DeleteSteeringPolicyRequest wrapper for the DeleteSteeringPolicy operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteSteeringPolicy.go.html to see an example of how to use DeleteSteeringPolicyRequest.
type DeleteSteeringPolicyRequest struct {

	// The OCID of the target steering policy.
	SteeringPolicyId *string `mandatory:"true" contributesTo:"path" name:"steeringPolicyId"`

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
	Scope DeleteSteeringPolicyScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteSteeringPolicyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteSteeringPolicyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteSteeringPolicyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteSteeringPolicyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeleteSteeringPolicyRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeleteSteeringPolicyScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetDeleteSteeringPolicyScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeleteSteeringPolicyResponse wrapper for the DeleteSteeringPolicy operation
type DeleteSteeringPolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteSteeringPolicyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteSteeringPolicyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DeleteSteeringPolicyScopeEnum Enum with underlying type: string
type DeleteSteeringPolicyScopeEnum string

// Set of constants representing the allowable values for DeleteSteeringPolicyScopeEnum
const (
	DeleteSteeringPolicyScopeGlobal  DeleteSteeringPolicyScopeEnum = "GLOBAL"
	DeleteSteeringPolicyScopePrivate DeleteSteeringPolicyScopeEnum = "PRIVATE"
)

var mappingDeleteSteeringPolicyScopeEnum = map[string]DeleteSteeringPolicyScopeEnum{
	"GLOBAL":  DeleteSteeringPolicyScopeGlobal,
	"PRIVATE": DeleteSteeringPolicyScopePrivate,
}

var mappingDeleteSteeringPolicyScopeEnumLowerCase = map[string]DeleteSteeringPolicyScopeEnum{
	"global":  DeleteSteeringPolicyScopeGlobal,
	"private": DeleteSteeringPolicyScopePrivate,
}

// GetDeleteSteeringPolicyScopeEnumValues Enumerates the set of values for DeleteSteeringPolicyScopeEnum
func GetDeleteSteeringPolicyScopeEnumValues() []DeleteSteeringPolicyScopeEnum {
	values := make([]DeleteSteeringPolicyScopeEnum, 0)
	for _, v := range mappingDeleteSteeringPolicyScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeleteSteeringPolicyScopeEnumStringValues Enumerates the set of values in String for DeleteSteeringPolicyScopeEnum
func GetDeleteSteeringPolicyScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingDeleteSteeringPolicyScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeleteSteeringPolicyScopeEnum(val string) (DeleteSteeringPolicyScopeEnum, bool) {
	enum, ok := mappingDeleteSteeringPolicyScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
