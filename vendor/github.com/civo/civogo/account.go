package civogo

import (
	"bytes"
	"encoding/json"
)

// PaginatedAccounts returns a paginated list of Account object
type PaginatedAccounts struct {
	Page    int       `json:"page"`
	PerPage int       `json:"per_page"`
	Pages   int       `json:"pages"`
	Items   []Account `json:"items"`
}

// ListAccounts lists all accounts
func (c *Client) ListAccounts() (*PaginatedAccounts, error) {
	resp, err := c.SendGetRequest("/v2/accounts")
	if err != nil {
		return nil, decodeError(err)
	}

	accounts := &PaginatedAccounts{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&accounts); err != nil {
		return nil, decodeError(err)
	}

	return accounts, nil
}

// GetAccountID returns the account ID
func (c *Client) GetAccountID() string {
	accounts, err := c.ListAccounts()
	if err != nil {
		return ""
	}

	if len(accounts.Items) == 0 {
		return "No account found"
	}

	return accounts.Items[0].ID
}
