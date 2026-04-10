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

// ZskDnssecKeyVersion A zone signing key (ZSK) version. The version information contains timing and configuration data for the ZSK that is used to
// apply DNSSEC on the zone.
type ZskDnssecKeyVersion struct {

	// The UUID of the `DnssecKeyVersion`.
	Uuid *string `mandatory:"false" json:"uuid"`

	// The signing algorithm used for the key.
	Algorithm DnssecSigningAlgorithmEnum `mandatory:"false" json:"algorithm,omitempty"`

	// The length of the corresponding private key in bytes, expressed as an integer.
	LengthInBytes *int `mandatory:"false" json:"lengthInBytes"`

	// The date and time the key version was created, expressed in RFC 3339 timestamp format.
	// **Example:** `2016-07-22T17:23:59:00Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the key version was, or will be, published, expressed in RFC 3339 timestamp format. This is
	// when the zone contents will include a DNSKEY record corresponding to the key material.
	// **Example:** `2016-07-22T17:23:59:00Z`
	TimePublished *common.SDKTime `mandatory:"false" json:"timePublished"`

	// The date and time the key version went, or will go, active, expressed in RFC 3339 timestamp format. This is
	// when the key material will be used to generate RRSIGs.
	// **Example:** `2016-07-22T17:23:59:00Z`
	TimeActivated *common.SDKTime `mandatory:"false" json:"timeActivated"`

	// The date and time the key version went, or will go, inactive, expressed in RFC 3339 timestamp format. This
	// is when the key material will no longer be used to generate RRSIGs. For a key signing key (KSK) `DnssecKeyVersion`, this is
	// populated after `PromoteZoneDnssecKeyVersion` has been called on its successor `DnssecKeyVersion`.
	// **Example:** `2016-07-22T17:23:59:00Z`
	TimeInactivated *common.SDKTime `mandatory:"false" json:"timeInactivated"`

	// The date and time the key version was, or will be, unpublished, expressed in RFC 3339 timestamp format. This
	// is when the corresponding DNSKEY will be removed from zone contents. For a key signing key (KSK) `DnssecKeyVersion`, this is
	// populated after `PromoteZoneDnssecKeyVersion` has been called on its successor `DnssecKeyVersion`.
	// **Example:** `2016-07-22T17:23:59:00Z`
	TimeUnpublished *common.SDKTime `mandatory:"false" json:"timeUnpublished"`

	// The date and time at which the recommended key version publication/activation lifetime ends, expressed in RFC
	// 3339 timestamp format. This is when the corresponding DNSKEY should no longer exist in zone contents and no
	// longer be used to generate RRSIGs. For a key sigining key (KSK), if `PromoteZoneDnssecKeyVersion` has not been called on this
	// `DnssecKeyVersion`'s successor then it will remain active for arbitrarily long past its recommended lifetime.
	// This prevents service disruption at the potential increased risk of key compromise.
	// **Example:** `2016-07-22T17:23:59:00Z`
	TimeExpired *common.SDKTime `mandatory:"false" json:"timeExpired"`

	// The date and time the key version was promoted expressed in RFC 3339 timestamp format.
	// **Example:** `2016-07-22T17:23:59:00Z`
	TimePromoted *common.SDKTime `mandatory:"false" json:"timePromoted"`

	// When populated, this is the UUID of the `DnssecKeyVersion` that this `DnssecKeyVersion` will replace or has
	// replaced.
	PredecessorDnssecKeyVersionUuid *string `mandatory:"false" json:"predecessorDnssecKeyVersionUuid"`

	// When populated, this is the UUID of the `DnssecKeyVersion` that will replace, or has replaced, this
	// `DnssecKeyVersion`.
	SuccessorDnssecKeyVersionUuid *string `mandatory:"false" json:"successorDnssecKeyVersionUuid"`

	// The key tag associated with the `DnssecKeyVersion`. This key tag will be present in the RRSIG and DS records
	// associated with the key material for this `DnssecKeyVersion`. For more information about key tags, see
	// RFC 4034 (https://tools.ietf.org/html/rfc4034).
	KeyTag *int `mandatory:"false" json:"keyTag"`
}

func (m ZskDnssecKeyVersion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZskDnssecKeyVersion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDnssecSigningAlgorithmEnum(string(m.Algorithm)); !ok && m.Algorithm != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Algorithm: %s. Supported values are: %s.", m.Algorithm, strings.Join(GetDnssecSigningAlgorithmEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
