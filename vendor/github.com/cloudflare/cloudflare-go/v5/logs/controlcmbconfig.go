// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ControlCmbConfigService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewControlCmbConfigService] method instead.
type ControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewControlCmbConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewControlCmbConfigService(opts ...option.RequestOption) (r *ControlCmbConfigService) {
	r = &ControlCmbConfigService{}
	r.Options = opts
	return
}

// Updates CMB config.
func (r *ControlCmbConfigService) New(ctx context.Context, params ControlCmbConfigNewParams, opts ...option.RequestOption) (res *CmbConfig, err error) {
	var env ControlCmbConfigNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes CMB config.
func (r *ControlCmbConfigService) Delete(ctx context.Context, body ControlCmbConfigDeleteParams, opts ...option.RequestOption) (res *ControlCmbConfigDeleteResponse, err error) {
	var env ControlCmbConfigDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets CMB config.
func (r *ControlCmbConfigService) Get(ctx context.Context, query ControlCmbConfigGetParams, opts ...option.RequestOption) (res *CmbConfig, err error) {
	var env ControlCmbConfigGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CmbConfig struct {
	// Allow out of region access
	AllowOutOfRegionAccess bool `json:"allow_out_of_region_access"`
	// Name of the region.
	Regions string        `json:"regions"`
	JSON    cmbConfigJSON `json:"-"`
}

// cmbConfigJSON contains the JSON metadata for the struct [CmbConfig]
type cmbConfigJSON struct {
	AllowOutOfRegionAccess apijson.Field
	Regions                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CmbConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cmbConfigJSON) RawJSON() string {
	return r.raw
}

type CmbConfigParam struct {
	// Allow out of region access
	AllowOutOfRegionAccess param.Field[bool] `json:"allow_out_of_region_access"`
	// Name of the region.
	Regions param.Field[string] `json:"regions"`
}

func (r CmbConfigParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ControlCmbConfigDeleteResponse = interface{}

type ControlCmbConfigNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	CmbConfig CmbConfigParam      `json:"cmb_config,required"`
}

func (r ControlCmbConfigNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.CmbConfig)
}

type ControlCmbConfigNewResponseEnvelope struct {
	Errors   []ControlCmbConfigNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ControlCmbConfigNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ControlCmbConfigNewResponseEnvelopeSuccess `json:"success,required"`
	Result  CmbConfig                                  `json:"result,nullable"`
	JSON    controlCmbConfigNewResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigNewResponseEnvelope]
type controlCmbConfigNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ControlCmbConfigNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             controlCmbConfigNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ControlCmbConfigNewResponseEnvelopeErrors]
type controlCmbConfigNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    controlCmbConfigNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ControlCmbConfigNewResponseEnvelopeErrorsSource]
type controlCmbConfigNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ControlCmbConfigNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             controlCmbConfigNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ControlCmbConfigNewResponseEnvelopeMessages]
type controlCmbConfigNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    controlCmbConfigNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// controlCmbConfigNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ControlCmbConfigNewResponseEnvelopeMessagesSource]
type controlCmbConfigNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ControlCmbConfigNewResponseEnvelopeSuccess bool

const (
	ControlCmbConfigNewResponseEnvelopeSuccessTrue ControlCmbConfigNewResponseEnvelopeSuccess = true
)

func (r ControlCmbConfigNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlCmbConfigNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ControlCmbConfigDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ControlCmbConfigDeleteResponseEnvelope struct {
	Errors   []ControlCmbConfigDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ControlCmbConfigDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ControlCmbConfigDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ControlCmbConfigDeleteResponse                `json:"result,nullable"`
	JSON    controlCmbConfigDeleteResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigDeleteResponseEnvelope]
type controlCmbConfigDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigDeleteResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           ControlCmbConfigDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             controlCmbConfigDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ControlCmbConfigDeleteResponseEnvelopeErrors]
type controlCmbConfigDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    controlCmbConfigDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ControlCmbConfigDeleteResponseEnvelopeErrorsSource]
type controlCmbConfigDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigDeleteResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           ControlCmbConfigDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             controlCmbConfigDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ControlCmbConfigDeleteResponseEnvelopeMessages]
type controlCmbConfigDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    controlCmbConfigDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// controlCmbConfigDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ControlCmbConfigDeleteResponseEnvelopeMessagesSource]
type controlCmbConfigDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ControlCmbConfigDeleteResponseEnvelopeSuccess bool

const (
	ControlCmbConfigDeleteResponseEnvelopeSuccessTrue ControlCmbConfigDeleteResponseEnvelopeSuccess = true
)

func (r ControlCmbConfigDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlCmbConfigDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ControlCmbConfigGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ControlCmbConfigGetResponseEnvelope struct {
	Errors   []ControlCmbConfigGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ControlCmbConfigGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ControlCmbConfigGetResponseEnvelopeSuccess `json:"success,required"`
	Result  CmbConfig                                  `json:"result,nullable"`
	JSON    controlCmbConfigGetResponseEnvelopeJSON    `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ControlCmbConfigGetResponseEnvelope]
type controlCmbConfigGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigGetResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ControlCmbConfigGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             controlCmbConfigGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ControlCmbConfigGetResponseEnvelopeErrors]
type controlCmbConfigGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigGetResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    controlCmbConfigGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ControlCmbConfigGetResponseEnvelopeErrorsSource]
type controlCmbConfigGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigGetResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ControlCmbConfigGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             controlCmbConfigGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ControlCmbConfigGetResponseEnvelopeMessages]
type controlCmbConfigGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ControlCmbConfigGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    controlCmbConfigGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// controlCmbConfigGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ControlCmbConfigGetResponseEnvelopeMessagesSource]
type controlCmbConfigGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ControlCmbConfigGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r controlCmbConfigGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ControlCmbConfigGetResponseEnvelopeSuccess bool

const (
	ControlCmbConfigGetResponseEnvelopeSuccessTrue ControlCmbConfigGetResponseEnvelopeSuccess = true
)

func (r ControlCmbConfigGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ControlCmbConfigGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
