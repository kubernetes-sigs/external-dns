package cloudflare

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSubscriptions retrieves a list of subscriptions
func (s *Service) GetSubscriptions(parameters connection.APIRequestParameters) ([]Subscription, error) {
	return connection.InvokeRequestAll(s.GetSubscriptionsPaginated, parameters)
}

// GetSubscriptionsPaginated retrieves a paginated list of subscriptions
func (s *Service) GetSubscriptionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Subscription], error) {
	body, err := s.getSubscriptionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetSubscriptionsPaginated), err
}

func (s *Service) getSubscriptionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Subscription], error) {
	body := &connection.APIResponseBodyData[[]Subscription]{}

	response, err := s.connection.Get("/cloudflare/v1/subscriptions", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
