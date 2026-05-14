// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3

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

// HostnameIPFSUniversalPathContentListService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameIPFSUniversalPathContentListService] method instead.
type HostnameIPFSUniversalPathContentListService struct {
	Options []option.RequestOption
	Entries *HostnameIPFSUniversalPathContentListEntryService
}

// NewHostnameIPFSUniversalPathContentListService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewHostnameIPFSUniversalPathContentListService(opts ...option.RequestOption) (r *HostnameIPFSUniversalPathContentListService) {
	r = &HostnameIPFSUniversalPathContentListService{}
	r.Options = opts
	r.Entries = NewHostnameIPFSUniversalPathContentListEntryService(opts...)
	return
}

// Update IPFS Universal Path Gateway Content List
func (r *HostnameIPFSUniversalPathContentListService) Update(ctx context.Context, identifier string, params HostnameIPFSUniversalPathContentListUpdateParams, opts ...option.RequestOption) (res *ContentList, err error) {
	var env HostnameIPFSUniversalPathContentListUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", params.ZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// IPFS Universal Path Gateway Content List Details
func (r *HostnameIPFSUniversalPathContentListService) Get(ctx context.Context, identifier string, query HostnameIPFSUniversalPathContentListGetParams, opts ...option.RequestOption) (res *ContentList, err error) {
	var env HostnameIPFSUniversalPathContentListGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list", query.ZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ContentList struct {
	// Behavior of the content list.
	Action ContentListAction `json:"action"`
	JSON   contentListJSON   `json:"-"`
}

// contentListJSON contains the JSON metadata for the struct [ContentList]
type contentListJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentListJSON) RawJSON() string {
	return r.raw
}

// Behavior of the content list.
type ContentListAction string

const (
	ContentListActionBlock ContentListAction = "block"
)

func (r ContentListAction) IsKnown() bool {
	switch r {
	case ContentListActionBlock:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListUpdateParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Behavior of the content list.
	Action param.Field[HostnameIPFSUniversalPathContentListUpdateParamsAction] `json:"action,required"`
	// Provides content list entries.
	Entries param.Field[[]HostnameIPFSUniversalPathContentListUpdateParamsEntry] `json:"entries,required"`
}

func (r HostnameIPFSUniversalPathContentListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Behavior of the content list.
type HostnameIPFSUniversalPathContentListUpdateParamsAction string

const (
	HostnameIPFSUniversalPathContentListUpdateParamsActionBlock HostnameIPFSUniversalPathContentListUpdateParamsAction = "block"
)

func (r HostnameIPFSUniversalPathContentListUpdateParamsAction) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListUpdateParamsActionBlock:
		return true
	}
	return false
}

// Specify a content list entry to block.
type HostnameIPFSUniversalPathContentListUpdateParamsEntry struct {
	// Specify the CID or content path of content to block.
	Content param.Field[string] `json:"content"`
	// Specify an optional description of the content list entry.
	Description param.Field[string] `json:"description"`
	// Specify the type of content list entry to block.
	Type param.Field[HostnameIPFSUniversalPathContentListUpdateParamsEntriesType] `json:"type"`
}

func (r HostnameIPFSUniversalPathContentListUpdateParamsEntry) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the type of content list entry to block.
type HostnameIPFSUniversalPathContentListUpdateParamsEntriesType string

const (
	HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid         HostnameIPFSUniversalPathContentListUpdateParamsEntriesType = "cid"
	HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeContentPath HostnameIPFSUniversalPathContentListUpdateParamsEntriesType = "content_path"
)

func (r HostnameIPFSUniversalPathContentListUpdateParamsEntriesType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeCid, HostnameIPFSUniversalPathContentListUpdateParamsEntriesTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ContentList           `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                                                    `json:"result_info"`
	JSON       hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListUpdateResponseEnvelope]
type hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListGetParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameIPFSUniversalPathContentListGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ContentList           `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                                                 `json:"result_info"`
	JSON       hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListGetResponseEnvelope]
type hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
