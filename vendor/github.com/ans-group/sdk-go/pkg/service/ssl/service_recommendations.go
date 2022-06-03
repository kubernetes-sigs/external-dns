package ssl

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetRecommendations retrieves SSL recommendations for a domain
func (s *Service) GetRecommendations(domainName string) (Recommendations, error) {
	body, err := s.getRecommendationsResponseBody(domainName)

	return body.Data, err
}

func (s *Service) getRecommendationsResponseBody(domainName string) (*connection.APIResponseBodyData[Recommendations], error) {
	body := &connection.APIResponseBodyData[Recommendations]{}

	if domainName == "" {
		return body, fmt.Errorf("invalid domain name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ssl/v1/recommendations/%s", domainName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
