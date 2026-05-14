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

// HostnameCertificateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameCertificateService] method instead.
type HostnameCertificateService struct {
	Options []option.RequestOption
}

// NewHostnameCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHostnameCertificateService(opts ...option.RequestOption) (r *HostnameCertificateService) {
	r = &HostnameCertificateService{}
	r.Options = opts
	return
}

// Upload a certificate to be used for client authentication on a hostname. 10
// hostname certificates per zone are allowed.
func (r *HostnameCertificateService) New(ctx context.Context, params HostnameCertificateNewParams, opts ...option.RequestOption) (res *HostnameCertificateNewResponse, err error) {
	var env HostnameCertificateNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Certificates
func (r *HostnameCertificateService) List(ctx context.Context, query HostnameCertificateListParams, opts ...option.RequestOption) (res *pagination.SinglePage[HostnameCertificateListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates", query.ZoneID)
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
func (r *HostnameCertificateService) ListAutoPaging(ctx context.Context, query HostnameCertificateListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[HostnameCertificateListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Hostname Client Certificate
func (r *HostnameCertificateService) Delete(ctx context.Context, certificateID string, body HostnameCertificateDeleteParams, opts ...option.RequestOption) (res *HostnameCertificateDeleteResponse, err error) {
	var env HostnameCertificateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", body.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the certificate by ID to be used for client authentication on a hostname.
func (r *HostnameCertificateService) Get(ctx context.Context, certificateID string, query HostnameCertificateGetParams, opts ...option.RequestOption) (res *HostnameCertificateGetResponse, err error) {
	var env HostnameCertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/certificates/%s", query.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HostnameCertificateNewResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status HostnameCertificateNewResponseStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                          `json:"uploaded_on" format:"date-time"`
	JSON       hostnameCertificateNewResponseJSON `json:"-"`
}

// hostnameCertificateNewResponseJSON contains the JSON metadata for the struct
// [HostnameCertificateNewResponse]
type hostnameCertificateNewResponseJSON struct {
	ID           apijson.Field
	Certificate  apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	Status       apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HostnameCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type HostnameCertificateNewResponseStatus string

const (
	HostnameCertificateNewResponseStatusInitializing       HostnameCertificateNewResponseStatus = "initializing"
	HostnameCertificateNewResponseStatusPendingDeployment  HostnameCertificateNewResponseStatus = "pending_deployment"
	HostnameCertificateNewResponseStatusPendingDeletion    HostnameCertificateNewResponseStatus = "pending_deletion"
	HostnameCertificateNewResponseStatusActive             HostnameCertificateNewResponseStatus = "active"
	HostnameCertificateNewResponseStatusDeleted            HostnameCertificateNewResponseStatus = "deleted"
	HostnameCertificateNewResponseStatusDeploymentTimedOut HostnameCertificateNewResponseStatus = "deployment_timed_out"
	HostnameCertificateNewResponseStatusDeletionTimedOut   HostnameCertificateNewResponseStatus = "deletion_timed_out"
)

func (r HostnameCertificateNewResponseStatus) IsKnown() bool {
	switch r {
	case HostnameCertificateNewResponseStatusInitializing, HostnameCertificateNewResponseStatusPendingDeployment, HostnameCertificateNewResponseStatusPendingDeletion, HostnameCertificateNewResponseStatusActive, HostnameCertificateNewResponseStatusDeleted, HostnameCertificateNewResponseStatusDeploymentTimedOut, HostnameCertificateNewResponseStatusDeletionTimedOut:
		return true
	}
	return false
}

type HostnameCertificateListResponse struct {
	// Identifier.
	ID string `json:"id"`
	// Identifier.
	CERTID string `json:"cert_id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The hostname certificate's private key.
	PrivateKey string                              `json:"private_key"`
	JSON       hostnameCertificateListResponseJSON `json:"-"`
	AuthenticatedOriginPull
}

// hostnameCertificateListResponseJSON contains the JSON metadata for the struct
// [HostnameCertificateListResponse]
type hostnameCertificateListResponseJSON struct {
	ID          apijson.Field
	CERTID      apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	Hostname    apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateDeleteResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status HostnameCertificateDeleteResponseStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                             `json:"uploaded_on" format:"date-time"`
	JSON       hostnameCertificateDeleteResponseJSON `json:"-"`
}

// hostnameCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [HostnameCertificateDeleteResponse]
type hostnameCertificateDeleteResponseJSON struct {
	ID           apijson.Field
	Certificate  apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	Status       apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type HostnameCertificateDeleteResponseStatus string

const (
	HostnameCertificateDeleteResponseStatusInitializing       HostnameCertificateDeleteResponseStatus = "initializing"
	HostnameCertificateDeleteResponseStatusPendingDeployment  HostnameCertificateDeleteResponseStatus = "pending_deployment"
	HostnameCertificateDeleteResponseStatusPendingDeletion    HostnameCertificateDeleteResponseStatus = "pending_deletion"
	HostnameCertificateDeleteResponseStatusActive             HostnameCertificateDeleteResponseStatus = "active"
	HostnameCertificateDeleteResponseStatusDeleted            HostnameCertificateDeleteResponseStatus = "deleted"
	HostnameCertificateDeleteResponseStatusDeploymentTimedOut HostnameCertificateDeleteResponseStatus = "deployment_timed_out"
	HostnameCertificateDeleteResponseStatusDeletionTimedOut   HostnameCertificateDeleteResponseStatus = "deletion_timed_out"
)

func (r HostnameCertificateDeleteResponseStatus) IsKnown() bool {
	switch r {
	case HostnameCertificateDeleteResponseStatusInitializing, HostnameCertificateDeleteResponseStatusPendingDeployment, HostnameCertificateDeleteResponseStatusPendingDeletion, HostnameCertificateDeleteResponseStatusActive, HostnameCertificateDeleteResponseStatusDeleted, HostnameCertificateDeleteResponseStatusDeploymentTimedOut, HostnameCertificateDeleteResponseStatusDeletionTimedOut:
		return true
	}
	return false
}

type HostnameCertificateGetResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status HostnameCertificateGetResponseStatus `json:"status"`
	// The time when the certificate was uploaded.
	UploadedOn time.Time                          `json:"uploaded_on" format:"date-time"`
	JSON       hostnameCertificateGetResponseJSON `json:"-"`
}

// hostnameCertificateGetResponseJSON contains the JSON metadata for the struct
// [HostnameCertificateGetResponse]
type hostnameCertificateGetResponseJSON struct {
	ID           apijson.Field
	Certificate  apijson.Field
	ExpiresOn    apijson.Field
	Issuer       apijson.Field
	SerialNumber apijson.Field
	Signature    apijson.Field
	Status       apijson.Field
	UploadedOn   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HostnameCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type HostnameCertificateGetResponseStatus string

const (
	HostnameCertificateGetResponseStatusInitializing       HostnameCertificateGetResponseStatus = "initializing"
	HostnameCertificateGetResponseStatusPendingDeployment  HostnameCertificateGetResponseStatus = "pending_deployment"
	HostnameCertificateGetResponseStatusPendingDeletion    HostnameCertificateGetResponseStatus = "pending_deletion"
	HostnameCertificateGetResponseStatusActive             HostnameCertificateGetResponseStatus = "active"
	HostnameCertificateGetResponseStatusDeleted            HostnameCertificateGetResponseStatus = "deleted"
	HostnameCertificateGetResponseStatusDeploymentTimedOut HostnameCertificateGetResponseStatus = "deployment_timed_out"
	HostnameCertificateGetResponseStatusDeletionTimedOut   HostnameCertificateGetResponseStatus = "deletion_timed_out"
)

func (r HostnameCertificateGetResponseStatus) IsKnown() bool {
	switch r {
	case HostnameCertificateGetResponseStatusInitializing, HostnameCertificateGetResponseStatusPendingDeployment, HostnameCertificateGetResponseStatusPendingDeletion, HostnameCertificateGetResponseStatusActive, HostnameCertificateGetResponseStatusDeleted, HostnameCertificateGetResponseStatusDeploymentTimedOut, HostnameCertificateGetResponseStatusDeletionTimedOut:
		return true
	}
	return false
}

type HostnameCertificateNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The hostname certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The hostname certificate's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r HostnameCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameCertificateNewResponseEnvelope struct {
	Errors   []HostnameCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success HostnameCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  HostnameCertificateNewResponse                `json:"result"`
	JSON    hostnameCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameCertificateNewResponseEnvelope]
type hostnameCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateNewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           HostnameCertificateNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             hostnameCertificateNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HostnameCertificateNewResponseEnvelopeErrors]
type hostnameCertificateNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    hostnameCertificateNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [HostnameCertificateNewResponseEnvelopeErrorsSource]
type hostnameCertificateNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateNewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           HostnameCertificateNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             hostnameCertificateNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameCertificateNewResponseEnvelopeMessages]
type hostnameCertificateNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    hostnameCertificateNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// hostnameCertificateNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [HostnameCertificateNewResponseEnvelopeMessagesSource]
type hostnameCertificateNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type HostnameCertificateNewResponseEnvelopeSuccess bool

const (
	HostnameCertificateNewResponseEnvelopeSuccessTrue HostnameCertificateNewResponseEnvelopeSuccess = true
)

func (r HostnameCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameCertificateListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameCertificateDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameCertificateDeleteResponseEnvelope struct {
	Errors   []HostnameCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success HostnameCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  HostnameCertificateDeleteResponse                `json:"result"`
	JSON    hostnameCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameCertificateDeleteResponseEnvelope]
type hostnameCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateDeleteResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           HostnameCertificateDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             hostnameCertificateDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [HostnameCertificateDeleteResponseEnvelopeErrors]
type hostnameCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    hostnameCertificateDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [HostnameCertificateDeleteResponseEnvelopeErrorsSource]
type hostnameCertificateDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateDeleteResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           HostnameCertificateDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             hostnameCertificateDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameCertificateDeleteResponseEnvelopeMessages]
type hostnameCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    hostnameCertificateDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// hostnameCertificateDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [HostnameCertificateDeleteResponseEnvelopeMessagesSource]
type hostnameCertificateDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type HostnameCertificateDeleteResponseEnvelopeSuccess bool

const (
	HostnameCertificateDeleteResponseEnvelopeSuccessTrue HostnameCertificateDeleteResponseEnvelopeSuccess = true
)

func (r HostnameCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameCertificateGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameCertificateGetResponseEnvelope struct {
	Errors   []HostnameCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []HostnameCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success HostnameCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  HostnameCertificateGetResponse                `json:"result"`
	JSON    hostnameCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HostnameCertificateGetResponseEnvelope]
type hostnameCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           HostnameCertificateGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             hostnameCertificateGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [HostnameCertificateGetResponseEnvelopeErrors]
type hostnameCertificateGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    hostnameCertificateGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [HostnameCertificateGetResponseEnvelopeErrorsSource]
type hostnameCertificateGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           HostnameCertificateGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             hostnameCertificateGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [HostnameCertificateGetResponseEnvelopeMessages]
type hostnameCertificateGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type HostnameCertificateGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    hostnameCertificateGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// hostnameCertificateGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [HostnameCertificateGetResponseEnvelopeMessagesSource]
type hostnameCertificateGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameCertificateGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameCertificateGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type HostnameCertificateGetResponseEnvelopeSuccess bool

const (
	HostnameCertificateGetResponseEnvelopeSuccessTrue HostnameCertificateGetResponseEnvelopeSuccess = true
)

func (r HostnameCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
