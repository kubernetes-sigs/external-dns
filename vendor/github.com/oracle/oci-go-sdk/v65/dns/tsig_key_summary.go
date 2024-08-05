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

// TsigKeySummary A TSIG key.
type TsigKeySummary struct {

	// TSIG key algorithms are encoded as domain names, but most consist of only one
	// non-empty label, which is not required to be explicitly absolute.
	// Applicable algorithms include: hmac-sha1, hmac-sha224, hmac-sha256,
	// hmac-sha512. For more information on these algorithms, see RFC 4635 (https://tools.ietf.org/html/rfc4635#section-2).
	Algorithm *string `mandatory:"true" json:"algorithm"`

	// A globally unique domain name identifying the key for a given pair of hosts.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment containing the TSIG key.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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
	LifecycleState TsigKeySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m TsigKeySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TsigKeySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTsigKeySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTsigKeySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TsigKeySummaryLifecycleStateEnum Enum with underlying type: string
type TsigKeySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for TsigKeySummaryLifecycleStateEnum
const (
	TsigKeySummaryLifecycleStateActive   TsigKeySummaryLifecycleStateEnum = "ACTIVE"
	TsigKeySummaryLifecycleStateCreating TsigKeySummaryLifecycleStateEnum = "CREATING"
	TsigKeySummaryLifecycleStateDeleted  TsigKeySummaryLifecycleStateEnum = "DELETED"
	TsigKeySummaryLifecycleStateDeleting TsigKeySummaryLifecycleStateEnum = "DELETING"
	TsigKeySummaryLifecycleStateFailed   TsigKeySummaryLifecycleStateEnum = "FAILED"
	TsigKeySummaryLifecycleStateUpdating TsigKeySummaryLifecycleStateEnum = "UPDATING"
)

var mappingTsigKeySummaryLifecycleStateEnum = map[string]TsigKeySummaryLifecycleStateEnum{
	"ACTIVE":   TsigKeySummaryLifecycleStateActive,
	"CREATING": TsigKeySummaryLifecycleStateCreating,
	"DELETED":  TsigKeySummaryLifecycleStateDeleted,
	"DELETING": TsigKeySummaryLifecycleStateDeleting,
	"FAILED":   TsigKeySummaryLifecycleStateFailed,
	"UPDATING": TsigKeySummaryLifecycleStateUpdating,
}

var mappingTsigKeySummaryLifecycleStateEnumLowerCase = map[string]TsigKeySummaryLifecycleStateEnum{
	"active":   TsigKeySummaryLifecycleStateActive,
	"creating": TsigKeySummaryLifecycleStateCreating,
	"deleted":  TsigKeySummaryLifecycleStateDeleted,
	"deleting": TsigKeySummaryLifecycleStateDeleting,
	"failed":   TsigKeySummaryLifecycleStateFailed,
	"updating": TsigKeySummaryLifecycleStateUpdating,
}

// GetTsigKeySummaryLifecycleStateEnumValues Enumerates the set of values for TsigKeySummaryLifecycleStateEnum
func GetTsigKeySummaryLifecycleStateEnumValues() []TsigKeySummaryLifecycleStateEnum {
	values := make([]TsigKeySummaryLifecycleStateEnum, 0)
	for _, v := range mappingTsigKeySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTsigKeySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for TsigKeySummaryLifecycleStateEnum
func GetTsigKeySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingTsigKeySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTsigKeySummaryLifecycleStateEnum(val string) (TsigKeySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingTsigKeySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
