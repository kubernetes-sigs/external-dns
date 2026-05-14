// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DEXCommandService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXCommandService] method instead.
type DEXCommandService struct {
	Options   []option.RequestOption
	Devices   *DEXCommandDeviceService
	Downloads *DEXCommandDownloadService
	Quota     *DEXCommandQuotaService
}

// NewDEXCommandService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXCommandService(opts ...option.RequestOption) (r *DEXCommandService) {
	r = &DEXCommandService{}
	r.Options = opts
	r.Devices = NewDEXCommandDeviceService(opts...)
	r.Downloads = NewDEXCommandDownloadService(opts...)
	r.Quota = NewDEXCommandQuotaService(opts...)
	return
}

// Initiate commands for up to 10 devices per account
func (r *DEXCommandService) New(ctx context.Context, params DEXCommandNewParams, opts ...option.RequestOption) (res *DEXCommandNewResponse, err error) {
	var env DEXCommandNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a paginated list of commands issued to devices under the specified
// account, optionally filtered by time range, device, or other parameters
func (r *DEXCommandService) List(ctx context.Context, params DEXCommandListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[DEXCommandListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/commands", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Retrieves a paginated list of commands issued to devices under the specified
// account, optionally filtered by time range, device, or other parameters
func (r *DEXCommandService) ListAutoPaging(ctx context.Context, params DEXCommandListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[DEXCommandListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type DEXCommandNewResponse struct {
	// List of created commands
	Commands []DEXCommandNewResponseCommand `json:"commands"`
	JSON     dexCommandNewResponseJSON      `json:"-"`
}

// dexCommandNewResponseJSON contains the JSON metadata for the struct
// [DEXCommandNewResponse]
type dexCommandNewResponseJSON struct {
	Commands    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseCommand struct {
	// Unique identifier for the command
	ID string `json:"id"`
	// Command arguments
	Args map[string]string `json:"args"`
	// Identifier for the device associated with the command
	DeviceID string `json:"device_id"`
	// Current status of the command
	Status DEXCommandNewResponseCommandsStatus `json:"status"`
	// Type of the command (e.g., "pcap" or "warp-diag")
	Type string                           `json:"type"`
	JSON dexCommandNewResponseCommandJSON `json:"-"`
}

// dexCommandNewResponseCommandJSON contains the JSON metadata for the struct
// [DEXCommandNewResponseCommand]
type dexCommandNewResponseCommandJSON struct {
	ID          apijson.Field
	Args        apijson.Field
	DeviceID    apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseCommandJSON) RawJSON() string {
	return r.raw
}

// Current status of the command
type DEXCommandNewResponseCommandsStatus string

const (
	DEXCommandNewResponseCommandsStatusPendingExec   DEXCommandNewResponseCommandsStatus = "PENDING_EXEC"
	DEXCommandNewResponseCommandsStatusPendingUpload DEXCommandNewResponseCommandsStatus = "PENDING_UPLOAD"
	DEXCommandNewResponseCommandsStatusSuccess       DEXCommandNewResponseCommandsStatus = "SUCCESS"
	DEXCommandNewResponseCommandsStatusFailed        DEXCommandNewResponseCommandsStatus = "FAILED"
)

func (r DEXCommandNewResponseCommandsStatus) IsKnown() bool {
	switch r {
	case DEXCommandNewResponseCommandsStatusPendingExec, DEXCommandNewResponseCommandsStatusPendingUpload, DEXCommandNewResponseCommandsStatusSuccess, DEXCommandNewResponseCommandsStatusFailed:
		return true
	}
	return false
}

type DEXCommandListResponse struct {
	Commands []DEXCommandListResponseCommand `json:"commands"`
	JSON     dexCommandListResponseJSON      `json:"-"`
}

// dexCommandListResponseJSON contains the JSON metadata for the struct
// [DEXCommandListResponse]
type dexCommandListResponseJSON struct {
	Commands    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXCommandListResponseCommand struct {
	ID            string                            `json:"id"`
	CompletedDate time.Time                         `json:"completed_date,nullable" format:"date-time"`
	CreatedDate   time.Time                         `json:"created_date" format:"date-time"`
	DeviceID      string                            `json:"device_id"`
	Filename      string                            `json:"filename,nullable"`
	Status        string                            `json:"status"`
	Type          string                            `json:"type"`
	UserEmail     string                            `json:"user_email"`
	JSON          dexCommandListResponseCommandJSON `json:"-"`
}

// dexCommandListResponseCommandJSON contains the JSON metadata for the struct
// [DEXCommandListResponseCommand]
type dexCommandListResponseCommandJSON struct {
	ID            apijson.Field
	CompletedDate apijson.Field
	CreatedDate   apijson.Field
	DeviceID      apijson.Field
	Filename      apijson.Field
	Status        apijson.Field
	Type          apijson.Field
	UserEmail     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXCommandListResponseCommand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandListResponseCommandJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// List of device-level commands to execute
	Commands param.Field[[]DEXCommandNewParamsCommand] `json:"commands,required"`
}

func (r DEXCommandNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DEXCommandNewParamsCommand struct {
	// Type of command to execute on the device
	CommandType param.Field[DEXCommandNewParamsCommandsCommandType] `json:"command_type,required"`
	// Unique identifier for the device
	DeviceID param.Field[string] `json:"device_id,required"`
	// Email tied to the device
	UserEmail   param.Field[string]                                 `json:"user_email,required"`
	CommandArgs param.Field[DEXCommandNewParamsCommandsCommandArgs] `json:"command_args"`
}

func (r DEXCommandNewParamsCommand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of command to execute on the device
type DEXCommandNewParamsCommandsCommandType string

const (
	DEXCommandNewParamsCommandsCommandTypePCAP     DEXCommandNewParamsCommandsCommandType = "pcap"
	DEXCommandNewParamsCommandsCommandTypeWARPDiag DEXCommandNewParamsCommandsCommandType = "warp-diag"
)

func (r DEXCommandNewParamsCommandsCommandType) IsKnown() bool {
	switch r {
	case DEXCommandNewParamsCommandsCommandTypePCAP, DEXCommandNewParamsCommandsCommandTypeWARPDiag:
		return true
	}
	return false
}

type DEXCommandNewParamsCommandsCommandArgs struct {
	// List of interfaces to capture packets on
	Interfaces param.Field[[]DEXCommandNewParamsCommandsCommandArgsInterface] `json:"interfaces"`
	// Maximum file size (in MB) for the capture file. Specifies the maximum file size
	// of the warp-diag zip artifact that can be uploaded. If the zip artifact exceeds
	// the specified max file size, it will NOT be uploaded
	MaxFileSizeMB param.Field[float64] `json:"max-file-size-mb"`
	// Maximum number of bytes to save for each packet
	PacketSizeBytes param.Field[float64] `json:"packet-size-bytes"`
	// Test an IP address from all included or excluded ranges. Tests an IP address
	// from all included or excluded ranges. Essentially the same as running 'route get
	// <ip>” and collecting the results. This option may increase the time taken to
	// collect the warp-diag
	TestAllRoutes param.Field[bool] `json:"test-all-routes"`
	// Limit on capture duration (in minutes)
	TimeLimitMin param.Field[float64] `json:"time-limit-min"`
}

func (r DEXCommandNewParamsCommandsCommandArgs) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DEXCommandNewParamsCommandsCommandArgsInterface string

const (
	DEXCommandNewParamsCommandsCommandArgsInterfaceDefault DEXCommandNewParamsCommandsCommandArgsInterface = "default"
	DEXCommandNewParamsCommandsCommandArgsInterfaceTunnel  DEXCommandNewParamsCommandsCommandArgsInterface = "tunnel"
)

func (r DEXCommandNewParamsCommandsCommandArgsInterface) IsKnown() bool {
	switch r {
	case DEXCommandNewParamsCommandsCommandArgsInterfaceDefault, DEXCommandNewParamsCommandsCommandArgsInterfaceTunnel:
		return true
	}
	return false
}

type DEXCommandNewResponseEnvelope struct {
	Errors   []DEXCommandNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXCommandNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    DEXCommandNewResponseEnvelopeSuccess    `json:"success,required"`
	Result     DEXCommandNewResponse                   `json:"result"`
	ResultInfo DEXCommandNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       dexCommandNewResponseEnvelopeJSON       `json:"-"`
}

// dexCommandNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXCommandNewResponseEnvelope]
type dexCommandNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DEXCommandNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexCommandNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexCommandNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeErrors]
type dexCommandNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dexCommandNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexCommandNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeErrorsSource]
type dexCommandNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DEXCommandNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexCommandNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexCommandNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeMessages]
type dexCommandNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXCommandNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dexCommandNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexCommandNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DEXCommandNewResponseEnvelopeMessagesSource]
type dexCommandNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXCommandNewResponseEnvelopeSuccess bool

const (
	DEXCommandNewResponseEnvelopeSuccessTrue DEXCommandNewResponseEnvelopeSuccess = true
)

func (r DEXCommandNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXCommandNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXCommandNewResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                     `json:"total_count"`
	JSON       dexCommandNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// dexCommandNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [DEXCommandNewResponseEnvelopeResultInfo]
type dexCommandNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXCommandNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexCommandNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DEXCommandListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number for pagination
	Page param.Field[float64] `query:"page,required"`
	// Number of results per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Optionally filter executed commands by command type
	CommandType param.Field[string] `query:"command_type"`
	// Unique identifier for a device
	DeviceID param.Field[string] `query:"device_id"`
	// Start time for the query in ISO (RFC3339 - ISO 8601) format
	From param.Field[time.Time] `query:"from" format:"date-time"`
	// Optionally filter executed commands by status
	Status param.Field[DEXCommandListParamsStatus] `query:"status"`
	// End time for the query in ISO (RFC3339 - ISO 8601) format
	To param.Field[time.Time] `query:"to" format:"date-time"`
	// Email tied to the device
	UserEmail param.Field[string] `query:"user_email"`
}

// URLQuery serializes [DEXCommandListParams]'s query parameters as `url.Values`.
func (r DEXCommandListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Optionally filter executed commands by status
type DEXCommandListParamsStatus string

const (
	DEXCommandListParamsStatusPendingExec   DEXCommandListParamsStatus = "PENDING_EXEC"
	DEXCommandListParamsStatusPendingUpload DEXCommandListParamsStatus = "PENDING_UPLOAD"
	DEXCommandListParamsStatusSuccess       DEXCommandListParamsStatus = "SUCCESS"
	DEXCommandListParamsStatusFailed        DEXCommandListParamsStatus = "FAILED"
)

func (r DEXCommandListParamsStatus) IsKnown() bool {
	switch r {
	case DEXCommandListParamsStatusPendingExec, DEXCommandListParamsStatusPendingUpload, DEXCommandListParamsStatusSuccess, DEXCommandListParamsStatusFailed:
		return true
	}
	return false
}
