package v2

import (
	"context"

	apiv2 "github.com/exoscale/egoscale/v2/api"
	"github.com/exoscale/egoscale/v2/oapi"
)

// IAMAccessKeyResource represents an API resource accessible to an IAM access key.
type IAMAccessKeyResource struct {
	Domain       string
	ResourceName string
	ResourceType string
}

func iamAccessKeyResourceFromAPI(r *oapi.AccessKeyResource) *IAMAccessKeyResource {
	return &IAMAccessKeyResource{
		Domain:       string(*r.Domain),
		ResourceName: *r.ResourceName,
		ResourceType: string(*r.ResourceType),
	}
}

// IAMAccessKeyOperation represents an API operation supported by an IAM access key.
type IAMAccessKeyOperation struct {
	Name string
	Tags []string
}

func iamAccessKeyOperationFromAPI(o *oapi.AccessKeyOperation) *IAMAccessKeyOperation {
	return &IAMAccessKeyOperation{
		Name: *o.Operation,
		Tags: func() (v []string) {
			if o.Tags != nil {
				v = *o.Tags
			}
			return
		}(),
	}
}

// CreateIAMAccessKeyOpt represents a CreateIAMAccessKey operation option.
type CreateIAMAccessKeyOpt func(*oapi.CreateAccessKeyJSONRequestBody)

// CreateIAMAccessKeyWithOperations sets a restricted list of API operations to the IAM access key.
func CreateIAMAccessKeyWithOperations(v []string) CreateIAMAccessKeyOpt {
	return func(b *oapi.CreateAccessKeyJSONRequestBody) {
		if len(v) > 0 {
			b.Operations = &v
		}
	}
}

// CreateIAMAccessKeyWithResources sets a restricted list of API resources to the IAM access key.
func CreateIAMAccessKeyWithResources(v []IAMAccessKeyResource) CreateIAMAccessKeyOpt {
	return func(b *oapi.CreateAccessKeyJSONRequestBody) {
		if len(v) > 0 {
			b.Resources = func() *[]oapi.AccessKeyResource {
				list := make([]oapi.AccessKeyResource, len(v))
				for i, r := range v {
					r := r
					list[i] = oapi.AccessKeyResource{
						Domain:       (*oapi.AccessKeyResourceDomain)(&r.Domain),
						ResourceName: &r.ResourceName,
						ResourceType: (*oapi.AccessKeyResourceResourceType)(&r.ResourceType),
					}
				}
				return &list
			}()
		}
	}
}

// CreateIAMAccessKeyWithTags sets a restricted list of API operation tags to the IAM access key.
func CreateIAMAccessKeyWithTags(v []string) CreateIAMAccessKeyOpt {
	return func(b *oapi.CreateAccessKeyJSONRequestBody) {
		if len(v) > 0 {
			b.Tags = &v
		}
	}
}

// IAMAccessKey represents an IAM access key.
type IAMAccessKey struct {
	Key        *string `req-for:"delete"`
	Name       *string
	Operations *[]string
	Resources  *[]IAMAccessKeyResource
	Secret     *string
	Tags       *[]string
	Type       *string
	Version    *string
}

func iamAccessKeyFromAPI(k *oapi.AccessKey) *IAMAccessKey {
	return &IAMAccessKey{
		Key:        k.Key,
		Name:       k.Name,
		Operations: k.Operations,
		Resources: func() *[]IAMAccessKeyResource {
			if k.Resources != nil {
				list := make([]IAMAccessKeyResource, len(*k.Resources))
				for i, r := range *k.Resources {
					list[i] = *iamAccessKeyResourceFromAPI(&r)
				}
				return &list
			}
			return nil
		}(),
		Secret:  k.Secret,
		Tags:    k.Tags,
		Type:    (*string)(k.Type),
		Version: (*string)(k.Version),
	}
}

// CreateIAMAccessKey creates a new IAM access key.
func (c *Client) CreateIAMAccessKey(
	ctx context.Context,
	zone string,
	name string,
	opts ...CreateIAMAccessKeyOpt,
) (*IAMAccessKey, error) {
	body := oapi.CreateAccessKeyJSONRequestBody{Name: &name}
	for _, opt := range opts {
		opt(&body)
	}

	res, err := c.CreateAccessKeyWithResponse(apiv2.WithZone(ctx, zone), body)
	if err != nil {
		return nil, err
	}

	// Contrary to other CreateXXX() methods, we don't chain the CreateAccessKeyWithResponse
	// call with GetAccessKeyWithResponse as the IAM access key secret is only returned once
	// in the CreateAccessKeyResponse body, therefore we have to return it directly to the
	// caller.
	return iamAccessKeyFromAPI(res.JSON200), nil
}

// GetIAMAccessKey returns the IAM access key corresponding to the specified key.
func (c *Client) GetIAMAccessKey(ctx context.Context, zone, key string) (*IAMAccessKey, error) {
	resp, err := c.GetAccessKeyWithResponse(apiv2.WithZone(ctx, zone), key)
	if err != nil {
		return nil, err
	}

	return iamAccessKeyFromAPI(resp.JSON200), nil
}

// ListIAMAccessKeyOperations returns the list of all available API operations supported
// by IAM access keys.
func (c *Client) ListIAMAccessKeyOperations(ctx context.Context, zone string) ([]*IAMAccessKeyOperation, error) {
	list := make([]*IAMAccessKeyOperation, 0)

	resp, err := c.ListAccessKeyKnownOperationsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.AccessKeyOperations != nil {
		for i := range *resp.JSON200.AccessKeyOperations {
			list = append(list, iamAccessKeyOperationFromAPI(&(*resp.JSON200.AccessKeyOperations)[i]))
		}
	}

	return list, nil
}

// ListIAMAccessKeys returns the list of existing IAM access keys.
func (c *Client) ListIAMAccessKeys(ctx context.Context, zone string) ([]*IAMAccessKey, error) {
	list := make([]*IAMAccessKey, 0)

	resp, err := c.ListAccessKeysWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.AccessKeys != nil {
		for i := range *resp.JSON200.AccessKeys {
			list = append(list, iamAccessKeyFromAPI(&(*resp.JSON200.AccessKeys)[i]))
		}
	}

	return list, nil
}

// ListMyIAMAccessKeyOperations returns the list of operations the current API access key
// performing the request is restricted to.
func (c *Client) ListMyIAMAccessKeyOperations(ctx context.Context, zone string) ([]*IAMAccessKeyOperation, error) {
	list := make([]*IAMAccessKeyOperation, 0)

	resp, err := c.ListAccessKeyOperationsWithResponse(apiv2.WithZone(ctx, zone))
	if err != nil {
		return nil, err
	}

	if resp.JSON200.AccessKeyOperations != nil {
		for i := range *resp.JSON200.AccessKeyOperations {
			list = append(list, iamAccessKeyOperationFromAPI(&(*resp.JSON200.AccessKeyOperations)[i]))
		}
	}

	return list, nil
}

// RevokeIAMAccessKey revokes an IAM access key.
func (c *Client) RevokeIAMAccessKey(ctx context.Context, zone string, iamAccessKey *IAMAccessKey) error {
	if err := validateOperationParams(iamAccessKey, "delete"); err != nil {
		return err
	}

	resp, err := c.RevokeAccessKeyWithResponse(apiv2.WithZone(ctx, zone), *iamAccessKey.Key)
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
