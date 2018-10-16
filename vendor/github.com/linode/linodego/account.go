package linodego

import "context"

// Account associated with the token in use
type Account struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Email      string
	Company    string
	Address1   string
	Address2   string
	Balance    float32
	City       string
	State      string
	Zip        string
	Country    string
	TaxID      string      `json:"tax_id"`
	CreditCard *CreditCard `json:"credit_card"`
}

// CreditCard information associated with the Account.
type CreditCard struct {
	LastFour string `json:"last_four"`
	Expiry   string
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *Account) fixDates() *Account {
	return v
}

// GetAccount gets the contact and billing information related to the Account
func (c *Client) GetAccount(ctx context.Context) (*Account, error) {
	e, err := c.Account.Endpoint()
	if err != nil {
		return nil, err
	}
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&Account{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*Account).fixDates(), nil
}
