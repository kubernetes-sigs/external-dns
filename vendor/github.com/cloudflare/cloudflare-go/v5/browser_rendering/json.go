// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// JsonService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJsonService] method instead.
type JsonService struct {
	Options []option.RequestOption
}

// NewJsonService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJsonService(opts ...option.RequestOption) (r *JsonService) {
	r = &JsonService{}
	r.Options = opts
	return
}

// Gets json from a webpage from a provided URL or HTML. Pass `prompt` or `schema`
// in the body. Control page loading with `gotoOptions` and `waitFor*` options.
func (r *JsonService) New(ctx context.Context, params JsonNewParams, opts ...option.RequestOption) (res *JsonNewResponse, err error) {
	var env JsonNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/json", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type JsonNewResponse map[string]interface{}

type JsonNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cache TTL default is 5s. Set to 0 to disable.
	CacheTTL param.Field[float64] `query:"cacheTTL"`
	// The maximum duration allowed for the browser action to complete after the page
	// has loaded (such as taking screenshots, extracting content, or generating PDFs).
	// If this time limit is exceeded, the action stops and returns a timeout error.
	ActionTimeout param.Field[float64] `json:"actionTimeout"`
	// Adds a `<script>` tag into the page with the desired URL or content.
	AddScriptTag param.Field[[]JsonNewParamsAddScriptTag] `json:"addScriptTag"`
	// Adds a `<link rel="stylesheet">` tag into the page with the desired URL or a
	// `<style type="text/css">` tag with the content.
	AddStyleTag param.Field[[]JsonNewParamsAddStyleTag] `json:"addStyleTag"`
	// Only allow requests that match the provided regex patterns, eg. '/^.\*\.(css)'.
	AllowRequestPattern param.Field[[]string] `json:"allowRequestPattern"`
	// Only allow requests that match the provided resource types, eg. 'image' or
	// 'script'.
	AllowResourceTypes param.Field[[]JsonNewParamsAllowResourceType] `json:"allowResourceTypes"`
	// Provide credentials for HTTP authentication.
	Authenticate param.Field[JsonNewParamsAuthenticate] `json:"authenticate"`
	// Attempt to proceed when 'awaited' events fail or timeout.
	BestAttempt param.Field[bool] `json:"bestAttempt"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setcookie).
	Cookies param.Field[[]JsonNewParamsCookie] `json:"cookies"`
	// Optional list of custom AI models to use for the request. The models will be
	// tried in the order provided, and in case a model returns an error, the next one
	// will be used as fallback.
	CustomAI         param.Field[[]JsonNewParamsCustomAI] `json:"custom_ai"`
	EmulateMediaType param.Field[string]                  `json:"emulateMediaType"`
	// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
	GotoOptions param.Field[JsonNewParamsGotoOptions] `json:"gotoOptions"`
	// Set the content of the page, eg: `<h1>Hello World!!</h1>`. Either `html` or
	// `url` must be set.
	HTML   param.Field[string] `json:"html"`
	Prompt param.Field[string] `json:"prompt"`
	// Block undesired requests that match the provided regex patterns, eg.
	// '/^.\*\.(css)'.
	RejectRequestPattern param.Field[[]string] `json:"rejectRequestPattern"`
	// Block undesired requests that match the provided resource types, eg. 'image' or
	// 'script'.
	RejectResourceTypes  param.Field[[]JsonNewParamsRejectResourceType] `json:"rejectResourceTypes"`
	ResponseFormat       param.Field[JsonNewParamsResponseFormat]       `json:"response_format"`
	SetExtraHTTPHeaders  param.Field[map[string]string]                 `json:"setExtraHTTPHeaders"`
	SetJavaScriptEnabled param.Field[bool]                              `json:"setJavaScriptEnabled"`
	// URL to navigate to, eg. `https://example.com`.
	URL       param.Field[string] `json:"url" format:"uri"`
	UserAgent param.Field[string] `json:"userAgent"`
	// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
	Viewport param.Field[JsonNewParamsViewport] `json:"viewport"`
	// Wait for the selector to appear in page. Check
	// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
	WaitForSelector param.Field[JsonNewParamsWaitForSelector] `json:"waitForSelector"`
	// Waits for a specified timeout before continuing.
	WaitForTimeout param.Field[float64] `json:"waitForTimeout"`
}

func (r JsonNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [JsonNewParams]'s query parameters as `url.Values`.
func (r JsonNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type JsonNewParamsAddScriptTag struct {
	ID      param.Field[string] `json:"id"`
	Content param.Field[string] `json:"content"`
	Type    param.Field[string] `json:"type"`
	URL     param.Field[string] `json:"url"`
}

func (r JsonNewParamsAddScriptTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsAddStyleTag struct {
	Content param.Field[string] `json:"content"`
	URL     param.Field[string] `json:"url"`
}

func (r JsonNewParamsAddStyleTag) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsAllowResourceType string

const (
	JsonNewParamsAllowResourceTypeDocument           JsonNewParamsAllowResourceType = "document"
	JsonNewParamsAllowResourceTypeStylesheet         JsonNewParamsAllowResourceType = "stylesheet"
	JsonNewParamsAllowResourceTypeImage              JsonNewParamsAllowResourceType = "image"
	JsonNewParamsAllowResourceTypeMedia              JsonNewParamsAllowResourceType = "media"
	JsonNewParamsAllowResourceTypeFont               JsonNewParamsAllowResourceType = "font"
	JsonNewParamsAllowResourceTypeScript             JsonNewParamsAllowResourceType = "script"
	JsonNewParamsAllowResourceTypeTexttrack          JsonNewParamsAllowResourceType = "texttrack"
	JsonNewParamsAllowResourceTypeXHR                JsonNewParamsAllowResourceType = "xhr"
	JsonNewParamsAllowResourceTypeFetch              JsonNewParamsAllowResourceType = "fetch"
	JsonNewParamsAllowResourceTypePrefetch           JsonNewParamsAllowResourceType = "prefetch"
	JsonNewParamsAllowResourceTypeEventsource        JsonNewParamsAllowResourceType = "eventsource"
	JsonNewParamsAllowResourceTypeWebsocket          JsonNewParamsAllowResourceType = "websocket"
	JsonNewParamsAllowResourceTypeManifest           JsonNewParamsAllowResourceType = "manifest"
	JsonNewParamsAllowResourceTypeSignedexchange     JsonNewParamsAllowResourceType = "signedexchange"
	JsonNewParamsAllowResourceTypePing               JsonNewParamsAllowResourceType = "ping"
	JsonNewParamsAllowResourceTypeCspviolationreport JsonNewParamsAllowResourceType = "cspviolationreport"
	JsonNewParamsAllowResourceTypePreflight          JsonNewParamsAllowResourceType = "preflight"
	JsonNewParamsAllowResourceTypeOther              JsonNewParamsAllowResourceType = "other"
)

func (r JsonNewParamsAllowResourceType) IsKnown() bool {
	switch r {
	case JsonNewParamsAllowResourceTypeDocument, JsonNewParamsAllowResourceTypeStylesheet, JsonNewParamsAllowResourceTypeImage, JsonNewParamsAllowResourceTypeMedia, JsonNewParamsAllowResourceTypeFont, JsonNewParamsAllowResourceTypeScript, JsonNewParamsAllowResourceTypeTexttrack, JsonNewParamsAllowResourceTypeXHR, JsonNewParamsAllowResourceTypeFetch, JsonNewParamsAllowResourceTypePrefetch, JsonNewParamsAllowResourceTypeEventsource, JsonNewParamsAllowResourceTypeWebsocket, JsonNewParamsAllowResourceTypeManifest, JsonNewParamsAllowResourceTypeSignedexchange, JsonNewParamsAllowResourceTypePing, JsonNewParamsAllowResourceTypeCspviolationreport, JsonNewParamsAllowResourceTypePreflight, JsonNewParamsAllowResourceTypeOther:
		return true
	}
	return false
}

// Provide credentials for HTTP authentication.
type JsonNewParamsAuthenticate struct {
	Password param.Field[string] `json:"password,required"`
	Username param.Field[string] `json:"username,required"`
}

func (r JsonNewParamsAuthenticate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsCookie struct {
	Name         param.Field[string]                           `json:"name,required"`
	Value        param.Field[string]                           `json:"value,required"`
	Domain       param.Field[string]                           `json:"domain"`
	Expires      param.Field[float64]                          `json:"expires"`
	HTTPOnly     param.Field[bool]                             `json:"httpOnly"`
	PartitionKey param.Field[string]                           `json:"partitionKey"`
	Path         param.Field[string]                           `json:"path"`
	Priority     param.Field[JsonNewParamsCookiesPriority]     `json:"priority"`
	SameParty    param.Field[bool]                             `json:"sameParty"`
	SameSite     param.Field[JsonNewParamsCookiesSameSite]     `json:"sameSite"`
	Secure       param.Field[bool]                             `json:"secure"`
	SourcePort   param.Field[float64]                          `json:"sourcePort"`
	SourceScheme param.Field[JsonNewParamsCookiesSourceScheme] `json:"sourceScheme"`
	URL          param.Field[string]                           `json:"url"`
}

func (r JsonNewParamsCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsCookiesPriority string

const (
	JsonNewParamsCookiesPriorityLow    JsonNewParamsCookiesPriority = "Low"
	JsonNewParamsCookiesPriorityMedium JsonNewParamsCookiesPriority = "Medium"
	JsonNewParamsCookiesPriorityHigh   JsonNewParamsCookiesPriority = "High"
)

func (r JsonNewParamsCookiesPriority) IsKnown() bool {
	switch r {
	case JsonNewParamsCookiesPriorityLow, JsonNewParamsCookiesPriorityMedium, JsonNewParamsCookiesPriorityHigh:
		return true
	}
	return false
}

type JsonNewParamsCookiesSameSite string

const (
	JsonNewParamsCookiesSameSiteStrict JsonNewParamsCookiesSameSite = "Strict"
	JsonNewParamsCookiesSameSiteLax    JsonNewParamsCookiesSameSite = "Lax"
	JsonNewParamsCookiesSameSiteNone   JsonNewParamsCookiesSameSite = "None"
)

func (r JsonNewParamsCookiesSameSite) IsKnown() bool {
	switch r {
	case JsonNewParamsCookiesSameSiteStrict, JsonNewParamsCookiesSameSiteLax, JsonNewParamsCookiesSameSiteNone:
		return true
	}
	return false
}

type JsonNewParamsCookiesSourceScheme string

const (
	JsonNewParamsCookiesSourceSchemeUnset     JsonNewParamsCookiesSourceScheme = "Unset"
	JsonNewParamsCookiesSourceSchemeNonSecure JsonNewParamsCookiesSourceScheme = "NonSecure"
	JsonNewParamsCookiesSourceSchemeSecure    JsonNewParamsCookiesSourceScheme = "Secure"
)

func (r JsonNewParamsCookiesSourceScheme) IsKnown() bool {
	switch r {
	case JsonNewParamsCookiesSourceSchemeUnset, JsonNewParamsCookiesSourceSchemeNonSecure, JsonNewParamsCookiesSourceSchemeSecure:
		return true
	}
	return false
}

type JsonNewParamsCustomAI struct {
	// Authorization token for the AI model: `Bearer <token>`.
	Authorization param.Field[string] `json:"authorization,required"`
	// AI model to use for the request. Must be formed as `<provider>/<model_name>`,
	// e.g. `workers-ai/@cf/meta/llama-3.3-70b-instruct-fp8-fast`
	Model param.Field[string] `json:"model,required"`
}

func (r JsonNewParamsCustomAI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Check [options](https://pptr.dev/api/puppeteer.gotooptions).
type JsonNewParamsGotoOptions struct {
	Referer        param.Field[string]                                 `json:"referer"`
	ReferrerPolicy param.Field[string]                                 `json:"referrerPolicy"`
	Timeout        param.Field[float64]                                `json:"timeout"`
	WaitUntil      param.Field[JsonNewParamsGotoOptionsWaitUntilUnion] `json:"waitUntil"`
}

func (r JsonNewParamsGotoOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [browser_rendering.JsonNewParamsGotoOptionsWaitUntilString],
// [browser_rendering.JsonNewParamsGotoOptionsWaitUntilArray].
type JsonNewParamsGotoOptionsWaitUntilUnion interface {
	implementsJsonNewParamsGotoOptionsWaitUntilUnion()
}

type JsonNewParamsGotoOptionsWaitUntilString string

const (
	JsonNewParamsGotoOptionsWaitUntilStringLoad             JsonNewParamsGotoOptionsWaitUntilString = "load"
	JsonNewParamsGotoOptionsWaitUntilStringDomcontentloaded JsonNewParamsGotoOptionsWaitUntilString = "domcontentloaded"
	JsonNewParamsGotoOptionsWaitUntilStringNetworkidle0     JsonNewParamsGotoOptionsWaitUntilString = "networkidle0"
	JsonNewParamsGotoOptionsWaitUntilStringNetworkidle2     JsonNewParamsGotoOptionsWaitUntilString = "networkidle2"
)

func (r JsonNewParamsGotoOptionsWaitUntilString) IsKnown() bool {
	switch r {
	case JsonNewParamsGotoOptionsWaitUntilStringLoad, JsonNewParamsGotoOptionsWaitUntilStringDomcontentloaded, JsonNewParamsGotoOptionsWaitUntilStringNetworkidle0, JsonNewParamsGotoOptionsWaitUntilStringNetworkidle2:
		return true
	}
	return false
}

func (r JsonNewParamsGotoOptionsWaitUntilString) implementsJsonNewParamsGotoOptionsWaitUntilUnion() {}

type JsonNewParamsGotoOptionsWaitUntilArray []JsonNewParamsGotoOptionsWaitUntilArrayItem

func (r JsonNewParamsGotoOptionsWaitUntilArray) implementsJsonNewParamsGotoOptionsWaitUntilUnion() {}

type JsonNewParamsGotoOptionsWaitUntilArrayItem string

const (
	JsonNewParamsGotoOptionsWaitUntilArrayItemLoad             JsonNewParamsGotoOptionsWaitUntilArrayItem = "load"
	JsonNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded JsonNewParamsGotoOptionsWaitUntilArrayItem = "domcontentloaded"
	JsonNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0     JsonNewParamsGotoOptionsWaitUntilArrayItem = "networkidle0"
	JsonNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2     JsonNewParamsGotoOptionsWaitUntilArrayItem = "networkidle2"
)

func (r JsonNewParamsGotoOptionsWaitUntilArrayItem) IsKnown() bool {
	switch r {
	case JsonNewParamsGotoOptionsWaitUntilArrayItemLoad, JsonNewParamsGotoOptionsWaitUntilArrayItemDomcontentloaded, JsonNewParamsGotoOptionsWaitUntilArrayItemNetworkidle0, JsonNewParamsGotoOptionsWaitUntilArrayItemNetworkidle2:
		return true
	}
	return false
}

type JsonNewParamsRejectResourceType string

const (
	JsonNewParamsRejectResourceTypeDocument           JsonNewParamsRejectResourceType = "document"
	JsonNewParamsRejectResourceTypeStylesheet         JsonNewParamsRejectResourceType = "stylesheet"
	JsonNewParamsRejectResourceTypeImage              JsonNewParamsRejectResourceType = "image"
	JsonNewParamsRejectResourceTypeMedia              JsonNewParamsRejectResourceType = "media"
	JsonNewParamsRejectResourceTypeFont               JsonNewParamsRejectResourceType = "font"
	JsonNewParamsRejectResourceTypeScript             JsonNewParamsRejectResourceType = "script"
	JsonNewParamsRejectResourceTypeTexttrack          JsonNewParamsRejectResourceType = "texttrack"
	JsonNewParamsRejectResourceTypeXHR                JsonNewParamsRejectResourceType = "xhr"
	JsonNewParamsRejectResourceTypeFetch              JsonNewParamsRejectResourceType = "fetch"
	JsonNewParamsRejectResourceTypePrefetch           JsonNewParamsRejectResourceType = "prefetch"
	JsonNewParamsRejectResourceTypeEventsource        JsonNewParamsRejectResourceType = "eventsource"
	JsonNewParamsRejectResourceTypeWebsocket          JsonNewParamsRejectResourceType = "websocket"
	JsonNewParamsRejectResourceTypeManifest           JsonNewParamsRejectResourceType = "manifest"
	JsonNewParamsRejectResourceTypeSignedexchange     JsonNewParamsRejectResourceType = "signedexchange"
	JsonNewParamsRejectResourceTypePing               JsonNewParamsRejectResourceType = "ping"
	JsonNewParamsRejectResourceTypeCspviolationreport JsonNewParamsRejectResourceType = "cspviolationreport"
	JsonNewParamsRejectResourceTypePreflight          JsonNewParamsRejectResourceType = "preflight"
	JsonNewParamsRejectResourceTypeOther              JsonNewParamsRejectResourceType = "other"
)

func (r JsonNewParamsRejectResourceType) IsKnown() bool {
	switch r {
	case JsonNewParamsRejectResourceTypeDocument, JsonNewParamsRejectResourceTypeStylesheet, JsonNewParamsRejectResourceTypeImage, JsonNewParamsRejectResourceTypeMedia, JsonNewParamsRejectResourceTypeFont, JsonNewParamsRejectResourceTypeScript, JsonNewParamsRejectResourceTypeTexttrack, JsonNewParamsRejectResourceTypeXHR, JsonNewParamsRejectResourceTypeFetch, JsonNewParamsRejectResourceTypePrefetch, JsonNewParamsRejectResourceTypeEventsource, JsonNewParamsRejectResourceTypeWebsocket, JsonNewParamsRejectResourceTypeManifest, JsonNewParamsRejectResourceTypeSignedexchange, JsonNewParamsRejectResourceTypePing, JsonNewParamsRejectResourceTypeCspviolationreport, JsonNewParamsRejectResourceTypePreflight, JsonNewParamsRejectResourceTypeOther:
		return true
	}
	return false
}

type JsonNewParamsResponseFormat struct {
	Type param.Field[string] `json:"type,required"`
	// Schema for the response format. More information here:
	// https://developers.cloudflare.com/workers-ai/json-mode/
	Schema param.Field[map[string]interface{}] `json:"schema"`
}

func (r JsonNewParamsResponseFormat) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Check [options](https://pptr.dev/api/puppeteer.page.setviewport).
type JsonNewParamsViewport struct {
	Height            param.Field[float64] `json:"height,required"`
	Width             param.Field[float64] `json:"width,required"`
	DeviceScaleFactor param.Field[float64] `json:"deviceScaleFactor"`
	HasTouch          param.Field[bool]    `json:"hasTouch"`
	IsLandscape       param.Field[bool]    `json:"isLandscape"`
	IsMobile          param.Field[bool]    `json:"isMobile"`
}

func (r JsonNewParamsViewport) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Wait for the selector to appear in page. Check
// [options](https://pptr.dev/api/puppeteer.page.waitforselector).
type JsonNewParamsWaitForSelector struct {
	Selector param.Field[string]                              `json:"selector,required"`
	Hidden   param.Field[JsonNewParamsWaitForSelectorHidden]  `json:"hidden"`
	Timeout  param.Field[float64]                             `json:"timeout"`
	Visible  param.Field[JsonNewParamsWaitForSelectorVisible] `json:"visible"`
}

func (r JsonNewParamsWaitForSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JsonNewParamsWaitForSelectorHidden bool

const (
	JsonNewParamsWaitForSelectorHiddenTrue JsonNewParamsWaitForSelectorHidden = true
)

func (r JsonNewParamsWaitForSelectorHidden) IsKnown() bool {
	switch r {
	case JsonNewParamsWaitForSelectorHiddenTrue:
		return true
	}
	return false
}

type JsonNewParamsWaitForSelectorVisible bool

const (
	JsonNewParamsWaitForSelectorVisibleTrue JsonNewParamsWaitForSelectorVisible = true
)

func (r JsonNewParamsWaitForSelectorVisible) IsKnown() bool {
	switch r {
	case JsonNewParamsWaitForSelectorVisibleTrue:
		return true
	}
	return false
}

type JsonNewResponseEnvelope struct {
	Result JsonNewResponse `json:"result,required"`
	// Response status
	Status bool                            `json:"status,required"`
	Errors []JsonNewResponseEnvelopeErrors `json:"errors"`
	JSON   jsonNewResponseEnvelopeJSON     `json:"-"`
}

// jsonNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [JsonNewResponseEnvelope]
type jsonNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Status      apijson.Field
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JsonNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jsonNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JsonNewResponseEnvelopeErrors struct {
	// Error code
	Code float64 `json:"code,required"`
	// Error Message
	Message string                            `json:"message,required"`
	JSON    jsonNewResponseEnvelopeErrorsJSON `json:"-"`
}

// jsonNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JsonNewResponseEnvelopeErrors]
type jsonNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JsonNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jsonNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}
