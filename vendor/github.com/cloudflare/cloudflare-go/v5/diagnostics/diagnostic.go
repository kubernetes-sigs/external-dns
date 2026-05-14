// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package diagnostics

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DiagnosticService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDiagnosticService] method instead.
type DiagnosticService struct {
	Options              []option.RequestOption
	Traceroutes          *TracerouteService
	EndpointHealthchecks *EndpointHealthcheckService
}

// NewDiagnosticService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDiagnosticService(opts ...option.RequestOption) (r *DiagnosticService) {
	r = &DiagnosticService{}
	r.Options = opts
	r.Traceroutes = NewTracerouteService(opts...)
	r.EndpointHealthchecks = NewEndpointHealthcheckService(opts...)
	return
}
