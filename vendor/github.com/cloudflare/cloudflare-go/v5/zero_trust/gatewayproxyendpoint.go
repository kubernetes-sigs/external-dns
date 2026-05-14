// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// GatewayProxyEndpointService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayProxyEndpointService] method instead.
type GatewayProxyEndpointService struct {
	Options []option.RequestOption
}

// NewGatewayProxyEndpointService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewGatewayProxyEndpointService(opts ...option.RequestOption) (r *GatewayProxyEndpointService) {
	r = &GatewayProxyEndpointService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) New(ctx context.Context, params GatewayProxyEndpointNewParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	var env GatewayProxyEndpointNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all Zero Trust Gateway proxy endpoints for an account.
func (r *GatewayProxyEndpointService) List(ctx context.Context, query GatewayProxyEndpointListParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	var env GatewayProxyEndpointListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Delete(ctx context.Context, proxyEndpointID string, body GatewayProxyEndpointDeleteParams, opts ...option.RequestOption) (res *GatewayProxyEndpointDeleteResponse, err error) {
	var env GatewayProxyEndpointDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", body.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Edit(ctx context.Context, proxyEndpointID string, params GatewayProxyEndpointEditParams, opts ...option.RequestOption) (res *ProxyEndpoint, err error) {
	var env GatewayProxyEndpointEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", params.AccountID, proxyEndpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) Get(ctx context.Context, proxyEndpointID string, query GatewayProxyEndpointGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProxyEndpoint], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if proxyEndpointID == "" {
		err = errors.New("missing required proxy_endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/proxy_endpoints/%s", query.AccountID, proxyEndpointID)
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

// Fetches a single Zero Trust Gateway proxy endpoint.
func (r *GatewayProxyEndpointService) GetAutoPaging(ctx context.Context, proxyEndpointID string, query GatewayProxyEndpointGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProxyEndpoint] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, proxyEndpointID, query, opts...))
}

type GatewayIPs = string

type GatewayIPsParam = string

type ProxyEndpoint struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// A list of CIDRs to restrict ingress connections.
	IPs []GatewayIPs `json:"ips"`
	// The name of the proxy endpoint.
	Name string `json:"name"`
	// The subdomain to be used as the destination in the proxy client.
	Subdomain string            `json:"subdomain"`
	UpdatedAt time.Time         `json:"updated_at" format:"date-time"`
	JSON      proxyEndpointJSON `json:"-"`
}

// proxyEndpointJSON contains the JSON metadata for the struct [ProxyEndpoint]
type proxyEndpointJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IPs         apijson.Field
	Name        apijson.Field
	Subdomain   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProxyEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r proxyEndpointJSON) RawJSON() string {
	return r.raw
}

type GatewayProxyEndpointDeleteResponse = interface{}

type GatewayProxyEndpointNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]GatewayIPsParam] `json:"ips,required"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name,required"`
}

func (r GatewayProxyEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ProxyEndpoint                                  `json:"result"`
	JSON    gatewayProxyEndpointNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointNewResponseEnvelope]
type gatewayProxyEndpointNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointNewResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointNewResponseEnvelopeSuccessTrue GatewayProxyEndpointNewResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayProxyEndpointListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayProxyEndpointListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointListResponseEnvelopeSuccess `json:"success,required"`
	Result  ProxyEndpoint                                   `json:"result"`
	JSON    gatewayProxyEndpointListResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointListResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointListResponseEnvelope]
type gatewayProxyEndpointListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointListResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointListResponseEnvelopeSuccessTrue GatewayProxyEndpointListResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayProxyEndpointDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayProxyEndpointDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayProxyEndpointDeleteResponse                `json:"result"`
	JSON    gatewayProxyEndpointDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayProxyEndpointDeleteResponseEnvelope]
type gatewayProxyEndpointDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointDeleteResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointDeleteResponseEnvelopeSuccessTrue GatewayProxyEndpointDeleteResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayProxyEndpointEditParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of CIDRs to restrict ingress connections.
	IPs param.Field[[]GatewayIPsParam] `json:"ips"`
	// The name of the proxy endpoint.
	Name param.Field[string] `json:"name"`
}

func (r GatewayProxyEndpointEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayProxyEndpointEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayProxyEndpointEditResponseEnvelopeSuccess `json:"success,required"`
	Result  ProxyEndpoint                                   `json:"result"`
	JSON    gatewayProxyEndpointEditResponseEnvelopeJSON    `json:"-"`
}

// gatewayProxyEndpointEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [GatewayProxyEndpointEditResponseEnvelope]
type gatewayProxyEndpointEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayProxyEndpointEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayProxyEndpointEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayProxyEndpointEditResponseEnvelopeSuccess bool

const (
	GatewayProxyEndpointEditResponseEnvelopeSuccessTrue GatewayProxyEndpointEditResponseEnvelopeSuccess = true
)

func (r GatewayProxyEndpointEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayProxyEndpointEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayProxyEndpointGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
