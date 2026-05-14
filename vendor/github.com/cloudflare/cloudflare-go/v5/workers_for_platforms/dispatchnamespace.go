// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

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

// DispatchNamespaceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchNamespaceService] method instead.
type DispatchNamespaceService struct {
	Options []option.RequestOption
	Scripts *DispatchNamespaceScriptService
}

// NewDispatchNamespaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDispatchNamespaceService(opts ...option.RequestOption) (r *DispatchNamespaceService) {
	r = &DispatchNamespaceService{}
	r.Options = opts
	r.Scripts = NewDispatchNamespaceScriptService(opts...)
	return
}

// Create a new Workers for Platforms namespace.
func (r *DispatchNamespaceService) New(ctx context.Context, params DispatchNamespaceNewParams, opts ...option.RequestOption) (res *DispatchNamespaceNewResponse, err error) {
	var env DispatchNamespaceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of Workers for Platforms namespaces.
func (r *DispatchNamespaceService) List(ctx context.Context, query DispatchNamespaceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DispatchNamespaceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces", query.AccountID)
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

// Fetch a list of Workers for Platforms namespaces.
func (r *DispatchNamespaceService) ListAutoPaging(ctx context.Context, query DispatchNamespaceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DispatchNamespaceListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a Workers for Platforms namespace.
func (r *DispatchNamespaceService) Delete(ctx context.Context, dispatchNamespace string, body DispatchNamespaceDeleteParams, opts ...option.RequestOption) (res *DispatchNamespaceDeleteResponse, err error) {
	var env DispatchNamespaceDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", body.AccountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a Workers for Platforms namespace.
func (r *DispatchNamespaceService) Get(ctx context.Context, dispatchNamespace string, query DispatchNamespaceGetParams, opts ...option.RequestOption) (res *DispatchNamespaceGetResponse, err error) {
	var env DispatchNamespaceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dispatchNamespace == "" {
		err = errors.New("missing required dispatch_namespace parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s", query.AccountID, dispatchNamespace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DispatchNamespaceNewResponse struct {
	// Identifier.
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier.
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string `json:"namespace_name"`
	// The current number of scripts in this Dispatch Namespace.
	ScriptCount int64 `json:"script_count"`
	// Whether the Workers in the namespace are executed in a "trusted" manner. When a
	// Worker is trusted, it has access to the shared caches for the zone in the Cache
	// API, and has access to the `request.cf` object on incoming Requests. When a
	// Worker is untrusted, caches are not shared across the zone, and `request.cf` is
	// undefined. By default, Workers in a namespace are "untrusted".
	TrustedWorkers bool                             `json:"trusted_workers"`
	JSON           dispatchNamespaceNewResponseJSON `json:"-"`
}

// dispatchNamespaceNewResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceNewResponse]
type dispatchNamespaceNewResponseJSON struct {
	CreatedBy      apijson.Field
	CreatedOn      apijson.Field
	ModifiedBy     apijson.Field
	ModifiedOn     apijson.Field
	NamespaceID    apijson.Field
	NamespaceName  apijson.Field
	ScriptCount    apijson.Field
	TrustedWorkers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceListResponse struct {
	// Identifier.
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier.
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string `json:"namespace_name"`
	// The current number of scripts in this Dispatch Namespace.
	ScriptCount int64 `json:"script_count"`
	// Whether the Workers in the namespace are executed in a "trusted" manner. When a
	// Worker is trusted, it has access to the shared caches for the zone in the Cache
	// API, and has access to the `request.cf` object on incoming Requests. When a
	// Worker is untrusted, caches are not shared across the zone, and `request.cf` is
	// undefined. By default, Workers in a namespace are "untrusted".
	TrustedWorkers bool                              `json:"trusted_workers"`
	JSON           dispatchNamespaceListResponseJSON `json:"-"`
}

// dispatchNamespaceListResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceListResponse]
type dispatchNamespaceListResponseJSON struct {
	CreatedBy      apijson.Field
	CreatedOn      apijson.Field
	ModifiedBy     apijson.Field
	ModifiedOn     apijson.Field
	NamespaceID    apijson.Field
	NamespaceName  apijson.Field
	ScriptCount    apijson.Field
	TrustedWorkers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DispatchNamespaceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceListResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponse = interface{}

type DispatchNamespaceGetResponse struct {
	// Identifier.
	CreatedBy string `json:"created_by"`
	// When the script was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// Identifier.
	ModifiedBy string `json:"modified_by"`
	// When the script was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// API Resource UUID tag.
	NamespaceID string `json:"namespace_id"`
	// Name of the Workers for Platforms dispatch namespace.
	NamespaceName string `json:"namespace_name"`
	// The current number of scripts in this Dispatch Namespace.
	ScriptCount int64 `json:"script_count"`
	// Whether the Workers in the namespace are executed in a "trusted" manner. When a
	// Worker is trusted, it has access to the shared caches for the zone in the Cache
	// API, and has access to the `request.cf` object on incoming Requests. When a
	// Worker is untrusted, caches are not shared across the zone, and `request.cf` is
	// undefined. By default, Workers in a namespace are "untrusted".
	TrustedWorkers bool                             `json:"trusted_workers"`
	JSON           dispatchNamespaceGetResponseJSON `json:"-"`
}

// dispatchNamespaceGetResponseJSON contains the JSON metadata for the struct
// [DispatchNamespaceGetResponse]
type dispatchNamespaceGetResponseJSON struct {
	CreatedBy      apijson.Field
	CreatedOn      apijson.Field
	ModifiedBy     apijson.Field
	ModifiedOn     apijson.Field
	NamespaceID    apijson.Field
	NamespaceName  apijson.Field
	ScriptCount    apijson.Field
	TrustedWorkers apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The name of the dispatch namespace.
	Name param.Field[string] `json:"name"`
}

func (r DispatchNamespaceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DispatchNamespaceNewResponseEnvelope struct {
	Errors   []DispatchNamespaceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceNewResponse                `json:"result"`
	JSON    dispatchNamespaceNewResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceNewResponseEnvelope]
type dispatchNamespaceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DispatchNamespaceNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatchNamespaceNewResponseEnvelopeErrors]
type dispatchNamespaceNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dispatchNamespaceNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DispatchNamespaceNewResponseEnvelopeErrorsSource]
type dispatchNamespaceNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DispatchNamespaceNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DispatchNamespaceNewResponseEnvelopeMessages]
type dispatchNamespaceNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dispatchNamespaceNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DispatchNamespaceNewResponseEnvelopeMessagesSource]
type dispatchNamespaceNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceNewResponseEnvelopeSuccess bool

const (
	DispatchNamespaceNewResponseEnvelopeSuccessTrue DispatchNamespaceNewResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceDeleteResponseEnvelope struct {
	Errors   []DispatchNamespaceDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceDeleteResponse                `json:"result,nullable"`
	JSON    dispatchNamespaceDeleteResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceDeleteResponseEnvelope]
type dispatchNamespaceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           DispatchNamespaceDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatchNamespaceDeleteResponseEnvelopeErrors]
type dispatchNamespaceDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    dispatchNamespaceDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DispatchNamespaceDeleteResponseEnvelopeErrorsSource]
type dispatchNamespaceDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DispatchNamespaceDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DispatchNamespaceDeleteResponseEnvelopeMessages]
type dispatchNamespaceDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dispatchNamespaceDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DispatchNamespaceDeleteResponseEnvelopeMessagesSource]
type dispatchNamespaceDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceDeleteResponseEnvelopeSuccess bool

const (
	DispatchNamespaceDeleteResponseEnvelopeSuccessTrue DispatchNamespaceDeleteResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DispatchNamespaceGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DispatchNamespaceGetResponseEnvelope struct {
	Errors   []DispatchNamespaceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DispatchNamespaceGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DispatchNamespaceGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DispatchNamespaceGetResponse                `json:"result"`
	JSON    dispatchNamespaceGetResponseEnvelopeJSON    `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DispatchNamespaceGetResponseEnvelope]
type dispatchNamespaceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceGetResponseEnvelopeErrors struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DispatchNamespaceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dispatchNamespaceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DispatchNamespaceGetResponseEnvelopeErrors]
type dispatchNamespaceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceGetResponseEnvelopeErrorsSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dispatchNamespaceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DispatchNamespaceGetResponseEnvelopeErrorsSource]
type dispatchNamespaceGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceGetResponseEnvelopeMessages struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DispatchNamespaceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dispatchNamespaceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DispatchNamespaceGetResponseEnvelopeMessages]
type dispatchNamespaceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DispatchNamespaceGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dispatchNamespaceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dispatchNamespaceGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DispatchNamespaceGetResponseEnvelopeMessagesSource]
type dispatchNamespaceGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DispatchNamespaceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dispatchNamespaceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DispatchNamespaceGetResponseEnvelopeSuccess bool

const (
	DispatchNamespaceGetResponseEnvelopeSuccessTrue DispatchNamespaceGetResponseEnvelopeSuccess = true
)

func (r DispatchNamespaceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DispatchNamespaceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
