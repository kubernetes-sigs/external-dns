package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	return connection.InvokeRequestAll(s.GetSolutionsPaginated, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Solution], error) {
	body, err := s.getSolutionsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetSolutionsPaginated), err
}

func (s *Service) getSolutionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Solution], error) {
	body := &connection.APIResponseBodyData[[]Solution]{}

	response, err := s.connection.Get("/ecloud/v1/solutions", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetSolution retrieves a single Solution by ID
func (s *Service) GetSolution(solutionID int) (Solution, error) {
	body, err := s.getSolutionResponseBody(solutionID)

	return body.Data, err
}

func (s *Service) getSolutionResponseBody(solutionID int) (*connection.APIResponseBodyData[Solution], error) {
	body := &connection.APIResponseBodyData[Solution]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d", solutionID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// PatchSolution patches an eCloud solution
func (s *Service) PatchSolution(solutionID int, patch PatchSolutionRequest) (int, error) {
	body, err := s.patchSolutionResponseBody(solutionID, patch)

	return body.Data.ID, err
}

func (s *Service) patchSolutionResponseBody(solutionID int, patch PatchSolutionRequest) (*connection.APIResponseBodyData[Solution], error) {
	body := &connection.APIResponseBodyData[Solution]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v1/solutions/%d", solutionID), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionVirtualMachines retrieves a list of vms
func (s *Service) GetSolutionVirtualMachines(solutionID int, parameters connection.APIRequestParameters) ([]VirtualMachine, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[VirtualMachine], error) {
		return s.GetSolutionVirtualMachinesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionVirtualMachinesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionVirtualMachinesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[VirtualMachine], error) {
	body, err := s.getSolutionVirtualMachinesPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[VirtualMachine], error) {
		return s.GetSolutionVirtualMachinesPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionVirtualMachinesPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VirtualMachine], error) {
	body := &connection.APIResponseBodyData[[]VirtualMachine]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/vms", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionSites retrieves a list of sites
func (s *Service) GetSolutionSites(solutionID int, parameters connection.APIRequestParameters) ([]Site, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Site], error) {
		return s.GetSolutionSitesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionSitesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionSitesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Site], error) {
	body, err := s.getSolutionSitesPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Site], error) {
		return s.GetSolutionSitesPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionSitesPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Site], error) {
	body := &connection.APIResponseBodyData[[]Site]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/sites", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionDatastores retrieves a list of datastores
func (s *Service) GetSolutionDatastores(solutionID int, parameters connection.APIRequestParameters) ([]Datastore, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
		return s.GetSolutionDatastoresPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionDatastoresPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionDatastoresPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
	body, err := s.getSolutionDatastoresPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Datastore], error) {
		return s.GetSolutionDatastoresPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionDatastoresPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Datastore], error) {
	body := &connection.APIResponseBodyData[[]Datastore]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/datastores", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionHosts retrieves a list of hosts
func (s *Service) GetSolutionHosts(solutionID int, parameters connection.APIRequestParameters) ([]V1Host, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
		return s.GetSolutionHostsPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionHostsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionHostsPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
	body, err := s.getSolutionHostsPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[V1Host], error) {
		return s.GetSolutionHostsPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionHostsPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]V1Host], error) {
	body := &connection.APIResponseBodyData[[]V1Host]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/hosts", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionNetworks retrieves a list of networks
func (s *Service) GetSolutionNetworks(solutionID int, parameters connection.APIRequestParameters) ([]V1Network, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[V1Network], error) {
		return s.GetSolutionNetworksPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionNetworksPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionNetworksPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[V1Network], error) {
	body, err := s.getSolutionNetworksPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[V1Network], error) {
		return s.GetSolutionNetworksPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionNetworksPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]V1Network], error) {
	body := &connection.APIResponseBodyData[[]V1Network]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/networks", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionFirewalls retrieves a list of firewalls
func (s *Service) GetSolutionFirewalls(solutionID int, parameters connection.APIRequestParameters) ([]Firewall, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
		return s.GetSolutionFirewallsPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionFirewallsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionFirewallsPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
	body, err := s.getSolutionFirewallsPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Firewall], error) {
		return s.GetSolutionFirewallsPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionFirewallsPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Firewall], error) {
	body := &connection.APIResponseBodyData[[]Firewall]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/firewalls", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionTemplates retrieves a list of templates
func (s *Service) GetSolutionTemplates(solutionID int, parameters connection.APIRequestParameters) ([]Template, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetSolutionTemplatesPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionTemplatesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionTemplatesPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Template], error) {
	body, err := s.getSolutionTemplatesPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Template], error) {
		return s.GetSolutionTemplatesPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionTemplatesPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Template], error) {
	body := &connection.APIResponseBodyData[[]Template]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/templates", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionTemplate retrieves a single solution template by name
func (s *Service) GetSolutionTemplate(solutionID int, templateName string) (Template, error) {
	body, err := s.getSolutionTemplateResponseBody(solutionID, templateName)

	return body.Data, err
}

func (s *Service) getSolutionTemplateResponseBody(solutionID int, templateName string) (*connection.APIResponseBodyData[Template], error) {
	body := &connection.APIResponseBodyData[Template]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}
	if templateName == "" {
		return body, fmt.Errorf("invalid template name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/templates/%s", solutionID, templateName), connection.APIRequestParameters{})
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

// RenameSolutionTemplate renames a solution template
func (s *Service) RenameSolutionTemplate(solutionID int, templateName string, req RenameTemplateRequest) error {
	_, err := s.renameSolutionTemplateResponseBody(solutionID, templateName, req)

	return err
}

func (s *Service) renameSolutionTemplateResponseBody(solutionID int, templateName string, req RenameTemplateRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}
	if templateName == "" {
		return body, fmt.Errorf("invalid template name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v1/solutions/%d/templates/%s/move", solutionID, templateName), &req)
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

// DeleteSolutionTemplate removes a solution template
func (s *Service) DeleteSolutionTemplate(solutionID int, templateName string) error {
	_, err := s.deleteSolutionTemplateResponseBody(solutionID, templateName)

	return err
}

func (s *Service) deleteSolutionTemplateResponseBody(solutionID int, templateName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}
	if templateName == "" {
		return body, fmt.Errorf("invalid template name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v1/solutions/%d/templates/%s", solutionID, templateName), nil)
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

// GetSolutionTags retrieves a list of tags
func (s *Service) GetSolutionTags(solutionID int, parameters connection.APIRequestParameters) ([]Tag, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Tag], error) {
		return s.GetSolutionTagsPaginated(solutionID, p)
	}, parameters)
}

// GetSolutionTagsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionTagsPaginated(solutionID int, parameters connection.APIRequestParameters) (*connection.Paginated[Tag], error) {
	body, err := s.getSolutionTagsPaginatedResponseBody(solutionID, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Tag], error) {
		return s.GetSolutionTagsPaginated(solutionID, p)
	}), err
}

func (s *Service) getSolutionTagsPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Tag], error) {
	body := &connection.APIResponseBodyData[[]Tag]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/tags", solutionID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// GetSolutionTag retrieves a single solution tag by key
func (s *Service) GetSolutionTag(solutionID int, tagKey string) (Tag, error) {
	body, err := s.getSolutionTagResponseBody(solutionID, tagKey)

	return body.Data, err
}

func (s *Service) getSolutionTagResponseBody(solutionID int, tagKey string) (*connection.APIResponseBodyData[Tag], error) {
	body := &connection.APIResponseBodyData[Tag]{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}
	if tagKey == "" {
		return body, fmt.Errorf("invalid tag key")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v1/solutions/%d/tags/%s", solutionID, tagKey), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{Key: tagKey}
		}

		return nil
	})
}

// CreateSolutionTag creates a new solution tag
func (s *Service) CreateSolutionTag(solutionID int, req CreateTagRequest) error {
	_, err := s.createSolutionTagResponseBody(solutionID, req)

	return err
}

func (s *Service) createSolutionTagResponseBody(solutionID int, req CreateTagRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/ecloud/v1/solutions/%d/tags", solutionID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SolutionNotFoundError{ID: solutionID}
		}

		return nil
	})
}

// PatchSolutionTag patches an eCloud solution tag
func (s *Service) PatchSolutionTag(solutionID int, tagKey string, patch PatchTagRequest) error {
	_, err := s.patchSolutionTagResponseBody(solutionID, tagKey, patch)

	return err
}

func (s *Service) patchSolutionTagResponseBody(solutionID int, tagKey string, patch PatchTagRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}
	if tagKey == "" {
		return body, fmt.Errorf("invalid tag key")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v1/solutions/%d/tags/%s", solutionID, tagKey), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{Key: tagKey}
		}

		return nil
	})
}

// DeleteSolutionTag removes a solution tag
func (s *Service) DeleteSolutionTag(solutionID int, tagKey string) error {
	_, err := s.deleteSolutionTagResponseBody(solutionID, tagKey)

	return err
}

func (s *Service) deleteSolutionTagResponseBody(solutionID int, tagKey string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if solutionID < 1 {
		return body, fmt.Errorf("invalid solution id")
	}
	if tagKey == "" {
		return body, fmt.Errorf("invalid tag key")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v1/solutions/%d/tags/%s", solutionID, tagKey), nil)
	if err != nil {
		return body, err
	}
	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TagNotFoundError{Key: tagKey}
		}

		return nil
	})
}
