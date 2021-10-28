//go:generate go run ../../gen/model_response/main.go -package billing -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package billing -source model.go -destination model_paginated_generated.go

package billing

import "github.com/ukfast/sdk-go/pkg/connection"

// Card represents a credit/debit card
// +genie:model_response
// +genie:model_paginated
type Card struct {
	ID           int    `json:"id"`
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

// CloudCost represents a cloud cost item
// +genie:model_response
// +genie:model_paginated
type CloudCost struct {
	ID       int `json:"id"`
	ServerID int `json:"server_id"`
	Resource struct {
		Type                   string              `json:"type"`
		Quantity               int                 `json:"quantity"`
		Price                  float32             `json:"price"`
		Period                 string              `json:"period"`
		UsageSinceLastInvoice  int                 `json:"usage_since_last_invoice"`
		CostSinceLastInvoice   float32             `json:"cost_since_last_invoice"`
		UsageForPeriodEstimate int                 `json:"usage_for_period_estimate"`
		CostForPeriodEstimate  float32             `json:"cost_for_period_estimate"`
		BillingStart           connection.DateTime `json:"billing_start"`
		BillingEnd             connection.DateTime `json:"billing_end"`
		BillingDueDate         connection.DateTime `json:"billing_due_date"`
	} `json:"resource"`
}

// DirectDebit represents a direct debit
// +genie:model_response
type DirectDebit struct {
	Name           string              `json:"name"`
	Number         string              `json:"number"`
	SortCode       string              `json:"sortcode"`
	IsActivated    bool                `json:"is_activated"`
	Status         string              `json:"status"`
	DaysCredit     int                 `json:"days_credit"`
	SignupDateTime connection.DateTime `json:"signup_datetime"`
	SignupSource   string              `json:"signup_source"`
}

// InvoiceQuery represents an invoice query
// +genie:model_response
// +genie:model_paginated
type InvoiceQuery struct {
	ID               int                 `json:"id"`
	ContactID        int                 `json:"contact_id"`
	Amount           float32             `json:"amount"`
	WhatWasExpected  string              `json:"what_was_expected"`
	WhatWasReceived  string              `json:"what_was_received"`
	ProposedSolution string              `json:"proposed_solution"`
	InvoiceIDs       []int               `json:"invoice_ids"`
	Resolution       bool                `json:"resolution"`
	ResolutionDate   connection.DateTime `json:"resolution_date"`
	Status           string              `json:"status"`
	Date             connection.DateTime `json:"date"`
}

// Invoice represents an invoice
// +genie:model_response
// +genie:model_paginated
type Invoice struct {
	ID             int                 `json:"id"`
	Date           connection.DateTime `json:"date"`
	Paid           bool                `json:"paid"`
	Gross          float32             `json:"gross"`
	VAT            float32             `json:"vat"`
	Net            float32             `json:"net"`
	Outstanding    float32             `json:"outstanding"`
	ViaDirectDebit bool                `json:"via_direct_debit"`
}

// Payment represents a payment
// +genie:model_response
// +genie:model_paginated
type Payment struct {
	ID          int                 `json:"id"`
	Category    string              `json:"category"`
	Quantity    int                 `json:"quantity"`
	Description string              `json:"description"`
	Date        connection.DateTime `json:"date"`
	DateFrom    connection.DateTime `json:"date_from"`
	DateTo      connection.DateTime `json:"date_to"`
	Cost        float32             `json:"cost"`
	VAT         float32             `json:"vat"`
	Gross       float32             `json:"gross"`
	Discount    float32             `json:"discount"`
}

// RecurringCost represents a recurring cost
// +genie:model_response
// +genie:model_paginated
type RecurringCost struct {
	ID   int `json:"id"`
	Type struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"type"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	OrderID         string `json:"order_id"`
	PurchaseOrderID string `json:"purchase_order_id"`
	CostCentreID    int    `json:"cost_centre_id"`
	Product         struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"product"`
	Cost                     float32             `json:"cost"`
	Period                   string              `json:"period"`
	Interval                 int                 `json:"interval"`
	ByCard                   bool                `json:"by_card"`
	NextPaymentAt            connection.Date     `json:"next_payment_at"`
	EndDate                  connection.Date     `json:"end_date"`
	ContractEndDate          connection.Date     `json:"contract_end_date"`
	FrozenEndDate            connection.Date     `json:"frozen_end_date"`
	MigrationEndDate         connection.Date     `json:"migration_end_date"`
	ExtendedMigrationEndDate connection.Date     `json:"extended_migration_end_date"`
	CreatedAt                connection.DateTime `json:"created_at"`
	Partner                  struct {
		ID   int    `json:"id"`
		Cost string `json:"cost"`
	} `json:"partner"`
	ProjectID int `json:"project_id"`
}
