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

// SteeringPolicySummary A DNS steering policy.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type SteeringPolicySummary struct {

	// The OCID of the compartment containing the steering policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name for the steering policy. Does not have to be unique and can be changed.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The Time To Live (TTL) for responses from the steering policy, in seconds.
	// If not specified during creation, a value of 30 seconds will be used.
	Ttl *int `mandatory:"true" json:"ttl"`

	// A set of predefined rules based on the desired purpose of the steering policy. Each
	// template utilizes Traffic Management's rules in a different order to produce the desired
	// results when answering DNS queries.
	//
	// **Example:** The `FAILOVER` template determines answers by filtering the policy's answers
	// using the `FILTER` rule first, then the following rules in succession: `HEALTH`, `PRIORITY`,
	// and `LIMIT`. This gives the domain dynamic failover capability.
	//
	// It is **strongly recommended** to use a template other than `CUSTOM` when creating
	// a steering policy.
	//
	// All templates require the rule order to begin with an unconditional `FILTER` rule that keeps
	// answers contingent upon `answer.isDisabled != true`, except for `CUSTOM`. A defined
	// `HEALTH` rule must follow the `FILTER` rule if the policy references a `healthCheckMonitorId`.
	// The last rule of a template must must be a `LIMIT` rule. For more information about templates
	// and code examples, see Traffic Management API Guide (https://docs.cloud.oracle.com/iaas/Content/TrafficManagement/Concepts/trafficmanagementapi.htm).
	// **Template Types**
	// * `FAILOVER` - Uses health check information on your endpoints to determine which DNS answers
	// to serve. If an endpoint fails a health check, the answer for that endpoint will be removed
	// from the list of available answers until the endpoint is detected as healthy.
	//
	// * `LOAD_BALANCE` - Distributes web traffic to specified endpoints based on defined weights.
	//
	// * `ROUTE_BY_GEO` - Answers DNS queries based on the query's geographic location. For a list of geographic
	// locations to route by, see Traffic Management Geographic Locations (https://docs.cloud.oracle.com/iaas/Content/TrafficManagement/Reference/trafficmanagementgeo.htm).
	//
	// * `ROUTE_BY_ASN` - Answers DNS queries based on the query's originating ASN.
	//
	// * `ROUTE_BY_IP` - Answers DNS queries based on the query's IP address.
	//
	// * `CUSTOM` - Allows a customized configuration of rules.
	Template SteeringPolicySummaryTemplateEnum `mandatory:"true" json:"template"`

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

	// The canonical absolute URL of the resource.
	Self *string `mandatory:"true" json:"self"`

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the resource was created, expressed in RFC 3339 timestamp format.
	// **Example:** `2016-07-22T17:23:59:60Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the resource.
	LifecycleState SteeringPolicySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the health check monitor providing health data about the answers of the
	// steering policy. A steering policy answer with `rdata` matching a monitored endpoint
	// will use the health data of that endpoint. A steering policy answer with `rdata` not
	// matching any monitored endpoint will be assumed healthy.
	//
	// **Note:** To use the Health Check monitoring feature in a steering policy, a monitor
	// must be created using the Health Checks service first. For more information on how to
	// create a monitor, please see Managing Health Checks (https://docs.cloud.oracle.com/iaas/Content/HealthChecks/Tasks/managinghealthchecks.htm).
	HealthCheckMonitorId *string `mandatory:"false" json:"healthCheckMonitorId"`
}

func (m SteeringPolicySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SteeringPolicySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSteeringPolicySummaryTemplateEnum(string(m.Template)); !ok && m.Template != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Template: %s. Supported values are: %s.", m.Template, strings.Join(GetSteeringPolicySummaryTemplateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSteeringPolicySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSteeringPolicySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SteeringPolicySummaryTemplateEnum Enum with underlying type: string
type SteeringPolicySummaryTemplateEnum string

// Set of constants representing the allowable values for SteeringPolicySummaryTemplateEnum
const (
	SteeringPolicySummaryTemplateFailover    SteeringPolicySummaryTemplateEnum = "FAILOVER"
	SteeringPolicySummaryTemplateLoadBalance SteeringPolicySummaryTemplateEnum = "LOAD_BALANCE"
	SteeringPolicySummaryTemplateRouteByGeo  SteeringPolicySummaryTemplateEnum = "ROUTE_BY_GEO"
	SteeringPolicySummaryTemplateRouteByAsn  SteeringPolicySummaryTemplateEnum = "ROUTE_BY_ASN"
	SteeringPolicySummaryTemplateRouteByIp   SteeringPolicySummaryTemplateEnum = "ROUTE_BY_IP"
	SteeringPolicySummaryTemplateCustom      SteeringPolicySummaryTemplateEnum = "CUSTOM"
)

var mappingSteeringPolicySummaryTemplateEnum = map[string]SteeringPolicySummaryTemplateEnum{
	"FAILOVER":     SteeringPolicySummaryTemplateFailover,
	"LOAD_BALANCE": SteeringPolicySummaryTemplateLoadBalance,
	"ROUTE_BY_GEO": SteeringPolicySummaryTemplateRouteByGeo,
	"ROUTE_BY_ASN": SteeringPolicySummaryTemplateRouteByAsn,
	"ROUTE_BY_IP":  SteeringPolicySummaryTemplateRouteByIp,
	"CUSTOM":       SteeringPolicySummaryTemplateCustom,
}

var mappingSteeringPolicySummaryTemplateEnumLowerCase = map[string]SteeringPolicySummaryTemplateEnum{
	"failover":     SteeringPolicySummaryTemplateFailover,
	"load_balance": SteeringPolicySummaryTemplateLoadBalance,
	"route_by_geo": SteeringPolicySummaryTemplateRouteByGeo,
	"route_by_asn": SteeringPolicySummaryTemplateRouteByAsn,
	"route_by_ip":  SteeringPolicySummaryTemplateRouteByIp,
	"custom":       SteeringPolicySummaryTemplateCustom,
}

// GetSteeringPolicySummaryTemplateEnumValues Enumerates the set of values for SteeringPolicySummaryTemplateEnum
func GetSteeringPolicySummaryTemplateEnumValues() []SteeringPolicySummaryTemplateEnum {
	values := make([]SteeringPolicySummaryTemplateEnum, 0)
	for _, v := range mappingSteeringPolicySummaryTemplateEnum {
		values = append(values, v)
	}
	return values
}

// GetSteeringPolicySummaryTemplateEnumStringValues Enumerates the set of values in String for SteeringPolicySummaryTemplateEnum
func GetSteeringPolicySummaryTemplateEnumStringValues() []string {
	return []string{
		"FAILOVER",
		"LOAD_BALANCE",
		"ROUTE_BY_GEO",
		"ROUTE_BY_ASN",
		"ROUTE_BY_IP",
		"CUSTOM",
	}
}

// GetMappingSteeringPolicySummaryTemplateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSteeringPolicySummaryTemplateEnum(val string) (SteeringPolicySummaryTemplateEnum, bool) {
	enum, ok := mappingSteeringPolicySummaryTemplateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SteeringPolicySummaryLifecycleStateEnum Enum with underlying type: string
type SteeringPolicySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SteeringPolicySummaryLifecycleStateEnum
const (
	SteeringPolicySummaryLifecycleStateActive   SteeringPolicySummaryLifecycleStateEnum = "ACTIVE"
	SteeringPolicySummaryLifecycleStateCreating SteeringPolicySummaryLifecycleStateEnum = "CREATING"
	SteeringPolicySummaryLifecycleStateDeleted  SteeringPolicySummaryLifecycleStateEnum = "DELETED"
	SteeringPolicySummaryLifecycleStateDeleting SteeringPolicySummaryLifecycleStateEnum = "DELETING"
)

var mappingSteeringPolicySummaryLifecycleStateEnum = map[string]SteeringPolicySummaryLifecycleStateEnum{
	"ACTIVE":   SteeringPolicySummaryLifecycleStateActive,
	"CREATING": SteeringPolicySummaryLifecycleStateCreating,
	"DELETED":  SteeringPolicySummaryLifecycleStateDeleted,
	"DELETING": SteeringPolicySummaryLifecycleStateDeleting,
}

var mappingSteeringPolicySummaryLifecycleStateEnumLowerCase = map[string]SteeringPolicySummaryLifecycleStateEnum{
	"active":   SteeringPolicySummaryLifecycleStateActive,
	"creating": SteeringPolicySummaryLifecycleStateCreating,
	"deleted":  SteeringPolicySummaryLifecycleStateDeleted,
	"deleting": SteeringPolicySummaryLifecycleStateDeleting,
}

// GetSteeringPolicySummaryLifecycleStateEnumValues Enumerates the set of values for SteeringPolicySummaryLifecycleStateEnum
func GetSteeringPolicySummaryLifecycleStateEnumValues() []SteeringPolicySummaryLifecycleStateEnum {
	values := make([]SteeringPolicySummaryLifecycleStateEnum, 0)
	for _, v := range mappingSteeringPolicySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSteeringPolicySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SteeringPolicySummaryLifecycleStateEnum
func GetSteeringPolicySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"DELETING",
	}
}

// GetMappingSteeringPolicySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSteeringPolicySummaryLifecycleStateEnum(val string) (SteeringPolicySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSteeringPolicySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
