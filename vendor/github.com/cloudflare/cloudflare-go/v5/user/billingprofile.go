// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// BillingProfileService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingProfileService] method instead.
type BillingProfileService struct {
	Options []option.RequestOption
}

// NewBillingProfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillingProfileService(opts ...option.RequestOption) (r *BillingProfileService) {
	r = &BillingProfileService{}
	r.Options = opts
	return
}

// Accesses your billing profile object.
//
// Deprecated: deprecated
func (r *BillingProfileService) Get(ctx context.Context, opts ...option.RequestOption) (res *BillingProfileGetResponse, err error) {
	var env BillingProfileGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "user/billing/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BillingProfileGetResponse struct {
	// Billing item identifier tag.
	ID                     string                        `json:"id"`
	AccountType            string                        `json:"account_type"`
	Address                string                        `json:"address"`
	Address2               string                        `json:"address2"`
	Balance                string                        `json:"balance"`
	CardExpiryMonth        int64                         `json:"card_expiry_month"`
	CardExpiryYear         int64                         `json:"card_expiry_year"`
	CardNumber             string                        `json:"card_number"`
	City                   string                        `json:"city"`
	Company                string                        `json:"company"`
	Country                string                        `json:"country"`
	CreatedOn              time.Time                     `json:"created_on" format:"date-time"`
	DeviceData             string                        `json:"device_data"`
	EditedOn               time.Time                     `json:"edited_on" format:"date-time"`
	EnterpriseBillingEmail string                        `json:"enterprise_billing_email"`
	EnterprisePrimaryEmail string                        `json:"enterprise_primary_email"`
	FirstName              string                        `json:"first_name"`
	IsPartner              bool                          `json:"is_partner"`
	LastName               string                        `json:"last_name"`
	NextBillDate           time.Time                     `json:"next_bill_date" format:"date-time"`
	PaymentAddress         string                        `json:"payment_address"`
	PaymentAddress2        string                        `json:"payment_address2"`
	PaymentCity            string                        `json:"payment_city"`
	PaymentCountry         string                        `json:"payment_country"`
	PaymentEmail           string                        `json:"payment_email"`
	PaymentFirstName       string                        `json:"payment_first_name"`
	PaymentGateway         string                        `json:"payment_gateway"`
	PaymentLastName        string                        `json:"payment_last_name"`
	PaymentNonce           string                        `json:"payment_nonce"`
	PaymentState           string                        `json:"payment_state"`
	PaymentZipcode         string                        `json:"payment_zipcode"`
	PrimaryEmail           string                        `json:"primary_email"`
	State                  string                        `json:"state"`
	TaxIDType              string                        `json:"tax_id_type"`
	Telephone              string                        `json:"telephone"`
	UseLegacy              bool                          `json:"use_legacy"`
	ValidationCode         string                        `json:"validation_code"`
	Vat                    string                        `json:"vat"`
	Zipcode                string                        `json:"zipcode"`
	JSON                   billingProfileGetResponseJSON `json:"-"`
}

// billingProfileGetResponseJSON contains the JSON metadata for the struct
// [BillingProfileGetResponse]
type billingProfileGetResponseJSON struct {
	ID                     apijson.Field
	AccountType            apijson.Field
	Address                apijson.Field
	Address2               apijson.Field
	Balance                apijson.Field
	CardExpiryMonth        apijson.Field
	CardExpiryYear         apijson.Field
	CardNumber             apijson.Field
	City                   apijson.Field
	Company                apijson.Field
	Country                apijson.Field
	CreatedOn              apijson.Field
	DeviceData             apijson.Field
	EditedOn               apijson.Field
	EnterpriseBillingEmail apijson.Field
	EnterprisePrimaryEmail apijson.Field
	FirstName              apijson.Field
	IsPartner              apijson.Field
	LastName               apijson.Field
	NextBillDate           apijson.Field
	PaymentAddress         apijson.Field
	PaymentAddress2        apijson.Field
	PaymentCity            apijson.Field
	PaymentCountry         apijson.Field
	PaymentEmail           apijson.Field
	PaymentFirstName       apijson.Field
	PaymentGateway         apijson.Field
	PaymentLastName        apijson.Field
	PaymentNonce           apijson.Field
	PaymentState           apijson.Field
	PaymentZipcode         apijson.Field
	PrimaryEmail           apijson.Field
	State                  apijson.Field
	TaxIDType              apijson.Field
	Telephone              apijson.Field
	UseLegacy              apijson.Field
	ValidationCode         apijson.Field
	Vat                    apijson.Field
	Zipcode                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *BillingProfileGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingProfileGetResponseJSON) RawJSON() string {
	return r.raw
}

type BillingProfileGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   BillingProfileGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success BillingProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    billingProfileGetResponseEnvelopeJSON    `json:"-"`
}

// billingProfileGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BillingProfileGetResponseEnvelope]
type billingProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r billingProfileGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type BillingProfileGetResponseEnvelopeSuccess bool

const (
	BillingProfileGetResponseEnvelopeSuccessTrue BillingProfileGetResponseEnvelopeSuccess = true
)

func (r BillingProfileGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case BillingProfileGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
