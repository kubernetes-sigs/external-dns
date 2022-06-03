package cloudflare

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAccounts retrieves a list of accounts
func (s *Service) GetAccounts(parameters connection.APIRequestParameters) ([]Account, error) {
	return connection.InvokeRequestAll(s.GetAccountsPaginated, parameters)
}

// GetAccountsPaginated retrieves a paginated list of accounts
func (s *Service) GetAccountsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Account], error) {
	body, err := s.getAccountsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetAccountsPaginated), err
}

func (s *Service) getAccountsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Account], error) {
	body := &connection.APIResponseBodyData[[]Account]{}

	response, err := s.connection.Get("/cloudflare/v1/accounts", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAccount retrieves a single account by id
func (s *Service) GetAccount(accountID string) (Account, error) {
	body, err := s.getAccountResponseBody(accountID)

	return body.Data, err
}

func (s *Service) getAccountResponseBody(accountID string) (*connection.APIResponseBodyData[Account], error) {
	body := &connection.APIResponseBodyData[Account]{}

	if accountID == "" {
		return body, fmt.Errorf("invalid account id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/cloudflare/v1/accounts/%s", accountID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AccountNotFoundError{ID: accountID}
		}

		return nil
	})
}

// CreateAccount creates a new account
func (s *Service) CreateAccount(req CreateAccountRequest) (string, error) {
	body, err := s.createAccountResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createAccountResponseBody(req CreateAccountRequest) (*connection.APIResponseBodyData[Account], error) {
	body := &connection.APIResponseBodyData[Account]{}

	response, err := s.connection.Post("/cloudflare/v1/accounts", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchAccount updates an account
func (s *Service) PatchAccount(accountID string, req PatchAccountRequest) error {
	_, err := s.patchAccountResponseBody(accountID, req)

	return err
}

func (s *Service) patchAccountResponseBody(accountID string, req PatchAccountRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if accountID == "" {
		return body, fmt.Errorf("invalid account id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/cloudflare/v1/accounts/%s", accountID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// CreateAccount creates a new account member
func (s *Service) CreateAccountMember(accountID string, req CreateAccountMemberRequest) error {
	_, err := s.createAccountMemberResponseBody(accountID, req)

	return err
}

func (s *Service) createAccountMemberResponseBody(accountID string, req CreateAccountMemberRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if accountID == "" {
		return body, fmt.Errorf("invalid account id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/cloudflare/v1/accounts/%s/members", accountID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
