// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

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

// KeyService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyService] method instead.
type KeyService struct {
	Options []option.RequestOption
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r *KeyService) {
	r = &KeyService{}
	r.Options = opts
	return
}

// Creates an RSA private key in PEM and JWK formats. Key files are only displayed
// once after creation. Keys are created, used, and deleted independently of
// videos, and every key can sign any video.
func (r *KeyService) New(ctx context.Context, params KeyNewParams, opts ...option.RequestOption) (res *Keys, err error) {
	var env KeyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes signing keys and revokes all signed URLs generated with the key.
func (r *KeyService) Delete(ctx context.Context, identifier string, body KeyDeleteParams, opts ...option.RequestOption) (res *string, err error) {
	var env KeyDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys/%s", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the video ID and creation date and time when a signing key was created.
func (r *KeyService) Get(ctx context.Context, query KeyGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[KeyGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/keys", query.AccountID)
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

// Lists the video ID and creation date and time when a signing key was created.
func (r *KeyService) GetAutoPaging(ctx context.Context, query KeyGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[KeyGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

type Keys struct {
	// Identifier.
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time `json:"created" format:"date-time"`
	// The signing key in JWK format.
	Jwk string `json:"jwk"`
	// The signing key in PEM format.
	Pem  string   `json:"pem"`
	JSON keysJSON `json:"-"`
}

// keysJSON contains the JSON metadata for the struct [Keys]
type keysJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Jwk         apijson.Field
	Pem         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Keys) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keysJSON) RawJSON() string {
	return r.raw
}

type KeyGetResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The date and time a signing key was created.
	Created time.Time          `json:"created" format:"date-time"`
	JSON    keyGetResponseJSON `json:"-"`
}

// keyGetResponseJSON contains the JSON metadata for the struct [KeyGetResponse]
type keyGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyGetResponseJSON) RawJSON() string {
	return r.raw
}

type KeyNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r KeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type KeyNewResponseEnvelope struct {
	Errors   []KeyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeyNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success KeyNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Keys                          `json:"result"`
	JSON    keyNewResponseEnvelopeJSON    `json:"-"`
}

// keyNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelope]
type keyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeyNewResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           KeyNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             keyNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// keyNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelopeErrors]
type keyNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeyNewResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    keyNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// keyNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelopeErrorsSource]
type keyNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type KeyNewResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           KeyNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             keyNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// keyNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [KeyNewResponseEnvelopeMessages]
type keyNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type KeyNewResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    keyNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// keyNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [KeyNewResponseEnvelopeMessagesSource]
type keyNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type KeyNewResponseEnvelopeSuccess bool

const (
	KeyNewResponseEnvelopeSuccessTrue KeyNewResponseEnvelopeSuccess = true
)

func (r KeyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type KeyDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type KeyDeleteResponseEnvelope struct {
	Errors   []KeyDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeyDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success KeyDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  string                           `json:"result"`
	JSON    keyDeleteResponseEnvelopeJSON    `json:"-"`
}

// keyDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [KeyDeleteResponseEnvelope]
type keyDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeyDeleteResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           KeyDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             keyDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// keyDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [KeyDeleteResponseEnvelopeErrors]
type keyDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeyDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    keyDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// keyDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [KeyDeleteResponseEnvelopeErrorsSource]
type keyDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type KeyDeleteResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           KeyDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             keyDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// keyDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [KeyDeleteResponseEnvelopeMessages]
type keyDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type KeyDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    keyDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// keyDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [KeyDeleteResponseEnvelopeMessagesSource]
type keyDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type KeyDeleteResponseEnvelopeSuccess bool

const (
	KeyDeleteResponseEnvelopeSuccessTrue KeyDeleteResponseEnvelopeSuccess = true
)

func (r KeyDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case KeyDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type KeyGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
