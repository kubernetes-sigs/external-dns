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

// UpdateResolverEndpointDetails The body for updating an existing resolver endpoint.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateResolverEndpointDetails interface {
}

type updateresolverendpointdetails struct {
	JsonData     []byte
	EndpointType string `json:"endpointType"`
}

// UnmarshalJSON unmarshals json
func (m *updateresolverendpointdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdateresolverendpointdetails updateresolverendpointdetails
	s := struct {
		Model Unmarshalerupdateresolverendpointdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.EndpointType = s.Model.EndpointType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updateresolverendpointdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EndpointType {
	case "VNIC":
		mm := UpdateResolverVnicEndpointDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for UpdateResolverEndpointDetails: %s.", m.EndpointType)
		return *m, nil
	}
}

func (m updateresolverendpointdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updateresolverendpointdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateResolverEndpointDetailsEndpointTypeEnum Enum with underlying type: string
type UpdateResolverEndpointDetailsEndpointTypeEnum string

// Set of constants representing the allowable values for UpdateResolverEndpointDetailsEndpointTypeEnum
const (
	UpdateResolverEndpointDetailsEndpointTypeVnic UpdateResolverEndpointDetailsEndpointTypeEnum = "VNIC"
)

var mappingUpdateResolverEndpointDetailsEndpointTypeEnum = map[string]UpdateResolverEndpointDetailsEndpointTypeEnum{
	"VNIC": UpdateResolverEndpointDetailsEndpointTypeVnic,
}

var mappingUpdateResolverEndpointDetailsEndpointTypeEnumLowerCase = map[string]UpdateResolverEndpointDetailsEndpointTypeEnum{
	"vnic": UpdateResolverEndpointDetailsEndpointTypeVnic,
}

// GetUpdateResolverEndpointDetailsEndpointTypeEnumValues Enumerates the set of values for UpdateResolverEndpointDetailsEndpointTypeEnum
func GetUpdateResolverEndpointDetailsEndpointTypeEnumValues() []UpdateResolverEndpointDetailsEndpointTypeEnum {
	values := make([]UpdateResolverEndpointDetailsEndpointTypeEnum, 0)
	for _, v := range mappingUpdateResolverEndpointDetailsEndpointTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateResolverEndpointDetailsEndpointTypeEnumStringValues Enumerates the set of values in String for UpdateResolverEndpointDetailsEndpointTypeEnum
func GetUpdateResolverEndpointDetailsEndpointTypeEnumStringValues() []string {
	return []string{
		"VNIC",
	}
}

// GetMappingUpdateResolverEndpointDetailsEndpointTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateResolverEndpointDetailsEndpointTypeEnum(val string) (UpdateResolverEndpointDetailsEndpointTypeEnum, bool) {
	enum, ok := mappingUpdateResolverEndpointDetailsEndpointTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
