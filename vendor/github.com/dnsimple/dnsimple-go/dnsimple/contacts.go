package dnsimple

import (
	"context"
	"fmt"
)

// ContactsService handles communication with the contact related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/contacts/
type ContactsService struct {
	client *Client
}

// Contact represents a Contact in DNSimple.
type Contact struct {
	ID            int64  `json:"id,omitempty"`
	AccountID     int64  `json:"account_id,omitempty"`
	Label         string `json:"label,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	JobTitle      string `json:"job_title,omitempty"`
	Organization  string `json:"organization_name,omitempty"`
	Address1      string `json:"address1,omitempty"`
	Address2      string `json:"address2,omitempty"`
	City          string `json:"city,omitempty"`
	StateProvince string `json:"state_province,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
	Country       string `json:"country,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Fax           string `json:"fax,omitempty"`
	Email         string `json:"email,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

func contactPath(accountID string, contactID int64) (path string) {
	path = fmt.Sprintf("/%v/contacts", accountID)
	if contactID != 0 {
		path += fmt.Sprintf("/%v", contactID)
	}
	return
}

// ContactResponse represents a response from an API method that returns a Contact struct.
type ContactResponse struct {
	Response
	Data *Contact `json:"data"`
}

// ContactsResponse represents a response from an API method that returns a collection of Contact struct.
type ContactsResponse struct {
	Response
	Data []Contact `json:"data"`
}

// ListContacts list the contacts for an account.
//
// See https://developer.dnsimple.com/v2/contacts/#list
func (s *ContactsService) ListContacts(ctx context.Context, accountID string, options *ListOptions) (*ContactsResponse, error) {
	path := versioned(contactPath(accountID, 0))
	contactsResponse := &ContactsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, contactsResponse)
	if err != nil {
		return contactsResponse, err
	}

	contactsResponse.HTTPResponse = resp
	return contactsResponse, nil
}

// CreateContact creates a new contact.
//
// See https://developer.dnsimple.com/v2/contacts/#create
func (s *ContactsService) CreateContact(ctx context.Context, accountID string, contactAttributes Contact) (*ContactResponse, error) {
	path := versioned(contactPath(accountID, 0))
	contactResponse := &ContactResponse{}

	resp, err := s.client.post(ctx, path, contactAttributes, contactResponse)
	if err != nil {
		return nil, err
	}

	contactResponse.HTTPResponse = resp
	return contactResponse, nil
}

// GetContact fetches a contact.
//
// See https://developer.dnsimple.com/v2/contacts/#get
func (s *ContactsService) GetContact(ctx context.Context, accountID string, contactID int64) (*ContactResponse, error) {
	path := versioned(contactPath(accountID, contactID))
	contactResponse := &ContactResponse{}

	resp, err := s.client.get(ctx, path, contactResponse)
	if err != nil {
		return nil, err
	}

	contactResponse.HTTPResponse = resp
	return contactResponse, nil
}

// UpdateContact updates a contact.
//
// See https://developer.dnsimple.com/v2/contacts/#update
func (s *ContactsService) UpdateContact(ctx context.Context, accountID string, contactID int64, contactAttributes Contact) (*ContactResponse, error) {
	path := versioned(contactPath(accountID, contactID))
	contactResponse := &ContactResponse{}

	resp, err := s.client.patch(ctx, path, contactAttributes, contactResponse)
	if err != nil {
		return nil, err
	}

	contactResponse.HTTPResponse = resp
	return contactResponse, nil
}

// DeleteContact PERMANENTLY deletes a contact from the account.
//
// See https://developer.dnsimple.com/v2/contacts/#delete
func (s *ContactsService) DeleteContact(ctx context.Context, accountID string, contactID int64) (*ContactResponse, error) {
	path := versioned(contactPath(accountID, contactID))
	contactResponse := &ContactResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	contactResponse.HTTPResponse = resp
	return contactResponse, nil
}
