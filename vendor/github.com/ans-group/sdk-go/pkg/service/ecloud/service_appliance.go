package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetAppliances retrieves a list of appliances
func (s *Service) GetAppliances(parameters connection.APIRequestParameters) ([]Appliance, error) {
	return connection.InvokeRequestAll(s.GetAppliancesPaginated, parameters)
}

// GetAppliancesPaginated retrieves a paginated list of appliances
func (s *Service) GetAppliancesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
	body, err := s.getAppliancesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetAppliancesPaginated), err
}

func (s *Service) getAppliancesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Appliance], error) {
	body := &connection.APIResponseBodyData[[]Appliance]{}

	response, err := s.connection.Get("/ecloud/v1/appliances", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetAppliance retrieves a single Appliance by ID
func (s *Service) GetAppliance(applianceID string) (Appliance, error) {
	body, err := s.getApplianceResponseBody(applianceID)

	return body.Data, err
}

func (s *Service) getApplianceResponseBody(applianceID string) (*connection.APIResponseBodyData[Appliance], error) {
	body := &connection.APIResponseBodyData[Appliance]{}

	if applianceID == "" {
		return body, fmt.Errorf("invalid appliance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/appliances/%s", applianceID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplianceNotFoundError{ID: applianceID}
		}

		return nil
	})
}

// GetApplianceParameters retrieves a list of parameters
func (s *Service) GetApplianceParameters(applianceID string, parameters connection.APIRequestParameters) ([]ApplianceParameter, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[ApplianceParameter], error) {
		return s.GetApplianceParametersPaginated(applianceID, p)
	}, parameters)
}

// GetApplianceParametersPaginated retrieves a paginated list of domains
func (s *Service) GetApplianceParametersPaginated(applianceID string, parameters connection.APIRequestParameters) (*connection.Paginated[ApplianceParameter], error) {
	body, err := s.getApplianceParametersPaginatedResponseBody(applianceID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[ApplianceParameter], error) {
		return s.GetApplianceParametersPaginated(applianceID, p)
	}), err
}

func (s *Service) getApplianceParametersPaginatedResponseBody(applianceID string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]ApplianceParameter], error) {
	body := &connection.APIResponseBodyData[[]ApplianceParameter]{}

	if applianceID == "" {
		return body, fmt.Errorf("invalid appliance id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/appliances/%s/parameters", applianceID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ApplianceNotFoundError{ID: applianceID}
		}

		return nil
	})
}
