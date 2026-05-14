// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package registrar

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	return
}

// Update individual domain.
func (r *DomainService) Update(ctx context.Context, domainName string, params DomainUpdateParams, opts ...option.RequestOption) (res *DomainUpdateResponse, err error) {
	var env DomainUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", params.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List domains handled by Registrar.
func (r *DomainService) List(ctx context.Context, query DomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Domain], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains", query.AccountID)
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

// List domains handled by Registrar.
func (r *DomainService) ListAutoPaging(ctx context.Context, query DomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Domain] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Show individual domain.
func (r *DomainService) Get(ctx context.Context, domainName string, query DomainGetParams, opts ...option.RequestOption) (res *DomainGetResponse, err error) {
	var env DomainGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/registrar/domains/%s", query.AccountID, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Domain struct {
	// Domain identifier.
	ID string `json:"id"`
	// Shows if a domain is available for transferring into Cloudflare Registrar.
	Available bool `json:"available"`
	// Indicates if the domain can be registered as a new domain.
	CanRegister bool `json:"can_register"`
	// Shows time of creation.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Shows name of current registrar.
	CurrentRegistrar string `json:"current_registrar"`
	// Shows when domain name registration expires.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Shows whether a registrar lock is in place for a domain.
	Locked bool `json:"locked"`
	// Shows contact information for domain registrant.
	RegistrantContact DomainRegistrantContact `json:"registrant_contact"`
	// A comma-separated list of registry status codes. A full list of status codes can
	// be found at
	// [EPP Status Codes](https://www.icann.org/resources/pages/epp-status-codes-2014-06-16-en).
	RegistryStatuses string `json:"registry_statuses"`
	// Whether a particular TLD is currently supported by Cloudflare Registrar. Refer
	// to [TLD Policies](https://www.cloudflare.com/tld-policies/) for a list of
	// supported TLDs.
	SupportedTld bool `json:"supported_tld"`
	// Statuses for domain transfers into Cloudflare Registrar.
	TransferIn DomainTransferIn `json:"transfer_in"`
	// Last updated.
	UpdatedAt time.Time  `json:"updated_at" format:"date-time"`
	JSON      domainJSON `json:"-"`
}

// domainJSON contains the JSON metadata for the struct [Domain]
type domainJSON struct {
	ID                apijson.Field
	Available         apijson.Field
	CanRegister       apijson.Field
	CreatedAt         apijson.Field
	CurrentRegistrar  apijson.Field
	ExpiresAt         apijson.Field
	Locked            apijson.Field
	RegistrantContact apijson.Field
	RegistryStatuses  apijson.Field
	SupportedTld      apijson.Field
	TransferIn        apijson.Field
	UpdatedAt         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *Domain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainJSON) RawJSON() string {
	return r.raw
}

// Shows contact information for domain registrant.
type DomainRegistrantContact struct {
	// Address.
	Address string `json:"address,required"`
	// City.
	City string `json:"city,required"`
	// The country in which the user lives.
	Country string `json:"country,required,nullable"`
	// User's first name
	FirstName string `json:"first_name,required,nullable"`
	// User's last name
	LastName string `json:"last_name,required,nullable"`
	// Name of organization.
	Organization string `json:"organization,required"`
	// User's telephone number
	Phone string `json:"phone,required,nullable"`
	// State.
	State string `json:"state,required"`
	// The zipcode or postal code where the user lives.
	Zip string `json:"zip,required,nullable"`
	// Contact Identifier.
	ID string `json:"id"`
	// Optional address line for unit, floor, suite, etc.
	Address2 string `json:"address2"`
	// The contact email address of the user.
	Email string `json:"email"`
	// Contact fax number.
	Fax  string                      `json:"fax"`
	JSON domainRegistrantContactJSON `json:"-"`
}

// domainRegistrantContactJSON contains the JSON metadata for the struct
// [DomainRegistrantContact]
type domainRegistrantContactJSON struct {
	Address      apijson.Field
	City         apijson.Field
	Country      apijson.Field
	FirstName    apijson.Field
	LastName     apijson.Field
	Organization apijson.Field
	Phone        apijson.Field
	State        apijson.Field
	Zip          apijson.Field
	ID           apijson.Field
	Address2     apijson.Field
	Email        apijson.Field
	Fax          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DomainRegistrantContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainRegistrantContactJSON) RawJSON() string {
	return r.raw
}

// Statuses for domain transfers into Cloudflare Registrar.
type DomainTransferIn struct {
	// Form of authorization has been accepted by the registrant.
	AcceptFoa DomainTransferInAcceptFoa `json:"accept_foa"`
	// Shows transfer status with the registry.
	ApproveTransfer DomainTransferInApproveTransfer `json:"approve_transfer"`
	// Indicates if cancellation is still possible.
	CanCancelTransfer bool `json:"can_cancel_transfer"`
	// Privacy guards are disabled at the foreign registrar.
	DisablePrivacy DomainTransferInDisablePrivacy `json:"disable_privacy"`
	// Auth code has been entered and verified.
	EnterAuthCode DomainTransferInEnterAuthCode `json:"enter_auth_code"`
	// Domain is unlocked at the foreign registrar.
	UnlockDomain DomainTransferInUnlockDomain `json:"unlock_domain"`
	JSON         domainTransferInJSON         `json:"-"`
}

// domainTransferInJSON contains the JSON metadata for the struct
// [DomainTransferIn]
type domainTransferInJSON struct {
	AcceptFoa         apijson.Field
	ApproveTransfer   apijson.Field
	CanCancelTransfer apijson.Field
	DisablePrivacy    apijson.Field
	EnterAuthCode     apijson.Field
	UnlockDomain      apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DomainTransferIn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainTransferInJSON) RawJSON() string {
	return r.raw
}

// Form of authorization has been accepted by the registrant.
type DomainTransferInAcceptFoa string

const (
	DomainTransferInAcceptFoaNeeded DomainTransferInAcceptFoa = "needed"
	DomainTransferInAcceptFoaOk     DomainTransferInAcceptFoa = "ok"
)

func (r DomainTransferInAcceptFoa) IsKnown() bool {
	switch r {
	case DomainTransferInAcceptFoaNeeded, DomainTransferInAcceptFoaOk:
		return true
	}
	return false
}

// Shows transfer status with the registry.
type DomainTransferInApproveTransfer string

const (
	DomainTransferInApproveTransferNeeded   DomainTransferInApproveTransfer = "needed"
	DomainTransferInApproveTransferOk       DomainTransferInApproveTransfer = "ok"
	DomainTransferInApproveTransferPending  DomainTransferInApproveTransfer = "pending"
	DomainTransferInApproveTransferTrying   DomainTransferInApproveTransfer = "trying"
	DomainTransferInApproveTransferRejected DomainTransferInApproveTransfer = "rejected"
	DomainTransferInApproveTransferUnknown  DomainTransferInApproveTransfer = "unknown"
)

func (r DomainTransferInApproveTransfer) IsKnown() bool {
	switch r {
	case DomainTransferInApproveTransferNeeded, DomainTransferInApproveTransferOk, DomainTransferInApproveTransferPending, DomainTransferInApproveTransferTrying, DomainTransferInApproveTransferRejected, DomainTransferInApproveTransferUnknown:
		return true
	}
	return false
}

// Privacy guards are disabled at the foreign registrar.
type DomainTransferInDisablePrivacy string

const (
	DomainTransferInDisablePrivacyNeeded  DomainTransferInDisablePrivacy = "needed"
	DomainTransferInDisablePrivacyOk      DomainTransferInDisablePrivacy = "ok"
	DomainTransferInDisablePrivacyUnknown DomainTransferInDisablePrivacy = "unknown"
)

func (r DomainTransferInDisablePrivacy) IsKnown() bool {
	switch r {
	case DomainTransferInDisablePrivacyNeeded, DomainTransferInDisablePrivacyOk, DomainTransferInDisablePrivacyUnknown:
		return true
	}
	return false
}

// Auth code has been entered and verified.
type DomainTransferInEnterAuthCode string

const (
	DomainTransferInEnterAuthCodeNeeded   DomainTransferInEnterAuthCode = "needed"
	DomainTransferInEnterAuthCodeOk       DomainTransferInEnterAuthCode = "ok"
	DomainTransferInEnterAuthCodePending  DomainTransferInEnterAuthCode = "pending"
	DomainTransferInEnterAuthCodeTrying   DomainTransferInEnterAuthCode = "trying"
	DomainTransferInEnterAuthCodeRejected DomainTransferInEnterAuthCode = "rejected"
)

func (r DomainTransferInEnterAuthCode) IsKnown() bool {
	switch r {
	case DomainTransferInEnterAuthCodeNeeded, DomainTransferInEnterAuthCodeOk, DomainTransferInEnterAuthCodePending, DomainTransferInEnterAuthCodeTrying, DomainTransferInEnterAuthCodeRejected:
		return true
	}
	return false
}

// Domain is unlocked at the foreign registrar.
type DomainTransferInUnlockDomain string

const (
	DomainTransferInUnlockDomainNeeded  DomainTransferInUnlockDomain = "needed"
	DomainTransferInUnlockDomainOk      DomainTransferInUnlockDomain = "ok"
	DomainTransferInUnlockDomainPending DomainTransferInUnlockDomain = "pending"
	DomainTransferInUnlockDomainTrying  DomainTransferInUnlockDomain = "trying"
	DomainTransferInUnlockDomainUnknown DomainTransferInUnlockDomain = "unknown"
)

func (r DomainTransferInUnlockDomain) IsKnown() bool {
	switch r {
	case DomainTransferInUnlockDomainNeeded, DomainTransferInUnlockDomainOk, DomainTransferInUnlockDomainPending, DomainTransferInUnlockDomainTrying, DomainTransferInUnlockDomainUnknown:
		return true
	}
	return false
}

type DomainUpdateResponse = interface{}

type DomainGetResponse = interface{}

type DomainUpdateParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Auto-renew controls whether subscription is automatically renewed upon domain
	// expiration.
	AutoRenew param.Field[bool] `json:"auto_renew"`
	// Shows whether a registrar lock is in place for a domain.
	Locked param.Field[bool] `json:"locked"`
	// Privacy option controls redacting WHOIS information.
	Privacy param.Field[bool] `json:"privacy"`
}

func (r DomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DomainUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DomainUpdateResponse  `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainUpdateResponseEnvelopeJSON    `json:"-"`
}

// domainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainUpdateResponseEnvelope]
type domainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainUpdateResponseEnvelopeSuccess bool

const (
	DomainUpdateResponseEnvelopeSuccessTrue DomainUpdateResponseEnvelopeSuccess = true
)

func (r DomainUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DomainListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type DomainGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   DomainGetResponse     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    domainGetResponseEnvelopeJSON    `json:"-"`
}

// domainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DomainGetResponseEnvelope]
type domainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r domainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DomainGetResponseEnvelopeSuccess bool

const (
	DomainGetResponseEnvelopeSuccessTrue DomainGetResponseEnvelopeSuccess = true
)

func (r DomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
