package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// SSHKey represents an SSH key.
type SSHKey struct {
	Fingerprint *string
	Name        *string `req-for:"delete"`
}

func sshKeyFromAPI(k *oapi.SshKey) *SSHKey {
	return &SSHKey{
		Fingerprint: k.Fingerprint,
		Name:        k.Name,
	}
}

// DeleteSSHKey deletes an SSH key.
func (c *Client) DeleteSSHKey(ctx context.Context, zone string, sshKey *SSHKey) error {
	if err := validateOperationParams(sshKey, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteSshKeyWithResponse(apiv2.WithZone(ctx, zone), *sshKey.Name)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, c.OperationPoller(zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}

// GetSSHKey returns the SSH key corresponding to the specified name.
func (c *Client) GetSSHKey(ctx context.Context, zone, name string) (*SSHKey, error) {
	resp, err := c.GetSshKeyWithResponse(apiv2.WithZone(ctx, zone), name)
	if err != nil {
		return nil, err
	}

	return sshKeyFromAPI(resp.JSON200), nil
}

// ListSSHKeys returns the list of existing SSH keys.
func (c *Client) ListSSHKeys(ctx context.Context, zone string) ([]*SSHKey, error) {
	list := make([]*SSHKey, 0)

	resp, err := c.ListSshKeysWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.SshKeys != nil {
		for i := range *resp.JSON200.SshKeys {
			list = append(list, sshKeyFromAPI(&(*resp.JSON200.SshKeys)[i]))
		}
	}

	return list, nil
}

// RegisterSSHKey registers a new SSH key.
func (c *Client) RegisterSSHKey(ctx context.Context, zone, name, publicKey string) (*SSHKey, error) {
	_, err := c.RegisterSshKeyWithResponse(
		apiv2.WithZone(ctx, zone),
		oapi.RegisterSshKeyJSONRequestBody{
			Name:      name,
			PublicKey: publicKey,
		})
	if err != nil {
		return nil, err
	}

	return c.GetSSHKey(ctx, zone, name)
}
