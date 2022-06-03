package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type PostgresDatabaseTarget string

const (
	PostgresDatabaseTargetPrimary   PostgresDatabaseTarget = "primary"
	PostgresDatabaseTargetSecondary PostgresDatabaseTarget = "secondary"
)

type PostgresCommitType string

const (
	PostgresCommitTrue        PostgresCommitType = "true"
	PostgresCommitFalse       PostgresCommitType = "false"
	PostgresCommitLocal       PostgresCommitType = "local"
	PostgresCommitRemoteWrite PostgresCommitType = "remote_write"
	PostgresCommitRemoteApply PostgresCommitType = "remote_apply"
)

type PostgresReplicationType string

const (
	PostgresReplicationNone      PostgresReplicationType = "none"
	PostgresReplicationAsynch    PostgresReplicationType = "asynch"
	PostgresReplicationSemiSynch PostgresReplicationType = "semi_synch"
)

// A PostgresDatabase is a instance of Linode Postgres Managed Databases
type PostgresDatabase struct {
	ID                    int                       `json:"id"`
	Status                DatabaseStatus            `json:"status"`
	Label                 string                    `json:"label"`
	Region                string                    `json:"region"`
	Type                  string                    `json:"type"`
	Engine                string                    `json:"engine"`
	Version               string                    `json:"version"`
	Encrypted             bool                      `json:"encrypted"`
	AllowList             []string                  `json:"allow_list"`
	Port                  int                       `json:"port"`
	SSLConnection         bool                      `json:"ssl_connection"`
	ClusterSize           int                       `json:"cluster_size"`
	ReplicationCommitType PostgresCommitType        `json:"replication_commit_type"`
	ReplicationType       PostgresReplicationType   `json:"replication_type"`
	Hosts                 DatabaseHost              `json:"hosts"`
	Updates               DatabaseMaintenanceWindow `json:"updates"`
	Created               *time.Time                `json:"-"`
	Updated               *time.Time                `json:"-"`
}

func (d *PostgresDatabase) UnmarshalJSON(b []byte) error {
	type Mask PostgresDatabase

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
		Updated *parseabletime.ParseableTime `json:"updated"`
	}{
		Mask: (*Mask)(d),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	d.Created = (*time.Time)(p.Created)
	d.Updated = (*time.Time)(p.Updated)
	return nil
}

// PostgresCreateOptions fields are used when creating a new Postgres Database
type PostgresCreateOptions struct {
	Label                 string                  `json:"label"`
	Region                string                  `json:"region"`
	Type                  string                  `json:"type"`
	Engine                string                  `json:"engine"`
	AllowList             []string                `json:"allow_list,omitempty"`
	ClusterSize           int                     `json:"cluster_size,omitempty"`
	Encrypted             bool                    `json:"encrypted,omitempty"`
	SSLConnection         bool                    `json:"ssl_connection,omitempty"`
	ReplicationType       PostgresReplicationType `json:"replication_type,omitempty"`
	ReplicationCommitType PostgresCommitType      `json:"replication_commit_type,omitempty"`
}

// PostgresUpdateOptions fields are used when altering the existing Postgres Database
type PostgresUpdateOptions struct {
	Label     string                     `json:"label,omitempty"`
	AllowList *[]string                  `json:"allow_list,omitempty"`
	Updates   *DatabaseMaintenanceWindow `json:"updates,omitempty"`
}

// PostgresDatabaseSSL is the SSL Certificate to access the Linode Managed Postgres Database
type PostgresDatabaseSSL struct {
	CACertificate []byte `json:"ca_certificate"`
}

// PostgresDatabaseCredential is the Root Credentials to access the Linode Managed Database
type PostgresDatabaseCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostgresDatabasesPagedResponse struct {
	*PageOptions
	Data []PostgresDatabase `json:"data"`
}

func (PostgresDatabasesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

func (resp *PostgresDatabasesPagedResponse) appendData(r *PostgresDatabasesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListPostgresDatabases lists all Postgres Databases associated with the account
func (c *Client) ListPostgresDatabases(ctx context.Context, opts *ListOptions) ([]PostgresDatabase, error) {
	response := PostgresDatabasesPagedResponse{}

	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// PostgresDatabaseBackup is information for interacting with a backup for the existing Postgres Database
type PostgresDatabaseBackup struct {
	ID      int        `json:"id"`
	Label   string     `json:"label"`
	Type    string     `json:"type"`
	Created *time.Time `json:"-"`
}

func (d *PostgresDatabaseBackup) UnmarshalJSON(b []byte) error {
	type Mask PostgresDatabaseBackup

	p := struct {
		*Mask
		Created *parseabletime.ParseableTime `json:"created"`
	}{
		Mask: (*Mask)(d),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	d.Created = (*time.Time)(p.Created)
	return nil
}

// PostgresBackupCreateOptions are options used for CreatePostgresDatabaseBackup(...)
type PostgresBackupCreateOptions struct {
	Label  string                 `json:"label"`
	Target PostgresDatabaseTarget `json:"target"`
}

type PostgresDatabaseBackupsPagedResponse struct {
	*PageOptions
	Data []PostgresDatabaseBackup `json:"data"`
}

func (PostgresDatabaseBackupsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%d/backups", endpoint, id)
}

func (resp *PostgresDatabaseBackupsPagedResponse) appendData(r *PostgresDatabaseBackupsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListPostgresDatabaseBackups lists all Postgres Database Backups associated with the given Postgres Database
func (c *Client) ListPostgresDatabaseBackups(ctx context.Context, databaseID int, opts *ListOptions) ([]PostgresDatabaseBackup, error) {
	response := PostgresDatabaseBackupsPagedResponse{}

	err := c.listHelperWithID(ctx, &response, databaseID, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetPostgresDatabase returns a single Postgres Database matching the id
func (c *Client) GetPostgresDatabase(ctx context.Context, id int) (*PostgresDatabase, error) {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(req.SetResult(&PostgresDatabase{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*PostgresDatabase), nil
}

// CreatePostgresDatabase creates a new Postgres Database using the createOpts as configuration, returns the new Postgres Database
func (c *Client) CreatePostgresDatabase(ctx context.Context, createOpts PostgresCreateOptions) (*PostgresDatabase, error) {
	var body string
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&PostgresDatabase{})

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
	return r.Result().(*PostgresDatabase), nil
}

// DeletePostgresDatabase deletes an existing Postgres Database with the given id
func (c *Client) DeletePostgresDatabase(ctx context.Context, id int) error {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d", e, id)
	_, err = coupleAPIErrors(req.Delete(e))
	return err
}

// UpdatePostgresDatabase updates the given Postgres Database with the provided opts, returns the PostgresDatabase with the new settings
func (c *Client) UpdatePostgresDatabase(ctx context.Context, id int, opts PostgresUpdateOptions) (*PostgresDatabase, error) {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return nil, err
	}
	req := c.R(ctx).SetResult(&PostgresDatabase{})

	bodyData, err := json.Marshal(opts)
	if err != nil {
		return nil, NewError(err)
	}

	body := string(bodyData)

	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(req.SetBody(body).Put(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*PostgresDatabase), nil
}

// PatchPostgresDatabase applies security patches and updates to the underlying operating system of the Managed Postgres Database
func (c *Client) PatchPostgresDatabase(ctx context.Context, databaseID int) error {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/patch", e, databaseID)
	_, err = coupleAPIErrors(req.Post(e))
	if err != nil {
		return err
	}

	return nil
}

// GetPostgresDatabaseCredentials returns the Root Credentials for the given Postgres Database
func (c *Client) GetPostgresDatabaseCredentials(ctx context.Context, id int) (*PostgresDatabaseCredential, error) {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/credentials", e, id)
	r, err := coupleAPIErrors(req.SetResult(&PostgresDatabaseCredential{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*PostgresDatabaseCredential), nil
}

// ResetPostgresDatabaseCredentials returns the Root Credentials for the given Postgres Database (may take a few seconds to work)
func (c *Client) ResetPostgresDatabaseCredentials(ctx context.Context, id int) error {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/credentials/reset", e, id)
	_, err = coupleAPIErrors(req.Post(e))
	if err != nil {
		return err
	}

	return nil
}

// GetPostgresDatabaseSSL returns the SSL Certificate for the given Postgres Database
func (c *Client) GetPostgresDatabaseSSL(ctx context.Context, id int) (*PostgresDatabaseSSL, error) {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/ssl", e, id)
	r, err := coupleAPIErrors(req.SetResult(&PostgresDatabaseSSL{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*PostgresDatabaseSSL), nil
}

// GetPostgresDatabaseBackup returns a specific Postgres Database Backup with the given ids
func (c *Client) GetPostgresDatabaseBackup(ctx context.Context, databaseID int, backupID int) (*PostgresDatabaseBackup, error) {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/backups/%d", e, databaseID, backupID)
	r, err := coupleAPIErrors(req.SetResult(&PostgresDatabaseBackup{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*PostgresDatabaseBackup), nil
}

// RestorePostgresDatabaseBackup returns the given Postgres Database with the given Backup
func (c *Client) RestorePostgresDatabaseBackup(ctx context.Context, databaseID int, backupID int) error {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/backups/%d/restore", e, databaseID, backupID)
	_, err = coupleAPIErrors(req.Post(e))
	if err != nil {
		return err
	}
	return nil
}

// CreatePostgresDatabaseBackup creates a snapshot for the given Postgres database
func (c *Client) CreatePostgresDatabaseBackup(ctx context.Context, databaseID int, options PostgresBackupCreateOptions) error {
	e, err := c.DatabasePostgresInstances.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	bodyData, err := json.Marshal(options)
	if err != nil {
		return NewError(err)
	}

	body := string(bodyData)

	e = fmt.Sprintf("%s/%d/backups", e, databaseID)
	_, err = coupleAPIErrors(req.SetBody(body).Post(e))
	if err != nil {
		return err
	}

	return nil
}
