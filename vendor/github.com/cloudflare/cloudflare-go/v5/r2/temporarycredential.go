// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// TemporaryCredentialService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemporaryCredentialService] method instead.
type TemporaryCredentialService struct {
	Options []option.RequestOption
}

// NewTemporaryCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTemporaryCredentialService(opts ...option.RequestOption) (r *TemporaryCredentialService) {
	r = &TemporaryCredentialService{}
	r.Options = opts
	return
}

// Creates temporary access credentials on a bucket that can be optionally scoped
// to prefixes or objects.
func (r *TemporaryCredentialService) New(ctx context.Context, params TemporaryCredentialNewParams, opts ...option.RequestOption) (res *TemporaryCredentialNewResponse, err error) {
	var env TemporaryCredentialNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/r2/temp-access-credentials", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TemporaryCredentialParam struct {
	// Name of the R2 bucket.
	Bucket param.Field[string] `json:"bucket,required"`
	// The parent access key id to use for signing.
	ParentAccessKeyID param.Field[string] `json:"parentAccessKeyId,required"`
	// Permissions allowed on the credentials.
	Permission param.Field[TemporaryCredentialPermission] `json:"permission,required"`
	// How long the credentials will live for in seconds.
	TTLSeconds param.Field[float64] `json:"ttlSeconds,required"`
	// Optional object paths to scope the credentials to.
	Objects param.Field[[]string] `json:"objects"`
	// Optional prefix paths to scope the credentials to.
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r TemporaryCredentialParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Permissions allowed on the credentials.
type TemporaryCredentialPermission string

const (
	TemporaryCredentialPermissionAdminReadWrite  TemporaryCredentialPermission = "admin-read-write"
	TemporaryCredentialPermissionAdminReadOnly   TemporaryCredentialPermission = "admin-read-only"
	TemporaryCredentialPermissionObjectReadWrite TemporaryCredentialPermission = "object-read-write"
	TemporaryCredentialPermissionObjectReadOnly  TemporaryCredentialPermission = "object-read-only"
)

func (r TemporaryCredentialPermission) IsKnown() bool {
	switch r {
	case TemporaryCredentialPermissionAdminReadWrite, TemporaryCredentialPermissionAdminReadOnly, TemporaryCredentialPermissionObjectReadWrite, TemporaryCredentialPermissionObjectReadOnly:
		return true
	}
	return false
}

type TemporaryCredentialNewResponse struct {
	// ID for new access key.
	AccessKeyID string `json:"accessKeyId"`
	// Secret access key.
	SecretAccessKey string `json:"secretAccessKey"`
	// Security token.
	SessionToken string                             `json:"sessionToken"`
	JSON         temporaryCredentialNewResponseJSON `json:"-"`
}

// temporaryCredentialNewResponseJSON contains the JSON metadata for the struct
// [TemporaryCredentialNewResponse]
type temporaryCredentialNewResponseJSON struct {
	AccessKeyID     apijson.Field
	SecretAccessKey apijson.Field
	SessionToken    apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TemporaryCredentialNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r temporaryCredentialNewResponseJSON) RawJSON() string {
	return r.raw
}

type TemporaryCredentialNewParams struct {
	// Account ID.
	AccountID           param.Field[string]      `path:"account_id,required"`
	TemporaryCredential TemporaryCredentialParam `json:"temporary_credential,required"`
}

func (r TemporaryCredentialNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.TemporaryCredential)
}

type TemporaryCredentialNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []string                       `json:"messages,required"`
	Result   TemporaryCredentialNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success TemporaryCredentialNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    temporaryCredentialNewResponseEnvelopeJSON    `json:"-"`
}

// temporaryCredentialNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TemporaryCredentialNewResponseEnvelope]
type temporaryCredentialNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TemporaryCredentialNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r temporaryCredentialNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TemporaryCredentialNewResponseEnvelopeSuccess bool

const (
	TemporaryCredentialNewResponseEnvelopeSuccessTrue TemporaryCredentialNewResponseEnvelopeSuccess = true
)

func (r TemporaryCredentialNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TemporaryCredentialNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
