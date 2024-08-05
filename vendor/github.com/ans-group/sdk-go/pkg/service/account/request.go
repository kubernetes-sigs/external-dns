package account

// CreateInvoiceQueryRequest represents a request to create an invoice query
type CreateInvoiceQueryRequest struct {
	ContactID        int     `json:"contact_id"`
	ContactMethod    string  `json:"contact_method"`
	Amount           float32 `json:"amount"`
	WhatWasExpected  string  `json:"what_was_expected"`
	WhatWasReceived  string  `json:"what_was_received"`
	ProposedSolution string  `json:"proposed_solution"`
	InvoiceIDs       []int   `json:"invoice_ids"`
}
<<<<<<< HEAD
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======

// CreateClientRequest represents a request to create a client
type CreateClientRequest struct {
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
}

// PatchClientRequest represents a request to update a client
type PatchClientRequest struct {
	CompanyName      string `json:"company_name,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	EmailAddress     string `json:"email_address,omitempty"`
	LimitedNumber    string `json:"limited_number,omitempty"`
	VATNumber        string `json:"vat_number,omitempty"`
	Address          string `json:"address,omitempty"`
	Address1         string `json:"address1,omitempty"`
	City             string `json:"city,omitempty"`
	County           string `json:"county,omitempty"`
	Country          string `json:"country,omitempty"`
	Postcode         string `json:"postcode,omitempty"`
	Phone            string `json:"phone,omitempty"`
	Fax              string `json:"fax,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
	Type             string `json:"type,omitempty"`
	UserName         string `json:"user_name,omitempty"`
	IDReference      string `json:"id_reference,omitempty"`
	NominetContactID string `json:"nominet_contact_id,omitempty"`
}
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
