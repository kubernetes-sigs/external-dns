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

// Resolver An OCI DNS resolver. If the resolver has an attached VCN, the VCN will attempt to answer queries based on the
// attached views in priority order. If the query does not match any of the attached views, the query will be
// evaluated against the default view. If the default view does not match, the rules will be evaluated in
// priority order. If no rules match the query, answers come from Internet DNS. A resolver may have a maximum of 10
// resolver endpoints.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type Resolver struct {

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
	LifecycleState ResolverLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// A Boolean flag indicating whether or not parts of the resource are unable to be explicitly managed.
	IsProtected *bool `mandatory:"true" json:"isProtected"`

	// Read-only array of endpoints for the resolver.
	Endpoints []ResolverEndpointSummary `mandatory:"true" json:"endpoints"`

	// The attached views. Views are evaluated in order.
	AttachedViews []AttachedView `mandatory:"true" json:"attachedViews"`

	// The OCID of the attached VCN.
	AttachedVcnId *string `mandatory:"false" json:"attachedVcnId"`

	// The OCID of the default view.
	DefaultViewId *string `mandatory:"false" json:"defaultViewId"`

	// Rules for the resolver. Rules are evaluated in order.
	Rules []ResolverRule `mandatory:"false" json:"rules"`
}

func (m Resolver) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Resolver) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResolverLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResolverLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Resolver) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		AttachedVcnId  *string                           `json:"attachedVcnId"`
		DefaultViewId  *string                           `json:"defaultViewId"`
		Rules          []resolverrule                    `json:"rules"`
		CompartmentId  *string                           `json:"compartmentId"`
		DisplayName    *string                           `json:"displayName"`
		FreeformTags   map[string]string                 `json:"freeformTags"`
		DefinedTags    map[string]map[string]interface{} `json:"definedTags"`
		Id             *string                           `json:"id"`
		TimeCreated    *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated    *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState ResolverLifecycleStateEnum        `json:"lifecycleState"`
		Self           *string                           `json:"self"`
		IsProtected    *bool                             `json:"isProtected"`
		Endpoints      []resolverendpointsummary         `json:"endpoints"`
		AttachedViews  []AttachedView                    `json:"attachedViews"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.AttachedVcnId = model.AttachedVcnId

	m.DefaultViewId = model.DefaultViewId

	m.Rules = make([]ResolverRule, len(model.Rules))
	for i, n := range model.Rules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Rules[i] = nn.(ResolverRule)
		} else {
			m.Rules[i] = nil
		}
	}
	m.CompartmentId = model.CompartmentId

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.Id = model.Id

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.Self = model.Self

	m.IsProtected = model.IsProtected

	m.Endpoints = make([]ResolverEndpointSummary, len(model.Endpoints))
	for i, n := range model.Endpoints {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Endpoints[i] = nn.(ResolverEndpointSummary)
		} else {
			m.Endpoints[i] = nil
		}
	}
	m.AttachedViews = make([]AttachedView, len(model.AttachedViews))
	copy(m.AttachedViews, model.AttachedViews)
	return
}

// ResolverLifecycleStateEnum Enum with underlying type: string
type ResolverLifecycleStateEnum string

// Set of constants representing the allowable values for ResolverLifecycleStateEnum
const (
	ResolverLifecycleStateActive   ResolverLifecycleStateEnum = "ACTIVE"
	ResolverLifecycleStateCreating ResolverLifecycleStateEnum = "CREATING"
	ResolverLifecycleStateDeleted  ResolverLifecycleStateEnum = "DELETED"
	ResolverLifecycleStateDeleting ResolverLifecycleStateEnum = "DELETING"
	ResolverLifecycleStateFailed   ResolverLifecycleStateEnum = "FAILED"
	ResolverLifecycleStateUpdating ResolverLifecycleStateEnum = "UPDATING"
)

var mappingResolverLifecycleStateEnum = map[string]ResolverLifecycleStateEnum{
	"ACTIVE":   ResolverLifecycleStateActive,
	"CREATING": ResolverLifecycleStateCreating,
	"DELETED":  ResolverLifecycleStateDeleted,
	"DELETING": ResolverLifecycleStateDeleting,
	"FAILED":   ResolverLifecycleStateFailed,
	"UPDATING": ResolverLifecycleStateUpdating,
}

var mappingResolverLifecycleStateEnumLowerCase = map[string]ResolverLifecycleStateEnum{
	"active":   ResolverLifecycleStateActive,
	"creating": ResolverLifecycleStateCreating,
	"deleted":  ResolverLifecycleStateDeleted,
	"deleting": ResolverLifecycleStateDeleting,
	"failed":   ResolverLifecycleStateFailed,
	"updating": ResolverLifecycleStateUpdating,
}

// GetResolverLifecycleStateEnumValues Enumerates the set of values for ResolverLifecycleStateEnum
func GetResolverLifecycleStateEnumValues() []ResolverLifecycleStateEnum {
	values := make([]ResolverLifecycleStateEnum, 0)
	for _, v := range mappingResolverLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResolverLifecycleStateEnumStringValues Enumerates the set of values in String for ResolverLifecycleStateEnum
func GetResolverLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingResolverLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResolverLifecycleStateEnum(val string) (ResolverLifecycleStateEnum, bool) {
	enum, ok := mappingResolverLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
