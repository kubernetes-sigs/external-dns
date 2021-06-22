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

// ExternalMaster An external master name server used as the source of zone data.
// May either have a zone-embedded TSIG or reference a TSIG key by OCID,
// but not both.
type ExternalMaster struct {

	// The server's IP address (IPv4 or IPv6).
	Address *string `mandatory:"true" json:"address"`

	// The server's port. Port value must be a value of 53, otherwise omit
	// the port value.
	Port *int `mandatory:"false" json:"port"`

	Tsig *Tsig `mandatory:"false" json:"tsig"`

	// The OCID of the TSIG key.
	TsigKeyId *string `mandatory:"false" json:"tsigKeyId"`
}

func (m ExternalMaster) String() string {
	return common.PointerString(m)
}
