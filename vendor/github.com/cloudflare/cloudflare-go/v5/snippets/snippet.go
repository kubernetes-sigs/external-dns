// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package snippets

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// SnippetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnippetService] method instead.
type SnippetService struct {
	Options []option.RequestOption
	Content *ContentService
	Rules   *RuleService
}

// NewSnippetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSnippetService(opts ...option.RequestOption) (r *SnippetService) {
	r = &SnippetService{}
	r.Options = opts
	r.Content = NewContentService(opts...)
	r.Rules = NewRuleService(opts...)
	return
}

// Creates or updates a snippet belonging to the zone.
func (r *SnippetService) Update(ctx context.Context, snippetName string, params SnippetUpdateParams, opts ...option.RequestOption) (res *SnippetUpdateResponse, err error) {
	var env SnippetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", params.ZoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches all snippets belonging to the zone.
func (r *SnippetService) List(ctx context.Context, params SnippetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SnippetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets", params.ZoneID)
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

// Fetches all snippets belonging to the zone.
func (r *SnippetService) ListAutoPaging(ctx context.Context, params SnippetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SnippetListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a snippet belonging to the zone.
func (r *SnippetService) Delete(ctx context.Context, snippetName string, body SnippetDeleteParams, opts ...option.RequestOption) (res *string, err error) {
	var env SnippetDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", body.ZoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a snippet belonging to the zone.
func (r *SnippetService) Get(ctx context.Context, snippetName string, query SnippetGetParams, opts ...option.RequestOption) (res *SnippetGetResponse, err error) {
	var env SnippetGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if snippetName == "" {
		err = errors.New("missing required snippet_name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/snippets/%s", query.ZoneID, snippetName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type SnippetUpdateResponse struct {
	// The timestamp of when the snippet was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// The timestamp of when the snippet was last modified.
	ModifiedOn time.Time                 `json:"modified_on" format:"date-time"`
	JSON       snippetUpdateResponseJSON `json:"-"`
}

// snippetUpdateResponseJSON contains the JSON metadata for the struct
// [SnippetUpdateResponse]
type snippetUpdateResponseJSON struct {
	CreatedOn   apijson.Field
	SnippetName apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// A snippet object.
type SnippetListResponse struct {
	// The timestamp of when the snippet was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// The timestamp of when the snippet was last modified.
	ModifiedOn time.Time               `json:"modified_on" format:"date-time"`
	JSON       snippetListResponseJSON `json:"-"`
}

// snippetListResponseJSON contains the JSON metadata for the struct
// [SnippetListResponse]
type snippetListResponseJSON struct {
	CreatedOn   apijson.Field
	SnippetName apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetListResponseJSON) RawJSON() string {
	return r.raw
}

// A result.
type SnippetGetResponse struct {
	// The timestamp of when the snippet was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The identifying name of the snippet.
	SnippetName string `json:"snippet_name,required"`
	// The timestamp of when the snippet was last modified.
	ModifiedOn time.Time              `json:"modified_on" format:"date-time"`
	JSON       snippetGetResponseJSON `json:"-"`
}

// snippetGetResponseJSON contains the JSON metadata for the struct
// [SnippetGetResponse]
type snippetGetResponseJSON struct {
	CreatedOn   apijson.Field
	SnippetName apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetGetResponseJSON) RawJSON() string {
	return r.raw
}

type SnippetUpdateParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The list of files belonging to the snippet.
	Files param.Field[[]io.Reader] `json:"files,required" format:"binary"`
	// Metadata about the snippet.
	Metadata param.Field[SnippetUpdateParamsMetadata] `json:"metadata,required"`
}

func (r SnippetUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Metadata about the snippet.
type SnippetUpdateParamsMetadata struct {
	// Name of the file that contains the main module of the snippet.
	MainModule param.Field[string] `json:"main_module,required"`
}

func (r SnippetUpdateParamsMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type SnippetUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []SnippetUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []SnippetUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result SnippetUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SnippetUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetUpdateResponseEnvelopeJSON    `json:"-"`
}

// snippetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetUpdateResponseEnvelope]
type snippetUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type SnippetUpdateResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                   `json:"code"`
	JSON snippetUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SnippetUpdateResponseEnvelopeErrors]
type snippetUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message.
type SnippetUpdateResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                     `json:"code"`
	JSON snippetUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetUpdateResponseEnvelopeMessages]
type snippetUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SnippetUpdateResponseEnvelopeSuccess bool

const (
	SnippetUpdateResponseEnvelopeSuccessTrue SnippetUpdateResponseEnvelopeSuccess = true
)

func (r SnippetUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SnippetUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SnippetListParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The current page number.
	Page param.Field[int64] `query:"page"`
	// The number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [SnippetListParams]'s query parameters as `url.Values`.
func (r SnippetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SnippetDeleteParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// A response object.
type SnippetDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []SnippetDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []SnippetDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result string `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success SnippetDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetDeleteResponseEnvelopeJSON    `json:"-"`
}

// snippetDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetDeleteResponseEnvelope]
type snippetDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type SnippetDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                   `json:"code"`
	JSON snippetDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SnippetDeleteResponseEnvelopeErrors]
type snippetDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message.
type SnippetDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                     `json:"code"`
	JSON snippetDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SnippetDeleteResponseEnvelopeMessages]
type snippetDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SnippetDeleteResponseEnvelopeSuccess bool

const (
	SnippetDeleteResponseEnvelopeSuccessTrue SnippetDeleteResponseEnvelopeSuccess = true
)

func (r SnippetDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SnippetDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SnippetGetParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// A response object.
type SnippetGetResponseEnvelope struct {
	// A list of error messages.
	Errors []SnippetGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []SnippetGetResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result SnippetGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SnippetGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    snippetGetResponseEnvelopeJSON    `json:"-"`
}

// snippetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelope]
type snippetGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type SnippetGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                `json:"code"`
	JSON snippetGetResponseEnvelopeErrorsJSON `json:"-"`
}

// snippetGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelopeErrors]
type snippetGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message.
type SnippetGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64                                  `json:"code"`
	JSON snippetGetResponseEnvelopeMessagesJSON `json:"-"`
}

// snippetGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SnippetGetResponseEnvelopeMessages]
type snippetGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SnippetGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r snippetGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SnippetGetResponseEnvelopeSuccess bool

const (
	SnippetGetResponseEnvelopeSuccessTrue SnippetGetResponseEnvelopeSuccess = true
)

func (r SnippetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SnippetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
