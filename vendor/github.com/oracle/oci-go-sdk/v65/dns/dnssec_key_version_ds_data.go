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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DnssecKeyVersionDsData Data for a parent zone DS record corresponding to this key-signing key (KSK).
type DnssecKeyVersionDsData struct {

	// Presentation-format DS record data that must be added to the parent zone. For more information about RDATA,
	// see Supported DNS Resource Record Types (https://docs.oracle.com/iaas/Content/DNS/Reference/supporteddnsresource.htm)
	Rdata *string `mandatory:"false" json:"rdata"`

	// The type of the digest associated with the rdata.
	DigestType DnssecDigestTypeEnum `mandatory:"false" json:"digestType,omitempty"`
}

func (m DnssecKeyVersionDsData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DnssecKeyVersionDsData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDnssecDigestTypeEnum(string(m.DigestType)); !ok && m.DigestType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DigestType: %s. Supported values are: %s.", m.DigestType, strings.Join(GetDnssecDigestTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
