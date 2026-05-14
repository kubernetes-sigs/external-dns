// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// IPSECTunnelService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPSECTunnelService] method instead.
type IPSECTunnelService struct {
	Options []option.RequestOption
}

// NewIPSECTunnelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIPSECTunnelService(opts ...option.RequestOption) (r *IPSECTunnelService) {
	r = &IPSECTunnelService{}
	r.Options = opts
	return
}

// Creates a new IPsec tunnel associated with an account. Use `?validate_only=true`
// as an optional query parameter to only run validation without persisting
// changes.
func (r *IPSECTunnelService) New(ctx context.Context, params IPSECTunnelNewParams, opts ...option.RequestOption) (res *IPSECTunnelNewResponse, err error) {
	var env IPSECTunnelNewResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a specific IPsec tunnel associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *IPSECTunnelService) Update(ctx context.Context, ipsecTunnelID string, params IPSECTunnelUpdateParams, opts ...option.RequestOption) (res *IPSECTunnelUpdateResponse, err error) {
	var env IPSECTunnelUpdateResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", params.AccountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists IPsec tunnels associated with an account.
func (r *IPSECTunnelService) List(ctx context.Context, params IPSECTunnelListParams, opts ...option.RequestOption) (res *IPSECTunnelListResponse, err error) {
	var env IPSECTunnelListResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables and removes a specific static IPsec Tunnel associated with an account.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *IPSECTunnelService) Delete(ctx context.Context, ipsecTunnelID string, params IPSECTunnelDeleteParams, opts ...option.RequestOption) (res *IPSECTunnelDeleteResponse, err error) {
	var env IPSECTunnelDeleteResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", params.AccountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update multiple IPsec tunnels associated with an account. Use
// `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes.
func (r *IPSECTunnelService) BulkUpdate(ctx context.Context, params IPSECTunnelBulkUpdateParams, opts ...option.RequestOption) (res *IPSECTunnelBulkUpdateResponse, err error) {
	var env IPSECTunnelBulkUpdateResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists details for a specific IPsec tunnel.
func (r *IPSECTunnelService) Get(ctx context.Context, ipsecTunnelID string, params IPSECTunnelGetParams, opts ...option.RequestOption) (res *IPSECTunnelGetResponse, err error) {
	var env IPSECTunnelGetResponseEnvelope
	if params.XMagicNewHcTarget.Present {
		opts = append(opts, option.WithHeader("x-magic-new-hc-target", fmt.Sprintf("%s", params.XMagicNewHcTarget)))
	}
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s", params.AccountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a Pre Shared Key for a specific IPsec tunnel used in the IKE session.
// Use `?validate_only=true` as an optional query parameter to only run validation
// without persisting changes. After a PSK is generated, the PSK is immediately
// persisted to Cloudflare's edge and cannot be retrieved later. Note the PSK in a
// safe place.
func (r *IPSECTunnelService) PSKGenerate(ctx context.Context, ipsecTunnelID string, params IPSECTunnelPSKGenerateParams, opts ...option.RequestOption) (res *IPSECTunnelPSKGenerateResponse, err error) {
	var env IPSECTunnelPSKGenerateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipsecTunnelID == "" {
		err = errors.New("missing required ipsec_tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/ipsec_tunnels/%s/psk_generate", params.AccountID, ipsecTunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The PSK metadata that includes when the PSK was generated.
type PSKMetadata struct {
	// The date and time the tunnel was last modified.
	LastGeneratedOn time.Time       `json:"last_generated_on" format:"date-time"`
	JSON            pskMetadataJSON `json:"-"`
}

// pskMetadataJSON contains the JSON metadata for the struct [PSKMetadata]
type pskMetadataJSON struct {
	LastGeneratedOn apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PSKMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pskMetadataJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string                            `json:"description"`
	HealthCheck IPSECTunnelNewResponseHealthCheck `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection bool                       `json:"replay_protection"`
	JSON             ipsecTunnelNewResponseJSON `json:"-"`
}

// ipsecTunnelNewResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelNewResponse]
type ipsecTunnelNewResponseJSON struct {
	ID                 apijson.Field
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	HealthCheck        apijson.Field
	InterfaceAddress6  apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewResponseHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction IPSECTunnelNewResponseHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target IPSECTunnelNewResponseHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                       `json:"type"`
	JSON ipsecTunnelNewResponseHealthCheckJSON `json:"-"`
}

// ipsecTunnelNewResponseHealthCheckJSON contains the JSON metadata for the struct
// [IPSECTunnelNewResponseHealthCheck]
type ipsecTunnelNewResponseHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelNewResponseHealthCheckDirection string

const (
	IPSECTunnelNewResponseHealthCheckDirectionUnidirectional IPSECTunnelNewResponseHealthCheckDirection = "unidirectional"
	IPSECTunnelNewResponseHealthCheckDirectionBidirectional  IPSECTunnelNewResponseHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelNewResponseHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelNewResponseHealthCheckDirectionUnidirectional, IPSECTunnelNewResponseHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [IPSECTunnelNewResponseHealthCheckTargetMagicHealthCheckTarget] or
// [shared.UnionString].
type IPSECTunnelNewResponseHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelNewResponseHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPSECTunnelNewResponseHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPSECTunnelNewResponseHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelNewResponseHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                            `json:"saved"`
	JSON  ipsecTunnelNewResponseHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// ipsecTunnelNewResponseHealthCheckTargetMagicHealthCheckTargetJSON contains the
// JSON metadata for the struct
// [IPSECTunnelNewResponseHealthCheckTargetMagicHealthCheckTarget]
type ipsecTunnelNewResponseHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r IPSECTunnelNewResponseHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelNewResponseHealthCheckTargetUnion() {
}

type IPSECTunnelUpdateResponse struct {
	Modified            bool                                         `json:"modified"`
	ModifiedIPSECTunnel IPSECTunnelUpdateResponseModifiedIPSECTunnel `json:"modified_ipsec_tunnel"`
	JSON                ipsecTunnelUpdateResponseJSON                `json:"-"`
}

// ipsecTunnelUpdateResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelUpdateResponse]
type ipsecTunnelUpdateResponseJSON struct {
	Modified            apijson.Field
	ModifiedIPSECTunnel apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelUpdateResponseModifiedIPSECTunnel struct {
	// Identifier
	ID string `json:"id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string                                                  `json:"description"`
	HealthCheck IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheck `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection bool                                             `json:"replay_protection"`
	JSON             ipsecTunnelUpdateResponseModifiedIPSECTunnelJSON `json:"-"`
}

// ipsecTunnelUpdateResponseModifiedIPSECTunnelJSON contains the JSON metadata for
// the struct [IPSECTunnelUpdateResponseModifiedIPSECTunnel]
type ipsecTunnelUpdateResponseModifiedIPSECTunnelJSON struct {
	ID                 apijson.Field
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	HealthCheck        apijson.Field
	InterfaceAddress6  apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponseModifiedIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseModifiedIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                             `json:"type"`
	JSON ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckJSON `json:"-"`
}

// ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckJSON contains the JSON
// metadata for the struct
// [IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheck]
type ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirection string

const (
	IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirectionUnidirectional IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirection = "unidirectional"
	IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirectionBidirectional  IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirectionUnidirectional, IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                                  `json:"saved"`
	JSON  ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget]
type ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r IPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelUpdateResponseModifiedIPSECTunnelHealthCheckTargetUnion() {
}

type IPSECTunnelListResponse struct {
	IPSECTunnels []IPSECTunnelListResponseIPSECTunnel `json:"ipsec_tunnels"`
	JSON         ipsecTunnelListResponseJSON          `json:"-"`
}

// ipsecTunnelListResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelListResponse]
type ipsecTunnelListResponseJSON struct {
	IPSECTunnels apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPSECTunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelListResponseIPSECTunnel struct {
	// Identifier
	ID string `json:"id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string                                         `json:"description"`
	HealthCheck IPSECTunnelListResponseIPSECTunnelsHealthCheck `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection bool                                   `json:"replay_protection"`
	JSON             ipsecTunnelListResponseIPSECTunnelJSON `json:"-"`
}

// ipsecTunnelListResponseIPSECTunnelJSON contains the JSON metadata for the struct
// [IPSECTunnelListResponseIPSECTunnel]
type ipsecTunnelListResponseIPSECTunnelJSON struct {
	ID                 apijson.Field
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	HealthCheck        apijson.Field
	InterfaceAddress6  apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelListResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelListResponseIPSECTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction IPSECTunnelListResponseIPSECTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                    `json:"type"`
	JSON ipsecTunnelListResponseIPSECTunnelsHealthCheckJSON `json:"-"`
}

// ipsecTunnelListResponseIPSECTunnelsHealthCheckJSON contains the JSON metadata
// for the struct [IPSECTunnelListResponseIPSECTunnelsHealthCheck]
type ipsecTunnelListResponseIPSECTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelListResponseIPSECTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseIPSECTunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelListResponseIPSECTunnelsHealthCheckDirection string

const (
	IPSECTunnelListResponseIPSECTunnelsHealthCheckDirectionUnidirectional IPSECTunnelListResponseIPSECTunnelsHealthCheckDirection = "unidirectional"
	IPSECTunnelListResponseIPSECTunnelsHealthCheckDirectionBidirectional  IPSECTunnelListResponseIPSECTunnelsHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelListResponseIPSECTunnelsHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelListResponseIPSECTunnelsHealthCheckDirectionUnidirectional, IPSECTunnelListResponseIPSECTunnelsHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget] or
// [shared.UnionString].
type IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                         `json:"saved"`
	JSON  ipsecTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// ipsecTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget]
type ipsecTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r IPSECTunnelListResponseIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelListResponseIPSECTunnelsHealthCheckTargetUnion() {
}

type IPSECTunnelDeleteResponse struct {
	Deleted            bool                                        `json:"deleted"`
	DeletedIPSECTunnel IPSECTunnelDeleteResponseDeletedIPSECTunnel `json:"deleted_ipsec_tunnel"`
	JSON               ipsecTunnelDeleteResponseJSON               `json:"-"`
}

// ipsecTunnelDeleteResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelDeleteResponse]
type ipsecTunnelDeleteResponseJSON struct {
	Deleted            apijson.Field
	DeletedIPSECTunnel apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelDeleteResponseDeletedIPSECTunnel struct {
	// Identifier
	ID string `json:"id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string                                                 `json:"description"`
	HealthCheck IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheck `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection bool                                            `json:"replay_protection"`
	JSON             ipsecTunnelDeleteResponseDeletedIPSECTunnelJSON `json:"-"`
}

// ipsecTunnelDeleteResponseDeletedIPSECTunnelJSON contains the JSON metadata for
// the struct [IPSECTunnelDeleteResponseDeletedIPSECTunnel]
type ipsecTunnelDeleteResponseDeletedIPSECTunnelJSON struct {
	ID                 apijson.Field
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	HealthCheck        apijson.Field
	InterfaceAddress6  apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponseDeletedIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseDeletedIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                            `json:"type"`
	JSON ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckJSON `json:"-"`
}

// ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckJSON contains the JSON
// metadata for the struct [IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheck]
type ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirection string

const (
	IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirectionUnidirectional IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirection = "unidirectional"
	IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirectionBidirectional  IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirectionUnidirectional, IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                                 `json:"saved"`
	JSON  ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget]
type ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r IPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelDeleteResponseDeletedIPSECTunnelHealthCheckTargetUnion() {
}

type IPSECTunnelBulkUpdateResponse struct {
	Modified             bool                                               `json:"modified"`
	ModifiedIPSECTunnels []IPSECTunnelBulkUpdateResponseModifiedIPSECTunnel `json:"modified_ipsec_tunnels"`
	JSON                 ipsecTunnelBulkUpdateResponseJSON                  `json:"-"`
}

// ipsecTunnelBulkUpdateResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelBulkUpdateResponse]
type ipsecTunnelBulkUpdateResponseJSON struct {
	Modified             apijson.Field
	ModifiedIPSECTunnels apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *IPSECTunnelBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelBulkUpdateResponseModifiedIPSECTunnel struct {
	// Identifier
	ID string `json:"id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string                                                       `json:"description"`
	HealthCheck IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheck `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection bool                                                 `json:"replay_protection"`
	JSON             ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelJSON `json:"-"`
}

// ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelJSON contains the JSON metadata
// for the struct [IPSECTunnelBulkUpdateResponseModifiedIPSECTunnel]
type ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelJSON struct {
	ID                 apijson.Field
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	HealthCheck        apijson.Field
	InterfaceAddress6  apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelBulkUpdateResponseModifiedIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                                  `json:"type"`
	JSON ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckJSON `json:"-"`
}

// ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckJSON contains the
// JSON metadata for the struct
// [IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheck]
type ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirection string

const (
	IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirectionUnidirectional IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirection = "unidirectional"
	IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirectionBidirectional  IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirectionUnidirectional, IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget]
// or [shared.UnionString].
type IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                                       `json:"saved"`
	JSON  ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget]
type ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r IPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelBulkUpdateResponseModifiedIPSECTunnelsHealthCheckTargetUnion() {
}

type IPSECTunnelGetResponse struct {
	IPSECTunnel IPSECTunnelGetResponseIPSECTunnel `json:"ipsec_tunnel"`
	JSON        ipsecTunnelGetResponseJSON        `json:"-"`
}

// ipsecTunnelGetResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelGetResponse]
type ipsecTunnelGetResponseJSON struct {
	IPSECTunnel apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelGetResponseIPSECTunnel struct {
	// Identifier
	ID string `json:"id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint string `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress string `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name string `json:"name,required"`
	// When `true`, the tunnel can use a null-cipher (`ENCR_NULL`) in the ESP tunnel
	// (Phase 2).
	AllowNullCipher bool `json:"allow_null_cipher"`
	// The date and time the tunnel was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint string `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description string                                       `json:"description"`
	HealthCheck IPSECTunnelGetResponseIPSECTunnelHealthCheck `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 string `json:"interface_address6"`
	// The date and time the tunnel was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata `json:"psk_metadata"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection bool                                  `json:"replay_protection"`
	JSON             ipsecTunnelGetResponseIPSECTunnelJSON `json:"-"`
}

// ipsecTunnelGetResponseIPSECTunnelJSON contains the JSON metadata for the struct
// [IPSECTunnelGetResponseIPSECTunnel]
type ipsecTunnelGetResponseIPSECTunnelJSON struct {
	ID                 apijson.Field
	CloudflareEndpoint apijson.Field
	InterfaceAddress   apijson.Field
	Name               apijson.Field
	AllowNullCipher    apijson.Field
	CreatedOn          apijson.Field
	CustomerEndpoint   apijson.Field
	Description        apijson.Field
	HealthCheck        apijson.Field
	InterfaceAddress6  apijson.Field
	ModifiedOn         apijson.Field
	PSKMetadata        apijson.Field
	ReplayProtection   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *IPSECTunnelGetResponseIPSECTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseIPSECTunnelJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelGetResponseIPSECTunnelHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction IPSECTunnelGetResponseIPSECTunnelHealthCheckDirection `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled bool `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate HealthCheckRate `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type HealthCheckType                                  `json:"type"`
	JSON ipsecTunnelGetResponseIPSECTunnelHealthCheckJSON `json:"-"`
}

// ipsecTunnelGetResponseIPSECTunnelHealthCheckJSON contains the JSON metadata for
// the struct [IPSECTunnelGetResponseIPSECTunnelHealthCheck]
type ipsecTunnelGetResponseIPSECTunnelHealthCheckJSON struct {
	Direction   apijson.Field
	Enabled     apijson.Field
	Rate        apijson.Field
	Target      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponseIPSECTunnelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseIPSECTunnelHealthCheckJSON) RawJSON() string {
	return r.raw
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelGetResponseIPSECTunnelHealthCheckDirection string

const (
	IPSECTunnelGetResponseIPSECTunnelHealthCheckDirectionUnidirectional IPSECTunnelGetResponseIPSECTunnelHealthCheckDirection = "unidirectional"
	IPSECTunnelGetResponseIPSECTunnelHealthCheckDirectionBidirectional  IPSECTunnelGetResponseIPSECTunnelHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelGetResponseIPSECTunnelHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelGetResponseIPSECTunnelHealthCheckDirectionUnidirectional, IPSECTunnelGetResponseIPSECTunnelHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Union satisfied by
// [IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTarget] or
// [shared.UnionString].
type IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTarget{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTarget struct {
	// The effective health check target. If 'saved' is empty, then this field will be
	// populated with the calculated default value on GET requests. Ignored in POST,
	// PUT, and PATCH requests.
	Effective string `json:"effective"`
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved string                                                                       `json:"saved"`
	JSON  ipsecTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON `json:"-"`
}

// ipsecTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON
// contains the JSON metadata for the struct
// [IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTarget]
type ipsecTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON struct {
	Effective   apijson.Field
	Saved       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTargetJSON) RawJSON() string {
	return r.raw
}

func (r IPSECTunnelGetResponseIPSECTunnelHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelGetResponseIPSECTunnelHealthCheckTargetUnion() {
}

type IPSECTunnelPSKGenerateResponse struct {
	// Identifier
	IPSECTunnelID string `json:"ipsec_tunnel_id"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK string `json:"psk"`
	// The PSK metadata that includes when the PSK was generated.
	PSKMetadata PSKMetadata                        `json:"psk_metadata"`
	JSON        ipsecTunnelPSKGenerateResponseJSON `json:"-"`
}

// ipsecTunnelPSKGenerateResponseJSON contains the JSON metadata for the struct
// [IPSECTunnelPSKGenerateResponse]
type ipsecTunnelPSKGenerateResponseJSON struct {
	IPSECTunnelID apijson.Field
	PSK           apijson.Field
	PSKMetadata   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IPSECTunnelPSKGenerateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelPSKGenerateResponseJSON) RawJSON() string {
	return r.raw
}

type IPSECTunnelNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name param.Field[string] `json:"name,required"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint param.Field[string] `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description param.Field[string]                          `json:"description"`
	HealthCheck param.Field[IPSECTunnelNewParamsHealthCheck] `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 param.Field[string] `json:"interface_address6"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  param.Field[bool] `json:"replay_protection"`
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

func (r IPSECTunnelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPSECTunnelNewParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction param.Field[IPSECTunnelNewParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[HealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[IPSECTunnelNewParamsHealthCheckTargetUnion] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[HealthCheckType] `json:"type"`
}

func (r IPSECTunnelNewParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelNewParamsHealthCheckDirection string

const (
	IPSECTunnelNewParamsHealthCheckDirectionUnidirectional IPSECTunnelNewParamsHealthCheckDirection = "unidirectional"
	IPSECTunnelNewParamsHealthCheckDirectionBidirectional  IPSECTunnelNewParamsHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelNewParamsHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelNewParamsHealthCheckDirectionUnidirectional, IPSECTunnelNewParamsHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Satisfied by
// [magic_transit.IPSECTunnelNewParamsHealthCheckTargetMagicHealthCheckTarget],
// [shared.UnionString].
type IPSECTunnelNewParamsHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelNewParamsHealthCheckTargetUnion()
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelNewParamsHealthCheckTargetMagicHealthCheckTarget struct {
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved param.Field[string] `json:"saved"`
}

func (r IPSECTunnelNewParamsHealthCheckTargetMagicHealthCheckTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPSECTunnelNewParamsHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelNewParamsHealthCheckTargetUnion() {
}

type IPSECTunnelNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   IPSECTunnelNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelNewResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelNewResponseEnvelope]
type ipsecTunnelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelNewResponseEnvelopeSuccess bool

const (
	IPSECTunnelNewResponseEnvelopeSuccessTrue IPSECTunnelNewResponseEnvelopeSuccess = true
)

func (r IPSECTunnelNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPSECTunnelUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The IP address assigned to the Cloudflare side of the IPsec tunnel.
	CloudflareEndpoint param.Field[string] `json:"cloudflare_endpoint,required"`
	// A 31-bit prefix (/31 in CIDR notation) supporting two hosts, one for each side
	// of the tunnel. Select the subnet from the following private IP space:
	// 10.0.0.0–10.255.255.255, 172.16.0.0–172.31.255.255, 192.168.0.0–192.168.255.255.
	InterfaceAddress param.Field[string] `json:"interface_address,required"`
	// The name of the IPsec tunnel. The name cannot share a name with other tunnels.
	Name param.Field[string] `json:"name,required"`
	// The IP address assigned to the customer side of the IPsec tunnel. Not required,
	// but must be set for proactive traceroutes to work.
	CustomerEndpoint param.Field[string] `json:"customer_endpoint"`
	// An optional description forthe IPsec tunnel.
	Description param.Field[string]                             `json:"description"`
	HealthCheck param.Field[IPSECTunnelUpdateParamsHealthCheck] `json:"health_check"`
	// A 127 bit IPV6 prefix from within the virtual_subnet6 prefix space with the
	// address being the first IP of the subnet and not same as the address of
	// virtual_subnet6. Eg if virtual_subnet6 is 2606:54c1:7:0:a9fe:12d2::/127 ,
	// interface_address6 could be 2606:54c1:7:0:a9fe:12d2:1:200/127
	InterfaceAddress6 param.Field[string] `json:"interface_address6"`
	// A randomly generated or provided string for use in the IPsec tunnel.
	PSK param.Field[string] `json:"psk"`
	// If `true`, then IPsec replay protection will be supported in the
	// Cloudflare-to-customer direction.
	ReplayProtection  param.Field[bool] `json:"replay_protection"`
	XMagicNewHcTarget param.Field[bool] `header:"x-magic-new-hc-target"`
}

func (r IPSECTunnelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type IPSECTunnelUpdateParamsHealthCheck struct {
	// The direction of the flow of the healthcheck. Either unidirectional, where the
	// probe comes to you via the tunnel and the result comes back to Cloudflare via
	// the open Internet, or bidirectional where both the probe and result come and go
	// via the tunnel.
	Direction param.Field[IPSECTunnelUpdateParamsHealthCheckDirection] `json:"direction"`
	// Determines whether to run healthchecks for a tunnel.
	Enabled param.Field[bool] `json:"enabled"`
	// How frequent the health check is run. The default value is `mid`.
	Rate param.Field[HealthCheckRate] `json:"rate"`
	// The destination address in a request type health check. After the healthcheck is
	// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
	// to this address. This field defaults to `customer_gre_endpoint address`. This
	// field is ignored for bidirectional healthchecks as the interface_address (not
	// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
	// object form if the x-magic-new-hc-target header is set to true and string form
	// if x-magic-new-hc-target is absent or set to false.
	Target param.Field[IPSECTunnelUpdateParamsHealthCheckTargetUnion] `json:"target"`
	// The type of healthcheck to run, reply or request. The default value is `reply`.
	Type param.Field[HealthCheckType] `json:"type"`
}

func (r IPSECTunnelUpdateParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The direction of the flow of the healthcheck. Either unidirectional, where the
// probe comes to you via the tunnel and the result comes back to Cloudflare via
// the open Internet, or bidirectional where both the probe and result come and go
// via the tunnel.
type IPSECTunnelUpdateParamsHealthCheckDirection string

const (
	IPSECTunnelUpdateParamsHealthCheckDirectionUnidirectional IPSECTunnelUpdateParamsHealthCheckDirection = "unidirectional"
	IPSECTunnelUpdateParamsHealthCheckDirectionBidirectional  IPSECTunnelUpdateParamsHealthCheckDirection = "bidirectional"
)

func (r IPSECTunnelUpdateParamsHealthCheckDirection) IsKnown() bool {
	switch r {
	case IPSECTunnelUpdateParamsHealthCheckDirectionUnidirectional, IPSECTunnelUpdateParamsHealthCheckDirectionBidirectional:
		return true
	}
	return false
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target. Must be in
// object form if the x-magic-new-hc-target header is set to true and string form
// if x-magic-new-hc-target is absent or set to false.
//
// Satisfied by
// [magic_transit.IPSECTunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget],
// [shared.UnionString].
type IPSECTunnelUpdateParamsHealthCheckTargetUnion interface {
	ImplementsIPSECTunnelUpdateParamsHealthCheckTargetUnion()
}

// The destination address in a request type health check. After the healthcheck is
// decapsulated at the customer end of the tunnel, the ICMP echo will be forwarded
// to this address. This field defaults to `customer_gre_endpoint address`. This
// field is ignored for bidirectional healthchecks as the interface_address (not
// assigned to the Cloudflare side of the tunnel) is used as the target.
type IPSECTunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget struct {
	// The saved health check target. Setting the value to the empty string indicates
	// that the calculated default value will be used.
	Saved param.Field[string] `json:"saved"`
}

func (r IPSECTunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPSECTunnelUpdateParamsHealthCheckTargetMagicHealthCheckTarget) ImplementsIPSECTunnelUpdateParamsHealthCheckTargetUnion() {
}

type IPSECTunnelUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   IPSECTunnelUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelUpdateResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelUpdateResponseEnvelope]
type ipsecTunnelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelUpdateResponseEnvelopeSuccess bool

const (
	IPSECTunnelUpdateResponseEnvelopeSuccessTrue IPSECTunnelUpdateResponseEnvelopeSuccess = true
)

func (r IPSECTunnelUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPSECTunnelListParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type IPSECTunnelListResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   IPSECTunnelListResponse `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelListResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelListResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelListResponseEnvelope]
type ipsecTunnelListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelListResponseEnvelopeSuccess bool

const (
	IPSECTunnelListResponseEnvelopeSuccessTrue IPSECTunnelListResponseEnvelopeSuccess = true
)

func (r IPSECTunnelListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPSECTunnelDeleteParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type IPSECTunnelDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   IPSECTunnelDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelDeleteResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelDeleteResponseEnvelope]
type ipsecTunnelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelDeleteResponseEnvelopeSuccess bool

const (
	IPSECTunnelDeleteResponseEnvelopeSuccessTrue IPSECTunnelDeleteResponseEnvelopeSuccess = true
)

func (r IPSECTunnelDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPSECTunnelBulkUpdateParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	Body              interface{}         `json:"body,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

func (r IPSECTunnelBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type IPSECTunnelBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   IPSECTunnelBulkUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelBulkUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelBulkUpdateResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelBulkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IPSECTunnelBulkUpdateResponseEnvelope]
type ipsecTunnelBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelBulkUpdateResponseEnvelopeSuccess bool

const (
	IPSECTunnelBulkUpdateResponseEnvelopeSuccessTrue IPSECTunnelBulkUpdateResponseEnvelopeSuccess = true
)

func (r IPSECTunnelBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPSECTunnelGetParams struct {
	// Identifier
	AccountID         param.Field[string] `path:"account_id,required"`
	XMagicNewHcTarget param.Field[bool]   `header:"x-magic-new-hc-target"`
}

type IPSECTunnelGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   IPSECTunnelGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelGetResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPSECTunnelGetResponseEnvelope]
type ipsecTunnelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelGetResponseEnvelopeSuccess bool

const (
	IPSECTunnelGetResponseEnvelopeSuccessTrue IPSECTunnelGetResponseEnvelopeSuccess = true
)

func (r IPSECTunnelGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type IPSECTunnelPSKGenerateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r IPSECTunnelPSKGenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type IPSECTunnelPSKGenerateResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   IPSECTunnelPSKGenerateResponse `json:"result,required"`
	// Whether the API call was successful
	Success IPSECTunnelPSKGenerateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ipsecTunnelPSKGenerateResponseEnvelopeJSON    `json:"-"`
}

// ipsecTunnelPSKGenerateResponseEnvelopeJSON contains the JSON metadata for the
// struct [IPSECTunnelPSKGenerateResponseEnvelope]
type ipsecTunnelPSKGenerateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPSECTunnelPSKGenerateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsecTunnelPSKGenerateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IPSECTunnelPSKGenerateResponseEnvelopeSuccess bool

const (
	IPSECTunnelPSKGenerateResponseEnvelopeSuccessTrue IPSECTunnelPSKGenerateResponseEnvelopeSuccess = true
)

func (r IPSECTunnelPSKGenerateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPSECTunnelPSKGenerateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
