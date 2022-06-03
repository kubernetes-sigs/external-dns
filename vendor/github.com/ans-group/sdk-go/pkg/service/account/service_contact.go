package account

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetContacts retrieves a list of contacts
func (s *Service) GetContacts(parameters connection.APIRequestParameters) ([]Contact, error) {
	return connection.InvokeRequestAll(s.GetContactsPaginated, parameters)
}

// GetContactsPaginated retrieves a paginated list of contacts
func (s *Service) GetContactsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Contact], error) {
	body, err := s.getContactsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetContactsPaginated), err
}

func (s *Service) getContactsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Contact], error) {
	body := &connection.APIResponseBodyData[[]Contact]{}

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

func (s *Service) getContactResponseBody(contactID int) (*connection.APIResponseBodyData[Contact], error) {
	body := &connection.APIResponseBodyData[Contact]{}

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
