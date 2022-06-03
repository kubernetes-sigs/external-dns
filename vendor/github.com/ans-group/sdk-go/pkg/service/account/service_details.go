package account

import "github.com/ans-group/sdk-go/pkg/connection"

// GetDetails retrieves account details
func (s *Service) GetDetails() (Details, error) {
	body, err := s.getDetailsResponseBody()

	return body.Data, err
}

func (s *Service) getDetailsResponseBody() (*connection.APIResponseBodyData[Details], error) {
	body := &connection.APIResponseBodyData[Details]{}

	response, err := s.connection.Get("/account/v1/details", connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
