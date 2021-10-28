package account

import "github.com/ukfast/sdk-go/pkg/connection"

// GetDetails retrieves account details
func (s *Service) GetDetails() (Details, error) {
	body, err := s.getDetailsResponseBody()

	return body.Data, err
}

func (s *Service) getDetailsResponseBody() (*GetDetailsResponseBody, error) {
	body := &GetDetailsResponseBody{}

	response, err := s.connection.Get("/account/v1/details", connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
