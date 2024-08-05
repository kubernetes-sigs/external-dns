package civogo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// SSHKey represents an SSH public key, uploaded to access instances
type SSHKey struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Fingerprint string    `json:"fingerprint"`
	PublicKey   string    `json:"public_key"`
	CreatedAt   time.Time `json:"created_at"`
}

// ListSSHKeys list all SSH key for an account
func (c *Client) ListSSHKeys() ([]SSHKey, error) {
	resp, err := c.SendGetRequest("/v2/sshkeys")
	if err != nil {
		return nil, err
	}

	sshKeys := make([]SSHKey, 0)
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(&sshKeys); err != nil {
		return nil, decodeError(err)
	}

	return sshKeys, nil
}

// NewSSHKey creates a new SSH key record
func (c *Client) NewSSHKey(name string, publicKey string) (*SimpleResponse, error) {
	resp, err := c.SendPostRequest("/v2/sshkeys", map[string]string{
		"name":       name,
		"public_key": publicKey,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}

// UpdateSSHKey update a SSH key record
func (c *Client) UpdateSSHKey(name string, sshKeyID string) (*SSHKey, error) {
	resp, err := c.SendPutRequest(fmt.Sprintf("/v2/sshkeys/%s", sshKeyID), map[string]string{
		"name": name,
	})
	if err != nil {
		return nil, decodeError(err)
	}

	result := &SSHKey{}
	if err := json.NewDecoder(bytes.NewReader(resp)).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}

// FindSSHKey finds an SSH key by either part of the ID or part of the name
func (c *Client) FindSSHKey(search string) (*SSHKey, error) {
	keys, err := c.ListSSHKeys()
	if err != nil {
		return nil, decodeError(err)
	}

	exactMatch := false
	partialMatchesCount := 0
	result := SSHKey{}

	for _, value := range keys {
		if value.Name == search || value.ID == search {
			exactMatch = true
			result = value
		} else if strings.Contains(value.Name, search) || strings.Contains(value.ID, search) {
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

// DeleteSSHKey deletes an SSH key
func (c *Client) DeleteSSHKey(id string) (*SimpleResponse, error) {
	resp, err := c.SendDeleteRequest(fmt.Sprintf("/v2/sshkeys/%s", id))
	if err != nil {
		return nil, decodeError(err)
	}

	return c.DecodeSimpleResponse(resp)
}
