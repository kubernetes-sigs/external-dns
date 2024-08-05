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

// ResolverEndpoint An OCI DNS resolver endpoint.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ResolverEndpoint interface {

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
	GetLifecycleState() ResolverEndpointLifecycleStateEnum

	// The canonical absolute URL of the resource.
	GetSelf() *string

	// An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part
	// of the subnet and will be assigned by the system if unspecified when isForwarding is true.
	GetForwardingAddress() *string

	// An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the
	// subnet and will be assigned by the system if unspecified when isListening is true.
	GetListeningAddress() *string
}

type resolverendpoint struct {
	JsonData          []byte
	ForwardingAddress *string                            `mandatory:"false" json:"forwardingAddress"`
	ListeningAddress  *string                            `mandatory:"false" json:"listeningAddress"`
	Name              *string                            `mandatory:"true" json:"name"`
	IsForwarding      *bool                              `mandatory:"true" json:"isForwarding"`
	IsListening       *bool                              `mandatory:"true" json:"isListening"`
	CompartmentId     *string                            `mandatory:"true" json:"compartmentId"`
	TimeCreated       *common.SDKTime                    `mandatory:"true" json:"timeCreated"`
	TimeUpdated       *common.SDKTime                    `mandatory:"true" json:"timeUpdated"`
	LifecycleState    ResolverEndpointLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
	Self              *string                            `mandatory:"true" json:"self"`
	EndpointType      string                             `json:"endpointType"`
}

// UnmarshalJSON unmarshals json
func (m *resolverendpoint) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresolverendpoint resolverendpoint
	s := struct {
		Model Unmarshalerresolverendpoint
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
func (m *resolverendpoint) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EndpointType {
	case "VNIC":
		mm := ResolverVnicEndpoint{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResolverEndpoint: %s.", m.EndpointType)
		return *m, nil
	}
}

// GetForwardingAddress returns ForwardingAddress
func (m resolverendpoint) GetForwardingAddress() *string {
	return m.ForwardingAddress
}

// GetListeningAddress returns ListeningAddress
func (m resolverendpoint) GetListeningAddress() *string {
	return m.ListeningAddress
}

// GetName returns Name
func (m resolverendpoint) GetName() *string {
	return m.Name
}

// GetIsForwarding returns IsForwarding
func (m resolverendpoint) GetIsForwarding() *bool {
	return m.IsForwarding
}

// GetIsListening returns IsListening
func (m resolverendpoint) GetIsListening() *bool {
	return m.IsListening
}

// GetCompartmentId returns CompartmentId
func (m resolverendpoint) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m resolverendpoint) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m resolverendpoint) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m resolverendpoint) GetLifecycleState() ResolverEndpointLifecycleStateEnum {
	return m.LifecycleState
}

// GetSelf returns Self
func (m resolverendpoint) GetSelf() *string {
	return m.Self
}

func (m resolverendpoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resolverendpoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResolverEndpointLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResolverEndpointLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResolverEndpointLifecycleStateEnum Enum with underlying type: string
type ResolverEndpointLifecycleStateEnum string

// Set of constants representing the allowable values for ResolverEndpointLifecycleStateEnum
const (
	ResolverEndpointLifecycleStateActive   ResolverEndpointLifecycleStateEnum = "ACTIVE"
	ResolverEndpointLifecycleStateCreating ResolverEndpointLifecycleStateEnum = "CREATING"
	ResolverEndpointLifecycleStateDeleted  ResolverEndpointLifecycleStateEnum = "DELETED"
	ResolverEndpointLifecycleStateDeleting ResolverEndpointLifecycleStateEnum = "DELETING"
	ResolverEndpointLifecycleStateFailed   ResolverEndpointLifecycleStateEnum = "FAILED"
	ResolverEndpointLifecycleStateUpdating ResolverEndpointLifecycleStateEnum = "UPDATING"
)

var mappingResolverEndpointLifecycleStateEnum = map[string]ResolverEndpointLifecycleStateEnum{
	"ACTIVE":   ResolverEndpointLifecycleStateActive,
	"CREATING": ResolverEndpointLifecycleStateCreating,
	"DELETED":  ResolverEndpointLifecycleStateDeleted,
	"DELETING": ResolverEndpointLifecycleStateDeleting,
	"FAILED":   ResolverEndpointLifecycleStateFailed,
	"UPDATING": ResolverEndpointLifecycleStateUpdating,
}

var mappingResolverEndpointLifecycleStateEnumLowerCase = map[string]ResolverEndpointLifecycleStateEnum{
	"active":   ResolverEndpointLifecycleStateActive,
	"creating": ResolverEndpointLifecycleStateCreating,
	"deleted":  ResolverEndpointLifecycleStateDeleted,
	"deleting": ResolverEndpointLifecycleStateDeleting,
	"failed":   ResolverEndpointLifecycleStateFailed,
	"updating": ResolverEndpointLifecycleStateUpdating,
}

// GetResolverEndpointLifecycleStateEnumValues Enumerates the set of values for ResolverEndpointLifecycleStateEnum
func GetResolverEndpointLifecycleStateEnumValues() []ResolverEndpointLifecycleStateEnum {
	values := make([]ResolverEndpointLifecycleStateEnum, 0)
	for _, v := range mappingResolverEndpointLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetResolverEndpointLifecycleStateEnumStringValues Enumerates the set of values in String for ResolverEndpointLifecycleStateEnum
func GetResolverEndpointLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingResolverEndpointLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResolverEndpointLifecycleStateEnum(val string) (ResolverEndpointLifecycleStateEnum, bool) {
	enum, ok := mappingResolverEndpointLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ResolverEndpointEndpointTypeEnum Enum with underlying type: string
type ResolverEndpointEndpointTypeEnum string

// Set of constants representing the allowable values for ResolverEndpointEndpointTypeEnum
const (
	ResolverEndpointEndpointTypeVnic ResolverEndpointEndpointTypeEnum = "VNIC"
)

var mappingResolverEndpointEndpointTypeEnum = map[string]ResolverEndpointEndpointTypeEnum{
	"VNIC": ResolverEndpointEndpointTypeVnic,
}

var mappingResolverEndpointEndpointTypeEnumLowerCase = map[string]ResolverEndpointEndpointTypeEnum{
	"vnic": ResolverEndpointEndpointTypeVnic,
}

// GetResolverEndpointEndpointTypeEnumValues Enumerates the set of values for ResolverEndpointEndpointTypeEnum
func GetResolverEndpointEndpointTypeEnumValues() []ResolverEndpointEndpointTypeEnum {
	values := make([]ResolverEndpointEndpointTypeEnum, 0)
	for _, v := range mappingResolverEndpointEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResolverEndpointEndpointTypeEnumStringValues Enumerates the set of values in String for ResolverEndpointEndpointTypeEnum
func GetResolverEndpointEndpointTypeEnumStringValues() []string {
	return []string{
		"VNIC",
	}
}

// GetMappingResolverEndpointEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResolverEndpointEndpointTypeEnum(val string) (ResolverEndpointEndpointTypeEnum, bool) {
	enum, ok := mappingResolverEndpointEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
