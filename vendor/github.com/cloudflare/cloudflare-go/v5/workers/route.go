// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

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

// RouteService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRouteService] method instead.
type RouteService struct {
	Options []option.RequestOption
}

// NewRouteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRouteService(opts ...option.RequestOption) (r *RouteService) {
	r = &RouteService{}
	r.Options = opts
	return
}

// Creates a route that maps a URL pattern to a Worker.
func (r *RouteService) New(ctx context.Context, params RouteNewParams, opts ...option.RequestOption) (res *RouteNewResponse, err error) {
	var env RouteNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the URL pattern or Worker associated with a route.
func (r *RouteService) Update(ctx context.Context, routeID string, params RouteUpdateParams, opts ...option.RequestOption) (res *RouteUpdateResponse, err error) {
	var env RouteUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes/%s", params.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns routes for a zone.
func (r *RouteService) List(ctx context.Context, query RouteListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RouteListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes", query.ZoneID)
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

// Returns routes for a zone.
func (r *RouteService) ListAutoPaging(ctx context.Context, query RouteListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RouteListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a route.
func (r *RouteService) Delete(ctx context.Context, routeID string, body RouteDeleteParams, opts ...option.RequestOption) (res *RouteDeleteResponse, err error) {
	var env RouteDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes/%s", body.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns information about a route, including URL pattern and Worker.
func (r *RouteService) Get(ctx context.Context, routeID string, query RouteGetParams, opts ...option.RequestOption) (res *RouteGetResponse, err error) {
	var env RouteGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if routeID == "" {
		err = errors.New("missing required route_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/workers/routes/%s", query.ZoneID, routeID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RouteNewResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string               `json:"script"`
	JSON   routeNewResponseJSON `json:"-"`
}

// routeNewResponseJSON contains the JSON metadata for the struct
// [RouteNewResponse]
type routeNewResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseJSON) RawJSON() string {
	return r.raw
}

type RouteUpdateResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string                  `json:"script"`
	JSON   routeUpdateResponseJSON `json:"-"`
}

// routeUpdateResponseJSON contains the JSON metadata for the struct
// [RouteUpdateResponse]
type routeUpdateResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type RouteListResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string                `json:"script"`
	JSON   routeListResponseJSON `json:"-"`
}

// routeListResponseJSON contains the JSON metadata for the struct
// [RouteListResponse]
type routeListResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeListResponseJSON) RawJSON() string {
	return r.raw
}

type RouteDeleteResponse struct {
	// Identifier.
	ID   string                  `json:"id"`
	JSON routeDeleteResponseJSON `json:"-"`
}

// routeDeleteResponseJSON contains the JSON metadata for the struct
// [RouteDeleteResponse]
type routeDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern string `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script string               `json:"script"`
	JSON   routeGetResponseJSON `json:"-"`
}

// routeGetResponseJSON contains the JSON metadata for the struct
// [RouteGetResponse]
type routeGetResponseJSON struct {
	ID          apijson.Field
	Pattern     apijson.Field
	Script      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseJSON) RawJSON() string {
	return r.raw
}

type RouteNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script param.Field[string] `json:"script"`
}

func (r RouteNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteNewResponseEnvelope struct {
	Errors   []RouteNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RouteNewResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success RouteNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeNewResponseEnvelopeJSON    `json:"-"`
}

// routeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteNewResponseEnvelope]
type routeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RouteNewResponseEnvelopeErrors struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           RouteNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             routeNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// routeNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteNewResponseEnvelopeErrors]
type routeNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteNewResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    routeNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// routeNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RouteNewResponseEnvelopeErrorsSource]
type routeNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RouteNewResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           RouteNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             routeNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// routeNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RouteNewResponseEnvelopeMessages]
type routeNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RouteNewResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    routeNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// routeNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RouteNewResponseEnvelopeMessagesSource]
type routeNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RouteNewResponseEnvelopeSuccess bool

const (
	RouteNewResponseEnvelopeSuccessTrue RouteNewResponseEnvelopeSuccess = true
)

func (r RouteNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifier.
	ID param.Field[string] `json:"id,required"`
	// Pattern to match incoming requests against.
	// [Learn more](https://developers.cloudflare.com/workers/configuration/routing/routes/#matching-behavior).
	Pattern param.Field[string] `json:"pattern,required"`
	// Name of the script to run if the route matches.
	Script param.Field[string] `json:"script"`
}

func (r RouteUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteUpdateResponseEnvelope struct {
	Errors   []RouteUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   RouteUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success RouteUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeUpdateResponseEnvelopeJSON    `json:"-"`
}

// routeUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteUpdateResponseEnvelope]
type routeUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RouteUpdateResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           RouteUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             routeUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// routeUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteUpdateResponseEnvelopeErrors]
type routeUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    routeUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// routeUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RouteUpdateResponseEnvelopeErrorsSource]
type routeUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RouteUpdateResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           RouteUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             routeUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// routeUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RouteUpdateResponseEnvelopeMessages]
type routeUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RouteUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    routeUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// routeUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RouteUpdateResponseEnvelopeMessagesSource]
type routeUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RouteUpdateResponseEnvelopeSuccess bool

const (
	RouteUpdateResponseEnvelopeSuccessTrue RouteUpdateResponseEnvelopeSuccess = true
)

func (r RouteUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteDeleteResponseEnvelope struct {
	Errors   []RouteDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RouteDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success RouteDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeDeleteResponseEnvelopeJSON    `json:"-"`
}

// routeDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteDeleteResponseEnvelope]
type routeDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RouteDeleteResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           RouteDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             routeDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// routeDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteDeleteResponseEnvelopeErrors]
type routeDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    routeDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// routeDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RouteDeleteResponseEnvelopeErrorsSource]
type routeDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RouteDeleteResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           RouteDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             routeDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// routeDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RouteDeleteResponseEnvelopeMessages]
type routeDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RouteDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    routeDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// routeDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RouteDeleteResponseEnvelopeMessagesSource]
type routeDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RouteDeleteResponseEnvelopeSuccess bool

const (
	RouteDeleteResponseEnvelopeSuccessTrue RouteDeleteResponseEnvelopeSuccess = true
)

func (r RouteDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RouteGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RouteGetResponseEnvelope struct {
	Errors   []RouteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RouteGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RouteGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success RouteGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    routeGetResponseEnvelopeJSON    `json:"-"`
}

// routeGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RouteGetResponseEnvelope]
type routeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponseEnvelopeErrors struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           RouteGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             routeGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// routeGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RouteGetResponseEnvelopeErrors]
type routeGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    routeGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// routeGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RouteGetResponseEnvelopeErrorsSource]
type routeGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           RouteGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             routeGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// routeGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RouteGetResponseEnvelopeMessages]
type routeGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RouteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RouteGetResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    routeGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// routeGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RouteGetResponseEnvelopeMessagesSource]
type routeGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RouteGetResponseEnvelopeSuccess bool

const (
	RouteGetResponseEnvelopeSuccessTrue RouteGetResponseEnvelopeSuccess = true
)

func (r RouteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RouteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
