package dnsimple

import (
	"context"
	"fmt"
)

type CreateRegistrantChangeInput struct {
	DomainId           string            `json:"domain_id"`
	ContactId          string            `json:"contact_id"`
	ExtendedAttributes map[string]string `json:"extended_attributes"`
}

type RegistrantChange struct {
	Id        int `json:"id"`
	AccountId int `json:"account_id"`
	ContactId int `json:"contact_id"`
	DomainId  int `json:"domain_id"`
	// One of: "new", "pending", "cancelling", "cancelled", "completed".
	State               string            `json:"state"`
	ExtendedAttributes  map[string]string `json:"extended_attributes"`
	RegistryOwnerChange bool              `json:"registry_owner_change"`
	IrtLockLiftedBy     string            `json:"irt_lock_lifted_by"`
	CreatedAt           string            `json:"created_at"`
	UpdatedAt           string            `json:"updated_at"`
}

type RegistrantChangeResponse struct {
	Response
	Data *RegistrantChange `json:"data"`
}

type RegistrantChangesListResponse struct {
	Response
	Data []RegistrantChange `json:"data"`
}

type RegistrantChangeListOptions struct {
	// Only include results with a state field exactly matching the given string
	State *string `url:"state,omitempty"`
	// Only include results with a domain_id field exactly matching the given string
	DomainId *string `url:"domain_id,omitempty"`
	// Only include results with a contact_id field exactly matching the given string
	ContactId *string `url:"contact_id,omitempty"`

	ListOptions
}

type CheckRegistrantChangeInput struct {
	DomainId  string `json:"domain_id"`
	ContactId string `json:"contact_id"`
}

type ExtendedAttribute struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Required    bool                      `json:"required"`
	Options     []ExtendedAttributeOption `json:"options"`
}

type ExtendedAttributeOption struct {
	Title       string `json:"title"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

type RegistrantChangeCheck struct {
	DomainId            int                 `json:"domain_id"`
	ContactId           int                 `json:"contact_id"`
	ExtendedAttributes  []ExtendedAttribute `json:"extended_attributes"`
	RegistryOwnerChange bool                `json:"registry_owner_change"`
}

type RegistrantChangeCheckResponse struct {
	Response
	Data *RegistrantChangeCheck `json:"data"`
}

type RegistrantChangeDeleteResponse struct {
	Response
}

// ListRegistrantChange lists registrant changes in the account.
//
// See https://developer.dnsimple.com/v2/registrar/#listRegistrantChanges
func (s *RegistrarService) ListRegistrantChange(ctx context.Context, accountID string, options *RegistrantChangeListOptions) (*RegistrantChangesListResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/registrant_changes", accountID))
	changeResponse := &RegistrantChangesListResponse{}

	resp, err := s.client.get(ctx, path, changeResponse)
	if err != nil {
		return nil, err
	}

	changeResponse.HTTPResponse = resp
	return changeResponse, nil
}

// CreateRegistrantChange starts a registrant change.
//
// See https://developer.dnsimple.com/v2/registrar/#createRegistrantChange
func (s *RegistrarService) CreateRegistrantChange(ctx context.Context, accountID string, input *CreateRegistrantChangeInput) (*RegistrantChangeResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/registrant_changes", accountID))
	changeResponse := &RegistrantChangeResponse{}

	resp, err := s.client.post(ctx, path, input, changeResponse)
	if err != nil {
		return nil, err
	}

	changeResponse.HTTPResponse = resp
	return changeResponse, nil
}

// CheckRegistrantChange retrieves the requirements of a registrant change.
//
// See https://developer.dnsimple.com/v2/registrar/#checkRegistrantChange
func (s *RegistrarService) CheckRegistrantChange(ctx context.Context, accountID string, input *CheckRegistrantChangeInput) (*RegistrantChangeCheckResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/registrant_changes/check", accountID))
	checkResponse := &RegistrantChangeCheckResponse{}

	resp, err := s.client.post(ctx, path, input, checkResponse)
	if err != nil {
		return nil, err
	}

	checkResponse.HTTPResponse = resp
	return checkResponse, nil
}

// GetRegistrantChange retrieves the details of an existing registrant change.
//
// See https://developer.dnsimple.com/v2/registrar/#getRegistrantChange
func (s *RegistrarService) GetRegistrantChange(ctx context.Context, accountID string, registrantChange int) (*RegistrantChangeResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/registrant_changes/%v", accountID, registrantChange))
	checkResponse := &RegistrantChangeResponse{}

	resp, err := s.client.get(ctx, path, checkResponse)
	if err != nil {
		return nil, err
	}

	checkResponse.HTTPResponse = resp
	return checkResponse, nil
}

// DeleteRegistrantChange cancels an ongoing registrant change from the account.
//
// See https://developer.dnsimple.com/v2/registrar/#deleteRegistrantChange
func (s *RegistrarService) DeleteRegistrantChange(ctx context.Context, accountID string, registrantChange int) (*RegistrantChangeDeleteResponse, error) {
	path := versioned(fmt.Sprintf("/%v/registrar/registrant_changes/%v", accountID, registrantChange))
	deleteResponse := &RegistrantChangeDeleteResponse{}

	resp, err := s.client.delete(ctx, path, nil, deleteResponse)
	if err != nil {
		return nil, err
	}

	deleteResponse.HTTPResponse = resp
	return deleteResponse, nil
}
