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

// ResolverSummary An OCI DNS resolver. If the resolver has an attached VCN, the VCN will attempt to answer queries based on the
// attached views in priority order. If the query does not match any of the attached views, the query will be
// evaluated against the default view. If the default view does not match, the rules will be evaluated in
// priority order. If no rules match the query, answers come from Internet DNS. A resolver may have a maximum of 10
// resolver endpoints.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ResolverSummary struct {

	// The OCID of the owning compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the resolver.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	//
	// **Example:** `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The OCID of the resolver.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ"
	// format with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the resource.
	LifecycleState ResolverSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `mandatory:"true" json:"isProtected"`

	// The OCID of the attached VCN.
	AttachedVcnId *string `mandatory:"false" json:"attachedVcnId"`

	// The OCID of the default view.
	DefaultViewId *string `mandatory:"false" json:"defaultViewId"`
}

func (m ResolverSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResolverSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResolverSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResolverSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResolverSummaryLifecycleStateEnum Enum with underlying type: string
type ResolverSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ResolverSummaryLifecycleStateEnum
const (
	ResolverSummaryLifecycleStateActive   ResolverSummaryLifecycleStateEnum = "ACTIVE"
	ResolverSummaryLifecycleStateCreating ResolverSummaryLifecycleStateEnum = "CREATING"
	ResolverSummaryLifecycleStateDeleted  ResolverSummaryLifecycleStateEnum = "DELETED"
	ResolverSummaryLifecycleStateDeleting ResolverSummaryLifecycleStateEnum = "DELETING"
	ResolverSummaryLifecycleStateFailed   ResolverSummaryLifecycleStateEnum = "FAILED"
	ResolverSummaryLifecycleStateUpdating ResolverSummaryLifecycleStateEnum = "UPDATING"
)

var mappingResolverSummaryLifecycleStateEnum = map[string]ResolverSummaryLifecycleStateEnum{
	"ACTIVE":   ResolverSummaryLifecycleStateActive,
	"CREATING": ResolverSummaryLifecycleStateCreating,
	"DELETED":  ResolverSummaryLifecycleStateDeleted,
	"DELETING": ResolverSummaryLifecycleStateDeleting,
	"FAILED":   ResolverSummaryLifecycleStateFailed,
	"UPDATING": ResolverSummaryLifecycleStateUpdating,
}

var mappingResolverSummaryLifecycleStateEnumLowerCase = map[string]ResolverSummaryLifecycleStateEnum{
	"active":   ResolverSummaryLifecycleStateActive,
	"creating": ResolverSummaryLifecycleStateCreating,
	"deleted":  ResolverSummaryLifecycleStateDeleted,
	"deleting": ResolverSummaryLifecycleStateDeleting,
	"failed":   ResolverSummaryLifecycleStateFailed,
	"updating": ResolverSummaryLifecycleStateUpdating,
}

// GetResolverSummaryLifecycleStateEnumValues Enumerates the set of values for ResolverSummaryLifecycleStateEnum
func GetResolverSummaryLifecycleStateEnumValues() []ResolverSummaryLifecycleStateEnum {
	values := make([]ResolverSummaryLifecycleStateEnum, 0)
	for _, v := range mappingResolverSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResolverSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ResolverSummaryLifecycleStateEnum
func GetResolverSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingResolverSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResolverSummaryLifecycleStateEnum(val string) (ResolverSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingResolverSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
