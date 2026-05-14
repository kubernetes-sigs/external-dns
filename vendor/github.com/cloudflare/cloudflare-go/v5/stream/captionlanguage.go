// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// CaptionLanguageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCaptionLanguageService] method instead.
type CaptionLanguageService struct {
	Options []option.RequestOption
	Vtt     *CaptionLanguageVttService
}

// NewCaptionLanguageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCaptionLanguageService(opts ...option.RequestOption) (r *CaptionLanguageService) {
	r = &CaptionLanguageService{}
	r.Options = opts
	r.Vtt = NewCaptionLanguageVttService(opts...)
	return
}

// Generate captions or subtitles for provided language via AI.
func (r *CaptionLanguageService) New(ctx context.Context, identifier string, language string, body CaptionLanguageNewParams, opts ...option.RequestOption) (res *Caption, err error) {
	var env CaptionLanguageNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s/generate", body.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Uploads the caption or subtitle file to the endpoint for a specific BCP47
// language. One caption or subtitle file per language is allowed.
func (r *CaptionLanguageService) Update(ctx context.Context, identifier string, language string, params CaptionLanguageUpdateParams, opts ...option.RequestOption) (res *Caption, err error) {
	var env CaptionLanguageUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", params.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Removes the captions or subtitles from a video.
func (r *CaptionLanguageService) Delete(ctx context.Context, identifier string, language string, body CaptionLanguageDeleteParams, opts ...option.RequestOption) (res *string, err error) {
	var env CaptionLanguageDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", body.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the captions or subtitles for provided language.
func (r *CaptionLanguageService) Get(ctx context.Context, identifier string, language string, query CaptionLanguageGetParams, opts ...option.RequestOption) (res *Caption, err error) {
	var env CaptionLanguageGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s", query.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CaptionLanguageNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CaptionLanguageNewResponseEnvelope struct {
	Errors   []CaptionLanguageNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CaptionLanguageNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CaptionLanguageNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Caption                                   `json:"result"`
	JSON    captionLanguageNewResponseEnvelopeJSON    `json:"-"`
}

// captionLanguageNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [CaptionLanguageNewResponseEnvelope]
type captionLanguageNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           CaptionLanguageNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             captionLanguageNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// captionLanguageNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CaptionLanguageNewResponseEnvelopeErrors]
type captionLanguageNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    captionLanguageNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// captionLanguageNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CaptionLanguageNewResponseEnvelopeErrorsSource]
type captionLanguageNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           CaptionLanguageNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             captionLanguageNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// captionLanguageNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CaptionLanguageNewResponseEnvelopeMessages]
type captionLanguageNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    captionLanguageNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// captionLanguageNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CaptionLanguageNewResponseEnvelopeMessagesSource]
type captionLanguageNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CaptionLanguageNewResponseEnvelopeSuccess bool

const (
	CaptionLanguageNewResponseEnvelopeSuccessTrue CaptionLanguageNewResponseEnvelopeSuccess = true
)

func (r CaptionLanguageNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionLanguageNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CaptionLanguageUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The WebVTT file containing the caption or subtitle content.
	File param.Field[string] `json:"file,required"`
}

func (r CaptionLanguageUpdateParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type CaptionLanguageUpdateResponseEnvelope struct {
	Errors   []CaptionLanguageUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CaptionLanguageUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CaptionLanguageUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Caption                                      `json:"result"`
	JSON    captionLanguageUpdateResponseEnvelopeJSON    `json:"-"`
}

// captionLanguageUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [CaptionLanguageUpdateResponseEnvelope]
type captionLanguageUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageUpdateResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           CaptionLanguageUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             captionLanguageUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// captionLanguageUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CaptionLanguageUpdateResponseEnvelopeErrors]
type captionLanguageUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    captionLanguageUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// captionLanguageUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CaptionLanguageUpdateResponseEnvelopeErrorsSource]
type captionLanguageUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageUpdateResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           CaptionLanguageUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             captionLanguageUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// captionLanguageUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CaptionLanguageUpdateResponseEnvelopeMessages]
type captionLanguageUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    captionLanguageUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// captionLanguageUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CaptionLanguageUpdateResponseEnvelopeMessagesSource]
type captionLanguageUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CaptionLanguageUpdateResponseEnvelopeSuccess bool

const (
	CaptionLanguageUpdateResponseEnvelopeSuccessTrue CaptionLanguageUpdateResponseEnvelopeSuccess = true
)

func (r CaptionLanguageUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionLanguageUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CaptionLanguageDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CaptionLanguageDeleteResponseEnvelope struct {
	Errors   []CaptionLanguageDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CaptionLanguageDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CaptionLanguageDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  string                                       `json:"result"`
	JSON    captionLanguageDeleteResponseEnvelopeJSON    `json:"-"`
}

// captionLanguageDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [CaptionLanguageDeleteResponseEnvelope]
type captionLanguageDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           CaptionLanguageDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             captionLanguageDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// captionLanguageDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [CaptionLanguageDeleteResponseEnvelopeErrors]
type captionLanguageDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    captionLanguageDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// captionLanguageDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CaptionLanguageDeleteResponseEnvelopeErrorsSource]
type captionLanguageDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           CaptionLanguageDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             captionLanguageDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// captionLanguageDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CaptionLanguageDeleteResponseEnvelopeMessages]
type captionLanguageDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    captionLanguageDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// captionLanguageDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [CaptionLanguageDeleteResponseEnvelopeMessagesSource]
type captionLanguageDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CaptionLanguageDeleteResponseEnvelopeSuccess bool

const (
	CaptionLanguageDeleteResponseEnvelopeSuccessTrue CaptionLanguageDeleteResponseEnvelopeSuccess = true
)

func (r CaptionLanguageDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionLanguageDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CaptionLanguageGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CaptionLanguageGetResponseEnvelope struct {
	Errors   []CaptionLanguageGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CaptionLanguageGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CaptionLanguageGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Caption                                   `json:"result"`
	JSON    captionLanguageGetResponseEnvelopeJSON    `json:"-"`
}

// captionLanguageGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [CaptionLanguageGetResponseEnvelope]
type captionLanguageGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           CaptionLanguageGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             captionLanguageGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// captionLanguageGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CaptionLanguageGetResponseEnvelopeErrors]
type captionLanguageGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    captionLanguageGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// captionLanguageGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CaptionLanguageGetResponseEnvelopeErrorsSource]
type captionLanguageGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           CaptionLanguageGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             captionLanguageGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// captionLanguageGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CaptionLanguageGetResponseEnvelopeMessages]
type captionLanguageGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CaptionLanguageGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CaptionLanguageGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    captionLanguageGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// captionLanguageGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CaptionLanguageGetResponseEnvelopeMessagesSource]
type captionLanguageGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CaptionLanguageGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r captionLanguageGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CaptionLanguageGetResponseEnvelopeSuccess bool

const (
	CaptionLanguageGetResponseEnvelopeSuccessTrue CaptionLanguageGetResponseEnvelopeSuccess = true
)

func (r CaptionLanguageGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CaptionLanguageGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
