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

// CreateTsigKeyRequest wrapper for the CreateTsigKey operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateTsigKey.go.html to see an example of how to use CreateTsigKeyRequest.
type CreateTsigKeyRequest struct {

	// Details for creating a new TSIG key.
	CreateTsigKeyDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope CreateTsigKeyScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateTsigKeyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateTsigKeyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateTsigKeyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateTsigKeyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateTsigKeyRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateTsigKeyScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetCreateTsigKeyScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateTsigKeyResponse wrapper for the CreateTsigKey operation
type CreateTsigKeyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TsigKey instance
	TsigKey `presentIn:"body"`

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

func (response CreateTsigKeyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateTsigKeyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateTsigKeyScopeEnum Enum with underlying type: string
type CreateTsigKeyScopeEnum string

// Set of constants representing the allowable values for CreateTsigKeyScopeEnum
const (
	CreateTsigKeyScopeGlobal  CreateTsigKeyScopeEnum = "GLOBAL"
	CreateTsigKeyScopePrivate CreateTsigKeyScopeEnum = "PRIVATE"
)

var mappingCreateTsigKeyScopeEnum = map[string]CreateTsigKeyScopeEnum{
	"GLOBAL":  CreateTsigKeyScopeGlobal,
	"PRIVATE": CreateTsigKeyScopePrivate,
}

var mappingCreateTsigKeyScopeEnumLowerCase = map[string]CreateTsigKeyScopeEnum{
	"global":  CreateTsigKeyScopeGlobal,
	"private": CreateTsigKeyScopePrivate,
}

// GetCreateTsigKeyScopeEnumValues Enumerates the set of values for CreateTsigKeyScopeEnum
func GetCreateTsigKeyScopeEnumValues() []CreateTsigKeyScopeEnum {
	values := make([]CreateTsigKeyScopeEnum, 0)
	for _, v := range mappingCreateTsigKeyScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTsigKeyScopeEnumStringValues Enumerates the set of values in String for CreateTsigKeyScopeEnum
func GetCreateTsigKeyScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingCreateTsigKeyScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTsigKeyScopeEnum(val string) (CreateTsigKeyScopeEnum, bool) {
	enum, ok := mappingCreateTsigKeyScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
