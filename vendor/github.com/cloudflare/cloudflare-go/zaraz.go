package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type ZarazConfig struct {
	DebugKey      string                   `json:"debugKey"`
	Tools         map[string]ZarazTool     `json:"tools"`
	Triggers      map[string]ZarazTrigger  `json:"triggers"`
	ZarazVersion  int64                    `json:"zarazVersion"`
	Consent       ZarazConsent             `json:"consent,omitempty"`
	DataLayer     *bool                    `json:"dataLayer,omitempty"`
	Dlp           []any                    `json:"dlp,omitempty"`
	HistoryChange *bool                    `json:"historyChange,omitempty"`
	Settings      ZarazConfigSettings      `json:"settings,omitempty"`
	Variables     map[string]ZarazVariable `json:"variables,omitempty"`
}

type ZarazWorker struct {
	EscapedWorkerName string `json:"escapedWorkerName"`
	WorkerTag         string `json:"workerTag"`
	MutableId         string `json:"mutableId,omitempty"`
}
type ZarazConfigSettings struct {
	AutoInjectScript    *bool       `json:"autoInjectScript"`
	InjectIframes       *bool       `json:"injectIframes,omitempty"`
	Ecommerce           *bool       `json:"ecommerce,omitempty"`
	HideQueryParams     *bool       `json:"hideQueryParams,omitempty"`
	HideIpAddress       *bool       `json:"hideIPAddress,omitempty"`
	HideUserAgent       *bool       `json:"hideUserAgent,omitempty"`
	HideExternalReferer *bool       `json:"hideExternalReferer,omitempty"`
	CookieDomain        string      `json:"cookieDomain,omitempty"`
	InitPath            string      `json:"initPath,omitempty"`
	ScriptPath          string      `json:"scriptPath,omitempty"`
	TrackPath           string      `json:"trackPath,omitempty"`
	EventsApiPath       string      `json:"eventsApiPath,omitempty"`
	McRootPath          string      `json:"mcRootPath,omitempty"`
	ContextEnricher     ZarazWorker `json:"contextEnricher,omitempty"`
}

// Deprecated: To be removed pending migration of existing configs.
type ZarazNeoEvent struct {
	BlockingTriggers []string       `json:"blockingTriggers"`
	FiringTriggers   []string       `json:"firingTriggers"`
	Data             map[string]any `json:"data"`
	ActionType       string         `json:"actionType,omitempty"`
}

type ZarazAction struct {
	BlockingTriggers []string       `json:"blockingTriggers"`
	FiringTriggers   []string       `json:"firingTriggers"`
	Data             map[string]any `json:"data"`
	ActionType       string         `json:"actionType,omitempty"`
}

type ZarazToolType string

const (
	ZarazToolLibrary   ZarazToolType = "library"
	ZarazToolComponent ZarazToolType = "component"
	ZarazToolCustomMc  ZarazToolType = "custom-mc"
)

type ZarazTool struct {
	BlockingTriggers []string               `json:"blockingTriggers"`
	Enabled          *bool                  `json:"enabled"`
	DefaultFields    map[string]any         `json:"defaultFields"`
	Name             string                 `json:"name"`
	NeoEvents        []ZarazNeoEvent        `json:"neoEvents"`
	Actions          map[string]ZarazAction `json:"actions"`
	Type             ZarazToolType          `json:"type"`
	DefaultPurpose   string                 `json:"defaultPurpose,omitempty"`
	Library          string                 `json:"library,omitempty"`
	Component        string                 `json:"component,omitempty"`
	Permissions      []string               `json:"permissions"`
	Settings         map[string]any         `json:"settings"`
	Worker           ZarazWorker            `json:"worker,omitempty"`
}

type ZarazTriggerSystem string

const ZarazPageload ZarazTriggerSystem = "pageload"

type ZarazLoadRuleOp string

type ZarazRuleType string

const (
	ZarazClickListener     ZarazRuleType = "clickListener"
	ZarazTimer             ZarazRuleType = "timer"
	ZarazFormSubmission    ZarazRuleType = "formSubmission"
	ZarazVariableMatch     ZarazRuleType = "variableMatch"
	ZarazScrollDepth       ZarazRuleType = "scrollDepth"
	ZarazElementVisibility ZarazRuleType = "elementVisibility"
	ZarazClientEval        ZarazRuleType = "clientEval"
)

type ZarazSelectorType string

const (
	ZarazXPath ZarazSelectorType = "xpath"
	ZarazCSS   ZarazSelectorType = "css"
)

type ZarazRuleSettings struct {
	Type        ZarazSelectorType `json:"type,omitempty"`
	Selector    string            `json:"selector,omitempty"`
	WaitForTags int               `json:"waitForTags,omitempty"`
	Interval    int               `json:"interval,omitempty"`
	Limit       int               `json:"limit,omitempty"`
	Validate    *bool             `json:"validate,omitempty"`
	Variable    string            `json:"variable,omitempty"`
	Match       string            `json:"match,omitempty"`
	Positions   string            `json:"positions,omitempty"`
	Op          ZarazLoadRuleOp   `json:"op,omitempty"`
	Value       string            `json:"value,omitempty"`
}

type ZarazTriggerRule struct {
	Id       string            `json:"id"`
	Match    string            `json:"match,omitempty"`
	Op       ZarazLoadRuleOp   `json:"op,omitempty"`
	Value    string            `json:"value,omitempty"`
	Action   ZarazRuleType     `json:"action"`
	Settings ZarazRuleSettings `json:"settings"`
}

type ZarazTrigger struct {
	Name         string             `json:"name"`
	Description  string             `json:"description,omitempty"`
	LoadRules    []ZarazTriggerRule `json:"loadRules"`
	ExcludeRules []ZarazTriggerRule `json:"excludeRules"`
	ClientRules  []any              `json:"clientRules,omitempty"` // what is this?
	System       ZarazTriggerSystem `json:"system,omitempty"`
}

type ZarazVariableType string

const (
	ZarazVarString ZarazVariableType = "string"
	ZarazVarSecret ZarazVariableType = "secret"
	ZarazVarWorker ZarazVariableType = "worker"
)

type ZarazVariable struct {
	Name  string            `json:"name"`
	Type  ZarazVariableType `json:"type"`
	Value interface{}       `json:"value"`
}

type ZarazButtonTextTranslations struct {
	AcceptAll        map[string]string `json:"accept_all"`
	RejectAll        map[string]string `json:"reject_all"`
	ConfirmMyChoices map[string]string `json:"confirm_my_choices"`
}

type ZarazPurpose struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ZarazPurposeWithTranslations struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Order       int               `json:"order"`
}

type ZarazConsent struct {
	Enabled                               *bool                                   `json:"enabled"`
	ButtonTextTranslations                ZarazButtonTextTranslations             `json:"buttonTextTranslations,omitempty"`
	CompanyEmail                          string                                  `json:"companyEmail,omitempty"`
	CompanyName                           string                                  `json:"companyName,omitempty"`
	CompanyStreetAddress                  string                                  `json:"companyStreetAddress,omitempty"`
	ConsentModalIntroHTML                 string                                  `json:"consentModalIntroHTML,omitempty"`
	ConsentModalIntroHTMLWithTranslations map[string]string                       `json:"consentModalIntroHTMLWithTranslations,omitempty"`
	CookieName                            string                                  `json:"cookieName,omitempty"`
	CustomCSS                             string                                  `json:"customCSS,omitempty"`
	CustomIntroDisclaimerDismissed        *bool                                   `json:"customIntroDisclaimerDismissed,omitempty"`
	DefaultLanguage                       string                                  `json:"defaultLanguage,omitempty"`
	HideModal                             *bool                                   `json:"hideModal,omitempty"`
	Purposes                              map[string]ZarazPurpose                 `json:"purposes,omitempty"`
	PurposesWithTranslations              map[string]ZarazPurposeWithTranslations `json:"purposesWithTranslations,omitempty"`
}

type ZarazConfigResponse struct {
	Result ZarazConfig `json:"result"`
	Response
}

type ZarazWorkflowResponse struct {
	Result string `json:"result"`
	Response
}

type ZarazPublishResponse struct {
	Result string `json:"result"`
	Response
}

type UpdateZarazConfigParams struct {
	DebugKey      string                   `json:"debugKey"`
	Tools         map[string]ZarazTool     `json:"tools"`
	Triggers      map[string]ZarazTrigger  `json:"triggers"`
	ZarazVersion  int64                    `json:"zarazVersion"`
	Consent       ZarazConsent             `json:"consent,omitempty"`
	DataLayer     *bool                    `json:"dataLayer,omitempty"`
	Dlp           []any                    `json:"dlp,omitempty"`
	HistoryChange *bool                    `json:"historyChange,omitempty"`
	Settings      ZarazConfigSettings      `json:"settings,omitempty"`
	Variables     map[string]ZarazVariable `json:"variables,omitempty"`
}

type UpdateZarazWorkflowParams struct {
	Workflow string `json:"workflow"`
}

type PublishZarazConfigParams struct {
	Description string `json:"description"`
}

type ZarazHistoryRecord struct {
	ID          int64      `json:"id,omitempty"`
	UserID      string     `json:"userId,omitempty"`
	Description string     `json:"description,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

type ZarazConfigHistoryListResponse struct {
	Result []ZarazHistoryRecord `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type ListZarazConfigHistoryParams struct {
	ResultInfo
}

type GetZarazConfigsByIdResponse = map[string]interface{}

// listZarazConfigHistoryDefaultPageSize represents the default per_page size of the API.
var listZarazConfigHistoryDefaultPageSize int = 100

func (api *API) GetZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var recordResp ZarazConfigResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) UpdateZarazConfig(ctx context.Context, rc *ResourceContainer, params UpdateZarazConfigParams) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/config", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var updateResp ZarazConfigResponse
	err = json.Unmarshal(res, &updateResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return updateResp, nil
}

func (api *API) GetZarazWorkflow(ctx context.Context, rc *ResourceContainer) (ZarazWorkflowResponse, error) {
	if rc.Identifier == "" {
		return ZarazWorkflowResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/workflow", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazWorkflowResponse{}, err
	}

	var response ZarazWorkflowResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazWorkflowResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) UpdateZarazWorkflow(ctx context.Context, rc *ResourceContainer, params UpdateZarazWorkflowParams) (ZarazWorkflowResponse, error) {
	if rc.Identifier == "" {
		return ZarazWorkflowResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/workflow", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.Workflow)
	if err != nil {
		return ZarazWorkflowResponse{}, err
	}

	var response ZarazWorkflowResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazWorkflowResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) PublishZarazConfig(ctx context.Context, rc *ResourceContainer, params PublishZarazConfigParams) (ZarazPublishResponse, error) {
	if rc.Identifier == "" {
		return ZarazPublishResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/publish", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Description)
	if err != nil {
		return ZarazPublishResponse{}, err
	}

	var response ZarazPublishResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZarazPublishResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response, nil
}

func (api *API) ListZarazConfigHistory(ctx context.Context, rc *ResourceContainer, params ListZarazConfigHistoryParams) ([]ZarazHistoryRecord, *ResultInfo, error) {
	if rc.Identifier == "" {
		return nil, nil, ErrMissingZoneID
	}

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = listZarazConfigHistoryDefaultPageSize
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var records []ZarazHistoryRecord
	var lastResultInfo ResultInfo

	for {
		uri := buildURI(fmt.Sprintf("/zones/%s/settings/zaraz/v2/history", rc.Identifier), params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []ZarazHistoryRecord{}, &ResultInfo{}, err
		}
		var listResponse ZarazConfigHistoryListResponse
		err = json.Unmarshal(res, &listResponse)
		if err != nil {
			return []ZarazHistoryRecord{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		records = append(records, listResponse.Result...)
		lastResultInfo = listResponse.ResultInfo
		params.ResultInfo = listResponse.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}
	return records, &lastResultInfo, nil
}

func (api *API) GetDefaultZarazConfig(ctx context.Context, rc *ResourceContainer) (ZarazConfigResponse, error) {
	if rc.Identifier == "" {
		return ZarazConfigResponse{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/default", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZarazConfigResponse{}, err
	}

	var recordResp ZarazConfigResponse
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return ZarazConfigResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return recordResp, nil
}

func (api *API) ExportZarazConfig(ctx context.Context, rc *ResourceContainer) error {
	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/settings/zaraz/v2/export", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return err
	}

	var recordResp ZarazConfig
	err = json.Unmarshal(res, &recordResp)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}
