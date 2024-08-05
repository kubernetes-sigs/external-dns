package civogo

import (
	"bytes"
	"encoding/json"
	"time"
)

// User is the user struct
type User struct {
	ID               string    `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CompanyName      string    `json:"company_name"`
	EmailAddress     string    `json:"email_address"`
	Status           string    `json:"status"`
	Flags            string    `json:"flags"`
	Token            string    `json:"token"`
	MarketingAllowed int       `json:"marketing_allowed"`
	DefaultAccountID string    `json:"default_account_id"`
	// DefaultAccountID string      `json:"account_id"`
	PasswordDigest   string `json:"password_digest"`
	Partner          string `json:"partner"`
	PartnerUserID    string `json:"partner_user_id"`
	ReferralID       string `json:"referral_id"`
	LastChosenRegion string `json:"last_chosen_region"`
}

// UserEverything is the combination structure for all team related data for the current user and account
type UserEverything struct {
	User          User
	Accounts      []Account
	Organisations []Organisation
	Teams         []Team
	Roles         []Role
}

// GetUserEverything returns the organisation associated with the current account
func (c *Client) GetUserEverything(userID string) (*UserEverything, error) {
	resp, err := c.SendGetRequest("/v2/users/" + userID + "/everything")
	if err != nil {
		return nil, decodeError(err)
	}

	everything := &UserEverything{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(everything); err != nil {
		return nil, err
	}

	return everything, nil
}
