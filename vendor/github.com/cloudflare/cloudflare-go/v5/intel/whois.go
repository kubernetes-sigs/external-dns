// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// WhoisService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhoisService] method instead.
type WhoisService struct {
	Options []option.RequestOption
}

// NewWhoisService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWhoisService(opts ...option.RequestOption) (r *WhoisService) {
	r = &WhoisService{}
	r.Options = opts
	return
}

// Get WHOIS Record
func (r *WhoisService) Get(ctx context.Context, params WhoisGetParams, opts ...option.RequestOption) (res *WhoisGetResponse, err error) {
	var env WhoisGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/whois", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WhoisGetResponse struct {
	DNSSEC                    bool                 `json:"dnssec,required"`
	Domain                    string               `json:"domain,required"`
	Extension                 string               `json:"extension,required"`
	Found                     bool                 `json:"found,required"`
	Nameservers               []string             `json:"nameservers,required"`
	Punycode                  string               `json:"punycode,required"`
	Registrant                string               `json:"registrant,required"`
	Registrar                 string               `json:"registrar,required"`
	ID                        string               `json:"id"`
	AdministrativeCity        string               `json:"administrative_city"`
	AdministrativeCountry     string               `json:"administrative_country"`
	AdministrativeEmail       string               `json:"administrative_email"`
	AdministrativeFax         string               `json:"administrative_fax"`
	AdministrativeFaxExt      string               `json:"administrative_fax_ext"`
	AdministrativeID          string               `json:"administrative_id"`
	AdministrativeName        string               `json:"administrative_name"`
	AdministrativeOrg         string               `json:"administrative_org"`
	AdministrativePhone       string               `json:"administrative_phone"`
	AdministrativePhoneExt    string               `json:"administrative_phone_ext"`
	AdministrativePostalCode  string               `json:"administrative_postal_code"`
	AdministrativeProvince    string               `json:"administrative_province"`
	AdministrativeReferralURL string               `json:"administrative_referral_url"`
	AdministrativeStreet      string               `json:"administrative_street"`
	BillingCity               string               `json:"billing_city"`
	BillingCountry            string               `json:"billing_country"`
	BillingEmail              string               `json:"billing_email"`
	BillingFax                string               `json:"billing_fax"`
	BillingFaxExt             string               `json:"billing_fax_ext"`
	BillingID                 string               `json:"billing_id"`
	BillingName               string               `json:"billing_name"`
	BillingOrg                string               `json:"billing_org"`
	BillingPhone              string               `json:"billing_phone"`
	BillingPhoneExt           string               `json:"billing_phone_ext"`
	BillingPostalCode         string               `json:"billing_postal_code"`
	BillingProvince           string               `json:"billing_province"`
	BillingReferralURL        string               `json:"billing_referral_url"`
	BillingStreet             string               `json:"billing_street"`
	CreatedDate               time.Time            `json:"created_date" format:"date-time"`
	CreatedDateRaw            string               `json:"created_date_raw"`
	ExpirationDate            time.Time            `json:"expiration_date" format:"date-time"`
	ExpirationDateRaw         string               `json:"expiration_date_raw"`
	RegistrantCity            string               `json:"registrant_city"`
	RegistrantCountry         string               `json:"registrant_country"`
	RegistrantEmail           string               `json:"registrant_email"`
	RegistrantFax             string               `json:"registrant_fax"`
	RegistrantFaxExt          string               `json:"registrant_fax_ext"`
	RegistrantID              string               `json:"registrant_id"`
	RegistrantName            string               `json:"registrant_name"`
	RegistrantOrg             string               `json:"registrant_org"`
	RegistrantPhone           string               `json:"registrant_phone"`
	RegistrantPhoneExt        string               `json:"registrant_phone_ext"`
	RegistrantPostalCode      string               `json:"registrant_postal_code"`
	RegistrantProvince        string               `json:"registrant_province"`
	RegistrantReferralURL     string               `json:"registrant_referral_url"`
	RegistrantStreet          string               `json:"registrant_street"`
	RegistrarCity             string               `json:"registrar_city"`
	RegistrarCountry          string               `json:"registrar_country"`
	RegistrarEmail            string               `json:"registrar_email"`
	RegistrarFax              string               `json:"registrar_fax"`
	RegistrarFaxExt           string               `json:"registrar_fax_ext"`
	RegistrarID               string               `json:"registrar_id"`
	RegistrarName             string               `json:"registrar_name"`
	RegistrarOrg              string               `json:"registrar_org"`
	RegistrarPhone            string               `json:"registrar_phone"`
	RegistrarPhoneExt         string               `json:"registrar_phone_ext"`
	RegistrarPostalCode       string               `json:"registrar_postal_code"`
	RegistrarProvince         string               `json:"registrar_province"`
	RegistrarReferralURL      string               `json:"registrar_referral_url"`
	RegistrarStreet           string               `json:"registrar_street"`
	Status                    []string             `json:"status"`
	TechnicalCity             string               `json:"technical_city"`
	TechnicalCountry          string               `json:"technical_country"`
	TechnicalEmail            string               `json:"technical_email"`
	TechnicalFax              string               `json:"technical_fax"`
	TechnicalFaxExt           string               `json:"technical_fax_ext"`
	TechnicalID               string               `json:"technical_id"`
	TechnicalName             string               `json:"technical_name"`
	TechnicalOrg              string               `json:"technical_org"`
	TechnicalPhone            string               `json:"technical_phone"`
	TechnicalPhoneExt         string               `json:"technical_phone_ext"`
	TechnicalPostalCode       string               `json:"technical_postal_code"`
	TechnicalProvince         string               `json:"technical_province"`
	TechnicalReferralURL      string               `json:"technical_referral_url"`
	TechnicalStreet           string               `json:"technical_street"`
	UpdatedDate               time.Time            `json:"updated_date" format:"date-time"`
	UpdatedDateRaw            string               `json:"updated_date_raw"`
	WhoisServer               string               `json:"whois_server"`
	JSON                      whoisGetResponseJSON `json:"-"`
}

// whoisGetResponseJSON contains the JSON metadata for the struct
// [WhoisGetResponse]
type whoisGetResponseJSON struct {
	DNSSEC                    apijson.Field
	Domain                    apijson.Field
	Extension                 apijson.Field
	Found                     apijson.Field
	Nameservers               apijson.Field
	Punycode                  apijson.Field
	Registrant                apijson.Field
	Registrar                 apijson.Field
	ID                        apijson.Field
	AdministrativeCity        apijson.Field
	AdministrativeCountry     apijson.Field
	AdministrativeEmail       apijson.Field
	AdministrativeFax         apijson.Field
	AdministrativeFaxExt      apijson.Field
	AdministrativeID          apijson.Field
	AdministrativeName        apijson.Field
	AdministrativeOrg         apijson.Field
	AdministrativePhone       apijson.Field
	AdministrativePhoneExt    apijson.Field
	AdministrativePostalCode  apijson.Field
	AdministrativeProvince    apijson.Field
	AdministrativeReferralURL apijson.Field
	AdministrativeStreet      apijson.Field
	BillingCity               apijson.Field
	BillingCountry            apijson.Field
	BillingEmail              apijson.Field
	BillingFax                apijson.Field
	BillingFaxExt             apijson.Field
	BillingID                 apijson.Field
	BillingName               apijson.Field
	BillingOrg                apijson.Field
	BillingPhone              apijson.Field
	BillingPhoneExt           apijson.Field
	BillingPostalCode         apijson.Field
	BillingProvince           apijson.Field
	BillingReferralURL        apijson.Field
	BillingStreet             apijson.Field
	CreatedDate               apijson.Field
	CreatedDateRaw            apijson.Field
	ExpirationDate            apijson.Field
	ExpirationDateRaw         apijson.Field
	RegistrantCity            apijson.Field
	RegistrantCountry         apijson.Field
	RegistrantEmail           apijson.Field
	RegistrantFax             apijson.Field
	RegistrantFaxExt          apijson.Field
	RegistrantID              apijson.Field
	RegistrantName            apijson.Field
	RegistrantOrg             apijson.Field
	RegistrantPhone           apijson.Field
	RegistrantPhoneExt        apijson.Field
	RegistrantPostalCode      apijson.Field
	RegistrantProvince        apijson.Field
	RegistrantReferralURL     apijson.Field
	RegistrantStreet          apijson.Field
	RegistrarCity             apijson.Field
	RegistrarCountry          apijson.Field
	RegistrarEmail            apijson.Field
	RegistrarFax              apijson.Field
	RegistrarFaxExt           apijson.Field
	RegistrarID               apijson.Field
	RegistrarName             apijson.Field
	RegistrarOrg              apijson.Field
	RegistrarPhone            apijson.Field
	RegistrarPhoneExt         apijson.Field
	RegistrarPostalCode       apijson.Field
	RegistrarProvince         apijson.Field
	RegistrarReferralURL      apijson.Field
	RegistrarStreet           apijson.Field
	Status                    apijson.Field
	TechnicalCity             apijson.Field
	TechnicalCountry          apijson.Field
	TechnicalEmail            apijson.Field
	TechnicalFax              apijson.Field
	TechnicalFaxExt           apijson.Field
	TechnicalID               apijson.Field
	TechnicalName             apijson.Field
	TechnicalOrg              apijson.Field
	TechnicalPhone            apijson.Field
	TechnicalPhoneExt         apijson.Field
	TechnicalPostalCode       apijson.Field
	TechnicalProvince         apijson.Field
	TechnicalReferralURL      apijson.Field
	TechnicalStreet           apijson.Field
	UpdatedDate               apijson.Field
	UpdatedDateRaw            apijson.Field
	WhoisServer               apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WhoisGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseJSON) RawJSON() string {
	return r.raw
}

type WhoisGetParams struct {
	// Use to uniquely identify or reference the resource.
	AccountID param.Field[string] `path:"account_id,required"`
	Domain    param.Field[string] `query:"domain"`
}

// URLQuery serializes [WhoisGetParams]'s query parameters as `url.Values`.
func (r WhoisGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WhoisGetResponseEnvelope struct {
	Errors   []WhoisGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WhoisGetResponseEnvelopeMessages `json:"messages,required"`
	// Returns a boolean for the success/failure of the API call.
	Success WhoisGetResponseEnvelopeSuccess `json:"success,required"`
	Result  WhoisGetResponse                `json:"result"`
	JSON    whoisGetResponseEnvelopeJSON    `json:"-"`
}

// whoisGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WhoisGetResponseEnvelope]
type whoisGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WhoisGetResponseEnvelopeErrors struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           WhoisGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             whoisGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// whoisGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WhoisGetResponseEnvelopeErrors]
type whoisGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WhoisGetResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    whoisGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// whoisGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [WhoisGetResponseEnvelopeErrorsSource]
type whoisGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type WhoisGetResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           WhoisGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             whoisGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// whoisGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [WhoisGetResponseEnvelopeMessages]
type whoisGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type WhoisGetResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    whoisGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// whoisGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [WhoisGetResponseEnvelopeMessagesSource]
type whoisGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Returns a boolean for the success/failure of the API call.
type WhoisGetResponseEnvelopeSuccess bool

const (
	WhoisGetResponseEnvelopeSuccessTrue WhoisGetResponseEnvelopeSuccess = true
)

func (r WhoisGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WhoisGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
