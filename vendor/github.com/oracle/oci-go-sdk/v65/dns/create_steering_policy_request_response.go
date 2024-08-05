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

// CreateSteeringPolicyRequest wrapper for the CreateSteeringPolicy operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateSteeringPolicy.go.html to see an example of how to use CreateSteeringPolicyRequest.
type CreateSteeringPolicyRequest struct {

	// Details for creating a new steering policy.
	CreateSteeringPolicyDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case
	// of a timeout or server error without risk of executing that same action
	// again. Retry tokens expire after 24 hours, but can be invalidated before
	// then due to conflicting operations (for example, if a resource has been
	// deleted and purged from the system, then a retry of the original creation
	// request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope CreateSteeringPolicyScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateSteeringPolicyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateSteeringPolicyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateSteeringPolicyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateSteeringPolicyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateSteeringPolicyRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateSteeringPolicyScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetCreateSteeringPolicyScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateSteeringPolicyResponse wrapper for the CreateSteeringPolicy operation
type CreateSteeringPolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SteeringPolicy instance
	SteeringPolicy `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	ETag *string `presentIn:"header" name:"etag"`

	// The full URI of the resource related to the request.
	Location *string `presentIn:"header" name:"location"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateSteeringPolicyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateSteeringPolicyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateSteeringPolicyScopeEnum Enum with underlying type: string
type CreateSteeringPolicyScopeEnum string

// Set of constants representing the allowable values for CreateSteeringPolicyScopeEnum
const (
	CreateSteeringPolicyScopeGlobal  CreateSteeringPolicyScopeEnum = "GLOBAL"
	CreateSteeringPolicyScopePrivate CreateSteeringPolicyScopeEnum = "PRIVATE"
)

var mappingCreateSteeringPolicyScopeEnum = map[string]CreateSteeringPolicyScopeEnum{
	"GLOBAL":  CreateSteeringPolicyScopeGlobal,
	"PRIVATE": CreateSteeringPolicyScopePrivate,
}

var mappingCreateSteeringPolicyScopeEnumLowerCase = map[string]CreateSteeringPolicyScopeEnum{
	"global":  CreateSteeringPolicyScopeGlobal,
	"private": CreateSteeringPolicyScopePrivate,
}

// GetCreateSteeringPolicyScopeEnumValues Enumerates the set of values for CreateSteeringPolicyScopeEnum
func GetCreateSteeringPolicyScopeEnumValues() []CreateSteeringPolicyScopeEnum {
	values := make([]CreateSteeringPolicyScopeEnum, 0)
	for _, v := range mappingCreateSteeringPolicyScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateSteeringPolicyScopeEnumStringValues Enumerates the set of values in String for CreateSteeringPolicyScopeEnum
func GetCreateSteeringPolicyScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingCreateSteeringPolicyScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateSteeringPolicyScopeEnum(val string) (CreateSteeringPolicyScopeEnum, bool) {
	enum, ok := mappingCreateSteeringPolicyScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
