package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// DatabaseUserInfo represents the user information
type DatabaseUserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

// Database holds the database information
type Database struct {
	ID               string             `json:"id"`
	Name             string             `json:"name"`
	Nodes            int                `json:"nodes"`
	Size             string             `json:"size"`
	Software         string             `json:"software"`
	SoftwareVersion  string             `json:"software_version"`
	PublicIPv4       string             `json:"public_ipv4"`
	PrivateIPv4      string             `json:"private_ipv4"`
	NetworkID        string             `json:"network_id"`
	FirewallID       string             `json:"firewall_id"`
	Port             int                `json:"port"`
	Username         string             `json:"username"`
	Password         string             `json:"password"`
	DatabaseUserInfo []DatabaseUserInfo `json:"database_user_info"`
	DNSEntry         string             `json:"dns_entry,omitempty"`
	Status           string             `json:"status"`
}

// PaginatedDatabases is the structure for list response from DB endpoint
type PaginatedDatabases struct {
	Page    int        `json:"page"`
	PerPage int        `json:"per_page"`
	Pages   int        `json:"pages"`
	Items   []Database `json:"items"`
}

// CreateDatabaseRequest holds fields required to creates a new database
type CreateDatabaseRequest struct {
	Name            string `json:"name" validate:"required"`
	Size            string `json:"size" validate:"required"`
	Software        string `json:"software" validate:"required"`
	SoftwareVersion string `json:"software_version"`
	NetworkID       string `json:"network_id"`
	Nodes           int    `json:"nodes"`
	FirewallID      string `json:"firewall_id"`
	FirewallRules   string `json:"firewall_rule"`
	Region          string `json:"region"`
}

// UpdateDatabaseRequest holds fields required to update a database
type UpdateDatabaseRequest struct {
	Name       string `json:"name"`
	Nodes      *int   `json:"nodes"`
	FirewallID string `json:"firewall_id"`
	Region     string `json:"region"`
}

// SupportedSoftwareVersion contains the information related to a specific software version
type SupportedSoftwareVersion struct {
	SoftwareVersion string `json:"software_version"`
	Default         bool   `json:"default"`
}

// RestoreDatabaseRequest is the request body for restoring a database
type RestoreDatabaseRequest struct {
	// Name is the name of the database restore
	Name string `json:"name"`
	// Backup is the name of the database backup
	Backup string `json:"backup"`
	// Region is the name of the region
	Region string `json:"region"`
}

// ListDatabases returns a list of all databases
func (c *Client) ListDatabases() (*PaginatedDatabases, error) {
	resp, err := c.SendGetRequest("/v2/databases")
	if err != nil {
		return nil, decodeError(err)
	}

	databases := &PaginatedDatabases{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&databases); err != nil {
		return nil, err
	}

	return databases, nil
}

// GetDatabase finds a database by the database UUID
func (c *Client) GetDatabase(id string) (*Database, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/databases/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	db := &Database{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(db); err != nil {
		return nil, err
	}

	return db, nil
}

// DeleteDatabase deletes a database
func (c *Client) DeleteDatabase(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/databases/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// NewDatabase creates a new database
func (c *Client) NewDatabase(v *CreateDatabaseRequest) (*Database, error) {
	body, err := c.SendPostRequest("/v2/databases", v)
	if err != nil {
		return nil, decodeError(err)
	}

	result := &Database{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateDatabase updates a database
func (c *Client) UpdateDatabase(id string, v *UpdateDatabaseRequest) (*Database, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/databases/%s", id), v)
	if err != nil {
		return nil, decodeError(err)
	}

	result := &Database{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// FindDatabase finds a database by either part of the ID or part of the name
func (c *Client) FindDatabase(search string) (*Database, error) {
	databases, err := c.ListDatabases()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := Database{}

	for _, value := range databases.Items {
		if strings.EqualFold(value.Name, search) || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(strings.ToUpper(value.Name), strings.ToUpper(search)) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// ListDBVersions returns a list of all database versions
func (c *Client) ListDBVersions() (map[string][]SupportedSoftwareVersion, error) {
	resp, err := c.SendGetRequest("/v2/databases/versions")
	if err != nil {
		return nil, decodeError(err)
	}

	versions := make(map[string][]SupportedSoftwareVersion, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&versions); err != nil {
		return nil, err
	}

	return versions, nil
}

// RestoreDatabase restore a database
func (c *Client) RestoreDatabase(id string, v *RestoreDatabaseRequest) (*SimpleResponse, error) {
	resp, err := c.SendPostRequest(fmt.Sprintf("/v2/databases/%s/restore", id), v)
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
