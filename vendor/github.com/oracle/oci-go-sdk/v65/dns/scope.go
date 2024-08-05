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
	"strings"
)

// ScopeEnum Enum with underlying type: string
type ScopeEnum string

// Set of constants representing the allowable values for ScopeEnum
const (
	ScopeGlobal  ScopeEnum = "GLOBAL"
	ScopePrivate ScopeEnum = "PRIVATE"
)

var mappingScopeEnum = map[string]ScopeEnum{
	"GLOBAL":  ScopeGlobal,
	"PRIVATE": ScopePrivate,
}

var mappingScopeEnumLowerCase = map[string]ScopeEnum{
	"global":  ScopeGlobal,
	"private": ScopePrivate,
}

// GetScopeEnumValues Enumerates the set of values for ScopeEnum
func GetScopeEnumValues() []ScopeEnum {
	values := make([]ScopeEnum, 0)
	for _, v := range mappingScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetScopeEnumStringValues Enumerates the set of values in String for ScopeEnum
func GetScopeEnumStringValues() []string {
	return []string{
		"GLOBAL",
		"PRIVATE",
	}
}

// GetMappingScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScopeEnum(val string) (ScopeEnum, bool) {
	enum, ok := mappingScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
