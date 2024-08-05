package account

import "github.com/ans-group/sdk-go/pkg/connection"

type ContactType string

func (t ContactType) String() string {
	return string(t)
}

const (
	ContactTypePrimaryContact ContactType = "Primary Contact"
	ContactTypeAccounts       ContactType = "Accounts"
	ContactTypeTechnical      ContactType = "Technical"
	ContactTypeThirdParty     ContactType = "Third Party"
	ContactTypeOther          ContactType = "Other"
)

// Contact represents a UKFast account contact
type Contact struct {
	ID        int         `json:"id"`
	Type      ContactType `json:"type"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
}

// Details represents a UKFast account details
type Details struct {
	CompanyRegistrationNumber string `json:"company_registration_number"`
	VATIdentificationNumber   string `json:"vat_identification_number"`
	PrimaryContactID          int    `json:"primary_contact_id"`
}

// Credit represents a UKFast account credit
type Credit struct {
	Type      string `json:"type"`
	Total     int    `json:"total"`
	Remaining int    `json:"remaining"`
}

// Invoice represents a UKFast account invoice
type Invoice struct {
	ID    int             `json:"id"`
	Date  connection.Date `json:"date"`
	Paid  bool            `json:"paid"`
	Net   float32         `json:"net"`
	VAT   float32         `json:"vat"`
	Gross float32         `json:"gross"`
}

// InvoiceQuery represents a UKFast account invoice query
type InvoiceQuery struct {
	ID               int     `json:"id"`
	ContactID        int     `json:"contact_id"`
	Amount           float32 `json:"amount"`
	WhatWasExpected  string  `json:"what_was_expected"`
	WhatWasReceived  string  `json:"what_was_received"`
	ProposedSolution string  `json:"proposed_solution"`
	InvoiceIDs       []int   `json:"invoice_ids"`
}
<<<<<<< HEAD
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======

// Client represents an account client
type Client struct {
	ID               int    `json:"id"`
	CompanyName      string `json:"company_name"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	EmailAddress     string `json:"email_address"`
	LimitedNumber    string `json:"limited_number"`
	VATNumber        string `json:"vat_number"`
	Address          string `json:"address"`
	Address1         string `json:"address1"`
	City             string `json:"city"`
	County           string `json:"county"`
	Country          string `json:"country"`
	Postcode         string `json:"postcode"`
	Phone            string `json:"phone"`
	Fax              string `json:"fax"`
	Mobile           string `json:"mobile"`
	Type             string `json:"type"`
	UserName         string `json:"user_name"`
	IDReference      string `json:"id_reference"`
	NominetContactID string `json:"nominet_contact_id"`
	CreatedDate      string `json:"created_date"`
}
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
