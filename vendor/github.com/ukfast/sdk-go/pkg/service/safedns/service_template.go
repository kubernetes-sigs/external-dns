package safedns

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetTemplates retrieves a list of templates
func (s *Service) GetTemplates(parameters connection.APIRequestParameters) ([]Template, error) {
	var templates []Template

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTemplatesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, template := range response.(*PaginatedTemplate).Items {
			templates = append(templates, template)
		}
	}

	return templates, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTemplatesPaginated retrieves a paginated list of templates
func (s *Service) GetTemplatesPaginated(parameters connection.APIRequestParameters) (*PaginatedTemplate, error) {
	body, err := s.getTemplatesPaginatedResponseBody(parameters)

	return NewPaginatedTemplate(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTemplatesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTemplatesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetTemplateSliceResponseBody, error) {
	body := &GetTemplateSliceResponseBody{}

	response, err := s.connection.Get("/safedns/v1/templates", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetTemplate retrieves a single template by ID
func (s *Service) GetTemplate(templateID int) (Template, error) {
	body, err := s.getTemplateResponseBody(templateID)

	return body.Data, err
}

func (s *Service) getTemplateResponseBody(templateID int) (*GetTemplateResponseBody, error) {
	body := &GetTemplateResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/templates/%d", templateID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{TemplateID: templateID}
		}

		return nil
	})
}

// CreateTemplate creates a new SafeDNS template
func (s *Service) CreateTemplate(req CreateTemplateRequest) (int, error) {
	body, err := s.createTemplateResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createTemplateResponseBody(req CreateTemplateRequest) (*GetTemplateResponseBody, error) {
	body := &GetTemplateResponseBody{}

	response, err := s.connection.Post("/safedns/v1/templates", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchTemplate patches a SafeDNS template
func (s *Service) PatchTemplate(templateID int, patch PatchTemplateRequest) (int, error) {
	body, err := s.patchTemplateResponseBody(templateID, patch)

	return body.Data.ID, err
}

func (s *Service) patchTemplateResponseBody(templateID int, patch PatchTemplateRequest) (*GetTemplateResponseBody, error) {
	body := &GetTemplateResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/safedns/v1/templates/%d", templateID), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{TemplateID: templateID}
		}

		return nil
	})
}

// DeleteTemplate removes a SafeDNS template
func (s *Service) DeleteTemplate(templateID int) error {
	_, err := s.deleteTemplateResponseBody(templateID)

	return err
}

func (s *Service) deleteTemplateResponseBody(templateID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/safedns/v1/templates/%d", templateID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{TemplateID: templateID}
		}

		return nil
	})
}

// GetTemplateRecords retrieves a list of records
func (s *Service) GetTemplateRecords(templateID int, parameters connection.APIRequestParameters) ([]Record, error) {
	var records []Record

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTemplateRecordsPaginated(templateID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, record := range response.(*PaginatedRecord).Items {
			records = append(records, record)
		}
	}

	return records, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetTemplateRecordsPaginated retrieves a paginated list of templates
func (s *Service) GetTemplateRecordsPaginated(templateID int, parameters connection.APIRequestParameters) (*PaginatedRecord, error) {
	body, err := s.getTemplateRecordsPaginatedResponseBody(templateID, parameters)

	return NewPaginatedRecord(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetTemplateRecordsPaginated(templateID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getTemplateRecordsPaginatedResponseBody(templateID int, parameters connection.APIRequestParameters) (*GetRecordSliceResponseBody, error) {
	body := &GetRecordSliceResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/templates/%d/records", templateID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{TemplateID: templateID}
		}

		return nil
	})
}

// GetTemplateRecord retrieves a single zone record by ID
func (s *Service) GetTemplateRecord(templateID int, recordID int) (Record, error) {
	body, err := s.getTemplateRecordResponseBody(templateID, recordID)

	return body.Data, err
}

func (s *Service) getTemplateRecordResponseBody(templateID int, recordID int) (*GetRecordResponseBody, error) {
	body := &GetRecordResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}
	if recordID < 1 {
		return body, fmt.Errorf("invalid record id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/templates/%d/records/%d", templateID, recordID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateRecordNotFoundError{TemplateID: templateID, RecordID: recordID}
		}

		return nil
	})
}

// CreateTemplateRecord creates a new SafeDNS zone record
func (s *Service) CreateTemplateRecord(templateID int, req CreateRecordRequest) (int, error) {
	body, err := s.createTemplateRecordResponseBody(templateID, req)

	return body.Data.ID, err
}

func (s *Service) createTemplateRecordResponseBody(templateID int, req CreateRecordRequest) (*GetTemplateResponseBody, error) {
	body := &GetTemplateResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/safedns/v1/templates/%d/records", templateID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateNotFoundError{TemplateID: templateID}
		}

		return nil
	})
}

// PatchTemplateRecord patches a SafeDNS template record
func (s *Service) PatchTemplateRecord(templateID int, recordID int, patch PatchRecordRequest) (int, error) {
	body, err := s.patchTemplateRecordResponseBody(templateID, recordID, patch)

	return body.Data.ID, err
}

func (s *Service) patchTemplateRecordResponseBody(templateID int, recordID int, patch PatchRecordRequest) (*GetTemplateResponseBody, error) {
	body := &GetTemplateResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}
	if recordID < 1 {
		return body, fmt.Errorf("invalid record id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/safedns/v1/templates/%d/records/%d", templateID, recordID), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateRecordNotFoundError{TemplateID: templateID, RecordID: recordID}
		}

		return nil
	})
}

// DeleteTemplateRecord removes a SafeDNS template record
func (s *Service) DeleteTemplateRecord(templateID int, recordID int) error {
	_, err := s.deleteTemplateRecordResponseBody(templateID, recordID)

	return err
}

func (s *Service) deleteTemplateRecordResponseBody(templateID int, recordID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if templateID < 1 {
		return body, fmt.Errorf("invalid template id")
	}
	if recordID < 1 {
		return body, fmt.Errorf("invalid record id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/safedns/v1/templates/%d/records/%d", templateID, recordID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &TemplateRecordNotFoundError{TemplateID: templateID, RecordID: recordID}
		}

		return nil
	})
}
