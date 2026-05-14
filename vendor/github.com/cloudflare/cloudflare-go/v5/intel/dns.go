// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options []option.RequestOption
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	return
}

// Gets a list of all the domains that have resolved to a specific IP address.
func (r *DNSService) List(ctx context.Context, params DNSListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[DNS], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/dns", params.AccountID)
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

// Gets a list of all the domains that have resolved to a specific IP address.
func (r *DNSService) ListAutoPaging(ctx context.Context, params DNSListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[DNS] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type DNS struct {
	// Total results returned based on your search parameters.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Reverse DNS look-ups observed during the time period.
	ReverseRecords []DNSReverseRecord `json:"reverse_records"`
	JSON           dnsJSON            `json:"-"`
}

// dnsJSON contains the JSON metadata for the struct [DNS]
type dnsJSON struct {
	Count          apijson.Field
	Page           apijson.Field
	PerPage        apijson.Field
	ReverseRecords apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsJSON) RawJSON() string {
	return r.raw
}

type DNSReverseRecord struct {
	// First seen date of the DNS record during the time period.
	FirstSeen time.Time `json:"first_seen" format:"date"`
	// Hostname that the IP was observed resolving to.
	Hostname string `json:"hostname"`
	// Last seen date of the DNS record during the time period.
	LastSeen time.Time            `json:"last_seen" format:"date"`
	JSON     dnsReverseRecordJSON `json:"-"`
}

// dnsReverseRecordJSON contains the JSON metadata for the struct
// [DNSReverseRecord]
type dnsReverseRecordJSON struct {
	FirstSeen   apijson.Field
	Hostname    apijson.Field
	LastSeen    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSReverseRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsReverseRecordJSON) RawJSON() string {
	return r.raw
}

type DNSListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	IPV4      param.Field[string] `query:"ipv4"`
	// Requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results requested.
	PerPage        param.Field[float64]                     `query:"per_page"`
	StartEndParams param.Field[DNSListParamsStartEndParams] `query:"start_end_params"`
}

// URLQuery serializes [DNSListParams]'s query parameters as `url.Values`.
func (r DNSListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DNSListParamsStartEndParams struct {
	// Defaults to the current date.
	End param.Field[time.Time] `query:"end" format:"date"`
	// Defaults to 30 days before the end parameter value.
	Start param.Field[time.Time] `query:"start" format:"date"`
}

// URLQuery serializes [DNSListParamsStartEndParams]'s query parameters as
// `url.Values`.
func (r DNSListParamsStartEndParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
