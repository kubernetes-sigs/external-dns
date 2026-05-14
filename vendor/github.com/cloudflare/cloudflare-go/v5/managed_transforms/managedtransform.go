// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package managed_transforms

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

// ManagedTransformService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewManagedTransformService] method instead.
type ManagedTransformService struct {
	Options []option.RequestOption
}

// NewManagedTransformService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewManagedTransformService(opts ...option.RequestOption) (r *ManagedTransformService) {
	r = &ManagedTransformService{}
	r.Options = opts
	return
}

// Fetches a list of all Managed Transforms.
func (r *ManagedTransformService) List(ctx context.Context, query ManagedTransformListParams, opts ...option.RequestOption) (res *ManagedTransformListResponse, err error) {
	var env ManagedTransformListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disables all Managed Transforms.
func (r *ManagedTransformService) Delete(ctx context.Context, body ManagedTransformDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Updates the status of one or more Managed Transforms.
func (r *ManagedTransformService) Edit(ctx context.Context, params ManagedTransformEditParams, opts ...option.RequestOption) (res *ManagedTransformEditResponse, err error) {
	var env ManagedTransformEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/managed_headers", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A result.
type ManagedTransformListResponse struct {
	// The list of Managed Request Transforms.
	ManagedRequestHeaders []ManagedTransformListResponseManagedRequestHeader `json:"managed_request_headers,required"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders []ManagedTransformListResponseManagedResponseHeader `json:"managed_response_headers,required"`
	JSON                   managedTransformListResponseJSON                    `json:"-"`
}

// managedTransformListResponseJSON contains the JSON metadata for the struct
// [ManagedTransformListResponse]
type managedTransformListResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedTransformListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ManagedTransformListResponseManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled,required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict,required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                             `json:"conflicts_with"`
	JSON          managedTransformListResponseManagedRequestHeaderJSON `json:"-"`
}

// managedTransformListResponseManagedRequestHeaderJSON contains the JSON metadata
// for the struct [ManagedTransformListResponseManagedRequestHeader]
type managedTransformListResponseManagedRequestHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformListResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ManagedTransformListResponseManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled,required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict,required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                              `json:"conflicts_with"`
	JSON          managedTransformListResponseManagedResponseHeaderJSON `json:"-"`
}

// managedTransformListResponseManagedResponseHeaderJSON contains the JSON metadata
// for the struct [ManagedTransformListResponseManagedResponseHeader]
type managedTransformListResponseManagedResponseHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformListResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

// A result.
type ManagedTransformEditResponse struct {
	// The list of Managed Request Transforms.
	ManagedRequestHeaders []ManagedTransformEditResponseManagedRequestHeader `json:"managed_request_headers,required"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders []ManagedTransformEditResponseManagedResponseHeader `json:"managed_response_headers,required"`
	JSON                   managedTransformEditResponseJSON                    `json:"-"`
}

// managedTransformEditResponseJSON contains the JSON metadata for the struct
// [ManagedTransformEditResponse]
type managedTransformEditResponseJSON struct {
	ManagedRequestHeaders  apijson.Field
	ManagedResponseHeaders apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ManagedTransformEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ManagedTransformEditResponseManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled,required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict,required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                             `json:"conflicts_with"`
	JSON          managedTransformEditResponseManagedRequestHeaderJSON `json:"-"`
}

// managedTransformEditResponseManagedRequestHeaderJSON contains the JSON metadata
// for the struct [ManagedTransformEditResponseManagedRequestHeader]
type managedTransformEditResponseManagedRequestHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformEditResponseManagedRequestHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseManagedRequestHeaderJSON) RawJSON() string {
	return r.raw
}

// A Managed Transform object.
type ManagedTransformEditResponseManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID string `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled bool `json:"enabled,required"`
	// Whether the Managed Transform conflicts with the currently-enabled Managed
	// Transforms.
	HasConflict bool `json:"has_conflict,required"`
	// The Managed Transforms that this Managed Transform conflicts with.
	ConflictsWith []string                                              `json:"conflicts_with"`
	JSON          managedTransformEditResponseManagedResponseHeaderJSON `json:"-"`
}

// managedTransformEditResponseManagedResponseHeaderJSON contains the JSON metadata
// for the struct [ManagedTransformEditResponseManagedResponseHeader]
type managedTransformEditResponseManagedResponseHeaderJSON struct {
	ID            apijson.Field
	Enabled       apijson.Field
	HasConflict   apijson.Field
	ConflictsWith apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ManagedTransformEditResponseManagedResponseHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseManagedResponseHeaderJSON) RawJSON() string {
	return r.raw
}

type ManagedTransformListParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

// A response object.
type ManagedTransformListResponseEnvelope struct {
	// A list of error messages.
	Errors []ManagedTransformListResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []ManagedTransformListResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result ManagedTransformListResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ManagedTransformListResponseEnvelopeSuccess `json:"success,required"`
	JSON    managedTransformListResponseEnvelopeJSON    `json:"-"`
}

// managedTransformListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ManagedTransformListResponseEnvelope]
type managedTransformListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformListResponseEnvelopeErrorsSource `json:"source"`
	JSON   managedTransformListResponseEnvelopeErrorsJSON   `json:"-"`
}

// managedTransformListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ManagedTransformListResponseEnvelopeErrors]
type managedTransformListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                               `json:"pointer,required"`
	JSON    managedTransformListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// managedTransformListResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ManagedTransformListResponseEnvelopeErrorsSource]
type managedTransformListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformListResponseEnvelopeMessagesSource `json:"source"`
	JSON   managedTransformListResponseEnvelopeMessagesJSON   `json:"-"`
}

// managedTransformListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ManagedTransformListResponseEnvelopeMessages]
type managedTransformListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer,required"`
	JSON    managedTransformListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// managedTransformListResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ManagedTransformListResponseEnvelopeMessagesSource]
type managedTransformListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ManagedTransformListResponseEnvelopeSuccess bool

const (
	ManagedTransformListResponseEnvelopeSuccessTrue ManagedTransformListResponseEnvelopeSuccess = true
)

func (r ManagedTransformListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ManagedTransformListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ManagedTransformDeleteParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ManagedTransformEditParams struct {
	// The unique ID of the zone.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The list of Managed Request Transforms.
	ManagedRequestHeaders param.Field[[]ManagedTransformEditParamsManagedRequestHeader] `json:"managed_request_headers,required"`
	// The list of Managed Response Transforms.
	ManagedResponseHeaders param.Field[[]ManagedTransformEditParamsManagedResponseHeader] `json:"managed_response_headers,required"`
}

func (r ManagedTransformEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Managed Transform object.
type ManagedTransformEditParamsManagedRequestHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ManagedTransformEditParamsManagedRequestHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A Managed Transform object.
type ManagedTransformEditParamsManagedResponseHeader struct {
	// The human-readable identifier of the Managed Transform.
	ID param.Field[string] `json:"id,required"`
	// Whether the Managed Transform is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r ManagedTransformEditParamsManagedResponseHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A response object.
type ManagedTransformEditResponseEnvelope struct {
	// A list of error messages.
	Errors []ManagedTransformEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []ManagedTransformEditResponseEnvelopeMessages `json:"messages,required"`
	// A result.
	Result ManagedTransformEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ManagedTransformEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    managedTransformEditResponseEnvelopeJSON    `json:"-"`
}

// managedTransformEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [ManagedTransformEditResponseEnvelope]
type managedTransformEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   managedTransformEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// managedTransformEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ManagedTransformEditResponseEnvelopeErrors]
type managedTransformEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                               `json:"pointer,required"`
	JSON    managedTransformEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// managedTransformEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ManagedTransformEditResponseEnvelopeErrorsSource]
type managedTransformEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type ManagedTransformEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source ManagedTransformEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   managedTransformEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// managedTransformEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ManagedTransformEditResponseEnvelopeMessages]
type managedTransformEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type ManagedTransformEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                                 `json:"pointer,required"`
	JSON    managedTransformEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// managedTransformEditResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ManagedTransformEditResponseEnvelopeMessagesSource]
type managedTransformEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ManagedTransformEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedTransformEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ManagedTransformEditResponseEnvelopeSuccess bool

const (
	ManagedTransformEditResponseEnvelopeSuccessTrue ManagedTransformEditResponseEnvelopeSuccess = true
)

func (r ManagedTransformEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ManagedTransformEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
