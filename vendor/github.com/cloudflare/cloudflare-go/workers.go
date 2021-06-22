package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

	"github.com/pkg/errors"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

// WorkerScriptParams provides a worker script and the associated bindings
type WorkerScriptParams struct {
	Script string

	// Bindings should be a map where the keys are the binding name, and the
	// values are the binding content
	Bindings map[string]WorkerBinding
}

// WorkerRoute is used to map traffic matching a URL pattern to a workers
//
// API reference: https://api.cloudflare.com/#worker-routes-properties
type WorkerRoute struct {
	ID      string `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Enabled bool   `json:"enabled"` // this is deprecated: https://api.cloudflare.com/#worker-filters-deprecated--properties
	Script  string `json:"script,omitempty"`
}

// WorkerRoutesResponse embeds Response struct and slice of WorkerRoutes
type WorkerRoutesResponse struct {
	Response
	Routes []WorkerRoute `json:"result"`
}

// WorkerRouteResponse embeds Response struct and a single WorkerRoute
type WorkerRouteResponse struct {
	Response
	WorkerRoute `json:"result"`
}

// WorkerScript Cloudflare Worker struct with metadata
type WorkerScript struct {
	WorkerMetaData
	Script string `json:"script"`
}

// WorkerMetaData contains worker script information such as size, creation & modification dates
type WorkerMetaData struct {
	ID         string    `json:"id,omitempty"`
	ETAG       string    `json:"etag,omitempty"`
	Size       int       `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

// WorkerListResponse wrapper struct for API response to worker script list API call
type WorkerListResponse struct {
	Response
	WorkerList []WorkerMetaData `json:"result"`
}

// WorkerScriptResponse wrapper struct for API response to worker script calls
type WorkerScriptResponse struct {
	Response
	WorkerScript `json:"result"`
}

// WorkerBindingType represents a particular type of binding
type WorkerBindingType string

func (b WorkerBindingType) String() string {
	return string(b)
}

const (
	// WorkerInheritBindingType is the type for inherited bindings
	WorkerInheritBindingType WorkerBindingType = "inherit"
	// WorkerKvNamespaceBindingType is the type for KV Namespace bindings
	WorkerKvNamespaceBindingType WorkerBindingType = "kv_namespace"
	// WorkerWebAssemblyBindingType is the type for Web Assembly module bindings
	WorkerWebAssemblyBindingType WorkerBindingType = "wasm_module"
	// WorkerSecretTextBindingType is the type for secret text bindings
	WorkerSecretTextBindingType WorkerBindingType = "secret_text"
	// WorkerPlainTextBindingType is the type for plain text bindings
	WorkerPlainTextBindingType WorkerBindingType = "plain_text"
)

// WorkerBindingListItem a struct representing an individual binding in a list of bindings
type WorkerBindingListItem struct {
	Name    string `json:"name"`
	Binding WorkerBinding
}

// WorkerBindingListResponse wrapper struct for API response to worker binding list API call
type WorkerBindingListResponse struct {
	Response
	BindingList []WorkerBindingListItem
}

// Workers supports multiple types of bindings, e.g. KV namespaces or WebAssembly modules, and each type
// of binding will be represented differently in the upload request body. At a high-level, every binding
// will specify metadata, which is a JSON object with the properties "name" and "type". Some types of bindings
// will also have additional metadata properties. For example, KV bindings also specify the KV namespace.
// In addition to the metadata, some binding types may need to include additional data as part of the
// multipart form. For example, WebAssembly bindings will include the contents of the WebAssembly module.

// WorkerBinding is the generic interface implemented by all of
// the various binding types
type WorkerBinding interface {
	Type() WorkerBindingType

	// serialize is responsible for returning the binding metadata as well as an optionally
	// returning a function that can modify the multipart form body. For example, this is used
	// by WebAssembly bindings to add a new part containing the WebAssembly module contents.
	serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error)
}

// workerBindingMeta is the metadata portion of the binding
type workerBindingMeta = map[string]interface{}

// workerBindingBodyWriter allows for a binding to add additional parts to the multipart body
type workerBindingBodyWriter func(*multipart.Writer) error

// WorkerInheritBinding will just persist whatever binding content was previously uploaded
type WorkerInheritBinding struct {
	// Optional parameter that allows for renaming a binding without changing
	// its contents. If `OldName` is empty, the binding name will not be changed.
	OldName string
}

// Type returns the type of the binding
func (b WorkerInheritBinding) Type() WorkerBindingType {
	return WorkerInheritBindingType
}

func (b WorkerInheritBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	meta := workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
	}

	if b.OldName != "" {
		meta["old_name"] = b.OldName
	}

	return meta, nil, nil
}

// WorkerKvNamespaceBinding is a binding to a Workers KV Namespace
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/kv-namespaces/
type WorkerKvNamespaceBinding struct {
	NamespaceID string
}

// Type returns the type of the binding
func (b WorkerKvNamespaceBinding) Type() WorkerBindingType {
	return WorkerKvNamespaceBindingType
}

func (b WorkerKvNamespaceBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.NamespaceID == "" {
		return nil, nil, errors.Errorf(`NamespaceID for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":         bindingName,
		"type":         b.Type(),
		"namespace_id": b.NamespaceID,
	}, nil, nil
}

// WorkerWebAssemblyBinding is a binding to a WebAssembly module
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/webassembly-modules/
type WorkerWebAssemblyBinding struct {
	Module io.Reader
}

// Type returns the type of the binding
func (b WorkerWebAssemblyBinding) Type() WorkerBindingType {
	return WorkerWebAssemblyBindingType
}

func (b WorkerWebAssemblyBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	partName := getRandomPartName()

	bodyWriter := func(mpw *multipart.Writer) error {
		var hdr = textproto.MIMEHeader{}
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, partName))
		hdr.Set("content-type", "application/wasm")
		pw, err := mpw.CreatePart(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(pw, b.Module)
		return err
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"part": partName,
	}, bodyWriter, nil
}

// WorkerPlainTextBinding is a binding to plain text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-plain-text-binding
type WorkerPlainTextBinding struct {
	Text string
}

// Type returns the type of the binding
func (b WorkerPlainTextBinding) Type() WorkerBindingType {
	return WorkerPlainTextBindingType
}

func (b WorkerPlainTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, errors.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// WorkerSecretTextBinding is a binding to secret text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-secret-text-binding
type WorkerSecretTextBinding struct {
	Text string
}

// Type returns the type of the binding
func (b WorkerSecretTextBinding) Type() WorkerBindingType {
	return WorkerSecretTextBindingType
}

func (b WorkerSecretTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, errors.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// Each binding that adds a part to the multipart form body will need
// a unique part name so we just generate a random 128bit hex string
func getRandomPartName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}

// DeleteWorker deletes worker for a zone.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	// if ScriptName is provided we will treat as org request
	if requestParams.ScriptName != "" {
		return api.deleteWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerWithName deletes worker for a zone.
// Sccount must be specified as api option https://godoc.org/github.com/cloudflare/cloudflare-go#UsingAccount
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) deleteWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DownloadWorker fetch raw script content for your worker returns []byte containing worker code js
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) DownloadWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.downloadWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// DownloadWorkerWithName fetch raw script content for your worker returns string containing worker code js
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) downloadWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// ListWorkerBindings returns all the bindings for a particular worker
func (api *API) ListWorkerBindings(requestParams *WorkerRequestParams) (WorkerBindingListResponse, error) {
	if requestParams.ScriptName == "" {
		return WorkerBindingListResponse{}, errors.New("ScriptName is required")
	}
	if api.AccountID == "" {
		return WorkerBindingListResponse{}, errors.New("account ID required")
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings", api.AccountID, requestParams.ScriptName)

	var jsonRes struct {
		Response
		Bindings []workerBindingMeta `json:"result"`
	}
	var r WorkerBindingListResponse
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &jsonRes)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}

	r = WorkerBindingListResponse{
		Response:    jsonRes.Response,
		BindingList: make([]WorkerBindingListItem, 0, len(jsonRes.Bindings)),
	}
	for _, jsonBinding := range jsonRes.Bindings {
		name, ok := jsonBinding["name"].(string)
		if !ok {
			return r, errors.Errorf("Binding missing name %v", jsonBinding)
		}
		bType, ok := jsonBinding["type"].(string)
		if !ok {
			return r, errors.Errorf("Binding missing type %v", jsonBinding)
		}
		bindingListItem := WorkerBindingListItem{
			Name: name,
		}

		switch WorkerBindingType(bType) {
		case WorkerKvNamespaceBindingType:
			namespaceID := jsonBinding["namespace_id"].(string)
			bindingListItem.Binding = WorkerKvNamespaceBinding{
				NamespaceID: namespaceID,
			}
		case WorkerWebAssemblyBindingType:
			bindingListItem.Binding = WorkerWebAssemblyBinding{
				Module: &bindingContentReader{
					api:           api,
					requestParams: requestParams,
					bindingName:   name,
				},
			}
		case WorkerPlainTextBindingType:
			text := jsonBinding["text"].(string)
			bindingListItem.Binding = WorkerPlainTextBinding{
				Text: text,
			}
		case WorkerSecretTextBindingType:
			bindingListItem.Binding = WorkerSecretTextBinding{}
		default:
			bindingListItem.Binding = WorkerInheritBinding{}
		}
		r.BindingList = append(r.BindingList, bindingListItem)
	}

	return r, nil
}

// bindingContentReader is an io.Reader that will lazily load the
// raw bytes for a binding from the API when the Read() method
// is first called. This is only useful for binding types
// that store raw bytes, like WebAssembly modules
type bindingContentReader struct {
	api           *API
	requestParams *WorkerRequestParams
	bindingName   string
	content       []byte
	position      int
}

func (b *bindingContentReader) Read(p []byte) (n int, err error) {
	// Lazily load the content when Read() is first called
	if b.content == nil {
		uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings/%s/content", b.api.AccountID, b.requestParams.ScriptName, b.bindingName)
		res, err := b.api.makeRequest("GET", uri, nil)
		if err != nil {
			return 0, errors.Wrap(err, errMakeRequestError)
		}
		b.content = res
	}

	if b.position >= len(b.content) {
		return 0, io.EOF
	}

	bytesRemaining := len(b.content) - b.position
	bytesToProcess := 0
	if len(p) < bytesRemaining {
		bytesToProcess = len(p)
	} else {
		bytesToProcess = bytesRemaining
	}

	for i := 0; i < bytesToProcess; i++ {
		p[i] = b.content[b.position]
		b.position = b.position + 1
	}

	return bytesToProcess, nil
}

// ListWorkerScripts returns list of worker scripts for given account.
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) ListWorkerScripts() (WorkerListResponse, error) {
	if api.AccountID == "" {
		return WorkerListResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorker push raw script content for your worker.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(requestParams *WorkerRequestParams, data string) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, "application/javascript", []byte(data))
	}
	return api.uploadWorkerForZone(requestParams.ZoneID, "application/javascript", []byte(data))
}

// UploadWorkerWithBindings push raw script content and bindings for your worker
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorkerWithBindings(requestParams *WorkerRequestParams, data *WorkerScriptParams) (WorkerScriptResponse, error) {
	contentType, body, err := formatMultipartBody(data)
	if err != nil {
		return WorkerScriptResponse{}, err
	}
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, contentType, body)
	}
	return api.uploadWorkerForZone(requestParams.ZoneID, contentType, body)
}

func (api *API) uploadWorkerForZone(zoneID, contentType string, body []byte) (WorkerScriptResponse, error) {
	uri := "/zones/" + zoneID + "/workers/script"
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestWithHeaders("PUT", uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

func (api *API) uploadWorkerWithName(scriptName, contentType string, body []byte) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestWithHeaders("PUT", uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// Returns content-type, body, error
func formatMultipartBody(params *WorkerScriptParams) (string, []byte, error) {
	var buf = &bytes.Buffer{}
	var mpw = multipart.NewWriter(buf)
	defer mpw.Close()

	// Write metadata part
	scriptPartName := "script"
	meta := struct {
		BodyPart string              `json:"body_part"`
		Bindings []workerBindingMeta `json:"bindings"`
	}{
		BodyPart: scriptPartName,
		Bindings: make([]workerBindingMeta, 0, len(params.Bindings)),
	}

	bodyWriters := make([]workerBindingBodyWriter, 0, len(params.Bindings))
	for name, b := range params.Bindings {
		bindingMeta, bodyWriter, err := b.serialize(name)
		if err != nil {
			return "", nil, err
		}

		meta.Bindings = append(meta.Bindings, bindingMeta)
		bodyWriters = append(bodyWriters, bodyWriter)
	}

	var hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, "metadata"))
	hdr.Set("content-type", "application/json")
	pw, err := mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write(metaJSON)
	if err != nil {
		return "", nil, err
	}

	// Write script part
	hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, scriptPartName))
	hdr.Set("content-type", "application/javascript")
	pw, err = mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write([]byte(params.Script))
	if err != nil {
		return "", nil, err
	}

	// Write other bindings with parts
	for _, w := range bodyWriters {
		if w != nil {
			err = w(mpw)
			if err != nil {
				return "", nil, err
			}
		}
	}

	mpw.Close()

	return mpw.FormDataContentType(), buf.Bytes(), nil
}

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter, https://api.cloudflare.com/#worker-routes-create-route
func (api *API) CreateWorkerRoute(zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(api, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}

	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("POST", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerRoute deletes worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-routes-delete-route
func (api *API) DeleteWorkerRoute(zoneID string, routeID string) (WorkerRouteResponse, error) {
	uri := "/zones/" + zoneID + "/workers/routes/" + routeID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// ListWorkerRoutes returns list of worker routes
//
// API reference: https://api.cloudflare.com/#worker-filters-list-filters, https://api.cloudflare.com/#worker-routes-list-routes
func (api *API) ListWorkerRoutes(zoneID string) (WorkerRoutesResponse, error) {
	pathComponent := "filters"
	// Unfortunately we don't have a good signal of whether the user is wanting
	// to use the deprecated filters endpoint (https://api.cloudflare.com/#worker-filters-list-filters)
	// or the multi-script routes endpoint (https://api.cloudflare.com/#worker-script-list-workers)
	//
	// The filters endpoint does not support API tokens, so if an API token is specified we need to use
	// the routes endpoint. Otherwise, since the multi-script API endpoints that operate on a script
	// require an AccountID, we assume that anyone specifying an AccountID is using the routes endpoint.
	// This is likely too presumptuous. In the next major version, we should just remove the deprecated
	// filter endpoints entirely to avoid this ambiguity.
	if api.AccountID != "" || api.APIToken != "" {
		pathComponent = "routes"
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	for i := range r.Routes {
		route := &r.Routes[i]
		// The Enabled flag will not be set in the multi-script API response
		// so we manually set it to true if the script name is not empty
		// in case any multi-script customers rely on the Enabled field
		if route.Script != "" {
			route.Enabled = true
		}
	}
	return r, nil
}

// UpdateWorkerRoute updates worker route for a zone.
//
// API reference: https://api.cloudflare.com/#worker-filters-update-filter, https://api.cloudflare.com/#worker-routes-update-route
func (api *API) UpdateWorkerRoute(zoneID string, routeID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(api, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent + "/" + routeID
	res, err := api.makeRequest("PUT", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

func getRouteEndpoint(api *API, route WorkerRoute) (string, error) {
	if route.Script != "" && route.Enabled == true {
		return "", errors.New("Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
	}

	// For backwards-compatibility, fallback to the deprecated filter
	// endpoint if Enabled == true
	// https://api.cloudflare.com/#worker-filters-deprecated--properties
	if route.Enabled == true {
		return "filters", nil
	}

	return "routes", nil
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	"bytes"
	"encoding/hex"
>>>>>>> 5ce8c7613 (update vendored files)
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

	"github.com/pkg/errors"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

// WorkerScriptParams provides a worker script and the associated bindings
type WorkerScriptParams struct {
	Script string

	// Bindings should be a map where the keys are the binding name, and the
	// values are the binding content
	Bindings map[string]WorkerBinding
}

// WorkerRoute is used to map traffic matching a URL pattern to a workers
//
// API reference: https://api.cloudflare.com/#worker-routes-properties
type WorkerRoute struct {
	ID      string `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Enabled bool   `json:"enabled"` // this is deprecated: https://api.cloudflare.com/#worker-filters-deprecated--properties
	Script  string `json:"script,omitempty"`
}

// WorkerRoutesResponse embeds Response struct and slice of WorkerRoutes
type WorkerRoutesResponse struct {
	Response
	Routes []WorkerRoute `json:"result"`
}

// WorkerRouteResponse embeds Response struct and a single WorkerRoute
type WorkerRouteResponse struct {
	Response
	WorkerRoute `json:"result"`
}

// WorkerScript Cloudflare Worker struct with metadata
type WorkerScript struct {
	WorkerMetaData
	Script string `json:"script"`
}

// WorkerMetaData contains worker script information such as size, creation & modification dates
type WorkerMetaData struct {
	ID         string    `json:"id,omitempty"`
	ETAG       string    `json:"etag,omitempty"`
	Size       int       `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

// WorkerListResponse wrapper struct for API response to worker script list API call
type WorkerListResponse struct {
	Response
	WorkerList []WorkerMetaData `json:"result"`
}

// WorkerScriptResponse wrapper struct for API response to worker script calls
type WorkerScriptResponse struct {
	Response
	WorkerScript `json:"result"`
}

// WorkerBindingType represents a particular type of binding
type WorkerBindingType string

func (b WorkerBindingType) String() string {
	return string(b)
}

const (
	// WorkerInheritBindingType is the type for inherited bindings
	WorkerInheritBindingType WorkerBindingType = "inherit"
	// WorkerKvNamespaceBindingType is the type for KV Namespace bindings
	WorkerKvNamespaceBindingType WorkerBindingType = "kv_namespace"
	// WorkerWebAssemblyBindingType is the type for Web Assembly module bindings
	WorkerWebAssemblyBindingType WorkerBindingType = "wasm_module"
	// WorkerSecretTextBindingType is the type for secret text bindings
	WorkerSecretTextBindingType WorkerBindingType = "secret_text"
	// WorkerPlainTextBindingType is the type for plain text bindings
	WorkerPlainTextBindingType WorkerBindingType = "plain_text"
)

// WorkerBindingListItem a struct representing an individual binding in a list of bindings
type WorkerBindingListItem struct {
	Name    string `json:"name"`
	Binding WorkerBinding
}

// WorkerBindingListResponse wrapper struct for API response to worker binding list API call
type WorkerBindingListResponse struct {
	Response
	BindingList []WorkerBindingListItem
}

// Workers supports multiple types of bindings, e.g. KV namespaces or WebAssembly modules, and each type
// of binding will be represented differently in the upload request body. At a high-level, every binding
// will specify metadata, which is a JSON object with the properties "name" and "type". Some types of bindings
// will also have additional metadata properties. For example, KV bindings also specify the KV namespace.
// In addition to the metadata, some binding types may need to include additional data as part of the
// multipart form. For example, WebAssembly bindings will include the contents of the WebAssembly module.

// WorkerBinding is the generic interface implemented by all of
// the various binding types
type WorkerBinding interface {
	Type() WorkerBindingType

	// serialize is responsible for returning the binding metadata as well as an optionally
	// returning a function that can modify the multipart form body. For example, this is used
	// by WebAssembly bindings to add a new part containing the WebAssembly module contents.
	serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error)
}

// workerBindingMeta is the metadata portion of the binding
type workerBindingMeta = map[string]interface{}

// workerBindingBodyWriter allows for a binding to add additional parts to the multipart body
type workerBindingBodyWriter func(*multipart.Writer) error

// WorkerInheritBinding will just persist whatever binding content was previously uploaded
type WorkerInheritBinding struct {
	// Optional parameter that allows for renaming a binding without changing
	// its contents. If `OldName` is empty, the binding name will not be changed.
	OldName string
}

// Type returns the type of the binding
func (b WorkerInheritBinding) Type() WorkerBindingType {
	return WorkerInheritBindingType
}

func (b WorkerInheritBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	meta := workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
	}

	if b.OldName != "" {
		meta["old_name"] = b.OldName
	}

	return meta, nil, nil
}

// WorkerKvNamespaceBinding is a binding to a Workers KV Namespace
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/kv-namespaces/
type WorkerKvNamespaceBinding struct {
	NamespaceID string
}

// Type returns the type of the binding
func (b WorkerKvNamespaceBinding) Type() WorkerBindingType {
	return WorkerKvNamespaceBindingType
}

func (b WorkerKvNamespaceBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.NamespaceID == "" {
		return nil, nil, errors.Errorf(`NamespaceID for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":         bindingName,
		"type":         b.Type(),
		"namespace_id": b.NamespaceID,
	}, nil, nil
}

// WorkerWebAssemblyBinding is a binding to a WebAssembly module
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/webassembly-modules/
type WorkerWebAssemblyBinding struct {
	Module io.Reader
}

// Type returns the type of the binding
func (b WorkerWebAssemblyBinding) Type() WorkerBindingType {
	return WorkerWebAssemblyBindingType
}

func (b WorkerWebAssemblyBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	partName := getRandomPartName()

	bodyWriter := func(mpw *multipart.Writer) error {
		var hdr = textproto.MIMEHeader{}
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, partName))
		hdr.Set("content-type", "application/wasm")
		pw, err := mpw.CreatePart(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(pw, b.Module)
		return err
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"part": partName,
	}, bodyWriter, nil
}

// WorkerPlainTextBinding is a binding to plain text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-plain-text-binding
type WorkerPlainTextBinding struct {
	Text string
}

// Type returns the type of the binding
func (b WorkerPlainTextBinding) Type() WorkerBindingType {
	return WorkerPlainTextBindingType
}

func (b WorkerPlainTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, errors.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// WorkerSecretTextBinding is a binding to secret text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-secret-text-binding
type WorkerSecretTextBinding struct {
	Text string
}

// Type returns the type of the binding
func (b WorkerSecretTextBinding) Type() WorkerBindingType {
	return WorkerSecretTextBindingType
}

func (b WorkerSecretTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, errors.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// Each binding that adds a part to the multipart form body will need
// a unique part name so we just generate a random 128bit hex string
func getRandomPartName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}

// DeleteWorker deletes worker for a zone.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	// if ScriptName is provided we will treat as org request
	if requestParams.ScriptName != "" {
		return api.deleteWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerWithName deletes worker for a zone.
// Sccount must be specified as api option https://godoc.org/github.com/cloudflare/cloudflare-go#UsingAccount
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) deleteWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DownloadWorker fetch raw script content for your worker returns []byte containing worker code js
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) DownloadWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.downloadWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// DownloadWorkerWithName fetch raw script content for your worker returns string containing worker code js
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) downloadWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// ListWorkerBindings returns all the bindings for a particular worker
func (api *API) ListWorkerBindings(requestParams *WorkerRequestParams) (WorkerBindingListResponse, error) {
	if requestParams.ScriptName == "" {
		return WorkerBindingListResponse{}, errors.New("ScriptName is required")
	}
	if api.AccountID == "" {
		return WorkerBindingListResponse{}, errors.New("account ID required")
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings", api.AccountID, requestParams.ScriptName)

	var jsonRes struct {
		Response
		Bindings []workerBindingMeta `json:"result"`
	}
	var r WorkerBindingListResponse
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &jsonRes)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}

	r = WorkerBindingListResponse{
		Response:    jsonRes.Response,
		BindingList: make([]WorkerBindingListItem, 0, len(jsonRes.Bindings)),
	}
	for _, jsonBinding := range jsonRes.Bindings {
		name, ok := jsonBinding["name"].(string)
		if !ok {
			return r, errors.Errorf("Binding missing name %v", jsonBinding)
		}
		bType, ok := jsonBinding["type"].(string)
		if !ok {
			return r, errors.Errorf("Binding missing type %v", jsonBinding)
		}
		bindingListItem := WorkerBindingListItem{
			Name: name,
		}

		switch WorkerBindingType(bType) {
		case WorkerKvNamespaceBindingType:
			namespaceID := jsonBinding["namespace_id"].(string)
			bindingListItem.Binding = WorkerKvNamespaceBinding{
				NamespaceID: namespaceID,
			}
		case WorkerWebAssemblyBindingType:
			bindingListItem.Binding = WorkerWebAssemblyBinding{
				Module: &bindingContentReader{
					api:           api,
					requestParams: requestParams,
					bindingName:   name,
				},
			}
		case WorkerPlainTextBindingType:
			text := jsonBinding["text"].(string)
			bindingListItem.Binding = WorkerPlainTextBinding{
				Text: text,
			}
		case WorkerSecretTextBindingType:
			bindingListItem.Binding = WorkerSecretTextBinding{}
		default:
			bindingListItem.Binding = WorkerInheritBinding{}
		}
		r.BindingList = append(r.BindingList, bindingListItem)
	}

	return r, nil
}

// bindingContentReader is an io.Reader that will lazily load the
// raw bytes for a binding from the API when the Read() method
// is first called. This is only useful for binding types
// that store raw bytes, like WebAssembly modules
type bindingContentReader struct {
	api           *API
	requestParams *WorkerRequestParams
	bindingName   string
	content       []byte
	position      int
}

func (b *bindingContentReader) Read(p []byte) (n int, err error) {
	// Lazily load the content when Read() is first called
	if b.content == nil {
		uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings/%s/content", b.api.AccountID, b.requestParams.ScriptName, b.bindingName)
		res, err := b.api.makeRequest("GET", uri, nil)
		if err != nil {
			return 0, errors.Wrap(err, errMakeRequestError)
		}
		b.content = res
	}

	if b.position >= len(b.content) {
		return 0, io.EOF
	}

	bytesRemaining := len(b.content) - b.position
	bytesToProcess := 0
	if len(p) < bytesRemaining {
		bytesToProcess = len(p)
	} else {
		bytesToProcess = bytesRemaining
	}

	for i := 0; i < bytesToProcess; i++ {
		p[i] = b.content[b.position]
		b.position = b.position + 1
	}

	return bytesToProcess, nil
}

// ListWorkerScripts returns list of worker scripts for given account.
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) ListWorkerScripts() (WorkerListResponse, error) {
	if api.AccountID == "" {
		return WorkerListResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorker push raw script content for your worker.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(requestParams *WorkerRequestParams, data string) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, "application/javascript", []byte(data))
	}
	return api.uploadWorkerForZone(requestParams.ZoneID, "application/javascript", []byte(data))
}

// UploadWorkerWithBindings push raw script content and bindings for your worker
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorkerWithBindings(requestParams *WorkerRequestParams, data *WorkerScriptParams) (WorkerScriptResponse, error) {
	contentType, body, err := formatMultipartBody(data)
	if err != nil {
		return WorkerScriptResponse{}, err
	}
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, contentType, body)
	}
	return api.uploadWorkerForZone(requestParams.ZoneID, contentType, body)
}

func (api *API) uploadWorkerForZone(zoneID, contentType string, body []byte) (WorkerScriptResponse, error) {
	uri := "/zones/" + zoneID + "/workers/script"
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestWithHeaders("PUT", uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

func (api *API) uploadWorkerWithName(scriptName, contentType string, body []byte) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestWithHeaders("PUT", uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// Returns content-type, body, error
func formatMultipartBody(params *WorkerScriptParams) (string, []byte, error) {
	var buf = &bytes.Buffer{}
	var mpw = multipart.NewWriter(buf)
	defer mpw.Close()

	// Write metadata part
	scriptPartName := "script"
	meta := struct {
		BodyPart string              `json:"body_part"`
		Bindings []workerBindingMeta `json:"bindings"`
	}{
		BodyPart: scriptPartName,
		Bindings: make([]workerBindingMeta, 0, len(params.Bindings)),
	}

	bodyWriters := make([]workerBindingBodyWriter, 0, len(params.Bindings))
	for name, b := range params.Bindings {
		bindingMeta, bodyWriter, err := b.serialize(name)
		if err != nil {
			return "", nil, err
		}

		meta.Bindings = append(meta.Bindings, bindingMeta)
		bodyWriters = append(bodyWriters, bodyWriter)
	}

	var hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, "metadata"))
	hdr.Set("content-type", "application/json")
	pw, err := mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write(metaJSON)
	if err != nil {
		return "", nil, err
	}

	// Write script part
	hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, scriptPartName))
	hdr.Set("content-type", "application/javascript")
	pw, err = mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write([]byte(params.Script))
	if err != nil {
		return "", nil, err
	}

	// Write other bindings with parts
	for _, w := range bodyWriters {
		if w != nil {
			err = w(mpw)
			if err != nil {
				return "", nil, err
			}
		}
	}

	mpw.Close()

	return mpw.FormDataContentType(), buf.Bytes(), nil
}

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter, https://api.cloudflare.com/#worker-routes-create-route
func (api *API) CreateWorkerRoute(zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(api, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}

	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("POST", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerRoute deletes worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-routes-delete-route
func (api *API) DeleteWorkerRoute(zoneID string, routeID string) (WorkerRouteResponse, error) {
	uri := "/zones/" + zoneID + "/workers/routes/" + routeID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// ListWorkerRoutes returns list of worker routes
//
// API reference: https://api.cloudflare.com/#worker-filters-list-filters, https://api.cloudflare.com/#worker-routes-list-routes
func (api *API) ListWorkerRoutes(zoneID string) (WorkerRoutesResponse, error) {
	pathComponent := "filters"
	// Unfortunately we don't have a good signal of whether the user is wanting
	// to use the deprecated filters endpoint (https://api.cloudflare.com/#worker-filters-list-filters)
	// or the multi-script routes endpoint (https://api.cloudflare.com/#worker-script-list-workers)
	//
	// The filters endpoint does not support API tokens, so if an API token is specified we need to use
	// the routes endpoint. Otherwise, since the multi-script API endpoints that operate on a script
	// require an AccountID, we assume that anyone specifying an AccountID is using the routes endpoint.
	// This is likely too presumptuous. In the next major version, we should just remove the deprecated
	// filter endpoints entirely to avoid this ambiguity.
	if api.AccountID != "" || api.APIToken != "" {
		pathComponent = "routes"
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	for i := range r.Routes {
		route := &r.Routes[i]
		// The Enabled flag will not be set in the multi-script API response
		// so we manually set it to true if the script name is not empty
		// in case any multi-script customers rely on the Enabled field
		if route.Script != "" {
			route.Enabled = true
		}
	}
	return r, nil
}

// UpdateWorkerRoute updates worker route for a zone.
//
// API reference: https://api.cloudflare.com/#worker-filters-update-filter, https://api.cloudflare.com/#worker-routes-update-route
func (api *API) UpdateWorkerRoute(zoneID string, routeID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(api, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent + "/" + routeID
	res, err := api.makeRequest("PUT", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

func getRouteEndpoint(api *API, route WorkerRoute) (string, error) {
	if route.Script != "" && route.Enabled == true {
		return "", errors.New("Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
	}

	// For backwards-compatibility, fallback to the deprecated filter
	// endpoint if Enabled == true
	// https://api.cloudflare.com/#worker-filters-deprecated--properties
	if route.Enabled == true {
		return "filters", nil
	}

	return "routes", nil
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	"bytes"
	"context"
	rand "crypto/rand"
	"encoding/hex"
>>>>>>> 6b7ce455e (update vendored files)
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"strings"
	"time"

	"errors"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests.
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

// WorkerScriptParams provides a worker script and the associated bindings.
type WorkerScriptParams struct {
	Script string

	// Module changes the Content-Type header to specify the script is an
	// ES Module syntax script.
	Module bool

	// Bindings should be a map where the keys are the binding name, and the
	// values are the binding content
	Bindings map[string]WorkerBinding
}

// WorkerRoute is used to map traffic matching a URL pattern to a workers
//
// API reference: https://api.cloudflare.com/#worker-routes-properties
type WorkerRoute struct {
	ID      string `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Enabled bool   `json:"enabled"` // this is deprecated: https://api.cloudflare.com/#worker-filters-deprecated--properties
	Script  string `json:"script,omitempty"`
}

// WorkerRoutesResponse embeds Response struct and slice of WorkerRoutes.
type WorkerRoutesResponse struct {
	Response
	Routes []WorkerRoute `json:"result"`
}

// WorkerRouteResponse embeds Response struct and a single WorkerRoute.
type WorkerRouteResponse struct {
	Response
	WorkerRoute `json:"result"`
}

// WorkerScript Cloudflare Worker struct with metadata.
type WorkerScript struct {
	WorkerMetaData
	Script     string `json:"script"`
	UsageModel string `json:"usage_model,omitempty"`
}

// WorkerMetaData contains worker script information such as size, creation & modification dates.
type WorkerMetaData struct {
	ID         string    `json:"id,omitempty"`
	ETAG       string    `json:"etag,omitempty"`
	Size       int       `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

// WorkerListResponse wrapper struct for API response to worker script list API call.
type WorkerListResponse struct {
	Response
	WorkerList []WorkerMetaData `json:"result"`
}

// WorkerScriptResponse wrapper struct for API response to worker script calls.
type WorkerScriptResponse struct {
	Response
	Module       bool
	WorkerScript `json:"result"`
}

// WorkerBindingType represents a particular type of binding.
type WorkerBindingType string

func (b WorkerBindingType) String() string {
	return string(b)
}

const (
	// WorkerDurableObjectBindingType is the type for Durable Object bindings.
	WorkerDurableObjectBindingType WorkerBindingType = "durable_object_namespace"
	// WorkerInheritBindingType is the type for inherited bindings.
	WorkerInheritBindingType WorkerBindingType = "inherit"
	// WorkerKvNamespaceBindingType is the type for KV Namespace bindings.
	WorkerKvNamespaceBindingType WorkerBindingType = "kv_namespace"
	// WorkerWebAssemblyBindingType is the type for Web Assembly module bindings.
	WorkerWebAssemblyBindingType WorkerBindingType = "wasm_module"
	// WorkerSecretTextBindingType is the type for secret text bindings.
	WorkerSecretTextBindingType WorkerBindingType = "secret_text"
	// WorkerPlainTextBindingType is the type for plain text bindings.
	WorkerPlainTextBindingType WorkerBindingType = "plain_text"
	// WorkerServiceBindingType is the type for service bindings.
	WorkerServiceBindingType WorkerBindingType = "service"
	// WorkerR2BucketBindingType is the type for R2 bucket bindings.
	WorkerR2BucketBindingType WorkerBindingType = "r2_bucket"
)

// WorkerBindingListItem a struct representing an individual binding in a list of bindings.
type WorkerBindingListItem struct {
	Name    string `json:"name"`
	Binding WorkerBinding
}

// WorkerBindingListResponse wrapper struct for API response to worker binding list API call.
type WorkerBindingListResponse struct {
	Response
	BindingList []WorkerBindingListItem
}

// Workers supports multiple types of bindings, e.g. KV namespaces or WebAssembly modules, and each type
// of binding will be represented differently in the upload request body. At a high-level, every binding
// will specify metadata, which is a JSON object with the properties "name" and "type". Some types of bindings
// will also have additional metadata properties. For example, KV bindings also specify the KV namespace.
// In addition to the metadata, some binding types may need to include additional data as part of the
// multipart form. For example, WebAssembly bindings will include the contents of the WebAssembly module.

// WorkerBinding is the generic interface implemented by all of
// the various binding types.
type WorkerBinding interface {
	Type() WorkerBindingType

	// serialize is responsible for returning the binding metadata as well as an optionally
	// returning a function that can modify the multipart form body. For example, this is used
	// by WebAssembly bindings to add a new part containing the WebAssembly module contents.
	serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error)
}

// workerBindingMeta is the metadata portion of the binding.
type workerBindingMeta = map[string]interface{}

// workerBindingBodyWriter allows for a binding to add additional parts to the multipart body.
type workerBindingBodyWriter func(*multipart.Writer) error

// WorkerInheritBinding will just persist whatever binding content was previously uploaded.
type WorkerInheritBinding struct {
	// Optional parameter that allows for renaming a binding without changing
	// its contents. If `OldName` is empty, the binding name will not be changed.
	OldName string
}

// Type returns the type of the binding.
func (b WorkerInheritBinding) Type() WorkerBindingType {
	return WorkerInheritBindingType
}

func (b WorkerInheritBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	meta := workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
	}

	if b.OldName != "" {
		meta["old_name"] = b.OldName
	}

	return meta, nil, nil
}

// WorkerKvNamespaceBinding is a binding to a Workers KV Namespace
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/kv-namespaces/
type WorkerKvNamespaceBinding struct {
	NamespaceID string
}

// Type returns the type of the binding.
func (b WorkerKvNamespaceBinding) Type() WorkerBindingType {
	return WorkerKvNamespaceBindingType
}

func (b WorkerKvNamespaceBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.NamespaceID == "" {
		return nil, nil, fmt.Errorf(`NamespaceID for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":         bindingName,
		"type":         b.Type(),
		"namespace_id": b.NamespaceID,
	}, nil, nil
}

// WorkerDurableObjectBinding is a binding to a Workers Durable Object
//
// https://api.cloudflare.com/#durable-objects-namespace-properties
type WorkerDurableObjectBinding struct {
	ClassName  string
	ScriptName string
}

// Type returns the type of the binding.
func (b WorkerDurableObjectBinding) Type() WorkerBindingType {
	return WorkerDurableObjectBindingType
}

func (b WorkerDurableObjectBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.ClassName == "" {
		return nil, nil, fmt.Errorf(`ClassName for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":        bindingName,
		"type":        b.Type(),
		"class_name":  b.ClassName,
		"script_name": b.ScriptName,
	}, nil, nil
}

// WorkerWebAssemblyBinding is a binding to a WebAssembly module
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/webassembly-modules/
type WorkerWebAssemblyBinding struct {
	Module io.Reader
}

// Type returns the type of the binding.
func (b WorkerWebAssemblyBinding) Type() WorkerBindingType {
	return WorkerWebAssemblyBindingType
}

func (b WorkerWebAssemblyBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	partName := getRandomPartName()

	bodyWriter := func(mpw *multipart.Writer) error {
		var hdr = textproto.MIMEHeader{}
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, partName))
		hdr.Set("content-type", "application/wasm")
		pw, err := mpw.CreatePart(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(pw, b.Module)
		return err
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"part": partName,
	}, bodyWriter, nil
}

// WorkerPlainTextBinding is a binding to plain text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-plain-text-binding
type WorkerPlainTextBinding struct {
	Text string
}

// Type returns the type of the binding.
func (b WorkerPlainTextBinding) Type() WorkerBindingType {
	return WorkerPlainTextBindingType
}

func (b WorkerPlainTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, fmt.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// WorkerSecretTextBinding is a binding to secret text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-secret-text-binding
type WorkerSecretTextBinding struct {
	Text string
}

// Type returns the type of the binding.
func (b WorkerSecretTextBinding) Type() WorkerBindingType {
	return WorkerSecretTextBindingType
}

func (b WorkerSecretTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, fmt.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

type WorkerServiceBinding struct {
	Service     string
	Environment *string
}

func (b WorkerServiceBinding) Type() WorkerBindingType {
	return WorkerServiceBindingType
}

func (b WorkerServiceBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Service == "" {
		return nil, nil, fmt.Errorf(`Service for binding "%s" cannot be empty`, bindingName)
	}

	meta := workerBindingMeta{
		"name":    bindingName,
		"type":    b.Type(),
		"service": b.Service,
	}

	if b.Environment != nil {
		meta["environment"] = *b.Environment
	}

	return meta, nil, nil
}

// WorkerR2BucketBinding is a binding to an R2 bucket.
type WorkerR2BucketBinding struct {
	BucketName string
}

// Type returns the type of the binding.
func (b WorkerR2BucketBinding) Type() WorkerBindingType {
	return WorkerR2BucketBindingType
}

func (b WorkerR2BucketBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.BucketName == "" {
		return nil, nil, fmt.Errorf(`BucketName for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":        bindingName,
		"type":        b.Type(),
		"bucket_name": b.BucketName,
	}, nil, nil
}

// Each binding that adds a part to the multipart form body will need
// a unique part name so we just generate a random 128bit hex string.
func getRandomPartName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes) //nolint:errcheck
	return hex.EncodeToString(randBytes)
}

// DeleteWorker deletes worker for a zone.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(ctx context.Context, requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	// if ScriptName is provided we will treat as org request
	if requestParams.ScriptName != "" {
		return api.deleteWorkerWithName(ctx, requestParams.ScriptName)
	}
	uri := fmt.Sprintf("/zones/%s/workers/script", requestParams.ZoneID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// DeleteWorkerWithName deletes worker for a zone.
// Sccount must be specified as api option https://godoc.org/github.com/cloudflare/cloudflare-go#UsingAccount
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) deleteWorkerWithName(ctx context.Context, scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", api.AccountID, scriptName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// DownloadWorker fetch raw script content for your worker returns []byte containing worker code js
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) DownloadWorker(ctx context.Context, requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.downloadWorkerWithName(ctx, requestParams.ScriptName)
	}
	uri := fmt.Sprintf("/zones/%s/workers/script", requestParams.ZoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	r.Script = string(res)
	r.Module = false
	r.Success = true
	return r, nil
}

// DownloadWorkerWithName fetch raw script content for your worker returns string containing worker code js
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) downloadWorkerWithName(ctx context.Context, scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", api.AccountID, scriptName)
	res, err := api.makeRequestContextWithHeadersComplete(ctx, http.MethodGet, uri, nil, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}

	// Check if the response type is multipart, in which case this was a module worker
	mediaType, mediaParams, _ := mime.ParseMediaType(res.Headers.Get("content-type"))
	if strings.HasPrefix(mediaType, "multipart/") {
		bytesReader := bytes.NewReader(res.Body)
		mimeReader := multipart.NewReader(bytesReader, mediaParams["boundary"])
		mimePart, err := mimeReader.NextPart()
		if err != nil {
			return r, fmt.Errorf("could not get multipart response body: %w", err)
		}
		mimePartBody, err := ioutil.ReadAll(mimePart)
		if err != nil {
			return r, fmt.Errorf("could not read multipart response body: %w", err)
		}
		r.Script = string(mimePartBody)
		r.Module = true
	} else {
		r.Script = string(res.Body)
		r.Module = false
	}

	r.Success = true
	return r, nil
}

// ListWorkerBindings returns all the bindings for a particular worker.
func (api *API) ListWorkerBindings(ctx context.Context, requestParams *WorkerRequestParams) (WorkerBindingListResponse, error) {
	if requestParams.ScriptName == "" {
		return WorkerBindingListResponse{}, errors.New("ScriptName is required")
	}
	if api.AccountID == "" {
		return WorkerBindingListResponse{}, errors.New("account ID required")
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings", api.AccountID, requestParams.ScriptName)

	var jsonRes struct {
		Response
		Bindings []workerBindingMeta `json:"result"`
	}
	var r WorkerBindingListResponse
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &jsonRes)
	if err != nil {
		return r, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	r = WorkerBindingListResponse{
		Response:    jsonRes.Response,
		BindingList: make([]WorkerBindingListItem, 0, len(jsonRes.Bindings)),
	}
	for _, jsonBinding := range jsonRes.Bindings {
		name, ok := jsonBinding["name"].(string)
		if !ok {
			return r, fmt.Errorf("Binding missing name %v", jsonBinding)
		}
		bType, ok := jsonBinding["type"].(string)
		if !ok {
			return r, fmt.Errorf("Binding missing type %v", jsonBinding)
		}
		bindingListItem := WorkerBindingListItem{
			Name: name,
		}

		switch WorkerBindingType(bType) {
		case WorkerDurableObjectBindingType:
			class_name := jsonBinding["class_name"].(string)
			script_name := jsonBinding["script_name"].(string)
			bindingListItem.Binding = WorkerDurableObjectBinding{
				ClassName:  class_name,
				ScriptName: script_name,
			}
		case WorkerKvNamespaceBindingType:
			namespaceID := jsonBinding["namespace_id"].(string)
			bindingListItem.Binding = WorkerKvNamespaceBinding{
				NamespaceID: namespaceID,
			}
		case WorkerWebAssemblyBindingType:
			bindingListItem.Binding = WorkerWebAssemblyBinding{
				Module: &bindingContentReader{
					ctx:           ctx,
					api:           api,
					requestParams: requestParams,
					bindingName:   name,
				},
			}
		case WorkerPlainTextBindingType:
			text := jsonBinding["text"].(string)
			bindingListItem.Binding = WorkerPlainTextBinding{
				Text: text,
			}
		case WorkerServiceBindingType:
			service := jsonBinding["service"].(string)
			environment := jsonBinding["environment"].(string)
			bindingListItem.Binding = WorkerServiceBinding{
				Service:     service,
				Environment: &environment,
			}
		case WorkerSecretTextBindingType:
			bindingListItem.Binding = WorkerSecretTextBinding{}
		case WorkerR2BucketBindingType:
			bucketName := jsonBinding["bucket_name"].(string)
			bindingListItem.Binding = WorkerR2BucketBinding{
				BucketName: bucketName,
			}
		default:
			bindingListItem.Binding = WorkerInheritBinding{}
		}
		r.BindingList = append(r.BindingList, bindingListItem)
	}

	return r, nil
}

// bindingContentReader is an io.Reader that will lazily load the
// raw bytes for a binding from the API when the Read() method
// is first called. This is only useful for binding types
// that store raw bytes, like WebAssembly modules.
type bindingContentReader struct {
	api           *API
	requestParams *WorkerRequestParams
	ctx           context.Context
	bindingName   string
	content       []byte
	position      int
}

func (b *bindingContentReader) Read(p []byte) (n int, err error) {
	// Lazily load the content when Read() is first called
	if b.content == nil {
		uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings/%s/content", b.api.AccountID, b.requestParams.ScriptName, b.bindingName)
		res, err := b.api.makeRequestContext(b.ctx, http.MethodGet, uri, nil)
		if err != nil {
			return 0, err
		}
		b.content = res
	}

	if b.position >= len(b.content) {
		return 0, io.EOF
	}

	bytesRemaining := len(b.content) - b.position
	bytesToProcess := 0
	if len(p) < bytesRemaining {
		bytesToProcess = len(p)
	} else {
		bytesToProcess = bytesRemaining
	}

	for i := 0; i < bytesToProcess; i++ {
		p[i] = b.content[b.position]
		b.position = b.position + 1
	}

	return bytesToProcess, nil
}

// ListWorkerScripts returns list of worker scripts for given account.
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) ListWorkerScripts(ctx context.Context) (WorkerListResponse, error) {
	if api.AccountID == "" {
		return WorkerListResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerListResponse{}, err
	}
	var r WorkerListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerListResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// UploadWorker push raw script content for your worker.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(ctx context.Context, requestParams *WorkerRequestParams, params *WorkerScriptParams) (WorkerScriptResponse, error) {
	if params.Module {
		return api.UploadWorkerWithBindings(ctx, requestParams, params)
	}

	contentType := "application/javascript"
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(ctx, requestParams.ScriptName, contentType, []byte(params.Script))
	}
	return api.uploadWorkerForZone(ctx, requestParams.ZoneID, contentType, []byte(params.Script))
}

// UploadWorkerWithBindings push raw script content and bindings for your worker
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorkerWithBindings(ctx context.Context, requestParams *WorkerRequestParams, data *WorkerScriptParams) (WorkerScriptResponse, error) {
	contentType, body, err := formatMultipartBody(data)
	if err != nil {
		return WorkerScriptResponse{}, err
	}
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(ctx, requestParams.ScriptName, contentType, body)
	}
	return api.uploadWorkerForZone(ctx, requestParams.ZoneID, contentType, body)
}

func (api *API) uploadWorkerForZone(ctx context.Context, zoneID, contentType string, body []byte) (WorkerScriptResponse, error) {
	uri := fmt.Sprintf("/zones/%s/workers/script", zoneID)
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

func (api *API) uploadWorkerWithName(ctx context.Context, scriptName, contentType string, body []byte) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", api.AccountID, scriptName)
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// Returns content-type, body, error.
func formatMultipartBody(params *WorkerScriptParams) (string, []byte, error) {
	var buf = &bytes.Buffer{}
	var mpw = multipart.NewWriter(buf)
	defer mpw.Close()

	// Write metadata part
	var scriptPartName string
	meta := struct {
		BodyPart   string              `json:"body_part,omitempty"`
		MainModule string              `json:"main_module,omitempty"`
		Bindings   []workerBindingMeta `json:"bindings"`
	}{
		Bindings: make([]workerBindingMeta, 0, len(params.Bindings)),
	}

	if params.Module {
		scriptPartName = "worker.mjs"
		meta.MainModule = scriptPartName
	} else {
		scriptPartName = "script"
		meta.BodyPart = scriptPartName
	}

	bodyWriters := make([]workerBindingBodyWriter, 0, len(params.Bindings))
	for name, b := range params.Bindings {
		bindingMeta, bodyWriter, err := b.serialize(name)
		if err != nil {
			return "", nil, err
		}

		meta.Bindings = append(meta.Bindings, bindingMeta)
		bodyWriters = append(bodyWriters, bodyWriter)
	}

	var hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, "metadata"))
	hdr.Set("content-type", "application/json")
	pw, err := mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write(metaJSON)
	if err != nil {
		return "", nil, err
	}

	// Write script part
	hdr = textproto.MIMEHeader{}

	contentType := "application/javascript"
	if params.Module {
		contentType = "application/javascript+module"
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"; filename="%[1]s"`, scriptPartName))
	} else {
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, scriptPartName))
	}
	hdr.Set("content-type", contentType)

	pw, err = mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write([]byte(params.Script))
	if err != nil {
		return "", nil, err
	}

	// Write other bindings with parts
	for _, w := range bodyWriters {
		if w != nil {
			err = w(mpw)
			if err != nil {
				return "", nil, err
			}
		}
	}

	mpw.Close()

	return mpw.FormDataContentType(), buf.Bytes(), nil
}

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter, https://api.cloudflare.com/#worker-routes-create-route
func (api *API) CreateWorkerRoute(ctx context.Context, zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}

	uri := fmt.Sprintf("/zones/%s/workers/%s", zoneID, pathComponent)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// DeleteWorkerRoute deletes worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-routes-delete-route
func (api *API) DeleteWorkerRoute(ctx context.Context, zoneID string, routeID string) (WorkerRouteResponse, error) {
	uri := fmt.Sprintf("/zones/%s/workers/routes/%s", zoneID, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// ListWorkerRoutes returns list of worker routes
//
// API reference: https://api.cloudflare.com/#worker-filters-list-filters, https://api.cloudflare.com/#worker-routes-list-routes
func (api *API) ListWorkerRoutes(ctx context.Context, zoneID string) (WorkerRoutesResponse, error) {
	pathComponent := "filters"
	// Unfortunately we don't have a good signal of whether the user is wanting
	// to use the deprecated filters endpoint (https://api.cloudflare.com/#worker-filters-list-filters)
	// or the multi-script routes endpoint (https://api.cloudflare.com/#worker-script-list-workers)
	//
	// The filters endpoint does not support API tokens, so if an API token is specified we need to use
	// the routes endpoint. Otherwise, since the multi-script API endpoints that operate on a script
	// require an AccountID, we assume that anyone specifying an AccountID is using the routes endpoint.
	// This is likely too presumptuous. In the next major version, we should just remove the deprecated
	// filter endpoints entirely to avoid this ambiguity.
	if api.AccountID != "" || api.APIToken != "" {
		pathComponent = "routes"
	}
	uri := fmt.Sprintf("/zones/%s/workers/%s", zoneID, pathComponent)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, err
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	for i := range r.Routes {
		route := &r.Routes[i]
		// The Enabled flag will not be set in the multi-script API response
		// so we manually set it to true if the script name is not empty
		// in case any multi-script customers rely on the Enabled field
		if route.Script != "" {
			route.Enabled = true
		}
	}
	return r, nil
}

// GetWorkerRoute returns a worker route.
//
// API reference: https://api.cloudflare.com/#worker-routes-get-route
func (api *API) GetWorkerRoute(ctx context.Context, zoneID string, routeID string) (WorkerRouteResponse, error) {
	uri := fmt.Sprintf("/zones/%s/workers/routes/%s", zoneID, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// UpdateWorkerRoute updates worker route for a zone.
//
// API reference: https://api.cloudflare.com/#worker-filters-update-filter, https://api.cloudflare.com/#worker-routes-update-route
func (api *API) UpdateWorkerRoute(ctx context.Context, zoneID string, routeID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(api, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	uri := fmt.Sprintf("/zones/%s/workers/%s/%s", zoneID, pathComponent, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

func getRouteEndpoint(api *API, route WorkerRoute) (string, error) {
	if route.Script != "" && route.Enabled {
		return "", errors.New("Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
	}

	// For backwards-compatibility, fallback to the deprecated filter
	// endpoint if Enabled == true
	// https://api.cloudflare.com/#worker-filters-deprecated--properties
	if route.Enabled {
		return "filters", nil
	}

	return "routes", nil
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"bytes"
	"context"
	"encoding/hex"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"

	"github.com/pkg/errors"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

// WorkerScriptParams provides a worker script and the associated bindings
type WorkerScriptParams struct {
	Script string

	// Bindings should be a map where the keys are the binding name, and the
	// values are the binding content
	Bindings map[string]WorkerBinding
}

// WorkerRoute is used to map traffic matching a URL pattern to a workers
//
// API reference: https://api.cloudflare.com/#worker-routes-properties
type WorkerRoute struct {
	ID      string `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Enabled bool   `json:"enabled"` // this is deprecated: https://api.cloudflare.com/#worker-filters-deprecated--properties
	Script  string `json:"script,omitempty"`
}

// WorkerRoutesResponse embeds Response struct and slice of WorkerRoutes
type WorkerRoutesResponse struct {
	Response
	Routes []WorkerRoute `json:"result"`
}

// WorkerRouteResponse embeds Response struct and a single WorkerRoute
type WorkerRouteResponse struct {
	Response
	WorkerRoute `json:"result"`
}

// WorkerScript Cloudflare Worker struct with metadata
type WorkerScript struct {
	WorkerMetaData
	Script string `json:"script"`
}

// WorkerMetaData contains worker script information such as size, creation & modification dates
type WorkerMetaData struct {
	ID         string    `json:"id,omitempty"`
	ETAG       string    `json:"etag,omitempty"`
	Size       int       `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

// WorkerListResponse wrapper struct for API response to worker script list API call
type WorkerListResponse struct {
	Response
	WorkerList []WorkerMetaData `json:"result"`
}

// WorkerScriptResponse wrapper struct for API response to worker script calls
type WorkerScriptResponse struct {
	Response
	WorkerScript `json:"result"`
}

// WorkerBindingType represents a particular type of binding
type WorkerBindingType string

func (b WorkerBindingType) String() string {
	return string(b)
}

const (
	// WorkerInheritBindingType is the type for inherited bindings
	WorkerInheritBindingType WorkerBindingType = "inherit"
	// WorkerKvNamespaceBindingType is the type for KV Namespace bindings
	WorkerKvNamespaceBindingType WorkerBindingType = "kv_namespace"
	// WorkerWebAssemblyBindingType is the type for Web Assembly module bindings
	WorkerWebAssemblyBindingType WorkerBindingType = "wasm_module"
	// WorkerSecretTextBindingType is the type for secret text bindings
	WorkerSecretTextBindingType WorkerBindingType = "secret_text"
	// WorkerPlainTextBindingType is the type for plain text bindings
	WorkerPlainTextBindingType WorkerBindingType = "plain_text"
)

// WorkerBindingListItem a struct representing an individual binding in a list of bindings
type WorkerBindingListItem struct {
	Name    string `json:"name"`
	Binding WorkerBinding
}

// WorkerBindingListResponse wrapper struct for API response to worker binding list API call
type WorkerBindingListResponse struct {
	Response
	BindingList []WorkerBindingListItem
}

// Workers supports multiple types of bindings, e.g. KV namespaces or WebAssembly modules, and each type
// of binding will be represented differently in the upload request body. At a high-level, every binding
// will specify metadata, which is a JSON object with the properties "name" and "type". Some types of bindings
// will also have additional metadata properties. For example, KV bindings also specify the KV namespace.
// In addition to the metadata, some binding types may need to include additional data as part of the
// multipart form. For example, WebAssembly bindings will include the contents of the WebAssembly module.

// WorkerBinding is the generic interface implemented by all of
// the various binding types
type WorkerBinding interface {
	Type() WorkerBindingType

	// serialize is responsible for returning the binding metadata as well as an optionally
	// returning a function that can modify the multipart form body. For example, this is used
	// by WebAssembly bindings to add a new part containing the WebAssembly module contents.
	serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error)
}

// workerBindingMeta is the metadata portion of the binding
type workerBindingMeta = map[string]interface{}

// workerBindingBodyWriter allows for a binding to add additional parts to the multipart body
type workerBindingBodyWriter func(*multipart.Writer) error

// WorkerInheritBinding will just persist whatever binding content was previously uploaded
type WorkerInheritBinding struct {
	// Optional parameter that allows for renaming a binding without changing
	// its contents. If `OldName` is empty, the binding name will not be changed.
	OldName string
}

// Type returns the type of the binding
func (b WorkerInheritBinding) Type() WorkerBindingType {
	return WorkerInheritBindingType
}

func (b WorkerInheritBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	meta := workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
	}

	if b.OldName != "" {
		meta["old_name"] = b.OldName
	}

	return meta, nil, nil
}

// WorkerKvNamespaceBinding is a binding to a Workers KV Namespace
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/kv-namespaces/
type WorkerKvNamespaceBinding struct {
	NamespaceID string
}

// Type returns the type of the binding
func (b WorkerKvNamespaceBinding) Type() WorkerBindingType {
	return WorkerKvNamespaceBindingType
}

func (b WorkerKvNamespaceBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.NamespaceID == "" {
		return nil, nil, errors.Errorf(`NamespaceID for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":         bindingName,
		"type":         b.Type(),
		"namespace_id": b.NamespaceID,
	}, nil, nil
}

// WorkerWebAssemblyBinding is a binding to a WebAssembly module
//
// https://developers.cloudflare.com/workers/archive/api/resource-bindings/webassembly-modules/
type WorkerWebAssemblyBinding struct {
	Module io.Reader
}

// Type returns the type of the binding
func (b WorkerWebAssemblyBinding) Type() WorkerBindingType {
	return WorkerWebAssemblyBindingType
}

func (b WorkerWebAssemblyBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	partName := getRandomPartName()

	bodyWriter := func(mpw *multipart.Writer) error {
		var hdr = textproto.MIMEHeader{}
		hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, partName))
		hdr.Set("content-type", "application/wasm")
		pw, err := mpw.CreatePart(hdr)
		if err != nil {
			return err
		}
		_, err = io.Copy(pw, b.Module)
		return err
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"part": partName,
	}, bodyWriter, nil
}

// WorkerPlainTextBinding is a binding to plain text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-plain-text-binding
type WorkerPlainTextBinding struct {
	Text string
}

// Type returns the type of the binding
func (b WorkerPlainTextBinding) Type() WorkerBindingType {
	return WorkerPlainTextBindingType
}

func (b WorkerPlainTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, errors.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// WorkerSecretTextBinding is a binding to secret text
//
// https://developers.cloudflare.com/workers/tooling/api/scripts/#add-a-secret-text-binding
type WorkerSecretTextBinding struct {
	Text string
}

// Type returns the type of the binding
func (b WorkerSecretTextBinding) Type() WorkerBindingType {
	return WorkerSecretTextBindingType
}

func (b WorkerSecretTextBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Text == "" {
		return nil, nil, errors.Errorf(`Text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// Each binding that adds a part to the multipart form body will need
// a unique part name so we just generate a random 128bit hex string
func getRandomPartName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}

// DeleteWorker deletes worker for a zone.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(ctx context.Context, requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	// if ScriptName is provided we will treat as org request
	if requestParams.ScriptName != "" {
		return api.deleteWorkerWithName(ctx, requestParams.ScriptName)
	}
	uri := fmt.Sprintf("/zones/%s/workers/script", requestParams.ZoneID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerWithName deletes worker for a zone.
// Sccount must be specified as api option https://godoc.org/github.com/cloudflare/cloudflare-go#UsingAccount
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) deleteWorkerWithName(ctx context.Context, scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", api.AccountID, scriptName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DownloadWorker fetch raw script content for your worker returns []byte containing worker code js
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) DownloadWorker(ctx context.Context, requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.downloadWorkerWithName(ctx, requestParams.ScriptName)
	}
	uri := fmt.Sprintf("/zones/%s/workers/script", requestParams.ZoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// DownloadWorkerWithName fetch raw script content for your worker returns string containing worker code js
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) downloadWorkerWithName(ctx context.Context, scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", api.AccountID, scriptName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// ListWorkerBindings returns all the bindings for a particular worker
func (api *API) ListWorkerBindings(ctx context.Context, requestParams *WorkerRequestParams) (WorkerBindingListResponse, error) {
	if requestParams.ScriptName == "" {
		return WorkerBindingListResponse{}, errors.New("ScriptName is required")
	}
	if api.AccountID == "" {
		return WorkerBindingListResponse{}, errors.New("account ID required")
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings", api.AccountID, requestParams.ScriptName)

	var jsonRes struct {
		Response
		Bindings []workerBindingMeta `json:"result"`
	}
	var r WorkerBindingListResponse
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &jsonRes)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}

	r = WorkerBindingListResponse{
		Response:    jsonRes.Response,
		BindingList: make([]WorkerBindingListItem, 0, len(jsonRes.Bindings)),
	}
	for _, jsonBinding := range jsonRes.Bindings {
		name, ok := jsonBinding["name"].(string)
		if !ok {
			return r, errors.Errorf("Binding missing name %v", jsonBinding)
		}
		bType, ok := jsonBinding["type"].(string)
		if !ok {
			return r, errors.Errorf("Binding missing type %v", jsonBinding)
		}
		bindingListItem := WorkerBindingListItem{
			Name: name,
		}

		switch WorkerBindingType(bType) {
		case WorkerKvNamespaceBindingType:
			namespaceID := jsonBinding["namespace_id"].(string)
			bindingListItem.Binding = WorkerKvNamespaceBinding{
				NamespaceID: namespaceID,
			}
		case WorkerWebAssemblyBindingType:
			bindingListItem.Binding = WorkerWebAssemblyBinding{
				Module: &bindingContentReader{
					api:           api,
					requestParams: requestParams,
					bindingName:   name,
				},
			}
		case WorkerPlainTextBindingType:
			text := jsonBinding["text"].(string)
			bindingListItem.Binding = WorkerPlainTextBinding{
				Text: text,
			}
		case WorkerSecretTextBindingType:
			bindingListItem.Binding = WorkerSecretTextBinding{}
		default:
			bindingListItem.Binding = WorkerInheritBinding{}
		}
		r.BindingList = append(r.BindingList, bindingListItem)
	}

	return r, nil
}

// bindingContentReader is an io.Reader that will lazily load the
// raw bytes for a binding from the API when the Read() method
// is first called. This is only useful for binding types
// that store raw bytes, like WebAssembly modules
type bindingContentReader struct {
	api           *API
	requestParams *WorkerRequestParams
	bindingName   string
	content       []byte
	position      int
}

func (b *bindingContentReader) Read(p []byte) (n int, err error) {
	// Lazily load the content when Read() is first called
	if b.content == nil {
		uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings/%s/content", b.api.AccountID, b.requestParams.ScriptName, b.bindingName)
		res, err := b.api.makeRequest(http.MethodGet, uri, nil)
		if err != nil {
			return 0, err
		}
		b.content = res
	}

	if b.position >= len(b.content) {
		return 0, io.EOF
	}

	bytesRemaining := len(b.content) - b.position
	bytesToProcess := 0
	if len(p) < bytesRemaining {
		bytesToProcess = len(p)
	} else {
		bytesToProcess = bytesRemaining
	}

	for i := 0; i < bytesToProcess; i++ {
		p[i] = b.content[b.position]
		b.position = b.position + 1
	}

	return bytesToProcess, nil
}

// ListWorkerScripts returns list of worker scripts for given account.
//
// API reference: https://developers.cloudflare.com/workers/tooling/api/scripts/
func (api *API) ListWorkerScripts(ctx context.Context) (WorkerListResponse, error) {
	if api.AccountID == "" {
		return WorkerListResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts", api.AccountID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerListResponse{}, err
	}
	var r WorkerListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorker push raw script content for your worker.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(ctx context.Context, requestParams *WorkerRequestParams, data string) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(ctx, requestParams.ScriptName, "application/javascript", []byte(data))
	}
	return api.uploadWorkerForZone(ctx, requestParams.ZoneID, "application/javascript", []byte(data))
}

// UploadWorkerWithBindings push raw script content and bindings for your worker
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorkerWithBindings(ctx context.Context, requestParams *WorkerRequestParams, data *WorkerScriptParams) (WorkerScriptResponse, error) {
	contentType, body, err := formatMultipartBody(data)
	if err != nil {
		return WorkerScriptResponse{}, err
	}
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(ctx, requestParams.ScriptName, contentType, body)
	}
	return api.uploadWorkerForZone(ctx, requestParams.ZoneID, contentType, body)
}

func (api *API) uploadWorkerForZone(ctx context.Context, zoneID, contentType string, body []byte) (WorkerScriptResponse, error) {
	uri := fmt.Sprintf("/zones/%s/workers/script", zoneID)
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

func (api *API) uploadWorkerWithName(ctx context.Context, scriptName, contentType string, body []byte) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required")
	}
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s", api.AccountID, scriptName)
	headers := make(http.Header)
	headers.Set("Content-Type", contentType)
	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPut, uri, body, headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, err
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// Returns content-type, body, error
func formatMultipartBody(params *WorkerScriptParams) (string, []byte, error) {
	var buf = &bytes.Buffer{}
	var mpw = multipart.NewWriter(buf)
	defer mpw.Close()

	// Write metadata part
	scriptPartName := "script"
	meta := struct {
		BodyPart string              `json:"body_part"`
		Bindings []workerBindingMeta `json:"bindings"`
	}{
		BodyPart: scriptPartName,
		Bindings: make([]workerBindingMeta, 0, len(params.Bindings)),
	}

	bodyWriters := make([]workerBindingBodyWriter, 0, len(params.Bindings))
	for name, b := range params.Bindings {
		bindingMeta, bodyWriter, err := b.serialize(name)
		if err != nil {
			return "", nil, err
		}

		meta.Bindings = append(meta.Bindings, bindingMeta)
		bodyWriters = append(bodyWriters, bodyWriter)
	}

	var hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, "metadata"))
	hdr.Set("content-type", "application/json")
	pw, err := mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write(metaJSON)
	if err != nil {
		return "", nil, err
	}

	// Write script part
	hdr = textproto.MIMEHeader{}
	hdr.Set("content-disposition", fmt.Sprintf(`form-data; name="%s"`, scriptPartName))
	hdr.Set("content-type", "application/javascript")
	pw, err = mpw.CreatePart(hdr)
	if err != nil {
		return "", nil, err
	}
	_, err = pw.Write([]byte(params.Script))
	if err != nil {
		return "", nil, err
	}

	// Write other bindings with parts
	for _, w := range bodyWriters {
		if w != nil {
			err = w(mpw)
			if err != nil {
				return "", nil, err
			}
		}
	}

	mpw.Close()

	return mpw.FormDataContentType(), buf.Bytes(), nil
}

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter, https://api.cloudflare.com/#worker-routes-create-route
func (api *API) CreateWorkerRoute(ctx context.Context, zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(api, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}

	uri := fmt.Sprintf("/zones/%s/workers/%s", zoneID, pathComponent)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerRoute deletes worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-routes-delete-route
func (api *API) DeleteWorkerRoute(ctx context.Context, zoneID string, routeID string) (WorkerRouteResponse, error) {
	uri := fmt.Sprintf("/zones/%s/workers/routes/%s", zoneID, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// ListWorkerRoutes returns list of worker routes
//
// API reference: https://api.cloudflare.com/#worker-filters-list-filters, https://api.cloudflare.com/#worker-routes-list-routes
func (api *API) ListWorkerRoutes(ctx context.Context, zoneID string) (WorkerRoutesResponse, error) {
	pathComponent := "filters"
	// Unfortunately we don't have a good signal of whether the user is wanting
	// to use the deprecated filters endpoint (https://api.cloudflare.com/#worker-filters-list-filters)
	// or the multi-script routes endpoint (https://api.cloudflare.com/#worker-script-list-workers)
	//
	// The filters endpoint does not support API tokens, so if an API token is specified we need to use
	// the routes endpoint. Otherwise, since the multi-script API endpoints that operate on a script
	// require an AccountID, we assume that anyone specifying an AccountID is using the routes endpoint.
	// This is likely too presumptuous. In the next major version, we should just remove the deprecated
	// filter endpoints entirely to avoid this ambiguity.
	if api.AccountID != "" || api.APIToken != "" {
		pathComponent = "routes"
	}
	uri := fmt.Sprintf("/zones/%s/workers/%s", zoneID, pathComponent)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, err
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	for i := range r.Routes {
		route := &r.Routes[i]
		// The Enabled flag will not be set in the multi-script API response
		// so we manually set it to true if the script name is not empty
		// in case any multi-script customers rely on the Enabled field
		if route.Script != "" {
			route.Enabled = true
		}
	}
	return r, nil
}

// GetWorkerRoute returns a worker route.
//
// API reference: https://api.cloudflare.com/#worker-routes-get-route
func (api *API) GetWorkerRoute(ctx context.Context, zoneID string, routeID string) (WorkerRouteResponse, error) {
	uri := fmt.Sprintf("/zones/%s/workers/routes/%s", zoneID, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// UpdateWorkerRoute updates worker route for a zone.
//
// API reference: https://api.cloudflare.com/#worker-filters-update-filter, https://api.cloudflare.com/#worker-routes-update-route
func (api *API) UpdateWorkerRoute(ctx context.Context, zoneID string, routeID string, route WorkerRoute) (WorkerRouteResponse, error) {
	pathComponent, err := getRouteEndpoint(route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	uri := fmt.Sprintf("/zones/%s/workers/%s/%s", zoneID, pathComponent, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, route)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

func getRouteEndpoint(route WorkerRoute) (string, error) {
	if route.Script != "" && route.Enabled {
		return "", errors.New("Only `Script` or `Enabled` may be specified for a WorkerRoute, not both")
	}

	// For backwards-compatibility, fallback to the deprecated filter
	// endpoint if Enabled == true
	// https://api.cloudflare.com/#worker-filters-deprecated--properties
	if route.Enabled {
		return "filters", nil
	}

	return "routes", nil
}

type WorkerDomainParams struct {
	ZoneID      string `json:"zone_id"`
	Hostname    string `json:"hostname"`
	Service     string `json:"service"`
	Environment string `json:"environment,omitempty"`
}

type WorkerDomainResult struct {
	ID          string `json:"id"`
	ZoneID      string `json:"zone_id"`
	ZoneName    string `json:"zone_name"`
	Hostname    string `json:"hostname"`
	Service     string `json:"service"`
	Environment string `json:"environment"`
}

type WorkerDomainResponse struct {
	Response
	WorkerDomainResult `json:"result"`
}

// AttachWorkerToDomain attaches a worker to a zone and hostname
//
// API reference: https://api.cloudflare.com/#worker-domain-attach-to-domain
func (api *API) AttachWorkerToDomain(ctx context.Context, rc *ResourceContainer, params *WorkerDomainParams) (WorkerDomainResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/domains", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return WorkerDomainResponse{}, err
	}

	var r WorkerDomainResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerDomainResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// WorkerRequestParams provides parameters for worker requests for both enterprise and standard requests
type WorkerRequestParams struct {
	ZoneID     string
	ScriptName string
}

// WorkerRoute aka filters are patterns used to enable or disable workers that match requests.
//
// API reference: https://api.cloudflare.com/#worker-filters-properties
type WorkerRoute struct {
	ID      string `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Enabled bool   `json:"enabled"`
	Script  string `json:"script,omitempty"`
}

// WorkerRoutesResponse embeds Response struct and slice of WorkerRoutes
type WorkerRoutesResponse struct {
	Response
	Routes []WorkerRoute `json:"result"`
}

// WorkerRouteResponse embeds Response struct and a single WorkerRoute
type WorkerRouteResponse struct {
	Response
	WorkerRoute `json:"result"`
}

// WorkerScript Cloudflare Worker struct with metadata
type WorkerScript struct {
	WorkerMetaData
	Script string `json:"script"`
}

// WorkerMetaData contains worker script information such as size, creation & modification dates
type WorkerMetaData struct {
	ID         string    `json:"id,omitempty"`
	ETAG       string    `json:"etag,omitempty"`
	Size       int       `json:"size,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

// WorkerListResponse wrapper struct for API response to worker script list API call
type WorkerListResponse struct {
	Response
	WorkerList []WorkerMetaData `json:"result"`
}

// WorkerScriptResponse wrapper struct for API response to worker script calls
type WorkerScriptResponse struct {
	Response
	WorkerScript `json:"result"`
}

// DeleteWorker deletes worker for a zone.
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) DeleteWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	// if ScriptName is provided we will treat as org request
	if requestParams.ScriptName != "" {
		return api.deleteWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerWithName deletes worker for a zone.
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise
// account must be specified as api option https://godoc.org/github.com/cloudflare/cloudflare-go#UsingAccount
//
// API reference: https://api.cloudflare.com/#worker-script-delete-worker
func (api *API) deleteWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("DELETE", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DownloadWorker fetch raw script content for your worker returns []byte containing worker code js
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) DownloadWorker(requestParams *WorkerRequestParams) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.downloadWorkerWithName(requestParams.ScriptName)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// DownloadWorkerWithName fetch raw script content for your worker returns string containing worker code js
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
//
// API reference: https://api.cloudflare.com/#worker-script-download-worker
func (api *API) downloadWorkerWithName(scriptName string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	res, err := api.makeRequest("GET", uri, nil)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	r.Script = string(res)
	r.Success = true
	return r, nil
}

// ListWorkerScripts returns list of worker scripts for given account.
//
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise
//
// API reference: https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
func (api *API) ListWorkerScripts() (WorkerListResponse, error) {
	if api.AccountID == "" {
		return WorkerListResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerListResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerListResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorker push raw script content for your worker.
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) UploadWorker(requestParams *WorkerRequestParams, data string) (WorkerScriptResponse, error) {
	if requestParams.ScriptName != "" {
		return api.uploadWorkerWithName(requestParams.ScriptName, data)
	}
	uri := "/zones/" + requestParams.ZoneID + "/workers/script"
	headers := make(http.Header)
	headers.Set("Content-Type", "application/javascript")
	res, err := api.makeRequestWithHeaders("PUT", uri, []byte(data), headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// UploadWorkerWithName push raw script content for your worker.
//
// This is an enterprise only feature https://developers.cloudflare.com/workers/api/config-api-for-enterprise/
//
// API reference: https://api.cloudflare.com/#worker-script-upload-worker
func (api *API) uploadWorkerWithName(scriptName string, data string) (WorkerScriptResponse, error) {
	if api.AccountID == "" {
		return WorkerScriptResponse{}, errors.New("account ID required for enterprise only request")
	}
	uri := "/accounts/" + api.AccountID + "/workers/scripts/" + scriptName
	headers := make(http.Header)
	headers.Set("Content-Type", "application/javascript")
	res, err := api.makeRequestWithHeaders("PUT", uri, []byte(data), headers)
	var r WorkerScriptResponse
	if err != nil {
		return r, errors.Wrap(err, errMakeRequestError)
	}
	err = json.Unmarshal(res, &r)
	if err != nil {
		return r, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// CreateWorkerRoute creates worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-create-filter
func (api *API) CreateWorkerRoute(zoneID string, route WorkerRoute) (WorkerRouteResponse, error) {
	// Check whether a script name is defined in order to determine whether
	// to use the single-script or multi-script endpoint.
	pathComponent := "filters"
	if route.Script != "" {
		if api.AccountID == "" {
			return WorkerRouteResponse{}, errors.New("account ID required for enterprise only request")
		}
		pathComponent = "routes"
	}

	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("POST", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// DeleteWorkerRoute deletes worker route for a zone
//
// API reference: https://api.cloudflare.com/#worker-filters-delete-filter
func (api *API) DeleteWorkerRoute(zoneID string, routeID string) (WorkerRouteResponse, error) {
	// For deleting a route, it doesn't matter whether we use the
	// single-script or multi-script endpoint
	uri := "/zones/" + zoneID + "/workers/filters/" + routeID
	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	return r, nil
}

// ListWorkerRoutes returns list of worker routes
//
// API reference: https://api.cloudflare.com/#worker-filters-list-filters
func (api *API) ListWorkerRoutes(zoneID string) (WorkerRoutesResponse, error) {
	pathComponent := "filters"
	if api.AccountID != "" {
		pathComponent = "routes"
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, errors.Wrap(err, errUnmarshalError)
	}
	for i := range r.Routes {
		route := &r.Routes[i]
		// The Enabled flag will not be set in the multi-script API response
		// so we manually set it to true if the script name is not empty
		// in case any multi-script customers rely on the Enabled field
		if route.Script != "" {
			route.Enabled = true
		}
	}
	return r, nil
}

// UpdateWorkerRoute updates worker route for a zone.
//
// API reference: https://api.cloudflare.com/#worker-filters-update-filter
func (api *API) UpdateWorkerRoute(zoneID string, routeID string, route WorkerRoute) (WorkerRouteResponse, error) {
	// Check whether a script name is defined in order to determine whether
	// to use the single-script or multi-script endpoint.
	pathComponent := "filters"
	if route.Script != "" {
		if api.AccountID == "" {
			return WorkerRouteResponse{}, errors.New("account ID required for enterprise only request")
		}
		pathComponent = "routes"
	}
	uri := "/zones/" + zoneID + "/workers/" + pathComponent + "/" + routeID
	res, err := api.makeRequest("PUT", uri, route)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errMakeRequestError)
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, errors.Wrap(err, errUnmarshalError)
	}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	return r, nil
}
