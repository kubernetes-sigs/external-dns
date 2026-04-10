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

// ZoneDnssecStateEnum Enum with underlying type: string
type ZoneDnssecStateEnum string

// Set of constants representing the allowable values for ZoneDnssecStateEnum
const (
	ZoneDnssecStateEnabled  ZoneDnssecStateEnum = "ENABLED"
	ZoneDnssecStateDisabled ZoneDnssecStateEnum = "DISABLED"
)

var mappingZoneDnssecStateEnum = map[string]ZoneDnssecStateEnum{
	"ENABLED":  ZoneDnssecStateEnabled,
	"DISABLED": ZoneDnssecStateDisabled,
}

var mappingZoneDnssecStateEnumLowerCase = map[string]ZoneDnssecStateEnum{
	"enabled":  ZoneDnssecStateEnabled,
	"disabled": ZoneDnssecStateDisabled,
}

// GetZoneDnssecStateEnumValues Enumerates the set of values for ZoneDnssecStateEnum
func GetZoneDnssecStateEnumValues() []ZoneDnssecStateEnum {
	values := make([]ZoneDnssecStateEnum, 0)
	for _, v := range mappingZoneDnssecStateEnum {
		values = append(values, v)
	}
	return values
}

// GetZoneDnssecStateEnumStringValues Enumerates the set of values in String for ZoneDnssecStateEnum
func GetZoneDnssecStateEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingZoneDnssecStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingZoneDnssecStateEnum(val string) (ZoneDnssecStateEnum, bool) {
	enum, ok := mappingZoneDnssecStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
