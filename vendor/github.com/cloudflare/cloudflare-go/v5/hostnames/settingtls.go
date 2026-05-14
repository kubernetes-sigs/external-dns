// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hostnames

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// SettingTLSService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingTLSService] method instead.
type SettingTLSService struct {
	Options []option.RequestOption
}

// NewSettingTLSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingTLSService(opts ...option.RequestOption) (r *SettingTLSService) {
	r = &SettingTLSService{}
	r.Options = opts
	return
}

// Update the tls setting value for the hostname.
func (r *SettingTLSService) Update(ctx context.Context, settingID SettingTLSUpdateParamsSettingID, hostname string, params SettingTLSUpdateParams, opts ...option.RequestOption) (res *Setting, err error) {
	var env SettingTLSUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", params.ZoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the tls setting value for the hostname.
func (r *SettingTLSService) Delete(ctx context.Context, settingID SettingTLSDeleteParamsSettingID, hostname string, body SettingTLSDeleteParams, opts ...option.RequestOption) (res *SettingTLSDeleteResponse, err error) {
	var env SettingTLSDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v/%s", body.ZoneID, settingID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List the requested TLS setting for the hostnames under this zone.
func (r *SettingTLSService) Get(ctx context.Context, settingID SettingTLSGetParamsSettingID, query SettingTLSGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[SettingTLSGetResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hostnames/settings/%v", query.ZoneID, settingID)
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

// List the requested TLS setting for the hostnames under this zone.
func (r *SettingTLSService) GetAutoPaging(ctx context.Context, settingID SettingTLSGetParamsSettingID, query SettingTLSGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SettingTLSGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, settingID, query, opts...))
}

type Setting struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value SettingValueUnion `json:"value"`
	JSON  settingJSON       `json:"-"`
}

// settingJSON contains the JSON metadata for the struct [Setting]
type settingJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Setting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingJSON) RawJSON() string {
	return r.raw
}

// The tls setting value.
//
// Union satisfied by [shared.UnionFloat], [shared.UnionString] or
// [SettingValueArray].
type SettingValueUnion interface {
	ImplementsSettingValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SettingValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SettingValueArray{}),
		},
	)
}

type SettingValueArray []string

func (r SettingValueArray) ImplementsSettingValueUnion() {}

// The tls setting value.
//
// Satisfied by [shared.UnionFloat], [shared.UnionString],
// [hostnames.SettingValueArrayParam].
type SettingValueUnionParam interface {
	ImplementsSettingValueUnionParam()
}

type SettingValueArrayParam []string

func (r SettingValueArrayParam) ImplementsSettingValueUnionParam() {}

type SettingTLSDeleteResponse struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value SettingValueUnion            `json:"value"`
	JSON  settingTLSDeleteResponseJSON `json:"-"`
}

// settingTLSDeleteResponseJSON contains the JSON metadata for the struct
// [SettingTLSDeleteResponse]
type settingTLSDeleteResponseJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTLSGetResponse struct {
	// This is the time the tls setting was originally created for this hostname.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The hostname for which the tls settings are set.
	Hostname string `json:"hostname"`
	// Deployment status for the given tls setting.
	Status string `json:"status"`
	// This is the time the tls setting was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// The tls setting value.
	Value SettingValueUnion         `json:"value"`
	JSON  settingTLSGetResponseJSON `json:"-"`
}

// settingTLSGetResponseJSON contains the JSON metadata for the struct
// [SettingTLSGetResponse]
type settingTLSGetResponseJSON struct {
	CreatedAt   apijson.Field
	Hostname    apijson.Field
	Status      apijson.Field
	UpdatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingTLSUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The tls setting value.
	Value param.Field[SettingValueUnionParam] `json:"value,required"`
}

func (r SettingTLSUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The TLS Setting name.
type SettingTLSUpdateParamsSettingID string

const (
	SettingTLSUpdateParamsSettingIDCiphers       SettingTLSUpdateParamsSettingID = "ciphers"
	SettingTLSUpdateParamsSettingIDMinTLSVersion SettingTLSUpdateParamsSettingID = "min_tls_version"
	SettingTLSUpdateParamsSettingIDHTTP2         SettingTLSUpdateParamsSettingID = "http2"
)

func (r SettingTLSUpdateParamsSettingID) IsKnown() bool {
	switch r {
	case SettingTLSUpdateParamsSettingIDCiphers, SettingTLSUpdateParamsSettingIDMinTLSVersion, SettingTLSUpdateParamsSettingIDHTTP2:
		return true
	}
	return false
}

type SettingTLSUpdateResponseEnvelope struct {
	Errors   []SettingTLSUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTLSUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingTLSUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Setting                                 `json:"result"`
	JSON    settingTLSUpdateResponseEnvelopeJSON    `json:"-"`
}

// settingTLSUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingTLSUpdateResponseEnvelope]
type settingTLSUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTLSUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SettingTLSUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingTLSUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingTLSUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTLSUpdateResponseEnvelopeErrors]
type settingTLSUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTLSUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    settingTLSUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingTLSUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingTLSUpdateResponseEnvelopeErrorsSource]
type settingTLSUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingTLSUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SettingTLSUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingTLSUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingTLSUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingTLSUpdateResponseEnvelopeMessages]
type settingTLSUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTLSUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    settingTLSUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingTLSUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingTLSUpdateResponseEnvelopeMessagesSource]
type settingTLSUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingTLSUpdateResponseEnvelopeSuccess bool

const (
	SettingTLSUpdateResponseEnvelopeSuccessTrue SettingTLSUpdateResponseEnvelopeSuccess = true
)

func (r SettingTLSUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTLSUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingTLSDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// The TLS Setting name.
type SettingTLSDeleteParamsSettingID string

const (
	SettingTLSDeleteParamsSettingIDCiphers       SettingTLSDeleteParamsSettingID = "ciphers"
	SettingTLSDeleteParamsSettingIDMinTLSVersion SettingTLSDeleteParamsSettingID = "min_tls_version"
	SettingTLSDeleteParamsSettingIDHTTP2         SettingTLSDeleteParamsSettingID = "http2"
)

func (r SettingTLSDeleteParamsSettingID) IsKnown() bool {
	switch r {
	case SettingTLSDeleteParamsSettingIDCiphers, SettingTLSDeleteParamsSettingIDMinTLSVersion, SettingTLSDeleteParamsSettingIDHTTP2:
		return true
	}
	return false
}

type SettingTLSDeleteResponseEnvelope struct {
	Errors   []SettingTLSDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingTLSDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingTLSDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingTLSDeleteResponse                `json:"result"`
	JSON    settingTLSDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingTLSDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingTLSDeleteResponseEnvelope]
type settingTLSDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingTLSDeleteResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SettingTLSDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingTLSDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingTLSDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingTLSDeleteResponseEnvelopeErrors]
type settingTLSDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingTLSDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    settingTLSDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingTLSDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingTLSDeleteResponseEnvelopeErrorsSource]
type settingTLSDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingTLSDeleteResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SettingTLSDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingTLSDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingTLSDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingTLSDeleteResponseEnvelopeMessages]
type settingTLSDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingTLSDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    settingTLSDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingTLSDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingTLSDeleteResponseEnvelopeMessagesSource]
type settingTLSDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingTLSDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingTLSDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingTLSDeleteResponseEnvelopeSuccess bool

const (
	SettingTLSDeleteResponseEnvelopeSuccessTrue SettingTLSDeleteResponseEnvelopeSuccess = true
)

func (r SettingTLSDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingTLSDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingTLSGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// The TLS Setting name.
type SettingTLSGetParamsSettingID string

const (
	SettingTLSGetParamsSettingIDCiphers       SettingTLSGetParamsSettingID = "ciphers"
	SettingTLSGetParamsSettingIDMinTLSVersion SettingTLSGetParamsSettingID = "min_tls_version"
	SettingTLSGetParamsSettingIDHTTP2         SettingTLSGetParamsSettingID = "http2"
)

func (r SettingTLSGetParamsSettingID) IsKnown() bool {
	switch r {
	case SettingTLSGetParamsSettingIDCiphers, SettingTLSGetParamsSettingIDMinTLSVersion, SettingTLSGetParamsSettingIDHTTP2:
		return true
	}
	return false
}
