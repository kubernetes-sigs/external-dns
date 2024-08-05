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

// DeleteZoneRequest wrapper for the DeleteZone operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteZone.go.html to see an example of how to use DeleteZoneRequest.
type DeleteZoneRequest struct {

	// The name or OCID of the target zone.
	ZoneNameOrId *string `mandatory:"true" contributesTo:"path" name:"zoneNameOrId"`

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
	Scope DeleteZoneScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The OCID of the view the zone is associated with. Required when accessing a private zone by name.
	ViewId *string `mandatory:"false" contributesTo:"query" name:"viewId"`

	// The OCID of the compartment the zone belongs to.
	// This parameter is deprecated and should be omitted.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteZoneRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteZoneRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteZoneRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteZoneRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeleteZoneRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeleteZoneScopeEnum(string(request.Scope)); !ok && request.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", request.Scope, strings.Join(GetDeleteZoneScopeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeleteZoneResponse wrapper for the DeleteZone operation
type DeleteZoneResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response DeleteZoneResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteZoneResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DeleteZoneScopeEnum Enum with underlying type: string
type DeleteZoneScopeEnum string

// Set of constants representing the allowable values for DeleteZoneScopeEnum
const (
	DeleteZoneScopeGlobal  DeleteZoneScopeEnum = "GLOBAL"
	DeleteZoneScopePrivate DeleteZoneScopeEnum = "PRIVATE"
)

var mappingDeleteZoneScopeEnum = map[string]DeleteZoneScopeEnum{
	"GLOBAL":  DeleteZoneScopeGlobal,
	"PRIVATE": DeleteZoneScopePrivate,
}

var mappingDeleteZoneScopeEnumLowerCase = map[string]DeleteZoneScopeEnum{
	"global":  DeleteZoneScopeGlobal,
	"private": DeleteZoneScopePrivate,
}

// GetDeleteZoneScopeEnumValues Enumerates the set of values for DeleteZoneScopeEnum
func GetDeleteZoneScopeEnumValues() []DeleteZoneScopeEnum {
	values := make([]DeleteZoneScopeEnum, 0)
	for _, v := range mappingDeleteZoneScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetDeleteZoneScopeEnumStringValues Enumerates the set of values in String for DeleteZoneScopeEnum
func GetDeleteZoneScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingDeleteZoneScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeleteZoneScopeEnum(val string) (DeleteZoneScopeEnum, bool) {
	enum, ok := mappingDeleteZoneScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
