package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/linode/linodego/internal/parseabletime"
)

type (
	DatabaseEngineType           string
	DatabaseDayOfWeek            int
	DatabaseMaintenanceFrequency string
	DatabaseStatus               string
)

const (
	DatabaseMaintenanceDaySunday DatabaseDayOfWeek = iota + 1
	DatabaseMaintenanceDayMonday
	DatabaseMaintenanceDayTuesday
	DatabaseMaintenanceDayWednesday
	DatabaseMaintenanceDayThursday
	DatabaseMaintenanceDayFriday
	DatabaseMaintenanceDaySaturday
)

const (
	DatabaseMaintenanceFrequencyWeekly  DatabaseMaintenanceFrequency = "weekly"
	DatabaseMaintenanceFrequencyMonthly DatabaseMaintenanceFrequency = "monthly"
)

const (
	DatabaseEngineTypeMySQL    DatabaseEngineType = "mysql"
	DatabaseEngineTypeMongo    DatabaseEngineType = "mongodb"
	DatabaseEngineTypePostgres DatabaseEngineType = "postgresql"
)

const (
	DatabaseStatusProvisioning DatabaseStatus = "provisioning"
	DatabaseStatusActive       DatabaseStatus = "active"
	DatabaseStatusDeleting     DatabaseStatus = "deleting"
	DatabaseStatusDeleted      DatabaseStatus = "deleted"
	DatabaseStatusSuspending   DatabaseStatus = "suspending"
	DatabaseStatusSuspended    DatabaseStatus = "suspended"
	DatabaseStatusResuming     DatabaseStatus = "resuming"
	DatabaseStatusRestoring    DatabaseStatus = "restoring"
	DatabaseStatusFailed       DatabaseStatus = "failed"
	DatabaseStatusDegraded     DatabaseStatus = "degraded"
	DatabaseStatusUpdating     DatabaseStatus = "updating"
	DatabaseStatusBackingUp    DatabaseStatus = "backing_up"
)

type DatabasesPagedResponse struct {
	*PageOptions
	Data []Database `json:"data"`
}

func (DatabasesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Databases.Endpoint()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/instances", endpoint)
}

func (resp *DatabasesPagedResponse) appendData(r *DatabasesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

type DatabaseEnginesPagedResponse struct {
	*PageOptions
	Data []DatabaseEngine `json:"data"`
}

func (DatabaseEnginesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Databases.Endpoint()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/engines", endpoint)
}

func (resp *DatabaseEnginesPagedResponse) appendData(r *DatabaseEnginesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

type DatabaseTypesPagedResponse struct {
	*PageOptions
	Data []DatabaseType `json:"data"`
}

func (DatabaseTypesPagedResponse) endpoint(c *Client) string {
	endpoint, err := c.Databases.Endpoint()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/types", endpoint)
}

func (resp *DatabaseTypesPagedResponse) appendData(r *DatabaseTypesPagedResponse) {
	resp.Data = append(resp.Data, r.Data...)
}

// A Database is a instance of Linode Managed Databases
type Database struct {
	ID              int            `json:"id"`
	Status          DatabaseStatus `json:"status"`
	Label           string         `json:"label"`
	Hosts           DatabaseHost   `json:"hosts"`
	Region          string         `json:"region"`
	Type            string         `json:"type"`
	Engine          string         `json:"engine"`
	Version         string         `json:"version"`
	ClusterSize     int            `json:"cluster_size"`
	ReplicationType string         `json:"replication_type"`
	SSLConnection   bool           `json:"ssl_connection"`
	Encrypted       bool           `json:"encrypted"`
	AllowList       []string       `json:"allow_list"`
	InstanceURI     string         `json:"instance_uri"`
	Created         *time.Time     `json:"-"`
	Updated         *time.Time     `json:"-"`
}

// DatabaseHost for Primary/Secondary of Database
type DatabaseHost struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary,omitempty"`
}

// DatabaseEngine is information about Engines supported by Linode Managed Databases
type DatabaseEngine struct {
	ID      string `json:"id"`
	Engine  string `json:"engine"`
	Version string `json:"version"`
}

// DatabaseMaintenanceWindow stores information about a MySQL cluster's maintenance window
type DatabaseMaintenanceWindow struct {
	DayOfWeek   DatabaseDayOfWeek            `json:"day_of_week"`
	Duration    int                          `json:"duration"`
	Frequency   DatabaseMaintenanceFrequency `json:"frequency"`
	HourOfDay   int                          `json:"hour_of_day"`
	WeekOfMonth *int                         `json:"week_of_month"`
}

// DatabaseType is information about the supported Database Types by Linode Managed Databases
type DatabaseType struct {
	ID          string                `json:"id"`
	Label       string                `json:"label"`
	Class       string                `json:"class"`
	VirtualCPUs int                   `json:"vcpus"`
	Disk        int                   `json:"disk"`
	Memory      int                   `json:"memory"`
	Engines     DatabaseTypeEngineMap `json:"engines"`
}

// DatabaseTypeEngineMap stores a list of Database Engine types by engine
type DatabaseTypeEngineMap struct {
	MySQL []DatabaseTypeEngine `json:"mysql"`
}

// DatabaseTypeEngine Sizes and Prices
type DatabaseTypeEngine struct {
	Quantity int          `json:"quantity"`
	Price    ClusterPrice `json:"price"`
}

// ClusterPrice for Hourly and Monthly price models
type ClusterPrice struct {
	Hourly  float32 `json:"hourly"`
	Monthly float32 `json:"monthly"`
}

func (d *Database) UnmarshalJSON(b []byte) error {
	type Mask Database

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

// ListDatabases lists all Database instances in Linode Managed Databases for the account
func (c *Client) ListDatabases(ctx context.Context, opts *ListOptions) ([]Database, error) {
	response := DatabasesPagedResponse{}

	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// ListDatabaseEngines lists all Database Engines
func (c *Client) ListDatabaseEngines(ctx context.Context, opts *ListOptions) ([]DatabaseEngine, error) {
	response := DatabaseEnginesPagedResponse{}

	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetDatabaseEngine returns a specific Database Engine
func (c *Client) GetDatabaseEngine(ctx context.Context, opts *ListOptions, id string) (*DatabaseEngine, error) {
	e, err := c.Databases.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/engines/%s", e, id)
	r, err := coupleAPIErrors(req.SetResult(&DatabaseEngine{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*DatabaseEngine), nil
}

// ListDatabaseTypes lists all Types of Database provided in Linode Managed Databases
func (c *Client) ListDatabaseTypes(ctx context.Context, opts *ListOptions) ([]DatabaseType, error) {
	response := DatabaseTypesPagedResponse{}

	err := c.listHelper(ctx, &response, opts)
	if err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetDatabaseType returns a specific Database Type
func (c *Client) GetDatabaseType(ctx context.Context, opts *ListOptions, id string) (*DatabaseType, error) {
	e, err := c.Databases.Endpoint()
	if err != nil {
		return nil, err
	}

	req := c.R(ctx)

	e = fmt.Sprintf("%s/types/%s", e, id)
	r, err := coupleAPIErrors(req.SetResult(&DatabaseType{}).Get(e))
	if err != nil {
		return nil, err
	}

	return r.Result().(*DatabaseType), nil
}
