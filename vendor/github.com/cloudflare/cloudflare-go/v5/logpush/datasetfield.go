// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

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

// DatasetFieldService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetFieldService] method instead.
type DatasetFieldService struct {
	Options []option.RequestOption
}

// NewDatasetFieldService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasetFieldService(opts ...option.RequestOption) (r *DatasetFieldService) {
	r = &DatasetFieldService{}
	r.Options = opts
	return
}

// Lists all fields available for a dataset. The response result is. an object with
// key-value pairs, where keys are field names, and values are descriptions.
func (r *DatasetFieldService) Get(ctx context.Context, datasetID DatasetFieldGetParamsDatasetID, query DatasetFieldGetParams, opts ...option.RequestOption) (res *DatasetFieldGetResponse, err error) {
	var env DatasetFieldGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/datasets/%v/fields", accountOrZone, accountOrZoneID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DatasetFieldGetResponse = interface{}

type DatasetFieldGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Name of the dataset. A list of supported datasets can be found on the
// [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
type DatasetFieldGetParamsDatasetID string

const (
	DatasetFieldGetParamsDatasetIDAccessRequests           DatasetFieldGetParamsDatasetID = "access_requests"
	DatasetFieldGetParamsDatasetIDAuditLogs                DatasetFieldGetParamsDatasetID = "audit_logs"
	DatasetFieldGetParamsDatasetIDAuditLogsV2              DatasetFieldGetParamsDatasetID = "audit_logs_v2"
	DatasetFieldGetParamsDatasetIDBISOUserActions          DatasetFieldGetParamsDatasetID = "biso_user_actions"
	DatasetFieldGetParamsDatasetIDCasbFindings             DatasetFieldGetParamsDatasetID = "casb_findings"
	DatasetFieldGetParamsDatasetIDDevicePostureResults     DatasetFieldGetParamsDatasetID = "device_posture_results"
	DatasetFieldGetParamsDatasetIDDLPForensicCopies        DatasetFieldGetParamsDatasetID = "dlp_forensic_copies"
	DatasetFieldGetParamsDatasetIDDNSFirewallLogs          DatasetFieldGetParamsDatasetID = "dns_firewall_logs"
	DatasetFieldGetParamsDatasetIDDNSLogs                  DatasetFieldGetParamsDatasetID = "dns_logs"
	DatasetFieldGetParamsDatasetIDEmailSecurityAlerts      DatasetFieldGetParamsDatasetID = "email_security_alerts"
	DatasetFieldGetParamsDatasetIDFirewallEvents           DatasetFieldGetParamsDatasetID = "firewall_events"
	DatasetFieldGetParamsDatasetIDGatewayDNS               DatasetFieldGetParamsDatasetID = "gateway_dns"
	DatasetFieldGetParamsDatasetIDGatewayHTTP              DatasetFieldGetParamsDatasetID = "gateway_http"
	DatasetFieldGetParamsDatasetIDGatewayNetwork           DatasetFieldGetParamsDatasetID = "gateway_network"
	DatasetFieldGetParamsDatasetIDHTTPRequests             DatasetFieldGetParamsDatasetID = "http_requests"
	DatasetFieldGetParamsDatasetIDMagicIDsDetections       DatasetFieldGetParamsDatasetID = "magic_ids_detections"
	DatasetFieldGetParamsDatasetIDNELReports               DatasetFieldGetParamsDatasetID = "nel_reports"
	DatasetFieldGetParamsDatasetIDNetworkAnalyticsLogs     DatasetFieldGetParamsDatasetID = "network_analytics_logs"
	DatasetFieldGetParamsDatasetIDPageShieldEvents         DatasetFieldGetParamsDatasetID = "page_shield_events"
	DatasetFieldGetParamsDatasetIDSinkholeHTTPLogs         DatasetFieldGetParamsDatasetID = "sinkhole_http_logs"
	DatasetFieldGetParamsDatasetIDSpectrumEvents           DatasetFieldGetParamsDatasetID = "spectrum_events"
	DatasetFieldGetParamsDatasetIDSSHLogs                  DatasetFieldGetParamsDatasetID = "ssh_logs"
	DatasetFieldGetParamsDatasetIDWorkersTraceEvents       DatasetFieldGetParamsDatasetID = "workers_trace_events"
	DatasetFieldGetParamsDatasetIDZarazEvents              DatasetFieldGetParamsDatasetID = "zaraz_events"
	DatasetFieldGetParamsDatasetIDZeroTrustNetworkSessions DatasetFieldGetParamsDatasetID = "zero_trust_network_sessions"
)

func (r DatasetFieldGetParamsDatasetID) IsKnown() bool {
	switch r {
	case DatasetFieldGetParamsDatasetIDAccessRequests, DatasetFieldGetParamsDatasetIDAuditLogs, DatasetFieldGetParamsDatasetIDAuditLogsV2, DatasetFieldGetParamsDatasetIDBISOUserActions, DatasetFieldGetParamsDatasetIDCasbFindings, DatasetFieldGetParamsDatasetIDDevicePostureResults, DatasetFieldGetParamsDatasetIDDLPForensicCopies, DatasetFieldGetParamsDatasetIDDNSFirewallLogs, DatasetFieldGetParamsDatasetIDDNSLogs, DatasetFieldGetParamsDatasetIDEmailSecurityAlerts, DatasetFieldGetParamsDatasetIDFirewallEvents, DatasetFieldGetParamsDatasetIDGatewayDNS, DatasetFieldGetParamsDatasetIDGatewayHTTP, DatasetFieldGetParamsDatasetIDGatewayNetwork, DatasetFieldGetParamsDatasetIDHTTPRequests, DatasetFieldGetParamsDatasetIDMagicIDsDetections, DatasetFieldGetParamsDatasetIDNELReports, DatasetFieldGetParamsDatasetIDNetworkAnalyticsLogs, DatasetFieldGetParamsDatasetIDPageShieldEvents, DatasetFieldGetParamsDatasetIDSinkholeHTTPLogs, DatasetFieldGetParamsDatasetIDSpectrumEvents, DatasetFieldGetParamsDatasetIDSSHLogs, DatasetFieldGetParamsDatasetIDWorkersTraceEvents, DatasetFieldGetParamsDatasetIDZarazEvents, DatasetFieldGetParamsDatasetIDZeroTrustNetworkSessions:
		return true
	}
	return false
}

type DatasetFieldGetResponseEnvelope struct {
	Errors   []DatasetFieldGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DatasetFieldGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DatasetFieldGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DatasetFieldGetResponse                `json:"result"`
	JSON    datasetFieldGetResponseEnvelopeJSON    `json:"-"`
}

// datasetFieldGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetFieldGetResponseEnvelope]
type datasetFieldGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DatasetFieldGetResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DatasetFieldGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             datasetFieldGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// datasetFieldGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DatasetFieldGetResponseEnvelopeErrors]
type datasetFieldGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DatasetFieldGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    datasetFieldGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// datasetFieldGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DatasetFieldGetResponseEnvelopeErrorsSource]
type datasetFieldGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DatasetFieldGetResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DatasetFieldGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             datasetFieldGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// datasetFieldGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DatasetFieldGetResponseEnvelopeMessages]
type datasetFieldGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DatasetFieldGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    datasetFieldGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// datasetFieldGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DatasetFieldGetResponseEnvelopeMessagesSource]
type datasetFieldGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetFieldGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetFieldGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DatasetFieldGetResponseEnvelopeSuccess bool

const (
	DatasetFieldGetResponseEnvelopeSuccessTrue DatasetFieldGetResponseEnvelopeSuccess = true
)

func (r DatasetFieldGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatasetFieldGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
