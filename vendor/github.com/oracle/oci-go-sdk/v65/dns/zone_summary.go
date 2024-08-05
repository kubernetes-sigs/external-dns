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

// ZoneSummary A DNS zone.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ZoneSummary struct {

	// The name of the zone.
	Name *string `mandatory:"true" json:"name"`

	// The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones.
	ZoneType ZoneSummaryZoneTypeEnum `mandatory:"true" json:"zoneType"`

	// The OCID of the compartment containing the zone.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The scope of the zone.
	Scope ScopeEnum `mandatory:"true" json:"scope"`

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

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// The OCID of the zone.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Version is the never-repeating, totally-orderable, version of the
	// zone, from which the serial field of the zone's SOA record is
	// derived.
	Version *string `mandatory:"true" json:"version"`

	// The current serial of the zone. As seen in the zone's SOA record.
	Serial *int64 `mandatory:"true" json:"serial"`

	// The current state of the zone resource.
	LifecycleState ZoneSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `mandatory:"true" json:"isProtected"`

	// The OCID of the private view containing the zone. This value will
	// be null for zones in the global DNS, which are publicly resolvable and
	// not part of a private view.
	ViewId *string `mandatory:"false" json:"viewId"`
}

func (m ZoneSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZoneSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingZoneSummaryZoneTypeEnum(string(m.ZoneType)); !ok && m.ZoneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ZoneType: %s. Supported values are: %s.", m.ZoneType, strings.Join(GetZoneSummaryZoneTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingZoneSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZoneSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ZoneSummaryZoneTypeEnum Enum with underlying type: string
type ZoneSummaryZoneTypeEnum string

// Set of constants representing the allowable values for ZoneSummaryZoneTypeEnum
const (
	ZoneSummaryZoneTypePrimary   ZoneSummaryZoneTypeEnum = "PRIMARY"
	ZoneSummaryZoneTypeSecondary ZoneSummaryZoneTypeEnum = "SECONDARY"
)

var mappingZoneSummaryZoneTypeEnum = map[string]ZoneSummaryZoneTypeEnum{
	"PRIMARY":   ZoneSummaryZoneTypePrimary,
	"SECONDARY": ZoneSummaryZoneTypeSecondary,
}

var mappingZoneSummaryZoneTypeEnumLowerCase = map[string]ZoneSummaryZoneTypeEnum{
	"primary":   ZoneSummaryZoneTypePrimary,
	"secondary": ZoneSummaryZoneTypeSecondary,
}

// GetZoneSummaryZoneTypeEnumValues Enumerates the set of values for ZoneSummaryZoneTypeEnum
func GetZoneSummaryZoneTypeEnumValues() []ZoneSummaryZoneTypeEnum {
	values := make([]ZoneSummaryZoneTypeEnum, 0)
	for _, v := range mappingZoneSummaryZoneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneSummaryZoneTypeEnumStringValues Enumerates the set of values in String for ZoneSummaryZoneTypeEnum
func GetZoneSummaryZoneTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
	}
}

// GetMappingZoneSummaryZoneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZoneSummaryZoneTypeEnum(val string) (ZoneSummaryZoneTypeEnum, bool) {
	enum, ok := mappingZoneSummaryZoneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ZoneSummaryLifecycleStateEnum Enum with underlying type: string
type ZoneSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ZoneSummaryLifecycleStateEnum
const (
	ZoneSummaryLifecycleStateActive   ZoneSummaryLifecycleStateEnum = "ACTIVE"
	ZoneSummaryLifecycleStateCreating ZoneSummaryLifecycleStateEnum = "CREATING"
	ZoneSummaryLifecycleStateDeleted  ZoneSummaryLifecycleStateEnum = "DELETED"
	ZoneSummaryLifecycleStateDeleting ZoneSummaryLifecycleStateEnum = "DELETING"
	ZoneSummaryLifecycleStateFailed   ZoneSummaryLifecycleStateEnum = "FAILED"
	ZoneSummaryLifecycleStateUpdating ZoneSummaryLifecycleStateEnum = "UPDATING"
)

var mappingZoneSummaryLifecycleStateEnum = map[string]ZoneSummaryLifecycleStateEnum{
	"ACTIVE":   ZoneSummaryLifecycleStateActive,
	"CREATING": ZoneSummaryLifecycleStateCreating,
	"DELETED":  ZoneSummaryLifecycleStateDeleted,
	"DELETING": ZoneSummaryLifecycleStateDeleting,
	"FAILED":   ZoneSummaryLifecycleStateFailed,
	"UPDATING": ZoneSummaryLifecycleStateUpdating,
}

var mappingZoneSummaryLifecycleStateEnumLowerCase = map[string]ZoneSummaryLifecycleStateEnum{
	"active":   ZoneSummaryLifecycleStateActive,
	"creating": ZoneSummaryLifecycleStateCreating,
	"deleted":  ZoneSummaryLifecycleStateDeleted,
	"deleting": ZoneSummaryLifecycleStateDeleting,
	"failed":   ZoneSummaryLifecycleStateFailed,
	"updating": ZoneSummaryLifecycleStateUpdating,
}

// GetZoneSummaryLifecycleStateEnumValues Enumerates the set of values for ZoneSummaryLifecycleStateEnum
func GetZoneSummaryLifecycleStateEnumValues() []ZoneSummaryLifecycleStateEnum {
	values := make([]ZoneSummaryLifecycleStateEnum, 0)
	for _, v := range mappingZoneSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ZoneSummaryLifecycleStateEnum
func GetZoneSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingZoneSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZoneSummaryLifecycleStateEnum(val string) (ZoneSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingZoneSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
