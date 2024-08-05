package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// ObjectStoreCredential holds the credential of an object store
type ObjectStoreCredential struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	AccessKeyID       string `json:"access_key_id"`
	SecretAccessKeyID string `json:"secret_access_key_id"`
	MaxSizeGB         int    `json:"max_size_gb,omitempty"`
	Suspended         bool   `json:"suspended"`
	Status            string `json:"status"`
}

// PaginatedObjectStoreCredentials is a paginated list of Objectstore credentials
type PaginatedObjectStoreCredentials struct {
	Page    int                     `json:"page"`
	PerPage int                     `json:"per_page"`
	Pages   int                     `json:"pages"`
	Items   []ObjectStoreCredential `json:"items"`
}

// CreateObjectStoreCredentialRequest holds the request to create a new object store credential
type CreateObjectStoreCredentialRequest struct {
	Name              string  `json:"name" validate:"required"`
	AccessKeyID       *string `json:"access_key_id"`
	SecretAccessKeyID *string `json:"secret_access_key_id"`
	MaxSizeGB         *int    `json:"max_size_gb,omitempty"`
	Region            string  `json:"region,omitempty"`
}

// UpdateObjectStoreCredentialRequest holds the request to update a specified object store credential's details
type UpdateObjectStoreCredentialRequest struct {
	AccessKeyID       *string `json:"access_key_id"`
	SecretAccessKeyID *string `json:"secret_access_key_id"`
	MaxSizeGB         *int    `json:"max_size_gb,omitempty"`
	Region            string  `json:"region,omitempty"`
}

// ListObjectStoreCredentials returns all object store credentials in that specific region
func (c *Client) ListObjectStoreCredentials() (*PaginatedObjectStoreCredentials, error) {
	resp, err := c.SendGetRequest("/v2/objectstore/credentials")
	if err != nil {
		return nil, decodeError(err)
	}

	creds := &PaginatedObjectStoreCredentials{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&creds); err != nil {
		return nil, err
	}

	return creds, nil
}

// GetObjectStoreCredential finds an objectstore credential by the full ID
func (c *Client) GetObjectStoreCredential(id string) (*ObjectStoreCredential, error) {
	resp, err := c.SendGetRequest(fmt.Sprintf("/v2/objectstore/credentials/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	var oscr = ObjectStoreCredential{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&oscr); err != nil {
		return nil, err
	}

	return &oscr, nil
}

// FindObjectStoreCredential finds an objectstore credential by name or by accesskeyID
func (c *Client) FindObjectStoreCredential(search string) (*ObjectStoreCredential, error) {
	creds, err := c.ListObjectStoreCredentials()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := ObjectStoreCredential{}

	for _, value := range creds.Items {
		if value.AccessKeyID == search || value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.AccessKeyID, search) || strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
			if !exactMatch {
				result = value
				partialMatchesCount++
			}
		}
	}

	if exactMatch || partialMatchesCount == 1 {
		return &result, nil
	} else if partialMatchesCount > 1 {
		err := fmt.Errorf("unable to find %s because there were multiple matches", search)
		return nil, MultipleMatchesError.wrap(err)
	} else {
		err := fmt.Errorf("unable to find %s, zero matches", search)
		return nil, ZeroMatchesError.wrap(err)
	}
}

// NewObjectStoreCredential creates a new objectstore credential
func (c *Client) NewObjectStoreCredential(v *CreateObjectStoreCredentialRequest) (*ObjectStoreCredential, error) {
	body, err := c.SendPostRequest("/v2/objectstore/credentials", v)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &ObjectStoreCredential{}
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// UpdateObjectStoreCredential updates an objectstore credential
func (c *Client) UpdateObjectStoreCredential(id string, v *UpdateObjectStoreCredentialRequest) (*ObjectStoreCredential, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/objectstore/credentials/%s", id), v)
	if err != nil {
		return nil, decodeError(err)
	}

	var result = &ObjectStoreCredential{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// DeleteObjectStoreCredential deletes an objectstore credential
func (c *Client) DeleteObjectStoreCredential(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/objectstore/credentials/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
