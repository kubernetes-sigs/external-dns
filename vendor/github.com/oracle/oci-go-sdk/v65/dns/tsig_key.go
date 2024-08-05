// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.cloud.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TsigKey A TSIG key.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type TsigKey struct {

	// TSIG key algorithms are encoded as domain names, but most consist of only one
	// non-empty label, which is not required to be explicitly absolute.
	// Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256,
	// hmac-sha512. For more information on these algorithms, see RFC 4635 (https://tools.ietf.org/html/rfc4635#section-2).
	Algorithm *string `mandatory:"true" json:"algorithm"`

	// A globally unique domain name identifying the key for a given pair of hosts.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment containing the TSIG key.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A base64 string encoding the binary shared secret.
	Secret *string `mandatory:"true" json:"secret"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// The date and time the resource was created, expressed in RFC 3339 timestamp format.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the resource.
	LifecycleState TsigKeyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the resource was last updated, expressed in RFC 3339 timestamp format.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m TsigKey) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TsigKey) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTsigKeyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTsigKeyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TsigKeyLifecycleStateEnum Enum with underlying type: string
type TsigKeyLifecycleStateEnum string

// Set of constants representing the allowable values for TsigKeyLifecycleStateEnum
const (
	TsigKeyLifecycleStateActive   TsigKeyLifecycleStateEnum = "ACTIVE"
	TsigKeyLifecycleStateCreating TsigKeyLifecycleStateEnum = "CREATING"
	TsigKeyLifecycleStateDeleted  TsigKeyLifecycleStateEnum = "DELETED"
	TsigKeyLifecycleStateDeleting TsigKeyLifecycleStateEnum = "DELETING"
	TsigKeyLifecycleStateFailed   TsigKeyLifecycleStateEnum = "FAILED"
	TsigKeyLifecycleStateUpdating TsigKeyLifecycleStateEnum = "UPDATING"
)

var mappingTsigKeyLifecycleStateEnum = map[string]TsigKeyLifecycleStateEnum{
	"ACTIVE":   TsigKeyLifecycleStateActive,
	"CREATING": TsigKeyLifecycleStateCreating,
	"DELETED":  TsigKeyLifecycleStateDeleted,
	"DELETING": TsigKeyLifecycleStateDeleting,
	"FAILED":   TsigKeyLifecycleStateFailed,
	"UPDATING": TsigKeyLifecycleStateUpdating,
}

var mappingTsigKeyLifecycleStateEnumLowerCase = map[string]TsigKeyLifecycleStateEnum{
	"active":   TsigKeyLifecycleStateActive,
	"creating": TsigKeyLifecycleStateCreating,
	"deleted":  TsigKeyLifecycleStateDeleted,
	"deleting": TsigKeyLifecycleStateDeleting,
	"failed":   TsigKeyLifecycleStateFailed,
	"updating": TsigKeyLifecycleStateUpdating,
}

// GetTsigKeyLifecycleStateEnumValues Enumerates the set of values for TsigKeyLifecycleStateEnum
func GetTsigKeyLifecycleStateEnumValues() []TsigKeyLifecycleStateEnum {
	values := make([]TsigKeyLifecycleStateEnum, 0)
	for _, v := range mappingTsigKeyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTsigKeyLifecycleStateEnumStringValues Enumerates the set of values in String for TsigKeyLifecycleStateEnum
func GetTsigKeyLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingTsigKeyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTsigKeyLifecycleStateEnum(val string) (TsigKeyLifecycleStateEnum, bool) {
	enum, ok := mappingTsigKeyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
