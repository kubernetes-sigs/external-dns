package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/linode/linodego/internal/parseabletime"
)

// InstanceDisk represents an Instance Disk object
type InstanceDisk struct {
	ID         int            `json:"id"`
	Label      string         `json:"label"`
	Status     DiskStatus     `json:"status"`
	Size       int            `json:"size"`
	Filesystem DiskFilesystem `json:"filesystem"`
	Created    *time.Time     `json:"-"`
	Updated    *time.Time     `json:"-"`
}

// DiskFilesystem constants start with Filesystem and include Linode API Filesystems
type DiskFilesystem string

// DiskFilesystem constants represent the filesystems types an Instance Disk may use
const (
	FilesystemRaw    DiskFilesystem = "raw"
	FilesystemSwap   DiskFilesystem = "swap"
	FilesystemExt3   DiskFilesystem = "ext3"
	FilesystemExt4   DiskFilesystem = "ext4"
	FilesystemInitrd DiskFilesystem = "initrd"
)

// DiskStatus constants have the prefix "Disk" and include Linode API Instance Disk Status
type DiskStatus string

// DiskStatus constants represent the status values an Instance Disk may have
const (
	DiskReady    DiskStatus = "ready"
	DiskNotReady DiskStatus = "not ready"
	DiskDeleting DiskStatus = "deleting"
)

// InstanceDisksPagedResponse represents a paginated InstanceDisk API response
type InstanceDisksPagedResponse struct {
	*PageOptions
	Data []InstanceDisk `json:"data"`
}

// InstanceDiskCreateOptions are InstanceDisk settings that can be used at creation
type InstanceDiskCreateOptions struct {
	Label string `json:"label"`
	Size  int    `json:"size"`

	// Image is optional, but requires RootPass if provided
	Image    string `json:"image,omitempty"`
	RootPass string `json:"root_pass,omitempty"`

	Filesystem      string            `json:"filesystem,omitempty"`
	AuthorizedKeys  []string          `json:"authorized_keys,omitempty"`
	AuthorizedUsers []string          `json:"authorized_users,omitempty"`
	StackscriptID   int               `json:"stackscript_id,omitempty"`
	StackscriptData map[string]string `json:"stackscript_data,omitempty"`
}

// InstanceDiskUpdateOptions are InstanceDisk settings that can be used in updates
type InstanceDiskUpdateOptions struct {
	Label string `json:"label"`
}

<<<<<<< HEAD
// endpointWithID gets the endpoint URL for InstanceDisks of a given Instance
func (InstanceDisksPagedResponse) endpointWithID(c *Client, id int) string {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	endpoint, err := c.InstanceDisks.endpointWithParams(id)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceDisks when processing paginated InstanceDisk responses
func (resp *InstanceDisksPagedResponse) appendData(r *InstanceDisksPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceDisks lists InstanceDisks
func (c *Client) ListInstanceDisks(ctx context.Context, linodeID int, opts *ListOptions) ([]InstanceDisk, error) {
	response := InstanceDisksPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *InstanceDisk) UnmarshalJSON(b []byte) error {
	type Mask InstanceDisk

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetInstanceDisk gets the template with the provided ID
func (c *Client) GetInstanceDisk(ctx context.Context, linodeID int, configID int) (*InstanceDisk, error) {
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, configID)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceDisk{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceDisk), nil
}

// CreateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) CreateInstanceDisk(ctx context.Context, linodeID int, createOpts InstanceDiskCreateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// UpdateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) UpdateInstanceDisk(ctx context.Context, linodeID int, diskID int, updateOpts InstanceDiskUpdateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// RenameInstanceDisk renames an InstanceDisk
func (c *Client) RenameInstanceDisk(ctx context.Context, linodeID int, diskID int, label string) (*InstanceDisk, error) {
	return c.UpdateInstanceDisk(ctx, linodeID, diskID, InstanceDiskUpdateOptions{Label: label})
}

// ResizeInstanceDisk resizes the size of the Instance disk
func (c *Client) ResizeInstanceDisk(ctx context.Context, linodeID int, diskID int, size int) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/resize", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"size": size,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// PasswordResetInstanceDisk resets the "root" account password on the Instance disk
func (c *Client) PasswordResetInstanceDisk(ctx context.Context, linodeID int, diskID int, password string) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/password", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"password": password,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// DeleteInstanceDisk deletes a Linode Instance Disk
func (c *Client) DeleteInstanceDisk(ctx context.Context, linodeID int, diskID int) error {
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceDisks.endpointWithID(id)
||||||| parent of 5ce8c7613 (update vendored files)
	endpoint, err := c.InstanceDisks.endpointWithID(id)
=======
	endpoint, err := c.InstanceDisks.endpointWithParams(id)
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceDisks when processing paginated InstanceDisk responses
func (resp *InstanceDisksPagedResponse) appendData(r *InstanceDisksPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceDisks lists InstanceDisks
func (c *Client) ListInstanceDisks(ctx context.Context, linodeID int, opts *ListOptions) ([]InstanceDisk, error) {
	response := InstanceDisksPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *InstanceDisk) UnmarshalJSON(b []byte) error {
	type Mask InstanceDisk

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetInstanceDisk gets the template with the provided ID
func (c *Client) GetInstanceDisk(ctx context.Context, linodeID int, configID int) (*InstanceDisk, error) {
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, configID)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceDisk{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceDisk), nil
}

// CreateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) CreateInstanceDisk(ctx context.Context, linodeID int, createOpts InstanceDiskCreateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// UpdateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) UpdateInstanceDisk(ctx context.Context, linodeID int, diskID int, updateOpts InstanceDiskUpdateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// RenameInstanceDisk renames an InstanceDisk
func (c *Client) RenameInstanceDisk(ctx context.Context, linodeID int, diskID int, label string) (*InstanceDisk, error) {
	return c.UpdateInstanceDisk(ctx, linodeID, diskID, InstanceDiskUpdateOptions{Label: label})
}

// ResizeInstanceDisk resizes the size of the Instance disk
func (c *Client) ResizeInstanceDisk(ctx context.Context, linodeID int, diskID int, size int) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/resize", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"size": size,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// PasswordResetInstanceDisk resets the "root" account password on the Instance disk
func (c *Client) PasswordResetInstanceDisk(ctx context.Context, linodeID int, diskID int, password string) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/password", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"password": password,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// DeleteInstanceDisk deletes a Linode Instance Disk
func (c *Client) DeleteInstanceDisk(ctx context.Context, linodeID int, diskID int) error {
<<<<<<< HEAD
	e, err := c.InstanceDisks.endpointWithID(linodeID)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	e, err := c.InstanceDisks.endpointWithID(linodeID)
=======
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceDisks.endpointWithID(id)
||||||| parent of 6b7ce455e (update vendored files)
	endpoint, err := c.InstanceDisks.endpointWithID(id)
=======
	endpoint, err := c.InstanceDisks.endpointWithParams(id)
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceDisks when processing paginated InstanceDisk responses
func (resp *InstanceDisksPagedResponse) appendData(r *InstanceDisksPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceDisks lists InstanceDisks
func (c *Client) ListInstanceDisks(ctx context.Context, linodeID int, opts *ListOptions) ([]InstanceDisk, error) {
	response := InstanceDisksPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *InstanceDisk) UnmarshalJSON(b []byte) error {
	type Mask InstanceDisk

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetInstanceDisk gets the template with the provided ID
func (c *Client) GetInstanceDisk(ctx context.Context, linodeID int, configID int) (*InstanceDisk, error) {
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, configID)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceDisk{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceDisk), nil
}

// CreateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) CreateInstanceDisk(ctx context.Context, linodeID int, createOpts InstanceDiskCreateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// UpdateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) UpdateInstanceDisk(ctx context.Context, linodeID int, diskID int, updateOpts InstanceDiskUpdateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// RenameInstanceDisk renames an InstanceDisk
func (c *Client) RenameInstanceDisk(ctx context.Context, linodeID int, diskID int, label string) (*InstanceDisk, error) {
	return c.UpdateInstanceDisk(ctx, linodeID, diskID, InstanceDiskUpdateOptions{Label: label})
}

// ResizeInstanceDisk resizes the size of the Instance disk
func (c *Client) ResizeInstanceDisk(ctx context.Context, linodeID int, diskID int, size int) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/resize", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"size": size,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// PasswordResetInstanceDisk resets the "root" account password on the Instance disk
func (c *Client) PasswordResetInstanceDisk(ctx context.Context, linodeID int, diskID int, password string) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/password", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"password": password,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// DeleteInstanceDisk deletes a Linode Instance Disk
func (c *Client) DeleteInstanceDisk(ctx context.Context, linodeID int, diskID int) error {
<<<<<<< HEAD
	e, err := c.InstanceDisks.endpointWithID(linodeID)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	e, err := c.InstanceDisks.endpointWithID(linodeID)
=======
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceDisks.endpointWithID(id)
||||||| parent of 4d7e5ad26 (update vendored files)
	endpoint, err := c.InstanceDisks.endpointWithID(id)
=======
	endpoint, err := c.InstanceDisks.endpointWithParams(id)
>>>>>>> 4d7e5ad26 (update vendored files)
	if err != nil {
		panic(err)
	}
	return endpoint
}

// appendData appends InstanceDisks when processing paginated InstanceDisk responses
func (resp *InstanceDisksPagedResponse) appendData(r *InstanceDisksPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListInstanceDisks lists InstanceDisks
func (c *Client) ListInstanceDisks(ctx context.Context, linodeID int, opts *ListOptions) ([]InstanceDisk, error) {
	response := InstanceDisksPagedResponse{}
	err := c.listHelperWithID(ctx, &response, linodeID, opts)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *InstanceDisk) UnmarshalJSON(b []byte) error {
	type Mask InstanceDisk

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetInstanceDisk gets the template with the provided ID
func (c *Client) GetInstanceDisk(ctx context.Context, linodeID int, configID int) (*InstanceDisk, error) {
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, configID)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceDisk{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceDisk), nil
}

// CreateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) CreateInstanceDisk(ctx context.Context, linodeID int, createOpts InstanceDiskCreateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(createOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// UpdateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) UpdateInstanceDisk(ctx context.Context, linodeID int, diskID int, updateOpts InstanceDiskUpdateOptions) (*InstanceDisk, error) {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return nil, err
	}

	e = fmt.Sprintf("%s/%d", e, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{})

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return nil, NewError(err)
	}

	r, err := coupleAPIErrors(req.
		SetBody(body).
		Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// RenameInstanceDisk renames an InstanceDisk
func (c *Client) RenameInstanceDisk(ctx context.Context, linodeID int, diskID int, label string) (*InstanceDisk, error) {
	return c.UpdateInstanceDisk(ctx, linodeID, diskID, InstanceDiskUpdateOptions{Label: label})
}

// ResizeInstanceDisk resizes the size of the Instance disk
func (c *Client) ResizeInstanceDisk(ctx context.Context, linodeID int, diskID int, size int) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/resize", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"size": size,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// PasswordResetInstanceDisk resets the "root" account password on the Instance disk
func (c *Client) PasswordResetInstanceDisk(ctx context.Context, linodeID int, diskID int, password string) error {
	var body string
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d/password", e, diskID)

	req := c.R(ctx).SetResult(&InstanceDisk{})
	updateOpts := map[string]interface{}{
		"password": password,
	}

	if bodyData, err := json.Marshal(updateOpts); err == nil {
		body = string(bodyData)
	} else {
		return NewError(err)
	}

	_, err = coupleAPIErrors(req.
		SetBody(body).
		Post(e))

	return err
}

// DeleteInstanceDisk deletes a Linode Instance Disk
func (c *Client) DeleteInstanceDisk(ctx context.Context, linodeID int, diskID int) error {
<<<<<<< HEAD
	e, err := c.InstanceDisks.endpointWithID(linodeID)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	e, err := c.InstanceDisks.endpointWithID(linodeID)
=======
	e, err := c.InstanceDisks.endpointWithParams(linodeID)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	endpoint, err := c.InstanceDisks.endpointWithID(id)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// endpointWithID gets the endpoint URL for InstanceDisks of a given Instance
func (InstanceDisksPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.InstanceDisks.endpointWithID(id)
=======
// endpoint gets the endpoint URL for InstanceDisks of a given Instance
func (InstanceDisksPagedResponse) endpoint(ids ...any) string {
	id := ids[0].(int)
	return fmt.Sprintf("linode/instances/%d/disks", id)
}

func (resp *InstanceDisksPagedResponse) castResult(r *resty.Request, e string) (int, int, error) {
	res, err := coupleAPIErrors(r.SetResult(InstanceDisksPagedResponse{}).Get(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err != nil {
		return 0, 0, err
	}
	castedRes := res.Result().(*InstanceDisksPagedResponse)
	resp.Data = append(resp.Data, castedRes.Data...)
	return castedRes.Pages, castedRes.Results, nil
}

// ListInstanceDisks lists InstanceDisks
func (c *Client) ListInstanceDisks(ctx context.Context, linodeID int, opts *ListOptions) ([]InstanceDisk, error) {
	response := InstanceDisksPagedResponse{}
	err := c.listHelper(ctx, &response, opts, linodeID)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (i *InstanceDisk) UnmarshalJSON(b []byte) error {
	type Mask InstanceDisk

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(i),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	i.Created = (*time.Time)(p.Created)
	i.Updated = (*time.Time)(p.Updated)

	return nil
}

// GetInstanceDisk gets the template with the provided ID
func (c *Client) GetInstanceDisk(ctx context.Context, linodeID int, diskID int) (*InstanceDisk, error) {
	e := fmt.Sprintf("linode/instances/%d/disks/%d", linodeID, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{})
	r, err := coupleAPIErrors(req.Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceDisk), nil
}

// CreateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) CreateInstanceDisk(ctx context.Context, linodeID int, opts InstanceDiskCreateOptions) (*InstanceDisk, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf("linode/instances/%d/disks", linodeID)
	req := c.R(ctx).SetResult(&InstanceDisk{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Post(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// UpdateInstanceDisk creates a new InstanceDisk for the given Instance
func (c *Client) UpdateInstanceDisk(ctx context.Context, linodeID int, diskID int, opts InstanceDiskUpdateOptions) (*InstanceDisk, error) {
	body, err := json.Marshal(opts)
	if err != nil {
		return nil, err
	}

	e := fmt.Sprintf("linode/instances/%d/disks/%d", linodeID, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{}).SetBody(string(body))
	r, err := coupleAPIErrors(req.Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*InstanceDisk), nil
}

// RenameInstanceDisk renames an InstanceDisk
func (c *Client) RenameInstanceDisk(ctx context.Context, linodeID int, diskID int, label string) (*InstanceDisk, error) {
	return c.UpdateInstanceDisk(ctx, linodeID, diskID, InstanceDiskUpdateOptions{Label: label})
}

// ResizeInstanceDisk resizes the size of the Instance disk
func (c *Client) ResizeInstanceDisk(ctx context.Context, linodeID int, diskID int, size int) error {
	opts := map[string]any{
		"size": size,
	}

	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	e := fmt.Sprintf("linode/instances/%d/disks/%d/resize", linodeID, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{}).SetBody(body)
	_, err = coupleAPIErrors(req.Post(e))
	return err
}

// PasswordResetInstanceDisk resets the "root" account password on the Instance disk
func (c *Client) PasswordResetInstanceDisk(ctx context.Context, linodeID int, diskID int, password string) error {
	opts := map[string]any{
		"password": password,
	}

	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}

	e := fmt.Sprintf("linode/instances/%d/disks/%d/password", linodeID, diskID)
	req := c.R(ctx).SetResult(&InstanceDisk{}).SetBody(string(body))
	_, err = coupleAPIErrors(req.SetBody(body).Post(e))
	return err
}

// DeleteInstanceDisk deletes a Linode Instance Disk
func (c *Client) DeleteInstanceDisk(ctx context.Context, linodeID int, diskID int) error {
<<<<<<< HEAD
	e, err := c.InstanceDisks.endpointWithID(linodeID)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, diskID)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	e, err := c.InstanceDisks.endpointWithID(linodeID)
	if err != nil {
		return err
	}
	e = fmt.Sprintf("%s/%d", e, diskID)

	_, err = coupleAPIErrors(c.R(ctx).Delete(e))
=======
	e := fmt.Sprintf("linode/instances/%d/disks/%d", linodeID, diskID)
	_, err := coupleAPIErrors(c.R(ctx).Delete(e))
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return err
}
