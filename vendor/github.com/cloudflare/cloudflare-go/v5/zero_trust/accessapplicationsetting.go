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
)

// AccessApplicationSettingService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationSettingService] method instead.
type AccessApplicationSettingService struct {
	Options []option.RequestOption
}

// NewAccessApplicationSettingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessApplicationSettingService(opts ...option.RequestOption) (r *AccessApplicationSettingService) {
	r = &AccessApplicationSettingService{}
	r.Options = opts
	return
}

// Updates Access application settings.
func (r *AccessApplicationSettingService) Update(ctx context.Context, appID AppIDParam, params AccessApplicationSettingUpdateParams, opts ...option.RequestOption) (res *AccessApplicationSettingUpdateResponse, err error) {
	var env AccessApplicationSettingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/settings", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates Access application settings.
func (r *AccessApplicationSettingService) Edit(ctx context.Context, appID AppIDParam, params AccessApplicationSettingEditParams, opts ...option.RequestOption) (res *AccessApplicationSettingEditResponse, err error) {
	var env AccessApplicationSettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/settings", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessApplicationSettingUpdateResponse struct {
	// Enables loading application content in an iFrame.
	AllowIframe bool `json:"allow_iframe"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool                                       `json:"skip_interstitial"`
	JSON             accessApplicationSettingUpdateResponseJSON `json:"-"`
}

// accessApplicationSettingUpdateResponseJSON contains the JSON metadata for the
// struct [AccessApplicationSettingUpdateResponse]
type accessApplicationSettingUpdateResponseJSON struct {
	AllowIframe      apijson.Field
	SkipInterstitial apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationSettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingEditResponse struct {
	// Enables loading application content in an iFrame.
	AllowIframe bool `json:"allow_iframe"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial bool                                     `json:"skip_interstitial"`
	JSON             accessApplicationSettingEditResponseJSON `json:"-"`
}

// accessApplicationSettingEditResponseJSON contains the JSON metadata for the
// struct [AccessApplicationSettingEditResponse]
type accessApplicationSettingEditResponseJSON struct {
	AllowIframe      apijson.Field
	SkipInterstitial apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationSettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingEditResponseJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Enables loading application content in an iFrame.
	AllowIframe param.Field[bool] `json:"allow_iframe"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool] `json:"skip_interstitial"`
}

func (r AccessApplicationSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationSettingUpdateResponseEnvelope struct {
	Errors   []AccessApplicationSettingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationSettingUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationSettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationSettingUpdateResponse                `json:"result"`
	JSON    accessApplicationSettingUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationSettingUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessApplicationSettingUpdateResponseEnvelope]
type accessApplicationSettingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationSettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingUpdateResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessApplicationSettingUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationSettingUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationSettingUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationSettingUpdateResponseEnvelopeErrors]
type accessApplicationSettingUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationSettingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessApplicationSettingUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationSettingUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationSettingUpdateResponseEnvelopeErrorsSource]
type accessApplicationSettingUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationSettingUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingUpdateResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AccessApplicationSettingUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationSettingUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationSettingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationSettingUpdateResponseEnvelopeMessages]
type accessApplicationSettingUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationSettingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    accessApplicationSettingUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationSettingUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessApplicationSettingUpdateResponseEnvelopeMessagesSource]
type accessApplicationSettingUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationSettingUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationSettingUpdateResponseEnvelopeSuccess bool

const (
	AccessApplicationSettingUpdateResponseEnvelopeSuccessTrue AccessApplicationSettingUpdateResponseEnvelopeSuccess = true
)

func (r AccessApplicationSettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationSettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationSettingEditParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Enables loading application content in an iFrame.
	AllowIframe param.Field[bool] `json:"allow_iframe"`
	// Enables automatic authentication through cloudflared.
	SkipInterstitial param.Field[bool] `json:"skip_interstitial"`
}

func (r AccessApplicationSettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessApplicationSettingEditResponseEnvelope struct {
	Errors   []AccessApplicationSettingEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationSettingEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationSettingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationSettingEditResponse                `json:"result"`
	JSON    accessApplicationSettingEditResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationSettingEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessApplicationSettingEditResponseEnvelope]
type accessApplicationSettingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationSettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingEditResponseEnvelopeErrors struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessApplicationSettingEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationSettingEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationSettingEditResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationSettingEditResponseEnvelopeErrors]
type accessApplicationSettingEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationSettingEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingEditResponseEnvelopeErrorsSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessApplicationSettingEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationSettingEditResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationSettingEditResponseEnvelopeErrorsSource]
type accessApplicationSettingEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationSettingEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingEditResponseEnvelopeMessages struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessApplicationSettingEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationSettingEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationSettingEditResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationSettingEditResponseEnvelopeMessages]
type accessApplicationSettingEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationSettingEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationSettingEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessApplicationSettingEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationSettingEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationSettingEditResponseEnvelopeMessagesSource]
type accessApplicationSettingEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationSettingEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationSettingEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationSettingEditResponseEnvelopeSuccess bool

const (
	AccessApplicationSettingEditResponseEnvelopeSuccessTrue AccessApplicationSettingEditResponseEnvelopeSuccess = true
)

func (r AccessApplicationSettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationSettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
