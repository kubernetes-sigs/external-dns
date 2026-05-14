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

// ReceivedService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReceivedService] method instead.
type ReceivedService struct {
	Options []option.RequestOption
	Fields  *ReceivedFieldService
}

// NewReceivedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReceivedService(opts ...option.RequestOption) (r *ReceivedService) {
	r = &ReceivedService{}
	r.Options = opts
	r.Fields = NewReceivedFieldService(opts...)
	return
}

// The `/received` api route allows customers to retrieve their edge HTTP logs. The
// basic access pattern is "give me all the logs for zone Z for minute M", where
// the minute M refers to the time records were received at Cloudflare's central
// data center. `start` is inclusive, and `end` is exclusive. Because of that, to
// get all data, at minutely cadence, starting at 10AM, the proper values are:
// `start=2018-05-20T10:00:00Z&end=2018-05-20T10:01:00Z`, then
// `start=2018-05-20T10:01:00Z&end=2018-05-20T10:02:00Z` and so on; the overlap
// will be handled properly.
func (r *ReceivedService) Get(ctx context.Context, params ReceivedGetParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env apijson.UnionUnmarshaler[interface{}]
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/logs/received", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Value
	return
}

type ReceivedGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Sets the (exclusive) end of the requested time frame. This can be a unix
	// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
	// RFC 3339. `end` must be at least five minutes earlier than now and must be later
	// than `start`. Difference between `start` and `end` must be not greater than one
	// hour.
	End param.Field[ReceivedGetParamsEndUnion] `query:"end,required"`
	// When `?count=` is provided, the response will contain up to `count` results.
	// Since results are not sorted, you are likely to get different data for repeated
	// requests. `count` must be an integer > 0.
	Count param.Field[int64] `query:"count"`
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
	// When `?sample=` is provided, a sample of matching records is returned. If
	// `sample=0.1` then 10% of records will be returned. Sampling is random: repeated
	// calls will not only return different records, but likely will also vary slightly
	// in number of returned records. When `?count=` is also specified, `count` is
	// applied to the number of returned records, not the sampled records. So, with
	// `sample=0.05` and `count=7`, when there is a total of 100 records available,
	// approximately five will be returned. When there are 1000 records, seven will be
	// returned. When there are 10,000 records, seven will be returned.
	Sample param.Field[float64] `query:"sample"`
	// Sets the (inclusive) beginning of the requested time frame. This can be a unix
	// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
	// RFC 3339. At this point in time, it cannot exceed a time in the past greater
	// than seven days.
	Start param.Field[ReceivedGetParamsStartUnion] `query:"start"`
	// By default, timestamps in responses are returned as Unix nanosecond integers.
	// The `?timestamps=` argument can be set to change the format in which response
	// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
	// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
	// timestamps as strings.
	Timestamps param.Field[ReceivedGetParamsTimestamps] `query:"timestamps"`
}

// URLQuery serializes [ReceivedGetParams]'s query parameters as `url.Values`.
func (r ReceivedGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sets the (exclusive) end of the requested time frame. This can be a unix
// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
// RFC 3339. `end` must be at least five minutes earlier than now and must be later
// than `start`. Difference between `start` and `end` must be not greater than one
// hour.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ReceivedGetParamsEndUnion interface {
	ImplementsReceivedGetParamsEndUnion()
}

// Sets the (inclusive) beginning of the requested time frame. This can be a unix
// timestamp (in seconds or nanoseconds), or an absolute timestamp that conforms to
// RFC 3339. At this point in time, it cannot exceed a time in the past greater
// than seven days.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ReceivedGetParamsStartUnion interface {
	ImplementsReceivedGetParamsStartUnion()
}

// By default, timestamps in responses are returned as Unix nanosecond integers.
// The `?timestamps=` argument can be set to change the format in which response
// timestamps are returned. Possible values are: `unix`, `unixnano`, `rfc3339`.
// Note that `unix` and `unixnano` return timestamps as integers; `rfc3339` returns
// timestamps as strings.
type ReceivedGetParamsTimestamps string

const (
	ReceivedGetParamsTimestampsUnix     ReceivedGetParamsTimestamps = "unix"
	ReceivedGetParamsTimestampsUnixnano ReceivedGetParamsTimestamps = "unixnano"
	ReceivedGetParamsTimestampsRfc3339  ReceivedGetParamsTimestamps = "rfc3339"
)

func (r ReceivedGetParamsTimestamps) IsKnown() bool {
	switch r {
	case ReceivedGetParamsTimestampsUnix, ReceivedGetParamsTimestampsUnixnano, ReceivedGetParamsTimestampsRfc3339:
		return true
	}
	return false
}
