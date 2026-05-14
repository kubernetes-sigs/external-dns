// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3

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

// HostnameService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameService] method instead.
type HostnameService struct {
	Options            []option.RequestOption
	IPFSUniversalPaths *HostnameIPFSUniversalPathService
}

// NewHostnameService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHostnameService(opts ...option.RequestOption) (r *HostnameService) {
	r = &HostnameService{}
	r.Options = opts
	r.IPFSUniversalPaths = NewHostnameIPFSUniversalPathService(opts...)
	return
}

// Create Web3 Hostname
func (r *HostnameService) New(ctx context.Context, params HostnameNewParams, opts ...option.RequestOption) (res *Hostname, err error) {
	var env HostnameNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Web3 Hostnames
func (r *HostnameService) List(ctx context.Context, query HostnameListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Hostname], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames", query.ZoneID)
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

// List Web3 Hostnames
func (r *HostnameService) ListAutoPaging(ctx context.Context, query HostnameListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Hostname] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Web3 Hostname
func (r *HostnameService) Delete(ctx context.Context, identifier string, body HostnameDeleteParams, opts ...option.RequestOption) (res *HostnameDeleteResponse, err error) {
	var env HostnameDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", body.ZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit Web3 Hostname
func (r *HostnameService) Edit(ctx context.Context, identifier string, params HostnameEditParams, opts ...option.RequestOption) (res *Hostname, err error) {
	var env HostnameEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", params.ZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Web3 Hostname Details
func (r *HostnameService) Get(ctx context.Context, identifier string, query HostnameGetParams, opts ...option.RequestOption) (res *Hostname, err error) {
	var env HostnameGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s", query.ZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Hostname struct {
	// Specify the identifier of the hostname.
	ID        string    `json:"id"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specify an optional description of the hostname.
	Description string `json:"description"`
	// Specify the DNSLink value used if the target is ipfs.
	Dnslink    string    `json:"dnslink"`
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Specify the hostname that points to the target gateway via CNAME.
	Name string `json:"name"`
	// Specifies the status of the hostname's activation.
	Status HostnameStatus `json:"status"`
	// Specify the target gateway of the hostname.
	Target HostnameTarget `json:"target"`
	JSON   hostnameJSON   `json:"-"`
}

// hostnameJSON contains the JSON metadata for the struct [Hostname]
type hostnameJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	Dnslink     apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Status      apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Hostname) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameJSON) RawJSON() string {
	return r.raw
}

// Specifies the status of the hostname's activation.
type HostnameStatus string

const (
	HostnameStatusActive   HostnameStatus = "active"
	HostnameStatusPending  HostnameStatus = "pending"
	HostnameStatusDeleting HostnameStatus = "deleting"
	HostnameStatusError    HostnameStatus = "error"
)

func (r HostnameStatus) IsKnown() bool {
	switch r {
	case HostnameStatusActive, HostnameStatusPending, HostnameStatusDeleting, HostnameStatusError:
		return true
	}
	return false
}

// Specify the target gateway of the hostname.
type HostnameTarget string

const (
	HostnameTargetEthereum          HostnameTarget = "ethereum"
	HostnameTargetIPFS              HostnameTarget = "ipfs"
	HostnameTargetIPFSUniversalPath HostnameTarget = "ipfs_universal_path"
)

func (r HostnameTarget) IsKnown() bool {
	switch r {
	case HostnameTargetEthereum, HostnameTargetIPFS, HostnameTargetIPFSUniversalPath:
		return true
	}
	return false
}

type HostnameDeleteResponse struct {
	// Specify the identifier of the hostname.
	ID   string                     `json:"id,required"`
	JSON hostnameDeleteResponseJSON `json:"-"`
}

// hostnameDeleteResponseJSON contains the JSON metadata for the struct
// [HostnameDeleteResponse]
type hostnameDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type HostnameNewParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Specify the hostname that points to the target gateway via CNAME.
	Name param.Field[string] `json:"name,required"`
	// Specify the target gateway of the hostname.
	Target param.Field[HostnameNewParamsTarget] `json:"target,required"`
	// Specify an optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// Specify the DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r HostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the target gateway of the hostname.
type HostnameNewParamsTarget string

const (
	HostnameNewParamsTargetEthereum          HostnameNewParamsTarget = "ethereum"
	HostnameNewParamsTargetIPFS              HostnameNewParamsTarget = "ipfs"
	HostnameNewParamsTargetIPFSUniversalPath HostnameNewParamsTarget = "ipfs_universal_path"
)

func (r HostnameNewParamsTarget) IsKnown() bool {
	switch r {
	case HostnameNewParamsTargetEthereum, HostnameNewParamsTargetIPFS, HostnameNewParamsTargetIPFSUniversalPath:
		return true
	}
	return false
}

type HostnameNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Hostname              `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameNewResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                     `json:"result_info"`
	JSON       hostnameNewResponseEnvelopeJSON `json:"-"`
}

// hostnameNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameNewResponseEnvelope]
type hostnameNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameNewResponseEnvelopeSuccess bool

const (
	HostnameNewResponseEnvelopeSuccessTrue HostnameNewResponseEnvelopeSuccess = true
)

func (r HostnameNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameListParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameDeleteParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   HostnameDeleteResponse `json:"result,required,nullable"`
	// Specifies whether the API call was successful.
	Success HostnameDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameDeleteResponseEnvelopeJSON    `json:"-"`
}

// hostnameDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameDeleteResponseEnvelope]
type hostnameDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameDeleteResponseEnvelopeSuccess bool

const (
	HostnameDeleteResponseEnvelopeSuccessTrue HostnameDeleteResponseEnvelopeSuccess = true
)

func (r HostnameDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameEditParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Specify an optional description of the hostname.
	Description param.Field[string] `json:"description"`
	// Specify the DNSLink value used if the target is ipfs.
	Dnslink param.Field[string] `json:"dnslink"`
}

func (r HostnameEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Hostname              `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameEditResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                      `json:"result_info"`
	JSON       hostnameEditResponseEnvelopeJSON `json:"-"`
}

// hostnameEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameEditResponseEnvelope]
type hostnameEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameEditResponseEnvelopeSuccess bool

const (
	HostnameEditResponseEnvelopeSuccessTrue HostnameEditResponseEnvelopeSuccess = true
)

func (r HostnameEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameGetParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Hostname              `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameGetResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                     `json:"result_info"`
	JSON       hostnameGetResponseEnvelopeJSON `json:"-"`
}

// hostnameGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameGetResponseEnvelope]
type hostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameGetResponseEnvelopeSuccess bool

const (
	HostnameGetResponseEnvelopeSuccessTrue HostnameGetResponseEnvelopeSuccess = true
)

func (r HostnameGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
