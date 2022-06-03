package safedns

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetZones retrieves a list of zones
func (s *Service) GetZones(parameters connection.APIRequestParameters) ([]Zone, error) {
	return connection.InvokeRequestAll(s.GetZonesPaginated, parameters)
}

// GetZonesPaginated retrieves a paginated list of zones
func (s *Service) GetZonesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Zone], error) {
	body, err := s.getZonesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetZonesPaginated), err
}

func (s *Service) getZonesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Zone], error) {
	body := &connection.APIResponseBodyData[[]Zone]{}

	response, err := s.connection.Get("/safedns/v1/zones", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetZone retrieves a single zone by name
func (s *Service) GetZone(zoneName string) (Zone, error) {
	body, err := s.getZoneResponseBody(zoneName)

	return body.Data, err
}

func (s *Service) getZoneResponseBody(zoneName string) (*connection.APIResponseBodyData[Zone], error) {
	body := &connection.APIResponseBodyData[Zone]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/zones/%s", zoneName), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ZoneName: zoneName}
		}

		return nil
	})
}

// CreateZone creates a new SafeDNS zone
func (s *Service) CreateZone(req CreateZoneRequest) error {
	_, err := s.createZoneResponseBody(req)

	return err
}

func (s *Service) createZoneResponseBody(req CreateZoneRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	response, err := s.connection.Post("/safedns/v1/zones", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchZone patches a SafeDNS zone
func (s *Service) PatchZone(zoneName string, req PatchZoneRequest) error {
	_, err := s.patchZoneResponseBody(zoneName, req)

	return err
}

func (s *Service) patchZoneResponseBody(zoneName string, req PatchZoneRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/safedns/v1/zones/%s", zoneName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// DeleteZone removes a SafeDNS zone
func (s *Service) DeleteZone(zoneName string) error {
	_, err := s.deleteZoneResponseBody(zoneName)

	return err
}

func (s *Service) deleteZoneResponseBody(zoneName string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/safedns/v1/zones/%s", zoneName), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ZoneName: zoneName}
		}

		return nil
	})
}

// GetZoneRecords retrieves a list of records
func (s *Service) GetZoneRecords(zoneName string, parameters connection.APIRequestParameters) ([]Record, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetZoneRecordsPaginated(zoneName, p)
	}, parameters)
}

// GetZoneRecordsPaginated retrieves a paginated list of zones
func (s *Service) GetZoneRecordsPaginated(zoneName string, parameters connection.APIRequestParameters) (*connection.Paginated[Record], error) {
	body, err := s.getZoneRecordsPaginatedResponseBody(zoneName, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Record], error) {
		return s.GetZoneRecordsPaginated(zoneName, p)
	}), err
}

func (s *Service) getZoneRecordsPaginatedResponseBody(zoneName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Record], error) {
	body := &connection.APIResponseBodyData[[]Record]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/zones/%s/records", zoneName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ZoneName: zoneName}
		}

		return nil
	})
}

// GetZoneRecord retrieves a single zone record by ID
func (s *Service) GetZoneRecord(zoneName string, recordID int) (Record, error) {
	body, err := s.getZoneRecordResponseBody(zoneName, recordID)

	return body.Data, err
}

func (s *Service) getZoneRecordResponseBody(zoneName string, recordID int) (*connection.APIResponseBodyData[Record], error) {
	body := &connection.APIResponseBodyData[Record]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}
	if recordID < 1 {
		return body, fmt.Errorf("invalid record id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, recordID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: recordID}
		}

		return nil
	})
}

// CreateZoneRecord creates a new SafeDNS zone record
func (s *Service) CreateZoneRecord(zoneName string, req CreateRecordRequest) (int, error) {
	body, err := s.createZoneRecordResponseBody(zoneName, req)

	return body.Data.ID, err
}

func (s *Service) createZoneRecordResponseBody(zoneName string, req CreateRecordRequest) (*connection.APIResponseBodyData[Record], error) {
	body := &connection.APIResponseBodyData[Record]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/safedns/v1/zones/%s/records", zoneName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ZoneName: zoneName}
		}

		return nil
	})
}

// UpdateZoneRecord updates a SafeDNS zone record
func (s *Service) UpdateZoneRecord(zoneName string, record Record) (int, error) {
	body, err := s.updateZoneRecordResponseBody(zoneName, record)

	return body.Data.ID, err
}

func (s *Service) updateZoneRecordResponseBody(zoneName string, record Record) (*connection.APIResponseBodyData[Record], error) {
	body := &connection.APIResponseBodyData[Record]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}
	if record.ID < 1 {
		return body, fmt.Errorf("invalid record id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, record.ID), &record)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: record.ID}
		}

		return nil
	})
}

// PatchZoneRecord patches a SafeDNS zone record
func (s *Service) PatchZoneRecord(zoneName string, recordID int, patch PatchRecordRequest) (int, error) {
	body, err := s.patchZoneRecordResponseBody(zoneName, recordID, patch)

	return body.Data.ID, err
}

func (s *Service) patchZoneRecordResponseBody(zoneName string, recordID int, patch PatchRecordRequest) (*connection.APIResponseBodyData[Record], error) {
	body := &connection.APIResponseBodyData[Record]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}
	if recordID < 1 {
		return body, fmt.Errorf("invalid record id")
	}

	response, err := s.connection.Put(fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, recordID), &patch)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: recordID}
		}

		return nil
	})
}

// DeleteZoneRecord removes a SafeDNS zone record
func (s *Service) DeleteZoneRecord(zoneName string, recordID int) error {
	_, err := s.deleteZoneRecordResponseBody(zoneName, recordID)

	return err
}

func (s *Service) deleteZoneRecordResponseBody(zoneName string, recordID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}
	if recordID < 1 {
		return body, fmt.Errorf("invalid record id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/safedns/v1/zones/%s/records/%d", zoneName, recordID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneRecordNotFoundError{ZoneName: zoneName, RecordID: recordID}
		}

		return nil
	})
}

// GetZoneNotes retrieves a list of notes
func (s *Service) GetZoneNotes(zoneName string, parameters connection.APIRequestParameters) ([]Note, error) {
	return connection.InvokeRequestAll(func(p connection.APIRequestParameters) (*connection.Paginated[Note], error) {
		return s.GetZoneNotesPaginated(zoneName, p)
	}, parameters)
}

// GetZoneNotesPaginated retrieves a paginated list of zones
func (s *Service) GetZoneNotesPaginated(zoneName string, parameters connection.APIRequestParameters) (*connection.Paginated[Note], error) {
	body, err := s.getZoneNotesPaginatedResponseBody(zoneName, parameters)

	return connection.NewPaginated(body, parameters, func(p connection.APIRequestParameters) (*connection.Paginated[Note], error) {
		return s.GetZoneNotesPaginated(zoneName, p)
	}), err
}

func (s *Service) getZoneNotesPaginatedResponseBody(zoneName string, parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Note], error) {
	body := &connection.APIResponseBodyData[[]Note]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/zones/%s/notes", zoneName), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ZoneName: zoneName}
		}

		return nil
	})
}

// GetZoneNote retrieves a single zone note by ID
func (s *Service) GetZoneNote(zoneName string, noteID int) (Note, error) {
	body, err := s.getZoneNoteResponseBody(zoneName, noteID)

	return body.Data, err
}

func (s *Service) getZoneNoteResponseBody(zoneName string, noteID int) (*connection.APIResponseBodyData[Note], error) {
	body := &connection.APIResponseBodyData[Note]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}
	if noteID < 1 {
		return body, fmt.Errorf("invalid note id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/safedns/v1/zones/%s/notes/%d", zoneName, noteID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNoteNotFoundError{ZoneName: zoneName, NoteID: noteID}
		}

		return nil
	})
}

// CreateZoneNote creates a new SafeDNS zone note
func (s *Service) CreateZoneNote(zoneName string, req CreateNoteRequest) (int, error) {
	body, err := s.createZoneNote(zoneName, req)

	return body.Data.ID, err
}

func (s *Service) createZoneNote(zoneName string, req CreateNoteRequest) (*connection.APIResponseBodyData[Note], error) {
	body := &connection.APIResponseBodyData[Note]{}

	if zoneName == "" {
		return body, fmt.Errorf("invalid zone name")
	}

	response, err := s.connection.Post(fmt.Sprintf("/safedns/v1/zones/%s/notes", zoneName), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &ZoneNotFoundError{ZoneName: zoneName}
		}

		return nil
	})
}
