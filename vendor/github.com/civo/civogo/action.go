package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

// PaginateActionList is a struct for a page of actions
type PaginateActionList struct {
	Page    int      `json:"page"`
	PerPage int      `json:"per_page"`
	Pages   int      `json:"pages"`
	Items   []Action `json:"items"`
}

// Action is a struct for an individual action within the database and when serialized
type Action struct {
	ID          int       `json:"id" gorm:"autoIncrement"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	AccountID   string    `json:"account_id"`
	UserID      string    `json:"user_id"`
	Type        string    `json:"type"`
	Details     string    `json:"details,omitempty"`
	RelatedID   string    `json:"related_id,omitempty"`
	RelatedType string    `json:"related_type,omitempty"`
	Debug       bool      `json:"debug"`
}

// ActionListRequest is a struct for the request to list actions
type ActionListRequest struct {
	PerPage      int    `json:"per_page,omitempty" url:"per_page,omitempty"`
	Page         int    `json:"page,omitempty" url:"page,omitempty"`
	IncludeDebug bool   `json:"include_debug,omitempty" url:"include_debug,omitempty"`
	ResourceID   string `json:"resource_id,omitempty" url:"resource_id,omitempty"`
	Details      string `json:"details,omitempty" url:"details,omitempty"`
	RelatedID    string `json:"related_id,omitempty" url:"related_id,omitempty"`
	ResourceType string `json:"resource_type,omitempty" url:"resource_type,omitempty"`
	ActionType   string `json:"action_type,omitempty" url:"action_type,omitempty"`
	CreatedAt    string `json:"created_at,omitempty" url:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty" url:"updated_at,omitempty"`
	UserID       string `json:"user_id,omitempty" url:"user_id,omitempty"`
}

// ListActions returns a page of actions
func (c *Client) ListActions(listRequest *ActionListRequest) (*PaginateActionList, error) {
	url := "/v2/actions"

	vals, err := query.Values(listRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.SendGetRequest(fmt.Sprintf("%s?%s", url, vals.Encode()))
	if err != nil {
		return nil, decodeError(err)
	}

	paginateActionList := PaginateActionList{}
	err = json.NewDecoder(bytes.NewReader(resp)).Decode(&paginateActionList)
	return &paginateActionList, err
}
