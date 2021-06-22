package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

// BackupService is the interface to interact with the backup endpoint on the Vultr API
// Link : https://www.vultr.com/api/#tag/backup
type BackupService interface {
	Get(ctx context.Context, backupID string) (*Backup, error)
	List(ctx context.Context, options *ListOptions) ([]Backup, *Meta, error)
}

// BackupServiceHandler handles interaction with the backup methods for the Vultr API
type BackupServiceHandler struct {
	client *Client
}

// Backup represents a Vultr backup
type Backup struct {
	ID          string `json:"id"`
	DateCreated string `json:"date_created"`
	Description string `json:"description"`
	Size        int    `json:"size"`
	Status      string `json:"status"`
}

type backupsBase struct {
	Backups []Backup `json:"backups"`
	Meta    *Meta    `json:"meta"`
}

type backupBase struct {
	Backup *Backup `json:"backup"`
}

// Get retrieves a backup that matches the given backupID
func (b *BackupServiceHandler) Get(ctx context.Context, backupID string) (*Backup, error) {
	uri := fmt.Sprintf("/v2/backups/%s", backupID)
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, err
	}

	backup := new(backupBase)
	if err := b.client.DoWithContext(ctx, req, backup); err != nil {
		return nil, err
	}

	return backup.Backup, nil
}

// List retrieves a list of all backups on the current account
func (b *BackupServiceHandler) List(ctx context.Context, options *ListOptions) ([]Backup, *Meta, error) {
	uri := "/v2/backups"
	req, err := b.client.NewRequest(ctx, http.MethodGet, uri, nil)

	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}

	req.URL.RawQuery = newValues.Encode()

	backups := new(backupsBase)
	if err = b.client.DoWithContext(ctx, req, backups); err != nil {
		return nil, nil, err
	}

	return backups.Backups, backups.Meta, nil
}
