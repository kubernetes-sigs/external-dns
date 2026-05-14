// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dcv_delegation

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DCVDelegationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDCVDelegationService] method instead.
type DCVDelegationService struct {
	Options []option.RequestOption
}

// NewDCVDelegationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDCVDelegationService(opts ...option.RequestOption) (r *DCVDelegationService) {
	r = &DCVDelegationService{}
	r.Options = opts
	return
}

// Retrieve the account and zone specific unique identifier used as part of the
// CNAME target for DCV Delegation.
func (r *DCVDelegationService) Get(ctx context.Context, query DCVDelegationGetParams, opts ...option.RequestOption) (res *DCVDelegationUUID, err error) {
	var env DCVDelegationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dcv_delegation/uuid", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DCVDelegationUUID struct {
	// The DCV Delegation unique identifier.
	UUID string                `json:"uuid"`
	JSON dcvDelegationUUIDJSON `json:"-"`
}

// dcvDelegationUUIDJSON contains the JSON metadata for the struct
// [DCVDelegationUUID]
type dcvDelegationUUIDJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationUUID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dcvDelegationUUIDJSON) RawJSON() string {
	return r.raw
}

type DCVDelegationGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DCVDelegationGetResponseEnvelope struct {
	Errors   []DCVDelegationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DCVDelegationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DCVDelegationGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DCVDelegationUUID                       `json:"result"`
	JSON    dcvDelegationGetResponseEnvelopeJSON    `json:"-"`
}

// dcvDelegationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DCVDelegationGetResponseEnvelope]
type dcvDelegationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dcvDelegationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DCVDelegationGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DCVDelegationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dcvDelegationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dcvDelegationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DCVDelegationGetResponseEnvelopeErrors]
type dcvDelegationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DCVDelegationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dcvDelegationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DCVDelegationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dcvDelegationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dcvDelegationGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DCVDelegationGetResponseEnvelopeErrorsSource]
type dcvDelegationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dcvDelegationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DCVDelegationGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DCVDelegationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dcvDelegationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dcvDelegationGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DCVDelegationGetResponseEnvelopeMessages]
type dcvDelegationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DCVDelegationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dcvDelegationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DCVDelegationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dcvDelegationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dcvDelegationGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DCVDelegationGetResponseEnvelopeMessagesSource]
type dcvDelegationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DCVDelegationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dcvDelegationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DCVDelegationGetResponseEnvelopeSuccess bool

const (
	DCVDelegationGetResponseEnvelopeSuccessTrue DCVDelegationGetResponseEnvelopeSuccess = true
)

func (r DCVDelegationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DCVDelegationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
