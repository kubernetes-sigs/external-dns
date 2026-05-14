// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth

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

// OriginTLSClientAuthService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginTLSClientAuthService] method instead.
type OriginTLSClientAuthService struct {
	Options   []option.RequestOption
	Hostnames *HostnameService
	Settings  *SettingService
}

// NewOriginTLSClientAuthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginTLSClientAuthService(opts ...option.RequestOption) (r *OriginTLSClientAuthService) {
	r = &OriginTLSClientAuthService{}
	r.Options = opts
	r.Hostnames = NewHostnameService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}

// Upload your own certificate you want Cloudflare to use for edge-to-origin
// communication to override the shared certificate. Please note that it is
// important to keep only one certificate active. Also, make sure to enable
// zone-level authenticated origin pulls by making a PUT call to settings endpoint
// to see the uploaded certificate in use.
func (r *OriginTLSClientAuthService) New(ctx context.Context, params OriginTLSClientAuthNewParams, opts ...option.RequestOption) (res *OriginTLSClientAuthNewResponse, err error) {
	var env OriginTLSClientAuthNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Certificates
func (r *OriginTLSClientAuthService) List(ctx context.Context, query OriginTLSClientAuthListParams, opts ...option.RequestOption) (res *pagination.SinglePage[OriginTLSClientAuthListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", query.ZoneID)
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

// List Certificates
func (r *OriginTLSClientAuthService) ListAutoPaging(ctx context.Context, query OriginTLSClientAuthListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OriginTLSClientAuthListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Certificate
func (r *OriginTLSClientAuthService) Delete(ctx context.Context, certificateID string, body OriginTLSClientAuthDeleteParams, opts ...option.RequestOption) (res *OriginTLSClientAuthDeleteResponse, err error) {
	var env OriginTLSClientAuthDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", body.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Certificate Details
func (r *OriginTLSClientAuthService) Get(ctx context.Context, certificateID string, query OriginTLSClientAuthGetParams, opts ...option.RequestOption) (res *OriginTLSClientAuthGetResponse, err error) {
	var env OriginTLSClientAuthGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", query.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneAuthenticatedOriginPull struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// When the certificate from the authority expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate activation.
	Status ZoneAuthenticatedOriginPullStatus `json:"status"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                       `json:"uploaded_on" format:"date-time"`
	JSON       zoneAuthenticatedOriginPullJSON `json:"-"`
}

// zoneAuthenticatedOriginPullJSON contains the JSON metadata for the struct
// [ZoneAuthenticatedOriginPull]
type zoneAuthenticatedOriginPullJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	ExpiresOn   apijson.Field
	Issuer      apijson.Field
	Signature   apijson.Field
	Status      apijson.Field
	UploadedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAuthenticatedOriginPull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAuthenticatedOriginPullJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate activation.
type ZoneAuthenticatedOriginPullStatus string

const (
	ZoneAuthenticatedOriginPullStatusInitializing       ZoneAuthenticatedOriginPullStatus = "initializing"
	ZoneAuthenticatedOriginPullStatusPendingDeployment  ZoneAuthenticatedOriginPullStatus = "pending_deployment"
	ZoneAuthenticatedOriginPullStatusPendingDeletion    ZoneAuthenticatedOriginPullStatus = "pending_deletion"
	ZoneAuthenticatedOriginPullStatusActive             ZoneAuthenticatedOriginPullStatus = "active"
	ZoneAuthenticatedOriginPullStatusDeleted            ZoneAuthenticatedOriginPullStatus = "deleted"
	ZoneAuthenticatedOriginPullStatusDeploymentTimedOut ZoneAuthenticatedOriginPullStatus = "deployment_timed_out"
	ZoneAuthenticatedOriginPullStatusDeletionTimedOut   ZoneAuthenticatedOriginPullStatus = "deletion_timed_out"
)

func (r ZoneAuthenticatedOriginPullStatus) IsKnown() bool {
	switch r {
	case ZoneAuthenticatedOriginPullStatusInitializing, ZoneAuthenticatedOriginPullStatusPendingDeployment, ZoneAuthenticatedOriginPullStatusPendingDeletion, ZoneAuthenticatedOriginPullStatusActive, ZoneAuthenticatedOriginPullStatusDeleted, ZoneAuthenticatedOriginPullStatusDeploymentTimedOut, ZoneAuthenticatedOriginPullStatusDeletionTimedOut:
		return true
	}
	return false
}

type OriginTLSClientAuthNewResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                             `json:"private_key"`
	JSON       originTLSClientAuthNewResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// originTLSClientAuthNewResponseJSON contains the JSON metadata for the struct
// [OriginTLSClientAuthNewResponse]
type originTLSClientAuthNewResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthListResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                              `json:"private_key"`
	JSON       originTLSClientAuthListResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// originTLSClientAuthListResponseJSON contains the JSON metadata for the struct
// [OriginTLSClientAuthListResponse]
type originTLSClientAuthListResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthListResponseJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthDeleteResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                                `json:"private_key"`
	JSON       originTLSClientAuthDeleteResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// originTLSClientAuthDeleteResponseJSON contains the JSON metadata for the struct
// [OriginTLSClientAuthDeleteResponse]
type originTLSClientAuthDeleteResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthGetResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                             `json:"private_key"`
	JSON       originTLSClientAuthGetResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// originTLSClientAuthGetResponseJSON contains the JSON metadata for the struct
// [OriginTLSClientAuthGetResponse]
type originTLSClientAuthGetResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The zone's leaf certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r OriginTLSClientAuthNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OriginTLSClientAuthNewResponseEnvelope struct {
	Errors   []OriginTLSClientAuthNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OriginTLSClientAuthNewResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginTLSClientAuthNewResponse                `json:"result"`
	JSON    originTLSClientAuthNewResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthNewResponseEnvelope]
type originTLSClientAuthNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthNewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           OriginTLSClientAuthNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             originTLSClientAuthNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthNewResponseEnvelopeErrors]
type originTLSClientAuthNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    originTLSClientAuthNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthNewResponseEnvelopeErrorsSource]
type originTLSClientAuthNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthNewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           OriginTLSClientAuthNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             originTLSClientAuthNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthNewResponseEnvelopeMessages]
type originTLSClientAuthNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    originTLSClientAuthNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// originTLSClientAuthNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthNewResponseEnvelopeMessagesSource]
type originTLSClientAuthNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginTLSClientAuthNewResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthNewResponseEnvelopeSuccessTrue OriginTLSClientAuthNewResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSClientAuthListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthDeleteResponseEnvelope struct {
	Errors   []OriginTLSClientAuthDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OriginTLSClientAuthDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginTLSClientAuthDeleteResponse                `json:"result"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthDeleteResponseEnvelope]
type originTLSClientAuthDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthDeleteResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           OriginTLSClientAuthDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             originTLSClientAuthDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthDeleteResponseEnvelopeErrors]
type originTLSClientAuthDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthDeleteResponseEnvelopeErrorsSource]
type originTLSClientAuthDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthDeleteResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           OriginTLSClientAuthDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             originTLSClientAuthDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthDeleteResponseEnvelopeMessages]
type originTLSClientAuthDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    originTLSClientAuthDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// originTLSClientAuthDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [OriginTLSClientAuthDeleteResponseEnvelopeMessagesSource]
type originTLSClientAuthDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginTLSClientAuthDeleteResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthDeleteResponseEnvelopeSuccessTrue OriginTLSClientAuthDeleteResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginTLSClientAuthGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginTLSClientAuthGetResponseEnvelope struct {
	Errors   []OriginTLSClientAuthGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OriginTLSClientAuthGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OriginTLSClientAuthGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginTLSClientAuthGetResponse                `json:"result"`
	JSON    originTLSClientAuthGetResponseEnvelopeJSON    `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginTLSClientAuthGetResponseEnvelope]
type originTLSClientAuthGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           OriginTLSClientAuthGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             originTLSClientAuthGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [OriginTLSClientAuthGetResponseEnvelopeErrors]
type originTLSClientAuthGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    originTLSClientAuthGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthGetResponseEnvelopeErrorsSource]
type originTLSClientAuthGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           OriginTLSClientAuthGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             originTLSClientAuthGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [OriginTLSClientAuthGetResponseEnvelopeMessages]
type originTLSClientAuthGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OriginTLSClientAuthGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    originTLSClientAuthGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// originTLSClientAuthGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [OriginTLSClientAuthGetResponseEnvelopeMessagesSource]
type originTLSClientAuthGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginTLSClientAuthGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originTLSClientAuthGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginTLSClientAuthGetResponseEnvelopeSuccess bool

const (
	OriginTLSClientAuthGetResponseEnvelopeSuccessTrue OriginTLSClientAuthGetResponseEnvelopeSuccess = true
)

func (r OriginTLSClientAuthGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginTLSClientAuthGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
