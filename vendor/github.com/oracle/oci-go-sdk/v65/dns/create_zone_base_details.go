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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateZoneBaseDetails The body for either defining a new zone or migrating a zone from migrationSource. This is determined by the migrationSource discriminator.
// NONE indicates creation of a new zone (default). DYNECT indicates migration from a DynECT zone.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateZoneBaseDetails interface {

	// The name of the zone.
	// Global zone names must be unique across all other zones within the realm. Private zone names must be unique
	// within their view.
	// Unicode characters will be converted into punycode, see RFC 3492 (https://tools.ietf.org/html/rfc3492).
	GetName() *string

	// The OCID of the compartment containing the zone.
	GetCompartmentId() *string

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createzonebasedetails struct {
	JsonData        []byte
	FreeformTags    map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags     map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	Name            *string                           `mandatory:"true" json:"name"`
	CompartmentId   *string                           `mandatory:"true" json:"compartmentId"`
	MigrationSource string                            `json:"migrationSource"`
}

// UnmarshalJSON unmarshals json
func (m *createzonebasedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatezonebasedetails createzonebasedetails
	s := struct {
		Model Unmarshalercreatezonebasedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.CompartmentId = s.Model.CompartmentId
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.MigrationSource = s.Model.MigrationSource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createzonebasedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MigrationSource {
	case "NONE":
		mm := CreateZoneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNECT":
		mm := CreateMigratedDynectZoneDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateZoneBaseDetails: %s.", m.MigrationSource)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m createzonebasedetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createzonebasedetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetName returns Name
func (m createzonebasedetails) GetName() *string {
	return m.Name
}

// GetCompartmentId returns CompartmentId
func (m createzonebasedetails) GetCompartmentId() *string {
	return m.CompartmentId
}

func (m createzonebasedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createzonebasedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateZoneBaseDetailsMigrationSourceEnum Enum with underlying type: string
type CreateZoneBaseDetailsMigrationSourceEnum string

// Set of constants representing the allowable values for CreateZoneBaseDetailsMigrationSourceEnum
const (
	CreateZoneBaseDetailsMigrationSourceNone   CreateZoneBaseDetailsMigrationSourceEnum = "NONE"
	CreateZoneBaseDetailsMigrationSourceDynect CreateZoneBaseDetailsMigrationSourceEnum = "DYNECT"
)

var mappingCreateZoneBaseDetailsMigrationSourceEnum = map[string]CreateZoneBaseDetailsMigrationSourceEnum{
	"NONE":   CreateZoneBaseDetailsMigrationSourceNone,
	"DYNECT": CreateZoneBaseDetailsMigrationSourceDynect,
}

var mappingCreateZoneBaseDetailsMigrationSourceEnumLowerCase = map[string]CreateZoneBaseDetailsMigrationSourceEnum{
	"none":   CreateZoneBaseDetailsMigrationSourceNone,
	"dynect": CreateZoneBaseDetailsMigrationSourceDynect,
}

// GetCreateZoneBaseDetailsMigrationSourceEnumValues Enumerates the set of values for CreateZoneBaseDetailsMigrationSourceEnum
func GetCreateZoneBaseDetailsMigrationSourceEnumValues() []CreateZoneBaseDetailsMigrationSourceEnum {
	values := make([]CreateZoneBaseDetailsMigrationSourceEnum, 0)
	for _, v := range mappingCreateZoneBaseDetailsMigrationSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateZoneBaseDetailsMigrationSourceEnumStringValues Enumerates the set of values in String for CreateZoneBaseDetailsMigrationSourceEnum
func GetCreateZoneBaseDetailsMigrationSourceEnumStringValues() []string {
	return []string{
		"NONE",
		"DYNECT",
	}
}

// GetMappingCreateZoneBaseDetailsMigrationSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateZoneBaseDetailsMigrationSourceEnum(val string) (CreateZoneBaseDetailsMigrationSourceEnum, bool) {
	enum, ok := mappingCreateZoneBaseDetailsMigrationSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
