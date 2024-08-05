package cloudflare

import (
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
)

// WorkersPutSecretRequest provides parameters for creating and updating secrets.
type WorkersPutSecretRequest struct {
	Name string            `json:"name"`
	Text string            `json:"text"`
	Type WorkerBindingType `json:"type"`
}

// WorkersSecret contains the name and type of the secret.
type WorkersSecret struct {
	Name string `json:"name"`
	Type string `json:"secret_text"`
}

// WorkersPutSecretResponse is the response received when creating or updating a secret.
type WorkersPutSecretResponse struct {
	Response
	Result WorkersSecret `json:"result"`
}

// WorkersListSecretsResponse is the response received when listing secrets.
type WorkersListSecretsResponse struct {
	Response
	Result []WorkersSecret `json:"result"`
}

// SetWorkersSecret creates or updates a secret
// API reference: https://api.cloudflare.com/
func (api *API) SetWorkersSecret(ctx context.Context, script string, req *WorkersPutSecretRequest) (WorkersPutSecretResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets", api.AccountID, script)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, req)
	if err != nil {
<<<<<<< HEAD
<<<<<<< HEAD
		return WorkersPutSecretResponse{}, errors.Wrap(err, errMakeRequestError)
	}

	result := WorkersPutSecretResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// DeleteWorkersSecret deletes a secret
// API reference: https://api.cloudflare.com/
func (api *API) DeleteWorkersSecret(ctx context.Context, script, secretName string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets/%s", api.AccountID, script, secretName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, errors.Wrap(err, errMakeRequestError)
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersSecrets lists secrets for a given worker
// API reference: https://api.cloudflare.com/
func (api *API) ListWorkersSecrets(ctx context.Context, script string) (WorkersListSecretsResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets", api.AccountID, script)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersListSecretsResponse{}, errors.Wrap(err, errMakeRequestError)
||||||| parent of 6b7ce455e (update vendored files)
=======
		return WorkersPutSecretResponse{}, err
	}

	result := WorkersPutSecretResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// DeleteWorkersSecret deletes a secret
// API reference: https://api.cloudflare.com/
func (api *API) DeleteWorkersSecret(ctx context.Context, script, secretName string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets/%s", api.AccountID, script, secretName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// ListWorkersSecrets lists secrets for a given worker
// API reference: https://api.cloudflare.com/
func (api *API) ListWorkersSecrets(ctx context.Context, script string) (WorkersListSecretsResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets", api.AccountID, script)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersListSecretsResponse{}, err
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
		return WorkersPutSecretResponse{}, err
	}

	result := WorkersPutSecretResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// DeleteWorkersSecret deletes a secret
// API reference: https://api.cloudflare.com/
func (api *API) DeleteWorkersSecret(ctx context.Context, script, secretName string) (Response, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets/%s", api.AccountID, script, secretName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, errors.Wrap(err, errUnmarshalError)
	}

	return result, err
}

// ListWorkersSecrets lists secrets for a given worker
// API reference: https://api.cloudflare.com/
func (api *API) ListWorkersSecrets(ctx context.Context, script string) (WorkersListSecretsResponse, error) {
	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets", api.AccountID, script)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersListSecretsResponse{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// WorkersPutSecretRequest provides parameters for creating and updating secrets.
type WorkersPutSecretRequest struct {
	Name string            `json:"name"`
	Text string            `json:"text"`
	Type WorkerBindingType `json:"type"`
}

// WorkersSecret contains the name and type of the secret.
type WorkersSecret struct {
	Name string `json:"name"`
	Type string `json:"secret_text"`
}

// WorkersPutSecretResponse is the response received when creating or updating a secret.
type WorkersPutSecretResponse struct {
	Response
	Result WorkersSecret `json:"result"`
}

// WorkersListSecretsResponse is the response received when listing secrets.
type WorkersListSecretsResponse struct {
	Response
	Result []WorkersSecret `json:"result"`
}

type SetWorkersSecretParams struct {
	ScriptName string
	Secret     *WorkersPutSecretRequest
}

type DeleteWorkersSecretParams struct {
	ScriptName string
	SecretName string
}

type ListWorkersSecretsParams struct {
	ScriptName string
}

// SetWorkersSecret creates or updates a secret.
//
// API reference: https://api.cloudflare.com/
func (api *API) SetWorkersSecret(ctx context.Context, rc *ResourceContainer, params SetWorkersSecretParams) (WorkersPutSecretResponse, error) {
	if rc.Level != AccountRouteLevel {
		return WorkersPutSecretResponse{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkersPutSecretResponse{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets", rc.Identifier, params.ScriptName)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params.Secret)
	if err != nil {
		return WorkersPutSecretResponse{}, err
	}

	result := WorkersPutSecretResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// DeleteWorkersSecret deletes a secret.
//
// API reference: https://api.cloudflare.com/
func (api *API) DeleteWorkersSecret(ctx context.Context, rc *ResourceContainer, params DeleteWorkersSecretParams) (Response, error) {
	if rc.Level != AccountRouteLevel {
		return Response{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return Response{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets/%s", rc.Identifier, params.ScriptName, params.SecretName)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// ListWorkersSecrets lists secrets for a given worker
// API reference: https://api.cloudflare.com/
func (api *API) ListWorkersSecrets(ctx context.Context, rc *ResourceContainer, params ListWorkersSecretsParams) (WorkersListSecretsResponse, error) {
	if rc.Level != AccountRouteLevel {
		return WorkersListSecretsResponse{}, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return WorkersListSecretsResponse{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/scripts/%s/secrets", rc.Identifier, params.ScriptName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersListSecretsResponse{}, err
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}

	result := WorkersListSecretsResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}
