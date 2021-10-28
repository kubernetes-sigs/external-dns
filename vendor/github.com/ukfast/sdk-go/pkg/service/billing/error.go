package billing

import "fmt"

// CardNotFoundError indicates a card was not found
type CardNotFoundError struct {
	ID int
}

func (e *CardNotFoundError) Error() string {
	return fmt.Sprintf("Card not found with ID [%d]", e.ID)
}

// CloudCostNotFoundError indicates a cloud cost was not found
type CloudCostNotFoundError struct {
	ID int
}

func (e *CloudCostNotFoundError) Error() string {
	return fmt.Sprintf("Cloud cost not found with ID [%d]", e.ID)
}

// DirectDebitNotFoundError indicates direct debit details were not found
type DirectDebitNotFoundError struct {
}

func (e *DirectDebitNotFoundError) Error() string {
	return "Direct debit details not found"
}

// InvoiceNotFoundError indicates an invoice was not found
type InvoiceNotFoundError struct {
	ID int
}

func (e *InvoiceNotFoundError) Error() string {
	return fmt.Sprintf("Invoice not found with ID [%d]", e.ID)
}

// InvoiceQueryNotFoundError indicates an invoice query was not found
type InvoiceQueryNotFoundError struct {
	ID int
}

func (e *InvoiceQueryNotFoundError) Error() string {
	return fmt.Sprintf("Invoice query not found with ID [%d]", e.ID)
}

// PaymentNotFoundError indicates a payment was not found
type PaymentNotFoundError struct {
	ID int
}

func (e *PaymentNotFoundError) Error() string {
	return fmt.Sprintf("Payment not found with ID [%d]", e.ID)
}

// RecurringCostNotFoundError indicates a recurring cost was not found
type RecurringCostNotFoundError struct {
	ID int
}

func (e *RecurringCostNotFoundError) Error() string {
	return fmt.Sprintf("Recurring cost not found with ID [%d]", e.ID)
}
