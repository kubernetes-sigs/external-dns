// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// ConnectorEventLatestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorEventLatestService] method instead.
type ConnectorEventLatestService struct {
	Options []option.RequestOption
}

// NewConnectorEventLatestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorEventLatestService(opts ...option.RequestOption) (r *ConnectorEventLatestService) {
	r = &ConnectorEventLatestService{}
	r.Options = opts
	return
}

// Get latest Events
func (r *ConnectorEventLatestService) List(ctx context.Context, connectorID string, query ConnectorEventLatestListParams, opts ...option.RequestOption) (res *ConnectorEventLatestListResponse, err error) {
	var env ConnectorEventLatestListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events/latest", query.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectorEventLatestListResponse struct {
	Count float64                                `json:"count,required"`
	Items []ConnectorEventLatestListResponseItem `json:"items,required"`
	JSON  connectorEventLatestListResponseJSON   `json:"-"`
}

// connectorEventLatestListResponseJSON contains the JSON metadata for the struct
// [ConnectorEventLatestListResponse]
type connectorEventLatestListResponseJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseJSON) RawJSON() string {
	return r.raw
}

// Recorded Event
type ConnectorEventLatestListResponseItem struct {
	E ConnectorEventLatestListResponseItemsE `json:"e,required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n,required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                                  `json:"t,required"`
	JSON connectorEventLatestListResponseItemJSON `json:"-"`
}

// connectorEventLatestListResponseItemJSON contains the JSON metadata for the
// struct [ConnectorEventLatestListResponseItem]
type connectorEventLatestListResponseItemJSON struct {
	E           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventLatestListResponseItemsE struct {
	// Initialized process
	K ConnectorEventLatestListResponseItemsEK `json:"k,required"`
	// Location of upgrade bundle
	URL   string                                     `json:"url"`
	JSON  connectorEventLatestListResponseItemsEJSON `json:"-"`
	union ConnectorEventLatestListResponseItemsEUnion
}

// connectorEventLatestListResponseItemsEJSON contains the JSON metadata for the
// struct [ConnectorEventLatestListResponseItemsE]
type connectorEventLatestListResponseItemsEJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r connectorEventLatestListResponseItemsEJSON) RawJSON() string {
	return r.raw
}

func (r *ConnectorEventLatestListResponseItemsE) UnmarshalJSON(data []byte) (err error) {
	*r = ConnectorEventLatestListResponseItemsE{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ConnectorEventLatestListResponseItemsEUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ConnectorEventLatestListResponseItemsEInit],
// [ConnectorEventLatestListResponseItemsELeave],
// [ConnectorEventLatestListResponseItemsEStartAttestation],
// [ConnectorEventLatestListResponseItemsEFinishAttestationSuccess],
// [ConnectorEventLatestListResponseItemsEFinishAttestationFailure],
// [ConnectorEventLatestListResponseItemsEStartRotateCryptKey],
// [ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccess],
// [ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailure],
// [ConnectorEventLatestListResponseItemsEStartRotatePki],
// [ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccess],
// [ConnectorEventLatestListResponseItemsEFinishRotatePkiFailure],
// [ConnectorEventLatestListResponseItemsEStartUpgrade],
// [ConnectorEventLatestListResponseItemsEFinishUpgradeSuccess],
// [ConnectorEventLatestListResponseItemsEFinishUpgradeFailure],
// [ConnectorEventLatestListResponseItemsEReconcile],
// [ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnel].
func (r ConnectorEventLatestListResponseItemsE) AsUnion() ConnectorEventLatestListResponseItemsEUnion {
	return r.union
}

// Union satisfied by [ConnectorEventLatestListResponseItemsEInit],
// [ConnectorEventLatestListResponseItemsELeave],
// [ConnectorEventLatestListResponseItemsEStartAttestation],
// [ConnectorEventLatestListResponseItemsEFinishAttestationSuccess],
// [ConnectorEventLatestListResponseItemsEFinishAttestationFailure],
// [ConnectorEventLatestListResponseItemsEStartRotateCryptKey],
// [ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccess],
// [ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailure],
// [ConnectorEventLatestListResponseItemsEStartRotatePki],
// [ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccess],
// [ConnectorEventLatestListResponseItemsEFinishRotatePkiFailure],
// [ConnectorEventLatestListResponseItemsEStartUpgrade],
// [ConnectorEventLatestListResponseItemsEFinishUpgradeSuccess],
// [ConnectorEventLatestListResponseItemsEFinishUpgradeFailure],
// [ConnectorEventLatestListResponseItemsEReconcile] or
// [ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnel].
type ConnectorEventLatestListResponseItemsEUnion interface {
	implementsConnectorEventLatestListResponseItemsE()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ConnectorEventLatestListResponseItemsEUnion)(nil)).Elem(),
		"k",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEInit{}),
			DiscriminatorValue: "Init",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsELeave{}),
			DiscriminatorValue: "Leave",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEStartAttestation{}),
			DiscriminatorValue: "StartAttestation",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishAttestationSuccess{}),
			DiscriminatorValue: "FinishAttestationSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishAttestationFailure{}),
			DiscriminatorValue: "FinishAttestationFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEStartRotateCryptKey{}),
			DiscriminatorValue: "StartRotateCryptKey",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccess{}),
			DiscriminatorValue: "FinishRotateCryptKeySuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailure{}),
			DiscriminatorValue: "FinishRotateCryptKeyFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEStartRotatePki{}),
			DiscriminatorValue: "StartRotatePki",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccess{}),
			DiscriminatorValue: "FinishRotatePkiSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishRotatePkiFailure{}),
			DiscriminatorValue: "FinishRotatePkiFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEStartUpgrade{}),
			DiscriminatorValue: "StartUpgrade",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishUpgradeSuccess{}),
			DiscriminatorValue: "FinishUpgradeSuccess",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEFinishUpgradeFailure{}),
			DiscriminatorValue: "FinishUpgradeFailure",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEReconcile{}),
			DiscriminatorValue: "Reconcile",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnel{}),
			DiscriminatorValue: "ConfigureCloudflaredTunnel",
		},
	)
}

type ConnectorEventLatestListResponseItemsEInit struct {
	// Initialized process
	K    ConnectorEventLatestListResponseItemsEInitK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEInitJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEInitJSON contains the JSON metadata for
// the struct [ConnectorEventLatestListResponseItemsEInit]
type connectorEventLatestListResponseItemsEInitJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEInit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEInitJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEInit) implementsConnectorEventLatestListResponseItemsE() {
}

// Initialized process
type ConnectorEventLatestListResponseItemsEInitK string

const (
	ConnectorEventLatestListResponseItemsEInitKInit ConnectorEventLatestListResponseItemsEInitK = "Init"
)

func (r ConnectorEventLatestListResponseItemsEInitK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEInitKInit:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsELeave struct {
	// Stopped process
	K    ConnectorEventLatestListResponseItemsELeaveK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsELeaveJSON `json:"-"`
}

// connectorEventLatestListResponseItemsELeaveJSON contains the JSON metadata for
// the struct [ConnectorEventLatestListResponseItemsELeave]
type connectorEventLatestListResponseItemsELeaveJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsELeave) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsELeaveJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsELeave) implementsConnectorEventLatestListResponseItemsE() {
}

// Stopped process
type ConnectorEventLatestListResponseItemsELeaveK string

const (
	ConnectorEventLatestListResponseItemsELeaveKLeave ConnectorEventLatestListResponseItemsELeaveK = "Leave"
)

func (r ConnectorEventLatestListResponseItemsELeaveK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsELeaveKLeave:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEStartAttestation struct {
	// Started attestation
	K    ConnectorEventLatestListResponseItemsEStartAttestationK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEStartAttestationJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEStartAttestationJSON contains the JSON
// metadata for the struct [ConnectorEventLatestListResponseItemsEStartAttestation]
type connectorEventLatestListResponseItemsEStartAttestationJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEStartAttestation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEStartAttestationJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEStartAttestation) implementsConnectorEventLatestListResponseItemsE() {
}

// Started attestation
type ConnectorEventLatestListResponseItemsEStartAttestationK string

const (
	ConnectorEventLatestListResponseItemsEStartAttestationKStartAttestation ConnectorEventLatestListResponseItemsEStartAttestationK = "StartAttestation"
)

func (r ConnectorEventLatestListResponseItemsEStartAttestationK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEStartAttestationKStartAttestation:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishAttestationSuccess struct {
	// Finished attestation
	K    ConnectorEventLatestListResponseItemsEFinishAttestationSuccessK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishAttestationSuccessJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishAttestationSuccessJSON contains the
// JSON metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishAttestationSuccess]
type connectorEventLatestListResponseItemsEFinishAttestationSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishAttestationSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishAttestationSuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishAttestationSuccess) implementsConnectorEventLatestListResponseItemsE() {
}

// Finished attestation
type ConnectorEventLatestListResponseItemsEFinishAttestationSuccessK string

const (
	ConnectorEventLatestListResponseItemsEFinishAttestationSuccessKFinishAttestationSuccess ConnectorEventLatestListResponseItemsEFinishAttestationSuccessK = "FinishAttestationSuccess"
)

func (r ConnectorEventLatestListResponseItemsEFinishAttestationSuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishAttestationSuccessKFinishAttestationSuccess:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishAttestationFailure struct {
	// Failed attestation
	K    ConnectorEventLatestListResponseItemsEFinishAttestationFailureK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishAttestationFailureJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishAttestationFailureJSON contains the
// JSON metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishAttestationFailure]
type connectorEventLatestListResponseItemsEFinishAttestationFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishAttestationFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishAttestationFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishAttestationFailure) implementsConnectorEventLatestListResponseItemsE() {
}

// Failed attestation
type ConnectorEventLatestListResponseItemsEFinishAttestationFailureK string

const (
	ConnectorEventLatestListResponseItemsEFinishAttestationFailureKFinishAttestationFailure ConnectorEventLatestListResponseItemsEFinishAttestationFailureK = "FinishAttestationFailure"
)

func (r ConnectorEventLatestListResponseItemsEFinishAttestationFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishAttestationFailureKFinishAttestationFailure:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEStartRotateCryptKey struct {
	// Started crypt key rotation
	K    ConnectorEventLatestListResponseItemsEStartRotateCryptKeyK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEStartRotateCryptKeyJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEStartRotateCryptKeyJSON contains the JSON
// metadata for the struct
// [ConnectorEventLatestListResponseItemsEStartRotateCryptKey]
type connectorEventLatestListResponseItemsEStartRotateCryptKeyJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEStartRotateCryptKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEStartRotateCryptKeyJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEStartRotateCryptKey) implementsConnectorEventLatestListResponseItemsE() {
}

// Started crypt key rotation
type ConnectorEventLatestListResponseItemsEStartRotateCryptKeyK string

const (
	ConnectorEventLatestListResponseItemsEStartRotateCryptKeyKStartRotateCryptKey ConnectorEventLatestListResponseItemsEStartRotateCryptKeyK = "StartRotateCryptKey"
)

func (r ConnectorEventLatestListResponseItemsEStartRotateCryptKeyK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEStartRotateCryptKeyKStartRotateCryptKey:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccess struct {
	// Finished crypt key rotation
	K    ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessJSON contains
// the JSON metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccess]
type connectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccess) implementsConnectorEventLatestListResponseItemsE() {
}

// Finished crypt key rotation
type ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessK string

const (
	ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessKFinishRotateCryptKeySuccess ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessK = "FinishRotateCryptKeySuccess"
)

func (r ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishRotateCryptKeySuccessKFinishRotateCryptKeySuccess:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailure struct {
	// Failed crypt key rotation
	K    ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureJSON contains
// the JSON metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailure]
type connectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailure) implementsConnectorEventLatestListResponseItemsE() {
}

// Failed crypt key rotation
type ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureK string

const (
	ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureKFinishRotateCryptKeyFailure ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureK = "FinishRotateCryptKeyFailure"
)

func (r ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishRotateCryptKeyFailureKFinishRotateCryptKeyFailure:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEStartRotatePki struct {
	// Started PKI rotation
	K    ConnectorEventLatestListResponseItemsEStartRotatePkiK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEStartRotatePkiJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEStartRotatePkiJSON contains the JSON
// metadata for the struct [ConnectorEventLatestListResponseItemsEStartRotatePki]
type connectorEventLatestListResponseItemsEStartRotatePkiJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEStartRotatePki) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEStartRotatePkiJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEStartRotatePki) implementsConnectorEventLatestListResponseItemsE() {
}

// Started PKI rotation
type ConnectorEventLatestListResponseItemsEStartRotatePkiK string

const (
	ConnectorEventLatestListResponseItemsEStartRotatePkiKStartRotatePki ConnectorEventLatestListResponseItemsEStartRotatePkiK = "StartRotatePki"
)

func (r ConnectorEventLatestListResponseItemsEStartRotatePkiK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEStartRotatePkiKStartRotatePki:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccess struct {
	// Finished PKI rotation
	K    ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccessK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishRotatePkiSuccessJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishRotatePkiSuccessJSON contains the
// JSON metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccess]
type connectorEventLatestListResponseItemsEFinishRotatePkiSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishRotatePkiSuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccess) implementsConnectorEventLatestListResponseItemsE() {
}

// Finished PKI rotation
type ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccessK string

const (
	ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccessKFinishRotatePkiSuccess ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccessK = "FinishRotatePkiSuccess"
)

func (r ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishRotatePkiSuccessKFinishRotatePkiSuccess:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishRotatePkiFailure struct {
	// Failed PKI rotation
	K    ConnectorEventLatestListResponseItemsEFinishRotatePkiFailureK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishRotatePkiFailureJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishRotatePkiFailureJSON contains the
// JSON metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishRotatePkiFailure]
type connectorEventLatestListResponseItemsEFinishRotatePkiFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishRotatePkiFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishRotatePkiFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishRotatePkiFailure) implementsConnectorEventLatestListResponseItemsE() {
}

// Failed PKI rotation
type ConnectorEventLatestListResponseItemsEFinishRotatePkiFailureK string

const (
	ConnectorEventLatestListResponseItemsEFinishRotatePkiFailureKFinishRotatePkiFailure ConnectorEventLatestListResponseItemsEFinishRotatePkiFailureK = "FinishRotatePkiFailure"
)

func (r ConnectorEventLatestListResponseItemsEFinishRotatePkiFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishRotatePkiFailureKFinishRotatePkiFailure:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEStartUpgrade struct {
	// Started upgrade
	K ConnectorEventLatestListResponseItemsEStartUpgradeK `json:"k,required"`
	// Location of upgrade bundle
	URL  string                                                 `json:"url,required"`
	JSON connectorEventLatestListResponseItemsEStartUpgradeJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEStartUpgradeJSON contains the JSON
// metadata for the struct [ConnectorEventLatestListResponseItemsEStartUpgrade]
type connectorEventLatestListResponseItemsEStartUpgradeJSON struct {
	K           apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEStartUpgrade) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEStartUpgradeJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEStartUpgrade) implementsConnectorEventLatestListResponseItemsE() {
}

// Started upgrade
type ConnectorEventLatestListResponseItemsEStartUpgradeK string

const (
	ConnectorEventLatestListResponseItemsEStartUpgradeKStartUpgrade ConnectorEventLatestListResponseItemsEStartUpgradeK = "StartUpgrade"
)

func (r ConnectorEventLatestListResponseItemsEStartUpgradeK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEStartUpgradeKStartUpgrade:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishUpgradeSuccess struct {
	// Finished upgrade
	K    ConnectorEventLatestListResponseItemsEFinishUpgradeSuccessK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishUpgradeSuccessJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishUpgradeSuccessJSON contains the JSON
// metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishUpgradeSuccess]
type connectorEventLatestListResponseItemsEFinishUpgradeSuccessJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishUpgradeSuccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishUpgradeSuccessJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishUpgradeSuccess) implementsConnectorEventLatestListResponseItemsE() {
}

// Finished upgrade
type ConnectorEventLatestListResponseItemsEFinishUpgradeSuccessK string

const (
	ConnectorEventLatestListResponseItemsEFinishUpgradeSuccessKFinishUpgradeSuccess ConnectorEventLatestListResponseItemsEFinishUpgradeSuccessK = "FinishUpgradeSuccess"
)

func (r ConnectorEventLatestListResponseItemsEFinishUpgradeSuccessK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishUpgradeSuccessKFinishUpgradeSuccess:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEFinishUpgradeFailure struct {
	// Failed upgrade
	K    ConnectorEventLatestListResponseItemsEFinishUpgradeFailureK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEFinishUpgradeFailureJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEFinishUpgradeFailureJSON contains the JSON
// metadata for the struct
// [ConnectorEventLatestListResponseItemsEFinishUpgradeFailure]
type connectorEventLatestListResponseItemsEFinishUpgradeFailureJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEFinishUpgradeFailure) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEFinishUpgradeFailureJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEFinishUpgradeFailure) implementsConnectorEventLatestListResponseItemsE() {
}

// Failed upgrade
type ConnectorEventLatestListResponseItemsEFinishUpgradeFailureK string

const (
	ConnectorEventLatestListResponseItemsEFinishUpgradeFailureKFinishUpgradeFailure ConnectorEventLatestListResponseItemsEFinishUpgradeFailureK = "FinishUpgradeFailure"
)

func (r ConnectorEventLatestListResponseItemsEFinishUpgradeFailureK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEFinishUpgradeFailureKFinishUpgradeFailure:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEReconcile struct {
	// Reconciled
	K    ConnectorEventLatestListResponseItemsEReconcileK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEReconcileJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEReconcileJSON contains the JSON metadata
// for the struct [ConnectorEventLatestListResponseItemsEReconcile]
type connectorEventLatestListResponseItemsEReconcileJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEReconcile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEReconcileJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEReconcile) implementsConnectorEventLatestListResponseItemsE() {
}

// Reconciled
type ConnectorEventLatestListResponseItemsEReconcileK string

const (
	ConnectorEventLatestListResponseItemsEReconcileKReconcile ConnectorEventLatestListResponseItemsEReconcileK = "Reconcile"
)

func (r ConnectorEventLatestListResponseItemsEReconcileK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEReconcileKReconcile:
		return true
	}
	return false
}

type ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnel struct {
	// Configured Cloudflared tunnel
	K    ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnelK    `json:"k,required"`
	JSON connectorEventLatestListResponseItemsEConfigureCloudflaredTunnelJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEConfigureCloudflaredTunnelJSON contains
// the JSON metadata for the struct
// [ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnel]
type connectorEventLatestListResponseItemsEConfigureCloudflaredTunnelJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEConfigureCloudflaredTunnelJSON) RawJSON() string {
	return r.raw
}

func (r ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnel) implementsConnectorEventLatestListResponseItemsE() {
}

// Configured Cloudflared tunnel
type ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnelK string

const (
	ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnelKConfigureCloudflaredTunnel ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnelK = "ConfigureCloudflaredTunnel"
)

func (r ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnelK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEConfigureCloudflaredTunnelKConfigureCloudflaredTunnel:
		return true
	}
	return false
}

// Initialized process
type ConnectorEventLatestListResponseItemsEK string

const (
	ConnectorEventLatestListResponseItemsEKInit                        ConnectorEventLatestListResponseItemsEK = "Init"
	ConnectorEventLatestListResponseItemsEKLeave                       ConnectorEventLatestListResponseItemsEK = "Leave"
	ConnectorEventLatestListResponseItemsEKStartAttestation            ConnectorEventLatestListResponseItemsEK = "StartAttestation"
	ConnectorEventLatestListResponseItemsEKFinishAttestationSuccess    ConnectorEventLatestListResponseItemsEK = "FinishAttestationSuccess"
	ConnectorEventLatestListResponseItemsEKFinishAttestationFailure    ConnectorEventLatestListResponseItemsEK = "FinishAttestationFailure"
	ConnectorEventLatestListResponseItemsEKStartRotateCryptKey         ConnectorEventLatestListResponseItemsEK = "StartRotateCryptKey"
	ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeySuccess ConnectorEventLatestListResponseItemsEK = "FinishRotateCryptKeySuccess"
	ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeyFailure ConnectorEventLatestListResponseItemsEK = "FinishRotateCryptKeyFailure"
	ConnectorEventLatestListResponseItemsEKStartRotatePki              ConnectorEventLatestListResponseItemsEK = "StartRotatePki"
	ConnectorEventLatestListResponseItemsEKFinishRotatePkiSuccess      ConnectorEventLatestListResponseItemsEK = "FinishRotatePkiSuccess"
	ConnectorEventLatestListResponseItemsEKFinishRotatePkiFailure      ConnectorEventLatestListResponseItemsEK = "FinishRotatePkiFailure"
	ConnectorEventLatestListResponseItemsEKStartUpgrade                ConnectorEventLatestListResponseItemsEK = "StartUpgrade"
	ConnectorEventLatestListResponseItemsEKFinishUpgradeSuccess        ConnectorEventLatestListResponseItemsEK = "FinishUpgradeSuccess"
	ConnectorEventLatestListResponseItemsEKFinishUpgradeFailure        ConnectorEventLatestListResponseItemsEK = "FinishUpgradeFailure"
	ConnectorEventLatestListResponseItemsEKReconcile                   ConnectorEventLatestListResponseItemsEK = "Reconcile"
	ConnectorEventLatestListResponseItemsEKConfigureCloudflaredTunnel  ConnectorEventLatestListResponseItemsEK = "ConfigureCloudflaredTunnel"
)

func (r ConnectorEventLatestListResponseItemsEK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEKInit, ConnectorEventLatestListResponseItemsEKLeave, ConnectorEventLatestListResponseItemsEKStartAttestation, ConnectorEventLatestListResponseItemsEKFinishAttestationSuccess, ConnectorEventLatestListResponseItemsEKFinishAttestationFailure, ConnectorEventLatestListResponseItemsEKStartRotateCryptKey, ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeySuccess, ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeyFailure, ConnectorEventLatestListResponseItemsEKStartRotatePki, ConnectorEventLatestListResponseItemsEKFinishRotatePkiSuccess, ConnectorEventLatestListResponseItemsEKFinishRotatePkiFailure, ConnectorEventLatestListResponseItemsEKStartUpgrade, ConnectorEventLatestListResponseItemsEKFinishUpgradeSuccess, ConnectorEventLatestListResponseItemsEKFinishUpgradeFailure, ConnectorEventLatestListResponseItemsEKReconcile, ConnectorEventLatestListResponseItemsEKConfigureCloudflaredTunnel:
		return true
	}
	return false
}

type ConnectorEventLatestListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectorEventLatestListResponseEnvelope struct {
	Result   ConnectorEventLatestListResponse                   `json:"result,required"`
	Success  bool                                               `json:"success,required"`
	Errors   []ConnectorEventLatestListResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventLatestListResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventLatestListResponseEnvelopeJSON       `json:"-"`
}

// connectorEventLatestListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConnectorEventLatestListResponseEnvelope]
type connectorEventLatestListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventLatestListResponseEnvelopeErrors struct {
	Code    float64                                            `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    connectorEventLatestListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventLatestListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ConnectorEventLatestListResponseEnvelopeErrors]
type connectorEventLatestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventLatestListResponseEnvelopeMessages struct {
	Code    float64                                              `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    connectorEventLatestListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventLatestListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ConnectorEventLatestListResponseEnvelopeMessages]
type connectorEventLatestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
