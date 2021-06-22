// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// CreateMigratedDynectZoneDetails The body for migrating a zone from DynECT.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateMigratedDynectZoneDetails struct {

	// The name of the zone.
	Name *string `mandatory:"true" json:"name"`

	// The OCID of the compartment containing the zone.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	DynectMigrationDetails *DynectMigrationDetails `mandatory:"false" json:"dynectMigrationDetails"`
}

//GetName returns Name
func (m CreateMigratedDynectZoneDetails) GetName() *string {
	return m.Name
}

//GetCompartmentId returns CompartmentId
func (m CreateMigratedDynectZoneDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetFreeformTags returns FreeformTags
func (m CreateMigratedDynectZoneDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateMigratedDynectZoneDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateMigratedDynectZoneDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateMigratedDynectZoneDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateMigratedDynectZoneDetails CreateMigratedDynectZoneDetails
	s := struct {
		DiscriminatorParam string `json:"migrationSource"`
		MarshalTypeCreateMigratedDynectZoneDetails
	}{
		"DYNECT",
		(MarshalTypeCreateMigratedDynectZoneDetails)(m),
	}

	return json.Marshal(&s)
}
