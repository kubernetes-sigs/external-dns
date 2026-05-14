// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DEXService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXService] method instead.
type DEXService struct {
	Options               []option.RequestOption
	WARPChangeEvents      *DEXWARPChangeEventService
	Commands              *DEXCommandService
	Colos                 *DEXColoService
	FleetStatus           *DEXFleetStatusService
	HTTPTests             *DEXHTTPTestService
	Tests                 *DEXTestService
	TracerouteTestResults *DEXTracerouteTestResultService
	TracerouteTests       *DEXTracerouteTestService
}

// NewDEXService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXService(opts ...option.RequestOption) (r *DEXService) {
	r = &DEXService{}
	r.Options = opts
	r.WARPChangeEvents = NewDEXWARPChangeEventService(opts...)
	r.Commands = NewDEXCommandService(opts...)
	r.Colos = NewDEXColoService(opts...)
	r.FleetStatus = NewDEXFleetStatusService(opts...)
	r.HTTPTests = NewDEXHTTPTestService(opts...)
	r.Tests = NewDEXTestService(opts...)
	r.TracerouteTestResults = NewDEXTracerouteTestResultService(opts...)
	r.TracerouteTests = NewDEXTracerouteTestService(opts...)
	return
}

type DigitalExperienceMonitor struct {
	ID string `json:"id,required"`
	// Whether the policy is the default for the account
	Default bool                         `json:"default,required"`
	Name    string                       `json:"name,required"`
	JSON    digitalExperienceMonitorJSON `json:"-"`
}

// digitalExperienceMonitorJSON contains the JSON metadata for the struct
// [DigitalExperienceMonitor]
type digitalExperienceMonitorJSON struct {
	ID          apijson.Field
	Default     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DigitalExperienceMonitor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r digitalExperienceMonitorJSON) RawJSON() string {
	return r.raw
}

type NetworkPath struct {
	Slots []NetworkPathSlot `json:"slots,required"`
	// Specifies the sampling applied, if any, to the slots response. When sampled,
	// results shown represent the first test run to the start of each sampling
	// interval.
	Sampling NetworkPathSampling `json:"sampling,nullable"`
	JSON     networkPathJSON     `json:"-"`
}

// networkPathJSON contains the JSON metadata for the struct [NetworkPath]
type networkPathJSON struct {
	Slots       apijson.Field
	Sampling    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkPath) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkPathJSON) RawJSON() string {
	return r.raw
}

type NetworkPathSlot struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// Round trip time in ms of the client to app mile
	ClientToAppRTTMs int64 `json:"clientToAppRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare egress mile
	ClientToCfEgressRTTMs int64 `json:"clientToCfEgressRttMs,required,nullable"`
	// Round trip time in ms of the client to Cloudflare ingress mile
	ClientToCfIngressRTTMs int64  `json:"clientToCfIngressRttMs,required,nullable"`
	Timestamp              string `json:"timestamp,required"`
	// Round trip time in ms of the client to ISP mile
	ClientToISPRTTMs int64               `json:"clientToIspRttMs,nullable"`
	JSON             networkPathSlotJSON `json:"-"`
}

// networkPathSlotJSON contains the JSON metadata for the struct [NetworkPathSlot]
type networkPathSlotJSON struct {
	ID                     apijson.Field
	ClientToAppRTTMs       apijson.Field
	ClientToCfEgressRTTMs  apijson.Field
	ClientToCfIngressRTTMs apijson.Field
	Timestamp              apijson.Field
	ClientToISPRTTMs       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *NetworkPathSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkPathSlotJSON) RawJSON() string {
	return r.raw
}

// Specifies the sampling applied, if any, to the slots response. When sampled,
// results shown represent the first test run to the start of each sampling
// interval.
type NetworkPathSampling struct {
	Unit  NetworkPathSamplingUnit `json:"unit,required"`
	Value int64                   `json:"value,required"`
	JSON  networkPathSamplingJSON `json:"-"`
}

// networkPathSamplingJSON contains the JSON metadata for the struct
// [NetworkPathSampling]
type networkPathSamplingJSON struct {
	Unit        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkPathSampling) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkPathSamplingJSON) RawJSON() string {
	return r.raw
}

type NetworkPathSamplingUnit string

const (
	NetworkPathSamplingUnitHours NetworkPathSamplingUnit = "hours"
)

func (r NetworkPathSamplingUnit) IsKnown() bool {
	switch r {
	case NetworkPathSamplingUnitHours:
		return true
	}
	return false
}

type NetworkPathResponse struct {
	// API Resource UUID tag.
	ID         string `json:"id,required"`
	DeviceName string `json:"deviceName"`
	// The interval at which the Traceroute synthetic application test is set to run.
	Interval    string                  `json:"interval"`
	Kind        NetworkPathResponseKind `json:"kind"`
	Name        string                  `json:"name"`
	NetworkPath NetworkPath             `json:"networkPath,nullable"`
	// The host of the Traceroute synthetic application test
	URL  string                  `json:"url"`
	JSON networkPathResponseJSON `json:"-"`
}

// networkPathResponseJSON contains the JSON metadata for the struct
// [NetworkPathResponse]
type networkPathResponseJSON struct {
	ID          apijson.Field
	DeviceName  apijson.Field
	Interval    apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	NetworkPath apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkPathResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkPathResponseJSON) RawJSON() string {
	return r.raw
}

type NetworkPathResponseKind string

const (
	NetworkPathResponseKindTraceroute NetworkPathResponseKind = "traceroute"
)

func (r NetworkPathResponseKind) IsKnown() bool {
	switch r {
	case NetworkPathResponseKindTraceroute:
		return true
	}
	return false
}

type Percentiles struct {
	// p50 observed in the time period
	P50 float64 `json:"p50,nullable"`
	// p90 observed in the time period
	P90 float64 `json:"p90,nullable"`
	// p95 observed in the time period
	P95 float64 `json:"p95,nullable"`
	// p99 observed in the time period
	P99  float64         `json:"p99,nullable"`
	JSON percentilesJSON `json:"-"`
}

// percentilesJSON contains the JSON metadata for the struct [Percentiles]
type percentilesJSON struct {
	P50         apijson.Field
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Percentiles) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r percentilesJSON) RawJSON() string {
	return r.raw
}
