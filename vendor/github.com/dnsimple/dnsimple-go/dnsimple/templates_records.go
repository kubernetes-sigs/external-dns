package dnsimple

import (
	"context"
	"fmt"
)

// TemplateRecord represents a DNS record for a template in DNSimple.
type TemplateRecord struct {
	ID         int64  `json:"id,omitempty"`
	TemplateID int64  `json:"template_id,omitempty"`
	Name       string `json:"name"`
	Content    string `json:"content,omitempty"`
	TTL        int    `json:"ttl,omitempty"`
	Type       string `json:"type,omitempty"`
	Priority   int    `json:"priority,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

func templateRecordPath(accountID string, templateIdentifier string, templateRecordID int64) string {
	if templateRecordID != 0 {
		return fmt.Sprintf("%v/records/%v", templatePath(accountID, templateIdentifier), templateRecordID)
	}

	return templatePath(accountID, templateIdentifier) + "/records"
}

// TemplateRecordResponse represents a response from an API method that returns a TemplateRecord struct.
type TemplateRecordResponse struct {
	Response
	Data *TemplateRecord `json:"data"`
}

// TemplateRecordsResponse represents a response from an API method that returns a collection of TemplateRecord struct.
type TemplateRecordsResponse struct {
	Response
	Data []TemplateRecord `json:"data"`
}

// ListTemplateRecords list the templates for an account.
//
// See https://developer.dnsimple.com/v2/templates/records/#list
func (s *TemplatesService) ListTemplateRecords(ctx context.Context, accountID string, templateIdentifier string, options *ListOptions) (*TemplateRecordsResponse, error) {
	path := versioned(templateRecordPath(accountID, templateIdentifier, 0))
	templateRecordsResponse := &TemplateRecordsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, templateRecordsResponse)
	if err != nil {
		return templateRecordsResponse, err
	}

	templateRecordsResponse.HTTPResponse = resp
	return templateRecordsResponse, nil
}

// CreateTemplateRecord creates a new template record.
//
// See https://developer.dnsimple.com/v2/templates/records/#create
func (s *TemplatesService) CreateTemplateRecord(ctx context.Context, accountID string, templateIdentifier string, templateRecordAttributes TemplateRecord) (*TemplateRecordResponse, error) {
	path := versioned(templateRecordPath(accountID, templateIdentifier, 0))
	templateRecordResponse := &TemplateRecordResponse{}

	resp, err := s.client.post(ctx, path, templateRecordAttributes, templateRecordResponse)
	if err != nil {
		return nil, err
	}

	templateRecordResponse.HTTPResponse = resp
	return templateRecordResponse, nil
}

// GetTemplateRecord fetches a template record.
//
// See https://developer.dnsimple.com/v2/templates/records/#get
func (s *TemplatesService) GetTemplateRecord(ctx context.Context, accountID string, templateIdentifier string, templateRecordID int64) (*TemplateRecordResponse, error) {
	path := versioned(templateRecordPath(accountID, templateIdentifier, templateRecordID))
	templateRecordResponse := &TemplateRecordResponse{}

	resp, err := s.client.get(ctx, path, templateRecordResponse)
	if err != nil {
		return nil, err
	}

	templateRecordResponse.HTTPResponse = resp
	return templateRecordResponse, nil
}

// DeleteTemplateRecord deletes a template record.
//
// See https://developer.dnsimple.com/v2/templates/records/#delete
func (s *TemplatesService) DeleteTemplateRecord(ctx context.Context, accountID string, templateIdentifier string, templateRecordID int64) (*TemplateRecordResponse, error) {
	path := versioned(templateRecordPath(accountID, templateIdentifier, templateRecordID))
	templateRecordResponse := &TemplateRecordResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	templateRecordResponse.HTTPResponse = resp
	return templateRecordResponse, nil
}
