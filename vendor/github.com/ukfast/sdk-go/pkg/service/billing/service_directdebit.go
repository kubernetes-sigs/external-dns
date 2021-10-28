package billing

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetDirectDebit retrieves direct debit details
func (s *Service) GetDirectDebit() (DirectDebit, error) {
	body, err := s.getDirectDebitResponseBody()

	return body.Data, err
}

func (s *Service) getDirectDebitResponseBody() (*GetDirectDebitResponseBody, error) {
	body := &GetDirectDebitResponseBody{}

	response, err := s.connection.Get(fmt.Sprintf("/billing/v1/direct-debit"), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DirectDebitNotFoundError{}
		}

		return nil
	})
}
