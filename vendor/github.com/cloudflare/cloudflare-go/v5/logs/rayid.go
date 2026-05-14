// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logs

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

// RayIDService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRayIDService] method instead.
type RayIDService struct {
	Options []option.RequestOption
}

// NewRayIDService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRayIDService(opts ...option.RequestOption) (r *RayIDService) {
	r = &RayIDService{}
	r.Options = opts
	return
}

// The `/rayids` api route allows lookups by specific rayid. The rayids route will
// return zero, one, or more records (ray ids are not unique).
func (r *RayIDService) Get(ctx context.Context, RayID string, params RayIDGetParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env apijson.UnionUnmarshaler[interface{}]
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if RayID == "" {
		err = errors.New("missing required ray_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logs/rayids/%s", params.ZoneID, RayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Value
	return
}

type RayIDGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The `/received` route by default returns a limited set of fields, and allows
	// customers to override the default field set by specifying individual fields. The
	// reasons for this are: 1. Most customers require only a small subset of fields,
	// but that subset varies from customer to customer; 2. Flat schema is much easier
	// to work with downstream (importing into BigTable etc); 3. Performance (time to
	// process, file size). If `?fields=` is not specified, default field set is
	// returned. This default field set may change at any time. When `?fields=` is
	// provided, each record is returned with the specified fields. `fields` must be
	// specified as a comma separated list without any whitespaces, and all fields must
	// exist. The order in which fields are specified does not matter, and the order of
	// fields in the response is not specified.
	Fields param.Field[string] `query:"fields"`
	// By default, timestamps in responses are returned as Unix nanosecond integers.
	// The `?timestamps=` argument can be set to change the format in which response
	// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
	// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
	// timestamps as strings.
	Timestamps param.Field[RayIDGetParamsTimestamps] `query:"timestamps"`
}

// URLQuery serializes [RayIDGetParams]'s query parameters as `url.Values`.
func (r RayIDGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// By default, timestamps in responses are returned as Unix nanosecond integers.
// The `?timestamps=` argument can be set to change the format in which response
// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
// timestamps as strings.
type RayIDGetParamsTimestamps string

const (
	RayIDGetParamsTimestampsUnix     RayIDGetParamsTimestamps = "unix"
	RayIDGetParamsTimestampsUnixnano RayIDGetParamsTimestamps = "unixnano"
	RayIDGetParamsTimestampsRfc3339  RayIDGetParamsTimestamps = "rfc3339"
)

func (r RayIDGetParamsTimestamps) IsKnown() bool {
	switch r {
	case RayIDGetParamsTimestampsUnix, RayIDGetParamsTimestampsUnixnano, RayIDGetParamsTimestampsRfc3339:
		return true
	}
	return false
}
