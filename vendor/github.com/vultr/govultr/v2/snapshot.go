package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// SnapshotService is the interface to interact with Snapshot endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/snapshot
type SnapshotService interface {
	Create(ctx context.Context, snapshotReq *SnapshotReq) (*Snapshot, error)
	CreateFromURL(ctx context.Context, snapshotURLReq *SnapshotURLReq) (*Snapshot, error)
	Get(ctx context.Context, snapshotID string) (*Snapshot, error)
	Delete(ctx context.Context, snapshotID string) error
	List(ctx context.Context, options *ListOptions) ([]Snapshot, *Meta, error)
}

// SnapshotServiceHandler handles interaction with the snapshot methods for the Vultr API
type SnapshotServiceHandler struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	client *Client
}

// Snapshot represents a Vultr snapshot
type Snapshot struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	ID             string `json:"id"`
	DateCreated    string `json:"date_created"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	CompressedSize int    `json:"compressed_size"`
	Status         string `json:"status"`
	OsID           int    `json:"os_id"`
	AppID          int    `json:"app_id"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Status      string `json:"status"`
	OsID        int    `json:"os_id"`
	AppID       int    `json:"app_id"`
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Status      string `json:"status"`
	OsID        int    `json:"os_id"`
	AppID       int    `json:"app_id"`
=======
	ID             string `json:"id"`
	DateCreated    string `json:"date_created"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	CompressedSize int    `json:"compressed_size"`
	Status         string `json:"status"`
	OsID           int    `json:"os_id"`
	AppID          int    `json:"app_id"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Status      string `json:"status"`
	OsID        int    `json:"os_id"`
	AppID       int    `json:"app_id"`
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Status      string `json:"status"`
	OsID        int    `json:"os_id"`
	AppID       int    `json:"app_id"`
=======
	ID             string `json:"id"`
	DateCreated    string `json:"date_created"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	CompressedSize int    `json:"compressed_size"`
	Status         string `json:"status"`
	OsID           int    `json:"os_id"`
	AppID          int    `json:"app_id"`
>>>>>>> 6b7ce455e (update vendored files)
}

// SnapshotReq struct is used to create snapshots.
type SnapshotReq struct {
	InstanceID  string `json:"instance_id,omitempty"`
	Description string `json:"description,omitempty"`
}

// SnapshotURLReq struct is used to create snapshots from a URL.
type SnapshotURLReq struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type snapshotsBase struct {
	Snapshots []Snapshot `json:"snapshots"`
	Meta      *Meta      `json:"meta"`
}

type snapshotBase struct {
	Snapshot *Snapshot `json:"snapshot"`
}

// Create makes a snapshot of a provided server
func (s *SnapshotServiceHandler) Create(ctx context.Context, snapshotReq *SnapshotReq) (*Snapshot, error) {
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// CreateFromURL will create a snapshot based on an image iso from a URL you provide
func (s *SnapshotServiceHandler) CreateFromURL(ctx context.Context, snapshotURLReq *SnapshotURLReq) (*Snapshot, error) {
	uri := "/v2/snapshots/create-from-url"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotURLReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Get a specific snapshot
func (s *SnapshotServiceHandler) Get(ctx context.Context, snapshotID string) (*Snapshot, error) {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Delete a snapshot.
func (s *SnapshotServiceHandler) Delete(ctx context.Context, snapshotID string) error {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return s.client.DoWithContext(ctx, req, nil)
}

// List all available snapshots.
func (s *SnapshotServiceHandler) List(ctx context.Context, options *ListOptions) ([]Snapshot, *Meta, error) {
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}
	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	snapshots := new(snapshotsBase)
	if err = s.client.DoWithContext(ctx, req, snapshots); err != nil {
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	Client *Client
||||||| parent of 4d7e5ad26 (update vendored files)
	Client *Client
=======
	client *Client
>>>>>>> 4d7e5ad26 (update vendored files)
}

// Snapshot represents a Vultr snapshot
type Snapshot struct {
	ID             string `json:"id"`
	DateCreated    string `json:"date_created"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	CompressedSize int    `json:"compressed_size"`
	Status         string `json:"status"`
	OsID           int    `json:"os_id"`
	AppID          int    `json:"app_id"`
}

// SnapshotReq struct is used to create snapshots.
type SnapshotReq struct {
	InstanceID  string `json:"instance_id,omitempty"`
	Description string `json:"description,omitempty"`
}

// SnapshotURLReq struct is used to create snapshots from a URL.
type SnapshotURLReq struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type snapshotsBase struct {
	Snapshots []Snapshot `json:"snapshots"`
	Meta      *Meta      `json:"meta"`
}

type snapshotBase struct {
	Snapshot *Snapshot `json:"snapshot"`
}

// Create makes a snapshot of a provided server
func (s *SnapshotServiceHandler) Create(ctx context.Context, snapshotReq *SnapshotReq) (*Snapshot, error) {
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// CreateFromURL will create a snapshot based on an image iso from a URL you provide
func (s *SnapshotServiceHandler) CreateFromURL(ctx context.Context, snapshotURLReq *SnapshotURLReq) (*Snapshot, error) {
	uri := "/v2/snapshots/create-from-url"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotURLReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Get a specific snapshot
func (s *SnapshotServiceHandler) Get(ctx context.Context, snapshotID string) (*Snapshot, error) {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Delete a snapshot.
func (s *SnapshotServiceHandler) Delete(ctx context.Context, snapshotID string) error {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return s.client.DoWithContext(ctx, req, nil)
}

// List all available snapshots.
func (s *SnapshotServiceHandler) List(ctx context.Context, options *ListOptions) ([]Snapshot, *Meta, error) {
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}
	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	snapshots := new(snapshotsBase)
<<<<<<< HEAD
	if err = s.Client.DoWithContext(ctx, req, snapshots); err != nil {
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	if err = s.Client.DoWithContext(ctx, req, snapshots); err != nil {
=======
	if err = s.client.DoWithContext(ctx, req, snapshots); err != nil {
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	Client *Client
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	Client *Client
=======
	client *Client
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

// Snapshot represents a Vultr snapshot
type Snapshot struct {
	ID             string `json:"id"`
	DateCreated    string `json:"date_created"`
	Description    string `json:"description"`
	Size           int    `json:"size"`
	CompressedSize int    `json:"compressed_size"`
	Status         string `json:"status"`
	OsID           int    `json:"os_id"`
	AppID          int    `json:"app_id"`
}

// SnapshotReq struct is used to create snapshots.
type SnapshotReq struct {
	InstanceID  string `json:"instance_id,omitempty"`
	Description string `json:"description,omitempty"`
}

// SnapshotURLReq struct is used to create snapshots from a URL.
type SnapshotURLReq struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

type snapshotsBase struct {
	Snapshots []Snapshot `json:"snapshots"`
	Meta      *Meta      `json:"meta"`
}

type snapshotBase struct {
	Snapshot *Snapshot `json:"snapshot"`
}

// Create makes a snapshot of a provided server
func (s *SnapshotServiceHandler) Create(ctx context.Context, snapshotReq *SnapshotReq) (*Snapshot, error) {
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// CreateFromURL will create a snapshot based on an image iso from a URL you provide
func (s *SnapshotServiceHandler) CreateFromURL(ctx context.Context, snapshotURLReq *SnapshotURLReq) (*Snapshot, error) {
	uri := "/v2/snapshots/create-from-url"

	req, err := s.client.NewRequest(ctx, http.MethodPost, uri, snapshotURLReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Get a specific snapshot
func (s *SnapshotServiceHandler) Get(ctx context.Context, snapshotID string) (*Snapshot, error) {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Delete a snapshot.
func (s *SnapshotServiceHandler) Delete(ctx context.Context, snapshotID string) error {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return s.client.DoWithContext(ctx, req, nil)
}

// List all available snapshots.
func (s *SnapshotServiceHandler) List(ctx context.Context, options *ListOptions) ([]Snapshot, *Meta, error) {
	uri := "/v2/snapshots"

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}
	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	snapshots := new(snapshotsBase)
<<<<<<< HEAD
	if err = s.Client.DoWithContext(ctx, req, snapshots); err != nil {
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if err = s.Client.DoWithContext(ctx, req, snapshots); err != nil {
=======
	if err = s.client.DoWithContext(ctx, req, snapshots); err != nil {
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		return nil, nil, err
	}

	return snapshots.Snapshots, snapshots.Meta, nil
}
