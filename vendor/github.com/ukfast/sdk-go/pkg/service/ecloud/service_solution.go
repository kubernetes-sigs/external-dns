package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetSolutions retrieves a list of solutions
func (s *Service) GetSolutions(parameters connection.APIRequestParameters) ([]Solution, error) {
	var solutions []Solution

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, solution := range response.(*PaginatedSolution).Items {
			solutions = append(solutions, solution)
		}
	}

	return solutions, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionsPaginated retrieves a paginated list of solutions
func (s *Service) GetSolutionsPaginated(parameters connection.APIRequestParameters) (*PaginatedSolution, error) {
	body, err := s.getSolutionsPaginatedResponseBody(parameters)

	return NewPaginatedSolution(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetSolutionSliceResponseBody, error) {
	body := &GetSolutionSliceResponseBody{}

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

func (s *Service) getSolutionResponseBody(solutionID int) (*GetSolutionResponseBody, error) {
	body := &GetSolutionResponseBody{}

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

func (s *Service) patchSolutionResponseBody(solutionID int, patch PatchSolutionRequest) (*GetSolutionResponseBody, error) {
	body := &GetSolutionResponseBody{}

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
	var vms []VirtualMachine

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionVirtualMachinesPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, vm := range response.(*PaginatedVirtualMachine).Items {
			vms = append(vms, vm)
		}
	}

	return vms, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionVirtualMachinesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionVirtualMachinesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedVirtualMachine, error) {
	body, err := s.getSolutionVirtualMachinesPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedVirtualMachine(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionVirtualMachinesPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionVirtualMachinesPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetVirtualMachineSliceResponseBody, error) {
	body := &GetVirtualMachineSliceResponseBody{}

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
	var sites []Site

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionSitesPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, site := range response.(*PaginatedSite).Items {
			sites = append(sites, site)
		}
	}

	return sites, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionSitesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionSitesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedSite, error) {
	body, err := s.getSolutionSitesPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedSite(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionSitesPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionSitesPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetSiteSliceResponseBody, error) {
	body := &GetSiteSliceResponseBody{}

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
	var datastores []Datastore

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionDatastoresPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, datastore := range response.(*PaginatedDatastore).Items {
			datastores = append(datastores, datastore)
		}
	}

	return datastores, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionDatastoresPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionDatastoresPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedDatastore, error) {
	body, err := s.getSolutionDatastoresPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedDatastore(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionDatastoresPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionDatastoresPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetDatastoreSliceResponseBody, error) {
	body := &GetDatastoreSliceResponseBody{}

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
	var hosts []V1Host

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionHostsPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, host := range response.(*PaginatedV1Host).Items {
			hosts = append(hosts, host)
		}
	}

	return hosts, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionHostsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionHostsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedV1Host, error) {
	body, err := s.getSolutionHostsPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedV1Host(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionHostsPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionHostsPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetV1HostSliceResponseBody, error) {
	body := &GetV1HostSliceResponseBody{}

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
	var networks []V1Network

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionNetworksPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, network := range response.(*PaginatedV1Network).Items {
			networks = append(networks, network)
		}
	}

	return networks, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionNetworksPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionNetworksPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedV1Network, error) {
	body, err := s.getSolutionNetworksPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedV1Network(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionNetworksPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionNetworksPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetV1NetworkSliceResponseBody, error) {
	body := &GetV1NetworkSliceResponseBody{}

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
	var firewalls []Firewall

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionFirewallsPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, firewall := range response.(*PaginatedFirewall).Items {
			firewalls = append(firewalls, firewall)
		}
	}

	return firewalls, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionFirewallsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionFirewallsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedFirewall, error) {
	body, err := s.getSolutionFirewallsPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedFirewall(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionFirewallsPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionFirewallsPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetFirewallSliceResponseBody, error) {
	body := &GetFirewallSliceResponseBody{}

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
	var templates []Template

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionTemplatesPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, template := range response.(*PaginatedTemplate).Items {
			templates = append(templates, template)
		}
	}

	return templates, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionTemplatesPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionTemplatesPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedTemplate, error) {
	body, err := s.getSolutionTemplatesPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedTemplate(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionTemplatesPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionTemplatesPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetTemplateSliceResponseBody, error) {
	body := &GetTemplateSliceResponseBody{}

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

func (s *Service) getSolutionTemplateResponseBody(solutionID int, templateName string) (*GetTemplateResponseBody, error) {
	body := &GetTemplateResponseBody{}

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
	var tags []Tag

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionTagsPaginated(solutionID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, tag := range response.(*PaginatedTag).Items {
			tags = append(tags, tag)
		}
	}

	return tags, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSolutionTagsPaginated retrieves a paginated list of domains
func (s *Service) GetSolutionTagsPaginated(solutionID int, parameters connection.APIRequestParameters) (*PaginatedTag, error) {
	body, err := s.getSolutionTagsPaginatedResponseBody(solutionID, parameters)

	return NewPaginatedTag(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSolutionTagsPaginated(solutionID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSolutionTagsPaginatedResponseBody(solutionID int, parameters connection.APIRequestParameters) (*GetTagSliceResponseBody, error) {
	body := &GetTagSliceResponseBody{}

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

func (s *Service) getSolutionTagResponseBody(solutionID int, tagKey string) (*GetTagResponseBody, error) {
	body := &GetTagResponseBody{}

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
