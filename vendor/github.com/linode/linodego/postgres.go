package linodego

import (
	"context"
	"encoding/json"
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

// A PostgresDatabase is an instance of Linode Postgres Managed Databases
type PostgresDatabase struct {
	ID        int            `json:"id"`
	Status    DatabaseStatus `json:"status"`
	Label     string         `json:"label"`
	Region    string         `json:"region"`
	Type      string         `json:"type"`
	Engine    string         `json:"engine"`
	Version   string         `json:"version"`
	AllowList []string       `json:"allow_list"`
	Port      int            `json:"port"`

	ClusterSize int              `json:"cluster_size"`
	Platform    DatabasePlatform `json:"platform"`

	// Members has dynamic keys so it is a map
	Members map[string]DatabaseMemberType `json:"members"`

	SSLConnection     bool                         `json:"ssl_connection"`
	Encrypted         bool                         `json:"encrypted"`
	Hosts             DatabaseHost                 `json:"hosts"`
	Updates           DatabaseMaintenanceWindow    `json:"updates"`
	Created           *time.Time                   `json:"-"`
	Updated           *time.Time                   `json:"-"`
	Fork              *DatabaseFork                `json:"fork"`
	OldestRestoreTime *time.Time                   `json:"-"`
	UsedDiskSizeGB    int                          `json:"used_disk_size_gb"`
	TotalDiskSizeGB   int                          `json:"total_disk_size_gb"`
	EngineConfig      PostgresDatabaseEngineConfig `json:"engine_config"`
	PrivateNetwork    *DatabasePrivateNetwork      `json:"private_network,omitempty"`
}

type PostgresDatabaseEngineConfig struct {
	PG                      *PostgresDatabaseEngineConfigPG        `json:"pg,omitempty"`
	PGStatMonitorEnable     *bool                                  `json:"pg_stat_monitor_enable,omitempty"`
	PGLookout               *PostgresDatabaseEngineConfigPGLookout `json:"pglookout,omitempty"`
	SharedBuffersPercentage *float64                               `json:"shared_buffers_percentage,omitempty"`
	WorkMem                 *int                                   `json:"work_mem,omitempty"`
}

type PostgresDatabaseEngineConfigPG struct {
	AutovacuumAnalyzeScaleFactor     *float64 `json:"autovacuum_analyze_scale_factor,omitempty"`
	AutovacuumAnalyzeThreshold       *int32   `json:"autovacuum_analyze_threshold,omitempty"`
	AutovacuumMaxWorkers             *int     `json:"autovacuum_max_workers,omitempty"`
	AutovacuumNaptime                *int     `json:"autovacuum_naptime,omitempty"`
	AutovacuumVacuumCostDelay        *int     `json:"autovacuum_vacuum_cost_delay,omitempty"`
	AutovacuumVacuumCostLimit        *int     `json:"autovacuum_vacuum_cost_limit,omitempty"`
	AutovacuumVacuumScaleFactor      *float64 `json:"autovacuum_vacuum_scale_factor,omitempty"`
	AutovacuumVacuumThreshold        *int32   `json:"autovacuum_vacuum_threshold,omitempty"`
	BGWriterDelay                    *int     `json:"bgwriter_delay,omitempty"`
	BGWriterFlushAfter               *int     `json:"bgwriter_flush_after,omitempty"`
	BGWriterLRUMaxPages              *int     `json:"bgwriter_lru_maxpages,omitempty"`
	BGWriterLRUMultiplier            *float64 `json:"bgwriter_lru_multiplier,omitempty"`
	DeadlockTimeout                  *int     `json:"deadlock_timeout,omitempty"`
	DefaultToastCompression          *string  `json:"default_toast_compression,omitempty"`
	IdleInTransactionSessionTimeout  *int     `json:"idle_in_transaction_session_timeout,omitempty"`
	JIT                              *bool    `json:"jit,omitempty"`
	MaxFilesPerProcess               *int     `json:"max_files_per_process,omitempty"`
	MaxLocksPerTransaction           *int     `json:"max_locks_per_transaction,omitempty"`
	MaxLogicalReplicationWorkers     *int     `json:"max_logical_replication_workers,omitempty"`
	MaxParallelWorkers               *int     `json:"max_parallel_workers,omitempty"`
	MaxParallelWorkersPerGather      *int     `json:"max_parallel_workers_per_gather,omitempty"`
	MaxPredLocksPerTransaction       *int     `json:"max_pred_locks_per_transaction,omitempty"`
	MaxReplicationSlots              *int     `json:"max_replication_slots,omitempty"`
	MaxSlotWALKeepSize               *int32   `json:"max_slot_wal_keep_size,omitempty"`
	MaxStackDepth                    *int     `json:"max_stack_depth,omitempty"`
	MaxStandbyArchiveDelay           *int     `json:"max_standby_archive_delay,omitempty"`
	MaxStandbyStreamingDelay         *int     `json:"max_standby_streaming_delay,omitempty"`
	MaxWALSenders                    *int     `json:"max_wal_senders,omitempty"`
	MaxWorkerProcesses               *int     `json:"max_worker_processes,omitempty"`
	PasswordEncryption               *string  `json:"password_encryption,omitempty"`
	PGPartmanBGWInterval             *int     `json:"pg_partman_bgw.interval,omitempty"`
	PGPartmanBGWRole                 *string  `json:"pg_partman_bgw.role,omitempty"`
	PGStatMonitorPGSMEnableQueryPlan *bool    `json:"pg_stat_monitor.pgsm_enable_query_plan,omitempty"`
	PGStatMonitorPGSMMaxBuckets      *int     `json:"pg_stat_monitor.pgsm_max_buckets,omitempty"`
	PGStatStatementsTrack            *string  `json:"pg_stat_statements.track,omitempty"`
	TempFileLimit                    *int32   `json:"temp_file_limit,omitempty"`
	Timezone                         *string  `json:"timezone,omitempty"`
	TrackActivityQuerySize           *int     `json:"track_activity_query_size,omitempty"`
	TrackCommitTimestamp             *string  `json:"track_commit_timestamp,omitempty"`
	TrackFunctions                   *string  `json:"track_functions,omitempty"`
	TrackIOTiming                    *string  `json:"track_io_timing,omitempty"`
	WALSenderTimeout                 *int     `json:"wal_sender_timeout,omitempty"`
	WALWriterDelay                   *int     `json:"wal_writer_delay,omitempty"`
}

type PostgresDatabaseEngineConfigPGLookout struct {
	MaxFailoverReplicationTimeLag *int64 `json:"max_failover_replication_time_lag,omitempty"`
}

type PostgresDatabaseConfigInfo struct {
	PG                      PostgresDatabaseConfigInfoPG                      `json:"pg"`
	PGStatMonitorEnable     PostgresDatabaseConfigInfoPGStatMonitorEnable     `json:"pg_stat_monitor_enable"`
	PGLookout               PostgresDatabaseConfigInfoPGLookout               `json:"pglookout"`
	SharedBuffersPercentage PostgresDatabaseConfigInfoSharedBuffersPercentage `json:"shared_buffers_percentage"`
	WorkMem                 PostgresDatabaseConfigInfoWorkMem                 `json:"work_mem"`
}

type PostgresDatabaseConfigInfoPG struct {
	AutovacuumAnalyzeScaleFactor     AutovacuumAnalyzeScaleFactor     `json:"autovacuum_analyze_scale_factor"`
	AutovacuumAnalyzeThreshold       AutovacuumAnalyzeThreshold       `json:"autovacuum_analyze_threshold"`
	AutovacuumMaxWorkers             AutovacuumMaxWorkers             `json:"autovacuum_max_workers"`
	AutovacuumNaptime                AutovacuumNaptime                `json:"autovacuum_naptime"`
	AutovacuumVacuumCostDelay        AutovacuumVacuumCostDelay        `json:"autovacuum_vacuum_cost_delay"`
	AutovacuumVacuumCostLimit        AutovacuumVacuumCostLimit        `json:"autovacuum_vacuum_cost_limit"`
	AutovacuumVacuumScaleFactor      AutovacuumVacuumScaleFactor      `json:"autovacuum_vacuum_scale_factor"`
	AutovacuumVacuumThreshold        AutovacuumVacuumThreshold        `json:"autovacuum_vacuum_threshold"`
	BGWriterDelay                    BGWriterDelay                    `json:"bgwriter_delay"`
	BGWriterFlushAfter               BGWriterFlushAfter               `json:"bgwriter_flush_after"`
	BGWriterLRUMaxPages              BGWriterLRUMaxPages              `json:"bgwriter_lru_maxpages"`
	BGWriterLRUMultiplier            BGWriterLRUMultiplier            `json:"bgwriter_lru_multiplier"`
	DeadlockTimeout                  DeadlockTimeout                  `json:"deadlock_timeout"`
	DefaultToastCompression          DefaultToastCompression          `json:"default_toast_compression"`
	IdleInTransactionSessionTimeout  IdleInTransactionSessionTimeout  `json:"idle_in_transaction_session_timeout"`
	JIT                              JIT                              `json:"jit"`
	MaxFilesPerProcess               MaxFilesPerProcess               `json:"max_files_per_process"`
	MaxLocksPerTransaction           MaxLocksPerTransaction           `json:"max_locks_per_transaction"`
	MaxLogicalReplicationWorkers     MaxLogicalReplicationWorkers     `json:"max_logical_replication_workers"`
	MaxParallelWorkers               MaxParallelWorkers               `json:"max_parallel_workers"`
	MaxParallelWorkersPerGather      MaxParallelWorkersPerGather      `json:"max_parallel_workers_per_gather"`
	MaxPredLocksPerTransaction       MaxPredLocksPerTransaction       `json:"max_pred_locks_per_transaction"`
	MaxReplicationSlots              MaxReplicationSlots              `json:"max_replication_slots"`
	MaxSlotWALKeepSize               MaxSlotWALKeepSize               `json:"max_slot_wal_keep_size"`
	MaxStackDepth                    MaxStackDepth                    `json:"max_stack_depth"`
	MaxStandbyArchiveDelay           MaxStandbyArchiveDelay           `json:"max_standby_archive_delay"`
	MaxStandbyStreamingDelay         MaxStandbyStreamingDelay         `json:"max_standby_streaming_delay"`
	MaxWALSenders                    MaxWALSenders                    `json:"max_wal_senders"`
	MaxWorkerProcesses               MaxWorkerProcesses               `json:"max_worker_processes"`
	PasswordEncryption               PasswordEncryption               `json:"password_encryption"`
	PGPartmanBGWInterval             PGPartmanBGWInterval             `json:"pg_partman_bgw.interval"`
	PGPartmanBGWRole                 PGPartmanBGWRole                 `json:"pg_partman_bgw.role"`
	PGStatMonitorPGSMEnableQueryPlan PGStatMonitorPGSMEnableQueryPlan `json:"pg_stat_monitor.pgsm_enable_query_plan"`
	PGStatMonitorPGSMMaxBuckets      PGStatMonitorPGSMMaxBuckets      `json:"pg_stat_monitor.pgsm_max_buckets"`
	PGStatStatementsTrack            PGStatStatementsTrack            `json:"pg_stat_statements.track"`
	TempFileLimit                    TempFileLimit                    `json:"temp_file_limit"`
	Timezone                         Timezone                         `json:"timezone"`
	TrackActivityQuerySize           TrackActivityQuerySize           `json:"track_activity_query_size"`
	TrackCommitTimestamp             TrackCommitTimestamp             `json:"track_commit_timestamp"`
	TrackFunctions                   TrackFunctions                   `json:"track_functions"`
	TrackIOTiming                    TrackIOTiming                    `json:"track_io_timing"`
	WALSenderTimeout                 WALSenderTimeout                 `json:"wal_sender_timeout"`
	WALWriterDelay                   WALWriterDelay                   `json:"wal_writer_delay"`
}

type AutovacuumAnalyzeScaleFactor struct {
	Description     string  `json:"description"`
	Maximum         float64 `json:"maximum"`
	Minimum         float64 `json:"minimum"`
	RequiresRestart bool    `json:"requires_restart"`
	Type            string  `json:"type"`
}

type AutovacuumAnalyzeThreshold struct {
	Description     string `json:"description"`
	Maximum         int32  `json:"maximum"`
	Minimum         int32  `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type AutovacuumMaxWorkers struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type AutovacuumNaptime struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type AutovacuumVacuumCostDelay struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type AutovacuumVacuumCostLimit struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type AutovacuumVacuumScaleFactor struct {
	Description     string  `json:"description"`
	Maximum         float64 `json:"maximum"`
	Minimum         float64 `json:"minimum"`
	RequiresRestart bool    `json:"requires_restart"`
	Type            string  `json:"type"`
}

type AutovacuumVacuumThreshold struct {
	Description     string `json:"description"`
	Maximum         int32  `json:"maximum"`
	Minimum         int32  `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type BGWriterDelay struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type BGWriterFlushAfter struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type BGWriterLRUMaxPages struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type BGWriterLRUMultiplier struct {
	Description     string  `json:"description"`
	Example         float64 `json:"example"`
	Maximum         float64 `json:"maximum"`
	Minimum         float64 `json:"minimum"`
	RequiresRestart bool    `json:"requires_restart"`
	Type            string  `json:"type"`
}

type DeadlockTimeout struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type DefaultToastCompression struct {
	Description     string   `json:"description"`
	Enum            []string `json:"enum"`
	Example         string   `json:"example"`
	RequiresRestart bool     `json:"requires_restart"`
	Type            string   `json:"type"`
}

type IdleInTransactionSessionTimeout struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type JIT struct {
	Description     string `json:"description"`
	Example         bool   `json:"example"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxFilesPerProcess struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxLocksPerTransaction struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxLogicalReplicationWorkers struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxParallelWorkers struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxParallelWorkersPerGather struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxPredLocksPerTransaction struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxReplicationSlots struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxSlotWALKeepSize struct {
	Description     string `json:"description"`
	Maximum         int32  `json:"maximum"`
	Minimum         int32  `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxStackDepth struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxStandbyArchiveDelay struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxStandbyStreamingDelay struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxWALSenders struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type MaxWorkerProcesses struct {
	Description     string `json:"description"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PasswordEncryption struct {
	Description     string   `json:"description"`
	Enum            []string `json:"enum"`
	Example         string   `json:"example"`
	RequiresRestart bool     `json:"requires_restart"`
	Type            string   `json:"type"`
}

type PGPartmanBGWInterval struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PGPartmanBGWRole struct {
	Description     string `json:"description"`
	Example         string `json:"example"`
	MaxLength       int    `json:"maxLength"`
	Pattern         string `json:"pattern"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PGStatMonitorPGSMEnableQueryPlan struct {
	Description     string `json:"description"`
	Example         bool   `json:"example"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PGStatMonitorPGSMMaxBuckets struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PGStatStatementsTrack struct {
	Description     string   `json:"description"`
	Enum            []string `json:"enum"`
	RequiresRestart bool     `json:"requires_restart"`
	Type            string   `json:"type"`
}

type TempFileLimit struct {
	Description     string `json:"description"`
	Example         int32  `json:"example"`
	Maximum         int32  `json:"maximum"`
	Minimum         int32  `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type Timezone struct {
	Description     string `json:"description"`
	Example         string `json:"example"`
	MaxLength       int    `json:"maxLength"`
	Pattern         string `json:"pattern"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type TrackActivityQuerySize struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type TrackCommitTimestamp struct {
	Description     string   `json:"description"`
	Enum            []string `json:"enum"`
	Example         string   `json:"example"`
	RequiresRestart bool     `json:"requires_restart"`
	Type            string   `json:"type"`
}

type TrackFunctions struct {
	Description     string   `json:"description"`
	Enum            []string `json:"enum"`
	RequiresRestart bool     `json:"requires_restart"`
	Type            string   `json:"type"`
}

type TrackIOTiming struct {
	Description     string   `json:"description"`
	Enum            []string `json:"enum"`
	Example         string   `json:"example"`
	RequiresRestart bool     `json:"requires_restart"`
	Type            string   `json:"type"`
}

type WALSenderTimeout struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type WALWriterDelay struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PostgresDatabaseConfigInfoPGStatMonitorEnable struct {
	Description     string `json:"description"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PostgresDatabaseConfigInfoPGLookout struct {
	PGLookoutMaxFailoverReplicationTimeLag PGLookoutMaxFailoverReplicationTimeLag `json:"max_failover_replication_time_lag"`
}

type PGLookoutMaxFailoverReplicationTimeLag struct {
	Description     string `json:"description"`
	Maximum         int64  `json:"maximum"`
	Minimum         int64  `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

type PostgresDatabaseConfigInfoSharedBuffersPercentage struct {
	Description     string  `json:"description"`
	Example         float64 `json:"example"`
	Maximum         float64 `json:"maximum"`
	Minimum         float64 `json:"minimum"`
	RequiresRestart bool    `json:"requires_restart"`
	Type            string  `json:"type"`
}

type PostgresDatabaseConfigInfoWorkMem struct {
	Description     string `json:"description"`
	Example         int    `json:"example"`
	Maximum         int    `json:"maximum"`
	Minimum         int    `json:"minimum"`
	RequiresRestart bool   `json:"requires_restart"`
	Type            string `json:"type"`
}

func (d *PostgresDatabase) UnmarshalJSON(b []byte) error {
	type Mask PostgresDatabase

	p := struct {
		*Mask

		Created           *parseabletime.ParseableTime `json:"created"`
		Updated           *parseabletime.ParseableTime `json:"updated"`
		OldestRestoreTime *parseabletime.ParseableTime `json:"oldest_restore_time"`
	}{
		Mask: (*Mask)(d),
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	d.Created = (*time.Time)(p.Created)
	d.Updated = (*time.Time)(p.Updated)
	d.OldestRestoreTime = (*time.Time)(p.OldestRestoreTime)

	return nil
}

// PostgresCreateOptions fields are used when creating a new Postgres Database
type PostgresCreateOptions struct {
	Label       string   `json:"label"`
	Region      string   `json:"region"`
	Type        string   `json:"type"`
	Engine      string   `json:"engine"`
	AllowList   []string `json:"allow_list,omitempty"`
	ClusterSize int      `json:"cluster_size,omitempty"`

	Fork *DatabaseFork `json:"fork,omitempty"`

	EngineConfig   *PostgresDatabaseEngineConfig `json:"engine_config,omitempty"`
	PrivateNetwork *DatabasePrivateNetwork       `json:"private_network,omitempty"`
}

// PostgresUpdateOptions fields are used when altering the existing Postgres Database
type PostgresUpdateOptions struct {
	Label          string                        `json:"label,omitempty"`
	AllowList      *[]string                     `json:"allow_list,omitempty"`
	Updates        *DatabaseMaintenanceWindow    `json:"updates,omitempty"`
	Type           string                        `json:"type,omitempty"`
	ClusterSize    int                           `json:"cluster_size,omitempty"`
	Version        string                        `json:"version,omitempty"`
	EngineConfig   *PostgresDatabaseEngineConfig `json:"engine_config,omitempty"`
	PrivateNetwork *DatabasePrivateNetwork       `json:"private_network,omitempty"`
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

// ListPostgresDatabases lists all Postgres Databases associated with the account
func (c *Client) ListPostgresDatabases(ctx context.Context, opts *ListOptions) ([]PostgresDatabase, error) {
	return getPaginatedResults[PostgresDatabase](ctx, c, "databases/postgresql/instances", opts)
}

// GetPostgresDatabase returns a single Postgres Database matching the id
func (c *Client) GetPostgresDatabase(ctx context.Context, databaseID int) (*PostgresDatabase, error) {
	e := formatAPIPath("databases/postgresql/instances/%d", databaseID)
	return doGETRequest[PostgresDatabase](ctx, c, e)
}

// CreatePostgresDatabase creates a new Postgres Database using the createOpts as configuration, returns the new Postgres Database
func (c *Client) CreatePostgresDatabase(ctx context.Context, opts PostgresCreateOptions) (*PostgresDatabase, error) {
	return doPOSTRequest[PostgresDatabase](ctx, c, "databases/postgresql/instances", opts)
}

// DeletePostgresDatabase deletes an existing Postgres Database with the given id
func (c *Client) DeletePostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d", databaseID)
	return doDELETERequest(ctx, c, e)
}

// UpdatePostgresDatabase updates the given Postgres Database with the provided opts, returns the PostgresDatabase with the new settings
func (c *Client) UpdatePostgresDatabase(ctx context.Context, databaseID int, opts PostgresUpdateOptions) (*PostgresDatabase, error) {
	e := formatAPIPath("databases/postgresql/instances/%d", databaseID)
	return doPUTRequest[PostgresDatabase](ctx, c, e, opts)
}

// PatchPostgresDatabase applies security patches and updates to the underlying operating system of the Managed Postgres Database
func (c *Client) PatchPostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/patch", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// GetPostgresDatabaseCredentials returns the Root Credentials for the given Postgres Database
func (c *Client) GetPostgresDatabaseCredentials(ctx context.Context, databaseID int) (*PostgresDatabaseCredential, error) {
	e := formatAPIPath("databases/postgresql/instances/%d/credentials", databaseID)
	return doGETRequest[PostgresDatabaseCredential](ctx, c, e)
}

// ResetPostgresDatabaseCredentials returns the Root Credentials for the given Postgres Database (may take a few seconds to work)
func (c *Client) ResetPostgresDatabaseCredentials(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/credentials/reset", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// GetPostgresDatabaseSSL returns the SSL Certificate for the given Postgres Database
func (c *Client) GetPostgresDatabaseSSL(ctx context.Context, databaseID int) (*PostgresDatabaseSSL, error) {
	e := formatAPIPath("databases/postgresql/instances/%d/ssl", databaseID)
	return doGETRequest[PostgresDatabaseSSL](ctx, c, e)
}

// SuspendPostgresDatabase suspends a PostgreSQL Managed Database, releasing idle resources and keeping only necessary data.
// All service data is lost if there are no backups available.
func (c *Client) SuspendPostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/suspend", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// ResumePostgresDatabase resumes a suspended PostgreSQL Managed Database
func (c *Client) ResumePostgresDatabase(ctx context.Context, databaseID int) error {
	e := formatAPIPath("databases/postgresql/instances/%d/resume", databaseID)
	return doPOSTRequestNoRequestResponseBody(ctx, c, e)
}

// GetPostgresDatabaseConfig returns a detailed list of all the configuration options for PostgreSQL Databases
func (c *Client) GetPostgresDatabaseConfig(ctx context.Context) (*PostgresDatabaseConfigInfo, error) {
	return doGETRequest[PostgresDatabaseConfigInfo](ctx, c, "databases/postgresql/config")
}
