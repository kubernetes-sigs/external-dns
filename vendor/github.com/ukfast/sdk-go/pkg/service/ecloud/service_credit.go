package ecloud

import (
	"github.com/ukfast/sdk-go/pkg/connection"
	"github.com/ukfast/sdk-go/pkg/service/account"
)

// GetCredits retrieves a list of credits
func (s *Service) GetCredits(parameters connection.APIRequestParameters) ([]account.Credit, error) {
	body, err := s.getCreditsResponseBody(parameters)

	return body.Data, err
}

func (s *Service) getCreditsResponseBody(parameters connection.APIRequestParameters) (*account.GetCreditSliceResponseBody, error) {
	body := &account.GetCreditSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v1/credits", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
