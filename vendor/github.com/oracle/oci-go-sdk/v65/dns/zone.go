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

// Zone A DNS zone.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Zone struct {

	// The name of the zone.
	Name *string `mandatory:"true" json:"name"`

	// The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL zones.
	ZoneType ZoneZoneTypeEnum `mandatory:"true" json:"zoneType"`

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

	// External master servers for the zone. `externalMasters` becomes a
	// required parameter when the `zoneType` value is `SECONDARY`.
	ExternalMasters []ExternalMaster `mandatory:"true" json:"externalMasters"`

	// External secondary servers for the zone.
	// This field is currently not supported when `zoneType` is `SECONDARY` or `scope` is `PRIVATE`.
	ExternalDownstreams []ExternalDownstream `mandatory:"true" json:"externalDownstreams"`

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
	LifecycleState ZoneLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `mandatory:"true" json:"isProtected"`

	// The authoritative nameservers for the zone.
	Nameservers []Nameserver `mandatory:"true" json:"nameservers"`

	// The OCID of the private view containing the zone. This value will
	// be null for zones in the global DNS, which are publicly resolvable and
	// not part of a private view.
	ViewId *string `mandatory:"false" json:"viewId"`

	// The OCI nameservers that transfer the zone data with external nameservers.
	ZoneTransferServers []ZoneTransferServer `mandatory:"false" json:"zoneTransferServers"`
}

func (m Zone) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Zone) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingZoneZoneTypeEnum(string(m.ZoneType)); !ok && m.ZoneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ZoneType: %s. Supported values are: %s.", m.ZoneType, strings.Join(GetZoneZoneTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingZoneLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetZoneLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ZoneZoneTypeEnum Enum with underlying type: string
type ZoneZoneTypeEnum string

// Set of constants representing the allowable values for ZoneZoneTypeEnum
const (
	ZoneZoneTypePrimary   ZoneZoneTypeEnum = "PRIMARY"
	ZoneZoneTypeSecondary ZoneZoneTypeEnum = "SECONDARY"
)

var mappingZoneZoneTypeEnum = map[string]ZoneZoneTypeEnum{
	"PRIMARY":   ZoneZoneTypePrimary,
	"SECONDARY": ZoneZoneTypeSecondary,
}

var mappingZoneZoneTypeEnumLowerCase = map[string]ZoneZoneTypeEnum{
	"primary":   ZoneZoneTypePrimary,
	"secondary": ZoneZoneTypeSecondary,
}

// GetZoneZoneTypeEnumValues Enumerates the set of values for ZoneZoneTypeEnum
func GetZoneZoneTypeEnumValues() []ZoneZoneTypeEnum {
	values := make([]ZoneZoneTypeEnum, 0)
	for _, v := range mappingZoneZoneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneZoneTypeEnumStringValues Enumerates the set of values in String for ZoneZoneTypeEnum
func GetZoneZoneTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
	}
}

// GetMappingZoneZoneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZoneZoneTypeEnum(val string) (ZoneZoneTypeEnum, bool) {
	enum, ok := mappingZoneZoneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ZoneLifecycleStateEnum Enum with underlying type: string
type ZoneLifecycleStateEnum string

// Set of constants representing the allowable values for ZoneLifecycleStateEnum
const (
	ZoneLifecycleStateActive   ZoneLifecycleStateEnum = "ACTIVE"
	ZoneLifecycleStateCreating ZoneLifecycleStateEnum = "CREATING"
	ZoneLifecycleStateDeleted  ZoneLifecycleStateEnum = "DELETED"
	ZoneLifecycleStateDeleting ZoneLifecycleStateEnum = "DELETING"
	ZoneLifecycleStateFailed   ZoneLifecycleStateEnum = "FAILED"
	ZoneLifecycleStateUpdating ZoneLifecycleStateEnum = "UPDATING"
)

var mappingZoneLifecycleStateEnum = map[string]ZoneLifecycleStateEnum{
	"ACTIVE":   ZoneLifecycleStateActive,
	"CREATING": ZoneLifecycleStateCreating,
	"DELETED":  ZoneLifecycleStateDeleted,
	"DELETING": ZoneLifecycleStateDeleting,
	"FAILED":   ZoneLifecycleStateFailed,
	"UPDATING": ZoneLifecycleStateUpdating,
}

var mappingZoneLifecycleStateEnumLowerCase = map[string]ZoneLifecycleStateEnum{
	"active":   ZoneLifecycleStateActive,
	"creating": ZoneLifecycleStateCreating,
	"deleted":  ZoneLifecycleStateDeleted,
	"deleting": ZoneLifecycleStateDeleting,
	"failed":   ZoneLifecycleStateFailed,
	"updating": ZoneLifecycleStateUpdating,
}

// GetZoneLifecycleStateEnumValues Enumerates the set of values for ZoneLifecycleStateEnum
func GetZoneLifecycleStateEnumValues() []ZoneLifecycleStateEnum {
	values := make([]ZoneLifecycleStateEnum, 0)
	for _, v := range mappingZoneLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneLifecycleStateEnumStringValues Enumerates the set of values in String for ZoneLifecycleStateEnum
func GetZoneLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingZoneLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZoneLifecycleStateEnum(val string) (ZoneLifecycleStateEnum, bool) {
	enum, ok := mappingZoneLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
