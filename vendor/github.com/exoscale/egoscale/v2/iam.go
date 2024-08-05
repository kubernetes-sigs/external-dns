package v2

import (
	"github.com/exoscale/egoscale/v2/oapi"
)

// IAMPolicy represents an IAM policy resource.
type IAMPolicy struct {
	DefaultServiceStrategy string
	Services               map[string]IAMPolicyService
}

// IAMPolicyService represents a service of IAM policy.
type IAMPolicyService struct {
	Type  *string
	Rules []IAMPolicyServiceRule
}

// IamPolicyServiceRule represents service rule of IAM policy.
type IAMPolicyServiceRule struct {
	Action     *string
	Expression *string
	Resources  []string
}

func iamPolicyFromAPI(r *oapi.IamPolicy) *IAMPolicy {
	services := make(map[string]IAMPolicyService, len(r.Services.AdditionalProperties))
	for name, service := range r.Services.AdditionalProperties {
		rules := []IAMPolicyServiceRule{}
		if service.Rules != nil && len(*service.Rules) > 0 {
			for _, rule := range *service.Rules {
				rules = append(rules, IAMPolicyServiceRule{
					Action:     (*string)(rule.Action),
					Expression: rule.Expression,
				})
			}
		}

		services[name] = IAMPolicyService{
			Type:  (*string)(service.Type),
			Rules: rules,
		}
	}

	return &IAMPolicy{
		DefaultServiceStrategy: string(r.DefaultServiceStrategy),
		Services:               services,
	}
}
