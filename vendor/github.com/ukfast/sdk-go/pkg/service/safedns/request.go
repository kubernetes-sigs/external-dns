package safedns

import "github.com/ukfast/sdk-go/pkg/connection"

// CreateZoneRequest represents a request to create a SafeDNS zone
type CreateZoneRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateZoneRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchZoneRequest represents a SafeDNS zone patch request
type PatchZoneRequest struct {
	Description string `json:"description,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchZoneRequest) Validate() *connection.ValidationError {
	return nil
}

// PatchRecordRequest represents a SafeDNS record patch request
type PatchRecordRequest struct {
	Name     string     `json:"name,omitempty"`
	Type     string     `json:"type,omitempty"`
	Content  string     `json:"content,omitempty"`
	TTL      *RecordTTL `json:"ttl,omitempty"`
	Priority *int       `json:"priority,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchRecordRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateRecordRequest represents a request to create a SafeDNS record
type CreateRecordRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name       string     `json:"name" validate:"required"`
	TemplateID int        `json:"template_id,omitempty"`
	Type       string     `json:"type" validate:"required"`
	Content    string     `json:"content" validate:"required"`
	TTL        *RecordTTL `json:"ttl,omitempty"`
	Priority   *int       `json:"priority,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateRecordRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// CreateNoteRequest represents a request to create a SafeDNS note
type CreateNoteRequest struct {
	connection.APIRequestBodyDefaultValidator

	ContactID int    `json:"contact_id,omitempty"`
	Notes     string `json:"notes" validate:"required"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateNoteRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}

// PatchTemplateRequest represents a SafeDNS template patch request
type PatchTemplateRequest struct {
	Name    string `json:"name,omitempty"`
	Default *bool  `json:"default,omitempty"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *PatchTemplateRequest) Validate() *connection.ValidationError {
	return nil
}

// CreateTemplateRequest represents a request to create a SafeDNS template
type CreateTemplateRequest struct {
	connection.APIRequestBodyDefaultValidator

	Name    string `json:"name" validate:"required"`
	Default bool   `json:"default"`
}

// Validate returns an error if struct properties are missing/invalid
func (c *CreateTemplateRequest) Validate() *connection.ValidationError {
	return c.APIRequestBodyDefaultValidator.Validate(c)
}
