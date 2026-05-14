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
)

// PrefixAdvertisementStatusService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixAdvertisementStatusService] method instead.
type PrefixAdvertisementStatusService struct {
	Options []option.RequestOption
}

// NewPrefixAdvertisementStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPrefixAdvertisementStatusService(opts ...option.RequestOption) (r *PrefixAdvertisementStatusService) {
	r = &PrefixAdvertisementStatusService{}
	r.Options = opts
	return
}

// Advertise or withdraw the BGP route for a prefix.
//
// **Deprecated:** Prefer the BGP Prefixes endpoints, which additionally allow for
// advertising and withdrawing subnets of an IP prefix.
//
// Deprecated: deprecated
func (r *PrefixAdvertisementStatusService) Edit(ctx context.Context, prefixID string, params PrefixAdvertisementStatusEditParams, opts ...option.RequestOption) (res *PrefixAdvertisementStatusEditResponse, err error) {
	var env PrefixAdvertisementStatusEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// View the current advertisement state for a prefix.
//
// **Deprecated:** Prefer the BGP Prefixes endpoints, which additionally allow for
// advertising and withdrawing subnets of an IP prefix.
//
// Deprecated: deprecated
func (r *PrefixAdvertisementStatusService) Get(ctx context.Context, prefixID string, query PrefixAdvertisementStatusGetParams, opts ...option.RequestOption) (res *PrefixAdvertisementStatusGetResponse, err error) {
	var env PrefixAdvertisementStatusGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/status", query.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PrefixAdvertisementStatusEditResponse struct {
	// Advertisement status of the prefix. If `true`, the BGP route for the prefix is
	// advertised to the Internet. If `false`, the BGP route is withdrawn.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                 `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixAdvertisementStatusEditResponseJSON `json:"-"`
}

// prefixAdvertisementStatusEditResponseJSON contains the JSON metadata for the
// struct [PrefixAdvertisementStatusEditResponse]
type prefixAdvertisementStatusEditResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusGetResponse struct {
	// Advertisement status of the prefix. If `true`, the BGP route for the prefix is
	// advertised to the Internet. If `false`, the BGP route is withdrawn.
	Advertised bool `json:"advertised"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time                                `json:"advertised_modified_at,nullable" format:"date-time"`
	JSON                 prefixAdvertisementStatusGetResponseJSON `json:"-"`
}

// prefixAdvertisementStatusGetResponseJSON contains the JSON metadata for the
// struct [PrefixAdvertisementStatusGetResponse]
type prefixAdvertisementStatusGetResponseJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusEditParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Advertisement status of the prefix. If `true`, the BGP route for the prefix is
	// advertised to the Internet. If `false`, the BGP route is withdrawn.
	Advertised param.Field[bool] `json:"advertised,required"`
}

func (r PrefixAdvertisementStatusEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixAdvertisementStatusEditResponseEnvelope struct {
	Errors   []PrefixAdvertisementStatusEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixAdvertisementStatusEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success PrefixAdvertisementStatusEditResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixAdvertisementStatusEditResponse                `json:"result"`
	JSON    prefixAdvertisementStatusEditResponseEnvelopeJSON    `json:"-"`
}

// prefixAdvertisementStatusEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [PrefixAdvertisementStatusEditResponseEnvelope]
type prefixAdvertisementStatusEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusEditResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           PrefixAdvertisementStatusEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             prefixAdvertisementStatusEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// prefixAdvertisementStatusEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [PrefixAdvertisementStatusEditResponseEnvelopeErrors]
type prefixAdvertisementStatusEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    prefixAdvertisementStatusEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// prefixAdvertisementStatusEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [PrefixAdvertisementStatusEditResponseEnvelopeErrorsSource]
type prefixAdvertisementStatusEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusEditResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           PrefixAdvertisementStatusEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             prefixAdvertisementStatusEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// prefixAdvertisementStatusEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [PrefixAdvertisementStatusEditResponseEnvelopeMessages]
type prefixAdvertisementStatusEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    prefixAdvertisementStatusEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// prefixAdvertisementStatusEditResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [PrefixAdvertisementStatusEditResponseEnvelopeMessagesSource]
type prefixAdvertisementStatusEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrefixAdvertisementStatusEditResponseEnvelopeSuccess bool

const (
	PrefixAdvertisementStatusEditResponseEnvelopeSuccessTrue PrefixAdvertisementStatusEditResponseEnvelopeSuccess = true
)

func (r PrefixAdvertisementStatusEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixAdvertisementStatusEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixAdvertisementStatusGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixAdvertisementStatusGetResponseEnvelope struct {
	Errors   []PrefixAdvertisementStatusGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixAdvertisementStatusGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success PrefixAdvertisementStatusGetResponseEnvelopeSuccess `json:"success,required"`
	Result  PrefixAdvertisementStatusGetResponse                `json:"result"`
	JSON    prefixAdvertisementStatusGetResponseEnvelopeJSON    `json:"-"`
}

// prefixAdvertisementStatusGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [PrefixAdvertisementStatusGetResponseEnvelope]
type prefixAdvertisementStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusGetResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           PrefixAdvertisementStatusGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             prefixAdvertisementStatusGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// prefixAdvertisementStatusGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [PrefixAdvertisementStatusGetResponseEnvelopeErrors]
type prefixAdvertisementStatusGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    prefixAdvertisementStatusGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// prefixAdvertisementStatusGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [PrefixAdvertisementStatusGetResponseEnvelopeErrorsSource]
type prefixAdvertisementStatusGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusGetResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           PrefixAdvertisementStatusGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             prefixAdvertisementStatusGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// prefixAdvertisementStatusGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [PrefixAdvertisementStatusGetResponseEnvelopeMessages]
type prefixAdvertisementStatusGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrefixAdvertisementStatusGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    prefixAdvertisementStatusGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// prefixAdvertisementStatusGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [PrefixAdvertisementStatusGetResponseEnvelopeMessagesSource]
type prefixAdvertisementStatusGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixAdvertisementStatusGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixAdvertisementStatusGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrefixAdvertisementStatusGetResponseEnvelopeSuccess bool

const (
	PrefixAdvertisementStatusGetResponseEnvelopeSuccessTrue PrefixAdvertisementStatusGetResponseEnvelopeSuccess = true
)

func (r PrefixAdvertisementStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixAdvertisementStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
