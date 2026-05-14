// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization

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

// URLNormalizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLNormalizationService] method instead.
type URLNormalizationService struct {
	Options []option.RequestOption
}

// NewURLNormalizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewURLNormalizationService(opts ...option.RequestOption) (r *URLNormalizationService) {
	r = &URLNormalizationService{}
	r.Options = opts
	return
}

// Updates the URL Normalization settings.
func (r *URLNormalizationService) Update(ctx context.Context, params URLNormalizationUpdateParams, opts ...option.RequestOption) (res *URLNormalizationUpdateResponse, err error) {
	var env URLNormalizationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes the URL Normalization settings.
func (r *URLNormalizationService) Delete(ctx context.Context, body URLNormalizationDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches the current URL Normalization settings.
func (r *URLNormalizationService) Get(ctx context.Context, query URLNormalizationGetParams, opts ...option.RequestOption) (res *URLNormalizationGetResponse, err error) {
	var env URLNormalizationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type URLNormalizationUpdateResponse struct {
	// The scope of the URL normalization.
	Scope URLNormalizationUpdateResponseScope `json:"scope,required"`
	// The type of URL normalization performed by Cloudflare.
	Type URLNormalizationUpdateResponseType `json:"type,required"`
	JSON urlNormalizationUpdateResponseJSON `json:"-"`
}

// urlNormalizationUpdateResponseJSON contains the JSON metadata for the struct
// [URLNormalizationUpdateResponse]
type urlNormalizationUpdateResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The scope of the URL normalization.
type URLNormalizationUpdateResponseScope string

const (
	URLNormalizationUpdateResponseScopeIncoming URLNormalizationUpdateResponseScope = "incoming"
	URLNormalizationUpdateResponseScopeBoth     URLNormalizationUpdateResponseScope = "both"
	URLNormalizationUpdateResponseScopeNone     URLNormalizationUpdateResponseScope = "none"
)

func (r URLNormalizationUpdateResponseScope) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseScopeIncoming, URLNormalizationUpdateResponseScopeBoth, URLNormalizationUpdateResponseScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationUpdateResponseType string

const (
	URLNormalizationUpdateResponseTypeCloudflare URLNormalizationUpdateResponseType = "cloudflare"
	URLNormalizationUpdateResponseTypeRfc3986    URLNormalizationUpdateResponseType = "rfc3986"
)

func (r URLNormalizationUpdateResponseType) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseTypeCloudflare, URLNormalizationUpdateResponseTypeRfc3986:
		return true
	}
	return false
}

// A result.
type URLNormalizationGetResponse struct {
	// The scope of the URL normalization.
	Scope URLNormalizationGetResponseScope `json:"scope,required"`
	// The type of URL normalization performed by Cloudflare.
	Type URLNormalizationGetResponseType `json:"type,required"`
	JSON urlNormalizationGetResponseJSON `json:"-"`
}

// urlNormalizationGetResponseJSON contains the JSON metadata for the struct
// [URLNormalizationGetResponse]
type urlNormalizationGetResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseJSON) RawJSON() string {
	return r.raw
}

// The scope of the URL normalization.
type URLNormalizationGetResponseScope string

const (
	URLNormalizationGetResponseScopeIncoming URLNormalizationGetResponseScope = "incoming"
	URLNormalizationGetResponseScopeBoth     URLNormalizationGetResponseScope = "both"
	URLNormalizationGetResponseScopeNone     URLNormalizationGetResponseScope = "none"
)

func (r URLNormalizationGetResponseScope) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseScopeIncoming, URLNormalizationGetResponseScopeBoth, URLNormalizationGetResponseScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationGetResponseType string

const (
	URLNormalizationGetResponseTypeCloudflare URLNormalizationGetResponseType = "cloudflare"
	URLNormalizationGetResponseTypeRfc3986    URLNormalizationGetResponseType = "rfc3986"
)

func (r URLNormalizationGetResponseType) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseTypeCloudflare, URLNormalizationGetResponseTypeRfc3986:
		return true
	}
	return false
}

type URLNormalizationUpdateParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The scope of the URL normalization.
	Scope param.Field[URLNormalizationUpdateParamsScope] `json:"scope,required"`
	// The type of URL normalization performed by Cloudflare.
	Type param.Field[URLNormalizationUpdateParamsType] `json:"type,required"`
}

func (r URLNormalizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The scope of the URL normalization.
type URLNormalizationUpdateParamsScope string

const (
	URLNormalizationUpdateParamsScopeIncoming URLNormalizationUpdateParamsScope = "incoming"
	URLNormalizationUpdateParamsScopeBoth     URLNormalizationUpdateParamsScope = "both"
	URLNormalizationUpdateParamsScopeNone     URLNormalizationUpdateParamsScope = "none"
)

func (r URLNormalizationUpdateParamsScope) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateParamsScopeIncoming, URLNormalizationUpdateParamsScopeBoth, URLNormalizationUpdateParamsScopeNone:
		return true
	}
	return false
}

// The type of URL normalization performed by Cloudflare.
type URLNormalizationUpdateParamsType string

const (
	URLNormalizationUpdateParamsTypeCloudflare URLNormalizationUpdateParamsType = "cloudflare"
	URLNormalizationUpdateParamsTypeRfc3986    URLNormalizationUpdateParamsType = "rfc3986"
)

func (r URLNormalizationUpdateParamsType) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateParamsTypeCloudflare, URLNormalizationUpdateParamsTypeRfc3986:
		return true
	}
	return false
}

// A response object.
type URLNormalizationUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []URLNormalizationUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []URLNormalizationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result URLNormalizationUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success URLNormalizationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    urlNormalizationUpdateResponseEnvelopeJSON    `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [URLNormalizationUpdateResponseEnvelope]
type urlNormalizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON   urlNormalizationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [URLNormalizationUpdateResponseEnvelopeErrors]
type urlNormalizationUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationUpdateResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer,required"`
	JSON    urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [URLNormalizationUpdateResponseEnvelopeErrorsSource]
type urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON   urlNormalizationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [URLNormalizationUpdateResponseEnvelopeMessages]
type urlNormalizationUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationUpdateResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                   `json:"pointer,required"`
	JSON    urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [URLNormalizationUpdateResponseEnvelopeMessagesSource]
type urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type URLNormalizationUpdateResponseEnvelopeSuccess bool

const (
	URLNormalizationUpdateResponseEnvelopeSuccessTrue URLNormalizationUpdateResponseEnvelopeSuccess = true
)

func (r URLNormalizationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case URLNormalizationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type URLNormalizationDeleteParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type URLNormalizationGetParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// A response object.
type URLNormalizationGetResponseEnvelope struct {
	// A list of error messages.
	Errors []URLNormalizationGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []URLNormalizationGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result URLNormalizationGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success URLNormalizationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    urlNormalizationGetResponseEnvelopeJSON    `json:"-"`
}

// urlNormalizationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [URLNormalizationGetResponseEnvelope]
type urlNormalizationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   urlNormalizationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// urlNormalizationGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [URLNormalizationGetResponseEnvelopeErrors]
type urlNormalizationGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                              `json:"pointer,required"`
	JSON    urlNormalizationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// urlNormalizationGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [URLNormalizationGetResponseEnvelopeErrorsSource]
type urlNormalizationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type URLNormalizationGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source URLNormalizationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   urlNormalizationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// urlNormalizationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [URLNormalizationGetResponseEnvelopeMessages]
type urlNormalizationGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type URLNormalizationGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                `json:"pointer,required"`
	JSON    urlNormalizationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// urlNormalizationGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [URLNormalizationGetResponseEnvelopeMessagesSource]
type urlNormalizationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type URLNormalizationGetResponseEnvelopeSuccess bool

const (
	URLNormalizationGetResponseEnvelopeSuccessTrue URLNormalizationGetResponseEnvelopeSuccess = true
)

func (r URLNormalizationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case URLNormalizationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
