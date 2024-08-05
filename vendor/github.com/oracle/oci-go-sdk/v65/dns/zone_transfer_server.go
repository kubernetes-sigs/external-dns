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

// ZoneTransferServer An OCI nameserver that transfers zone data with external nameservers.
type ZoneTransferServer struct {

	// The server's IP address (IPv4 or IPv6).
	Address *string `mandatory:"true" json:"address"`

	// The server's port.
	Port *int `mandatory:"false" json:"port"`

	// A Boolean flag indicating whether or not the server is a zone data transfer source.
	IsTransferSource *bool `mandatory:"false" json:"isTransferSource"`

	// A Boolean flag indicating whether or not the server is a zone data transfer destination.
	IsTransferDestination *bool `mandatory:"false" json:"isTransferDestination"`
}

func (m ZoneTransferServer) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ZoneTransferServer) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
