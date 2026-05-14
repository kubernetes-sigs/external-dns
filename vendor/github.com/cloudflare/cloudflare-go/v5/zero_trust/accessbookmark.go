// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// AccessBookmarkService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessBookmarkService] method instead.
type AccessBookmarkService struct {
	Options []option.RequestOption
}

// NewAccessBookmarkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessBookmarkService(opts ...option.RequestOption) (r *AccessBookmarkService) {
	r = &AccessBookmarkService{}
	r.Options = opts
	return
}

// Create a new Bookmark application.
//
// Deprecated: deprecated
func (r *AccessBookmarkService) New(ctx context.Context, bookmarkID string, params AccessBookmarkNewParams, opts ...option.RequestOption) (res *Bookmark, err error) {
	var env AccessBookmarkNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", params.AccountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Bookmark application.
//
// Deprecated: deprecated
func (r *AccessBookmarkService) Update(ctx context.Context, bookmarkID string, params AccessBookmarkUpdateParams, opts ...option.RequestOption) (res *Bookmark, err error) {
	var env AccessBookmarkUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", params.AccountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Bookmark applications.
//
// Deprecated: deprecated
func (r *AccessBookmarkService) List(ctx context.Context, query AccessBookmarkListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Bookmark], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists Bookmark applications.
//
// Deprecated: deprecated
func (r *AccessBookmarkService) ListAutoPaging(ctx context.Context, query AccessBookmarkListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Bookmark] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Bookmark application.
//
// Deprecated: deprecated
func (r *AccessBookmarkService) Delete(ctx context.Context, bookmarkID string, body AccessBookmarkDeleteParams, opts ...option.RequestOption) (res *AccessBookmarkDeleteResponse, err error) {
	var env AccessBookmarkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", body.AccountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Bookmark application.
//
// Deprecated: deprecated
func (r *AccessBookmarkService) Get(ctx context.Context, bookmarkID string, query AccessBookmarkGetParams, opts ...option.RequestOption) (res *Bookmark, err error) {
	var env AccessBookmarkGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if bookmarkID == "" {
		err = errors.New("missing required bookmark_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/bookmarks/%s", query.AccountID, bookmarkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Bookmark struct {
	// The unique identifier for the Bookmark application.
	ID string `json:"id"`
	// Displays the application in the App Launcher.
	AppLauncherVisible bool      `json:"app_launcher_visible"`
	CreatedAt          time.Time `json:"created_at" format:"date-time"`
	// The domain of the Bookmark application.
	Domain string `json:"domain"`
	// The image URL for the logo shown in the App Launcher dashboard.
	LogoURL string `json:"logo_url"`
	// The name of the Bookmark application.
	Name      string       `json:"name"`
	UpdatedAt time.Time    `json:"updated_at" format:"date-time"`
	JSON      bookmarkJSON `json:"-"`
}

// bookmarkJSON contains the JSON metadata for the struct [Bookmark]
type bookmarkJSON struct {
	ID                 apijson.Field
	AppLauncherVisible apijson.Field
	CreatedAt          apijson.Field
	Domain             apijson.Field
	LogoURL            apijson.Field
	Name               apijson.Field
	UpdatedAt          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Bookmark) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bookmarkJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponse struct {
	// UUID.
	ID   string                           `json:"id"`
	JSON accessBookmarkDeleteResponseJSON `json:"-"`
}

// accessBookmarkDeleteResponseJSON contains the JSON metadata for the struct
// [AccessBookmarkDeleteResponse]
type accessBookmarkDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AccessBookmarkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessBookmarkNewResponseEnvelope struct {
	Errors   []AccessBookmarkNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessBookmarkNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Bookmark                                 `json:"result"`
	JSON    accessBookmarkNewResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessBookmarkNewResponseEnvelope]
type accessBookmarkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccessBookmarkNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessBookmarkNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessBookmarkNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessBookmarkNewResponseEnvelopeErrors]
type accessBookmarkNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accessBookmarkNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessBookmarkNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessBookmarkNewResponseEnvelopeErrorsSource]
type accessBookmarkNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccessBookmarkNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessBookmarkNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessBookmarkNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessBookmarkNewResponseEnvelopeMessages]
type accessBookmarkNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkNewResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accessBookmarkNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessBookmarkNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccessBookmarkNewResponseEnvelopeMessagesSource]
type accessBookmarkNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessBookmarkNewResponseEnvelopeSuccess bool

const (
	AccessBookmarkNewResponseEnvelopeSuccessTrue AccessBookmarkNewResponseEnvelopeSuccess = true
)

func (r AccessBookmarkNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessBookmarkNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessBookmarkUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Body      interface{}         `json:"body,required"`
}

func (r AccessBookmarkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AccessBookmarkUpdateResponseEnvelope struct {
	Errors   []AccessBookmarkUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessBookmarkUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Bookmark                                    `json:"result"`
	JSON    accessBookmarkUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessBookmarkUpdateResponseEnvelope]
type accessBookmarkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkUpdateResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccessBookmarkUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessBookmarkUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessBookmarkUpdateResponseEnvelopeErrors]
type accessBookmarkUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accessBookmarkUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [AccessBookmarkUpdateResponseEnvelopeErrorsSource]
type accessBookmarkUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkUpdateResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccessBookmarkUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessBookmarkUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessBookmarkUpdateResponseEnvelopeMessages]
type accessBookmarkUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accessBookmarkUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessBookmarkUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessBookmarkUpdateResponseEnvelopeMessagesSource]
type accessBookmarkUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessBookmarkUpdateResponseEnvelopeSuccess bool

const (
	AccessBookmarkUpdateResponseEnvelopeSuccessTrue AccessBookmarkUpdateResponseEnvelopeSuccess = true
)

func (r AccessBookmarkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessBookmarkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessBookmarkListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessBookmarkDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessBookmarkDeleteResponseEnvelope struct {
	Errors   []AccessBookmarkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessBookmarkDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessBookmarkDeleteResponse                `json:"result"`
	JSON    accessBookmarkDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessBookmarkDeleteResponseEnvelope]
type accessBookmarkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AccessBookmarkDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessBookmarkDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessBookmarkDeleteResponseEnvelopeErrors]
type accessBookmarkDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    accessBookmarkDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [AccessBookmarkDeleteResponseEnvelopeErrorsSource]
type accessBookmarkDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           AccessBookmarkDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessBookmarkDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessBookmarkDeleteResponseEnvelopeMessages]
type accessBookmarkDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    accessBookmarkDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessBookmarkDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AccessBookmarkDeleteResponseEnvelopeMessagesSource]
type accessBookmarkDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessBookmarkDeleteResponseEnvelopeSuccess bool

const (
	AccessBookmarkDeleteResponseEnvelopeSuccessTrue AccessBookmarkDeleteResponseEnvelopeSuccess = true
)

func (r AccessBookmarkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessBookmarkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessBookmarkGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessBookmarkGetResponseEnvelope struct {
	Errors   []AccessBookmarkGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessBookmarkGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessBookmarkGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Bookmark                                 `json:"result"`
	JSON    accessBookmarkGetResponseEnvelopeJSON    `json:"-"`
}

// accessBookmarkGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessBookmarkGetResponseEnvelope]
type accessBookmarkGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           AccessBookmarkGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessBookmarkGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessBookmarkGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessBookmarkGetResponseEnvelopeErrors]
type accessBookmarkGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    accessBookmarkGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessBookmarkGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AccessBookmarkGetResponseEnvelopeErrorsSource]
type accessBookmarkGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           AccessBookmarkGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessBookmarkGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessBookmarkGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessBookmarkGetResponseEnvelopeMessages]
type accessBookmarkGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessBookmarkGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    accessBookmarkGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessBookmarkGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AccessBookmarkGetResponseEnvelopeMessagesSource]
type accessBookmarkGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessBookmarkGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessBookmarkGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessBookmarkGetResponseEnvelopeSuccess bool

const (
	AccessBookmarkGetResponseEnvelopeSuccessTrue AccessBookmarkGetResponseEnvelopeSuccess = true
)

func (r AccessBookmarkGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessBookmarkGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
