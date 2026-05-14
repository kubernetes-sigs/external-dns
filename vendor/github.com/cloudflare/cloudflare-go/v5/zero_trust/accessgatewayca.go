// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// AccessGatewayCAService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessGatewayCAService] method instead.
type AccessGatewayCAService struct {
	Options []option.RequestOption
}

// NewAccessGatewayCAService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessGatewayCAService(opts ...option.RequestOption) (r *AccessGatewayCAService) {
	r = &AccessGatewayCAService{}
	r.Options = opts
	return
}

// Adds a new SSH Certificate Authority (CA).
func (r *AccessGatewayCAService) New(ctx context.Context, body AccessGatewayCANewParams, opts ...option.RequestOption) (res *AccessGatewayCANewResponse, err error) {
	var env AccessGatewayCANewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/gateway_ca", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists SSH Certificate Authorities (CA).
func (r *AccessGatewayCAService) List(ctx context.Context, query AccessGatewayCAListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccessGatewayCAListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/gateway_ca", query.AccountID)
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

// Lists SSH Certificate Authorities (CA).
func (r *AccessGatewayCAService) ListAutoPaging(ctx context.Context, query AccessGatewayCAListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessGatewayCAListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an SSH Certificate Authority.
func (r *AccessGatewayCAService) Delete(ctx context.Context, certificateID string, body AccessGatewayCADeleteParams, opts ...option.RequestOption) (res *AccessGatewayCADeleteResponse, err error) {
	var env AccessGatewayCADeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/gateway_ca/%s", body.AccountID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessGatewayCANewResponse struct {
	// The key ID of this certificate.
	ID string `json:"id"`
	// The public key of this certificate.
	PublicKey string                         `json:"public_key"`
	JSON      accessGatewayCANewResponseJSON `json:"-"`
}

// accessGatewayCANewResponseJSON contains the JSON metadata for the struct
// [AccessGatewayCANewResponse]
type accessGatewayCANewResponseJSON struct {
	ID          apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCANewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCANewResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCAListResponse struct {
	// The key ID of this certificate.
	ID string `json:"id"`
	// The public key of this certificate.
	PublicKey string                          `json:"public_key"`
	JSON      accessGatewayCAListResponseJSON `json:"-"`
}

// accessGatewayCAListResponseJSON contains the JSON metadata for the struct
// [AccessGatewayCAListResponse]
type accessGatewayCAListResponseJSON struct {
	ID          apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCAListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCAListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCADeleteResponse struct {
	// UUID.
	ID   string                            `json:"id"`
	JSON accessGatewayCADeleteResponseJSON `json:"-"`
}

// accessGatewayCADeleteResponseJSON contains the JSON metadata for the struct
// [AccessGatewayCADeleteResponse]
type accessGatewayCADeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCADeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCADeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCANewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessGatewayCANewResponseEnvelope struct {
	Errors   []AccessGatewayCANewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGatewayCANewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessGatewayCANewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessGatewayCANewResponse                `json:"result"`
	JSON    accessGatewayCANewResponseEnvelopeJSON    `json:"-"`
}

// accessGatewayCANewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessGatewayCANewResponseEnvelope]
type accessGatewayCANewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCANewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCANewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCANewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AccessGatewayCANewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessGatewayCANewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessGatewayCANewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessGatewayCANewResponseEnvelopeErrors]
type accessGatewayCANewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGatewayCANewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCANewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCANewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    accessGatewayCANewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessGatewayCANewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [AccessGatewayCANewResponseEnvelopeErrorsSource]
type accessGatewayCANewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCANewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCANewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCANewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccessGatewayCANewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessGatewayCANewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessGatewayCANewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessGatewayCANewResponseEnvelopeMessages]
type accessGatewayCANewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGatewayCANewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCANewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCANewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accessGatewayCANewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessGatewayCANewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccessGatewayCANewResponseEnvelopeMessagesSource]
type accessGatewayCANewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCANewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCANewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessGatewayCANewResponseEnvelopeSuccess bool

const (
	AccessGatewayCANewResponseEnvelopeSuccessTrue AccessGatewayCANewResponseEnvelopeSuccess = true
)

func (r AccessGatewayCANewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGatewayCANewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessGatewayCAListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessGatewayCADeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessGatewayCADeleteResponseEnvelope struct {
	Errors   []AccessGatewayCADeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessGatewayCADeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessGatewayCADeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessGatewayCADeleteResponse                `json:"result"`
	JSON    accessGatewayCADeleteResponseEnvelopeJSON    `json:"-"`
}

// accessGatewayCADeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessGatewayCADeleteResponseEnvelope]
type accessGatewayCADeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCADeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCADeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCADeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           AccessGatewayCADeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessGatewayCADeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessGatewayCADeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessGatewayCADeleteResponseEnvelopeErrors]
type accessGatewayCADeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGatewayCADeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCADeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCADeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    accessGatewayCADeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessGatewayCADeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [AccessGatewayCADeleteResponseEnvelopeErrorsSource]
type accessGatewayCADeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCADeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCADeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCADeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AccessGatewayCADeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessGatewayCADeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessGatewayCADeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessGatewayCADeleteResponseEnvelopeMessages]
type accessGatewayCADeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessGatewayCADeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCADeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessGatewayCADeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    accessGatewayCADeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessGatewayCADeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessGatewayCADeleteResponseEnvelopeMessagesSource]
type accessGatewayCADeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessGatewayCADeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessGatewayCADeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessGatewayCADeleteResponseEnvelopeSuccess bool

const (
	AccessGatewayCADeleteResponseEnvelopeSuccessTrue AccessGatewayCADeleteResponseEnvelopeSuccess = true
)

func (r AccessGatewayCADeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessGatewayCADeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
