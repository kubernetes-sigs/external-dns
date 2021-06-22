package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// SSHKeyService is the interface to interact with the SSH Key endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/ssh
type SSHKeyService interface {
	Create(ctx context.Context, sshKeyReq *SSHKeyReq) (*SSHKey, error)
	Get(ctx context.Context, sshKeyID string) (*SSHKey, error)
	Update(ctx context.Context, sshKeyID string, sshKeyReq *SSHKeyReq) error
	Delete(ctx context.Context, sshKeyID string) error
	List(ctx context.Context, options *ListOptions) ([]SSHKey, *Meta, error)
}

// SSHKeyServiceHandler handles interaction with the SSH Key methods for the Vultr API
type SSHKeyServiceHandler struct {
	client *Client
}

// SSHKey represents an SSH Key on Vultr
type SSHKey struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	SSHKey      string `json:"ssh_key"`
	DateCreated string `json:"date_created"`
}

// SSHKeyReq is the ssh key struct for create and update calls
type SSHKeyReq struct {
	Name   string `json:"name,omitempty"`
	SSHKey string `json:"ssh_key,omitempty"`
}

type sshKeysBase struct {
	SSHKeys []SSHKey `json:"ssh_keys"`
	Meta    *Meta    `json:"meta"`
}

type sshKeyBase struct {
	SSHKey *SSHKey `json:"ssh_key"`
}

// Create a ssh key
func (s *SSHKeyServiceHandler) Create(ctx context.Context, sshKeyReq *SSHKeyReq) (*SSHKey, error) {
	uri := "/v2/ssh-keys"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, sshKeyReq)
	if err != nil {
		return nil, err
	}

	key := new(sshKeyBase)
	if err = s.client.DoWithContext(ctx, req, key); err != nil {
		return nil, err
	}

	return key.SSHKey, nil
}

// Get a specific ssh key.
func (s *SSHKeyServiceHandler) Get(ctx context.Context, sshKeyID string) (*SSHKey, error) {
	uri := fmt.Sprintf("/v2/ssh-keys/%s", sshKeyID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	sshKey := new(sshKeyBase)
	if err = s.client.DoWithContext(ctx, req, sshKey); err != nil {
		return nil, err
	}

	return sshKey.SSHKey, nil
}

// Update will update the given SSH Key. Empty strings will be ignored.
func (s *SSHKeyServiceHandler) Update(ctx context.Context, sshKeyID string, sshKeyReq *SSHKeyReq) error {
	uri := fmt.Sprintf("/v2/ssh-keys/%s", sshKeyID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, uri, sshKeyReq)
	if err != nil {
		return err
	}

	return s.client.DoWithContext(ctx, req, nil)
}

// Delete a specific ssh-key.
func (s *SSHKeyServiceHandler) Delete(ctx context.Context, sshKeyID string) error {
	uri := fmt.Sprintf("/v2/ssh-keys/%s", sshKeyID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return s.client.DoWithContext(ctx, req, nil)
}

// List all available SSH Keys.
func (s *SSHKeyServiceHandler) List(ctx context.Context, options *ListOptions) ([]SSHKey, *Meta, error) {
	uri := "/v2/ssh-keys"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	sshKeys := new(sshKeysBase)
	if err = s.client.DoWithContext(ctx, req, sshKeys); err != nil {
		return nil, nil, err
	}

	return sshKeys.SSHKeys, sshKeys.Meta, nil
}
