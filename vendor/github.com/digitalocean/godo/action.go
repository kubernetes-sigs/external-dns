package godo

import (
	"context"
	"fmt"
	"net/http"
)

const (
	actionsBasePath = "v2/actions"

	// ActionInProgress is an in progress action status
	ActionInProgress = "in-progress"

	//ActionCompleted is a completed action status
	ActionCompleted = "completed"
)

<<<<<<< HEAD
// ActionsService handles communction with action related methods of the
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// DigitalOcean API: https://docs.digitalocean.com/reference/api/api-reference/#tag/Actions
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// DigitalOcean API: https://developers.digitalocean.com/documentation/v2#actions
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// DigitalOcean API: https://developers.digitalocean.com/documentation/v2#actions
=======
// DigitalOcean API: https://docs.digitalocean.com/reference/api/api-reference/#tag/Actions
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// DigitalOcean API: https://developers.digitalocean.com/documentation/v2#actions
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// ActionsService handles communction with action related methods of the
// DigitalOcean API: https://developers.digitalocean.com/documentation/v2#actions
=======
// ActionsService handles communication with action related methods of the
// DigitalOcean API: https://docs.digitalocean.com/reference/api/api-reference/#tag/Actions
>>>>>>> 6b7ce455e (update vendored files)
type ActionsService interface {
	List(context.Context, *ListOptions) ([]Action, *Response, error)
	Get(context.Context, int) (*Action, *Response, error)
}

// ActionsServiceOp handles communication with the image action related methods of the
// DigitalOcean API.
type ActionsServiceOp struct {
	client *Client
}

var _ ActionsService = &ActionsServiceOp{}

type actionsRoot struct {
	Actions []Action `json:"actions"`
	Links   *Links   `json:"links"`
	Meta    *Meta    `json:"meta"`
}

type actionRoot struct {
	Event *Action `json:"action"`
}

// Action represents a DigitalOcean Action
type Action struct {
	ID           int        `json:"id"`
	Status       string     `json:"status"`
	Type         string     `json:"type"`
	StartedAt    *Timestamp `json:"started_at"`
	CompletedAt  *Timestamp `json:"completed_at"`
	ResourceID   int        `json:"resource_id"`
	ResourceType string     `json:"resource_type"`
	Region       *Region    `json:"region,omitempty"`
	RegionSlug   string     `json:"region_slug,omitempty"`
}

// List all actions
func (s *ActionsServiceOp) List(ctx context.Context, opt *ListOptions) ([]Action, *Response, error) {
	path := actionsBasePath
	path, err := addOptions(path, opt)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(actionsRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	if l := root.Links; l != nil {
		resp.Links = l
	}
	if m := root.Meta; m != nil {
		resp.Meta = m
	}

	return root.Actions, resp, err
}

// Get an action by ID.
func (s *ActionsServiceOp) Get(ctx context.Context, id int) (*Action, *Response, error) {
	if id < 1 {
		return nil, nil, NewArgError("id", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/%d", actionsBasePath, id)
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(actionRoot)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Event, resp, err
}

func (a Action) String() string {
	return Stringify(a)
}
