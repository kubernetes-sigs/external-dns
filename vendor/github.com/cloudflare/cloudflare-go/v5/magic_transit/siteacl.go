// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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

// SiteACLService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSiteACLService] method instead.
type SiteACLService struct {
	Options []option.RequestOption
}

// NewSiteACLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSiteACLService(opts ...option.RequestOption) (r *SiteACLService) {
	r = &SiteACLService{}
	r.Options = opts
	return
}

// Creates a new Site ACL.
func (r *SiteACLService) New(ctx context.Context, siteID string, params SiteACLNewParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env SiteACLNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls", params.AccountID, siteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a specific Site ACL.
func (r *SiteACLService) Update(ctx context.Context, siteID string, aclID string, params SiteACLUpdateParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env SiteACLUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", params.AccountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Site ACLs associated with an account.
func (r *SiteACLService) List(ctx context.Context, siteID string, query SiteACLListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ACL], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls", query.AccountID, siteID)
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

// Lists Site ACLs associated with an account.
func (r *SiteACLService) ListAutoPaging(ctx context.Context, siteID string, query SiteACLListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ACL] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, siteID, query, opts...))
}

// Remove a specific Site ACL.
func (r *SiteACLService) Delete(ctx context.Context, siteID string, aclID string, body SiteACLDeleteParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env SiteACLDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", body.AccountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Patch a specific Site ACL.
func (r *SiteACLService) Edit(ctx context.Context, siteID string, aclID string, params SiteACLEditParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env SiteACLEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", params.AccountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific Site ACL.
func (r *SiteACLService) Get(ctx context.Context, siteID string, aclID string, query SiteACLGetParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env SiteACLGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if siteID == "" {
		err = errors.New("missing required site_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/sites/%s/acls/%s", query.AccountID, siteID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Bidirectional ACL policy for network traffic within a site.
type ACL struct {
	// Identifier
	ID string `json:"id"`
	// Description for the ACL.
	Description string `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic Connector. If not included in request, will default
	// to false.
	ForwardLocally bool             `json:"forward_locally"`
	LAN1           ACLConfiguration `json:"lan_1"`
	LAN2           ACLConfiguration `json:"lan_2"`
	// The name of the ACL.
	Name      string            `json:"name"`
	Protocols []AllowedProtocol `json:"protocols"`
	// The desired traffic direction for this ACL policy. If set to "false", the policy
	// will allow bidirectional traffic. If set to "true", the policy will only allow
	// traffic in one direction. If not included in request, will default to false.
	Unidirectional bool    `json:"unidirectional"`
	JSON           aclJSON `json:"-"`
}

// aclJSON contains the JSON metadata for the struct [ACL]
type aclJSON struct {
	ID             apijson.Field
	Description    apijson.Field
	ForwardLocally apijson.Field
	LAN1           apijson.Field
	LAN2           apijson.Field
	Name           apijson.Field
	Protocols      apijson.Field
	Unidirectional apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclJSON) RawJSON() string {
	return r.raw
}

type ACLConfiguration struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID string `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName string `json:"lan_name"`
	// Array of port ranges on the provided LAN that will be included in the ACL. If no
	// ports or port rangess are provided, communication on any port on this LAN is
	// allowed.
	PortRanges []string `json:"port_ranges"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// or port ranges are provided, communication on any port on this LAN is allowed.
	Ports []int64 `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets []Subnet             `json:"subnets"`
	JSON    aclConfigurationJSON `json:"-"`
}

// aclConfigurationJSON contains the JSON metadata for the struct
// [ACLConfiguration]
type aclConfigurationJSON struct {
	LANID       apijson.Field
	LANName     apijson.Field
	PortRanges  apijson.Field
	Ports       apijson.Field
	Subnets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACLConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclConfigurationJSON) RawJSON() string {
	return r.raw
}

type ACLConfigurationParam struct {
	// The identifier for the LAN you want to create an ACL policy with.
	LANID param.Field[string] `json:"lan_id,required"`
	// The name of the LAN based on the provided lan_id.
	LANName param.Field[string] `json:"lan_name"`
	// Array of port ranges on the provided LAN that will be included in the ACL. If no
	// ports or port rangess are provided, communication on any port on this LAN is
	// allowed.
	PortRanges param.Field[[]string] `json:"port_ranges"`
	// Array of ports on the provided LAN that will be included in the ACL. If no ports
	// or port ranges are provided, communication on any port on this LAN is allowed.
	Ports param.Field[[]int64] `json:"ports"`
	// Array of subnet IPs within the LAN that will be included in the ACL. If no
	// subnets are provided, communication on any subnets on this LAN are allowed.
	Subnets param.Field[[]SubnetParam] `json:"subnets"`
}

func (r ACLConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Array of allowed communication protocols between configured LANs. If no
// protocols are provided, all protocols are allowed.
type AllowedProtocol string

const (
	AllowedProtocolTCP  AllowedProtocol = "tcp"
	AllowedProtocolUdp  AllowedProtocol = "udp"
	AllowedProtocolIcmp AllowedProtocol = "icmp"
)

func (r AllowedProtocol) IsKnown() bool {
	switch r {
	case AllowedProtocolTCP, AllowedProtocolUdp, AllowedProtocolIcmp:
		return true
	}
	return false
}

type Subnet = string

type SubnetParam = string

type SiteACLNewParams struct {
	// Identifier
	AccountID param.Field[string]                `path:"account_id,required"`
	LAN1      param.Field[ACLConfigurationParam] `json:"lan_1,required"`
	LAN2      param.Field[ACLConfigurationParam] `json:"lan_2,required"`
	// The name of the ACL.
	Name param.Field[string] `json:"name,required"`
	// Description for the ACL.
	Description param.Field[string] `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic Connector. If not included in request, will default
	// to false.
	ForwardLocally param.Field[bool]              `json:"forward_locally"`
	Protocols      param.Field[[]AllowedProtocol] `json:"protocols"`
	// The desired traffic direction for this ACL policy. If set to "false", the policy
	// will allow bidirectional traffic. If set to "true", the policy will only allow
	// traffic in one direction. If not included in request, will default to false.
	Unidirectional param.Field[bool] `json:"unidirectional"`
}

func (r SiteACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result ACL `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLNewResponseEnvelopeJSON    `json:"-"`
}

// siteACLNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLNewResponseEnvelope]
type siteACLNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLNewResponseEnvelopeSuccess bool

const (
	SiteACLNewResponseEnvelopeSuccessTrue SiteACLNewResponseEnvelopeSuccess = true
)

func (r SiteACLNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Description for the ACL.
	Description param.Field[string] `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic Connector. If not included in request, will default
	// to false.
	ForwardLocally param.Field[bool]                  `json:"forward_locally"`
	LAN1           param.Field[ACLConfigurationParam] `json:"lan_1"`
	LAN2           param.Field[ACLConfigurationParam] `json:"lan_2"`
	// The name of the ACL.
	Name      param.Field[string]            `json:"name"`
	Protocols param.Field[[]AllowedProtocol] `json:"protocols"`
	// The desired traffic direction for this ACL policy. If set to "false", the policy
	// will allow bidirectional traffic. If set to "true", the policy will only allow
	// traffic in one direction. If not included in request, will default to false.
	Unidirectional param.Field[bool] `json:"unidirectional"`
}

func (r SiteACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result ACL `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLUpdateResponseEnvelopeJSON    `json:"-"`
}

// siteACLUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLUpdateResponseEnvelope]
type siteACLUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLUpdateResponseEnvelopeSuccess bool

const (
	SiteACLUpdateResponseEnvelopeSuccessTrue SiteACLUpdateResponseEnvelopeSuccess = true
)

func (r SiteACLUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteACLDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteACLDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result ACL `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLDeleteResponseEnvelopeJSON    `json:"-"`
}

// siteACLDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLDeleteResponseEnvelope]
type siteACLDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLDeleteResponseEnvelopeSuccess bool

const (
	SiteACLDeleteResponseEnvelopeSuccessTrue SiteACLDeleteResponseEnvelopeSuccess = true
)

func (r SiteACLDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLEditParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Description for the ACL.
	Description param.Field[string] `json:"description"`
	// The desired forwarding action for this ACL policy. If set to "false", the policy
	// will forward traffic to Cloudflare. If set to "true", the policy will forward
	// traffic locally on the Magic Connector. If not included in request, will default
	// to false.
	ForwardLocally param.Field[bool]                  `json:"forward_locally"`
	LAN1           param.Field[ACLConfigurationParam] `json:"lan_1"`
	LAN2           param.Field[ACLConfigurationParam] `json:"lan_2"`
	// The name of the ACL.
	Name      param.Field[string]            `json:"name"`
	Protocols param.Field[[]AllowedProtocol] `json:"protocols"`
	// The desired traffic direction for this ACL policy. If set to "false", the policy
	// will allow bidirectional traffic. If set to "true", the policy will only allow
	// traffic in one direction. If not included in request, will default to false.
	Unidirectional param.Field[bool] `json:"unidirectional"`
}

func (r SiteACLEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SiteACLEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result ACL `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLEditResponseEnvelopeJSON    `json:"-"`
}

// siteACLEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLEditResponseEnvelope]
type siteACLEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLEditResponseEnvelopeSuccess bool

const (
	SiteACLEditResponseEnvelopeSuccessTrue SiteACLEditResponseEnvelopeSuccess = true
)

func (r SiteACLEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SiteACLGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type SiteACLGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Bidirectional ACL policy for network traffic within a site.
	Result ACL `json:"result,required"`
	// Whether the API call was successful
	Success SiteACLGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    siteACLGetResponseEnvelopeJSON    `json:"-"`
}

// siteACLGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SiteACLGetResponseEnvelope]
type siteACLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SiteACLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r siteACLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SiteACLGetResponseEnvelopeSuccess bool

const (
	SiteACLGetResponseEnvelopeSuccessTrue SiteACLGetResponseEnvelopeSuccess = true
)

func (r SiteACLGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SiteACLGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
