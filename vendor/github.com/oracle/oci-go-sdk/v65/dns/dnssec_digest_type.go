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

// DnssecDigestTypeEnum Enum with underlying type: string
type DnssecDigestTypeEnum string

// Set of constants representing the allowable values for DnssecDigestTypeEnum
const (
	DnssecDigestTypeSha256 DnssecDigestTypeEnum = "SHA_256"
)

var mappingDnssecDigestTypeEnum = map[string]DnssecDigestTypeEnum{
	"SHA_256": DnssecDigestTypeSha256,
}

var mappingDnssecDigestTypeEnumLowerCase = map[string]DnssecDigestTypeEnum{
	"sha_256": DnssecDigestTypeSha256,
}

// GetDnssecDigestTypeEnumValues Enumerates the set of values for DnssecDigestTypeEnum
func GetDnssecDigestTypeEnumValues() []DnssecDigestTypeEnum {
	values := make([]DnssecDigestTypeEnum, 0)
	for _, v := range mappingDnssecDigestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDnssecDigestTypeEnumStringValues Enumerates the set of values in String for DnssecDigestTypeEnum
func GetDnssecDigestTypeEnumStringValues() []string {
	return []string{
		"SHA_256",
	}
}

// GetMappingDnssecDigestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDnssecDigestTypeEnum(val string) (DnssecDigestTypeEnum, bool) {
	enum, ok := mappingDnssecDigestTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
