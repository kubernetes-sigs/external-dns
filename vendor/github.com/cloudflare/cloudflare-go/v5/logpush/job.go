// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpush

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

// JobService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r *JobService) {
	r = &JobService{}
	r.Options = opts
	return
}

// Creates a new Logpush job for an account or zone.
func (r *JobService) New(ctx context.Context, params JobNewParams, opts ...option.RequestOption) (res *LogpushJob, err error) {
	var env JobNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a Logpush job.
func (r *JobService) Update(ctx context.Context, jobID int64, params JobUpdateParams, opts ...option.RequestOption) (res *LogpushJob, err error) {
	var env JobUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists Logpush jobs for an account or zone.
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *pagination.SinglePage[LogpushJob], err error) {
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
	path := fmt.Sprintf("%s/%s/logpush/jobs", accountOrZone, accountOrZoneID)
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

// Lists Logpush jobs for an account or zone.
func (r *JobService) ListAutoPaging(ctx context.Context, query JobListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[LogpushJob] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Logpush job.
func (r *JobService) Delete(ctx context.Context, jobID int64, body JobDeleteParams, opts ...option.RequestOption) (res *JobDeleteResponse, err error) {
	var env JobDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of a Logpush job.
func (r *JobService) Get(ctx context.Context, jobID int64, query JobGetParams, opts ...option.RequestOption) (res *LogpushJob, err error) {
	var env JobGetResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/logpush/jobs/%v", accountOrZone, accountOrZoneID, jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LogpushJob struct {
	// Unique id of the job.
	ID int64 `json:"id"`
	// Name of the dataset. A list of supported datasets can be found on the
	// [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	Dataset LogpushJobDataset `json:"dataset,nullable"`
	// Uniquely identifies a resource (such as an s3 bucket) where data. will be
	// pushed. Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf string `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled bool `json:"enabled"`
	// If not null, the job is currently failing. Failures are usually. repetitive
	// (example: no permissions to write to destination bucket). Only the last failure
	// is recorded. On successful execution of a job the error_message and last_error
	// are set to null.
	ErrorMessage string `json:"error_message,nullable"`
	// This field is deprecated. Please use `max_upload_*` parameters instead. . The
	// frequency at which Cloudflare sends batches of logs to your destination. Setting
	// frequency to high sends your logs in larger quantities of smaller files. Setting
	// frequency to low sends logs in smaller quantities of larger files.
	//
	// Deprecated: deprecated
	Frequency LogpushJobFrequency `json:"frequency,nullable"`
	// The kind parameter (optional) is used to differentiate between Logpush and Edge
	// Log Delivery jobs (when supported by the dataset).
	Kind LogpushJobKind `json:"kind"`
	// Records the last time for which logs have been successfully pushed. If the last
	// successful push was for logs range 2018-07-23T10:00:00Z to 2018-07-23T10:01:00Z
	// then the value of this field will be 2018-07-23T10:01:00Z. If the job has never
	// run or has just been enabled and hasn't run yet then the field will be empty.
	LastComplete time.Time `json:"last_complete,nullable" format:"date-time"`
	// Records the last time the job failed. If not null, the job is currently.
	// failing. If null, the job has either never failed or has run successfully at
	// least once since last failure. See also the error_message field.
	LastError time.Time `json:"last_error,nullable" format:"date-time"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	//
	// Deprecated: deprecated
	LogpullOptions string `json:"logpull_options,nullable" format:"uri-reference"`
	// The maximum uncompressed file size of a batch of logs. This setting value must
	// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
	// minimum file size; this means that log files may be much smaller than this batch
	// size.
	MaxUploadBytes LogpushJobMaxUploadBytes `json:"max_upload_bytes,nullable"`
	// The maximum interval in seconds for log batches. This setting must be between 30
	// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
	// a minimum interval for log batches; this means that log files may be sent in
	// shorter intervals than this.
	MaxUploadIntervalSeconds LogpushJobMaxUploadIntervalSeconds `json:"max_upload_interval_seconds,nullable"`
	// The maximum number of log lines per batch. This setting must be between 1000 and
	// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
	// number of log lines per batch; this means that log files may contain many fewer
	// lines than this.
	MaxUploadRecords LogpushJobMaxUploadRecords `json:"max_upload_records,nullable"`
	// Optional human readable job name. Not unique. Cloudflare suggests. that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name string `json:"name,nullable"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions OutputOptions  `json:"output_options,nullable"`
	JSON          logpushJobJSON `json:"-"`
}

// logpushJobJSON contains the JSON metadata for the struct [LogpushJob]
type logpushJobJSON struct {
	ID                       apijson.Field
	Dataset                  apijson.Field
	DestinationConf          apijson.Field
	Enabled                  apijson.Field
	ErrorMessage             apijson.Field
	Frequency                apijson.Field
	Kind                     apijson.Field
	LastComplete             apijson.Field
	LastError                apijson.Field
	LogpullOptions           apijson.Field
	MaxUploadBytes           apijson.Field
	MaxUploadIntervalSeconds apijson.Field
	MaxUploadRecords         apijson.Field
	Name                     apijson.Field
	OutputOptions            apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *LogpushJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logpushJobJSON) RawJSON() string {
	return r.raw
}

// Name of the dataset. A list of supported datasets can be found on the
// [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
type LogpushJobDataset string

const (
	LogpushJobDatasetAccessRequests           LogpushJobDataset = "access_requests"
	LogpushJobDatasetAuditLogs                LogpushJobDataset = "audit_logs"
	LogpushJobDatasetAuditLogsV2              LogpushJobDataset = "audit_logs_v2"
	LogpushJobDatasetBISOUserActions          LogpushJobDataset = "biso_user_actions"
	LogpushJobDatasetCasbFindings             LogpushJobDataset = "casb_findings"
	LogpushJobDatasetDevicePostureResults     LogpushJobDataset = "device_posture_results"
	LogpushJobDatasetDLPForensicCopies        LogpushJobDataset = "dlp_forensic_copies"
	LogpushJobDatasetDNSFirewallLogs          LogpushJobDataset = "dns_firewall_logs"
	LogpushJobDatasetDNSLogs                  LogpushJobDataset = "dns_logs"
	LogpushJobDatasetEmailSecurityAlerts      LogpushJobDataset = "email_security_alerts"
	LogpushJobDatasetFirewallEvents           LogpushJobDataset = "firewall_events"
	LogpushJobDatasetGatewayDNS               LogpushJobDataset = "gateway_dns"
	LogpushJobDatasetGatewayHTTP              LogpushJobDataset = "gateway_http"
	LogpushJobDatasetGatewayNetwork           LogpushJobDataset = "gateway_network"
	LogpushJobDatasetHTTPRequests             LogpushJobDataset = "http_requests"
	LogpushJobDatasetMagicIDsDetections       LogpushJobDataset = "magic_ids_detections"
	LogpushJobDatasetNELReports               LogpushJobDataset = "nel_reports"
	LogpushJobDatasetNetworkAnalyticsLogs     LogpushJobDataset = "network_analytics_logs"
	LogpushJobDatasetPageShieldEvents         LogpushJobDataset = "page_shield_events"
	LogpushJobDatasetSinkholeHTTPLogs         LogpushJobDataset = "sinkhole_http_logs"
	LogpushJobDatasetSpectrumEvents           LogpushJobDataset = "spectrum_events"
	LogpushJobDatasetSSHLogs                  LogpushJobDataset = "ssh_logs"
	LogpushJobDatasetWorkersTraceEvents       LogpushJobDataset = "workers_trace_events"
	LogpushJobDatasetZarazEvents              LogpushJobDataset = "zaraz_events"
	LogpushJobDatasetZeroTrustNetworkSessions LogpushJobDataset = "zero_trust_network_sessions"
)

func (r LogpushJobDataset) IsKnown() bool {
	switch r {
	case LogpushJobDatasetAccessRequests, LogpushJobDatasetAuditLogs, LogpushJobDatasetAuditLogsV2, LogpushJobDatasetBISOUserActions, LogpushJobDatasetCasbFindings, LogpushJobDatasetDevicePostureResults, LogpushJobDatasetDLPForensicCopies, LogpushJobDatasetDNSFirewallLogs, LogpushJobDatasetDNSLogs, LogpushJobDatasetEmailSecurityAlerts, LogpushJobDatasetFirewallEvents, LogpushJobDatasetGatewayDNS, LogpushJobDatasetGatewayHTTP, LogpushJobDatasetGatewayNetwork, LogpushJobDatasetHTTPRequests, LogpushJobDatasetMagicIDsDetections, LogpushJobDatasetNELReports, LogpushJobDatasetNetworkAnalyticsLogs, LogpushJobDatasetPageShieldEvents, LogpushJobDatasetSinkholeHTTPLogs, LogpushJobDatasetSpectrumEvents, LogpushJobDatasetSSHLogs, LogpushJobDatasetWorkersTraceEvents, LogpushJobDatasetZarazEvents, LogpushJobDatasetZeroTrustNetworkSessions:
		return true
	}
	return false
}

// This field is deprecated. Please use `max_upload_*` parameters instead. . The
// frequency at which Cloudflare sends batches of logs to your destination. Setting
// frequency to high sends your logs in larger quantities of smaller files. Setting
// frequency to low sends logs in smaller quantities of larger files.
type LogpushJobFrequency string

const (
	LogpushJobFrequencyHigh LogpushJobFrequency = "high"
	LogpushJobFrequencyLow  LogpushJobFrequency = "low"
)

func (r LogpushJobFrequency) IsKnown() bool {
	switch r {
	case LogpushJobFrequencyHigh, LogpushJobFrequencyLow:
		return true
	}
	return false
}

// The kind parameter (optional) is used to differentiate between Logpush and Edge
// Log Delivery jobs (when supported by the dataset).
type LogpushJobKind string

const (
	LogpushJobKindEmpty LogpushJobKind = ""
	LogpushJobKindEdge  LogpushJobKind = "edge"
)

func (r LogpushJobKind) IsKnown() bool {
	switch r {
	case LogpushJobKindEmpty, LogpushJobKindEdge:
		return true
	}
	return false
}

// The maximum uncompressed file size of a batch of logs. This setting value must
// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
// minimum file size; this means that log files may be much smaller than this batch
// size.
type LogpushJobMaxUploadBytes int64

const (
	LogpushJobMaxUploadBytes0 LogpushJobMaxUploadBytes = 0
)

func (r LogpushJobMaxUploadBytes) IsKnown() bool {
	switch r {
	case LogpushJobMaxUploadBytes0:
		return true
	}
	return false
}

// The maximum interval in seconds for log batches. This setting must be between 30
// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
// a minimum interval for log batches; this means that log files may be sent in
// shorter intervals than this.
type LogpushJobMaxUploadIntervalSeconds int64

const (
	LogpushJobMaxUploadIntervalSeconds0 LogpushJobMaxUploadIntervalSeconds = 0
)

func (r LogpushJobMaxUploadIntervalSeconds) IsKnown() bool {
	switch r {
	case LogpushJobMaxUploadIntervalSeconds0:
		return true
	}
	return false
}

// The maximum number of log lines per batch. This setting must be between 1000 and
// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
// number of log lines per batch; this means that log files may contain many fewer
// lines than this.
type LogpushJobMaxUploadRecords int64

const (
	LogpushJobMaxUploadRecords0 LogpushJobMaxUploadRecords = 0
)

func (r LogpushJobMaxUploadRecords) IsKnown() bool {
	switch r {
	case LogpushJobMaxUploadRecords0:
		return true
	}
	return false
}

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type OutputOptions struct {
	// String to be prepended before each batch.
	BatchPrefix string `json:"batch_prefix,nullable"`
	// String to be appended after each batch.
	BatchSuffix string `json:"batch_suffix,nullable"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_44228 bool `json:"CVE-2021-44228,nullable"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter string `json:"field_delimiter,nullable"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames []string `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType OutputOptionsOutputType `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter string `json:"record_delimiter,nullable"`
	// String to be prepended before each record.
	RecordPrefix string `json:"record_prefix,nullable"`
	// String to be appended after each record.
	RecordSuffix string `json:"record_suffix,nullable"`
	// String to use as template for each record instead of the default json key value
	// mapping. All fields used in the template must be present in `field_names` as
	// well, otherwise they will end up as null. Format as a Go `text/template` without
	// any standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate string `json:"record_template,nullable"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate float64 `json:"sample_rate,nullable"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat OutputOptionsTimestampFormat `json:"timestamp_format"`
	JSON            outputOptionsJSON            `json:"-"`
}

// outputOptionsJSON contains the JSON metadata for the struct [OutputOptions]
type outputOptionsJSON struct {
	BatchPrefix     apijson.Field
	BatchSuffix     apijson.Field
	Cve2021_44228   apijson.Field
	FieldDelimiter  apijson.Field
	FieldNames      apijson.Field
	OutputType      apijson.Field
	RecordDelimiter apijson.Field
	RecordPrefix    apijson.Field
	RecordSuffix    apijson.Field
	RecordTemplate  apijson.Field
	SampleRate      apijson.Field
	TimestampFormat apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *OutputOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r outputOptionsJSON) RawJSON() string {
	return r.raw
}

// Specifies the output type, such as `ndjson` or `csv`. This sets default values
// for the rest of the settings, depending on the chosen output type. Some
// formatting rules, like string quoting, are different between output types.
type OutputOptionsOutputType string

const (
	OutputOptionsOutputTypeNdjson OutputOptionsOutputType = "ndjson"
	OutputOptionsOutputTypeCsv    OutputOptionsOutputType = "csv"
)

func (r OutputOptionsOutputType) IsKnown() bool {
	switch r {
	case OutputOptionsOutputTypeNdjson, OutputOptionsOutputTypeCsv:
		return true
	}
	return false
}

// String to specify the format for timestamps, such as `unixnano`, `unix`, or
// `rfc3339`.
type OutputOptionsTimestampFormat string

const (
	OutputOptionsTimestampFormatUnixnano OutputOptionsTimestampFormat = "unixnano"
	OutputOptionsTimestampFormatUnix     OutputOptionsTimestampFormat = "unix"
	OutputOptionsTimestampFormatRfc3339  OutputOptionsTimestampFormat = "rfc3339"
)

func (r OutputOptionsTimestampFormat) IsKnown() bool {
	switch r {
	case OutputOptionsTimestampFormatUnixnano, OutputOptionsTimestampFormatUnix, OutputOptionsTimestampFormatRfc3339:
		return true
	}
	return false
}

// The structured replacement for `logpull_options`. When including this field, the
// `logpull_option` field will be ignored.
type OutputOptionsParam struct {
	// String to be prepended before each batch.
	BatchPrefix param.Field[string] `json:"batch_prefix"`
	// String to be appended after each batch.
	BatchSuffix param.Field[string] `json:"batch_suffix"`
	// If set to true, will cause all occurrences of `${` in the generated files to be
	// replaced with `x{`.
	Cve2021_44228 param.Field[bool] `json:"CVE-2021-44228"`
	// String to join fields. This field be ignored when `record_template` is set.
	FieldDelimiter param.Field[string] `json:"field_delimiter"`
	// List of field names to be included in the Logpush output. For the moment, there
	// is no option to add all fields at once, so you must specify all the fields names
	// you are interested in.
	FieldNames param.Field[[]string] `json:"field_names"`
	// Specifies the output type, such as `ndjson` or `csv`. This sets default values
	// for the rest of the settings, depending on the chosen output type. Some
	// formatting rules, like string quoting, are different between output types.
	OutputType param.Field[OutputOptionsOutputType] `json:"output_type"`
	// String to be inserted in-between the records as separator.
	RecordDelimiter param.Field[string] `json:"record_delimiter"`
	// String to be prepended before each record.
	RecordPrefix param.Field[string] `json:"record_prefix"`
	// String to be appended after each record.
	RecordSuffix param.Field[string] `json:"record_suffix"`
	// String to use as template for each record instead of the default json key value
	// mapping. All fields used in the template must be present in `field_names` as
	// well, otherwise they will end up as null. Format as a Go `text/template` without
	// any standard functions, like conditionals, loops, sub-templates, etc.
	RecordTemplate param.Field[string] `json:"record_template"`
	// Floating number to specify sampling rate. Sampling is applied on top of
	// filtering, and regardless of the current `sample_interval` of the data.
	SampleRate param.Field[float64] `json:"sample_rate"`
	// String to specify the format for timestamps, such as `unixnano`, `unix`, or
	// `rfc3339`.
	TimestampFormat param.Field[OutputOptionsTimestampFormat] `json:"timestamp_format"`
}

func (r OutputOptionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type JobDeleteResponse struct {
	// Unique id of the job.
	ID   int64                 `json:"id"`
	JSON jobDeleteResponseJSON `json:"-"`
}

// jobDeleteResponseJSON contains the JSON metadata for the struct
// [JobDeleteResponse]
type jobDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type JobNewParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data. will be
	// pushed. Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Name of the dataset. A list of supported datasets can be found on the
	// [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
	Dataset param.Field[JobNewParamsDataset] `json:"dataset"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The filters to select the events to include and/or remove from your logs. For
	// more information, refer to
	// [Filters](https://developers.cloudflare.com/logs/reference/filters/).
	Filter param.Field[string] `json:"filter"`
	// This field is deprecated. Please use `max_upload_*` parameters instead. . The
	// frequency at which Cloudflare sends batches of logs to your destination. Setting
	// frequency to high sends your logs in larger quantities of smaller files. Setting
	// frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[JobNewParamsFrequency] `json:"frequency"`
	// The kind parameter (optional) is used to differentiate between Logpush and Edge
	// Log Delivery jobs (when supported by the dataset).
	Kind param.Field[JobNewParamsKind] `json:"kind"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// The maximum uncompressed file size of a batch of logs. This setting value must
	// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
	// minimum file size; this means that log files may be much smaller than this batch
	// size.
	MaxUploadBytes param.Field[JobNewParamsMaxUploadBytes] `json:"max_upload_bytes"`
	// The maximum interval in seconds for log batches. This setting must be between 30
	// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
	// a minimum interval for log batches; this means that log files may be sent in
	// shorter intervals than this.
	MaxUploadIntervalSeconds param.Field[JobNewParamsMaxUploadIntervalSeconds] `json:"max_upload_interval_seconds"`
	// The maximum number of log lines per batch. This setting must be between 1000 and
	// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
	// number of log lines per batch; this means that log files may contain many fewer
	// lines than this.
	MaxUploadRecords param.Field[JobNewParamsMaxUploadRecords] `json:"max_upload_records"`
	// Optional human readable job name. Not unique. Cloudflare suggests. that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name param.Field[string] `json:"name"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[OutputOptionsParam] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r JobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Name of the dataset. A list of supported datasets can be found on the
// [Developer Docs](https://developers.cloudflare.com/logs/reference/log-fields/).
type JobNewParamsDataset string

const (
	JobNewParamsDatasetAccessRequests           JobNewParamsDataset = "access_requests"
	JobNewParamsDatasetAuditLogs                JobNewParamsDataset = "audit_logs"
	JobNewParamsDatasetAuditLogsV2              JobNewParamsDataset = "audit_logs_v2"
	JobNewParamsDatasetBISOUserActions          JobNewParamsDataset = "biso_user_actions"
	JobNewParamsDatasetCasbFindings             JobNewParamsDataset = "casb_findings"
	JobNewParamsDatasetDevicePostureResults     JobNewParamsDataset = "device_posture_results"
	JobNewParamsDatasetDLPForensicCopies        JobNewParamsDataset = "dlp_forensic_copies"
	JobNewParamsDatasetDNSFirewallLogs          JobNewParamsDataset = "dns_firewall_logs"
	JobNewParamsDatasetDNSLogs                  JobNewParamsDataset = "dns_logs"
	JobNewParamsDatasetEmailSecurityAlerts      JobNewParamsDataset = "email_security_alerts"
	JobNewParamsDatasetFirewallEvents           JobNewParamsDataset = "firewall_events"
	JobNewParamsDatasetGatewayDNS               JobNewParamsDataset = "gateway_dns"
	JobNewParamsDatasetGatewayHTTP              JobNewParamsDataset = "gateway_http"
	JobNewParamsDatasetGatewayNetwork           JobNewParamsDataset = "gateway_network"
	JobNewParamsDatasetHTTPRequests             JobNewParamsDataset = "http_requests"
	JobNewParamsDatasetMagicIDsDetections       JobNewParamsDataset = "magic_ids_detections"
	JobNewParamsDatasetNELReports               JobNewParamsDataset = "nel_reports"
	JobNewParamsDatasetNetworkAnalyticsLogs     JobNewParamsDataset = "network_analytics_logs"
	JobNewParamsDatasetPageShieldEvents         JobNewParamsDataset = "page_shield_events"
	JobNewParamsDatasetSinkholeHTTPLogs         JobNewParamsDataset = "sinkhole_http_logs"
	JobNewParamsDatasetSpectrumEvents           JobNewParamsDataset = "spectrum_events"
	JobNewParamsDatasetSSHLogs                  JobNewParamsDataset = "ssh_logs"
	JobNewParamsDatasetWorkersTraceEvents       JobNewParamsDataset = "workers_trace_events"
	JobNewParamsDatasetZarazEvents              JobNewParamsDataset = "zaraz_events"
	JobNewParamsDatasetZeroTrustNetworkSessions JobNewParamsDataset = "zero_trust_network_sessions"
)

func (r JobNewParamsDataset) IsKnown() bool {
	switch r {
	case JobNewParamsDatasetAccessRequests, JobNewParamsDatasetAuditLogs, JobNewParamsDatasetAuditLogsV2, JobNewParamsDatasetBISOUserActions, JobNewParamsDatasetCasbFindings, JobNewParamsDatasetDevicePostureResults, JobNewParamsDatasetDLPForensicCopies, JobNewParamsDatasetDNSFirewallLogs, JobNewParamsDatasetDNSLogs, JobNewParamsDatasetEmailSecurityAlerts, JobNewParamsDatasetFirewallEvents, JobNewParamsDatasetGatewayDNS, JobNewParamsDatasetGatewayHTTP, JobNewParamsDatasetGatewayNetwork, JobNewParamsDatasetHTTPRequests, JobNewParamsDatasetMagicIDsDetections, JobNewParamsDatasetNELReports, JobNewParamsDatasetNetworkAnalyticsLogs, JobNewParamsDatasetPageShieldEvents, JobNewParamsDatasetSinkholeHTTPLogs, JobNewParamsDatasetSpectrumEvents, JobNewParamsDatasetSSHLogs, JobNewParamsDatasetWorkersTraceEvents, JobNewParamsDatasetZarazEvents, JobNewParamsDatasetZeroTrustNetworkSessions:
		return true
	}
	return false
}

// This field is deprecated. Please use `max_upload_*` parameters instead. . The
// frequency at which Cloudflare sends batches of logs to your destination. Setting
// frequency to high sends your logs in larger quantities of smaller files. Setting
// frequency to low sends logs in smaller quantities of larger files.
type JobNewParamsFrequency string

const (
	JobNewParamsFrequencyHigh JobNewParamsFrequency = "high"
	JobNewParamsFrequencyLow  JobNewParamsFrequency = "low"
)

func (r JobNewParamsFrequency) IsKnown() bool {
	switch r {
	case JobNewParamsFrequencyHigh, JobNewParamsFrequencyLow:
		return true
	}
	return false
}

// The kind parameter (optional) is used to differentiate between Logpush and Edge
// Log Delivery jobs (when supported by the dataset).
type JobNewParamsKind string

const (
	JobNewParamsKindEmpty JobNewParamsKind = ""
	JobNewParamsKindEdge  JobNewParamsKind = "edge"
)

func (r JobNewParamsKind) IsKnown() bool {
	switch r {
	case JobNewParamsKindEmpty, JobNewParamsKindEdge:
		return true
	}
	return false
}

// The maximum uncompressed file size of a batch of logs. This setting value must
// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
// minimum file size; this means that log files may be much smaller than this batch
// size.
type JobNewParamsMaxUploadBytes int64

const (
	JobNewParamsMaxUploadBytes0 JobNewParamsMaxUploadBytes = 0
)

func (r JobNewParamsMaxUploadBytes) IsKnown() bool {
	switch r {
	case JobNewParamsMaxUploadBytes0:
		return true
	}
	return false
}

// The maximum interval in seconds for log batches. This setting must be between 30
// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
// a minimum interval for log batches; this means that log files may be sent in
// shorter intervals than this.
type JobNewParamsMaxUploadIntervalSeconds int64

const (
	JobNewParamsMaxUploadIntervalSeconds0 JobNewParamsMaxUploadIntervalSeconds = 0
)

func (r JobNewParamsMaxUploadIntervalSeconds) IsKnown() bool {
	switch r {
	case JobNewParamsMaxUploadIntervalSeconds0:
		return true
	}
	return false
}

// The maximum number of log lines per batch. This setting must be between 1000 and
// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
// number of log lines per batch; this means that log files may contain many fewer
// lines than this.
type JobNewParamsMaxUploadRecords int64

const (
	JobNewParamsMaxUploadRecords0 JobNewParamsMaxUploadRecords = 0
)

func (r JobNewParamsMaxUploadRecords) IsKnown() bool {
	switch r {
	case JobNewParamsMaxUploadRecords0:
		return true
	}
	return false
}

type JobNewResponseEnvelope struct {
	Errors   []JobNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success JobNewResponseEnvelopeSuccess `json:"success,required"`
	Result  LogpushJob                    `json:"result,nullable"`
	JSON    jobNewResponseEnvelopeJSON    `json:"-"`
}

// jobNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobNewResponseEnvelope]
type jobNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobNewResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           JobNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             jobNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// jobNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobNewResponseEnvelopeErrors]
type jobNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobNewResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    jobNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// jobNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [JobNewResponseEnvelopeErrorsSource]
type jobNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type JobNewResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           JobNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             jobNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// jobNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobNewResponseEnvelopeMessages]
type jobNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type JobNewResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    jobNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// jobNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [JobNewResponseEnvelopeMessagesSource]
type jobNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type JobNewResponseEnvelopeSuccess bool

const (
	JobNewResponseEnvelopeSuccessTrue JobNewResponseEnvelopeSuccess = true
)

func (r JobNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type JobUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Uniquely identifies a resource (such as an s3 bucket) where data. will be
	// pushed. Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf" format:"uri"`
	// Flag that indicates if the job is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The filters to select the events to include and/or remove from your logs. For
	// more information, refer to
	// [Filters](https://developers.cloudflare.com/logs/reference/filters/).
	Filter param.Field[string] `json:"filter"`
	// This field is deprecated. Please use `max_upload_*` parameters instead. . The
	// frequency at which Cloudflare sends batches of logs to your destination. Setting
	// frequency to high sends your logs in larger quantities of smaller files. Setting
	// frequency to low sends logs in smaller quantities of larger files.
	Frequency param.Field[JobUpdateParamsFrequency] `json:"frequency"`
	// The kind parameter (optional) is used to differentiate between Logpush and Edge
	// Log Delivery jobs (when supported by the dataset).
	Kind param.Field[JobUpdateParamsKind] `json:"kind"`
	// This field is deprecated. Use `output_options` instead. Configuration string. It
	// specifies things like requested fields and timestamp formats. If migrating from
	// the logpull api, copy the url (full url or just the query string) of your call
	// here, and logpush will keep on making this call for you, setting start and end
	// times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options" format:"uri-reference"`
	// The maximum uncompressed file size of a batch of logs. This setting value must
	// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
	// minimum file size; this means that log files may be much smaller than this batch
	// size.
	MaxUploadBytes param.Field[JobUpdateParamsMaxUploadBytes] `json:"max_upload_bytes"`
	// The maximum interval in seconds for log batches. This setting must be between 30
	// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
	// a minimum interval for log batches; this means that log files may be sent in
	// shorter intervals than this.
	MaxUploadIntervalSeconds param.Field[JobUpdateParamsMaxUploadIntervalSeconds] `json:"max_upload_interval_seconds"`
	// The maximum number of log lines per batch. This setting must be between 1000 and
	// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
	// number of log lines per batch; this means that log files may contain many fewer
	// lines than this.
	MaxUploadRecords param.Field[JobUpdateParamsMaxUploadRecords] `json:"max_upload_records"`
	// Optional human readable job name. Not unique. Cloudflare suggests. that you set
	// this to a meaningful string, like the domain name, to make it easier to identify
	// your job.
	Name param.Field[string] `json:"name"`
	// The structured replacement for `logpull_options`. When including this field, the
	// `logpull_option` field will be ignored.
	OutputOptions param.Field[OutputOptionsParam] `json:"output_options"`
	// Ownership challenge token to prove destination ownership.
	OwnershipChallenge param.Field[string] `json:"ownership_challenge"`
}

func (r JobUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// This field is deprecated. Please use `max_upload_*` parameters instead. . The
// frequency at which Cloudflare sends batches of logs to your destination. Setting
// frequency to high sends your logs in larger quantities of smaller files. Setting
// frequency to low sends logs in smaller quantities of larger files.
type JobUpdateParamsFrequency string

const (
	JobUpdateParamsFrequencyHigh JobUpdateParamsFrequency = "high"
	JobUpdateParamsFrequencyLow  JobUpdateParamsFrequency = "low"
)

func (r JobUpdateParamsFrequency) IsKnown() bool {
	switch r {
	case JobUpdateParamsFrequencyHigh, JobUpdateParamsFrequencyLow:
		return true
	}
	return false
}

// The kind parameter (optional) is used to differentiate between Logpush and Edge
// Log Delivery jobs (when supported by the dataset).
type JobUpdateParamsKind string

const (
	JobUpdateParamsKindEmpty JobUpdateParamsKind = ""
	JobUpdateParamsKindEdge  JobUpdateParamsKind = "edge"
)

func (r JobUpdateParamsKind) IsKnown() bool {
	switch r {
	case JobUpdateParamsKindEmpty, JobUpdateParamsKindEdge:
		return true
	}
	return false
}

// The maximum uncompressed file size of a batch of logs. This setting value must
// be between `5 MB` and `1 GB`, or `0` to disable it. Note that you cannot set a
// minimum file size; this means that log files may be much smaller than this batch
// size.
type JobUpdateParamsMaxUploadBytes int64

const (
	JobUpdateParamsMaxUploadBytes0 JobUpdateParamsMaxUploadBytes = 0
)

func (r JobUpdateParamsMaxUploadBytes) IsKnown() bool {
	switch r {
	case JobUpdateParamsMaxUploadBytes0:
		return true
	}
	return false
}

// The maximum interval in seconds for log batches. This setting must be between 30
// and 300 seconds (5 minutes), or `0` to disable it. Note that you cannot specify
// a minimum interval for log batches; this means that log files may be sent in
// shorter intervals than this.
type JobUpdateParamsMaxUploadIntervalSeconds int64

const (
	JobUpdateParamsMaxUploadIntervalSeconds0 JobUpdateParamsMaxUploadIntervalSeconds = 0
)

func (r JobUpdateParamsMaxUploadIntervalSeconds) IsKnown() bool {
	switch r {
	case JobUpdateParamsMaxUploadIntervalSeconds0:
		return true
	}
	return false
}

// The maximum number of log lines per batch. This setting must be between 1000 and
// 1,000,000 lines, or `0` to disable it. Note that you cannot specify a minimum
// number of log lines per batch; this means that log files may contain many fewer
// lines than this.
type JobUpdateParamsMaxUploadRecords int64

const (
	JobUpdateParamsMaxUploadRecords0 JobUpdateParamsMaxUploadRecords = 0
)

func (r JobUpdateParamsMaxUploadRecords) IsKnown() bool {
	switch r {
	case JobUpdateParamsMaxUploadRecords0:
		return true
	}
	return false
}

type JobUpdateResponseEnvelope struct {
	Errors   []JobUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success JobUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  LogpushJob                       `json:"result,nullable"`
	JSON    jobUpdateResponseEnvelopeJSON    `json:"-"`
}

// jobUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobUpdateResponseEnvelope]
type jobUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobUpdateResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           JobUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             jobUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// jobUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobUpdateResponseEnvelopeErrors]
type jobUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    jobUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// jobUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [JobUpdateResponseEnvelopeErrorsSource]
type jobUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type JobUpdateResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           JobUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             jobUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// jobUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobUpdateResponseEnvelopeMessages]
type jobUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type JobUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    jobUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// jobUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [JobUpdateResponseEnvelopeMessagesSource]
type jobUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type JobUpdateResponseEnvelopeSuccess bool

const (
	JobUpdateResponseEnvelopeSuccessTrue JobUpdateResponseEnvelopeSuccess = true
)

func (r JobUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type JobListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobDeleteResponseEnvelope struct {
	Errors   []JobDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success JobDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  JobDeleteResponse                `json:"result"`
	JSON    jobDeleteResponseEnvelopeJSON    `json:"-"`
}

// jobDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobDeleteResponseEnvelope]
type jobDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobDeleteResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           JobDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             jobDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// jobDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobDeleteResponseEnvelopeErrors]
type jobDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    jobDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// jobDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [JobDeleteResponseEnvelopeErrorsSource]
type jobDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type JobDeleteResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           JobDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             jobDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// jobDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobDeleteResponseEnvelopeMessages]
type jobDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type JobDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    jobDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// jobDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [JobDeleteResponseEnvelopeMessagesSource]
type jobDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type JobDeleteResponseEnvelopeSuccess bool

const (
	JobDeleteResponseEnvelopeSuccessTrue JobDeleteResponseEnvelopeSuccess = true
)

func (r JobDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type JobGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type JobGetResponseEnvelope struct {
	Errors   []JobGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []JobGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success JobGetResponseEnvelopeSuccess `json:"success,required"`
	Result  LogpushJob                    `json:"result,nullable"`
	JSON    jobGetResponseEnvelopeJSON    `json:"-"`
}

// jobGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [JobGetResponseEnvelope]
type jobGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           JobGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             jobGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// jobGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [JobGetResponseEnvelopeErrors]
type jobGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    jobGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// jobGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [JobGetResponseEnvelopeErrorsSource]
type jobGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           JobGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             jobGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// jobGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [JobGetResponseEnvelopeMessages]
type jobGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *JobGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type JobGetResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    jobGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// jobGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [JobGetResponseEnvelopeMessagesSource]
type jobGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *JobGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jobGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type JobGetResponseEnvelopeSuccess bool

const (
	JobGetResponseEnvelopeSuccessTrue JobGetResponseEnvelopeSuccess = true
)

func (r JobGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case JobGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
