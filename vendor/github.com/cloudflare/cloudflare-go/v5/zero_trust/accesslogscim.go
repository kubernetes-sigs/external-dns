// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AccessLogSCIMService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessLogSCIMService] method instead.
type AccessLogSCIMService struct {
	Options []option.RequestOption
	Updates *AccessLogSCIMUpdateService
}

// NewAccessLogSCIMService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessLogSCIMService(opts ...option.RequestOption) (r *AccessLogSCIMService) {
	r = &AccessLogSCIMService{}
	r.Options = opts
	r.Updates = NewAccessLogSCIMUpdateService(opts...)
	return
}

type AccessRequest struct {
	// The event that occurred, such as a login attempt.
	Action string `json:"action"`
	// The result of the authentication event.
	Allowed bool `json:"allowed"`
	// The URL of the Access application.
	AppDomain string `json:"app_domain"`
	// The unique identifier for the Access application.
	AppUID string `json:"app_uid"`
	// The IdP used to authenticate.
	Connection string    `json:"connection"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// The IP address of the authenticating user.
	IPAddress string `json:"ip_address"`
	// The unique identifier for the request to Cloudflare.
	RayID string `json:"ray_id"`
	// The email address of the authenticating user.
	UserEmail string            `json:"user_email" format:"email"`
	JSON      accessRequestJSON `json:"-"`
}

// accessRequestJSON contains the JSON metadata for the struct [AccessRequest]
type accessRequestJSON struct {
	Action      apijson.Field
	Allowed     apijson.Field
	AppDomain   apijson.Field
	AppUID      apijson.Field
	Connection  apijson.Field
	CreatedAt   apijson.Field
	IPAddress   apijson.Field
	RayID       apijson.Field
	UserEmail   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRequestJSON) RawJSON() string {
	return r.raw
}
