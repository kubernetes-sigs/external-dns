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

// SteeringPolicyWeightedRuleCase The representation of SteeringPolicyWeightedRuleCase
type SteeringPolicyWeightedRuleCase struct {

	// An expression that uses conditions at the time of a DNS query to indicate
	// whether a case matches. Conditions may include the geographical location, IP
	// subnet, or ASN the DNS query originated. **Example:** If you have an
	// office that uses the subnet `192.0.2.0/24` you could use a `caseCondition`
	// expression `query.client.address in ('192.0.2.0/24')` to define a case that
	// matches queries from that office.
	CaseCondition *string `mandatory:"false" json:"caseCondition"`

	// An array of `SteeringPolicyWeightedAnswerData` objects.
	AnswerData []SteeringPolicyWeightedAnswerData `mandatory:"false" json:"answerData"`
}

func (m SteeringPolicyWeightedRuleCase) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SteeringPolicyWeightedRuleCase) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
