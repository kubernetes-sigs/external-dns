// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// ConnectorEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorEventService] method instead.
type ConnectorEventService struct {
	Options []option.RequestOption
	Latest  *ConnectorEventLatestService
}

// NewConnectorEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorEventService(opts ...option.RequestOption) (r *ConnectorEventService) {
	r = &ConnectorEventService{}
	r.Options = opts
	r.Latest = NewConnectorEventLatestService(opts...)
	return
}

// List Events
func (r *ConnectorEventService) List(ctx context.Context, connectorID string, params ConnectorEventListParams, opts ...option.RequestOption) (res *ConnectorEventListResponse, err error) {
	var env ConnectorEventListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Event
func (r *ConnectorEventService) Get(ctx context.Context, connectorID string, eventT float64, eventN float64, query ConnectorEventGetParams, opts ...option.RequestOption) (res *ConnectorEventGetResponse, err error) {
	var env ConnectorEventGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events/%v.%v", query.AccountID, connectorID, eventT, eventN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectorEventListResponse struct {
	Count  float64                          `json:"count,required"`
	Items  []ConnectorEventListResponseItem `json:"items,required"`
	Cursor string                           `json:"cursor"`
	JSON   connectorEventListResponseJSON   `json:"-"`
}

// connectorEventListResponseJSON contains the JSON metadata for the struct
// [ConnectorEventListResponse]
type connectorEventListResponseJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseItem struct {
	// Time the Event was collected (seconds since the Unix epoch)
	A float64 `json:"a,required"`
	// Kind
	K string `json:"k,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                            `json:"t,required"`
	JSON connectorEventListResponseItemJSON `json:"-"`
}

// connectorEventListResponseItemJSON contains the JSON metadata for the struct
// [ConnectorEventListResponseItem]
type connectorEventListResponseItemJSON struct {
	A           apijson.Field
	K           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Recorded Event
type ConnectorEventGetResponse struct {
	E ConnectorEventGetResponseE `json:"e,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                       `json:"t,required"`
	JSON connectorEventGetResponseJSON `json:"-"`
}

// connectorEventGetResponseJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponse]
type connectorEventGetResponseJSON struct {
	E           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseE struct {
	// Initialized process
	K ConnectorEventGetResponseEK `json:"k,required"`
	// Location of upgrade bundle
	URL   string                         `json:"url"`
	JSON  connectorEventGetResponseEJSON `json:"-"`
	union ConnectorEventGetResponseEUnion
}

// connectorEventGetResponseEJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseE]
type connectorEventGetResponseEJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r connectorEventGetResponseEJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectorEventGetResponseE) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectorEventGetResponseE{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectorEventGetResponseEUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [ConnectorEventGetResponseEInit],
// [ConnectorEventGetResponseELeave], [ConnectorEventGetResponseEStartAttestation],
// [ConnectorEventGetResponseEFinishAttestationSuccess],
// [ConnectorEventGetResponseEFinishAttestationFailure],
// [ConnectorEventGetResponseEStartRotateCryptKey],
// [ConnectorEventGetResponseEFinishRotateCryptKeySuccess],
// [ConnectorEventGetResponseEFinishRotateCryptKeyFailure],
// [ConnectorEventGetResponseEStartRotatePki],
// [ConnectorEventGetResponseEFinishRotatePkiSuccess],
// [ConnectorEventGetResponseEFinishRotatePkiFailure],
// [ConnectorEventGetResponseEStartUpgrade],
// [ConnectorEventGetResponseEFinishUpgradeSuccess],
// [ConnectorEventGetResponseEFinishUpgradeFailure],
// [ConnectorEventGetResponseEReconcile],
// [ConnectorEventGetResponseEConfigureCloudflaredTunnel].
func (r ConnectorEventGetResponseE) AsUnion() ConnectorEventGetResponseEUnion {
	return r.union
}

// Union satisfied by [ConnectorEventGetResponseEInit],
// [ConnectorEventGetResponseELeave], [ConnectorEventGetResponseEStartAttestation],
// [ConnectorEventGetResponseEFinishAttestationSuccess],
// [ConnectorEventGetResponseEFinishAttestationFailure],
// [ConnectorEventGetResponseEStartRotateCryptKey],
// [ConnectorEventGetResponseEFinishRotateCryptKeySuccess],
// [ConnectorEventGetResponseEFinishRotateCryptKeyFailure],
// [ConnectorEventGetResponseEStartRotatePki],
// [ConnectorEventGetResponseEFinishRotatePkiSuccess],
// [ConnectorEventGetResponseEFinishRotatePkiFailure],
// [ConnectorEventGetResponseEStartUpgrade],
// [ConnectorEventGetResponseEFinishUpgradeSuccess],
// [ConnectorEventGetResponseEFinishUpgradeFailure],
// [ConnectorEventGetResponseEReconcile] or
// [ConnectorEventGetResponseEConfigureCloudflaredTunnel].
type ConnectorEventGetResponseEUnion interface {
	implementsConnectorEventGetResponseE()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectorEventGetResponseEUnion)(nil)).Elem(),
		"k",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEInit{}),
			DiscriminatorValue: "Init",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseELeave{}),
			DiscriminatorValue: "Leave",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEStartAttestation{}),
			DiscriminatorValue: "StartAttestation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishAttestationSuccess{}),
			DiscriminatorValue: "FinishAttestationSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishAttestationFailure{}),
			DiscriminatorValue: "FinishAttestationFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEStartRotateCryptKey{}),
			DiscriminatorValue: "StartRotateCryptKey",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishRotateCryptKeySuccess{}),
			DiscriminatorValue: "FinishRotateCryptKeySuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishRotateCryptKeyFailure{}),
			DiscriminatorValue: "FinishRotateCryptKeyFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEStartRotatePki{}),
			DiscriminatorValue: "StartRotatePki",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishRotatePkiSuccess{}),
			DiscriminatorValue: "FinishRotatePkiSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishRotatePkiFailure{}),
			DiscriminatorValue: "FinishRotatePkiFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEStartUpgrade{}),
			DiscriminatorValue: "StartUpgrade",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishUpgradeSuccess{}),
			DiscriminatorValue: "FinishUpgradeSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEFinishUpgradeFailure{}),
			DiscriminatorValue: "FinishUpgradeFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEReconcile{}),
			DiscriminatorValue: "Reconcile",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventGetResponseEConfigureCloudflaredTunnel{}),
			DiscriminatorValue: "ConfigureCloudflaredTunnel",
		},
	)
}

type ConnectorEventGetResponseEInit struct {
	// Initialized process
	K    ConnectorEventGetResponseEInitK    `json:"k,required"`
	JSON connectorEventGetResponseEInitJSON `json:"-"`
}

// connectorEventGetResponseEInitJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseEInit]
type connectorEventGetResponseEInitJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEInit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEInitJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEInit) implementsConnectorEventGetResponseE() {}

// Initialized process
type ConnectorEventGetResponseEInitK string

const (
	ConnectorEventGetResponseEInitKInit ConnectorEventGetResponseEInitK = "Init"
)

func (r ConnectorEventGetResponseEInitK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEInitKInit:
		return true
	}
	return false
}

type ConnectorEventGetResponseELeave struct {
	// Stopped process
	K    ConnectorEventGetResponseELeaveK    `json:"k,required"`
	JSON connectorEventGetResponseELeaveJSON `json:"-"`
}

// connectorEventGetResponseELeaveJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseELeave]
type connectorEventGetResponseELeaveJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseELeave) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseELeaveJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseELeave) implementsConnectorEventGetResponseE() {}

// Stopped process
type ConnectorEventGetResponseELeaveK string

const (
	ConnectorEventGetResponseELeaveKLeave ConnectorEventGetResponseELeaveK = "Leave"
)

func (r ConnectorEventGetResponseELeaveK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseELeaveKLeave:
		return true
	}
	return false
}

type ConnectorEventGetResponseEStartAttestation struct {
	// Started attestation
	K    ConnectorEventGetResponseEStartAttestationK    `json:"k,required"`
	JSON connectorEventGetResponseEStartAttestationJSON `json:"-"`
}

// connectorEventGetResponseEStartAttestationJSON contains the JSON metadata for
// the struct [ConnectorEventGetResponseEStartAttestation]
type connectorEventGetResponseEStartAttestationJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEStartAttestation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEStartAttestationJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEStartAttestation) implementsConnectorEventGetResponseE() {}

// Started attestation
type ConnectorEventGetResponseEStartAttestationK string

const (
	ConnectorEventGetResponseEStartAttestationKStartAttestation ConnectorEventGetResponseEStartAttestationK = "StartAttestation"
)

func (r ConnectorEventGetResponseEStartAttestationK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEStartAttestationKStartAttestation:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishAttestationSuccess struct {
	// Finished attestation
	K    ConnectorEventGetResponseEFinishAttestationSuccessK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishAttestationSuccessJSON `json:"-"`
}

// connectorEventGetResponseEFinishAttestationSuccessJSON contains the JSON
// metadata for the struct [ConnectorEventGetResponseEFinishAttestationSuccess]
type connectorEventGetResponseEFinishAttestationSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishAttestationSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishAttestationSuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishAttestationSuccess) implementsConnectorEventGetResponseE() {}

// Finished attestation
type ConnectorEventGetResponseEFinishAttestationSuccessK string

const (
	ConnectorEventGetResponseEFinishAttestationSuccessKFinishAttestationSuccess ConnectorEventGetResponseEFinishAttestationSuccessK = "FinishAttestationSuccess"
)

func (r ConnectorEventGetResponseEFinishAttestationSuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishAttestationSuccessKFinishAttestationSuccess:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishAttestationFailure struct {
	// Failed attestation
	K    ConnectorEventGetResponseEFinishAttestationFailureK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishAttestationFailureJSON `json:"-"`
}

// connectorEventGetResponseEFinishAttestationFailureJSON contains the JSON
// metadata for the struct [ConnectorEventGetResponseEFinishAttestationFailure]
type connectorEventGetResponseEFinishAttestationFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishAttestationFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishAttestationFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishAttestationFailure) implementsConnectorEventGetResponseE() {}

// Failed attestation
type ConnectorEventGetResponseEFinishAttestationFailureK string

const (
	ConnectorEventGetResponseEFinishAttestationFailureKFinishAttestationFailure ConnectorEventGetResponseEFinishAttestationFailureK = "FinishAttestationFailure"
)

func (r ConnectorEventGetResponseEFinishAttestationFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishAttestationFailureKFinishAttestationFailure:
		return true
	}
	return false
}

type ConnectorEventGetResponseEStartRotateCryptKey struct {
	// Started crypt key rotation
	K    ConnectorEventGetResponseEStartRotateCryptKeyK    `json:"k,required"`
	JSON connectorEventGetResponseEStartRotateCryptKeyJSON `json:"-"`
}

// connectorEventGetResponseEStartRotateCryptKeyJSON contains the JSON metadata for
// the struct [ConnectorEventGetResponseEStartRotateCryptKey]
type connectorEventGetResponseEStartRotateCryptKeyJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEStartRotateCryptKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEStartRotateCryptKeyJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEStartRotateCryptKey) implementsConnectorEventGetResponseE() {}

// Started crypt key rotation
type ConnectorEventGetResponseEStartRotateCryptKeyK string

const (
	ConnectorEventGetResponseEStartRotateCryptKeyKStartRotateCryptKey ConnectorEventGetResponseEStartRotateCryptKeyK = "StartRotateCryptKey"
)

func (r ConnectorEventGetResponseEStartRotateCryptKeyK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEStartRotateCryptKeyKStartRotateCryptKey:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishRotateCryptKeySuccess struct {
	// Finished crypt key rotation
	K    ConnectorEventGetResponseEFinishRotateCryptKeySuccessK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishRotateCryptKeySuccessJSON `json:"-"`
}

// connectorEventGetResponseEFinishRotateCryptKeySuccessJSON contains the JSON
// metadata for the struct [ConnectorEventGetResponseEFinishRotateCryptKeySuccess]
type connectorEventGetResponseEFinishRotateCryptKeySuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishRotateCryptKeySuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishRotateCryptKeySuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishRotateCryptKeySuccess) implementsConnectorEventGetResponseE() {
}

// Finished crypt key rotation
type ConnectorEventGetResponseEFinishRotateCryptKeySuccessK string

const (
	ConnectorEventGetResponseEFinishRotateCryptKeySuccessKFinishRotateCryptKeySuccess ConnectorEventGetResponseEFinishRotateCryptKeySuccessK = "FinishRotateCryptKeySuccess"
)

func (r ConnectorEventGetResponseEFinishRotateCryptKeySuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishRotateCryptKeySuccessKFinishRotateCryptKeySuccess:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishRotateCryptKeyFailure struct {
	// Failed crypt key rotation
	K    ConnectorEventGetResponseEFinishRotateCryptKeyFailureK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishRotateCryptKeyFailureJSON `json:"-"`
}

// connectorEventGetResponseEFinishRotateCryptKeyFailureJSON contains the JSON
// metadata for the struct [ConnectorEventGetResponseEFinishRotateCryptKeyFailure]
type connectorEventGetResponseEFinishRotateCryptKeyFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishRotateCryptKeyFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishRotateCryptKeyFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishRotateCryptKeyFailure) implementsConnectorEventGetResponseE() {
}

// Failed crypt key rotation
type ConnectorEventGetResponseEFinishRotateCryptKeyFailureK string

const (
	ConnectorEventGetResponseEFinishRotateCryptKeyFailureKFinishRotateCryptKeyFailure ConnectorEventGetResponseEFinishRotateCryptKeyFailureK = "FinishRotateCryptKeyFailure"
)

func (r ConnectorEventGetResponseEFinishRotateCryptKeyFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishRotateCryptKeyFailureKFinishRotateCryptKeyFailure:
		return true
	}
	return false
}

type ConnectorEventGetResponseEStartRotatePki struct {
	// Started PKI rotation
	K    ConnectorEventGetResponseEStartRotatePkiK    `json:"k,required"`
	JSON connectorEventGetResponseEStartRotatePkiJSON `json:"-"`
}

// connectorEventGetResponseEStartRotatePkiJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEStartRotatePki]
type connectorEventGetResponseEStartRotatePkiJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEStartRotatePki) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEStartRotatePkiJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEStartRotatePki) implementsConnectorEventGetResponseE() {}

// Started PKI rotation
type ConnectorEventGetResponseEStartRotatePkiK string

const (
	ConnectorEventGetResponseEStartRotatePkiKStartRotatePki ConnectorEventGetResponseEStartRotatePkiK = "StartRotatePki"
)

func (r ConnectorEventGetResponseEStartRotatePkiK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEStartRotatePkiKStartRotatePki:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishRotatePkiSuccess struct {
	// Finished PKI rotation
	K    ConnectorEventGetResponseEFinishRotatePkiSuccessK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishRotatePkiSuccessJSON `json:"-"`
}

// connectorEventGetResponseEFinishRotatePkiSuccessJSON contains the JSON metadata
// for the struct [ConnectorEventGetResponseEFinishRotatePkiSuccess]
type connectorEventGetResponseEFinishRotatePkiSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishRotatePkiSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishRotatePkiSuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishRotatePkiSuccess) implementsConnectorEventGetResponseE() {}

// Finished PKI rotation
type ConnectorEventGetResponseEFinishRotatePkiSuccessK string

const (
	ConnectorEventGetResponseEFinishRotatePkiSuccessKFinishRotatePkiSuccess ConnectorEventGetResponseEFinishRotatePkiSuccessK = "FinishRotatePkiSuccess"
)

func (r ConnectorEventGetResponseEFinishRotatePkiSuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishRotatePkiSuccessKFinishRotatePkiSuccess:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishRotatePkiFailure struct {
	// Failed PKI rotation
	K    ConnectorEventGetResponseEFinishRotatePkiFailureK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishRotatePkiFailureJSON `json:"-"`
}

// connectorEventGetResponseEFinishRotatePkiFailureJSON contains the JSON metadata
// for the struct [ConnectorEventGetResponseEFinishRotatePkiFailure]
type connectorEventGetResponseEFinishRotatePkiFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishRotatePkiFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishRotatePkiFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishRotatePkiFailure) implementsConnectorEventGetResponseE() {}

// Failed PKI rotation
type ConnectorEventGetResponseEFinishRotatePkiFailureK string

const (
	ConnectorEventGetResponseEFinishRotatePkiFailureKFinishRotatePkiFailure ConnectorEventGetResponseEFinishRotatePkiFailureK = "FinishRotatePkiFailure"
)

func (r ConnectorEventGetResponseEFinishRotatePkiFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishRotatePkiFailureKFinishRotatePkiFailure:
		return true
	}
	return false
}

type ConnectorEventGetResponseEStartUpgrade struct {
	// Started upgrade
	K ConnectorEventGetResponseEStartUpgradeK `json:"k,required"`
	// Location of upgrade bundle
	URL  string                                     `json:"url,required"`
	JSON connectorEventGetResponseEStartUpgradeJSON `json:"-"`
}

// connectorEventGetResponseEStartUpgradeJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEStartUpgrade]
type connectorEventGetResponseEStartUpgradeJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEStartUpgrade) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEStartUpgradeJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEStartUpgrade) implementsConnectorEventGetResponseE() {}

// Started upgrade
type ConnectorEventGetResponseEStartUpgradeK string

const (
	ConnectorEventGetResponseEStartUpgradeKStartUpgrade ConnectorEventGetResponseEStartUpgradeK = "StartUpgrade"
)

func (r ConnectorEventGetResponseEStartUpgradeK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEStartUpgradeKStartUpgrade:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishUpgradeSuccess struct {
	// Finished upgrade
	K    ConnectorEventGetResponseEFinishUpgradeSuccessK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishUpgradeSuccessJSON `json:"-"`
}

// connectorEventGetResponseEFinishUpgradeSuccessJSON contains the JSON metadata
// for the struct [ConnectorEventGetResponseEFinishUpgradeSuccess]
type connectorEventGetResponseEFinishUpgradeSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishUpgradeSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishUpgradeSuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishUpgradeSuccess) implementsConnectorEventGetResponseE() {}

// Finished upgrade
type ConnectorEventGetResponseEFinishUpgradeSuccessK string

const (
	ConnectorEventGetResponseEFinishUpgradeSuccessKFinishUpgradeSuccess ConnectorEventGetResponseEFinishUpgradeSuccessK = "FinishUpgradeSuccess"
)

func (r ConnectorEventGetResponseEFinishUpgradeSuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishUpgradeSuccessKFinishUpgradeSuccess:
		return true
	}
	return false
}

type ConnectorEventGetResponseEFinishUpgradeFailure struct {
	// Failed upgrade
	K    ConnectorEventGetResponseEFinishUpgradeFailureK    `json:"k,required"`
	JSON connectorEventGetResponseEFinishUpgradeFailureJSON `json:"-"`
}

// connectorEventGetResponseEFinishUpgradeFailureJSON contains the JSON metadata
// for the struct [ConnectorEventGetResponseEFinishUpgradeFailure]
type connectorEventGetResponseEFinishUpgradeFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEFinishUpgradeFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEFinishUpgradeFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEFinishUpgradeFailure) implementsConnectorEventGetResponseE() {}

// Failed upgrade
type ConnectorEventGetResponseEFinishUpgradeFailureK string

const (
	ConnectorEventGetResponseEFinishUpgradeFailureKFinishUpgradeFailure ConnectorEventGetResponseEFinishUpgradeFailureK = "FinishUpgradeFailure"
)

func (r ConnectorEventGetResponseEFinishUpgradeFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEFinishUpgradeFailureKFinishUpgradeFailure:
		return true
	}
	return false
}

type ConnectorEventGetResponseEReconcile struct {
	// Reconciled
	K    ConnectorEventGetResponseEReconcileK    `json:"k,required"`
	JSON connectorEventGetResponseEReconcileJSON `json:"-"`
}

// connectorEventGetResponseEReconcileJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEReconcile]
type connectorEventGetResponseEReconcileJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEReconcile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEReconcileJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEReconcile) implementsConnectorEventGetResponseE() {}

// Reconciled
type ConnectorEventGetResponseEReconcileK string

const (
	ConnectorEventGetResponseEReconcileKReconcile ConnectorEventGetResponseEReconcileK = "Reconcile"
)

func (r ConnectorEventGetResponseEReconcileK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEReconcileKReconcile:
		return true
	}
	return false
}

type ConnectorEventGetResponseEConfigureCloudflaredTunnel struct {
	// Configured Cloudflared tunnel
	K    ConnectorEventGetResponseEConfigureCloudflaredTunnelK    `json:"k,required"`
	JSON connectorEventGetResponseEConfigureCloudflaredTunnelJSON `json:"-"`
}

// connectorEventGetResponseEConfigureCloudflaredTunnelJSON contains the JSON
// metadata for the struct [ConnectorEventGetResponseEConfigureCloudflaredTunnel]
type connectorEventGetResponseEConfigureCloudflaredTunnelJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEConfigureCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEConfigureCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventGetResponseEConfigureCloudflaredTunnel) implementsConnectorEventGetResponseE() {
}

// Configured Cloudflared tunnel
type ConnectorEventGetResponseEConfigureCloudflaredTunnelK string

const (
	ConnectorEventGetResponseEConfigureCloudflaredTunnelKConfigureCloudflaredTunnel ConnectorEventGetResponseEConfigureCloudflaredTunnelK = "ConfigureCloudflaredTunnel"
)

func (r ConnectorEventGetResponseEConfigureCloudflaredTunnelK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEConfigureCloudflaredTunnelKConfigureCloudflaredTunnel:
		return true
	}
	return false
}

// Initialized process
type ConnectorEventGetResponseEK string

const (
	ConnectorEventGetResponseEKInit                        ConnectorEventGetResponseEK = "Init"
	ConnectorEventGetResponseEKLeave                       ConnectorEventGetResponseEK = "Leave"
	ConnectorEventGetResponseEKStartAttestation            ConnectorEventGetResponseEK = "StartAttestation"
	ConnectorEventGetResponseEKFinishAttestationSuccess    ConnectorEventGetResponseEK = "FinishAttestationSuccess"
	ConnectorEventGetResponseEKFinishAttestationFailure    ConnectorEventGetResponseEK = "FinishAttestationFailure"
	ConnectorEventGetResponseEKStartRotateCryptKey         ConnectorEventGetResponseEK = "StartRotateCryptKey"
	ConnectorEventGetResponseEKFinishRotateCryptKeySuccess ConnectorEventGetResponseEK = "FinishRotateCryptKeySuccess"
	ConnectorEventGetResponseEKFinishRotateCryptKeyFailure ConnectorEventGetResponseEK = "FinishRotateCryptKeyFailure"
	ConnectorEventGetResponseEKStartRotatePki              ConnectorEventGetResponseEK = "StartRotatePki"
	ConnectorEventGetResponseEKFinishRotatePkiSuccess      ConnectorEventGetResponseEK = "FinishRotatePkiSuccess"
	ConnectorEventGetResponseEKFinishRotatePkiFailure      ConnectorEventGetResponseEK = "FinishRotatePkiFailure"
	ConnectorEventGetResponseEKStartUpgrade                ConnectorEventGetResponseEK = "StartUpgrade"
	ConnectorEventGetResponseEKFinishUpgradeSuccess        ConnectorEventGetResponseEK = "FinishUpgradeSuccess"
	ConnectorEventGetResponseEKFinishUpgradeFailure        ConnectorEventGetResponseEK = "FinishUpgradeFailure"
	ConnectorEventGetResponseEKReconcile                   ConnectorEventGetResponseEK = "Reconcile"
	ConnectorEventGetResponseEKConfigureCloudflaredTunnel  ConnectorEventGetResponseEK = "ConfigureCloudflaredTunnel"
)

func (r ConnectorEventGetResponseEK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEKInit, ConnectorEventGetResponseEKLeave, ConnectorEventGetResponseEKStartAttestation, ConnectorEventGetResponseEKFinishAttestationSuccess, ConnectorEventGetResponseEKFinishAttestationFailure, ConnectorEventGetResponseEKStartRotateCryptKey, ConnectorEventGetResponseEKFinishRotateCryptKeySuccess, ConnectorEventGetResponseEKFinishRotateCryptKeyFailure, ConnectorEventGetResponseEKStartRotatePki, ConnectorEventGetResponseEKFinishRotatePkiSuccess, ConnectorEventGetResponseEKFinishRotatePkiFailure, ConnectorEventGetResponseEKStartUpgrade, ConnectorEventGetResponseEKFinishUpgradeSuccess, ConnectorEventGetResponseEKFinishUpgradeFailure, ConnectorEventGetResponseEKReconcile, ConnectorEventGetResponseEKConfigureCloudflaredTunnel:
		return true
	}
	return false
}

type ConnectorEventListParams struct {
	// Account identifier
	AccountID param.Field[string]  `path:"account_id,required"`
	From      param.Field[float64] `query:"from,required"`
	To        param.Field[float64] `query:"to,required"`
	Cursor    param.Field[string]  `query:"cursor"`
	// Filter by event kind
	K     param.Field[string]  `query:"k"`
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [ConnectorEventListParams]'s query parameters as
// `url.Values`.
func (r ConnectorEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConnectorEventListResponseEnvelope struct {
	Result   ConnectorEventListResponse                   `json:"result,required"`
	Success  bool                                         `json:"success,required"`
	Errors   []ConnectorEventListResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventListResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventListResponseEnvelopeJSON       `json:"-"`
}

// connectorEventListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEventListResponseEnvelope]
type connectorEventListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseEnvelopeErrors struct {
	Code    float64                                      `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    connectorEventListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEventListResponseEnvelopeErrors]
type connectorEventListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseEnvelopeMessages struct {
	Code    float64                                        `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    connectorEventListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ConnectorEventListResponseEnvelopeMessages]
type connectorEventListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectorEventGetResponseEnvelope struct {
	// Recorded Event
	Result   ConnectorEventGetResponse                   `json:"result,required"`
	Success  bool                                        `json:"success,required"`
	Errors   []ConnectorEventGetResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventGetResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventGetResponseEnvelopeJSON       `json:"-"`
}

// connectorEventGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseEnvelope]
type connectorEventGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseEnvelopeErrors struct {
	Code    float64                                     `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    connectorEventGetResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEnvelopeErrors]
type connectorEventGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseEnvelopeMessages struct {
	Code    float64                                       `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    connectorEventGetResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEnvelopeMessages]
type connectorEventGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
