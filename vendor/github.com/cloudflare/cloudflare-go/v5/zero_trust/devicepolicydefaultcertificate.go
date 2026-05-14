// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DevicePolicyDefaultCertificateService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicePolicyDefaultCertificateService] method instead.
type DevicePolicyDefaultCertificateService struct {
	Options []option.RequestOption
}

// NewDevicePolicyDefaultCertificateService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicePolicyDefaultCertificateService(opts ...option.RequestOption) (r *DevicePolicyDefaultCertificateService) {
	r = &DevicePolicyDefaultCertificateService{}
	r.Options = opts
	return
}

// Enable Zero Trust Clients to provision a certificate, containing a x509 subject,
// and referenced by Access device posture policies when the client visits MTLS
// protected domains. This facilitates device posture without a WARP session.
func (r *DevicePolicyDefaultCertificateService) Edit(ctx context.Context, params DevicePolicyDefaultCertificateEditParams, opts ...option.RequestOption) (res *DevicePolicyCertificates, err error) {
	var env DevicePolicyDefaultCertificateEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/devices/policy/certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches device certificate provisioning.
func (r *DevicePolicyDefaultCertificateService) Get(ctx context.Context, query DevicePolicyDefaultCertificateGetParams, opts ...option.RequestOption) (res *DevicePolicyCertificates, err error) {
	var env DevicePolicyDefaultCertificateGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/devices/policy/certificates", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DevicePolicyDefaultCertificateEditParams struct {
	ZoneID                   param.Field[string]           `path:"zone_id,required"`
	DevicePolicyCertificates DevicePolicyCertificatesParam `json:"device_policy_certificates,required"`
}

func (r DevicePolicyDefaultCertificateEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.DevicePolicyCertificates)
}

type DevicePolicyDefaultCertificateEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   DevicePolicyCertificates `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyDefaultCertificateEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyDefaultCertificateEditResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyDefaultCertificateEditResponseEnvelopeJSON contains the JSON
// metadata for the struct [DevicePolicyDefaultCertificateEditResponseEnvelope]
type devicePolicyDefaultCertificateEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultCertificateEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultCertificateEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultCertificateEditResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultCertificateEditResponseEnvelopeSuccessTrue DevicePolicyDefaultCertificateEditResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultCertificateEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultCertificateEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DevicePolicyDefaultCertificateGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type DevicePolicyDefaultCertificateGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   DevicePolicyCertificates `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DevicePolicyDefaultCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    devicePolicyDefaultCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// devicePolicyDefaultCertificateGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DevicePolicyDefaultCertificateGetResponseEnvelope]
type devicePolicyDefaultCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevicePolicyDefaultCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devicePolicyDefaultCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DevicePolicyDefaultCertificateGetResponseEnvelopeSuccess bool

const (
	DevicePolicyDefaultCertificateGetResponseEnvelopeSuccessTrue DevicePolicyDefaultCertificateGetResponseEnvelopeSuccess = true
)

func (r DevicePolicyDefaultCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DevicePolicyDefaultCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
