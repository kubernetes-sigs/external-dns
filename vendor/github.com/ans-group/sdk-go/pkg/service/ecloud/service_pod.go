package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetPods retrieves a list of pods
func (s *Service) GetPods(parameters connection.APIRequestParameters) ([]Pod, error) {
	return connection.InvokeRequestAll(s.GetPodsPaginated, parameters)
}

// GetPodsPaginated retrieves a paginated list of pods
func (s *Service) GetPodsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Pod], error) {
	body, err := s.getPodsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetPodsPaginated), err
}

func (s *Service) getPodsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Pod], error) {
	body := &connection.APIResponseBodyData[[]Pod]{}

	response, err := s.connection.Get("/ecloud/v1/pods", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetPod retrieves a single pod by ID
func (s *Service) GetPod(podID int) (Pod, error) {
	body, err := s.getPodResponseBody(podID)

	return body.Data, err
}

func (s *Service) getPodResponseBody(podID int) (*connection.APIResponseBodyData[Pod], error) {
	body := &connection.APIResponseBodyData[Pod]{}

	if podID < 1 {
		return body, fmt.Errorf("invalid pod id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/pods/%d", podID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &PodNotFoundError{ID: podID}
		}

		return nil
	})
}

// GetPodTemplates retrieves a list of templates
func (s *Service) GetPodTemplates(podID int, parameters connection.APIRequestParameters) ([]Template, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetPodTemplatesPaginated(podID, p)
	}, parameters)
}

// GetPodTemplatesPaginated retrieves a paginated list of domains
func (s *Service) GetPodTemplatesPaginated(podID int, parameters connection.APIRequestParameters) (*connection.Paginated[Template], error) {
	body, err := s.getPodTemplatesPaginatedResponseBody(podID, parameters)
	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetPodTemplatesPaginated(podID, p)
	}), err
}

func (s *Service) getPodTemplatesPaginatedResponseBody(podID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Template], error) {
	body := &connection.APIResponseBodyData[[]Template]{}

	if podID < 1 {
		return body, fmt.Errorf("invalid pod id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/pods/%d/templates", podID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetPodTemplate retrieves a single pod template by name
func (s *Service) GetPodTemplate(podID int, templateName string) (Template, error) {
	body, err := s.getPodTemplateResponseBody(podID, templateName)

	return body.Data, err
}

func (s *Service) getPodTemplateResponseBody(podID int, templateName string) (*connection.APIResponseBodyData[Template], error) {
	body := &connection.APIResponseBodyData[Template]{}

	if podID < 1 {
		return body, fmt.Errorf("invalid pod id")
	}
	if templateName == "" {
		return body, fmt.Errorf("invalid template name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/pods/%d/templates/%s", podID, templateName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{Name: templateName}
		}

		return nil
	})
}

// RenamePodTemplate renames a pod template
func (s *Service) RenamePodTemplate(podID int, templateName string, req RenameTemplateRequest) error {
	_, err := s.renamePodTemplateResponseBody(podID, templateName, req)

	return err
}

func (s *Service) renamePodTemplateResponseBody(podID int, templateName string, req RenameTemplateRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if podID < 1 {
		return body, fmt.Errorf("invalid pod id")
	}
	if templateName == "" {
		return body, fmt.Errorf("invalid template name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v1/pods/%d/templates/%s/move", podID, templateName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{Name: templateName}
		}

		return nil
	})
}

// DeletePodTemplate removes a pod template
func (s *Service) DeletePodTemplate(podID int, templateName string) error {
	_, err := s.deletePodTemplateResponseBody(podID, templateName)

	return err
}

func (s *Service) deletePodTemplateResponseBody(podID int, templateName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if podID < 1 {
		return body, fmt.Errorf("invalid pod id")
	}
	if templateName == "" {
		return body, fmt.Errorf("invalid template name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v1/pods/%d/templates/%s", podID, templateName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{Name: templateName}
		}

		return nil
	})
}

// GetPodAppliances retrieves a list of appliances
func (s *Service) GetPodAppliances(podID int, parameters connection.APIRequestParameters) ([]Appliance, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
		return s.GetPodAppliancesPaginated(podID, p)
	}, parameters)
}

// GetPodAppliancesPaginated retrieves a paginated list of domains
func (s *Service) GetPodAppliancesPaginated(podID int, parameters connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
	body, err := s.getPodAppliancesPaginatedResponseBody(podID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Appliance], error) {
		return s.GetPodAppliancesPaginated(podID, p)
	}), err
}

func (s *Service) getPodAppliancesPaginatedResponseBody(podID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Appliance], error) {
	body := &connection.APIResponseBodyData[[]Appliance]{}

	if podID < 1 {
		return body, fmt.Errorf("invalid pod id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/pods/%d/appliances", podID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PodConsoleAvailable removes a pod template
func (s *Service) PodConsoleAvailable(podID int) (bool, error) {
	resp, err := s.podConsoleAvailableResponse(podID)
	if err != nil {
		return false, err
	}

	if resp.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}

func (s *Service) podConsoleAvailableResponse(podID int) (*connection.APIResponse, error) {
	if podID < 1 {
		return &connection.APIResponse{}, fmt.Errorf("invalid pod id")
	}

	return s.connection.Get(fmt.Sprintf("/ecloud/v1/pods/%d/console-available", podID), connection.APIRequestParameters{})
}
