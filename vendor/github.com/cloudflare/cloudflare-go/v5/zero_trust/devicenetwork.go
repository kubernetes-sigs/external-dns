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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DeviceNetworkService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceNetworkService] method instead.
type DeviceNetworkService struct {
	Options []option.RequestOption
}

// NewDeviceNetworkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceNetworkService(opts ...option.RequestOption) (r *DeviceNetworkService) {
	r = &DeviceNetworkService{}
	r.Options = opts
	return
}

// Creates a new device managed network.
func (r *DeviceNetworkService) New(ctx context.Context, params DeviceNetworkNewParams, opts ...option.RequestOption) (res *DeviceNetwork, err error) {
	var env DeviceNetworkNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured device managed network.
func (r *DeviceNetworkService) Update(ctx context.Context, networkID string, params DeviceNetworkUpdateParams, opts ...option.RequestOption) (res *DeviceNetwork, err error) {
	var env DeviceNetworkUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks/%s", params.AccountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a list of managed networks for an account.
func (r *DeviceNetworkService) List(ctx context.Context, query DeviceNetworkListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DeviceNetwork], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks", query.AccountID)
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

// Fetches a list of managed networks for an account.
func (r *DeviceNetworkService) ListAutoPaging(ctx context.Context, query DeviceNetworkListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DeviceNetwork] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a device managed network and fetches a list of the remaining device
// managed networks for an account.
func (r *DeviceNetworkService) Delete(ctx context.Context, networkID string, body DeviceNetworkDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[DeviceNetwork], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks/%s", body.AccountID, networkID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Deletes a device managed network and fetches a list of the remaining device
// managed networks for an account.
func (r *DeviceNetworkService) DeleteAutoPaging(ctx context.Context, networkID string, body DeviceNetworkDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DeviceNetwork] {
	return pagination.NewSinglePageAutoPager(r.Delete(ctx, networkID, body, opts...))
}

// Fetches details for a single managed network.
func (r *DeviceNetworkService) Get(ctx context.Context, networkID string, query DeviceNetworkGetParams, opts ...option.RequestOption) (res *DeviceNetwork, err error) {
	var env DeviceNetworkGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if networkID == "" {
		err = errors.New("missing required network_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/networks/%s", query.AccountID, networkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceNetwork struct {
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config DeviceNetworkConfig `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name string `json:"name"`
	// API UUID.
	NetworkID string `json:"network_id"`
	// The type of device managed network.
	Type DeviceNetworkType `json:"type"`
	JSON deviceNetworkJSON `json:"-"`
}

// deviceNetworkJSON contains the JSON metadata for the struct [DeviceNetwork]
type deviceNetworkJSON struct {
	Config      apijson.Field
	Name        apijson.Field
	NetworkID   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceNetworkJSON) RawJSON() string {
	return r.raw
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr string `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 string                  `json:"sha256"`
	JSON   deviceNetworkConfigJSON `json:"-"`
}

// deviceNetworkConfigJSON contains the JSON metadata for the struct
// [DeviceNetworkConfig]
type deviceNetworkConfigJSON struct {
	TLSSockaddr apijson.Field
	Sha256      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceNetworkConfigJSON) RawJSON() string {
	return r.raw
}

// The type of device managed network.
type DeviceNetworkType string

const (
	DeviceNetworkTypeTLS DeviceNetworkType = "tls"
)

func (r DeviceNetworkType) IsKnown() bool {
	switch r {
	case DeviceNetworkTypeTLS:
		return true
	}
	return false
}

type DeviceNetworkNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[DeviceNetworkNewParamsConfig] `json:"config,required"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name,required"`
	// The type of device managed network.
	Type param.Field[DeviceNetworkNewParamsType] `json:"type,required"`
}

func (r DeviceNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkNewParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r DeviceNetworkNewParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type DeviceNetworkNewParamsType string

const (
	DeviceNetworkNewParamsTypeTLS DeviceNetworkNewParamsType = "tls"
)

func (r DeviceNetworkNewParamsType) IsKnown() bool {
	switch r {
	case DeviceNetworkNewParamsTypeTLS:
		return true
	}
	return false
}

type DeviceNetworkNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DeviceNetwork         `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// deviceNetworkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceNetworkNewResponseEnvelope]
type deviceNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceNetworkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceNetworkNewResponseEnvelopeSuccess bool

const (
	DeviceNetworkNewResponseEnvelopeSuccessTrue DeviceNetworkNewResponseEnvelopeSuccess = true
)

func (r DeviceNetworkNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceNetworkNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceNetworkUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The configuration object containing information for the WARP client to detect
	// the managed network.
	Config param.Field[DeviceNetworkUpdateParamsConfig] `json:"config"`
	// The name of the device managed network. This name must be unique.
	Name param.Field[string] `json:"name"`
	// The type of device managed network.
	Type param.Field[DeviceNetworkUpdateParamsType] `json:"type"`
}

func (r DeviceNetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration object containing information for the WARP client to detect
// the managed network.
type DeviceNetworkUpdateParamsConfig struct {
	// A network address of the form "host:port" that the WARP client will use to
	// detect the presence of a TLS host.
	TLSSockaddr param.Field[string] `json:"tls_sockaddr,required"`
	// The SHA-256 hash of the TLS certificate presented by the host found at
	// tls_sockaddr. If absent, regular certificate verification (trusted roots, valid
	// timestamp, etc) will be used to validate the certificate.
	Sha256 param.Field[string] `json:"sha256"`
}

func (r DeviceNetworkUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of device managed network.
type DeviceNetworkUpdateParamsType string

const (
	DeviceNetworkUpdateParamsTypeTLS DeviceNetworkUpdateParamsType = "tls"
)

func (r DeviceNetworkUpdateParamsType) IsKnown() bool {
	switch r {
	case DeviceNetworkUpdateParamsTypeTLS:
		return true
	}
	return false
}

type DeviceNetworkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DeviceNetwork         `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceNetworkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceNetworkUpdateResponseEnvelopeJSON    `json:"-"`
}

// deviceNetworkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceNetworkUpdateResponseEnvelope]
type deviceNetworkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceNetworkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceNetworkUpdateResponseEnvelopeSuccess bool

const (
	DeviceNetworkUpdateResponseEnvelopeSuccessTrue DeviceNetworkUpdateResponseEnvelopeSuccess = true
)

func (r DeviceNetworkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceNetworkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceNetworkListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceNetworkDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceNetworkGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceNetworkGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DeviceNetwork         `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceNetworkGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceNetworkGetResponseEnvelopeJSON    `json:"-"`
}

// deviceNetworkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceNetworkGetResponseEnvelope]
type deviceNetworkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceNetworkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceNetworkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceNetworkGetResponseEnvelopeSuccess bool

const (
	DeviceNetworkGetResponseEnvelopeSuccessTrue DeviceNetworkGetResponseEnvelopeSuccess = true
)

func (r DeviceNetworkGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceNetworkGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
