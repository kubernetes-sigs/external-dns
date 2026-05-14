// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// ConnectorService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorService] method instead.
type ConnectorService struct {
	Options   []option.RequestOption
	Events    *ConnectorEventService
	Snapshots *ConnectorSnapshotService
}

// NewConnectorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorService(opts ...option.RequestOption) (r *ConnectorService) {
	r = &ConnectorService{}
	r.Options = opts
	r.Events = NewConnectorEventService(opts...)
	r.Snapshots = NewConnectorSnapshotService(opts...)
	return
}

// Add a connector to your account
func (r *ConnectorService) New(ctx context.Context, params ConnectorNewParams, opts ...option.RequestOption) (res *ConnectorNewResponse, err error) {
	var env ConnectorNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Replace Connector
func (r *ConnectorService) Update(ctx context.Context, connectorID string, params ConnectorUpdateParams, opts ...option.RequestOption) (res *ConnectorUpdateResponse, err error) {
	var env ConnectorUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Connectors
func (r *ConnectorService) List(ctx context.Context, query ConnectorListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ConnectorListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors", query.AccountID)
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

// List Connectors
func (r *ConnectorService) ListAutoPaging(ctx context.Context, query ConnectorListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ConnectorListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Remove a connector from your account
func (r *ConnectorService) Delete(ctx context.Context, connectorID string, body ConnectorDeleteParams, opts ...option.RequestOption) (res *ConnectorDeleteResponse, err error) {
	var env ConnectorDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", body.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit Connector to update specific properties
func (r *ConnectorService) Edit(ctx context.Context, connectorID string, params ConnectorEditParams, opts ...option.RequestOption) (res *ConnectorEditResponse, err error) {
	var env ConnectorEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch Connector
func (r *ConnectorService) Get(ctx context.Context, connectorID string, query ConnectorGetParams, opts ...option.RequestOption) (res *ConnectorGetResponse, err error) {
	var env ConnectorGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", query.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ConnectorNewResponse struct {
	ID                           string                     `json:"id,required"`
	Activated                    bool                       `json:"activated,required"`
	InterruptWindowDurationHours float64                    `json:"interrupt_window_duration_hours,required"`
	InterruptWindowHourOfDay     float64                    `json:"interrupt_window_hour_of_day,required"`
	LastUpdated                  string                     `json:"last_updated,required"`
	Notes                        string                     `json:"notes,required"`
	Timezone                     string                     `json:"timezone,required"`
	Device                       ConnectorNewResponseDevice `json:"device"`
	LastHeartbeat                string                     `json:"last_heartbeat"`
	LastSeenVersion              string                     `json:"last_seen_version"`
	JSON                         connectorNewResponseJSON   `json:"-"`
}

// connectorNewResponseJSON contains the JSON metadata for the struct
// [ConnectorNewResponse]
type connectorNewResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewResponseDevice struct {
	ID           string                         `json:"id,required"`
	SerialNumber string                         `json:"serial_number"`
	JSON         connectorNewResponseDeviceJSON `json:"-"`
}

// connectorNewResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorNewResponseDevice]
type connectorNewResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorNewResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponse struct {
	ID                           string                        `json:"id,required"`
	Activated                    bool                          `json:"activated,required"`
	InterruptWindowDurationHours float64                       `json:"interrupt_window_duration_hours,required"`
	InterruptWindowHourOfDay     float64                       `json:"interrupt_window_hour_of_day,required"`
	LastUpdated                  string                        `json:"last_updated,required"`
	Notes                        string                        `json:"notes,required"`
	Timezone                     string                        `json:"timezone,required"`
	Device                       ConnectorUpdateResponseDevice `json:"device"`
	LastHeartbeat                string                        `json:"last_heartbeat"`
	LastSeenVersion              string                        `json:"last_seen_version"`
	JSON                         connectorUpdateResponseJSON   `json:"-"`
}

// connectorUpdateResponseJSON contains the JSON metadata for the struct
// [ConnectorUpdateResponse]
type connectorUpdateResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponseDevice struct {
	ID           string                            `json:"id,required"`
	SerialNumber string                            `json:"serial_number"`
	JSON         connectorUpdateResponseDeviceJSON `json:"-"`
}

// connectorUpdateResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorUpdateResponseDevice]
type connectorUpdateResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorUpdateResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorListResponse struct {
	ID                           string                      `json:"id,required"`
	Activated                    bool                        `json:"activated,required"`
	InterruptWindowDurationHours float64                     `json:"interrupt_window_duration_hours,required"`
	InterruptWindowHourOfDay     float64                     `json:"interrupt_window_hour_of_day,required"`
	LastUpdated                  string                      `json:"last_updated,required"`
	Notes                        string                      `json:"notes,required"`
	Timezone                     string                      `json:"timezone,required"`
	Device                       ConnectorListResponseDevice `json:"device"`
	LastHeartbeat                string                      `json:"last_heartbeat"`
	LastSeenVersion              string                      `json:"last_seen_version"`
	JSON                         connectorListResponseJSON   `json:"-"`
}

// connectorListResponseJSON contains the JSON metadata for the struct
// [ConnectorListResponse]
type connectorListResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorListResponseDevice struct {
	ID           string                          `json:"id,required"`
	SerialNumber string                          `json:"serial_number"`
	JSON         connectorListResponseDeviceJSON `json:"-"`
}

// connectorListResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorListResponseDevice]
type connectorListResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorListResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorListResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponse struct {
	ID                           string                        `json:"id,required"`
	Activated                    bool                          `json:"activated,required"`
	InterruptWindowDurationHours float64                       `json:"interrupt_window_duration_hours,required"`
	InterruptWindowHourOfDay     float64                       `json:"interrupt_window_hour_of_day,required"`
	LastUpdated                  string                        `json:"last_updated,required"`
	Notes                        string                        `json:"notes,required"`
	Timezone                     string                        `json:"timezone,required"`
	Device                       ConnectorDeleteResponseDevice `json:"device"`
	LastHeartbeat                string                        `json:"last_heartbeat"`
	LastSeenVersion              string                        `json:"last_seen_version"`
	JSON                         connectorDeleteResponseJSON   `json:"-"`
}

// connectorDeleteResponseJSON contains the JSON metadata for the struct
// [ConnectorDeleteResponse]
type connectorDeleteResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponseDevice struct {
	ID           string                            `json:"id,required"`
	SerialNumber string                            `json:"serial_number"`
	JSON         connectorDeleteResponseDeviceJSON `json:"-"`
}

// connectorDeleteResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorDeleteResponseDevice]
type connectorDeleteResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorDeleteResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponse struct {
	ID                           string                      `json:"id,required"`
	Activated                    bool                        `json:"activated,required"`
	InterruptWindowDurationHours float64                     `json:"interrupt_window_duration_hours,required"`
	InterruptWindowHourOfDay     float64                     `json:"interrupt_window_hour_of_day,required"`
	LastUpdated                  string                      `json:"last_updated,required"`
	Notes                        string                      `json:"notes,required"`
	Timezone                     string                      `json:"timezone,required"`
	Device                       ConnectorEditResponseDevice `json:"device"`
	LastHeartbeat                string                      `json:"last_heartbeat"`
	LastSeenVersion              string                      `json:"last_seen_version"`
	JSON                         connectorEditResponseJSON   `json:"-"`
}

// connectorEditResponseJSON contains the JSON metadata for the struct
// [ConnectorEditResponse]
type connectorEditResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponseDevice struct {
	ID           string                          `json:"id,required"`
	SerialNumber string                          `json:"serial_number"`
	JSON         connectorEditResponseDeviceJSON `json:"-"`
}

// connectorEditResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorEditResponseDevice]
type connectorEditResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorEditResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponse struct {
	ID                           string                     `json:"id,required"`
	Activated                    bool                       `json:"activated,required"`
	InterruptWindowDurationHours float64                    `json:"interrupt_window_duration_hours,required"`
	InterruptWindowHourOfDay     float64                    `json:"interrupt_window_hour_of_day,required"`
	LastUpdated                  string                     `json:"last_updated,required"`
	Notes                        string                     `json:"notes,required"`
	Timezone                     string                     `json:"timezone,required"`
	Device                       ConnectorGetResponseDevice `json:"device"`
	LastHeartbeat                string                     `json:"last_heartbeat"`
	LastSeenVersion              string                     `json:"last_seen_version"`
	JSON                         connectorGetResponseJSON   `json:"-"`
}

// connectorGetResponseJSON contains the JSON metadata for the struct
// [ConnectorGetResponse]
type connectorGetResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponseDevice struct {
	ID           string                         `json:"id,required"`
	SerialNumber string                         `json:"serial_number"`
	JSON         connectorGetResponseDeviceJSON `json:"-"`
}

// connectorGetResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorGetResponseDevice]
type connectorGetResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorGetResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewParams struct {
	// Account identifier
	AccountID                    param.Field[string]                   `path:"account_id,required"`
	Device                       param.Field[ConnectorNewParamsDevice] `json:"device,required"`
	Activated                    param.Field[bool]                     `json:"activated"`
	InterruptWindowDurationHours param.Field[float64]                  `json:"interrupt_window_duration_hours"`
	InterruptWindowHourOfDay     param.Field[float64]                  `json:"interrupt_window_hour_of_day"`
	Notes                        param.Field[string]                   `json:"notes"`
	Timezone                     param.Field[string]                   `json:"timezone"`
}

func (r ConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorNewParamsDevice struct {
	ID           param.Field[string] `json:"id"`
	SerialNumber param.Field[string] `json:"serial_number"`
}

func (r ConnectorNewParamsDevice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorNewResponseEnvelope struct {
	Errors   []ConnectorNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectorNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ConnectorNewResponse                   `json:"result,required"`
	Success  bool                                   `json:"success,required"`
	JSON     connectorNewResponseEnvelopeJSON       `json:"-"`
}

// connectorNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorNewResponseEnvelope]
type connectorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewResponseEnvelopeErrors struct {
	Code    float64                                `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    connectorNewResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConnectorNewResponseEnvelopeErrors]
type connectorNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewResponseEnvelopeMessages struct {
	Code    float64                                  `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    connectorNewResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorNewResponseEnvelopeMessages]
type connectorNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateParams struct {
	// Account identifier
	AccountID                    param.Field[string]  `path:"account_id,required"`
	Activated                    param.Field[bool]    `json:"activated"`
	InterruptWindowDurationHours param.Field[float64] `json:"interrupt_window_duration_hours"`
	InterruptWindowHourOfDay     param.Field[float64] `json:"interrupt_window_hour_of_day"`
	Notes                        param.Field[string]  `json:"notes"`
	Timezone                     param.Field[string]  `json:"timezone"`
}

func (r ConnectorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorUpdateResponseEnvelope struct {
	Errors   []ConnectorUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectorUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ConnectorUpdateResponse                   `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     connectorUpdateResponseEnvelopeJSON       `json:"-"`
}

// connectorUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorUpdateResponseEnvelope]
type connectorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponseEnvelopeErrors struct {
	Code    float64                                   `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    connectorUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorUpdateResponseEnvelopeErrors]
type connectorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponseEnvelopeMessages struct {
	Code    float64                                     `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    connectorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorUpdateResponseEnvelopeMessages]
type connectorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectorDeleteParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectorDeleteResponseEnvelope struct {
	Errors   []ConnectorDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectorDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ConnectorDeleteResponse                   `json:"result,required"`
	Success  bool                                      `json:"success,required"`
	JSON     connectorDeleteResponseEnvelopeJSON       `json:"-"`
}

// connectorDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorDeleteResponseEnvelope]
type connectorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponseEnvelopeErrors struct {
	Code    float64                                   `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    connectorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorDeleteResponseEnvelopeErrors]
type connectorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponseEnvelopeMessages struct {
	Code    float64                                     `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    connectorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorDeleteResponseEnvelopeMessages]
type connectorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditParams struct {
	// Account identifier
	AccountID                    param.Field[string]  `path:"account_id,required"`
	Activated                    param.Field[bool]    `json:"activated"`
	InterruptWindowDurationHours param.Field[float64] `json:"interrupt_window_duration_hours"`
	InterruptWindowHourOfDay     param.Field[float64] `json:"interrupt_window_hour_of_day"`
	Notes                        param.Field[string]  `json:"notes"`
	Timezone                     param.Field[string]  `json:"timezone"`
}

func (r ConnectorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorEditResponseEnvelope struct {
	Errors   []ConnectorEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectorEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ConnectorEditResponse                   `json:"result,required"`
	Success  bool                                    `json:"success,required"`
	JSON     connectorEditResponseEnvelopeJSON       `json:"-"`
}

// connectorEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEditResponseEnvelope]
type connectorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponseEnvelopeErrors struct {
	Code    float64                                 `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    connectorEditResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEditResponseEnvelopeErrors]
type connectorEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponseEnvelopeMessages struct {
	Code    float64                                   `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    connectorEditResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorEditResponseEnvelopeMessages]
type connectorEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectorGetResponseEnvelope struct {
	Errors   []ConnectorGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ConnectorGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ConnectorGetResponse                   `json:"result,required"`
	Success  bool                                   `json:"success,required"`
	JSON     connectorGetResponseEnvelopeJSON       `json:"-"`
}

// connectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorGetResponseEnvelope]
type connectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponseEnvelopeErrors struct {
	Code    float64                                `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    connectorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConnectorGetResponseEnvelopeErrors]
type connectorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponseEnvelopeMessages struct {
	Code    float64                                  `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    connectorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorGetResponseEnvelopeMessages]
type connectorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
