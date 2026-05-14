// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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
)

// ScanConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanConfigService] method instead.
type ScanConfigService struct {
	Options []option.RequestOption
}

// NewScanConfigService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScanConfigService(opts ...option.RequestOption) (r *ScanConfigService) {
	r = &ScanConfigService{}
	r.Options = opts
	return
}

// Create a new Scan Config
func (r *ScanConfigService) New(ctx context.Context, params ScanConfigNewParams, opts ...option.RequestOption) (res *ScanConfigNewResponse, err error) {
	var env ScanConfigNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Scan Configs
func (r *ScanConfigService) List(ctx context.Context, query ScanConfigListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ScanConfigListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config", query.AccountID)
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

// List Scan Configs
func (r *ScanConfigService) ListAutoPaging(ctx context.Context, query ScanConfigListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ScanConfigListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Scan Config
func (r *ScanConfigService) Delete(ctx context.Context, configID string, body ScanConfigDeleteParams, opts ...option.RequestOption) (res *ScanConfigDeleteResponse, err error) {
	var env ScanConfigDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config/%s", body.AccountID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update an existing Scan Config
func (r *ScanConfigService) Edit(ctx context.Context, configID string, params ScanConfigEditParams, opts ...option.RequestOption) (res *ScanConfigEditResponse, err error) {
	var env ScanConfigEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if configID == "" {
		err = errors.New("missing required config_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/scans/config/%s", params.AccountID, configID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ScanConfigNewResponse struct {
	// Defines the Config ID.
	ID        string `json:"id,required"`
	AccountID string `json:"account_id,required"`
	// Defines the number of days between each scan (0 = One-off scan).
	Frequency float64 `json:"frequency,required"`
	// Defines a list of IP addresses or CIDR blocks to scan. The maximum number of
	// total IP addresses allowed is 5000.
	IPs []string `json:"ips,required"`
	// Defines a list of ports to scan. Valid values are:"default", "all", or a
	// comma-separated list of ports or range of ports (e.g. ["1-80", "443"]).
	// "default" scans the 100 most commonly open ports.
	Ports []string                  `json:"ports,required"`
	JSON  scanConfigNewResponseJSON `json:"-"`
}

// scanConfigNewResponseJSON contains the JSON metadata for the struct
// [ScanConfigNewResponse]
type scanConfigNewResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	Ports       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseJSON) RawJSON() string {
	return r.raw
}

type ScanConfigListResponse struct {
	// Defines the Config ID.
	ID        string `json:"id,required"`
	AccountID string `json:"account_id,required"`
	// Defines the number of days between each scan (0 = One-off scan).
	Frequency float64 `json:"frequency,required"`
	// Defines a list of IP addresses or CIDR blocks to scan. The maximum number of
	// total IP addresses allowed is 5000.
	IPs []string `json:"ips,required"`
	// Defines a list of ports to scan. Valid values are:"default", "all", or a
	// comma-separated list of ports or range of ports (e.g. ["1-80", "443"]).
	// "default" scans the 100 most commonly open ports.
	Ports []string                   `json:"ports,required"`
	JSON  scanConfigListResponseJSON `json:"-"`
}

// scanConfigListResponseJSON contains the JSON metadata for the struct
// [ScanConfigListResponse]
type scanConfigListResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	Ports       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigListResponseJSON) RawJSON() string {
	return r.raw
}

type ScanConfigDeleteResponse = interface{}

type ScanConfigEditResponse struct {
	// Defines the Config ID.
	ID        string `json:"id,required"`
	AccountID string `json:"account_id,required"`
	// Defines the number of days between each scan (0 = One-off scan).
	Frequency float64 `json:"frequency,required"`
	// Defines a list of IP addresses or CIDR blocks to scan. The maximum number of
	// total IP addresses allowed is 5000.
	IPs []string `json:"ips,required"`
	// Defines a list of ports to scan. Valid values are:"default", "all", or a
	// comma-separated list of ports or range of ports (e.g. ["1-80", "443"]).
	// "default" scans the 100 most commonly open ports.
	Ports []string                   `json:"ports,required"`
	JSON  scanConfigEditResponseJSON `json:"-"`
}

// scanConfigEditResponseJSON contains the JSON metadata for the struct
// [ScanConfigEditResponse]
type scanConfigEditResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Frequency   apijson.Field
	IPs         apijson.Field
	Ports       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigEditResponseJSON) RawJSON() string {
	return r.raw
}

type ScanConfigNewParams struct {
	// Defines the Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Defines a list of IP addresses or CIDR blocks to scan. The maximum number of
	// total IP addresses allowed is 5000.
	IPs param.Field[[]string] `json:"ips,required"`
	// Defines the number of days between each scan (0 = One-off scan).
	Frequency param.Field[float64] `json:"frequency"`
	// Defines a list of ports to scan. Valid values are:"default", "all", or a
	// comma-separated list of ports or range of ports (e.g. ["1-80", "443"]).
	// "default" scans the 100 most commonly open ports.
	Ports param.Field[[]string] `json:"ports"`
}

func (r ScanConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScanConfigNewResponseEnvelope struct {
	Errors   []ScanConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScanConfigNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScanConfigNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ScanConfigNewResponse                `json:"result"`
	JSON    scanConfigNewResponseEnvelopeJSON    `json:"-"`
}

// scanConfigNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanConfigNewResponseEnvelope]
type scanConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanConfigNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ScanConfigNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             scanConfigNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// scanConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScanConfigNewResponseEnvelopeErrors]
type scanConfigNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScanConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScanConfigNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    scanConfigNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scanConfigNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ScanConfigNewResponseEnvelopeErrorsSource]
type scanConfigNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScanConfigNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ScanConfigNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             scanConfigNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// scanConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScanConfigNewResponseEnvelopeMessages]
type scanConfigNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScanConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScanConfigNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    scanConfigNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scanConfigNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScanConfigNewResponseEnvelopeMessagesSource]
type scanConfigNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScanConfigNewResponseEnvelopeSuccess bool

const (
	ScanConfigNewResponseEnvelopeSuccessTrue ScanConfigNewResponseEnvelopeSuccess = true
)

func (r ScanConfigNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScanConfigNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ScanConfigListParams struct {
	// Defines the Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanConfigDeleteParams struct {
	// Defines the Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ScanConfigDeleteResponseEnvelope struct {
	Errors   []string                             `json:"errors,required"`
	Messages []string                             `json:"messages,required"`
	Result   ScanConfigDeleteResponse             `json:"result,required"`
	Success  bool                                 `json:"success,required"`
	JSON     scanConfigDeleteResponseEnvelopeJSON `json:"-"`
}

// scanConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanConfigDeleteResponseEnvelope]
type scanConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanConfigEditParams struct {
	// Defines the Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Defines the number of days between each scan (0 = One-off scan).
	Frequency param.Field[float64] `json:"frequency"`
	// Defines a list of IP addresses or CIDR blocks to scan. The maximum number of
	// total IP addresses allowed is 5000.
	IPs param.Field[[]string] `json:"ips"`
	// Defines a list of ports to scan. Valid values are:"default", "all", or a
	// comma-separated list of ports or range of ports (e.g. ["1-80", "443"]).
	// "default" scans the 100 most commonly open ports.
	Ports param.Field[[]string] `json:"ports"`
}

func (r ScanConfigEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScanConfigEditResponseEnvelope struct {
	Errors   []ScanConfigEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ScanConfigEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ScanConfigEditResponseEnvelopeSuccess `json:"success,required"`
	Result  ScanConfigEditResponse                `json:"result"`
	JSON    scanConfigEditResponseEnvelopeJSON    `json:"-"`
}

// scanConfigEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ScanConfigEditResponseEnvelope]
type scanConfigEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ScanConfigEditResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ScanConfigEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             scanConfigEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// scanConfigEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ScanConfigEditResponseEnvelopeErrors]
type scanConfigEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScanConfigEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ScanConfigEditResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    scanConfigEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// scanConfigEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ScanConfigEditResponseEnvelopeErrorsSource]
type scanConfigEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ScanConfigEditResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ScanConfigEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             scanConfigEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// scanConfigEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ScanConfigEditResponseEnvelopeMessages]
type scanConfigEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ScanConfigEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ScanConfigEditResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    scanConfigEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// scanConfigEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ScanConfigEditResponseEnvelopeMessagesSource]
type scanConfigEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScanConfigEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scanConfigEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ScanConfigEditResponseEnvelopeSuccess bool

const (
	ScanConfigEditResponseEnvelopeSuccessTrue ScanConfigEditResponseEnvelopeSuccess = true
)

func (r ScanConfigEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ScanConfigEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
