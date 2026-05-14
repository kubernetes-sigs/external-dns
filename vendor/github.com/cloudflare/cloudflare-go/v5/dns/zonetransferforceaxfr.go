// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// ZoneTransferForceAXFRService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferForceAXFRService] method instead.
type ZoneTransferForceAXFRService struct {
	Options []option.RequestOption
}

// NewZoneTransferForceAXFRService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneTransferForceAXFRService(opts ...option.RequestOption) (r *ZoneTransferForceAXFRService) {
	r = &ZoneTransferForceAXFRService{}
	r.Options = opts
	return
}

// Sends AXFR zone transfer request to primary nameserver(s).
func (r *ZoneTransferForceAXFRService) New(ctx context.Context, params ZoneTransferForceAXFRNewParams, opts ...option.RequestOption) (res *ForceAXFR, err error) {
	var env ZoneTransferForceAXFRNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/force_axfr", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ForceAXFR = string

type ZoneTransferForceAXFRNewParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r ZoneTransferForceAXFRNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneTransferForceAXFRNewResponseEnvelope struct {
	Errors   []ZoneTransferForceAXFRNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferForceAXFRNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferForceAXFRNewResponseEnvelopeSuccess `json:"success,required"`
	// When force_axfr query parameter is set to true, the response is a simple string.
	Result ForceAXFR                                    `json:"result"`
	JSON   zoneTransferForceAXFRNewResponseEnvelopeJSON `json:"-"`
}

// zoneTransferForceAXFRNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferForceAXFRNewResponseEnvelope]
type zoneTransferForceAXFRNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferForceAXFRNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferForceAXFRNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferForceAXFRNewResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           ZoneTransferForceAXFRNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferForceAXFRNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferForceAXFRNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneTransferForceAXFRNewResponseEnvelopeErrors]
type zoneTransferForceAXFRNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferForceAXFRNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferForceAXFRNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferForceAXFRNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    zoneTransferForceAXFRNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferForceAXFRNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferForceAXFRNewResponseEnvelopeErrorsSource]
type zoneTransferForceAXFRNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferForceAXFRNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferForceAXFRNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferForceAXFRNewResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ZoneTransferForceAXFRNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferForceAXFRNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferForceAXFRNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneTransferForceAXFRNewResponseEnvelopeMessages]
type zoneTransferForceAXFRNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferForceAXFRNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferForceAXFRNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferForceAXFRNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    zoneTransferForceAXFRNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferForceAXFRNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneTransferForceAXFRNewResponseEnvelopeMessagesSource]
type zoneTransferForceAXFRNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferForceAXFRNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferForceAXFRNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferForceAXFRNewResponseEnvelopeSuccess bool

const (
	ZoneTransferForceAXFRNewResponseEnvelopeSuccessTrue ZoneTransferForceAXFRNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferForceAXFRNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferForceAXFRNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
