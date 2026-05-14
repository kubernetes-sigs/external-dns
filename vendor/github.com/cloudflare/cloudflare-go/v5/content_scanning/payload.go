// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

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

// PayloadService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPayloadService] method instead.
type PayloadService struct {
	Options []option.RequestOption
}

// NewPayloadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPayloadService(opts ...option.RequestOption) (r *PayloadService) {
	r = &PayloadService{}
	r.Options = opts
	return
}

// Add custom scan expressions for Content Scanning.
func (r *PayloadService) New(ctx context.Context, params PayloadNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[PayloadNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Add custom scan expressions for Content Scanning.
func (r *PayloadService) NewAutoPaging(ctx context.Context, params PayloadNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PayloadNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, params, opts...))
}

// Get a list of existing custom scan expressions for Content Scanning.
func (r *PayloadService) List(ctx context.Context, query PayloadListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PayloadListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads", query.ZoneID)
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

// Get a list of existing custom scan expressions for Content Scanning.
func (r *PayloadService) ListAutoPaging(ctx context.Context, query PayloadListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PayloadListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Content Scan Custom Expression.
func (r *PayloadService) Delete(ctx context.Context, expressionID string, body PayloadDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[PayloadDeleteResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if expressionID == "" {
		err = errors.New("missing required expression_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/payloads/%s", body.ZoneID, expressionID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Delete a Content Scan Custom Expression.
func (r *PayloadService) DeleteAutoPaging(ctx context.Context, expressionID string, body PayloadDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PayloadDeleteResponse] {
	return pagination.NewSinglePageAutoPager(r.Delete(ctx, expressionID, body, opts...))
}

// Defines a custom scan expression to match Content Scanning on.
type PayloadNewResponse struct {
	// defines the unique ID for this custom scan expression.
	ID string `json:"id"`
	// Defines the ruleset expression to use in matching content objects.
	Payload string                 `json:"payload"`
	JSON    payloadNewResponseJSON `json:"-"`
}

// payloadNewResponseJSON contains the JSON metadata for the struct
// [PayloadNewResponse]
type payloadNewResponseJSON struct {
	ID          apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayloadNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payloadNewResponseJSON) RawJSON() string {
	return r.raw
}

// Defines a custom scan expression to match Content Scanning on.
type PayloadListResponse struct {
	// defines the unique ID for this custom scan expression.
	ID string `json:"id"`
	// Defines the ruleset expression to use in matching content objects.
	Payload string                  `json:"payload"`
	JSON    payloadListResponseJSON `json:"-"`
}

// payloadListResponseJSON contains the JSON metadata for the struct
// [PayloadListResponse]
type payloadListResponseJSON struct {
	ID          apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayloadListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payloadListResponseJSON) RawJSON() string {
	return r.raw
}

// Defines a custom scan expression to match Content Scanning on.
type PayloadDeleteResponse struct {
	// defines the unique ID for this custom scan expression.
	ID string `json:"id"`
	// Defines the ruleset expression to use in matching content objects.
	Payload string                    `json:"payload"`
	JSON    payloadDeleteResponseJSON `json:"-"`
}

// payloadDeleteResponseJSON contains the JSON metadata for the struct
// [PayloadDeleteResponse]
type payloadDeleteResponseJSON struct {
	ID          apijson.Field
	Payload     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PayloadDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r payloadDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PayloadNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string]    `path:"zone_id,required"`
	Body   []PayloadNewParamsBody `json:"body,required"`
}

func (r PayloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PayloadNewParamsBody struct {
	// Defines the ruleset expression to use in matching content objects.
	Payload param.Field[string] `json:"payload,required"`
}

func (r PayloadNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PayloadListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PayloadDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
