package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// DatabaseBackup represents a backup
type DatabaseBackup struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Software     string    `json:"software,omitempty"`
	Status       string    `json:"status,omitempty"`
	Schedule     string    `json:"schedule,omitempty"`
	DatabaseName string    `json:"database_name,omitempty"`
	DatabaseID   string    `json:"database_id,omitempty"`
	IsScheduled  bool      `json:"is_scheduled,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// PaginatedDatabaseBackup is the structure for list response from DB endpoint
type PaginatedDatabaseBackup struct {
	Page    int              `json:"page"`
	PerPage int              `json:"per_page"`
	Pages   int              `json:"pages"`
	Items   []DatabaseBackup `json:"items"`
}

// DatabaseBackupCreateRequest represents a backup create request
type DatabaseBackupCreateRequest struct {
	// Name is the name of the database backup to be created
	Name string `json:"name"`
	// Schedule is used for scheduled backup
	Schedule string `json:"schedule"`
	// Types is used for manual backup
	Type string `json:"type"`
	// Region is name of the region
	Region string `json:"region"`
}

// DatabaseBackupUpdateRequest represents a backup update request
type DatabaseBackupUpdateRequest struct {
	// Name is name name of the backup
	Name string `json:"name"`
	// Schedule is schedule for scheduled backup
	Schedule string `json:"schedule"`
	// Region is name of the region
	Region string `json:"region"`
}

// ListDatabaseBackup lists backups for database
func (c *Client) ListDatabaseBackup(did string) (*PaginatedDatabaseBackup, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/databases/%s/backups", did))
	if err != nil {
		return nil, decodeError(err)
	}

	back := &PaginatedDatabaseBackup{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&back); err != nil {
		return nil, decodeError(err)
	}

	return back, nil
}

// UpdateDatabaseBackup update database backup
func (c *Client) UpdateDatabaseBackup(did string, v *DatabaseBackupUpdateRequest) (*DatabaseBackup, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/databases/%s/backups", did), v)
	if err != nil {
		return nil, decodeError(err)
	}

	result := &DatabaseBackup{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateDatabaseBackup create database backup
func (c *Client) CreateDatabaseBackup(did string, v *DatabaseBackupCreateRequest) (*DatabaseBackup, error) {
	body, err := c.SendPostRequest(fmt.Sprintf("/v2/databases/%s/backups", did), v)
	if err != nil {
		return nil, decodeError(err)
	}

	result := &DatabaseBackup{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteDatabaseBackup deletes a database backup
func (c *Client) DeleteDatabaseBackup(dbid, id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/databases/%s/backups/%s", dbid, id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// GetDatabaseBackup finds a database by the database UUID
func (c *Client) GetDatabaseBackup(dbid, id string) (*DatabaseBackup, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/databases/%s/backups/%s", dbid, id))
	if err != nil {
		return nil, decodeError(err)
	}

	bk := &DatabaseBackup{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(bk); err != nil {
		return nil, err
	}

	return bk, nil
}

// FindDatabaseBackup finds a database by either part of the ID or part of the name
func (c *Client) FindDatabaseBackup(dbid, search string) (*DatabaseBackup, error) {
	backups, err := c.ListDatabaseBackup(dbid)
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := DatabaseBackup{}

	for _, value := range backups.Items {
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
