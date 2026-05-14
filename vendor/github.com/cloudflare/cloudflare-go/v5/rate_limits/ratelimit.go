// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rate_limits

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// RateLimitService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRateLimitService] method instead.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
type RateLimitService struct {
	Options []option.RequestOption
}

// NewRateLimitService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRateLimitService(opts ...option.RequestOption) (r *RateLimitService) {
	r = &RateLimitService{}
	r.Options = opts
	return
}

// Creates a new rate limit for a zone. Refer to the object definition for a list
// of required attributes.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) New(ctx context.Context, params RateLimitNewParams, opts ...option.RequestOption) (res *RateLimit, err error) {
	var env RateLimitNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the rate limits for a zone.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) List(ctx context.Context, params RateLimitListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[RateLimit], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Fetches the rate limits for a zone.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) ListAutoPaging(ctx context.Context, params RateLimitListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[RateLimit] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing rate limit.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Delete(ctx context.Context, rateLimitID string, body RateLimitDeleteParams, opts ...option.RequestOption) (res *RateLimitDeleteResponse, err error) {
	var env RateLimitDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", body.ZoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rate limit.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Edit(ctx context.Context, rateLimitID string, params RateLimitEditParams, opts ...option.RequestOption) (res *RateLimit, err error) {
	var env RateLimitEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", params.ZoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a rate limit.
//
// Deprecated: Rate limiting API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#rate-limiting-api-previous-version
// for full details.
func (r *RateLimitService) Get(ctx context.Context, rateLimitID string, query RateLimitGetParams, opts ...option.RequestOption) (res *RateLimit, err error) {
	var env RateLimitGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if rateLimitID == "" {
		err = errors.New("missing required rate_limit_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/rate_limits/%s", query.ZoneID, rateLimitID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// The action to apply to a matched request. The `log` action is only available on
// an Enterprise plan.
type Action string

const (
	ActionBlock            Action = "block"
	ActionChallenge        Action = "challenge"
	ActionJSChallenge      Action = "js_challenge"
	ActionManagedChallenge Action = "managed_challenge"
	ActionAllow            Action = "allow"
	ActionLog              Action = "log"
	ActionBypass           Action = "bypass"
)

func (r Action) IsKnown() bool {
	switch r {
	case ActionBlock, ActionChallenge, ActionJSChallenge, ActionManagedChallenge, ActionAllow, ActionLog, ActionBypass:
		return true
	}
	return false
}

type RateLimit struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action RateLimitAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []RateLimitBypass `json:"bypass"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match RateLimitMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64       `json:"threshold"`
	JSON      rateLimitJSON `json:"-"`
}

// rateLimitJSON contains the JSON metadata for the struct [RateLimit]
type rateLimitJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Bypass      apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	Match       apijson.Field
	Period      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitAction struct {
	// The action to perform.
	Mode RateLimitActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response RateLimitActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64             `json:"timeout"`
	JSON    rateLimitActionJSON `json:"-"`
}

// rateLimitActionJSON contains the JSON metadata for the struct [RateLimitAction]
type rateLimitActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitActionJSON) RawJSON() string {
	return r.raw
}

// The action to perform.
type RateLimitActionMode string

const (
	RateLimitActionModeSimulate         RateLimitActionMode = "simulate"
	RateLimitActionModeBan              RateLimitActionMode = "ban"
	RateLimitActionModeChallenge        RateLimitActionMode = "challenge"
	RateLimitActionModeJSChallenge      RateLimitActionMode = "js_challenge"
	RateLimitActionModeManagedChallenge RateLimitActionMode = "managed_challenge"
)

func (r RateLimitActionMode) IsKnown() bool {
	switch r {
	case RateLimitActionModeSimulate, RateLimitActionModeBan, RateLimitActionModeChallenge, RateLimitActionModeJSChallenge, RateLimitActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RateLimitActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                      `json:"content_type"`
	JSON        rateLimitActionResponseJSON `json:"-"`
}

// rateLimitActionResponseJSON contains the JSON metadata for the struct
// [RateLimitActionResponse]
type rateLimitActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitActionResponseJSON) RawJSON() string {
	return r.raw
}

type RateLimitBypass struct {
	Name RateLimitBypassName `json:"name"`
	// The URL to bypass.
	Value string              `json:"value"`
	JSON  rateLimitBypassJSON `json:"-"`
}

// rateLimitBypassJSON contains the JSON metadata for the struct [RateLimitBypass]
type rateLimitBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitBypassJSON) RawJSON() string {
	return r.raw
}

type RateLimitBypassName string

const (
	RateLimitBypassNameURL RateLimitBypassName = "url"
)

func (r RateLimitBypassName) IsKnown() bool {
	switch r {
	case RateLimitBypassNameURL:
		return true
	}
	return false
}

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitMatch struct {
	Headers  []RateLimitMatchHeader `json:"headers"`
	Request  RateLimitMatchRequest  `json:"request"`
	Response RateLimitMatchResponse `json:"response"`
	JSON     rateLimitMatchJSON     `json:"-"`
}

// rateLimitMatchJSON contains the JSON metadata for the struct [RateLimitMatch]
type rateLimitMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchJSON) RawJSON() string {
	return r.raw
}

type RateLimitMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op RateLimitMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                   `json:"value"`
	JSON  rateLimitMatchHeaderJSON `json:"-"`
}

// rateLimitMatchHeaderJSON contains the JSON metadata for the struct
// [RateLimitMatchHeader]
type rateLimitMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchHeaderJSON) RawJSON() string {
	return r.raw
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitMatchHeadersOp string

const (
	RateLimitMatchHeadersOpEq RateLimitMatchHeadersOp = "eq"
	RateLimitMatchHeadersOpNe RateLimitMatchHeadersOp = "ne"
)

func (r RateLimitMatchHeadersOp) IsKnown() bool {
	switch r {
	case RateLimitMatchHeadersOpEq, RateLimitMatchHeadersOpNe:
		return true
	}
	return false
}

type RateLimitMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []RateLimitMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                    `json:"url"`
	JSON rateLimitMatchRequestJSON `json:"-"`
}

// rateLimitMatchRequestJSON contains the JSON metadata for the struct
// [RateLimitMatchRequest]
type rateLimitMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchRequestJSON) RawJSON() string {
	return r.raw
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitMatchRequestMethod string

const (
	RateLimitMatchRequestMethodGet    RateLimitMatchRequestMethod = "GET"
	RateLimitMatchRequestMethodPost   RateLimitMatchRequestMethod = "POST"
	RateLimitMatchRequestMethodPut    RateLimitMatchRequestMethod = "PUT"
	RateLimitMatchRequestMethodDelete RateLimitMatchRequestMethod = "DELETE"
	RateLimitMatchRequestMethodPatch  RateLimitMatchRequestMethod = "PATCH"
	RateLimitMatchRequestMethodHead   RateLimitMatchRequestMethod = "HEAD"
	RateLimitMatchRequestMethod_All   RateLimitMatchRequestMethod = "_ALL_"
)

func (r RateLimitMatchRequestMethod) IsKnown() bool {
	switch r {
	case RateLimitMatchRequestMethodGet, RateLimitMatchRequestMethodPost, RateLimitMatchRequestMethodPut, RateLimitMatchRequestMethodDelete, RateLimitMatchRequestMethodPatch, RateLimitMatchRequestMethodHead, RateLimitMatchRequestMethod_All:
		return true
	}
	return false
}

type RateLimitMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                       `json:"origin_traffic"`
	JSON          rateLimitMatchResponseJSON `json:"-"`
}

// rateLimitMatchResponseJSON contains the JSON metadata for the struct
// [RateLimitMatchResponse]
type rateLimitMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RateLimitMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitMatchResponseJSON) RawJSON() string {
	return r.raw
}

type RateLimitDeleteResponse struct {
	// The unique identifier of the rate limit.
	ID string `json:"id"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action RateLimitDeleteResponseAction `json:"action"`
	// Criteria specifying when the current rate limit should be bypassed. You can
	// specify that the rate limit should not apply to one or more URLs.
	Bypass []RateLimitDeleteResponseBypass `json:"bypass"`
	// An informative summary of the rule. This value is sanitized and any tags will be
	// removed.
	Description string `json:"description"`
	// When true, indicates that the rate limit is currently disabled.
	Disabled bool `json:"disabled"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match RateLimitDeleteResponseMatch `json:"match"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period float64 `json:"period"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold float64                     `json:"threshold"`
	JSON      rateLimitDeleteResponseJSON `json:"-"`
}

// rateLimitDeleteResponseJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponse]
type rateLimitDeleteResponseJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Bypass      apijson.Field
	Description apijson.Field
	Disabled    apijson.Field
	Match       apijson.Field
	Period      apijson.Field
	Threshold   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitDeleteResponseAction struct {
	// The action to perform.
	Mode RateLimitDeleteResponseActionMode `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response RateLimitDeleteResponseActionResponse `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout float64                           `json:"timeout"`
	JSON    rateLimitDeleteResponseActionJSON `json:"-"`
}

// rateLimitDeleteResponseActionJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponseAction]
type rateLimitDeleteResponseActionJSON struct {
	Mode        apijson.Field
	Response    apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseActionJSON) RawJSON() string {
	return r.raw
}

// The action to perform.
type RateLimitDeleteResponseActionMode string

const (
	RateLimitDeleteResponseActionModeSimulate         RateLimitDeleteResponseActionMode = "simulate"
	RateLimitDeleteResponseActionModeBan              RateLimitDeleteResponseActionMode = "ban"
	RateLimitDeleteResponseActionModeChallenge        RateLimitDeleteResponseActionMode = "challenge"
	RateLimitDeleteResponseActionModeJSChallenge      RateLimitDeleteResponseActionMode = "js_challenge"
	RateLimitDeleteResponseActionModeManagedChallenge RateLimitDeleteResponseActionMode = "managed_challenge"
)

func (r RateLimitDeleteResponseActionMode) IsKnown() bool {
	switch r {
	case RateLimitDeleteResponseActionModeSimulate, RateLimitDeleteResponseActionModeBan, RateLimitDeleteResponseActionModeChallenge, RateLimitDeleteResponseActionModeJSChallenge, RateLimitDeleteResponseActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RateLimitDeleteResponseActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body string `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType string                                    `json:"content_type"`
	JSON        rateLimitDeleteResponseActionResponseJSON `json:"-"`
}

// rateLimitDeleteResponseActionResponseJSON contains the JSON metadata for the
// struct [RateLimitDeleteResponseActionResponse]
type rateLimitDeleteResponseActionResponseJSON struct {
	Body        apijson.Field
	ContentType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseActionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseActionResponseJSON) RawJSON() string {
	return r.raw
}

type RateLimitDeleteResponseBypass struct {
	Name RateLimitDeleteResponseBypassName `json:"name"`
	// The URL to bypass.
	Value string                            `json:"value"`
	JSON  rateLimitDeleteResponseBypassJSON `json:"-"`
}

// rateLimitDeleteResponseBypassJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponseBypass]
type rateLimitDeleteResponseBypassJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseBypass) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseBypassJSON) RawJSON() string {
	return r.raw
}

type RateLimitDeleteResponseBypassName string

const (
	RateLimitDeleteResponseBypassNameURL RateLimitDeleteResponseBypassName = "url"
)

func (r RateLimitDeleteResponseBypassName) IsKnown() bool {
	switch r {
	case RateLimitDeleteResponseBypassNameURL:
		return true
	}
	return false
}

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitDeleteResponseMatch struct {
	Headers  []RateLimitDeleteResponseMatchHeader `json:"headers"`
	Request  RateLimitDeleteResponseMatchRequest  `json:"request"`
	Response RateLimitDeleteResponseMatchResponse `json:"response"`
	JSON     rateLimitDeleteResponseMatchJSON     `json:"-"`
}

// rateLimitDeleteResponseMatchJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponseMatch]
type rateLimitDeleteResponseMatchJSON struct {
	Headers     apijson.Field
	Request     apijson.Field
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseMatch) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseMatchJSON) RawJSON() string {
	return r.raw
}

type RateLimitDeleteResponseMatchHeader struct {
	// The name of the response header to match.
	Name string `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op RateLimitDeleteResponseMatchHeadersOp `json:"op"`
	// The value of the response header, which must match exactly.
	Value string                                 `json:"value"`
	JSON  rateLimitDeleteResponseMatchHeaderJSON `json:"-"`
}

// rateLimitDeleteResponseMatchHeaderJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponseMatchHeader]
type rateLimitDeleteResponseMatchHeaderJSON struct {
	Name        apijson.Field
	Op          apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseMatchHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseMatchHeaderJSON) RawJSON() string {
	return r.raw
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitDeleteResponseMatchHeadersOp string

const (
	RateLimitDeleteResponseMatchHeadersOpEq RateLimitDeleteResponseMatchHeadersOp = "eq"
	RateLimitDeleteResponseMatchHeadersOpNe RateLimitDeleteResponseMatchHeadersOp = "ne"
)

func (r RateLimitDeleteResponseMatchHeadersOp) IsKnown() bool {
	switch r {
	case RateLimitDeleteResponseMatchHeadersOpEq, RateLimitDeleteResponseMatchHeadersOpNe:
		return true
	}
	return false
}

type RateLimitDeleteResponseMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods []RateLimitDeleteResponseMatchRequestMethod `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes []string `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL  string                                  `json:"url"`
	JSON rateLimitDeleteResponseMatchRequestJSON `json:"-"`
}

// rateLimitDeleteResponseMatchRequestJSON contains the JSON metadata for the
// struct [RateLimitDeleteResponseMatchRequest]
type rateLimitDeleteResponseMatchRequestJSON struct {
	Methods     apijson.Field
	Schemes     apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseMatchRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseMatchRequestJSON) RawJSON() string {
	return r.raw
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitDeleteResponseMatchRequestMethod string

const (
	RateLimitDeleteResponseMatchRequestMethodGet    RateLimitDeleteResponseMatchRequestMethod = "GET"
	RateLimitDeleteResponseMatchRequestMethodPost   RateLimitDeleteResponseMatchRequestMethod = "POST"
	RateLimitDeleteResponseMatchRequestMethodPut    RateLimitDeleteResponseMatchRequestMethod = "PUT"
	RateLimitDeleteResponseMatchRequestMethodDelete RateLimitDeleteResponseMatchRequestMethod = "DELETE"
	RateLimitDeleteResponseMatchRequestMethodPatch  RateLimitDeleteResponseMatchRequestMethod = "PATCH"
	RateLimitDeleteResponseMatchRequestMethodHead   RateLimitDeleteResponseMatchRequestMethod = "HEAD"
	RateLimitDeleteResponseMatchRequestMethod_All   RateLimitDeleteResponseMatchRequestMethod = "_ALL_"
)

func (r RateLimitDeleteResponseMatchRequestMethod) IsKnown() bool {
	switch r {
	case RateLimitDeleteResponseMatchRequestMethodGet, RateLimitDeleteResponseMatchRequestMethodPost, RateLimitDeleteResponseMatchRequestMethodPut, RateLimitDeleteResponseMatchRequestMethodDelete, RateLimitDeleteResponseMatchRequestMethodPatch, RateLimitDeleteResponseMatchRequestMethodHead, RateLimitDeleteResponseMatchRequestMethod_All:
		return true
	}
	return false
}

type RateLimitDeleteResponseMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic bool                                     `json:"origin_traffic"`
	JSON          rateLimitDeleteResponseMatchResponseJSON `json:"-"`
}

// rateLimitDeleteResponseMatchResponseJSON contains the JSON metadata for the
// struct [RateLimitDeleteResponseMatchResponse]
type rateLimitDeleteResponseMatchResponseJSON struct {
	OriginTraffic apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RateLimitDeleteResponseMatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseMatchResponseJSON) RawJSON() string {
	return r.raw
}

type RateLimitNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RateLimitNewParamsAction] `json:"action,required"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match param.Field[RateLimitNewParamsMatch] `json:"match,required"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period param.Field[float64] `json:"period,required"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold param.Field[float64] `json:"threshold,required"`
}

func (r RateLimitNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitNewParamsAction struct {
	// The action to perform.
	Mode param.Field[RateLimitNewParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RateLimitNewParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RateLimitNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RateLimitNewParamsActionMode string

const (
	RateLimitNewParamsActionModeSimulate         RateLimitNewParamsActionMode = "simulate"
	RateLimitNewParamsActionModeBan              RateLimitNewParamsActionMode = "ban"
	RateLimitNewParamsActionModeChallenge        RateLimitNewParamsActionMode = "challenge"
	RateLimitNewParamsActionModeJSChallenge      RateLimitNewParamsActionMode = "js_challenge"
	RateLimitNewParamsActionModeManagedChallenge RateLimitNewParamsActionMode = "managed_challenge"
)

func (r RateLimitNewParamsActionMode) IsKnown() bool {
	switch r {
	case RateLimitNewParamsActionModeSimulate, RateLimitNewParamsActionModeBan, RateLimitNewParamsActionModeChallenge, RateLimitNewParamsActionModeJSChallenge, RateLimitNewParamsActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RateLimitNewParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RateLimitNewParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitNewParamsMatch struct {
	Headers  param.Field[[]RateLimitNewParamsMatchHeader] `json:"headers"`
	Request  param.Field[RateLimitNewParamsMatchRequest]  `json:"request"`
	Response param.Field[RateLimitNewParamsMatchResponse] `json:"response"`
}

func (r RateLimitNewParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitNewParamsMatchHeader struct {
	// The name of the response header to match.
	Name param.Field[string] `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op param.Field[RateLimitNewParamsMatchHeadersOp] `json:"op"`
	// The value of the response header, which must match exactly.
	Value param.Field[string] `json:"value"`
}

func (r RateLimitNewParamsMatchHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitNewParamsMatchHeadersOp string

const (
	RateLimitNewParamsMatchHeadersOpEq RateLimitNewParamsMatchHeadersOp = "eq"
	RateLimitNewParamsMatchHeadersOpNe RateLimitNewParamsMatchHeadersOp = "ne"
)

func (r RateLimitNewParamsMatchHeadersOp) IsKnown() bool {
	switch r {
	case RateLimitNewParamsMatchHeadersOpEq, RateLimitNewParamsMatchHeadersOpNe:
		return true
	}
	return false
}

type RateLimitNewParamsMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods param.Field[[]RateLimitNewParamsMatchRequestMethod] `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes param.Field[[]string] `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL param.Field[string] `json:"url"`
}

func (r RateLimitNewParamsMatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitNewParamsMatchRequestMethod string

const (
	RateLimitNewParamsMatchRequestMethodGet    RateLimitNewParamsMatchRequestMethod = "GET"
	RateLimitNewParamsMatchRequestMethodPost   RateLimitNewParamsMatchRequestMethod = "POST"
	RateLimitNewParamsMatchRequestMethodPut    RateLimitNewParamsMatchRequestMethod = "PUT"
	RateLimitNewParamsMatchRequestMethodDelete RateLimitNewParamsMatchRequestMethod = "DELETE"
	RateLimitNewParamsMatchRequestMethodPatch  RateLimitNewParamsMatchRequestMethod = "PATCH"
	RateLimitNewParamsMatchRequestMethodHead   RateLimitNewParamsMatchRequestMethod = "HEAD"
	RateLimitNewParamsMatchRequestMethod_All   RateLimitNewParamsMatchRequestMethod = "_ALL_"
)

func (r RateLimitNewParamsMatchRequestMethod) IsKnown() bool {
	switch r {
	case RateLimitNewParamsMatchRequestMethodGet, RateLimitNewParamsMatchRequestMethodPost, RateLimitNewParamsMatchRequestMethodPut, RateLimitNewParamsMatchRequestMethodDelete, RateLimitNewParamsMatchRequestMethodPatch, RateLimitNewParamsMatchRequestMethodHead, RateLimitNewParamsMatchRequestMethod_All:
		return true
	}
	return false
}

type RateLimitNewParamsMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic param.Field[bool] `json:"origin_traffic"`
}

func (r RateLimitNewParamsMatchResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RateLimit             `json:"result,required"`
	// Defines whether the API call was successful.
	Success RateLimitNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitNewResponseEnvelopeJSON    `json:"-"`
}

// rateLimitNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitNewResponseEnvelope]
type rateLimitNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type RateLimitNewResponseEnvelopeSuccess bool

const (
	RateLimitNewResponseEnvelopeSuccessTrue RateLimitNewResponseEnvelopeSuccess = true
)

func (r RateLimitNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RateLimitListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines the page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Defines the maximum number of results per page. You can only set the value to
	// `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RateLimitListParams]'s query parameters as `url.Values`.
func (r RateLimitListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RateLimitDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RateLimitDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   RateLimitDeleteResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success RateLimitDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitDeleteResponseEnvelopeJSON    `json:"-"`
}

// rateLimitDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitDeleteResponseEnvelope]
type rateLimitDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type RateLimitDeleteResponseEnvelopeSuccess bool

const (
	RateLimitDeleteResponseEnvelopeSuccessTrue RateLimitDeleteResponseEnvelopeSuccess = true
)

func (r RateLimitDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RateLimitEditParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RateLimitEditParamsAction] `json:"action,required"`
	// Determines which traffic the rate limit counts towards the threshold.
	Match param.Field[RateLimitEditParamsMatch] `json:"match,required"`
	// The time in seconds (an integer value) to count matching traffic. If the count
	// exceeds the configured threshold within this period, Cloudflare will perform the
	// configured action.
	Period param.Field[float64] `json:"period,required"`
	// The threshold that will trigger the configured mitigation action. Configure this
	// value along with the `period` property to establish a threshold per period.
	Threshold param.Field[float64] `json:"threshold,required"`
}

func (r RateLimitEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RateLimitEditParamsAction struct {
	// The action to perform.
	Mode param.Field[RateLimitEditParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RateLimitEditParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RateLimitEditParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RateLimitEditParamsActionMode string

const (
	RateLimitEditParamsActionModeSimulate         RateLimitEditParamsActionMode = "simulate"
	RateLimitEditParamsActionModeBan              RateLimitEditParamsActionMode = "ban"
	RateLimitEditParamsActionModeChallenge        RateLimitEditParamsActionMode = "challenge"
	RateLimitEditParamsActionModeJSChallenge      RateLimitEditParamsActionMode = "js_challenge"
	RateLimitEditParamsActionModeManagedChallenge RateLimitEditParamsActionMode = "managed_challenge"
)

func (r RateLimitEditParamsActionMode) IsKnown() bool {
	switch r {
	case RateLimitEditParamsActionModeSimulate, RateLimitEditParamsActionModeBan, RateLimitEditParamsActionModeChallenge, RateLimitEditParamsActionModeJSChallenge, RateLimitEditParamsActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RateLimitEditParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RateLimitEditParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which traffic the rate limit counts towards the threshold.
type RateLimitEditParamsMatch struct {
	Headers  param.Field[[]RateLimitEditParamsMatchHeader] `json:"headers"`
	Request  param.Field[RateLimitEditParamsMatchRequest]  `json:"request"`
	Response param.Field[RateLimitEditParamsMatchResponse] `json:"response"`
}

func (r RateLimitEditParamsMatch) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitEditParamsMatchHeader struct {
	// The name of the response header to match.
	Name param.Field[string] `json:"name"`
	// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
	Op param.Field[RateLimitEditParamsMatchHeadersOp] `json:"op"`
	// The value of the response header, which must match exactly.
	Value param.Field[string] `json:"value"`
}

func (r RateLimitEditParamsMatchHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The operator used when matching: `eq` means "equal" and `ne` means "not equal".
type RateLimitEditParamsMatchHeadersOp string

const (
	RateLimitEditParamsMatchHeadersOpEq RateLimitEditParamsMatchHeadersOp = "eq"
	RateLimitEditParamsMatchHeadersOpNe RateLimitEditParamsMatchHeadersOp = "ne"
)

func (r RateLimitEditParamsMatchHeadersOp) IsKnown() bool {
	switch r {
	case RateLimitEditParamsMatchHeadersOpEq, RateLimitEditParamsMatchHeadersOpNe:
		return true
	}
	return false
}

type RateLimitEditParamsMatchRequest struct {
	// The HTTP methods to match. You can specify a subset (for example,
	// `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when
	// creating a rate limit.
	Methods param.Field[[]RateLimitEditParamsMatchRequestMethod] `json:"methods"`
	// The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both
	// schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is
	// optional.
	Schemes param.Field[[]string] `json:"schemes"`
	// The URL pattern to match, composed of a host and a path such as
	// `example.org/path*`. Normalization is applied before the pattern is matched. `*`
	// wildcards are expanded to match applicable traffic. Query strings are not
	// matched. Set the value to `*` to match all traffic to your zone.
	URL param.Field[string] `json:"url"`
}

func (r RateLimitEditParamsMatchRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An HTTP method or `_ALL_` to indicate all methods.
type RateLimitEditParamsMatchRequestMethod string

const (
	RateLimitEditParamsMatchRequestMethodGet    RateLimitEditParamsMatchRequestMethod = "GET"
	RateLimitEditParamsMatchRequestMethodPost   RateLimitEditParamsMatchRequestMethod = "POST"
	RateLimitEditParamsMatchRequestMethodPut    RateLimitEditParamsMatchRequestMethod = "PUT"
	RateLimitEditParamsMatchRequestMethodDelete RateLimitEditParamsMatchRequestMethod = "DELETE"
	RateLimitEditParamsMatchRequestMethodPatch  RateLimitEditParamsMatchRequestMethod = "PATCH"
	RateLimitEditParamsMatchRequestMethodHead   RateLimitEditParamsMatchRequestMethod = "HEAD"
	RateLimitEditParamsMatchRequestMethod_All   RateLimitEditParamsMatchRequestMethod = "_ALL_"
)

func (r RateLimitEditParamsMatchRequestMethod) IsKnown() bool {
	switch r {
	case RateLimitEditParamsMatchRequestMethodGet, RateLimitEditParamsMatchRequestMethodPost, RateLimitEditParamsMatchRequestMethodPut, RateLimitEditParamsMatchRequestMethodDelete, RateLimitEditParamsMatchRequestMethodPatch, RateLimitEditParamsMatchRequestMethodHead, RateLimitEditParamsMatchRequestMethod_All:
		return true
	}
	return false
}

type RateLimitEditParamsMatchResponse struct {
	// When true, only the uncached traffic served from your origin servers will count
	// towards rate limiting. In this case, any cached traffic served by Cloudflare
	// will not count towards rate limiting. This field is optional. Notes: This field
	// is deprecated. Instead, use response headers and set "origin_traffic" to "false"
	// to avoid legacy behaviour interacting with the "response_headers" property.
	OriginTraffic param.Field[bool] `json:"origin_traffic"`
}

func (r RateLimitEditParamsMatchResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RateLimitEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RateLimit             `json:"result,required"`
	// Defines whether the API call was successful.
	Success RateLimitEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitEditResponseEnvelopeJSON    `json:"-"`
}

// rateLimitEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitEditResponseEnvelope]
type rateLimitEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type RateLimitEditResponseEnvelopeSuccess bool

const (
	RateLimitEditResponseEnvelopeSuccessTrue RateLimitEditResponseEnvelopeSuccess = true
)

func (r RateLimitEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RateLimitGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RateLimitGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   RateLimit             `json:"result,required"`
	// Defines whether the API call was successful.
	Success RateLimitGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    rateLimitGetResponseEnvelopeJSON    `json:"-"`
}

// rateLimitGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RateLimitGetResponseEnvelope]
type rateLimitGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RateLimitGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rateLimitGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type RateLimitGetResponseEnvelopeSuccess bool

const (
	RateLimitGetResponseEnvelopeSuccessTrue RateLimitGetResponseEnvelopeSuccess = true
)

func (r RateLimitGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RateLimitGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
