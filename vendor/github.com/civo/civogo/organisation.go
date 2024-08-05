package civogo

import (
	"bytes"
	"encoding/json"
	"time"
)

// Organisation represents a group of accounts treated as a single entity
type Organisation struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Account is the owner of Civo resources such as instances, Kubernetes clusters, volumes, etc
// Really the Account should be defined with Account endpoints, but there aren't any that are
// publicly-useful
type Account struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	Label           string    `json:"label,omitempty"`
	EmailAddress    string    `json:"email_address,omitempty"`
	APIKey          string    `json:"api_key,omitempty"`
	Token           string    `json:"token,omitempty"`
	Flags           string    `json:"flags,omitempty"`
	Timezone        string    `json:"timezone,omitempty"`
	Partner         string    `json:"partner,omitempty"`
	DefaultUserID   string    `json:"default_user_id,omitempty"`
	Status          string    `json:"status,omitempty"`
	EmailConfirmed  bool      `json:"email_confirmed,omitempty"`
	CreditCardAdded bool      `json:"credit_card_added,omitempty"`
	Enabled         bool      `json:"enabled,omitempty"`
}

// GetOrganisation returns the organisation associated with the current account
func (c *Client) GetOrganisation() (*Organisation, error) {
	resp, err := c.SendGetRequest("/v2/organisation")
	if err != nil {
		return nil, decodeError(err)
	}

	organisation := &Organisation{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(organisation); err != nil {
		return nil, err
	}

	return organisation, nil
}

// CreateOrganisation creates an organisation with the current account as the only linked member (errors if it's already linked)
func (c *Client) CreateOrganisation(name string) (*Organisation, error) {
	data := map[string]string{"name": name}
	resp, err := c.SendPostRequest("/v2/organisation", data)
	if err != nil {
		return nil, decodeError(err)
	}

	organisation := &Organisation{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(organisation); err != nil {
		return nil, err
	}

	return organisation, nil
}

// RenameOrganisation changes the human set name of the organisation (e.g. for re-branding efforts)
func (c *Client) RenameOrganisation(name string) (*Organisation, error) {
	data := map[string]string{"name": name}
	resp, err := c.SendPutRequest("/v2/organisation", data)
	if err != nil {
		return nil, decodeError(err)
	}

	organisation := &Organisation{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(organisation); err != nil {
		return nil, err
	}

	return organisation, nil
}

// AddAccountToOrganisation sets the link between second, third, etc accounts and the existing organisation
func (c *Client) AddAccountToOrganisation(organisationID, organisationToken string) ([]Account, error) {
	data := map[string]string{"organisation_id": organisationID, "organisation_token": organisationToken}
	resp, err := c.SendPostRequest("/v2/organisation/accounts", data)
	if err != nil {
		return nil, decodeError(err)
	}

	accounts := make([]Account, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}

// ListAccountsInOrganisation returns all the accounts in the current account's organisation
func (c *Client) ListAccountsInOrganisation() ([]Account, error) {
	resp, err := c.SendGetRequest("/v2/organisation/accounts")
	if err != nil {
		return nil, decodeError(err)
	}

	accounts := make([]Account, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&accounts); err != nil {
		return nil, err
	}

	return accounts, nil
}
