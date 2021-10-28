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
	Client *Client
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
	URL string `json:"url"`
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

	req, err := s.Client.NewRequest(ctx, http.MethodPost, uri, snapshotReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.Client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// CreateFromURL will create a snapshot based on an image iso from a URL you provide
func (s *SnapshotServiceHandler) CreateFromURL(ctx context.Context, snapshotURLReq *SnapshotURLReq) (*Snapshot, error) {
	uri := "/v2/snapshots/create-from-url"

	req, err := s.Client.NewRequest(ctx, http.MethodPost, uri, snapshotURLReq)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.Client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Get a specific snapshot
func (s *SnapshotServiceHandler) Get(ctx context.Context, snapshotID string) (*Snapshot, error) {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.Client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	snapshot := new(snapshotBase)
	if err = s.Client.DoWithContext(ctx, req, snapshot); err != nil {
		return nil, err
	}

	return snapshot.Snapshot, nil
}

// Delete a snapshot.
func (s *SnapshotServiceHandler) Delete(ctx context.Context, snapshotID string) error {
	uri := fmt.Sprintf("/v2/snapshots/%s", snapshotID)

	req, err := s.Client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return s.Client.DoWithContext(ctx, req, nil)
}

// List all available snapshots.
func (s *SnapshotServiceHandler) List(ctx context.Context, options *ListOptions) ([]Snapshot, *Meta, error) {
	uri := "/v2/snapshots"

	req, err := s.Client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, nil, err
	}
	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	snapshots := new(snapshotsBase)
	if err = s.Client.DoWithContext(ctx, req, snapshots); err != nil {
		return nil, nil, err
	}

	return snapshots.Snapshots, snapshots.Meta, nil
}
