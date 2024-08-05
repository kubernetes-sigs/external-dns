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

// RecordOperation An extension of the existing record resource, describing either a
// precondition, an add, or a remove. Preconditions check all fields,
// including read-only data like `recordHash` and `rrsetVersion`.
type RecordOperation struct {

	// The fully qualified domain name where the record can be located.
	Domain *string `mandatory:"false" json:"domain"`

	// A unique identifier for the record within its zone.
	RecordHash *string `mandatory:"false" json:"recordHash"`

	// A Boolean flag indicating whether or not parts of the record
	// are unable to be explicitly managed.
	IsProtected *bool `mandatory:"false" json:"isProtected"`

	// The record's data, as whitespace-delimited tokens in
	// type-specific presentation format. All RDATA is normalized and the
	// returned presentation of your RDATA may differ from its initial input.
	// For more information about RDATA, see Supported DNS Resource Record Types (https://docs.cloud.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm)
	Rdata *string `mandatory:"false" json:"rdata"`

	// The latest version of the record's zone in which its RRSet differs
	// from the preceding version.
	RrsetVersion *string `mandatory:"false" json:"rrsetVersion"`

	// The type of DNS record, such as A or CNAME. For more information, see Resource Record (RR) TYPEs (https://www.iana.org/assignments/dns-parameters/dns-parameters.xhtml#dns-parameters-4).
	Rtype *string `mandatory:"false" json:"rtype"`

	// The Time To Live for the record, in seconds. Using a TTL lower than 30 seconds is not recommended.
	Ttl *int `mandatory:"false" json:"ttl"`

	// A description of how a record relates to a PATCH operation.
	//
	// - `REQUIRE` indicates a precondition that record data **must** already exist.
	// - `PROHIBIT` indicates a precondition that record data **must not** already exist.
	// - `ADD` indicates that record data **must** exist after successful application.
	// - `REMOVE` indicates that record data **must not** exist after successful application.
	//
	//   **Note:** `ADD` and `REMOVE` operations can succeed even if
	//   they require no changes when applied, such as when the described
	//   records are already present or absent.
	//
	//   **Note:** `ADD` and `REMOVE` operations can describe changes for
	//   more than one record.
	//
	//   **Example:** `{ "domain": "www.example.com", "rtype": "AAAA", "ttl": 60 }`
	//   specifies a new TTL for every record in the www.example.com AAAA RRSet.
	Operation RecordOperationOperationEnum `mandatory:"false" json:"operation,omitempty"`
}

func (m RecordOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RecordOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRecordOperationOperationEnum(string(m.Operation)); !ok && m.Operation != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operation: %s. Supported values are: %s.", m.Operation, strings.Join(GetRecordOperationOperationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RecordOperationOperationEnum Enum with underlying type: string
type RecordOperationOperationEnum string

// Set of constants representing the allowable values for RecordOperationOperationEnum
const (
	RecordOperationOperationRequire  RecordOperationOperationEnum = "REQUIRE"
	RecordOperationOperationProhibit RecordOperationOperationEnum = "PROHIBIT"
	RecordOperationOperationAdd      RecordOperationOperationEnum = "ADD"
	RecordOperationOperationRemove   RecordOperationOperationEnum = "REMOVE"
)

var mappingRecordOperationOperationEnum = map[string]RecordOperationOperationEnum{
	"REQUIRE":  RecordOperationOperationRequire,
	"PROHIBIT": RecordOperationOperationProhibit,
	"ADD":      RecordOperationOperationAdd,
	"REMOVE":   RecordOperationOperationRemove,
}

var mappingRecordOperationOperationEnumLowerCase = map[string]RecordOperationOperationEnum{
	"require":  RecordOperationOperationRequire,
	"prohibit": RecordOperationOperationProhibit,
	"add":      RecordOperationOperationAdd,
	"remove":   RecordOperationOperationRemove,
}

// GetRecordOperationOperationEnumValues Enumerates the set of values for RecordOperationOperationEnum
func GetRecordOperationOperationEnumValues() []RecordOperationOperationEnum {
	values := make([]RecordOperationOperationEnum, 0)
	for _, v := range mappingRecordOperationOperationEnum {
		values = append(values, v)
	}
	return values
}

// GetRecordOperationOperationEnumStringValues Enumerates the set of values in String for RecordOperationOperationEnum
func GetRecordOperationOperationEnumStringValues() []string {
	return []string{
		"REQUIRE",
		"PROHIBIT",
		"ADD",
		"REMOVE",
	}
}

// GetMappingRecordOperationOperationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRecordOperationOperationEnum(val string) (RecordOperationOperationEnum, bool) {
	enum, ok := mappingRecordOperationOperationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
