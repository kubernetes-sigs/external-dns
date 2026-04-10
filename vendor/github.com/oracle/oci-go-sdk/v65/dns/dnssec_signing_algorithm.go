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

// DnssecSigningAlgorithmEnum Enum with underlying type: string
type DnssecSigningAlgorithmEnum string

// Set of constants representing the allowable values for DnssecSigningAlgorithmEnum
const (
	DnssecSigningAlgorithmRsasha256 DnssecSigningAlgorithmEnum = "RSASHA256"
)

var mappingDnssecSigningAlgorithmEnum = map[string]DnssecSigningAlgorithmEnum{
	"RSASHA256": DnssecSigningAlgorithmRsasha256,
}

var mappingDnssecSigningAlgorithmEnumLowerCase = map[string]DnssecSigningAlgorithmEnum{
	"rsasha256": DnssecSigningAlgorithmRsasha256,
}

// GetDnssecSigningAlgorithmEnumValues Enumerates the set of values for DnssecSigningAlgorithmEnum
func GetDnssecSigningAlgorithmEnumValues() []DnssecSigningAlgorithmEnum {
	values := make([]DnssecSigningAlgorithmEnum, 0)
	for _, v := range mappingDnssecSigningAlgorithmEnum {
		values = append(values, v)
	}
	return values
}

// GetDnssecSigningAlgorithmEnumStringValues Enumerates the set of values in String for DnssecSigningAlgorithmEnum
func GetDnssecSigningAlgorithmEnumStringValues() []string {
	return []string{
		"RSASHA256",
	}
}

// GetMappingDnssecSigningAlgorithmEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnssecSigningAlgorithmEnum(val string) (DnssecSigningAlgorithmEnum, bool) {
	enum, ok := mappingDnssecSigningAlgorithmEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
