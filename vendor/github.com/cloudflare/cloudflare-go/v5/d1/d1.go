// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// D1Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewD1Service] method instead.
type D1Service struct {
	Options  []option.RequestOption
	Database *DatabaseService
}

// NewD1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewD1Service(opts ...option.RequestOption) (r *D1Service) {
	r = &D1Service{}
	r.Options = opts
	r.Database = NewDatabaseService(opts...)
	return
}

// The details of the D1 database.
type D1 struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The D1 database's size, in bytes.
	FileSize float64 `json:"file_size"`
	// D1 database name.
	Name      string  `json:"name"`
	NumTables float64 `json:"num_tables"`
	// Configuration for D1 read replication.
	ReadReplication D1ReadReplication `json:"read_replication"`
	// D1 database identifier (UUID).
	UUID    string `json:"uuid"`
	Version string `json:"version"`
	JSON    d1JSON `json:"-"`
}

// d1JSON contains the JSON metadata for the struct [D1]
type d1JSON struct {
	CreatedAt       apijson.Field
	FileSize        apijson.Field
	Name            apijson.Field
	NumTables       apijson.Field
	ReadReplication apijson.Field
	UUID            apijson.Field
	Version         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *D1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1JSON) RawJSON() string {
	return r.raw
}

// Configuration for D1 read replication.
type D1ReadReplication struct {
	// The read replication mode for the database. Use 'auto' to create replicas and
	// allow D1 automatically place them around the world, or 'disabled' to not use any
	// database replicas (it can take a few hours for all replicas to be deleted).
	Mode D1ReadReplicationMode `json:"mode,required"`
	JSON d1ReadReplicationJSON `json:"-"`
}

// d1ReadReplicationJSON contains the JSON metadata for the struct
// [D1ReadReplication]
type d1ReadReplicationJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1ReadReplication) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1ReadReplicationJSON) RawJSON() string {
	return r.raw
}

// The read replication mode for the database. Use 'auto' to create replicas and
// allow D1 automatically place them around the world, or 'disabled' to not use any
// database replicas (it can take a few hours for all replicas to be deleted).
type D1ReadReplicationMode string

const (
	D1ReadReplicationModeAuto     D1ReadReplicationMode = "auto"
	D1ReadReplicationModeDisabled D1ReadReplicationMode = "disabled"
)

func (r D1ReadReplicationMode) IsKnown() bool {
	switch r {
	case D1ReadReplicationModeAuto, D1ReadReplicationModeDisabled:
		return true
	}
	return false
}
