// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// DatabaseService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatabaseService] method instead.
type DatabaseService struct {
	Options []option.RequestOption
}

// NewDatabaseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatabaseService(opts ...option.RequestOption) (r *DatabaseService) {
	r = &DatabaseService{}
	r.Options = opts
	return
}

// Returns the created D1 database.
func (r *DatabaseService) New(ctx context.Context, params DatabaseNewParams, opts ...option.RequestOption) (res *D1, err error) {
	var env DatabaseNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the specified D1 database.
func (r *DatabaseService) Update(ctx context.Context, databaseID string, params DatabaseUpdateParams, opts ...option.RequestOption) (res *D1, err error) {
	var env DatabaseUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a list of D1 databases.
func (r *DatabaseService) List(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DatabaseListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns a list of D1 databases.
func (r *DatabaseService) ListAutoPaging(ctx context.Context, params DatabaseListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DatabaseListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes the specified D1 database.
func (r *DatabaseService) Delete(ctx context.Context, databaseID string, body DatabaseDeleteParams, opts ...option.RequestOption) (res *DatabaseDeleteResponse, err error) {
	var env DatabaseDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", body.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates partially the specified D1 database.
func (r *DatabaseService) Edit(ctx context.Context, databaseID string, params DatabaseEditParams, opts ...option.RequestOption) (res *D1, err error) {
	var env DatabaseEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns a URL where the SQL contents of your D1 can be downloaded. Note: this
// process may take some time for larger DBs, during which your D1 will be
// unavailable to serve queries. To avoid blocking your DB unnecessarily, an
// in-progress export must be continually polled or will automatically cancel.
func (r *DatabaseService) Export(ctx context.Context, databaseID string, params DatabaseExportParams, opts ...option.RequestOption) (res *DatabaseExportResponse, err error) {
	var env DatabaseExportResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/export", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the specified D1 database.
func (r *DatabaseService) Get(ctx context.Context, databaseID string, query DatabaseGetParams, opts ...option.RequestOption) (res *D1, err error) {
	var env DatabaseGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s", query.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a temporary URL for uploading an SQL file to, then instructing the D1
// to import it and polling it for status updates. Imports block the D1 for their
// duration.
func (r *DatabaseService) Import(ctx context.Context, databaseID string, params DatabaseImportParams, opts ...option.RequestOption) (res *DatabaseImportResponse, err error) {
	var env DatabaseImportResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/import", params.AccountID, databaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the query result as an object.
func (r *DatabaseService) Query(ctx context.Context, databaseID string, params DatabaseQueryParams, opts ...option.RequestOption) (res *pagination.SinglePage[QueryResult], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/query", params.AccountID, databaseID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns the query result as an object.
func (r *DatabaseService) QueryAutoPaging(ctx context.Context, databaseID string, params DatabaseQueryParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[QueryResult] {
	return pagination.NewSinglePageAutoPager(r.Query(ctx, databaseID, params, opts...))
}

// Returns the query result rows as arrays rather than objects. This is a
// performance-optimized version of the /query endpoint.
func (r *DatabaseService) Raw(ctx context.Context, databaseID string, params DatabaseRawParams, opts ...option.RequestOption) (res *pagination.SinglePage[DatabaseRawResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if databaseID == "" {
		err = errors.New("missing required database_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/d1/database/%s/raw", params.AccountID, databaseID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns the query result rows as arrays rather than objects. This is a
// performance-optimized version of the /query endpoint.
func (r *DatabaseService) RawAutoPaging(ctx context.Context, databaseID string, params DatabaseRawParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DatabaseRawResponse] {
	return pagination.NewSinglePageAutoPager(r.Raw(ctx, databaseID, params, opts...))
}

type QueryResult struct {
	Meta    QueryResultMeta `json:"meta"`
	Results []interface{}   `json:"results"`
	Success bool            `json:"success"`
	JSON    queryResultJSON `json:"-"`
}

// queryResultJSON contains the JSON metadata for the struct [QueryResult]
type queryResultJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResultJSON) RawJSON() string {
	return r.raw
}

type QueryResultMeta struct {
	// Denotes if the database has been altered in some way, like deleting rows.
	ChangedDB bool `json:"changed_db"`
	// Rough indication of how many rows were modified by the query, as provided by
	// SQLite's `sqlite3_total_changes()`.
	Changes float64 `json:"changes"`
	// The duration of the SQL query execution inside the database. Does not include
	// any network communication.
	Duration float64 `json:"duration"`
	// The row ID of the last inserted row in a table with an `INTEGER PRIMARY KEY` as
	// provided by SQLite. Tables created with `WITHOUT ROWID` do not populate this.
	LastRowID float64 `json:"last_row_id"`
	// Number of rows read during the SQL query execution, including indices (not all
	// rows are necessarily returned).
	RowsRead float64 `json:"rows_read"`
	// Number of rows written during the SQL query execution, including indices.
	RowsWritten float64 `json:"rows_written"`
	// Denotes if the query has been handled by the database primary instance.
	ServedByPrimary bool `json:"served_by_primary"`
	// Region location hint of the database instance that handled the query.
	ServedByRegion QueryResultMetaServedByRegion `json:"served_by_region"`
	// Size of the database after the query committed, in bytes.
	SizeAfter float64 `json:"size_after"`
	// Various durations for the query.
	Timings QueryResultMetaTimings `json:"timings"`
	JSON    queryResultMetaJSON    `json:"-"`
}

// queryResultMetaJSON contains the JSON metadata for the struct [QueryResultMeta]
type queryResultMetaJSON struct {
	ChangedDB       apijson.Field
	Changes         apijson.Field
	Duration        apijson.Field
	LastRowID       apijson.Field
	RowsRead        apijson.Field
	RowsWritten     apijson.Field
	ServedByPrimary apijson.Field
	ServedByRegion  apijson.Field
	SizeAfter       apijson.Field
	Timings         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *QueryResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResultMetaJSON) RawJSON() string {
	return r.raw
}

// Region location hint of the database instance that handled the query.
type QueryResultMetaServedByRegion string

const (
	QueryResultMetaServedByRegionWnam QueryResultMetaServedByRegion = "WNAM"
	QueryResultMetaServedByRegionEnam QueryResultMetaServedByRegion = "ENAM"
	QueryResultMetaServedByRegionWeur QueryResultMetaServedByRegion = "WEUR"
	QueryResultMetaServedByRegionEeur QueryResultMetaServedByRegion = "EEUR"
	QueryResultMetaServedByRegionApac QueryResultMetaServedByRegion = "APAC"
	QueryResultMetaServedByRegionOc   QueryResultMetaServedByRegion = "OC"
)

func (r QueryResultMetaServedByRegion) IsKnown() bool {
	switch r {
	case QueryResultMetaServedByRegionWnam, QueryResultMetaServedByRegionEnam, QueryResultMetaServedByRegionWeur, QueryResultMetaServedByRegionEeur, QueryResultMetaServedByRegionApac, QueryResultMetaServedByRegionOc:
		return true
	}
	return false
}

// Various durations for the query.
type QueryResultMetaTimings struct {
	// The duration of the SQL query execution inside the database. Does not include
	// any network communication.
	SqlDurationMs float64                    `json:"sql_duration_ms"`
	JSON          queryResultMetaTimingsJSON `json:"-"`
}

// queryResultMetaTimingsJSON contains the JSON metadata for the struct
// [QueryResultMetaTimings]
type queryResultMetaTimingsJSON struct {
	SqlDurationMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *QueryResultMetaTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryResultMetaTimingsJSON) RawJSON() string {
	return r.raw
}

type DatabaseListResponse struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// D1 database name.
	Name string `json:"name"`
	// D1 database identifier (UUID).
	UUID    string                   `json:"uuid"`
	Version string                   `json:"version"`
	JSON    databaseListResponseJSON `json:"-"`
}

// databaseListResponseJSON contains the JSON metadata for the struct
// [DatabaseListResponse]
type databaseListResponseJSON struct {
	CreatedAt   apijson.Field
	Name        apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseListResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseDeleteResponse = interface{}

type DatabaseExportResponse struct {
	// The current time-travel bookmark for your D1, used to poll for updates. Will not
	// change for the duration of the export task.
	AtBookmark string `json:"at_bookmark"`
	// Only present when status = 'error'. Contains the error message.
	Error string `json:"error"`
	// Logs since the last time you polled
	Messages []string `json:"messages"`
	// Only present when status = 'complete'
	Result  DatabaseExportResponseResult `json:"result"`
	Status  DatabaseExportResponseStatus `json:"status"`
	Success bool                         `json:"success"`
	Type    DatabaseExportResponseType   `json:"type"`
	JSON    databaseExportResponseJSON   `json:"-"`
}

// databaseExportResponseJSON contains the JSON metadata for the struct
// [DatabaseExportResponse]
type databaseExportResponseJSON struct {
	AtBookmark  apijson.Field
	Error       apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseExportResponseJSON) RawJSON() string {
	return r.raw
}

// Only present when status = 'complete'
type DatabaseExportResponseResult struct {
	// The generated SQL filename.
	Filename string `json:"filename"`
	// The URL to download the exported SQL. Available for one hour.
	SignedURL string                           `json:"signed_url"`
	JSON      databaseExportResponseResultJSON `json:"-"`
}

// databaseExportResponseResultJSON contains the JSON metadata for the struct
// [DatabaseExportResponseResult]
type databaseExportResponseResultJSON struct {
	Filename    apijson.Field
	SignedURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseExportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseExportResponseResultJSON) RawJSON() string {
	return r.raw
}

type DatabaseExportResponseStatus string

const (
	DatabaseExportResponseStatusComplete DatabaseExportResponseStatus = "complete"
	DatabaseExportResponseStatusError    DatabaseExportResponseStatus = "error"
)

func (r DatabaseExportResponseStatus) IsKnown() bool {
	switch r {
	case DatabaseExportResponseStatusComplete, DatabaseExportResponseStatusError:
		return true
	}
	return false
}

type DatabaseExportResponseType string

const (
	DatabaseExportResponseTypeExport DatabaseExportResponseType = "export"
)

func (r DatabaseExportResponseType) IsKnown() bool {
	switch r {
	case DatabaseExportResponseTypeExport:
		return true
	}
	return false
}

type DatabaseImportResponse struct {
	// The current time-travel bookmark for your D1, used to poll for updates. Will not
	// change for the duration of the import. Only returned if an import process is
	// currently running or recently finished.
	AtBookmark string `json:"at_bookmark"`
	// Only present when status = 'error'. Contains the error message that prevented
	// the import from succeeding.
	Error string `json:"error"`
	// Derived from the database ID and etag, to use in avoiding repeated uploads. Only
	// returned when for the 'init' action.
	Filename string `json:"filename"`
	// Logs since the last time you polled
	Messages []string `json:"messages"`
	// Only present when status = 'complete'
	Result  DatabaseImportResponseResult `json:"result"`
	Status  DatabaseImportResponseStatus `json:"status"`
	Success bool                         `json:"success"`
	Type    DatabaseImportResponseType   `json:"type"`
	// The R2 presigned URL to use for uploading. Only returned when for the 'init'
	// action.
	UploadURL string                     `json:"upload_url"`
	JSON      databaseImportResponseJSON `json:"-"`
}

// databaseImportResponseJSON contains the JSON metadata for the struct
// [DatabaseImportResponse]
type databaseImportResponseJSON struct {
	AtBookmark  apijson.Field
	Error       apijson.Field
	Filename    apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Status      apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	UploadURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseImportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseJSON) RawJSON() string {
	return r.raw
}

// Only present when status = 'complete'
type DatabaseImportResponseResult struct {
	// The time-travel bookmark if you need restore your D1 to directly after the
	// import succeeded.
	FinalBookmark string                           `json:"final_bookmark"`
	Meta          DatabaseImportResponseResultMeta `json:"meta"`
	// The total number of queries that were executed during the import.
	NumQueries float64                          `json:"num_queries"`
	JSON       databaseImportResponseResultJSON `json:"-"`
}

// databaseImportResponseResultJSON contains the JSON metadata for the struct
// [DatabaseImportResponseResult]
type databaseImportResponseResultJSON struct {
	FinalBookmark apijson.Field
	Meta          apijson.Field
	NumQueries    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DatabaseImportResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseResultJSON) RawJSON() string {
	return r.raw
}

type DatabaseImportResponseResultMeta struct {
	// Denotes if the database has been altered in some way, like deleting rows.
	ChangedDB bool `json:"changed_db"`
	// Rough indication of how many rows were modified by the query, as provided by
	// SQLite's `sqlite3_total_changes()`.
	Changes float64 `json:"changes"`
	// The duration of the SQL query execution inside the database. Does not include
	// any network communication.
	Duration float64 `json:"duration"`
	// The row ID of the last inserted row in a table with an `INTEGER PRIMARY KEY` as
	// provided by SQLite. Tables created with `WITHOUT ROWID` do not populate this.
	LastRowID float64 `json:"last_row_id"`
	// Number of rows read during the SQL query execution, including indices (not all
	// rows are necessarily returned).
	RowsRead float64 `json:"rows_read"`
	// Number of rows written during the SQL query execution, including indices.
	RowsWritten float64 `json:"rows_written"`
	// Denotes if the query has been handled by the database primary instance.
	ServedByPrimary bool `json:"served_by_primary"`
	// Region location hint of the database instance that handled the query.
	ServedByRegion DatabaseImportResponseResultMetaServedByRegion `json:"served_by_region"`
	// Size of the database after the query committed, in bytes.
	SizeAfter float64 `json:"size_after"`
	// Various durations for the query.
	Timings DatabaseImportResponseResultMetaTimings `json:"timings"`
	JSON    databaseImportResponseResultMetaJSON    `json:"-"`
}

// databaseImportResponseResultMetaJSON contains the JSON metadata for the struct
// [DatabaseImportResponseResultMeta]
type databaseImportResponseResultMetaJSON struct {
	ChangedDB       apijson.Field
	Changes         apijson.Field
	Duration        apijson.Field
	LastRowID       apijson.Field
	RowsRead        apijson.Field
	RowsWritten     apijson.Field
	ServedByPrimary apijson.Field
	ServedByRegion  apijson.Field
	SizeAfter       apijson.Field
	Timings         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DatabaseImportResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseResultMetaJSON) RawJSON() string {
	return r.raw
}

// Region location hint of the database instance that handled the query.
type DatabaseImportResponseResultMetaServedByRegion string

const (
	DatabaseImportResponseResultMetaServedByRegionWnam DatabaseImportResponseResultMetaServedByRegion = "WNAM"
	DatabaseImportResponseResultMetaServedByRegionEnam DatabaseImportResponseResultMetaServedByRegion = "ENAM"
	DatabaseImportResponseResultMetaServedByRegionWeur DatabaseImportResponseResultMetaServedByRegion = "WEUR"
	DatabaseImportResponseResultMetaServedByRegionEeur DatabaseImportResponseResultMetaServedByRegion = "EEUR"
	DatabaseImportResponseResultMetaServedByRegionApac DatabaseImportResponseResultMetaServedByRegion = "APAC"
	DatabaseImportResponseResultMetaServedByRegionOc   DatabaseImportResponseResultMetaServedByRegion = "OC"
)

func (r DatabaseImportResponseResultMetaServedByRegion) IsKnown() bool {
	switch r {
	case DatabaseImportResponseResultMetaServedByRegionWnam, DatabaseImportResponseResultMetaServedByRegionEnam, DatabaseImportResponseResultMetaServedByRegionWeur, DatabaseImportResponseResultMetaServedByRegionEeur, DatabaseImportResponseResultMetaServedByRegionApac, DatabaseImportResponseResultMetaServedByRegionOc:
		return true
	}
	return false
}

// Various durations for the query.
type DatabaseImportResponseResultMetaTimings struct {
	// The duration of the SQL query execution inside the database. Does not include
	// any network communication.
	SqlDurationMs float64                                     `json:"sql_duration_ms"`
	JSON          databaseImportResponseResultMetaTimingsJSON `json:"-"`
}

// databaseImportResponseResultMetaTimingsJSON contains the JSON metadata for the
// struct [DatabaseImportResponseResultMetaTimings]
type databaseImportResponseResultMetaTimingsJSON struct {
	SqlDurationMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DatabaseImportResponseResultMetaTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseResultMetaTimingsJSON) RawJSON() string {
	return r.raw
}

type DatabaseImportResponseStatus string

const (
	DatabaseImportResponseStatusComplete DatabaseImportResponseStatus = "complete"
	DatabaseImportResponseStatusError    DatabaseImportResponseStatus = "error"
)

func (r DatabaseImportResponseStatus) IsKnown() bool {
	switch r {
	case DatabaseImportResponseStatusComplete, DatabaseImportResponseStatusError:
		return true
	}
	return false
}

type DatabaseImportResponseType string

const (
	DatabaseImportResponseTypeImport DatabaseImportResponseType = "import"
)

func (r DatabaseImportResponseType) IsKnown() bool {
	switch r {
	case DatabaseImportResponseTypeImport:
		return true
	}
	return false
}

type DatabaseRawResponse struct {
	Meta    DatabaseRawResponseMeta    `json:"meta"`
	Results DatabaseRawResponseResults `json:"results"`
	Success bool                       `json:"success"`
	JSON    databaseRawResponseJSON    `json:"-"`
}

// databaseRawResponseJSON contains the JSON metadata for the struct
// [DatabaseRawResponse]
type databaseRawResponseJSON struct {
	Meta        apijson.Field
	Results     apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseRawResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseJSON) RawJSON() string {
	return r.raw
}

type DatabaseRawResponseMeta struct {
	// Denotes if the database has been altered in some way, like deleting rows.
	ChangedDB bool `json:"changed_db"`
	// Rough indication of how many rows were modified by the query, as provided by
	// SQLite's `sqlite3_total_changes()`.
	Changes float64 `json:"changes"`
	// The duration of the SQL query execution inside the database. Does not include
	// any network communication.
	Duration float64 `json:"duration"`
	// The row ID of the last inserted row in a table with an `INTEGER PRIMARY KEY` as
	// provided by SQLite. Tables created with `WITHOUT ROWID` do not populate this.
	LastRowID float64 `json:"last_row_id"`
	// Number of rows read during the SQL query execution, including indices (not all
	// rows are necessarily returned).
	RowsRead float64 `json:"rows_read"`
	// Number of rows written during the SQL query execution, including indices.
	RowsWritten float64 `json:"rows_written"`
	// Denotes if the query has been handled by the database primary instance.
	ServedByPrimary bool `json:"served_by_primary"`
	// Region location hint of the database instance that handled the query.
	ServedByRegion DatabaseRawResponseMetaServedByRegion `json:"served_by_region"`
	// Size of the database after the query committed, in bytes.
	SizeAfter float64 `json:"size_after"`
	// Various durations for the query.
	Timings DatabaseRawResponseMetaTimings `json:"timings"`
	JSON    databaseRawResponseMetaJSON    `json:"-"`
}

// databaseRawResponseMetaJSON contains the JSON metadata for the struct
// [DatabaseRawResponseMeta]
type databaseRawResponseMetaJSON struct {
	ChangedDB       apijson.Field
	Changes         apijson.Field
	Duration        apijson.Field
	LastRowID       apijson.Field
	RowsRead        apijson.Field
	RowsWritten     apijson.Field
	ServedByPrimary apijson.Field
	ServedByRegion  apijson.Field
	SizeAfter       apijson.Field
	Timings         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *DatabaseRawResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Region location hint of the database instance that handled the query.
type DatabaseRawResponseMetaServedByRegion string

const (
	DatabaseRawResponseMetaServedByRegionWnam DatabaseRawResponseMetaServedByRegion = "WNAM"
	DatabaseRawResponseMetaServedByRegionEnam DatabaseRawResponseMetaServedByRegion = "ENAM"
	DatabaseRawResponseMetaServedByRegionWeur DatabaseRawResponseMetaServedByRegion = "WEUR"
	DatabaseRawResponseMetaServedByRegionEeur DatabaseRawResponseMetaServedByRegion = "EEUR"
	DatabaseRawResponseMetaServedByRegionApac DatabaseRawResponseMetaServedByRegion = "APAC"
	DatabaseRawResponseMetaServedByRegionOc   DatabaseRawResponseMetaServedByRegion = "OC"
)

func (r DatabaseRawResponseMetaServedByRegion) IsKnown() bool {
	switch r {
	case DatabaseRawResponseMetaServedByRegionWnam, DatabaseRawResponseMetaServedByRegionEnam, DatabaseRawResponseMetaServedByRegionWeur, DatabaseRawResponseMetaServedByRegionEeur, DatabaseRawResponseMetaServedByRegionApac, DatabaseRawResponseMetaServedByRegionOc:
		return true
	}
	return false
}

// Various durations for the query.
type DatabaseRawResponseMetaTimings struct {
	// The duration of the SQL query execution inside the database. Does not include
	// any network communication.
	SqlDurationMs float64                            `json:"sql_duration_ms"`
	JSON          databaseRawResponseMetaTimingsJSON `json:"-"`
}

// databaseRawResponseMetaTimingsJSON contains the JSON metadata for the struct
// [DatabaseRawResponseMetaTimings]
type databaseRawResponseMetaTimingsJSON struct {
	SqlDurationMs apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DatabaseRawResponseMetaTimings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseMetaTimingsJSON) RawJSON() string {
	return r.raw
}

type DatabaseRawResponseResults struct {
	Columns []string                       `json:"columns"`
	Rows    [][]interface{}                `json:"rows"`
	JSON    databaseRawResponseResultsJSON `json:"-"`
}

// databaseRawResponseResultsJSON contains the JSON metadata for the struct
// [DatabaseRawResponseResults]
type databaseRawResponseResultsJSON struct {
	Columns     apijson.Field
	Rows        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseRawResponseResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseRawResponseResultsJSON) RawJSON() string {
	return r.raw
}

type DatabaseNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// D1 database name.
	Name param.Field[string] `json:"name,required"`
	// Specify the region to create the D1 primary, if available. If this option is
	// omitted, the D1 will be created as close as possible to the current user.
	PrimaryLocationHint param.Field[DatabaseNewParamsPrimaryLocationHint] `json:"primary_location_hint"`
}

func (r DatabaseNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify the region to create the D1 primary, if available. If this option is
// omitted, the D1 will be created as close as possible to the current user.
type DatabaseNewParamsPrimaryLocationHint string

const (
	DatabaseNewParamsPrimaryLocationHintWnam DatabaseNewParamsPrimaryLocationHint = "wnam"
	DatabaseNewParamsPrimaryLocationHintEnam DatabaseNewParamsPrimaryLocationHint = "enam"
	DatabaseNewParamsPrimaryLocationHintWeur DatabaseNewParamsPrimaryLocationHint = "weur"
	DatabaseNewParamsPrimaryLocationHintEeur DatabaseNewParamsPrimaryLocationHint = "eeur"
	DatabaseNewParamsPrimaryLocationHintApac DatabaseNewParamsPrimaryLocationHint = "apac"
	DatabaseNewParamsPrimaryLocationHintOc   DatabaseNewParamsPrimaryLocationHint = "oc"
)

func (r DatabaseNewParamsPrimaryLocationHint) IsKnown() bool {
	switch r {
	case DatabaseNewParamsPrimaryLocationHintWnam, DatabaseNewParamsPrimaryLocationHintEnam, DatabaseNewParamsPrimaryLocationHintWeur, DatabaseNewParamsPrimaryLocationHintEeur, DatabaseNewParamsPrimaryLocationHintApac, DatabaseNewParamsPrimaryLocationHintOc:
		return true
	}
	return false
}

type DatabaseNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The details of the D1 database.
	Result D1 `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseNewResponseEnvelopeJSON    `json:"-"`
}

// databaseNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseNewResponseEnvelope]
type databaseNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseNewResponseEnvelopeSuccess bool

const (
	DatabaseNewResponseEnvelopeSuccessTrue DatabaseNewResponseEnvelopeSuccess = true
)

func (r DatabaseNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Configuration for D1 read replication.
	ReadReplication param.Field[DatabaseUpdateParamsReadReplication] `json:"read_replication,required"`
}

func (r DatabaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for D1 read replication.
type DatabaseUpdateParamsReadReplication struct {
	// The read replication mode for the database. Use 'auto' to create replicas and
	// allow D1 automatically place them around the world, or 'disabled' to not use any
	// database replicas (it can take a few hours for all replicas to be deleted).
	Mode param.Field[DatabaseUpdateParamsReadReplicationMode] `json:"mode,required"`
}

func (r DatabaseUpdateParamsReadReplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The read replication mode for the database. Use 'auto' to create replicas and
// allow D1 automatically place them around the world, or 'disabled' to not use any
// database replicas (it can take a few hours for all replicas to be deleted).
type DatabaseUpdateParamsReadReplicationMode string

const (
	DatabaseUpdateParamsReadReplicationModeAuto     DatabaseUpdateParamsReadReplicationMode = "auto"
	DatabaseUpdateParamsReadReplicationModeDisabled DatabaseUpdateParamsReadReplicationMode = "disabled"
)

func (r DatabaseUpdateParamsReadReplicationMode) IsKnown() bool {
	switch r {
	case DatabaseUpdateParamsReadReplicationModeAuto, DatabaseUpdateParamsReadReplicationModeDisabled:
		return true
	}
	return false
}

type DatabaseUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The details of the D1 database.
	Result D1 `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseUpdateResponseEnvelopeJSON    `json:"-"`
}

// databaseUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseUpdateResponseEnvelope]
type databaseUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseUpdateResponseEnvelopeSuccess bool

const (
	DatabaseUpdateResponseEnvelopeSuccessTrue DatabaseUpdateResponseEnvelopeSuccess = true
)

func (r DatabaseUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// a database name to search for.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [DatabaseListParams]'s query parameters as `url.Values`.
func (r DatabaseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DatabaseDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DatabaseDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   DatabaseDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success DatabaseDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseDeleteResponseEnvelopeJSON    `json:"-"`
}

// databaseDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseDeleteResponseEnvelope]
type databaseDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseDeleteResponseEnvelopeSuccess bool

const (
	DatabaseDeleteResponseEnvelopeSuccessTrue DatabaseDeleteResponseEnvelopeSuccess = true
)

func (r DatabaseDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseEditParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Configuration for D1 read replication.
	ReadReplication param.Field[DatabaseEditParamsReadReplication] `json:"read_replication"`
}

func (r DatabaseEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for D1 read replication.
type DatabaseEditParamsReadReplication struct {
	// The read replication mode for the database. Use 'auto' to create replicas and
	// allow D1 automatically place them around the world, or 'disabled' to not use any
	// database replicas (it can take a few hours for all replicas to be deleted).
	Mode param.Field[DatabaseEditParamsReadReplicationMode] `json:"mode,required"`
}

func (r DatabaseEditParamsReadReplication) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The read replication mode for the database. Use 'auto' to create replicas and
// allow D1 automatically place them around the world, or 'disabled' to not use any
// database replicas (it can take a few hours for all replicas to be deleted).
type DatabaseEditParamsReadReplicationMode string

const (
	DatabaseEditParamsReadReplicationModeAuto     DatabaseEditParamsReadReplicationMode = "auto"
	DatabaseEditParamsReadReplicationModeDisabled DatabaseEditParamsReadReplicationMode = "disabled"
)

func (r DatabaseEditParamsReadReplicationMode) IsKnown() bool {
	switch r {
	case DatabaseEditParamsReadReplicationModeAuto, DatabaseEditParamsReadReplicationModeDisabled:
		return true
	}
	return false
}

type DatabaseEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The details of the D1 database.
	Result D1 `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseEditResponseEnvelopeJSON    `json:"-"`
}

// databaseEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseEditResponseEnvelope]
type databaseEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseEditResponseEnvelopeSuccess bool

const (
	DatabaseEditResponseEnvelopeSuccessTrue DatabaseEditResponseEnvelopeSuccess = true
)

func (r DatabaseEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseExportParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Specifies that you will poll this endpoint until the export completes
	OutputFormat param.Field[DatabaseExportParamsOutputFormat] `json:"output_format,required"`
	// To poll an in-progress export, provide the current bookmark (returned by your
	// first polling response)
	CurrentBookmark param.Field[string]                          `json:"current_bookmark"`
	DumpOptions     param.Field[DatabaseExportParamsDumpOptions] `json:"dump_options"`
}

func (r DatabaseExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies that you will poll this endpoint until the export completes
type DatabaseExportParamsOutputFormat string

const (
	DatabaseExportParamsOutputFormatPolling DatabaseExportParamsOutputFormat = "polling"
)

func (r DatabaseExportParamsOutputFormat) IsKnown() bool {
	switch r {
	case DatabaseExportParamsOutputFormatPolling:
		return true
	}
	return false
}

type DatabaseExportParamsDumpOptions struct {
	// Export only the table definitions, not their contents
	NoData param.Field[bool] `json:"no_data"`
	// Export only each table's contents, not its definition
	NoSchema param.Field[bool] `json:"no_schema"`
	// Filter the export to just one or more tables. Passing an empty array is the same
	// as not passing anything and means: export all tables.
	Tables param.Field[[]string] `json:"tables"`
}

func (r DatabaseExportParamsDumpOptions) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseExportResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   DatabaseExportResponse `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseExportResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseExportResponseEnvelopeJSON    `json:"-"`
}

// databaseExportResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseExportResponseEnvelope]
type databaseExportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseExportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseExportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseExportResponseEnvelopeSuccess bool

const (
	DatabaseExportResponseEnvelopeSuccessTrue DatabaseExportResponseEnvelopeSuccess = true
)

func (r DatabaseExportResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseExportResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DatabaseGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The details of the D1 database.
	Result D1 `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseGetResponseEnvelopeJSON    `json:"-"`
}

// databaseGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseGetResponseEnvelope]
type databaseGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseGetResponseEnvelopeSuccess bool

const (
	DatabaseGetResponseEnvelopeSuccessTrue DatabaseGetResponseEnvelopeSuccess = true
)

func (r DatabaseGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseImportParams struct {
	// Account identifier tag.
	AccountID param.Field[string]           `path:"account_id,required"`
	Body      DatabaseImportParamsBodyUnion `json:"body,required"`
}

func (r DatabaseImportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DatabaseImportParamsBody struct {
	// Indicates you have a new SQL file to upload.
	Action param.Field[DatabaseImportParamsBodyAction] `json:"action,required"`
	// This identifies the currently-running import, checking its status.
	CurrentBookmark param.Field[string] `json:"current_bookmark"`
	// Required when action is 'init' or 'ingest'. An md5 hash of the file you're
	// uploading. Used to check if it already exists, and validate its contents before
	// ingesting.
	Etag param.Field[string] `json:"etag"`
	// The filename you have successfully uploaded.
	Filename param.Field[string] `json:"filename"`
}

func (r DatabaseImportParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatabaseImportParamsBody) implementsDatabaseImportParamsBodyUnion() {}

// Satisfied by [d1.DatabaseImportParamsBodyInit],
// [d1.DatabaseImportParamsBodyIngest], [d1.DatabaseImportParamsBodyPoll],
// [DatabaseImportParamsBody].
type DatabaseImportParamsBodyUnion interface {
	implementsDatabaseImportParamsBodyUnion()
}

type DatabaseImportParamsBodyInit struct {
	// Indicates you have a new SQL file to upload.
	Action param.Field[DatabaseImportParamsBodyInitAction] `json:"action,required"`
	// Required when action is 'init' or 'ingest'. An md5 hash of the file you're
	// uploading. Used to check if it already exists, and validate its contents before
	// ingesting.
	Etag param.Field[string] `json:"etag,required"`
}

func (r DatabaseImportParamsBodyInit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatabaseImportParamsBodyInit) implementsDatabaseImportParamsBodyUnion() {}

// Indicates you have a new SQL file to upload.
type DatabaseImportParamsBodyInitAction string

const (
	DatabaseImportParamsBodyInitActionInit DatabaseImportParamsBodyInitAction = "init"
)

func (r DatabaseImportParamsBodyInitAction) IsKnown() bool {
	switch r {
	case DatabaseImportParamsBodyInitActionInit:
		return true
	}
	return false
}

type DatabaseImportParamsBodyIngest struct {
	// Indicates you've finished uploading to tell the D1 to start consuming it
	Action param.Field[DatabaseImportParamsBodyIngestAction] `json:"action,required"`
	// An md5 hash of the file you're uploading. Used to check if it already exists,
	// and validate its contents before ingesting.
	Etag param.Field[string] `json:"etag,required"`
	// The filename you have successfully uploaded.
	Filename param.Field[string] `json:"filename,required"`
}

func (r DatabaseImportParamsBodyIngest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatabaseImportParamsBodyIngest) implementsDatabaseImportParamsBodyUnion() {}

// Indicates you've finished uploading to tell the D1 to start consuming it
type DatabaseImportParamsBodyIngestAction string

const (
	DatabaseImportParamsBodyIngestActionIngest DatabaseImportParamsBodyIngestAction = "ingest"
)

func (r DatabaseImportParamsBodyIngestAction) IsKnown() bool {
	switch r {
	case DatabaseImportParamsBodyIngestActionIngest:
		return true
	}
	return false
}

type DatabaseImportParamsBodyPoll struct {
	// Indicates you've finished uploading to tell the D1 to start consuming it
	Action param.Field[DatabaseImportParamsBodyPollAction] `json:"action,required"`
	// This identifies the currently-running import, checking its status.
	CurrentBookmark param.Field[string] `json:"current_bookmark,required"`
}

func (r DatabaseImportParamsBodyPoll) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r DatabaseImportParamsBodyPoll) implementsDatabaseImportParamsBodyUnion() {}

// Indicates you've finished uploading to tell the D1 to start consuming it
type DatabaseImportParamsBodyPollAction string

const (
	DatabaseImportParamsBodyPollActionPoll DatabaseImportParamsBodyPollAction = "poll"
)

func (r DatabaseImportParamsBodyPollAction) IsKnown() bool {
	switch r {
	case DatabaseImportParamsBodyPollActionPoll:
		return true
	}
	return false
}

// Indicates you have a new SQL file to upload.
type DatabaseImportParamsBodyAction string

const (
	DatabaseImportParamsBodyActionInit   DatabaseImportParamsBodyAction = "init"
	DatabaseImportParamsBodyActionIngest DatabaseImportParamsBodyAction = "ingest"
	DatabaseImportParamsBodyActionPoll   DatabaseImportParamsBodyAction = "poll"
)

func (r DatabaseImportParamsBodyAction) IsKnown() bool {
	switch r {
	case DatabaseImportParamsBodyActionInit, DatabaseImportParamsBodyActionIngest, DatabaseImportParamsBodyActionPoll:
		return true
	}
	return false
}

type DatabaseImportResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   DatabaseImportResponse `json:"result,required"`
	// Whether the API call was successful
	Success DatabaseImportResponseEnvelopeSuccess `json:"success,required"`
	JSON    databaseImportResponseEnvelopeJSON    `json:"-"`
}

// databaseImportResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatabaseImportResponseEnvelope]
type databaseImportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatabaseImportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r databaseImportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DatabaseImportResponseEnvelopeSuccess bool

const (
	DatabaseImportResponseEnvelopeSuccessTrue DatabaseImportResponseEnvelopeSuccess = true
)

func (r DatabaseImportResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DatabaseImportResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DatabaseQueryParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Your SQL query. Supports multiple statements, joined by semicolons, which will
	// be executed as a batch.
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r DatabaseQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatabaseRawParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Your SQL query. Supports multiple statements, joined by semicolons, which will
	// be executed as a batch.
	Sql    param.Field[string]   `json:"sql,required"`
	Params param.Field[[]string] `json:"params"`
}

func (r DatabaseRawParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
