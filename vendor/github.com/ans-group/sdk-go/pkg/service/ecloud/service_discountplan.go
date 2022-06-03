package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetDiscountPlans retrieves a list of discount plans
func (s *Service) GetDiscountPlans(parameters connection.APIRequestParameters) ([]DiscountPlan, error) {
	return connection.InvokeRequestAll(s.GetDiscountPlansPaginated, parameters)
}

// GetDiscountPlansPaginated retrieves a paginated list of discount plans
func (s *Service) GetDiscountPlansPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[DiscountPlan], error) {
	body, err := s.getDiscountPlansPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetDiscountPlansPaginated), err
}

func (s *Service) getDiscountPlansPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]DiscountPlan], error) {
	body := &connection.APIResponseBodyData[[]DiscountPlan]{}

	response, err := s.connection.Get("/ecloud/v2/discount-plans", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetDiscountPlan retrieves a single discount plan by id
func (s *Service) GetDiscountPlan(discID string) (DiscountPlan, error) {
	body, err := s.getDiscountPlanResponseBody(discID)

	return body.Data, err
}

func (s *Service) getDiscountPlanResponseBody(discID string) (*connection.APIResponseBodyData[DiscountPlan], error) {
	body := &connection.APIResponseBodyData[DiscountPlan]{}

	if discID == "" {
		return body, fmt.Errorf("invalid discount plan id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/discount-plans/%s", discID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DiscountPlanNotFoundError{ID: discID}
		}

		return nil
	})
}

// ApproveDiscountPlan approves a floating IP to a resource
func (s *Service) ApproveDiscountPlan(discID string) error {
	_, err := s.approveDiscountPlanResponseBody(discID)

	return err
}

func (s *Service) approveDiscountPlanResponseBody(discID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if discID == "" {
		return body, fmt.Errorf("invalid floating IP id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/discount-plans/%s/approve", discID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DiscountPlanNotFoundError{ID: discID}
		}

		return nil
	})
}

// RejectDiscountPlan rejects a floating IP from a resource
func (s *Service) RejectDiscountPlan(discID string) error {
	_, err := s.rejectDiscountPlanResponseBody(discID)

	return err
}

func (s *Service) rejectDiscountPlanResponseBody(discID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if discID == "" {
		return body, fmt.Errorf("invalid floating IP id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v2/discount-plans/%s/reject", discID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &DiscountPlanNotFoundError{ID: discID}
		}

		return nil
	})
}
