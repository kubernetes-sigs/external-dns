// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_post_quantum_encryption

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// OriginPostQuantumEncryptionService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginPostQuantumEncryptionService] method instead.
type OriginPostQuantumEncryptionService struct {
	Options []option.RequestOption
}

// NewOriginPostQuantumEncryptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOriginPostQuantumEncryptionService(opts ...option.RequestOption) (r *OriginPostQuantumEncryptionService) {
	r = &OriginPostQuantumEncryptionService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Update(ctx context.Context, params OriginPostQuantumEncryptionUpdateParams, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionUpdateResponse, err error) {
	var env OriginPostQuantumEncryptionUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to use Post-Quantum (PQ) key agreement algorithms when
// connecting to your origin. Preferred instructs Cloudflare to opportunistically
// send a Post-Quantum keyshare in the first message to the origin (for fastest
// connections when the origin supports and prefers PQ), supported means that PQ
// algorithms are advertised but only used when requested by the origin, and off
// means that PQ algorithms are not advertised
func (r *OriginPostQuantumEncryptionService) Get(ctx context.Context, query OriginPostQuantumEncryptionGetParams, opts ...option.RequestOption) (res *OriginPostQuantumEncryptionGetResponse, err error) {
	var env OriginPostQuantumEncryptionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/origin_post_quantum_encryption", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OriginPostQuantumEncryptionUpdateResponse struct {
	// Value of the zone setting.
	ID OriginPostQuantumEncryptionUpdateResponseID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value OriginPostQuantumEncryptionUpdateResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                     `json:"modified_on,nullable" format:"date-time"`
	JSON       originPostQuantumEncryptionUpdateResponseJSON `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseJSON contains the JSON metadata for the
// struct [OriginPostQuantumEncryptionUpdateResponse]
type originPostQuantumEncryptionUpdateResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type OriginPostQuantumEncryptionUpdateResponseID string

const (
	OriginPostQuantumEncryptionUpdateResponseIDOriginPqe OriginPostQuantumEncryptionUpdateResponseID = "origin_pqe"
)

func (r OriginPostQuantumEncryptionUpdateResponseID) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateResponseIDOriginPqe:
		return true
	}
	return false
}

// The value of the feature
type OriginPostQuantumEncryptionUpdateResponseValue string

const (
	OriginPostQuantumEncryptionUpdateResponseValuePreferred OriginPostQuantumEncryptionUpdateResponseValue = "preferred"
	OriginPostQuantumEncryptionUpdateResponseValueSupported OriginPostQuantumEncryptionUpdateResponseValue = "supported"
	OriginPostQuantumEncryptionUpdateResponseValueOff       OriginPostQuantumEncryptionUpdateResponseValue = "off"
)

func (r OriginPostQuantumEncryptionUpdateResponseValue) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateResponseValuePreferred, OriginPostQuantumEncryptionUpdateResponseValueSupported, OriginPostQuantumEncryptionUpdateResponseValueOff:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionGetResponse struct {
	// Value of the zone setting.
	ID OriginPostQuantumEncryptionGetResponseID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value OriginPostQuantumEncryptionGetResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                                  `json:"modified_on,nullable" format:"date-time"`
	JSON       originPostQuantumEncryptionGetResponseJSON `json:"-"`
}

// originPostQuantumEncryptionGetResponseJSON contains the JSON metadata for the
// struct [OriginPostQuantumEncryptionGetResponse]
type originPostQuantumEncryptionGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionGetResponseJSON) RawJSON() string {
	return r.raw
}

// Value of the zone setting.
type OriginPostQuantumEncryptionGetResponseID string

const (
	OriginPostQuantumEncryptionGetResponseIDOriginPqe OriginPostQuantumEncryptionGetResponseID = "origin_pqe"
)

func (r OriginPostQuantumEncryptionGetResponseID) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionGetResponseIDOriginPqe:
		return true
	}
	return false
}

// The value of the feature
type OriginPostQuantumEncryptionGetResponseValue string

const (
	OriginPostQuantumEncryptionGetResponseValuePreferred OriginPostQuantumEncryptionGetResponseValue = "preferred"
	OriginPostQuantumEncryptionGetResponseValueSupported OriginPostQuantumEncryptionGetResponseValue = "supported"
	OriginPostQuantumEncryptionGetResponseValueOff       OriginPostQuantumEncryptionGetResponseValue = "off"
)

func (r OriginPostQuantumEncryptionGetResponseValue) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionGetResponseValuePreferred, OriginPostQuantumEncryptionGetResponseValueSupported, OriginPostQuantumEncryptionGetResponseValueOff:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Origin Post Quantum Encryption Setting.
	Value param.Field[OriginPostQuantumEncryptionUpdateParamsValue] `json:"value,required"`
}

func (r OriginPostQuantumEncryptionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Origin Post Quantum Encryption Setting.
type OriginPostQuantumEncryptionUpdateParamsValue string

const (
	OriginPostQuantumEncryptionUpdateParamsValuePreferred OriginPostQuantumEncryptionUpdateParamsValue = "preferred"
	OriginPostQuantumEncryptionUpdateParamsValueSupported OriginPostQuantumEncryptionUpdateParamsValue = "supported"
	OriginPostQuantumEncryptionUpdateParamsValueOff       OriginPostQuantumEncryptionUpdateParamsValue = "off"
)

func (r OriginPostQuantumEncryptionUpdateParamsValue) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateParamsValuePreferred, OriginPostQuantumEncryptionUpdateParamsValueSupported, OriginPostQuantumEncryptionUpdateParamsValueOff:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginPostQuantumEncryptionUpdateResponse                `json:"result"`
	JSON    originPostQuantumEncryptionUpdateResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionUpdateResponseEnvelope]
type originPostQuantumEncryptionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess = true
)

func (r OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginPostQuantumEncryptionGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OriginPostQuantumEncryptionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success OriginPostQuantumEncryptionGetResponseEnvelopeSuccess `json:"success,required"`
	Result  OriginPostQuantumEncryptionGetResponse                `json:"result"`
	JSON    originPostQuantumEncryptionGetResponseEnvelopeJSON    `json:"-"`
}

// originPostQuantumEncryptionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginPostQuantumEncryptionGetResponseEnvelope]
type originPostQuantumEncryptionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginPostQuantumEncryptionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originPostQuantumEncryptionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type OriginPostQuantumEncryptionGetResponseEnvelopeSuccess bool

const (
	OriginPostQuantumEncryptionGetResponseEnvelopeSuccessTrue OriginPostQuantumEncryptionGetResponseEnvelopeSuccess = true
)

func (r OriginPostQuantumEncryptionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginPostQuantumEncryptionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
