package account

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetContacts retrieves a list of contacts
func (s *Service) GetContacts(parameters connection.APIRequestParameters) ([]Contact, error) {
	var contacts []Contact

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetContactsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, contact := range response.(*PaginatedContact).Items {
			contacts = append(contacts, contact)
		}
	}

	return contacts, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetContactsPaginated retrieves a paginated list of contacts
func (s *Service) GetContactsPaginated(parameters connection.APIRequestParameters) (*PaginatedContact, error) {
	body, err := s.getContactsPaginatedResponseBody(parameters)

	return NewPaginatedContact(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetContactsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getContactsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetContactSliceResponseBody, error) {
	body := &GetContactSliceResponseBody{}

	response, err := s.connection.Get("/account/v1/contacts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetContact retrieves a single contact by id
func (s *Service) GetContact(contactID int) (Contact, error) {
	body, err := s.getContactResponseBody(contactID)

	return body.Data, err
}

func (s *Service) getContactResponseBody(contactID int) (*GetContactResponseBody, error) {
	body := &GetContactResponseBody{}

	if contactID < 1 {
		return body, fmt.Errorf("invalid contact id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/account/v1/contacts/%d", contactID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ContactNotFoundError{ID: contactID}
		}

		return nil
	})
}
