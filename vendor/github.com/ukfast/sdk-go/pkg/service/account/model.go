//go:generate go run ../../gen/model_response/main.go -package account -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package account -source model.go -destination model_paginated_generated.go

package account

import "github.com/ukfast/sdk-go/pkg/connection"

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
// +genie:model_response
// +genie:model_paginated
type Contact struct {
	ID        int         `json:"id"`
	Type      ContactType `json:"type"`
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
}

// Details represents a UKFast account details
// +genie:model_response
type Details struct {
	CompanyRegistrationNumber string `json:"company_registration_number"`
	VATIdentificationNumber   string `json:"vat_identification_number"`
	PrimaryContactID          int    `json:"primary_contact_id"`
}

// Credit represents a UKFast account credit
// +genie:model_response
// +genie:model_paginated
type Credit struct {
	Type      string `json:"type"`
	Total     int    `json:"total"`
	Remaining int    `json:"remaining"`
}

// Invoice represents a UKFast account invoice
// +genie:model_response
// +genie:model_paginated
type Invoice struct {
	ID    int             `json:"id"`
	Date  connection.Date `json:"date"`
	Paid  bool            `json:"paid"`
	Net   float32         `json:"net"`
	VAT   float32         `json:"vat"`
	Gross float32         `json:"gross"`
}

// InvoiceQuery represents a UKFast account invoice query
// +genie:model_response
// +genie:model_paginated
type InvoiceQuery struct {
	ID               int     `json:"id"`
	ContactID        int     `json:"contact_id"`
	Amount           float32 `json:"amount"`
	WhatWasExpected  string  `json:"what_was_expected"`
	WhatWasReceived  string  `json:"what_was_received"`
	ProposedSolution string  `json:"proposed_solution"`
	InvoiceIDs       []int   `json:"invoice_ids"`
}
