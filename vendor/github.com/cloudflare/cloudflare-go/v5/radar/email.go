// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// EmailService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailService] method instead.
type EmailService struct {
	Options  []option.RequestOption
	Routing  *EmailRoutingService
	Security *EmailSecurityService
}

// NewEmailService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmailService(opts ...option.RequestOption) (r *EmailService) {
	r = &EmailService{}
	r.Options = opts
	r.Routing = NewEmailRoutingService(opts...)
	r.Security = NewEmailSecurityService(opts...)
	return
}

type RadarEmailSeries struct {
	Fail []string             `json:"FAIL,required"`
	None []string             `json:"NONE,required"`
	Pass []string             `json:"PASS,required"`
	JSON radarEmailSeriesJSON `json:"-"`
}

// radarEmailSeriesJSON contains the JSON metadata for the struct
// [RadarEmailSeries]
type radarEmailSeriesJSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSeriesJSON) RawJSON() string {
	return r.raw
}

type RadarEmailSummary struct {
	// A numeric string.
	Fail string `json:"FAIL,required"`
	// A numeric string.
	None string `json:"NONE,required"`
	// A numeric string.
	Pass string                `json:"PASS,required"`
	JSON radarEmailSummaryJSON `json:"-"`
}

// radarEmailSummaryJSON contains the JSON metadata for the struct
// [RadarEmailSummary]
type radarEmailSummaryJSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSummary) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarEmailSummaryJSON) RawJSON() string {
	return r.raw
}
