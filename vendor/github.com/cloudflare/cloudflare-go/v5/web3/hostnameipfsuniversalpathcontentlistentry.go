// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// HostnameIPFSUniversalPathContentListEntryService contains methods and other
// services that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameIPFSUniversalPathContentListEntryService] method instead.
type HostnameIPFSUniversalPathContentListEntryService struct {
	Options []option.RequestOption
}

// NewHostnameIPFSUniversalPathContentListEntryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewHostnameIPFSUniversalPathContentListEntryService(opts ...option.RequestOption) (r *HostnameIPFSUniversalPathContentListEntryService) {
	r = &HostnameIPFSUniversalPathContentListEntryService{}
	r.Options = opts
	return
}

// Create IPFS Universal Path Gateway Content List Entry
func (r *HostnameIPFSUniversalPathContentListEntryService) New(ctx context.Context, identifier string, params HostnameIPFSUniversalPathContentListEntryNewParams, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryNewResponse, err error) {
	var env HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", params.ZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit IPFS Universal Path Gateway Content List Entry
func (r *HostnameIPFSUniversalPathContentListEntryService) Update(ctx context.Context, identifier string, contentListEntryIdentifier string, params HostnameIPFSUniversalPathContentListEntryUpdateParams, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryUpdateResponse, err error) {
	var env HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if contentListEntryIdentifier == "" {
		err = errors.New("missing required content_list_entry_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", params.ZoneID, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List IPFS Universal Path Gateway Content List Entries
func (r *HostnameIPFSUniversalPathContentListEntryService) List(ctx context.Context, identifier string, query HostnameIPFSUniversalPathContentListEntryListParams, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryListResponse, err error) {
	var env HostnameIPFSUniversalPathContentListEntryListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries", query.ZoneID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete IPFS Universal Path Gateway Content List Entry
func (r *HostnameIPFSUniversalPathContentListEntryService) Delete(ctx context.Context, identifier string, contentListEntryIdentifier string, body HostnameIPFSUniversalPathContentListEntryDeleteParams, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryDeleteResponse, err error) {
	var env HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if contentListEntryIdentifier == "" {
		err = errors.New("missing required content_list_entry_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", body.ZoneID, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// IPFS Universal Path Gateway Content List Entry Details
func (r *HostnameIPFSUniversalPathContentListEntryService) Get(ctx context.Context, identifier string, contentListEntryIdentifier string, query HostnameIPFSUniversalPathContentListEntryGetParams, opts ...option.RequestOption) (res *HostnameIPFSUniversalPathContentListEntryGetResponse, err error) {
	var env HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if contentListEntryIdentifier == "" {
		err = errors.New("missing required content_list_entry_identifier parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/web3/hostnames/%s/ipfs_universal_path/content_list/entries/%s", query.ZoneID, identifier, contentListEntryIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Specify a content list entry to block.
type HostnameIPFSUniversalPathContentListEntryNewResponse struct {
	// Specify the identifier of the hostname.
	ID string `json:"id"`
	// Specify the CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specify an optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Specify the type of content list entry to block.
	Type HostnameIPFSUniversalPathContentListEntryNewResponseType `json:"type"`
	JSON hostnameIPFSUniversalPathContentListEntryNewResponseJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryNewResponseJSON contains the JSON
// metadata for the struct [HostnameIPFSUniversalPathContentListEntryNewResponse]
type hostnameIPFSUniversalPathContentListEntryNewResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryNewResponseJSON) RawJSON() string {
	return r.raw
}

// Specify the type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryNewResponseType string

const (
	HostnameIPFSUniversalPathContentListEntryNewResponseTypeCid         HostnameIPFSUniversalPathContentListEntryNewResponseType = "cid"
	HostnameIPFSUniversalPathContentListEntryNewResponseTypeContentPath HostnameIPFSUniversalPathContentListEntryNewResponseType = "content_path"
)

func (r HostnameIPFSUniversalPathContentListEntryNewResponseType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryNewResponseTypeCid, HostnameIPFSUniversalPathContentListEntryNewResponseTypeContentPath:
		return true
	}
	return false
}

// Specify a content list entry to block.
type HostnameIPFSUniversalPathContentListEntryUpdateResponse struct {
	// Specify the identifier of the hostname.
	ID string `json:"id"`
	// Specify the CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specify an optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Specify the type of content list entry to block.
	Type HostnameIPFSUniversalPathContentListEntryUpdateResponseType `json:"type"`
	JSON hostnameIPFSUniversalPathContentListEntryUpdateResponseJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryUpdateResponseJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryUpdateResponse]
type hostnameIPFSUniversalPathContentListEntryUpdateResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Specify the type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryUpdateResponseType string

const (
	HostnameIPFSUniversalPathContentListEntryUpdateResponseTypeCid         HostnameIPFSUniversalPathContentListEntryUpdateResponseType = "cid"
	HostnameIPFSUniversalPathContentListEntryUpdateResponseTypeContentPath HostnameIPFSUniversalPathContentListEntryUpdateResponseType = "content_path"
)

func (r HostnameIPFSUniversalPathContentListEntryUpdateResponseType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryUpdateResponseTypeCid, HostnameIPFSUniversalPathContentListEntryUpdateResponseTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryListResponse struct {
	// Provides content list entries.
	Entries []HostnameIPFSUniversalPathContentListEntryListResponseEntry `json:"entries"`
	JSON    hostnameIPFSUniversalPathContentListEntryListResponseJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseJSON contains the JSON
// metadata for the struct [HostnameIPFSUniversalPathContentListEntryListResponse]
type hostnameIPFSUniversalPathContentListEntryListResponseJSON struct {
	Entries     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseJSON) RawJSON() string {
	return r.raw
}

// Specify a content list entry to block.
type HostnameIPFSUniversalPathContentListEntryListResponseEntry struct {
	// Specify the identifier of the hostname.
	ID string `json:"id"`
	// Specify the CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specify an optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Specify the type of content list entry to block.
	Type HostnameIPFSUniversalPathContentListEntryListResponseEntriesType `json:"type"`
	JSON hostnameIPFSUniversalPathContentListEntryListResponseEntryJSON   `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseEntryJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryListResponseEntry]
type hostnameIPFSUniversalPathContentListEntryListResponseEntryJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponseEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseEntryJSON) RawJSON() string {
	return r.raw
}

// Specify the type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryListResponseEntriesType string

const (
	HostnameIPFSUniversalPathContentListEntryListResponseEntriesTypeCid         HostnameIPFSUniversalPathContentListEntryListResponseEntriesType = "cid"
	HostnameIPFSUniversalPathContentListEntryListResponseEntriesTypeContentPath HostnameIPFSUniversalPathContentListEntryListResponseEntriesType = "content_path"
)

func (r HostnameIPFSUniversalPathContentListEntryListResponseEntriesType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryListResponseEntriesTypeCid, HostnameIPFSUniversalPathContentListEntryListResponseEntriesTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryDeleteResponse struct {
	// Specify the identifier of the hostname.
	ID   string                                                      `json:"id,required"`
	JSON hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON contains the JSON
// metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryDeleteResponse]
type hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Specify a content list entry to block.
type HostnameIPFSUniversalPathContentListEntryGetResponse struct {
	// Specify the identifier of the hostname.
	ID string `json:"id"`
	// Specify the CID or content path of content to block.
	Content   string    `json:"content"`
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Specify an optional description of the content list entry.
	Description string    `json:"description"`
	ModifiedOn  time.Time `json:"modified_on" format:"date-time"`
	// Specify the type of content list entry to block.
	Type HostnameIPFSUniversalPathContentListEntryGetResponseType `json:"type"`
	JSON hostnameIPFSUniversalPathContentListEntryGetResponseJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryGetResponseJSON contains the JSON
// metadata for the struct [HostnameIPFSUniversalPathContentListEntryGetResponse]
type hostnameIPFSUniversalPathContentListEntryGetResponseJSON struct {
	ID          apijson.Field
	Content     apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryGetResponseJSON) RawJSON() string {
	return r.raw
}

// Specify the type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryGetResponseType string

const (
	HostnameIPFSUniversalPathContentListEntryGetResponseTypeCid         HostnameIPFSUniversalPathContentListEntryGetResponseType = "cid"
	HostnameIPFSUniversalPathContentListEntryGetResponseTypeContentPath HostnameIPFSUniversalPathContentListEntryGetResponseType = "content_path"
)

func (r HostnameIPFSUniversalPathContentListEntryGetResponseType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryGetResponseTypeCid, HostnameIPFSUniversalPathContentListEntryGetResponseTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryNewParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Specify the CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Specify the type of content list entry to block.
	Type param.Field[HostnameIPFSUniversalPathContentListEntryNewParamsType] `json:"type,required"`
	// Specify an optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r HostnameIPFSUniversalPathContentListEntryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryNewParamsType string

const (
	HostnameIPFSUniversalPathContentListEntryNewParamsTypeCid         HostnameIPFSUniversalPathContentListEntryNewParamsType = "cid"
	HostnameIPFSUniversalPathContentListEntryNewParamsTypeContentPath HostnameIPFSUniversalPathContentListEntryNewParamsType = "content_path"
)

func (r HostnameIPFSUniversalPathContentListEntryNewParamsType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryNewParamsTypeCid, HostnameIPFSUniversalPathContentListEntryNewParamsTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Specify a content list entry to block.
	Result HostnameIPFSUniversalPathContentListEntryNewResponse `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                                                      `json:"result_info"`
	JSON       hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryUpdateParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Specify the CID or content path of content to block.
	Content param.Field[string] `json:"content,required"`
	// Specify the type of content list entry to block.
	Type param.Field[HostnameIPFSUniversalPathContentListEntryUpdateParamsType] `json:"type,required"`
	// Specify an optional description of the content list entry.
	Description param.Field[string] `json:"description"`
}

func (r HostnameIPFSUniversalPathContentListEntryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the type of content list entry to block.
type HostnameIPFSUniversalPathContentListEntryUpdateParamsType string

const (
	HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeCid         HostnameIPFSUniversalPathContentListEntryUpdateParamsType = "cid"
	HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeContentPath HostnameIPFSUniversalPathContentListEntryUpdateParamsType = "content_path"
)

func (r HostnameIPFSUniversalPathContentListEntryUpdateParamsType) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeCid, HostnameIPFSUniversalPathContentListEntryUpdateParamsTypeContentPath:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Specify a content list entry to block.
	Result HostnameIPFSUniversalPathContentListEntryUpdateResponse `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                                                         `json:"result_info"`
	JSON       hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryListParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameIPFSUniversalPathContentListEntryListResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                 `json:"errors,required"`
	Messages []shared.ResponseInfo                                 `json:"messages,required"`
	Result   HostnameIPFSUniversalPathContentListEntryListResponse `json:"result,required,nullable"`
	// Specifies whether the API call was successful.
	Success    HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON       `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryListResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo struct {
	// Specifies the total number of results for the requested service.
	Count float64 `json:"count"`
	// Specifies the current page within paginated list of results.
	Page float64 `json:"page"`
	// Specifies the number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Specifies the total results available without any search parameters.
	TotalCount float64                                                                     `json:"total_count"`
	JSON       hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo]
type hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type HostnameIPFSUniversalPathContentListEntryDeleteParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                                   `json:"errors,required"`
	Messages []shared.ResponseInfo                                   `json:"messages,required"`
	Result   HostnameIPFSUniversalPathContentListEntryDeleteResponse `json:"result,required,nullable"`
	// Specifies whether the API call was successful.
	Success HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON    `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameIPFSUniversalPathContentListEntryGetParams struct {
	// Specify the identifier of the hostname.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Specify a content list entry to block.
	Result HostnameIPFSUniversalPathContentListEntryGetResponse `json:"result,required"`
	// Specifies whether the API call was successful.
	Success HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess `json:"success,required"`
	// Provides the API response.
	ResultInfo interface{}                                                      `json:"result_info"`
	JSON       hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON `json:"-"`
}

// hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope]
type hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameIPFSUniversalPathContentListEntryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Specifies whether the API call was successful.
type HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess bool

const (
	HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccessTrue HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess = true
)

func (r HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameIPFSUniversalPathContentListEntryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
