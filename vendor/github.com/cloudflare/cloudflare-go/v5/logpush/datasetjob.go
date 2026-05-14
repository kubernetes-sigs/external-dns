// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DatasetJobService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetJobService] method instead.
type DatasetJobService struct {
	Options []option.RequestOption
}

// NewDatasetJobService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatasetJobService(opts ...option.RequestOption) (r *DatasetJobService) {
	r = &DatasetJobService{}
	r.Options = opts
	return
}

// Lists Logpush jobs for an account or zone for a dataset.
func (r *DatasetJobService) Get(ctx context.Context, datasetID DatasetJobGetParamsDatasetID, query DatasetJobGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[LogpushJob], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
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
	path := fmt.Sprintf("%s/%s/logpush/datasets/%v/jobs", accountOrZone, accountOrZoneID, datasetID)
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

// Lists Logpush jobs for an account or zone for a dataset.
func (r *DatasetJobService) GetAutoPaging(ctx context.Context, datasetID DatasetJobGetParamsDatasetID, query DatasetJobGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LogpushJob] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, datasetID, query, opts...))
}

type DatasetJobGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// Name of the dataset. A list of supported datasets can be found on the
// [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
type DatasetJobGetParamsDatasetID string

const (
	DatasetJobGetParamsDatasetIDAccessRequests           DatasetJobGetParamsDatasetID = "access_requests"
	DatasetJobGetParamsDatasetIDAuditLogs                DatasetJobGetParamsDatasetID = "audit_logs"
	DatasetJobGetParamsDatasetIDAuditLogsV2              DatasetJobGetParamsDatasetID = "audit_logs_v2"
	DatasetJobGetParamsDatasetIDBISOUserActions          DatasetJobGetParamsDatasetID = "biso_user_actions"
	DatasetJobGetParamsDatasetIDCasbFindings             DatasetJobGetParamsDatasetID = "casb_findings"
	DatasetJobGetParamsDatasetIDDevicePostureResults     DatasetJobGetParamsDatasetID = "device_posture_results"
	DatasetJobGetParamsDatasetIDDLPForensicCopies        DatasetJobGetParamsDatasetID = "dlp_forensic_copies"
	DatasetJobGetParamsDatasetIDDNSFirewallLogs          DatasetJobGetParamsDatasetID = "dns_firewall_logs"
	DatasetJobGetParamsDatasetIDDNSLogs                  DatasetJobGetParamsDatasetID = "dns_logs"
	DatasetJobGetParamsDatasetIDEmailSecurityAlerts      DatasetJobGetParamsDatasetID = "email_security_alerts"
	DatasetJobGetParamsDatasetIDFirewallEvents           DatasetJobGetParamsDatasetID = "firewall_events"
	DatasetJobGetParamsDatasetIDGatewayDNS               DatasetJobGetParamsDatasetID = "gateway_dns"
	DatasetJobGetParamsDatasetIDGatewayHTTP              DatasetJobGetParamsDatasetID = "gateway_http"
	DatasetJobGetParamsDatasetIDGatewayNetwork           DatasetJobGetParamsDatasetID = "gateway_network"
	DatasetJobGetParamsDatasetIDHTTPRequests             DatasetJobGetParamsDatasetID = "http_requests"
	DatasetJobGetParamsDatasetIDMagicIDsDetections       DatasetJobGetParamsDatasetID = "magic_ids_detections"
	DatasetJobGetParamsDatasetIDNELReports               DatasetJobGetParamsDatasetID = "nel_reports"
	DatasetJobGetParamsDatasetIDNetworkAnalyticsLogs     DatasetJobGetParamsDatasetID = "network_analytics_logs"
	DatasetJobGetParamsDatasetIDPageShieldEvents         DatasetJobGetParamsDatasetID = "page_shield_events"
	DatasetJobGetParamsDatasetIDSinkholeHTTPLogs         DatasetJobGetParamsDatasetID = "sinkhole_http_logs"
	DatasetJobGetParamsDatasetIDSpectrumEvents           DatasetJobGetParamsDatasetID = "spectrum_events"
	DatasetJobGetParamsDatasetIDSSHLogs                  DatasetJobGetParamsDatasetID = "ssh_logs"
	DatasetJobGetParamsDatasetIDWorkersTraceEvents       DatasetJobGetParamsDatasetID = "workers_trace_events"
	DatasetJobGetParamsDatasetIDZarazEvents              DatasetJobGetParamsDatasetID = "zaraz_events"
	DatasetJobGetParamsDatasetIDZeroTrustNetworkSessions DatasetJobGetParamsDatasetID = "zero_trust_network_sessions"
)

func (r DatasetJobGetParamsDatasetID) IsKnown() bool {
	switch r {
	case DatasetJobGetParamsDatasetIDAccessRequests, DatasetJobGetParamsDatasetIDAuditLogs, DatasetJobGetParamsDatasetIDAuditLogsV2, DatasetJobGetParamsDatasetIDBISOUserActions, DatasetJobGetParamsDatasetIDCasbFindings, DatasetJobGetParamsDatasetIDDevicePostureResults, DatasetJobGetParamsDatasetIDDLPForensicCopies, DatasetJobGetParamsDatasetIDDNSFirewallLogs, DatasetJobGetParamsDatasetIDDNSLogs, DatasetJobGetParamsDatasetIDEmailSecurityAlerts, DatasetJobGetParamsDatasetIDFirewallEvents, DatasetJobGetParamsDatasetIDGatewayDNS, DatasetJobGetParamsDatasetIDGatewayHTTP, DatasetJobGetParamsDatasetIDGatewayNetwork, DatasetJobGetParamsDatasetIDHTTPRequests, DatasetJobGetParamsDatasetIDMagicIDsDetections, DatasetJobGetParamsDatasetIDNELReports, DatasetJobGetParamsDatasetIDNetworkAnalyticsLogs, DatasetJobGetParamsDatasetIDPageShieldEvents, DatasetJobGetParamsDatasetIDSinkholeHTTPLogs, DatasetJobGetParamsDatasetIDSpectrumEvents, DatasetJobGetParamsDatasetIDSSHLogs, DatasetJobGetParamsDatasetIDWorkersTraceEvents, DatasetJobGetParamsDatasetIDZarazEvents, DatasetJobGetParamsDatasetIDZeroTrustNetworkSessions:
		return true
	}
	return false
}
