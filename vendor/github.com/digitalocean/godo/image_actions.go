package godo

import (
	"context"
	"fmt"
	"net/http"
)

// ImageActionsService is an interface for interfacing with the image actions
// endpoints of the DigitalOcean API
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Image-Actions
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2#image-actions
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2#image-actions
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Image-Actions
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2#image-actions
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2#image-actions
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Image-Actions
>>>>>>> 6b7ce455e (update vendored files)
type ImageActionsService interface {
	Get(context.Context, int, int) (*Action, *Response, error)
	Transfer(context.Context, int, *ActionRequest) (*Action, *Response, error)
	Convert(context.Context, int) (*Action, *Response, error)
}

// ImageActionsServiceOp handles communication with the image action related methods of the
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2#image-actions
||||||| parent of 4d7e5ad26 (update vendored files)
// See: https://developers.digitalocean.com/documentation/v2#image-actions
=======
// See: https://docs.digitalocean.com/reference/api/api-reference/#tag/Image-Actions
>>>>>>> 4d7e5ad26 (update vendored files)
type ImageActionsService interface {
	Get(context.Context, int, int) (*Action, *Response, error)
	Transfer(context.Context, int, *ActionRequest) (*Action, *Response, error)
	Convert(context.Context, int) (*Action, *Response, error)
}

<<<<<<< HEAD
// ImageActionsServiceOp handles communition with the image action related methods of the
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// ImageActionsServiceOp handles communition with the image action related methods of the
=======
// ImageActionsServiceOp handles communication with the image action related methods of the
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// See: https://developers.digitalocean.com/documentation/v2#image-actions
type ImageActionsService interface {
	Get(context.Context, int, int) (*Action, *Response, error)
	Transfer(context.Context, int, *ActionRequest) (*Action, *Response, error)
	Convert(context.Context, int) (*Action, *Response, error)
}

// ImageActionsServiceOp handles communition with the image action related methods of the
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// DigitalOcean API.
type ImageActionsServiceOp struct {
	client *Client
}

var _ ImageActionsService = &ImageActionsServiceOp{}

// Transfer an image
func (i *ImageActionsServiceOp) Transfer(ctx context.Context, imageID int, transferRequest *ActionRequest) (*Action, *Response, error) {
	if imageID < 1 {
		return nil, nil, NewArgError("imageID", "cannot be less than 1")
	}

	if transferRequest == nil {
		return nil, nil, NewArgError("transferRequest", "cannot be nil")
	}

	path := fmt.Sprintf("v2/images/%d/actions", imageID)

	req, err := i.client.NewRequest(ctx, http.MethodPost, path, transferRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(actionRoot)
	resp, err := i.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Event, resp, err
}

// Convert an image to a snapshot
func (i *ImageActionsServiceOp) Convert(ctx context.Context, imageID int) (*Action, *Response, error) {
	if imageID < 1 {
		return nil, nil, NewArgError("imageID", "cannont be less than 1")
	}

	path := fmt.Sprintf("v2/images/%d/actions", imageID)

	convertRequest := &ActionRequest{
		"type": "convert",
	}

	req, err := i.client.NewRequest(ctx, http.MethodPost, path, convertRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(actionRoot)
	resp, err := i.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Event, resp, err
}

// Get an action for a particular image by id.
func (i *ImageActionsServiceOp) Get(ctx context.Context, imageID, actionID int) (*Action, *Response, error) {
	if imageID < 1 {
		return nil, nil, NewArgError("imageID", "cannot be less than 1")
	}

	if actionID < 1 {
		return nil, nil, NewArgError("actionID", "cannot be less than 1")
	}

	path := fmt.Sprintf("v2/images/%d/actions/%d", imageID, actionID)

	req, err := i.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(actionRoot)
	resp, err := i.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Event, resp, err
}
