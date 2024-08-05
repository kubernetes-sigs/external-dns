package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// IAMRole represents an IAM Role resource.
type IAMRole struct {
	ID          *string `req-for:"delete,update"`
	Description *string
	Editable    *bool
	Labels      map[string]string
	Name        *string `req-for:"create"`
	Permissions []string
	Policy      *IAMPolicy
}

func iamRoleFromAPI(r *oapi.IamRole) *IAMRole {
	labels := map[string]string{}
	if r.Labels != nil && r.Labels.AdditionalProperties != nil {
		labels = r.Labels.AdditionalProperties
	}

	permissions := []string{}
	if r.Permissions != nil {
		for _, p := range *r.Permissions {
			permissions = append(permissions, string(p))
		}
	}

	return &IAMRole{
		ID:          r.Id,
		Description: r.Description,
		Editable:    r.Editable,
		Labels:      labels,
		Name:        r.Name,
		Permissions: permissions,

		Policy: iamPolicyFromAPI(r.Policy),
	}
}

// GetIAMRole returns the IAM Role.
func (c *Client) GetIAMRole(ctx context.Context, zone, id string) (*IAMRole, error) {
	resp, err := c.GetIamRoleWithResponse(apiv2.WithZone(ctx, zone), id)
	if err != nil {
		return nil, err
	}

	return iamRoleFromAPI(resp.JSON200), nil
}

// ListIAMRoles returns the list of existing IAM Roles.
func (c *Client) ListIAMRoles(ctx context.Context, zone string) ([]*IAMRole, error) {
	list := make([]*IAMRole, 0)

	resp, err := c.ListIamRolesWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.IamRoles != nil {
		for i := range *resp.JSON200.IamRoles {
			list = append(list, iamRoleFromAPI(&(*resp.JSON200.IamRoles)[i]))
		}
	}

	return list, nil
}

// CreateIAMRole creates a IAM Role.
func (c *Client) CreateIAMRole(
	ctx context.Context,
	zone string,
	role *IAMRole,
) (*IAMRole, error) {
	if err := validateOperationParams(role, "create"); err != nil {
		return nil, err
	}

	req := oapi.CreateIamRoleJSONRequestBody{
		Name:        *role.Name,
		Description: role.Description,
		Editable:    role.Editable,
	}

	if role.Labels != nil {
		req.Labels = &oapi.Labels{
			AdditionalProperties: role.Labels,
		}
	}

	if role.Permissions != nil {
		t := []oapi.CreateIamRoleJSONBodyPermissions{}
		for _, p := range role.Permissions {
			t = append(t, oapi.CreateIamRoleJSONBodyPermissions(p))
		}

		req.Permissions = &t
	}

	if role.Policy != nil {
		t := oapi.IamPolicy{
			DefaultServiceStrategy: oapi.IamPolicyDefaultServiceStrategy(role.Policy.DefaultServiceStrategy),
			Services: oapi.IamPolicy_Services{
				AdditionalProperties: map[string]oapi.IamServicePolicy{},
			},
		}

		if len(role.Policy.Services) > 0 {
			for name, service := range role.Policy.Services {
				s := oapi.IamServicePolicy{
					Type: (*oapi.IamServicePolicyType)(service.Type),
				}

				if service.Rules != nil {
					rules := []oapi.IamServicePolicyRule{}

					for _, rule := range service.Rules {
						r := oapi.IamServicePolicyRule{
							Action:     (*oapi.IamServicePolicyRuleAction)(rule.Action),
							Expression: rule.Expression,
						}

						rules = append(rules, r)
					}

					s.Rules = &rules
				}

				t.Services.AdditionalProperties[name] = s

			}
		}

		req.Policy = &t
	}

	resp, err := c.CreateIamRoleWithResponse(
		apiv2.WithZone(ctx, zone),
		req,
	)
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

	return c.GetIAMRole(ctx, zone, *res.(*struct {
		Command *string `json:"command,omitempty"`
		Id      *string `json:"id,omitempty"` // revive:disable-line
		Link    *string `json:"link,omitempty"`
	}).Id)
}

// DeleteIAMRole deletes IAM Role.
func (c *Client) DeleteIAMRole(ctx context.Context, zone string, role *IAMRole) error {
	if err := validateOperationParams(role, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteIamRoleWithResponse(apiv2.WithZone(ctx, zone), *role.ID)
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

// UpdateIAMRole updates existing IAM Role.
func (c *Client) UpdateIAMRole(ctx context.Context, zone string, role *IAMRole) error {
	if err := validateOperationParams(role, "update"); err != nil {
		return err
	}

	req := oapi.UpdateIamRoleJSONRequestBody{
		Description: role.Description,
	}

	if role.Labels != nil {
		req.Labels = &oapi.Labels{
			AdditionalProperties: role.Labels,
		}
	}

	if role.Permissions != nil {
		t := []oapi.UpdateIamRoleJSONBodyPermissions{}
		for _, p := range role.Permissions {
			t = append(t, oapi.UpdateIamRoleJSONBodyPermissions(p))
		}

		req.Permissions = &t
	}

	resp, err := c.UpdateIamRoleWithResponse(
		apiv2.WithZone(ctx, zone),
		*role.ID,
		req,
	)
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

// UpdateIAMRolePolicy updates existing IAM Role policy.
func (c *Client) UpdateIAMRolePolicy(ctx context.Context, zone string, role *IAMRole) error {
	if err := validateOperationParams(role, "update"); err != nil {
		return err
	}

	req := oapi.UpdateIamRolePolicyJSONRequestBody{
		DefaultServiceStrategy: oapi.IamPolicyDefaultServiceStrategy(role.Policy.DefaultServiceStrategy),
		Services: oapi.IamPolicy_Services{
			AdditionalProperties: map[string]oapi.IamServicePolicy{},
		},
	}

	if len(role.Policy.Services) > 0 {
		for name, service := range role.Policy.Services {
			t := oapi.IamServicePolicy{
				Type: (*oapi.IamServicePolicyType)(service.Type),
			}

			if service.Rules != nil {
				rules := []oapi.IamServicePolicyRule{}

				for _, rule := range service.Rules {
					r := oapi.IamServicePolicyRule{
						Action:     (*oapi.IamServicePolicyRuleAction)(rule.Action),
						Expression: rule.Expression,
					}

					rules = append(rules, r)
				}

				t.Rules = &rules
			}

			req.Services.AdditionalProperties[name] = t

		}
	}

	resp, err := c.UpdateIamRolePolicyWithResponse(
		apiv2.WithZone(ctx, zone),
		*role.ID,
		req,
	)
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
