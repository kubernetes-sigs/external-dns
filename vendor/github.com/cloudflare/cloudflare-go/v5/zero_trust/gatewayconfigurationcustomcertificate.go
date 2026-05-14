// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// GatewayConfigurationCustomCertificateService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayConfigurationCustomCertificateService] method instead.
type GatewayConfigurationCustomCertificateService struct {
	Options []option.RequestOption
}

// NewGatewayConfigurationCustomCertificateService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewGatewayConfigurationCustomCertificateService(opts ...option.RequestOption) (r *GatewayConfigurationCustomCertificateService) {
	r = &GatewayConfigurationCustomCertificateService{}
	r.Options = opts
	return
}

// Fetches the current Zero Trust certificate configuration.
//
// Deprecated: deprecated
func (r *GatewayConfigurationCustomCertificateService) Get(ctx context.Context, query GatewayConfigurationCustomCertificateGetParams, opts ...option.RequestOption) (res *CustomCertificateSettings, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/configuration/custom_certificate", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GatewayConfigurationCustomCertificateGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
