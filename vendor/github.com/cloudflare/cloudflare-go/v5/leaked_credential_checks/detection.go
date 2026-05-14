// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package leaked_credential_checks

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

// DetectionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDetectionService] method instead.
type DetectionService struct {
	Options []option.RequestOption
}

// NewDetectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDetectionService(opts ...option.RequestOption) (r *DetectionService) {
	r = &DetectionService{}
	r.Options = opts
	return
}

// Create user-defined detection pattern for Leaked Credential Checks.
func (r *DetectionService) New(ctx context.Context, params DetectionNewParams, opts ...option.RequestOption) (res *DetectionNewResponse, err error) {
	var env DetectionNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update user-defined detection pattern for Leaked Credential Checks.
func (r *DetectionService) Update(ctx context.Context, detectionID string, params DetectionUpdateParams, opts ...option.RequestOption) (res *DetectionUpdateResponse, err error) {
	var env DetectionUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if detectionID == "" {
		err = errors.New("missing required detection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections/%s", params.ZoneID, detectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List user-defined detection patterns for Leaked Credential Checks.
func (r *DetectionService) List(ctx context.Context, query DetectionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DetectionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections", query.ZoneID)
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

// List user-defined detection patterns for Leaked Credential Checks.
func (r *DetectionService) ListAutoPaging(ctx context.Context, query DetectionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DetectionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Remove user-defined detection pattern for Leaked Credential Checks.
func (r *DetectionService) Delete(ctx context.Context, detectionID string, body DetectionDeleteParams, opts ...option.RequestOption) (res *DetectionDeleteResponse, err error) {
	var env DetectionDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if detectionID == "" {
		err = errors.New("missing required detection_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/leaked-credential-checks/detections/%s", body.ZoneID, detectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Defines a custom set of username/password expressions to match Leaked Credential
// Checks on.
type DetectionNewResponse struct {
	// Defines the unique ID for this custom detection.
	ID string `json:"id"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	Password string `json:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	Username string                   `json:"username"`
	JSON     detectionNewResponseJSON `json:"-"`
}

// detectionNewResponseJSON contains the JSON metadata for the struct
// [DetectionNewResponse]
type detectionNewResponseJSON struct {
	ID          apijson.Field
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetectionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detectionNewResponseJSON) RawJSON() string {
	return r.raw
}

// Defines a custom set of username/password expressions to match Leaked Credential
// Checks on.
type DetectionUpdateResponse struct {
	// Defines the unique ID for this custom detection.
	ID string `json:"id"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	Password string `json:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	Username string                      `json:"username"`
	JSON     detectionUpdateResponseJSON `json:"-"`
}

// detectionUpdateResponseJSON contains the JSON metadata for the struct
// [DetectionUpdateResponse]
type detectionUpdateResponseJSON struct {
	ID          apijson.Field
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetectionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detectionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Defines a custom set of username/password expressions to match Leaked Credential
// Checks on.
type DetectionListResponse struct {
	// Defines the unique ID for this custom detection.
	ID string `json:"id"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	Password string `json:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	Username string                    `json:"username"`
	JSON     detectionListResponseJSON `json:"-"`
}

// detectionListResponseJSON contains the JSON metadata for the struct
// [DetectionListResponse]
type detectionListResponseJSON struct {
	ID          apijson.Field
	Password    apijson.Field
	Username    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetectionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detectionListResponseJSON) RawJSON() string {
	return r.raw
}

type DetectionDeleteResponse = interface{}

type DetectionNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	Password param.Field[string] `json:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	Username param.Field[string] `json:"username"`
}

func (r DetectionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DetectionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Defines a custom set of username/password expressions to match Leaked Credential
	// Checks on.
	Result DetectionNewResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success DetectionNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    detectionNewResponseEnvelopeJSON    `json:"-"`
}

// detectionNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DetectionNewResponseEnvelope]
type detectionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetectionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detectionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type DetectionNewResponseEnvelopeSuccess bool

const (
	DetectionNewResponseEnvelopeSuccessTrue DetectionNewResponseEnvelopeSuccess = true
)

func (r DetectionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DetectionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DetectionUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines ehe ruleset expression to use in matching the password in a request.
	Password param.Field[string] `json:"password"`
	// Defines the ruleset expression to use in matching the username in a request.
	Username param.Field[string] `json:"username"`
}

func (r DetectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DetectionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Defines a custom set of username/password expressions to match Leaked Credential
	// Checks on.
	Result DetectionUpdateResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success DetectionUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    detectionUpdateResponseEnvelopeJSON    `json:"-"`
}

// detectionUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DetectionUpdateResponseEnvelope]
type detectionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetectionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detectionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type DetectionUpdateResponseEnvelopeSuccess bool

const (
	DetectionUpdateResponseEnvelopeSuccessTrue DetectionUpdateResponseEnvelopeSuccess = true
)

func (r DetectionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DetectionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DetectionListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DetectionDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DetectionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   DetectionDeleteResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success DetectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    detectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// detectionDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DetectionDeleteResponseEnvelope]
type detectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DetectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r detectionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type DetectionDeleteResponseEnvelopeSuccess bool

const (
	DetectionDeleteResponseEnvelopeSuccessTrue DetectionDeleteResponseEnvelopeSuccess = true
)

func (r DetectionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DetectionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
