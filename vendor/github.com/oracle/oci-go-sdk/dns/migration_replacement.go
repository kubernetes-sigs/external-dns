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

// MigrationReplacement A record to add to a zone in replacement of contents that cannot be migrated.
type MigrationReplacement struct {

	// The canonical name for the type of the replacement record, such as A or CNAME.
	Rtype *string `mandatory:"true" json:"rtype"`

	// The Time To Live of the replacement record, in seconds.
	Ttl *int `mandatory:"true" json:"ttl"`

	// The record data of the replacement record, as whitespace-delimited tokens in
	// type-specific presentation format.
	Rdata *string `mandatory:"true" json:"rdata"`

	// The canonical name for a substitute type of the replacement record to be used if the specified `rtype` is not allowed at the domain. The specified `ttl` and `rdata` will still apply with the substitute type.
	SubstituteRtype *string `mandatory:"false" json:"substituteRtype"`
}

func (m MigrationReplacement) String() string {
	return common.PointerString(m)
}
