package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetAppliances retrieves a list of appliances
func (s *Service) GetAppliances(parameters connection.APIRequestParameters) ([]Appliance, error) {
	var appliances []Appliance

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAppliancesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, appliance := range response.(*PaginatedAppliance).Items {
			appliances = append(appliances, appliance)
		}
	}

	return appliances, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetAppliancesPaginated retrieves a paginated list of appliances
func (s *Service) GetAppliancesPaginated(parameters connection.APIRequestParameters) (*PaginatedAppliance, error) {
	body, err := s.getAppliancesPaginatedResponseBody(parameters)

	return NewPaginatedAppliance(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetAppliancesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getAppliancesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetApplianceSliceResponseBody, error) {
	body := &GetApplianceSliceResponseBody{}

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

func (s *Service) getApplianceResponseBody(applianceID string) (*GetApplianceResponseBody, error) {
	body := &GetApplianceResponseBody{}

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
	var appParameters []ApplianceParameter

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetApplianceParametersPaginated(applianceID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, parameter := range response.(*PaginatedApplianceParameter).Items {
			appParameters = append(appParameters, parameter)
		}
	}

	return appParameters, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetApplianceParametersPaginated retrieves a paginated list of domains
func (s *Service) GetApplianceParametersPaginated(applianceID string, parameters connection.APIRequestParameters) (*PaginatedApplianceParameter, error) {
	body, err := s.getApplianceParametersPaginatedResponseBody(applianceID, parameters)

	return NewPaginatedApplianceParameter(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetApplianceParametersPaginated(applianceID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getApplianceParametersPaginatedResponseBody(applianceID string, parameters connection.APIRequestParameters) (*GetApplianceParameterSliceResponseBody, error) {
	body := &GetApplianceParameterSliceResponseBody{}

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
