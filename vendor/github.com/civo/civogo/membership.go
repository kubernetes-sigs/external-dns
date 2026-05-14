package civogo

import (
	"bytes"
	"encoding/json"
)

// MembershipResponse is the response for the memberships of a user
type MembershipResponse struct {
	Accounts      []MembershipAccount
	Organisations []MembershipOrganisation
}

// MembershipAccount is the DTO for an account.
type MembershipAccount struct {
	ID             string `json:"id"`
	EmailAddress   string `json:"email_address"`
	Label          string `json:"label"`
	OrganisationID string `json:"organisation_id,omitempty"`
}

// MembershipOrganisation is the DTO for an organisation.
type MembershipOrganisation struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ListMemberships returns all the memberships(to accounts and organisations) for the user
func (c *Client) ListMemberships() (*MembershipResponse, error) {
	resp, err := c.SendGetRequest("/v2/memberships")
	if err != nil {
		return nil, decodeError(err)
	}

	mrs := &MembershipResponse{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&mrs); err != nil {
		return nil, err
	}

	return mrs, nil
}
