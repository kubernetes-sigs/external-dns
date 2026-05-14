// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DownloadService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDownloadService] method instead.
type DownloadService struct {
	Options []option.RequestOption
}

// NewDownloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDownloadService(opts ...option.RequestOption) (r *DownloadService) {
	r = &DownloadService{}
	r.Options = opts
	return
}

// Creates a download for a video when a video is ready to view.
func (r *DownloadService) New(ctx context.Context, identifier string, params DownloadNewParams, opts ...option.RequestOption) (res *DownloadNewResponse, err error) {
	var env DownloadNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", params.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete the downloads for a video.
func (r *DownloadService) Delete(ctx context.Context, identifier string, body DownloadDeleteParams, opts ...option.RequestOption) (res *string, err error) {
	var env DownloadDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", body.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists the downloads created for a video.
func (r *DownloadService) Get(ctx context.Context, identifier string, query DownloadGetParams, opts ...option.RequestOption) (res *DownloadGetResponse, err error) {
	var env DownloadGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/downloads", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DownloadNewResponse = interface{}

type DownloadGetResponse = interface{}

type DownloadNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r DownloadNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DownloadNewResponseEnvelope struct {
	Errors   []DownloadNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DownloadNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DownloadNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DownloadNewResponse                `json:"result"`
	JSON    downloadNewResponseEnvelopeJSON    `json:"-"`
}

// downloadNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DownloadNewResponseEnvelope]
type downloadNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DownloadNewResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DownloadNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             downloadNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// downloadNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DownloadNewResponseEnvelopeErrors]
type downloadNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DownloadNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DownloadNewResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    downloadNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// downloadNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DownloadNewResponseEnvelopeErrorsSource]
type downloadNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DownloadNewResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DownloadNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             downloadNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// downloadNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DownloadNewResponseEnvelopeMessages]
type downloadNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DownloadNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DownloadNewResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    downloadNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// downloadNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DownloadNewResponseEnvelopeMessagesSource]
type downloadNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DownloadNewResponseEnvelopeSuccess bool

const (
	DownloadNewResponseEnvelopeSuccessTrue DownloadNewResponseEnvelopeSuccess = true
)

func (r DownloadNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DownloadNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DownloadDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DownloadDeleteResponseEnvelope struct {
	Errors   []DownloadDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DownloadDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DownloadDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  string                                `json:"result"`
	JSON    downloadDeleteResponseEnvelopeJSON    `json:"-"`
}

// downloadDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DownloadDeleteResponseEnvelope]
type downloadDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DownloadDeleteResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DownloadDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             downloadDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// downloadDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DownloadDeleteResponseEnvelopeErrors]
type downloadDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DownloadDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DownloadDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    downloadDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// downloadDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DownloadDeleteResponseEnvelopeErrorsSource]
type downloadDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DownloadDeleteResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DownloadDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             downloadDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// downloadDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DownloadDeleteResponseEnvelopeMessages]
type downloadDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DownloadDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DownloadDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    downloadDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// downloadDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DownloadDeleteResponseEnvelopeMessagesSource]
type downloadDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DownloadDeleteResponseEnvelopeSuccess bool

const (
	DownloadDeleteResponseEnvelopeSuccessTrue DownloadDeleteResponseEnvelopeSuccess = true
)

func (r DownloadDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DownloadDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DownloadGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DownloadGetResponseEnvelope struct {
	Errors   []DownloadGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DownloadGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DownloadGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DownloadGetResponse                `json:"result"`
	JSON    downloadGetResponseEnvelopeJSON    `json:"-"`
}

// downloadGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DownloadGetResponseEnvelope]
type downloadGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DownloadGetResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           DownloadGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             downloadGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// downloadGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DownloadGetResponseEnvelopeErrors]
type downloadGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DownloadGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DownloadGetResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    downloadGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// downloadGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DownloadGetResponseEnvelopeErrorsSource]
type downloadGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DownloadGetResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DownloadGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             downloadGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// downloadGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DownloadGetResponseEnvelopeMessages]
type downloadGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DownloadGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DownloadGetResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    downloadGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// downloadGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DownloadGetResponseEnvelopeMessagesSource]
type downloadGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DownloadGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r downloadGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DownloadGetResponseEnvelopeSuccess bool

const (
	DownloadGetResponseEnvelopeSuccessTrue DownloadGetResponseEnvelopeSuccess = true
)

func (r DownloadGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DownloadGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
