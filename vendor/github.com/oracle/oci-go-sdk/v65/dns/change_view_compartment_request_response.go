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

// ChangeViewCompartmentRequest wrapper for the ChangeViewCompartment operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ChangeViewCompartment.go.html to see an example of how to use ChangeViewCompartmentRequest.
type ChangeViewCompartmentRequest struct {

	// The OCID of the target view.
	ViewId *string `mandatory:"true" contributesTo:"path" name:"viewId"`

	// Details for moving a view into a different compartment.
	ChangeViewCompartmentDetails `contributesTo:"body"`

	// The `If-Match` header field makes the request method conditional on the
	// existence of at least one current representation of the target resource,
	// when the field-value is `*`, or having a current representation of the
	// target resource that has an entity-tag matching a member of the list of
	// entity-tags provided in the field-value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"If-Match"`

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
	Scope ChangeViewCompartmentScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangeViewCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangeViewCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ChangeViewCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangeViewCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ChangeViewCompartmentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingChangeViewCompartmentScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetChangeViewCompartmentScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeViewCompartmentResponse wrapper for the ChangeViewCompartment operation
type ChangeViewCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response ChangeViewCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangeViewCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ChangeViewCompartmentScopeEnum Enum with underlying type: string
type ChangeViewCompartmentScopeEnum string

// Set of constants representing the allowable values for ChangeViewCompartmentScopeEnum
const (
	ChangeViewCompartmentScopeGlobal  ChangeViewCompartmentScopeEnum = "GLOBAL"
	ChangeViewCompartmentScopePrivate ChangeViewCompartmentScopeEnum = "PRIVATE"
)

var mappingChangeViewCompartmentScopeEnum = map[string]ChangeViewCompartmentScopeEnum{
	"GLOBAL":  ChangeViewCompartmentScopeGlobal,
	"PRIVATE": ChangeViewCompartmentScopePrivate,
}

var mappingChangeViewCompartmentScopeEnumLowerCase = map[string]ChangeViewCompartmentScopeEnum{
	"global":  ChangeViewCompartmentScopeGlobal,
	"private": ChangeViewCompartmentScopePrivate,
}

// GetChangeViewCompartmentScopeEnumValues Enumerates the set of values for ChangeViewCompartmentScopeEnum
func GetChangeViewCompartmentScopeEnumValues() []ChangeViewCompartmentScopeEnum {
	values := make([]ChangeViewCompartmentScopeEnum, 0)
	for _, v := range mappingChangeViewCompartmentScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeViewCompartmentScopeEnumStringValues Enumerates the set of values in String for ChangeViewCompartmentScopeEnum
func GetChangeViewCompartmentScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingChangeViewCompartmentScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeViewCompartmentScopeEnum(val string) (ChangeViewCompartmentScopeEnum, bool) {
	enum, ok := mappingChangeViewCompartmentScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
