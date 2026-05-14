// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// MonitorPreviewService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorPreviewService] method instead.
type MonitorPreviewService struct {
	Options []option.RequestOption
}

// NewMonitorPreviewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMonitorPreviewService(opts ...option.RequestOption) (r *MonitorPreviewService) {
	r = &MonitorPreviewService{}
	r.Options = opts
	return
}

// Preview pools using the specified monitor with provided monitor details. The
// returned preview_id can be used in the preview endpoint to retrieve the results.
func (r *MonitorPreviewService) New(ctx context.Context, monitorID string, params MonitorPreviewNewParams, opts ...option.RequestOption) (res *MonitorPreviewNewResponse, err error) {
	var env MonitorPreviewNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if monitorID == "" {
		err = errors.New("missing required monitor_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitors/%s/preview", params.AccountID, monitorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type MonitorPreviewNewResponse struct {
	// Monitored pool IDs mapped to their respective names.
	Pools     map[string]string             `json:"pools"`
	PreviewID string                        `json:"preview_id"`
	JSON      monitorPreviewNewResponseJSON `json:"-"`
}

// monitorPreviewNewResponseJSON contains the JSON metadata for the struct
// [MonitorPreviewNewResponse]
type monitorPreviewNewResponseJSON struct {
	Pools       apijson.Field
	PreviewID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorPreviewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorPreviewNewResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorPreviewNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Do not validate the certificate when monitor use HTTPS. This parameter is
	// currently only valid for HTTP and HTTPS monitors.
	AllowInsecure param.Field[bool] `json:"allow_insecure"`
	// To be marked unhealthy the monitored origin must fail this healthcheck N
	// consecutive times.
	ConsecutiveDown param.Field[int64] `json:"consecutive_down"`
	// To be marked healthy the monitored origin must pass this healthcheck N
	// consecutive times.
	ConsecutiveUp param.Field[int64] `json:"consecutive_up"`
	// Object description.
	Description param.Field[string] `json:"description"`
	// A case-insensitive sub-string to look for in the response body. If this string
	// is not found, the origin will be marked as unhealthy. This parameter is only
	// valid for HTTP and HTTPS monitors.
	ExpectedBody param.Field[string] `json:"expected_body"`
	// The expected HTTP response code or code range of the health check. This
	// parameter is only valid for HTTP and HTTPS monitors.
	ExpectedCodes param.Field[string] `json:"expected_codes"`
	// Follow redirects if returned by the origin. This parameter is only valid for
	// HTTP and HTTPS monitors.
	FollowRedirects param.Field[bool] `json:"follow_redirects"`
	// The HTTP request headers to send in the health check. It is recommended you set
	// a Host header by default. The User-Agent header cannot be overridden. This
	// parameter is only valid for HTTP and HTTPS monitors.
	Header param.Field[map[string][]string] `json:"header"`
	// The interval between each health check. Shorter intervals may improve failover
	// time, but will increase load on the origins as we check from multiple locations.
	Interval param.Field[int64] `json:"interval"`
	// The method to use for the health check. This defaults to 'GET' for HTTP/HTTPS
	// based checks and 'connection_established' for TCP based health checks.
	Method param.Field[string] `json:"method"`
	// The endpoint path you want to conduct a health check against. This parameter is
	// only valid for HTTP and HTTPS monitors.
	Path param.Field[string] `json:"path"`
	// The port number to connect to for the health check. Required for TCP, UDP, and
	// SMTP checks. HTTP and HTTPS checks should only define the port when using a
	// non-standard port (HTTP: default 80, HTTPS: default 443).
	Port param.Field[int64] `json:"port"`
	// Assign this monitor to emulate the specified zone while probing. This parameter
	// is only valid for HTTP and HTTPS monitors.
	ProbeZone param.Field[string] `json:"probe_zone"`
	// The number of retries to attempt in case of a timeout before marking the origin
	// as unhealthy. Retries are attempted immediately.
	Retries param.Field[int64] `json:"retries"`
	// The timeout (in seconds) before marking the health check as failed.
	Timeout param.Field[int64] `json:"timeout"`
	// The protocol to use for the health check. Currently supported protocols are
	// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
	Type param.Field[MonitorPreviewNewParamsType] `json:"type"`
}

func (r MonitorPreviewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol to use for the health check. Currently supported protocols are
// 'HTTP','HTTPS', 'TCP', 'ICMP-PING', 'UDP-ICMP', and 'SMTP'.
type MonitorPreviewNewParamsType string

const (
	MonitorPreviewNewParamsTypeHTTP     MonitorPreviewNewParamsType = "http"
	MonitorPreviewNewParamsTypeHTTPS    MonitorPreviewNewParamsType = "https"
	MonitorPreviewNewParamsTypeTCP      MonitorPreviewNewParamsType = "tcp"
	MonitorPreviewNewParamsTypeUdpIcmp  MonitorPreviewNewParamsType = "udp_icmp"
	MonitorPreviewNewParamsTypeIcmpPing MonitorPreviewNewParamsType = "icmp_ping"
	MonitorPreviewNewParamsTypeSmtp     MonitorPreviewNewParamsType = "smtp"
)

func (r MonitorPreviewNewParamsType) IsKnown() bool {
	switch r {
	case MonitorPreviewNewParamsTypeHTTP, MonitorPreviewNewParamsTypeHTTPS, MonitorPreviewNewParamsTypeTCP, MonitorPreviewNewParamsTypeUdpIcmp, MonitorPreviewNewParamsTypeIcmpPing, MonitorPreviewNewParamsTypeSmtp:
		return true
	}
	return false
}

type MonitorPreviewNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   MonitorPreviewNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success MonitorPreviewNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    monitorPreviewNewResponseEnvelopeJSON    `json:"-"`
}

// monitorPreviewNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MonitorPreviewNewResponseEnvelope]
type monitorPreviewNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MonitorPreviewNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorPreviewNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MonitorPreviewNewResponseEnvelopeSuccess bool

const (
	MonitorPreviewNewResponseEnvelopeSuccessTrue MonitorPreviewNewResponseEnvelopeSuccess = true
)

func (r MonitorPreviewNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MonitorPreviewNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
