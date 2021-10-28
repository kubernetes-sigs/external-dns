package billing

import "github.com/ukfast/sdk-go/pkg/connection"

// CreateCardRequest represents a request to create a card
type CreateCardRequest struct {
	FriendlyName string `json:"friendly_name"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Postcode     string `json:"postcode"`
	CardNumber   string `json:"card_number"`
	CardType     string `json:"card_type"`
	ValidFrom    string `json:"valid_from"`
	Expiry       string `json:"expiry"`
	IssueNumber  int    `json:"issue_number"`
	PrimaryCard  bool   `json:"primary_card"`
}

// PatchCardRequest represents a request to create a card
type PatchCardRequest struct {
	FriendlyName string `json:"friendly_name,omitempty"`
	Name         string `json:"name,omitempty"`
	Address      string `json:"address,omitempty"`
	Postcode     string `json:"postcode,omitempty"`
	CardNumber   string `json:"card_number,omitempty"`
	CardType     string `json:"card_type,omitempty"`
	ValidFrom    string `json:"valid_from,omitempty"`
	Expiry       string `json:"expiry,omitempty"`
	IssueNumber  int    `json:"issue_number,omitempty"`
	PrimaryCard  *bool  `json:"primary_card,omitempty"`
}

// CreateInvoiceQueryRequest represents a request to create an invoice query
type CreateInvoiceQueryRequest struct {
	ContactID        int                 `json:"contact_id"`
	Amount           float32             `json:"amount"`
	WhatWasExpected  string              `json:"what_was_expected"`
	WhatWasReceived  string              `json:"what_was_received"`
	ProposedSolution string              `json:"proposed_solution"`
	ContactMethod    string              `json:"contact_method"`
	InvoiceIDs       []int               `json:"invoice_ids"`
	Resolution       string              `json:"resolution"`
	ResolutionDate   connection.DateTime `json:"resolution_date"`
	Status           string              `json:"status"`
	Date             connection.DateTime `json:"date"`
}
