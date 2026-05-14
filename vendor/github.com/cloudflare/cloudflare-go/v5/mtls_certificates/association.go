// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package mtls_certificates

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

// AssociationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssociationService] method instead.
type AssociationService struct {
	Options []option.RequestOption
}

// NewAssociationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssociationService(opts ...option.RequestOption) (r *AssociationService) {
	r = &AssociationService{}
	r.Options = opts
	return
}

// Lists all active associations between the certificate and Cloudflare services.
func (r *AssociationService) Get(ctx context.Context, mtlsCertificateID string, query AssociationGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[CertificateAsssociation], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if mtlsCertificateID == "" {
		err = errors.New("missing required mtls_certificate_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mtls_certificates/%s/associations", query.AccountID, mtlsCertificateID)
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

// Lists all active associations between the certificate and Cloudflare services.
func (r *AssociationService) GetAutoPaging(ctx context.Context, mtlsCertificateID string, query AssociationGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CertificateAsssociation] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, mtlsCertificateID, query, opts...))
}

type CertificateAsssociation struct {
	// The service using the certificate.
	Service string `json:"service"`
	// Certificate deployment status for the given service.
	Status string                      `json:"status"`
	JSON   certificateAsssociationJSON `json:"-"`
}

// certificateAsssociationJSON contains the JSON metadata for the struct
// [CertificateAsssociation]
type certificateAsssociationJSON struct {
	Service     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateAsssociation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r certificateAsssociationJSON) RawJSON() string {
	return r.raw
}

type AssociationGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
