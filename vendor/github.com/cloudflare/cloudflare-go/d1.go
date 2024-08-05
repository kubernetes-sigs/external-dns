package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingDatabaseID = fmt.Errorf("required missing database ID")
)

type D1Database struct {
	Name      string     `json:"name"`
	NumTables int        `json:"num_tables"`
	UUID      string     `json:"uuid"`
	Version   string     `json:"version"`
	CreatedAt *time.Time `json:"created_at"`
	FileSize  int64      `json:"file_size"`
}

type ListD1DatabasesParams struct {
	Name string `url:"name,omitempty"`
	ResultInfo
}

type ListD1Response struct {
	Result []D1Database `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type CreateD1DatabaseParams struct {
	Name string `json:"name"`
}

type D1DatabaseResponse struct {
	Result D1Database `json:"result"`
	Response
}

type QueryD1DatabaseParams struct {
	DatabaseID string   `json:"-"`
	SQL        string   `json:"sql"`
	Parameters []string `json:"params"`
}

type D1DatabaseMetadata struct {
	ChangedDB   *bool   `json:"changed_db,omitempty"`
	Changes     int     `json:"changes"`
	Duration    float64 `json:"duration"`
	LastRowID   int     `json:"last_row_id"`
	RowsRead    int     `json:"rows_read"`
	RowsWritten int     `json:"rows_written"`
	SizeAfter   int     `json:"size_after"`
}

type D1Result struct {
	Success *bool              `json:"success"`
	Results []map[string]any   `json:"results"`
	Meta    D1DatabaseMetadata `json:"meta"`
}

type QueryD1Response struct {
	Result []D1Result `json:"result"`
	Response
}

// ListD1Databases returns all databases for an account.
//
// API reference: https://developers.cloudflare.com/api/operations/cloudflare-d1-list-databases
func (api *API) ListD1Databases(ctx context.Context, rc *ResourceContainer, params ListD1DatabasesParams) ([]D1Database, *ResultInfo, error) {
	if rc.Identifier == "" {
		return []D1Database{}, &ResultInfo{}, ErrMissingAccountID
	}
	baseURL := fmt.Sprintf("/accounts/%s/d1/database", rc.Identifier)
	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = 100
	}

	if params.Page < 1 {
		params.Page = 1
	}
	var databases []D1Database
	var r ListD1Response
	for {
		uri := buildURI(baseURL, params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []D1Database{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []D1Database{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		databases = append(databases, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return databases, &r.ResultInfo, nil
}

// CreateD1Database creates a new database for an account.
//
// API reference: https://developers.cloudflare.com/api/operations/cloudflare-d1-create-database
func (api *API) CreateD1Database(ctx context.Context, rc *ResourceContainer, params CreateD1DatabaseParams) (D1Database, error) {
	if rc.Identifier == "" {
		return D1Database{}, ErrMissingAccountID
	}
	uri := fmt.Sprintf("/accounts/%s/d1/database", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return D1Database{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r D1DatabaseResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return D1Database{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DeleteD1Database deletes a database for an account.
//
// API reference: https://developers.cloudflare.com/api/operations/cloudflare-d1-delete-database
func (api *API) DeleteD1Database(ctx context.Context, rc *ResourceContainer, databaseID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}
	if databaseID == "" {
		return ErrMissingDatabaseID
	}
	uri := fmt.Sprintf("/accounts/%s/d1/database/%s", rc.Identifier, databaseID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}
	return nil
}

// GetD1Database returns a database for an account.
//
// API reference: https://developers.cloudflare.com/api/operations/cloudflare-d1-get-database
func (api *API) GetD1Database(ctx context.Context, rc *ResourceContainer, databaseID string) (D1Database, error) {
	if rc.Identifier == "" {
		return D1Database{}, ErrMissingAccountID
	}
	uri := fmt.Sprintf("/accounts/%s/d1/database/%s", rc.Identifier, databaseID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return D1Database{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r D1DatabaseResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return D1Database{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// QueryD1Database queries a database for an account.
//
// API reference: https://developers.cloudflare.com/api/operations/cloudflare-d1-query-database
func (api *API) QueryD1Database(ctx context.Context, rc *ResourceContainer, params QueryD1DatabaseParams) ([]D1Result, error) {
	if rc.Identifier == "" {
		return []D1Result{}, ErrMissingAccountID
	}
	if params.DatabaseID == "" {
		return []D1Result{}, ErrMissingDatabaseID
	}
	uri := fmt.Sprintf("/accounts/%s/d1/database/%s/query", rc.Identifier, params.DatabaseID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return []D1Result{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r QueryD1Response
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []D1Result{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}
