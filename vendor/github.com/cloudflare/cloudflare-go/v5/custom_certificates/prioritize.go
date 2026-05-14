// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_certificates

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

// PrioritizeService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrioritizeService] method instead.
type PrioritizeService struct {
	Options []option.RequestOption
}

// NewPrioritizeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrioritizeService(opts ...option.RequestOption) (r *PrioritizeService) {
	r = &PrioritizeService{}
	r.Options = opts
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *PrioritizeService) Update(ctx context.Context, params PrioritizeUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[CustomCertificate], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *PrioritizeService) UpdateAutoPaging(ctx context.Context, params PrioritizeUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CustomCertificate] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, params, opts...))
}

type PrioritizeUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Array of ordered certificates.
	Certificates param.Field[[]PrioritizeUpdateParamsCertificate] `json:"certificates,required"`
}

func (r PrioritizeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrioritizeUpdateParamsCertificate struct {
	// Identifier.
	ID param.Field[string] `json:"id"`
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r PrioritizeUpdateParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
