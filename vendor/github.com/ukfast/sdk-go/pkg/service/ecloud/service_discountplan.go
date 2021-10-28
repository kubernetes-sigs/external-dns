package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetDiscountPlans retrieves a list of discount plans
func (s *Service) GetDiscountPlans(parameters connection.APIRequestParameters) ([]DiscountPlan, error) {
	var discs []DiscountPlan

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDiscountPlansPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, disc := range response.(*PaginatedDiscountPlan).Items {
			discs = append(discs, disc)
		}
	}

	return discs, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetDiscountPlansPaginated retrieves a paginated list of discount plans
func (s *Service) GetDiscountPlansPaginated(parameters connection.APIRequestParameters) (*PaginatedDiscountPlan, error) {
	body, err := s.getDiscountPlansPaginatedResponseBody(parameters)

	return NewPaginatedDiscountPlan(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetDiscountPlansPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getDiscountPlansPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetDiscountPlanSliceResponseBody, error) {
	body := &GetDiscountPlanSliceResponseBody{}

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

func (s *Service) getDiscountPlanResponseBody(discID string) (*GetDiscountPlanResponseBody, error) {
	body := &GetDiscountPlanResponseBody{}

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
