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

// ResolverEndpointSummary An OCI DNS resolver endpoint.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ResolverEndpointSummary interface {

	// The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver.
	GetName() *string

	// A Boolean flag indicating whether or not the resolver endpoint is for forwarding.
	GetIsForwarding() *bool

	// A Boolean flag indicating whether or not the resolver endpoint is for listening.
	GetIsListening() *bool

	// The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under
	// and will be updated if the resolver's compartment is changed.
	GetCompartmentId() *string

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	GetTimeCreated() *common.SDKTime

	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ"
	// format with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	GetTimeUpdated() *common.SDKTime

	// The current state of the resource.
	GetLifecycleState() ResolverEndpointSummaryLifecycleStateEnum

	// The canonical absolute URL of the resource.
	GetSelf() *string

	// An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part
	// of the subnet and will be assigned by the system if unspecified when isForwarding is true.
	GetForwardingAddress() *string

	// An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the
	// subnet and will be assigned by the system if unspecified when isListening is true.
	GetListeningAddress() *string
}

type resolverendpointsummary struct {
	JsonData          []byte
	ForwardingAddress *string                                   `mandatory:"false" json:"forwardingAddress"`
	ListeningAddress  *string                                   `mandatory:"false" json:"listeningAddress"`
	Name              *string                                   `mandatory:"true" json:"name"`
	IsForwarding      *bool                                     `mandatory:"true" json:"isForwarding"`
	IsListening       *bool                                     `mandatory:"true" json:"isListening"`
	CompartmentId     *string                                   `mandatory:"true" json:"compartmentId"`
	TimeCreated       *common.SDKTime                           `mandatory:"true" json:"timeCreated"`
	TimeUpdated       *common.SDKTime                           `mandatory:"true" json:"timeUpdated"`
	LifecycleState    ResolverEndpointSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	Self              *string                                   `mandatory:"true" json:"self"`
	EndpointType      string                                    `json:"endpointType"`
}

// UnmarshalJSON unmarshals json
func (m *resolverendpointsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresolverendpointsummary resolverendpointsummary
	s := struct {
		Model Unmarshalerresolverendpointsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.IsForwarding = s.Model.IsForwarding
	m.IsListening = s.Model.IsListening
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleState = s.Model.LifecycleState
	m.Self = s.Model.Self
	m.ForwardingAddress = s.Model.ForwardingAddress
	m.ListeningAddress = s.Model.ListeningAddress
	m.EndpointType = s.Model.EndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resolverendpointsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EndpointType {
	case "VNIC":
		mm := ResolverVnicEndpointSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResolverEndpointSummary: %s.", m.EndpointType)
		return *m, nil
	}
}

// GetForwardingAddress returns ForwardingAddress
func (m resolverendpointsummary) GetForwardingAddress() *string {
	return m.ForwardingAddress
}

// GetListeningAddress returns ListeningAddress
func (m resolverendpointsummary) GetListeningAddress() *string {
	return m.ListeningAddress
}

// GetName returns Name
func (m resolverendpointsummary) GetName() *string {
	return m.Name
}

// GetIsForwarding returns IsForwarding
func (m resolverendpointsummary) GetIsForwarding() *bool {
	return m.IsForwarding
}

// GetIsListening returns IsListening
func (m resolverendpointsummary) GetIsListening() *bool {
	return m.IsListening
}

// GetCompartmentId returns CompartmentId
func (m resolverendpointsummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m resolverendpointsummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m resolverendpointsummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m resolverendpointsummary) GetLifecycleState() ResolverEndpointSummaryLifecycleStateEnum {
	return m.LifecycleState
}

// GetSelf returns Self
func (m resolverendpointsummary) GetSelf() *string {
	return m.Self
}

func (m resolverendpointsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resolverendpointsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResolverEndpointSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResolverEndpointSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResolverEndpointSummaryLifecycleStateEnum Enum with underlying type: string
type ResolverEndpointSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ResolverEndpointSummaryLifecycleStateEnum
const (
	ResolverEndpointSummaryLifecycleStateActive   ResolverEndpointSummaryLifecycleStateEnum = "ACTIVE"
	ResolverEndpointSummaryLifecycleStateCreating ResolverEndpointSummaryLifecycleStateEnum = "CREATING"
	ResolverEndpointSummaryLifecycleStateDeleted  ResolverEndpointSummaryLifecycleStateEnum = "DELETED"
	ResolverEndpointSummaryLifecycleStateDeleting ResolverEndpointSummaryLifecycleStateEnum = "DELETING"
	ResolverEndpointSummaryLifecycleStateFailed   ResolverEndpointSummaryLifecycleStateEnum = "FAILED"
	ResolverEndpointSummaryLifecycleStateUpdating ResolverEndpointSummaryLifecycleStateEnum = "UPDATING"
)

var mappingResolverEndpointSummaryLifecycleStateEnum = map[string]ResolverEndpointSummaryLifecycleStateEnum{
	"ACTIVE":   ResolverEndpointSummaryLifecycleStateActive,
	"CREATING": ResolverEndpointSummaryLifecycleStateCreating,
	"DELETED":  ResolverEndpointSummaryLifecycleStateDeleted,
	"DELETING": ResolverEndpointSummaryLifecycleStateDeleting,
	"FAILED":   ResolverEndpointSummaryLifecycleStateFailed,
	"UPDATING": ResolverEndpointSummaryLifecycleStateUpdating,
}

var mappingResolverEndpointSummaryLifecycleStateEnumLowerCase = map[string]ResolverEndpointSummaryLifecycleStateEnum{
	"active":   ResolverEndpointSummaryLifecycleStateActive,
	"creating": ResolverEndpointSummaryLifecycleStateCreating,
	"deleted":  ResolverEndpointSummaryLifecycleStateDeleted,
	"deleting": ResolverEndpointSummaryLifecycleStateDeleting,
	"failed":   ResolverEndpointSummaryLifecycleStateFailed,
	"updating": ResolverEndpointSummaryLifecycleStateUpdating,
}

// GetResolverEndpointSummaryLifecycleStateEnumValues Enumerates the set of values for ResolverEndpointSummaryLifecycleStateEnum
func GetResolverEndpointSummaryLifecycleStateEnumValues() []ResolverEndpointSummaryLifecycleStateEnum {
	values := make([]ResolverEndpointSummaryLifecycleStateEnum, 0)
	for _, v := range mappingResolverEndpointSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResolverEndpointSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ResolverEndpointSummaryLifecycleStateEnum
func GetResolverEndpointSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingResolverEndpointSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResolverEndpointSummaryLifecycleStateEnum(val string) (ResolverEndpointSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingResolverEndpointSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResolverEndpointSummaryEndpointTypeEnum Enum with underlying type: string
type ResolverEndpointSummaryEndpointTypeEnum string

// Set of constants representing the allowable values for ResolverEndpointSummaryEndpointTypeEnum
const (
	ResolverEndpointSummaryEndpointTypeVnic ResolverEndpointSummaryEndpointTypeEnum = "VNIC"
)

var mappingResolverEndpointSummaryEndpointTypeEnum = map[string]ResolverEndpointSummaryEndpointTypeEnum{
	"VNIC": ResolverEndpointSummaryEndpointTypeVnic,
}

var mappingResolverEndpointSummaryEndpointTypeEnumLowerCase = map[string]ResolverEndpointSummaryEndpointTypeEnum{
	"vnic": ResolverEndpointSummaryEndpointTypeVnic,
}

// GetResolverEndpointSummaryEndpointTypeEnumValues Enumerates the set of values for ResolverEndpointSummaryEndpointTypeEnum
func GetResolverEndpointSummaryEndpointTypeEnumValues() []ResolverEndpointSummaryEndpointTypeEnum {
	values := make([]ResolverEndpointSummaryEndpointTypeEnum, 0)
	for _, v := range mappingResolverEndpointSummaryEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResolverEndpointSummaryEndpointTypeEnumStringValues Enumerates the set of values in String for ResolverEndpointSummaryEndpointTypeEnum
func GetResolverEndpointSummaryEndpointTypeEnumStringValues() []string {
	return []string{
		"VNIC",
	}
}

// GetMappingResolverEndpointSummaryEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResolverEndpointSummaryEndpointTypeEnum(val string) (ResolverEndpointSummaryEndpointTypeEnum, bool) {
	enum, ok := mappingResolverEndpointSummaryEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
