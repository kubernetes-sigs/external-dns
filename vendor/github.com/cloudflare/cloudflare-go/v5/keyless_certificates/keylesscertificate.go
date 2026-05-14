// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificates

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/custom_hostnames"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// KeylessCertificateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeylessCertificateService] method instead.
type KeylessCertificateService struct {
	Options []option.RequestOption
}

// NewKeylessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewKeylessCertificateService(opts ...option.RequestOption) (r *KeylessCertificateService) {
	r = &KeylessCertificateService{}
	r.Options = opts
	return
}

// Create Keyless SSL Configuration
func (r *KeylessCertificateService) New(ctx context.Context, params KeylessCertificateNewParams, opts ...option.RequestOption) (res *KeylessCertificate, err error) {
	var env KeylessCertificateNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Keyless SSL configurations for a given zone.
func (r *KeylessCertificateService) List(ctx context.Context, query KeylessCertificateListParams, opts ...option.RequestOption) (res *pagination.SinglePage[KeylessCertificate], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates", query.ZoneID)
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

// List all Keyless SSL configurations for a given zone.
func (r *KeylessCertificateService) ListAutoPaging(ctx context.Context, query KeylessCertificateListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[KeylessCertificate] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Keyless SSL Configuration
func (r *KeylessCertificateService) Delete(ctx context.Context, keylessCertificateID string, body KeylessCertificateDeleteParams, opts ...option.RequestOption) (res *KeylessCertificateDeleteResponse, err error) {
	var env KeylessCertificateDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if keylessCertificateID == "" {
		err = errors.New("missing required keyless_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", body.ZoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This will update attributes of a Keyless SSL. Consists of one or more of the
// following: host,name,port.
func (r *KeylessCertificateService) Edit(ctx context.Context, keylessCertificateID string, params KeylessCertificateEditParams, opts ...option.RequestOption) (res *KeylessCertificate, err error) {
	var env KeylessCertificateEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if keylessCertificateID == "" {
		err = errors.New("missing required keyless_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", params.ZoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for one Keyless SSL configuration.
func (r *KeylessCertificateService) Get(ctx context.Context, keylessCertificateID string, query KeylessCertificateGetParams, opts ...option.RequestOption) (res *KeylessCertificate, err error) {
	var env KeylessCertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if keylessCertificateID == "" {
		err = errors.New("missing required keyless_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", query.ZoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type KeylessCertificate struct {
	// Keyless certificate identifier tag.
	ID string `json:"id,required"`
	// When the Keyless SSL was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Whether or not the Keyless SSL is on or off.
	Enabled bool `json:"enabled,required"`
	// The keyless SSL name.
	Host string `json:"host,required" format:"hostname"`
	// When the Keyless SSL was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The keyless SSL name.
	Name string `json:"name,required"`
	// Available permissions for the Keyless SSL for the current user requesting the
	// item.
	Permissions []string `json:"permissions,required"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port float64 `json:"port,required"`
	// Status of the Keyless SSL.
	Status KeylessCertificateStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel Tunnel                 `json:"tunnel"`
	JSON   keylessCertificateJSON `json:"-"`
}

// keylessCertificateJSON contains the JSON metadata for the struct
// [KeylessCertificate]
type keylessCertificateJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	Host        apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Port        apijson.Field
	Status      apijson.Field
	Tunnel      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateJSON) RawJSON() string {
	return r.raw
}

// Status of the Keyless SSL.
type KeylessCertificateStatus string

const (
	KeylessCertificateStatusActive  KeylessCertificateStatus = "active"
	KeylessCertificateStatusDeleted KeylessCertificateStatus = "deleted"
)

func (r KeylessCertificateStatus) IsKnown() bool {
	switch r {
	case KeylessCertificateStatusActive, KeylessCertificateStatusDeleted:
		return true
	}
	return false
}

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type Tunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string     `json:"vnet_id,required"`
	JSON   tunnelJSON `json:"-"`
}

// tunnelJSON contains the JSON metadata for the struct [Tunnel]
type tunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Tunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelJSON) RawJSON() string {
	return r.raw
}

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type TunnelParam struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r TunnelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeylessCertificateDeleteResponse struct {
	// Identifier.
	ID   string                               `json:"id"`
	JSON keylessCertificateDeleteResponseJSON `json:"-"`
}

// keylessCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [KeylessCertificateDeleteResponse]
type keylessCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The zone's SSL certificate or SSL certificate and intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[custom_hostnames.BundleMethod] `json:"bundle_method"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[TunnelParam] `json:"tunnel"`
}

func (r KeylessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeylessCertificateNewResponseEnvelope struct {
	Errors   []KeylessCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success KeylessCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  KeylessCertificate                           `json:"result"`
	JSON    keylessCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateNewResponseEnvelope]
type keylessCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           KeylessCertificateNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             keylessCertificateNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// keylessCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KeylessCertificateNewResponseEnvelopeErrors]
type keylessCertificateNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    keylessCertificateNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// keylessCertificateNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [KeylessCertificateNewResponseEnvelopeErrorsSource]
type keylessCertificateNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           KeylessCertificateNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             keylessCertificateNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// keylessCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [KeylessCertificateNewResponseEnvelopeMessages]
type keylessCertificateNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    keylessCertificateNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// keylessCertificateNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [KeylessCertificateNewResponseEnvelopeMessagesSource]
type keylessCertificateNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type KeylessCertificateNewResponseEnvelopeSuccess bool

const (
	KeylessCertificateNewResponseEnvelopeSuccessTrue KeylessCertificateNewResponseEnvelopeSuccess = true
)

func (r KeylessCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeylessCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type KeylessCertificateListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type KeylessCertificateDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type KeylessCertificateDeleteResponseEnvelope struct {
	Errors   []KeylessCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success KeylessCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  KeylessCertificateDeleteResponse                `json:"result"`
	JSON    keylessCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateDeleteResponseEnvelope]
type keylessCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           KeylessCertificateDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             keylessCertificateDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [KeylessCertificateDeleteResponseEnvelopeErrors]
type keylessCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    keylessCertificateDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [KeylessCertificateDeleteResponseEnvelopeErrorsSource]
type keylessCertificateDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           KeylessCertificateDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             keylessCertificateDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KeylessCertificateDeleteResponseEnvelopeMessages]
type keylessCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    keylessCertificateDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [KeylessCertificateDeleteResponseEnvelopeMessagesSource]
type keylessCertificateDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type KeylessCertificateDeleteResponseEnvelopeSuccess bool

const (
	KeylessCertificateDeleteResponseEnvelopeSuccessTrue KeylessCertificateDeleteResponseEnvelopeSuccess = true
)

func (r KeylessCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeylessCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type KeylessCertificateEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether or not the Keyless SSL is on or off.
	Enabled param.Field[bool] `json:"enabled"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host" format:"hostname"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[TunnelParam] `json:"tunnel"`
}

func (r KeylessCertificateEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeylessCertificateEditResponseEnvelope struct {
	Errors   []KeylessCertificateEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success KeylessCertificateEditResponseEnvelopeSuccess `json:"success,required"`
	Result  KeylessCertificate                            `json:"result"`
	JSON    keylessCertificateEditResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateEditResponseEnvelope]
type keylessCertificateEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateEditResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           KeylessCertificateEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             keylessCertificateEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// keylessCertificateEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KeylessCertificateEditResponseEnvelopeErrors]
type keylessCertificateEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    keylessCertificateEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// keylessCertificateEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [KeylessCertificateEditResponseEnvelopeErrorsSource]
type keylessCertificateEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateEditResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           KeylessCertificateEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             keylessCertificateEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// keylessCertificateEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KeylessCertificateEditResponseEnvelopeMessages]
type keylessCertificateEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    keylessCertificateEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// keylessCertificateEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [KeylessCertificateEditResponseEnvelopeMessagesSource]
type keylessCertificateEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type KeylessCertificateEditResponseEnvelopeSuccess bool

const (
	KeylessCertificateEditResponseEnvelopeSuccessTrue KeylessCertificateEditResponseEnvelopeSuccess = true
)

func (r KeylessCertificateEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeylessCertificateEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type KeylessCertificateGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type KeylessCertificateGetResponseEnvelope struct {
	Errors   []KeylessCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success KeylessCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  KeylessCertificate                           `json:"result"`
	JSON    keylessCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateGetResponseEnvelope]
type keylessCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateGetResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           KeylessCertificateGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             keylessCertificateGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// keylessCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KeylessCertificateGetResponseEnvelopeErrors]
type keylessCertificateGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    keylessCertificateGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// keylessCertificateGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [KeylessCertificateGetResponseEnvelopeErrorsSource]
type keylessCertificateGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateGetResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           KeylessCertificateGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             keylessCertificateGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// keylessCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [KeylessCertificateGetResponseEnvelopeMessages]
type keylessCertificateGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    keylessCertificateGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// keylessCertificateGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [KeylessCertificateGetResponseEnvelopeMessagesSource]
type keylessCertificateGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type KeylessCertificateGetResponseEnvelopeSuccess bool

const (
	KeylessCertificateGetResponseEnvelopeSuccessTrue KeylessCertificateGetResponseEnvelopeSuccess = true
)

func (r KeylessCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeylessCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
