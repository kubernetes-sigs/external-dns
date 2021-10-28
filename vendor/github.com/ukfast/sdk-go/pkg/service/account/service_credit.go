package account

import "github.com/ukfast/sdk-go/pkg/connection"

// GetCredits retrieves a list of credits
func (s *Service) GetCredits(parameters connection.APIRequestParameters) ([]Credit, error) {
	body, err := s.getCreditsResponseBody(parameters)

	return body.Data, err
}

func (s *Service) getCreditsResponseBody(parameters connection.APIRequestParameters) (*GetCreditSliceResponseBody, error) {
	body := &GetCreditSliceResponseBody{}

	response, err := s.connection.Get("/account/v1/credits", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
