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
	"strings"
)

// ZoneResolutionModeEnum Enum with underlying type: string
type ZoneResolutionModeEnum string

// Set of constants representing the allowable values for ZoneResolutionModeEnum
const (
	ZoneResolutionModeStatic           ZoneResolutionModeEnum = "STATIC"
	ZoneResolutionModeTransparent      ZoneResolutionModeEnum = "TRANSPARENT"
	ZoneResolutionModeRtypeTransparent ZoneResolutionModeEnum = "RTYPE_TRANSPARENT"
)

var mappingZoneResolutionModeEnum = map[string]ZoneResolutionModeEnum{
	"STATIC":            ZoneResolutionModeStatic,
	"TRANSPARENT":       ZoneResolutionModeTransparent,
	"RTYPE_TRANSPARENT": ZoneResolutionModeRtypeTransparent,
}

var mappingZoneResolutionModeEnumLowerCase = map[string]ZoneResolutionModeEnum{
	"static":            ZoneResolutionModeStatic,
	"transparent":       ZoneResolutionModeTransparent,
	"rtype_transparent": ZoneResolutionModeRtypeTransparent,
}

// GetZoneResolutionModeEnumValues Enumerates the set of values for ZoneResolutionModeEnum
func GetZoneResolutionModeEnumValues() []ZoneResolutionModeEnum {
	values := make([]ZoneResolutionModeEnum, 0)
	for _, v := range mappingZoneResolutionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneResolutionModeEnumStringValues Enumerates the set of values in String for ZoneResolutionModeEnum
func GetZoneResolutionModeEnumStringValues() []string {
	return []string{
		"STATIC",
		"TRANSPARENT",
		"RTYPE_TRANSPARENT",
	}
}

// GetMappingZoneResolutionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZoneResolutionModeEnum(val string) (ZoneResolutionModeEnum, bool) {
	enum, ok := mappingZoneResolutionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
