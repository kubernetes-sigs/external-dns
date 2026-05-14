// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package content_scanning

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

// ContentScanningService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContentScanningService] method instead.
type ContentScanningService struct {
	Options  []option.RequestOption
	Payloads *PayloadService
	Settings *SettingService
}

// NewContentScanningService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContentScanningService(opts ...option.RequestOption) (r *ContentScanningService) {
	r = &ContentScanningService{}
	r.Options = opts
	r.Payloads = NewPayloadService(opts...)
	r.Settings = NewSettingService(opts...)
	return
}

// Disable Content Scanning.
func (r *ContentScanningService) Disable(ctx context.Context, body ContentScanningDisableParams, opts ...option.RequestOption) (res *ContentScanningDisableResponse, err error) {
	var env ContentScanningDisableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/disable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Enable Content Scanning.
func (r *ContentScanningService) Enable(ctx context.Context, body ContentScanningEnableParams, opts ...option.RequestOption) (res *ContentScanningEnableResponse, err error) {
	var env ContentScanningEnableResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/content-upload-scan/enable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ContentScanningDisableResponse = interface{}

type ContentScanningEnableResponse = interface{}

type ContentScanningDisableParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningDisableResponseEnvelope struct {
	Errors   []shared.ResponseInfo          `json:"errors,required"`
	Messages []shared.ResponseInfo          `json:"messages,required"`
	Result   ContentScanningDisableResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ContentScanningDisableResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningDisableResponseEnvelopeJSON    `json:"-"`
}

// contentScanningDisableResponseEnvelopeJSON contains the JSON metadata for the
// struct [ContentScanningDisableResponseEnvelope]
type contentScanningDisableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningDisableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningDisableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ContentScanningDisableResponseEnvelopeSuccess bool

const (
	ContentScanningDisableResponseEnvelopeSuccessTrue ContentScanningDisableResponseEnvelopeSuccess = true
)

func (r ContentScanningDisableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningDisableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ContentScanningEnableParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ContentScanningEnableResponseEnvelope struct {
	Errors   []shared.ResponseInfo         `json:"errors,required"`
	Messages []shared.ResponseInfo         `json:"messages,required"`
	Result   ContentScanningEnableResponse `json:"result,required"`
	// Whether the API call was successful.
	Success ContentScanningEnableResponseEnvelopeSuccess `json:"success,required"`
	JSON    contentScanningEnableResponseEnvelopeJSON    `json:"-"`
}

// contentScanningEnableResponseEnvelopeJSON contains the JSON metadata for the
// struct [ContentScanningEnableResponseEnvelope]
type contentScanningEnableResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ContentScanningEnableResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r contentScanningEnableResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ContentScanningEnableResponseEnvelopeSuccess bool

const (
	ContentScanningEnableResponseEnvelopeSuccessTrue ContentScanningEnableResponseEnvelopeSuccess = true
)

func (r ContentScanningEnableResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ContentScanningEnableResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
