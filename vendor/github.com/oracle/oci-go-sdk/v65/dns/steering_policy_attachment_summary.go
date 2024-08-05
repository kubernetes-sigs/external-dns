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

// SteeringPolicyAttachmentSummary An attachment between a steering policy and a domain.
type SteeringPolicyAttachmentSummary struct {

	// The OCID of the attached steering policy.
	SteeringPolicyId *string `mandatory:"true" json:"steeringPolicyId"`

	// The OCID of the attached zone.
	ZoneId *string `mandatory:"true" json:"zoneId"`

	// The attached domain within the attached zone.
	DomainName *string `mandatory:"true" json:"domainName"`

	// A user-friendly name for the steering policy attachment.
	// Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The record types covered by the attachment at the domain. The set of record types is
	// determined by aggregating the record types from the answers defined in the steering
	// policy.
	Rtypes []string `mandatory:"true" json:"rtypes"`

	// The OCID of the compartment containing the steering policy attachment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created, expressed in RFC 3339 timestamp format.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the resource.
	LifecycleState SteeringPolicyAttachmentSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

func (m SteeringPolicyAttachmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SteeringPolicyAttachmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSteeringPolicyAttachmentSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSteeringPolicyAttachmentSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SteeringPolicyAttachmentSummaryLifecycleStateEnum Enum with underlying type: string
type SteeringPolicyAttachmentSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SteeringPolicyAttachmentSummaryLifecycleStateEnum
const (
	SteeringPolicyAttachmentSummaryLifecycleStateCreating SteeringPolicyAttachmentSummaryLifecycleStateEnum = "CREATING"
	SteeringPolicyAttachmentSummaryLifecycleStateActive   SteeringPolicyAttachmentSummaryLifecycleStateEnum = "ACTIVE"
	SteeringPolicyAttachmentSummaryLifecycleStateDeleting SteeringPolicyAttachmentSummaryLifecycleStateEnum = "DELETING"
)

var mappingSteeringPolicyAttachmentSummaryLifecycleStateEnum = map[string]SteeringPolicyAttachmentSummaryLifecycleStateEnum{
	"CREATING": SteeringPolicyAttachmentSummaryLifecycleStateCreating,
	"ACTIVE":   SteeringPolicyAttachmentSummaryLifecycleStateActive,
	"DELETING": SteeringPolicyAttachmentSummaryLifecycleStateDeleting,
}

var mappingSteeringPolicyAttachmentSummaryLifecycleStateEnumLowerCase = map[string]SteeringPolicyAttachmentSummaryLifecycleStateEnum{
	"creating": SteeringPolicyAttachmentSummaryLifecycleStateCreating,
	"active":   SteeringPolicyAttachmentSummaryLifecycleStateActive,
	"deleting": SteeringPolicyAttachmentSummaryLifecycleStateDeleting,
}

// GetSteeringPolicyAttachmentSummaryLifecycleStateEnumValues Enumerates the set of values for SteeringPolicyAttachmentSummaryLifecycleStateEnum
func GetSteeringPolicyAttachmentSummaryLifecycleStateEnumValues() []SteeringPolicyAttachmentSummaryLifecycleStateEnum {
	values := make([]SteeringPolicyAttachmentSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSteeringPolicyAttachmentSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSteeringPolicyAttachmentSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SteeringPolicyAttachmentSummaryLifecycleStateEnum
func GetSteeringPolicyAttachmentSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
	}
}

// GetMappingSteeringPolicyAttachmentSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSteeringPolicyAttachmentSummaryLifecycleStateEnum(val string) (SteeringPolicyAttachmentSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSteeringPolicyAttachmentSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
