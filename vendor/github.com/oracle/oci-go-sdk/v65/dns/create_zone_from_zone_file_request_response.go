// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// CreateZoneFromZoneFileRequest wrapper for the CreateZoneFromZoneFile operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateZoneFromZoneFile.go.html to see an example of how to use CreateZoneFromZoneFileRequest.
type CreateZoneFromZoneFileRequest struct {

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The zone file contents.
	CreateZoneFromZoneFileDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope CreateZoneFromZoneFileScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the resource is associated with.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateZoneFromZoneFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateZoneFromZoneFileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateZoneFromZoneFileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.CreateZoneFromZoneFileDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateZoneFromZoneFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateZoneFromZoneFileRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateZoneFromZoneFileScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetCreateZoneFromZoneFileScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateZoneFromZoneFileResponse wrapper for the CreateZoneFromZoneFile operation
type CreateZoneFromZoneFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Zone instance
	Zone `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	Etag *string `presentIn:"header" name:"etag"`

	// The full URI of the resource related to the request.
	Location *string `presentIn:"header" name:"location"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response CreateZoneFromZoneFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateZoneFromZoneFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateZoneFromZoneFileScopeEnum Enum with underlying type: string
type CreateZoneFromZoneFileScopeEnum string

// Set of constants representing the allowable values for CreateZoneFromZoneFileScopeEnum
const (
	CreateZoneFromZoneFileScopeGlobal  CreateZoneFromZoneFileScopeEnum = "GLOBAL"
	CreateZoneFromZoneFileScopePrivate CreateZoneFromZoneFileScopeEnum = "PRIVATE"
)

var mappingCreateZoneFromZoneFileScopeEnum = map[string]CreateZoneFromZoneFileScopeEnum{
	"GLOBAL":  CreateZoneFromZoneFileScopeGlobal,
	"PRIVATE": CreateZoneFromZoneFileScopePrivate,
}

var mappingCreateZoneFromZoneFileScopeEnumLowerCase = map[string]CreateZoneFromZoneFileScopeEnum{
	"global":  CreateZoneFromZoneFileScopeGlobal,
	"private": CreateZoneFromZoneFileScopePrivate,
}

// GetCreateZoneFromZoneFileScopeEnumValues Enumerates the set of values for CreateZoneFromZoneFileScopeEnum
func GetCreateZoneFromZoneFileScopeEnumValues() []CreateZoneFromZoneFileScopeEnum {
	values := make([]CreateZoneFromZoneFileScopeEnum, 0)
	for _, v := range mappingCreateZoneFromZoneFileScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateZoneFromZoneFileScopeEnumStringValues Enumerates the set of values in String for CreateZoneFromZoneFileScopeEnum
func GetCreateZoneFromZoneFileScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingCreateZoneFromZoneFileScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateZoneFromZoneFileScopeEnum(val string) (CreateZoneFromZoneFileScopeEnum, bool) {
	enum, ok := mappingCreateZoneFromZoneFileScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
