package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Team is a named group of users (has many members)
type Team struct {
	ID        string    `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// TeamMember is a link record between User and Team.
type TeamMember struct {
	ID          string    `json:"id"`
	TeamID      string    `json:"team_id,omitempty"`
	UserID      string    `json:"user_id,omitempty"`
	Permissions string    `json:"permissions,omitempty"`
	Roles       string    `json:"roles,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// ListTeams returns all teams for the current account
func (c *Client) ListTeams() ([]Team, error) {
	resp, err := c.SendGetRequest("/v2/teams")
	if err != nil {
		return nil, decodeError(err)
	}

	teams := make([]Team, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&teams); err != nil {
		return nil, err
	}

	return teams, nil
}

// CreateTeam creates a new team in the account
func (c *Client) CreateTeam(name string) (*Team, error) {
	data := map[string]string{"name": name}
	resp, err := c.SendPostRequest("/v2/teams", data)
	if err != nil {
		return nil, decodeError(err)
	}

	team := &Team{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(team); err != nil {
		return nil, err
	}

	return team, nil
}

// FindTeam finds a team by either part of the ID or part of the name
func (c *Client) FindTeam(search string) (*Team, error) {
	teams, err := c.ListTeams()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Team{}

	for _, value := range teams {
		if value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s team because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s team, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// RenameTeam changes the human set name for a team
func (c *Client) RenameTeam(teamID, name string) (*Team, error) {
	data := map[string]string{"name": name}
	resp, err := c.SendPutRequest("/v2/teams/"+teamID, data)
	if err != nil {
		return nil, decodeError(err)
	}

	team := &Team{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(team); err != nil {
		return nil, err
	}

	return team, nil
}

// DeleteTeam removes a team (and therefore all team member access)
func (c *Client) DeleteTeam(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest("/v2/teams/" + id)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// ListTeamMembers returns a list of all team members (and their permissions) in the specified team
func (c *Client) ListTeamMembers(teamID string) ([]TeamMember, error) {
	resp, err := c.SendGetRequest("/v2/teams/" + teamID + "/members")
	if err != nil {
		return nil, decodeError(err)
	}

	teamMembers := make([]TeamMember, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&teamMembers); err != nil {
		return nil, err
	}

	return teamMembers, nil
}

// AddTeamMember adds a team member to the specified team, with permissions and roles (which are combinative)
func (c *Client) AddTeamMember(teamID, userID, permissions, roles string) ([]TeamMember, error) {
	data := map[string]string{"user_id": userID, "permissions": permissions, "roles": roles}
	_, err := c.SendPostRequest("/v2/teams/"+teamID+"/members", data)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.ListTeamMembers(teamID)
}

// UpdateTeamMember changes the permissions or roles for a specified team member
func (c *Client) UpdateTeamMember(teamID, teamMemberID, permissions, roles string) (*TeamMember, error) {
	data := map[string]string{"permissions": permissions, "roles": roles}
	resp, err := c.SendPostRequest("/v2/teams/"+teamID+"/members/"+teamMemberID, data)
	if err != nil {
		return nil, decodeError(err)
	}

	teamMember := &TeamMember{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(teamMember); err != nil {
		return nil, err
	}

	return teamMember, nil
}

// RemoveTeamMember removes the specified team member from the specified team
func (c *Client) RemoveTeamMember(teamID, teamMemberID string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest("/v2/teams/" + teamID + "/members/" + teamMemberID)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
