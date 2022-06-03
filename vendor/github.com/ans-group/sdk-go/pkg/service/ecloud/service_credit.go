package ecloud

import (
	"github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/account"
)

// GetCredits retrieves a list of credits
func (s *Service) GetCredits(parameters connection.APIRequestParameters) ([]account.Credit, error) {
	body, err := s.getCreditsResponseBody(parameters)

	return body.Data, err
}

func (s *Service) getCreditsResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]account.Credit], error) {
	body := &connection.APIResponseBodyData[[]account.Credit]{}

	response, err := s.connection.Get("/ecloud/v1/credits", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
