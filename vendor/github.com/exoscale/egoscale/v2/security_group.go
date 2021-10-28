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
	Description     *string
	EndPort         *uint16
	FlowDirection   *string `req-for:"create"`
	ICMPCode        *int64
	ICMPType        *int64
	ID              *string `req-for:"delete"`
	Network         *net.IPNet
	Protocol        *string `req-for:"create"`
	SecurityGroupID *string
	StartPort       *uint16
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
			if r.SecurityGroup != nil {
				v = &r.SecurityGroup.Id
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
	}
}

// SecurityGroup represents a Security Group.
type SecurityGroup struct {
	Description     *string
	ID              *string `req-for:"update,delete"`
	Name            *string `req-for:"create"`
	ExternalSources *[]string
	Rules           []*SecurityGroupRule
}

func securityGroupFromAPI(s *oapi.SecurityGroup) *SecurityGroup {
	return &SecurityGroup{
		Description:     s.Description,
		ID:              s.Id,
		Name:            s.Name,
		ExternalSources: s.ExternalSources,
		Rules: func() (rules []*SecurityGroupRule) {
			if s.Rules != nil {
				rules = make([]*SecurityGroupRule, 0)
				for _, rule := range *s.Rules {
					rule := rule
					rules = append(rules, securityGroupRuleFromAPI(&rule))
				}
			}
			return rules
		}(),
	}
}

// CreateSecurityGroup creates a Security Group.
func (c *Client) CreateSecurityGroup(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
) (*SecurityGroup, error) {
	if err := validateOperationParams(securityGroup, "create"); err != nil {
		return nil, err
	}

	resp, err := c.CreateSecurityGroupWithResponse(ctx, oapi.CreateSecurityGroupJSONRequestBody{
		Description: securityGroup.Description,
		Name:        *securityGroup.Name,
	})
	if err != nil {
		return nil, err
	}

	res, err := oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	return c.GetSecurityGroup(ctx, zone, *res.(*oapi.Reference).Id)
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
	// also compare the protocol/start port/end port parameters of the new rule to the
	// ones specified in the input rule parameter.
	rules := make(map[string]struct{})
	for _, r := range securityGroup.Rules {
		rules[*r.ID] = struct{}{}
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
					v = &oapi.SecurityGroupResource{Id: *rule.SecurityGroupID}
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
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return nil, err
	}

	sgUpdated, err := c.GetSecurityGroup(ctx, zone, *res.(*oapi.Reference).Id)
	if err != nil {
		return nil, err
	}

	// Look for an unknown rule: if we find one we hope it's the one we've just created.
	for _, r := range sgUpdated.Rules {
		if _, ok := rules[*r.ID]; !ok && (*r.FlowDirection == *rule.FlowDirection && *r.Protocol == *rule.Protocol) {
			return r, nil
		}
	}

	return nil, errors.New("unable to identify the rule created")
}

// DeleteSecurityGroup deletes a Security Group.
func (c *Client) DeleteSecurityGroup(ctx context.Context, zone string, securityGroup *SecurityGroup) error {
	if err := validateOperationParams(securityGroup, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteSecurityGroupWithResponse(apiv2.WithZone(ctx, zone), *securityGroup.ID)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
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
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// FindSecurityGroup attempts to find a Security Group by name or ID.
func (c *Client) FindSecurityGroup(ctx context.Context, zone, x string) (*SecurityGroup, error) {
	res, err := c.ListSecurityGroups(ctx, zone)
	if err != nil {
		return nil, err
	}

	for _, r := range res {
		if *r.ID == x || *r.Name == x {
			return c.GetSecurityGroup(ctx, zone, *r.ID)
		}
	}

	return nil, apiv2.ErrNotFound
}

// GetSecurityGroup returns the Security Group corresponding to the specified ID.
func (c *Client) GetSecurityGroup(ctx context.Context, zone, id string) (*SecurityGroup, error) {
	resp, err := c.GetSecurityGroupWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return securityGroupFromAPI(resp.JSON200), nil
}

// ListSecurityGroups returns the list of existing Security Groups.
func (c *Client) ListSecurityGroups(ctx context.Context, zone string) ([]*SecurityGroup, error) {
	list := make([]*SecurityGroup, 0)

	resp, err := c.ListSecurityGroupsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SecurityGroups != nil {
		for i := range *resp.JSON200.SecurityGroups {
			list = append(list, securityGroupFromAPI(&(*resp.JSON200.SecurityGroups)[i]))
		}
	}

	return list, nil
}

// AddExternalSourceToSecurityGroup adds a new external source to a
// Security Group. This operation is idempotent.
func (c *Client) AddExternalSourceToSecurityGroup(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
	cidr string,
) error {
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return err
	}

	resp, err := c.AddExternalSourceToSecurityGroupWithResponse(
		apiv2.WithZone(ctx, zone),
		*securityGroup.ID,
		oapi.AddExternalSourceToSecurityGroupJSONRequestBody{
			Cidr: cidr,
		})
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// RemoveExternalSourceFromSecurityGroup removes an external source from
// a Security Group. This operation is idempotent.
func (c *Client) RemoveExternalSourceFromSecurityGroup(
	ctx context.Context,
	zone string,
	securityGroup *SecurityGroup,
	cidr string,
) error {
	if err := validateOperationParams(securityGroup, "update"); err != nil {
		return err
	}

	resp, err := c.RemoveExternalSourceFromSecurityGroupWithResponse(
		apiv2.WithZone(ctx, zone),
		*securityGroup.ID,
		oapi.RemoveExternalSourceFromSecurityGroupJSONRequestBody{
			Cidr: cidr,
		})
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
