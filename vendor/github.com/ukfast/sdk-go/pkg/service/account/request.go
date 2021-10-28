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
