package v2

import (
	"context"
	"errors"
	"net"
	"strings"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// SecurityGroupRule represents a Security Group rule.
type SecurityGroupRule struct {
	Description       *string
	EndPort           *uint16
	FlowDirection     *string `req-for:"create"`
	ICMPCode          *int64
	ICMPType          *int64
	ID                *string `req-for:"delete"`
	Network           *net.IPNet
	Protocol          *string `req-for:"create"`
	SecurityGroupID   *string
	SecurityGroupName *string
	Visibility        *string
	StartPort         *uint16
}

func securityGroupRuleFromAPI(r *oapi.SecurityGroupRule) *SecurityGroupRule {
	return &SecurityGroupRule{
		Description: r.Description,
		EndPort: func() (v *uint16) {
			if r.EndPort != nil {
				port := uint16(*r.EndPort)
				v = &port
			}
			return
		}(),
		FlowDirection: (*string)(r.FlowDirection),
		ICMPCode: func() (v *int64) {
			if r.Icmp != nil {
				v = r.Icmp.Code
			}
			return
		}(),
		ICMPType: func() (v *int64) {
			if r.Icmp != nil {
				v = r.Icmp.Type
			}
			return
		}(),
		ID: r.Id,
		Network: func() (v *net.IPNet) {
			if r.Network != nil {
				_, v, _ = net.ParseCIDR(*r.Network)
			}
			return
		}(),
		Protocol: (*string)(r.Protocol),
		SecurityGroupID: func() (v *string) {
			if r.SecurityGroup != nil &&
				(r.SecurityGroup.Visibility == nil ||
					*r.SecurityGroup.Visibility != "public") {
				v = r.SecurityGroup.Id
			}
			return
		}(),
		SecurityGroupName: func() (v *string) {
			if r.SecurityGroup != nil && r.SecurityGroup.Visibility != nil &&
				*r.SecurityGroup.Visibility == "public" && r.SecurityGroup.Name != nil {
				v = r.SecurityGroup.Name
			}
			return
		}(),
		StartPort: func() (v *uint16) {
			if r.StartPort != nil {
				port := uint16(*r.StartPort)
				v = &port
			}
			return
		}(),
		Visibility: func() (v *string) {
			if r.SecurityGroup != nil && r.SecurityGroup.Visibility != nil {
				v = (*string)(r.SecurityGroup.Visibility)
			}
			return
		}(),
	}
}

// CreateSecurityGroupRule creates a Security Group rule.
func (c *Client) CreateSecurityGroupRule(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
	rule *SecurityGroupRule,
) (*SecurityGroupRule, error) {
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return nil, err
	}
	if err := validateOperationParams(rule, "create"); err != nil {
		return nil, err
	}

	var icmp *struct {
		Code *int64 `json:"code,omitempty"`
		Type *int64 `json:"type,omitempty"`
	}

	if strings.HasPrefix(*rule.Protocol, "icmp") {
		icmp = &struct {
			Code *int64 `json:"code,omitempty"`
			Type *int64 `json:"type,omitempty"`
		}{
			Code: rule.ICMPCode,
			Type: rule.ICMPType,
		}
	}

	// The API doesn't return the Security Group rule created directly, so in order to
	// return a *SecurityGroupRule corresponding to the new rule we have to manually
	// compare the list of rules in the SG before and after the rule creation, and
	// identify the rule that wasn't there before.
	// Note: in case of multiple rules creation in parallel this technique is subject
	// to race condition as we could return an unrelated rule. To prevent this, we
	// also compare the properties of the new rule to the ones specified in the input
	// rule parameter.
	sgCurrent, err := c.GetSecurityGroup(ctx, zone, *securityGroup.ID)
	if err != nil {
		return nil, err
	}

	currentRules := make(map[string]struct{})
	for _, r := range sgCurrent.Rules {
		currentRules[*r.ID] = struct{}{}
	}

	resp, err := c.AddRuleToSecurityGroupWithResponse(
		apiv2.WithZone(ctx, zone),
		*securityGroup.ID,
		oapi.AddRuleToSecurityGroupJSONRequestBody{
			Description: rule.Description,
			EndPort: func() (v *int64) {
				if rule.EndPort != nil {
					port := int64(*rule.EndPort)
					v = &port
				}
				return
			}(),
			FlowDirection: oapi.AddRuleToSecurityGroupJSONBodyFlowDirection(*rule.FlowDirection),
			Icmp:          icmp,
			Network: func() (v *string) {
				if rule.Network != nil {
					ip := rule.Network.String()
					v = &ip
				}
				return
			}(),
			Protocol: oapi.AddRuleToSecurityGroupJSONBodyProtocol(*rule.Protocol),
			SecurityGroup: func() (v *oapi.SecurityGroupResource) {
				if rule.SecurityGroupID != nil {
					v = &oapi.SecurityGroupResource{Id: rule.SecurityGroupID}
				}
				if rule.SecurityGroupName != nil {
					visibility := oapi.SecurityGroupResourceVisibilityPublic
					v = &oapi.SecurityGroupResource{Name: rule.SecurityGroupName, Visibility: &visibility}
				}
				return
			}(),
			StartPort: func() (v *int64) {
				if rule.StartPort != nil {
					port := int64(*rule.StartPort)
					v = &port
				}
				return
			}(),
		})
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	sgUpdated, err := c.GetSecurityGroup(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
	if err != nil {
		return nil, err
	}

	// Look for an unknown rule which properties match the one we've just created.
	for _, r := range sgUpdated.Rules {
		if _, ok := currentRules[*r.ID]; !ok {
			if *r.FlowDirection == *rule.FlowDirection && *r.Protocol == *rule.Protocol {
				if rule.Description != nil && r.Description != nil && *r.Description != *rule.Description {
					continue
				}

				if rule.StartPort != nil && r.StartPort != nil && *r.StartPort != *rule.StartPort {
					continue
				}

				if rule.EndPort != nil && r.EndPort != nil && *r.EndPort != *rule.EndPort {
					continue
				}

				if rule.Network != nil && r.Network != nil && r.Network.String() != rule.Network.String() {
					continue
				}

				if rule.SecurityGroupID != nil && r.SecurityGroupID != nil &&
					*r.SecurityGroupID != *rule.SecurityGroupID {
					continue
				}

				if rule.ICMPType != nil && r.ICMPType != nil && *r.ICMPType != *rule.ICMPType {
					continue
				}

				if rule.ICMPCode != nil && r.ICMPCode != nil && *r.ICMPCode != *rule.ICMPCode {
					continue
				}

				return r, nil
			}
		}
	}

	return nil, errors.New("unable to identify the rule created")
}

// DeleteSecurityGroupRule deletes a Security Group rule.
func (c *Client) DeleteSecurityGroupRule(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
	rule *SecurityGroupRule,
) error {
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return err
	}
	if err := validateOperationParams(rule, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteRuleFromSecurityGroupWithResponse(apiv2.WithZone(ctx, zone), *securityGroup.ID, *rule.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
