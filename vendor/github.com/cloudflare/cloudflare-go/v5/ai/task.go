// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// TaskService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r *TaskService) {
	r = &TaskService{}
	r.Options = opts
	return
}

// Task Search
func (r *TaskService) List(ctx context.Context, query TaskListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TaskListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/tasks/search", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Task Search
func (r *TaskService) ListAutoPaging(ctx context.Context, query TaskListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TaskListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

type TaskListResponse = interface{}

type TaskListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
