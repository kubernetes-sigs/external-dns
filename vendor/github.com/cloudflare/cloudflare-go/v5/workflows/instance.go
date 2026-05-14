// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workflows

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// InstanceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInstanceService] method instead.
type InstanceService struct {
	Options []option.RequestOption
	Status  *InstanceStatusService
	Events  *InstanceEventService
}

// NewInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewInstanceService(opts ...option.RequestOption) (r *InstanceService) {
	r = &InstanceService{}
	r.Options = opts
	r.Status = NewInstanceStatusService(opts...)
	r.Events = NewInstanceEventService(opts...)
	return
}

// Create a new workflow instance
func (r *InstanceService) New(ctx context.Context, workflowName string, params InstanceNewParams, opts ...option.RequestOption) (res *InstanceNewResponse, err error) {
	var env InstanceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances", params.AccountID, workflowName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List of workflow instances
func (r *InstanceService) List(ctx context.Context, workflowName string, params InstanceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[InstanceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances", params.AccountID, workflowName)
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

// List of workflow instances
func (r *InstanceService) ListAutoPaging(ctx context.Context, workflowName string, params InstanceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[InstanceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, workflowName, params, opts...))
}

// Batch create new Workflow instances
func (r *InstanceService) Bulk(ctx context.Context, workflowName string, params InstanceBulkParams, opts ...option.RequestOption) (res *pagination.SinglePage[InstanceBulkResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/batch", params.AccountID, workflowName)
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

// Batch create new Workflow instances
func (r *InstanceService) BulkAutoPaging(ctx context.Context, workflowName string, params InstanceBulkParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[InstanceBulkResponse] {
	return pagination.NewSinglePageAutoPager(r.Bulk(ctx, workflowName, params, opts...))
}

// Get logs and status from instance
func (r *InstanceService) Get(ctx context.Context, workflowName string, instanceID string, query InstanceGetParams, opts ...option.RequestOption) (res *InstanceGetResponse, err error) {
	var env InstanceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if workflowName == "" {
		err = errors.New("missing required workflow_name parameter")
		return
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workflows/%s/instances/%s", query.AccountID, workflowName, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type InstanceNewResponse struct {
	ID         string                    `json:"id,required"`
	Status     InstanceNewResponseStatus `json:"status,required"`
	VersionID  string                    `json:"version_id,required" format:"uuid"`
	WorkflowID string                    `json:"workflow_id,required" format:"uuid"`
	JSON       instanceNewResponseJSON   `json:"-"`
}

// instanceNewResponseJSON contains the JSON metadata for the struct
// [InstanceNewResponse]
type instanceNewResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseStatus string

const (
	InstanceNewResponseStatusQueued          InstanceNewResponseStatus = "queued"
	InstanceNewResponseStatusRunning         InstanceNewResponseStatus = "running"
	InstanceNewResponseStatusPaused          InstanceNewResponseStatus = "paused"
	InstanceNewResponseStatusErrored         InstanceNewResponseStatus = "errored"
	InstanceNewResponseStatusTerminated      InstanceNewResponseStatus = "terminated"
	InstanceNewResponseStatusComplete        InstanceNewResponseStatus = "complete"
	InstanceNewResponseStatusWaitingForPause InstanceNewResponseStatus = "waitingForPause"
	InstanceNewResponseStatusWaiting         InstanceNewResponseStatus = "waiting"
)

func (r InstanceNewResponseStatus) IsKnown() bool {
	switch r {
	case InstanceNewResponseStatusQueued, InstanceNewResponseStatusRunning, InstanceNewResponseStatusPaused, InstanceNewResponseStatusErrored, InstanceNewResponseStatusTerminated, InstanceNewResponseStatusComplete, InstanceNewResponseStatusWaitingForPause, InstanceNewResponseStatusWaiting:
		return true
	}
	return false
}

type InstanceListResponse struct {
	ID         string                     `json:"id,required"`
	CreatedOn  time.Time                  `json:"created_on,required" format:"date-time"`
	EndedOn    time.Time                  `json:"ended_on,required,nullable" format:"date-time"`
	ModifiedOn time.Time                  `json:"modified_on,required" format:"date-time"`
	StartedOn  time.Time                  `json:"started_on,required,nullable" format:"date-time"`
	Status     InstanceListResponseStatus `json:"status,required"`
	VersionID  string                     `json:"version_id,required" format:"uuid"`
	WorkflowID string                     `json:"workflow_id,required" format:"uuid"`
	JSON       instanceListResponseJSON   `json:"-"`
}

// instanceListResponseJSON contains the JSON metadata for the struct
// [InstanceListResponse]
type instanceListResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	EndedOn     apijson.Field
	ModifiedOn  apijson.Field
	StartedOn   apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceListResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceListResponseStatus string

const (
	InstanceListResponseStatusQueued          InstanceListResponseStatus = "queued"
	InstanceListResponseStatusRunning         InstanceListResponseStatus = "running"
	InstanceListResponseStatusPaused          InstanceListResponseStatus = "paused"
	InstanceListResponseStatusErrored         InstanceListResponseStatus = "errored"
	InstanceListResponseStatusTerminated      InstanceListResponseStatus = "terminated"
	InstanceListResponseStatusComplete        InstanceListResponseStatus = "complete"
	InstanceListResponseStatusWaitingForPause InstanceListResponseStatus = "waitingForPause"
	InstanceListResponseStatusWaiting         InstanceListResponseStatus = "waiting"
)

func (r InstanceListResponseStatus) IsKnown() bool {
	switch r {
	case InstanceListResponseStatusQueued, InstanceListResponseStatusRunning, InstanceListResponseStatusPaused, InstanceListResponseStatusErrored, InstanceListResponseStatusTerminated, InstanceListResponseStatusComplete, InstanceListResponseStatusWaitingForPause, InstanceListResponseStatusWaiting:
		return true
	}
	return false
}

type InstanceBulkResponse struct {
	ID         string                     `json:"id,required"`
	Status     InstanceBulkResponseStatus `json:"status,required"`
	VersionID  string                     `json:"version_id,required" format:"uuid"`
	WorkflowID string                     `json:"workflow_id,required" format:"uuid"`
	JSON       instanceBulkResponseJSON   `json:"-"`
}

// instanceBulkResponseJSON contains the JSON metadata for the struct
// [InstanceBulkResponse]
type instanceBulkResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	VersionID   apijson.Field
	WorkflowID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceBulkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceBulkResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceBulkResponseStatus string

const (
	InstanceBulkResponseStatusQueued          InstanceBulkResponseStatus = "queued"
	InstanceBulkResponseStatusRunning         InstanceBulkResponseStatus = "running"
	InstanceBulkResponseStatusPaused          InstanceBulkResponseStatus = "paused"
	InstanceBulkResponseStatusErrored         InstanceBulkResponseStatus = "errored"
	InstanceBulkResponseStatusTerminated      InstanceBulkResponseStatus = "terminated"
	InstanceBulkResponseStatusComplete        InstanceBulkResponseStatus = "complete"
	InstanceBulkResponseStatusWaitingForPause InstanceBulkResponseStatus = "waitingForPause"
	InstanceBulkResponseStatusWaiting         InstanceBulkResponseStatus = "waiting"
)

func (r InstanceBulkResponseStatus) IsKnown() bool {
	switch r {
	case InstanceBulkResponseStatusQueued, InstanceBulkResponseStatusRunning, InstanceBulkResponseStatusPaused, InstanceBulkResponseStatusErrored, InstanceBulkResponseStatusTerminated, InstanceBulkResponseStatusComplete, InstanceBulkResponseStatusWaitingForPause, InstanceBulkResponseStatusWaiting:
		return true
	}
	return false
}

type InstanceGetResponse struct {
	End       time.Time                      `json:"end,required,nullable" format:"date-time"`
	Error     InstanceGetResponseError       `json:"error,required,nullable"`
	Output    InstanceGetResponseOutputUnion `json:"output,required"`
	Params    interface{}                    `json:"params,required"`
	Queued    time.Time                      `json:"queued,required" format:"date-time"`
	Start     time.Time                      `json:"start,required,nullable" format:"date-time"`
	Status    InstanceGetResponseStatus      `json:"status,required"`
	Steps     []InstanceGetResponseStep      `json:"steps,required"`
	Success   bool                           `json:"success,required,nullable"`
	Trigger   InstanceGetResponseTrigger     `json:"trigger,required"`
	VersionID string                         `json:"versionId,required" format:"uuid"`
	JSON      instanceGetResponseJSON        `json:"-"`
}

// instanceGetResponseJSON contains the JSON metadata for the struct
// [InstanceGetResponse]
type instanceGetResponseJSON struct {
	End         apijson.Field
	Error       apijson.Field
	Output      apijson.Field
	Params      apijson.Field
	Queued      apijson.Field
	Start       apijson.Field
	Status      apijson.Field
	Steps       apijson.Field
	Success     apijson.Field
	Trigger     apijson.Field
	VersionID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseError struct {
	Message string                       `json:"message,required"`
	Name    string                       `json:"name,required"`
	JSON    instanceGetResponseErrorJSON `json:"-"`
}

// instanceGetResponseErrorJSON contains the JSON metadata for the struct
// [InstanceGetResponseError]
type instanceGetResponseErrorJSON struct {
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseErrorJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type InstanceGetResponseOutputUnion interface {
	ImplementsInstanceGetResponseOutputUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstanceGetResponseOutputUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type InstanceGetResponseStatus string

const (
	InstanceGetResponseStatusQueued          InstanceGetResponseStatus = "queued"
	InstanceGetResponseStatusRunning         InstanceGetResponseStatus = "running"
	InstanceGetResponseStatusPaused          InstanceGetResponseStatus = "paused"
	InstanceGetResponseStatusErrored         InstanceGetResponseStatus = "errored"
	InstanceGetResponseStatusTerminated      InstanceGetResponseStatus = "terminated"
	InstanceGetResponseStatusComplete        InstanceGetResponseStatus = "complete"
	InstanceGetResponseStatusWaitingForPause InstanceGetResponseStatus = "waitingForPause"
	InstanceGetResponseStatusWaiting         InstanceGetResponseStatus = "waiting"
)

func (r InstanceGetResponseStatus) IsKnown() bool {
	switch r {
	case InstanceGetResponseStatusQueued, InstanceGetResponseStatusRunning, InstanceGetResponseStatusPaused, InstanceGetResponseStatusErrored, InstanceGetResponseStatusTerminated, InstanceGetResponseStatusComplete, InstanceGetResponseStatusWaitingForPause, InstanceGetResponseStatusWaiting:
		return true
	}
	return false
}

type InstanceGetResponseStep struct {
	Type InstanceGetResponseStepsType `json:"type,required"`
	// This field can have the runtime type of
	// [[]InstanceGetResponseStepsObjectAttempt].
	Attempts interface{} `json:"attempts"`
	// This field can have the runtime type of [InstanceGetResponseStepsObjectConfig].
	Config interface{} `json:"config"`
	End    time.Time   `json:"end,nullable" format:"date-time"`
	// This field can have the runtime type of [InstanceGetResponseStepsObjectError].
	Error    interface{} `json:"error"`
	Finished bool        `json:"finished"`
	Name     string      `json:"name"`
	// This field can have the runtime type of [interface{}].
	Output  interface{} `json:"output"`
	Start   time.Time   `json:"start" format:"date-time"`
	Success bool        `json:"success,nullable"`
	// This field can have the runtime type of [InstanceGetResponseStepsObjectTrigger].
	Trigger interface{}                 `json:"trigger"`
	JSON    instanceGetResponseStepJSON `json:"-"`
	union   InstanceGetResponseStepsUnion
}

// instanceGetResponseStepJSON contains the JSON metadata for the struct
// [InstanceGetResponseStep]
type instanceGetResponseStepJSON struct {
	Type        apijson.Field
	Attempts    apijson.Field
	Config      apijson.Field
	End         apijson.Field
	Error       apijson.Field
	Finished    apijson.Field
	Name        apijson.Field
	Output      apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	Trigger     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r instanceGetResponseStepJSON) RawJSON() string {
	return r.raw
}

func (r *InstanceGetResponseStep) UnmarshalJSON(data []byte) (err error) {
	*r = InstanceGetResponseStep{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [InstanceGetResponseStepsUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [InstanceGetResponseStepsObject],
// [InstanceGetResponseStepsObject], [InstanceGetResponseStepsObject],
// [InstanceGetResponseStepsObject].
func (r InstanceGetResponseStep) AsUnion() InstanceGetResponseStepsUnion {
	return r.union
}

// Union satisfied by [InstanceGetResponseStepsObject],
// [InstanceGetResponseStepsObject], [InstanceGetResponseStepsObject] or
// [InstanceGetResponseStepsObject].
type InstanceGetResponseStepsUnion interface {
	implementsInstanceGetResponseStep()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*InstanceGetResponseStepsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceGetResponseStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceGetResponseStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceGetResponseStepsObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(InstanceGetResponseStepsObject{}),
		},
	)
}

type InstanceGetResponseStepsObject struct {
	Attempts []InstanceGetResponseStepsObjectAttempt `json:"attempts,required"`
	Config   InstanceGetResponseStepsObjectConfig    `json:"config,required"`
	End      time.Time                               `json:"end,required,nullable" format:"date-time"`
	Name     string                                  `json:"name,required"`
	Output   interface{}                             `json:"output,required"`
	Start    time.Time                               `json:"start,required" format:"date-time"`
	Success  bool                                    `json:"success,required,nullable"`
	Type     InstanceGetResponseStepsObjectType      `json:"type,required"`
	JSON     instanceGetResponseStepsObjectJSON      `json:"-"`
}

// instanceGetResponseStepsObjectJSON contains the JSON metadata for the struct
// [InstanceGetResponseStepsObject]
type instanceGetResponseStepsObjectJSON struct {
	Attempts    apijson.Field
	Config      apijson.Field
	End         apijson.Field
	Name        apijson.Field
	Output      apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseStepsObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseStepsObjectJSON) RawJSON() string {
	return r.raw
}

func (r InstanceGetResponseStepsObject) implementsInstanceGetResponseStep() {}

type InstanceGetResponseStepsObjectAttempt struct {
	End     time.Time                                   `json:"end,required,nullable" format:"date-time"`
	Error   InstanceGetResponseStepsObjectAttemptsError `json:"error,required,nullable"`
	Start   time.Time                                   `json:"start,required" format:"date-time"`
	Success bool                                        `json:"success,required,nullable"`
	JSON    instanceGetResponseStepsObjectAttemptJSON   `json:"-"`
}

// instanceGetResponseStepsObjectAttemptJSON contains the JSON metadata for the
// struct [InstanceGetResponseStepsObjectAttempt]
type instanceGetResponseStepsObjectAttemptJSON struct {
	End         apijson.Field
	Error       apijson.Field
	Start       apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseStepsObjectAttempt) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseStepsObjectAttemptJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseStepsObjectAttemptsError struct {
	Message string                                          `json:"message,required"`
	Name    string                                          `json:"name,required"`
	JSON    instanceGetResponseStepsObjectAttemptsErrorJSON `json:"-"`
}

// instanceGetResponseStepsObjectAttemptsErrorJSON contains the JSON metadata for
// the struct [InstanceGetResponseStepsObjectAttemptsError]
type instanceGetResponseStepsObjectAttemptsErrorJSON struct {
	Message     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseStepsObjectAttemptsError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseStepsObjectAttemptsErrorJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseStepsObjectConfig struct {
	Retries InstanceGetResponseStepsObjectConfigRetries `json:"retries,required"`
	Timeout interface{}                                 `json:"timeout,required"`
	JSON    instanceGetResponseStepsObjectConfigJSON    `json:"-"`
}

// instanceGetResponseStepsObjectConfigJSON contains the JSON metadata for the
// struct [InstanceGetResponseStepsObjectConfig]
type instanceGetResponseStepsObjectConfigJSON struct {
	Retries     apijson.Field
	Timeout     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseStepsObjectConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseStepsObjectConfigJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseStepsObjectConfigRetries struct {
	Delay   interface{}                                        `json:"delay,required"`
	Limit   float64                                            `json:"limit,required"`
	Backoff InstanceGetResponseStepsObjectConfigRetriesBackoff `json:"backoff"`
	JSON    instanceGetResponseStepsObjectConfigRetriesJSON    `json:"-"`
}

// instanceGetResponseStepsObjectConfigRetriesJSON contains the JSON metadata for
// the struct [InstanceGetResponseStepsObjectConfigRetries]
type instanceGetResponseStepsObjectConfigRetriesJSON struct {
	Delay       apijson.Field
	Limit       apijson.Field
	Backoff     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseStepsObjectConfigRetries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseStepsObjectConfigRetriesJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseStepsObjectConfigRetriesBackoff string

const (
	InstanceGetResponseStepsObjectConfigRetriesBackoffConstant    InstanceGetResponseStepsObjectConfigRetriesBackoff = "constant"
	InstanceGetResponseStepsObjectConfigRetriesBackoffLinear      InstanceGetResponseStepsObjectConfigRetriesBackoff = "linear"
	InstanceGetResponseStepsObjectConfigRetriesBackoffExponential InstanceGetResponseStepsObjectConfigRetriesBackoff = "exponential"
)

func (r InstanceGetResponseStepsObjectConfigRetriesBackoff) IsKnown() bool {
	switch r {
	case InstanceGetResponseStepsObjectConfigRetriesBackoffConstant, InstanceGetResponseStepsObjectConfigRetriesBackoffLinear, InstanceGetResponseStepsObjectConfigRetriesBackoffExponential:
		return true
	}
	return false
}

type InstanceGetResponseStepsObjectType string

const (
	InstanceGetResponseStepsObjectTypeStep InstanceGetResponseStepsObjectType = "step"
)

func (r InstanceGetResponseStepsObjectType) IsKnown() bool {
	switch r {
	case InstanceGetResponseStepsObjectTypeStep:
		return true
	}
	return false
}

type InstanceGetResponseStepsType string

const (
	InstanceGetResponseStepsTypeStep         InstanceGetResponseStepsType = "step"
	InstanceGetResponseStepsTypeSleep        InstanceGetResponseStepsType = "sleep"
	InstanceGetResponseStepsTypeTermination  InstanceGetResponseStepsType = "termination"
	InstanceGetResponseStepsTypeWaitForEvent InstanceGetResponseStepsType = "waitForEvent"
)

func (r InstanceGetResponseStepsType) IsKnown() bool {
	switch r {
	case InstanceGetResponseStepsTypeStep, InstanceGetResponseStepsTypeSleep, InstanceGetResponseStepsTypeTermination, InstanceGetResponseStepsTypeWaitForEvent:
		return true
	}
	return false
}

type InstanceGetResponseTrigger struct {
	Source InstanceGetResponseTriggerSource `json:"source,required"`
	JSON   instanceGetResponseTriggerJSON   `json:"-"`
}

// instanceGetResponseTriggerJSON contains the JSON metadata for the struct
// [InstanceGetResponseTrigger]
type instanceGetResponseTriggerJSON struct {
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseTrigger) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseTriggerJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseTriggerSource string

const (
	InstanceGetResponseTriggerSourceUnknown InstanceGetResponseTriggerSource = "unknown"
	InstanceGetResponseTriggerSourceAPI     InstanceGetResponseTriggerSource = "api"
	InstanceGetResponseTriggerSourceBinding InstanceGetResponseTriggerSource = "binding"
	InstanceGetResponseTriggerSourceEvent   InstanceGetResponseTriggerSource = "event"
	InstanceGetResponseTriggerSourceCron    InstanceGetResponseTriggerSource = "cron"
)

func (r InstanceGetResponseTriggerSource) IsKnown() bool {
	switch r {
	case InstanceGetResponseTriggerSourceUnknown, InstanceGetResponseTriggerSourceAPI, InstanceGetResponseTriggerSourceBinding, InstanceGetResponseTriggerSourceEvent, InstanceGetResponseTriggerSourceCron:
		return true
	}
	return false
}

type InstanceNewParams struct {
	AccountID         param.Field[string]      `path:"account_id,required"`
	InstanceID        param.Field[string]      `json:"instance_id"`
	InstanceRetention param.Field[interface{}] `json:"instance_retention"`
	Params            param.Field[interface{}] `json:"params"`
}

func (r InstanceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceNewResponseEnvelope struct {
	Errors     []InstanceNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []InstanceNewResponseEnvelopeMessages `json:"messages,required"`
	Result     InstanceNewResponse                   `json:"result,required"`
	Success    InstanceNewResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo InstanceNewResponseEnvelopeResultInfo `json:"result_info"`
	JSON       instanceNewResponseEnvelopeJSON       `json:"-"`
}

// instanceNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceNewResponseEnvelope]
type instanceNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseEnvelopeErrors struct {
	Code    float64                               `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    instanceNewResponseEnvelopeErrorsJSON `json:"-"`
}

// instanceNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [InstanceNewResponseEnvelopeErrors]
type instanceNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseEnvelopeMessages struct {
	Code    float64                                 `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    instanceNewResponseEnvelopeMessagesJSON `json:"-"`
}

// instanceNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [InstanceNewResponseEnvelopeMessages]
type instanceNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InstanceNewResponseEnvelopeSuccess bool

const (
	InstanceNewResponseEnvelopeSuccessTrue InstanceNewResponseEnvelopeSuccess = true
)

func (r InstanceNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InstanceNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InstanceNewResponseEnvelopeResultInfo struct {
	Count      float64                                   `json:"count,required"`
	PerPage    float64                                   `json:"per_page,required"`
	TotalCount float64                                   `json:"total_count,required"`
	NextCursor string                                    `json:"next_cursor"`
	Page       float64                                   `json:"page"`
	JSON       instanceNewResponseEnvelopeResultInfoJSON `json:"-"`
}

// instanceNewResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [InstanceNewResponseEnvelopeResultInfo]
type instanceNewResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	NextCursor  apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceNewResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceNewResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type InstanceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// `page` and `cursor` are mutually exclusive, use one or the other.
	Cursor param.Field[string] `query:"cursor"`
	// Accepts ISO 8601 with no timezone offsets and in UTC.
	DateEnd param.Field[time.Time] `query:"date_end" format:"date-time"`
	// Accepts ISO 8601 with no timezone offsets and in UTC.
	DateStart param.Field[time.Time] `query:"date_start" format:"date-time"`
	// `page` and `cursor` are mutually exclusive, use one or the other.
	Page    param.Field[float64]                  `query:"page"`
	PerPage param.Field[float64]                  `query:"per_page"`
	Status  param.Field[InstanceListParamsStatus] `query:"status"`
}

// URLQuery serializes [InstanceListParams]'s query parameters as `url.Values`.
func (r InstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type InstanceListParamsStatus string

const (
	InstanceListParamsStatusQueued          InstanceListParamsStatus = "queued"
	InstanceListParamsStatusRunning         InstanceListParamsStatus = "running"
	InstanceListParamsStatusPaused          InstanceListParamsStatus = "paused"
	InstanceListParamsStatusErrored         InstanceListParamsStatus = "errored"
	InstanceListParamsStatusTerminated      InstanceListParamsStatus = "terminated"
	InstanceListParamsStatusComplete        InstanceListParamsStatus = "complete"
	InstanceListParamsStatusWaitingForPause InstanceListParamsStatus = "waitingForPause"
	InstanceListParamsStatusWaiting         InstanceListParamsStatus = "waiting"
)

func (r InstanceListParamsStatus) IsKnown() bool {
	switch r {
	case InstanceListParamsStatusQueued, InstanceListParamsStatusRunning, InstanceListParamsStatusPaused, InstanceListParamsStatusErrored, InstanceListParamsStatusTerminated, InstanceListParamsStatusComplete, InstanceListParamsStatusWaitingForPause, InstanceListParamsStatusWaiting:
		return true
	}
	return false
}

type InstanceBulkParams struct {
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      []InstanceBulkParamsBody `json:"body"`
}

func (r InstanceBulkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type InstanceBulkParamsBody struct {
	InstanceID        param.Field[string]      `json:"instance_id"`
	InstanceRetention param.Field[interface{}] `json:"instance_retention"`
	Params            param.Field[interface{}] `json:"params"`
}

func (r InstanceBulkParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type InstanceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type InstanceGetResponseEnvelope struct {
	Errors     []InstanceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages   []InstanceGetResponseEnvelopeMessages `json:"messages,required"`
	Result     InstanceGetResponse                   `json:"result,required"`
	Success    InstanceGetResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo InstanceGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       instanceGetResponseEnvelopeJSON       `json:"-"`
}

// instanceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InstanceGetResponseEnvelope]
type instanceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseEnvelopeErrors struct {
	Code    float64                               `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    instanceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// instanceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [InstanceGetResponseEnvelopeErrors]
type instanceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseEnvelopeMessages struct {
	Code    float64                                 `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    instanceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// instanceGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [InstanceGetResponseEnvelopeMessages]
type instanceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InstanceGetResponseEnvelopeSuccess bool

const (
	InstanceGetResponseEnvelopeSuccessTrue InstanceGetResponseEnvelopeSuccess = true
)

func (r InstanceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InstanceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InstanceGetResponseEnvelopeResultInfo struct {
	Count      float64                                   `json:"count,required"`
	PerPage    float64                                   `json:"per_page,required"`
	TotalCount float64                                   `json:"total_count,required"`
	NextCursor string                                    `json:"next_cursor"`
	Page       float64                                   `json:"page"`
	JSON       instanceGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// instanceGetResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [InstanceGetResponseEnvelopeResultInfo]
type instanceGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	NextCursor  apijson.Field
	Page        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InstanceGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r instanceGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
