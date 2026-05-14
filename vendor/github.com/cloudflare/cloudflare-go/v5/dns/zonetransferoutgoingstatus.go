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

// ZoneTransferOutgoingStatusService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferOutgoingStatusService] method instead.
type ZoneTransferOutgoingStatusService struct {
	Options []option.RequestOption
}

// NewZoneTransferOutgoingStatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneTransferOutgoingStatusService(opts ...option.RequestOption) (r *ZoneTransferOutgoingStatusService) {
	r = &ZoneTransferOutgoingStatusService{}
	r.Options = opts
	return
}

// Get primary zone transfer status.
func (r *ZoneTransferOutgoingStatusService) Get(ctx context.Context, query ZoneTransferOutgoingStatusGetParams, opts ...option.RequestOption) (res *EnableTransfer, err error) {
	var env ZoneTransferOutgoingStatusGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/status", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneTransferOutgoingStatusGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferOutgoingStatusGetResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingStatusGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingStatusGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess `json:"success,required"`
	// The zone transfer status of a primary zone.
	Result EnableTransfer                                    `json:"result"`
	JSON   zoneTransferOutgoingStatusGetResponseEnvelopeJSON `json:"-"`
}

// zoneTransferOutgoingStatusGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingStatusGetResponseEnvelope]
type zoneTransferOutgoingStatusGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingStatusGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingStatusGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingStatusGetResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           ZoneTransferOutgoingStatusGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingStatusGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingStatusGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingStatusGetResponseEnvelopeErrors]
type zoneTransferOutgoingStatusGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingStatusGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingStatusGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingStatusGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    zoneTransferOutgoingStatusGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingStatusGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferOutgoingStatusGetResponseEnvelopeErrorsSource]
type zoneTransferOutgoingStatusGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingStatusGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingStatusGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingStatusGetResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           ZoneTransferOutgoingStatusGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingStatusGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingStatusGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingStatusGetResponseEnvelopeMessages]
type zoneTransferOutgoingStatusGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingStatusGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingStatusGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingStatusGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    zoneTransferOutgoingStatusGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingStatusGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ZoneTransferOutgoingStatusGetResponseEnvelopeMessagesSource]
type zoneTransferOutgoingStatusGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingStatusGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingStatusGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingStatusGetResponseEnvelopeSuccessTrue ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingStatusGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingStatusGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
