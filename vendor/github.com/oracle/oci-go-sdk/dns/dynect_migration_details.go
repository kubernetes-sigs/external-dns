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
	"github.com/oracle/oci-go-sdk/common"
)

// DynectMigrationDetails Details specific to performing a DynECT zone migration.
type DynectMigrationDetails struct {

	// DynECT customer name the zone belongs to.
	CustomerName *string `mandatory:"true" json:"customerName"`

	// DynECT API username to perform the migration with.
	Username *string `mandatory:"true" json:"username"`

	// DynECT API password for the provided username.
	Password *string `mandatory:"true" json:"password"`

	// A map of fully-qualified domain names (FQDNs) to an array of `MigrationReplacement` objects.
	HttpRedirectReplacements map[string][]MigrationReplacement `mandatory:"false" json:"httpRedirectReplacements"`
}

func (m DynectMigrationDetails) String() string {
	return common.PointerString(m)
}
