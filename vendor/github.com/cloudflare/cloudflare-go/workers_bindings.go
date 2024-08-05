package cloudflare

import (
	"context"
	rand "crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"

	"github.com/goccy/go-json"
)

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
	// WorkerAnalyticsEngineBindingType is the type for Analytics Engine dataset bindings.
	WorkerAnalyticsEngineBindingType WorkerBindingType = "analytics_engine"
	// WorkerQueueBindingType is the type for queue bindings.
	WorkerQueueBindingType WorkerBindingType = "queue"
	// DispatchNamespaceBindingType is the type for WFP namespace bindings.
	DispatchNamespaceBindingType WorkerBindingType = "dispatch_namespace"
	// WorkerD1DataseBindingType is for D1 databases.
	WorkerD1DataseBindingType WorkerBindingType = "d1"
)

type ListWorkerBindingsParams struct {
	ScriptName        string
	DispatchNamespace *string
}

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

// WorkerKvNamespaceBinding is a binding to a Workers KV Namespace.
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
		return nil, nil, fmt.Errorf(`namespace ID for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":         bindingName,
		"type":         b.Type(),
		"namespace_id": b.NamespaceID,
	}, nil, nil
}

// WorkerDurableObjectBinding is a binding to a Workers Durable Object.
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

// WorkerWebAssemblyBinding is a binding to a WebAssembly module.
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

// WorkerPlainTextBinding is a binding to plain text.
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
		return nil, nil, fmt.Errorf(`text for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"text": b.Text,
	}, nil, nil
}

// WorkerSecretTextBinding is a binding to secret text.
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
		return nil, nil, fmt.Errorf(`text for binding "%s" cannot be empty`, bindingName)
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
		return nil, nil, fmt.Errorf(`service for binding "%s" cannot be empty`, bindingName)
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

// WorkerAnalyticsEngineBinding is a binding to an Analytics Engine dataset.
type WorkerAnalyticsEngineBinding struct {
	Dataset string
}

// Type returns the type of the binding.
func (b WorkerAnalyticsEngineBinding) Type() WorkerBindingType {
	return WorkerAnalyticsEngineBindingType
}

func (b WorkerAnalyticsEngineBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Dataset == "" {
		return nil, nil, fmt.Errorf(`dataset for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name":    bindingName,
		"type":    b.Type(),
		"dataset": b.Dataset,
	}, nil, nil
}

// WorkerQueueBinding is a binding to a Workers Queue.
//
// https://developers.cloudflare.com/workers/platform/bindings/#queue-bindings
type WorkerQueueBinding struct {
	Binding string
	Queue   string
}

// Type returns the type of the binding.
func (b WorkerQueueBinding) Type() WorkerBindingType {
	return WorkerQueueBindingType
}

func (b WorkerQueueBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Binding == "" {
		return nil, nil, fmt.Errorf(`binding name for binding "%s" cannot be empty`, bindingName)
	}
	if b.Queue == "" {
		return nil, nil, fmt.Errorf(`queue name for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"type":       b.Type(),
		"name":       b.Binding,
		"queue_name": b.Queue,
	}, nil, nil
}

// DispatchNamespaceBinding is a binding to a Workers for Platforms namespace
//
// https://developers.cloudflare.com/workers/platform/bindings/#dispatch-namespace-bindings-workers-for-platforms
type DispatchNamespaceBinding struct {
	Binding   string
	Namespace string
	Outbound  *NamespaceOutboundOptions
}

type NamespaceOutboundOptions struct {
	Worker WorkerReference
	Params []OutboundParamSchema
}

type WorkerReference struct {
	Service     string
	Environment *string
}

type OutboundParamSchema struct {
	Name string
}

// Type returns the type of the binding.
func (b DispatchNamespaceBinding) Type() WorkerBindingType {
	return DispatchNamespaceBindingType
}

func (b DispatchNamespaceBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.Binding == "" {
		return nil, nil, fmt.Errorf(`binding name for binding "%s" cannot be empty`, bindingName)
	}
	if b.Namespace == "" {
		return nil, nil, fmt.Errorf(`namespace name for binding "%s" cannot be empty`, bindingName)
	}

	meta := workerBindingMeta{
		"type":      b.Type(),
		"name":      b.Binding,
		"namespace": b.Namespace,
	}

	if b.Outbound != nil {
		if b.Outbound.Worker.Service == "" {
			return nil, nil, fmt.Errorf(`outbound options for binding "%s" must have a service name`, bindingName)
		}

		var params []map[string]interface{}
		for _, param := range b.Outbound.Params {
			params = append(params, map[string]interface{}{
				"name": param.Name,
			})
		}

		meta["outbound"] = map[string]interface{}{
			"worker": map[string]interface{}{
				"service":     b.Outbound.Worker.Service,
				"environment": b.Outbound.Worker.Environment,
			},
			"params": params,
		}
	}

	return meta, nil, nil
}

// WorkerD1DatabaseBinding is a binding to a D1 instance.
type WorkerD1DatabaseBinding struct {
	DatabaseID string
}

// Type returns the type of the binding.
func (b WorkerD1DatabaseBinding) Type() WorkerBindingType {
	return WorkerD1DataseBindingType
}

func (b WorkerD1DatabaseBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	if b.DatabaseID == "" {
		return nil, nil, fmt.Errorf(`database ID for binding "%s" cannot be empty`, bindingName)
	}

	return workerBindingMeta{
		"name": bindingName,
		"type": b.Type(),
		"id":   b.DatabaseID,
	}, nil, nil
}

// UnsafeBinding is for experimental or deprecated bindings, and allows specifying any binding type or property.
type UnsafeBinding map[string]interface{}

// Type returns the type of the binding.
func (b UnsafeBinding) Type() WorkerBindingType {
	return ""
}

func (b UnsafeBinding) serialize(bindingName string) (workerBindingMeta, workerBindingBodyWriter, error) {
	b["name"] = bindingName
	return b, nil, nil
}

// Each binding that adds a part to the multipart form body will need
// a unique part name so we just generate a random 128bit hex string.
func getRandomPartName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes) //nolint:errcheck
	return hex.EncodeToString(randBytes)
}

// ListWorkerBindings returns all the bindings for a particular worker.
func (api *API) ListWorkerBindings(ctx context.Context, rc *ResourceContainer, params ListWorkerBindingsParams) (WorkerBindingListResponse, error) {
	if params.ScriptName == "" {
		return WorkerBindingListResponse{}, errors.New("script name is required")
	}

	if rc.Level != AccountRouteLevel {
		return WorkerBindingListResponse{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkerBindingListResponse{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings", rc.Identifier, params.ScriptName)
	if params.DispatchNamespace != nil && *params.DispatchNamespace != "" {
		uri = fmt.Sprintf("/accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", rc.Identifier, *params.DispatchNamespace, params.ScriptName)
	}

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
			return r, fmt.Errorf("binding missing name %v", jsonBinding)
		}
		bType, ok := jsonBinding["type"].(string)
		if !ok {
			return r, fmt.Errorf("binding missing type %v", jsonBinding)
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
		case WorkerQueueBindingType:
			queueName := jsonBinding["queue_name"].(string)
			bindingListItem.Binding = WorkerQueueBinding{
				Binding: name,
				Queue:   queueName,
			}
		case WorkerWebAssemblyBindingType:
			bindingListItem.Binding = WorkerWebAssemblyBinding{
				Module: &bindingContentReader{
					api:         api,
					ctx:         ctx,
					accountID:   rc.Identifier,
					params:      &params,
					bindingName: name,
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
		case WorkerAnalyticsEngineBindingType:
			dataset := jsonBinding["dataset"].(string)
			bindingListItem.Binding = WorkerAnalyticsEngineBinding{
				Dataset: dataset,
			}
		case WorkerD1DataseBindingType:
			database_id := jsonBinding["database_id"].(string)
			bindingListItem.Binding = WorkerD1DatabaseBinding{
				DatabaseID: database_id,
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
	api         *API
	accountID   string
	params      *ListWorkerBindingsParams
	ctx         context.Context
	bindingName string
	content     []byte
	position    int
}

func (b *bindingContentReader) Read(p []byte) (n int, err error) {
	// Lazily load the content when Read() is first called
	if b.content == nil {
		uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/bindings/%s/content", b.accountID, b.params.ScriptName, b.bindingName)
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
