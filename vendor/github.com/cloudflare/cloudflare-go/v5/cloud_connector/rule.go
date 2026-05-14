// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloud_connector

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
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Put Rules
func (r *RuleService) Update(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleUpdateResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cloud_connector/rules", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Put Rules
func (r *RuleService) UpdateAutoPaging(ctx context.Context, params RuleUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleUpdateResponse] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, params, opts...))
}

// Rules
func (r *RuleService) List(ctx context.Context, query RuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cloud_connector/rules", query.ZoneID)
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

// Rules
func (r *RuleService) ListAutoPaging(ctx context.Context, query RuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RuleListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type RuleUpdateResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters RuleUpdateResponseParameters `json:"parameters"`
	// Cloud Provider type
	Provider RuleUpdateResponseProvider `json:"provider"`
	JSON     ruleUpdateResponseJSON     `json:"-"`
}

// ruleUpdateResponseJSON contains the JSON metadata for the struct
// [RuleUpdateResponse]
type ruleUpdateResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Parameters  apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Parameters of Cloud Connector Rule
type RuleUpdateResponseParameters struct {
	// Host to perform Cloud Connection to
	Host string                           `json:"host"`
	JSON ruleUpdateResponseParametersJSON `json:"-"`
}

// ruleUpdateResponseParametersJSON contains the JSON metadata for the struct
// [RuleUpdateResponseParameters]
type ruleUpdateResponseParametersJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseParametersJSON) RawJSON() string {
	return r.raw
}

// Cloud Provider type
type RuleUpdateResponseProvider string

const (
	RuleUpdateResponseProviderAwsS3        RuleUpdateResponseProvider = "aws_s3"
	RuleUpdateResponseProviderCloudflareR2 RuleUpdateResponseProvider = "cloudflare_r2"
	RuleUpdateResponseProviderGcpStorage   RuleUpdateResponseProvider = "gcp_storage"
	RuleUpdateResponseProviderAzureStorage RuleUpdateResponseProvider = "azure_storage"
)

func (r RuleUpdateResponseProvider) IsKnown() bool {
	switch r {
	case RuleUpdateResponseProviderAwsS3, RuleUpdateResponseProviderCloudflareR2, RuleUpdateResponseProviderGcpStorage, RuleUpdateResponseProviderAzureStorage:
		return true
	}
	return false
}

type RuleListResponse struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Enabled     bool   `json:"enabled"`
	Expression  string `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters RuleListResponseParameters `json:"parameters"`
	// Cloud Provider type
	Provider RuleListResponseProvider `json:"provider"`
	JSON     ruleListResponseJSON     `json:"-"`
}

// ruleListResponseJSON contains the JSON metadata for the struct
// [RuleListResponse]
type ruleListResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Parameters  apijson.Field
	Provider    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseJSON) RawJSON() string {
	return r.raw
}

// Parameters of Cloud Connector Rule
type RuleListResponseParameters struct {
	// Host to perform Cloud Connection to
	Host string                         `json:"host"`
	JSON ruleListResponseParametersJSON `json:"-"`
}

// ruleListResponseParametersJSON contains the JSON metadata for the struct
// [RuleListResponseParameters]
type ruleListResponseParametersJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleListResponseParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleListResponseParametersJSON) RawJSON() string {
	return r.raw
}

// Cloud Provider type
type RuleListResponseProvider string

const (
	RuleListResponseProviderAwsS3        RuleListResponseProvider = "aws_s3"
	RuleListResponseProviderCloudflareR2 RuleListResponseProvider = "cloudflare_r2"
	RuleListResponseProviderGcpStorage   RuleListResponseProvider = "gcp_storage"
	RuleListResponseProviderAzureStorage RuleListResponseProvider = "azure_storage"
)

func (r RuleListResponseProvider) IsKnown() bool {
	switch r {
	case RuleListResponseProviderAwsS3, RuleListResponseProviderCloudflareR2, RuleListResponseProviderGcpStorage, RuleListResponseProviderAzureStorage:
		return true
	}
	return false
}

type RuleUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string]    `path:"zone_id,required"`
	Rules  []RuleUpdateParamsRule `json:"rules"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Rules)
}

type RuleUpdateParamsRule struct {
	ID          param.Field[string] `json:"id"`
	Description param.Field[string] `json:"description"`
	Enabled     param.Field[bool]   `json:"enabled"`
	Expression  param.Field[string] `json:"expression"`
	// Parameters of Cloud Connector Rule
	Parameters param.Field[RuleUpdateParamsRulesParameters] `json:"parameters"`
	// Cloud Provider type
	Provider param.Field[RuleUpdateParamsRulesProvider] `json:"provider"`
}

func (r RuleUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Parameters of Cloud Connector Rule
type RuleUpdateParamsRulesParameters struct {
	// Host to perform Cloud Connection to
	Host param.Field[string] `json:"host"`
}

func (r RuleUpdateParamsRulesParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloud Provider type
type RuleUpdateParamsRulesProvider string

const (
	RuleUpdateParamsRulesProviderAwsS3        RuleUpdateParamsRulesProvider = "aws_s3"
	RuleUpdateParamsRulesProviderCloudflareR2 RuleUpdateParamsRulesProvider = "cloudflare_r2"
	RuleUpdateParamsRulesProviderGcpStorage   RuleUpdateParamsRulesProvider = "gcp_storage"
	RuleUpdateParamsRulesProviderAzureStorage RuleUpdateParamsRulesProvider = "azure_storage"
)

func (r RuleUpdateParamsRulesProvider) IsKnown() bool {
	switch r {
	case RuleUpdateParamsRulesProviderAwsS3, RuleUpdateParamsRulesProviderCloudflareR2, RuleUpdateParamsRulesProviderGcpStorage, RuleUpdateParamsRulesProviderAzureStorage:
		return true
	}
	return false
}

type RuleListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
