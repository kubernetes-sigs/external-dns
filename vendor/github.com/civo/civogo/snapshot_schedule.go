package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// SnapshotSchedule represents a snapshot schedule configuration
type SnapshotSchedule struct {
	ID             string                 `json:"id"`
	Name           string                 `json:"name"`
	Description    string                 `json:"description"`
	CronExpression string                 `json:"cron_expression"`
	Paused         bool                   `json:"paused"`
	Retention      SnapshotRetention      `json:"retention"`
	Instances      []SnapshotInstance     `json:"instances"`
	Status         SnapshotScheduleStatus `json:"status"`
	CreatedAt      time.Time              `json:"created_at"`
}

// SnapshotRetention defines how snapshots should be retained
type SnapshotRetention struct {
	Period       string `json:"period,omitempty"`
	MaxSnapshots int    `json:"max_snapshots,omitempty"`
}

// SnapshotInstance represents an instance to be snapshotted
type SnapshotInstance struct {
	ID              string   `json:"id"`
	Size            string   `json:"size,omitempty"`
	IncludedVolumes []string `json:"included_volumes,omitempty"`
}

// SnapshotScheduleStatus represents the current status of a snapshot schedule
type SnapshotScheduleStatus struct {
	State        string           `json:"state"`
	LastSnapshot LastSnapshotInfo `json:"last_snapshot"`
}

// LastSnapshotInfo contains information about the last snapshot taken
type LastSnapshotInfo struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

// CreateSnapshotScheduleRequest represents the request to create a new snapshot schedule
type CreateSnapshotScheduleRequest struct {
	Name           string                   `json:"name"`
	Description    string                   `json:"description,omitempty"`
	CronExpression string                   `json:"cron_expression"`
	Retention      SnapshotRetention        `json:"retention"`
	Instances      []CreateSnapshotInstance `json:"instances"`
}

// CreateSnapshotInstance represents an instance in the create request
type CreateSnapshotInstance struct {
	InstanceID     string `json:"instance_id"`
	IncludeVolumes bool   `json:"include_volumes"`
}

// UpdateSnapshotScheduleRequest represents the request to update a snapshot schedule
type UpdateSnapshotScheduleRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Paused      *bool  `json:"paused,omitempty"`
}

// CreateSnapshotSchedule creates a new snapshot schedule
func (c *Client) CreateSnapshotSchedule(r *CreateSnapshotScheduleRequest) (*SnapshotSchedule, error) {
	body, err := c.SendPostRequest("/v2/resourcesnapshotschedules", r)
	if err != nil {
		return nil, decodeError(err)
	}

	var schedule = &SnapshotSchedule{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

// ListSnapshotSchedules returns all snapshot schedules
func (c *Client) ListSnapshotSchedules() ([]SnapshotSchedule, error) {
	resp, err := c.SendGetRequest("/v2/resourcesnapshotschedules")
	if err != nil {
		return nil, decodeError(err)
	}

	schedules := make([]SnapshotSchedule, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&schedules); err != nil {
		return nil, err
	}

	return schedules, nil
}

// FindSnapshotSchedule finds a snapshot schedule by ID or name
func (c *Client) FindSnapshotSchedule(search string) (*SnapshotSchedule, error) {
	schedules, err := c.ListSnapshotSchedules()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := SnapshotSchedule{}

	for _, value := range schedules {
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
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// GetSnapshotSchedule retrieves a specific snapshot schedule by ID
func (c *Client) GetSnapshotSchedule(id string) (*SnapshotSchedule, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/resourcesnapshotschedules/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	schedule := &SnapshotSchedule{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

// DeleteSnapshotSchedule deletes a snapshot schedule
func (c *Client) DeleteSnapshotSchedule(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/resourcesnapshotschedules/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UpdateSnapshotSchedule updates a snapshot schedule
func (c *Client) UpdateSnapshotSchedule(id string, r *UpdateSnapshotScheduleRequest) (*SnapshotSchedule, error) {
	body, err := c.SendPutRequest(fmt.Sprintf("/v2/resourcesnapshotschedules/%s", id), r)
	if err != nil {
		return nil, decodeError(err)
	}

	var schedule = &SnapshotSchedule{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}
