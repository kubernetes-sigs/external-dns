// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// ZoneTransferIncomingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferIncomingService] method instead.
type ZoneTransferIncomingService struct {
	Options []option.RequestOption
}

// NewZoneTransferIncomingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneTransferIncomingService(opts ...option.RequestOption) (r *ZoneTransferIncomingService) {
	r = &ZoneTransferIncomingService{}
	r.Options = opts
	return
}

// Create secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) New(ctx context.Context, params ZoneTransferIncomingNewParams, opts ...option.RequestOption) (res *ZoneTransferIncomingNewResponse, err error) {
	var env ZoneTransferIncomingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) Update(ctx context.Context, params ZoneTransferIncomingUpdateParams, opts ...option.RequestOption) (res *ZoneTransferIncomingUpdateResponse, err error) {
	var env ZoneTransferIncomingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) Delete(ctx context.Context, body ZoneTransferIncomingDeleteParams, opts ...option.RequestOption) (res *ZoneTransferIncomingDeleteResponse, err error) {
	var env ZoneTransferIncomingDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) Get(ctx context.Context, query ZoneTransferIncomingGetParams, opts ...option.RequestOption) (res *ZoneTransferIncomingGetResponse, err error) {
	var env ZoneTransferIncomingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneTransferIncomingNewResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                             `json:"soa_serial"`
	JSON      zoneTransferIncomingNewResponseJSON `json:"-"`
}

// zoneTransferIncomingNewResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingNewResponse]
type zoneTransferIncomingNewResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingUpdateResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                                `json:"soa_serial"`
	JSON      zoneTransferIncomingUpdateResponseJSON `json:"-"`
}

// zoneTransferIncomingUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingUpdateResponse]
type zoneTransferIncomingUpdateResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingDeleteResponse struct {
	ID   string                                 `json:"id"`
	JSON zoneTransferIncomingDeleteResponseJSON `json:"-"`
}

// zoneTransferIncomingDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingDeleteResponse]
type zoneTransferIncomingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingGetResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                             `json:"soa_serial"`
	JSON      zoneTransferIncomingGetResponseJSON `json:"-"`
}

// zoneTransferIncomingGetResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingGetResponse]
type zoneTransferIncomingGetResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingNewParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r ZoneTransferIncomingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferIncomingNewResponseEnvelope struct {
	Errors   []ZoneTransferIncomingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferIncomingNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferIncomingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingNewResponse                `json:"result"`
	JSON    zoneTransferIncomingNewResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferIncomingNewResponseEnvelope]
type zoneTransferIncomingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingNewResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ZoneTransferIncomingNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferIncomingNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferIncomingNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneTransferIncomingNewResponseEnvelopeErrors]
type zoneTransferIncomingNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    zoneTransferIncomingNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferIncomingNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingNewResponseEnvelopeErrorsSource]
type zoneTransferIncomingNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingNewResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ZoneTransferIncomingNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferIncomingNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferIncomingNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneTransferIncomingNewResponseEnvelopeMessages]
type zoneTransferIncomingNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    zoneTransferIncomingNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferIncomingNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingNewResponseEnvelopeMessagesSource]
type zoneTransferIncomingNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferIncomingNewResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingNewResponseEnvelopeSuccessTrue ZoneTransferIncomingNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferIncomingUpdateParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r ZoneTransferIncomingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferIncomingUpdateResponseEnvelope struct {
	Errors   []ZoneTransferIncomingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferIncomingUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferIncomingUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingUpdateResponse                `json:"result"`
	JSON    zoneTransferIncomingUpdateResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferIncomingUpdateResponseEnvelope]
type zoneTransferIncomingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingUpdateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ZoneTransferIncomingUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferIncomingUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferIncomingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneTransferIncomingUpdateResponseEnvelopeErrors]
type zoneTransferIncomingUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    zoneTransferIncomingUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferIncomingUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingUpdateResponseEnvelopeErrorsSource]
type zoneTransferIncomingUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingUpdateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           ZoneTransferIncomingUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferIncomingUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferIncomingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingUpdateResponseEnvelopeMessages]
type zoneTransferIncomingUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    zoneTransferIncomingUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferIncomingUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferIncomingUpdateResponseEnvelopeMessagesSource]
type zoneTransferIncomingUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferIncomingUpdateResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingUpdateResponseEnvelopeSuccessTrue ZoneTransferIncomingUpdateResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferIncomingDeleteParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferIncomingDeleteResponseEnvelope struct {
	Errors   []ZoneTransferIncomingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferIncomingDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferIncomingDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingDeleteResponse                `json:"result"`
	JSON    zoneTransferIncomingDeleteResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferIncomingDeleteResponseEnvelope]
type zoneTransferIncomingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ZoneTransferIncomingDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferIncomingDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferIncomingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneTransferIncomingDeleteResponseEnvelopeErrors]
type zoneTransferIncomingDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    zoneTransferIncomingDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferIncomingDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingDeleteResponseEnvelopeErrorsSource]
type zoneTransferIncomingDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           ZoneTransferIncomingDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferIncomingDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferIncomingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingDeleteResponseEnvelopeMessages]
type zoneTransferIncomingDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    zoneTransferIncomingDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferIncomingDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferIncomingDeleteResponseEnvelopeMessagesSource]
type zoneTransferIncomingDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferIncomingDeleteResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingDeleteResponseEnvelopeSuccessTrue ZoneTransferIncomingDeleteResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferIncomingGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferIncomingGetResponseEnvelope struct {
	Errors   []ZoneTransferIncomingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferIncomingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferIncomingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingGetResponse                `json:"result"`
	JSON    zoneTransferIncomingGetResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferIncomingGetResponseEnvelope]
type zoneTransferIncomingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ZoneTransferIncomingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferIncomingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferIncomingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneTransferIncomingGetResponseEnvelopeErrors]
type zoneTransferIncomingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    zoneTransferIncomingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferIncomingGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingGetResponseEnvelopeErrorsSource]
type zoneTransferIncomingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ZoneTransferIncomingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferIncomingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferIncomingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneTransferIncomingGetResponseEnvelopeMessages]
type zoneTransferIncomingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    zoneTransferIncomingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferIncomingGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneTransferIncomingGetResponseEnvelopeMessagesSource]
type zoneTransferIncomingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferIncomingGetResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingGetResponseEnvelopeSuccessTrue ZoneTransferIncomingGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
