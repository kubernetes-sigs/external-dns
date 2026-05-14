// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_shield

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// PolicyService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyService] method instead.
type PolicyService struct {
	Options []option.RequestOption
}

// NewPolicyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPolicyService(opts ...option.RequestOption) (r *PolicyService) {
	r = &PolicyService{}
	r.Options = opts
	return
}

// Create a Page Shield policy.
func (r *PolicyService) New(ctx context.Context, params PolicyNewParams, opts ...option.RequestOption) (res *PolicyNewResponse, err error) {
	var env PolicyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Page Shield policy by ID.
func (r *PolicyService) Update(ctx context.Context, policyID string, params PolicyUpdateParams, opts ...option.RequestOption) (res *PolicyUpdateResponse, err error) {
	var env PolicyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", params.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Page Shield policies.
func (r *PolicyService) List(ctx context.Context, query PolicyListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PolicyListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists all Page Shield policies.
func (r *PolicyService) ListAutoPaging(ctx context.Context, query PolicyListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PolicyListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Page Shield policy by ID.
func (r *PolicyService) Delete(ctx context.Context, policyID string, body PolicyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", body.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches a Page Shield policy by ID.
func (r *PolicyService) Get(ctx context.Context, policyID string, query PolicyGetParams, opts ...option.RequestOption) (res *PolicyGetResponse, err error) {
	var env PolicyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/page_shield/policies/%s", query.ZoneID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PolicyParam struct {
	// The action to take if the expression matches
	Action param.Field[PolicyAction] `json:"action,required"`
	// A description for the policy
	Description param.Field[string] `json:"description,required"`
	// Whether the policy is enabled
	Enabled param.Field[bool] `json:"enabled,required"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression param.Field[string] `json:"expression,required"`
	// The policy which will be applied
	Value param.Field[string] `json:"value,required"`
}

func (r PolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type PolicyAction string

const (
	PolicyActionAllow PolicyAction = "allow"
	PolicyActionLog   PolicyAction = "log"
)

func (r PolicyAction) IsKnown() bool {
	switch r {
	case PolicyActionAllow, PolicyActionLog:
		return true
	}
	return false
}

type PolicyNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The action to take if the expression matches
	Action PolicyNewResponseAction `json:"action,required"`
	// A description for the policy
	Description string `json:"description,required"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled,required"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression,required"`
	// The policy which will be applied
	Value string                `json:"value,required"`
	JSON  policyNewResponseJSON `json:"-"`
}

// policyNewResponseJSON contains the JSON metadata for the struct
// [PolicyNewResponse]
type policyNewResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take if the expression matches
type PolicyNewResponseAction string

const (
	PolicyNewResponseActionAllow PolicyNewResponseAction = "allow"
	PolicyNewResponseActionLog   PolicyNewResponseAction = "log"
)

func (r PolicyNewResponseAction) IsKnown() bool {
	switch r {
	case PolicyNewResponseActionAllow, PolicyNewResponseActionLog:
		return true
	}
	return false
}

type PolicyUpdateResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The action to take if the expression matches
	Action PolicyUpdateResponseAction `json:"action,required"`
	// A description for the policy
	Description string `json:"description,required"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled,required"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression,required"`
	// The policy which will be applied
	Value string                   `json:"value,required"`
	JSON  policyUpdateResponseJSON `json:"-"`
}

// policyUpdateResponseJSON contains the JSON metadata for the struct
// [PolicyUpdateResponse]
type policyUpdateResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take if the expression matches
type PolicyUpdateResponseAction string

const (
	PolicyUpdateResponseActionAllow PolicyUpdateResponseAction = "allow"
	PolicyUpdateResponseActionLog   PolicyUpdateResponseAction = "log"
)

func (r PolicyUpdateResponseAction) IsKnown() bool {
	switch r {
	case PolicyUpdateResponseActionAllow, PolicyUpdateResponseActionLog:
		return true
	}
	return false
}

type PolicyListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The action to take if the expression matches
	Action PolicyListResponseAction `json:"action,required"`
	// A description for the policy
	Description string `json:"description,required"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled,required"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression,required"`
	// The policy which will be applied
	Value string                 `json:"value,required"`
	JSON  policyListResponseJSON `json:"-"`
}

// policyListResponseJSON contains the JSON metadata for the struct
// [PolicyListResponse]
type policyListResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take if the expression matches
type PolicyListResponseAction string

const (
	PolicyListResponseActionAllow PolicyListResponseAction = "allow"
	PolicyListResponseActionLog   PolicyListResponseAction = "log"
)

func (r PolicyListResponseAction) IsKnown() bool {
	switch r {
	case PolicyListResponseActionAllow, PolicyListResponseActionLog:
		return true
	}
	return false
}

type PolicyGetResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The action to take if the expression matches
	Action PolicyGetResponseAction `json:"action,required"`
	// A description for the policy
	Description string `json:"description,required"`
	// Whether the policy is enabled
	Enabled bool `json:"enabled,required"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression string `json:"expression,required"`
	// The policy which will be applied
	Value string                `json:"value,required"`
	JSON  policyGetResponseJSON `json:"-"`
}

// policyGetResponseJSON contains the JSON metadata for the struct
// [PolicyGetResponse]
type policyGetResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyGetResponseJSON) RawJSON() string {
	return r.raw
}

// The action to take if the expression matches
type PolicyGetResponseAction string

const (
	PolicyGetResponseActionAllow PolicyGetResponseAction = "allow"
	PolicyGetResponseActionLog   PolicyGetResponseAction = "log"
)

func (r PolicyGetResponseAction) IsKnown() bool {
	switch r {
	case PolicyGetResponseActionAllow, PolicyGetResponseActionLog:
		return true
	}
	return false
}

type PolicyNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	Policy PolicyParam         `json:"policy,required"`
}

func (r PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Policy)
}

type PolicyNewResponseEnvelope struct {
	Result PolicyNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  PolicyNewResponseEnvelopeSuccess `json:"success,required"`
	Errors   []shared.ResponseInfo            `json:"errors"`
	Messages []shared.ResponseInfo            `json:"messages"`
	JSON     policyNewResponseEnvelopeJSON    `json:"-"`
}

// policyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyNewResponseEnvelope]
type policyNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyNewResponseEnvelopeSuccess bool

const (
	PolicyNewResponseEnvelopeSuccessTrue PolicyNewResponseEnvelopeSuccess = true
)

func (r PolicyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PolicyUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to take if the expression matches
	Action param.Field[PolicyUpdateParamsAction] `json:"action"`
	// A description for the policy
	Description param.Field[string] `json:"description"`
	// Whether the policy is enabled
	Enabled param.Field[bool] `json:"enabled"`
	// The expression which must match for the policy to be applied, using the
	// Cloudflare Firewall rule expression syntax
	Expression param.Field[string] `json:"expression"`
	// The policy which will be applied
	Value param.Field[string] `json:"value"`
}

func (r PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to take if the expression matches
type PolicyUpdateParamsAction string

const (
	PolicyUpdateParamsActionAllow PolicyUpdateParamsAction = "allow"
	PolicyUpdateParamsActionLog   PolicyUpdateParamsAction = "log"
)

func (r PolicyUpdateParamsAction) IsKnown() bool {
	switch r {
	case PolicyUpdateParamsActionAllow, PolicyUpdateParamsActionLog:
		return true
	}
	return false
}

type PolicyUpdateResponseEnvelope struct {
	Result PolicyUpdateResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  PolicyUpdateResponseEnvelopeSuccess `json:"success,required"`
	Errors   []shared.ResponseInfo               `json:"errors"`
	Messages []shared.ResponseInfo               `json:"messages"`
	JSON     policyUpdateResponseEnvelopeJSON    `json:"-"`
}

// policyUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyUpdateResponseEnvelope]
type policyUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyUpdateResponseEnvelopeSuccess bool

const (
	PolicyUpdateResponseEnvelopeSuccessTrue PolicyUpdateResponseEnvelopeSuccess = true
)

func (r PolicyUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PolicyListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PolicyDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PolicyGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PolicyGetResponseEnvelope struct {
	Result PolicyGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success  PolicyGetResponseEnvelopeSuccess `json:"success,required"`
	Errors   []shared.ResponseInfo            `json:"errors"`
	Messages []shared.ResponseInfo            `json:"messages"`
	JSON     policyGetResponseEnvelopeJSON    `json:"-"`
}

// policyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PolicyGetResponseEnvelope]
type policyGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolicyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r policyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PolicyGetResponseEnvelopeSuccess bool

const (
	PolicyGetResponseEnvelopeSuccessTrue PolicyGetResponseEnvelopeSuccess = true
)

func (r PolicyGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PolicyGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
