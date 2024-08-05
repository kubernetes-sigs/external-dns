package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// APIKey represents an IAM API Key resource.
type APIKey struct {
	Key    *string `req-for:"delete"`
	Name   *string `req-for:"create"`
	RoleID *string `req-for:"create"`
}

func apiKeyFromAPI(r *oapi.IamApiKey) *APIKey {
	return &APIKey{
		Key:    r.Key,
		Name:   r.Name,
		RoleID: r.RoleId,
	}
}

// GetAPIKey returns the IAM API Key.
func (c *Client) GetAPIKey(ctx context.Context, zone, key string) (*APIKey, error) {
	resp, err := c.GetApiKeyWithResponse(apiv2.WithZone(ctx, zone), key)
	if err != nil {
		return nil, err
	}

	return apiKeyFromAPI(resp.JSON200), nil
}

// ListAPIKeys returns the list of existing IAM API Keys.
func (c *Client) ListAPIKeys(ctx context.Context, zone string) ([]*APIKey, error) {
	list := make([]*APIKey, 0)

	resp, err := c.ListApiKeysWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.ApiKeys != nil {
		for i := range *resp.JSON200.ApiKeys {
			list = append(list, apiKeyFromAPI(&(*resp.JSON200.ApiKeys)[i]))
		}
	}

	return list, nil
}

// CreateAPIKey creates a IAM API Key.
func (c *Client) CreateAPIKey(
	ctx context.Context,
	zone string,
	apiKey *APIKey,
) (key *APIKey, secret string, err error) {
	if err = validateOperationParams(apiKey, "create"); err != nil {
		return
	}

	req := oapi.CreateApiKeyJSONRequestBody{
		Name:   *apiKey.Name,
		RoleId: *apiKey.RoleID,
	}

	resp, err := c.CreateApiKeyWithResponse(
		apiv2.WithZone(ctx, zone),
		req,
	)
	if err != nil {
		return
	}

	key = &APIKey{
		Key:    resp.JSON200.Key,
		Name:   resp.JSON200.Name,
		RoleID: resp.JSON200.RoleId,
	}
	secret = *resp.JSON200.Secret

	return
}

// DeleteAPIKey deletes IAM API Key.
func (c *Client) DeleteAPIKey(ctx context.Context, zone string, apiKey *APIKey) error {
	if err := validateOperationParams(apiKey, "delete"); err != nil {
		return err
	}

	resp, err := c.DeleteApiKeyWithResponse(apiv2.WithZone(ctx, zone), *apiKey.Key)
	if err != nil {
		return err
	}

	_, err = oapi.NewPoller().
		WithTimeout(c.timeout).
		WithInterval(c.pollInterval).
		Poll(ctx, oapi.OperationPoller(c, zone, *resp.JSON200.Id))
	if err != nil {
		return err
	}

	return nil
}
