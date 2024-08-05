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

// ResolverRule A rule for a resolver. Specifying both qnameCoverConditions and clientAddressConditions is not allowed.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type ResolverRule interface {

	// A list of CIDR blocks. The query must come from a client within one of the blocks in order for the rule action
	// to apply.
	GetClientAddressConditions() []string

	// A list of domain names. The query must be covered by one of the domains in order for the rule action to apply.
	GetQnameCoverConditions() []string
}

type resolverrule struct {
	JsonData                []byte
	ClientAddressConditions []string `mandatory:"true" json:"clientAddressConditions"`
	QnameCoverConditions    []string `mandatory:"true" json:"qnameCoverConditions"`
	Action                  string   `json:"action"`
}

// UnmarshalJSON unmarshals json
func (m *resolverrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerresolverrule resolverrule
	s := struct {
		Model Unmarshalerresolverrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ClientAddressConditions = s.Model.ClientAddressConditions
	m.QnameCoverConditions = s.Model.QnameCoverConditions
	m.Action = s.Model.Action

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *resolverrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Action {
	case "FORWARD":
		mm := ResolverForwardRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ResolverRule: %s.", m.Action)
		return *m, nil
	}
}

// GetClientAddressConditions returns ClientAddressConditions
func (m resolverrule) GetClientAddressConditions() []string {
	return m.ClientAddressConditions
}

// GetQnameCoverConditions returns QnameCoverConditions
func (m resolverrule) GetQnameCoverConditions() []string {
	return m.QnameCoverConditions
}

func (m resolverrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m resolverrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResolverRuleActionEnum Enum with underlying type: string
type ResolverRuleActionEnum string

// Set of constants representing the allowable values for ResolverRuleActionEnum
const (
	ResolverRuleActionForward ResolverRuleActionEnum = "FORWARD"
)

var mappingResolverRuleActionEnum = map[string]ResolverRuleActionEnum{
	"FORWARD": ResolverRuleActionForward,
}

var mappingResolverRuleActionEnumLowerCase = map[string]ResolverRuleActionEnum{
	"forward": ResolverRuleActionForward,
}

// GetResolverRuleActionEnumValues Enumerates the set of values for ResolverRuleActionEnum
func GetResolverRuleActionEnumValues() []ResolverRuleActionEnum {
	values := make([]ResolverRuleActionEnum, 0)
	for _, v := range mappingResolverRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetResolverRuleActionEnumStringValues Enumerates the set of values in String for ResolverRuleActionEnum
func GetResolverRuleActionEnumStringValues() []string {
	return []string{
		"FORWARD",
	}
}

// GetMappingResolverRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResolverRuleActionEnum(val string) (ResolverRuleActionEnum, bool) {
	enum, ok := mappingResolverRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
