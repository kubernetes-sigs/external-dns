// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// RegionalHostnameService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionalHostnameService] method instead.
type RegionalHostnameService struct {
	Options []option.RequestOption
	Regions *RegionalHostnameRegionService
}

// NewRegionalHostnameService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegionalHostnameService(opts ...option.RequestOption) (r *RegionalHostnameService) {
	r = &RegionalHostnameService{}
	r.Options = opts
	r.Regions = NewRegionalHostnameRegionService(opts...)
	return
}

// Create a new Regional Hostname entry. Cloudflare will only use data centers that
// are physically located within the chosen region to decrypt and service HTTPS
// traffic. Learn more about
// [Regional Services](https://developers.cloudflare.com/data-localization/regional-services/get-started/).
func (r *RegionalHostnameService) New(ctx context.Context, params RegionalHostnameNewParams, opts ...option.RequestOption) (res *RegionalHostnameNewResponse, err error) {
	var env RegionalHostnameNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Regional Hostnames within a zone.
func (r *RegionalHostnameService) List(ctx context.Context, query RegionalHostnameListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RegionalHostnameListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames", query.ZoneID)
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

// List all Regional Hostnames within a zone.
func (r *RegionalHostnameService) ListAutoPaging(ctx context.Context, query RegionalHostnameListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RegionalHostnameListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete the region configuration for a specific Regional Hostname.
func (r *RegionalHostnameService) Delete(ctx context.Context, hostname string, body RegionalHostnameDeleteParams, opts ...option.RequestOption) (res *RegionalHostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", body.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update the configuration for a specific Regional Hostname. Only the region_key
// of a hostname is mutable.
func (r *RegionalHostnameService) Edit(ctx context.Context, hostname string, params RegionalHostnameEditParams, opts ...option.RequestOption) (res *RegionalHostnameEditResponse, err error) {
	var env RegionalHostnameEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", params.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the configuration for a specific Regional Hostname, within a zone.
func (r *RegionalHostnameService) Get(ctx context.Context, hostname string, query RegionalHostnameGetParams, opts ...option.RequestOption) (res *RegionalHostnameGetResponse, err error) {
	var env RegionalHostnameGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", query.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RegionalHostnameNewResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string `json:"region_key,required"`
	// Configure which routing method to use for the regional hostname
	Routing string                          `json:"routing"`
	JSON    regionalHostnameNewResponseJSON `json:"-"`
}

// regionalHostnameNewResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameNewResponse]
type regionalHostnameNewResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	Routing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameListResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string `json:"region_key,required"`
	// Configure which routing method to use for the regional hostname
	Routing string                           `json:"routing"`
	JSON    regionalHostnameListResponseJSON `json:"-"`
}

// regionalHostnameListResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameListResponse]
type regionalHostnameListResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	Routing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameDeleteResponse struct {
	Errors   []RegionalHostnameDeleteResponseError   `json:"errors,required"`
	Messages []RegionalHostnameDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success RegionalHostnameDeleteResponseSuccess `json:"success,required"`
	JSON    regionalHostnameDeleteResponseJSON    `json:"-"`
}

// regionalHostnameDeleteResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameDeleteResponse]
type regionalHostnameDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameDeleteResponseError struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           RegionalHostnameDeleteResponseErrorsSource `json:"source"`
	JSON             regionalHostnameDeleteResponseErrorJSON    `json:"-"`
}

// regionalHostnameDeleteResponseErrorJSON contains the JSON metadata for the
// struct [RegionalHostnameDeleteResponseError]
type regionalHostnameDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameDeleteResponseErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    regionalHostnameDeleteResponseErrorsSourceJSON `json:"-"`
}

// regionalHostnameDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [RegionalHostnameDeleteResponseErrorsSource]
type regionalHostnameDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameDeleteResponseMessage struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           RegionalHostnameDeleteResponseMessagesSource `json:"source"`
	JSON             regionalHostnameDeleteResponseMessageJSON    `json:"-"`
}

// regionalHostnameDeleteResponseMessageJSON contains the JSON metadata for the
// struct [RegionalHostnameDeleteResponseMessage]
type regionalHostnameDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameDeleteResponseMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    regionalHostnameDeleteResponseMessagesSourceJSON `json:"-"`
}

// regionalHostnameDeleteResponseMessagesSourceJSON contains the JSON metadata for
// the struct [RegionalHostnameDeleteResponseMessagesSource]
type regionalHostnameDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegionalHostnameDeleteResponseSuccess bool

const (
	RegionalHostnameDeleteResponseSuccessTrue RegionalHostnameDeleteResponseSuccess = true
)

func (r RegionalHostnameDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RegionalHostnameEditResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string `json:"region_key,required"`
	// Configure which routing method to use for the regional hostname
	Routing string                           `json:"routing"`
	JSON    regionalHostnameEditResponseJSON `json:"-"`
}

// regionalHostnameEditResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameEditResponse]
type regionalHostnameEditResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	Routing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameGetResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string `json:"region_key,required"`
	// Configure which routing method to use for the regional hostname
	Routing string                          `json:"routing"`
	JSON    regionalHostnameGetResponseJSON `json:"-"`
}

// regionalHostnameGetResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameGetResponse]
type regionalHostnameGetResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	Routing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname param.Field[string] `json:"hostname,required"`
	// Identifying key for the region
	RegionKey param.Field[string] `json:"region_key,required"`
	// Configure which routing method to use for the regional hostname
	Routing param.Field[string] `json:"routing"`
}

func (r RegionalHostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegionalHostnameNewResponseEnvelope struct {
	Errors   []RegionalHostnameNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegionalHostnameNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RegionalHostnameNewResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalHostnameNewResponse                `json:"result"`
	JSON    regionalHostnameNewResponseEnvelopeJSON    `json:"-"`
}

// regionalHostnameNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalHostnameNewResponseEnvelope]
type regionalHostnameNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           RegionalHostnameNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             regionalHostnameNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// regionalHostnameNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegionalHostnameNewResponseEnvelopeErrors]
type regionalHostnameNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    regionalHostnameNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// regionalHostnameNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RegionalHostnameNewResponseEnvelopeErrorsSource]
type regionalHostnameNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           RegionalHostnameNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             regionalHostnameNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// regionalHostnameNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RegionalHostnameNewResponseEnvelopeMessages]
type regionalHostnameNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    regionalHostnameNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// regionalHostnameNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RegionalHostnameNewResponseEnvelopeMessagesSource]
type regionalHostnameNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegionalHostnameNewResponseEnvelopeSuccess bool

const (
	RegionalHostnameNewResponseEnvelopeSuccessTrue RegionalHostnameNewResponseEnvelopeSuccess = true
)

func (r RegionalHostnameNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegionalHostnameListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalHostnameDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalHostnameEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifying key for the region
	RegionKey param.Field[string] `json:"region_key,required"`
}

func (r RegionalHostnameEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegionalHostnameEditResponseEnvelope struct {
	Errors   []RegionalHostnameEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegionalHostnameEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RegionalHostnameEditResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalHostnameEditResponse                `json:"result"`
	JSON    regionalHostnameEditResponseEnvelopeJSON    `json:"-"`
}

// regionalHostnameEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalHostnameEditResponseEnvelope]
type regionalHostnameEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameEditResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           RegionalHostnameEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             regionalHostnameEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// regionalHostnameEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RegionalHostnameEditResponseEnvelopeErrors]
type regionalHostnameEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameEditResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    regionalHostnameEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// regionalHostnameEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RegionalHostnameEditResponseEnvelopeErrorsSource]
type regionalHostnameEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameEditResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           RegionalHostnameEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             regionalHostnameEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// regionalHostnameEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RegionalHostnameEditResponseEnvelopeMessages]
type regionalHostnameEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    regionalHostnameEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// regionalHostnameEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [RegionalHostnameEditResponseEnvelopeMessagesSource]
type regionalHostnameEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegionalHostnameEditResponseEnvelopeSuccess bool

const (
	RegionalHostnameEditResponseEnvelopeSuccessTrue RegionalHostnameEditResponseEnvelopeSuccess = true
)

func (r RegionalHostnameEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegionalHostnameGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalHostnameGetResponseEnvelope struct {
	Errors   []RegionalHostnameGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RegionalHostnameGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success RegionalHostnameGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalHostnameGetResponse                `json:"result"`
	JSON    regionalHostnameGetResponseEnvelopeJSON    `json:"-"`
}

// regionalHostnameGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalHostnameGetResponseEnvelope]
type regionalHostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           RegionalHostnameGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             regionalHostnameGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// regionalHostnameGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RegionalHostnameGetResponseEnvelopeErrors]
type regionalHostnameGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    regionalHostnameGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// regionalHostnameGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [RegionalHostnameGetResponseEnvelopeErrorsSource]
type regionalHostnameGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           RegionalHostnameGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             regionalHostnameGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// regionalHostnameGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RegionalHostnameGetResponseEnvelopeMessages]
type regionalHostnameGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RegionalHostnameGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    regionalHostnameGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// regionalHostnameGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [RegionalHostnameGetResponseEnvelopeMessagesSource]
type regionalHostnameGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegionalHostnameGetResponseEnvelopeSuccess bool

const (
	RegionalHostnameGetResponseEnvelopeSuccessTrue RegionalHostnameGetResponseEnvelopeSuccess = true
)

func (r RegionalHostnameGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
