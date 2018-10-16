// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
//
// Example code for Identity and Access Management Service API
//

package example

import (
	"context"
	"fmt"
	"log"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/identity"
)

// ExampleListAvailabilityDomains Lists the Availability Domains in your tenancy.
// Specify the OCID of either the tenancy or another of your compartments as
// the value for the compartment ID (remember that the tenancy is simply the root compartment).
func ExampleListAvailabilityDomains() {
	c, err := identity.NewIdentityClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(err)

	// The OCID of the tenancy containing the compartment.
	tenancyID, err := common.DefaultConfigProvider().TenancyOCID()
	helpers.FatalIfError(err)

	request := identity.ListAvailabilityDomainsRequest{
		CompartmentId: &tenancyID,
	}

	r, err := c.ListAvailabilityDomains(context.Background(), request)
	helpers.FatalIfError(err)

	log.Printf("list of available domains: %v", r.Items)
	fmt.Println("list available domains completed")

	// Output:
	// list available domains completed
}
