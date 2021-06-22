package govultr

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

const scriptPath = "/v2/startup-scripts"

// StartupScriptService is the interface to interact with the startup script endpoints on the Vultr API
// Link : https://www.vultr.com/api/#tag/startup
type StartupScriptService interface {
	Create(ctx context.Context, req *StartupScriptReq) (*StartupScript, error)
	Get(ctx context.Context, scriptID string) (*StartupScript, error)
	Update(ctx context.Context, scriptID string, scriptReq *StartupScriptReq) error
	Delete(ctx context.Context, scriptID string) error
	List(ctx context.Context, options *ListOptions) ([]StartupScript, *Meta, error)
}

// StartupScriptServiceHandler handles interaction with the startup script methods for the Vultr API
type StartupScriptServiceHandler struct {
	client *Client
}

// StartupScript represents an startup script on Vultr
type StartupScript struct {
	ID           string `json:"id"`
	DateCreated  string `json:"date_created"`
	DateModified string `json:"date_modified"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Script       string `json:"script"`
}

// StartupScriptReq is the user struct for create and update calls
type StartupScriptReq struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Script string `json:"script"`
}

type startupScriptsBase struct {
	StartupScripts []StartupScript `json:"startup_scripts"`
	Meta           *Meta           `json:"meta"`
}

type startupScriptBase struct {
	StartupScript *StartupScript `json:"startup_script"`
}

var _ StartupScriptService = &StartupScriptServiceHandler{}

// Create a startup script
func (s *StartupScriptServiceHandler) Create(ctx context.Context, scriptReq *StartupScriptReq) (*StartupScript, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, scriptPath, scriptReq)
	if err != nil {
		return nil, err
	}

	script := new(startupScriptBase)
	if err = s.client.DoWithContext(ctx, req, script); err != nil {
		return nil, err
	}

	return script.StartupScript, nil
}

// Get a single startup script
func (s *StartupScriptServiceHandler) Get(ctx context.Context, scriptID string) (*StartupScript, error) {
	uri := fmt.Sprintf("%s/%s", scriptPath, scriptID)

	req, err := s.client.NewRequest(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	script := new(startupScriptBase)
	if err = s.client.DoWithContext(ctx, req, script); err != nil {
		return nil, err
	}

	return script.StartupScript, nil
}

// Update will update the given startup script. Empty strings will be ignored.
func (s *StartupScriptServiceHandler) Update(ctx context.Context, scriptID string, scriptReq *StartupScriptReq) error {
	uri := fmt.Sprintf("%s/%s", scriptPath, scriptID)

	req, err := s.client.NewRequest(ctx, http.MethodPatch, uri, scriptReq)
	if err != nil {
		return err
	}

	return s.client.DoWithContext(ctx, req, nil)
}

// Delete the specified startup script from your account.
func (s *StartupScriptServiceHandler) Delete(ctx context.Context, scriptID string) error {
	uri := fmt.Sprintf("%s/%s", scriptPath, scriptID)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return s.client.DoWithContext(ctx, req, nil)
}

// List all the startup scripts associated with your Vultr account
func (s *StartupScriptServiceHandler) List(ctx context.Context, options *ListOptions) ([]StartupScript, *Meta, error) {
	req, err := s.client.NewRequest(ctx, http.MethodGet, scriptPath, nil)
	if err != nil {
		return nil, nil, err
	}

	newValues, err := query.Values(options)
	if err != nil {
		return nil, nil, err
	}
	req.URL.RawQuery = newValues.Encode()

	scripts := new(startupScriptsBase)
	if err = s.client.DoWithContext(ctx, req, scripts); err != nil {
		return nil, nil, err
	}

	return scripts.StartupScripts, scripts.Meta, nil
}
