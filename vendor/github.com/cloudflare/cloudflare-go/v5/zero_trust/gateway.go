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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// GatewayService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayService] method instead.
type GatewayService struct {
	Options          []option.RequestOption
	AuditSSHSettings *GatewayAuditSSHSettingService
	Categories       *GatewayCategoryService
	AppTypes         *GatewayAppTypeService
	Configurations   *GatewayConfigurationService
	Lists            *GatewayListService
	Locations        *GatewayLocationService
	Logging          *GatewayLoggingService
	ProxyEndpoints   *GatewayProxyEndpointService
	Rules            *GatewayRuleService
	Certificates     *GatewayCertificateService
}

// NewGatewayService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGatewayService(opts ...option.RequestOption) (r *GatewayService) {
	r = &GatewayService{}
	r.Options = opts
	r.AuditSSHSettings = NewGatewayAuditSSHSettingService(opts...)
	r.Categories = NewGatewayCategoryService(opts...)
	r.AppTypes = NewGatewayAppTypeService(opts...)
	r.Configurations = NewGatewayConfigurationService(opts...)
	r.Lists = NewGatewayListService(opts...)
	r.Locations = NewGatewayLocationService(opts...)
	r.Logging = NewGatewayLoggingService(opts...)
	r.ProxyEndpoints = NewGatewayProxyEndpointService(opts...)
	r.Rules = NewGatewayRuleService(opts...)
	r.Certificates = NewGatewayCertificateService(opts...)
	return
}

// Creates a Zero Trust account with an existing Cloudflare account.
func (r *GatewayService) New(ctx context.Context, body GatewayNewParams, opts ...option.RequestOption) (res *GatewayNewResponse, err error) {
	var env GatewayNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information about the current Zero Trust account.
func (r *GatewayService) List(ctx context.Context, query GatewayListParams, opts ...option.RequestOption) (res *GatewayListResponse, err error) {
	var env GatewayListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type GatewayNewResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                 `json:"provider_name"`
	JSON         gatewayNewResponseJSON `json:"-"`
}

// gatewayNewResponseJSON contains the JSON metadata for the struct
// [GatewayNewResponse]
type gatewayNewResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayNewResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayListResponse struct {
	// Cloudflare account ID.
	ID string `json:"id"`
	// Gateway internal ID.
	GatewayTag string `json:"gateway_tag"`
	// The name of the provider. Usually Cloudflare.
	ProviderName string                  `json:"provider_name"`
	JSON         gatewayListResponseJSON `json:"-"`
}

// gatewayListResponseJSON contains the JSON metadata for the struct
// [GatewayListResponse]
type gatewayListResponseJSON struct {
	ID           apijson.Field
	GatewayTag   apijson.Field
	ProviderName apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GatewayListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListResponseJSON) RawJSON() string {
	return r.raw
}

type GatewayNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayNewResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayNewResponse                `json:"result"`
	JSON    gatewayNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayNewResponseEnvelope]
type gatewayNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayNewResponseEnvelopeSuccess bool

const (
	GatewayNewResponseEnvelopeSuccessTrue GatewayNewResponseEnvelopeSuccess = true
)

func (r GatewayNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayListResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayListResponse                `json:"result"`
	JSON    gatewayListResponseEnvelopeJSON    `json:"-"`
}

// gatewayListResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayListResponseEnvelope]
type gatewayListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayListResponseEnvelopeSuccess bool

const (
	GatewayListResponseEnvelopeSuccessTrue GatewayListResponseEnvelopeSuccess = true
)

func (r GatewayListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
