package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type MySQLDatabaseTarget string

type MySQLDatabaseMaintenanceWindow = DatabaseMaintenanceWindow

const (
	MySQLDatabaseTargetPrimary   MySQLDatabaseTarget = "primary"
	MySQLDatabaseTargetSecondary MySQLDatabaseTarget = "secondary"
)

// A MySQLDatabase is a instance of Linode MySQL Managed Databases
type MySQLDatabase struct {
	ID              int                       `json:"id"`
	Status          DatabaseStatus            `json:"status"`
	Label           string                    `json:"label"`
	Hosts           DatabaseHost              `json:"hosts"`
	Region          string                    `json:"region"`
	Type            string                    `json:"type"`
	Engine          string                    `json:"engine"`
	Version         string                    `json:"version"`
	ClusterSize     int                       `json:"cluster_size"`
	ReplicationType string                    `json:"replication_type"`
	SSLConnection   bool                      `json:"ssl_connection"`
	Encrypted       bool                      `json:"encrypted"`
	AllowList       []string                  `json:"allow_list"`
	InstanceURI     string                    `json:"instance_uri"`
	Created         *time.Time                `json:"-"`
	Updated         *time.Time                `json:"-"`
	Updates         DatabaseMaintenanceWindow `json:"updates"`
}

func (d *MySQLDatabase) UnmarshalJSON(b []byte) error {
	type Mask MySQLDatabase

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

// MySQLCreateOptions fields are used when creating a new MySQL Database
type MySQLCreateOptions struct {
	Label           string   `json:"label"`
	Region          string   `json:"region"`
	Type            string   `json:"type"`
	Engine          string   `json:"engine"`
	AllowList       []string `json:"allow_list,omitempty"`
	ReplicationType string   `json:"replication_type,omitempty"`
	ClusterSize     int      `json:"cluster_size,omitempty"`
	Encrypted       bool     `json:"encrypted,omitempty"`
	SSLConnection   bool     `json:"ssl_connection,omitempty"`
}

// MySQLUpdateOptions fields are used when altering the existing MySQL Database
type MySQLUpdateOptions struct {
	Label     string                     `json:"label,omitempty"`
	AllowList *[]string                  `json:"allow_list,omitempty"`
	Updates   *DatabaseMaintenanceWindow `json:"updates,omitempty"`
}

// MySQLDatabaseBackup is information for interacting with a backup for the existing MySQL Database
type MySQLDatabaseBackup struct {
	ID      int        `json:"id"`
	Label   string     `json:"label"`
	Type    string     `json:"type"`
	Created *time.Time `json:"-"`
}

// MySQLBackupCreateOptions are options used for CreateMySQLDatabaseBackup(...)
type MySQLBackupCreateOptions struct {
	Label  string              `json:"label"`
	Target MySQLDatabaseTarget `json:"target"`
}

func (d *MySQLDatabaseBackup) UnmarshalJSON(b []byte) error {
	type Mask MySQLDatabaseBackup

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

type MySQLDatabasesPagedResponse struct {
	*PageOptions
	Data []MySQLDatabase `json:"data"`
}

func (MySQLDatabasesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		panic(err)
	}
	return endpoint
}

func (resp *MySQLDatabasesPagedResponse) appendData(r *MySQLDatabasesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// MySQLDatabaseCredential is the Root Credentials to access the Linode Managed Database
type MySQLDatabaseCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// MySQLDatabaseSSL is the SSL Certificate to access the Linode Managed MySQL Database
type MySQLDatabaseSSL struct {
	CACertificate []byte `json:"ca_certificate"`
}

// ListMySQLDatabases lists all MySQL Databases associated with the account
func (c *Client) ListMySQLDatabases(ctx context.Context, opts *ListOptions) ([]MySQLDatabase, error) {
	response := MySQLDatabasesPagedResponse{}

	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

type MySQLDatabaseBackupsPagedResponse struct {
	*PageOptions
	Data []MySQLDatabaseBackup `json:"data"`
}

func (MySQLDatabaseBackupsPagedResponse) endpointWithID(c *Client, id int) string {
	endpoint, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%d/backups", endpoint, id)
}

func (resp *MySQLDatabaseBackupsPagedResponse) appendData(r *MySQLDatabaseBackupsPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// ListMySQLDatabaseBackups lists all MySQL Database Backups associated with the given MySQL Database
func (c *Client) ListMySQLDatabaseBackups(ctx context.Context, databaseID int, opts *ListOptions) ([]MySQLDatabaseBackup, error) {
	response := MySQLDatabaseBackupsPagedResponse{}

	err := c.listHelperWithID(ctx, &response, databaseID, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetMySQLDatabase returns a single MySQL Database matching the id
func (c *Client) GetMySQLDatabase(ctx context.Context, id int) (*MySQLDatabase, error) {
	e, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d", e, id)
	r, err := coupleAPIErrors(req.SetResult(&MySQLDatabase{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*MySQLDatabase), nil
}

// CreateMySQLDatabase creates a new MySQL Database using the createOpts as configuration, returns the new MySQL Database
func (c *Client) CreateMySQLDatabase(ctx context.Context, createOpts MySQLCreateOptions) (*MySQLDatabase, error) {
	var body string
	e, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx).SetResult(&MySQLDatabase{})

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
	return r.Result().(*MySQLDatabase), nil
}

// DeleteMySQLDatabase deletes an existing MySQL Database with the given id
func (c *Client) DeleteMySQLDatabase(ctx context.Context, id int) error {
	e, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		return err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d", e, id)
	_, err = coupleAPIErrors(req.Delete(e))
	return err
}

// UpdateMySQLDatabase updates the given MySQL Database with the provided opts, returns the MySQLDatabase with the new settings
func (c *Client) UpdateMySQLDatabase(ctx context.Context, id int, opts MySQLUpdateOptions) (*MySQLDatabase, error) {
	e, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		return nil, err
	}
	req := c.R(ctx).SetResult(&MySQLDatabase{})

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

	return r.Result().(*MySQLDatabase), nil
}

// GetMySQLDatabaseSSL returns the SSL Certificate for the given MySQL Database
func (c *Client) GetMySQLDatabaseSSL(ctx context.Context, id int) (*MySQLDatabaseSSL, error) {
	e, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/ssl", e, id)
	r, err := coupleAPIErrors(req.SetResult(&MySQLDatabaseSSL{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*MySQLDatabaseSSL), nil
}

// GetMySQLDatabaseCredentials returns the Root Credentials for the given MySQL Database
func (c *Client) GetMySQLDatabaseCredentials(ctx context.Context, id int) (*MySQLDatabaseCredential, error) {
	e, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/credentials", e, id)
	r, err := coupleAPIErrors(req.SetResult(&MySQLDatabaseCredential{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*MySQLDatabaseCredential), nil
}

// ResetMySQLDatabaseCredentials returns the Root Credentials for the given MySQL Database (may take a few seconds to work)
func (c *Client) ResetMySQLDatabaseCredentials(ctx context.Context, id int) error {
	e, err := c.DatabaseMySQLInstances.Endpoint()
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

// GetMySQLDatabaseBackup returns a specific MySQL Database Backup with the given ids
func (c *Client) GetMySQLDatabaseBackup(ctx context.Context, databaseID int, backupID int) (*MySQLDatabaseBackup, error) {
	e, err := c.DatabaseMySQLInstances.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/%d/backups/%d", e, databaseID, backupID)
	r, err := coupleAPIErrors(req.SetResult(&MySQLDatabaseBackup{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*MySQLDatabaseBackup), nil
}

// RestoreMySQLDatabaseBackup returns the given MySQL Database with the given Backup
func (c *Client) RestoreMySQLDatabaseBackup(ctx context.Context, databaseID int, backupID int) error {
	e, err := c.DatabaseMySQLInstances.Endpoint()
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

// CreateMySQLDatabaseBackup creates a snapshot for the given MySQL database
func (c *Client) CreateMySQLDatabaseBackup(ctx context.Context, databaseID int, options MySQLBackupCreateOptions) error {
	e, err := c.DatabaseMySQLInstances.Endpoint()
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

// PatchMySQLDatabase applies security patches and updates to the underlying operating system of the Managed MySQL Database
func (c *Client) PatchMySQLDatabase(ctx context.Context, databaseID int) error {
	e, err := c.DatabaseMySQLInstances.Endpoint()
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
