//go:generate go run ../../gen/model_response/main.go -package pss -source model.go -destination model_response_generated.go
//go:generate go run ../../gen/model_paginated/main.go -package pss -source model.go -destination model_paginated_generated.go

package pss

import "github.com/ukfast/sdk-go/pkg/connection"

type AuthorType string

func (s AuthorType) String() string {
	return string(s)
}

const (
	AuthorTypeClient  AuthorType = "Client"
	AuthorTypeAuto    AuthorType = "Auto"
	AuthorTypeSupport AuthorType = "Support"
)

var AuthorTypeEnum connection.EnumSlice = []connection.Enum{AuthorTypeClient, AuthorTypeAuto, AuthorTypeSupport}

// ParseAuthorType attempts to parse a AuthorType from string
func ParseAuthorType(s string) (AuthorType, error) {
	e, err := connection.ParseEnum(s, AuthorTypeEnum)
	if err != nil {
		return "", err
	}

	return e.(AuthorType), err
}

type RequestPriority string

func (s RequestPriority) String() string {
	return string(s)
}

const (
	RequestPriorityNormal   RequestPriority = "Normal"
	RequestPriorityHigh     RequestPriority = "High"
	RequestPriorityCritical RequestPriority = "Critical"
)

var RequestPriorityEnum connection.EnumSlice = []connection.Enum{RequestPriorityNormal, RequestPriorityHigh, RequestPriorityCritical}

// ParseRequestPriority attempts to parse a RequestPriority from string
func ParseRequestPriority(s string) (RequestPriority, error) {
	e, err := connection.ParseEnum(s, RequestPriorityEnum)
	if err != nil {
		return "", err
	}

	return e.(RequestPriority), err
}

type RequestStatus string

func (s RequestStatus) String() string {
	return string(s)
}

const (
	RequestStatusCompleted                RequestStatus = "Completed"
	RequestStatusAwaitingCustomerResponse RequestStatus = "Awaiting Customer Response"
	RequestStatusRepliedAndCompleted      RequestStatus = "Replied and Completed"
	RequestStatusSubmitted                RequestStatus = "Submitted"
)

var RequestStatusEnum connection.EnumSlice = []connection.Enum{
	RequestStatusCompleted,
	RequestStatusAwaitingCustomerResponse,
	RequestStatusRepliedAndCompleted,
	RequestStatusSubmitted,
}

// ParseRequestStatus attempts to parse a RequestStatus from string
func ParseRequestStatus(s string) (RequestStatus, error) {
	e, err := connection.ParseEnum(s, RequestStatusEnum)
	if err != nil {
		return "", err
	}

	return e.(RequestStatus), err
}

// Request represents a PSS request
// +genie:model_response
// +genie:model_paginated
type Request struct {
	ID                int                 `json:"id"`
	Author            Author              `json:"author"`
	Type              string              `json:"type"`
	Secure            bool                `json:"secure"`
	Subject           string              `json:"subject"`
	CreatedAt         connection.DateTime `json:"created_at"`
	Priority          RequestPriority     `json:"priority"`
	Archived          bool                `json:"archived"`
	Status            RequestStatus       `json:"status"`
	RequestSMS        bool                `json:"request_sms"`
	Version           int                 `json:"version"`
	CustomerReference string              `json:"customer_reference"`
	Product           Product             `json:"product"`
	LastRepliedAt     connection.DateTime `json:"last_replied_at"`
	CC                []string            `json:"cc"`
	UnreadReplies     int                 `json:"unread_replies"`
	ContactMethod     string              `json:"contact_method"`
}

// Author represents a PSS request author
type Author struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Type AuthorType `json:"type"`
}

// Reply represents a PSS reply
// +genie:model_response
// +genie:model_paginated
type Reply struct {
	ID          string              `json:"id"`
	RequestID   int                 `json:"request_id"`
	Author      Author              `json:"author"`
	Description string              `json:"description"`
	Attachments []Attachment        `json:"attachments"`
	Read        bool                `json:"read"`
	CreatedAt   connection.DateTime `json:"created_at"`
}

// Attachment represents a PSS attachment
type Attachment struct {
	Name string `json:"name"`
}

// Product represents a product to which the request applies to
type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Feedback represents PSS feedback
// +genie:model_response
type Feedback struct {
	ID               int                 `json:"id"`
	ContactID        int                 `json:"contact_id"`
	Score            int                 `json:"score"`
	Comment          string              `json:"comment"`
	SpeedResolved    int                 `json:"speed_resolved"`
	Quality          int                 `json:"quality"`
	NPSScore         int                 `json:"nps_score"`
	ThirdPartConsent bool                `json:"thirdparty_consent"`
	CreatedAt        connection.DateTime `json:"created_at"`
}
