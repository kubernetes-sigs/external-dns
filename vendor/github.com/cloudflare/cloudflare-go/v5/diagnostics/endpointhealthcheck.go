// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package diagnostics

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

// EndpointHealthcheckService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEndpointHealthcheckService] method instead.
type EndpointHealthcheckService struct {
	Options []option.RequestOption
}

// NewEndpointHealthcheckService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEndpointHealthcheckService(opts ...option.RequestOption) (r *EndpointHealthcheckService) {
	r = &EndpointHealthcheckService{}
	r.Options = opts
	return
}

// Create Endpoint Health Check.
func (r *EndpointHealthcheckService) New(ctx context.Context, params EndpointHealthcheckNewParams, opts ...option.RequestOption) (res *EndpointHealthcheckNewResponse, err error) {
	var env EndpointHealthcheckNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/diagnostics/endpoint-healthchecks", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Endpoint Health Check.
func (r *EndpointHealthcheckService) Update(ctx context.Context, id string, params EndpointHealthcheckUpdateParams, opts ...option.RequestOption) (res *EndpointHealthcheckUpdateResponse, err error) {
	var env EndpointHealthcheckUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/diagnostics/endpoint-healthchecks/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Endpoint Health Checks.
func (r *EndpointHealthcheckService) List(ctx context.Context, query EndpointHealthcheckListParams, opts ...option.RequestOption) (res *EndpointHealthcheckListResponse, err error) {
	var env EndpointHealthcheckListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/diagnostics/endpoint-healthchecks", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Endpoint Health Check.
func (r *EndpointHealthcheckService) Delete(ctx context.Context, id string, body EndpointHealthcheckDeleteParams, opts ...option.RequestOption) (res *EndpointHealthcheckDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/diagnostics/endpoint-healthchecks/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a single Endpoint Health Check.
func (r *EndpointHealthcheckService) Get(ctx context.Context, id string, query EndpointHealthcheckGetParams, opts ...option.RequestOption) (res *EndpointHealthcheckGetResponse, err error) {
	var env EndpointHealthcheckGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/diagnostics/endpoint-healthchecks/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EndpointHealthcheckParam struct {
	// type of check to perform
	CheckType param.Field[EndpointHealthcheckCheckType] `json:"check_type,required"`
	// the IP address of the host to perform checks against
	Endpoint param.Field[string] `json:"endpoint,required"`
	// Optional name associated with this check
	Name param.Field[string] `json:"name"`
}

func (r EndpointHealthcheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// type of check to perform
type EndpointHealthcheckCheckType string

const (
	EndpointHealthcheckCheckTypeIcmp EndpointHealthcheckCheckType = "icmp"
)

func (r EndpointHealthcheckCheckType) IsKnown() bool {
	switch r {
	case EndpointHealthcheckCheckTypeIcmp:
		return true
	}
	return false
}

type EndpointHealthcheckNewResponse struct {
	// type of check to perform
	CheckType EndpointHealthcheckNewResponseCheckType `json:"check_type,required"`
	// the IP address of the host to perform checks against
	Endpoint string `json:"endpoint,required"`
	// UUID.
	ID string `json:"id"`
	// Optional name associated with this check
	Name string                             `json:"name"`
	JSON endpointHealthcheckNewResponseJSON `json:"-"`
}

// endpointHealthcheckNewResponseJSON contains the JSON metadata for the struct
// [EndpointHealthcheckNewResponse]
type endpointHealthcheckNewResponseJSON struct {
	CheckType   apijson.Field
	Endpoint    apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckNewResponseJSON) RawJSON() string {
	return r.raw
}

// type of check to perform
type EndpointHealthcheckNewResponseCheckType string

const (
	EndpointHealthcheckNewResponseCheckTypeIcmp EndpointHealthcheckNewResponseCheckType = "icmp"
)

func (r EndpointHealthcheckNewResponseCheckType) IsKnown() bool {
	switch r {
	case EndpointHealthcheckNewResponseCheckTypeIcmp:
		return true
	}
	return false
}

type EndpointHealthcheckUpdateResponse struct {
	// type of check to perform
	CheckType EndpointHealthcheckUpdateResponseCheckType `json:"check_type,required"`
	// the IP address of the host to perform checks against
	Endpoint string `json:"endpoint,required"`
	// UUID.
	ID string `json:"id"`
	// Optional name associated with this check
	Name string                                `json:"name"`
	JSON endpointHealthcheckUpdateResponseJSON `json:"-"`
}

// endpointHealthcheckUpdateResponseJSON contains the JSON metadata for the struct
// [EndpointHealthcheckUpdateResponse]
type endpointHealthcheckUpdateResponseJSON struct {
	CheckType   apijson.Field
	Endpoint    apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// type of check to perform
type EndpointHealthcheckUpdateResponseCheckType string

const (
	EndpointHealthcheckUpdateResponseCheckTypeIcmp EndpointHealthcheckUpdateResponseCheckType = "icmp"
)

func (r EndpointHealthcheckUpdateResponseCheckType) IsKnown() bool {
	switch r {
	case EndpointHealthcheckUpdateResponseCheckTypeIcmp:
		return true
	}
	return false
}

type EndpointHealthcheckListResponse struct {
	// type of check to perform
	CheckType EndpointHealthcheckListResponseCheckType `json:"check_type,required"`
	// the IP address of the host to perform checks against
	Endpoint string `json:"endpoint,required"`
	// UUID.
	ID string `json:"id"`
	// Optional name associated with this check
	Name string                              `json:"name"`
	JSON endpointHealthcheckListResponseJSON `json:"-"`
}

// endpointHealthcheckListResponseJSON contains the JSON metadata for the struct
// [EndpointHealthcheckListResponse]
type endpointHealthcheckListResponseJSON struct {
	CheckType   apijson.Field
	Endpoint    apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckListResponseJSON) RawJSON() string {
	return r.raw
}

// type of check to perform
type EndpointHealthcheckListResponseCheckType string

const (
	EndpointHealthcheckListResponseCheckTypeIcmp EndpointHealthcheckListResponseCheckType = "icmp"
)

func (r EndpointHealthcheckListResponseCheckType) IsKnown() bool {
	switch r {
	case EndpointHealthcheckListResponseCheckTypeIcmp:
		return true
	}
	return false
}

type EndpointHealthcheckDeleteResponse struct {
	Errors   []EndpointHealthcheckDeleteResponseError   `json:"errors,required"`
	Messages []EndpointHealthcheckDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success EndpointHealthcheckDeleteResponseSuccess `json:"success,required"`
	JSON    endpointHealthcheckDeleteResponseJSON    `json:"-"`
}

// endpointHealthcheckDeleteResponseJSON contains the JSON metadata for the struct
// [EndpointHealthcheckDeleteResponse]
type endpointHealthcheckDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckDeleteResponseError struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           EndpointHealthcheckDeleteResponseErrorsSource `json:"source"`
	JSON             endpointHealthcheckDeleteResponseErrorJSON    `json:"-"`
}

// endpointHealthcheckDeleteResponseErrorJSON contains the JSON metadata for the
// struct [EndpointHealthcheckDeleteResponseError]
type endpointHealthcheckDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckDeleteResponseErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    endpointHealthcheckDeleteResponseErrorsSourceJSON `json:"-"`
}

// endpointHealthcheckDeleteResponseErrorsSourceJSON contains the JSON metadata for
// the struct [EndpointHealthcheckDeleteResponseErrorsSource]
type endpointHealthcheckDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckDeleteResponseMessage struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           EndpointHealthcheckDeleteResponseMessagesSource `json:"source"`
	JSON             endpointHealthcheckDeleteResponseMessageJSON    `json:"-"`
}

// endpointHealthcheckDeleteResponseMessageJSON contains the JSON metadata for the
// struct [EndpointHealthcheckDeleteResponseMessage]
type endpointHealthcheckDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckDeleteResponseMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    endpointHealthcheckDeleteResponseMessagesSourceJSON `json:"-"`
}

// endpointHealthcheckDeleteResponseMessagesSourceJSON contains the JSON metadata
// for the struct [EndpointHealthcheckDeleteResponseMessagesSource]
type endpointHealthcheckDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EndpointHealthcheckDeleteResponseSuccess bool

const (
	EndpointHealthcheckDeleteResponseSuccessTrue EndpointHealthcheckDeleteResponseSuccess = true
)

func (r EndpointHealthcheckDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case EndpointHealthcheckDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type EndpointHealthcheckGetResponse struct {
	// type of check to perform
	CheckType EndpointHealthcheckGetResponseCheckType `json:"check_type,required"`
	// the IP address of the host to perform checks against
	Endpoint string `json:"endpoint,required"`
	// UUID.
	ID string `json:"id"`
	// Optional name associated with this check
	Name string                             `json:"name"`
	JSON endpointHealthcheckGetResponseJSON `json:"-"`
}

// endpointHealthcheckGetResponseJSON contains the JSON metadata for the struct
// [EndpointHealthcheckGetResponse]
type endpointHealthcheckGetResponseJSON struct {
	CheckType   apijson.Field
	Endpoint    apijson.Field
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckGetResponseJSON) RawJSON() string {
	return r.raw
}

// type of check to perform
type EndpointHealthcheckGetResponseCheckType string

const (
	EndpointHealthcheckGetResponseCheckTypeIcmp EndpointHealthcheckGetResponseCheckType = "icmp"
)

func (r EndpointHealthcheckGetResponseCheckType) IsKnown() bool {
	switch r {
	case EndpointHealthcheckGetResponseCheckTypeIcmp:
		return true
	}
	return false
}

type EndpointHealthcheckNewParams struct {
	// Identifier
	AccountID           param.Field[string]      `path:"account_id,required"`
	EndpointHealthcheck EndpointHealthcheckParam `json:"endpoint_healthcheck,required"`
}

func (r EndpointHealthcheckNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EndpointHealthcheck)
}

type EndpointHealthcheckNewResponseEnvelope struct {
	Errors   []EndpointHealthcheckNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EndpointHealthcheckNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success EndpointHealthcheckNewResponseEnvelopeSuccess `json:"success,required"`
	Result  EndpointHealthcheckNewResponse                `json:"result"`
	JSON    endpointHealthcheckNewResponseEnvelopeJSON    `json:"-"`
}

// endpointHealthcheckNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [EndpointHealthcheckNewResponseEnvelope]
type endpointHealthcheckNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckNewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           EndpointHealthcheckNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             endpointHealthcheckNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// endpointHealthcheckNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EndpointHealthcheckNewResponseEnvelopeErrors]
type endpointHealthcheckNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    endpointHealthcheckNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// endpointHealthcheckNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [EndpointHealthcheckNewResponseEnvelopeErrorsSource]
type endpointHealthcheckNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckNewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           EndpointHealthcheckNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             endpointHealthcheckNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// endpointHealthcheckNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EndpointHealthcheckNewResponseEnvelopeMessages]
type endpointHealthcheckNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    endpointHealthcheckNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// endpointHealthcheckNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [EndpointHealthcheckNewResponseEnvelopeMessagesSource]
type endpointHealthcheckNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EndpointHealthcheckNewResponseEnvelopeSuccess bool

const (
	EndpointHealthcheckNewResponseEnvelopeSuccessTrue EndpointHealthcheckNewResponseEnvelopeSuccess = true
)

func (r EndpointHealthcheckNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EndpointHealthcheckNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EndpointHealthcheckUpdateParams struct {
	// Identifier
	AccountID           param.Field[string]      `path:"account_id,required"`
	EndpointHealthcheck EndpointHealthcheckParam `json:"endpoint_healthcheck,required"`
}

func (r EndpointHealthcheckUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.EndpointHealthcheck)
}

type EndpointHealthcheckUpdateResponseEnvelope struct {
	Errors   []EndpointHealthcheckUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EndpointHealthcheckUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success EndpointHealthcheckUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  EndpointHealthcheckUpdateResponse                `json:"result"`
	JSON    endpointHealthcheckUpdateResponseEnvelopeJSON    `json:"-"`
}

// endpointHealthcheckUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [EndpointHealthcheckUpdateResponseEnvelope]
type endpointHealthcheckUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckUpdateResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           EndpointHealthcheckUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             endpointHealthcheckUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// endpointHealthcheckUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [EndpointHealthcheckUpdateResponseEnvelopeErrors]
type endpointHealthcheckUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    endpointHealthcheckUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// endpointHealthcheckUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [EndpointHealthcheckUpdateResponseEnvelopeErrorsSource]
type endpointHealthcheckUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckUpdateResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           EndpointHealthcheckUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             endpointHealthcheckUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// endpointHealthcheckUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EndpointHealthcheckUpdateResponseEnvelopeMessages]
type endpointHealthcheckUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    endpointHealthcheckUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// endpointHealthcheckUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [EndpointHealthcheckUpdateResponseEnvelopeMessagesSource]
type endpointHealthcheckUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EndpointHealthcheckUpdateResponseEnvelopeSuccess bool

const (
	EndpointHealthcheckUpdateResponseEnvelopeSuccessTrue EndpointHealthcheckUpdateResponseEnvelopeSuccess = true
)

func (r EndpointHealthcheckUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EndpointHealthcheckUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EndpointHealthcheckListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type EndpointHealthcheckListResponseEnvelope struct {
	Errors   []EndpointHealthcheckListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EndpointHealthcheckListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success EndpointHealthcheckListResponseEnvelopeSuccess `json:"success,required"`
	Result  EndpointHealthcheckListResponse                `json:"result"`
	JSON    endpointHealthcheckListResponseEnvelopeJSON    `json:"-"`
}

// endpointHealthcheckListResponseEnvelopeJSON contains the JSON metadata for the
// struct [EndpointHealthcheckListResponseEnvelope]
type endpointHealthcheckListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckListResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           EndpointHealthcheckListResponseEnvelopeErrorsSource `json:"source"`
	JSON             endpointHealthcheckListResponseEnvelopeErrorsJSON   `json:"-"`
}

// endpointHealthcheckListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EndpointHealthcheckListResponseEnvelopeErrors]
type endpointHealthcheckListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckListResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    endpointHealthcheckListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// endpointHealthcheckListResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [EndpointHealthcheckListResponseEnvelopeErrorsSource]
type endpointHealthcheckListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckListResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           EndpointHealthcheckListResponseEnvelopeMessagesSource `json:"source"`
	JSON             endpointHealthcheckListResponseEnvelopeMessagesJSON   `json:"-"`
}

// endpointHealthcheckListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EndpointHealthcheckListResponseEnvelopeMessages]
type endpointHealthcheckListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckListResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    endpointHealthcheckListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// endpointHealthcheckListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [EndpointHealthcheckListResponseEnvelopeMessagesSource]
type endpointHealthcheckListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EndpointHealthcheckListResponseEnvelopeSuccess bool

const (
	EndpointHealthcheckListResponseEnvelopeSuccessTrue EndpointHealthcheckListResponseEnvelopeSuccess = true
)

func (r EndpointHealthcheckListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EndpointHealthcheckListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type EndpointHealthcheckDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type EndpointHealthcheckGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type EndpointHealthcheckGetResponseEnvelope struct {
	Errors   []EndpointHealthcheckGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EndpointHealthcheckGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success EndpointHealthcheckGetResponseEnvelopeSuccess `json:"success,required"`
	Result  EndpointHealthcheckGetResponse                `json:"result"`
	JSON    endpointHealthcheckGetResponseEnvelopeJSON    `json:"-"`
}

// endpointHealthcheckGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [EndpointHealthcheckGetResponseEnvelope]
type endpointHealthcheckGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code,required"`
	Message          string                                             `json:"message,required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           EndpointHealthcheckGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             endpointHealthcheckGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// endpointHealthcheckGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [EndpointHealthcheckGetResponseEnvelopeErrors]
type endpointHealthcheckGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    endpointHealthcheckGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// endpointHealthcheckGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [EndpointHealthcheckGetResponseEnvelopeErrorsSource]
type endpointHealthcheckGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           EndpointHealthcheckGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             endpointHealthcheckGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// endpointHealthcheckGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [EndpointHealthcheckGetResponseEnvelopeMessages]
type endpointHealthcheckGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *EndpointHealthcheckGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type EndpointHealthcheckGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    endpointHealthcheckGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// endpointHealthcheckGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [EndpointHealthcheckGetResponseEnvelopeMessagesSource]
type endpointHealthcheckGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EndpointHealthcheckGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r endpointHealthcheckGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type EndpointHealthcheckGetResponseEnvelopeSuccess bool

const (
	EndpointHealthcheckGetResponseEnvelopeSuccessTrue EndpointHealthcheckGetResponseEnvelopeSuccess = true
)

func (r EndpointHealthcheckGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case EndpointHealthcheckGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
