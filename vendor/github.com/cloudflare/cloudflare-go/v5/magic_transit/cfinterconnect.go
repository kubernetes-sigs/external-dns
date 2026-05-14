// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// CfInterconnectService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCfInterconnectService] method instead.
type CfInterconnectService struct {
	Options []option.RequestOption
}

// NewCfInterconnectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCfInterconnectService(opts ...option.RequestOption) (r *CfInterconnectService) {
	r = &CfInterconnectService{}
	r.Options = opts
	return
}

// Updates a specific interconnect associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *CfInterconnectService) Update(ctx context.Context, cfInterconnectID string, params CfInterconnectUpdateParams, opts ...option.RequestOption) (res *CfInterconnectUpdateResponse, err error) {
	var env CfInterconnectUpdateResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", params.AccountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists interconnects associated with an account.
func (r *CfInterconnectService) List(ctx context.Context, params CfInterconnectListParams, opts ...option.RequestOption) (res *CfInterconnectListResponse, err error) {
	var env CfInterconnectListResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates multiple interconnects associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *CfInterconnectService) BulkUpdate(ctx context.Context, params CfInterconnectBulkUpdateParams, opts ...option.RequestOption) (res *CfInterconnectBulkUpdateResponse, err error) {
	var env CfInterconnectBulkUpdateResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific interconnect.
func (r *CfInterconnectService) Get(ctx context.Context, cfInterconnectID string, params CfInterconnectGetParams, opts ...option.RequestOption) (res *CfInterconnectGetResponse, err error) {
	var env CfInterconnectGetResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cfInterconnectID == "" {
		err = errors.New("missing required cf_interconnect_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/cf_interconnects/%s", params.AccountID, cfInterconnectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CfInterconnectUpdateResponse struct {
	Modified             bool                                             `json:"modified"`
	ModifiedInterconnect CfInterconnectUpdateResponseModifiedInterconnect `json:"modified_interconnect"`
	JSON                 cfInterconnectUpdateResponseJSON                 `json:"-"`
}

// cfInterconnectUpdateResponseJSON contains the JSON metadata for the struct
// [CfInterconnectUpdateResponse]
type cfInterconnectUpdateResponseJSON struct {
	Modified             apijson.Field
	ModifiedInterconnect apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectUpdateResponseModifiedInterconnect struct {
	// Identifier
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         CfInterconnectUpdateResponseModifiedInterconnectGRE `json:"gre"`
	HealthCheck HealthCheck                                         `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                                               `json:"name"`
	JSON cfInterconnectUpdateResponseModifiedInterconnectJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectJSON contains the JSON metadata
// for the struct [CfInterconnectUpdateResponseModifiedInterconnect]
type cfInterconnectUpdateResponseModifiedInterconnectJSON struct {
	ID                apijson.Field
	ColoName          apijson.Field
	CreatedOn         apijson.Field
	Description       apijson.Field
	GRE               apijson.Field
	HealthCheck       apijson.Field
	InterfaceAddress  apijson.Field
	InterfaceAddress6 apijson.Field
	ModifiedOn        apijson.Field
	Mtu               apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseModifiedInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type CfInterconnectUpdateResponseModifiedInterconnectGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                  `json:"cloudflare_endpoint"`
	JSON               cfInterconnectUpdateResponseModifiedInterconnectGREJSON `json:"-"`
}

// cfInterconnectUpdateResponseModifiedInterconnectGREJSON contains the JSON
// metadata for the struct [CfInterconnectUpdateResponseModifiedInterconnectGRE]
type cfInterconnectUpdateResponseModifiedInterconnectGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseModifiedInterconnectGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseModifiedInterconnectGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectListResponse struct {
	Interconnects []CfInterconnectListResponseInterconnect `json:"interconnects"`
	JSON          cfInterconnectListResponseJSON           `json:"-"`
}

// cfInterconnectListResponseJSON contains the JSON metadata for the struct
// [CfInterconnectListResponse]
type cfInterconnectListResponseJSON struct {
	Interconnects apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CfInterconnectListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectListResponseInterconnect struct {
	// Identifier
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         CfInterconnectListResponseInterconnectsGRE `json:"gre"`
	HealthCheck HealthCheck                                `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                                     `json:"name"`
	JSON cfInterconnectListResponseInterconnectJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectJSON contains the JSON metadata for the
// struct [CfInterconnectListResponseInterconnect]
type cfInterconnectListResponseInterconnectJSON struct {
	ID                apijson.Field
	ColoName          apijson.Field
	CreatedOn         apijson.Field
	Description       apijson.Field
	GRE               apijson.Field
	HealthCheck       apijson.Field
	InterfaceAddress  apijson.Field
	InterfaceAddress6 apijson.Field
	ModifiedOn        apijson.Field
	Mtu               apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CfInterconnectListResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type CfInterconnectListResponseInterconnectsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                         `json:"cloudflare_endpoint"`
	JSON               cfInterconnectListResponseInterconnectsGREJSON `json:"-"`
}

// cfInterconnectListResponseInterconnectsGREJSON contains the JSON metadata for
// the struct [CfInterconnectListResponseInterconnectsGRE]
type cfInterconnectListResponseInterconnectsGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectListResponseInterconnectsGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseInterconnectsGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectBulkUpdateResponse struct {
	Modified              bool                                                   `json:"modified"`
	ModifiedInterconnects []CfInterconnectBulkUpdateResponseModifiedInterconnect `json:"modified_interconnects"`
	JSON                  cfInterconnectBulkUpdateResponseJSON                   `json:"-"`
}

// cfInterconnectBulkUpdateResponseJSON contains the JSON metadata for the struct
// [CfInterconnectBulkUpdateResponse]
type cfInterconnectBulkUpdateResponseJSON struct {
	Modified              apijson.Field
	ModifiedInterconnects apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectBulkUpdateResponseModifiedInterconnect struct {
	// Identifier
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE `json:"gre"`
	HealthCheck HealthCheck                                              `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                                                   `json:"name"`
	JSON cfInterconnectBulkUpdateResponseModifiedInterconnectJSON `json:"-"`
}

// cfInterconnectBulkUpdateResponseModifiedInterconnectJSON contains the JSON
// metadata for the struct [CfInterconnectBulkUpdateResponseModifiedInterconnect]
type cfInterconnectBulkUpdateResponseModifiedInterconnectJSON struct {
	ID                apijson.Field
	ColoName          apijson.Field
	CreatedOn         apijson.Field
	Description       apijson.Field
	GRE               apijson.Field
	HealthCheck       apijson.Field
	InterfaceAddress  apijson.Field
	InterfaceAddress6 apijson.Field
	ModifiedOn        apijson.Field
	Mtu               apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseModifiedInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseModifiedInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                                       `json:"cloudflare_endpoint"`
	JSON               cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON `json:"-"`
}

// cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON contains the JSON
// metadata for the struct
// [CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE]
type cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseModifiedInterconnectsGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseModifiedInterconnectsGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectGetResponse struct {
	Interconnect CfInterconnectGetResponseInterconnect `json:"interconnect"`
	JSON         cfInterconnectGetResponseJSON         `json:"-"`
}

// cfInterconnectGetResponseJSON contains the JSON metadata for the struct
// [CfInterconnectGetResponse]
type cfInterconnectGetResponseJSON struct {
	Interconnect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CfInterconnectGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectGetResponseInterconnect struct {
	// Identifier
	ID string `json:"id"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	ColoName string `json:"colo_name"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// An optional description of the interconnect.
	Description string `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         CfInterconnectGetResponseInterconnectGRE `json:"gre"`
	HealthCheck HealthCheck                              `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu int64 `json:"mtu"`
	// The name of the interconnect. The name cannot share a name with other tunnels.
	Name string                                    `json:"name"`
	JSON cfInterconnectGetResponseInterconnectJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectJSON contains the JSON metadata for the
// struct [CfInterconnectGetResponseInterconnect]
type cfInterconnectGetResponseInterconnectJSON struct {
	ID                apijson.Field
	ColoName          apijson.Field
	CreatedOn         apijson.Field
	Description       apijson.Field
	GRE               apijson.Field
	HealthCheck       apijson.Field
	InterfaceAddress  apijson.Field
	InterfaceAddress6 apijson.Field
	ModifiedOn        apijson.Field
	Mtu               apijson.Field
	Name              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *CfInterconnectGetResponseInterconnect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectJSON) RawJSON() string {
	return r.raw
}

// The configuration specific to GRE interconnects.
type CfInterconnectGetResponseInterconnectGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint string                                       `json:"cloudflare_endpoint"`
	JSON               cfInterconnectGetResponseInterconnectGREJSON `json:"-"`
}

// cfInterconnectGetResponseInterconnectGREJSON contains the JSON metadata for the
// struct [CfInterconnectGetResponseInterconnectGRE]
type cfInterconnectGetResponseInterconnectGREJSON struct {
	CloudflareEndpoint apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CfInterconnectGetResponseInterconnectGRE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseInterconnectGREJSON) RawJSON() string {
	return r.raw
}

type CfInterconnectUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description of the interconnect.
	Description param.Field[string] `json:"description"`
	// The configuration specific to GRE interconnects.
	GRE         param.Field[CfInterconnectUpdateParamsGRE] `json:"gre"`
	HealthCheck param.Field[HealthCheckParam]              `json:"health_check"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 param.Field[string] `json:"interface_address6"`
	// The Maximum Transmission Unit (MTU) in bytes for the interconnect. The minimum
	// value is 576.
	Mtu               param.Field[int64] `json:"mtu"`
	XMagicNewHcTarget param.Field[bool]  `header:"x-magic-new-hc-target"`
}

func (r CfInterconnectUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration specific to GRE interconnects.
type CfInterconnectUpdateParamsGRE struct {
	// The IP address assigned to the Cloudflare side of the GRE tunnel created as part
	// of the Interconnect.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint"`
}

func (r CfInterconnectUpdateParamsGRE) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CfInterconnectUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo        `json:"errors,required"`
	Messages []shared.ResponseInfo        `json:"messages,required"`
	Result   CfInterconnectUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfInterconnectUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfInterconnectUpdateResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CfInterconnectUpdateResponseEnvelope]
type cfInterconnectUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectUpdateResponseEnvelopeSuccess bool

const (
	CfInterconnectUpdateResponseEnvelopeSuccessTrue CfInterconnectUpdateResponseEnvelopeSuccess = true
)

func (r CfInterconnectUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CfInterconnectListParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type CfInterconnectListResponseEnvelope struct {
	Errors   []shared.ResponseInfo      `json:"errors,required"`
	Messages []shared.ResponseInfo      `json:"messages,required"`
	Result   CfInterconnectListResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfInterconnectListResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfInterconnectListResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectListResponseEnvelopeJSON contains the JSON metadata for the struct
// [CfInterconnectListResponseEnvelope]
type cfInterconnectListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectListResponseEnvelopeSuccess bool

const (
	CfInterconnectListResponseEnvelopeSuccessTrue CfInterconnectListResponseEnvelopeSuccess = true
)

func (r CfInterconnectListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CfInterconnectBulkUpdateParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	Body              interface{}         `json:"body,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

func (r CfInterconnectBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CfInterconnectBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo            `json:"errors,required"`
	Messages []shared.ResponseInfo            `json:"messages,required"`
	Result   CfInterconnectBulkUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfInterconnectBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfInterconnectBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CfInterconnectBulkUpdateResponseEnvelope]
type cfInterconnectBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectBulkUpdateResponseEnvelopeSuccess bool

const (
	CfInterconnectBulkUpdateResponseEnvelopeSuccessTrue CfInterconnectBulkUpdateResponseEnvelopeSuccess = true
)

func (r CfInterconnectBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CfInterconnectGetParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type CfInterconnectGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   CfInterconnectGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success CfInterconnectGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    cfInterconnectGetResponseEnvelopeJSON    `json:"-"`
}

// cfInterconnectGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CfInterconnectGetResponseEnvelope]
type cfInterconnectGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CfInterconnectGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cfInterconnectGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type CfInterconnectGetResponseEnvelopeSuccess bool

const (
	CfInterconnectGetResponseEnvelopeSuccessTrue CfInterconnectGetResponseEnvelopeSuccess = true
)

func (r CfInterconnectGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CfInterconnectGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
