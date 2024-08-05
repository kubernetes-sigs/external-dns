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

// ResolverVnicEndpointSummary An OCI DNS resolver VNIC endpoint. A VNIC is created for each ResolverVnicEndpoint. VCNs and subnets cannot be
// deleted while ResolverVnicEndpoints exist in them due to the VNIC.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ResolverVnicEndpointSummary struct {

	// The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver.
	Name *string `mandatory:"true" json:"name"`

	// A Boolean flag indicating whether or not the resolver endpoint is for forwarding.
	IsForwarding *bool `mandatory:"true" json:"isForwarding"`

	// A Boolean flag indicating whether or not the resolver endpoint is for listening.
	IsListening *bool `mandatory:"true" json:"isListening"`

	// The OCID of the owning compartment. This will match the resolver that the resolver endpoint is under
	// and will be updated if the resolver's compartment is changed.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the resource was created in "YYYY-MM-ddThh:mm:ssZ" format
	// with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was last updated in "YYYY-MM-ddThh:mm:ssZ"
	// format with a Z offset, as defined by RFC 3339.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// The OCID of a subnet. Must be part of the VCN that the resolver is attached to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part
	// of the subnet and will be assigned by the system if unspecified when isForwarding is true.
	ForwardingAddress *string `mandatory:"false" json:"forwardingAddress"`

	// An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the
	// subnet and will be assigned by the system if unspecified when isListening is true.
	ListeningAddress *string `mandatory:"false" json:"listeningAddress"`

	// The current state of the resource.
	LifecycleState ResolverEndpointSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetName returns Name
func (m ResolverVnicEndpointSummary) GetName() *string {
	return m.Name
}

// GetForwardingAddress returns ForwardingAddress
func (m ResolverVnicEndpointSummary) GetForwardingAddress() *string {
	return m.ForwardingAddress
}

// GetIsForwarding returns IsForwarding
func (m ResolverVnicEndpointSummary) GetIsForwarding() *bool {
	return m.IsForwarding
}

// GetIsListening returns IsListening
func (m ResolverVnicEndpointSummary) GetIsListening() *bool {
	return m.IsListening
}

// GetListeningAddress returns ListeningAddress
func (m ResolverVnicEndpointSummary) GetListeningAddress() *string {
	return m.ListeningAddress
}

// GetCompartmentId returns CompartmentId
func (m ResolverVnicEndpointSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m ResolverVnicEndpointSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ResolverVnicEndpointSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m ResolverVnicEndpointSummary) GetLifecycleState() ResolverEndpointSummaryLifecycleStateEnum {
	return m.LifecycleState
}

// GetSelf returns Self
func (m ResolverVnicEndpointSummary) GetSelf() *string {
	return m.Self
}

func (m ResolverVnicEndpointSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResolverVnicEndpointSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResolverEndpointSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetResolverEndpointSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ResolverVnicEndpointSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeResolverVnicEndpointSummary ResolverVnicEndpointSummary
	s := struct {
		DiscriminatorParam string `json:"endpointType"`
		MarshalTypeResolverVnicEndpointSummary
	}{
		"VNIC",
		(MarshalTypeResolverVnicEndpointSummary)(m),
	}

	return json.Marshal(&s)
}
