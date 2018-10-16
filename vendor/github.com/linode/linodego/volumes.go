package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty"
)

// VolumeStatus indicates the status of the Volume
type VolumeStatus string

const (
	// VolumeCreating indicates the Volume is being created and is not yet available for use
	VolumeCreating VolumeStatus = "creating"

	// VolumeActive indicates the Volume is online and available for use
	VolumeActive VolumeStatus = "active"

	// VolumeResizing indicates the Volume is in the process of upgrading its current capacity
	VolumeResizing VolumeStatus = "resizing"

	// VolumeContactSupport indicates there is a problem with the Volume. A support ticket must be opened to resolve the issue
	VolumeContactSupport VolumeStatus = "contact_support"
)

// Volume represents a linode volume object
type Volume struct {
	CreatedStr string `json:"created"`
	UpdatedStr string `json:"updated"`

	ID             int
	Label          string
	Status         VolumeStatus
	Region         string
	Size           int
	LinodeID       *int      `json:"linode_id"`
	FilesystemPath string    `json:"filesystem_path"`
	Created        time.Time `json:"-"`
	Updated        time.Time `json:"-"`
}

type VolumeCreateOptions struct {
	Label    string `json:"label,omitempty"`
	Region   string `json:"region,omitempty"`
	LinodeID int    `json:"linode_id,omitempty"`
	ConfigID int    `json:"config_id,omitempty"`
	// The Volume's size, in GiB. Minimum size is 10GiB, maximum size is 10240GiB. A "0" value will result in the default size.
	Size int `json:"size,omitempty"`
}

type VolumeAttachOptions struct {
	LinodeID int `json:"linode_id"`
	ConfigID int `json:"config_id,omitempty"`
}

// VolumesPagedResponse represents a linode API response for listing of volumes
type VolumesPagedResponse struct {
	*PageOptions
	Data []*Volume
}

// endpoint gets the endpoint URL for Volume
func (VolumesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Volumes.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends Volumes when processing paginated Volume responses
func (resp *VolumesPagedResponse) appendData(r *VolumesPagedResponse) {
	(*resp).Data = append(resp.Data, r.Data...)
}

// setResult sets the Resty response type of Volume
func (VolumesPagedResponse) setResult(r *resty.Request) {
	r.SetResult(VolumesPagedResponse{})
}

// ListVolumes lists Volumes
func (c *Client) ListVolumes(ctx context.Context, opts *ListOptions) ([]*Volume, error) {
	response := VolumesPagedResponse{}
	err := c.listHelper(ctx, &response, opts)
	for _, el := range response.Data {
		el.fixDates()
	}
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// fixDates converts JSON timestamps to Go time.Time values
func (v *Volume) fixDates() *Volume {
	if parsed, err := parseDates(v.CreatedStr); err != nil {
		v.Created = *parsed
	}
	if parsed, err := parseDates(v.UpdatedStr); err != nil {
		v.Updated = *parsed
	}
	return v
}

// GetVolume gets the template with the provided ID
func (c *Client) GetVolume(ctx context.Context, id int) (*Volume, error) {
	e, err := c.Volumes.Endpoint()
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&Volume{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*Volume).fixDates(), nil
}

// AttachVolume attaches a volume to a Linode instance
func (c *Client) AttachVolume(ctx context.Context, id int, options *VolumeAttachOptions) (*Volume, error) {
	body := ""
	if bodyData, err := json.Marshal(options); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	e, err := c.Volumes.Endpoint()
	if err != nil {
		return nil, NewError(err)
	}

	e = fmt.Sprintf("%s/%d/attach", e, id)
	resp, err := coupleAPIErrors(c.R(ctx).
		SetResult(&Volume{}).
		SetBody(body).
		Post(e))

	return resp.Result().(*Volume).fixDates(), nil
}

// CreateVolume creates a Linode Volume
func (c *Client) CreateVolume(ctx context.Context, createOpts VolumeCreateOptions) (*Volume, error) {
	body := ""
	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	e, err := c.Volumes.Endpoint()
	if err != nil {
		return nil, NewError(err)
	}

	resp, err := coupleAPIErrors(c.R(ctx).
		SetResult(&Volume{}).
		SetBody(body).
		Post(e))

	if err != nil {
		return nil, err
	}

	return resp.Result().(*Volume).fixDates(), nil
}

// RenameVolume renames the label of a Linode volume
// There is no UpdateVolume because the label is the only alterable field.
func (c *Client) RenameVolume(ctx context.Context, id int, label string) (*Volume, error) {
	body, _ := json.Marshal(map[string]string{"label": label})

	e, err := c.Volumes.Endpoint()
	if err != nil {
		return nil, NewError(err)
	}
	e = fmt.Sprintf("%s/%d", e, id)

	resp, err := coupleAPIErrors(c.R(ctx).
		SetResult(&Volume{}).
		SetBody(body).
		Put(e))

	if err != nil {
		return nil, err
	}

	return resp.Result().(*Volume).fixDates(), nil
}

// CloneVolume clones a Linode volume
func (c *Client) CloneVolume(ctx context.Context, id int, label string) (*Volume, error) {
	body := fmt.Sprintf("{\"label\":\"%s\"}", label)

	e, err := c.Volumes.Endpoint()
	if err != nil {
		return nil, NewError(err)
	}
	e = fmt.Sprintf("%s/%d/clone", e, id)

	resp, err := coupleAPIErrors(c.R(ctx).
		SetResult(&Volume{}).
		SetBody(body).
		Post(e))

	if err != nil {
		return nil, err
	}

	return resp.Result().(*Volume).fixDates(), nil
}

// DetachVolume detaches a Linode volume
func (c *Client) DetachVolume(ctx context.Context, id int) (bool, error) {
	body := ""

	e, err := c.Volumes.Endpoint()
	if err != nil {
		return false, NewError(err)
	}

	e = fmt.Sprintf("%s/%d/detach", e, id)

	resp, err := coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))

	return settleBoolResponseOrError(resp, err)
}

// ResizeVolume resizes an instance to new Linode type
func (c *Client) ResizeVolume(ctx context.Context, id int, size int) (bool, error) {
	body := fmt.Sprintf("{\"size\": %d}", size)

	e, err := c.Volumes.Endpoint()
	if err != nil {
		return false, NewError(err)
	}
	e = fmt.Sprintf("%s/%d/resize", e, id)

	resp, err := coupleAPIErrors(c.R(ctx).
		SetBody(body).
		Post(e))

	return settleBoolResponseOrError(resp, err)
}

// DeleteVolume deletes the Volume with the specified id
func (c *Client) DeleteVolume(ctx context.Context, id int) error {
	e, err := c.Volumes.Endpoint()
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, id)

	if _, err := coupleAPIErrors(c.R(ctx).Delete(e)); err != nil {
		return err
	}

	return nil
}
