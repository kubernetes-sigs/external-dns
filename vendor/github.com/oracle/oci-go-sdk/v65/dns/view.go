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

// View An OCI DNS view.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type View struct {

	// The OCID of the owning compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the view.
	DisplayName *string `mandatory:"true" json:"displayName"`

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

	// The OCID of the view.
	Id *string `mandatory:"true" json:"id"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ"
	// format with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the resource.
	LifecycleState ViewLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `mandatory:"true" json:"isProtected"`
}

func (m View) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m View) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingViewLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetViewLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ViewLifecycleStateEnum Enum with underlying type: string
type ViewLifecycleStateEnum string

// Set of constants representing the allowable values for ViewLifecycleStateEnum
const (
	ViewLifecycleStateActive   ViewLifecycleStateEnum = "ACTIVE"
	ViewLifecycleStateDeleted  ViewLifecycleStateEnum = "DELETED"
	ViewLifecycleStateDeleting ViewLifecycleStateEnum = "DELETING"
	ViewLifecycleStateUpdating ViewLifecycleStateEnum = "UPDATING"
)

var mappingViewLifecycleStateEnum = map[string]ViewLifecycleStateEnum{
	"ACTIVE":   ViewLifecycleStateActive,
	"DELETED":  ViewLifecycleStateDeleted,
	"DELETING": ViewLifecycleStateDeleting,
	"UPDATING": ViewLifecycleStateUpdating,
}

var mappingViewLifecycleStateEnumLowerCase = map[string]ViewLifecycleStateEnum{
	"active":   ViewLifecycleStateActive,
	"deleted":  ViewLifecycleStateDeleted,
	"deleting": ViewLifecycleStateDeleting,
	"updating": ViewLifecycleStateUpdating,
}

// GetViewLifecycleStateEnumValues Enumerates the set of values for ViewLifecycleStateEnum
func GetViewLifecycleStateEnumValues() []ViewLifecycleStateEnum {
	values := make([]ViewLifecycleStateEnum, 0)
	for _, v := range mappingViewLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetViewLifecycleStateEnumStringValues Enumerates the set of values in String for ViewLifecycleStateEnum
func GetViewLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"DELETING",
		"UPDATING",
	}
}

// GetMappingViewLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingViewLifecycleStateEnum(val string) (ViewLifecycleStateEnum, bool) {
	enum, ok := mappingViewLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
