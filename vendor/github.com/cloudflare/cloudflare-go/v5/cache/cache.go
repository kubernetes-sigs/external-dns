// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// CacheService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCacheService] method instead.
type CacheService struct {
	Options             []option.RequestOption
	CacheReserve        *CacheReserveService
	SmartTieredCache    *SmartTieredCacheService
	Variants            *VariantService
	RegionalTieredCache *RegionalTieredCacheService
}

// NewCacheService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCacheService(opts ...option.RequestOption) (r *CacheService) {
	r = &CacheService{}
	r.Options = opts
	r.CacheReserve = NewCacheReserveService(opts...)
	r.SmartTieredCache = NewSmartTieredCacheService(opts...)
	r.Variants = NewVariantService(opts...)
	r.RegionalTieredCache = NewRegionalTieredCacheService(opts...)
	return
}

// ### Purge All Cached Content
//
// Removes ALL files from Cloudflare's cache. All tiers can purge everything.
//
// ```
// {"purge_everything": true}
// ```
//
// ### Purge Cached Content by URL
//
// Granularly removes one or more files from Cloudflare's cache by specifying URLs.
// All tiers can purge by URL.
//
// To purge files with custom cache keys, include the headers used to compute the
// cache key as in the example. If you have a device type or geo in your cache key,
// you will need to include the CF-Device-Type or CF-IPCountry headers. If you have
// lang in your cache key, you will need to include the Accept-Language header.
//
// **NB:** When including the Origin header, be sure to include the **scheme** and
// **hostname**. The port number can be omitted if it is the default port (80 for
// http, 443 for https), but must be included otherwise.
//
// Single file purge example with files:
//
// ```
// {"files": ["http://www.example.com/css/styles.css", "http://www.example.com/js/index.js"]}
// ```
//
// Single file purge example with url and header pairs:
//
// ```
// {"files": [{url: "http://www.example.com/cat_picture.jpg", headers: { "CF-IPCountry": "US", "CF-Device-Type": "desktop", "Accept-Language": "zh-CN" }}, {url: "http://www.example.com/dog_picture.jpg", headers: { "CF-IPCountry": "EU", "CF-Device-Type": "mobile", "Accept-Language": "en-US" }}]}
// ```
//
// ### Purge Cached Content by Tag, Host or Prefix
//
// Granularly removes one or more files from Cloudflare's cache either by
// specifying the host, the associated Cache-Tag, or a Prefix.
//
// Flex purge with tags:
//
// ```
// {"tags": ["a-cache-tag", "another-cache-tag"]}
// ```
//
// Flex purge with hosts:
//
// ```
// {"hosts": ["www.example.com", "images.example.com"]}
// ```
//
// Flex purge with prefixes:
//
// ```
// {"prefixes": ["www.example.com/foo", "images.example.com/bar/baz"]}
// ```
//
// ### Availability and limits
//
// please refer to
// [purge cache availability and limits documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/#availability-and-limits).
func (r *CacheService) Purge(ctx context.Context, params CachePurgeParams, opts ...option.RequestOption) (res *CachePurgeResponse, err error) {
	var env CachePurgeResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/purge_cache", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CachePurgeResponse struct {
	ID   string                 `json:"id,required"`
	JSON cachePurgeResponseJSON `json:"-"`
}

// cachePurgeResponseJSON contains the JSON metadata for the struct
// [CachePurgeResponse]
type cachePurgeResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeResponseJSON) RawJSON() string {
	return r.raw
}

type CachePurgeParams struct {
	ZoneID param.Field[string]       `path:"zone_id,required"`
	Body   CachePurgeParamsBodyUnion `json:"body,required"`
}

func (r CachePurgeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type CachePurgeParamsBody struct {
	Files    param.Field[interface{}] `json:"files"`
	Hosts    param.Field[interface{}] `json:"hosts"`
	Prefixes param.Field[interface{}] `json:"prefixes"`
	// For more information, please refer to
	// [purge everything documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool]        `json:"purge_everything"`
	Tags            param.Field[interface{}] `json:"tags"`
}

func (r CachePurgeParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBody) implementsCachePurgeParamsBodyUnion() {}

// Satisfied by [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByTags],
// [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames],
// [cache.CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes],
// [cache.CachePurgeParamsBodyCachePurgeEverything],
// [cache.CachePurgeParamsBodyCachePurgeSingleFile],
// [cache.CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders],
// [CachePurgeParamsBody].
type CachePurgeParamsBodyUnion interface {
	implementsCachePurgeParamsBodyUnion()
}

type CachePurgeParamsBodyCachePurgeFlexPurgeByTags struct {
	// For more information on cache tags and purging by tags, please refer to
	// [purge by cache-tags documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-tags/).
	Tags param.Field[[]string] `json:"tags"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByTags) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByTags) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames struct {
	// For more information purging by hostnames, please refer to
	// [purge by hostname documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-hostname/).
	Hosts param.Field[[]string] `json:"hosts"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByHostnames) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes struct {
	// For more information on purging by prefixes, please refer to
	// [purge by prefix documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge_by_prefix/).
	Prefixes param.Field[[]string] `json:"prefixes"`
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeFlexPurgeByPrefixes) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeEverything struct {
	// For more information, please refer to
	// [purge everything documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-everything/).
	PurgeEverything param.Field[bool] `json:"purge_everything"`
}

func (r CachePurgeParamsBodyCachePurgeEverything) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeEverything) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeSingleFile struct {
	// For more information on purging files, please refer to
	// [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]string] `json:"files"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeSingleFile) implementsCachePurgeParamsBodyUnion() {}

type CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders struct {
	// For more information on purging files with URL and headers, please refer to
	// [purge by single-file documentation page](https://developers.cloudflare.com/cache/how-to/purge-cache/purge-by-single-file/).
	Files param.Field[[]CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile] `json:"files"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeaders) implementsCachePurgeParamsBodyUnion() {
}

type CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile struct {
	Headers param.Field[map[string]string] `json:"headers"`
	URL     param.Field[string]            `json:"url"`
}

func (r CachePurgeParamsBodyCachePurgeSingleFileWithURLAndHeadersFile) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CachePurgeResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Indicates the API call's success or failure.
	Success bool                           `json:"success,required"`
	Result  CachePurgeResponse             `json:"result,nullable"`
	JSON    cachePurgeResponseEnvelopeJSON `json:"-"`
}

// cachePurgeResponseEnvelopeJSON contains the JSON metadata for the struct
// [CachePurgeResponseEnvelope]
type cachePurgeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CachePurgeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePurgeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
