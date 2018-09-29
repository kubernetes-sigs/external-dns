package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

// InstanceBackupsResponse response struct for backup snapshot
type InstanceBackupsResponse struct {
	Automatic []*InstanceSnapshot
	Snapshot  *InstanceBackupSnapshotResponse
}

type InstanceBackupSnapshotResponse struct {
	Current    *InstanceSnapshot
	InProgress *InstanceSnapshot `json:"in_progress"`
}

type RestoreInstanceOptions struct {
	LinodeID  int  `json:"linode_id"`
	Overwrite bool `json:"overwrite"`
}

// InstanceSnapshot represents a linode backup snapshot
type InstanceSnapshot struct {
	CreatedStr  string `json:"created"`
	UpdatedStr  string `json:"updated"`
	FinishedStr string `json:"finished"`

	ID       int
	Label    string
	Status   InstanceSnapshotStatus
	Type     string
	Created  *time.Time `json:"-"`
	Updated  *time.Time `json:"-"`
	Finished *time.Time `json:"-"`
	Configs  []string
	Disks    []*InstanceSnapshotDisk
}

type InstanceSnapshotDisk struct {
	Label      string
	Size       int
	Filesystem string
}

type InstanceSnapshotStatus string

var (
	SnapshotPaused              InstanceSnapshotStatus = "paused"
	SnapshotPending             InstanceSnapshotStatus = "pending"
	SnapshotRunning             InstanceSnapshotStatus = "running"
	SnapshotNeedsPostProcessing InstanceSnapshotStatus = "needsPostProcessing"
	SnapshotSuccessful          InstanceSnapshotStatus = "successful"
	SnapshotFailed              InstanceSnapshotStatus = "failed"
	SnapshotUserAborted         InstanceSnapshotStatus = "userAborted"
)

func (l *InstanceSnapshot) fixDates() *InstanceSnapshot {
	l.Created, _ = parseDates(l.CreatedStr)
	l.Updated, _ = parseDates(l.UpdatedStr)
	l.Finished, _ = parseDates(l.FinishedStr)
	return l
}

// GetInstanceSnapshot gets the snapshot with the provided ID
func (c *Client) GetInstanceSnapshot(ctx context.Context, linodeID int, snapshotID int) (*InstanceSnapshot, error) {
	e, err := c.InstanceSnapshots.endpointWithID(linodeID)
	if err != nil {
		return nil, err
	}
	e = fmt.Sprintf("%s/%d", e, snapshotID)
	r, err := coupleAPIErrors(c.R(ctx).SetResult(&InstanceSnapshot{}).Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceSnapshot).fixDates(), nil
}

// CreateInstanceSnapshot Creates or Replaces the snapshot Backup of a Linode. If a previous snapshot exists for this Linode, it will be deleted.
func (c *Client) CreateInstanceSnapshot(ctx context.Context, linodeID int, label string) (*InstanceSnapshot, error) {
	o, err := json.Marshal(map[string]string{"label": label})
	if err != nil {
		return nil, err
	}
	body := string(o)
	e, err := c.InstanceSnapshots.endpointWithID(linodeID)
	if err != nil {
		return nil, err
	}

	r, err := coupleAPIErrors(c.R(ctx).
		SetBody(body).
		SetResult(&InstanceSnapshot{}).
		Post(e))
	return r.Result().(*InstanceSnapshot).fixDates(), nil
}

// GetInstanceBackups gets the Instance's available Backups.
// This is not called ListInstanceBackups because a single object is returned, matching the API response.
func (c *Client) GetInstanceBackups(ctx context.Context, linodeID int) (*InstanceBackupsResponse, error) {
	e, err := c.InstanceSnapshots.endpointWithID(linodeID)
	if err != nil {
		return nil, err
	}
	r, err := coupleAPIErrors(c.R(ctx).
		SetResult(&InstanceBackupsResponse{}).
		Get(e))
	if err != nil {
		return nil, err
	}
	return r.Result().(*InstanceBackupsResponse).fixDates(), nil
}

// EnableInstanceBackups Enables backups for the specified Linode.
func (c *Client) EnableInstanceBackups(ctx context.Context, linodeID int) (bool, error) {
	e, err := c.InstanceSnapshots.endpointWithID(linodeID)
	if err != nil {
		return false, err
	}
	e = fmt.Sprintf("%s/enable", e)

	r, err := coupleAPIErrors(c.R(ctx).Post(e))
	return settleBoolResponseOrError(r, err)
}

// CancelInstanceBackups Cancels backups for the specified Linode.
func (c *Client) CancelInstanceBackups(ctx context.Context, linodeID int) (bool, error) {
	e, err := c.InstanceSnapshots.endpointWithID(linodeID)
	if err != nil {
		return false, err
	}
	e = fmt.Sprintf("%s/cancel", e)

	r, err := coupleAPIErrors(c.R(ctx).Post(e))
	return settleBoolResponseOrError(r, err)
}

// RestoreInstanceBackup Restores a Linode's Backup to the specified Linode.
func (c *Client) RestoreInstanceBackup(ctx context.Context, linodeID int, backupID int, opts RestoreInstanceOptions) (bool, error) {
	o, err := json.Marshal(opts)
	if err != nil {
		return false, NewError(err)
	}
	body := string(o)
	e, err := c.InstanceSnapshots.endpointWithID(linodeID)
	if err != nil {
		return false, err
	}
	e = fmt.Sprintf("%s/%d/restore", e, backupID)

	r, err := coupleAPIErrors(c.R(ctx).SetBody(body).Post(e))
	if err != nil {
		return false, err
	}

	return settleBoolResponseOrError(r, err)

}

func (l *InstanceBackupSnapshotResponse) fixDates() *InstanceBackupSnapshotResponse {
	if l.Current != nil {
		l.Current.fixDates()
	}
	if l.InProgress != nil {
		l.InProgress.fixDates()
	}
	return l
}

func (l *InstanceBackupsResponse) fixDates() *InstanceBackupsResponse {
	for _, el := range l.Automatic {
		el.fixDates()
	}
	if l.Snapshot != nil {
		l.Snapshot.fixDates()
	}
	return l
}
