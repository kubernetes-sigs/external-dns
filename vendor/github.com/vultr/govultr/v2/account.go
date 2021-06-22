package govultr

import (
	"context"
	"net/http"
)

// AccountService is the interface to interact with Accounts endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/account
type AccountService interface {
	Get(ctx context.Context) (*Account, error)
}

// AccountServiceHandler handles interaction with the account methods for the Vultr API
type AccountServiceHandler struct {
	client *Client
}

type accountBase struct {
	Account *Account `json:"account"`
}

// Account represents a Vultr account
type Account struct {
	Balance           float32  `json:"balance"`
	PendingCharges    float32  `json:"pending_charges"`
	LastPaymentDate   string   `json:"last_payment_date"`
	LastPaymentAmount float32  `json:"last_payment_amount"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	ACL               []string `json:"acls"`
}

// Get Vultr account info
func (a *AccountServiceHandler) Get(ctx context.Context) (*Account, error) {
	uri := "/v2/account"
	req, err := a.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	account := new(accountBase)
	if err = a.client.DoWithContext(ctx, req, account); err != nil {
		return nil, err
	}

	return account.Account, nil
}
