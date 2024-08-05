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

// CreateResolverVnicEndpointDetails The body for defining a new resolver VNIC endpoint. Either isForwarding or isListening must be true, but not both.
// If isListening is true, a listeningAddress may be provided. If isForwarding is true, a forwardingAddress
// may be provided. When not provided, an address will be chosen automatically.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateResolverVnicEndpointDetails struct {

	// The name of the resolver endpoint. Must be unique, case-insensitive, within the resolver.
	Name *string `mandatory:"true" json:"name"`

	// A Boolean flag indicating whether or not the resolver endpoint is for forwarding.
	IsForwarding *bool `mandatory:"true" json:"isForwarding"`

	// A Boolean flag indicating whether or not the resolver endpoint is for listening.
	IsListening *bool `mandatory:"true" json:"isListening"`

	// The OCID of a subnet. Must be part of the VCN that the resolver is attached to.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// An IP address from which forwarded queries may be sent. For VNIC endpoints, this IP address must be part
	// of the subnet and will be assigned by the system if unspecified when isForwarding is true.
	ForwardingAddress *string `mandatory:"false" json:"forwardingAddress"`

	// An IP address to listen to queries on. For VNIC endpoints this IP address must be part of the
	// subnet and will be assigned by the system if unspecified when isListening is true.
	ListeningAddress *string `mandatory:"false" json:"listeningAddress"`

	// An array of network security group OCIDs for the resolver endpoint. These must be part of the VCN that the
	// resolver endpoint is a part of.
	NsgIds []string `mandatory:"false" json:"nsgIds"`
}

// GetName returns Name
func (m CreateResolverVnicEndpointDetails) GetName() *string {
	return m.Name
}

// GetForwardingAddress returns ForwardingAddress
func (m CreateResolverVnicEndpointDetails) GetForwardingAddress() *string {
	return m.ForwardingAddress
}

// GetIsForwarding returns IsForwarding
func (m CreateResolverVnicEndpointDetails) GetIsForwarding() *bool {
	return m.IsForwarding
}

// GetIsListening returns IsListening
func (m CreateResolverVnicEndpointDetails) GetIsListening() *bool {
	return m.IsListening
}

// GetListeningAddress returns ListeningAddress
func (m CreateResolverVnicEndpointDetails) GetListeningAddress() *string {
	return m.ListeningAddress
}

func (m CreateResolverVnicEndpointDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateResolverVnicEndpointDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateResolverVnicEndpointDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateResolverVnicEndpointDetails CreateResolverVnicEndpointDetails
	s := struct {
		DiscriminatorParam string `json:"endpointType"`
		MarshalTypeCreateResolverVnicEndpointDetails
	}{
		"VNIC",
		(MarshalTypeCreateResolverVnicEndpointDetails)(m),
	}

	return json.Marshal(&s)
}
