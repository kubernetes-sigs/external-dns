package linodego

import (
	"context"
	"fmt"
)

// ObjectStorageTransfer is an object matching the response of object-storage/transfer
type ObjectStorageTransfer struct {
	AmmountUsed int `json:"used"`
}

// CancelObjectStorage cancels and removes all object storage from the Account
func (c *Client) CancelObjectStorage(ctx context.Context) error {
	e, err := c.ObjectStorage.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/cancel", e)
	_, err = coupleAPIErrors(req.Post(e))
	if err != nil {
		return err
	}

	return nil
}

// GetObjectStorageTransfer returns the amount of outbound data transferred used by the Account
func (c *Client) GetObjectStorageTransfer(ctx context.Context) (*ObjectStorageTransfer, error) {
	e, err := c.ObjectStorage.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/transfer", e)
	r, err := coupleAPIErrors(req.SetResult(&ObjectStorageTransfer{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*ObjectStorageTransfer), nil
}
