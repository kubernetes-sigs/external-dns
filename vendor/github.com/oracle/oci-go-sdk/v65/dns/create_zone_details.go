// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DNS API
//
// API for the DNS service. Use this API to manage DNS zones, records, and other DNS resources.
// For more information, see Overview of the DNS Service (https://docs.oracle.com/iaas/Content/DNS/Concepts/dnszonemanagement.htm).
//

package dns

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateZoneDetails The body for defining a new zone.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateZoneDetails struct {

	// The name of the zone.
	// Global zone names must be unique across all other zones within the realm. Private zone names must be unique
	// within their view.
	// Unicode characters will be converted into punycode, see RFC 3492 (https://tools.ietf.org/html/rfc3492).
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment containing the zone.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// This value will be null for zones in the global DNS.
	ViewId *string `mandatory:"false" json:"viewId"`

	// External master servers for the zone. `externalMasters` becomes a
	// required parameter when the `zoneType` value is `SECONDARY`.
	ExternalMasters []ExternalMaster `mandatory:"false" json:"externalMasters"`

	// External secondary servers for the zone.
	// This field is currently not supported when `zoneType` is `SECONDARY` or `scope` is `PRIVATE`.
	ExternalDownstreams []ExternalDownstream `mandatory:"false" json:"externalDownstreams"`

	// The type of the zone. Must be either `PRIMARY` or `SECONDARY`. `SECONDARY` is only supported for GLOBAL
	// zones.
	ZoneType CreateZoneDetailsZoneTypeEnum `mandatory:"false" json:"zoneType,omitempty"`

	// The scope of the zone.
	Scope ScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// The resolution mode of a zone defines behavior related to how query responses can be handled.
	ResolutionMode ZoneResolutionModeEnum `mandatory:"false" json:"resolutionMode,omitempty"`

	// The state of DNSSEC on the zone.
	// For DNSSEC to function, every parent zone in the DNS tree up to the top-level domain (or an independent
	// trust anchor) must also have DNSSEC correctly set up.
	// After enabling DNSSEC, you must add a DS record to the zone's parent zone containing the
	// `KskDnssecKeyVersion` data. You can find the DS data in the `dsData` attribute of the `KskDnssecKeyVersion`.
	// Then, use the `PromoteZoneDnssecKeyVersion` operation to promote the `KskDnssecKeyVersion`.
	// New `KskDnssecKeyVersion`s are generated annually, a week before the existing `KskDnssecKeyVersion`'s expiration.
	// To rollover a `KskDnssecKeyVersion`, you must replace the parent zone's DS record containing the old
	// `KskDnssecKeyVersion` data with the data from the new `KskDnssecKeyVersion`.
	// To remove the old DS record without causing service disruption, wait until the old DS record's TTL has
	// expired, and the new DS record has propagated. After the DS replacement has been completed, then the
	// `PromoteZoneDnssecKeyVersion` operation must be called.
	// Metrics are emitted in the `oci_dns` namespace daily for each `KskDnssecKeyVersion` indicating how many
	// days are left until expiration.
	// We recommend that you set up alarms and notifications for KskDnssecKeyVersion expiration so that the
	// necessary parent zone updates can be made and the `PromoteZoneDnssecKeyVersion` operation can be called.
	// Enabling DNSSEC results in additional records in DNS responses which increases their size and can
	// cause higher response latency.
	// For more information, see DNSSEC (https://docs.oracle.com/iaas/Content/DNS/Concepts/dnssec.htm).
	DnssecState ZoneDnssecStateEnum `mandatory:"false" json:"dnssecState,omitempty"`
}

// GetName returns Name
func (m CreateZoneDetails) GetName() *string {
	return m.Name
}

// GetCompartmentId returns CompartmentId
func (m CreateZoneDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateZoneDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateZoneDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateZoneDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateZoneDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateZoneDetailsZoneTypeEnum(string(m.ZoneType)); !ok && m.ZoneType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ZoneType: %s. Supported values are: %s.", m.ZoneType, strings.Join(GetCreateZoneDetailsZoneTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingZoneResolutionModeEnum(string(m.ResolutionMode)); !ok && m.ResolutionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResolutionMode: %s. Supported values are: %s.", m.ResolutionMode, strings.Join(GetZoneResolutionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingZoneDnssecStateEnum(string(m.DnssecState)); !ok && m.DnssecState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DnssecState: %s. Supported values are: %s.", m.DnssecState, strings.Join(GetZoneDnssecStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateZoneDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateZoneDetails CreateZoneDetails
	s := struct {
		DiscriminatorParam string `json:"migrationSource"`
		MarshalTypeCreateZoneDetails
	}{
		"NONE",
		(MarshalTypeCreateZoneDetails)(m),
	}

	return json.Marshal(&s)
}

// CreateZoneDetailsZoneTypeEnum Enum with underlying type: string
type CreateZoneDetailsZoneTypeEnum string

// Set of constants representing the allowable values for CreateZoneDetailsZoneTypeEnum
const (
	CreateZoneDetailsZoneTypePrimary   CreateZoneDetailsZoneTypeEnum = "PRIMARY"
	CreateZoneDetailsZoneTypeSecondary CreateZoneDetailsZoneTypeEnum = "SECONDARY"
)

var mappingCreateZoneDetailsZoneTypeEnum = map[string]CreateZoneDetailsZoneTypeEnum{
	"PRIMARY":   CreateZoneDetailsZoneTypePrimary,
	"SECONDARY": CreateZoneDetailsZoneTypeSecondary,
}

var mappingCreateZoneDetailsZoneTypeEnumLowerCase = map[string]CreateZoneDetailsZoneTypeEnum{
	"primary":   CreateZoneDetailsZoneTypePrimary,
	"secondary": CreateZoneDetailsZoneTypeSecondary,
}

// GetCreateZoneDetailsZoneTypeEnumValues Enumerates the set of values for CreateZoneDetailsZoneTypeEnum
func GetCreateZoneDetailsZoneTypeEnumValues() []CreateZoneDetailsZoneTypeEnum {
	values := make([]CreateZoneDetailsZoneTypeEnum, 0)
	for _, v := range mappingCreateZoneDetailsZoneTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateZoneDetailsZoneTypeEnumStringValues Enumerates the set of values in String for CreateZoneDetailsZoneTypeEnum
func GetCreateZoneDetailsZoneTypeEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"SECONDARY",
	}
}

// GetMappingCreateZoneDetailsZoneTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateZoneDetailsZoneTypeEnum(val string) (CreateZoneDetailsZoneTypeEnum, bool) {
	enum, ok := mappingCreateZoneDetailsZoneTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
