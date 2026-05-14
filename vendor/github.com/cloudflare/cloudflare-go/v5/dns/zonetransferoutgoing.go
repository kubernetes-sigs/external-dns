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

// ZoneTransferOutgoingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferOutgoingService] method instead.
type ZoneTransferOutgoingService struct {
	Options []option.RequestOption
	Status  *ZoneTransferOutgoingStatusService
}

// NewZoneTransferOutgoingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneTransferOutgoingService(opts ...option.RequestOption) (r *ZoneTransferOutgoingService) {
	r = &ZoneTransferOutgoingService{}
	r.Options = opts
	r.Status = NewZoneTransferOutgoingStatusService(opts...)
	return
}

// Create primary zone configuration for outgoing zone transfers.
func (r *ZoneTransferOutgoingService) New(ctx context.Context, params ZoneTransferOutgoingNewParams, opts ...option.RequestOption) (res *ZoneTransferOutgoingNewResponse, err error) {
	var env ZoneTransferOutgoingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update primary zone configuration for outgoing zone transfers.
func (r *ZoneTransferOutgoingService) Update(ctx context.Context, params ZoneTransferOutgoingUpdateParams, opts ...option.RequestOption) (res *ZoneTransferOutgoingUpdateResponse, err error) {
	var env ZoneTransferOutgoingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete primary zone configuration for outgoing zone transfers.
func (r *ZoneTransferOutgoingService) Delete(ctx context.Context, body ZoneTransferOutgoingDeleteParams, opts ...option.RequestOption) (res *ZoneTransferOutgoingDeleteResponse, err error) {
	var env ZoneTransferOutgoingDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Disable outgoing zone transfers for primary zone and clears IXFR backlog of
// primary zone.
func (r *ZoneTransferOutgoingService) Disable(ctx context.Context, params ZoneTransferOutgoingDisableParams, opts ...option.RequestOption) (res *DisableTransfer, err error) {
	var env ZoneTransferOutgoingDisableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/disable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable outgoing zone transfers for primary zone.
func (r *ZoneTransferOutgoingService) Enable(ctx context.Context, params ZoneTransferOutgoingEnableParams, opts ...option.RequestOption) (res *EnableTransfer, err error) {
	var env ZoneTransferOutgoingEnableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Notifies the secondary nameserver(s) and clears IXFR backlog of primary zone.
func (r *ZoneTransferOutgoingService) ForceNotify(ctx context.Context, params ZoneTransferOutgoingForceNotifyParams, opts ...option.RequestOption) (res *string, err error) {
	var env ZoneTransferOutgoingForceNotifyResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing/force_notify", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get primary zone configuration for outgoing zone transfers.
func (r *ZoneTransferOutgoingService) Get(ctx context.Context, query ZoneTransferOutgoingGetParams, opts ...option.RequestOption) (res *ZoneTransferOutgoingGetResponse, err error) {
	var env ZoneTransferOutgoingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/outgoing", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DisableTransfer = string

type EnableTransfer = string

type ZoneTransferOutgoingNewResponse struct {
	ID string `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                             `json:"soa_serial"`
	JSON      zoneTransferOutgoingNewResponseJSON `json:"-"`
}

// zoneTransferOutgoingNewResponseJSON contains the JSON metadata for the struct
// [ZoneTransferOutgoingNewResponse]
type zoneTransferOutgoingNewResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SOASerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneTransferOutgoingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingUpdateResponse struct {
	ID string `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                                `json:"soa_serial"`
	JSON      zoneTransferOutgoingUpdateResponseJSON `json:"-"`
}

// zoneTransferOutgoingUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneTransferOutgoingUpdateResponse]
type zoneTransferOutgoingUpdateResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SOASerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneTransferOutgoingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDeleteResponse struct {
	ID   string                                 `json:"id"`
	JSON zoneTransferOutgoingDeleteResponseJSON `json:"-"`
}

// zoneTransferOutgoingDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneTransferOutgoingDeleteResponse]
type zoneTransferOutgoingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingGetResponse struct {
	ID string `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                             `json:"soa_serial"`
	JSON      zoneTransferOutgoingGetResponseJSON `json:"-"`
}

// zoneTransferOutgoingGetResponseJSON contains the JSON metadata for the struct
// [ZoneTransferOutgoingGetResponse]
type zoneTransferOutgoingGetResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SOASerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ZoneTransferOutgoingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingNewParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r ZoneTransferOutgoingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferOutgoingNewResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferOutgoingNewResponse                `json:"result"`
	JSON    zoneTransferOutgoingNewResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferOutgoingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferOutgoingNewResponseEnvelope]
type zoneTransferOutgoingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingNewResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ZoneTransferOutgoingNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingNewResponseEnvelopeErrors]
type zoneTransferOutgoingNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    zoneTransferOutgoingNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingNewResponseEnvelopeErrorsSource]
type zoneTransferOutgoingNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingNewResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ZoneTransferOutgoingNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneTransferOutgoingNewResponseEnvelopeMessages]
type zoneTransferOutgoingNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    zoneTransferOutgoingNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingNewResponseEnvelopeMessagesSource]
type zoneTransferOutgoingNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingNewResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingNewResponseEnvelopeSuccessTrue ZoneTransferOutgoingNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferOutgoingUpdateParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r ZoneTransferOutgoingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferOutgoingUpdateResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferOutgoingUpdateResponse                `json:"result"`
	JSON    zoneTransferOutgoingUpdateResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferOutgoingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingUpdateResponseEnvelope]
type zoneTransferOutgoingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingUpdateResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ZoneTransferOutgoingUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneTransferOutgoingUpdateResponseEnvelopeErrors]
type zoneTransferOutgoingUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    zoneTransferOutgoingUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingUpdateResponseEnvelopeErrorsSource]
type zoneTransferOutgoingUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingUpdateResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           ZoneTransferOutgoingUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingUpdateResponseEnvelopeMessages]
type zoneTransferOutgoingUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    zoneTransferOutgoingUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferOutgoingUpdateResponseEnvelopeMessagesSource]
type zoneTransferOutgoingUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingUpdateResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingUpdateResponseEnvelopeSuccessTrue ZoneTransferOutgoingUpdateResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferOutgoingDeleteParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferOutgoingDeleteResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferOutgoingDeleteResponse                `json:"result"`
	JSON    zoneTransferOutgoingDeleteResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferOutgoingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingDeleteResponseEnvelope]
type zoneTransferOutgoingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDeleteResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ZoneTransferOutgoingDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneTransferOutgoingDeleteResponseEnvelopeErrors]
type zoneTransferOutgoingDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    zoneTransferOutgoingDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingDeleteResponseEnvelopeErrorsSource]
type zoneTransferOutgoingDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDeleteResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           ZoneTransferOutgoingDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingDeleteResponseEnvelopeMessages]
type zoneTransferOutgoingDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    zoneTransferOutgoingDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferOutgoingDeleteResponseEnvelopeMessagesSource]
type zoneTransferOutgoingDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingDeleteResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingDeleteResponseEnvelopeSuccessTrue ZoneTransferOutgoingDeleteResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferOutgoingDisableParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r ZoneTransferOutgoingDisableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneTransferOutgoingDisableResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingDisableResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingDisableResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingDisableResponseEnvelopeSuccess `json:"success,required"`
	// The zone transfer status of a primary zone.
	Result DisableTransfer                                 `json:"result"`
	JSON   zoneTransferOutgoingDisableResponseEnvelopeJSON `json:"-"`
}

// zoneTransferOutgoingDisableResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingDisableResponseEnvelope]
type zoneTransferOutgoingDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDisableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDisableResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code,required"`
	Message          string                                                  `json:"message,required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           ZoneTransferOutgoingDisableResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingDisableResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingDisableResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneTransferOutgoingDisableResponseEnvelopeErrors]
type zoneTransferOutgoingDisableResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDisableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDisableResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDisableResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    zoneTransferOutgoingDisableResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingDisableResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferOutgoingDisableResponseEnvelopeErrorsSource]
type zoneTransferOutgoingDisableResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDisableResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDisableResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDisableResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           ZoneTransferOutgoingDisableResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingDisableResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingDisableResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingDisableResponseEnvelopeMessages]
type zoneTransferOutgoingDisableResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDisableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDisableResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingDisableResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    zoneTransferOutgoingDisableResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingDisableResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferOutgoingDisableResponseEnvelopeMessagesSource]
type zoneTransferOutgoingDisableResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingDisableResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingDisableResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingDisableResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingDisableResponseEnvelopeSuccessTrue ZoneTransferOutgoingDisableResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingDisableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingDisableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferOutgoingEnableParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r ZoneTransferOutgoingEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneTransferOutgoingEnableResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingEnableResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingEnableResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingEnableResponseEnvelopeSuccess `json:"success,required"`
	// The zone transfer status of a primary zone.
	Result EnableTransfer                                 `json:"result"`
	JSON   zoneTransferOutgoingEnableResponseEnvelopeJSON `json:"-"`
}

// zoneTransferOutgoingEnableResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingEnableResponseEnvelope]
type zoneTransferOutgoingEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingEnableResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           ZoneTransferOutgoingEnableResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingEnableResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingEnableResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ZoneTransferOutgoingEnableResponseEnvelopeErrors]
type zoneTransferOutgoingEnableResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingEnableResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingEnableResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingEnableResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    zoneTransferOutgoingEnableResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingEnableResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingEnableResponseEnvelopeErrorsSource]
type zoneTransferOutgoingEnableResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingEnableResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingEnableResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingEnableResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           ZoneTransferOutgoingEnableResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingEnableResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingEnableResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingEnableResponseEnvelopeMessages]
type zoneTransferOutgoingEnableResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingEnableResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingEnableResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingEnableResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    zoneTransferOutgoingEnableResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingEnableResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [ZoneTransferOutgoingEnableResponseEnvelopeMessagesSource]
type zoneTransferOutgoingEnableResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingEnableResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingEnableResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingEnableResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingEnableResponseEnvelopeSuccessTrue ZoneTransferOutgoingEnableResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingEnableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingEnableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferOutgoingForceNotifyParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r ZoneTransferOutgoingForceNotifyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type ZoneTransferOutgoingForceNotifyResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingForceNotifyResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingForceNotifyResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingForceNotifyResponseEnvelopeSuccess `json:"success,required"`
	// When force_notify query parameter is set to true, the response is a simple
	// string.
	Result string                                              `json:"result"`
	JSON   zoneTransferOutgoingForceNotifyResponseEnvelopeJSON `json:"-"`
}

// zoneTransferOutgoingForceNotifyResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZoneTransferOutgoingForceNotifyResponseEnvelope]
type zoneTransferOutgoingForceNotifyResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingForceNotifyResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingForceNotifyResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingForceNotifyResponseEnvelopeErrors struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           ZoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingForceNotifyResponseEnvelopeErrors]
type zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingForceNotifyResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [ZoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSource]
type zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingForceNotifyResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingForceNotifyResponseEnvelopeMessages struct {
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           ZoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZoneTransferOutgoingForceNotifyResponseEnvelopeMessages]
type zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingForceNotifyResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [ZoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSource]
type zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingForceNotifyResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingForceNotifyResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingForceNotifyResponseEnvelopeSuccessTrue ZoneTransferOutgoingForceNotifyResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingForceNotifyResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingForceNotifyResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferOutgoingGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferOutgoingGetResponseEnvelope struct {
	Errors   []ZoneTransferOutgoingGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferOutgoingGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferOutgoingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferOutgoingGetResponse                `json:"result"`
	JSON    zoneTransferOutgoingGetResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferOutgoingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferOutgoingGetResponseEnvelope]
type zoneTransferOutgoingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ZoneTransferOutgoingGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferOutgoingGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferOutgoingGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneTransferOutgoingGetResponseEnvelopeErrors]
type zoneTransferOutgoingGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    zoneTransferOutgoingGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferOutgoingGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingGetResponseEnvelopeErrorsSource]
type zoneTransferOutgoingGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           ZoneTransferOutgoingGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferOutgoingGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferOutgoingGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZoneTransferOutgoingGetResponseEnvelopeMessages]
type zoneTransferOutgoingGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferOutgoingGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferOutgoingGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    zoneTransferOutgoingGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferOutgoingGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneTransferOutgoingGetResponseEnvelopeMessagesSource]
type zoneTransferOutgoingGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferOutgoingGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferOutgoingGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferOutgoingGetResponseEnvelopeSuccess bool

const (
	ZoneTransferOutgoingGetResponseEnvelopeSuccessTrue ZoneTransferOutgoingGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferOutgoingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferOutgoingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
