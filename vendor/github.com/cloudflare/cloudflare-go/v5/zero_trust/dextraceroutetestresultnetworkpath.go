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
)

// DEXTracerouteTestResultNetworkPathService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXTracerouteTestResultNetworkPathService] method instead.
type DEXTracerouteTestResultNetworkPathService struct {
	Options []option.RequestOption
}

// NewDEXTracerouteTestResultNetworkPathService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDEXTracerouteTestResultNetworkPathService(opts ...option.RequestOption) (r *DEXTracerouteTestResultNetworkPathService) {
	r = &DEXTracerouteTestResultNetworkPathService{}
	r.Options = opts
	return
}

// Get a breakdown of hops and performance metrics for a specific traceroute test
// run
func (r *DEXTracerouteTestResultNetworkPathService) Get(ctx context.Context, testResultID string, query DEXTracerouteTestResultNetworkPathGetParams, opts ...option.RequestOption) (res *DEXTracerouteTestResultNetworkPathGetResponse, err error) {
	var env DEXTracerouteTestResultNetworkPathGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testResultID == "" {
		err = errors.New("missing required test_result_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/traceroute-test-results/%s/network-path", query.AccountID, testResultID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXTracerouteTestResultNetworkPathGetResponse struct {
	// an array of the hops taken by the device to reach the end destination
	Hops []DEXTracerouteTestResultNetworkPathGetResponseHop `json:"hops,required"`
	// API Resource UUID tag.
	ResultID string `json:"resultId,required"`
	// name of the device associated with this network path response
	DeviceName string `json:"deviceName"`
	// API Resource UUID tag.
	TestID string `json:"testId"`
	// name of the tracroute test
	TestName string                                            `json:"testName"`
	JSON     dexTracerouteTestResultNetworkPathGetResponseJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseJSON contains the JSON metadata for
// the struct [DEXTracerouteTestResultNetworkPathGetResponse]
type dexTracerouteTestResultNetworkPathGetResponseJSON struct {
	Hops        apijson.Field
	ResultID    apijson.Field
	DeviceName  apijson.Field
	TestID      apijson.Field
	TestName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathGetResponseHop struct {
	TTL           int64                                                     `json:"ttl,required"`
	ASN           int64                                                     `json:"asn,nullable"`
	Aso           string                                                    `json:"aso,nullable"`
	IPAddress     string                                                    `json:"ipAddress,nullable"`
	Location      DEXTracerouteTestResultNetworkPathGetResponseHopsLocation `json:"location,nullable"`
	Mile          DEXTracerouteTestResultNetworkPathGetResponseHopsMile     `json:"mile,nullable"`
	Name          string                                                    `json:"name,nullable"`
	PacketLossPct float64                                                   `json:"packetLossPct,nullable"`
	RTTMs         int64                                                     `json:"rttMs,nullable"`
	JSON          dexTracerouteTestResultNetworkPathGetResponseHopJSON      `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseHopJSON contains the JSON metadata
// for the struct [DEXTracerouteTestResultNetworkPathGetResponseHop]
type dexTracerouteTestResultNetworkPathGetResponseHopJSON struct {
	TTL           apijson.Field
	ASN           apijson.Field
	Aso           apijson.Field
	IPAddress     apijson.Field
	Location      apijson.Field
	Mile          apijson.Field
	Name          apijson.Field
	PacketLossPct apijson.Field
	RTTMs         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseHop) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseHopJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathGetResponseHopsLocation struct {
	City  string                                                        `json:"city,nullable"`
	State string                                                        `json:"state,nullable"`
	Zip   string                                                        `json:"zip,nullable"`
	JSON  dexTracerouteTestResultNetworkPathGetResponseHopsLocationJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseHopsLocationJSON contains the JSON
// metadata for the struct
// [DEXTracerouteTestResultNetworkPathGetResponseHopsLocation]
type dexTracerouteTestResultNetworkPathGetResponseHopsLocationJSON struct {
	City        apijson.Field
	State       apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseHopsLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseHopsLocationJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathGetResponseHopsMile string

const (
	DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToApp       DEXTracerouteTestResultNetworkPathGetResponseHopsMile = "client-to-app"
	DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToCfEgress  DEXTracerouteTestResultNetworkPathGetResponseHopsMile = "client-to-cf-egress"
	DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToCfIngress DEXTracerouteTestResultNetworkPathGetResponseHopsMile = "client-to-cf-ingress"
	DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToISP       DEXTracerouteTestResultNetworkPathGetResponseHopsMile = "client-to-isp"
)

func (r DEXTracerouteTestResultNetworkPathGetResponseHopsMile) IsKnown() bool {
	switch r {
	case DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToApp, DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToCfEgress, DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToCfIngress, DEXTracerouteTestResultNetworkPathGetResponseHopsMileClientToISP:
		return true
	}
	return false
}

type DEXTracerouteTestResultNetworkPathGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DEXTracerouteTestResultNetworkPathGetResponseEnvelope struct {
	Errors   []DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DEXTracerouteTestResultNetworkPathGetResponse                `json:"result"`
	JSON    dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON    `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [DEXTracerouteTestResultNetworkPathGetResponseEnvelope]
type dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrors struct {
	Code             int64                                                             `json:"code,required"`
	Message          string                                                            `json:"message,required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrors]
type dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSource]
type dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessages struct {
	Code             int64                                                               `json:"code,required"`
	Message          string                                                              `json:"message,required"`
	DocumentationURL string                                                              `json:"documentation_url"`
	Source           DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessages]
type dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                                  `json:"pointer"`
	JSON    dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSource]
type dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexTracerouteTestResultNetworkPathGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess bool

const (
	DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccessTrue DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess = true
)

func (r DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXTracerouteTestResultNetworkPathGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
